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
	debug   = "abcde"
)

func findDuplicateLetters(slice []byte) map[byte][]int {
	positions := make(map[byte][]int)

	for i, letter := range slice {
		if _, ok := positions[letter]; !ok {
			positions[letter] = []int{i}
		} else {
			positions[letter] = append(positions[letter], i)
		}
	}

	// Remover a primeira ocorrência de cada letra
	for letter, pos := range positions {
		if len(pos) <= 1 {
			delete(positions, letter)
		} else {
			positions[letter] = pos[1:]
		}
	}

	return positions
}

func main() {

	var sizePassw = 10
	var password = make([]byte, sizePassw)
	var charSource string
	charSource = lower + upper + nums + special
	var r = len(charSource) - 1

	//gera a senha
	for i := 0; i < sizePassw; i++ {
		randNum := rand.Intn(r)
		password[i] = byte(randNum)
	}

	//verifica letras duplicadas e substitui as letras duplicadas por novas letras
	for {
		mapDup := findDuplicateLetters(password)
		fmt.Println(mapDup)
		fmt.Println(password)

		for key, value := range mapDup {
			fmt.Println(key, value)
			for i := 0; i < len(value); i++ {
				var j = value[i]
				fmt.Println(j)
				randNum := rand.Intn(r)
				fmt.Println(randNum)
				password[j] = byte(randNum)
				fmt.Println(password[j])
			}
		}
		var dup bool
		if len(mapDup) == 0 {
			dup = false
		} else {
			dup = true
		}

		if dup == false {
			break
		}
	}

	//imprime a senha
	fmt.Println(password)
	for i := 0; i < len(password); i++ {
		var j = password[i]
		fmt.Printf("%c", charSource[j])
	}
	fmt.Printf("\n")

}
