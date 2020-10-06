package Week_08

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	arr1 := strings.Split(version1, ".")
	arr2 := strings.Split(version2, ".")
	m := len(arr1)
	n := len(arr2)
	var l int
	if m > n {
		l = m
	} else {
		l = n
	}
	a1 := make([]int, l)
	a2 := make([]int, l)
	for idx, s1 := range arr1 {
		a1[idx], _ = strconv.Atoi(s1)
	}
	for idx, s2 := range arr2 {
		a2[idx], _ = strconv.Atoi(s2)
	}
	for i := 0; i < l; i++ {
		if a1[i] < a2[i] {
			return -1
		}
		if a1[i] > a2[i] {
			return 1
		}
	}
	return 0
}
