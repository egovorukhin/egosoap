package soap

import "encoding/xml"

type Input struct {
	Envelope Request
}

type Request struct {
	XMLName xml.Name    `xml:"soapenv:Envelope"`
	SoapEnv string      `xml:"xmlns:soapenv,attr"`
	Ns      *string     `xml:"xmlns:ns,attr,omitempty"`
	Cus     *string     `xml:"xmlns:cus,attr,omitempty"`
	Agent   *string     `xml:"xmlns:agent,attr,omitempty"`
	Header  interface{} `xml:"soapenv:Header"`
	Body    interface{} `xml:"soapenv:Body"`
}
