package amadeus

import "encoding/xml"

type SearchAirMultiAvailabitiyResponse struct {
	XMLName xml.Name                               `xml:"soapenv:Envelope"`
	Env     *string                                `xml:"xmlns:soapenv,attr"`
	Sec     *string                                `xml:"xmlns:sec,attr"`
	Link    *string                                `xml:"xmlns:link,attr"`
	Ses     *string                                `xml:"xmlns:ses,attr"`
	Header  *ResponseHeader                        `xml:"soapenv:Header"`
	Body    *SearchAirMultiAvailabitiyResponseBody `xml:"soapenv:Body"`
}

type SearchAirMultiAvailabitiyResponseBody struct {
	AirMultiAvailability *AirMultiAvailability `xml:"Air_MultiAvailability"`
}
