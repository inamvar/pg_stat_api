package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"pg_stats/database"
	"pg_stats/handlers"
	"testing"
)

func TestStatRoute(t *testing.T) {


	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "get without any query params",
			route:        "/",
			expectedCode: 200,
		},
		// Second test case
		{
			description:  "get HTTP status 404, when route is not exists",
			route:        "/not-found",
			expectedCode: 404,
		},
		// third test case
		{
			description:  "get HTTP status 400, when  query param filter value is not valid",
			route:        "/?filter=invalid-value",
			expectedCode: 400,
		},
	}

	//create mock connection
	_ = database.CreateMockConnection()
	// Define Fiber app.
	app := fiber.New()
	app.Get("/", handlers.StatsHandler)

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest("GET", test.route, nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
	}

}
