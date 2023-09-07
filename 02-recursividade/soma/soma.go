package main

import "fmt"

func soma(array []int, size int) (int) {
	if (size == 0) {
		return 0;
 	} else if (size == 1) {
		return array[0];
	} else {
		return array[0] + soma(array[1:size], size -1)
	}
}

func main() {

	var array = [5]int{1, 2, 3, 4, 5}
	var total = soma(array[:], 5)

	fmt.Println("Array: ", array)
	fmt.Println("Soma: ", total)
}