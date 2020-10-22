package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	var n int = 45
	summation := 0
	start := time.Now()
	for i := 0; i < n; i++ {
		summation += sumEuler(i + 1)
	}
	fmt.Println(summation)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)

}

func splitNumbers(n int) []int {
	ns := strconv.Itoa(n)
	arr := strings.Split(ns, "")
	var intArr []int
	for _, s := range arr {
		n, _ := strconv.Atoi(s)
		intArr = append(intArr, n)

	}
	return intArr
}

func factorial(n int) int {
	result := 0
	if n == 0 {
		return 1
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			result += n
		} else {
			result *= n - i
		}

	}

	return result
}

// este es mi cuello de botella
func minFactorial(n int) int {
	var i int = 0
	for {
		arr := splitNumbers(i)
		result := 0
		resultFactorial := 0
		for _, s := range arr {
			result += factorial(s)
		}
		arrFactorial := splitNumbers(result)
		for _, s := range arrFactorial {
			resultFactorial += s
		}
		if n == resultFactorial {
			if n == 0 || n == 1 {
				return 1
			}
			return i
		}
		i++
	}
}

func sumFactorial(n int) int {
	arr := splitNumbers(n)
	result := 0
	for _, s := range arr {
		result += s
	}
	return result
}

func sumEuler(i int) int {
	n := minFactorial(i)
	sum := sumFactorial(n)
	return sum
}
