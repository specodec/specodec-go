package ryu

func pow5bits(e int32) int32 {
	return e*1217359/524288 + 1
}

func log10Pow2(e int32) int32 {
	return e*78913/262144
}

func log10Pow5(e int32) int32 {
	return e*732923/1048576
}

func decimalLength17(v uint64) int32 {
	if v >= 10000000000000000 { return 17 }
	if v >= 1000000000000000 { return 16 }
	if v >= 100000000000000 { return 15 }
	if v >= 10000000000000 { return 14 }
	if v >= 1000000000000 { return 13 }
	if v >= 100000000000 { return 12 }
	if v >= 10000000000 { return 11 }
	if v >= 1000000000 { return 10 }
	if v >= 100000000 { return 9 }
	if v >= 10000000 { return 8 }
	if v >= 1000000 { return 7 }
	if v >= 100000 { return 6 }
	if v >= 10000 { return 5 }
	if v >= 1000 { return 4 }
	if v >= 100 { return 3 }
	if v >= 10 { return 2 }
	return 1
}

func mulShift32(m uint64, factor uint64, shift int32) uint64 {
	factorLo := factor & 0xFFFFFFFF
	factorHi := factor >> 32
	bits0 := m * factorLo
	bits1 := m * factorHi
	sumVal := (bits0 >> 32) + bits1
	return (sumVal >> (shift - 32)) & 0xFFFFFFFF
}

func mulShift64(m uint64, mul []uint64, shift int32) uint64 {
	hi0, _ := umul128(m, mul[0])
	hi2, lo2 := umul128(m, mul[1])
	
	sum_lo := lo2 + hi0
	carry := uint64(0)
	if sum_lo < lo2 {
		carry = 1
	}
	sum_hi := hi2 + carry
	
	shiftAmount := shift - 64
	if shiftAmount >= 128 {
		return 0
	} else if shiftAmount >= 64 {
		return sum_hi >> (shiftAmount - 64)
	} else if shiftAmount == 0 {
		return sum_lo
	} else {
		return (sum_hi << (64 - shiftAmount)) | (sum_lo >> shiftAmount)
	}
}

func umul128(a, b uint64) (uint64, uint64) {
	aLo := a & 0xFFFFFFFF
	aHi := a >> 32
	bLo := b & 0xFFFFFFFF
	bHi := b >> 32
	
	ll := aLo * bLo
	lh := aLo * bHi
	hl := aHi * bLo
	hh := aHi * bHi
	
	mid := lh + hl
	overflow := uint64(0)
	if mid < lh {
		overflow = 1
	}
	
	hi := hh + (mid >> 32) + overflow
	lo := ll + ((mid & 0xFFFFFFFF) << 32)
	if lo < ll {
		hi++
	}
	
	return hi, lo
}

func multipleOfPowerOf5_64(value uint64, q int32) bool {
	if q == 0 { return true }
	if q >= 64 { return value == 0 }
	pow5 := uint64(5)
	for i := int32(1); i < q; i++ {
		pow5 *= 5
	}
	return (value % pow5) == 0
}

func multipleOfPowerOf2_64(value uint64, q int32) bool {
	if q == 0 { return true }
	if q >= 64 { return value == 0 }
	return (value & ((1 << q) - 1)) == 0
}