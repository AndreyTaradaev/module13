package graph

import "fmt"

type Graph struct {
	list [][]int
}

func CreateGraph(node int) *Graph {
	elem := Graph{}
	elem.list = make([][]int, node)
	return &elem
}

func  LoadGraf ( l [][]int)  *Graph{
g := CreateGraph(len(l))
g.list = l
return g
}





/* func FillGraph() (*Graph,error){
fmt.Print("Введите количество вершин:")
var v uint 
_,er:=fmt.Scan(&v)
if(er !=nil){
	return nil ,er
}
fmt.Println( " Заполнение связей графа( введите список соседних вершин без пробелов )")
g := CreateGraph(int(v))
for i:=uint(0) ; i < v ;i++{
	fmt.Println( "Вершина", i,":" )
	var w uint
	_,err := fmt.Scan(&w)	
	if(err!=nil){
		return nil, err
	}	
	for i:=w; i>0; i/=10{
		a+= i%10    
		}



}
 */
//}


func (g Graph) String() string {
	var ret string
	for i := 0; i < len(g.list); i++ {
		ret = ret +"вершина: " + fmt.Sprintf("%d",i)+"\n\t" 
		ar:=  g.list[i]	
			for y:=0; y< len(ar);y++{
			ret = ret + fmt.Sprintf("%d ",ar[y])	
			}
			ret = ret +"\n"
	}
	return ret
}

//   1---------5
//	 |\	      /
//	 | \-----2---- \
//	 |				\
//   0-----------3---4-----6


func CreateTestingGraf()  *Graph{ 			
	g := make([][]int ,7)
	//0
	ar := [] int{1,3}
	g[0] = ar
	//1
	ar = [] int{0,2,5}
	g[1] = ar

	//2
	ar = [] int{1,4,5}
	g[2] = ar

	//3
	ar = [] int{0,4}
	g[3] = ar
	//4
	ar = [] int{2,3,6}
	g[4] = ar
	//5
	ar = [] int{1,2}
	g[5] = ar
	//6
	ar = [] int{4}
	g[6] = ar
	return LoadGraf(g)
}

func (g *Graph) BFS(startnode ,endnode int) error{
if( startnode<0 || startnode>=len(g.list) ){
	return fmt.Errorf("start node out range")
}

if( endnode<0 || endnode>=len(g.list) ){
	return fmt.Errorf("end node out range")
}
// очередь
var q  Queue //   очередь вершин
q.Enqueue(startnode)
used := make( [] bool, len (g.list),len (g.list)) // признак  посещености вершины
d :=  make( [] int , len (g.list),len (g.list))  // расстояние
p :=  make( [] int , len (g.list),len (g.list))  // предки

used [startnode] = true
p[startnode] =-1
for q.Len()!=0{
	v,_ := q.Peek()//Извлекаем из головы очереди вершину
	q.Dequeue()
	size := len(g.list[v]) //получаем количество связей у вершины
	for i:=0; i<size ;i++{
		to := g.list[v][i]
			if(!used[to]){//Если вершина не посещена
				used[to] = true //посещаем ее
				q.Enqueue(to) //и добавляем к концу очереди
				d[to] = d[v]+1 //Считаем расстояние до вершины
				p[to] = v //Запоминаем предка
			}
	}
}

if(!used[endnode]){
fmt.Println("Path from  ",startnode ," to ",endnode, "not found" )	
return fmt.Errorf("not found path")	
} else{
 path := []int{}
 for v:=endnode ; v!=-1; v=p[v]{
	path = append(path,v)
 }
 fmt.Println("Path from  ",startnode ," to ",endnode)

 for i:= len(path)-1;i>=0; i--{
fmt.Print( path[i]," " )
 }
 fmt.Print( "\n" )

}
return nil
}

type Queue struct {
	values [] int
}

// добавление в очередь
func (q *Queue) Enqueue(element int) {

	q.values = append(q.values, element) // Simply append to enqueue.
}

// значение пераого в очереди
func (q *Queue) Peek() (int, error) {

	if len(q.values) == 0 {
		return 0, fmt.Errorf("Queue is emply!")
	}
	return q.values[0], nil
}

// удаление из очереди
func (q *Queue) Dequeue() (int, error) {

	if len(q.values) == 0 {
		return 0, fmt.Errorf("Queue is emply!")
	}
	element := q.values[0]
	if len(q.values) == 1 {
		q.values = []int{}
		return element, nil
	}
	q.values = q.values[1:]
	return element, nil
}

func (q Queue) Len() int {
	return len(q.values)
}

func (q Queue) Get() []int {
	return q.values
}

func (q *Queue) String() string {
	var ret string
	ar := q.values
	for _, val := range ar {
		ret = ret + fmt.Sprintf("%d ", val)
	}
	return ret
}

/*

   1-----------*5
   * \         *
   |  \       /
   |   \----*2*--- \
   |				*
   0*------------3--*4---*6

*/

func CreateDirectGraf()  *Graph{ 			
	g := make([][]int ,7)
	//0
	ar := [] int{1}
	g[0] = ar
	//1
	ar = [] int{2,5}
	g[1] = ar

	//2
	ar = [] int{4,5}
	g[2] = ar

	//3
	ar = [] int{0,4}
	g[3] = ar
	//4
	ar = [] int{2,6}
	g[4] = ar
	//5
	ar = [] int{}
	g[5] = ar
	//6
	ar = [] int{}
	g[6] = ar
	return LoadGraf(g)
}


/*

   ----------*3 -----*5 *- 
   |   1 ---* *       *  |
   |  * \-----+--*4---|  |
   | /        |          |
   |/         |          |
   0---------*2---------*6

*/

func CreateDirectGraf1()  *Graph{ 			
	g := make([][]int ,7)
	//0
	ar := [] int{1,2,3}
	g[0] = ar
	//1
	ar = [] int{3,4}
	g[1] = ar

	//2
	ar = [] int{3,6}
	g[2] = ar

	//3
	ar = [] int{5}
	g[3] = ar
	//4
	ar = [] int{5}
	g[4] = ar
	//5
	ar = [] int{}
	g[5] = ar	
	//6
	ar = [] int{5}
	g[6] = ar	
	return LoadGraf(g)
}