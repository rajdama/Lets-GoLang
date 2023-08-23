package main

import "fmt"

func main() {
//one way to use variables 

var var1 int
var var2 float64
var var3 bool
var var4 string

var1 = 2
var2 = 0.1
var3 = false
var4 = "hello"

fmt.Printf("%v %f %v %q\n",var1, var2, var3, var4 ) // Using Printf needs format specifying

//second way to use variables (preffered)

var5 := 2
var6 := 0.1
var7 := false
var8 := "hello"

fmt.Printf("%v %f %v %q\n",var5, var6, var7, var8 )

//declaring two variables simultaneously

var9, var10 := .2, "hello"
fmt.Println(var9, var10) // Using Println does not needs format specifying

//typecasting

var11 := 5.7
var12 := int(var11) //Float to Int
fmt.Println(var12) 

//constants
const var13 = 1
const var14 = 2
const var15 = var13 + var14 // will be computed during compiletime and not runtime unlike in javascript
fmt.Println(var15) 


}