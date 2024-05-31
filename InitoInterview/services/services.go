package services

import (
	"fmt"
	"os"
	"time"
)

type logsReaderService struct{}

type LogsReaderServiceI interface {
	ReadLogs() (string, error)
	GenerateLogs() (string, error)
}

func NewLogsReaderService() LogsReaderServiceI {
	return &logsReaderService{}
}

func (s *logsReaderService) ReadLogs() (string, error) {
	logs, err := s.GenerateLogs()
	if err != nil {
		return "", err
	}
	return logs, nil
}

func (s *logsReaderService) GenerateLogs() (string, error) {

	file, err := os.OpenFile("./logs.txt", os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		fmt.Println("Failed to open file ", err)
		file, err = os.Create("logs.txt")
		if err != nil {
			fmt.Println("Failed to create file:", err)
			return "", err
		}
	}

	logs := time.Now().String()
	_, err = file.WriteString(logs + "\r\n")
	if err != nil {
		fmt.Println("Failed to write to file:", err) //print the failed message
		return "", err
	}

	defer file.Close()

	return logs, nil
}
