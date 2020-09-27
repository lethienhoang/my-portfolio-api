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
	API_HOST  = ""
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

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")

	DBDRIVER = os.Getenv("DB_DRIVER")
	DBURL = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	ISSUER = os.Getenv("ISSUER")
	SECRETKEY = []byte(os.Getenv("API_SECRET"))
	API_HOST = os.Getenv("API_HOST")
	PORT, err = strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		PORT = 9000
	}

	REDIS_DSN = os.Getenv("REDIS_DSN")
	PASSWRORD = os.Getenv("REDIS_DSN_PASSWRORD")
}
