package main;

import "fmt";

//Function declaration
func add(a,b int) int{
	return a+b;
}

//function with multiple return values
func swap(a,b string) (string, string){
	return b,a;
}



func main(){
	result := add(5,10);
	fmt.Println("Result of addition:", result);
	output1, output2 := swap("Hello", "World");
	fmt.Println("Output of swap:", output1, output2);
}