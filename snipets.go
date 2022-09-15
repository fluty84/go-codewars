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
