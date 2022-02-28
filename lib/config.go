package lib

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	DbHost  string
	DbPort  int
	DbUser  string
	DbPass  string
	DbName  string
	WebPort int
}

func (config *Config) Initialize(requireLoadingFile bool) {
	viper.SetDefault("db_host", "localhost")
	viper.SetDefault("db_port", 12790)
	viper.SetDefault("db_user", "superuser")
	viper.SetDefault("db_pass", "fake-password")
	viper.SetDefault("db_name", "mastermind")
	viper.SetDefault("web_port", 18080)

	viper.SetConfigName("config.yaml")    // name of config file (without extension)
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	confFileError := viper.ReadInConfig() // Find and read the config file

	if confFileError != nil && requireLoadingFile { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %w \n", confFileError))
	}

	config.DbHost = viper.Get("db_host").(string)
	config.DbUser = viper.Get("db_user").(string)
	config.DbPass = viper.Get("db_pass").(string)
	config.DbName = viper.Get("db_name").(string)
	config.DbPort = viper.Get("db_port").(int)

	config.WebPort = viper.Get("web_port").(int)
}
