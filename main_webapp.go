package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"taxi-tracker-webapp/controller"
	"time"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", controller.Index).Methods("GET")
	r.HandleFunc("/login", controller.Login).Methods("POST")

	// Static files
	staticPath := os.Getenv("GOPATH") + "/src/taxi-tracker-webapp/static/"
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticPath))))

	cssPath := os.Getenv("GOPATH") + "/src/taxi-tracker-webapp/css/"
	r.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(cssPath))))

	fontsPath := os.Getenv("GOPATH") + "/src/taxi-tracker-webapp/fonts/"
	r.PathPrefix("/fonts/").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir(fontsPath))))

	imgPath := os.Getenv("GOPATH") + "/src/taxi-tracker-webapp/img/"
	r.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir(imgPath))))

	jsPath := os.Getenv("GOPATH") + "/src/taxi-tracker-webapp/js/"
	r.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(jsPath))))


	server := &http.Server{
		Addr:           ":3000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Listening on port 3000")
	server.ListenAndServe()
}
