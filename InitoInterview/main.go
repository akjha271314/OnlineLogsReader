package main

import (
	"fmt"
	"net/http"
	"onlineLogsReader/routes"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting web server at :8080")
	r := mux.NewRouter()
	routes.Init(r)

	http.ListenAndServe(":8080", nil)
}
