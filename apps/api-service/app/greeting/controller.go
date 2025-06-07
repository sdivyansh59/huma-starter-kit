package greeting

import (
	"context"
	"fmt"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{
		// initialize logger
	}
}

func (c *Controller) GetGreeting(ctx context.Context, input *GetGreetingInput) (*GetGreetingOutput, error) {
	resp := &GetGreetingOutput{}
	resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
	return resp, nil
}

func (c *Controller) PostReview(ctx context.Context, input *PostReviewInput) (*struct{}, error) {
	return nil, nil // TODO: save review in data store.
}
