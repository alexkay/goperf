package main

import (
	"flag"
	"log"
	"math"
	"os"
	"runtime/pprof"

	"github.com/alexkay/goperf/fft"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

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
