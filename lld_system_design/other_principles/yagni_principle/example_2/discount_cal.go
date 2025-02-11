package main

import "fmt"

// Extra function for a feature that is not yet needed
func calculateDiscount(price float64, discount float64, isHoliday bool) float64 {
	if isHoliday {
		return price * (1 - discount/100) // Additional logic for holidays
	}
	return price * (1 - discount/100)
}

// Not needed yet, but implemented anyway
func applyLoyaltyPoints(userID int, points int) {
	fmt.Println("Loyalty points system coming soon...") // Placeholder
}

func main() {
	price := 100.0
	fmt.Println("Discounted Price:", calculateDiscount(price, 10, false))
}
