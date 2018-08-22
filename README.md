# Currency Exchange Application

This repository contains a forex application using golang programming language.

# Installation
```shell
go get github.com/chandraanwar91/exchange-currency
```

# Usage and Demo

Create the Docker image according to [Dockerfile](Dockerfile).
This step uses Maven to build, test, and package the [Go application](app.go).
The resulting image is 7MB in size.

```shell
# This may take a few minutes.
$ docker-compose up
```

# Notes

For API Documentation I use swagger

```shell
# run swagger 
$ <base_url>/swagger/index.html
```
