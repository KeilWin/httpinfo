package defaults

const (
	defaultHomeHandlerBodyLimitInBytes  int64 = 1000
	defaultHomeHandlerHeadersCountLimit int64 = 1000
)

func GetHomeHandlerBodyBytesLimitInBytes() int64 {
	return defaultHomeHandlerBodyLimitInBytes
}

func GetHomeHandlerHeadersCountLimit() int64 {
	return defaultHomeHandlerHeadersCountLimit
}
