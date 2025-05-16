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
	ContextLoggerKey       = "logger"
	ContextUserKey         = "user"
	ContextRefreshTokenKey = "refresh_token_payload"
	ContentAPIRecordKey    = "api_key_record"
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

type RoleType string

const (
	RoleTypeUser  RoleType = "user"
	RoleTypeAdmin RoleType = "admin"
	RoleTypeSuper RoleType = "super"
)

type RoleLevel int

const (
	RoleLevelLow RoleLevel = iota
	RoleLevelMedium
	RoleLevelHigh
)

var RoleLevelMap = map[RoleType]RoleLevel{
	RoleTypeUser:  RoleLevelLow,
	RoleTypeAdmin: RoleLevelMedium,
	RoleTypeSuper: RoleLevelHigh,
}

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

const (
	AuthTypeBasic  = "basic"
	AuthTypeBearer = "bearer"
)
