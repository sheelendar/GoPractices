package main

// ************************************************** ShoppingCart*****************************************
type ShoppingCart struct {
	productList []ProductPrice
}

func ShoppingCartConstractor() *ShoppingCart {
	return &ShoppingCart{}
}

func (shoppingcart *ShoppingCart) GetProductList() []ProductPrice {
	return shoppingcart.productList
}

func (shoppingcart *ShoppingCart) addProductInList(product ProductPrice, productType ProductType) {
	tcc := TypeCouponDecoratorConstractor(PercentageCouponDecoratorConstractor(product, 10), 3, productType)

	shoppingcart.productList = append(shoppingcart.productList, tcc)
}

func (shoppingcart *ShoppingCart) getTotalPrice() float64 {
	totalPrice := float64(0)
	for _, product := range shoppingcart.productList {
		totalPrice += product.getPrice()
	}
	return totalPrice
}
