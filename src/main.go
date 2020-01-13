package main

import (
	"fmt"
	"jltools"
)

func main() {
	keys := make(map[interface{}]uint32)
	keys["A"] = 2
	keys["B"] = 8
	keys["C"] = 5
	keys["D"] = 7
	keys["E"] = 3
	keys["F"] = 1

	pTree := jltools.CreatHuffman(keys)
	if result,ok := jltools.GetHuffmanCode("F", pTree);ok{
		fmt.Println(result)
	}
}
