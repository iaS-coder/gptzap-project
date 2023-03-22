package main

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/sashabaranov/go-openai"
)

func GenerateGPTText(query string) (string, error) {
	client := openai.NewClient("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")
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: query,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", err
	}

	return resp.Choices[0].Message.Content, nil

}

func parseBase64RequestData(r string) (string, error) {
	// encodedStr := base64.StdEncoding.EncodeToString([]byte(r))
	dataBytes, err := base64.URLEncoding.DecodeString(r)
	if err != nil {
		return "", err
	}

	data, err := url.ParseQuery(string(dataBytes))
	if data.Has("Body") {
		return data.Get("Body"), nil
	}

	return "", errors.New("body not found")
}

func process(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Printf(request.Body)
	result, err := parseBase64RequestData(request.Body)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	text, err := GenerateGPTText(result)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       text,
	}, nil

}

func main() {
	lambda.Start(process)
}
