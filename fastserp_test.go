// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomnimblewaynimblego_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Nimbleway/nimble-go"
	"github.com/Nimbleway/nimble-go/internal/testutil"
	"github.com/Nimbleway/nimble-go/option"
)

func TestFastSerpRunWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomnimblewaynimblego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.FastSerp.Run(context.TODO(), githubcomnimblewaynimblego.FastSerpRunParams{
		SearchEngine:      githubcomnimblewaynimblego.FastSerpRunParamsSearchEngineGoogleSearch,
		Country:           githubcomnimblewaynimblego.String("US"),
		Device:            githubcomnimblewaynimblego.FastSerpRunParamsDeviceDesktop,
		Domain:            githubcomnimblewaynimblego.String("com"),
		Locale:            githubcomnimblewaynimblego.String("en"),
		Location:          githubcomnimblewaynimblego.String("New York, New York, United States"),
		NumResults:        githubcomnimblewaynimblego.Int(10),
		Page:              githubcomnimblewaynimblego.Int(1),
		Parse:             githubcomnimblewaynimblego.Bool(true),
		Query:             githubcomnimblewaynimblego.String("nimble web data"),
		Render:            githubcomnimblewaynimblego.Bool(false),
		ResolveURL:        githubcomnimblewaynimblego.Bool(true),
		ShowHiddenResults: githubcomnimblewaynimblego.Bool(false),
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
