package main

import "fmt"

func main() {
	fmt.Println("Error Handling")
	fmt.Println(dev(10, 5))
	a, _ := dev(-10, 0) // use _ to ignore the error
	fmt.Println(a)

	b, err := dev(10, 0) // to access the error
	fmt.Println(b, err)

	var (
		result float64
		er     error
	)

	result, er = dev(10, 0)

	fmt.Println(result, er)

	var m int
	fmt.Scan(&m)

	fmt.Println(10 / m)

}

func dev(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Denominator must not be zero")
	}
	return a / b, nil
}
