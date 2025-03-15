package llms

import (
	"context"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/redisvector"
	"os"
)

var (
	baseUrlEnv  = "KPT_OLLAMA_URL"
	modelEnv    = "KPT_MODEL"
	redisUrlEnv = "KPT_REDIS_URL"
	vectorIndex = "kernal-vector"
)

func getLocalEnv(key string) string {
	return os.Getenv(key)
}

func CreateOpenAILLM() (*openai.LLM, error) {
	modelLLM, err := openai.New(
		openai.WithToken("ollama"),
		openai.WithBaseURL(getLocalEnv(baseUrlEnv)+"/v1"),
		openai.WithModel(getLocalEnv(modelEnv)),
	)
	if err != nil {
		return nil, err
	}
	return modelLLM, err
}

func CreateOllamaLLM() (*ollama.LLM, error) {
	modelLLM, err := ollama.New(
		ollama.WithServerURL(getLocalEnv(baseUrlEnv)),
		ollama.WithModel(getLocalEnv(modelEnv)),
	)
	if err != nil {
		return nil, err
	}
	return modelLLM, nil
}

func GetVectorStore(embedding *embeddings.EmbedderImpl) (vectorstores.VectorStore, error) {
	ctx := context.Background()
	store, err := redisvector.New(ctx,
		redisvector.WithConnectionURL(getLocalEnv(redisUrlEnv)),
		redisvector.WithIndexName(vectorIndex, true),
		redisvector.WithEmbedder(embedding),
	)
	if err != nil {
		return nil, err
	}
	return store, nil

}
