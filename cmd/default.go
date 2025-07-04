package main

const defaultHomeHandlerBodyLimitInBytes int64 = 1000
const defaultHomeHandlerHeadersCountLimit int64 = 1000

const defaultAppPort = ":8080"
const defaultIndexTemplate = "web/template/index.html"
const defaultCrtPath = "/etc/ssl/server.crt"
const defaultKeyPath = "/etc/ssl/server.key"
const defaultDumpPath = "/app/stats/dump.json"

func GetHomeHandlerBodyBytesLimitInBytes() int64 {
	return defaultHomeHandlerBodyLimitInBytes
}

func GetHomeHandlerHeadersCountLimit() int64 {
	return defaultHomeHandlerHeadersCountLimit
}
