package gpt

import (
	"context"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/prompts"
	this "kernal-gpt/llms"
	graph "kernal-gpt/utils"
	"log"
)

var shouldDecideInstruct = func(ctx context.Context, state []llms.MessageContent) string {
	lastMsg := state[len(state)-1]

	for _, part := range lastMsg.Parts {
		return part.(llms.TextContent).Text
	}

	return graph.END
}

func RunRagWorkflow(input string) string {
	llm, err := this.CreateOllamaLLM()
	if err != nil {
		log.Fatal(err)
	}

	workflow := graph.NewMessageGraph()
	workflow.AddNode("quert_instruct", func(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error) {
		prompt := prompts.NewPromptTemplate(entityPrompt, []string{"user_input"})
		llmChain := chains.NewLLMChain(llm, prompt)

		out, err := chains.Run(ctx, llmChain, input)
		if err != nil {
			log.Fatal(err)
		}
		return append(state,
			llms.TextParts("rag_type", out+"_node"),
		), nil

	})

	workflow.AddNode("oscmd_node", func(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error) {
		prompt := prompts.NewPromptTemplate(osCmdPrompt, []string{"user_input"})
		llmChain := chains.NewLLMChain(llm, prompt)

		out, err := chains.Run(ctx, llmChain, input)
		if err != nil {
			log.Fatal(err)
		}
		return append(state,
			llms.TextParts("stdout", out),
		), nil

	})

	workflow.AddNode("ebpf_node", func(ctx context.Context, state []llms.MessageContent) ([]llms.MessageContent, error) {
		prompt := ConstructRunningPrompt(input)
		out := RunBpftrace(prompt)
		return append(state,
			llms.TextParts("stdout", out),
		), nil
	})

	workflow.SetEntryPoint("quert_instruct")
	workflow.AddConditionalEdge("quert_instruct", shouldDecideInstruct)
	workflow.AddEdge("oscmd_node", graph.END)
	workflow.AddEdge("ebpf_node", graph.END)

	app, err := workflow.Compile()
	if err != nil {
		log.Fatal("error: %v", err)
	}

	intialState := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, input),
	}

	response, err := app.Invoke(context.Background(), intialState)
	if err != nil {
		log.Fatal("error: %v", err)
	}

	lastMsg := response[len(response)-1]
	return lastMsg.Parts[0].(llms.TextContent).Text

}
