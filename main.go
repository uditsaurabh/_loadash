package main

import (
	"fmt"

	arr "github.com/uditsaurabh/_loadash/array"
)

func main() {
	res, err := arr.ChunkIntArray([]int{1, 2, 3, 4, 5}, 3)
	fmt.Println(res, err)
}
