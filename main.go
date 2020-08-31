package main

import (
	"challenge-go/routes"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("challenge - Javier Prozapas")

	r := routes.SetupRouter()
	r.Use(gin.Logger())
	if err := r.Run(":7007"); err != nil {
		log.Fatal("Unable to start:", err)
	}
}
