package common

import (
	"gosync/logger"
	"gosync/utils"
	"os"
	"path"
)

const AppName = "gosync"

var RunningHome = ""

func ResolveRunningHome(specifiedPath string) {
	if specifiedPath == "" {
		// use the default path [${user-home}/.gosync]
		userHomeDir, err := os.UserHomeDir()
		if err != nil {
			logger.Fatal("Fail to get the user home dir!", err)
		}
		RunningHome = path.Join(userHomeDir, "."+AppName)
	} else {
		RunningHome = specifiedPath
	}
	// ensure the dir exists
	err := utils.EnsureDir(RunningHome)
	if err != nil {
		logger.Fatal("Fail to check the running home dir!", err)
	}
	logger.Infoln("Finally setting running home ->", RunningHome)
}
