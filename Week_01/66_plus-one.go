package Week_01

func plusOne(digits []int) []int {

	l := len(digits)
	if l == 0 {
		return []int{}
	}
	carry := 0
	for i := l - 1; i >= 0; i-- {
		var sum int
		if i == l-1 {
			sum = digits[i] + 1
		} else {
			sum = digits[i] + carry
		}
		digits[i] = sum % 10
		carry = sum / 10
	}

	if carry == 0 {
		return digits
	}
	return append([]int{1}, digits...)

}
