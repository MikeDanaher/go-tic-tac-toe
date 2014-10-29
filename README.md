Tic-Tac-Toe in Go
=================

An unbeatable Tic-Tac-Toe game written in Go

### Setup

1. [Install Go](https://golang.org/doc/install) by downloading the distribution.
  * Can also be done using homebrew: brew install go
2. `git clone git@github.com:MikeDanaher/go-tic-tac-toe.git`

### Playing the game

While in the 'go-tic-tac-toe' directory start the game with:

`./go-tic-tac-toe`

*This was compiled on a Mac running OS X. You may need to recompile by using `go build`. Then run the above command.*

### Running the tests

The tests were written using [Ginkgo](http://onsi.github.io/ginkgo/). You'll need to set your $GOPATH and get Ginkgo to run the tests.

1. Set $GOPATH, Go recomends `$ export GOPATH=$HOME/go`
2. Get Ginkgo `$ go get github.com/onsi/ginkgo/ginkgo`
2. Get Gomega `$ go get github.com/onsi/gomega`

After installing Ginkgo, move into the tic-tac-toe directory and run them using `ginkgo` or `go test`.

1. `cd tic-tac-toe`
2. `ginkgo`

### Additional Information

For more information about Go or to setup a development environment on your machine, check out the [Go Docs](https://golang.org/).
