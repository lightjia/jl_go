package jlalgorithm

import "fmt"

// Copyright (C) 2019-2020 by lightjia <lightjia@163.com>
//define the data struct
type BinaryTree struct {
	pLeft  *BinaryTree //The Left Branch
	pRight *BinaryTree //The Right Branch
	Value  interface{} //Store The Value Of Tree Node
}

//create the BinaryTree node
func NewBinaryTree(value interface{}) *BinaryTree {
	return &BinaryTree{nil, nil, value}
}

//Add left BinaryTree node
func (pSelf *BinaryTree) AddLeft(pBinaryTree *BinaryTree) {
	pSelf.pLeft = pBinaryTree
}

func (pSelf *BinaryTree) GetLeft() *BinaryTree{
	return pSelf.pLeft
}

func (pSelf *BinaryTree) GetRight() *BinaryTree{
	return pSelf.pRight
}

//Add right BinaryTree node
func (pSelf *BinaryTree) AddRight(pBinaryTree *BinaryTree) {
	pSelf.pRight = pBinaryTree
}

//PreOrder traversal binary tree
func PreTraversal(pRoot *BinaryTree) {
	if nil != pRoot {
		fmt.Println(pRoot.Value)
		PreTraversal(pRoot.pLeft)
		PreTraversal(pRoot.pRight)
	}
}

//Middle order traversal binary tree
func MidTraversal(pRoot *BinaryTree) {
	if nil != pRoot {
		MidTraversal(pRoot.pLeft)
		fmt.Println(pRoot.Value)
		MidTraversal(pRoot.pRight)
	}
}

//Post order traversal binary tree
func PostTraversal(pRoot *BinaryTree) {
	if nil != pRoot {
		PostTraversal(pRoot.pLeft)
		PostTraversal(pRoot.pRight)
		fmt.Println(pRoot.Value)
	}
}
//find common parent node
func FindCommonParant(pRoot *BinaryTree, pNodes []*BinaryTree) (int, []*BinaryTree) {
	if nil == pRoot || len(pNodes) <= 0 {
		return 0, nil
	}

	iLeftRet, leftnodes := FindCommonParant(pRoot.pLeft, pNodes)
	iRightRet, rightnodes := FindCommonParant(pRoot.pRight, pNodes)

	var result []*BinaryTree
	iFlag := iLeftRet + iRightRet
	if len(pNodes) == iFlag {
		result = append(result, leftnodes...)
		result = append(result, rightnodes...)
		result = append(result, pRoot)
	} else {
		for i := 0; i < len(pNodes); i++ {
			if pRoot == pNodes[i] {
				iFlag++ //current node as the target node
				break
			}
		}
	}

	return iFlag, result[:]
}