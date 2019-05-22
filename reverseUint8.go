package main

import (
	"fmt"
)

func main() {
	var test1 uint8
	test1 = 0x52
	reversed := Reverse(test1)
}

func Reverse(num uint8) uint8 {
	var s uint8 = 8

var mask uint8
	mask = 0xFF
	for {
		s >>= 1
	 	if s <= 0 {
		 	break
		}
		mask ^= (mask << s)
		num = ((num >> s) & mask) | ((num << s) & ^mask)
	}
	return num
}
