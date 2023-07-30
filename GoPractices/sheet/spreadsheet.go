package main

import (
	"fmt"
	"strings"
)

/*Prompt:
Today we’re going to build a basic spreadsheet application like Google sheets or Excel but much simpler. Our spreadsheet, let’s call it OpenSheet, will only support cells which hold either integers or formulas that sum two cells.

You are tasked with writing a program that handles this functionality for OpenSheet. You can make any decisions you want regarding how this program is organized, but there must be some pancake_sort of setter/getter methods that can be called by the application for any given cell. All inputs will be strings.

For setting you can expect two inputs: the cell location and the cell value.

Example of how your setter could look
set_cell("C1", "45")
set_cell("B1", "10")
set_cell("A1", "=C1+B1")

For getting you will be provided one input that is the cell location.

Example of how your getter could look
get_cell("A1") # should return 55 in this case*/

var spreadsheet = make(map[string]int)

func main() {

	set_cell("C1", 45)
	set_cell("B1", 10)
	set_cell("A1", "=C1+B1")
	fmt.Println(get_cell("A1")) //55
	set_cell("B1", 20)
	set_cell("A1", "=C1+B1")
	fmt.Println(get_cell("A1")) // 65
}

func get_cell(col string) int {
	if v, ok := spreadsheet[col]; ok {
		return v
	}
	return 0
}

func set_cell(col string, input interface{}) {
	if value, ok := input.(int); ok {
		spreadsheet[col] = value
	} else if op, ok := input.(string); ok {
		operand := strings.Split(op, "+")
		cel2 := operand[1]
		operandWithExtra := strings.Split(operand[0], "=")
		cel1 := operandWithExtra[1]
		calculateOperation(col, cel1, cel2)
		/*fmt.Println(cel1)
		fmt.Println(cel2)
		fmt.Println(spreadsheet[col])*/
	}

}

func calculateOperation(col string, cel1 string, cel2 string) {
	if len(spreadsheet) <= 0 {
		return
	}
	val1 := 0
	val2 := 0

	if v, ok := spreadsheet[cel1]; ok {
		val1 = v
	}

	if v, ok := spreadsheet[cel2]; ok {
		val2 = v
	}

	spreadsheet[col] = val1 + val2
}
