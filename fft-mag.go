package main

import (
	"math"
)

//#cgo CFLAGS: -I/usr/local/include
//#cgo LDFLAGS: -lavformat -lavcodec -lavutil -L/usr/local/lib
//#include <libavcodec/avfft.h>
import "C"

func main() {
	const N = 1000000 // 1M
	const nbits = 11
	const inputSize = 1 << nbits
	const outputSize = (1 << (nbits - 1)) + 1

	input := [inputSize]float32{}
	output := [outputSize]float32{}

	cx := C.av_rdft_init(nbits, C.DFT_R2C)

	var f = math.Pi
	for i := 0; i < inputSize; i++ {
		f = math.Floor(f * math.Pi)
		input[i] = float32(f)
	}

	for k := 0; k < N; k++ {
		C.av_rdft_calc(cx, (*C.FFTSample)(&input[0]))

		// Calculate magnitudes.
		const n = inputSize
		const n2 = n * n
		output[0] = float32(10.0 * math.Log10(float64(input[0]*input[0]/n2)))
		output[n/2] = float32(10.0 * math.Log10(float64(input[1]*input[1]/n2)))
		for i := 1; i < n/2; i++ {
			re := input[i*2]
			im := input[i*2+1]
			output[i] = float32(10.0 * math.Log10(float64((re*re+im*im)/n2)))
		}
	}

	C.av_rdft_end(cx)
}
