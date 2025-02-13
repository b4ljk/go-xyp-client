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
	AddressDetail      string `xml:"addressDetail" json:"address_detail"`
	AddressStreetName  string `xml:"addressStreetName" json:"address_street_name"`
	AimagCityCode      string `xml:"aimagCityCode" json:"aimag_city_code"`
	AimagCityName      string `xml:"aimagCityName" json:"aimag_city_name"`
	BagKhorooCode      string `xml:"bagKhorooCode" json:"bag_khoroo_code"`
	BagKhorooName      string `xml:"bagKhorooName" json:"bag_khoroo_name"`
	BirthDateAsText    string `xml:"birthDateAsText" json:"birth_date_as_text"`
	BirthPlace         string `xml:"birthPlace" json:"birth_place"`
	CivilID            string `xml:"civilId" json:"civil_id"`
	Firstname          string `xml:"firstname" json:"firstname"`
	Gender             string `xml:"gender" json:"gender"`
	Lastname           string `xml:"lastname" json:"lastname"`
	Nationality        string `xml:"nationality" json:"nationality"`
	PassTime           string `xml:"passTime" json:"pass_time"`
	PassportAddress    string `xml:"passportAddress" json:"passport_address"`
	PassportExpireDate string `xml:"passportExpireDate" json:"passport_expire_date"`
	PassportIssueDate  string `xml:"passportIssueDate" json:"passport_issue_date"`
	PassportNum        string `xml:"passportNum" json:"passport_num"`
	Regnum             string `xml:"regnum" json:"regnum"`
	SoumDistrictCode   string `xml:"soumDistrictCode" json:"soum_district_code"`
	SoumDistrictName   string `xml:"soumDistrictName" json:"soum_district_name"`
	Surname            string `xml:"surname" json:"surname"`
	Image              string `xml:"image" json:"image"`
}
