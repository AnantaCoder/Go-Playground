package main

import (
	"fmt"
	"bufio"
	"math"
	"os"
	"sort"
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

	sort.Ints(left)
	sort.Ints(right)


	// fmt.Println(left)
	// fmt.Println(right)

	var merger []int
	for i := 0; i < len(right); i++ {
		temp := math.Abs(float64(left[i] - right[i]))
		merger = append(merger, int(temp))
	}
	sum := 0
	for i := 0; i < len(merger); i++ {
		sum = sum+ merger[i]
	}

	fmt.Println("The sum is ", sum)
	// The code is  2086478
     


	// it is right but it worked because i have 2 array of equal sizes .
	

}
