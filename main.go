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

func printTimeLog(dir, file string) {
	logPath, err := getFullPath(dir, file)
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
	for {
		logger.Println("")
		time.Sleep(10 * time.Second)
	}
}

func disableSleep() error {
	cmd := exec.Command("pmset", "disablesleep", "1")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func enableSleep() {
	cmd := exec.Command("pmset", "disablesleep", "0")
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func main() {
	err := disableSleep()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer enableSleep()

	printTimeLog(LogDirectory, LogFilePath)
}
