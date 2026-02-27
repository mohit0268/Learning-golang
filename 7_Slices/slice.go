package main

import (
	"fmt"
	"slices"
)

//slice -> dynamically sized array
//most used construct in Go lang
//useful methods for slice
func main(){
	//uninitialized slice is nil
	//syntax for slice declaration
	
	var s[]int;
	fmt.Println("Slice s:", s);
	fmt.Println(len(s))

	//make([]<arrayitem>, length, capacity)
	s = make([]int, 0, 5);
	fmt.Println("Slice s:", s);
	fmt.Println("Length of slice s:", len(s));
	fmt.Println("Capacity of slice s:", cap(s));

	//append method to add elements to slice
	s = append(s,1)
	s = append(s,2)
	s = append(s,3)
	s = append(s,4)
	s = append(s,5)
	s = append(s,6) //appending 6 will increase the capacity of slice s
	fmt.Println("Slice s after appending 1:", s);
	fmt.Println("Length of slice s after appending 1:", len(s));
	fmt.Println("Capacity of slice s after appending 1:", cap(s));

	//copy method to copy elements from one slice to another

	nums := make([]int, 0, 5);
	nums = append(nums,1,2);
	nums2 := make([]int, len(nums));
	copy(nums2, nums);
	fmt.Println("Slice nums:", nums);
	fmt.Println("Slice nums2:", nums2);

	//slice operations
	var a = []int{1,2,3,4,5};
	slice1 := a[1:4]; //slice from index 1 to 3
	fmt.Println("Slice slice1:", slice1);

	var b = []int{1,2,3,4,5};
	var compare = slices.Equal(a,b);
	fmt.Println("Are slices a and b equal?", compare);

}