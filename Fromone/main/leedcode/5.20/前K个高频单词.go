package main

import "sort"

func main() {

}

type T struct {
	name string
	cnt  int
}
type S []T

func topKFrequent(words []string, k int) []string {
	s := make(map[string]int)
	ans := make(S, 0)
	slen := len(words)
	for i := 0; i < slen; i++ {
		s[words[i]]++
	}
	slen = len(s)
	for i, j := range s {
		temp := make(S, 1)
		temp[0].name = i
		temp[0].cnt = j
		ans = append(ans, temp...)
	}
	sort.Sort(ans)
	a := make([]string, 0)
	for i := 0; i < k; i++ {
		a = append(a, ans[i].name)
	}
	return a
}
func (s S) Len() int {
	return len(s)
}
func (s S) Less(i, j int) bool {
	if s[i].cnt > s[j].cnt {
		return s[i].cnt > s[j].cnt
	}
	return s[i].name<s[j].name
}
func (s S) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
