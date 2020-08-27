package Day16

func replaceSpace(s string) string {

	bs := []byte(s)
	t := make([]byte, 0)
	for _, b := range bs {
		if b == ' ' {
			t = append(t, '%', '2', '0')
		} else {
			t = append(t, b)
		}
	}
	return string(t)
}
