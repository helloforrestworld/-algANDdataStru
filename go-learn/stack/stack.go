package stack

type Create []int

func (q *Create) Push(v int) {
	*q = append(*q, v)
}

func (q *Create) Pop() int {
	lastIndex := len(*q) - 1
	if lastIndex < 0 {
		return 0
	}
	trail := (*q)[lastIndex]
	*q = (*q)[:lastIndex]
	return trail
}

func (q *Create) IsEmpty() bool {
	return len(*q) == 0
}
