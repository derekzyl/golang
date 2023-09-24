package main

import "fmt"


  func test (a string, b string){
    fmt.Println("a: ", a, "b: ", b)
    temp := a
    a=b
    b=temp
    
    fmt.Println("a: ", a, "b: ", b)}


    func slic (product string, feat ...map[string]string){

if len(feat) !=0{

if len(feat) ==1{
  fmt.Println("product",product, "feat", feat )
}

if len(feat) >1{
  for k,v := range feat{
    fmt.Println("product", product, "[", "property", k, "value", v)
  }
}
}

    }

func main(){
  
test("hello", "world")
val := map[string]string {
  "ingredient" :"sugar",
  "name":"chocolate",
}

slic("cake", val)

}
