package node

import (
	"fmt"
)



type Node struct{
Key  int
Value  interface{}
count uint // dublicates
left  *Node
right *Node
}

func (n Node) String() string {
return  fmt.Sprintf("key: %d , Value: %#v(%T), count %d, left:  %p, rigth %p", n.Key,n.Value,n.Value,n.count,n.left,n.right)
} 

func (n Node) Print () string{
	var ret string 
	for i:=0 ;i< int(n.count);i++{
		ret = ret + fmt.Sprintf("%d(%#v),",n.Key,n.Value)
	}
	return  ret 
}
// getters 
func (n *Node) Right ()  *Node {
	if (n == nil){
		return nil
	}
	return n.right
}

func (n *Node) Left()  *Node {
	if (n == nil){
		return nil
	}
	return n.left
}



func deleteChild(parent ,item,child *Node){
	if(parent==nil) {// root
		if(item.right == nil) {//у корня нет правой ветки
			l:= item.left
			item.Key = l.Key
			item.Value = l.Value
			item.left = l.left
			item.right = l.right
			l.left = nil
			l.right = nil
			return
		}
	item.Key = child.Key
	item.Value = child.Value
	item.right = child.right
	return
	}
	if(parent.right == item) {
		parent.right = child
		if(item.left !=nil ){
			child.left = item.left
		}
		item.left = nil
		item.right = nil
	}else {
		parent.left = child
		item.left = nil
		item.right = nil
	}	
}

func deleteTwoChild(parent, node *Node){
//надо найти справа минимум
minnode := node.right.Min()
//если min равен  правому потомку 
	if(minnode == node.right){
		deleteChild(parent,node,node.right)	
		return				
	}
	node.Key = minnode.Key
	node.Value = minnode.Value
	r:=  minnode.right
	if(r!=nil){
		minnode.Key =  r.Key
		minnode.Value = r.Value				
		minnode.right = r.right
		minnode.left = 	r.left			
		r.left = nil
		r.right = nil
	}
}


func (n *Node) Delete (parent  *Node )  {
	if(n.count>1){ //если есть дубликаты
		n.count--
		return
	}
	
	if(parent == nil){ //надо удалить корень
		deleteTwoChild(parent, n)
		return
	}
	if(n.left == nil){ // если левого потомка нет
		deleteChild(parent,n,n.right)
	} else	if(n.right == nil){ // если  правого потомка нет
		deleteChild(parent,n,n.left)
	}else{
		deleteTwoChild(parent, n)
	}
	return 
}

func (n *Node) replace(parent, target *Node ) {

}

//ctor

func New (newkey int, newvalue interface{}) *Node {
	elem :=  Node{Key:newkey,Value: newvalue, count: 1, left: nil,right: nil}
	return &elem
}

func (n *Node) Add(unit *Node)   {
	current := n	
	for  current !=nil {
		if(current.Key == unit.Key){
			current.count++	
			return 
		 } 
		 if current.Key>unit.Key {
			if(current.left == nil ){
				current.left = unit
				break
			}
			current  = current.left	
		 }else{
			if(current.right == nil ){
				current.right = unit
				break
			}
			current = current.right
		 }	 		
	}		
}

func (n *Node) Max()   *Node {
	if n==nil {
		return nil
	}
	current :=  n
	for current.right !=nil{
		current = current.right	
	}
	return current
}


func (n *Node) Min( )   *Node {
	if n==nil {
		return nil
	}
	current :=  n
	for current.left !=nil{
		current = current.left
	}
return current
}
