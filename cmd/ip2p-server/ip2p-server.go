package main

import (
	"log"
	"net/http"

	"github.com/manuelpepe/ip2p-server/internal/db"
	"github.com/manuelpepe/ip2p-server/internal/handlers"

	"github.com/gorilla/mux"
)

func main() {
	db_, err := db.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db_.DB.Close()
	server := handlers.Server{DB: db_}
	router := mux.NewRouter()
	router.HandleFunc("/isp", server.HandleISP)
	router.HandleFunc("/country/{cc}", server.HandleCountry)
	router.HandleFunc("/ip/{ip}", server.HandleIP)
	log.Fatal(http.ListenAndServe(":8888", router))
}
