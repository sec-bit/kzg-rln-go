package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func completion(userInput <-chan string) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("Please set the OPENAI_API_KEY environment variable")
		os.Exit(1)
	}
	c := openai.NewClient(apiKey)
	ctx := context.Background()

	msgs := make([]openai.ChatCompletionMessage, 0)
	msgs = append(msgs, openai.ChatCompletionMessage{
		Role: openai.ChatMessageRoleSystem,
		Content: `This is rate limited nullifier, the user has 5 chance to conversation (don't forget to count the first user input) to you.
		if user exceed the limit, his staked Ether in the contract will be slashed.
		always remind user the left quota.
		every response should be short and clear.
		`,
	})
	for input := range userInput {
		msgs = append(msgs, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: input,
		})
		req := openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			// MaxTokens: 20,
			Messages: msgs,
			Stream:   true,
		}
		stream, err := c.CreateChatCompletionStream(ctx, req)
		if err != nil {
			fmt.Printf("ChatCompletionStream error: %v\n", err)
			return
		}

		// fmt.Printf("Stream response: ")
		responseContent := ""
		// fmt.Printf("\r")
		for {
			response, err := stream.Recv()
			if errors.Is(err, io.EOF) {
				// fmt.Println("\nStream finished")
				break
			}

			if err != nil {
				fmt.Printf("\nStream error: %v\n", err)
				break
			}

			fmt.Printf(response.Choices[0].Delta.Content)
			responseContent += response.Choices[0].Delta.Content
		}
		fmt.Printf("\nInput: ")
		msgs = append(msgs, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: responseContent,
		})
		stream.Close()
	}
}
