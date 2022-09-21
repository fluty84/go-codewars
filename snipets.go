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

Exercise: Errors
Copy your Sqrt function from the earlier exercise and modify it to return an error value.

Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string{
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
	}

func Sqrt(x float64) (float64, error) {
	
	if x < 0 {
	return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

Implement a Reader type that emits an infinite stream of the ASCII character 'A'.

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func (m MyReader) Read(b []byte) (n int, err error) {
	for i := range b {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}

Exercise: rot13Reader
A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.

For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).

Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.

The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method.

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (payload *rot13Reader) Read(b []byte) (n int, err error) {
	
	n, err = payload.r.Read(b)
	
	for i := range b {
	
		if b[i] >= 'a' && b[i] <= 'z' {
        // Rotate lowercase letters 13 places.
        	if b[i] >= 'm' {
            b[i] = b[i] - 13
        	} else {
            b[i] = b[i] + 13
        	}
    	}else if b[i] >= 'A' && b[i] <= 'Z'{ 
			// Rotate upercase letters 13 places.
        	if b[i] >= 'M' {
            b[i] = b[i] - 13
        	} else {
            b[i] = b[i] + 13
        	}
		}
	
	}
	return 
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

