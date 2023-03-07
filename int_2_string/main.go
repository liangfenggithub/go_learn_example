package main

import "strconv"

func main() {

	// 10 => 'A'

	i := int64(10)
	s := strconv.FormatInt(i, 16)
	println(s)
}
