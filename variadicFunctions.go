//functions with any number of trailing arguments
//Println is an Example

package main

import "fmt"

func sum(nums ...int) {		// nums will be stored as a slice
    fmt.Println(nums, " ")
}

func sum2(a int ,b int){
	fmt.Println("Sum= ",a+b)
}

func main() {

    sum(1, 2)
    sum(1, 2, 3)

    nums := []int{1, 2, 3, 4}
    sum(nums...)

    sum2(2,3)
}
