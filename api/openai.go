package api

import (
	"context"
	"log"

	"github.com/tmc/langchaingo/llms"
	"github.com/tmc/langchaingo/llms/openai"
)

func OpenAI(input string) (string, error) {
	llm, err := openai.New()
	if err != nil {
		return "", err
	}

	// Call the OpenAI API with the input and specified options.
	completion, err := llm.Call(
		context.Background(),
		input,
		llms.WithTemperature(0.8),
		llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
			// Print streaming chunks to the console.
			log.Print(string(chunk))
			return nil
		}),
	)
	if err != nil {
		return "", err
	}

	return completion, nil
}
