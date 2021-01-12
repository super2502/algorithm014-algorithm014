package Day142

import "strings"

func simplifyPath(path string) string {
	stack := make([]string, 0)
	paths := strings.Split(path, "/")
	for _, p := range paths {
		switch p {
		case ".", "":
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, p)
		}
	}
	return "/" + strings.Join(stack, "/")
}
