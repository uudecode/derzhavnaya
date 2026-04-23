package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(Load),
)

type Config struct {
	App          AppConfig         `mapstructure:"app"`
	Database     DatabaseConfig    `mapstructure:"database"`
	S3           S3Config          `mapstructure:"s3"`
	InitialAdmin AdminConfig       `mapstructure:"initial_admin"`
	Translation  TranslationConfig `mapstructure:"translation"`
}

type AppConfig struct {
	Port           int      `mapstructure:"port"`
	LogLevel       string   `mapstructure:"log_level"`
	DefaultLang    string   `mapstructure:"default_lang"`
	SupportedLangs []string `mapstructure:"supported_langs"`
}

type DatabaseConfig struct {
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	Name            string        `mapstructure:"name"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	SSLMode         string        `mapstructure:"sslmode"`
	MaxOpenConns    int           `mapstructure:"max_open_conns"`
	MaxIdleConns    int           `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
}

func (db DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		db.Host, db.Port, db.User, db.Password, db.Name, db.SSLMode)
}

func (db DatabaseConfig) String() string {
	return fmt.Sprintf("host=%s, port=%d, name=%s, user=%s, sslmode=%s", db.Host, db.Port, db.Name, db.User, db.SSLMode)
}

type S3Config struct {
	Region          string `mapstructure:"region"`
	Endpoint        string `mapstructure:"endpoint"`
	Bucket          string `mapstructure:"bucket"`
	PublicBaseURL   string `mapstructure:"public_base_url"`
	AccessKeyID     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
	DebugS3         bool   `mapstructure:"debug_s3"`
}

func (s S3Config) String() string {
	return fmt.Sprintf("region=%s, endpoint=%s, bucket=%s, public_base_url=%s", s.Region, s.Endpoint, s.Bucket, s.PublicBaseURL)
}

type AdminConfig struct {
	Email    string `mapstructure:"email"`
	Password string `mapstructure:"password"`
}

func (a AdminConfig) String() string {
	return fmt.Sprintf("email=****, password=****")
}

type TranslationConfig struct {
	LiteLLMUrl string            `mapstructure:"litellm_url"`
	LiteLLMKey string            `mapstructure:"litellm_key"`
	ModelName  string            `mapstructure:"model_name"`
	Prompts    map[string]string `mapstructure:"prompts"`
	TimeoutSec int               `mapstructure:"timeout_sec"`
}

func (t TranslationConfig) String() string {
	return fmt.Sprintf("url=%s, key=****", t.LiteLLMUrl)
}

func Load() (*Config, error) {

	v := viper.New()
	v.SetDefault("app.log_level", "info")
	v.SetDefault("app.port", 8080)
	v.SetConfigName("config")
	v.SetConfigType("toml")
	v.AddConfigPath(".")

	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	bindings := map[string]string{
		"app.log_level":              "LOG_LEVEL",
		"database.host":              "DB_HOST",
		"database.port":              "DB_PORT",
		"database.name":              "DB_NAME",
		"database.user":              "DB_USER",
		"database.password":          "DB_PASSWORD",
		"database.sslmode":           "DB_SSLMODE",
		"database.max_open_conns":    "DB_MAX_OPEN_CONNS",
		"database.max_idle_conns":    "DB_MAX_IDLE_CONNS",
		"database.conn_max_lifetime": "DB_CONN_MAX_LIFETIME",
		"s3.region":                  "AWS_REGION",
		"s3.endpoint":                "AWS_ENDPOINT_URL",
		"s3.bucket":                  "MEDIA_BUCKET",
		"s3.public_base_url":         "MEDIA_PUBLIC_BASE_URL",
		"s3.access_key_id":           "AWS_ACCESS_KEY_ID",
		"s3.secret_access_key":       "AWS_SECRET_ACCESS_KEY",
		"s3.debug_s3":                "DEBUG_S3",
		"initial_admin.email":        "INITIAL_ADMIN_EMAIL",
		"initial_admin.password":     "INITIAL_ADMIN_PASSWORD",
		"translation.litellm_url":    "TRANSLATION_LITELLM_URL",
		"translation.litellm_key":    "TRANSLATION_LITELLM_KEY",
		"translation.model_name":     "TRANSLATION_MODEL_NAME",
		"translation.timeout_sec":    "TRANSLATION_TIMEOUT_SEC",
	}

	for key, env := range bindings {
		if err := v.BindEnv(key, env); err != nil {
			return nil, fmt.Errorf("failed to bind env var %s: %w", env, err)
		}
	}

	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Config file not found (%s), proceeding with environment variables\n", err)
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
