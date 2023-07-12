package otp

import "math/rand"

var numberSet = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func GenereteOtpNum(length int) string {
	if length <= 0 {
		return ""
	}

	res := ""

	for i := 0; i < length; i++ {
		index := rand.Intn(10)
		res += numberSet[index]
	}

	return res
}
