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
	res, err := json.Marshal(isps)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(res))
}

func (s *Server) HandleCountry(w http.ResponseWriter, r *http.Request) {
	// Amount of IPs per country
	vars := mux.Vars(r)
	countryCode := vars["cc"]
	ipCount := s.DB.GetIPCountInCountry(countryCode)
	fmt.Fprintf(w, fmt.Sprint(ipCount))
}

func (s *Server) HandleIP(w http.ResponseWriter, r *http.Request) {
	// All data for a given IP
	vars := mux.Vars(r)
	ip := net.ParseIP(vars["ip"])
	converted := ipconv.Ip2int(ip)
	ipinfo := s.DB.GetDataForIP(converted)
	res, err := json.Marshal(ipinfo)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(res))
}
