https://github.com/sivchari/go2nim

[How to return a slice in Go and calling from C?](https://stackoverflow.com/questions/52309383/how-to-return-a-slice-in-go-and-calling-from-c)

```sh
go build -buildmode=c-shared -o fib.so fib.go
nim c -r fib
```
```sh
go build -buildmode=c-shared -o struct.so struct.go
nim c -r struct
```
