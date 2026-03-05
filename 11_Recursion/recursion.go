package main;

import "fmt";

//returning a recursive function to calculate factorial of a number which return int and takes an int as argument
func fact(n int) int{
	if(n == 0){
		return 1;
	}
	return n * fact(n-1);
}

func main(){
	fmt.Println("Factorial of 5 is:", fact(5));
	var fib func(n int) int;

	fib = func(n int) int {
		if n <= 1{
			return n;
		}
		return fib(n-1) + fib(n-2);
	}

	fmt.Println("Fibonacci of 10 is:", fib(10));
}