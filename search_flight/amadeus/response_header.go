package amadeus

import "encoding/xml"

type Response struct {
	XMLName    xml.Name       `xml:"Envelope"`
	XMLNSs     string         `xml:"xmlns:soapenv,attr"`
	XMLNSawsl  string         `xml:"xmlns:awsl,attr"`
	XMLNSawsse string         `xml:"xmlns:awsse,attr"`
	XMLNSwsa   string         `xml:"xmlns:wsa,attr"`
	Header     ResponseHeader `xml:"Header"`
}

type ResponseHeader struct {
	To                  string                            `xml:"wsa:To"`
	From                ResponseHeaderFrom                `xml:"wsa:From"`
	Action              string                            `xml:"wsa:Action"`
	MessageID           string                            `xml:"wsa:MessageID"`
	RelatesTo           ResponseHeaderRelatesTo           `xml:"wsa:RelatesTo"`
	TransactionFlowLink ResponseHeaderTransactionFlowLink `xml:"awsl:TransactionFlowLink"`
	Session             ResponseHeaderSession             `xml:"awsse:Session"`
}

type ResponseHeaderFrom struct {
	Address string `xml:"wsa:Address"`
}

type ResponseHeaderRelatesTo struct {
	RelationshipType string `xml:"RelationshipType,attr"`
	Value            string `xml:",chardata"`
}

type ResponseHeaderTransactionFlowLink struct {
	Consumer ResponseHeaderConsumer `xml:"Consumer"`
	Receiver ResponseHeaderReceiver `xml:"Receiver"`
}

type ResponseHeaderSession struct {
	TransactionStatusCode string `xml:"TransactionStatusCode,attr"`
	SessionId             string `xml:"awsse:SessionId"`
	SequenceNumber        int    `xml:"awsse:SequenceNumber"`
	SecurityToken         string `xml:"awsse:SecurityToken"`
}

type ResponseHeaderConsumer struct {
	UniqueID string `xml:"awsl:UniqueID"`
}

type ResponseHeaderReceiver struct {
	ServerID string `xml:"awsl:ServerID"`
}
