package Day91

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	ret := make([]string, 0)
	i := 0
	for i < len(words) {
		j := i
		l := len(words[j])
		for j < len(words) {
			if j == len(words)-1 {
				ret = append(ret, genLine(words[i:], maxWidth, l, true))
				i = j + 1
				break
			} else {
				if l+len(words[j+1])+1 > maxWidth {
					ret = append(ret, genLine(words[i:j+1], maxWidth, l, false))
					break
				} else {
					l += len(words[j+1]) + 1
					j++
				}
			}
		}
		i = j + 1
	}
	//for _, s := range ret {
	//	fmt.Printf("[%s]\n", s)
	//}
	return ret
}

func genLine(words []string, maxWidth int, oneSpaceLength int, lastLine bool) string {
	if lastLine {
		words = []string{strings.Join(words, " ")}
	}
	if len(words) == 1 {
		s := words[0]
		for i := 0; i < maxWidth-len(words[0]); i++ {
			s += " "
		}
		//fmt.Printf("one word line (%v)(%v)(%v)(%v)\n", words[0], s, len(words[0]), maxWidth)
		return s
	}

	spaces := maxWidth - oneSpaceLength
	addSpaceNum := spaces / (len(words) - 1)
	addOneMoreSpaceNum := spaces % (len(words) - 1)
	for i := 0; i < len(words)-1; i++ {
		for j := 0; j < addSpaceNum; j++ {
			words[i] += " "
		}
	}
	for i := 0; i < addOneMoreSpaceNum; i++ {
		words[i] += " "
	}
	return strings.Join(words, " ")
}
