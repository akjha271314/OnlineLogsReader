package controllers

import (
	"encoding/json"
	"net/http"
	"onlineLogsReader/services"
)

type logsReaderController struct {
	service services.LogsReaderServiceI
}

type LogsReaderControllerI interface {
	ReadLogs(w http.ResponseWriter, r *http.Request)
}

func NewLogsReaderController() LogsReaderControllerI {
	return &logsReaderController{
		service: services.NewLogsReaderService(),
	}
}

func (c *logsReaderController) ReadLogs(w http.ResponseWriter, r *http.Request) {

	logs, err := c.service.ReadLogs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Encode the response struct to JSON
	jsonData, err := json.Marshal(logs)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	// Set the Content-Type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.Write(jsonData)
}
