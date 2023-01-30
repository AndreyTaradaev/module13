package main 
import (
	"graph/graph"
	"fmt"
)

func main (){
//fmt.Print ("Введите количество вершин графа: ")
//var count int
//fmt.Scan(&count) 
g:= graph.CreateTestingGraf()
fmt.Println(g)

fmt.Println(`
 1----------5
 |\       /
 | \-----2---- \
 |              \
 0-----------3---4-----6`)


g.BFS(1,6)
g.BFS(5,6)
g.BFS(6,0)
g.BFS(5,0)
 fmt.Println( "Directed graph")
g = graph.CreateDirectGraf()
fmt.Println(`

1-----------*5
* \         *
|  \       /
|   \----*2*--- \
|                *
0*------------3--*4---*6
`)
g.BFS(1,6)

g.BFS(0,6)
 g.BFS(6,0)
 fmt.Println( "Directed graph #2")
g = graph.CreateDirectGraf1()
fmt.Println(`


----------*3 -----*5 *- 
|   1 ---* *       *  |
|  * \-----+--*4---|  |
| /        |          |
|/         |          |
0---------*2---------*6
`)
g.BFS(0,5)
g.BFS(2,5)
g.BFS(3,0)
g.BFS(1,2)
g.BFS(5,0)
}