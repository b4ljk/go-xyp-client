package ndsh

import (
	"github.com/b4ljk/xyp-go/internal/models"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup, bc models.Controller) {
	// common auth ashiglaj jwt uguh shit
	// AuthController{Controller: bc}.Register(router.Group("auth"))

	{
		NdshController{Controller: bc}.Register(router.Group("ndsh"))
	}
}
