package handlers

import (
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestCountryHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/country/SC", nil)
    w := httptest.NewRecorder()
    server := Server{}
    server.HandleCountry(w, req)
    res := w.Result()
    defer res.Body.Close()
    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        t.Errorf("expected error to be nil got %v", err)
    }
    if string(data) != "1" {
        t.Errorf("expected 1 got %v", string(data))
    }

}