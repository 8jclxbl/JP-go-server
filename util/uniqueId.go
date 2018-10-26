package util

import (
	"math/rand"
	"strconv"
	"time"
)

const randSup = 4096

func GenerateId() string {
	t := time.Now()
	timeStamp := t.Unix()

	rand.Seed(timeStamp)
	rnd := rand.Intn(randSup)

	mix := strconv.Itoa(int(timeStamp)) + strconv.Itoa(rnd)

	return Cipher(mix)

}
