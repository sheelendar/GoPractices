package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const GetUserURL = "https://jsonplaceholder.typicode.com/users"

type UserResponse []User

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}
type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}
type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}
type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

func main() {
	userFetchDetailsHandler(GetUserURL)
}
func userFetchDetailsHandler(url string) {
	client := http.DefaultClient
	res, err := client.Get(url)
	fmt.Println(res.Body)
	if res.StatusCode != http.StatusOK {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	var usersResponse UserResponse
	err = json.NewDecoder(res.Body).Decode(&usersResponse)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(usersResponse)
	for _, user := range usersResponse {
		fmt.Println("Name:  ", user.Name, " UserName: ", user.Username)
		fmt.Println("Add:  ", user.Address)
		fmt.Println()
	}
}
