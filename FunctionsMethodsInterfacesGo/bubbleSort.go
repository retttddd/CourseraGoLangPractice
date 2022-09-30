package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input []int
	fmt.Print("Enter up to 10 integers or type x to finalise sequence \n")
	for i := 0; i < 10; i++ {

		var value string
		fmt.Print("Enter integer # ", i, "   ")
		fmt.Scan(&value)
		if value == "x" || value == "X" {
			break
		}

		valueint, err := strconv.Atoi(value)
		if err != nil {
			i--
			fmt.Println("Invalid syntax. You can enter only integers")
			continue
		}
		input = append(input, valueint)

	}
	fmt.Println("\n--- Unsorted --- \n\n", input)
	BubbleSort(input)
	fmt.Println("\n--- Sorted ---\n\n", input, "\n")
}

func BubbleSort(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j)
			}
		}
	}

}

func Swap(arr []int, j int) {

	arr[j], arr[j+1] = arr[j+1], arr[j]

}
