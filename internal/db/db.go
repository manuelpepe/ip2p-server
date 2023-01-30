package db

import (
	"os"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

type IPInfo struct {
	Ip uint32
	ProxyType string
	CountryCode string
	CountryName string
	RegionName string
	CityName string
	ISP string
	Domain string
	UsageType string
	ASN int
	AS string
}

func ConnectDB() (*sql.DB, error) {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@localhost/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return &sql.DB{}, err
	}
	
	// Test DB connection
	_, err = db.Query("SELECT 1")
	if err != nil {
		return &sql.DB{}, err
	}
	return db, nil
	
}


func GetTopISPsInCountry(countryCode string, count int) []string {
	return []string{"Some ISP", "other"}
}

func GetIPCountInCountry(countryCode string) uint {
	return 1
}

func GetDataForIP(ip int) *IPInfo {
	return &IPInfo{}
}
