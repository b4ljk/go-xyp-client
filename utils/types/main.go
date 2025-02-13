package types

type PassportDataType struct {
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
	AddressDetail      string `xml:"addressDetail" json:"addressDetail"`
	AddressStreetName  string `xml:"addressStreetName" json:"addressStreetName"`
	AimagCityCode      string `xml:"aimagCityCode" json:"aimagCityCode"`
	AimagCityName      string `xml:"aimagCityName" json:"aimagCityName"`
	BagKhorooCode      string `xml:"bagKhorooCode" json:"bagKhorooCode"`
	BagKhorooName      string `xml:"bagKhorooName" json:"bagKhorooName"`
	BirthDateAsText    string `xml:"birthDateAsText" json:"birthDateAsText"`
	BirthPlace         string `xml:"birthPlace" json:"birthPlace"`
	CivilID            string `xml:"civilId" json:"civilId"`
	Firstname          string `xml:"firstname" json:"firstname"`
	Gender             string `xml:"gender" json:"gender"`
	Lastname           string `xml:"lastname" json:"lastname"`
	Nationality        string `xml:"nationality" json:"nationality"`
	PassTime           string `xml:"passTime" json:"passTime"`
	PassportAddress    string `xml:"passportAddress" json:"passportAddress"`
	PassportExpireDate string `xml:"passportExpireDate" json:"passportExpireDate"`
	PassportIssueDate  string `xml:"passportIssueDate" json:"passportIssueDate"`
	PassportNum        string `xml:"passportNum" json:"passportNum"`
	Regnum             string `xml:"regnum" json:"regnum"`
	SoumDistrictCode   string `xml:"soumDistrictCode" json:"soumDistrictCode"`
	SoumDistrictName   string `xml:"soumDistrictName" json:"soumDistrictName"`
	Surname            string `xml:"surname" json:"surname"`
	Image              string `xml:"image" json:"image"`
}
