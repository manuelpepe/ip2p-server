package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/manuelpepe/ip2p-server/internal/db"
	"github.com/manuelpepe/ip2p-server/internal/ipconv"
)

var ExpectedISPInfo = []db.ISPInfo{
	{Name: "ISP-A", TotalIPs: 10},
	{Name: "ISP-B", TotalIPs: 23},
	{Name: "ISP-C", TotalIPs: 13},
}

var ExpectedCountryInfo = 14012

var ExpectedIPInfo = &db.IPInfo{
	IPBlockFrom: 1000001,
	IPBlockTo:   1000002,
	ProxyType:   "proxy_type",
	CountryCode: "SC",
	CountryName: "some_country",
	RegionName:  "some_region",
	CityName:    "some_city",
	ISP:         "some_isp",
	Domain:      "some_domain",
	UsageType:   "usage_type",
	ASN:         "some_asn",
	AS:          "some_as",
}

type MockDB struct {
}

func (m *MockDB) GetTopISPsInCountry(string, int) []db.ISPInfo {
	return ExpectedISPInfo
}

func (m *MockDB) GetIPCountInCountry(countryCode string) uint {
	return uint(ExpectedCountryInfo)
}

func (m *MockDB) GetDataForIP(ip uint32) (*db.IPInfo, error) {
	return ExpectedIPInfo, nil
}

func TestCountryHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/country/SC", nil)
	w := httptest.NewRecorder()
	server := Server{DB: &MockDB{}}
	server.HandleCountry(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != fmt.Sprint(ExpectedCountryInfo) {
		t.Errorf("expected 14012 got %v", string(data))
	}
}

func TestISPHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/isp/SC", nil)
	w := httptest.NewRecorder()
	server := Server{DB: &MockDB{}}
	server.HandleISP(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	dataExpected, err := json.Marshal(ExpectedISPInfo)
	if err != nil {
		panic(err)
	}
	if string(data) != string(dataExpected) {
		t.Errorf("expected %s got %s", dataExpected, data)
	}
}

func TestIPHandler(t *testing.T) {
	type TC struct {
		arg string
		exp any // expected value will be marshaled
	}

	var tests = []TC{
		{"23.227.38.83", ExpectedIPInfo},
		{"23.1", ipconv.WrongIPFormatError.Error()},
		{"23.227.38.83.123", ipconv.WrongIPFormatError.Error()},
		{"255.255.255.256", ipconv.WrongIPFormatError.Error()},
	}

	for _, test := range tests {
		req := httptest.NewRequest(http.MethodGet, "/ip/"+test.arg, nil)
		// Required to manually set Vars as Router is not used (gorilla/mux/issues/373)
		req = mux.SetURLVars(req, map[string]string{
			"ip": test.arg,
		})
		w := httptest.NewRecorder()
		server := Server{DB: &MockDB{}}
		server.HandleIP(w, req)
		res := w.Result()
		defer res.Body.Close()
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}
		dataExpected, err := json.Marshal(test.exp)
		if err != nil {
			panic(err)
		}
		if string(data) != string(dataExpected) {
			t.Errorf("expected %s got %s", dataExpected, data)
		}
	}
}
