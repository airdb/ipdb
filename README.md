# IPDB

IP Database service refer to the collection, analysis, and provision of data related to Internet Protocol (IP) addresses. This information can be used for a variety of purposes, such as improving website analytics, identifying and preventing fraudulent activities, personalizing user experiences, and more.

For example, few provide can offers GeoIP2 data services and databases that cover 99.9999% of IP addresses in use and are updated weekly.

## Build and Run

### Step 1: Purchase and dowload the IP dataset or IP database from IPIP, IPPlus360, MaxMind, IP2location, etc

- [IPIP](https://www.ipip.net/product/ip.html)
- [IPPlus360](https://www.ipplus360.com/)
- [MaxMind](https://www.maxmind.com/en/home)
- [IP2location](https://www.ip2location.com/)

### Step 2: Unzip IP Database and setup the config file

- Save the file to ./data directory
- Update the config.yaml file

### Step 3: Build ipdb service

```bash
git clone https://github.com/airdb/ipdb

cd ipdb

make build

make start
```

### Step 4: Use the IPdb service

```bash
$ curl 127.0.0.1:8080/

==== Welcome to ipdb service ====

Usage: visit or curl the url directly, then you will get the ip information.

IP2Location: https://www.ip2location.com/
  GET /v1/ip2location/
  GET /v1/ip2location/{ip}

IPIP: https://www.ipip.net/
  GET /v1/ipip/
  GET /v1/ipip/{ip}

MaxMind: https://www.maxmind.com/
  GET /v1/maxmind/
  GET /v1/maxmind/{ip}

IPPlus360: https://www.ipplus360.com/
  GET /v1/ipplus360/
  GET /v1/ipplus360/{ip}


Thank you for following us!!
https://github.com/airdb
```

```bash

```bash

## Reference

### IPDB Database API Document

[![IPDB Database API Document](https://godoc.org/github.com/ipipdotnet/ipdb-go?status.svg)](https://godoc.org/github.com/ipipdotnet/ipdb-go)

### IP2Location Database API Document

[![IP2Location Database API Document](https://godoc.org/github.com/ipipdotnet/ipdb-go/ip2location?status.svg)](https://godoc.org/github.com/ipipdotnet/ipdb-go/ip2location)
