type
  User = pointer

proc newUser(id:int, name:cstring):User {.dynlib: "./struct.so", importc: "NewUser".}
proc id(p:User):int {.dynlib: "./struct.so", importc: "GetID".}
proc name(p:User):cstring {.dynlib: "./struct.so", importc: "GetName".}

when isMainModule:
  var user = newUser(1, "abc")
  echo user.id()
  echo user.name()

  user = newUser(2, "あいう")
  echo user.id()
  echo user.name()
