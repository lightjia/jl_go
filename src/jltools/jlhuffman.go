package jltools

import (
	"fmt"
	"jlalgorithm"
)

type huffManNode struct {
	value  interface{} //store the value of huffman tree node
	weight uint32      //store the weight of node
}

func (hmn *huffManNode) Print() {
	fmt.Printf("Value:%c Weight:%d\n", hmn.value, hmn.weight)
}

func (hmn *huffManNode) Cmp(h interface{}) int {
	if hmn1, ok := h.(*huffManNode); ok {
		if hmn.weight > hmn1.weight {
			return 1
		} else if hmn.weight == hmn1.weight {
			return 0
		} else {
			return -1
		}
	}

	return -2
}

func GetHuffmanValue(pTree *jlalgorithm.BinaryTree) interface{} {
	if nil != pTree {
		huffman, _ := pTree.Value.(*huffManNode)
		return huffman.value
	}

	return nil
}

func GetHuffmanCode(key interface{}, pRoot *jlalgorithm.BinaryTree) ([]byte, bool) {
	if pRoot != nil {
		ret := make([]byte, 0)
		if pRoot.GetLeft() != nil {
			if node, ok := pRoot.GetLeft().Value.(*huffManNode); ok {
				if node.value == key {
					return append(ret, 0), true
				}
			} else {
				if result, ok := GetHuffmanCode(key, pRoot.GetLeft()); ok {
					ret = append(ret, 0)
					return append(ret, result...), true
				}
			}
		}

		if pRoot.GetRight() != nil {
			if node, ok := pRoot.GetRight().Value.(*huffManNode); ok {
				if node.value == key {
					return append(ret, 1), true
				}
			} else {
				if result, ok := GetHuffmanCode(key, pRoot.GetRight()); ok {
					ret = append(ret, 1)
					return append(ret, result...), true
				}
			}
		}
	}

	return nil, false
}

func CreatHuffman(keys map[interface{}]uint32) *jlalgorithm.BinaryTree {
	var pRet *jlalgorithm.BinaryTree
	var pList *jlalgorithm.LinkList
	for key, value := range keys {
		pList = jlalgorithm.OrderInsert(pList, jlalgorithm.NewLinkList(&huffManNode{key, value}))
	}

	jlalgorithm.Dump(pList)
	for node := pList; node != nil; node = pList {
		node1, _ := node.Value.(*huffManNode) //Get First Node Of The Order LinkList
		if tmpTree1, ok := node1.value.(*jlalgorithm.BinaryTree); ok {
			pRet = tmpTree1 //Judge The Node Is Tree Node
		} else {
			pRet = jlalgorithm.NewBinaryTree(node1)
		}

		if node.GetNext() != nil {
			node2, _ := node.GetNext().Value.(*huffManNode) //Get The Second Node Of The Order LinkList
			tmpTree := jlalgorithm.NewBinaryTree(nil)
			tmpTree.AddLeft(pRet) //The Smaller Weight Node As The Left Child Node
			if tmpTree2, ok := node2.value.(*jlalgorithm.BinaryTree); ok {
				tmpTree.AddRight(tmpTree2)
			} else {
				tmpTree.AddRight(jlalgorithm.NewBinaryTree(node2))
			}

			//Create New Node Weight Is Sum Of The Two Node Weight
			newnode := jlalgorithm.NewLinkList(&huffManNode{tmpTree, node1.weight + node2.weight})
			pList = jlalgorithm.OrderInsert(node.GetNext().GetNext(), newnode)
		} else {
			break
		}
	}

	//Handle Only One Node Of Huffman
	if nil != pRet && nil != pRet.Value {
		tmp := jlalgorithm.NewBinaryTree(nil)
		tmp.AddLeft(pRet)
		pRet = tmp
	}

	return pRet
}
