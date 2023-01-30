package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/manuelpepe/ip2p-server/internal/db"
)

type Server struct {
	DB *sql.DB
}


func (s *Server) HandleISP(w http.ResponseWriter, r *http.Request) {
	// Top 10 ISP in Switzerland
	// isps := db.GetTopISPsInCountry(countryCode, 10)
	fmt.Fprintf(w, "handle isp")
}


func (s *Server) HandleCountry(w http.ResponseWriter, r *http.Request) {
	// Amount of IPs per country
	vars := mux.Vars(r)
	countryCode := vars["cc"]
	ipCount := db.GetIPCountInCountry(countryCode)
	fmt.Fprintf(w, fmt.Sprint(ipCount))
}


func (s *Server) HandleIP(w http.ResponseWriter, r *http.Request) {
	// All data for a given IP
	vars := mux.Vars(r)
	ip := vars["ip"]
	fmt.Fprintf(w, ip)
}