package main

import (
	"math/rand"
	"time"
)

func main() {
	//This code creates random number and place as variable
	rand.Seed(time.Now().UnixNano())
	hiddenNumber := int64(rand.Intn(100))

	//Insert code here
}
