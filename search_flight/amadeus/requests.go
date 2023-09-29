package amadeus

import "encoding/xml"

type SearchAirMultiAvailabitiyRequest struct {
	XMLName xml.Name                       `xml:"soapenv:Envelope"`
	Env     *string                        `xml:"xmlns:soapenv,attr"`
	Sec     *string                        `xml:"xmlns:sec,attr"`
	Link    *string                        `xml:"xmlns:link,attr"`
	Ses     *string                        `xml:"xmlns:ses,attr"`
	Header  *Header                        `xml:"soapenv:Header"`
	Body    *SearchAirMultiAvailabitiyBody `xml:"soapenv:Body"`
}

type SearchAirMultiAvailabitiyBody struct {
	AirMultiAvailability *AirMultiAvailability `xml:"Air_MultiAvailability"`
}

type AirMultiAvailability struct {
	MessageActionDetails *MessageActionDetails `xml:"messageActionDetails"`
	RequestSection       *RequestSection       `xml:"requestSection"`
}

type MessageActionDetails struct {
	FunctionDetails *FunctionDetails `xml:"functionDetails"`
}

type FunctionDetails struct {
	BusinessFunction *string `xml:"businessFunction"`
	ActionCode       *string `xml:"actionCode"`
}

type RequestSection struct {
	AvailabilityProductInfo *AvailabilityProductInfo `xml:"availabilityProductInfo"`
	AvailabilityOptions     *AvailabilityOptions     `xml:"availabilityOptions"`
}

type AvailabilityProductInfo struct {
	AvailabilityDetails   *AvailabilityDetails   `xml:"availabilityDetails"`
	DepartureLocationInfo *DepartureLocationInfo `xml:"departureLocationInfo"`
	ArrivalLocationInfo   *ArrivalLocationInfo   `xml:"arrivalLocationInfo"`
}

type AvailabilityDetails struct {
	DepartureDate *string `xml:"departureDate"`
}

type DepartureLocationInfo struct {
	CityAirport *string `xml:"cityAirport"`
}

type ArrivalLocationInfo struct {
	CityAirport *string `xml:"cityAirport"`
}

type AvailabilityOptions struct {
	TypeOfRequest *string `xml:"typeOfRequest"`
}
