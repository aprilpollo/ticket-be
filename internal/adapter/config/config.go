package core

import (
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
)

// Env holds the application configuration loaded from environment variables.
var Env *Config

// Config holds the application configuration loaded from environment variables.
type Config struct {
	App
	JWT
	GoogleAuth
	RabbitMQ
	Redis
	Postgre
}

// App holds the application configuration.
type App struct {
	AppName                  string `env:"APP_NAME,default=MyApp"`
	AppVersion               string `env:"APP_VERSION"`
	ApiPort                  string `env:"API_PORT,default=8760"`
	ShutdownTimeout          uint   `env:"API_SHUTDOWN_TIMEOUT_SECONDS,default=30"`
	AllowedCredentialOrigins string `env:"ALLOWED_CREDENTIAL_ORIGINS"`

	LogLevel    string `env:"LOG_LEVEL,default=info"`
	Development bool   `env:"DEVELOPMENT,default=false"`
}

// JWT holds the configuration for JSON Web Token.
type JWT struct {
	SecretKey          string `env:"JWT_SECRET_KEY,required"`
	JwtExpireDaysCount int    `env:"JWT_EXPIRE_DAYS_COUNT"`
	Issuer             string `env:"JWT_ISSUER,default=MyApp"`
	Subject            string `env:"JWT_SUBJECT,default=User"`
	SigningMethod      string `env:"JWT_SIGNING_METHOD,default=HS256"`
}

// GoogleAuth holds the configuration for Google OAuth authentication.
type GoogleAuth struct {
	GoogleClientID     string `env:"GOOGLE_CLIENT_ID"`
	GoogleClientSecret string `env:"GOOGLE_CLIENT_SECRET"`
	GoogleRedirectURI  string `env:"GOOGLE_REDIRECT_URI"`
	GoogleTokenURL     string `env:"GOOGLE_TOKEN_URL"`
	GoogleGrantType    string `env:"GOOGLE_GRANT_TYPE"`
}

// Postgre holds the configuration for PostgreSQL database connection.
type Postgre struct {
	URI             string `env:"POSTGRE_URI,default="`
	MaxIdleConns    int    `env:"POSTGRE_MAX_IDLE_CONNS,default=10"`
	MaxOpenConns    int    `env:"POSTGRE_MAX_OPEN_CONNS,default=100"`
	ConnMaxLifetime int    `env:"POSTGRE_CONN_MAX_LIFETIME,default=0"`
}

// RabbitMQ holds the configuration for RabbitMQ messaging service.
type RabbitMQ struct {
	Uri             string `env:"RABBITMQ_URI,required"`
	Exchange        string `env:"RABBITMQ_EXCHANGE,default=events"`
	QueueType       string `env:"RABBITMQ_QUEUE_TYPE,default=topic"`
	QueuePrefix     string `env:"RABBITMQ_QUEUE_PREFIX,default=Ngorder API"`
	QueueRetryCount int    `env:"RABBITMQ_RETRY_COUNT,default=3"`
}

// Redis holds the configuration for Redis caching service.
type Redis struct {
	Host           string `env:"REDIS_HOST,required"`
	Password       string `env:"REDIS_PASSWORD"`
	ReadTimeoutMs  int16  `env:"REDIS_READ_TIMEOUT,required"`
	WriteTimeoutMs int16  `env:"REDIS_WRITE_TIMEOUT,required"`
}

// LoadConfig loads the application configuration from environment variables and .env file.
func LoadConfig() {
	// Try to load .env file, but don't fail if it doesn't exist
	// This is useful for development, but not required in production
	if err := godotenv.Load(); err != nil {
		// Only log as debug/info, not as error since it's optional
		// log.Println("Info: No .env file found, using environment variables only")
	}
	var cfg Config
	if err := envdecode.Decode(&cfg); err != nil {
		panic(err)
	}
	Env = &cfg
}
