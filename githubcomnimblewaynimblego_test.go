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

func TestMapWithOptionalParams(t *testing.T) {
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
	_, err := client.Map(context.TODO(), githubcomnimblewaynimblego.MapParams{
		URL:          "url",
		Country:      githubcomnimblewaynimblego.MapParamsCountryUs,
		DomainFilter: githubcomnimblewaynimblego.MapParamsDomainFilterAll,
		Limit:        githubcomnimblewaynimblego.Int(1000),
		Locale:       githubcomnimblewaynimblego.MapParamsLocaleEnUs,
		Sitemap:      githubcomnimblewaynimblego.MapParamsSitemapInclude,
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Search(context.TODO(), githubcomnimblewaynimblego.SearchParams{
		Query:          "x",
		ContentType:    []string{"string"},
		Country:        githubcomnimblewaynimblego.String("country"),
		DeepSearch:     githubcomnimblewaynimblego.Bool(true),
		EndDate:        githubcomnimblewaynimblego.String("end_date"),
		ExcludeDomains: []string{"string"},
		Focus: githubcomnimblewaynimblego.SearchParamsFocusUnion{
			OfString: githubcomnimblewaynimblego.String("string"),
		},
		IncludeAnswer:  githubcomnimblewaynimblego.Bool(true),
		IncludeDomains: []string{"string"},
		Locale:         githubcomnimblewaynimblego.String("locale"),
		MaxResults:     githubcomnimblewaynimblego.Int(1),
		MaxSubagents:   githubcomnimblewaynimblego.Int(1),
		OutputFormat:   githubcomnimblewaynimblego.SearchParamsOutputFormatPlainText,
		SearchDepth:    githubcomnimblewaynimblego.SearchParamsSearchDepthLite,
		StartDate:      githubcomnimblewaynimblego.String("start_date"),
		TimeRange:      githubcomnimblewaynimblego.SearchParamsTimeRangeHour,
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
