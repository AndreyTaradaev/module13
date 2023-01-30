package main

import (
	"fmt"
	"math/rand"
	"time"
	"tree/bst"
	
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	tree := bst.CreateTree(200)
	fmt.Println("tree: ", tree)

	fmt.Println("fill tree random value...")
	size := 20
	for i := 0; i < size; i++ {
		y := rand.Intn(size*10) - size*5
		switch {
		case y%3 == 0:
			tree.Add(y, y/2)
		case y%2 == 0:
			tree.Add(y, fmt.Sprintf("value_%d", y))
		default:
			tree.Add(y, y/10)
		}
	}
	fmt.Println("result fill tree ", tree)
	fmt.Println("Max value: ", tree.Max())
	fmt.Println("Min value: ", tree.Min())

	fmt.Println("1. Preorder short out")
	tree.Print(bst.Preorder, false)
	fmt.Println("2.  Inorder short out")
	tree.Print(bst.Inorder, false)
	fmt.Println("3. Postorder short out")
	tree.Print(bst.Postorder, false)
	fmt.Println("4.  BFS short out")
	tree.Print(bst.BFS, false)

	fmt.Println("\n")

	fmt.Print("Добавление ключа, введите значение ключа:  ")
	var w int
	fmt.Scanln(&w)
	tree.Add(w, float64(rand.Intn(size*10)-size*5)/3.5)
	tree.Print(bst.Inorder, false)
	fmt.Print("Поиск ключа ,введите значение ключа: ")
	fmt.Scanln(&w)
	if tree.IsExists(w) {
		fmt.Println("Есть такое:", tree.Find(w), "\n")
	} else {
		fmt.Println(" нет такого ключа")
	}
	fmt.Println(" \nДо удаления ")
	fmt.Println("tree: ", tree)
	tree.Print(bst.Inorder, false)
	fmt.Println("Удаление ключа , введите Значение ключа")
	fmt.Scan(&w)
	er := tree.Delete(w)
	if er != nil {
		fmt.Println(er)
		//return
	}
	fmt.Println(" После удаления ")
	fmt.Println("tree: ", tree)
	tree.Print(bst.Inorder, false)
	tree  = bst.CreateTree(200)
	fmt.Println("tree: ", tree)

	fmt.Println("fill tree random value  size 100000...")
	size = 100000
	for i := 0; i < size; i++ {
		y := rand.Intn(size*10) - size*5
		switch {
		case y%3 == 0:
			tree.Add(y, y/2)
		case y%2 == 0:
			tree.Add(y, fmt.Sprintf("value_%d", y))
		case y%5 == 0:
			tree.Add(y,  float64 (y)/7)
		default:
			tree.Add(y, y/10)
		}
	}
	fmt.Println("fill end")
	tree.Print(bst.Inorder, false)		
	fmt.Println("tree: ", tree)
	fmt.Print("Поиск ключа ,введите значение ключа: ")
	fmt.Scan(&w)
	if tree.IsExists(w) {
		fmt.Println("\nЕсть такое:", tree.Find(w), "\n")
	} else {
		fmt.Println(" \n нет такого ключа")
	}
}
