import times

type PGoStrSlice = pointer

proc genStrSlice(n:int):PGoStrSlice {.dynlib: "./strarr.so", importc: "genStrSlice".}
proc `[]`(p: PGoStrSlice, i:int):cstring {.dynlib: "./strarr.so", importc: "getGoStrSliceVal".}
proc len(p: PGoStrSlice):int {.dynlib: "./strarr.so", importc: "getGoStrSliceLen".}

proc `@`(p: PGoStrSlice):seq[cstring] =
  let sliceLen = p.len()
  var s = newSeq[cstring](sliceLen)
  for i in 0..<sliceLen:
    s[i] = p[i]
  return s

when isMainModule:
  const length = 10000
  let start = cpuTime()
  let strSlice = @( genStrSlice(length) )
  echo cpuTime() - start
  echo strSlice.len
  echo strSlice[0]
  echo strSlice[length-1]
