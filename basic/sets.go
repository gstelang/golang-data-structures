package basic

import "fmt"

type Set[T any] struct {
	items map[any]bool
}

func (s *Set[T]) add(item T) {
	if s.items == nil {
		s.items = make(map[any]bool)
	}
	s.items[item] = true
}

func (s *Set[T]) remove(item T) {
	delete(s.items, item)
}

func (s *Set[T]) size() int {
	return len(s.items)
}

func (s *Set[T]) has(item T) bool {
	_, exists := s.items[item]
	return exists
}

func main() {
	seq := Set[string]{}
	seq.add("abc")
	seq.add("xyz")
	fmt.Println("set size", seq.size())
	fmt.Println("set has abc? ", seq.has("abc"))
	fmt.Println("Removing abc")
	seq.remove("abc")
	fmt.Println("set has abc? ", seq.has("abc"))
	fmt.Println("set size", seq.size())

	seqInts := Set[int]{}
	seqInts.add(1)
	seqInts.add(2)
	fmt.Println("set size", seqInts.size())
	fmt.Println("set has 1?", seqInts.has(1))
	fmt.Println("Removing 1")
	seqInts.remove(1)
	fmt.Println("set has 1? ", seqInts.has(1))
	fmt.Println("set size", seqInts.size())
}
