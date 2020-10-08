package Week_08

import "fmt"

func isUnique(astr string) bool {
	fmt.Printf("astr: %v\n", astr)
	var bitMap int
	for i := 0; i < len(astr); i++ {
		b := astr[i]
		bbit := 1 << int(astr[i]-'a')
		fmt.Printf("%b, %b, %v, %v\n", bitMap, bbit, bitMap, b)
		if bitMap&bbit > 0 {
			return false
		}
		bitMap += bbit
	}
	return true
}
