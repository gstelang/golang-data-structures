package main

import "fmt"

type set struct {
	items map[string]bool
}

func (s *set) add(item string) {
	if s.items == nil {
		s.items = make(map[string]bool)
	}
	s.items[item] = true
}

func (s *set) remove(item string) {
	delete(s.items, item)
}

func (s set) size() int {
	return len(s.items)
}

func (s set) has(item string) bool {
	_, exists := s.items[item]
	return exists
}

func main() {
	seq := set{}
	seq.add("abc")
	seq.add("xyz")
	fmt.Println("set size", seq.size())
	fmt.Println("set has abc? ", seq.has("abc"))
	fmt.Println("Removing abc")
	seq.remove("abc")
	fmt.Println("set has abc? ", seq.has("abc"))
	fmt.Println("set size", seq.size())
}
