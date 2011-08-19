## Installing Go

On Mac, run `brew install go --use-git`. Other package managers should work similarly.

If your package manager doesn't have Go, it can be installed via [these directions](http://golang.org/doc/install.html).

If using MacPorts run `port install go` and add `export GOROOT=/opt/local` to your environment file

## Installing GoBDD (required)

Clone it from [this github repo](https://github.com/sdegutis/gobdd) and inside the directory, run `make test install`

If using MacPorts, run `sudo GOROOT=/opt/local/ make test install`

## Running tests

At the root level, type `gotest` to run all the test.
