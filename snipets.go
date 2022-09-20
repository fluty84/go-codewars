// Multipling two numbers

package kata

func Multiply(n1 int, n2 int) int {

	total := 0
	total = n1 * n2
	return total
}

// Summation
//Write a program that finds the summation of every number from 1 to num.
//The number will always be a positive integer greater than 0.

func Summation(n int) int {
	// the sleeper must awaken!
	count := 0
	for i := 0; i < n; i++ {
		count = count + i + 1
	}
	return count
}

// What is between
//Complete the function that takes two integers (a, b, where a < b) and return an array of all integers between the input parameters, including them.

func Between(a, b int) []int {

	numbArr := []int{a}

	for i := a; i < b; i++ {
		numbArr = append(numbArr, i+1)
	}
	return numbArr
}

// Given a non-empty array of integers, return the result of multiplying the values together in order.

func Grow(arr []int) int {

	mult := 1

	for i := 0; i < len(arr); i++ {
		mult = arr[i] * mult
	}

	return mult
}

//Consider an array/list of sheep where some sheep may be missing from their place.
// We need a function that counts the number of sheep present in the array (true means present).

func CountSheeps(numbers []bool) int {

	sheeps := 0

	for _, sheep := range numbers {
		if sheep == true {
			sheeps = sheeps + 1
		}
	}

	return sheeps
}

Complete the square sum function so that it squares each number passed into it and then sums the results together.

package kata
import "fmt"

func SquareSum(numbers []int) int {
    // your code here
  sum:=0
  for i:=0; i < len(numbers); i++{
    fmt.Println(numbers[i])
    sum += numbers[i]*numbers[i]
  }
  return sum
}

Exercise: Slices
Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit unsigned integers. When you run the program, it will display your picture, interpreting the integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)

package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {

	picture := make([][]uint8, dy)

	for y, _ := range picture {

		data := make([]uint8, dx)

		for x, _ := range data {
			data[x] = uint8(x ^ y)
		}
		
		picture[y] = data
	}

	return picture
}

func main() {
	pic.Show(Pic)

}

Exercise: Maps
Implement WordCount. It should return a map of the counts of each “word” in the string s. The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.

package main

import (
	"golang.org/x/tour/wc"
	//"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	
	sliced := strings.Fields(s)
	
	countMap := make(map[string]int)
	
	for _, value := range sliced {
	
		countMap[value] += 1
	}
	
	return countMap
}

func main() {
	wc.Test(WordCount)
}


Exercise: Fibonacci closure
Let's have some fun with functions.

Implement a fibonacci function that returns a function (a closure) that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).

package main

import "fmt"


func fibonacci() func(int) int {	
	
	solution := 0
	
	return func(x int) int {
	  solution = FibonacciRecursion(x)
	  
	  return solution
	}
}


// fibonacci is a function that returns
// a function that returns an int.

func fibonacci() func(int) int {
	
	return func(x int) int {
	  return FibonacciRecursion(x)
	}		
}

func main() {
	f := fibonacci()
	for i := 0; i < 15; i++ {
		fmt.Println(f(i))
	}
}


