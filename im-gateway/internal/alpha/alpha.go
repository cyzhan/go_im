package alpha

import (
	// this package is to make sure .env can be load before any other internal file
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func InitEnv() {
	godotenv.Load()
}
