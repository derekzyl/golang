package main

import (
	"fmt"
	"strconv"
)



func main (){


  var v string  = "$344"

  var dol string = string(v[0])
  srb := []string{ "hello", "world", "15"}

  data, err := strconv.ParseFloat(srb[2], 64)
if err != nil{
  fmt.Print(err)
}else{
  fmt.Printf(dol,data)
}
}
