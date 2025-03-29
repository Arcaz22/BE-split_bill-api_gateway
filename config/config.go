package config

import (
	"api-gateway/pkg/constant"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port         string        `mapstructure:"port"`
		ReadTimeout  time.Duration `mapstructure:"read_timeout"`
		WriteTimeout time.Duration `mapstructure:"write_timeout"`
	} `mapstructure:"server"`

	Services struct {
		Auth struct {
			URL     string        `mapstructure:"url"`
			Timeout time.Duration `mapstructure:"timeout"`
		} `mapstructure:"auth"`
		Transaction struct {
			URL     string        `mapstructure:"url"`
			Timeout time.Duration `mapstructure:"timeout"`
		} `mapstructure:"transaction"`
		Notification struct {
			URL     string        `mapstructure:"url"`
			Timeout time.Duration `mapstructure:"timeout"`
		} `mapstructure:"notification"`
	} `mapstructure:"services"`

	JWT struct {
		Secret               string        `mapstructure:"secret"`
		AccessTokenDuration  time.Duration `mapstructure:"access_token_duration"`
		RefreshTokenDuration time.Duration `mapstructure:"refresh_token_duration"`
	} `mapstructure:"jwt"`
}

var config Config

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	viper.SetDefault("server.port", "8080")
	viper.SetDefault("server.read_timeout", "5s")
	viper.SetDefault("server.write_timeout", "10s")

	viper.SetDefault("jwt.access_token_duration", time.Minute*constant.JWTAccessTokenExpiry)
	viper.SetDefault("jwt.refresh_token_duration", time.Hour*24*constant.JWTRefreshTokenExpiry)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&config)
}

func GetConfig() Config {
	return config
}
