package defaults

const (
	appPort  = ":8080"
	dumpPath = "/app/stats/dump.json"
)

func GetAppPort() string {
	return appPort
}

func GetDumpPath() string {
	return dumpPath
}
