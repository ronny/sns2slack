package main

import (
	"github.com/gorilla/mux"
	"log/syslog"
	"net/http"
	"os"
	"fmt"
)

var (
	log *syslog.Writer
)

func main() {
	var err error
	log, err = syslog.New(syslog.LOG_INFO, "sns2slack")
	if err != nil {
		panic(err)
	}
	defer log.Close()

	router := mux.NewRouter()
	router.HandleFunc(HookHandlerUrl, hookHandler).Methods("POST")
	http.Handle("/", router)

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}
	log.Info("Listening on :"+port+"...")
	fmt.Println("Listening on :"+port+"...")
	http.ListenAndServe(":"+port, nil)
}
