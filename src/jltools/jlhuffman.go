package jltools

import (
	"jlalgorithm"
)

type huffmannode struct {
	value  interface{}	//store the value of huffman tree node
	weight uint32	//store the weight of node
}

func insertnode(pList, pNode *jlalgorithm.LinkList) *jlalgorithm.LinkList {
	pRet := jlalgorithm.NewLinkList(nil)
	pRet.SetNext(pList)
	if nil != pNode {
		node, _ := pNode.Value.(huffmannode)
		if nil != pList {
			tmp := pRet
			for ; tmp.GetNext() != nil; tmp = tmp.GetNext() {
				tmpnode, _ := tmp.GetNext().Value.(huffmannode)
				if tmpnode.weight >= node.weight {
					if tmpnode.weight == node.weight{
						switch tmpnode.value.(type) {
						case byte:
							b1,_:=tmpnode.value.(byte)
							b2,_:=node.value.(byte)
							if b1>b2{
								continue
							}
						}
					}

					pNode.SetNext(tmp.GetNext())
					tmp.SetNext(pNode)
					break
				}
			}

			if tmp.GetNext() == nil {
				tmp.SetNext(pNode)
			}
		} else {
			pRet.SetNext(pNode)
		}
	}

	return pRet.GetNext()
}

func GetHuffmanValue(pTree *jlalgorithm.BinaryTree)interface{}{
	if nil!= pTree{
		huffman,_ := pTree.Value.(huffmannode)
		return huffman.value
	}

	return nil
}

func GetHuffmanCode(key interface{}, pRoot *jlalgorithm.BinaryTree) ([]byte, bool) {
	if pRoot != nil {
		ret := make([]byte, 0)
		if pRoot.GetLeft() != nil {
			if node, ok := pRoot.GetLeft().Value.(huffmannode); ok {
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
			if node, ok := pRoot.GetRight().Value.(huffmannode); ok {
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
	for key, value := range (keys) {
		node := huffmannode{key, value}
		pList = insertnode(pList, jlalgorithm.NewLinkList(node))
	}

	for node := pList; node != nil; node = pList {
		node1, _ := node.Value.(huffmannode)//Get First Node Of The Order LinkList
		if tmpTree1, ok := node1.value.(*jlalgorithm.BinaryTree); ok {
			pRet = tmpTree1	//Judge The Node Is Tree Node
		} else {
			pRet = jlalgorithm.NewBinaryTree(node1)
		}

		if node.GetNext() != nil {
			node2, _ := node.GetNext().Value.(huffmannode)	//Get The Second Node Of The Order LinkList
			tmpTree := jlalgorithm.NewBinaryTree(nil)
			tmpTree.AddLeft(pRet)	//The Smaller Weight Node As The Left Child Node
			if tmpTree2, ok := node2.value.(*jlalgorithm.BinaryTree); ok {
				tmpTree.AddRight(tmpTree2)
			} else {
				tmpTree.AddRight(jlalgorithm.NewBinaryTree(node2))
			}

			//Create New Node Weight Is Sum Of The Two Node Weight
			newnode := jlalgorithm.NewLinkList(huffmannode{tmpTree, node1.weight + node2.weight})
			pList = insertnode(node.GetNext().GetNext(), newnode)
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
