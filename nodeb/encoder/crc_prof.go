// Copyright (c) 2010 Joseph D Poirier
// Distributable under the terms of The New BSD License
// that can be found in the LICENSE file.

package main

import (
    "lte/l1enc"
    "fmt"
    "runtime"
    "time"
)

// 50 MB/s = 100,000 @ 2ms intervals
// 40 MB/s =  80,000 @ 2ms intervals
// 30 MB/s =  60,000 @ 2ms intervals
// 20 MB/s =  40,000 @ 2ms intervals
// 10 MB/s =  20,000 @ 2ms intervals
//  5 MB/s =  10,000 @ 2ms intervals

var bitCnts = []int {100000, 80000, 60000, 40000, 20000, 10000}
func main() {
	runtime.GOMAXPROCS(2)
	var tBeg, tEnd int64

	// 20 MB @ 10ms intervals = 25,000 * 8 = 200,000 bits
	data := make([]uint8, bitCnts[0]/8)

	for i := 0; i < bitCnts[0] / 8; i+=4 {
		// known data associated with ref crcs
		// 0x12345678
		data[i+0] = 0x12
		data[i+1] = 0x34
		data[i+2] = 0x56
		data[i+3] = 0x78
	}

	tBeg = time.Nanoseconds()
	_,_ = encoder.Crc(data, &encoder.Crc24_B, 24)
	tEnd = time.Nanoseconds()

	fmt.Printf( "Test took %f seconds to run\n", float(tEnd - tBeg) / 1e9)
}
