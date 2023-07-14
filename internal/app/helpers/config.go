package helpers

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost      string `mapstructure:"DATABASE_HOST"`
	DBPort      string `mapstructure:"DATABASE_PORT"`
	DBUsername  string `mapstructure:"DATABASE_USERNAME"`
	DBPasssword string `mapstructure:"DATABASE_PASSWORD"`
	DBName      string `mapstructure:"DATABASE_NAME"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	err = viper.ReadInConfig()

	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
