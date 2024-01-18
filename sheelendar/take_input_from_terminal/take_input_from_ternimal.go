package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("input text:")
	var w1, w2, w3 string
	n, err := fmt.Scanln(&w1, &w2, &w3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("number of items read: %d\n", n)
	fmt.Printf("read line: %s %s %s\n", w1, w2, w3)

	fmt.Println("input text:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err = scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	lineString := scanner.Text()
	fmt.Printf("read line: %s \n", lineString)

}
