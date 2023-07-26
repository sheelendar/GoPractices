package stringStack

type Stack struct {
	stack []byte
}

func (s *Stack) Push(item uint8) {
	s.stack = append(s.stack, item)
}
func (s *Stack) Pop() uint8 {
	size := len(s.stack)
	if size > 0 {
		temp := s.stack[size-1]
		s.stack = s.stack[0 : size-1]
		if size == 1 {
			s.stack = nil
		}
		return temp
	}
	return '0'
}
func (s *Stack) Top() uint8 {
	size := len(s.stack)
	if size > 0 {
		return s.stack[size-1]
	}
	return '0'
}
func (s *Stack) IsEmpty() bool {
	size := len(s.stack)
	if size <= 0 {
		return true
	}
	return false
}
