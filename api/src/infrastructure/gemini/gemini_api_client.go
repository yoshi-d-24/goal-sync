package gemini

import (
	"context"
	"fmt"
	"os"

	"encoding/json"

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

func (c *GeminiApiClient) GenerateText(ctx context.Context, prompt string) (string, error) {
	if len(prompt) == 0 {
		return "", fmt.Errorf("prompt should not be empty")
	}

	client, err := c.generateClient()
	if err != nil {
		return "", fmt.Errorf("failed to create gemini client. err=%s", err)
	}
	defer client.Close()

	model := client.GenerativeModel(modelCode)
	res, err := model.GenerateContent(ctx, genai.Text(prompt))
	if err != nil {
		return "", fmt.Errorf("failed to call gemini api. err=%s", err)
	}

	bs, err := json.Marshal(res.Candidates[0].Content.Parts[0])
	if err != nil {
		return "", fmt.Errorf("failed to parse gemini api result. err=%s", err)
	}
	result := string(bs)
	return result, nil
}
