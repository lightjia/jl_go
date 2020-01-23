package jlalgorithm

func getLpsArray(pat string) []int {
	ret := make([]int, len(pat))
	lps := 0 //最大前缀匹配长度
	//首字母匹配长度为0
	for i := 1; i < len(pat); i++ {
		if pat[i] == pat[lps] {
			lps++
			ret[i] = lps
		} else {
			if lps != 0 {
				lps = ret[lps-1]
				i--
			}
		}
	}

	//fmt.Println(ret)
	return ret
}

//src为源串  pat匹配串  返回值为匹配串在源串中的索引位置
func KmpSearch(src, pat string) int {
	if len(pat) < 1 || len(src) < len(pat) {
		return -1
	}

	iRet := -1
	for i, j, lps := 0, 0, getLpsArray(pat); i < len(src); i++ {
		if src[i] == pat[j] {
			j++
			if j == len(pat) {
				//找到相匹配的
				iRet = i - len(pat) + 1
				break
			}
		} else {
			if j > 0 {
				j = lps[j-1]
				i--
			}
		}
	}

	return iRet
}
