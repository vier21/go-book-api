package config

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoDBURL string
	SecretKey  []byte
	ServerPort string
	UserDBName string
	BookDBName string
}

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		return
	}
}

func GetConfig() *Config {
	return &Config{
		MongoDBURL: getDBURL(),
		SecretKey:  getSecretKey(),
		ServerPort: os.Getenv("SERVER_PORT"),
		BookDBName: getDBName("BOOK_DB"),
	}
}

func getSecretKey() []byte {
	return []byte(os.Getenv("SECRET_KEY"))
}

func getDBName(dburl string) string {
	rmv := strings.Replace(dburl, "//", "", 1)
	split := strings.Split(rmv, "/")

	if len(split) < 2 {
		return ""
	}
	dbname := split[1]
	return dbname
}

func getDBURL() string {
	dburl := os.Getenv("MONGODB_URI")
	if getDBName(dburl) == "" {
		return dburl
	}
	re := regexp.MustCompile(`\/[^/]+$`)
	result := re.ReplaceAllString(dburl, "")
	return result
}
