package main

import (
	"fmt"
	controller "gameton/Controller"
	"gameton/Graph"
)

func main() {
	Graph.FillGraph(controller.PlayerUniversal().Universe)
	fmt.Println(Graph.Dijkstra("Earth", "Kilback"))
	fmt.Println(controller.PlayerUniversal().Ship.Planet)
	fmt.Print(Graph.Adjacency["PinkDrum"])
}
