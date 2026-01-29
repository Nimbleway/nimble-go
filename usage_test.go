// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nimbleway_test

import (
	"context"
	"os"
	"testing"

	"github.com/stainless-sdks/nimbleway-go"
	"github.com/stainless-sdks/nimbleway-go/internal/testutil"
	"github.com/stainless-sdks/nimbleway-go/option"
)

func TestUsage(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := nimbleway.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	t.Skip("Prism tests are disabled")
	response, err := client.Extract(context.TODO(), nimbleway.ExtractParams{
		DebugOptions: nimbleway.ExtractParamsDebugOptions{},
		URL:          "https://example.com",
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", response.ID)
}
