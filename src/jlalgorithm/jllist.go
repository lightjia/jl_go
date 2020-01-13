package jlalgorithm

// Copyright (C) 2019-2020 by lightjia <lightjia@163.com>
//define the data struct
type LinkList struct {
	pPrev *LinkList
	pNext *LinkList
	Value interface{}
}

func NewLinkList(value interface{}) *LinkList {
	return &LinkList{nil, nil, value}
}

func (pSelf *LinkList) SetPrev(pNode *LinkList) {
	pSelf.pPrev = pNode
}

func (pSelf *LinkList) GetPrev() *LinkList {
	return pSelf.pPrev
}

func (pSelf *LinkList) SetNext(pNode *LinkList) {
	pSelf.pNext = pNode
}

func (pSelf *LinkList) GetNext() *LinkList {
	return pSelf.pNext
}
