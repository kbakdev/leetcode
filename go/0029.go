package _go

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > -2147483648 {
			return -dividend
		}
		return 2147483647
	}
	if dividend == divisor {
		return 1
	}
	if dividend == -divisor {
		return -1
	}
	if dividend > 0 && divisor > 0 {
		if dividend < divisor {
			return 0
		}
		var result int
		for dividend >= divisor {
			dividend -= divisor
			result++
		}
		return result
	}
	if dividend < 0 && divisor < 0 {
		if dividend > divisor {
			return 0
		}
		var result int
		for dividend <= divisor {
			dividend -= divisor
			result++
		}
		return result
	}
	if dividend > 0 && divisor < 0 {
		if dividend < -divisor {
			return 0
		}
		var result int
		for dividend >= -divisor {
			dividend += divisor
			result--
		}
		return result
	}
	if dividend < 0 && divisor > 0 {
		if dividend > -divisor {
			return 0
		}
		var result int
		for dividend <= -divisor {
			dividend += divisor
			result--
		}
		return result
	}
	return 0
}