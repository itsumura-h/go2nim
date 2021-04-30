package main

import (
	"C"
	// "fmt"
	// "reflect"
	"unsafe"
	// "github.com/k0kubun/pp"
)
import (
	"fmt"

	"github.com/k0kubun/pp"
)

func fibLogic(n int) int {
	if n < 2 {
		return n
	}
	return fibLogic(n-2) + fibLogic(n-1)
}

var boxInstances map[uintptr]*[]int

//export fibo
func fibo(n int) uintptr {
	var r []int
	for i := 0; i < n; i++ {
		r = append(r, fibLogic(i))
	}

	p := uintptr(unsafe.Pointer(&r))
	if boxInstances == nil {
		boxInstances = make(map[uintptr]*[]int)
	}
	boxInstances[p] = &r
	return p
}

//export getSliceIntVal
func getSliceIntVal(p uintptr, i int) int {
	s := (*[]int)(unsafe.Pointer(p))
	return (*s)[i]
}

//export getSliceLen
func getSliceLen(p uintptr) int {
	s := (*[]int)(unsafe.Pointer(p))
	return len(*s)
}

func main() {
	p := fibo(10)
	pp.Println(p)
	fmt.Println(getSliceLen(p))
	fmt.Println(getSliceIntVal(p, 0))
	fmt.Println(getSliceIntVal(p, 1))
	fmt.Println(getSliceIntVal(p, 2))
	fmt.Println(getSliceIntVal(p, 3))
	fmt.Println(getSliceIntVal(p, 4))
}
