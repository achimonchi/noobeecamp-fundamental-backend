package main

import (
	"fmt"
	"nbcamp_project/server"
	"nbcamp_project/server/config"

	"github.com/NooBeeID/go-logging/messages"
	"github.com/gorilla/mux"
)

func main() {
	run()
}

func run() {
	fmt.Println("Starting...")
	db := config.StartPostgres()
	fmt.Println("connect DB success")
	if db == nil {
		messages.SysError("error connection")
	}

	router := mux.NewRouter()

	port := ":8890"

	server.StartServer(router, db, port)
}
