package repository

import (
	"fmt"
	"strings"
	"time"

	"llmapi/src/internal/config"
	"llmapi/src/internal/constants"
	"llmapi/src/internal/utils/log"
	"llmapi/src/pkg/logger"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	gormlog "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type Options struct {
	DSN        string // Data Source Name for the database connection
	DBLogLevel int    // Log level for GORM logger
}

func CreateDB(opts *Options) (*gorm.DB, error) {
	// get db type from dsn prefix find first "://"

	dsn := opts.DSN

	idx := strings.Index(dsn, "://")
	if idx == -1 {
		return nil, fmt.Errorf("invalid DSN format: missing '://' scheme prefix in %s", dsn)
	}

	logWithDsn := log.Sys().With("dsn", dsn)

	dbType := strings.ToLower(dsn[:idx]) // Ensure lowercase comparison
	var dialector gorm.Dialector
	switch dbType {
	case constants.SqliteType:
		connStr := dsn[idx+3:]
		if connStr == "" {
			return nil, fmt.Errorf("invalid SQLite DSN: missing path after sqlite://")
		}
		dialector = sqlite.Open(connStr)
		logWithDsn.Info("Initializing SQLite database")
	case constants.PostgresType, constants.PostgresType + "ql":
		dialector = postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		})
		logWithDsn.Info("Initializing PostgreSQL database from DSN")
	case constants.MysqlType:
		connStr := dsn[idx+3:]
		if connStr == "" {
			return nil, fmt.Errorf("invalid MySQL DSN: missing connection string after mysql://")
		}
		dialector = mysql.New(mysql.Config{
			DSN:                       connStr,
			DefaultStringSize:         256, // string type default size
			DisableDatetimePrecision:  true,
			DontSupportRenameIndex:    true,
			DontSupportRenameColumn:   true,
			SkipInitializeWithVersion: false,
		})
		logWithDsn.Info("Initializing MySQL database from DSN")
	case constants.SqlServerType:
		dialector = sqlserver.Open(dsn)
		logWithDsn.Info("Initializing SQL Server database from DSN")
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}

	dbLog := gormlog.New(
		logger.NewGormLogger(log.WithType(log.GormType)),
		gormlog.Config{
			SlowThreshold:             time.Second,                       // Slow SQL threshold
			LogLevel:                  gormlog.LogLevel(opts.DBLogLevel), // Log level
			IgnoreRecordNotFoundError: true,                              // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,                              // Disable colorful log
		},
	)

	db, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "llmapi_",
			// SingularTable: true, // Use singular table names (e.g., "user" instead of "users")
		},
		Logger: dbLog,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect the database using DSN `%s`: %w", dsn, err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get underlying *sql.DB from GORM: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)                  // Max idle connections in the pool
	sqlDB.SetMaxOpenConns(100)                 // Max open connections to the database
	sqlDB.SetConnMaxLifetime(time.Hour)        // Max lifetime of a connection
	sqlDB.SetConnMaxIdleTime(time.Minute * 15) // Max idle time for a connection

	logWithDsn.Info("Pinging database to verify connection...")
	err = sqlDB.Ping()
	if err != nil {
		sqlDB.Close()
		return nil, fmt.Errorf("failed to ping %s database after connection: %w", dbType, err)
	}
	logWithDsn.Info("Successfully connected to database.", "database_type", dbType)
	return db, nil
}

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	log := log.Sys().With("dsn", cfg.DSN)
	log.Info("Initializing database connection...")

	db, err := CreateDB(&Options{
		DSN:        cfg.DSN,
		DBLogLevel: cfg.DBLogLevel,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create database: %w", err)
	}

	// Auto-migrate database schema if enabled
	if cfg.DBAutoMigrate {
		log.Info("Auto-migrating database schema...")
		if err := AutoMigrate(db); err != nil {
			return nil, fmt.Errorf("failed to auto-migrate database: %w", err)
		}
		log.Info("Database auto-migration completed")
	}
	return db, nil
}
