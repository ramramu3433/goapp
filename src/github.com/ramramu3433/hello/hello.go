package main

import "fmt"

func pow(x int) int{
var y int
y = x*x
return y
}

func main(){

fmt.Printf("Hello world \t Google developed this lang \n");
fmt.Println("Here i am writing some basic math ")
var i,j int
i=10;
j=20;
fmt.Printf("The value of x+y is %d \n",i+j)
fmt.Printf("The Square of %d is %d \n" ,i,pow(i))

}



