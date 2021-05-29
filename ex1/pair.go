package main

type pair struct {
	v1 int
	v2 int
}

type pairList []pair

func (p pairList) Len() int {
	return len(p)
}

func (p pairList) Less(i, j int) bool {
	return p[i].v1 > p[j].v1
}

func (p pairList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
