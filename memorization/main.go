package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const LIM = 60

var facts [LIM]uint64

func main() {
	var n uint64 = 60
	summation := uint64(0)
	start := time.Now()
	for i := uint64(0); i < n; i++ {
		summation += sumEuler(i + uint64(1))
	}
	fmt.Println(summation)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)

}

func splitNumbers(n uint64) []uint64 {
	ns := strconv.FormatUint(n, 10)
	arr := strings.Split(ns, "")
	var intArr []uint64
	for _, s := range arr {
		n, _ := strconv.ParseUint(s, 10, 64)
		intArr = append(intArr, n)

	}
	return intArr
}

func factorial(n uint64) (res uint64) {
	if facts[n] != 0 {
		res = facts[n]
		return res
	}

	if n > 0 {
		res = n * factorial(n-1)
		return res
	}

	return 1
}

// este es mi cuello de botella
func minFactorial(n uint64) uint64 {
	var i uint64 = 0
	fmt.Println(n)
	for {
		//		fmt.Print("min factorial")
		//		fmt.Println(i)
		var arr []uint64
		arr = splitNumbers(i)
		result := uint64(0)
		resultFactorial := uint64(0)
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

func sumFactorial(n uint64) uint64 {
	arr := splitNumbers(n)
	result := uint64(0)
	for _, s := range arr {
		result += s
	}
	return result
}

func sumEuler(i uint64) uint64 {
	var n uint64 = minFactorial(i)
	var sum uint64 = sumFactorial(n)
	return sum
}
