type
  GoSlice = pointer

proc fibo(n:int):GoSlice {.dynlib: "./fib.so", importc: "fibo".}
proc `[]`(p: GoSlice, i:int):int {.dynlib: "./fib.so", importc: "getSliceIntVal".}
proc len(p: GoSlice):int {.dynlib: "./fib.so", importc: "getSliceLen".}

proc `@`(p: GoSlice):seq[int] =
  let sliceLen = p.len()
  var s = newSeq[int](sliceLen)
  for i in 0..<sliceLen:
    s[i] = p[i]
  return s

let p = fibo(10)
echo @p
