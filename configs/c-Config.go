package configs

import (
	"github.com/spf13/viper"
)

// MainConfigs main config struct
type MainConfigs struct {
	Server struct {
		MainPort string
	}

	DataBase struct {
		Addr     string
		User     string
		Password string
		DB       string
	}
}

// Cfg config
var Cfg MainConfigs

// InitConfigs Initializes the main programm settings
func InitConfigs(name string) error {
	initDefaults()
	viper.SetConfigName(name)
	viper.AddConfigPath("/configs/")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		return err
	}
	return nil
}

//initDefaults Initializes default values
func initDefaults() {
	viper.SetDefault("Server.MainPort", ":1323")
	viper.SetDefault("DataBase.Addr", "localhost:5432")
	viper.SetDefault("DataBase.User", "testu1")
	viper.SetDefault("DataBase.Password", "testpass1")
	viper.SetDefault("DataBase.DB", "vscale_db")
}
