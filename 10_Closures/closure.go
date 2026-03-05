package main;

import "fmt";

func Counter() func() int{
	count := 0;
	return func() int{
		count+=1;
		return count;
	}
}

func main(){
	increment := Counter();
	fmt.Println(increment());
	 // Output: 1

}