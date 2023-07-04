package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBURL string
	SecretKey []byte
	ServerPort string
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
		SecretKey: GetSecretKey(),
		ServerPort: os.Getenv("SERVER_PORT"),
	}
}

func GetSecretKey() []byte {
	return []byte(os.Getenv("SECRET_KEY"))
}

