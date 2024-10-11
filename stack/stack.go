package stack

type Stack struct {
	items []string
	len   int
}

func (st *Stack) Push(s string) bool {
	st.items = append(st.items, s)
	st.len++
	return true
}

func (st *Stack) Pop() string {
	if st.len > 0 {
		lastItem := st.items[st.len-1]
		st.items = st.items[0 : st.len-1]
		st.len--
		return lastItem
	}
	return ""
}

func (st *Stack) Top() string {
	if st.IsEmpty() {
		return ""
	}
	return st.items[st.len-1]
}

func (st *Stack) IsEmpty() bool {
	return st.len == 0
}

// ================== Sample Problem: (Leetcode) 20. Valid Parenthesis ==================

func IsValidParenthesis(s string) bool {
	st := &Stack{}

	for i := 0; i < len(s); i++ {
		if st.IsEmpty() {
			st.Push(s[i : i+1])
			continue
		}

		if st.Top() == "(" && s[i:i+1] == ")" {
			st.Pop()
			continue
		}

		if st.Top() == "{" && s[i:i+1] == "}" {
			st.Pop()
			continue
		}

		if st.Top() == "[" && s[i:i+1] == "]" {
			st.Pop()
			continue
		}

		st.Push(s[i : i+1])
	}

	return st.IsEmpty()
}
