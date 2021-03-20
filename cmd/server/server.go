package main

import (
	"fmt"
	"os"

	srv "github.com/cardenasrjl/eth-stats/pkg/cmd/server"
)

func main() {
	if err := srv.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
