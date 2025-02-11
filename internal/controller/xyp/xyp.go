package xyp

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
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

func (co XYPController) Get(c *gin.Context) {

	REGNUM := viper.GetString("REGNUM")
	XYP_TOKEN := viper.GetString("XYP_TOKEN")
	XYP_KEY := viper.GetString("XYP_KEY")
	// time as string
	time := time.Now().Format("2006-01-02T15:04:05Z")

	soapRequest := fmt.Sprintf(`<?xml version="1.0" encoding="utf-8"?>
	<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cit="https://xyp.gov.mn/citizen">
	    <soapenv:Header/>
	    <soapenv:Body>
	        <cit:WS100101_getCitizenIDCardInfo>
	            <CitizenID>%s</CitizenID>
	        </cit:WS100101_getCitizenIDCardInfo>
	    </soapenv:Body>
	</soapenv:Envelope>`, REGNUM)

	xypSign := utils.XypSign{KeyPath: XYP_KEY}
	signed, err := xypSign.Generate(XYP_TOKEN, time)
	if err != nil {
		fmt.Println("Error signing:", err)
		return
	}

	// Define the SOAP endpoint
	url := "https://xyp.gov.mn/citizen-1.5.0/ws"

	// Create HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(soapRequest)))
	if err != nil {
		log.Fatal(err)
	}

	// Set headers
	req.Header.Set("Content-Type", "text/xml;charset=UTF-8")
	req.Header.Set("SOAPAction", "WS100101_getCitizenIDCardInfo")
	req.Header.Set("accessToken", signed.AccessToken)
	req.Header.Set("timeStamp", signed.Timestamp)
	req.Header.Set("signature", signed.Signature)

	// Send request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Print response
	fmt.Println("SOAP Response:", string(body))

	response.Success(c, 200, gin.H{
		"data": string(body),
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
