package main

import "fmt"

func showBackingArray(varName string, slice *[]string) {
	fmt.Printf("[%s]\taddr[%p] len[%d] cap[%d] ", varName, slice, len(*slice), cap(*slice))

	for i, s := range *slice {
		fmt.Printf("[%d] addr:[%p] value:[%#v]\n", i, &(*slice)[i], s)
	}
}

func main() {
	//1.declare using var
	var slice1 []string
	showBackingArray("slice1", &slice1)
	//fmt.Printf("[%s]\t addr[%p]len[%d]cap[%d] value[%#v]\n", "slice1", &slice1, len(slice1), cap(slice1), (slice1))

	//2.short variable
	slice2 := []string{}
	showBackingArray("slice2", &slice2)
	//fmt.Printf("[%s]\t addr[%p]len[%d]cap[%d] value[%#v]\n", "slice2", &slice2, len(slice2), cap(slice2), (slice2))

	//2.make slice
	slice3 := make([]string, 5)
	showBackingArray("slice3", &slice3)
	//fmt.Printf("[%s]\t addr[%p]len[%d]cap[%d] value[%#v]\n", "slice3", &slice3, len(slice3), cap(slice3), (slice3))

	//4. declare with make & leng
	slice4 := make([]string, 5, 8)
	showBackingArray("slice4", &slice4)

	//5. declare with initival value
	slice5 := []string{"A", "B", "C", "D", "E"}
	showBackingArray("slice5", &slice5)

	slice5 = append(slice5, "F")
	showBackingArray("slice5", &slice5)

}
