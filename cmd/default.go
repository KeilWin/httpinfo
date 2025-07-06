package main

const (
	defaultHomeHandlerBodyLimitInBytes  int64 = 1000
	defaultHomeHandlerHeadersCountLimit int64 = 1000
)

const defaultAppPort = ":8080"

const (
	defaultIndexTemplate     = "web/template/index.html"
	defaultAppTemplate       = "web/template/app.html"
	defaultHeaderTemplate    = "web/template/header.html"
	defaultContentTemplate   = "web/template/content.html"
	defaultFooterTemplate    = "web/template/footer.html"
	defaultLeftSideTemplate  = "web/template/leftSide.html"
	defaultRightSideTemplate = "web/template/rightSide.html"
)

const (
	defaultCrtPath = "/etc/ssl/server.crt"
	defaultKeyPath = "/etc/ssl/server.key"
)

const defaultDumpPath = "/app/stats/dump.json"

func GetHomeHandlerBodyBytesLimitInBytes() int64 {
	return defaultHomeHandlerBodyLimitInBytes
}

func GetHomeHandlerHeadersCountLimit() int64 {
	return defaultHomeHandlerHeadersCountLimit
}
