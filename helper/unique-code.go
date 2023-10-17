package helper

import (
	"math"
	"math/rand"
	"strconv"
)

func UniqueCode() string {
	randomInt := math.MaxInt32
	rand := rand.Intn(randomInt)
	res := strconv.Itoa(rand)
	return res
}
