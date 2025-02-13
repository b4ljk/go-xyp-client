package constants

const (
	XYP_PASSPORT_URL       = "https://xyp.gov.mn/citizen-1.5.0/ws?WSDL"
	XYP_PASSPORT_SOAP_BODY = `<?xml version="1.0" encoding="UTF-8"?>
    <soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/"
		xmlns:web="http://citizen.xyp.gov.mn/">
       <soapenv:Header/>
       <soapenv:Body>
          <web:WS100101_getCitizenIDCardInfo>
             <request>
                <regnum>%s</regnum>
                <auth>
                   <citizen>
                      <civilId></civilId>
                      <regnum>%s</regnum>
                      <fingerprint></fingerprint>
                      <authType>1</authType>
                   </citizen>
                   <operator>
                      <regnum></regnum>
                      <fingerprint></fingerprint>
                   </operator>
                </auth>
             </request>
          </web:WS100101_getCitizenIDCardInfo>
       </soapenv:Body>
    </soapenv:Envelope>`
)
