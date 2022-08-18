package main

import (
	"errors"
	"time"

	readability "github.com/go-shiori/go-readability"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	ErrNoURL = errors.New("No URL provided")
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	url := request.QueryStringParameters["url"]

	if len(url) <= 0 {
		return events.APIGatewayProxyResponse{}, ErrNoURL
	}

	article, err := readability.FromURL(url, 30*time.Second)

	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

	html := "<!DOCTYPE html><html><head><meta charset='utf-8'/><title>GO READER</title><meta name='viewport' content='width=device-width, initial-scale=1'/><link rel='stylesheet' href='https://unpkg.com/bamboo.css'/></head><body><h2 id='title'>" + article.Title + "</h2><h3>" + article.Excerpt + "</h3><p id='output'>" + article.Content + "</p></body></html>"
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "origin,Accept,Authorization,Content-Type",
			"Content-Type":                 "text/html; charset=utf-8",
		},
		Body:            html,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	lambda.Start(handler)
}
