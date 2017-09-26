package main

import (
	"os"
	"fmt"
	client "github.com/andrepinto/goway-sidecar/cmd/agent/app"
)

func main() {
	app := client.NewCliApp()
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}