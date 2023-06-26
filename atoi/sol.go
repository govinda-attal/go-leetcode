package atoi

import (
	"math"
	"strings"
)

func myAtoi(s string) int {
	sign := 1
	i := 0
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return 0
	}
	if s[0] == '-' {
		sign = -1
		i = 1
	} else if s[0] == '+' {
		i = 1
	}

	var rs int
	for ; i < len(s); i++ {
		val := int(s[i] - '0')
		if val < 0 || val > 9 {
			return rs * sign
		}
		rs = rs*10 + val
		if rs*sign < math.MinInt32 {
			return math.MinInt32
		}
		if rs*sign > math.MaxInt32 {
			return math.MaxInt32
		}
	}
	return rs * sign

}
