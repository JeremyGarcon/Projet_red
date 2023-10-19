package main

import (
	"math/rand"
	"time"
)

func randomBool() bool {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 0
}
