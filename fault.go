package egosoap

type OutputBody struct {
	Fault Fault `xml:"Fault"`
}

type Fault struct {
	FaultCode   string   `xml:"faultcode"`
	FaultString string   `xml:"faultstring"`
	Details     []Detail `xml:"detail"`
}

type Detail struct {
	AxlError AxlError `xml:"axlError"`
}

type AxlError struct {
	AxlCode    int    `xml:"axlcode"`
	AxlMessage string `xml:"axlmessage"`
	Request    string `xml:"request"`
}

func FaultOutputBody() *OutputBody {
	return &OutputBody{
		Fault: Fault{},
	}
}
