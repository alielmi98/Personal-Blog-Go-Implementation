package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	BlogTitle string `json:"blog_title"`
	AdminUser string `json:"admin_user"`
	AdminPass string `json:"admin_pass"`
	Port      string `json:"port"`
}

var AppConfig Config

func LoadConfig() {
	file, err := os.Open("config/config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&AppConfig)
	if err != nil {
		panic(err)
	}
}
