package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"time"
)


const (
	INFO  = "INFO"
	WARN  = "WARN"
	DEBUG = "DEBUG"
	ERROR = "ERROR"
)

func getLogFilePath() string {
	if runtime.GOOS == "windows" {
		return "C:\\logs\\app.log" 
	}
	return "/var/log/app.log" 
}

func logMessage(level, message string) {
	logMsg := fmt.Sprintf("[%s] %s: %s", level, time.Now().Format(time.RFC3339), message)
	fmt.Println(logMsg) 
	log.Println(logMsg) 
}

func main() {
	logFilePath := getLogFilePath()

	logDir := logFilePath[:len(logFilePath)-8] 
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

	logFile, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer logFile.Close()

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)

	fmt.Println("✅ Golang Logging Service Running...")

	for {
		logMessage(INFO, "Application is running")
		logMessage(WARN, "This is a warning message")
		logMessage(DEBUG, "Debugging application state")
		logMessage(ERROR, "An error occurred in the application")
		time.Sleep(5 * time.Second) 
	}
}
