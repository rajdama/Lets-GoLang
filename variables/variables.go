package main

import "fmt"

func main() {
//one way to use variables (not recomended)

var var1 int
var var2 float64
var var3 bool
var var4 string

var1 = 2
var2 = 0.1
var3 = false
var4 = "hello"

fmt.Printf("%v %f %v %q\n",var1, var2, var3, var4 )

//second way to use variables (recommended)

var5 := 2
var6 := 0.1
var7 := false
var8 := "hello"

fmt.Printf("%v %f %v %q\n",var5, var6, var7, var8 )
}