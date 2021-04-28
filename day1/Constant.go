package main

import (
    "fmt"
    "math"
)

const constant string = "Hello"

func main() {
    fmt.Println(constant)
	var something = constant + " " + "World"
	fmt.Println(something)

    const n = 500000000

    const d = 3e20 / n
    fmt.Println(d)

    fmt.Println(int64(d))

    fmt.Println(math.Sin(n))

	var testInt int64 =  int64( 10 * math.Sin(n))
	fmt.Println(testInt)
	fmt.Println(float32(testInt*1.0)/10.0)
}