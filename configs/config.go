package configs

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type conf struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg conf
	viper.AutomaticEnv()
	for _, k := range []string{
		"DB_DRIVER", "DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME",
		"WEB_SERVER_PORT", "GRPC_SERVER_PORT", "GRAPHQL_SERVER_PORT",
	} {
		_ = viper.BindEnv(k)
	}
	envFile := filepath.Join(path, ".env")
	if _, err := os.Stat(envFile); err == nil {
		viper.SetConfigFile(envFile)
		if err := viper.ReadInConfig(); err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, err
	}
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
