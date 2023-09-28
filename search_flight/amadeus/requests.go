package amadeus

import "encoding/xml"

type Envelope struct {
	XMLName xml.Name `xml:"soapenv:Envelope"`
	Sec     *string  `xml:"xmlns:sec,attr"`
	Link    *string  `xml:"xmlns:link,attr"`
	Ses     *string  `xml:"xmlns:ses,attr"`
	Header  *Header  `xml:"soapenv:Header"`
}

type Header struct {
	Wsa                   *string                `xml:"xmlns:wsa,attr"`
	Typ                   *string                `xml:"xmlns:typ,attr"`
	Iat                   *string                `xml:"xmlns:iat,attr"`
	AMASecurityHostedUser *AMASecurityHostedUser `xml:"sec:AMA_SecurityHostedUser"`
	Security              *Security              `xml:"wsse:Security"`

	To        *string `xml:"http://www.w3.org/2005/08/addressing To"`
	Action    *string `xml:"http://www.w3.org/2005/08/addressing Action"`
	MessageID *string `xml:"http://www.w3.org/2005/08/addressing MessageID"`

	TransactionFlowLink struct {
		Consumer struct {
			UniqueID string `xml:"http://wsdl.amadeus.com/2010/06/ws/Link_v1 UniqueID"`
		} `xml:"http://wsdl.amadeus.com/2010/06/ws/Link_v1 Consumer"`
	} `xml:"http://wsdl.amadeus.com/2010/06/ws/Link_v1 TransactionFlowLink"`

	Session struct {
		TransactionStatusCode string `xml:"TransactionStatusCode,attr"`
	} `xml:"http://xml.amadeus.com/2010/06/Session_v3 Session"`
}

type AMASecurityHostedUser struct {
	UserID UserID `xml:"sec:UserID"`
}

type UserID struct {
	PosType        string      `xml:"POS_Type,attr"`
	RequestorType  string      `xml:"RequestorType,attr"`
	PseudoCityCode string      `xml:"PseudoCityCode,attr"`
	AgentDutyCode  *string     `xml:"AgentDutyCode,attr"`
	RequestorID    RequestorID `xml:"typ:RequestorID"`
}

type RequestorID struct {
	CompanyName *string `xml:"iat:CompanyName"`
}

type Security struct {
	Wsse          string        `xml:"xmlns:wsse,attr"`
	Wsu           string        `xml:"xmlns:wsu,attr"`
	UsernameToken UsernameToken `xml:"wsse:UsernameToken"`
}

type UsernameToken struct {
	Username string   `xml:"wsse:Username"`
	Password Password `xml:"wsse:Password"`
	Nonce    Nonce    `xml:"wsse:Nonce"`
	Created  string   `xml:"wsu:Created"`
}

type Password struct {
	Type string `xml:"Type,attr"`
	Text string `xml:",chardata"`
}

type Nonce struct {
	EncodingType string `xml:"EncodingType,attr"`
	Text         string `xml:",chardata"`
}
