package xyp

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/b4ljk/xyp-go/internal/models"
	"github.com/b4ljk/xyp-go/pkg/response"
	"github.com/b4ljk/xyp-go/utils"
	"github.com/b4ljk/xyp-go/utils/constants"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type XYPController struct {
	models.Controller
}

type AnotherResponse struct {
	XMLName xml.Name `xml:"another_response"`
	Code    int      `xml:"code"`
	Details string   `xml:"details"`
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
	time := fmt.Sprintf("%d", time.Now().Unix())

	soapBody := fmt.Sprintf(constants.XYP_PASSPORT_SOAP_BODY, REGNUM, REGNUM)

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

	req, err := http.NewRequest("POST", constants.XYP_PASSPORT_URL, bytes.NewBuffer([]byte(soapBody)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		response.Error(c, 500, err.Error())
		return
	}

	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("accessToken", signedData.AccessToken)
	req.Header.Set("timeStamp", signedData.Timestamp)
	req.Header.Set("signature", signedData.Signature)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Error on request:", err)
		response.Error(c, 500, err.Error())
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error on read response body:", err)
		response.Error(c, 500, err.Error())
		return
	}

	var another AnotherResponse
	jsonData, err := utils.XMLToJSON(body, &another)

	response.Success(c, 200, gin.H{
		"data": jsonData,
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
