# Comparing C to Go

A few performance tests to evaluate [Go](http://golang.org/) for [Spek](http://spek.cc/).

## FFT calculations

Call `av_rdf_calc()` 1M times.

* **C**: 7.465
* **Go**: 7.548 (using Cgo)

Using [go-dsp](https://github.com/mjibson/go-dsp).

* **Go**: 164.15

## FFT + magnitudes

* **C**: 14.361
* **Go**: 18.592 (mag calc alone is 1.6 slower than in C)

## Building and running

* `gcc48 -O3 -std=c99 -lavformat -lavcodec -lavutil -lm fft.c -o fft && time ./fft`
* `go fmt fft.go && go build fft.go && time ./fft`


