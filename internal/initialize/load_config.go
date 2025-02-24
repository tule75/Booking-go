package initialize

import (
	"ecommerce_go/global"
	"fmt"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./configs/")  // path to config folder
	viper.SetConfigName("development") // config file name
	viper.SetConfigType("yaml")        // config file type

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fail to read configuration file: %w \n", err))
	}

	fmt.Println("Server port: ", viper.GetInt("server.port"))

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to unmarshal config %w", err)
	}

	// for _, v := range config.Mysql {
	// 	fmt.Printf("Database Username: %v, Password: %v, Host: %v \n", v.User, v.Password, v.Host)
	// }
}
