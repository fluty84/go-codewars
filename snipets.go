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
