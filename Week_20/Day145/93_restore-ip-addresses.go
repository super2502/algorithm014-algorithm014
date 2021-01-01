package Day145

import (
	"strconv"
	"strings"
)

func restoreIpAddresses(s string) []string {
	ret := make([]string, 0)
	var dfs func(path []string, s string, step int)
	dfs = func(path []string, s string, step int) {
		if s == "" && step == 4 {
			ret = append(ret, strings.Join(path, "."))
			return
		}
		if s == "" || step == 4 {
			return
		}
		path = append(path, string(s[0]))
		dfs(path, s[1:], step+1)
		path = path[:len(path)-1]
		if len(s) >= 2 && s[0] != '0' {
			path = append(path, s[:2])
			dfs(path, s[2:], step+1)
			path = path[:len(path)-1]
		}
		if len(s) >= 3 && s[0] != '0' {
			num, _ := strconv.Atoi(s[:3])
			if num <= 255 {
				path = append(path, s[:3])
				dfs(path, s[3:], step+1)
				path = path[:len(path)-1]
			}
		}
	}
	dfs([]string{}, s, 0)
	return ret
}
