package main;

import "fmt" ;
import "time" ;


type order struct {
	id int
	amount float64
	status string
	createdAt time.Time
}
//receiver function to update the status of the order we are using a pointer receiver so that we can modify the original order struct
func (o *order) updateStatus(newStatus string) {
	o.status = newStatus
}

//while getting the value of the order struct we don't need to use a pointer receiver because we are not modifying the original struct, we are just reading its value

func main(){
	//default value that it will have if we don't initialize it
	//int = 0 float = 0 string = "" bool =false time = 0001-01-01 00:00:00 +0000 UTC
	myOrder:= order{
		id: 1,
		amount: 99.99,
		status: "pending",
	}
	

	myOrder.updateStatus("confirmed");
	fmt.Printf("Order ID: %d, Amount: %.2f, Status: %s, Created At: %s\n",
		myOrder.id, myOrder.amount, myOrder.status, myOrder.createdAt.Format(time.RFC1123))
}