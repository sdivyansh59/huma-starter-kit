package greeting

import (
	"context"
	"fmt"
)

func GetGreeting(ctx context.Context, input *GetGreetingInput) (*GetGreetingOutput, error) {
	resp := &GetGreetingOutput{}
	resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
	return resp, nil
}

func PostReview(ctx context.Context, input *PostReviewInput) (*struct{}, error) {
	return nil, nil // TODO: save review in data store.
}
