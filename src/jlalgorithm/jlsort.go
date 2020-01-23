package jlalgorithm

func Jl_quicksort(array []int) {
	if len(array) > 1 {
		iStart := 1
		iEnd := len(array) - 1
		for ; iStart < iEnd; iEnd-- {
			if array[iEnd] < array[0] {
				for ; iStart < iEnd; iStart++ {
					if array[iStart] > array[0] {
						array[iStart], array[iEnd] = array[iEnd], array[iStart]
						iStart++
						break
					}
				}
			}
		}

		if array[0] > array[iStart] {
			array[iStart], array[0] = array[0], array[iStart]
		}

		Jl_quicksort(array[0:iStart])
		Jl_quicksort(array[iStart:len(array)])
	}
}
