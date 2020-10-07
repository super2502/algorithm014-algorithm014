package Day59

func reverseStr(s string, k int) string {
	bs := []byte(s)
	for i := 0; i < len(bs); i = i + 2*k {
		end := i + k - 1
		if end >= len(bs) {
			end = len(bs) - 1
		}
		reverse(bs, i, end)
	}
	return string(bs)
}
func reverse(bs []byte, i, j int) {
	for i < j {
		bs[i], bs[j] = bs[j], bs[i]
		i++
		j--
	}
}
