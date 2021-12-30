package configs

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress        string        `mapstructure:"SERVER_ADDRESS"`
	DBDriver             string        `mapstructure:"DB_DRIVER"`
	DBUser               string        `mapstructure:"DB_USER"`
	DBPassword           string        `mapstructure:"DB_PASSWORD"`
	DBSource             string        `mapstructure:"DB_SOURCE"`
	TokenSymmetricKey    string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	DataSourceName       string
	SenderEmail          string `mapstructure:"SENDER_EMAIL"`
	SenderPassword       string `mapstructure:"SENDER_PASSWORD"`
	SMTPHost             string `mapstructure:"SMTP_HOST"`
	SMTPPort             string `mapstructure:"SMTP_PORT"`
	RedisURL             string `mapstructure:"REDIS_URL"`
}

func LoadConfig(path string) (Config, error) {
	config := Config{}
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	if err != nil {
		return config, err
	}
	config.DataSourceName = config.DBUser + ":" + config.DBPassword + config.DBSource
	return config, nil
}
