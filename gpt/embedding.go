package gpt

import (
	"fmt"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/ollama"
	"log"
)

func getEmbedding(model string, connectionStr ...string) (llms.Model, *embeddings.EmbedderImpl) {
	opts := []ollama.Option{ollama.WithModel(model)}
	if len(connectionStr) > 0 {
		fmt.Println(4211, connectionStr[0])
		opts = append(opts, ollama.WithServerURL(connectionStr[0]))
	}

	llm, err := ollama.New(opts...)
	if err != nil {
		fmt.Println(1, err)
		log.Fatal(err)
	}

	e, err := embeddings.NewEmbedder(llm)
	if err != nil {
		fmt.Println(2, err)
		log.Fatal(err)
	}
	return llms.Model(llm), e
}
