package gpt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
	"log"
	"os"
)

func RunBpftrace(prompt string) {

	opts := []openai.Option{openai.WithModel("qwen2.5:1.5b")}
	opts = append(opts, openai.WithBaseURL("http://10.55.1.57:11434/v1"))
	opts = append(opts, openai.WithToken("ollama"))

	llm, err := openai.New(opts...)
	if err != nil {
		fmt.Println(1, err)
	}
	ctx := context.Background()

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeHuman, prompt),
	}

	resp, err := llm.GenerateContent(ctx, content, llms.WithTools(availableTools))

	if err != nil {
		fmt.Println(2, err)
	}

	executeToolCalls(ctx, llm, content, resp)
}

func constructCommand(args string) []string {
	var (
		operation map[string]interface{}
		cmd       []string
	)

	if err := json.Unmarshal([]byte(args), &operation); err != nil {
		log.Fatal("Error parsing JSON:", err)
	}
	if _, exists := operation["bufferingMode"]; exists {
		cmd = append(cmd, []string{"-B", operation["bufferingMode"].(string)}...)
	}

	if _, exists := operation["format"]; exists {
		cmd = append(cmd, []string{"-f", operation["format"].(string)}...)
	}

	if _, exists := operation["format"]; exists {
		cmd = append(cmd, []string{"-f", operation["format"].(string)}...)
	}

	if _, exists := operation["outputFile"]; exists {
		cmd = append(cmd, []string{"-o", operation["outputFile"].(string)}...)
	}

	if _, exists := operation["debugInfo"]; exists {
		cmd = append(cmd, []string{"-d"}...)
	}

	if _, exists := operation["verboseDebugInfo"]; exists {
		cmd = append(cmd, []string{"-dd"}...)
	}

	if _, exists := operation["program"]; exists {
		cmd = append(cmd, []string{"-e", operation["program"].(string)}...)
	}

	if _, exists := operation["includeDir"]; exists {
		for _, dir := range operation["includeDir"].([]interface{}) {
			cmd = append(cmd, []string{"-I", dir.(string)}...)
		}
	}

	if _, exists := operation["usdtFileActivation"]; exists {
		cmd = append(cmd, []string{"--usdt-file-activation"}...)
	}

	if _, exists := operation["unsafe"]; exists {
		cmd = append(cmd, []string{"--unsafe"}...)
	}

	if _, exists := operation["quiet"]; exists {
		cmd = append(cmd, []string{"-q"}...)
	}

	if _, exists := operation["verbose"]; exists {
		cmd = append(cmd, []string{"-v"}...)
	}

	if _, exists := operation["noWarnings"]; exists {
		cmd = append(cmd, []string{"--no-warnings"}...)
	}

	return cmd
}

func executeToolCalls(ctx context.Context, llm llms.Model, messageHistory []llms.MessageContent, resp *llms.ContentResponse) []llms.MessageContent {
	fmt.Println("Executing", len(resp.Choices[0].ToolCalls), "tool calls")
	for _, toolCall := range resp.Choices[0].ToolCalls {
		full_command := []string{"sudo"}
		switch toolCall.FunctionCall.Name {
		case "bpftrace":
			full_command = append(full_command, "bpftrace")
			args := toolCall.FunctionCall.Arguments
			constructCommand(args)
		case "SaveFile":
			var data map[string]interface{}
			args := toolCall.FunctionCall.Arguments
			if err := json.Unmarshal([]byte(args), &data); err != nil {
				log.Fatal("Error parsing JSON:", err)
			}
			filename := data["filename"].(string)
			fmt.Println("Save to file: " + filename)
			err := os.WriteFile(filename, data["content"].([]byte), 0644)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	return messageHistory
}

var availableTools = []llms.Tool{
	{
		Type: "function",
		Function: &llms.FunctionDefinition{
			Name:        "bpftrace",
			Description: "A tool to run bpftrace eBPF programs",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"bufferingMode": map[string]any{
						"type":        "string",
						"description": "output buffering mode",
					},
					"format": map[string]any{
						"type":        "string",
						"description": "output format",
					},
					"outputFile": map[string]any{
						"type":        "string",
						"description": "redirect bpftrace output to file",
					},
					"debugInfo": map[string]any{
						"type":        "boolean",
						"description": "debug info dry run",
					},
					"verboseDebugInfo": map[string]any{
						"type":        "boolean",
						"description": "verbose debug info dry run",
					},
					"program": map[string]any{
						"type":        "string",
						"description": "program to execute",
					},
					"includeDir": map[string]any{
						"type": "array",
						"items": map[string]any{
							"type": "string",
						},
						"description": "directories to add to the include search path",
					},
					"usdtFileActivation": map[string]any{
						"type":        "boolean",
						"description": "activate usdt semaphores based on file path",
					},
					"unsafe": map[string]any{
						"type":        "boolean",
						"description": "allow unsafe builtin functions",
					},
					"quiet": map[string]any{
						"type":        "boolean",
						"description": "keep messages quiet",
					},
					"verbose": map[string]any{
						"type":        "boolean",
						"description": "verbose messages",
					},
					"noWarnings": map[string]any{
						"type":        "boolean",
						"description": "disable all warning messages",
					},
					"timeout": map[string]any{
						"type":        "integer",
						"description": "seconds to run the command",
					},
					"continue": map[string]any{
						"type":        "boolean",
						"description": "finish conversation and not continue.",
					},
				},
				"required": []string{"program"},
			},
		},
	},
	{
		Type: "function",
		Function: &llms.FunctionDefinition{
			Name:        "SaveFile",
			Description: "Save the eBPF program to file",
			Parameters: map[string]any{
				"type": "object",
				"properties": map[string]any{
					"filename": map[string]any{
						"type":        "string",
						"description": "the file name to save to",
					},
					"content": map[string]any{
						"type":        "string",
						"description": "the file content",
					},
				},
				"required": []string{"filename", "content"},
			},
		},
	},
}
