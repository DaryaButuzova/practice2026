package stack

import "testing"

func TestStack(t *testing.T) {

	st := NewStack[int](30)

	st.Push(10)
	st.Push(20)

	if st.Peek() != 20 {
		t.Errorf("ожидали 20, получили %v", st.Peek())
	}

	if st.Pop() != 20 {
		t.Errorf("ожидали 20 при pop")
	}

	if st.Pop() != 10 {
		t.Errorf("ожидали 10 при pop")
	}
}
