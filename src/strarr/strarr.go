package main

import (
	"C"
	"strconv"
	"unsafe"
)

type GoStrSlice = []string
type PGoStrSlice = uintptr

var boxInstances map[PGoStrSlice]*GoStrSlice

//export genStrSlice
func genStrSlice(n int) PGoStrSlice {
	var s []string
	for i := 0; i < n; i++ {
		s = append(s, "ユーザー" + strconv.Itoa(i))
	}

	p := PGoStrSlice(unsafe.Pointer(&s))
	if boxInstances == nil {
		boxInstances = make(map[PGoStrSlice]*GoStrSlice)
	}
	boxInstances[p] = &s
	return p
}

//export getGoStrSliceVal
func getGoStrSliceVal(p PGoStrSlice, i int) *C.char {
	slice := (*GoStrSlice)(unsafe.Pointer(p))
	return C.CString((*slice)[i])
}

//export getGoStrSliceLen
func getGoStrSliceLen(p PGoStrSlice) int {
	s := (*GoStrSlice)(unsafe.Pointer(p))
	return len(*s)
}

func main() {}
