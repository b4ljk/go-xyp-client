package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SOAPEnvelope represents the SOAP envelope structure
type SOAPEnvelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  SOAPHeader
	Body    SOAPBody
}

// SOAPHeader represents the SOAP header
type SOAPHeader struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Content interface{}
}

// SOAPBody represents the SOAP body
type SOAPBody struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Content interface{}
}

// SOAPClient represents a SOAP client
type SOAPClient struct {
	url     string
	tls     bool
	headers map[string]string
	// httpClient *http.Client
	httpClient *http.Client
}

// NewSOAPClient creates a new SOAP client
func NewSOAPClient(url string, httpClient *http.Client) *SOAPClient {
	return &SOAPClient{
		url:        url,
		headers:    make(map[string]string),
		httpClient: httpClient,
	}
}

// AddHeader adds a header to the SOAP request
func (s *SOAPClient) AddHeader(key, value string) {
	s.headers[key] = value
}

// Call makes a SOAP request
func (s *SOAPClient) Call(action string, request, response interface{}) error {
	envelope := SOAPEnvelope{
		Body: SOAPBody{
			Content: request,
		},
	}

	// Marshal the request to XML
	requestXML, err := xml.MarshalIndent(envelope, "", "    ")
	if err != nil {
		return err
	}

	// if httpclient exists, use it, otherwise create a new one
	httpClient := s.httpClient
	if httpClient == nil {
		httpClient = &http.Client{}
	}

	req, err := http.NewRequest("POST", s.url, bytes.NewBuffer(requestXML))
	if err != nil {
		return err
	}

	// Set headers
	req.Header.Set("Content-Type", "text/xml; charset=utf-8")
	req.Header.Set("SOAPAction", action)
	for key, value := range s.headers {
		req.Header.Set(key, value)
	}

	// Make the request
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Read response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// Unmarshal response
	respEnvelope := SOAPEnvelope{
		Body: SOAPBody{
			Content: response,
		},
	}
	return xml.Unmarshal(bodyBytes, &respEnvelope)
}

// Request structures
type CitizenAuth struct {
	CivilID     string `xml:"civilId"`
	Regnum      string `xml:"regnum"`
	Fingerprint string `xml:"fingerprint"`
	AuthType    int    `xml:"authType"`
}

type OperatorAuth struct {
	Regnum      string `xml:"regnum"`
	Fingerprint string `xml:"fingerprint"`
}

type Auth struct {
	Citizen  CitizenAuth  `xml:"citizen"`
	Operator OperatorAuth `xml:"operator"`
}

type Request struct {
	XMLName xml.Name `xml:"request"`
	Regnum  string   `xml:"regnum"`
	Auth    Auth     `xml:"auth"`
}

type WS100101Request struct {
	Request Request `xml:"request"`
}

// Response structure - adjust fields based on your actual response
type WS100101Response struct {
	XMLName xml.Name    `xml:"WS100101_getCitizenIDCardInfoResponse"`
	Result  interface{} `xml:"return"`
}

func (w *WS100101Response) Success(c *gin.Context, i int, h gin.H) {
	panic("unimplemented")
}

// SignData structure

// XypClient structure
type XypClient struct {
	soapClient *SOAPClient
}

// NewXypClient creates a new XYP client
func NewXypClient(url string) *XypClient {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
	return &XypClient{
		soapClient: NewSOAPClient(url, httpClient),
	}
}

// GetCitizenIDCardInfo implements the WS100101_getCitizenIDCardInfo call
func (c *XypClient) GetCitizenIDCardInfo(registration string, signData SignatureData) (*WS100101Response, error) {
	// Prepare request
	request := &WS100101Request{
		Request: Request{
			Regnum: registration,
			Auth: Auth{
				Citizen: CitizenAuth{
					CivilID:     "",
					Regnum:      registration,
					Fingerprint: "",
					AuthType:    1, // 1-OTP, 2-Digital signature, 3-Fingerprint
				},
				Operator: OperatorAuth{
					Regnum:      "",
					Fingerprint: "",
				},
			},
		},
	}

	// Add headers
	c.soapClient.AddHeader("accessToken", signData.AccessToken)
	c.soapClient.AddHeader("timeStamp", fmt.Sprintf("%s", signData.Timestamp))
	c.soapClient.AddHeader("signature", signData.Signature)

	// Prepare response
	response := &WS100101Response{}

	// Make the call
	err := c.soapClient.Call("WS100101_getCitizenIDCardInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
