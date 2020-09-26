package Day47

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	ret := make([][]string, 0)
	frontWordMap := make(map[string]struct{})
	backWordMap := make(map[string]struct{})
	for _, word := range wordList {
		frontWordMap[word] = struct{}{}
		backWordMap[word] = struct{}{}
	}
	if _, ok := backWordMap[endWord]; !ok {
		return ret
	}
	delete(backWordMap, endWord)
	frontMap := map[string][]string{beginWord: {beginWord}}
	backMap := map[string][]string{endWord: {endWord}}
	forward := 1
	meet := false
	for len(frontMap) > 0 {
		//fmt.Printf("front: %+v, back %+v\n", frontMap, backMap)
		tmp := make(map[string][]string)
		for word, frontList := range frontMap {
			//if word =="maris" {
			//	//fmt.Printf(" ==== maris with back (%+v)\n", backMap)
			//}
			for i := 0; i < len(word); i++ {
				for b := 'a'; b < 'z'; b++ {
					key := word[:i] + string(b) + word[i+1:]
					if backList, ok := backMap[key]; ok {
						//fmt.Printf("we meet (%v)\n", key)
						meet = true
						f := make([]string, len(frontList))
						copy(f, frontList)
						b := make([]string, len(backList))
						copy(b, backList)
						if forward == 1 {
							r := append(f, b...)
							ret = append(ret, r)
						} else {
							r := append(b, f...)
							ret = append(ret, r)
						}
					}
					if _, ok := frontWordMap[key]; !ok {
						continue
					}
					if _, ok := frontMap[key]; ok {
						continue
					}
					//if word =="maris" {
					//	fmt.Printf(" ==== maris find key (%+v)(%+v)\n", key, frontList)
					//}
					f := make([]string, len(frontList))
					copy(f, frontList)
					if forward == 1 {
						tmp[key] = append(f, key)
					} else {
						tmp[key] = append([]string{key}, f...)
					}
					//if word =="maris" {
					//	fmt.Printf(" ==== maris after insert key (%+v)(%+v)\n", key, tmp[key])
					//}
					delete(frontWordMap, key)
				}
			}
		}
		if meet {
			break
		}
		frontMap = tmp
		//fmt.Printf("==== fk what is frontMap %+v", frontMap)
		if len(frontMap) > len(backMap) {
			frontMap, backMap = backMap, frontMap
			forward *= -1
			frontWordMap, backWordMap = backWordMap, frontWordMap
			//fmt.Printf("switch cause \n")
		}
	}
	return ret
}
