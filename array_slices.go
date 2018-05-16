package main 

import "fmt"

 func main() {
	var a [5]int
	fmt.Println("Arr: ",a)
	//initial values if not assigned will be
	// 0 for int and float
	// false for bool
	// "" for string
	 b := [4]int{1,2,3,4}
	fmt.Println(b[2])
	fmt.Println(len(b))

	var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)

    //Slices

    s := make([]string,3)
    fmt.Println("Slice:",s)
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println(s[2],len(s))
    s = append(s,"d")
    fmt.Println(s,len(s))

    //copy  a slice

    c := make([]string,len(s))
    copy(c,s)
    fmt.Println("C:",c)

    l := c      //copy
    fmt.Println("L:  ",l)
    l = c[1:3] 				// 1 included 3 Not
    fmt.Println("L:  ",l)
    fmt.Println(c[1:])
    fmt.Println(c[:3])

    t := []string{"g", "h", "i"}
    fmt.Println("dcl:", t)
    t = append(t,"j")
    fmt.Println("dcl:", t)
    
}