package main;

import (
	"fmt"
)

func main(){
	//map declaration and initialization
	//if key does't exist in map, it will return zero value of value type
	//syntax : make(map[keytype]valuetype)
	var users = make(map[string]int);
	//adding key-value pairs to map
	users["Alice"] = 25;
	users["Bob"] = 30;
	fmt.Println("Users map:", users);
	fmt.Println(users["Alice"]);

	//deleteing key-value pair from map
	delete(users, "Alice");
	fmt.Println("Users map after deleting Alice:", users);

	//clearing map
	clear(users);
	fmt.Println("Users map after clearing:", users);

}