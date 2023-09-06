package main

import (
	"fmt"
	"math/rand"
        "strings"
        "time"
)

func main() {
lc:="abcdefghijklmnopqrstuvwxyz"
uc:="ABCDEFGHIJKLMNOPQRSTUVWXYZ"
lu:=lc+uc
numb:="1234567890"
numbuc:=numb+uc
numblc:=numb+lc
numblu:=numb+lu
sp:="!@#$%^&*()`~-=_+[];',./<>?:{}|"
spuc:=uc+sp
splc:=lc+sp
splu:=lu+sp
alluc:=uc+numb+sp
alllc:=lc+numb+sp
all:=lc+uc+numb+sp
numbsp:=numb+sp
var choicea int
var choiceb int
var l int
var loop int
var loopchoice int
var password strings.Builder

rand.Seed(time.Now().UnixNano())

loop=0
for loop==0{

fmt.Println("welcome to random password generator")
fmt.Println("1. Letters only \n2. Numbers only \n3. Letters with Numbers \n4. Letters and special chars \n5. Numbers and special chars \n6. Special chars only \n7. Use All")
fmt.Scanln(&choicea)

if choicea==1||choicea==3||choicea==4||choicea==7{
fmt.Println("1. Upper case only \n2. Lower case only \n3. Both")
fmt.Scanln(&choiceb)
if choicea==1 && choiceb==1{
fmt.Println("Length of password: ")
fmt.Scanln(&l)
for i:=1;i<=l;i++{
r1:=rand.Intn(len(uc))
password.WriteString(string(uc[r1]))}
fmt.Println("generated password: ",password.String())
} else if choicea==1 && choiceb==2{
fmt.Println("Length of password: ")
fmt.Scanln(&l)
for i:=1;i<=l;i++{
r1:=rand.Intn(len(lc))
password.WriteString(string(lc[r1]))}
fmt.Println("generated password: ",password.String())
} else if choicea==1 && choiceb==3{
fmt.Println("Length of password: ")
fmt.Scanln(&l)
for i:=1;i<=l;i++{
r1:=rand.Intn(len(lu))
password.WriteString(string(lu[r1]))}
fmt.Println("generated password: ",password.String())
} else if choicea==3 && choiceb==1{
fmt.Println("Length of password: ")
fmt.Scanln(&l)
for i:=1;i<=l;i++{
r1:=rand.Intn(len(numbuc))
password.WriteString(string(numbuc[r1]))}
fmt.Println("generated password: ",password.String())
} else if choicea==3 && choiceb==2{
fmt.Println("Length of password: ")
fmt.Scanln(&l)
for i:=1;i<=l;i++{
r1:=rand.Intn(len(numblc))
password.WriteString(string(numblc[r1]))}
fmt.Println("generated password: ",password.String())
} else if choicea==3 && choiceb==3{
fmt.Println("Length of password: ")
fmt.Scanln(&l)
for i:=1;i<=l;i++{
r1:=rand.Intn(len(numblu))
password.WriteString(string(numblu[r1]))}
fmt.Println("generated password: ",password.String())
} else if choicea==4 && choiceb==1{
fmt.Println("Length of password: ")
fmt.Scanln(&l)
for i:=1;i<=l;i++{
r1:=rand.Intn(len(spuc))
password.WriteString(string(spuc[r1]))}
fmt.Println("generated password: ",password.String())
} else if choicea==4 && choiceb==2{
fmt.Println("Length of password: ")
fmt.Scanln(&l)
for i:=1;i<=l;i++{
r1:=rand.Intn(len(splc))
password.WriteString(string(splc[r1]))}
fmt.Println("generated password: ",password.String())
} else if choicea==4 && choiceb==3{
fmt.Println("Length of password: ")
fmt.Scanln(&l)
for i:=1;i<=l;i++{
r1:=rand.Intn(len(splu))
password.WriteString(string(splu[r1]))}
fmt.Println("generated password: ",password.String())
} else if choicea==7 && choiceb==1{
fmt.Println("Length of password: ")
fmt.Scanln(&l)
for i:=1;i<=l;i++{
r1:=rand.Intn(len(alluc))
password.WriteString(string(alluc[r1]))}
fmt.Println("generated password: ",password.String())
} else if choicea==7 && choiceb==2{
fmt.Println("Length of password: ")
fmt.Scanln(&l)
for i:=1;i<=l;i++{
r1:=rand.Intn(len(alllc))
password.WriteString(string(alllc[r1]))}
fmt.Println("generated password: ",password.String())
} else if choicea==7 && choiceb==3{
fmt.Println("Length of password: ")
fmt.Scanln(&l)
for i:=1;i<=l;i++{
r1:=rand.Intn(len(all))
password.WriteString(string(all[r1]))}
fmt.Println("generated password: ",password.String())
}
} else if choicea==2 {
fmt.Println("Length of password: ")
fmt.Scanln(&l)
for i:=1;i<=l;i++{
r1:=rand.Intn(len(numb))
password.WriteString(string(numb[r1]))}
fmt.Println("generated password: ",password.String())
} else if choicea==5 {
fmt.Println("Length of password: ")
fmt.Scanln(&l)
for i:=1;i<=l;i++{
r1:=rand.Intn(len(numbsp))
password.WriteString(string(numbsp[r1]))}
fmt.Println("generated password: ",password.String())
} else if choicea==6 {
fmt.Println("Length of password: ")
fmt.Scanln(&l)
for i:=1;i<=l;i++{
r1:=rand.Intn(len(sp))
password.WriteString(string(sp[r1]))}
fmt.Println("generated password: ",password.String())
}
fmt.Println("generate another? (enter 1 for yes): ")
fmt.Scanln(&loopchoice)
if loopchoice==1{
loop=0
} else {
loop=loop+1
}
}
}