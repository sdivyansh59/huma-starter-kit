package routes

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/sdivyansh59/huma-project-starter/app/setup"
	"net/http"
)

// RegisterRoutes registers all greeting-related routes to the API
func RegisterRoutes(api *huma.API, c *setup.Controllers) {
	huma.Register(*api, huma.Operation{
		OperationID: "get-greeting",
		Method:      http.MethodGet,
		Path:        "/greeting/{name}",
		Summary:     "Get a greeting",
		Description: "Get a greeting for a person by name.",
		Tags:        []string{"Greetings"},
	}, c.Greeting.GetGreeting)

	huma.Register(*api, huma.Operation{
		OperationID:   "post-review",
		Method:        http.MethodPost,
		Path:          "/reviews",
		Summary:       "Post a review",
		Tags:          []string{"Reviews"},
		DefaultStatus: http.StatusCreated,
	}, c.Greeting.PostReview)
}
