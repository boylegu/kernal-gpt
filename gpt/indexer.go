package gpt

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores"
	"github.com/tmc/langchaingo/vectorstores/redisvector"
	"io"
	"log"
	"os"
	"strings"
)

func getExampleJson() *JSONData {
	filePath := "/home/admin/kernal-gpt/gpt/examples.json"
	file, err := os.Open(filePath)
	defer file.Close()
	bytes, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	var jsonData JSONData
	json.Unmarshal(bytes, &jsonData)

	return &jsonData
}

type ContentData struct {
	Content string `json:"content"`
}

type JSONData struct {
	Data []ContentData `json:"data"`
}

func Retriever(input string) string {
	redisURL := "redis://172.19.109.20:6379"
	index := "kernal-vector"

	_, e := getEmbedding("qwen2.5:1.5b", "http://10.55.1.57:11434")
	fmt.Println(3, e)
	ctx := context.Background()

	store, err := redisvector.New(ctx,
		redisvector.WithConnectionURL(redisURL),
		redisvector.WithIndexName(index, true),
		redisvector.WithEmbedder(e),
	)
	if err != nil {
		fmt.Println(err)
	}
	documents := make([]schema.Document, 0, len(getExampleJson().Data))

	for _, item := range getExampleJson().Data {
		doc := schema.Document{
			PageContent: item.Content,
			Score:       0.0,
		}
		documents = append(documents, doc)
	}

	_, err = store.AddDocuments(ctx, documents)
	docs, err := store.SimilaritySearch(ctx, input, 2,
		vectorstores.WithScoreThreshold(0.5),
	)

	var output strings.Builder
	for _, doc := range docs {
		output.WriteString(doc.PageContent)
		output.WriteString("\n")
	}

	//result, err := chains.Run(
	//	ctx,
	//	chains.NewRetrievalQAFromLLM(
	//		llm,
	//		vectorstores.ToRetriever(store, 5, vectorstores.WithScoreThreshold(0.8)),
	//	),
	//	input,
	//)
	return output.String()
}
