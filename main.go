// datastructure project main.go
package main

import (
	//"strings"
	//"strconv"
	//"os"
	//"bufio"
	"fmt"
	"github.com/Comdex/datastructure/ds"
)

func main() {
	
//	list := ds.NewLinkedList(4,false)
//	list.Display()
	
//	fmt.Println("please insert: index,data")
//	scanner := bufio.NewScanner(os.Stdin)
//	t := ""
//      	if scanner.Scan() {
//		t = scanner.Text()	
//	}
//	s := strings.Split(t,",")
//	i,_ := strconv.Atoi(s[0])
//	j,_ := strconv.Atoi(s[1])
//	list.Insert(i,j)
//	list.Display()
	
//	fmt.Println("please insert: index")
//	scanner := bufio.NewScanner(os.Stdin)
//	t := ""
//      	if scanner.Scan() {
//		t = scanner.Text()	
//	}
	
//	i,_ := strconv.Atoi(t)
//	fmt.Println(list.IndexOf(i))
//	list.Display()
	
//	fmt.Println("is Empty: ",list.IsEmpty())
//	fmt.Println("length :",list.Length())
//	fmt.Println("clear :")
//	list.Clear()
//	fmt.Println("length :",list.Length())
	list := ds.NewLinkedList2()
	list.Insert(0,55)
	list.Insert(1,66)
	list.Insert(2,77)
	list.Display()
	fmt.Println("翻转后:")
	list.Reverse()
	list.Display()
	
	//线性表顺序表
	list2 := ds.NewSqList()
	list2.Insert(0,1)
	list2.Insert(1,2)
	list2.Insert(2,3)
	list2.Insert(3,4)
	list2.Display()
	d,_ := list2.Get(2)
	fmt.Println("Get :",d)
	fmt.Println("IndexOf :",list2.IndexOf(4))
	list.Remove(2)
	fmt.Println("Remove : ")
	list2.Display()
	fmt.Println(list2.Length())
}
