package gemini

import (
	"context"
	"fmt"
	"os"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

const (
	modelCode = "models/gemini-2.0-flash-lite"
)

type GeminiApiClient struct {
	generateClient func() (*genai.Client, error)
}

func NewGeminiApiClient(ctx context.Context) (*GeminiApiClient, error) {
	apiKey, exists := os.LookupEnv("GEMINI_API_KEY")

	if !exists {
		return nil, fmt.Errorf("GEMINI_API_KEY is unset")
	}
	generateClient := func() (*genai.Client, error) {
		client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
		if err != nil {
			return nil, err
		}
		return client, nil
	}

	return &GeminiApiClient{generateClient: generateClient}, nil
}

func (c *GeminiApiClient) GenerateText(ctx context.Context, prompt string) ([]*genai.Candidate, error) {
	if len(prompt) == 0 {
		return nil, fmt.Errorf("prompt should not be empty")
	}

	client, err := c.generateClient()
	if err != nil {
		return nil, fmt.Errorf("failed to create gemini client", err)
	}
	defer client.Close()

	model := client.GenerativeModel(modelCode)
	text := genai.Text(prompt)
	res, err := model.GenerateContent(ctx, text)
	if err != nil {
		return nil, fmt.Errorf("failed to call gemini api", err)
	}

	// TODO: return merged string
	printResponse(res)
	return res.Candidates, nil
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content == nil {
			continue
		}
		for _, part := range cand.Content.Parts {
			fmt.Println(part)
		}
	}
}
