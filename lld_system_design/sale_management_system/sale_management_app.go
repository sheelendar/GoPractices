package main

import "fmt"

/*
Implement a sales management system that allows one to track the sales at various levels in an organization.
Here, we have a company that sells snacks - food items and beverages at retail outlets. These outlets are all over India.
We want to have a system that tracks sales details at following levels
1) Store
2) City
3) State
4) All India

As an example, stores sell following items
1) Food items
a) Sandwich
b) Poha
c) Vada Paav
d) Burger

2) Beverages
a) Tea
b) Coffee
c) Bottled water

This system mainly tracks following:
1) Price of items on the menu (Prices will vary across cities. Within the city the prices will remain the same)
2) Stock left at a particular store
3) Number of units sold at a given store
Please do not focus on persistence layer i.e. data store. Use in-memory store for now. Use file system if you prefer that. Do not use any database or NoSQL store.
It is absolutely OK if certain values (e.g. Prices of items, Stock at a particular store) that you need for calculation are hardcoded.
Please make sure that you have a working code at the end of the session.
*/

func main() {
	foodPrices := make(map[FoodItem]int)
	beveragePrices := make(map[BeverageItem]int)
	foodPrices[Burger] = 10
	foodPrices[Poha] = 15
	foodPrices[Vada] = 18
	foodPrices[Sandwich] = 20

	beveragePrices[Coffee] = 25
	beveragePrices[Tea] = 21
	beveragePrices[Water] = 15

	foodSupply := make(map[FoodItem]int)
	beverageSupply := make(map[BeverageItem]int)
	foodSupply[Burger] = 10
	foodSupply[Poha] = 10
	foodSupply[Vada] = 10
	foodSupply[Sandwich] = 10

	beverageSupply[Coffee] = 10
	beverageSupply[Tea] = 10
	beverageSupply[Water] = 10

	store := GetStoreConstractor(foodSupply, foodPrices, beverageSupply, beveragePrices)
	city := GetCityConstractor(foodPrices, beveragePrices)
	city.addStore(store)

	state := GetStateConstractor()
	state.addCity(city)

	system := GetSystemConstractor(1)
	system.addState(state)

	system.purchageFood(state.id, city.id, store.id, Burger, 7)
	DisplayProductItemUnit(system)

	system.purchageFood(state.id, city.id, store.id, Burger, 5)

	system.purchageBeverage(state.id, city.id, store.id, Coffee, 5)
	system.purchageFood(state.id, city.id, store.id, Poha, 4)
	DisplayProductItemUnit(system)
}

func DisplayProductItemUnit(system *System) {
	for _, state := range system.states {

		for city, _ := range state.cities {

			for _, store := range city.stores {

				for product, qty := range store.foodSupply {
					fmt.Println(product, "  ", qty)
				}
				fmt.Println()
				for product, qty := range store.beverageSupply {
					fmt.Println(product, "  ", qty)
				}
			}
		}
	}
}
