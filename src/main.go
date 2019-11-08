package main

import (
	"fmt"
	"jlalgorithm"
	"jlanalyse"
	"math/rand"
)

func main() {
	const arraylen = 2000000
	array := make([]int, arraylen)
	rand.Seed(99)
	for i:=0;i <arraylen;i++{
		array[i]=rand.Intn(0xfffff)
	}

	//fmt.Println(array)
	fmt.Println(jlanalyse.GetJlSortStatic().GetLoopNum())
	jlalgorithm.Jl_quicksort(array[:], 0, arraylen)
	//fmt.Println(array)
	fmt.Println(jlanalyse.GetJlSortStatic().GetLoopNum())
}
