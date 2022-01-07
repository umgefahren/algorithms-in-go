package algorithms_in_go

import (
	"fmt"
	"github.com/yourbasic/graph"
	"math/rand"
	"testing"
)

func TestDijkstra(t *testing.T) {

	for i := 2; i < 5; i++ {
		t.Run(fmt.Sprintf("Graph-%v", i), func(t *testing.T) {
			for j := i; j < i*5; j++ {
				gr := graph.New(i)

				// fmt.Printf("i => %v j => %v \n", i, j)

				for g := 0; g < j; g++ {
					v := rand.Intn(i)
					w := rand.Intn(i)
					if v == w {
						for {
							w = rand.Intn(i)
							if w != v {
								break
							}
						}
					}
					c := rand.Int63n(10000)
					gr.AddBothCost(v, w, c)
					// fmt.Println(gr)
				}
				for w := 0; w < i; w++ {
					for v := 0; v < i; v++ {
						if v == w {
							continue
						}
						_, dist := graph.ShortestPath(gr, v, w)

						// fmt.Printf("The correct distance between %v and %v is %v \n", v, w, dist)

						if dist == -1 {
							continue
						}

						distTest := Dijkstra(gr, v, w)

						if dist != distTest && dist != -1 {
							t.Errorf("the graph is %v", gr)
							t.Errorf("The distance between %v and %v in the graph is %v not %v", v, w, dist, distTest)
						}
					}
				}
			}
		})
	}
}
