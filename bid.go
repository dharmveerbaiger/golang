package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
)

func abc(c *gin.Context) {
	c.JSON(200, gin.H{
		"price": rand.Intn(100),
	})
}
func main() {
	r := gin.Default()
	r.GET("/bid", abc)

	r.Run(":8081")
	//os.Args[1], listen and serve on port from command line
}
