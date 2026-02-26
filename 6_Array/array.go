package main;

import "fmt";

func main(){
	// //array declaration 
	// var a [5]int;
	// //initialization of array
	// a[1] = 11;
	// 	fmt.Println("Array a:", a);

	//shorhand declaration and initialization of array
	a:= [5]int {1,2,3,4,5};
	fmt.Println("Array a:", a);

	//array with ellipsis
	b := [...]string {"Go", "is", "awesome", "language"};
	fmt.Println("Array b:", b);

}