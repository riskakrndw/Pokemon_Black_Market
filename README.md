# Pokemon Black Market

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)
[![Go.Dev reference](https://img.shields.io/badge/gorm-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/gorm.io/gorm?tab=doc)
[![Go.Dev reference](https://img.shields.io/badge/echo-reference-blue?logo=go&logoColor=white)](https://github.com/labstack/echo)

## Table of Content

  - [About](#about)
    - [Features](#features)
    - [API Documentation](#api-documentation)
  - [System Design](#system-design)
  - [Getting started](#getting-started)
    - [Installing](#installing)


## About

Pokemon Black Market providing REST API for processing and reporting transaction data of Pokemon. There are 3 types of user in this application: 'Bos', 'Operasional' and 'Pengedar'

### Features

- [x] Register
- [x] Login
- [x] Profile
- [x] Logout
- [x] Search all pokemon from pokeapi.co
- [x] Search pokemon by pokemon's name from pokeapi.co
- [x] Create pokemon
- [x] Update pokemon
- [x] Delete pokemon
- [x] Get all pokemons
- [x] Search pokemon by id
- [x] Search pokemons by name
- [x] Create transaction
- [x] Cancel transaction
- [x] Get transaction by id
- [x] Get all success transaction
- [x] Get all cancelled transaction

### API Documentation

Application Programming Interface is available at [API page.](API.md)

## System Design

ERD
![Pokemon Black Market-ERD](https://user-images.githubusercontent.com/11256042/137578139-d968c6ff-2c42-4829-816a-4cb49325caf3.jpg)


Use Case
![Pokemon Black Market-Use Case](https://user-images.githubusercontent.com/11256042/137578144-53aa21b3-5ddd-4c50-92c8-262fbc9dd6eb.jpg)


## Getting Started

Below we describe how to start this project

### Installing

You must download and install `Go`, follow [this instruction](https://golang.org/doc/install) to install.

After Golang installed, Follow this instructions
```bash
$ git clone https://github.com/riskakrndw/Pokemon_Black_Market.git
$ go run main.go
```

Go to `http://localhost:8000/` to [start this application.](http://localhost:8000/)
