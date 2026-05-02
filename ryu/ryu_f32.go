package ryu

// Float32 configuration constants
const FLOAT_MANTISSA_BITS = 23
const FLOAT_BIAS = 127
const FLOAT_POW5_INV_BITCOUNT = 59
const FLOAT_POW5_BITCOUNT = 61

func Float32ToString(f float32) string {
	bits := math.Float32bits(f)
	
	sign := (bits >> 31) != 0
	ieeeMantissa := bits & 0x7FFFFF
	ieeeExponent := (bits >> 23) & 0xFF
	
	if ieeeExponent == 255 {
		if ieeeMantissa == 0 {
			if sign { return "-Infinity" }
			return "Infinity"
		}
		return "NaN"
	}
	if ieeeExponent == 0 && ieeeMantissa == 0 {
		if sign { return "-0E0" }
		return "0E0"
	}
	
	var e2 int32
	var m2 uint64
	if ieeeExponent == 0 {
		e2 = 1 - FLOAT_BIAS - FLOAT_MANTISSA_BITS - 2
		m2 = uint64(ieeeMantissa)
	} else {
		e2 = int32(ieeeExponent) - FLOAT_BIAS - FLOAT_MANTISSA_BITS - 2
		m2 = uint64((1 << FLOAT_MANTISSA_BITS) | ieeeMantissa)
	}
	
	even := (m2 & 1) == 0
	acceptBounds := even
	
	mv := m2 * 4
	mp := mv + 2
	mmShift := int32(0)
	if ieeeMantissa != 0 || ieeeExponent <= 1 {
		mmShift = 1
	}
	mm := mv - 1 - uint64(mmShift)
	
	vrIsTrailingZeros := false
	vmIsTrailingZeros := false
	lastDigit := uint64(0)
	var e10 int32
	var vr, vp, vm_ uint64
	
	if e2 >= 0 {
		q := log10Pow2(e2)
		e10 = q
		k := FLOAT_POW5_INV_BITCOUNT + pow5bits(q) - 1
		i := -e2 + q + k
		
		vr = mulShift32(mv, FLOAT_POW5_INV_SPLIT[q]+1, i)
		vp = mulShift32(mp, FLOAT_POW5_INV_SPLIT[q]+1, i)
		vm_ = mulShift32(mm, FLOAT_POW5_INV_SPLIT[q]+1, i)
		
		if q != 0 && (vp-1)/10 <= vm_/10 {
			l := FLOAT_POW5_INV_BITCOUNT + pow5bits(q-1) - 1
			lastDigit = mulShift32(mv, FLOAT_POW5_INV_SPLIT[q-1]+1, -e2+q-1+l) % 10
		}
		
		if q <= 9 {
			if mv % 5 == 0 {
				vrIsTrailingZeros = multipleOfPowerOf5_64(mv, q)
			} else if acceptBounds {
				vmIsTrailingZeros = multipleOfPowerOf5_64(mm, q)
			} else {
				if multipleOfPowerOf5_64(mp, q) {
					vp--
				}
			}
		}
	} else {
		q := log10Pow5(-e2)
		e10 = q + e2
		i := -e2 - q
		k := pow5bits(i) - FLOAT_POW5_BITCOUNT
		j := q - k
		
		vr = mulShift32(mv, FLOAT_POW5_SPLIT[i], j)
		vp = mulShift32(mp, FLOAT_POW5_SPLIT[i], j)
		vm_ = mulShift32(mm, FLOAT_POW5_SPLIT[i], j)
		
		if q != 0 && (vp-1)/10 <= vm_/10 {
			j2 := q - 1 - (pow5bits(i+1) - FLOAT_POW5_BITCOUNT)
			lastDigit = mulShift32(mv, FLOAT_POW5_SPLIT[i+1], j2) % 10
		}
		
		if q <= 1 {
			vrIsTrailingZeros = true
			if acceptBounds {
				vmIsTrailingZeros = mmShift == 1
			} else {
				vp--
			}
		} else if q < 31 {
			vrIsTrailingZeros = multipleOfPowerOf2_64(mv, q-1)
			if acceptBounds {
				vmIsTrailingZeros = multipleOfPowerOf5_64(mm, q)
			} else {
				if multipleOfPowerOf5_64(mp, q) {
					vp--
				}
			}
		}
	}
	
	removed := int32(0)
	vr2, vp2, vm2 := vr, vp, vm_
	
	if vmIsTrailingZeros || vrIsTrailingZeros {
		for vp2/10 > vm2/10 {
			vmIsTrailingZeros = vmIsTrailingZeros && (vm2%10 == 0)
			vrIsTrailingZeros = vrIsTrailingZeros && (lastDigit == 0)
			lastDigit = vr2 % 10
			vr2 /= 10
			vp2 /= 10
			vm2 /= 10
			removed++
		}
		
		if vmIsTrailingZeros {
			for vm2%10 == 0 {
				vrIsTrailingZeros = vrIsTrailingZeros && (lastDigit == 0)
				lastDigit = vr2 % 10
				vr2 /= 10
				vp2 /= 10
				vm2 /= 10
				removed++
			}
		}
		
		if vrIsTrailingZeros && lastDigit == 5 && (vr2&1) == 0 {
			lastDigit = 4
		}
		
		roundUp := (vr2 == vm2 && (!acceptBounds || !vmIsTrailingZeros)) || lastDigit >= 5
		output := vr2
		if roundUp {
			output++
		}
		exp := e10 + removed
		olength := decimalLength17(output)
		
		result := ""
		if sign { result = "-" }
		digits := strconv.FormatUint(output, 10)
		if olength == 1 {
			result += digits
		} else {
			result += digits[:1] + "." + digits[1:]
		}
		result += "E" + strconv.FormatInt(int64(exp+olength-1), 10)
		return result
	} else {
		for vp2/10 > vm2/10 {
			lastDigit = vr2 % 10
			vr2 /= 10
			vp2 /= 10
			vm2 /= 10
			removed++
		}
		
		output := vr2
		if vr2 == vm2 || lastDigit >= 5 {
			output++
		}
		exp := e10 + removed
		olength := decimalLength17(output)
		
		result := ""
		if sign { result = "-" }
		digits := strconv.FormatUint(output, 10)
		if olength == 1 {
			result += digits
		} else {
			result += digits[:1] + "." + digits[1:]
		}
		result += "E" + strconv.FormatInt(int64(exp+olength-1), 10)
		return result
	}
}

