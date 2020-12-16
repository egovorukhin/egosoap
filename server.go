package soap

type Server struct {
	Hostname string
	Port int
	Timeout int
	Secure bool
	Username *string
	Password *string
	Route string
}
