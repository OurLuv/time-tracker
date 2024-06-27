package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort   string `env:"SERVER_PORT" end-default:":8080"`
	User         string `env:"DB_USER"`
	Password     string `env:"DB_PASSWORD"`
	DatabaseName string `env:"DB_NAME"`
}

func MustLoad() *Config {
	var cfg Config
	Init()
	err := cleanenv.ReadEnv(&cfg)
	if err != nil {
		panic("Can't load config")
	}

	return &cfg
}

func Init() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}
