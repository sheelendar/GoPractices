package main

/*Fires in the Amazon
Programming challenge description:

When data is imported into any software environment, there is a risk of introducing bad or invalid data. An essential
step before doing any data analysis is to first validate the data to be used. This can save time and avoid problems later.
Imagine a data science team investigating forest fires in the Tropical Forests of the Amazon. They have two datasets in CSV format:
1. The first dataset represents the gathered statistics of forest fires by the month and year for different states in Brazil
2. The second dataset has a summary of the total number of fires per year across the entire country
Write a program to validate the first dataset using the summary data from the second.
Input:
line.
Two datasets in CSV format separated by one empty
The first dataset has the following columns:
1. year - the year when forest fires happened
2. state - Brazilian state
3. month - the month when forest fires happened
4. number - the number of reported fires in that year, state, and month
The second dataset has the following columns:
1. year
-the year when forest fires happened
2. number - the total number of reported fires in that*/

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	arr1 := []string{"year,state,month,number",
		"2001,sdfsddf,sdfs,20",
		"2002,sdfsddf,sdfs,10",
		"2005,sdfsddf,sdfs,30",
		"2001,sdfsddf,sdfs,2",
		"2005,sdfsddf,sdfs,2"}
	arr2 := []string{
		"year,number",
		"2001, 22",
		"2002, 10",
		"2005, 32",
	}
	//first := true
	/*scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data := scanner.Text()
		if data == "\n" {
			first = false
		}
		if first {
			arr1 = append(arr1, data)
		} else {
			arr2 = append(arr2, data)
		}
	}*/
	//	fmt.Println(arr1)
	//	fmt.Println(arr2)
	m := make(map[string]int)

	for i := 0; i < len(arr2); i++ {
		if i == 0 {
			continue
		}
		str := strings.Split(arr2[i], ",")
		num, err := strconv.Atoi(strings.TrimSpace(str[1]))
		if err != nil {
			fmt.Println(err)
		}
		m[str[0]] = num
	}
	for i := 0; i < len(arr1); i++ {
		if i == 0 {
			continue
		}
		str := strings.Split(arr1[i], ",")
		num, _ := strconv.Atoi(str[3])
		v, ok := m[str[0]]
		if !ok {
			fmt.Println(false)
			return
		} else {
			v = v - num
			if v < 0 {
				fmt.Println(false)
				return
			} else if v == 0 {
				delete(m, str[0])
			} else {
				m[str[0]] = v
			}
		}
	}
	if len(m) == 0 {
		fmt.Println(true)
		return
	}
	fmt.Println(false)
}
