package egosoap

import (
	"fmt"
	"github.com/egovorukhin/egosoap/ecxecute_sql_query"
	"testing"
)

type CallingSearchSpace struct {
	PkId                string `xml:"pkid"`
	Name                string `xml:"name"`
	Description         string `xml:"description"`
	DialPlanWizardGenId string `xml:"dialplanwizardgenid" db:"dial_plan_wizard_genid"`
	Clause              string `xml:"clause"`
	ResetToggle         string `xml:"resettoggle" db:"reset_toggle"`
	TkReset             int    `xml:"tkreset" db:"tk_reset"`
	TkPartitionUsage    string `xml:"tkpartitionusage" db:"tk_partition_usage"`
}

type CallingSearchSpaces []CallingSearchSpace

func TestInit(t *testing.T) {
	soap := Init(
		InitAttribute("http://www.cisco.com/AXL/API/11.5", "", ""),
		FaultOutputBody(),
	)
	user := "user"
	pass := "pass"
	soap.Server = Server{
		Hostname: "server",
		Port:     8443,
		Timeout:  30,
		Secure:   true,
		Username: &user,
		Password: &pass,
		Route:    "/route/",
	}
	css := CallingSearchSpaces{}
	_, err := soap.Connection(
		ecxecute_sql_query.SoapAction,
		ecxecute_sql_query.SetInputBody(
			fmt.Sprintf("select * from callingsearchspace where 1=1 %s", ""),
			"?"),
		ecxecute_sql_query.SetOutputBody(&css),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, value := range css {
		fmt.Printf("%v\n", value)
	}
}
