package main

import "fmt"

var uniqueID = 1

type User struct {
	id               int
	name             string
	userExpenseSheet map[*User]float64
	totalExpenses    float64
}

func UserConstractor(name string) *User {
	return &User{name: name, id: getUniqueID()}
}

func getUniqueID() int {
	id := uniqueID
	uniqueID++
	return id
}

type Pair struct {
	user    User
	expense float64
}

func (user *User) addUserExpensesSheet(u1 *User, expense float64) {
	if user == u1 {
		return
	}
	user.totalExpenses += expense
	user.userExpenseSheet[u1] += expense
}

func (user *User) printTotalBalance() {
	if user.totalExpenses > 0 {
		fmt.Println(user.name, " owes a total of ", user.totalExpenses)
	} else {
		fmt.Println(user.name, " gets back to total of ", user.totalExpenses*(1))
	}
}

func (user *User) GetId() int {
	return user.id
}

func (user *User) GetName() string {
	return user.name
}

func (user *User) GetUserExpenseSheet() map[*User]float64 {
	return user.userExpenseSheet
}

func (user *User) GetTotalExpenses() float64 {
	return user.totalExpenses
}

func (user *User) SetId(id int) *User {
	user.id = id
	return user
}

func (user *User) SetName(name string) *User {
	user.name = name
	return user
}

func (user *User) SetUserExpenseSheet(userExpenseSheet map[*User]float64) *User {
	user.userExpenseSheet = userExpenseSheet
	return user
}

func (user *User) SetTotalExpenses(totalExpenses float64) *User {
	user.totalExpenses = totalExpenses
	return user
}
