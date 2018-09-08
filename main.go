package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

// Handler is your Lambda function handler
func Handler(c context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	// If no name is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}

	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("Hello %s", request.Body),
		StatusCode: http.StatusOK,
	}, nil

}

func main() {
	lambda.Start(Handler)
}
