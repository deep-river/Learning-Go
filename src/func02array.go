package main

import "fmt" 

func main() {
	var list[5] int  
	fmt.Println("empty list", list)  
	list[4] = 100
	fmt.Println("get list[4]: ", list[4])
	fmt.Println("list: ", list)
	fmt.Println("len: ", len(list))

	b := [5] int {1, 2, 3, 4, 5}  
	fmt.Println("dcl: ", b) 
}