package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")

	var list = []int{2, 801, 53, 330, 9, 4, -1}

	radixSort(list);

	fmt.Println(list);
}

