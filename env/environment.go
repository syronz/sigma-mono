package env

// Environment directly fetch os envs with getting help from env
type Environment struct {
	Node      `json:"node"`
	Cloud     `json:"cloud"`
	Setting   `json:"setting"`
	Database  `json:"database"`
	Log       `json:"log"`
	MachineID string
}

// Node hold gin and tls configuration
type Node struct {
	Port     string `env:"SIGMA_NODE_PORT" json:"port"`
	ADDR     string `env:"SIGMA_NODE_ADDR" json:"addr"`
	TLSKey   string `env:"SIGMA_TLS_KEY" json:"tls_key"`
	TLSCert  string `env:"SIGMA_TLS_CERT" json:"tls_cert"`
	TimeZone string `env:"SIGMA_TIME_ZONE" json:"time_zone"`
}

// Cloud hold gin and tls configuration
type Cloud struct {
	Port               string `env:"SIGMA_CLOUD_PORT" json:"port"`
	ADDR               string `env:"SIGMA_CLOUD_ADDR" json:"addr"`
	HostURL            string `env:"SIGMA_CLOUD_HOST_URL,required" json:"host_url"`
	TLSKey             string `env:"SIGMA_CLOUD_TLS_KEY" json:"tls_key"`
	TLSCert            string `env:"SIGMA_CLOUD_TLS_CERT" json:"tls_cert"`
	TimeZone           string `env:"SIGMA_CLOUD_TIME_ZONE" json:"time_zone"`
	SuperAdminUsername string `env:"SIGMA_CLOUD_SUPER_ADMIN_USERNAME"`
	SuperAdminPassword string `env:"SIGMA_CLOUD_SUPER_ADMIN_PASSWORD"`
}

// Setting hold pass-keys and JWT, used for security
type Setting struct {
	PasswordSalt       string `env:"SIGMA_PASSWORD_SALT" json:"password_salt"`
	AutoMigrate        bool   `env:"SIGMA_AUTO_MIGRATE" json:"auto_migrate"`
	JWTSecretKey       string `env:"SIGMA_JWT_SECRET_KEY,required" json:"jwt_secret_key"`
	JWTExpiration      int    `env:"SIGMA_JWT_EXPIRATION,required" json:"jwt_expiration"`
	RecordRead         bool   `env:"SIGMA_RECORD_READ" json:"record_read"`
	RecordWrite        bool   `env:"SIGMA_RECORD_WRITE" json:"record_write"`
	TermsPath          string `env:"SIGMA_TERMS_PATH" json:"terms_path"`
	DefaultLanguage    string `env:"SIGMA_DEFAULT_LANGUAGE" json:"default_language"`
	TranslateInBackend bool   `env:"SIGMA_TRANSLATE_IN_BACKEND" json:"translate_in_backend"`
	ExcelMaxRows       uint64 `env:"SIGMA_EXCEL_MAX_ROWS" json:"excel_max_rows"`
}

// Database hold DB connections, in case we just have one database use same DSN for both
type Database struct {
	Data     Data     `json:"data"`
	Activity Activity `json:"activity"`
}

// Data is used inside the Database struct
type Data struct {
	DSN  string `env:"SIGMA_DATABASE_DATA_URL,required" json:"dsn"`
	Type string `env:"SIGMA_DATABASE_DATA_TYPE,required" json:"type"`
}

// Activity is used inside the Database struct
type Activity struct {
	DSN  string `env:"SIGMA_DATABASE_ACTIVITY_URL,required" json:"dsn"`
	Type string `env:"SIGMA_DATABASE_ACTIVITY_TYPE,required" json:"type"`
}

// Log configuration terms hold here
type Log struct {
	ServerLog ServerLog `json:"server_log"`
	APILog    APILog    `json:"api_log"`
}

// ServerLog is used inside the Log
type ServerLog struct {
	Format     string `env:"SIGMA_SERVER_LOG_FORMAT,required" json:"format"`
	Output     string `env:"SIGMA_SERVER_LOG_OUTPUT,required" json:"output"`
	Level      string `env:"SIGMA_SERVER_LOG_LEVEL,required" json:"level"`
	JSONIndent bool   `env:"SIGMA_SERVER_LOG_JSON_INDENT,required" json:"json_indent"`
}

// APILog is used inside the Log
type APILog struct {
	Format     string `env:"SIGMA_API_LOG_FORMAT,required" json:"format"`
	Output     string `env:"SIGMA_API_LOG_OUTPUT,required" json:"output"`
	Level      string `env:"SIGMA_API_LOG_LEVEL,required" json:"level"`
	JSONIndent bool   `env:"SIGMA_API_LOG_JSON_INDENT,required" json:"json_indent"`
}
