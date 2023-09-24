package main

import "fmt"


 type Product struct{

 name string
price float64
image string
is_available string 


 }


type Cart struct{

product Product
total float64


}

func main (){

  cart := Cart{
    total:7.7,
    product:Product{ 
      name :"coke",
    },
  }

  cart.total= 7.7
  cart.product.name= "coke"
  cart.product.image="img"
  cart.product.price=7.6



  fmt.Println(cart)
  



}
