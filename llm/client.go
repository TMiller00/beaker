package llm

import "context"

// Client is the interface for any LLM provider
type Client interface {
	CreateMessage(ctx context.Context, model string, maxTokens int, prompt string) (string, error)
}
