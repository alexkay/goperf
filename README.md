# Comparing C to Go

A few performance tests to evaluate [Go](http://golang.org/) for [Spek](http://spek.cc/).

## FFT calculations

Call `av_rdf_calc()` 1M times.

* **C**: 3.692 (using lavc)
* **Go**: 3.811 (using lavc via Cgo)
* **Go**: 0.978 (using lavc via Cgo in 4 goroutines)

Call `fft.FFTReal()` from [go-dsp](https://github.com/mjibson/go-dsp) 1M times.

* **Go**: 56.587 (stripped down)

## FFT + magnitudes

* **C**: 7.113
* **Go**: 10.127
* **Go**: 2.469 (4 goroutines)

## Building and running

* `gcc -O3 -std=c99 -lavformat -lavcodec -lavutil -lm fft-lavc.c -o fft-lavc && time ./fft-lavc`
* `go build fft-lavc.go && time ./fft-lavc`


