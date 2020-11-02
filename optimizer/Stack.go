package optimizer

type Stack struct {
	stack []interface{}
}

func (s Stack) Reverse() Stack {
	for i, j := 0, len(s.stack)-1; i < j; i, j = i+1, j-1 {
		s.stack[i], s.stack[j] = s.stack[j], s.stack[i]
	}
	return Stack{s.stack}
}

func (s *Stack) Push(value interface{}) {
	if s.stack == nil {
		s.stack = []interface{}{value}
	} else {
		s.stack = append(s.stack, value)
	}
}

func (s *Stack) Pop() interface{} {
	if s.stack == nil {
		return nil
	} else if len(s.stack) == 0 {
		return nil
	}
	value := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return value
}

func (s Stack) Peek() interface{} {
	if s.stack == nil {
		return nil
	} else if len(s.stack) == 0 {
		return nil
	}
	return s.stack[len(s.stack)-1]
}

func (s Stack) Slice() []interface{} {
	return s.stack
}

func (s Stack) Length() int {
	return len(s.stack)
}
