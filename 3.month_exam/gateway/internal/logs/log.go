package logs

import (
	"log"
	"os"
)


func GetLogger(filePath string) *log.Logger {
	logFile, err := os.OpenFile(filePath, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatalf("Failed to create log file: %v", err)
	}
	
	return log.New(logFile, "", log.LstdFlags)
}
