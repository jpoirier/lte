/*
	3GPP TS 36.212 v8.3.0 Multiplexing and channel coding (Release 8)
	Section 5.1.2 Code block segmentation and code block CRC attachment
*/
package l1enc

import (
	"testing"
)

func TestBlkSegParams(t *testing.T) {
	var bitCnts = []int {40, 500, 2000, 4000, 6000, 10000, 15000, 25000, 51000}

	for _, v := range bitCnts {
		if err, _ := BlkSegParams(v); err {
			t.Errorf("Bit cnt: %d \n", v)
		}
//fmt.Printf("C: %v, Ce: %v, Co: %v, Ke: %v, Ko: %v, F: %v, L: %v \n\n", nfo.C, nfo.Ce, nfo.Co, nfo.Ke, nfo.Ko, nfo.F, nfo.L)
	}
}
