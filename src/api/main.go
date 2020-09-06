package main

import (
	"my-portfolio-api/config"
	"my-portfolio-api/utils/redis_server"
)

func init() {
	config.Load()
	redis_server.Load()
}

func Run() {

}

func Listen(port int) {

}

func main() {

}
