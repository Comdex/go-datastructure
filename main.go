// datastructure project main.go
package main

import (
	//"strings"
	"strconv"
	"os"
	"bufio"
	"fmt"
	"github.com/Comdex/datastructure/ds"
)

func main() {
	
	list := ds.NewLinkedList(4,false)
	list.Display()
	
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
	
	fmt.Println("please insert: index")
	scanner := bufio.NewScanner(os.Stdin)
	t := ""
      	if scanner.Scan() {
		t = scanner.Text()	
	}
	
	i,_ := strconv.Atoi(t)
	fmt.Println(list.IndexOf(i))
	list.Display()
	
	fmt.Println("is Empty: ",list.IsEmpty())
	fmt.Println("length :",list.Length())
	fmt.Println("clear :")
	list.Clear()
	fmt.Println("length :",list.Length())
}
