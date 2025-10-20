package handlers

type ServerConfig struct {
	Port string
	Log  string
	Dump string
	Spa  string
	// SSL
	Crt string
	Key string
}

type Request struct {
	Method           string
	Url              string
	Proto            string
	Header           map[string][]string
	Body             string
	ContentLength    int64
	Host             string
	RemoteAddr       string
	RemotePort       string
	RequestURI       string
	RequestedCounter uint64
}
