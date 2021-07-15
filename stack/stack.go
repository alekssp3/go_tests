package stack

type Stack struct {
	data []string
}

type Stacker interface {
	Push(string)
	Pop() string
	Last() string
	len() int
}

func (st *Stack) Push(s string) {
	st.data = append(st.data, s)
}

func (st *Stack) Pop() string {
	val := st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]
	return val
}

func (st *Stack) len() int {
	return len(st.data)
}

func (st *Stack) Last() string {
	if st.len() > 0 {
		return st.data[len(st.data)-1]
	} else {
		return ""
	}
}
