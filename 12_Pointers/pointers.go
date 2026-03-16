// Pointers in Go are variables that store the memory address of another variable. 
// They allow you to manipulate data directly in memory, which can be more efficient and can also enable certain programming techniques. 
// In Go, you can use the `*` operator to declare a pointer and the `&` operator to get the address of a variable.
// 
 package main;

import "fmt";

func changeNum(num *int) int{
	*num = 20; // changing the value at the memory address
	return *num; // returning the value at the memory address
}

func main(){
	num := 10;
	fmt.Println("Original value:", num);
	fmt.Println("Changed value:", changeNum(&num));
}