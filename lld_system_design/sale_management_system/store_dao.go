package main

import (
	"fmt"
)

type FoodItem int
type BeverageItem int

const (
	Sandwich = iota + 1
	Poha
	Vada
	Burger
)
const (
	Tea = iota + 1
	Coffee
	Water
)

var uniqueID = 1

func getuniqueID() int {
	id := uniqueID
	uniqueID++
	return id
}

type Store struct {
	id               int
	foodPrice        map[FoodItem]int
	beveragePrice    map[BeverageItem]int
	foodSupply       map[FoodItem]int
	beverageSupply   map[BeverageItem]int
	foodUnitSold     map[FoodItem]int
	beverageUnitSold map[BeverageItem]int
}

func GetStoreConstractor(food, foodProice map[FoodItem]int, beverage, beveragePrice map[BeverageItem]int) *Store {
	return &Store{id: getuniqueID(), foodSupply: food, beverageSupply: beverage, foodPrice: foodProice,
		beveragePrice: beveragePrice, foodUnitSold: make(map[FoodItem]int), beverageUnitSold: make(map[BeverageItem]int)}
}

func (store *Store) purchaseFood(foodItem FoodItem, qty int) {
	if store.foodSupply[foodItem] < qty {
		fmt.Println("Not enough stocks")
		return
	}
	total := store.foodSupply[foodItem]
	store.foodSupply[foodItem] = total - qty
	store.foodUnitSold[foodItem] = qty + store.foodUnitSold[foodItem]
}

func (store *Store) purchaseBeverage(beverageItem BeverageItem, qty int) {
	if store.beverageSupply[beverageItem] < qty {
		fmt.Println("Not enough stocks")
		return
	}
	total := store.beverageSupply[beverageItem]
	store.beverageSupply[beverageItem] = total - qty
	store.beverageUnitSold[beverageItem] = qty + store.beverageUnitSold[beverageItem]
}

func (store *Store) Store(a, b map[int]int) {

}

func (store *Store) GetId() int {
	return store.id
}

func (store *Store) GetFoodPrice() map[FoodItem]int {
	return store.foodPrice
}

func (store *Store) GetBeveragePrice() map[BeverageItem]int {
	return store.beveragePrice
}

func (store *Store) SetId(id int) *Store {
	store.id = id
	return store
}

func (store *Store) SetFoodPrice(foodRates map[FoodItem]int) {
	store.foodPrice = foodRates
}

func (store *Store) SetBeveragePrice(beverageRage map[BeverageItem]int) {
	store.beveragePrice = beverageRage
}
