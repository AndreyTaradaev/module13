package bst

/*
попытка  реализации двоичного дерева без рекурсии , чтобы  исключить переполнение стека

*/
import (
	"fmt"
	"tree/node"
	"tree/queue"
)

type TypeTraverse byte

const ( // типы  обхода
	Preorder TypeTraverse = iota + 1 //Симметричный обход
	Inorder
	Postorder
	BFS
)

type bst struct {
	root  *node.Node
	max   uint
	count uint
}

func CreateTree(countchange uint) *bst {
	elem := bst{root: nil, max: countchange, count: countchange}
	return &elem
}

func (b bst) String() string {
if(b.root !=nil){
	return fmt.Sprintf(" root %p, limit change %d, remained %d, root %s ", b.root, b.max, b.count, b.root.Print())
}
return fmt.Sprintf(" root %p, limit change %d, remained %d", b.root, b.max, b.count )
}

func (b *bst) Add(key int, newvalue interface{}) {

	unit := node.New(key, newvalue)
	if b.root == nil {
		b.root = unit
		return
	}
	b.root.Add(unit)
	b.count--
	b.rebalance()
}

func (b *bst) rebalance(){
	if (b.count==0){
		b.count = b.max // need rebalance
	}
}

func (b *bst) Find(key int) *node.Node {
 item ,_:=find (b.root,key)
 return item
}


func  find(root *node.Node, key int)  (elem ,parent *node.Node ){ // поиск узла и родителя
	if root == nil {
		return nil,nil
	}
	elem = root	
	for elem.Key != key {
		if elem.Key > key {
			parent = elem
			elem = elem.Left()
		} else {
			parent = elem
			elem = elem.Right()
		}
		if elem == nil {
			return elem,nil
		}
	}
	return elem	,parent
}


func (b *bst) IsExists(key int) bool {
	return (b.Find(key) != nil)
}

func (b *bst) Min() *node.Node {
	if b == nil {
		return nil
	}
	return b.root.Min()
}

func (b *bst) Max() *node.Node {
	if b == nil {
		return nil
	}
	return b.root.Max()
}

func (b *bst) Print(Order TypeTraverse, detail bool) {
	var f func(*node.Node) (queue.Array, error)

	switch Order {
	case Preorder:
		//fallthroug
		f = CreatePreorderArray
	case Inorder:
		f = CreateInorderArray
	case Postorder:
		f = CreatePostorderArray
	default:
		f = CreateBFSArray
	}
	ar, er := f(b.root)
	if er != nil {
		fmt.Println("Error")
	}
	ar.SetDetail(detail)
	fmt.Println(ar)
}

func (b *bst) Delete(delme int) error{
	if (b.root == nil){
	return	fmt.Errorf("tree is emply")
	}
	delNode,parent  :=  find(b.root,delme)
	if(delNode==nil){
		return	fmt.Errorf("key not found")
	}

	fmt.Println("found ", delNode, " parent ",parent)
	delNode.Delete(parent)
	b.count-- 
	if(b.count==0){
		b.count = b.max
		b.rebalance()
	}
	return nil
}

func CreatePreorderArray(root *node.Node) (queue.Array, error) {
	st := new(queue.Stack)
	arr := queue.NewArray()
	if root == nil {
		return *arr, fmt.Errorf("root is null")
	}
	st.Push(root)
	for st.Len() != 0 {
		item, er := st.Pop()
		if er != nil {
			return *arr, fmt.Errorf("Pop error ")
		}
		arr.Add(item)
		if item.Right() != nil {
			st.Push(item.Right())
		}
		if item.Left() != nil {
			st.Push(item.Left())
		}
	}
	return *arr, nil
}

func CreateInorderArray(root *node.Node) (queue.Array, error) {
	st := new(queue.Stack)
	arr := queue.NewArray()
	if root == nil {
		return *arr, fmt.Errorf("root is null")
	}
	item := root
	for st.Len() != 0 || item != nil {
		if item != nil {
			st.Push(item)
			item = item.Left()
		} else {
			n, _ := st.Pop()
			arr.Add(n)
			item = n.Right()
		}
	}
	return *arr, nil
}

func CreatePostorderArray(root *node.Node) (queue.Array, error) {
	st := new(queue.Stack)
	arr := queue.NewArray()
	//var lastVisit *node.Node = nil
	if root == nil {
		return *arr, fmt.Errorf("root is null")
	}
	prev := root
	st.Push(root)
	for st.Len() != 0{
		item,_ := st.Peek()
		if(item.Left()!=nil && item.Left() !=prev && item.Right() !=prev ){
			st.Push(item.Left())
		} else if (item.Right()!=nil && item.Right()!=prev ){
			st.Push(item.Right())
		} else{
			prev = item
			st.Pop()
			arr.Add(item)
		}
	}
/* 	for st.Len() != 0 || item != nil {
		if item != nil {
			st.Push(item)
			item = item.Left()
		} else {
			peekNode, _ := st.Peek()			 
			if peekNode.Right() != nil && *lastVisit != *peekNode.Right() {
				item = peekNode.Right()
			} else {
				arr.Add(peekNode)
				st.Pop()
				lastVisit = peekNode
			}
		}
	} */
	return *arr, nil
}

func CreateBFSArray(root *node.Node) (queue.Array, error) {
	q := new(queue.Queue)
	arr := queue.NewArray()
	if root == nil {
		return *arr, fmt.Errorf("root is null")
	}
	q.Enqueue(root)
	for q.Len()!=0{
		if elem,err :=  q.Peek(); err !=nil{
			return	*arr, fmt.Errorf("Error Peek")	
			} else{
			if(elem.Left()!=nil) {
				q.Enqueue(elem.Left())
			}
			if(elem.Right()!=nil) {
				q.Enqueue(elem.Right())
			}
			q,_:= q.Peek()
			arr.Add(q)
			}
			q.Dequeue()
	}

	return *arr, nil
}




