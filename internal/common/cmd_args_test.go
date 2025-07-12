package common_test

import (
	"httpinfo/internal/common"
	df "httpinfo/internal/defaults"
	"httpinfo/internal/handlers"
	"os"
	"testing"
)

var testServerCfg *handlers.ServerConfig

func init() {
	testServerCfg = handlers.NewServerConfig()
	common.InitCmdArgs(testServerCfg)
}

func TestCreateServerConfigWithEmptyCmdArgs(t *testing.T) {
	cmdArgs := os.Args
	defer func() {
		os.Args = cmdArgs
	}()
	t.Logf("cmdArgs: %v", cmdArgs)

	os.Args = []string{cmdArgs[0]}
	t.Logf("os.Args: %v", os.Args)

	serverCfg := testServerCfg
	common.ParseCmdArgs()

	if serverCfg.Port != df.GetAppPort() {
		t.Errorf("Port is unexpected: %s", serverCfg.Port)
	}

	if serverCfg.Dump != df.GetDumpPath() {
		t.Errorf("Dump path is unexpected: %s", serverCfg.Dump)
	}

	if serverCfg.Crt != df.GetCrtPath() {
		t.Errorf("SSL crt path is unexpected: %s", serverCfg.Crt)
	}

	if serverCfg.Key != df.GetKeyPath() {
		t.Errorf("SSL key path is unexpected: %s", serverCfg.Key)
	}

	if serverCfg.TemplateCfg.Index != df.GetIndexTemplatePath() {
		t.Errorf("Index template path is unexpected: %s", serverCfg.TemplateCfg.Index)
	}

	if serverCfg.Log != df.GetLogPath() {
		t.Errorf("Log path is unexpected: %s", serverCfg.Log)
	}
}

func TestCreateServerConfigWithNonEmptyCmdArgs(t *testing.T) {
	cmdArgs := os.Args
	defer func() {
		os.Args = cmdArgs
	}()
	t.Logf("cmdArgs: %v", cmdArgs)

	const (
		expectedAppPort           = ":9000"
		expectedDumpPath          = "/TEST/DUMP_PATH.json"
		expectedCrtPath           = "/TEST/CRT_PATH.crt"
		expectedKeyPath           = "/TEST/KEY_PATH.key"
		expectedIndexTemplatePath = "/TEST/INDEX_TEMPLATE_PATH.html"
		expectedLogPath           = "/TEST/LOG_PATH.LOG"
	)
	os.Args = append(os.Args,
		"--app-port", expectedAppPort,
		"--dump-path", expectedDumpPath,
		"--crt-path", expectedCrtPath,
		"--key-path", expectedKeyPath,
		"--index-template-path", expectedIndexTemplatePath,
		"--log-path", expectedLogPath,
	)
	t.Logf("os.Args: %v", os.Args)

	serverCfg := testServerCfg
	common.ParseCmdArgs()

	if serverCfg.Port != expectedAppPort {
		t.Errorf("Port is unexpected: %s", serverCfg.Port)
	}

	if serverCfg.Dump != expectedDumpPath {
		t.Errorf("Dump path is unexpected: %s", serverCfg.Dump)
	}

	if serverCfg.Crt != expectedCrtPath {
		t.Errorf("SSL crt path is unexpected: %s", serverCfg.Crt)
	}

	if serverCfg.Key != expectedKeyPath {
		t.Errorf("SSL key path is unexpected: %s", serverCfg.Key)
	}

	if serverCfg.TemplateCfg.Index != expectedIndexTemplatePath {
		t.Errorf("Index template path is unexpected: %s", serverCfg.TemplateCfg.Index)
	}

	if serverCfg.Log != expectedLogPath {
		t.Errorf("Log path is unexpected: %s", serverCfg.Log)
	}
}
