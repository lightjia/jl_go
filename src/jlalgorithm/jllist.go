package jlalgorithm

// Copyright (C) 2019-2020 by lightjia <lightjia@163.com>
//define the data struct
type LinkList struct {
	pPrev *LinkList	//Previous node pointer
	pNext *LinkList	//Next node pointer
	Value interface{}	//Store The Value of Node
}

//Create Node
func NewLinkList(value interface{}) *LinkList {
	return &LinkList{nil, nil, value}
}

//Set Previous node pointer
func (pSelf *LinkList) SetPrev(pNode *LinkList) {
	pSelf.pPrev = pNode
}

//Get Previous node pointer
func (pSelf *LinkList) GetPrev() *LinkList {
	return pSelf.pPrev
}

//Set Next node pointer
func (pSelf *LinkList) SetNext(pNode *LinkList) {
	pSelf.pNext = pNode
}

//Get Next node pointer
func (pSelf *LinkList) GetNext() *LinkList {
	return pSelf.pNext
}
