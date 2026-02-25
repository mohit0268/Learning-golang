package main

import "fmt";

func main(){
	//there is nothing like while loop in Go lang, we can use for loop to achieve the same functionality as while loop.
	//forLoop
	// var i int = 0;
	// for i<5 {
	// 	fmt.Println("Value of i:", i);
	// 	i++;
	// }

	//for loop with initialization and post statement
	// for i := 1; i < 5; i++ {
	// 	fmt.Println("Value of i:", i);
	//}

	//for loop with range
	for i := range 10{
		fmt.Println("Value of i:", i);
	}

	
}