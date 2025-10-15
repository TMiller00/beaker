package providers

import (
	"context"

	"github.com/anthropics/anthropic-sdk-go"
	"github.com/anthropics/anthropic-sdk-go/option"
)

type AnthropicClient struct {
	client anthropic.Client
}

func NewAnthropicClient(apiKey string) *AnthropicClient {
	return &AnthropicClient{
		client: anthropic.NewClient(option.WithAPIKey(apiKey)),
	}
}

func (a *AnthropicClient) CreateMessage(ctx context.Context, model string, maxTokens int, prompt string, systemPrompt ...string) (string, error) {
	params := anthropic.MessageNewParams{
		Model:     anthropic.Model(model),
		MaxTokens: int64(maxTokens),
		Messages: []anthropic.MessageParam{
			anthropic.NewUserMessage(anthropic.NewTextBlock(prompt)),
		},
	}

	if len(systemPrompt) > 0 && systemPrompt[0] != "" {
		params.System = []anthropic.TextBlockParam{{Text: systemPrompt[0]}}
	}

	message, err := a.client.Messages.New(ctx, params)
	if err != nil {
		return "", err
	}

	// Extract text content
	var content string
	for _, block := range message.Content {
		if block.Type == "text" {
			content += block.Text
		}
	}

	return content, nil
}
