package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

var CONFIG Config

type Config struct {
	Env   string `yaml:"env"`
	Redis struct {
		URL string `yaml:"url"`
	} `yaml:"redis"`
	Flood struct {
		Max int64 `yaml:"max"`
		Dur int64 `yaml:"dur"`
	} `yaml:"flood"`
}

func InitConfig() error {
	if err := cleanenv.ReadConfig("./config/config.yaml", &CONFIG); err != nil {
		log.Fatal(err)
		return err
	}
	log.Println("Init config")
	return nil
}

func GetConf() Config {
	return CONFIG
}
