// Copyright (c) 2010 Joseph D Poirier
// Distributable under the terms of The New BSD License
// that can be found in the LICENSE file.

// Slices are a really quick and efficient way to work with chunks of
// data; this file contains some builtin methods that add functionality (???)
// to the "Data" data type.

package l1enc

//import "fmt"

type Data []uint8

// Len is an alias for len()
func (d Data) Len() int {
	return len(d)
}

// Cap is an alias for cap()
func (d Data) Cap() int {
	return cap(d)
}

// CapGrow returns a copy of the source slice
// with its capacity increased by cnt bytes.
func (s Data) CapGrow(cnt int) (d Data) {
	d = make(Data, s.Len(), s.Cap() + cnt)
	copy(d, s)

	return
}

// Copy returns a copy of the length of the source slice.
func (s Data) Copy() (d Data) {
	d = make(Data, s.Len())
	copy(d, s)

	return
}

// Append returns a slice with c appended to s.
func (s Data) Append(c Data) (d Data) {
	l := s.Len()

	d = make(Data, l + c.Len())
	copy(d, s)

	x := d[l:]
	copy(x, c)

	return
}

// Rsh returns a copy of the source slice with bitCnt bits
// right-shifted by shiftCnt bits.
func (s Data) Rsh(bitCnt, shiftCnt int) (d Data) {
	ovflwByteCntShft := shiftCnt / 8
	orphanCnt := shiftCnt % 0x8

	byteCntSrc := bitCnt / 8
	orphanCntSrc := bitCnt % 8
	if orphanCntSrc > 0 { byteCntSrc += 1 }

	z := orphanCnt
	if orphanCntSrc > 0 { z -= (8 - orphanCntSrc) }

	if z > 0 {
		z = 1
	} else {
		z = 0
	}

	d = make(Data, byteCntSrc + ovflwByteCntShft + z)
	dNdx := ovflwByteCntShft

	rShft := uint32(orphanCnt)
	lShft := uint32(8 - orphanCnt)
	t := uint8(0)

// (TODO) is processing an extra byte for orphan bits is really innocuous
	for i := 0; i < byteCntSrc; i++ {
		x := s[i]
		y := x >> rShft
		d[dNdx + i] = y | t
		t = x << lShft
	}

	return
}

// Lsh returns a copy of the source slice shifted left by shiftCnt bits
func (s Data) Lsh(shiftCnt int) (d Data) {
// - handle if all the bits of the slice are being shifted out

	offset := shiftCnt / 8
	orphanCnt := shiftCnt % 0x8

	// default to src length then reset later
	d = make(Data, s.Len(), s.Cap())

	lShft := uint32(orphanCnt)
	rShft := uint32(8 - orphanCnt)
	bytesToProc := s.Len() - offset

	i := 0
	x := s[offset]
	y := uint8(0)
	offset += 1

	for i = 1; i < bytesToProc; i++ {
		y = s[offset]
		d[i - 1] = (x << lShft) | (y >> rShft)
		x = y
		offset += 1
	}

	d[i - 1] = x << lShft

	return
}

