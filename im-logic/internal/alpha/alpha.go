package alpha

import (
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func InitEnv() {
	godotenv.Load()
}
