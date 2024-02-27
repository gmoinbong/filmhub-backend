package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Eror loading .env")
	}

}
func GoPort() (port string) {
	Port := os.Getenv("PORT")
	return Port
}
