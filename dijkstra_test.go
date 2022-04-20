package dijkstra

import (
	"fmt"
	"log"
	"testing"
)

func TestSalesman(t *testing.T) {
	graph := NewGraph()
	graph.AddVertex(0)
	graph.AddVertex(1)
	graph.AddVertex(2)
	graph.AddVertex(3)
	graph.AddVertex(4)
	graph.AddVertex(5)
	graph.AddVertex(6)
	graph.AddVertex(7)
	graph.AddVertex(8)
	graph.AddVertex(9)
	graph.AddVertex(10)
	graph.AddVertex(11)
	graph.AddVertex(12)
	graph.AddVertex(13)
	graph.AddVertex(14)
	graph.AddVertex(15)
	graph.AddVertex(16)
	graph.AddVertex(17)
	graph.AddVertex(18)
	graph.AddVertex(19)

	graph.AddArc(0, 1, 71)
	graph.AddArc(0, 7, 151)
	graph.AddArc(1, 0, 71)
	graph.AddArc(1, 2, 75)
	graph.AddArc(2, 1, 75)
	graph.AddArc(2, 3, 118)
	graph.AddArc(2, 7, 140)
	graph.AddArc(3, 2, 118)
	graph.AddArc(3, 4, 111)
	graph.AddArc(4, 3, 111)
	graph.AddArc(4, 5, 70)
	graph.AddArc(5, 4, 70)
	graph.AddArc(5, 6, 75)
	graph.AddArc(6, 6, 75)
	graph.AddArc(6, 7, 120)
	graph.AddArc(7, 0, 151)
	graph.AddArc(7, 2, 140)
	graph.AddArc(7, 8, 80)
	graph.AddArc(7, 10, 99)
	graph.AddArc(8, 7, 80)
	graph.AddArc(8, 11, 97)
	graph.AddArc(8, 9, 146)
	graph.AddArc(9, 6, 120)
	graph.AddArc(9, 8, 146)
	graph.AddArc(9, 11, 138)
	graph.AddArc(10, 8, 99)
	graph.AddArc(10, 13, 211)
	graph.AddArc(11, 8, 97)
	graph.AddArc(11, 9, 138)
	graph.AddArc(11, 13, 101)
	graph.AddArc(12, 13, 90)
	graph.AddArc(13, 10, 211)
	graph.AddArc(13, 11, 101)
	graph.AddArc(13, 12, 90)
	graph.AddArc(14, 16, 87)
	graph.AddArc(15, 17, 142)
	graph.AddArc(15, 18, 98)
	graph.AddArc(16, 14, 87)
	graph.AddArc(16, 17, 92)
	graph.AddArc(17, 16, 92)
	graph.AddArc(17, 14, 142)
	graph.AddArc(18, 14, 98)
	graph.AddArc(18, 19, 86)
	graph.AddArc(19, 18, 86)
	findPath(*graph)
}
func findPath(graph Graph) {
	best, err := graph.Shortest(3, 13)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Shortest distance ", best.Distance, " following path ", best.Path)

	best, err = graph.Longest(3, 13)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Longest distance ", best.Distance, " following path ", best.Path)

}
