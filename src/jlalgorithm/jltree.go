package jlalgorithm

// Copyright (C) 2019-2020 by lightjia <lightjia@163.com>
//define the data struct
type BinaryTree struct {
	pLeft  *BinaryTree //The Left Branch
	pRight *BinaryTree //The Right Branch
	Value  interface{} //Store The Value Of Tree Node
}

//二叉树节点实现接口
type Bter interface {
	Print() //节点的打印
}

//create the BinaryTree node
func NewBinaryTree(value interface{}) *BinaryTree {
	return &BinaryTree{nil, nil, value}
}

//Add left BinaryTree node
func (pSelf *BinaryTree) AddLeft(pBinaryTree *BinaryTree) {
	pSelf.pLeft = pBinaryTree
}

//Get Left Node
func (pSelf *BinaryTree) GetLeft() *BinaryTree {
	return pSelf.pLeft
}

//Get Right Node
func (pSelf *BinaryTree) GetRight() *BinaryTree {
	return pSelf.pRight
}

//Add right BinaryTree node
func (pSelf *BinaryTree) AddRight(pBinaryTree *BinaryTree) {
	pSelf.pRight = pBinaryTree
}

//PreOrder traversal binary tree
func PreTraversal(pRoot *BinaryTree) {
	if nil != pRoot {
		if v, ok := pRoot.Value.(Bter); ok {
			v.Print()
		}

		PreTraversal(pRoot.pLeft)
		PreTraversal(pRoot.pRight)
	}
}

//非递归先序遍历
func PreTraversal1(pRoot *BinaryTree) {
	stack := make([]*BinaryTree, 0)
	for nil != pRoot {
		if nil != pRoot {
			//访问根节点
			if v, ok := pRoot.Value.(Bter); ok {
				v.Print()
			}

			//右节点先入栈
			if nil != pRoot.pRight {
				stack = append(stack, pRoot.pRight)
			}

			//访问左节点
			pRoot = pRoot.pLeft
		}

		//左节点访问完  右节点出栈
		if nil == pRoot && len(stack) > 0 {
			pRoot = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
	}
}

//Middle order traversal binary tree
func MidTraversal(pRoot *BinaryTree) {
	if nil != pRoot {
		MidTraversal(pRoot.pLeft)
		if v, ok := pRoot.Value.(Bter); ok {
			v.Print()
		}
		MidTraversal(pRoot.pRight)
	}
}

//非递归中序遍历
func MidTraversal1(pRoot *BinaryTree) {
	stack := make([]*BinaryTree, 0)
	for nil != pRoot {
		//根节点入栈
		stack = append(stack, pRoot)
		pRoot = pRoot.pLeft//先访问左节点
		for nil == pRoot && len(stack) > 0 {
			pRoot = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if v, ok := pRoot.Value.(Bter); ok {
				v.Print()
			}

			pRoot = pRoot.pRight
		}
	}
}

//Post order traversal binary tree
func PostTraversal(pRoot *BinaryTree) {
	if nil != pRoot {
		PostTraversal(pRoot.pLeft)
		PostTraversal(pRoot.pRight)
		if v, ok := pRoot.Value.(Bter); ok {
			v.Print()
		}
	}
}

//非递归后序遍历
func PostTraversal1(pRoot *BinaryTree) {
	stack := make([]*BinaryTree, 0)
	//后序遍历需要访问父节点的左右孩子节点  需要辅助HASH标记是否访问过
	m := make(map[*BinaryTree]struct{})
	for nil != pRoot {
		stack = append(stack, pRoot)
		pRoot = pRoot.pLeft
		for nil == pRoot && len(stack) > 0 {
			pRoot = stack[len(stack)-1]
			if _, ok := m[pRoot]; !ok {
				m[pRoot] = struct{}{}
				pRoot = pRoot.pRight
			} else {
				stack = stack[:len(stack)-1]
				if v, ok := pRoot.Value.(Bter); ok {
					v.Print()
				}

				pRoot = nil
			}
		}
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
