package main

import (
	"fmt"
	"sort"
)

func main() {
	a := Personal{{"chong",7},{"465",13},{"wang",14},{"12312",15},{"8987",12}}
	fmt.Println(a)
	sort.Sort(a)
	fmt.Println(a)

	n := sort.Search(len(a), func(i int) bool {
		return a[i].age<=15
	})
	fmt.Println(n)
}

type personal struct {
	name string
	age int
}
type Personal []personal
func (c Personal) Len() int {
	return len(c)
}
func (c Personal) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}
func (c Personal) Less(i, j int) bool {
	return c[i].age > c[j].age
}
