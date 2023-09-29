package amadeus

import (
	core "api_server_core"
	"crypto/sha1"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

var (
	Env                   = "http://schemas.xmlsoap.org/soap/envelope/"
	Sec                   = "http://xml.amadeus.com/2010/06/Security_v1"
	Link                  = "http://wsdl.amadeus.com/2010/06/ws/Link_v1"
	Ses                   = "http://xml.amadeus.com/2010/06/Session_v3"
	Wsa                   = "http://www.w3.org/2005/08/addressing"
	Typ                   = "http://xml.amadeus.com/2010/06/Types_v1"
	Iat                   = "http://www.iata.org/IATA/2007/00/IATA2010.1"
	PosType               = "1"
	RequestorType         = "U"
	PseudoCityCode        = "HANVN08AA"
	AgentDutyCode         = "SU"
	CompanyName           = "VN"
	Wsse                  = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-secext-1.0.xsd"
	Wsu                   = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd"
	Username              = "WSVNGEN"
	PasswordType          = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-username-token-profile-1.0#PasswordDigest"
	EncodingType          = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary"
	To                    = "https://nodeA1.test.webservices.amadeus.com/1ASIWGENVN"
	Action                = "http://webservices.amadeus.com/SATRQT_19_1_1A"
	MessageID             = "urn:uuid:e88e9b9d-7a0f-46e0-8f01-e0cf6325cb67"
	UniqueID              = "oef5izfzuaawy3ergzk5005h"
	TransactionStatusCode = "Start"
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

func buildBody() *SearchAirMultiAvailabitiyBody {
	businessFunction := "1"
	actionCode := "44"
	departureDate := "151023"
	typeOfRequest := "TD"
	departureAirport := "HAN"
	arrivalAirport := "SGN"

	return &SearchAirMultiAvailabitiyBody{
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

func toBase64(inputBytes []byte) string {
	return base64.StdEncoding.EncodeToString(inputBytes)
}

func generateRandomString() string {
	var text string
	var possible = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	for i := 0; i < 12; i++ {
		text += string(possible[rand.Intn(len(possible))])
	}
	return text
}

func saltPassword(nonce string, timestamp string, password string) string {
	h := sha1.New()
	h.Write([]byte(nonce))
	h.Write([]byte(timestamp))
	h.Write(h.Sum(nil))
	h.Write([]byte(h.Sum(nil)))
	h.Write([]byte(password))
	return toBase64(h.Sum(nil))
}

func buildHeader() *Header {
	now := time.Now().UTC().Format("2006-01-02T15:04:05Z")
	nonce := generateRandomString()
	password := saltPassword(nonce, now, "Amadeus@24")
	nonceBase64 := toBase64([]byte(nonce))

	return &Header{
		Wsa: &Wsa,
		Typ: &Typ,
		Iat: &Iat,
		AMASecurityHostedUser: &AMASecurityHostedUser{
			UserID: &UserID{
				PosType:        &PosType,
				RequestorType:  &RequestorType,
				PseudoCityCode: &PseudoCityCode,
				AgentDutyCode:  &AgentDutyCode,
				RequestorID: &RequestorID{
					CompanyName: &CompanyName,
				},
			},
		},
		Security: &Security{
			Wsse: &Wsse,
			Wsu:  &Wsu,
			UsernameToken: &UsernameToken{
				Username: &Username,
				Password: &Password{
					Type: &PasswordType,
					Text: &password,
				},
				Nonce: &Nonce{
					EncodingType: &EncodingType,
					Text:         &nonceBase64,
				},
				Created: &now,
			},
		},
		To:        &To,
		Action:    &Action,
		MessageID: &MessageID,
		TransactionFlowLink: &TransactionFlowLink{
			Consumer: &Consumer{
				UniqueID: &UniqueID,
			},
		},
		Session: &Session{
			TransactionStatusCode: &TransactionStatusCode,
		},
	}
}
