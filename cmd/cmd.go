package cmd

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
	"kernal-gpt/gpt"
	"kernal-gpt/utils"
	"log"
	"os"
	"time"
)

var (
	ollamaURL       string
	model           string
	redisURL        string
	compilerVersion = "0.3.0.dev"
)

var RootCmd = &cobra.Command{
	Use:   "kernal-gpt",
	Short: "The operating system assistant",
	Long:  `A tool that converts natural language commands into OS command-line utilities and kernel hooks for automated system operations.`,
}

func init() {
	RootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		banner := figure.NewFigure("kernal-gpt", "", true)
		cmd.Println(banner.String())
		cmd.Println(cmd.Long)
		cmd.Println(cmd.UsageString())

	})
}

func RunGPTCommand() *cobra.Command {
	var inputFile string
	runCmd := &cobra.Command{
		Use:   "run",
		Short: "Run the kernal-gpt tool",
		Long:  `Run the kernal-gpt tool with the specified input instruct.`,

		PreRun: func(cmd *cobra.Command, args []string) {
			var err error

			model, err = getParam(model, "KPT_MODEL", "Model")
			if err != nil {
				os.Exit(1)
			}

			ollamaURL, err = getParam(ollamaURL, "KPT_OLLAMA_URL", "Ollama URL")
			if err != nil {
				os.Exit(1)
			}

			ollamaURL, err = getParam(redisURL, "KPT_REDIS_URL", "Redis URL")
			if err != nil {
				os.Exit(1)
			}
		},

		Run: func(cmd *cobra.Command, args []string) {
			if inputFile == "" {
				log.Println("Error: --input flag is required")
				cmd.Help()
				os.Exit(1)
			}
			banner := figure.NewColorFigure("kernal-gpt", "", "purple", true)
			banner.Print()
			t := utils.TabbyNew()
			t.AddLine("Version:", compilerVersion)
			t.Print()

			done := make(chan bool)
			go spinner(100*time.Millisecond, done)
			resp := gpt.RunRagWorkflow(inputFile)
			done <- true
			log.Println(resp)
		},
	}
	runCmd.Flags().StringVarP(&inputFile, "input", "i", "", "input file (required)")
	runCmd.MarkFlagRequired("input")
	return runCmd
}

func getParam(inputValue, envVar, paramName string) (string, error) {
	inputValue = os.Getenv(envVar)
	if inputValue == "" {
		return "", fmt.Errorf("Error: %s is required. Please provide it via --%s flag or %s environment variable.",
			paramName, envVar, envVar)
	}
	return inputValue, nil
}
