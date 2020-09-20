package Week_06

func minWindow(s string, t string) string {
	needMap := make(map[byte]int)
	needs := len(t)
	for _, b := range []byte(t) {
		needMap[b]++
	}
	//fmt.Printf("%s, %s\n",s, t)
	var i, i0, j0 int
	minLength := len(s) + 1
	for j := 0; j < len(s); j++ {
		b := s[j]
		if _, ok := needMap[b]; !ok {
			continue
		}
		//fmt.Printf("j, b, need (%v)(%v)(%v)(%+v)\n", j, b, needs, needMap)
		needMap[b]--
		if needMap[b] >= 0 {
			needs--
		}
		for needs == 0 {
			//fmt.Printf("i is (%v)\n", i)
			if j-i < minLength {
				minLength = j - i
				i0, j0 = i, j
				//fmt.Printf("got min i,j (%v)(%v)\n", i , j)
			}
			b := s[i]
			if _, ok := needMap[b]; !ok {
				i++
				continue
			}
			needMap[b]++
			//fmt.Printf("drop from i (%v)(%v), count(%v)\n", i, b, needMap[b])
			if needMap[b] <= 0 {
				//if j-i < minLength {
				//	minLength = j - i
				//	i0, j0 = i, j
				//	//fmt.Printf("got min i,j (%v)(%v)\n", i , j)
				//}
			} else {
				needs++
			}
			i++
		}
	}
	if minLength == len(s)+1 {
		return ""
	}
	return s[i0 : j0+1]
}
