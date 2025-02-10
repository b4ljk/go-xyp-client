package ndsh

import (
	"net/http"

	"github.com/b4ljk/xyp-go/internal/models"
	"github.com/b4ljk/xyp-go/pkg/response"
	"github.com/gin-gonic/gin"
)

type NdshController struct {
	models.Controller
}

type SSNCreateInput struct {
	RegisterNumber string `json:"register_number" binding:"required"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
}

func (co NdshController) Register(router *gin.RouterGroup) {
	router.GET("ssn-number/:id", co.GetById)
	router.POST("/", co.Create)
}

func (co NdshController) GetById(c *gin.Context) {
	params := c.Param("id")

	response.Success(c, 200, gin.H{
		"message": "success",
		"id":      params,
	})
}

func (co NdshController) Create(c *gin.Context) {

	var params SSNCreateInput

	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, http.StatusBadRequest, "Bad request you fucking dog")
		return
	}

	response.Success(c, http.StatusOK, params)
}
