# IP2Proxy API

A simple Go API to query data from the [IP2Proxy](https://lite.ip2location.com/database/px7-ip-proxytype-country-region-city-isp-domain-usagetype-asn) database

This is the solution to the Dreamlab Challenge

## Challenge (simplified)

1. Download [IP2Proxy IPv4](https://lite.ip2location.com/database/px7-ip-proxytype-country-region-city-isp-domain-usagetype-asn) data.
2. Populate a PostgreSQL database with the downloaded data.
3. Build a Go API with endpoints for:

    1. TOP 10 ISP from Switzerland (Country code: CH)
    2. Number of IPs per country (Country code given as a parameter)
    3. All available information from a given IP Address (IP Address given as a parameter)

4. Add Unit Tests and Documentation
5. Extra points for Docker/Docker-Compose

## How to Run

To start Docker containers run:

```
docker-compose up
```

After that, you'll need to populate the Database with the IP2Proxy data.
Download the file from [here](https://lite.ip2location.com/database/px7-ip-proxytype-country-region-city-isp-domain-usagetype-asn) and populate it with the following commands:

```
# cp/mv the file to the mounted volume
mv ~/Downloads/IP2PROXY-LITE-PX7.CSV ./pgdata
# connect to the database
docker exec -it ip2p_dl_challenge-postgres-1 psql -U ip2p -d ip2p 
# run the copy command
COPY ip2location_px7 FROM '/root/IP2PROXY-LITE-PX7.CSV' WITH CSV QUOTE AS '"';
```

After a few seconds the database should be populated and the number of inserted rows should be displayed in the
terminal (currently 2317954 rows) and you can exit the psql prompt with the `exit` command.
