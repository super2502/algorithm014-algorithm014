package Day11

import "strconv"

func fizzBuzz(n int) []string {
	strOrder := []int{3, 5}
	strs := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}
	ret := make([]string, n)
	for i := 1; i <= n; i++ {
		s := ""
		for _, k := range strOrder {
			if i%k == 0 {
				s += strs[k]
			}
		}
		if s == "" {
			s = strconv.Itoa(i)
		}

		ret[i-1] = s
	}
	return ret
}
