package main

type stack struct {
	len  int
	top  int
	data []int
}

func main()  {

}
func initStack(cap int) *stack {
	s := new(stack)
	s.len = cap
	s.top = -1

	return s

}

func (s *stack) isFull() {

}
func (s *stack) push(item int) bool {
	if s.top+1 == s.len {
		return false
	}

	s.top++
	s.data[s.top] = item

	return true
}

func (s *stack) pop() int {
	if s.top == -1 {
		return 0
	}
	res := s.data[s.top]
	s.top--

	return res
}

