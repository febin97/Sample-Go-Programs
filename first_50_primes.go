package main

import (
	"fmt"
	"math"
)
func isPrime(x int) int{
	for j := 2; float64(j)<= math.Sqrt(float64(x)); j++{
		if x%j == 0{
			return -1
		}
	}
	return 1
}
func main() {
	primeCount := 0
	primeSum := 0
	for i:=2 ; i>0 ; i++{
		if isPrime(i) == 1{
			fmt.Printf("%d ",i)
			primeSum = primeSum+i
			primeCount = primeCount+1
			if primeCount >= 50{
				break;
			}
		}
	}
	fmt.Println()
	fmt.Println("Sum=",primeSum)
}
