package specodec

import "specodec/ryu"

func FormatFloat32(value float32) string {
	return ryu.Float32ToString(value)
}

func FormatFloat64(value float64) string {
	return ryu.Float64ToString(value)
}
	absVal := math.Abs(value)

	if absVal == 0 {
		return sign + "0." + strings.Repeat("0", n-1) + "E+00"
	}

	expStr := fmt.Sprintf("%." + fmt.Sprintf("%d", n+2) + "e", absVal)
	
	var intPart string
	var fracPart string
	var expSign string
	var expVal int
	
	parts := strings.Split(expStr, "e")
	if len(parts) != 2 {
		panic("unexpected exponential format")
	}
	coeff := parts[0]
	expPart := parts[1]
	
	intPart = coeff[:1]
	fracPart = coeff[2:]
	expSign = expPart[:1]
	expValStr := expPart[1:]
	expVal = 0
	for _, c := range expValStr {
		expVal = expVal*10 + int(c-'0')
	}

	allDigits := intPart + fracPart

	lastKeptDigit := int(allDigits[n-1] - '0')
	firstDroppedDigit := int(allDigits[n] - '0')

	roundingDigits := allDigits[:n]

	roundUp := false
	if firstDroppedDigit > 5 {
		roundUp = true
	} else if firstDroppedDigit < 5 {
		roundUp = false
	} else {
		tailAfter5 := allDigits[n+1:]
		hasNonZero := false
		for _, c := range tailAfter5 {
			if c != '0' {
				hasNonZero = true
				break
			}
		}
		if hasNonZero {
			roundUp = true
		} else {
			roundUp = lastKeptDigit%2 == 1
		}
	}

	if roundUp {
		digitsArr := strings.Split(roundingDigits, "")
		carry := 1
		for i := n - 1; i >= 0 && carry > 0; i-- {
			newDigit := int(digitsArr[i][0]-'0') + carry
			digitsArr[i] = string('0' + newDigit%10)
			carry = newDigit / 10
		}
		if carry > 0 {
			digitsArr = append([]string{"1"}, digitsArr...)
			expVal++
			digitsArr = digitsArr[:n]
		}
		roundingDigits = strings.Join(digitsArr, "")
	}

	finalExp := expVal
	if expSign == "-" {
		finalExp = -expVal
	}
	expSignOut := "+"
	if finalExp < 0 {
		expSignOut = "-"
		finalExp = -finalExp
	}
	expStrOut := fmt.Sprintf("%02d", finalExp)

	if n == 1 {
		return sign + roundingDigits + "E" + expSignOut + expStrOut
	}
	return sign + roundingDigits[:1] + "." + roundingDigits[1:] + "E" + expSignOut + expStrOut
}

func formatFloat32(value float32) string {
	return formatSigFigs(float64(value), 9)
}

func formatFloat64(value float64) string {
	return formatSigFigs(value, 17)
}