package controllers

import (
	"math/rand"
	"time"
)

func GetRandomN(n int) (num int) {
	rand.Seed(time.Now().Unix())
	num = rand.Int() % n

	return
}
