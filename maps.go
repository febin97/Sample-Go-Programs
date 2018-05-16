package main 

import "fmt"

func main() {
	// make maps 

	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println("Map : ",m)

	n := map[string]int{"f1" : -1, "f2" : 1}
	fmt.Println("Map2 :" ,n)

	fmt.Println(m["k1"])
	fmt.Println(len(m))
	delete(m,"k2")
	fmt.Println(m)

	_,prs := m["k2"]
	fmt.Println(prs)
}