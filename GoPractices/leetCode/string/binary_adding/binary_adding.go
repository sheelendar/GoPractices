package main

import (
	"fmt"
)

func main() {
	fmt.Println(addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	num1 := []byte(a)
	num2 := []byte(b)
	i := len(num1) - 1
	j := len(num2) - 1
	var result []byte
	carry := 0
	for i >= 0 && j >= 0 {
		sum := int(num1[i]-'0') + int(num2[j]-'0') + carry
		rem := sum % 2
		result = append(result, byte(rem+'0'))
		carry = sum / 2
		i--
		j--
	}
	for i >= 0 {
		sum := int(num1[i]-'0') + carry
		rem := sum % 2
		result = append(result, byte(rem+'0'))
		carry = sum / 2
		i--
	}
	for j >= 0 {
		sum := int(num2[j]-'0') + carry
		rem := sum % 2
		result = append(result, byte(rem+'0'))
		carry = sum / 2
		j--
	}
	if carry > 0 {
		result = append(result, '1')
	}
	k := 0
	r := len(result) - 1
	for k < r {
		result[k], result[r] = result[r], result[k]
		k++
		r--
	}
	return string(result)
}
