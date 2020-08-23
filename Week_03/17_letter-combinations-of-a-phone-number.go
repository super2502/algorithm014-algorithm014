package Week_03

var sret []string
var numChars map[int][]byte

func letterCombinations(digits string) []string {
	l := len(digits)
	if l == 0 {
		return []string{}
	}
	nums := make([]int, l)
	for i := 0; i < l; i++ {
		nums[i] = int(digits[i] - '0')
	}
	sret = make([]string, 0)
	numChars = map[int][]byte{
		2: {'a', 'b', 'c'},
		3: {'d', 'e', 'f'},
		4: {'g', 'h', 'i'},
		5: {'j', 'k', 'l'},
		6: {'m', 'n', 'o'},
		7: {'p', 'q', 'r', 's'},
		8: {'t', 'u', 'v'},
		9: {'w', 'x', 'y', 'z'},
	}
	_dfs17(nums, []byte{})
	return sret
}

func dfs17(path []int, nums []int, ret []byte) {
	//fmt.Printf("called with path(%+v), ret(%+v)\n", path, ret)
	if len(path) == len(nums) {
		sret = append(sret, string(ret))
		//fmt.Printf("returned (%+v)\n", string(ret))
		return
	}
	num := nums[len(path)]
	bs := numChars[num]
	path = append(path, num)
	for _, b := range bs {
		retCopy := make([]byte, len(ret))
		copy(retCopy, ret)
		retCopy = append(retCopy, b)
		//fmt.Printf("call next %+v (%+v)\n", path, retCopy)
		dfs17(path, nums, retCopy)
	}
}

func _dfs17(nums []int, ret []byte) {
	if len(ret) == len(nums) {
		sret = append(sret, string(ret))
		return
	}
	num := nums[len(ret)]
	bs := numChars[num]

	for _, b := range bs {
		ret = append(ret, b)
		_dfs17(nums, ret)
		ret = ret[:len(ret)-1]
	}
}
