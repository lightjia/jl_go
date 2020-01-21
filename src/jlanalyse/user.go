package jlanalyse

import "math/rand"

type user struct {
	Id     uint64
	Age    uint8
	Name   string
	Passwd string
}

func NewUser() *user {
	n := make([]byte, 120)
	p := make([]byte, 120)
	for i := 0; i < 120; i++ {
		n[i] = 'A' + byte(rand.Uint32()%26)
		p[i] = 'a' + byte(rand.Uint32()%26)
	}

	return &user{uint64(rand.Int63()), uint8(rand.Uint32() % 128), string(n), string(p)}
}
