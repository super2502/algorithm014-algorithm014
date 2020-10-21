package Week_10

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	ret := make([][]string, 0)
	emailMap := make(map[string]string)
	for _, account := range accounts {
		for i := 1; i < len(account); i++ {
			emailMap[account[i]] = account[0]
		}
	}
	emails := make([]string, 0, len(emailMap))
	for email := range emailMap {
		emails = append(emails, email)
	}
	uf := &unionFind{}
	uf.init(emails)

	for _, account := range accounts {
		for i := 1; i < len(account)-1; i++ {
			uf.union(account[i], account[i+1])
		}
	}

	bossMap := make(map[string][]string)
	for _, email := range emails {
		bossMap[uf.find(email)] = append(bossMap[uf.find(email)], email)
	}

	for boss, emails := range bossMap {
		user := emailMap[boss]
		sort.Strings(emails)
		ret = append(ret, append([]string{user}, emails...))
	}

	return ret
}

type unionFind struct {
	parents map[string]string
}

func (uf *unionFind) init(emails []string) {
	uf.parents = make(map[string]string)
	for i := 0; i < len(emails); i++ {
		uf.parents[emails[i]] = emails[i]
	}
}
func (uf *unionFind) find(email string) string {
	tmp := email
	for uf.parents[email] != email {
		email = uf.parents[email]
	}
	for uf.parents[tmp] != email {
		tmp1 := uf.parents[tmp]
		uf.parents[tmp] = email
		tmp = tmp1
	}
	return email
}
func (uf *unionFind) union(email1, email2 string) {
	boss1 := uf.find(email1)
	boss2 := uf.find(email2)
	if boss1 == boss2 {
		return
	}
	uf.parents[boss1] = boss2
}
