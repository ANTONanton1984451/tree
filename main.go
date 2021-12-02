package main

import (
	"command"
	"fmt"
)

func main() {
	c := command.New()
	err := c.Exec()
	if err != nil {
		fmt.Printf("err")
	}
}
