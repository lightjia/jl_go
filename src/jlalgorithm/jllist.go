package jlalgorithm

// Copyright (C) 2019-2020 by lightjia <lightjia@163.com>
//define the data struct
type LinkList struct {
	pPrev *LinkList   //Previous node pointer
	pNext *LinkList   //Next node pointer
	Value interface{} //Store The Value of Node
}

//链表节点需要实现的接口
type LinkLister interface {
	Cmp(v interface{}) int //链表节点比较函数 1表示大于 0表示等于  -1表示小于
	Print()                //链表节点打印函数
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

//递归实现两有序链表合并 也可以实现有序链表插入节点   当链表较长时，递归耗用堆栈空间大，需改为非递归优化
func OrderInsert(root, node *LinkList) *LinkList {
	if nil == root {
		return node
	}

	if nil == node {
		return root
	}

	pRet := root
	if r, ok := root.Value.(LinkLister); ok {
		if r.Cmp(node.Value) < 0 {
			root.pNext = OrderInsert(root.pNext, node)
		} else {
			pRet = node
			node.pNext = OrderInsert(root, node.pNext)
		}
	}

	return pRet
}

//链表的遍历打印
func Dump(root *LinkList) {
	for root != nil {
		if r, ok := root.Value.(LinkLister); ok {
			r.Print()
		}

		root = root.pNext
	}
}
