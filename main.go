package main

import (
	"fmt"
	"os"

	"github.com/samar2170/portfolio-manager-v4/api"
	mutualfund "github.com/samar2170/portfolio-manager-v4/security/mutual-fund"
)

func main() {
	arg := os.Args[1]
	fmt.Println(arg)
	switch arg {
	case "setup":
		fmt.Println("setting up")
		setup()
	default:
		fmt.Println("starting")
		api.StartServer()
	}
}

func setup() {
	// stock.LoadData()
	mutualfund.LoadData()
}
