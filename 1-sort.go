package main

import "fmt"

func main() {
	var scores = []int{90, 70, 50, 80, 60}
	var length = len(scores)

	sort(scores, length)

	for i := 0; i < length; i++ {
		fmt.Printf("%d\n", scores[i])
	}
}

func sort(arrays []int, length int) {

	var minIndex int

	for i := 0; i < length; i++ {
		minIndex = 1

		var minValue = arrays[minIndex]
		for j := i; j < length-1; j++ {
			if minValue > arrays[j+1] {
				minValue = arrays[j+1]
				minIndex = j + 1
			}
		}

		if i != minIndex {
			var temp = arrays[i]
			arrays[i] = arrays[minIndex]
			arrays[minIndex] = temp
		}
	}
}
