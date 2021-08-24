package main

import "fmt" 

func main() {
	nums := []int{2,3,4}
	sum := 0
	// calculate the sums 
	for _,num := range nums {
		sum += num 
	}
	fmt.Println("sum: ", sum) 
	// look up for the index of value 3
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}
	kvs := map[string]string{"a":"apple","b":"banana"} 
	// taraversal of map kvs
	for k, v := range kvs {
		fmt.Printf("%v -> %v\n", k, v) 
	}
	for k := range kvs {
		fmt.Println("key of map kvs: ", k)
	}
	// print each character[ASCII code] of a string 
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}