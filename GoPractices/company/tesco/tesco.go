package main

import (
	"fmt"
	//"net/http"
)

type Shift struct {
	start int
	end   int
}

func main() {

	//s := http.ListenAndServe("", nil)
	//router:=mux.NewRouter()
	shifts := []Shift{{8, 10}, {10, 12}, {14, 19}, {19, 21}}

	result := calculateShiftTiming(shifts)
	for i := 0; i < len(result); i++ {
		fmt.Print(result[i].start)
		fmt.Print(" -> ")
		fmt.Print(result[i].end)
		fmt.Println()
	}
}

func calculateShiftTiming(shifts []Shift) []Shift {
	size := len(shifts)
	if size == 0 {
		return nil
	}
	if size == 1 {
		return shifts
	}
	var result []Shift

	start := shifts[0].start
	end := shifts[0].end

	for i := 1; i < size; i++ {
		if end >= shifts[i].start {
			end = shifts[i].end

		} else {
			result = append(result, Shift{start: start, end: end})
			start = shifts[i].start
			end = shifts[i].end
		}
	}
	result = append(result, Shift{start: start, end: end})
	return result
}
