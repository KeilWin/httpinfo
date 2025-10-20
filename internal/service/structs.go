package service

type ServerConfig struct {
	Port string
	Log  string
	Dump string
	// SSL
	Crt string
	Key string

	TemplateCfg *TemplatesConfig
}

type TemplatesConfig struct {
	Index string

	App string

	Header string

	Content   string
	LeftSide  string
	RightSide string

	Footer string
}

type Request struct {
	Method           string              `json:"method"`
	Url              string              `json:"url"`
	Proto            string              `json:"proto"`
	Header           map[string][]string `json:"header"`
	Body             string              `json:"body"`
	ContentLength    int64               `json:"contentLength"`
	Host             string              `json:"host"`
	RemoteAddr       string              `json:"remoteAddr"`
	RemotePort       string              `json:"remotePort"`
	RequestURI       string              `json:"requestURI"`
	RequestedCounter uint64              `json:"requestCounter"`
}
