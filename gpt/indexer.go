package gpt

import (
	"context"
	_ "embed"
	"encoding/json"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores"
	"io"
	this "kernal-gpt/llms"
	"log"
	"os"
	"strings"
)

func getExampleJson() *JSONData {
	filePath := "./gpt/examples.json"
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

	_, e := getEmbedding()
	ctx := context.Background()

	store, err := this.GetVectorStore(e)
	if err != nil {
		log.Fatal(err)
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
	return output.String()
}
