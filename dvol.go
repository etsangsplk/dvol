package main

import (
	"fmt"
	"os"

	"github.com/ClusterHQ/dvol/pkg/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}
}
