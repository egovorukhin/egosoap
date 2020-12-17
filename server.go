package egosoap

type Server struct {
	Url      string
	Hostname string
	Port     int
	Timeout  int
	Secure   bool
	Username *string
	Password *string
	Route    string
}
