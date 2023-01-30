package db

import (
	"os"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

type DB struct {
	DB *sql.DB
}


func ConnectDB() (*DB, error) {
	connStr := fmt.Sprintf(
		"postgresql://%s:%s@localhost/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return &DB{}, err
	}
	
	// Test DB connection
	_, err = db.Query("SELECT 1")
	if err != nil {
		return &DB{}, err
	}
	return &DB{DB: db}, nil
	
}

type IPInfo struct {
	IPBlockFrom uint32		`json:"ip_block_from"`
	IPBlockTo 	uint32		`json:"ip_block_to"`
	ProxyType 	string 		`json:"proxy_type"`
	CountryCode string		`json:"country_code"`
	CountryName string		`json:"country_name"`
	RegionName 	string		`json:"region_name"`
	CityName 	string		`json:"city_name"`
	ISP 		string		`json:"isp"`
	Domain 		string		`json:"domain"`
	UsageType 	string		`json:"usage_type"`
	ASN 		int			`json:"asn"`
	AS 			string		`json:"as"`
}

type ISPInfo struct {
	Name string `json:"name"`
	TotalIPs int `json:"total_ips"`
}

func (db *DB) GetTopISPsInCountry(countryCode string, count int) []ISPInfo {
	rows, err := db.DB.Query(
		`SELECT isp, SUM(ip_to - ip_from + 1) total_ips
		FROM ip2location_px7
		WHERE country_code = $1
		GROUP BY isp
		ORDER BY total_ips DESC
		LIMIT $2`, countryCode, count)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var (
		isps []ISPInfo
		isp string
		usageCount int
	)
	for rows.Next() {
		err := rows.Scan(&isp, &usageCount)
		if err != nil {
			panic(err)
		}
		isps = append(isps, ISPInfo{isp, usageCount})
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	return isps
}

func (db *DB) GetIPCountInCountry(countryCode string) uint {
	var total_ips uint
	err := db.DB.QueryRow(
		`SELECT SUM(ip_to - ip_from + 1) total_ips
		FROM ip2location_px7
		WHERE country_code = $1`, countryCode).Scan(&total_ips)
	if err != nil {
		panic(err)
	}
	return total_ips
}

func (db *DB) GetDataForIP(ip uint32) *IPInfo {
	var ipinfo IPInfo
	err := db.DB.QueryRow(
		`SELECT *
		FROM ip2location_px7
		WHERE ip_from = $1
			OR ip_to = $1
			OR (ip_from < $1 and ip_to > $1)`, ip).Scan(
				&ipinfo.IPBlockFrom,
				&ipinfo.IPBlockTo,
				&ipinfo.ProxyType,
				&ipinfo.CountryCode,
				&ipinfo.CountryName,
				&ipinfo.RegionName,
				&ipinfo.CityName,
				&ipinfo.ISP,
				&ipinfo.Domain,
				&ipinfo.UsageType,
				&ipinfo.ASN,
				&ipinfo.AS,
			)
	if err != nil {
		panic(err)
	}
	return &ipinfo
}
