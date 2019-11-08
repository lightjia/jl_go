package main

import (
	"fmt"
	"jlalgorithm"
	"math/rand"
)

func main() {
	const arraylen = 0xFF
	array := make([]int, arraylen)
	rand.Seed(99)
	for i := 0; i < arraylen; i++ {
		array[i] = rand.Intn(99)
	}

	fmt.Println("before sort", array)
	tmp := jlalgorithm.Jl_MergeSort(array)
	fmt.Println("after sort", tmp)
}
