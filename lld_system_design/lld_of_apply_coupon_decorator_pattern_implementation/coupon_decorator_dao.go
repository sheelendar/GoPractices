package main

// ************************************************** CouponDecorator*****************************************

// ************************************************** PercentageCouponDecorator*****************************************
type PercentageCouponDecorator struct {
	product            ProductPrice
	discountPercentage int
}

func PercentageCouponDecoratorConstractor(product ProductPrice, discountPercentage int) *PercentageCouponDecorator {
	return &PercentageCouponDecorator{product: product, discountPercentage: discountPercentage}
}

func (pcd *PercentageCouponDecorator) getPrice() float64 {
	price := pcd.product.getPrice()
	return price - (float64(pcd.discountPercentage)*price)/100
}

// ************************************************** TypeCouponDecorator*****************************************
type TypeCouponDecorator struct {
	product            ProductPrice
	discountPercentage int
	productType        ProductType
}

var eligibleTypes = map[ProductType]bool{Decorative: true, Funiture: true}

func TypeCouponDecoratorConstractor(product ProductPrice, discountPercentage int, productType ProductType) *TypeCouponDecorator {
	return &TypeCouponDecorator{product: product, discountPercentage: discountPercentage, productType: productType}
}

func (tcd *TypeCouponDecorator) getPrice() float64 {
	price := tcd.product.getPrice()
	if eligibleTypes[tcd.productType] {
		return price - (price * float64(tcd.discountPercentage) / 100)
	}
	return price
}
