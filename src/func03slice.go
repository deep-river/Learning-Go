package main

import "fmt" 

func main() {
	s := make([]string, 3)
	fmt.Println("empty clice s: ", s)  
	s[0] = "a" 
	s[1] = "b" 
	s[2] = "c" 
	fmt.Println("set value of s: ", s)
	fmt.Println("get value of s[2]: ", s[2])
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("append s: ", s)
	c := make([]string, len(s)) 
	copy(c, s)  
	fmt.Println("copy s to c: ", c)
	l := s[2:5] 
	fmt.Println("slice1 of s[2:5]: ", l)
	l = s[:5] 
	fmt.Println("slice2 of s[:5]: ", l)
	l = s[2:] 
	fmt.Println("slice3 of s[2:]: ", l)
}