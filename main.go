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

	service.GaussFilter(options.File, options.Raduis)
	fmt.Println("Data", options)
}
