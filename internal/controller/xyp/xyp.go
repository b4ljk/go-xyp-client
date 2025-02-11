package xyp

import (
	"fmt"
	"net/http"
	"time"

	"github.com/b4ljk/xyp-go/internal/models"
	"github.com/b4ljk/xyp-go/pkg/response"
	"github.com/b4ljk/xyp-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type XYPController struct {
	models.Controller
}

type XYPCreateInput struct {
	RegisterNumber string `json:"register_number" binding:"required"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
}

func (co XYPController) Register(router *gin.RouterGroup) {
	router.GET("ssn-number/:id", co.GetById)
	router.GET("/", co.Get)
	router.POST("/", co.Create)
}

func (co XYPController) GetById(c *gin.Context) {
	params := c.Param("id")

	response.Success(c, 200, gin.H{
		"message": "success",
		"id":      params,
	})
}

type RequestArgs struct {
	Request struct {
		Regnum string `json:"regnum"`
		Auth   struct {
			Citizen struct {
				CivilID     string `json:"civilId"`
				Regnum      string `json:"regnum"`
				Fingerprint string `json:"fingerprint"`
				AuthType    int    `json:"authType"`
			} `json:"citizen"`
			Operator struct {
				Regnum      string `json:"regnum"`
				Fingerprint string `json:"fingerprint"`
			} `json:"operator"`
		} `json:"auth"`
	} `json:"request"`
}

func (co XYPController) Get(c *gin.Context) {

	REGNUM := viper.GetString("REGNUM")
	XYP_TOKEN := viper.GetString("XYP_TOKEN")
	XYP_KEY := viper.GetString("XYP_KEY")
	// time as string
	time := time.Now().Format("2006-01-02T15:04:05Z")

	xypSign := utils.XypSign{KeyPath: XYP_KEY}

	signed, err := xypSign.Generate(XYP_TOKEN, time)
	if err != nil {
		fmt.Println("Error signing:", err)
		return
	}

	url := "https://xyp.gov.mn/citizen-1.5.0/ws?WSDL"

	// client := utils.NewSOAPClient(url, httpClient)
	client := utils.NewXypClient(url)
	response, err := client.GetCitizenIDCardInfo(REGNUM, signed)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println(response)

	response.Success(c, 200, gin.H{
		"data": response.Result,
	})

}

func (co XYPController) Create(c *gin.Context) {

	var params XYPCreateInput

	if err := c.ShouldBindJSON(&params); err != nil {
		response.Error(c, http.StatusBadRequest, "Bad request you fucking dog")
		return
	}

	response.Success(c, http.StatusOK, params)
}
