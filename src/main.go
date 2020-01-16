package main

import "jltools"

func main() {
	jltools.Compress("E:/codeart/jl_go/src/github.com/garyburd/redigo/LICENSE", "F:/tmp.b")
	jltools.UnCompress("F:/tmp.b", "F://tt")
}
