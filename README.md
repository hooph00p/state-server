# @hooph00p

## Plan

- [x] Get JSON Loaded into memory
- [x] Add some Tests with Contains, Doesn't Contain and fringe cases ("On-the-line")
- [x] Create REST Endpoint that accepts a POST request with a longitude and latitude argument
- [x] Add Insomnia tests to make sure the endpoint works
- [x] Migrates tests from Insomnia to main_test.go
- [x] Makefile for dependencies

## Dependencies:

- [gin](http://github.com/gin-gonic/gin), A Web Framework in Go
- [golang-geo](https://github.com/kellydunn/golang-geo/), A Math Library to help with Polygon Contains
