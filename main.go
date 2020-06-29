package main

import (
	"gotest/router"
)

func main()  {
	r := router.RouterInit()

	r.Run(":8081")
}
