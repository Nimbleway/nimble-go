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

func TestExtractWithOptionalParams(t *testing.T) {
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
	_, err := client.Extract(context.TODO(), githubcomnimblewaynimblego.ExtractParams{
		URL: "url",
		Browser: githubcomnimblewaynimblego.ExtractParamsBrowserUnion{
			OfExtractsBrowserString: githubcomnimblewaynimblego.String("chrome"),
		},
		BrowserActions: []githubcomnimblewaynimblego.ExtractParamsBrowserActionUnion{{
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
		Cookies: githubcomnimblewaynimblego.ExtractParamsCookiesUnion{
			OfString: githubcomnimblewaynimblego.String("sessionId=abc123; userId=user456"),
		},
		Country:             githubcomnimblewaynimblego.ExtractParamsCountryUs,
		Device:              githubcomnimblewaynimblego.ExtractParamsDeviceDesktop,
		Driver:              githubcomnimblewaynimblego.ExtractParamsDriverVx8,
		ExpectedStatusCodes: []int64{200, 201},
		Formats:             []string{"html"},
		Headers: map[string]githubcomnimblewaynimblego.ExtractParamsHeaderUnion{
			"Accept-Language": {
				OfString: githubcomnimblewaynimblego.String("en-US"),
			},
			"User-Agent": {
				OfString: githubcomnimblewaynimblego.String("CustomBot/1.0"),
			},
		},
		Http2:           githubcomnimblewaynimblego.Bool(true),
		IsXhr:           githubcomnimblewaynimblego.Bool(true),
		Locale:          githubcomnimblewaynimblego.ExtractParamsLocaleEnUs,
		MarkdownBackend: githubcomnimblewaynimblego.ExtractParamsMarkdownBackendFullPage,
		Method:          githubcomnimblewaynimblego.ExtractParamsMethodGet,
		NetworkCapture: []githubcomnimblewaynimblego.ExtractParamsNetworkCapture{{
			Method: "GET",
			ResourceType: githubcomnimblewaynimblego.ExtractParamsNetworkCaptureResourceTypeUnion{
				OfString: githubcomnimblewaynimblego.String("document"),
			},
			StatusCode: githubcomnimblewaynimblego.ExtractParamsNetworkCaptureStatusCodeUnion{
				OfFloat: githubcomnimblewaynimblego.Float(100),
			},
			StopOnRenderFlowFailure: githubcomnimblewaynimblego.Bool(true),
			URL: githubcomnimblewaynimblego.ExtractParamsNetworkCaptureURL{
				Value: "value",
				Type:  "exact",
			},
			Validation:                  githubcomnimblewaynimblego.Bool(true),
			WaitForRequestsCount:        githubcomnimblewaynimblego.Float(0),
			WaitForRequestsCountTimeout: githubcomnimblewaynimblego.Float(1),
		}},
		Os:    githubcomnimblewaynimblego.ExtractParamsOsWindows,
		Parse: githubcomnimblewaynimblego.Bool(true),
		Parser: githubcomnimblewaynimblego.ExtractParamsParserUnion{
			OfAnyMap: map[string]any{
				"myParser": "bar",
			},
		},
		ReferrerType:   githubcomnimblewaynimblego.ExtractParamsReferrerTypeRandom,
		Render:         githubcomnimblewaynimblego.Bool(true),
		RequestTimeout: githubcomnimblewaynimblego.Float(30000),
		Session: githubcomnimblewaynimblego.ExtractParamsSession{
			ID:                  githubcomnimblewaynimblego.String("id"),
			PrefetchUserbrowser: githubcomnimblewaynimblego.Bool(true),
			Retry:               githubcomnimblewaynimblego.Bool(true),
			Timeout:             githubcomnimblewaynimblego.Float(1),
		},
		Skill: githubcomnimblewaynimblego.ExtractParamsSkillUnion{
			OfString: githubcomnimblewaynimblego.String("dynamic-content"),
		},
		State: githubcomnimblewaynimblego.ExtractParamsStateCa,
		Tag:   githubcomnimblewaynimblego.String("campaign-2024-q1"),
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtractAsyncWithOptionalParams(t *testing.T) {
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
	_, err := client.ExtractAsync(context.TODO(), githubcomnimblewaynimblego.ExtractAsyncParams{
		URL: "url",
		Browser: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserUnion{
			OfExtractAsyncsBrowserString: githubcomnimblewaynimblego.String("chrome"),
		},
		BrowserActions: []githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionUnion{{
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
		CallbackURL:   githubcomnimblewaynimblego.String("https://example.com/webhook/callback"),
		City:          githubcomnimblewaynimblego.String("Los Angeles"),
		ConsentHeader: githubcomnimblewaynimblego.Bool(true),
		Cookies: githubcomnimblewaynimblego.ExtractAsyncParamsCookiesUnion{
			OfString: githubcomnimblewaynimblego.String("sessionId=abc123; userId=user456"),
		},
		Country:             githubcomnimblewaynimblego.ExtractAsyncParamsCountryUs,
		Device:              githubcomnimblewaynimblego.ExtractAsyncParamsDeviceDesktop,
		Driver:              githubcomnimblewaynimblego.ExtractAsyncParamsDriverVx8,
		ExpectedStatusCodes: []int64{200, 201},
		Formats:             []string{"html"},
		Headers: map[string]githubcomnimblewaynimblego.ExtractAsyncParamsHeaderUnion{
			"Accept-Language": {
				OfString: githubcomnimblewaynimblego.String("en-US"),
			},
			"User-Agent": {
				OfString: githubcomnimblewaynimblego.String("CustomBot/1.0"),
			},
		},
		Http2:           githubcomnimblewaynimblego.Bool(true),
		IsXhr:           githubcomnimblewaynimblego.Bool(true),
		Locale:          githubcomnimblewaynimblego.ExtractAsyncParamsLocaleEnUs,
		MarkdownBackend: githubcomnimblewaynimblego.ExtractAsyncParamsMarkdownBackendFullPage,
		Method:          githubcomnimblewaynimblego.ExtractAsyncParamsMethodGet,
		NetworkCapture: []githubcomnimblewaynimblego.ExtractAsyncParamsNetworkCapture{{
			Method: "GET",
			ResourceType: githubcomnimblewaynimblego.ExtractAsyncParamsNetworkCaptureResourceTypeUnion{
				OfString: githubcomnimblewaynimblego.String("document"),
			},
			StatusCode: githubcomnimblewaynimblego.ExtractAsyncParamsNetworkCaptureStatusCodeUnion{
				OfFloat: githubcomnimblewaynimblego.Float(100),
			},
			StopOnRenderFlowFailure: githubcomnimblewaynimblego.Bool(true),
			URL: githubcomnimblewaynimblego.ExtractAsyncParamsNetworkCaptureURL{
				Value: "value",
				Type:  "exact",
			},
			Validation:                  githubcomnimblewaynimblego.Bool(true),
			WaitForRequestsCount:        githubcomnimblewaynimblego.Float(0),
			WaitForRequestsCountTimeout: githubcomnimblewaynimblego.Float(1),
		}},
		Os:    githubcomnimblewaynimblego.ExtractAsyncParamsOsWindows,
		Parse: githubcomnimblewaynimblego.Bool(true),
		Parser: githubcomnimblewaynimblego.ExtractAsyncParamsParserUnion{
			OfAnyMap: map[string]any{
				"myParser": "bar",
			},
		},
		ReferrerType:   githubcomnimblewaynimblego.ExtractAsyncParamsReferrerTypeRandom,
		Render:         githubcomnimblewaynimblego.Bool(true),
		RequestTimeout: githubcomnimblewaynimblego.Float(30000),
		Session: githubcomnimblewaynimblego.ExtractAsyncParamsSession{
			ID:                  githubcomnimblewaynimblego.String("id"),
			PrefetchUserbrowser: githubcomnimblewaynimblego.Bool(true),
			Retry:               githubcomnimblewaynimblego.Bool(true),
			Timeout:             githubcomnimblewaynimblego.Float(1),
		},
		Skill: githubcomnimblewaynimblego.ExtractAsyncParamsSkillUnion{
			OfString: githubcomnimblewaynimblego.String("dynamic-content"),
		},
		State:             githubcomnimblewaynimblego.ExtractAsyncParamsStateCa,
		StorageCompress:   githubcomnimblewaynimblego.Bool(true),
		StorageObjectName: githubcomnimblewaynimblego.String("result-2024-01-15.json"),
		StorageType:       githubcomnimblewaynimblego.String("s3"),
		StorageURL:        githubcomnimblewaynimblego.String("s3://bucket-name/path/to/object"),
		Tag:               githubcomnimblewaynimblego.String("campaign-2024-q1"),
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtractBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.ExtractBatch(context.TODO(), githubcomnimblewaynimblego.ExtractBatchParams{
		Inputs: []githubcomnimblewaynimblego.ExtractBatchParamsInput{{
			Browser: githubcomnimblewaynimblego.ExtractBatchParamsInputBrowserUnion{
				OfExtractBatchsInputBrowserString: githubcomnimblewaynimblego.String("chrome"),
			},
			BrowserActions: []githubcomnimblewaynimblego.ExtractBatchParamsInputBrowserActionUnion{{
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
			CallbackURL:   githubcomnimblewaynimblego.String("https://example.com/webhook/callback"),
			City:          githubcomnimblewaynimblego.String("Los Angeles"),
			ConsentHeader: githubcomnimblewaynimblego.Bool(true),
			Cookies: githubcomnimblewaynimblego.ExtractBatchParamsInputCookiesUnion{
				OfString: githubcomnimblewaynimblego.String("sessionId=abc123; userId=user456"),
			},
			Country:             githubcomnimblewaynimblego.ExtractBatchParamsInputCountryUs,
			Device:              "desktop",
			Driver:              "vx8",
			ExpectedStatusCodes: []int64{200, 201},
			Formats:             []string{"html"},
			Headers: map[string]githubcomnimblewaynimblego.ExtractBatchParamsInputHeaderUnion{
				"Accept-Language": {
					OfString: githubcomnimblewaynimblego.String("en-US"),
				},
				"User-Agent": {
					OfString: githubcomnimblewaynimblego.String("CustomBot/1.0"),
				},
			},
			Http2:           githubcomnimblewaynimblego.Bool(true),
			IsXhr:           githubcomnimblewaynimblego.Bool(true),
			Locale:          githubcomnimblewaynimblego.ExtractBatchParamsInputLocaleEnUs,
			MarkdownBackend: "full_page",
			Method:          "GET",
			NetworkCapture: []githubcomnimblewaynimblego.ExtractBatchParamsInputNetworkCapture{{
				Method: "GET",
				ResourceType: githubcomnimblewaynimblego.ExtractBatchParamsInputNetworkCaptureResourceTypeUnion{
					OfString: githubcomnimblewaynimblego.String("document"),
				},
				StatusCode: githubcomnimblewaynimblego.ExtractBatchParamsInputNetworkCaptureStatusCodeUnion{
					OfFloat: githubcomnimblewaynimblego.Float(100),
				},
				StopOnRenderFlowFailure: githubcomnimblewaynimblego.Bool(true),
				URL: githubcomnimblewaynimblego.ExtractBatchParamsInputNetworkCaptureURL{
					Value: "value",
					Type:  "exact",
				},
				Validation:                  githubcomnimblewaynimblego.Bool(true),
				WaitForRequestsCount:        githubcomnimblewaynimblego.Float(0),
				WaitForRequestsCountTimeout: githubcomnimblewaynimblego.Float(1),
			}},
			Os:    "windows",
			Parse: githubcomnimblewaynimblego.Bool(true),
			Parser: githubcomnimblewaynimblego.ExtractBatchParamsInputParserUnion{
				OfAnyMap: map[string]any{
					"myParser": "bar",
				},
			},
			ReferrerType:   githubcomnimblewaynimblego.ExtractBatchParamsInputReferrerTypeRandom,
			Render:         githubcomnimblewaynimblego.Bool(true),
			RequestTimeout: githubcomnimblewaynimblego.Float(30000),
			Session: githubcomnimblewaynimblego.ExtractBatchParamsInputSession{
				ID:                  githubcomnimblewaynimblego.String("id"),
				PrefetchUserbrowser: githubcomnimblewaynimblego.Bool(true),
				Retry:               githubcomnimblewaynimblego.Bool(true),
				Timeout:             githubcomnimblewaynimblego.Float(1),
			},
			Skill: githubcomnimblewaynimblego.ExtractBatchParamsInputSkillUnion{
				OfString: githubcomnimblewaynimblego.String("dynamic-content"),
			},
			State:             "CA",
			StorageCompress:   githubcomnimblewaynimblego.Bool(true),
			StorageObjectName: githubcomnimblewaynimblego.String("result-2024-01-15.json"),
			StorageType:       githubcomnimblewaynimblego.String("s3"),
			StorageURL:        githubcomnimblewaynimblego.String("s3://bucket-name/path/to/object"),
			Tag:               githubcomnimblewaynimblego.String("campaign-2024-q1"),
			URL:               githubcomnimblewaynimblego.String("url"),
		}},
		SharedInputs: githubcomnimblewaynimblego.ExtractBatchParamsSharedInputs{
			Browser: githubcomnimblewaynimblego.ExtractBatchParamsSharedInputsBrowserUnion{
				OfExtractBatchsSharedInputsBrowserString: githubcomnimblewaynimblego.String("chrome"),
			},
			BrowserActions: []githubcomnimblewaynimblego.ExtractBatchParamsSharedInputsBrowserActionUnion{{
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
			CallbackURL:   githubcomnimblewaynimblego.String("https://example.com/webhook/callback"),
			City:          githubcomnimblewaynimblego.String("Los Angeles"),
			ConsentHeader: githubcomnimblewaynimblego.Bool(true),
			Cookies: githubcomnimblewaynimblego.ExtractBatchParamsSharedInputsCookiesUnion{
				OfString: githubcomnimblewaynimblego.String("sessionId=abc123; userId=user456"),
			},
			Country:             githubcomnimblewaynimblego.ExtractBatchParamsSharedInputsCountryUs,
			Device:              "desktop",
			Driver:              "vx8",
			ExpectedStatusCodes: []int64{200, 201},
			Formats:             []string{"html"},
			Headers: map[string]githubcomnimblewaynimblego.ExtractBatchParamsSharedInputsHeaderUnion{
				"Accept-Language": {
					OfString: githubcomnimblewaynimblego.String("en-US"),
				},
				"User-Agent": {
					OfString: githubcomnimblewaynimblego.String("CustomBot/1.0"),
				},
			},
			Http2:           githubcomnimblewaynimblego.Bool(true),
			IsXhr:           githubcomnimblewaynimblego.Bool(true),
			Locale:          githubcomnimblewaynimblego.ExtractBatchParamsSharedInputsLocaleEnUs,
			MarkdownBackend: "full_page",
			Method:          "GET",
			NetworkCapture: []githubcomnimblewaynimblego.ExtractBatchParamsSharedInputsNetworkCapture{{
				Method: "GET",
				ResourceType: githubcomnimblewaynimblego.ExtractBatchParamsSharedInputsNetworkCaptureResourceTypeUnion{
					OfString: githubcomnimblewaynimblego.String("document"),
				},
				StatusCode: githubcomnimblewaynimblego.ExtractBatchParamsSharedInputsNetworkCaptureStatusCodeUnion{
					OfFloat: githubcomnimblewaynimblego.Float(100),
				},
				StopOnRenderFlowFailure: githubcomnimblewaynimblego.Bool(true),
				URL: githubcomnimblewaynimblego.ExtractBatchParamsSharedInputsNetworkCaptureURL{
					Value: "value",
					Type:  "exact",
				},
				Validation:                  githubcomnimblewaynimblego.Bool(true),
				WaitForRequestsCount:        githubcomnimblewaynimblego.Float(0),
				WaitForRequestsCountTimeout: githubcomnimblewaynimblego.Float(1),
			}},
			Os:    "windows",
			Parse: githubcomnimblewaynimblego.Bool(true),
			Parser: githubcomnimblewaynimblego.ExtractBatchParamsSharedInputsParserUnion{
				OfAnyMap: map[string]any{
					"myParser": "bar",
				},
			},
			ReferrerType:   githubcomnimblewaynimblego.ExtractBatchParamsSharedInputsReferrerTypeRandom,
			Render:         githubcomnimblewaynimblego.Bool(true),
			RequestTimeout: githubcomnimblewaynimblego.Float(30000),
			Session: githubcomnimblewaynimblego.ExtractBatchParamsSharedInputsSession{
				ID:                  githubcomnimblewaynimblego.String("id"),
				PrefetchUserbrowser: githubcomnimblewaynimblego.Bool(true),
				Retry:               githubcomnimblewaynimblego.Bool(true),
				Timeout:             githubcomnimblewaynimblego.Float(1),
			},
			Skill: githubcomnimblewaynimblego.ExtractBatchParamsSharedInputsSkillUnion{
				OfString: githubcomnimblewaynimblego.String("dynamic-content"),
			},
			State:             "CA",
			StorageCompress:   githubcomnimblewaynimblego.Bool(true),
			StorageObjectName: githubcomnimblewaynimblego.String("result-2024-01-15.json"),
			StorageType:       githubcomnimblewaynimblego.String("s3"),
			StorageURL:        githubcomnimblewaynimblego.String("s3://bucket-name/path/to/object"),
			Tag:               githubcomnimblewaynimblego.String("campaign-2024-q1"),
			URL:               githubcomnimblewaynimblego.String("url"),
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
