package util

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func GetExecFilePath() string {
	path, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err)
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	index := strings.LastIndex(absPath, string(os.PathSeparator))
	resultPath := absPath[:index]
	return resultPath
}
