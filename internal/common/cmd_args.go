package common

import (
	"flag"

	df "httpinfo/internal/defaults"
	"httpinfo/internal/handlers"
)

func InitCmdArgs(serverCfg *handlers.ServerConfig) {
	flag.StringVar(&serverCfg.Port, "app-port", df.GetAppPort(), "Application port")
	flag.StringVar(&serverCfg.Dump, "dump-path", df.GetDumpPath(), "Dump statistic path")
	flag.StringVar(&serverCfg.Crt, "crt-path", df.GetCrtPath(), "Certificate SSL path")
	flag.StringVar(&serverCfg.Key, "key-path", df.GetKeyPath(), "Key SSL path")
	flag.StringVar(&serverCfg.TemplateCfg.Index, "index-template-path", df.GetIndexTemplatePath(), "Index template path")
}

func ParseCmdArgs() {
	flag.Parse()
}
