package astar

import "uasgo/data"

type AStarNode struct {
	Node    *data.Node
	G, H, F float64
	Parent  *AStarNode
}

// heuristic sederhana
func heuristic(a, b *data.Node) float64 {
	return 1
}

func AStar(start, goal *data.Node) []*data.Node {
	openSet := []*AStarNode{}
	closedSet := map[*data.Node]bool{}

	startNode := &AStarNode{
		Node: start,
		G:    0,
		H:    heuristic(start, goal),
	}
	startNode.F = startNode.G + startNode.H
	openSet = append(openSet, startNode)

	for len(openSet) > 0 {

		// cari F terkecil
		currentIndex := 0
		for i := range openSet {
			if openSet[i].F < openSet[currentIndex].F {
				currentIndex = i
			}
		}

		current := openSet[currentIndex]

		// goal ketemu
		if current.Node == goal {
			return reconstructPath(current)
		}

		// pindah ke closed
		openSet = append(openSet[:currentIndex], openSet[currentIndex+1:]...)
		closedSet[current.Node] = true

		for neighbor, cost := range current.Node.Neighbors {
			if closedSet[neighbor] {
				continue
			}

			g := current.G + cost
			n := &AStarNode{
				Node:   neighbor,
				G:      g,
				H:      heuristic(neighbor, goal),
				Parent: current,
			}
			n.F = n.G + n.H

			openSet = append(openSet, n)
		}
	}

	return nil
}

func reconstructPath(node *AStarNode) []*data.Node {
	path := []*data.Node{}

	for node != nil {
		path = append([]*data.Node{node.Node}, path...)
		node = node.Parent
	}
	return path
}
