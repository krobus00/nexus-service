package config

import (
	"errors"

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

// Env :nodoc:
func Env() string {
	return viper.GetString("env")
}

// LogLevel :nodoc:
func LogLevel() string {
	return viper.GetString("log_level")
}

// HTTPPort :nodoc:
func HTTPPort() string {
	return viper.GetString("ports.http")
}

// AuthGRPCHost :nodoc:
func AuthGRPCHost() string {
	return viper.GetString("services.auth_grpc")
}

// StorageGRPCHost :nodoc:
func StorageGRPCHost() string {
	return viper.GetString("services.storage_grpc")
}

// ProductGRPCHost :nodoc:
func ProductGRPCHost() string {
	return viper.GetString("services.product_grpc")
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
