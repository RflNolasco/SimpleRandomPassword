package main

import (
	"fmt"
	"math/rand"
)

const (
	lower   = "abcdefghijklmnopqrstuvwxyz"
	upper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	nums    = "0123456789"
	special = "!@#$%&-_+=.,/?~"
)

func main() {

	var sizePassw = 12
	var password = make([]byte, sizePassw)
	var charSource string
	charSource = lower + upper + nums + special
	var r = len(charSource) - 1

	for i := 0; i < sizePassw; i++ {
		randNum := rand.Intn(r)
		password[i] = byte(randNum)
	}

	fmt.Println(password)

	for i := 0; i < len(password); i++ {
		var j = password[i]
		fmt.Printf("%c", charSource[j])
	}
	fmt.Printf("\n")

}
