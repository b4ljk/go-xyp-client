package main

import (
	"log"

	controllers "github.com/b4ljk/xyp-go/internal/controller"
	"github.com/b4ljk/xyp-go/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadConfig()

	router := gin.Default()
	// router.Use(gzip.Gzip(gzip.DefaultCompression))

	api := router.Group("/api/v1")
	controllers.Register(api)

	log.Fatal(router.Run(":8080"))
}
