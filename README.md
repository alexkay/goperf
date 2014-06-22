# Comparing C to Go

## FFT calculations

Call `av_rdf_calc()` 1M times.

* **C**: 7.465
* **Go**: 7.548 (using Cgo)


## Building and running

* `gcc48 -O3 -std=c99 -lavformat -lavcodec -lavutil -lm fft.c -o fft && time ./fft`
* `go fmt fft.go && go build fft.go && time ./fft`


