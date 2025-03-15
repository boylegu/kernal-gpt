package gpt

import (
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms"
	this "kernal-gpt/llms"
	"log"
)

func getEmbedding() (llms.Model, *embeddings.EmbedderImpl) {
	llm, err := this.CreateOllamaLLM()
	if err != nil {
		log.Fatal(err)
	}

	e, err := embeddings.NewEmbedder(llm)
	if err != nil {
		log.Fatal(err)
	}
	return llms.Model(llm), e
}
