package main

import  (
	"fmt"
  "github.com/gin-gonic/gin"
)


func orderSetUp(c *gin.Context) {
  restID := c.PostForm("RestID")
  token := c.PostForm("token")

  fmt.Println(restID)
  fmt.Println(token)
}
