package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBURL string
}

func init()  {
	if err := godotenv.Load() ; err != nil {
		fmt.Println(err)
		return 
	}
}

func GetConfig() *Config  {
	return &Config{
		MongoDBURL: os.Getenv("MONGODB_URI"),
	}
}

