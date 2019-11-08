package jlanalyse

import (
	"sync"
)

type jlsortstatic struct {
	iLoopNum uint64
}

var(
pjlsortstatic *jlsortstatic
jlsortonce sync.Once
)

func newSortStatic() {
	pjlsortstatic = new(jlsortstatic)
	pjlsortstatic.iLoopNum=0
}

func (this *jlsortstatic) AddLoopNum(){
	this.iLoopNum++
}

func (this *jlsortstatic) ResetLoopNum(){
	this.iLoopNum = 0
}

func (this *jlsortstatic)GetLoopNum()uint64{
	return this.iLoopNum
}

func GetJlSortStatic() *jlsortstatic {
	jlsortonce.Do(newSortStatic)
	return pjlsortstatic
}

