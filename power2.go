package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	N = 1000000
)

func noChoice() {
	var bins [N]int
	var max int
	var EXPECTED_MAX = math.Log(N) / math.Log(math.Log(N))
	for i := 0; i < N; i++ {
		bins[rand.Intn(N)]++
	}
	for i := 0; i < N; i++ {
		if bins[i] > max {
			max = bins[i]
		}
	}
	fmt.Println(max, EXPECTED_MAX)
}

func twoChoice() {
	var bins [N]int
	var max int
	var EXPECTED_MAX = math.Log(math.Log(N))/math.Log(2) + 1
	for i := 0; i < N; i++ {
		x := rand.Intn(N)
		y := rand.Intn(N)
		if bins[x] > bins[y] {
			bins[y]++
		} else {
			bins[x]++
		}
	}
	for i := 0; i < N; i++ {
		if bins[i] > max {
			max = bins[i]
		}
	}
	fmt.Println(max, EXPECTED_MAX)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	noChoice()
	twoChoice()
}
