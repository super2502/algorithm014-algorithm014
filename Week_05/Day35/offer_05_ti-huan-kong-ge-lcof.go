package Day35

func replaceSpace(s string) string {
	t := make([]byte, 0)
	for _, b := range []byte(s) {
		if b == ' ' {
			t = append(t, '%', '2', '0')
		} else {
			t = append(t, b)
		}
	}
	return string(t)
}
