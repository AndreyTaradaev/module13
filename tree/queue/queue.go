package queue

import (
	"fmt"
	"tree/node"
)

// очередь  FIFO
type Queue struct {
	values []*node.Node
}

// добавление в очередь
func (q *Queue) Enqueue(element *node.Node) {

	q.values = append(q.values, element) // Simply append to enqueue.
}

// значение пераого в очереди
func (q *Queue) Peek() (*node.Node, error) {

	if len(q.values) == 0 {
		return &node.Node{}, fmt.Errorf("Queue is emply!")
	}
	return q.values[0], nil
}

// удаление из очереди
func (q *Queue) Dequeue() (*node.Node, error) {

	if len(q.values) == 0 {
		return &node.Node{}, fmt.Errorf("Queue is emply!")
	}
	element := q.values[0]
	if len(q.values) == 1 {
		q.values = []*node.Node{}
		return element, nil
	}
	q.values = q.values[1:]
	return element, nil
}

func (q Queue) Len() int {
	return len(q.values)
}

func (q Queue) Get() []*node.Node {
	return q.values
}

func (q *Queue) String() string {
	var ret string
	ar := q.values
	for _, val := range ar {
		ret = ret + fmt.Sprintf("%s ", val.String())
	}
	return ret
}

func (q Queue) Print() string {
	var ret string
	for _, val := range q.values {
		ret = ret + fmt.Sprintf("%s ", val.Print())
	}
	return ret
}

// очередь LIFO
type Stack struct {
	values []*node.Node
}

// добавление в очередь
func (q *Stack) Push(element *node.Node) {

	q.values = append(q.values, element) // Simply append to enqueue.
}

// значение пераого в очереди
func (q *Stack) Peek() (*node.Node, error) {
	s := len(q.values)
	if s == 0 {
		return &node.Node{}, fmt.Errorf("Stack is emply!")
	}
	return q.values[s-1], nil
}

// удаление из очереди
func (q *Stack) Pop() (*node.Node, error) {
	s := len(q.values)
	if s == 0 {
		return &node.Node{}, fmt.Errorf("Queue is emply!")
	}
	element := q.values[s-1]
	if len(q.values) == 1 {
		q.values = []*node.Node{}
		return element, nil
	}
	q.values = q.values[:s-1]
	return element, nil
}

func (q Stack) Len() int {
	return len(q.values)
}

func (q *Stack) Get() []*node.Node {
	return q.values
}

func (q *Stack) String() string {
	var ret string
	ar := q.values
	for _, val := range ar {
		ret = ret + fmt.Sprintf("%#v ", val)
	}
	return ret
}

func (q *Stack) Print() string {
	var ret string
	for _, val := range q.values {
		ret = ret + fmt.Sprintf("%s ", val.Print())
	}
	return ret 
}

///	==================================================

type Array struct {
	values []*node.Node
	detail bool
}

func NewArray() *Array {
	a := new(Array)
	a.values = make([]*node.Node, 0, 1)
	a.detail = false
	return a
}

func (a *Array) SetDetail(d bool) {
	a.detail = d
}

func (a *Array) Add(elem *node.Node) {
	a.values = append(a.values, elem)
}

func (a Array) String() string {
	var ret string
	for _, el := range a.values {
		if a.detail {
			ret =ret +" "+ el.String()
		}else{
			ret = ret + " " + el.Print()
		}
	}
	return ret
}