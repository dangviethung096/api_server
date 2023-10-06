package amadeus

import (
	core "api_server_core"
	"encoding/xml"
	"fmt"
	"net/http"
	"testing"
)

func TestXmlRequest(t *testing.T) {
	envelope := SearchAirMultiAvailabitiyRequest{
		Env:    &Env,
		Sec:    &Sec,
		Link:   &Link,
		Ses:    &Ses,
		Header: buildHeader(),
		Body:   buildBody(),
	}

	dataByte, err := xml.Marshal(envelope)
	if err != nil {
		t.Errorf("Build body request: %v\n", err)
		return
	}

	url := "https://nodeA1.test.webservices.amadeus.com/1ASIWGENVN"
	headers := map[string]string{
		"Content-Type": "text/xml; charset=utf-8",
		"SOAPAction":   "http://webservices.amadeus.com/SATRQT_19_1_1A",
	}

	// Request to 1A http server
	res, err := core.HttpRequest(url, http.MethodPost, headers, dataByte)
	if err != nil {
		t.Errorf("Request to 1A: %v\n", err)
	}

	fmt.Printf("Data in byte: \n%s\n", string(res.Body))
}

func buildBody() *SearchAirMultiAvailabitiyRequestBody {
	businessFunction := "1"
	actionCode := "44"
	departureDate := "151023"
	typeOfRequest := "TD"
	departureAirport := "HAN"
	arrivalAirport := "SGN"

	return &SearchAirMultiAvailabitiyRequestBody{
		AirMultiAvailability: &AirMultiAvailability{
			MessageActionDetails: &MessageActionDetails{
				FunctionDetails: &FunctionDetails{
					BusinessFunction: &businessFunction,
					ActionCode:       &actionCode,
				},
			},
			RequestSection: &RequestSection{
				AvailabilityProductInfo: &AvailabilityProductInfo{
					AvailabilityDetails: &AvailabilityDetails{
						DepartureDate: &departureDate,
					},
					DepartureLocationInfo: &DepartureLocationInfo{
						CityAirport: &departureAirport,
					},
					ArrivalLocationInfo: &ArrivalLocationInfo{
						CityAirport: &arrivalAirport,
					},
				},
				AvailabilityOptions: &AvailabilityOptions{
					TypeOfRequest: &typeOfRequest,
				},
			},
		},
	}
}
