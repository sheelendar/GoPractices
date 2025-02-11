package example1

type WaiterInterFace interface {
	SarveCustomer()
	TakeOrder()
}

type ChefInterface interface {
	DecideMenu()
	CookFood()
}

type HelperInterface interface {
	WashDishes()
	SweepFloor()
}
