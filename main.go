package main

import (
	"fmt"
	"uasgo/astar"
	"uasgo/data"
)

func main() {
	start, goal := data.LoadGraph()

	path := astar.AStar(start, goal)

	fmt.Println("Jalur terpendek (A*):")
	for i, node := range path {
		fmt.Printf("%d. %s\n", i+1, node.Name)
	}
}
