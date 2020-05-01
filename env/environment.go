package env

// Environment directly fetch os envs with getting help from env
type Environment struct {
	RadiusAuth `json:"radius_auth"`
	RestAPI    `json:"rest_api"`
	Setting    `json:"setting"`
	Database   `json:"database"`
	Log        `json:"log"`
	MachineID  string
}

// RadiusAuth hold gin and tls configuration
type RadiusAuth struct {
	Port string `env:"RADIUS_BILLING_RADIUS_AUTH_PORT" json:"port"`
	ADDR string `env:"RADIUS_BILLING_RADIUS_AUTH_ADDR" json:"addr"`
}

// RestAPI hold gin and tls configuration
type RestAPI struct {
	Port               string `env:"RADIUS_BILLING_REST_API_PORT" json:"port"`
	ADDR               string `env:"RADIUS_BILLING_REST_API_ADDR" json:"addr"`
	TLSKey             string `env:"RADIUS_BILLING_REST_API_TLS_KEY" json:"tls_key"`
	TLSCert            string `env:"RADIUS_BILLING_REST_API_TLS_CERT" json:"tls_cert"`
	TimeZone           string `env:"RADIUS_BILLING_REST_API_TIME_ZONE" json:"time_zone"`
	SuperAdminUsername string `env:"RADIUS_BILLING_REST_API_SUPER_ADMIN_USERNAME"`
	SuperAdminPassword string `env:"RADIUS_BILLING_REST_API_SUPER_ADMIN_PASSWORD"`
}

// Setting hold pass-keys and JWT, used for security
type Setting struct {
	PasswordSalt       string `env:"RADIUS_BILLING_PASSWORD_SALT" json:"password_salt"`
	AutoMigrate        bool   `env:"RADIUS_BILLING_AUTO_MIGRATE" json:"auto_migrate"`
	JWTSecretKey       string `env:"RADIUS_BILLING_JWT_SECRET_KEY,required" json:"jwt_secret_key"`
	JWTExpiration      int    `env:"RADIUS_BILLING_JWT_EXPIRATION,required" json:"jwt_expiration"`
	RecordRead         bool   `env:"RADIUS_BILLING_RECORD_READ" json:"record_read"`
	RecordWrite        bool   `env:"RADIUS_BILLING_RECORD_WRITE" json:"record_write"`
	TermsPath          string `env:"RADIUS_BILLING_TERMS_PATH" json:"terms_path"`
	DefaultLanguage    string `env:"RADIUS_BILLING_DEFAULT_LANGUAGE" json:"default_language"`
	TranslateInBackend bool   `env:"RADIUS_BILLING_TRANSLATE_IN_BACKEND" json:"translate_in_backend"`
	ExcelMaxRows       uint64 `env:"RADIUS_BILLING_EXCEL_MAX_ROWS" json:"excel_max_rows"`
}

// Database hold DB connections, in case we just have one database use same DSN for both
type Database struct {
	Data     Data     `json:"data"`
	Activity Activity `json:"activity"`
}

// Data is used inside the Database struct
type Data struct {
	DSN  string `env:"RADIUS_BILLING_DATABASE_DATA_URL,required" json:"dsn"`
	Type string `env:"RADIUS_BILLING_DATABASE_DATA_TYPE,required" json:"type"`
}

// Activity is used inside the Database struct
type Activity struct {
	DSN  string `env:"RADIUS_BILLING_DATABASE_ACTIVITY_URL,required" json:"dsn"`
	Type string `env:"RADIUS_BILLING_DATABASE_ACTIVITY_TYPE,required" json:"type"`
}

// Log configuration terms hold here
type Log struct {
	ServerLog ServerLog `json:"server_log"`
	APILog    APILog    `json:"api_log"`
}

// ServerLog is used inside the Log
type ServerLog struct {
	Format     string `env:"RADIUS_BILLING_SERVER_LOG_FORMAT,required" json:"format"`
	Output     string `env:"RADIUS_BILLING_SERVER_LOG_OUTPUT,required" json:"output"`
	Level      string `env:"RADIUS_BILLING_SERVER_LOG_LEVEL,required" json:"level"`
	JSONIndent bool   `env:"RADIUS_BILLING_SERVER_LOG_JSON_INDENT,required" json:"json_indent"`
}

// APILog is used inside the Log
type APILog struct {
	Format     string `env:"RADIUS_BILLING_API_LOG_FORMAT,required" json:"format"`
	Output     string `env:"RADIUS_BILLING_API_LOG_OUTPUT,required" json:"output"`
	Level      string `env:"RADIUS_BILLING_API_LOG_LEVEL,required" json:"level"`
	JSONIndent bool   `env:"RADIUS_BILLING_API_LOG_JSON_INDENT,required" json:"json_indent"`
}
