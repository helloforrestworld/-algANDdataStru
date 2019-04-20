package main

import (
	"fmt"
	"math"
	"reflect"
	"runtime"
)

func calculate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)
	}
}

//func div(a, b int) (int, int) {
//	return a / b, a % b
//}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling funtion %s with args "+"(%d, %d) \n", opName, a, b)
	return op(a, b)
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func sumArgs(values ...int) int {
	sum := 0
	for i := range values {
		sum += values[i]
	}
	return sum
}

func main() {
	ret, err := calculate(3, 4, "*")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}

	q, r := div(13, 3)
	fmt.Println(q, r)
	fmt.Println(apply(
		func(a, b int) int {
			return int(math.Pow(float64(a), float64(b)))
		}, 2, 3))
	fmt.Println(apply(pow, 2, 2))
	fmt.Println(sumArgs(1, 2, 3))
}
