// Copyright (c) 2010 Joseph D Poirier
// Distributable under the terms of The New BSD License
// that can be found in the LICENSE file.

// GPP TS 36.212 V8.6.0 (2009-03) Multiplexing and channel coding (Release 8)
// Section 5.1.2	Code block segmentation and code block CRC attachment

package l1enc

import (
	"math"
)

const Z = 6144		// turbo coder max block size

type BlkSegNfo struct {
	BlkCnt  int		// code block count
	FillCnt int		// fill bit count
	EblkCnt int		// even code block count
	EbyteCnt int	// bytes per block
	EbitCnt int		// bit count even code blocks
	Eorphans int	// orphan bits in the last byte
	OblkCnt int		// odd code block count
	ObyteCnt int	// bytes per block
	ObitCnt int		// bit count odd code blocks
	Oorphans int	// orphan bits in the last byte
	CrcSz int		// individual block segment crc size
}

//Crc(data []uint8, table *[256]uint32, bitCnt int, crcSz int) (err bool, crc uint32) {

func BlkSeg(data Data, nfo BlkSegNfo) []Data {
// (todo) add some sort of sanity check of the src buffer bit count!
// sanity check
//	t1 := (srcByteCnt * 8) + (nfo.crcSz * nfo.blkCnt) + nfo.fillCnt
//	t2 := (nfo.blkCnt * eBitCnt) + (nfo.blkCnt * oBitCnt)
// 	t1 should equal t2

// 20 MB user throughput @ 2ms TTI = 5,000 bytes
// 6114 / 8 = 765 bytes

	// slice buffer
	segCnt := nfo.EblkCnt + nfo.OblkCnt
	blkSegs := make([]Data, segCnt)

	// even/odd segment bit counts, excluding the crc
	eRawBits := nfo.EbitCnt - nfo.CrcSz
	oRawBits := nfo.ObitCnt - nfo.CrcSz

	// special case, process the first block segment
	b:= (eRawBits - nfo.FillCnt) / 8
	o := (eRawBits - nfo.FillCnt) % 8
	if o > 0 { b += 1 }

	s := data[0:b]	// src slice
	d := s.Copy()

	// right shift implicitly adds zeros for the fill bits
	blkSegs[0] = d.Rsh(eRawBits - nfo.FillCnt, nfo.FillCnt)
//	blkSegs[0] = d

	oZ := oRawBits / 8
	if (oRawBits % 8) > 0 { oZ += 1 }

	eZ := oRawBits / 8
	if (oRawBits % 8) > 0 { eZ += 1 }

	z := 0
	orphans := 0
	x := (eRawBits - nfo.FillCnt)

	for i := 1; i < segCnt; i++ {
		if (i & 1) == 1 { // odd case
			z = oZ
		} else { // even case
			z = eZ
		}

		// calc the start and end of
		// the src data to be copied
		y := x / 8
		orphans = x % 8
		s = data[y : y + z]

		// dst slice
//		d = Data {}
//		d.Copy(s)
//		d.Lsh(orphans)
//		blkSegs[i] = d

		d = s.Copy()
		blkSegs[i] = d.Lsh(orphans)

		if (i & 1) == 1 { // odd case
			x += oRawBits
		} else { // even case
			x += eRawBits
		}
	}

	return blkSegs
}

func BlkSegParams(BlkSz int) (err bool, nfo BlkSegNfo) {
	if BlkSz <= 0 { err = true; return }

	var Bpri int = 0
	var K int = 0

	// number of code blocks
	if BlkSz <= Z {
		nfo.CrcSz = 0
		nfo.BlkCnt = 1
		Bpri = BlkSz
	} else {
		nfo.CrcSz = 24
		nfo.BlkCnt = int(math.Ceil((float64(BlkSz) / float64(Z - nfo.CrcSz))))
		Bpri = BlkSz + nfo.BlkCnt * nfo.CrcSz
	}

	// segmentation size
	// first segmentation size: eBitCnt = minimum K in table  5.1.3-3 such that blkCnt * K >= Bpri
	if nfo.BlkCnt == 1 {
		nfo.EbitCnt = blkSizeMin(Bpri)
		nfo.EblkCnt = 1
		nfo.ObitCnt = 0
		nfo.OblkCnt = 0
	} else if nfo.BlkCnt > 1 {
		sanity := 5
		K = Bpri / nfo.BlkCnt

		for {
			sanity -= 1
			nfo.EbitCnt = blkSizeMin(K)

			if nfo.BlkCnt * nfo.EbitCnt >= Bpri {
				break
			}

			if sanity == 0 {
				err = true
				return
			}
		}

		// second segmentation size: oBitCnt = maximum K in table 5.1.3-3 such that K < eBitCnt
		nfo.ObitCnt = blkSizeMax(nfo.EbitCnt)
		DeltaK := nfo.EbitCnt - nfo.ObitCnt
		nfo.OblkCnt = int(math.Floor((float64(nfo.BlkCnt * nfo.EbitCnt - Bpri) / float64(DeltaK))))
		nfo.EblkCnt = nfo.BlkCnt - nfo.OblkCnt
	}

	// fill-bit count
	nfo.FillCnt = nfo.EblkCnt * nfo.EbitCnt + nfo.OblkCnt * nfo.ObitCnt - Bpri

	nfo.Eorphans = nfo.EbitCnt % 0x8
	nfo.Oorphans = nfo.ObitCnt & 0x8
	nfo.EbyteCnt = nfo.EbitCnt / 8
	nfo.ObyteCnt = nfo.ObitCnt / 8
	if nfo.Eorphans > 0 { nfo.EbyteCnt += 1 }
	if nfo.Oorphans > 0 { nfo.ObyteCnt += 1 }

	return
}

func blkSizeMax(bitCnt int) (blkSz int) {
	switch {
		case bitCnt == 40:
			blkSz = 39
		case bitCnt <= 512:
			blkSz = bitCnt - 8
		case bitCnt <= 1024:
			blkSz = bitCnt -  16
		case bitCnt <= 2048:
			blkSz = bitCnt -  32
		case bitCnt <= 4096:
			blkSz = bitCnt -  64
		case bitCnt <= 6144:
			blkSz = bitCnt -  128
	}

	return
}

func blkSizeMin(bitCnt int) (blkSz int) {
	switch {
		case bitCnt < 40:
			blkSz = 40
		case bitCnt <= 512:
			// 40   -  512 incre by 8
			blkSz = ((bitCnt + 7) / 8) * 8
		case bitCnt <= 1024:
			// 513  - 1024 incre by 16
			blkSz = ((bitCnt + 15) / 16) * 16
		case bitCnt <= 2048:
			// 1025 - 2048 incre by 32
			blkSz = ((bitCnt + 31) / 32) * 32
		case bitCnt <= 4096:
			// 2049 - 4096 incre by 64
			blkSz = ((bitCnt + 63) / 64) * 64
		case bitCnt <= 6144:
			// 4097 - 6144 incre by 128
			blkSz = ((bitCnt + 127) / 128) * 128
	}

	return
}

