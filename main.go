package main

import (
   "C"
   "fmt"
)

//export helloWorld
func helloWorld(){
   fmt.Println("Hello World")
}

//export hello
func hello(namePtr *C.char){
   name := C.GoString(namePtr)
   fmt.Println("Hello", name)
}


func main(){

}
