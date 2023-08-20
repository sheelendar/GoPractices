package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//s := "1:23am-1:08am"
	//s := "1:23am-3:23am"
	s := "12:30pm-12:00am"
	result := findMinutes(s)
	fmt.Println(result) // Print the result
}

func findMinutes(s string) string {
	times := strings.Split(s, "-")
	h1 := getMinutes(times[0])
	h2 := getMinutes(times[1])
	if h2-h1 < 0 {
		return strconv.Itoa(h2 + 1440 - h1)
	}
	return strconv.Itoa(h2 - h1)
}

func getMinutes(time string) int {
	times := strings.Split(time, ":")
	hour, _ := strconv.Atoi(times[0])
	minute, _ := strconv.Atoi(times[1][:2])
	zone := times[1][2:]
	if zone == "am" {
		if hour == 12 {
			return minute
		} else {
			return 60*hour + minute
		}
	} else {
		buffer := 720
		if hour == 12 {
			return minute + buffer
		} else {
			return 60*hour + minute + buffer
		}
	}
}
