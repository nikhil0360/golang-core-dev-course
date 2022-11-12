package main

// NOTE: you can also import like this
/*
import "fmt"
import "math"
)
*/
import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
)

// example of var
var c, java, python bool

// := (short variable declaration is not available outside functions)

// NOTE: variables can also be "factored" into blocks like imports
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// NOTE: Numeric constants are high-precision values.
const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

// NOTE: Constants cannot be declared using the := syntax.
const Pi = 3.14

// NOTE: type comes after the variable
// NOTE: you can also write the below function as:
/*
func add(x, y int) int {
	return x + y
}
*/
func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// NOTE: named return values, naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println("Hello, 世界")
	fmt.Printf("Square root of 7 is %g.\n", math.Sqrt(7))
	fmt.Printf("My favourite random number is %d.\n", rand.Intn(10))

	// NOTE: next line will not work because all exported variables in go start with captial letters
	// fmt.Println(math.pi)
	fmt.Println(math.Pi)

	fmt.Println(add(11, 11))
	fmt.Println(swap("first", "second"))

	a, b := split(100)
	fmt.Println("This is a example of split function", a, b)

	// var can be used to declare variable, they are initialised with their corresponding 0 value
	var i int
	fmt.Println(i, c, java, python)

	// use %T for type, and %v for variables
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	fmt.Println(Pi)

}
