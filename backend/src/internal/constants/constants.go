package constants

const (
	AppName         = "LLMAPI"
	Version         = "0.0.1"
	Author          = "Pzyyll"
	License         = "MIT"
	DashboardPrefix = "/dashboard"
)

const (
	HttpRequestIDHeader = "X-Request-ID"
	HttpRequestIDKey    = "request_id"
)

const (
	ContextLoggerKey = "logger"
)

const (
	SqliteType    = "sqlite"
	PostgresType  = "postgres"
	MysqlType     = "mysql"
	SqlServerType = "sqlserver"
)

const (
	SecretFilePath   = ".jwt_secret.key"
	DesiredKeyLength = 64
	EnvPrefix        = "LA_"
)

const (
	RoleTypeUser  = "user"
	RoleTypeAdmin = "admin"
	RoleTypeSuper = "super"
)

const (
	PasswordMinLength = 8
	PasswordMaxLength = 32
	UsernameMinLength = 3
	UsernameMaxLength = 32
	UsernameRegex     = "^[a-zA-Z0-9_]+$"
)

const (
	JWT_SIGNED_HS256 = "HS256"
	JWT_SIGNED_HS384 = "HS384"
	JWT_SIGNED_HS512 = "HS512"
)
