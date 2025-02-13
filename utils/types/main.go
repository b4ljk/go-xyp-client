package types

type PassportDataType struct {
	Envelope Envelope `xml:"Envelope"`
}

type Envelope struct {
	Body      Body   `xml:"Body"`
	XmlnsSoap string `xml:"_xmlns:soap"`
	Prefix    string `xml:"__prefix"`
}

type Body struct {
	WS100101GetCitizenIDCardInfoResponse WS100101GetCitizenIDCardInfoResponse `xml:"WS100101_getCitizenIDCardInfoResponse"`
	Prefix                               string                               `xml:"__prefix"`
}

type WS100101GetCitizenIDCardInfoResponse struct {
	Return   Return `xml:"return"`
	XmlnsNs2 string `xml:"_xmlns:ns2"`
	Prefix   string `xml:"__prefix"`
}

type Return struct {
	Request       Request  `xml:"request"`
	RequestID     string   `xml:"requestId"`
	Response      Response `xml:"response"`
	ResultCode    string   `xml:"resultCode"`
	ResultMessage string   `xml:"resultMessage"`
}

type Request struct {
	Auth     Auth   `xml:"auth"`
	Regnum   string `xml:"regnum"`
	XmlnsXsi string `xml:"_xmlns:xsi"`
	XsiType  string `xml:"_xsi:type"`
}

type Auth struct {
	Citizen  Citizen `xml:"citizen"`
	Operator Citizen `xml:"operator"`
}

type Citizen struct {
	AuthType    string `xml:"authType"`
	CivilID     string `xml:"civilId"`
	Fingerprint string `xml:"fingerprint"`
	Otp         string `xml:"otp"`
	Regnum      string `xml:"regnum"`
	Signature   string `xml:"signature"`
}

type Response struct {
	AddressDetail      string `xml:"addressDetail"`
	AddressStreetName  string `xml:"addressStreetName"`
	AimagCityCode      string `xml:"aimagCityCode"`
	AimagCityName      string `xml:"aimagCityName"`
	BagKhorooCode      string `xml:"bagKhorooCode"`
	BagKhorooName      string `xml:"bagKhorooName"`
	BirthDateAsText    string `xml:"birthDateAsText"`
	BirthPlace         string `xml:"birthPlace"`
	CivilID            string `xml:"civilId"`
	Firstname          string `xml:"firstname"`
	Gender             string `xml:"gender"`
	Lastname           string `xml:"lastname"`
	Nationality        string `xml:"nationality"`
	PassTime           string `xml:"passTime"`
	PassportAddress    string `xml:"passportAddress"`
	PassportExpireDate string `xml:"passportExpireDate"`
	PassportIssueDate  string `xml:"passportIssueDate"`
	PassportNum        string `xml:"passportNum"`
	Regnum             string `xml:"regnum"`
	SoumDistrictCode   string `xml:"soumDistrictCode"`
	SoumDistrictName   string `xml:"soumDistrictName"`
	Surname            string `xml:"surname"`
	XmlnsXsi           string `xml:"_xmlns:xsi"`
	XsiType            string `xml:"_xsi:type"`
	Image              string `xml:"image"`
}
