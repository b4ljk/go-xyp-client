package xyp

import (
	"github.com/b4ljk/xyp-go/internal/models"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.RouterGroup, bc models.Controller) {
	// common auth ashiglaj jwt uguh shit
	// AuthController{Controller: bc}.Register(router.Group("auth"))

	{
		XYPController{Controller: bc}.Register(router.Group("xyp"))
	}
}
