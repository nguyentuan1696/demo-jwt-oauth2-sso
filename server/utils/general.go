package utils

// Find the path name for the current directory.

import (
	"log"
	"os"
)

func CurrentDirectoryPathName() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

const (
	LogInfoColor   = "\033[1;34m%s\033[0m"
	LogNoticeColor = "\033[1;36m%s\033[0m"
	LogWaringColor = "\033[1;33m%s\033[0m"
	LogErrorColor  = "\033[1;31m%s\033[0m"
	LogDebugColor  = "\033[0;36m%s\033[0m"
)
