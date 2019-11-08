package jlalgorithm

import "jlanalyse"

func Jl_quicksort(array []int, start, end int) {
	if start >= end-1 {
		return
	}

	imid := start //哨兵参考值
	istart := start + 1
	iend := end - 1
	for istart < iend {
		for ; iend > istart; iend-- { //找出右边小于参考值的index
			jlanalyse.GetJlSortStatic().AddLoopNum()
			if array[imid] > array[iend] {
				break
			}
		}

		for ; istart < iend; istart++ { //找出左边大于参考值的index
			jlanalyse.GetJlSortStatic().AddLoopNum()
			if array[istart] > array[imid] {
				break
			}
		}

		if istart < iend { //左右交换
			array[istart], array[iend] = array[iend], array[istart]
		}
	}

	if array[imid] > array[istart] { //参考值归位
		array[istart], array[imid] = array[imid], array[istart]
	}

	Jl_quicksort(array[:], imid, istart) //递归进行左侧排序
	Jl_quicksort(array[:], istart, end)  //递归进行右侧排序
}
