package main

import (
	"fmt"
	"os"

	"github.com/samar2170/portfolio-manager-v4/api"
)

func main() {
	arg := os.Args[1]
	fmt.Println(arg)
	switch arg {
	case "setup":
		fmt.Println("setting up")
	default:
		fmt.Println("starting")
		api.StartServer()

	}
}
