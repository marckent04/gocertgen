# GoCertGen

GoCertGen is a CLI app which take in input a csv file and generate 
completion course certificates to HTML or PDF format


## Features
- [X] get and parse csv file 
- [X] generate html certificates
- [X] generate csv certificates
- [X] CLI compatibility
- [ ] unit tests
- [ ] e2e tests

## Requirements
have go installed and a CLI

## How to install
- go install

## How to use
run `go run main.go` with following args:
- `-format`: generated certificates format ( html / pdf )
- `-path`: csv file path
-  `-output`: output directory (default: outputs)
- `-template`: html template path

NB: we can see templates examples in the **template-examples** folder

## About me
Marc-Henry Nanguy, JS/TS/DART/GO developer and software craftsman