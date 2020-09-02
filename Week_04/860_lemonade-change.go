package Week_04

func lemonadeChange(bills []int) bool {
	five, ten := 0, 0

	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else if bill == 20 {
			if ten > 0 {
				if five == 0 {
					return false
				}
				ten--
				five--
			} else if five < 3 {
				return false
			} else {
				five -= 3
			}
		}
	}
	return true
}
