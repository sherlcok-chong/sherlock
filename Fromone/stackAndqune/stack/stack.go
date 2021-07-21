package stack

type Stack []interface{}

func (s Stack) Pop() ([]interface{}, interface{}) {
	val := s[len(s)-1]
	return s[:len(s)-1], val
}
func (s Stack) Push(val interface{}) []interface{} {
	return append(s, val)
}
func (s Stack)Empty() bool {
	return !(len(s)>0)
}
func (s Stack)Top() interface{} {
	return s[len(s)-1]
}
