package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	PORT      = 0
	SECRETKEY []byte
	DBURL     = ""
	DBDRIVER  = ""
	ISSUER    = ""
	REDIS_DSN = ""
	PASSWRORD = ""
)

// Load the env
func Load() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	host := os.Getenv("")
	user := os.Getenv("")
	dbname := os.Getenv("")
	password := os.Getenv("")

	DBDRIVER = os.Getenv("")
	DBURL = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, user, dbname, password)

	ISSUER = os.Getenv("ISSUER")
	SECRETKEY = []byte(os.Getenv("API_SECRET"))
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 9000
	}

	REDIS_DSN = os.Getenv("REDIS_DSN")
	PASSWRORD = os.Getenv("REDIS_DSN_PASSWRORD")
}
