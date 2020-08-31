package Week_04

import (
	"container/list"
	"fmt"
)

func minMutation(start string, end string, bank []string) int {
	sampleMap := make(map[string]struct{})
	for _, sample := range bank {
		sampleMap[sample] = struct{}{}
	}
	if _, ok := sampleMap[end]; !ok {
		return -1
	}
	visited := make(map[string]struct{})

	deque := list.New()
	deque.PushBack(&levelNode433{
		val:   start,
		level: 0,
	})

	for deque.Len() > 0 {
		cn := deque.Front().Value.(*levelNode433)
		deque.Remove(deque.Front())
		for i := 0; i < 8; i++ {
			fmt.Printf("===test bs (%s)\n", string(cn.val))
			for _, b := range []byte("ACGT") {
				if b == cn.val[i] {
					continue
				}
				word := cn.val[:i] + string(b) + cn.val[i+1:]
				fmt.Printf("test (%s) -> (%s)\n", cn.val, word)
				if word == end {
					return cn.level + 1
				}
				_, ok1 := sampleMap[word]
				_, ok2 := visited[word]
				if ok1 && !ok2 {
					fmt.Printf("push node into queue (%v)\n", word)
					deque.PushBack(&levelNode433{
						val:   word,
						level: cn.level + 1,
					})
				}
				visited[word] = struct{}{}
			}
		}
	}

	return -1
}

type levelNode433 struct {
	val   string
	level int
}
