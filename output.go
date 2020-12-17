package egosoap

import "encoding/json"

type Output struct {
	Envelope Response `xml:"Envelope"`
	Error    bool
}

type Response struct {
	//XMLName xml.Name `xml:"Envelope"`
	SoapEnv string      `xml:"soapenv,attr"`
	Body    interface{} `xml:"Body"`
}

func SetOutput(body interface{}) *Output {
	return &Output{
		Envelope: Response{
			SoapEnv: "",
			Body:    body,
		},
	}
}

func (output Output) String() string {
	b, _ := json.Marshal(output)
	return string(b)
}
