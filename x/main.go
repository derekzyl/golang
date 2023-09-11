package main

import (
	"fmt"
	"sort"
)


func main()  {
  
  dic := map[string]interface{} {
    "name": "derek",
    "age":"30",
    "goals":"billionaire",
  }

  d := make([]string,0, len(dic))

for key, _ := range dic{
 d= append(d, key)
}

sort.Strings(d)


hj :=make(map[string]interface{})

for _, value := range d{
  hj[value] = dic[value]
}

fmt.Print(hj)

}
