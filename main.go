package main

import (
	"fmt"
	"record-back/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
	err := r.Run(fmt.Sprintf(""))
	if err != nil {
		fmt.Printf("Run server failed, err: %s\n", err)
		return
	}
}
