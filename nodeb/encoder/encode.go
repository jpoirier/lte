// Copyright (c) 2010 Joseph D Poirier
// Distributable under the terms of The New BSD License
// that can be found in the LICENSE file.

package l1enc

import (
//	"fmt"	
//	"os"
)

// (TODO) jdp: the function should take a struct of info rather than blkSz

func Encode(src []uint8, blkSz int) (dst []uint8) {
    segNfo := encoder.BlkSegParams(blkSz)

}
