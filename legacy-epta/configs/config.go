package configs

import (
	"time"
)

type Config struct {
	HTTP HTTPConfig `yaml:"http" env:"HTTP"`

	Database DatabaseConfig `yaml:"database" env:"DATABASE"`

	Logging LoggingConfig `yaml:"logging" env:"LOGGING"`

	// функциональные флаги
	Features FeaturesConfig `yaml:"features" env:"FEATURES"`
}

type HTTPConfig struct {
	Addr         string        `yaml:"addr" env:"ADDR" env-default:":8080"`
	ReadTimeout  time.Duration `yaml:"read_timeout" env:"READ_TIMEOUT" env-default:"10s"`
	WriteTimeout time.Duration `yaml:"write_timeout" env:"WRITE_TIMEOUT" env-default:"10s"`
	IdleTimeout  time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"60s"`
	EnableCORS   bool          `yaml:"enable_cors" env:"ENABLE_CORS" env-default:"true"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	Port     int    `yaml:"port" env:"PORT" env-default:"5432"`
	User     string `yaml:"user" env:"USER" env-default:"postgres"`
	Password string `yaml:"password" env:"PASSWORD" env-default:"password"`
	Name     string `yaml:"name" env:"NAME" env-default:"university"`

	MaxOpenConns    int           `yaml:"max_open_conns" env:"MAX_OPEN_CONNS" env-default:"25"`
	MaxIdleConns    int           `yaml:"max_idle_conns" env:"MAX_IDLE_CONNS" env-default:"25"`
	ConnMaxLifetime time.Duration `yaml:"conn_max_lifetime" env:"CONN_MAX_LIFETIME" env-default:"5m"`
}

type LoggingConfig struct {
	Level    string `yaml:"level" env:"LEVEL" env-default:"info"`
	Format   string `yaml:"format" env:"FORMAT" env-default:"json"`
	FilePath string `yaml:"file_path" env:"FILE_PATH" env-default:""`
}

type FeaturesConfig struct {
	EnableRegistration      bool `yaml:"enable_registration" env:"ENABLE_REGISTRATION" env-default:"true"`
	EnableEmailVerification bool `yaml:"enable_email_verification" env:"ENABLE_EMAIL_VERIFICATION" env-default:"false"`
	EnableCaching           bool `yaml:"enable_caching" env:"ENABLE_CACHING" env-default:"true"`
	MaintenanceMode         bool `yaml:"maintenance_mode" env:"MAINTENANCE_MODE" env-default:"false"`
}
