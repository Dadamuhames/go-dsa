package main

import "math"


// get max value
func getMaxOfSlice(list []int) int {
	maxValue := 0;
	
	for i := 0; i < len(list); i++ {
		if list[i] > maxValue {
			maxValue = list[i];
		}
	}

	return maxValue;
}


// calculate length of max value
func getIterationsCount(list []int) int {
	maxValue := getMaxOfSlice(list);
	count := 0;

	for maxValue > 0 {
		count++;
		maxValue /= 10;
	}

	return count;
}

// get digit of number by index
func getDigitByIndex(num int, index int) int {
	firstPowerOfTen := math.Pow10(index+1);
	secondPowerOfTen := math.Pow10(index);

	result := num % int(firstPowerOfTen) / int(secondPowerOfTen);

	if result < 0 {result = 0};

	return result;
}


func radixSort(list []int) {
	var radixArray [10][]int = [10][]int {};

	iterationCount := getIterationsCount(list);

	for i := 0; i < iterationCount; i++ {
		for l := 0; l < len(list); l++ {
			num := getDigitByIndex(list[l], i);	
		
			radixArray[num] = append(radixArray[num], list[l]);
		}

		var newArray []int = []int {};

		for j := 0; j < 10; j++ {
			for l := 0; l < len(radixArray[j]); l++ {
				newArray = append(newArray, radixArray[j][l]);	
			}
		}

		copy(list, newArray);
		radixArray = [10][]int {};
	}	
}

