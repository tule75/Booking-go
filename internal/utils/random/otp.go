package random

import (
	"math/rand"
	"time"
)

func GenerateOTP() int {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	otp := 100000 + random.Intn(900000)

	return otp

}
