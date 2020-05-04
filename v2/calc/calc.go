package calc

//Add  , returns  a plus b.
func Add(args ...int) int {
	s := 0
	for _, v := range args {
		s += v
	}
	return s
}

//Sub  , returns a minus b.
func Sub(a, b int) int {
	return a - b
}
