package main

import (
	"fmt"
	"sort"
	"sync"
)

func findSpecialSliceWMinVal(specialSlices []SpecialSlice) int {
	min_index := -1

	var min_num int
	for i := 0; i < len(specialSlices); i++ {
		if specialSlices[i].index < specialSlices[i].length() {
			min_index = i
			break
		}
	}
	if min_index == -1 {
		return -1
	}
	min_num = specialSlices[min_index].data[specialSlices[min_index].index]

	for i := 0; i < len(specialSlices); i++ {
		if specialSlices[i].index < specialSlices[i].length() {
			nextValue := specialSlices[i].data[specialSlices[i].index]
			if nextValue < min_num {
				min_index = i
				min_num = nextValue
			}
		}
	}
	return min_index
}
func collectSortedArrays(arrOfSpecialSlices []SpecialSlice, startChan chan []int) []SpecialSlice {

	for arrayFromChan := range startChan {
		fmt.Print(arrayFromChan)
		arrOfSpecialSlices = append(arrOfSpecialSlices, SpecialSlice{
			index: 0,
			data:  arrayFromChan,
		})
	}
	return arrOfSpecialSlices
}

func goRoutine(sliceint []int, result chan []int, wg *sync.WaitGroup) {
	sort.Ints(sliceint)
	result <- sliceint
	wg.Done()
	return
}

type SpecialSlice struct {
	index int
	data  []int
}

func (s SpecialSlice) length() int {
	return len(s.data)
}

func main() {
	const chunksAmount = 4
	var numberOfSlice, sliceSize int
	var wg sync.WaitGroup
	var resultSlicebig []int
	nuberSlice := make([]int, 0)
	var specialSlices []SpecialSlice

	fmt.Println("Enter all ints you need to sort ( type integer and push enter) \nIf you want to end, push enter TWO TIMES")
	for {
		if sliceSize < 4 {
			_, err := fmt.Scanf("%d", &numberOfSlice)
			if err != nil {
				fmt.Print("Enter nums\n")
				continue
			}
			nuberSlice = append(nuberSlice, numberOfSlice)
			sliceSize = len(nuberSlice)
		} else {
			_, err := fmt.Scanf("%d", &numberOfSlice)
			if err != nil {
				break
			}
			nuberSlice = append(nuberSlice, numberOfSlice)
			sliceSize = len(nuberSlice)
		}

	}
	minChankSize := sliceSize / 4
	chunkSizes := [4]int{}

	for i := 0; i < chunksAmount; i++ {
		chunkSizes[i] = minChankSize
	}

	elementsLeft := sliceSize % chunksAmount
	for i := 0; i < elementsLeft; i++ {
		chunkSizes[i] = chunkSizes[i] + 1
	}

	c := make(chan []int, chunksAmount)
	var nextStartIndex = 0
	for i := 0; i < chunksAmount; i++ {
		wg.Add(1)
		go goRoutine(nuberSlice[nextStartIndex:nextStartIndex+chunkSizes[i]], c, &wg)
		nextStartIndex = nextStartIndex + chunkSizes[i]
	}
	wg.Wait()
	close(c)

	fmt.Print("Sorted chunks: ")
	specialSlices = collectSortedArrays(specialSlices, c)
	for {
		sliceIndex := findSpecialSliceWMinVal(specialSlices)
		if sliceIndex < 0 {
			break
		}
		specialElement := specialSlices[sliceIndex]
		resultSlicebig = append(resultSlicebig, specialElement.data[specialElement.index])
		specialSlices[sliceIndex].index = specialSlices[sliceIndex].index + 1

	}
	fmt.Println("\nSolid sorted slice:", resultSlicebig)

}
