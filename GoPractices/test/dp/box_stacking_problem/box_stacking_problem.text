package main

import (
	"fmt"
	"math"
	"sort"
)

/*
ou are given a set of n types of rectangular 3-D boxes, where the i^th box has height h(i), width w(i) and depth d(i)
(all real numbers). You want to create a stack of boxes which is as tall as possible, but you can only stack a box on top
of another box if the dimensions of the 2-D base of the lower box are each strictly larger than those of the 2-D base of
the higher box. Of course, you can rotate a box so that any side functions as its base. It is also allowable to use multiple
instances of the same type of box.
*/
type Box struct {
	h    int
	w    int
	d    int
	area int
}

func main() {

	arr := make([]Box, 4)
	arr[0] = Box{h: 4, w: 6, d: 7}
	arr[1] = Box{h: 1, w: 2, d: 3}
	arr[2] = Box{h: 4, w: 5, d: 6}
	arr[3] = Box{h: 10, w: 12, d: 32}

	fmt.Println("The maximum possible height of stack is ", maxStackHeight(arr, 4))
}

func maxStackHeight(arr []Box, size int) int {
	var box []Box
	//box := make([]Box, 3*size)
	for i := 0; i < size; i++ {
		box = append(box, Box{h: arr[i].h, w: arr[i].w, d: arr[i].d})
		box = append(box, Box{h: arr[i].w, w: arr[i].d, d: arr[i].h})
		box = append(box, Box{h: arr[i].d, w: arr[i].h, d: arr[i].w})
	}
	for i := 0; i < 3*size; i++ {
		box[i].area = box[i].w * box[i].d
	}
	sort.Slice(box, func(i, j int) bool {
		return box[i].area > box[j].area
	})
	count := 3 * size
	res := make([]int, count)
	for i := 0; i < count; i++ {
		res[i] = box[i].h
	}

	for i := 0; i < count; i++ {
		res[i] = 0
		val := 0
		for j := 0; j < i; j++ {
			if box[i].w < box[j].w && box[i].d < box[j].d {
				val = int(math.Max(float64(val), float64(res[j])))
			}
		}
		res[i] = val + box[i].h
	}
	max := 0
	for i := 0; i < count; i++ {
		max = int(math.Max(float64(max), float64(res[i])))

	}
	fmt.Print(max)
	fmt.Print(" ")
	return max
}
