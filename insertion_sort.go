package main


func insertionSort(list []int) {
	for i := 1; i < len(list); i++ {
		placeNumber(list, i);
	}
}


func placeNumber(list []int, position int) {
	index := position;

	for index > 0 && list[index - 1] > list[index] {
		list[index-1], list[index] = list[index], list[index-1];
		index--;
	}
}

