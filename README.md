# get-ze-weather 

## Introductions 
Hi folx! My name's Kent, and I like computers. This repo is a pretty simple little app that gets the weather for a given city. 

## Installation
- Download and install the go binary for your operating system from [here](https://golang.org/dl/)
- Clone the repo into `$GOPATH/src` (`echo $GOPATH` after installing the binary
- From the project root: `go get ./...`
- From the project root: `go build` 
- `export OPENWEATHERMAPAPIKEY=<api-key>`. You can get an api key from the openweathermap website for free.
- `get-ze-weather` 

## Notes
- I'm a rookie Go developer! In deciding what language to use for this project, I wanted something with the least amount of bootstrapping required, and Go was the perfect fit. All the libraries used in this project are native libraries.
- I'm inexperienced using Go's funky default testing framework. In the course of this project, I learned it from scratch; previously, I've used the Ginkgo and Gomega frameworks. I was considering using them as well, but I felt that they would overcomplicate the tests.
- There are some "idiomatic" things in here that are untested. These are `convertHTTPResponseToByteArray` and `promptUserForInput`. I don't like leaving them untested, but I also tried to keep them as small as possible and isolated away from the actual logic.
- The only exported function from the `models` package has a comment above it. This is a by-product of a go package that I use through Atom that does some aggressive style-checking. Don't interpret it as an endorsement or denouncement of comments. 
- In general, the code assumes a relatively happy path. My error-handling in Go is not particularly sophisticated, so there's a few hard exits. The code does assume that you're going to input a valid city in a somewhat expected format, e.g. `City,State` or `City Country`. 
- I try to be explicit with variable names, but you might see something glib or silly. I'm sorry that my brain loves the word "response" so much.

## Getting out in front of potential questions
- Can I run this on Windows? _Please don't. I had to write code on Windows towards the end of this exercise and it was awful._
