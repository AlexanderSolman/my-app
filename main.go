package main

import (
	"fmt"
	"time"
)

// Added comment!
// Another comment
// Another

func addition(num1, num2 int) int {
    return num1 + num2
}

func subtraction(num1, num2 int) int {
    return num1 - num2
}

func multiplication(num1, num2 float64) float64 {
    return num1 * num2
}

func dividen(num1, num2 float64) float64 {
    return num1 / num2
}

func assertType(i interface{}) error {
    if _, ok := i.(int); !ok { return fmt.Errorf("wrong type") }
    return nil
}

func main() {
    num1 := 20
    num2 := 20

    for {
        if err := assertType(num1); err != nil { panic(err) }
        if err := assertType(num2); err != nil { panic(err) }

        fmt.Printf("%v + %v = %v\n", num1, num2, addition(num1, num2))
        fmt.Printf("%v - %v = %v\n", num1, num2, subtraction(num1, num2))
        fmt.Printf("%v * %v = %f\n", num1, num2, multiplication(float64(num1), float64(num2)))
        fmt.Printf("%v / %v = %f\n", num1, num2, dividen(float64(num1), float64(num2)))
        time.Sleep(5 * time.Second)
    }
}
