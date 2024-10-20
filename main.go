package main

import (
	"fmt"
	"homework/views"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(fmt.Errorf("%v", r))
		}
	}()

	views.Home()
}
