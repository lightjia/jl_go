package jltest

import (
	"fmt"
	"jltools"
	"math/rand"
)

func Huffman_Test() {
	const LEN = 888
	array := make([]uint8, LEN)
	keys := make(map[interface{}]uint32)
	for i := 0; i < LEN; i++ {
		array[i] = 'A' + (uint8)(rand.Intn(9999)%10)
		if _, ok := keys[array[i]]; ok {
			keys[array[i]]++
		} else {
			keys[array[i]] = 1
		}
	}

	pTree := jltools.CreatHuffman(keys)
	if nil != pTree {
		for i := 0; i < 10; i++ {
			key := (uint8)('A' + i)
			weight, _ := keys[key]
			if code, ok := jltools.GetHuffmanCode(key, pTree); ok {
				fmt.Printf("%c:%v=>%v\n", key, weight, code)
			}
		}
	}
}
