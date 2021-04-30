type PGoStrSlice = pointer

proc genStrSlice():PGoStrSlice {.dynlib: "./strarr.so", importc: "genStrSlice".}
proc `[]`(p: PGoStrSlice, i:int):cstring {.dynlib: "./strarr.so", importc: "getGoStrSliceVal".}
proc len(p: PGoStrSlice):int {.dynlib: "./strarr.so", importc: "getGoStrSliceLen".}

proc `@`(p: PGoStrSlice):seq[cstring] =
  let sliceLen = p.len()
  var s = newSeq[cstring](sliceLen)
  for i in 0..<sliceLen:
    s[i] = p[i]
  return s

block:
  let sArr = @( genStrSlice() )
  echo sArr
  echo sArr[0]
