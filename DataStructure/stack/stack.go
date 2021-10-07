package stack

/**
@author Shitanford
@create 2021-04-15 11:11
*/

type Stack struct {
	data []interface{}
	len int
}

func New(len int) Stack {
	return Stack{data: make([]interface{}, len),
				len: len}
}
func (s Stack) Pop() interface{} {
	s.len -= 1
	return s.data[s.len]
}

func (s Stack) Push(d interface{}) {
	s.data[s.len] = d
	s.len += 1
}

func (s Stack) Top() interface{} {
	return s.data[s.len-1]
}

func (s Stack) Empty() bool {
	return s.len == 0
}