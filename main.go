package main

import (
	"fmt"
	"os"

	"github.com/hasanbakirci/api-observability-demo/cmd"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [api | consumer]")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "api":
		cmd.RunApi()
	case "consumer":
		cmd.RunConsumer()
	default:
		fmt.Println("Invalid argument. Use 'api' or 'consumer'")
		os.Exit(1)
	}
}
