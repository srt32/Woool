package main

import (
	"fmt"
	"os"

	"github.com/openai/openai-go"
)

// processReceipt calls the OpenAI API with the uploaded file content and extracts the purchase summary and amount.
func processReceipt(fileContent []byte) (string, string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	prompt := fmt.Sprintf("Please summarize this purchase and the dollar amount: %s", string(fileContent))

	resp, err := client.Completions.Create(openai.CompletionRequest{
		Model:     "text-davinci-003",
		Prompt:    prompt,
		MaxTokens: 100,
	})
	if err != nil {
		return "", "", err
	}

	// Assuming the response contains the summary and amount in the first text output
	// This part may need adjustment based on actual OpenAI response format
	text := resp.Choices[0].Text
	// Further parsing of text to extract summary and amount can be added here

	return "Purchase summary", "$100", nil // Placeholder return values
}
