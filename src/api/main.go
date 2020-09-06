package main

import (
	"my-portfolio-api/config"
	"my-portfolio-api/utils/redisserver"
)

func init() {
	config.Load()
	redisserver.Load()
}

func Run() {

}

func Listen(port int) {

}

func main() {

}
