// Copyright (c) 2010 Joseph D Poirier
// Distributable under the terms of The New BSD License
// that can be found in the LICENSE file.

package main

import (
	"fmt"
)

// Predefined polynomials.
const (
	crc24_a = 0x864cfb // byte width = 3
	crc24_b = 0x800063 // byte width = 3
	crc16   = 0x001021 // byte width = 2
	crc8    = 0x00009b // byte width = 1
)

// MakeTable returns the Table constructed from the specified polynomial.
func MakeTable(poly uint32, bitWidth uint32) []uint32 {
	t := make([]uint32, 256)

	topBit := uint32(1 << (bitWidth - 1))

	widthMask := uint32((((1 << (bitWidth - 1)) - 1) << 1) | 1)

	for i := 0; i < 256; i++ {
		crc := uint32(i) << (bitWidth - 8)

		for j := 0; j < 8; j++ {
			if (crc & topBit) > 0 {
				crc = (crc << 1) ^ poly
			} else {
				crc <<= 1
			}
		}

		t[i] = crc & widthMask
	}

	return t
}

func stdio(t []uint32) {
	for i := 0; i < 256; i++ {
		fmt.Printf("0x%.8x", t[i])

		if i != 255 {
			fmt.Printf(",")
		}

		if ((i + 1) % 10) == 0 {
			fmt.Printf("\n")
		}
	}

	fmt.Printf( "\n" )
}


func main() {

	// poly byte width times bits per byte
        t := MakeTable(crc24_a, 3 * 8)

	fmt.Printf("\n --- crc24_a table --- \n")
	stdio(t)
	
	// -----------------------------------------
	// poly byte width times bits per byte
        t = MakeTable(crc24_b, 3 * 8)

	fmt.Printf("\n --- crc24_b table --- \n")
	stdio(t)

	// -----------------------------------------
	// poly byte width times bits per byte
        t = MakeTable(crc16, 2 * 8)

	fmt.Printf("\n --- crc16 table --- \n")
	stdio(t)

	// -----------------------------------------
	// poly byte width times bits per byte
        t = MakeTable(crc8, 1 * 8)

	fmt.Printf("\n --- crc8 table --- \n")
	stdio(t)
}
