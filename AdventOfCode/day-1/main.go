package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	fmt.Println("This is a day 1 problem of advent of code 2024")
	file, err := os.Open("data.txt")
	check(err)
	var left, right []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var l, r int
		fmt.Sscanf(scanner.Text(), "%d %d", &l, &r)
		left = append(left, l)
		right = append(right, r)
	}
	// fmt.Println(left)
	// fmt.Println(right)

	fmt.Println(len(left))
	fmt.Println(len(right))

	// indexPop:= 0
	// l_min := left[0]
	// for i := 0; i < len(left); i++ {
	// if left[i] < l_min {
	// 	l_min = left[i]
	// 	indexPop = i
	// }
    // }
	// newLeftArray,poppedElement := removeByIndex(left,indexPop)



	// fmt.Println("The maximum element in left array is ", l_min)
	// fmt.Println("The index to be popped in left array", indexPop)
	// fmt.Println("New array ", newLeftArray,"new length ",len(newLeftArray))
	// fmt.Println("Popped element", poppedElement)




	var remaining []int

	for len(left) > 0 && len(right) > 0 {
		temp_left, indexpopl := minFinder(left)
		temp_right, indexpopr := minFinder(right)

		remaining = append(remaining, int(math.Abs(float64(temp_left-temp_right))))

		left,_ = removeByIndex(left, indexpopl)
		right,_ = removeByIndex(right, indexpopr) // Corrected: remove from right using indexpopr
	}
	sum := 0 
	for i := 0; i < len(remaining); i++ {
		sum = sum + remaining[i]
	}
	fmt.Println(sum)

}
func removeByIndex(slice []int,index int)([]int, int){
	if index<0 || index >= len(slice) {
		fmt.Println("Out of bounds ")
		return slice,index
	}
	removedElement := slice[index]
	newSlice := append(slice[:index],slice[index+1:]...)
	return newSlice,removedElement
}
// first (input)(output)
func minFinder(slice []int)(int,int){
	indexPop := 0
	min := slice[0]
	for i := 0; i < len(slice); i++ {
		if slice[i]<min {
			min = slice[i]
			indexPop = i 
		}
	}
	return min,indexPop
}