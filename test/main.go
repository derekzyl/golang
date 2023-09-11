package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)


func main()  {
  fmt.Print("hello world")
  app := gin.Default()


  app.GET("/", func(ctx *gin.Context){
    ctx.JSON(200, gin.H {"data":"hello"})
  })
  app.Run() 
}
