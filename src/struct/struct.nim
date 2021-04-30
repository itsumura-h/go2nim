type
  User = pointer

proc newUser(id:int, name:cstring):User {.dynlib: "./struct.so", importc: "NewUser".}
proc id(p:User):int {.dynlib: "./struct.so", importc: "GetID".}
proc name(p:User):cstring {.dynlib: "./struct.so", importc: "GetName".}

let user = newUser(2, "abc")
echo user.id()
echo user.name()
