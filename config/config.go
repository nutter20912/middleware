package config

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	basePath, _ := filepath.Abs(".")
	viper.SetConfigFile(fmt.Sprintf("%s/config/app.yaml", basePath))

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
}
