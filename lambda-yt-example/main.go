package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

type MyEvent struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

type Myresponse struct {
	Message string `json: "Answer"`
}

func HandleLambdaEvent(ctx context.Context, event MyEvent) (Myresponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return Myresponse{Message: fmt.Sprintf("%s %d", event.Name, event.Age)}, ctx.Err()
}

func main() {

	lambda.Start(HandleLambdaEvent)

}
