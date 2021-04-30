package main

import (
	"C"
	"fmt"
	"unsafe"

	"github.com/k0kubun/pp"
)

func fibLogic(n int) int {
	if n < 2 {
		return n
	}
	return fibLogic(n-2) + fibLogic(n-1)
}

type GoIntSlice = []int
type PGoIntSlice = uintptr

var boxInstances map[uintptr]*GoIntSlice

//export fibo
func fibo(n int) PGoIntSlice {
	var r GoIntSlice
	for i := 0; i < n; i++ {
		r = append(r, fibLogic(i))
	}

	p := PGoIntSlice(unsafe.Pointer(&r))
	if boxInstances == nil {
		boxInstances = make(map[PGoIntSlice]*GoIntSlice)
	}
	boxInstances[p] = &r
	return p
}

//export getGoIntSliceVal
func getGoIntSliceVal(p PGoIntSlice, i int) int {
	s := (*GoIntSlice)(unsafe.Pointer(p))
	return (*s)[i]
}

//export getGoIntSliceLen
func getGoIntSliceLen(p PGoIntSlice) int {
	s := (*GoIntSlice)(unsafe.Pointer(p))
	return len(*s)
}

func main() {
	p := fibo(10)
	pp.Println(p)
	fmt.Println(getGoIntSliceLen(p))
	fmt.Println(getGoIntSliceVal(p, 0))
	fmt.Println(getGoIntSliceVal(p, 1))
	fmt.Println(getGoIntSliceVal(p, 2))
	fmt.Println(getGoIntSliceVal(p, 3))
	fmt.Println(getGoIntSliceVal(p, 4))
}
