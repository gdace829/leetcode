package main

func maximumOddBinaryNumber(s string) string {
	num := make([]byte, len(s))
	for i := 0; i < len(num); i++ {
		num[i] = '0'
	}
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count++
		}
	}
	if count >= 1 {
		num[len(num)-1] = '1'

	}
	if count > 1 {
		for i := 0; i < count-1; i++ {
			num[i] = '1'
		}
	}
	return string(num)
}
