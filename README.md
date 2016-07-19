# hh-neural-network-demo

Demo of logistic regression that learns the XOR operation.

## Installation

To install, you will need to install Go, see [installation instructions](https://golang.org/doc/install).

In your GOPATH, run

	go get github.com/heldtogether/hh-neural-network-demo

to download the app. Then run

	cd $GOPATH/src/github.com/heldtogether/hh-neural-network-demo 
	go get -t ./...

to download all dependencies.

Finally, install the app with

	go install github.com/heldtogether/hh-neural-network-demo


## Usage

Run the app using

	hh-neural-network-demo [-rounds=100]

Available flags:

`-rounds`: Number of rounds of training to perform. The default is 100.
