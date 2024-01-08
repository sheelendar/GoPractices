package main

import "fmt"

/*
"Given Shopping cart with products and coupons and calculate the net price after applying coupons on products.
Coupons can be of different types with certain conditions.
1. N% off that is 10% off for all the individual
2.0% off on next item
3.D% off on Nth item of Type T.
Sequentially wants to apply all the coupons on the cart and get the Total amount."

*/

func main() {
	item1 := Item1Constractor("FAN", float64(1000), Electronic)
	item2 := Item1Constractor("SOFA", float64(2000), Funiture)

	shoppingCart := ShoppingCartConstractor()
	shoppingCart.addProductInList(item1, item1.productType)
	shoppingCart.addProductInList(item2, item2.productType)

	fmt.Println(shoppingCart.getTotalPrice())

}
