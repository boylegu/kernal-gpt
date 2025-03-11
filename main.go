package main

import (
	"kernal-gpt/cmd"
	"os"
)

func main() {

	cmd.RootCmd.AddCommand(cmd.RunGPTCommand())
	if err := cmd.RootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
