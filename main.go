package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

const (
	LogFilePath  = "time.log"
	LogDirectory = "./"
)

func getFullPath(dirName, fileName string) (string, error) {
	dirName, err := filepath.Abs(dirName)
	if err != nil {
		return "", err
	}

	path := filepath.Join(dirName, fileName)
	return path, nil
}

func main() {
	logPath, err := getFullPath(LogDirectory, LogFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer func() {
		err = logFile.Close()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}()
	logger := log.New(logFile, "", log.Ltime)

	cmd := exec.Command("pmset", "disablesleep", "1")
	err = cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for {
		logger.Println("")
		time.Sleep(10 * time.Second)
	}
}
