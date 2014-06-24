# Comparing C to Go

A few performance tests to evaluate [Go](http://golang.org/) for [Spek](http://spek.cc/).

## FFT calculations

Call `av_rdf_calc()` 1M times.

* **C**: 7.465 (using lavc)
* **Go**: 7.548 (using lavc via Cgo)
* **Go**: 1.892 (using lavc via Cgo in 4 goroutines)

Call `fft.FFTReal()` from [go-dsp](https://github.com/mjibson/go-dsp) 1M times.

* **Go**: 164.15 (stock)
* **Go**: 124.98 (stripped down)

## FFT + magnitudes

* **C**: 14.361
* **Go**: 18.592 (mag calc alone is 1.6 times slower than in C)

## Building and running

* `gcc48 -O3 -std=c99 -lavformat -lavcodec -lavutil -lm fft-lavc.c -o fft-lavc && time ./fft-lavc`
* `go fmt fft-lavc.go && go build fft-lavc.go && time ./fft-lavc`


