package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	log.Println("Server is running on port 8002")
	server.Run(":8002")

}
