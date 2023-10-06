package amadeus

import "encoding/xml"

type SearchFareInformativePricingWithoutPNR struct {
	XMLName xml.Name                              `xml:"soapenv:Envelope"`
	Env     *string                               `xml:"xmlns:soapenv,attr"`
	Sec     *string                               `xml:"xmlns:sec,attr"`
	Link    *string                               `xml:"xmlns:link,attr"`
	Ses     *string                               `xml:"xmlns:ses,attr"`
	Header  *RequestHeader                        `xml:"soapenv:Header"`
	Body    *SearchAirMultiAvailabitiyRequestBody `xml:"soapenv:Body"`
}

type SearchFareInformativePricingWithoutPNRBody struct {
	FareInformativePricingWithoutPNR FareInformativePricingWithoutPNR `xml:"Fare_InformativePricingWithoutPNR"`
}

type FareInformativePricingWithoutPNR struct {
	PassengersGroup    []PassengersGroup    `xml:"passengersGroup"`
	SegmentGroup       []SegmentGroup       `xml:"segmentGroup"`
	PricingOptionGroup []PricingOptionGroup `xml:"pricingOptionGroup"`
}

type PassengersGroup struct {
	SegmentRepetitionControl SegmentRepetitionControl `xml:"segmentRepetitionControl"`
	TravellersID             TravellersID             `xml:"travellersID"`
}

type SegmentRepetitionControl struct {
	SegmentControlDetails SegmentControlDetails `xml:"segmentControlDetails"`
}

type SegmentControlDetails struct {
	Quantity      int `xml:"quantity"`
	NumberOfUnits int `xml:"numberOfUnits"`
}

type TravellersID struct {
	TravellerDetails TravellerDetails `xml:"travellerDetails"`
}

type TravellerDetails struct {
	MeasurementValue int `xml:"measurementValue"`
}

type SegmentGroup struct {
	SegmentInformation SegmentInformation `xml:"segmentInformation"`
}

type SegmentInformation struct {
	FlightDate           FlightDate           `xml:"flightDate"`
	BoardPointDetails    BoardPointDetails    `xml:"boardPointDetails"`
	OffpointDetails      OffpointDetails      `xml:"offpointDetails"`
	CompanyDetails       CompanyDetails       `xml:"companyDetails"`
	FlightIdentification FlightIdentification `xml:"flightIdentification"`
}

type FlightDate struct {
	DepartureDate string `xml:"departureDate"`
	DepartureTime string `xml:"departureTime"`
	ArrivalDate   string `xml:"arrivalDate"`
	ArrivalTime   string `xml:"arrivalTime"`
}

type BoardPointDetails struct {
	TrueLocationID string `xml:"trueLocationId"`
}

type OffpointDetails struct {
	TrueLocationID string `xml:"trueLocationId"`
}

type CompanyDetails struct {
	MarketingCompany string `xml:"marketingCompany"`
	OperatingCompany string `xml:"operatingCompany"`
}

type FlightIdentification struct {
	FlightNumber string `xml:"flightNumber"`
	BookingClass string `xml:"bookingClass"`
}

type PricingOptionGroup struct {
	PricingOptionKey   string `xml:"pricingOptionKey>pricingOptionKey"`
	PaxSegTstReference string `xml:"paxSegTstReference"`
}
