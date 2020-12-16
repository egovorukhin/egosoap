package soap

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/egovorukhin/egorest"
)

type Soap struct {
	SoapAction string
	Input      Input
	Output     *Output
	Attributes Attribute
	Error      interface{}
}

type Attribute struct {
	SoapEnv string
	Ns      *string
	Cus     *string
	Agent   *string
}

func InitAttribute(Ns, Cus, Agent string) Attribute {
	var ns *string
	var cus *string
	var agent *string
	if Ns != "" {
		ns = &Ns
	}
	if Cus != "" {
		cus = &Cus
	}
	if Agent != "" {
		agent = &Agent
	}
	return Attribute{
		Ns:    ns,
		Cus:   cus,
		Agent: agent,
	}
}

//var soaps map[string]*Soap

func Init(attr Attribute, errBody interface{}) *Soap {

	s := &Soap{
		Input: Input{
			Envelope: Request{
				SoapEnv: "http://schemas.xmlsoap.org/soap/envelope/",
				Ns:      attr.Ns,
				Cus:     attr.Cus,
				Agent:   attr.Agent,
			},
		},
		Output: &Output{
			Envelope: Response{},
		},
		Error: errBody,
	}

	return s
}

func (s *Soap) SetHeader(h interface{}) Soap {
	s.Input.Envelope.Header = h
	return *s
}

func (s Soap) Connection(srv Server, soapAction string, inBody, outBody interface{}) (*Output, error) {

	s.SoapAction = soapAction
	s.Input.Envelope.Body = inBody
	s.Output.Envelope.Body = outBody

	output, err := s.connection(srv)
	if err != nil {
		return output, err
	}

	return output, nil
}

//Выполняем запрос до url, передаём xml soap иполучаем ответ
func (soap Soap) connection(srv Server) (*Output, error) {

	soap.Output.Error = false

	req := egorest.NewRequest(egorest.POST, srv.Route).
		SetHeader(
			egorest.SetHeader("Content-Type", "text/xml;charset=UTF-8"),
			egorest.SetHeader("SOAPAction", soap.SoapAction),
		)
	req.Data = &egorest.Data{
		ContentType: egorest.XML,
		Body:        soap.Input.Envelope,
	}

	client := egorest.NewClient(srv.Hostname, srv.Port, srv.Secure).SetTimeout(srv.Timeout)
	if srv.Username != nil || srv.Password != nil {
		client.SetBasicAuth(*srv.Username, *srv.Password)
	}

	err := client.Execute(req, &soap.Output.Envelope)
	//Если появилась ошибка, то нам нужно десериализовать
	//данные в Fault структуру и передать в качестве ошибки само значение
	if err != nil {
		soap.Output.Error = true
		soap.Output.Envelope.Body = soap.Error
		err = xml.Unmarshal([]byte(err.Error()), &soap.Output.Envelope)
		if err != nil {
			return nil, err
		}
		if soap.Error == nil {
			return nil, err
		}
		return nil, errors.New(fmt.Sprintf("%v", soap.Error))
	}

	return soap.Output, nil
}
