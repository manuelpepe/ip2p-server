# IP2Proxy API

A simple Go API to query data from the [IP2Proxy](https://lite.ip2location.com/database/px7-ip-proxytype-country-region-city-isp-domain-usagetype-asn) database

This is the solution to the Dreamlab Challenge

## Challenge (simplified)

1. Download [IP2Proxy IPv4](https://lite.ip2location.com/database/px7-ip-proxytype-country-region-city-isp-domain-usagetype-asn) data.
2. Populate a PostgreSQL database with the downloaded data.
3. Build a Go API with JSON endpoints for:

    1. TOP 10 ISP from Switzerland (Country code: CH)
    2. Number of IPs per country (Country code given as a parameter)
    3. All available information from a given IP Address (IP Address given as a parameter)

4. Add Unit Tests and Documentation
5. Extra points for Docker/Docker-Compose


### Assumptions

* For endpoint 1, I assume that only the ISP Name and the total amount of IPs in the ISP is enough data.
* For endpoint 2, only the number of IPs in the country is retrieved as an integer(which is valid JSON).
* For endpoint 3, I decided to include the ip_from and ip_to data to the returned JSON. Originally I was leaving it at that,
but later decided to add a list of all the IPs in IPv4 format to try out struct embedding and make use of the `Int2ip` function. It could be argued that the other IPs in the block are not necessarily information about the queried IP, but it does make it more fun.

## Endpoints

* **/isp**: TOP 10 ISP from Switzerland (ISO-3166: CH)

```
$ curl -s localhost:8888/isp | jq
[
  {
    "name": "Ivan Bulavkin",
    "total_ips": 160
  },
  {
    "name": "Sunrise GmbH",
    "total_ips": 134
  },
  {
    "name": "Private Layer Inc",
    "total_ips": 134
  },
  {
    "name": "Swisscom AG",
    "total_ips": 131
  },
  {
    "name": "Oracle Public Cloud",
    "total_ips": 122
  },
  {
    "name": "Cloud Innovation Ltd",
    "total_ips": 112
  },
  {
    "name": "Microsoft Corporation",
    "total_ips": 101
  },
  {
    "name": "RapidSeedbox Ltd",
    "total_ips": 93
  },
  {
    "name": "Rosite Equipment SRL",
    "total_ips": 84
  },
  {
    "name": "Hop Bilisim Teknolojileri Anonim Sirketi",
    "total_ips": 77
  }
]
```

* **/country/{countryCode}**: Number of IPs per country (Country code given as a parameter)
```
$ curl -s localhost:8888/country/AR | jq
79346
$ curl -s localhost:8888/country/AS | jq
0
```

* **/ip/{ipv4Address}**: All available information from a given IP Address
```
$ curl -s localhost:8888/ip/23.227.38.83 | jq
{
  "ip": "23.227.38.83",
  "ip_block_from": 400762449,
  "ip_block_to": 400762457,
  "proxy_type": "PUB",
  "country_code": "CA",
  "country_name": "Canada",
  "region_name": "Ontario",
  "city_name": "Ottawa",
  "isp": "Shopify Inc.",
  "domain": "shopify.com",
  "usage_type": "DCH",
  "asn": "13335",
  "as": "CloudFlare Inc.",
  "in_block_with": [
    "23.227.38.81",
    "23.227.38.82",
    "23.227.38.83",
    "23.227.38.84",
    "23.227.38.85",
    "23.227.38.86",
    "23.227.38.87",
    "23.227.38.88",
    "23.227.38.89"
  ]
}
$ curl -s localhost:8888/ip/1.1.1.256 | jq
"Wrong IP format"
$ curl -s localhost:8888/ip/1.1 | jq
"Wrong IP format"
$ curl -s localhost:8888/ip/1.1.1.1 | jq
"IP not found"
```
## Running with docker-compose

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
docker exec -it ip2p_postgres psql -U ip2p -d ip2p 
# run the copy command
COPY ip2location_px7 FROM '/root/IP2PROXY-LITE-PX7.CSV' WITH CSV QUOTE AS '"';
```

After a few seconds the database should be populated and the number of inserted rows should be displayed in the
terminal (currently 2317954 rows) and you can exit the psql prompt with the `exit` command.

Finally you can curl the endpoints specified in the Endpoints section:

``` 
curl -s localhost:8888/isp | jq
curl -s localhost:8888/ip/23.227.38.83 | jq
curl -s localhost:8888/country/AR | jq
``` 

## Running Unit Tests

Unit tests don't require a database connection so you can just run

```
go test -v ./...
```

in the project's root directory.
