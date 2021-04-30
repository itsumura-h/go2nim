package main

import "C"
import (
	"fmt"
	"unsafe"
)

type User struct {
	ID   int
	Name string
}

type PUser = uintptr

// GCで削除されないように対策する
// uintptrでメモリとして保持する
// https://qiita.com/74th/items/0362bea2012ef253c539
var boxInstances map[PUser]*User

func main() {
	u := NewUser(1, "hoge")
	fmt.Println(u)
}

//export NewUser
func NewUser(id int, name string) PUser {
	u := User{ID: id, Name: name}
	p := PUser(unsafe.Pointer(&u))
	if boxInstances == nil {
		boxInstances = make(map[PUser]*User)
	}
	boxInstances[p] = &u
	return p
}

//export GetID
func GetID(p PUser) int {
	u := (*User)(unsafe.Pointer(p))
	return u.ID
}

//export GetName
func GetName(p PUser) string {
	u := (*User)(unsafe.Pointer(p))
	return u.Name
}

// exec comamnd (m1 macなのでarm64で指定してます。)
// 1. GOARCH=arm64 CGO_ENABLED=1  go build -buildmode=c-shared -o main.dll main.go
// 2. nim c --r -d:nimDebugDlOpen main.nim
