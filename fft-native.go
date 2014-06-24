package main

import (
	"math"

	"github.com/alexkay/goperf/fft"
)

func main() {
	const N = 1000000 // 1M
	const nbits = 11
	const inputSize = 1 << nbits
	const outputSize = (1 << (nbits - 1)) + 1

	input := make([]float64, inputSize)

	var f = math.Pi
	for i := 0; i < inputSize; i++ {
		f = math.Floor(f * math.Pi)
		input[i] = f
	}

	for k := 0; k < N; k++ {
		fft.FFTReal(input)
	}
}
