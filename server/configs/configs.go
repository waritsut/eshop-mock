package configs

import (
	_ "embed"

	"github.com/spf13/viper"
)
 
var config = new(Config)

type Config struct {
	Port       string `mapstructure:"PORT"`
	DbDialect  string `mapstructure:"DB_DIALECT"`
	DbHost     string `mapstructure:"DB_HOST"`
	DbName     string `mapstructure:"DB_NAME"`
	DbUser     string `mapstructure:"DB_USER"`
	DbPassword string `mapstructure:"DB_PASSWORD"`
	DbPort     string `mapstructure:"DB_PORT"`
	APiReadTimeOut  int    `mapstructure:"API_READ_TIME_OUT"`
	APiWriteTimeOut int    `mapstructure:"API_WRITE_TIME_OUT"`
	LogPath         string `mapstructure:"LOG_PATH"`
}

func SetupConfig() {
	viper.AddConfigPath("./configs/")
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
}

func Get() Config {
	return *config
}
