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
	"github.com/Nimbleway/nimble-go/shared"
)

func TestCrawlListWithOptionalParams(t *testing.T) {
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
	_, err := client.Crawl.List(context.TODO(), githubcomnimblewaynimblego.CrawlListParams{
		Cursor: githubcomnimblewaynimblego.String("cursor"),
		Limit:  githubcomnimblewaynimblego.Int(10),
		Status: githubcomnimblewaynimblego.CrawlListParamsStatusQueued,
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCrawlRunWithOptionalParams(t *testing.T) {
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
	_, err := client.Crawl.Run(context.TODO(), githubcomnimblewaynimblego.CrawlRunParams{
		URL:                "url",
		AllowExternalLinks: githubcomnimblewaynimblego.Bool(false),
		AllowSubdomains:    githubcomnimblewaynimblego.Bool(false),
		Callback: githubcomnimblewaynimblego.CrawlRunParamsCallbackUnion{
			OfCrawlRunsCallbackObject: &githubcomnimblewaynimblego.CrawlRunParamsCallbackObject{
				URL:    "https://example.com",
				Events: []string{"started"},
				Headers: map[string]string{
					"foo": "string",
				},
				Metadata: map[string]any{
					"foo": "bar",
				},
			},
		},
		CrawlEntireDomain: githubcomnimblewaynimblego.Bool(false),
		ExcludePaths:      []string{"/exclude-this-path", "/and-this-path"},
		ExtractOptions: githubcomnimblewaynimblego.CrawlRunParamsExtractOptions{
			Body: map[string]any{
				"key": "value",
			},
			Browser: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserUnion{
				OfCrawlRunsExtractOptionsBrowserString: githubcomnimblewaynimblego.String("chrome"),
			},
			BrowserActions: []githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionUnion{{
				OfGotoAction: &shared.GotoActionParam{
					Goto: shared.GotoActionGotoUnionParam{
						OfString: githubcomnimblewaynimblego.String("https://example.com/login"),
					},
				},
			}, {
				OfWaitForElementAction: &shared.WaitForElementActionParam{
					WaitForElement: shared.WaitForElementActionWaitForElementUnionParam{
						OfString: githubcomnimblewaynimblego.String("#login-form"),
					},
				},
			}, {
				OfFillAction: &shared.FillActionParam{
					Fill: shared.FillActionFillUnionParam{
						OfType: &shared.FillActionFillTypeParam{
							Selector: shared.FillActionFillTypeSelectorUnionParam{
								OfString: githubcomnimblewaynimblego.String("#username"),
							},
							Value:          "user@example.com",
							ClickOnElement: githubcomnimblewaynimblego.Bool(true),
							Delay: shared.FillActionFillTypeDelayUnionParam{
								OfFloat: githubcomnimblewaynimblego.Float(1000),
							},
							Mode:                  "type",
							MouseMovementStrategy: "linear",
							Required: shared.FillActionFillTypeRequiredUnionParam{
								OfFillActionFillTypeRequiredString: githubcomnimblewaynimblego.Opt(shared.FillActionFillTypeRequiredStringTrue),
							},
							Scroll: githubcomnimblewaynimblego.Bool(true),
							Skip: shared.FillActionFillTypeSkipUnionParam{
								OfFillActionFillTypeSkipString: githubcomnimblewaynimblego.Opt(shared.FillActionFillTypeSkipStringTrue),
							},
							Timeout: githubcomnimblewaynimblego.Float(0),
							TypingInterval: shared.FillActionFillTypeTypingIntervalUnionParam{
								OfFloat: githubcomnimblewaynimblego.Float(1000),
							},
							TypingStrategy: "simple",
							Visible:        githubcomnimblewaynimblego.Bool(true),
						},
					},
				},
			}, {
				OfFillAction: &shared.FillActionParam{
					Fill: shared.FillActionFillUnionParam{
						OfType: &shared.FillActionFillTypeParam{
							Selector: shared.FillActionFillTypeSelectorUnionParam{
								OfString: githubcomnimblewaynimblego.String("#password"),
							},
							Value:          "password123",
							ClickOnElement: githubcomnimblewaynimblego.Bool(true),
							Delay: shared.FillActionFillTypeDelayUnionParam{
								OfFloat: githubcomnimblewaynimblego.Float(1000),
							},
							Mode:                  "type",
							MouseMovementStrategy: "linear",
							Required: shared.FillActionFillTypeRequiredUnionParam{
								OfFillActionFillTypeRequiredString: githubcomnimblewaynimblego.Opt(shared.FillActionFillTypeRequiredStringTrue),
							},
							Scroll: githubcomnimblewaynimblego.Bool(true),
							Skip: shared.FillActionFillTypeSkipUnionParam{
								OfFillActionFillTypeSkipString: githubcomnimblewaynimblego.Opt(shared.FillActionFillTypeSkipStringTrue),
							},
							Timeout: githubcomnimblewaynimblego.Float(0),
							TypingInterval: shared.FillActionFillTypeTypingIntervalUnionParam{
								OfFloat: githubcomnimblewaynimblego.Float(1000),
							},
							TypingStrategy: "simple",
							Visible:        githubcomnimblewaynimblego.Bool(true),
						},
					},
				},
			}, {
				OfClickAction: &shared.ClickActionParam{
					Click: shared.ClickActionClickUnionParam{
						OfString: githubcomnimblewaynimblego.String("#submit"),
					},
				},
			}, {
				OfScreenshotAction: &shared.ScreenshotActionParam{
					Screenshot: shared.ScreenshotActionScreenshotUnionParam{
						OfScreenshotActionScreenshotObject: &shared.ScreenshotActionScreenshotObjectParam{
							Format:   "png",
							FullPage: githubcomnimblewaynimblego.Bool(true),
							Quality:  githubcomnimblewaynimblego.Float(0),
							Required: shared.ScreenshotActionScreenshotObjectRequiredUnionParam{
								OfScreenshotActionScreenshotObjectRequiredString: githubcomnimblewaynimblego.Opt(shared.ScreenshotActionScreenshotObjectRequiredStringTrue),
							},
							Skip: shared.ScreenshotActionScreenshotObjectSkipUnionParam{
								OfScreenshotActionScreenshotObjectSkipString: githubcomnimblewaynimblego.Opt(shared.ScreenshotActionScreenshotObjectSkipStringTrue),
							},
						},
					},
				},
			}},
			City:          githubcomnimblewaynimblego.String("Los Angeles"),
			ConsentHeader: githubcomnimblewaynimblego.Bool(true),
			Cookies: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsCookiesUnion{
				OfString: githubcomnimblewaynimblego.String("sessionId=abc123; userId=user456"),
			},
			Country:             githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsCountryUs,
			Device:              "desktop",
			Driver:              "vx8",
			ExpectedStatusCodes: []int64{200, 201},
			Formats:             []string{"html"},
			Headers: map[string]githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsHeaderUnion{
				"Accept-Language": {
					OfString: githubcomnimblewaynimblego.String("en-US"),
				},
				"User-Agent": {
					OfString: githubcomnimblewaynimblego.String("CustomBot/1.0"),
				},
			},
			Http2:           githubcomnimblewaynimblego.Bool(true),
			IsXhr:           githubcomnimblewaynimblego.Bool(true),
			Locale:          githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsLocaleEnUs,
			MarkdownBackend: "full_page",
			Method:          "GET",
			NetworkCapture: []githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsNetworkCapture{{
				Method: "GET",
				ResourceType: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsNetworkCaptureResourceTypeUnion{
					OfString: githubcomnimblewaynimblego.String("document"),
				},
				StatusCode: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsNetworkCaptureStatusCodeUnion{
					OfFloat: githubcomnimblewaynimblego.Float(100),
				},
				URL: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsNetworkCaptureURL{
					Value: "value",
					Type:  "exact",
				},
				Validation:                  githubcomnimblewaynimblego.Bool(true),
				WaitForRequestsCount:        githubcomnimblewaynimblego.Float(0),
				WaitForRequestsCountTimeout: githubcomnimblewaynimblego.Float(1),
			}},
			Os:    "windows",
			Parse: githubcomnimblewaynimblego.Bool(true),
			Parser: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsParserUnion{
				OfAnyMap: map[string]any{
					"myParser": "bar",
				},
			},
			ReferrerType:   githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsReferrerTypeRandom,
			Render:         githubcomnimblewaynimblego.Bool(true),
			RequestTimeout: githubcomnimblewaynimblego.Float(30000),
			Session: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsSession{
				ID:                  githubcomnimblewaynimblego.String("id"),
				PrefetchUserbrowser: githubcomnimblewaynimblego.Bool(true),
				Retry:               githubcomnimblewaynimblego.Bool(true),
				Timeout:             githubcomnimblewaynimblego.Float(1),
			},
			Skill: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsSkillUnion{
				OfString: githubcomnimblewaynimblego.String("dynamic-content"),
			},
			State: "CA",
			Tag:   githubcomnimblewaynimblego.String("campaign-2024-q1"),
			URL:   githubcomnimblewaynimblego.String("url"),
		},
		IgnoreQueryParameters: githubcomnimblewaynimblego.Bool(false),
		IncludePaths:          []string{"/include-this-path", "/and-this-path"},
		Limit:                 githubcomnimblewaynimblego.Int(100),
		MaxDiscoveryDepth:     githubcomnimblewaynimblego.Int(3),
		Name:                  githubcomnimblewaynimblego.String("The best crawl ever"),
		Sitemap:               githubcomnimblewaynimblego.CrawlRunParamsSitemapInclude,
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCrawlStatus(t *testing.T) {
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
	_, err := client.Crawl.Status(context.TODO(), "123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCrawlTerminate(t *testing.T) {
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
	_, err := client.Crawl.Terminate(context.TODO(), "123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
