package main

func pow(a, b, m int) int {
	ans := 1
	for b > 0 {
		if (b & 1) > 0 {
			ans *= a
			ans %= m
		}
		a *= a
		b >>= 1
	}
	return ans
}

func mint(a, m int) int {
	return pow(a, m-2, m)
}
