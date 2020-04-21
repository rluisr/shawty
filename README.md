[![Build Status](https://cloud.drone.io/api/badges/rluisr/shawty/status.svg)](https://cloud.drone.io/rluisr/shawty)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/didip/shawty/master/LICENSE)

## Shawty: URL Shortener Service

This service encodes URL in base-36 and store them in filesystem.

It has 3 features: shorten, unshorten, and redirect.


### Can I use it in production?

You need to implement a storage that can scale beyond one application server.

Need these environment variables.

#### Redis
- `REDIS_ADDR`
- `REDIS_PASSWORD`
- `REDIS_DB`

If you have redis sentinel, set `REDIS_SENTINEL_ADDR`, `REDIS_SENTINEL_MASTER_NAME`,`REDIS_PASSWORD` and `REDIS_DB`.

#### REDIS_SENTINEL_MASTER_ADDR
split with `,` like `sentinel01.local,sentinel02.local`

#### Random string size
- `GENERATE_SIZE`


### Why?

By itself, URL shortening is quite useful.

But this project exists to demonstrate:

* How concise [Go](http://golang.org/) is. [cloc](http://cloc.sourceforge.net/) shows that this project contains only 125 lines.

* How slim Go is: 3MB RAM.

* How comprehensive Go standard library is.

* How easy it is to get up and running in Go. It took me about 1 hour from start to finish. Writing this README file took longer time.

* How performant Go is:
    ```
    # Command  : ab -n 100000 -c 200 -k http://localhost:8080/dec/1
    # Processor: 2.26 GHz Intel Core 2 Duo  <-- Crummy 6 years old laptop

    Concurrency Level:      200
    Time taken for tests:   8.610 seconds
    Complete requests:      100000
    Failed requests:        0
    Non-2xx responses:      100000
    Keep-Alive requests:    100000
    Total transferred:      22400000 bytes
    HTML transferred:       7600000 bytes
    Requests per second:    11614.80 [#/sec] (mean)
    Time per request:       17.219 [ms] (mean)
    Time per request:       0.086 [ms] (mean, across all concurrent requests)
    Transfer rate:          2540.74 [Kbytes/sec] received
    ```
