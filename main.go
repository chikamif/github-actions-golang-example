package main

import (
	"fmt"
	"github.com/chikamif/hello"
)

func main() {
	fmt.Println(say())
}

func say() string {
	return hello.Say()
}