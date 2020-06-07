package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/nick92/appstar/api"
	"github.com/nick92/appstar/common"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	db := common.Init()
	defer db.Close()

	r := gin.Default()
	rapi := r.Group("api")
	api.Routes(rapi)

	r.Run(":" + port) // listen and serve on 0.0.0.0:port
}
