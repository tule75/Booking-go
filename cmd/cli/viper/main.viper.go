package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type ConfigStruct struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		DbName   string `mapstructure:"dbname"`
	} `mapstructure:"databases`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")  // path to config folder
	viper.SetConfigName("development") // config file name
	viper.SetConfigType("yaml")        // config file type

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fail to read configuration file: %w \n", err))
	}

	fmt.Println("Server port: ", viper.GetInt("server.port"))

	var config ConfigStruct

	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to unmarshal config %w", err)
	}

	fmt.Printf("Config Port: %v \n Database: %v \n", config.Server.Port, config.Databases)

	for _, v := range config.Databases {
		fmt.Printf("Database Username: %v, Password: %v, Host: %v \n", v.User, v.Password, v.Host)
	}
}
