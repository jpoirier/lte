// Copyright (c) 2010 Joseph D Poirier
// Distributable under the terms of The New BSD License
// that can be found in the LICENSE file.

package l1enc

import (
	"testing"
)

type Table [256]uint32

var crcSzs 	= [3]int {24, 16, 8}
var tables 	= [3]*[256]uint32 {&Crc24_B, &Crc16, &Crc8}

//                                  3.840 bits
// non-reflected         0x97561000, 0x22F90000, 0x6A000000
var refCrcs = [3]uint32 {0x00E96A08, 0x0000449F, 0x00000056}

//							6113		4519		2013
// non-reflected         0x553D9500, 0xBB441A00, 0x7062B500
var refCRC24 = []uint32 {0x00953D55, 0x001A44BB, 0x00B56270}

// non-reflected         0xBCC30000, 0x2B190000, 0x98110000
var refCRC16 = []uint32 {0x0000C3BC, 0x0000192B, 0x00001198}

// non-reflected         0x68000000, 0xBF000000, 0x32000000
var refCRC8  = []uint32 {0x00000068, 0x000000BF, 0x00000032}

func TestCrc(t *testing.T) {
	var v uint32
	var c uint32

	//------------------- CRC bit counts with no orphan bits
	bitCnt := (480 * 8)

	// data buffer
	data := make(Data, 480)

	for i := 0; i < 480; i += 4 {
		// known data associated with ref crcs
		// 0x12345678
		data[i + 0] = 0x12
		data[i + 1] = 0x34
		data[i + 2] = 0x56
		data[i + 3] = 0x78
	}

	for i := 0; i < len(tables); i++ {
		if err, crc := Crc(data, tables[i], bitCnt, crcSzs[i]); err {
			t.Errorf("Iteration: %d, CRC size error \n", i)
		} else {
			switch crcSzs[i] {
				case  8:
					v =  uint32(crc[0])
				case 16:
					v = (uint32(crc[0]) << 8) | uint32(crc[1])
				case 24:
					v = (uint32(crc[0]) << 16) | (uint32(crc[1]) << 8) | uint32(crc[2])
			}

			if v != refCrcs[i] {
				t.Errorf("Expected: 0x%X, received: 0x%X \n", refCrcs[i], v)
			}
		}
	}

	//------------------- CRC bit counts with orphan bits
	d1 := make(Data, 800)
	d2 := make(Data, 800)
 	d3 := make(Data, 800)

	bitCnts := [3]int{6113, 4519, 2013}
	byteCnts := [3]int{6113/8+1, 4519/8+1, 2013/8+1}
	buffers := [3][]uint8{d1, d2, d3}

	for i, b := range buffers {
		for j := 0; j < byteCnts[i]; j++ {
			switch j % 0x4 {
				case 0:
					b[j] = 0x12
				case 1:
					b[j] = 0x34
				case 2:
					b[j] = 0x56
				case 3:
					b[j] = 0x78
			}
		}	
	}

	for i := 0; i < len(crcSzs); i++ {
		for j := 0; j < len(buffers); j++ {
			if err, crc := Crc(buffers[j], tables[i], bitCnts[j], crcSzs[i]); err {
				t.Errorf("CRC FUNC ERROR: CRC sz: %d, Bit cnt \n", crcSzs[i], bitCnts[j])
			} else {
				switch crcSzs[i] {
					case  8:
						c = refCRC8[j]
						v =  uint32(crc[0])
					case 16:
						c = refCRC16[j]
						v = (uint32(crc[0]) << 8) | uint32(crc[1])
					case 24:
						c = refCRC24[j]
						v = (uint32(crc[0]) << 16) | (uint32(crc[1]) << 8) | uint32(crc[2])
				}

				if v != c {
					t.Errorf("CRC: %d, Bit Cnt: %d, Expected: 0x%X, received: 0x%X \n", crcSzs[i], bitCnts[j], c, v)
				}
			}
		}
	}
}

func BenchmarkCrc(b *testing.B) {
	b.StopTimer()

	// data creation
	data := make([]uint8, 765)

	for i := 0; i < len(data); i += 4 {
		// 0x12345678
		data[i + 0] = 0x12
		data[i + 1] = 0x34
		data[i + 2] = 0x56
		data[i + 3] = 0x78
	}

	b.StartTimer()

	for i := 0; i < len(tables); i++ {
		_, _ = Crc(data, tables[i], (6114), crcSzs[i])
	}
}
