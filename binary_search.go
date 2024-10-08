package main


func binarySearchRecursive(list []int, search, start, end int) int {
	if end <= start { return -1 };

	middle := start + (end - start) / 2;

	result := list[middle];

	if result == search {
		return middle;
	}

	if result > search {
		return binarySearchRecursive(list, search, start, middle+1);
	}

	return binarySearchRecursive(list, search, middle+1, end);
}



func binarySearch(list []int, search int) int {
	return binarySearchRecursive(list, search, 0, len(list))
}

