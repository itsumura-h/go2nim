https://github.com/sivchari/go2nim

[How to return a slice in Go and calling from C?](https://stackoverflow.com/questions/52309383/how-to-return-a-slice-in-go-and-calling-from-c)

[cgo - Go references to C](https://golang.org/cmd/cgo/#hdr-Go_references_to_C)

[GoとCの間のポインタ渡し](https://qiita.com/74th/items/0362bea2012ef253c539)

```sh
go build -buildmode=c-shared -o fib.so fib.go
nim c -r fib
```

```sh
go build -buildmode=c-shared -o strarr.so strarr.go
nim c -r strarr
```

```sh
go build -buildmode=c-shared -o struct.so struct.go
nim c -r struct
```

## Nim → Go → Nim
### int
```
int(Nim) → int(Go) → int(Nim)
```

### string
```
cstring(Nim) → string(Go) → cstring(Nim)
```

## Go → Nim
### int
```
int(Go) → int(Nim)
```

### string
```
string(Go) → *C.char(Go) → cstring(Nim)
```
