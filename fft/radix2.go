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

package fft

import (
	"math"
	"sync"
)

var (
	radix2Lock    sync.RWMutex
	radix2Factors = map[int][]complex128{
		4: {complex(1, 0), complex(0, -1), complex(-1, 0), complex(0, 1)},
	}
)

// EnsureRadix2Factors ensures that all radix 2 factors are computed for inputs
// of length input_len. This is used to precompute needed factors for known
// sizes. Generally should only be used for benchmarks.
func EnsureRadix2Factors(input_len int) {
	getRadix2Factors(input_len)
}

func getRadix2Factors(input_len int) []complex128 {
	radix2Lock.RLock()

	if hasRadix2Factors(input_len) {
		defer radix2Lock.RUnlock()
		return radix2Factors[input_len]
	}

	radix2Lock.RUnlock()
	radix2Lock.Lock()
	defer radix2Lock.Unlock()

	if !hasRadix2Factors(input_len) {
		for i, p := 8, 4; i <= input_len; i, p = i<<1, i {
			if radix2Factors[i] == nil {
				radix2Factors[i] = make([]complex128, i)

				for n, j := 0, 0; n < i; n, j = n+2, j+1 {
					radix2Factors[i][n] = radix2Factors[p][j]
				}

				for n := 1; n < i; n += 2 {
					sin, cos := math.Sincos(-2 * math.Pi / float64(i) * float64(n))
					radix2Factors[i][n] = complex(cos, sin)
				}
			}
		}
	}

	return radix2Factors[input_len]
}

func hasRadix2Factors(idx int) bool {
	return radix2Factors[idx] != nil
}

// radix2FFT returns the FFT calculated using the radix-2 DIT Cooley-Tukey algorithm.
func radix2FFT(x []complex128) []complex128 {
	lx := len(x)
	factors := getRadix2Factors(lx)

	t := make([]complex128, lx) // temp
	r := reorderData(x)

	for stage := 2; stage <= lx; stage <<= 1 {
		blocks := lx / stage
		s_2 := stage / 2

		for nb := 0; nb < lx; nb += stage {
			if stage != 2 {
				for j := 0; j < s_2; j++ {
					idx := j + nb
					idx2 := idx + s_2
					ridx := r[idx]
					w_n := r[idx2] * factors[blocks*j]
					t[idx] = ridx + w_n
					t[idx2] = ridx - w_n
				}
			} else {
				n1 := nb + 1
				rn := r[nb]
				rn1 := r[n1]
				t[nb] = rn + rn1
				t[n1] = rn - rn1
			}
		}

		r, t = t, r
	}

	return r
}

// reorderData returns a copy of x reordered for the DFT.
func reorderData(x []complex128) []complex128 {
	lx := uint(len(x))
	r := make([]complex128, lx)
	s := log2(lx)

	var n uint
	for ; n < lx; n++ {
		r[reverseBits(n, s)] = x[n]
	}

	return r
}

// log2 returns the log base 2 of v
// from: http://graphics.stanford.edu/~seander/bithacks.html#IntegerLogObvious
func log2(v uint) uint {
	var r uint

	for v >>= 1; v != 0; v >>= 1 {
		r++
	}

	return r
}

// reverseBits returns the first s bits of v in reverse order
// from: http://graphics.stanford.edu/~seander/bithacks.html#BitReverseObvious
func reverseBits(v, s uint) uint {
	var r uint

	// Since we aren't reversing all the bits in v (just the first s bits),
	// we only need the first bit of v instead of a full copy.
	r = v & 1
	s--

	for v >>= 1; v != 0; v >>= 1 {
		r <<= 1
		r |= v & 1
		s--
	}

	return r << s
}
