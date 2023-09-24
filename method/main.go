package main

import "fmt"



func main() {
  dt:= dCon(6.0, 6.6,5.8) 
  fmt.Println("area", dt.area(), "perimeter", dt.perimeter())
}
