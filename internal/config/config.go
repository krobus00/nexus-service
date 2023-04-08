package config

import (
	"errors"
	"time"

	"github.com/spf13/viper"
)

var (
	serviceName    = ""
	serviceVersion = ""
)

func ServiceName() string {
	return serviceName
}

func ServiceVersion() string {
	return serviceVersion
}

func Env() string {
	return viper.GetString("env")
}

func LogLevel() string {
	return viper.GetString("log_level")
}

func HTTPPort() string {
	return viper.GetString("ports.http")
}

func AuthGRPCHost() string {
	return viper.GetString("services.auth_grpc")
}

func StorageGRPCHost() string {
	return viper.GetString("services.storage_grpc")
}

func ProductGRPCHost() string {
	return viper.GetString("services.product_grpc")
}

func GracefulShutdownTimeOut() time.Duration {
	cfg := viper.GetString("graceful_shutdown_timeout")
	return parseDuration(cfg, DefaultGracefulShutdownTimeOut)
}

func LoadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return errors.New("config not found")
		}
		return err
	}
	return nil
}

func parseDuration(in string, defaultDuration time.Duration) time.Duration {
	dur, err := time.ParseDuration(in)
	if err != nil {
		return defaultDuration
	}
	return dur
}
