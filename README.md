# Coursera GO

This is a simple repository to store assignments for the 
* [Getting Started with Go](https://www.coursera.org/learn/golang-getting-started) course on Coursera.
* [Functions, Methods, and Interfaces in Go](https://www.coursera.org/learn/golang-functions-methods) course on Coursera.

## Getting Started

### Prerequisites

* [Go](https://golang.org/dl/) - The Go Programming Language

## Personal Notes / Cheatsheet

### To Compile

go build -o ./bin/HelloWorld ./source/HelloWorld.go

### To Run

./bin/HelloWorld

### To run without compiling first

go run ./source/HelloWorld.go

### To run tests

go test ./source/HelloWorld.go ./source/HelloWorld_test.go

### To run tests with coverage

go test -cover ./source/HelloWorld.go ./source/HelloWorld_test.go

### To run tests with coverage and generate html report

go test -coverprofile=coverage.out ./source/HelloWorld.go ./source/HelloWorld_test.go