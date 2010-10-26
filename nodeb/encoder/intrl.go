// Copyright (c) 2010 Joseph D Poirier
// Distributable under the terms of The New BSD License
// that can be found in the LICENSE file.

// 3GPP TS 36.212 V8.6.0 (2009-03) Multiplexing and channel coding (Release 8)
// Section 5.1.3	Channel coding

package l1enc

type Data []uint8

type params struct {
   Ki uint32
   f1 uint32
   f2 uint32
}

var intrlParams = [188]params {
	// 	Ki   f1   f2       i
	params{  40,   3,  10}, // 1
	params{  48,   7,  12},
	params{  56,  19,  42},
	params{  64,   7,  16},
	params{  72,   7,  18},
	params{  80,  11,  20},
	params{  88,   5,  22},
	params{  96,  11,  24},
	params{ 104,   7,  26},
	params{ 112,  41,  84}, // 10
	params{ 120, 103,  90},
	params{ 128,  15,  32},
	params{ 136,   9 , 34},
	params{ 144,  17, 108},
	params{ 152,   9,  38},
	params{ 160,  21, 120},
	params{ 168, 101,  84},
	params{ 176,  21,  44},
	params{ 184,  57,  46},
	params{ 192,  23,  48}, // 20
	params{ 200,  13,  50},
	params{ 208,  27,  52},
	params{ 216,  11,  36},
	params{ 224,  27,  56},
	params{ 232,  85,  58},
	params{ 240,  29,  60},
	params{ 248,  33,  62},
	params{ 256,  15,  32},
	params{ 264,  17, 198},
	params{ 272,  33,  68}, // 30
	params{ 280, 103, 210},
	params{ 288,  19,  36},
	params{ 296,  19,  74},
	params{ 304,  37,  76},
	params{ 312,  19,  78},
	params{ 320,  21, 120},
	params{ 328,  21,  82},
	params{ 336, 115,  84},
	params{ 344, 193,  86},
	params{ 352,  21,  44}, // 40
	params{ 360, 133,  90},
	params{ 368,  81,  46},
	params{ 376,  45,  94},
	params{ 384,  23,  48},
	params{ 392, 243,  98},
	params{ 400, 151,  40},
	params{ 408, 155, 102},
	params{ 416,  25,  52},
	params{ 424,  51, 106},
	params{ 432,  47,  72}, // 50
	params{ 440,  91, 110},
	params{ 448,  29, 168},
	params{ 456,  29, 114},
	params{ 464, 247,  58},
	params{ 472,  29, 118},
	params{ 480,  89, 180},
	params{ 488,  91, 122},
	params{ 496, 157,  62},
	params{ 504,  55,  84},
	params{ 512,  31,  64}, // 60
	params{ 528,  17,  66},
	params{ 544,  35,  68},
	params{ 560, 227, 420},
	params{ 576,  65,  96},
	params{ 592,  19,  74},
	params{ 608,  37,  76},
	params{ 624,  41, 234},
	params{ 640,  39,  80},
	params{ 656, 185,  82},
	params{ 672,  43, 252}, // 70
	params{ 688,  21,  86},
	params{ 704, 155,  44},
	params{ 720,  79, 120},
	params{ 736, 139,  92},
	params{ 752,  23,  94},
	params{ 768, 217,  48},
	params{ 784,  25,  98},
	params{ 800,  17,  80},
	params{ 816, 127, 102},
	params{ 832,  25,  52}, // 80
	params{ 848, 239, 106},
	params{ 864,  17,  48},
	params{ 880, 137, 110},
	params{ 896, 215, 112},
	params{ 912,  29, 114},
	params{ 928,  15,  58},
	params{ 944, 147, 118},
	params{ 960,  29,  60},
	params{ 976,  59, 122},
	params{ 992,  65, 124}, // 90
	params{1008,  55,  84},
	params{1024,  31,  64},
	params{1056,  17,  66},
	params{1088, 171, 204},
	params{1120,  67, 140},
	params{1152,  35,  72},
	params{1184,  19,  74},
	params{1216,  39,  76},
	params{1248,  19,  78},
	params{1280, 199, 240}, // 100
	params{1312,  21,  82},
	params{1344, 211, 252},
	params{1376,  21,  86},
	params{1408,  43,  88},
	params{1440, 149,  60},
	params{1472,  45,  92},
	params{1504,  49, 846},
	params{1536,  71,  48},
	params{1568,  13,  28},
	params{1600,  17,  80}, // 100
	params{1632,  25, 102},
	params{1664, 183, 104},
	params{1696,  55, 954},
	params{1728, 127,  96},
	params{1760,  27, 110},
	params{1792,  29, 112},
	params{1824,  29, 114},
	params{1856,  57, 116},
	params{1888,  45, 354},
	params{1920,  31, 120}, // 120
	params{1952,  59, 610},
	params{1984, 185, 124},
	params{2016, 113, 420},
	params{2048,  31,  64},
	params{2112,  17,  66},
	params{2176, 171, 136},
	params{2240, 209, 420},
	params{2304, 253, 216},
	params{2368, 367, 444},
	params{2432, 265, 456}, // 130
	params{2496, 181, 468},
	params{2560,  39,  80},
	params{2624,  27, 164},
	params{2688, 127, 504},
	params{2752, 143, 172},
	params{2816,  43,  88},
	params{2880,  29, 300},
	params{2944,  45,  92},
	params{3008, 157, 188},
	params{3072,  47,  96}, // 140
	params{3136,  13,  28},
	params{3200, 111, 240},
	params{3264, 443, 204},
	params{3328,  51, 104},
	params{3392,  51, 212},
	params{3456, 451, 192},
	params{3520, 257, 220},
	params{3584,  57, 336},
	params{3648, 313, 228},
	params{3712, 271, 232}, // 150
	params{3776, 179, 236},
	params{3840, 331, 120},
	params{3904, 363, 244},
	params{3968, 375, 248},
	params{4032, 127, 168},
	params{4096,  31,  64},
	params{4160,  33, 130},
	params{4224,  43, 264},
	params{4288,  33, 134},
	params{4352, 477, 408}, // 160
	params{4416,  35, 138},
	params{4480, 233, 280},
	params{4544, 357, 142},
	params{4608, 337, 480},
	params{4672,  37, 146},
	params{4736,  71, 444},
	params{4800,  71, 120},
	params{4864,  37, 152},
	params{4928,  39, 462},
	params{4992, 127, 234}, // 170
	params{5056,  39, 158},
	params{5120,  39,  80},
	params{5184,  31,  96},
	params{5248, 113, 902},
	params{5312,  41, 166},
	params{5376, 251, 336},
	params{5440,  43, 170},
	params{5504,  21,  86},
	params{5568,  43, 174},
	params{5632,  45, 176}, // 180
	params{5696,  45, 178},
	params{5760, 161, 120},
	params{5824,  89, 182},
	params{5888, 323, 184},
	params{5952,  47, 186},
	params{6016,  23,  94},
	params{6080,  47, 190},
	params{6144, 263, 480}} // 188

func qppIntrl(src Data, dst Data, bitCnt uint32) (err bool) {

	var iSrc uint32
	var iDst uint32
	var i uint32
	var j uint32
	var aa uint32
	var bitPos uint32
	var bit	uint8

	K := bitCnt
	
	switch {
		default:
			err = true
			return
		case K < 40:
			i = 0
		case K < 512:
			i = (K / 8) - 5
		case K < 1024:
			i = (K / 16) + 27
		case K < 2048:
			i = (K / 32) + 59
		case K < 6144:
			i = (K / 64) + 91
	}

	f := intrlParams[i]

	for j = 0; j < K; j++ {
		// a % n  < same as >  a - (n * int(a / n))
		// calc the src and dst data indicies (sect 5.1.3.2.3)
		aa 	= (f.f1 * j) + (f.f2 * j * j)
		aa 	= aa - (K * (aa / K))
		iSrc 	= aa / 8
		iDst 	= j / 8

		// fetch the source bit
		bitPos	= aa & 0x7
		bit 	= (src[iSrc] >> (7 - bitPos)) & 0x1

		// store the output bit
		bitPos	= j & 0x7
		dst[iDst] |= (bit << (7 - bitPos))
	}

	return
}

