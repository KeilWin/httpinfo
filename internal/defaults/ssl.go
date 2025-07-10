package defaults

const (
	crtPath = "/etc/ssl/server.crt"
	keyPath = "/etc/ssl/server.key"
)

func GetCrtPath() string {
	return crtPath
}

func GetKeyPath() string {
	return keyPath
}
