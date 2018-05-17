package main 

import "fmt"

func add(a int, b int ) int {
	return a+b
}
func addAdd(a,b,c int ) int {
	return a+b+c
}
func vals2()(int,int) {			//Mutiple Return Values
	return 3,7
}
func vals3()(int,int,int) {			//Mutiple Return Values
	return 3,7,10
}
func main() {
	fmt.Println("1+2=",add(1,2))
	fmt.Println("1+2+3=",addAdd(1,2,3))
	//fmt.Println("1+2+3=",vals())
	a,b := vals2()
	fmt.Println(a," ",b)
	_,_,c := vals3()
	fmt.Println(c)


}