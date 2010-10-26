// Copyright (c) 2010 Joseph D Poirier
// Distributable under the terms of The New BSD License
// that can be found in the LICENSE file.

package l1enc

import (
//	"fmt"
	"testing"
)

func TestCopy(t *testing.T) {
	a  := make(Data, 5)
	b  := make(Data, 50)
	c  := make(Data, 500)
	z  := []Data { a, b, c }

	for _, o := range z {
		for i, _ := range o {
			o[i] = 0x56
		}
	}

	for i, o := range z {
		x := o.Copy()

		if x.Len() != o.Len() {
			t.Errorf("Len mismatch: %d != %d \n", x.Len(), o.Len())
		}

		for j, v := range o {
			if x[j] != v {
				t.Errorf("Iter: %d, index: %d, x[0x%X] != o[0x%X] \n", i, j, x[j], v)
			}
		}
	}
}

func TesAppend(t *testing.T) {
	a := Data {0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := Data {10, 11, 12, 13, 14, 15}
	c := Data {0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	x := a.Append(b)

	if x.Len() != (a.Len() + b.Len()) {
		t.Errorf("Len mismatch: %d != %d \n", x.Len(), a.Len() + b.Len())
	}

	for i := 0; i < x.Len(); i++ {
		if x[i] != c[i] {
			t.Errorf("Index %, x[%d] != c[%d] \n", i, x[i], c[i])
		}
	}
}

func TestRsh(t *testing.T) {
	a := Data {0xAA, 0xAA}
	b := Data {0x55, 0x55, 0x00}
	c := Data {0x2A, 0xAA, 0x80}
	d := Data {0x15, 0x55, 0x40}
	e := Data {0x0A, 0xAA, 0xA0}
	f := Data {0x05, 0x55, 0x50}
	g := Data {0x02, 0xAA, 0xA8}
	h := Data {0x01, 0x55, 0x54}
	k := Data {0x00, 0xAA, 0xAA}
	l := Data {0x00, 0x55, 0x55,0x00}
	z := []Data {b, c, d, e, f, g, h, k, l}

	bitCnt := 16 	// number of bits to be shifted

	for i, o := range z {
		a = a.Rsh(bitCnt, 1)	// recycle the reference variable with each iteration
		bitCnt += 1				// the data increases by 1 bit with each iteration

		// Rsh should handle the length correctly when
		// a right shift crosses a byte boundary
		if a.Len() < o.Len() {
			t.Errorf("Iter: %d, a.Len() !< o.Len() \n", i)
		}

		for j := 0; j < o.Len(); j++ {
			if a[j] != o[j] {
				t.Errorf("Iter: %d, index %d, a[0x%.2X] != o[0x%.2X] \n", i, j, a[i], o[i])
			}
		}
	}
}

func TestLshCopy(t *testing.T) {
	a := Data {0xAA, 0xAA}
	b := Data {0x55, 0x54}
	c := Data {0xAA, 0xA8}
	d := Data {0x55, 0x50}
	e := Data {0xAA, 0xA0}
	f := Data {0x55, 0x40}
	g := Data {0xAA, 0x80}
	h := Data {0x55, 0x00}
	k := Data {0xAA, 0x00}
	l := Data {0x54, 0x00}
	m := Data {0xA8, 0x00}
	n := Data {0x50, 0x00}
	p := Data {0xA0, 0x00}
	q := Data {0x40, 0x00}
	r := Data {0x80, 0x00}
	s := Data {0x00, 0x00}

	z := []Data {b, c, d, e, f, g, h, k, l, m, n, p, q, r, s}

	for i, o := range z {
		a = a.Lsh(1)	// recycle the reference variable with each iteration

		for j := 0; j < o.Len(); j++ {
			if a[j] != o[j] {
				t.Errorf("Iter: %d, index %d, a[0x%.2X] != o[0x%.2X] \n", i, j, a[j], o[j])
			}
		}
	}
}

func BenchmarkCopy(b *testing.B) {
	b.StopTimer()

	// data creation
	s := make(Data, 5000)
	d := Data {}

	b.StartTimer()
	d = s.Copy()
	d[0] = 1
}

func BenchmarkAppend(b *testing.B) {
	b.StopTimer()

	// data creation
	x := make(Data, 2500)
	y := Data {}

	b.StartTimer()
	y = x.Copy()
	y[0] = 1
}
