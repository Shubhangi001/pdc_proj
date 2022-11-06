package main

import (
	"os"

	"github.com/Shubhangi001/pdc_proj/cli"
)

func main() {
	defer os.Exit(0)
	cmd := cli.CommandLine{}
	cmd.Run()
}