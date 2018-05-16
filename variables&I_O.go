package main

import "fmt"
import "bufio"  //for standard I/O
import "os"
 func main() {
	var a= "first variable"
	fmt.Println(a)

	var b,c int = 1,2
	fmt.Println(b,c)

	f:= "short"
	fmt.Println(f)

	const n = 50000
	fmt.Println(n)

	// to input a string from user

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter Text: ")
	text,_ := reader.ReadString(' ')  // reads the string upto that character
	fmt.Println(text)

	fmt.Println("Enter Number: ")
    var i int
    _, err := fmt.Scanf("%d", &i)
    fmt.Println(i)		
    fmt.Println(err)		
}