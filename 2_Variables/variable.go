package main

import "fmt";

func main(){
	var name string = "Golang";
	fmt.Println("Hello", name);

	//infer type : Go lang compiler infer the type of variable based on the value assigned to it.
	var age = 10;
	fmt.Println("Age:", age);

	//short hand declaration
	isContain := true;
	fmt.Println("Is contain?", isContain);
}