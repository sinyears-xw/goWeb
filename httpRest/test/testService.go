package test

import "goweb/util"

func s1(name string) []byte {
	return util.Str2Bytes(name)
}

func s2(num []int) int {
	total := 0
	for i := 0; i < len(num); i++ {
		total += num[i]
	}
	return total
}