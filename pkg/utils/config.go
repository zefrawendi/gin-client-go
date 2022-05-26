package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Server Server `yaml:"server"`
}

type Server struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName(".gin-client-go")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&config)
	return config, err
}
