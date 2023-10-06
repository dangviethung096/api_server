package amadeus

import "encoding/xml"

type RequestEnvelope struct {
	XMLName xml.Name       `xml:"soapenv:Envelope"`
	Env     *string        `xml:"xmlns:soapenv,attr"`
	Sec     *string        `xml:"xmlns:sec,attr"`
	Link    *string        `xml:"xmlns:link,attr"`
	Ses     *string        `xml:"xmlns:ses,attr"`
	Header  *RequestHeader `xml:"soapenv:Header"`
	Body    interface{}    `xml:"soapenv:Body"`
}

type RequestHeader struct {
	Wsa                   *string                `xml:"xmlns:wsa,attr"`
	Typ                   *string                `xml:"xmlns:typ,attr"`
	Iat                   *string                `xml:"xmlns:iat,attr"`
	AMASecurityHostedUser *AMASecurityHostedUser `xml:"sec:AMA_SecurityHostedUser"`
	Security              *Security              `xml:"wsse:Security"`
	To                    *string                `xml:"wsa:To"`
	Action                *string                `xml:"wsa:Action"`
	MessageID             *string                `xml:"wsa:MessageID"`
	TransactionFlowLink   *TransactionFlowLink   `xml:"link:TransactionFlowLink"`
	Session               *Session               `xml:"ses:Session"`
}

type AMASecurityHostedUser struct {
	UserID *UserID `xml:"sec:UserID"`
}

type UserID struct {
	PosType        *string      `xml:"POS_Type,attr"`
	RequestorType  *string      `xml:"RequestorType,attr"`
	PseudoCityCode *string      `xml:"PseudoCityCode,attr"`
	AgentDutyCode  *string      `xml:"AgentDutyCode,attr"`
	RequestorID    *RequestorID `xml:"typ:RequestorID"`
}

type RequestorID struct {
	CompanyName *string `xml:"iat:CompanyName"`
}

type Security struct {
	Wsse          *string        `xml:"xmlns:wsse,attr"`
	Wsu           *string        `xml:"xmlns:wsu,attr"`
	UsernameToken *UsernameToken `xml:"wsse:UsernameToken"`
}

type UsernameToken struct {
	Username *string   `xml:"wsse:Username"`
	Password *Password `xml:"wsse:Password"`
	Nonce    *Nonce    `xml:"wsse:Nonce"`
	Created  *string   `xml:"wsu:Created"`
}

type Password struct {
	Type *string `xml:"Type,attr"`
	Text *string `xml:",chardata"`
}

type Nonce struct {
	EncodingType *string `xml:"EncodingType,attr"`
	Text         *string `xml:",chardata"`
}

type Session struct {
	TransactionStatusCode *string `xml:"TransactionStatusCode,attr"`
}

type TransactionFlowLink struct {
	Consumer *Consumer `xml:"link:Consumer"`
}

type Consumer struct {
	UniqueID *string `xml:"link:UniqueID"`
}
