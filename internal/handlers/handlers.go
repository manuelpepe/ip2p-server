package handlers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/manuelpepe/ip2p-server/internal/db"
	"github.com/manuelpepe/ip2p-server/internal/ipconv"
)

type Server struct {
	DB db.IDB
}

func (s *Server) HandleISP(w http.ResponseWriter, r *http.Request) {
	// Top 10 ISP in Switzerland
	isps := s.DB.GetTopISPsInCountry("CH", 10)
	respondWithJson(w, isps)
}

func (s *Server) HandleCountry(w http.ResponseWriter, r *http.Request) {
	// Amount of IPs per country
	vars := mux.Vars(r)
	countryCode := vars["cc"]
	ipCount := s.DB.GetIPCountInCountry(countryCode)
	respondWithJson(w, ipCount)
}

func (s *Server) HandleIP(w http.ResponseWriter, r *http.Request) {
	// All data for a given IP
	vars := mux.Vars(r)
	ip := net.ParseIP(vars["ip"])
	converted := ipconv.Ip2int(ip)
	ipinfo, err := s.DB.GetDataForIP(converted)
	if err != nil {
		respondWithJson(w, err.Error())
		return
	}
	respondWithJson(w, ipinfo)
}

func respondWithJson(w http.ResponseWriter, v any) {
	res, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(res))
}
