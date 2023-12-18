package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(checkIpAddress("12.34.56.oops"))
	fmt.Println(checkIpAddress("255.255.255.255"))
}
func checkIpAddress(ip string) bool {
	if len(ip) <= 0 {
		return false
	}

	arr := strings.Split(ip, ".")

	n := len(arr) // arr:= {255,255,255,255}

	if n != 4 {
		return false
	} else {
		for i := 0; i < n; i++ {
			v, err := strconv.Atoi(arr[i])
			if err != nil {
				return false
			}
			if !(v >= 0 && v <= 255) {
				return false
			}

		}
	}
	return true
}
