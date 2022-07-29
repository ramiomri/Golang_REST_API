package main

import (
	"example/golang-rest-api/middleware"
	"example/golang-rest-api/router"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)



func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	r := mux.NewRouter()
	r.Use(middleware.Header)
	router.RegisterRouter(r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "localhost:" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Println("Server started on port " + port)
	log.Fatal(srv.ListenAndServe())
}
