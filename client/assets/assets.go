package assets

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const (
	SliverClientDirName = ".sliver-client"
	versionFileName     = "version"
)

func GetRootAppDir() string {
	user, _ := user.Current()
	dir := filepath.Join(user.HomeDir, SliverClientDirName)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0700)
		if err != nil {
			log.Fatal(err)
		}
	}
	return dir
}

func GetClientLogsDir() string {
	logsDir := filepath.Join(GetRootAppDir(), "logs")
	if _, err := os.Stat(logsDir); os.IsNotExist(err) {
		err = os.MkdirAll(logsDir, 0700)
		if err != nil {
			log.Fatal(err)
		}
	}
	return logsDir
}

func GetConsoleLogsDir() string {
	consoleLogsDir := filepath.Join(GetClientLogsDir(), "console")
	if _, err := os.Stat(consoleLogsDir); os.IsNotExist(err) {
		err = os.MkdirAll(consoleLogsDir, 0700)
		if err != nil {
			log.Fatal(err)
		}
	}
	return consoleLogsDir
}

// 获取版本信息
func assetVersion() string {
	appDir := GetRootAppDir()
	data, err := os.ReadFile(filepath.Join(appDir, versionFileName))
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(data)) //去除两端空格
}
