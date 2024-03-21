package basic

import "fmt"

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) pop() (T, bool) {
	if len(s.items) == 0 {
		// Go does not have a concept of empty, null, or undefined for variable values.
		// Variables declared without an explicit initial value default to the zero value for their respective type.
		// The zero value for primitive types such as booleans, numeric types, and strings are false , 0 , and "" , respectively.
		var zero T
		return zero, false
	}
	lastIndex := len(s.items) - 1
	val := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return val, true
}

func (s *Stack[T]) size() int {
	return len(s.items)
}

func (s *Stack[T]) peek() T {
	return s.items[len(s.items)-1]
}

func main() {
	seq := Stack[string]{}
	val0, check0 := seq.pop()
	fmt.Println(val0)
	fmt.Println(check0)
	seq.push("abc")
	seq.push("xyz")
	fmt.Println("stack size:", seq.size())
	fmt.Println("stack peek:", seq.peek())
	val, check := seq.pop()
	fmt.Println(val)
	fmt.Println(check)
	fmt.Println("stack size:", seq.size())

	seqInts := Stack[int]{}
	valInt, checkInt := seqInts.pop()
	fmt.Println(valInt)
	fmt.Println(checkInt)
	seqInts.push(1)
	seqInts.push(2)
	fmt.Println("stack size:", seqInts.size())
	fmt.Println("stack peek:", seqInts.peek())
	valIntPopped, checkPopped := seqInts.pop()
	fmt.Println(valIntPopped)
	fmt.Println(checkPopped)
	fmt.Println("stack size:", seq.size())

}
