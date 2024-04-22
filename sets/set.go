package sets

type Set[T any] struct {
	items map[any]bool
}

func (s *Set[T]) Add(item T) {
	if s.items == nil {
		s.items = make(map[any]bool)
	}
	s.items[item] = true
}

func (s *Set[T]) Remove(item T) {
	delete(s.items, item)
}

func (s *Set[T]) Size() int {
	return len(s.items)
}

func (s *Set[T]) Has(item T) bool {
	_, exists := s.items[item]
	return exists
}

// func main() {
// 	seq := Set[string]{}
// 	seq.Add("abc")
// 	seq.Add("xyz")
// 	fmt.Println("set Size", seq.Size())
// 	fmt.Println("set Has abc? ", seq.Has("abc"))
// 	fmt.Println("Removing abc")
// 	seq.Remove("abc")
// 	fmt.Println("set Has abc? ", seq.Has("abc"))
// 	fmt.Println("set Size", seq.Size())

// 	seqInts := Set[int]{}
// 	seqInts.Add(1)
// 	seqInts.Add(2)
// 	fmt.Println("set Size", seqInts.Size())
// 	fmt.Println("set Has 1?", seqInts.Has(1))
// 	fmt.Println("Removing 1")
// 	seqInts.Remove(1)
// 	fmt.Println("set Has 1? ", seqInts.Has(1))
// 	fmt.Println("set Size", seqInts.Size())
// }
