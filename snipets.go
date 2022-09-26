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

Exercise: Images
Remember the picture generator you wrote earlier? Let's write another one, but this time it will return an implementation of image.Image instead of a slice of data.

Define your own Image type, implement the necessary methods, and call pic.ShowImage.

Bounds should return a image.Rectangle, like image.Rect(0, 0, w, h).

ColorModel should return color.RGBAModel.

At should return a color; the value v in the last picture generator corresponds to color.RGBA{v, v, 255, 255} in this one.

package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
	)

type Image struct{
	Width, Height int
	r, g, b, a uint8
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0,0,img.Width,img.Height)
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) At(x, y int) color.Color {
	return color.RGBA{img.r, img.g, img.b, img.a}
}

func main() {
	m := Image{220,340,60,144,150,200}
	pic.ShowImage(m)
}

Exercise: Equivalent Binary Trees
1. Implement the Walk function.

2. Test the Walk function.

The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.

Create a new channel ch and kick off the walker:

go Walk(tree.New(1), ch)
Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.

3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.

4. Test the Same function.

Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

The documentation for Tree can be found here.

package main

import (
	"golang.org/x/tour/tree"
	"fmt"
	)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int){
	
	var walker func (t * tree.Tree) // creo variable walker que es de tipo funcion 
	//fmt.Printf("t es ahora %v \n", t) 
	walker = func (t *tree.Tree) {  // inicio la variable con el contenido de la función 
		
		
        if (t == nil) { // mientras t tenga valor, el programa va a seguir ejecutándose
            return
        }
        walker(t.Left) // busca valores recursivamente hacia la izquierda, si no los hay sale de la función
        ch <- t.Value // devuelve al canal el valor de su iteracion recursiva
        walker(t.Right) // idem que left pero a la derecha
    }
    walker(t)
  	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	
	chT1 := make(chan int)
	chT2 := make(chan int)
	
	go Walk(t1, chT1)
	go Walk(t2, chT2)
	
	for {
		chT1response, okT1 := <- chT1 // almaceno la respuesta y la señal de canal activo
		//fmt.Println(chT1response, okT1)
		chT2response, okT2 := <- chT2
		
		if chT1response != chT2response  { // comparo si las respuestas van siendo las mismas, si alguna no lo es devuelvo false
            return false
        }

        if !okT1 || !okT2 { //si se cierra un canal salgo del bucle
            break
        }
	}
	
	return true // si no se ha salido de la funcion con un false, devuelvo true
}

func main() {

 ch := make(chan int)

 go Walk(tree.New(1), ch)
 
 fmt.Println("are trees equal?", Same(tree.New(1), tree.New(1)))
 
 for recived := range ch {
 	fmt.Println(recived)
 }

}

/* Instructions
In this exercise you'll be working with savings accounts. Each year, the balance of your savings account is updated based on its interest rate. The interest rate your bank gives you depends on the amount of money in your account (its balance):

3.213% for a balance less than 0 dollars (balance gets more negative).
0.5% for a balance greater than or equal to 0 dollars, and less than 1000 dollars.
1.621% for a balance greater than or equal to 1000 dollars, and less than 5000 dollars.
2.475% for a balance greater than or equal to 5000 dollars.
You have four tasks, each of which will deal your balance and its interest rate. */

package interest

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    switch {
        case balance < 0:
    	return 3.213
        case balance < 1000:
    	return 0.5
        case balance < 5000:
    	return 1.621
        default:
    	return 2.475
    }
	panic("Please implement the InterestRate function")
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    switch {
        case balance < 0:
    	return balance * 0.03213
        case balance < 1000:
    	return balance * 0.005
        case balance < 5000:
    	return balance * 0.01621
        default:
    	return balance * 0.02475
    }
	panic("Please implement the Interest function")
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
    switch {
        case balance < 0:
    	return balance * 0.03213 + balance
        case balance < 1000:
    	return balance * 0.005 + balance
        case balance < 5000:
    	return balance * 0.01621 + balance
        default:
    	return balance * 0.02475 + balance
    }
	panic("Please implement the AnnualBalanceUpdate function")
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    
    totalYears := 0 

    var AnnualBalanceUpdate func(balance float64, targetBalance float64) int 
    
    AnnualBalanceUpdate = func(balance float64, targetBalance float64) int {
    
    if balance > targetBalance {
        return totalYears
    }
    
	totalYears++
    switch {
        case balance < 0:
    	balance = balance * 0.03213 + balance
        case balance < 1000:
    	balance = balance * 0.005 + balance
        case balance < 5000:
    	balance = balance * 0.01621 + balance
        default:
    	balance = balance * 0.02475 + balance
    }
    return AnnualBalanceUpdate(balance, targetBalance)
    }
AnnualBalanceUpdate(balance, targetBalance)
    
return totalYears
    
	panic("Please implement the YearsBeforeDesiredBalance function")
}


