package Day156

type Employee struct {
	Id           int
	Importance   int
	Subordinates []int
}

func getImportance(employees []*Employee, id int) int {
	em := make(map[int]*Employee)
	for _, e := range employees {
		em[e.Id] = e
	}
	boss := em[id]
	queue := []*Employee{boss}
	ret := 0
	for len(queue) > 0 {
		tmp := make([]*Employee, 0)
		for _, e := range queue {
			ret += e.Importance
			for _, sub := range e.Subordinates {
				tmp = append(tmp, em[sub])
			}
		}
		queue = tmp
	}
	return ret
}
