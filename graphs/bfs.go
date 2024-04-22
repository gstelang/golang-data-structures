package main

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/gstelang/golang-data-structures/queue"
	"github.com/gstelang/golang-data-structures/sets"
)

type GraphService interface {
	// buildGraph([]string) *Graph
	reconstructPath(path map[string]string, end string) string
	bfs(start string, end string) string
}

type Graph struct {
	graphMap map[string][]string
}

func prettyPrintPath(hashMap map[string]string) {
	b, err := json.MarshalIndent(hashMap, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Print(string(b))
}

func (g *Graph) reconstructPath(path map[string]string, end string) string {

	// prettyPrintPath(path)
	current := path[end]
	result := []string{end}

	for current != "" {
		result = append(result, current)
		current = path[current]
	}
	slices.Reverse(result)

	return strings.Join(result, "->")
}

func (g *Graph) bfs(start string, end string) string {
	queue := queue.Queue[string]{}
	visited := sets.Set[string]{}

	queue.Enqueue(start)
	visited.Add(start)

	path := make(map[string]string)
	path[start] = ""

	for queue.Size() > 0 {
		current, dequeued := queue.Dequeue()
		if dequeued {
			visited.Add(current)
			neighbors, exist := g.graphMap[current]

			if exist {
				for _, neighbor := range neighbors {
					if visited.Has(neighbor) == false {
						visited.Add(neighbor)
						queue.Enqueue(neighbor)
						path[neighbor] = current
					}
				}
			}
		}
	}

	_, pathExists := path[end]
	if pathExists {
		return g.reconstructPath(path, end)
	}

	return ""
}

func main() {

	hashMap := make(map[string][]string)

	hashMap["A"] = []string{"B", "C", "D"}
	hashMap["B"] = []string{"E", "H"}
	hashMap["C"] = []string{"F", "K"}

	graph := &Graph{
		graphMap: hashMap,
	}

	fmt.Println("Path: ", graph.bfs("A", "F"))

}
