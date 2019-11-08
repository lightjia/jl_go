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

func jl_merge(arr1, arr2 []int) (r []int) {
	r = make([]int, len(arr1)+len(arr2))
	i, i1, i2 := 0, 0, 0
	for ; i < len(r); i++ {
		if i1 < len(arr1) {
			if (i2 < len(arr2) && arr1[i1] < arr2[i2]) || i2>=len(arr2) {
				r[i] = arr1[i1]
				i1++
				continue
			}
		}

		r[i] = arr2[i2]
		i2++
	}

	return r
}

func Jl_MergeSort(array []int) []int {
	if (len(array) >= 2) {
		r1 := Jl_MergeSort(array[:len(array)/2])
		r2 := Jl_MergeSort(array[len(array)/2:])

		return jl_merge(r1, r2)
	}

	return array
}
