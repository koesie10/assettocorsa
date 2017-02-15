# assettocorsa

## Description
This program will list the cars on an Assetto Corsa server 
based on the host and port of an Assetto Corsa server.

## Usage
```bash
$ assettocorsa status --host localhost --port 8081
```

The `port` parameters is the HTTP port in the server configuration.

## Install

### Requirements

* [Glide](https://github.com/Masterminds/glide)

To install, use `go get` and Glide:

```bash
$ go get -d github.com/koesie10/assettocorsa
$ cd $GOPATH/src/github.com/koesie10/assettocorsa
$ glide install
$ go install github.com/koesie10/assettocorsa
```

## Contribution

1. Fork ([https://github.com/koesie10/assettocorsa/fork](https://github.com/koesie10/assettocorsa/fork))
1. Create a feature branch
1. Commit your changes
1. Rebase your local changes against the master branch
1. Run `gofmt -s`
1. Create a new Pull Request

## Author

[koesie10](https://github.com/koesie10)
