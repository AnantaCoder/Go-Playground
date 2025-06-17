package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func main() {
	fmt.Println("This is a day 1 problem of advent of code 2024")
	file, err := os.Open("newdata.txt")
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




	


	// OneN:= similarityFounder(left,right)
	// // fmt.Println(OneN)
	// // new slice - left slice 
	// OneX := sliceDifference(left,OneN)

	// OneMul := redundantMultiplier(OneN)
	// targetSlice := append(OneX,OneMul...)
	// target := 0
	// for i := 0; i < len(targetSlice); i++ {
	// 	target = target + targetSlice[i]
	// }
	// fmt.Println("The secret code is ",target)


	// STEP 1: Get similarity-related slice
	similarSlice := similarityFounder(left, right) // This now gives value * count, per left[i]

	// STEP 2: Sum up that slice
	similarityScore := 0
	for _, val := range similarSlice {
		similarityScore += val
	}
	fmt.Println("Similarity Score:", similarityScore)

	// new approach 
	// fmt.Println("The secret code in new approach is ",similarityScore(left,right))



}


func similarityFounder(left []int, right []int) []int {
	countMap := make(map[int]int)
	for _, r := range right {
		countMap[r]++
	}

	var result []int
	for _, val := range left {
		result = append(result, val*countMap[val])
	}
	return result
}


func sliceDifference(slice1 []int, slice2[]int )([]int){
	result := slices.DeleteFunc(slice1, func(x int) bool {
		return slices.Contains(slice2,x)
	})
	return result
}

func redundantMultiplier(numbers []int) []int {
	counts := make(map[int]int)
	for _, num := range numbers {
		counts[num]++ 
	}

	result := make([]int, 0, len(numbers)) 
	for _, num := range numbers {
		
		result = append(result, num*counts[num])
	}

	return result
	
}


func similarityScore(left,right []int)int{
	rightCounts := make(map[int]int)
	for _, v := range right {
		rightCounts[v]++

	}
	score := 0 
	for _, v := range left {
		score += v*rightCounts[v]
	}
	return score
}