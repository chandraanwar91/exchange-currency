# Currency Exchange Application

This repository contains a forex application using golang programming language.

# Installation
```shell
go get github.com/chandraanwar91/exchange-currency
```
Copy file configuration
Copy File config.example.yaml to config.yaml and dbconf.example.yml to dbconf.yml in folder db

# Usage and Demo

Create the Docker image according to [Dockerfile](Dockerfile).

```shell
# This may take a few minutes.
$ docker-compose up
```

# Notes
For database I use bitbucket.org/liamstask/goose/src/master/ for migration

```shell
# run migration 
$ goose up
```

For API Documentation I use swagger

```shell
# run swagger 
$ <base_url>/swagger/index.html
```
