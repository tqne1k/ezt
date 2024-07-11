package bootstrap

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Env struct {
	ConnectionString string `mapstructure:"CONNECTION_STRING"`
	DbHost           string `mapstructure:"DB_HOST"`
	DbPort           string `mapstructure:"DB_PORT"`
	DbUser           string `mapstructure:"DB_USER"`
	DbPasswd         string `mapstructure:"DB_PASSWD"`
	DbName           string `mapstructure:"DB_NAME"`
	ContextTimeout   int    `mapstructure:"CONTEXT_TIMEOUT"`
	ServerAddress    string `mapstructure:"SERVER_ADDRESS"`
	Netcidr          int    `mapstructure:"NETCIDR"`
	ConfigDir        string `mapstructure:"CONFIG_DIR"`
}

func NewEnv() *Env {
	env := Env{}
	// Read .env in the root directory
	viper.SetConfigFile(".env")

	// viper.SetConfigFile(".env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	return &env
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}
