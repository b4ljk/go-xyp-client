package controllers

import (
	"net/http"

	"github.com/b4ljk/xyp-go/internal/controller/xyp"
	"github.com/b4ljk/xyp-go/internal/models"
	"github.com/gin-gonic/gin"
)

// Register Controller
func Alive(c *gin.Context) {
	defer func() {
		c.JSON(http.StatusOK, gin.H{"alive": true, "ready": true})
	}()
}

func Register(router *gin.RouterGroup) {
	// minio := integrations.MinioConnect()

	bc := models.Controller{
		// Minio: minio,
		// // Mail:      mail.CreateClient(),
		// Redis:     redis.CreateClient(),
		// Queue:     redis.InitAsynqWorker,
		// Asynq:     redis.InitAsynq(),
		// Inspector: redis.InitInspector(),
		// Response: &structs.Response{
		// 	StatusCode: http.StatusOK,
		// 	Body:       structs.ResponseBody{Message: "", Body: nil},
		// },
	}
	router.GET("/", Alive)
	xyp.Register(router, bc)

}
