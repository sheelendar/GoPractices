package main

import (
	"fmt"
	"sort"
)

func main() {
	author := []struct {
		a_name    string
		a_article int
		a_id      int
	}{
		{"Rahul", 304, 1098},
		{"Rohit", 634, 102},
		{"Anup", 104, 105},
		{"Sam", 10, 108},
		{"Tim", 234, 103},
	}

	sort.Slice(author, func(i, j int) bool {
		return author[i].a_article < author[j].a_article
	})
	fmt.Println(author)
}
