package defaults

const (
	logPath = "/app/logs/httpinfo.log"
)

func GetLogPath() string {
	return logPath
}
