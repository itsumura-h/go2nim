type PGoIntSlice = pointer

proc goFibo(n:int):PGoIntSlice {.dynlib: "./fib.so", importc: "fibo".}
proc `[]`(p: PGoIntSlice, i:int):int {.dynlib: "./fib.so", importc: "getGoIntSliceVal".}
proc len(p: PGoIntSlice):int {.dynlib: "./fib.so", importc: "getGoIntSliceLen".}

proc `@`(p: PGoIntSlice):seq[int] =
  let sliceLen = p.len()
  var s = newSeq[int](sliceLen)
  for i in 0..<sliceLen:
    s[i] = p[i]
  return s

proc fib(n:int):seq[int] =
  return @(goFibo(n))

echo fib(20)
