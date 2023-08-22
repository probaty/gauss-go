package main

import (
	"fmt"
	"gauss-go/service"
)

func main() {
	options, err := service.GetArgs()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Data", options)
}
