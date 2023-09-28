package amadeus

import "encoding/xml"

type SearchAirMultiAvailabitiy struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Env     *string  `xml:"xmlns:soapenv,attr"`
	Sec     *string  `xml:"xmlns:sec,attr"`
	Link    *string  `xml:"xmlns:link,attr"`
	Ses     *string  `xml:"xmlns:ses,attr"`
	Header  *Header  `xml:"soapenv:Header"`
}
