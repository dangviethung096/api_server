package amadeus

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestXmlRequest(t *testing.T) {
	postType := "1"
	// companyName := "Name"
	envelope := Envelope{
		Sec:  "http://xml.amadeus.com/2010/06/Security_v1",
		Link: "http://wsdl.amadeus.com/2010/06/ws/Link_v1",
		Ses:  "http://xml.amadeus.com/2010/06/Session_v3",
		Header: Header{
			Wsa: "http://www.w3.org/2005/08/addressing",
			Typ: "http://xml.amadeus.com/2010/06/Types_v1",
			Iat: "http://www.iata.org/IATA/2007/00/IATA2010.1",
			AMASecurityHostedUser: AMASecurityHostedUser{
				UserID: UserID{
					PosType:        postType,
					RequestorType:  "U",
					PseudoCityCode: "HANVN08AA",
					AgentDutyCode:  nil,
					RequestorID: RequestorID{
						CompanyName: nil,
					},
				},
			},
		},
	}

	dataByte, err := xml.Marshal(envelope)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Data in byte: \n%s\n", string(dataByte))
}
