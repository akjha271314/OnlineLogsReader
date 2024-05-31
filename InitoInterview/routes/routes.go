package routes

import (
	"net/http"
	"onlineLogsReader/controllers"

	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {
	logsReadeController := controllers.NewLogsReaderController()
	router.HandleFunc("/read_logs", logsReadeController.ReadLogs).Methods("GET")
	http.Handle("/", router)
}
