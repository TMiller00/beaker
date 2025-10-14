package main

import (
	"context"
	"fmt"
	"os"

	"github.com/TMiller00/beaker/providers"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading environment variables: %v\n", err)
	}

	anthropicApiKey := os.Getenv("ANTHROPIC_API_KEY")
	var client providers.Client = providers.NewAnthropicClient(anthropicApiKey)

	response, err := client.CreateMessage(
		context.TODO(),
		"claude-3-5-haiku-latest",
		1024,
		"When did the Tchaikovsky's 'Romeo And Juliet' fantasy overture become popular?",
	)

	if err != nil {
		fmt.Printf("Error creating new message: %v\n", err)
		return
	}

	fmt.Printf("%s\n", response)
}
