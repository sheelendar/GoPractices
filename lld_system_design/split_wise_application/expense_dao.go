package main

import "fmt"

type Split int

const (
	EQUAL = iota + 1
	EXACT
	PERCENT
)

var ExpenseUniqueID = 1

type SplitWise struct {
	users map[*User]bool
}

func (splitWise *SplitWise) calculateExpense(expense *Expense) {
	creator := expense.creater

	switch expense.split {
	case EQUAL:

	case EXACT:
	}

}

func (splitWise *SplitWise) verifyUser(user *User, expese *Expense) bool {
	isFind := false
	for _, u := range expese.defaultUsers {
		if u == user {
			isFind = true
		}
	}
	if !isFind {
		expese.defaultUsers = append(expese.defaultUsers, user)
	}
	for _, u := range expese.defaultUsers {
		if !splitWise.users[u] {
			return false
		}
	}
	return true
}

func (splitWise *SplitWise) addExpense(expese *Expense) {
	if !splitWise.verifyUser(expese.creater, expese) {
		fmt.Println("No user found")
		return
	}
	splitWise.calculateExpense(expese)
}

func (splitWise *SplitWise) printBalaceForAllUSer() {
	for user, _ := range splitWise.users {
		fmt.Println(user.name, "  ", user.totalExpenses)
	}
}
func (splitWise *SplitWise) registeredUSer(user *User) {
	if !splitWise.users[user] {
		splitWise.users[user] = true
	}
}

type Expense struct {
	id                  int
	description         string
	split               Split
	percentDistribution []float64
	exactDistribution   []float64
	creater             *User
	defaultUsers        []*User
	exactTotalAmount    float64
}

func ExpenseConstractor(user *User, split Split, users []*User, expense float64) *Expense {
	return &Expense{creater: user, split: split, defaultUsers: users, exactTotalAmount: expense, id: getExpenseUniqueID()}
}

func getExpenseUniqueID() int {
	id := ExpenseUniqueID
	ExpenseUniqueID++
	return id
}
