package main

import (
	"fmt"
	"strings"

	"github.com/go-cheatsheat/utils"
)

type Vertex struct {
	x int
	y int
}

func demoArray() {
	val_nums := []int{1, 2, 3, 4, 5}
	val_strings := []string{"a", "b", "c", "d", "e"}
	val_vertexs := []Vertex{{1, 2}, {2, 3}, {3, 4}}

	// Map
	fmt.Println(utils.Map(val_nums, func(c int, _ int) int { return c * c }))
	fmt.Println(utils.Map(val_strings, func(c string, _ int) string { return strings.ToUpper(c) }))
	fmt.Println(utils.Map(val_vertexs, func(c Vertex, _ int) int { return c.x + c.y }))

	// ForEach
	fmt.Println(val_vertexs)
	utils.ForEach(val_vertexs, func(c Vertex, index int) {
		val_vertexs[index] = Vertex{c.x * 2, c.y * 2}
	})
	fmt.Println(val_vertexs)

	// Filter
	fmt.Println(utils.Filter(val_nums, func(c int, _ int) bool { return c%2 == 0 }))
	fmt.Println(utils.Filter(val_strings, func(c string, _ int) bool { return c == "d" }))
	fmt.Println(utils.Filter(val_vertexs, func(c Vertex, _ int) bool { return c.x%4 == 0 }))

	// Reduce
	fmt.Println(utils.Reduce(val_nums, func(a int, b int) int {
		return a + b
	}, 0))
	fmt.Println(utils.Reduce(val_vertexs, func(a int, b Vertex) int {
		return a + b.x
	}, 0))

	// IndexOf
	fmt.Println(utils.IndexOf(val_nums, 4))
	fmt.Println(utils.IndexOf(val_strings, "b"))

	// ToMap
	fmt.Println(utils.ToMap(val_nums, func(c int, index int) (int, int) {
		return index, c
	}))
	fmt.Println(utils.ToMap(val_vertexs, func(c Vertex, index int) (int, Vertex) {
		return index, c
	}))
	fmt.Println(utils.ToMap(val_vertexs, func(c Vertex, _ int) (int, int) {
		return c.x, c.y
	}))
}

func main() {
	demoArray()
}
