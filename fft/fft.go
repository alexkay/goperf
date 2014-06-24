/*
 * Copyright (c) 2011 Matt Jibson <matt.jibson@gmail.com>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

// Package fft provides forward and inverse fast Fourier transform functions.
package fft

// FFTReal returns the forward FFT of the real-valued slice.
func FFTReal(x []float64) []complex128 {
	return FFT(ToComplex(x))
}

// FFT returns the forward FFT of the complex-valued slice.
func FFT(x []complex128) []complex128 {
	lx := len(x)

	// todo: non-hack handling length <= 1 cases
	if lx <= 1 {
		r := make([]complex128, lx)
		copy(r, x)
		return r
	}

	if IsPowerOf2(lx) {
		return radix2FFT(x)
	}

	panic("len must be the power of 2")
}

var (
	worker_pool_size = 0
)

// SetWorkerPoolSize sets the number of workers during FFT computation on multicore systems.
// If n is 0 (the default), then GOMAXPROCS workers will be created.
func SetWorkerPoolSize(n int) {
	if n < 0 {
		n = 0
	}

	worker_pool_size = n
}
