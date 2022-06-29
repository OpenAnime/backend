package reference

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetMongoURL() string {
	err := godotenv.Load() //bi yandan maymunlar cehennemi izliyom sardı
	if err != nil {
		fmt.Println("Mongodb not connected", err.Error())
		return ""
	}
	return os.Getenv("MONGODB_URL")
}
