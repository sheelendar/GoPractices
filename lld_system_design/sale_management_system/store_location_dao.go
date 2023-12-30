package main

var uniqueCityID = 1

func getUniqueCityID() int {
	id := uniqueCityID
	uniqueCityID++
	return id
}

type City struct {
	id             int
	foodPrices     map[FoodItem]int
	beveragePrices map[BeverageItem]int
	stores         []*Store
}

func GetCityConstractor(foodPrices map[FoodItem]int, beveragePrices map[BeverageItem]int) *City {
	return &City{id: getUniqueCityID(), foodPrices: foodPrices, beveragePrices: beveragePrices}
}

func (city *City) addStore(store *Store) {
	store.SetFoodPrice(city.foodPrices)
	store.SetBeveragePrice(city.beveragePrices)
	city.stores = append(city.stores, store)
}

func (city *City) purchageFood(storeId int, foodItem FoodItem, qty int) {
	for _, store := range city.stores {
		if store.GetId() == storeId {
			store.purchaseFood(foodItem, qty)
			break
		}
	}
}
func (city *City) purchageBeverage(storeId int, beverageItem BeverageItem, qty int) {
	for _, store := range city.stores {
		if store.GetId() == storeId {
			store.purchaseBeverage(beverageItem, qty)
			break
		}
	}
}

var uniqueStateID = 1

func getUniqueStateID() int {
	id := uniqueStateID
	uniqueStateID++
	return id
}

type State struct {
	id     int
	cities map[*City]bool
}

func GetStateConstractor() *State {
	return &State{id: getUniqueCityID(), cities: map[*City]bool{}}
}

func (state *State) addCity(city *City) {
	state.cities[city] = true
}

func (state *State) GetCities() map[*City]bool {
	return state.cities
}

func (state *State) purchageFood(cityID, storeId int, foodItem FoodItem, qty int) {
	for city, _ := range state.cities {
		if city.id == cityID {
			city.purchageFood(storeId, foodItem, qty)
			break
		}
	}
}
func (state *State) purchageBeverage(cityID, storeId int, beverageItem BeverageItem, qty int) {
	for city, _ := range state.cities {
		if city.id == cityID {
			city.purchageBeverage(storeId, beverageItem, qty)
			break
		}
	}
}

type System struct {
	id     int
	states []*State
}

func GetSystemConstractor(id int) *System {
	return &System{id: id}
}
func (system *System) addState(state *State) {
	system.states = append(system.states, state)
}

func (system *System) purchageFood(stateID, cityID, storeId int, foodItem FoodItem, qty int) {
	for _, state := range system.states {
		if state.id == stateID {
			state.purchageFood(cityID, storeId, foodItem, qty)
			break
		}
	}
}
func (system *System) purchageBeverage(stateID, cityID, storeId int, beverageItem BeverageItem, qty int) {
	for _, state := range system.states {
		if state.id == stateID {
			state.purchageBeverage(cityID, storeId, beverageItem, qty)
			break
		}
	}
}
