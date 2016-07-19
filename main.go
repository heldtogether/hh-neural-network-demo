package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"

	"github.com/goml/gobrain"
)

var (
	rounds       int
	learningRate = 0.6
	numberInputs = 2
	hiddenLayer  = 3
	input        []float64
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())

	flag.IntVar(&rounds, "rounds", 1000, "Number of rounds of training to perform.")
	flag.Parse()
}

func main() {

	patterns := [][][]float64{
		{{0, 0}, {0}},
		{{0, 1}, {1}},
		{{1, 0}, {1}},
		{{1, 1}, {0}},
	}

	ff := &gobrain.FeedForward{}
	ff.Init(numberInputs, hiddenLayer, 1)

	fmt.Printf("Before training using the \"xor\" pattern.\n")

	printResults(ff)

	ff.Train(patterns, rounds, learningRate, 0.4, false)

	fmt.Printf("After training with %d number of rounds using the \"xor\" pattern.\n", rounds)

	printResults(ff)
}

func printResults(ff *gobrain.FeedForward) {
	input = []float64{0, 0}
	fmt.Printf("Try Input: %v. Output: %v. \n", input, ff.Update(input)[0])

	input = []float64{0, 1}
	fmt.Printf("Try Input: %v. Output: %v. \n", input, ff.Update(input)[0])

	input = []float64{1, 0}
	fmt.Printf("Try Input: %v. Output: %v. \n", input, ff.Update(input)[0])

	input = []float64{1, 1}
	fmt.Printf("Try Input: %v. Output: %v. \n", input, ff.Update(input)[0])
}
