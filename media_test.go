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

func TestMediaRunWithOptionalParams(t *testing.T) {
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
	_, err := client.Media.Run(context.TODO(), githubcomnimblewaynimblego.MediaRunParams{
		URL:               "https://example.com",
		Country:           githubcomnimblewaynimblego.String("country"),
		ExpectedMimeTypes: []string{"string"},
		Locale:            githubcomnimblewaynimblego.String("locale"),
		Storage: githubcomnimblewaynimblego.MediaRunParamsStorage{
			URL:        "url",
			ObjectName: githubcomnimblewaynimblego.String("object_name"),
			Type:       "s3",
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

func TestMediaRunAsyncWithOptionalParams(t *testing.T) {
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
	_, err := client.Media.RunAsync(context.TODO(), githubcomnimblewaynimblego.MediaRunAsyncParams{
		URL:               "https://example.com",
		CallbackURL:       githubcomnimblewaynimblego.String("https://example.com/webhook/callback"),
		Country:           githubcomnimblewaynimblego.String("country"),
		ExpectedMimeTypes: []string{"string"},
		Locale:            githubcomnimblewaynimblego.String("locale"),
		Storage: githubcomnimblewaynimblego.MediaRunAsyncParamsStorage{
			URL:        "url",
			ObjectName: githubcomnimblewaynimblego.String("object_name"),
			Type:       "s3",
		},
		StorageCompress:   githubcomnimblewaynimblego.Bool(true),
		StorageObjectName: githubcomnimblewaynimblego.String("result-2024-01-15.json"),
		StorageType:       githubcomnimblewaynimblego.String("s3"),
		StorageURL:        githubcomnimblewaynimblego.String("s3://bucket-name/path/to/object"),
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
