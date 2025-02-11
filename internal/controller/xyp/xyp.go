package xyp

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/b4ljk/xyp-go/internal/models"
	"github.com/b4ljk/xyp-go/myservice"
	"github.com/b4ljk/xyp-go/pkg/response"
	"github.com/b4ljk/xyp-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/hooklift/gowsdl/soap"
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

func (co XYPController) Get(c *gin.Context) {

	REGNUM := viper.GetString("REGNUM")
	XYP_TOKEN := viper.GetString("XYP_TOKEN")
	XYP_KEY := viper.GetString("XYP_KEY")
	time := time.Now().Format("2006-01-02T15:04:05")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	signature := utils.XypSign{
		KeyPath: XYP_KEY,
	}

	signedData, err := signature.Generate(XYP_TOKEN, time)

	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	newClient := soap.NewClient("https://xyp.gov.mn/citizen-1.5.0/ws?WSDL", soap.WithHTTPHeaders(map[string]string{
		"accessToken": signedData.AccessToken,
		"signature":   signedData.Signature,
		"timeStamp":   signedData.Timestamp,
	}), soap.WithHTTPClient(client))

	citizenService := myservice.NewCitizenService(newClient)

	_test := &myservice.WS100101_getCitizenIDCardInfo{
		Request: &myservice.CitizenRequestData{
			Regnum:  REGNUM,
			CivilId: "",
			ServiceRequest: &myservice.ServiceRequest{
				Auth: &myservice.AuthorizationData{
					Citizen: &myservice.AuthorizationEntity{
						CivilId:         "",
						Regnum:          REGNUM,
						AppAuthToken:    "",
						AuthAppName:     "",
						CertFingerprint: "",
						Signature:       "",
					},
					Operator: nil,
				},
			},
		},
	}

	_response, err := citizenService.WS100101_getCitizenIDCardInfo(_test)
	if err != nil {
		response.Error(c, 500, err.Error())
		return
	}

	response.Success(c, 200, gin.H{
		"data": _response,
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
