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

func TestExtractTemplateGenerationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Extract.Templates.Generations.New(context.TODO(), githubcomnimblewaynimblego.ExtractTemplateGenerationNewParams{
		OfCreateExtractTemplateGenerationRequestPublicV2: &githubcomnimblewaynimblego.ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateGenerationRequestPublicV2{
			Prompt: "prompt",
			URL:    "url",
			InputSchema: map[string]any{
				"foo": "bar",
			},
			Metadata: githubcomnimblewaynimblego.ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateGenerationRequestPublicV2Metadata{
				Description: githubcomnimblewaynimblego.String("description"),
				DisplayName: githubcomnimblewaynimblego.String("display_name"),
				Tags:        []string{"string"},
			},
			Name: githubcomnimblewaynimblego.String("name"),
			OutputSchema: map[string]any{
				"foo": "bar",
			},
		},
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtractTemplateGenerationGet(t *testing.T) {
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
	_, err := client.Extract.Templates.Generations.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
