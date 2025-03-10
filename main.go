package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"kernal-gpt/gpt"
	"kernal-gpt/utils"
	"os"
	"time"
)

const compilerVersion string = "0.0.1"

var isDone bool

func spinner(delay time.Duration, done chan bool) {
	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	frames := []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}
	i := 0
	for {
		select {
		case <-done:
			fmt.Print("\r")
			return
		default:
			if isDone {
				fmt.Print("\r")
				done <- true
				return
			}
			fmt.Printf("\r%s Thinking...", frames[i])
			i = (i + 1) % len(frames)
			time.Sleep(delay)
		}
	}
}

var rootCmd = &cobra.Command{
	Use:   "kernal-gpt",
	Short: "The operating system assistant",
	Long:  `A tool that converts natural language commands into OS command-line utilities and kernel hooks for automated system operations.`,
}

func init() {
	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		banner := figure.NewFigure("kernal-gpt", "", true)
		cmd.Println(banner.String())
		cmd.Println(cmd.Long)
		cmd.Println(cmd.UsageString())

	})
}

func main() {

	rootCmd.AddCommand(runGPTCommand())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func runGPTCommand() *cobra.Command {
	var inputFile string
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run the kernal-gpt tool",
		Long:  `Run the kernal-gpt tool with the specified input instruct.`,
		Run: func(cmd *cobra.Command, args []string) {
			if inputFile == "" {
				fmt.Println("Error: --input flag is required")
				cmd.Help()
				os.Exit(1)
			}
			banner := figure.NewColorFigure("kernal-gpt", "", "purple", true)
			banner.Print()
			t := utils.TabbyNew()
			t.AddLine("Version:", compilerVersion)
			t.Print()
			prompt := gpt.ConstructRunningPrompt(inputFile)
			gpt.RunBpftrace(prompt)
		},
	}
	runCmd.Flags().StringVarP(&inputFile, "input", "i", "", "input file (required)")
	runCmd.MarkFlagRequired("input")
	return runCmd
}
