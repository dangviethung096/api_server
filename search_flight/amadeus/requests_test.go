package amadeus

import (
	"encoding/xml"
	"fmt"
	"testing"
)

var (
	Env                   = "http://schemas.xmlsoap.org/soap/envelope/"
	Sec                   = "http://schemas.xmlsoap.org/soap/envelope/"
	Link                  = "http://xml.amadeus.com/2010/06/Security_v1"
	Ses                   = "http://wsdl.amadeus.com/2010/06/ws/Link_v1"
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
	PasswordText          = "Irwhru5lS3mwKMyUods/p72MWDE="
	EncodingType          = "http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-soap-message-security-1.0#Base64Binary"
	NonceText             = "Z0ZHQWh1bmx6Vnpu"
	Created               = "2022-11-15T17:00:10.217Z"
	To                    = "https://nodeA1.test.webservices.amadeus.com/1ASIWGENVN"
	Action                = "http://webservices.amadeus.com/SATRQT_19_1_1A"
	MessageID             = "urn:uuid:e88e9b9d-7a0f-46e0-8f01-e0cf6325cb67"
	UniqueID              = "oef5izfzuaawy3ergzk5005h"
	TransactionStatusCode = "Start"
)

func TestXmlRequest(t *testing.T) {
	envelope := SearchAirMultiAvailabitiy{
		Env:  &Env,
		Sec:  &Sec,
		Link: &Link,
		Ses:  &Ses,
		Header: &Header{
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
						Text: &PasswordText,
					},
					Nonce: &Nonce{
						EncodingType: &EncodingType,
						Text:         &NonceText,
					},
					Created: &Created,
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
		},
	}

	dataByte, err := xml.Marshal(envelope)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data in byte: \n%s\n", string(dataByte))
}
