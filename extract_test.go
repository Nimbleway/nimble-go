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

func TestExtractAsyncWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Extract.Async(context.TODO(), githubcomnimblewaynimblego.ExtractAsyncParams{
		URL: "url",
		Browser: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserUnion{
			OfExtractAsyncsBrowserString: githubcomnimblewaynimblego.String("chrome"),
		},
		BrowserActions: []githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionUnion{{
			OfExtractAsyncsBrowserActionGotoAction: &githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionGotoAction{
				Goto: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionGotoActionGotoUnion{
					OfString: githubcomnimblewaynimblego.String("https://example.com/login"),
				},
			},
		}, {
			OfExtractAsyncsBrowserActionWaitForElementAction: &githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionWaitForElementAction{
				WaitForElement: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementUnion{
					OfString: githubcomnimblewaynimblego.String("#login-form"),
				},
			},
		}, {
			OfExtractAsyncsBrowserActionFillAction: &githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillAction{
				Fill: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillUnion{
					OfType: &githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillType{
						Selector: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillTypeSelectorUnion{
							OfString: githubcomnimblewaynimblego.String("#username"),
						},
						Value:          "user@example.com",
						ClickOnElement: githubcomnimblewaynimblego.Bool(true),
						Delay: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillTypeDelayUnion{
							OfFloat: githubcomnimblewaynimblego.Float(1000),
						},
						Mode:                  "type",
						MouseMovementStrategy: "linear",
						Required: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillTypeRequiredUnion{
							OfExtractAsyncsBrowserActionFillActionFillTypeRequiredString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillTypeRequiredStringTrue),
						},
						Scroll: githubcomnimblewaynimblego.Bool(true),
						Skip: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillTypeSkipUnion{
							OfExtractAsyncsBrowserActionFillActionFillTypeSkipString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillTypeSkipStringTrue),
						},
						Timeout: githubcomnimblewaynimblego.Float(0),
						TypingInterval: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillTypeTypingIntervalUnion{
							OfFloat: githubcomnimblewaynimblego.Float(1000),
						},
						TypingStrategy: "simple",
						Visible:        githubcomnimblewaynimblego.Bool(true),
					},
				},
			},
		}, {
			OfExtractAsyncsBrowserActionFillAction: &githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillAction{
				Fill: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillUnion{
					OfType: &githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillType{
						Selector: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillTypeSelectorUnion{
							OfString: githubcomnimblewaynimblego.String("#password"),
						},
						Value:          "password123",
						ClickOnElement: githubcomnimblewaynimblego.Bool(true),
						Delay: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillTypeDelayUnion{
							OfFloat: githubcomnimblewaynimblego.Float(1000),
						},
						Mode:                  "type",
						MouseMovementStrategy: "linear",
						Required: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillTypeRequiredUnion{
							OfExtractAsyncsBrowserActionFillActionFillTypeRequiredString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillTypeRequiredStringTrue),
						},
						Scroll: githubcomnimblewaynimblego.Bool(true),
						Skip: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillTypeSkipUnion{
							OfExtractAsyncsBrowserActionFillActionFillTypeSkipString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillTypeSkipStringTrue),
						},
						Timeout: githubcomnimblewaynimblego.Float(0),
						TypingInterval: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionFillActionFillTypeTypingIntervalUnion{
							OfFloat: githubcomnimblewaynimblego.Float(1000),
						},
						TypingStrategy: "simple",
						Visible:        githubcomnimblewaynimblego.Bool(true),
					},
				},
			},
		}, {
			OfExtractAsyncsBrowserActionClickAction: &githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionClickAction{
				Click: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionClickActionClickUnion{
					OfString: githubcomnimblewaynimblego.String("#submit"),
				},
			},
		}, {
			OfExtractAsyncsBrowserActionScreenshotAction: &githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionScreenshotAction{
				Screenshot: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionScreenshotActionScreenshotUnion{
					OfExtractAsyncsBrowserActionScreenshotActionScreenshotObject: &githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObject{
						Format:   "png",
						FullPage: githubcomnimblewaynimblego.Bool(true),
						Quality:  githubcomnimblewaynimblego.Float(0),
						Required: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion{
							OfExtractAsyncsBrowserActionScreenshotActionScreenshotObjectRequiredString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectRequiredStringTrue),
						},
						Skip: githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion{
							OfExtractAsyncsBrowserActionScreenshotActionScreenshotObjectSkipString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectSkipStringTrue),
						},
					},
				},
			},
		}},
		City:          githubcomnimblewaynimblego.String("Los Angeles"),
		ClientTimeout: githubcomnimblewaynimblego.Float(25000),
		ConsentHeader: githubcomnimblewaynimblego.Bool(true),
		Cookies: githubcomnimblewaynimblego.ExtractAsyncParamsCookiesUnion{
			OfExtractAsyncsCookiesArray: []githubcomnimblewaynimblego.ExtractAsyncParamsCookiesArrayItem{{
				Creation:     githubcomnimblewaynimblego.String("creation"),
				Domain:       githubcomnimblewaynimblego.String("domain"),
				Expires:      githubcomnimblewaynimblego.String("expires"),
				Extensions:   []string{"string"},
				HostOnly:     githubcomnimblewaynimblego.Bool(true),
				HTTPOnly:     githubcomnimblewaynimblego.Bool(true),
				LastAccessed: githubcomnimblewaynimblego.String("lastAccessed"),
				MaxAge: githubcomnimblewaynimblego.ExtractAsyncParamsCookiesArrayItemMaxAgeUnion{
					OfExtractAsyncsCookiesArrayItemMaxAgeString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractAsyncParamsCookiesArrayItemMaxAgeStringInfinity),
				},
				Name:          githubcomnimblewaynimblego.String("name"),
				Path:          githubcomnimblewaynimblego.String("path"),
				PathIsDefault: githubcomnimblewaynimblego.Bool(true),
				SameSite:      "strict",
				Secure:        githubcomnimblewaynimblego.Bool(true),
				Value:         githubcomnimblewaynimblego.String("value"),
			}},
		},
		Country:             githubcomnimblewaynimblego.ExtractAsyncParamsCountryUs,
		Device:              githubcomnimblewaynimblego.ExtractAsyncParamsDeviceDesktop,
		DisableIPCheck:      githubcomnimblewaynimblego.Bool(false),
		Driver:              githubcomnimblewaynimblego.ExtractAsyncParamsDriverVx8,
		ExpectedStatusCodes: []int64{200, 201},
		Formats:             []string{"html"},
		Headers: map[string]githubcomnimblewaynimblego.ExtractAsyncParamsHeaderUnion{
			"User-Agent": {
				OfString: githubcomnimblewaynimblego.String("CustomBot/1.0"),
			},
			"Accept-Language": {
				OfString: githubcomnimblewaynimblego.String("en-US"),
			},
		},
		Http2:  githubcomnimblewaynimblego.Bool(true),
		Ip6:    githubcomnimblewaynimblego.Bool(false),
		IsXhr:  githubcomnimblewaynimblego.Bool(true),
		Locale: githubcomnimblewaynimblego.ExtractAsyncParamsLocaleEnUs,
		Metadata: githubcomnimblewaynimblego.ExtractAsyncParamsMetadata{
			AccountName:         githubcomnimblewaynimblego.String("account_name"),
			APIType:             githubcomnimblewaynimblego.String("api_type"),
			CrawlDepth:          githubcomnimblewaynimblego.Int(-9007199254740991),
			CrawlID:             githubcomnimblewaynimblego.String("crawl_id"),
			DefinitionID:        githubcomnimblewaynimblego.Int(-9007199254740991),
			DefinitionName:      githubcomnimblewaynimblego.String("definition_name"),
			Endpoint:            githubcomnimblewaynimblego.String("endpoint"),
			ExecutionID:         githubcomnimblewaynimblego.String("execution_id"),
			FlowitTaskID:        githubcomnimblewaynimblego.String("flowit_task_id"),
			InputID:             githubcomnimblewaynimblego.String("input_id"),
			IsPublicWsa:         githubcomnimblewaynimblego.Bool(true),
			IsSitemap:           githubcomnimblewaynimblego.Bool(true),
			IsWsa:               githubcomnimblewaynimblego.Bool(true),
			ParserID:            githubcomnimblewaynimblego.String("parser_id"),
			PipelineExecutionID: githubcomnimblewaynimblego.Int(-9007199254740991),
			QueryTemplateID:     githubcomnimblewaynimblego.String("query_template_id"),
			Source:              githubcomnimblewaynimblego.String("source"),
			TemplateID:          githubcomnimblewaynimblego.Int(-9007199254740991),
			TemplateName:        githubcomnimblewaynimblego.String("template_name"),
			WsaID:               githubcomnimblewaynimblego.String("wsa_id"),
			WsaName:             githubcomnimblewaynimblego.String("wsa_name"),
			WsaVersion:          githubcomnimblewaynimblego.Float(0),
		},
		Method:     githubcomnimblewaynimblego.ExtractAsyncParamsMethodGet,
		NativeMode: githubcomnimblewaynimblego.ExtractAsyncParamsNativeModeRequester,
		NetworkCapture: []githubcomnimblewaynimblego.ExtractAsyncParamsNetworkCapture{{
			Method: "GET",
			ResourceType: githubcomnimblewaynimblego.ExtractAsyncParamsNetworkCaptureResourceTypeUnion{
				OfString: githubcomnimblewaynimblego.String("document"),
			},
			StatusCode: githubcomnimblewaynimblego.ExtractAsyncParamsNetworkCaptureStatusCodeUnion{
				OfFloat: githubcomnimblewaynimblego.Float(100),
			},
			URL: githubcomnimblewaynimblego.ExtractAsyncParamsNetworkCaptureURL{
				Value: "value",
				Type:  "exact",
			},
			Validation:                  githubcomnimblewaynimblego.Bool(true),
			WaitForRequestsCount:        githubcomnimblewaynimblego.Float(0),
			WaitForRequestsCountTimeout: githubcomnimblewaynimblego.Float(1),
		}},
		NoUserbrowser: githubcomnimblewaynimblego.Bool(false),
		Os:            githubcomnimblewaynimblego.ExtractAsyncParamsOsWindows,
		Parse:         githubcomnimblewaynimblego.Bool(true),
		Parser: githubcomnimblewaynimblego.ExtractAsyncParamsParserUnion{
			OfAnyMap: map[string]any{
				"myParser": "bar",
			},
		},
		ProxyProvider: githubcomnimblewaynimblego.ExtractAsyncParamsProxyProviderBrightdata,
		ProxyProviders: map[string]float64{
			"brightdata": 70,
			"oxylabs":    30,
		},
		QueryTemplate: githubcomnimblewaynimblego.ExtractAsyncParamsQueryTemplate{
			ID:      "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			APIType: "WEB",
			Pagination: githubcomnimblewaynimblego.ExtractAsyncParamsQueryTemplatePaginationUnion{
				OfExtractAsyncsQueryTemplatePaginationNextPageParams: &githubcomnimblewaynimblego.ExtractAsyncParamsQueryTemplatePaginationNextPageParams{
					NextPageParams: map[string]any{
						"foo": "bar",
					},
				},
			},
			Params: map[string]any{
				"foo": "bar",
			},
		},
		RawHeaders:   githubcomnimblewaynimblego.Bool(true),
		ReferrerType: githubcomnimblewaynimblego.ExtractAsyncParamsReferrerTypeRandom,
		Render:       githubcomnimblewaynimblego.Bool(true),
		RenderFlow: []map[string]any{{
			"wait": "bar",
		}, {
			"click": "bar",
		}},
		RenderOptions: githubcomnimblewaynimblego.ExtractAsyncParamsRenderOptions{
			Adblock:        githubcomnimblewaynimblego.Bool(true),
			BlockedDomains: []string{"ads.example.com", "tracker.com"},
			BrowserEngine: githubcomnimblewaynimblego.ExtractAsyncParamsRenderOptionsBrowserEngineUnion{
				OfExtractAsyncsRenderOptionsBrowserEngineString: githubcomnimblewaynimblego.String("chrome"),
			},
			Cache:             githubcomnimblewaynimblego.Bool(false),
			ConnectorType:     "puppeteer",
			DisabledResources: []string{"image", "stylesheet"},
			Enable2captcha:    githubcomnimblewaynimblego.Bool(true),
			Extensions:        []string{"extension-id-1", "extension-id-2"},
			FingerprintID:     githubcomnimblewaynimblego.String("fp-abc123"),
			HackiumConfiguration: githubcomnimblewaynimblego.ExtractAsyncParamsRenderOptionsHackiumConfiguration{
				CollectLogs:                 githubcomnimblewaynimblego.Bool(true),
				DoNotFixMathSalt:            githubcomnimblewaynimblego.Bool(true),
				EnableDocumentElementSpoof:  githubcomnimblewaynimblego.Bool(true),
				EnableDocumentHasFocus:      githubcomnimblewaynimblego.Bool(true),
				EnableFakeNavigationHistory: githubcomnimblewaynimblego.Bool(true),
				EnableKeyOrdering:           githubcomnimblewaynimblego.Bool(true),
				EnableSniffer:               githubcomnimblewaynimblego.Bool(true),
				EnableVerboseLogs:           githubcomnimblewaynimblego.Bool(true),
			},
			Headless:               githubcomnimblewaynimblego.Bool(true),
			IncludeIframes:         githubcomnimblewaynimblego.Bool(true),
			LoadLocalStorage:       githubcomnimblewaynimblego.Bool(true),
			LocalStorageKeysToLoad: []string{"authToken", "userId"},
			MouseStrategy:          "linear",
			NoAcceptEncoding:       githubcomnimblewaynimblego.Bool(true),
			OverridePermissions:    githubcomnimblewaynimblego.Bool(true),
			RandomHeaderOrder:      githubcomnimblewaynimblego.Bool(true),
			RenderType:             "domcontentloaded",
			StoreLocalStorage:      githubcomnimblewaynimblego.Bool(true),
			Timeout:                githubcomnimblewaynimblego.Float(30000),
			TypingInterval:         githubcomnimblewaynimblego.Float(100),
			TypingStrategy:         "simple",
			Userbrowser:            githubcomnimblewaynimblego.Bool(true),
			WaitUntil:              "networkidle2",
			WithPerformanceMetrics: githubcomnimblewaynimblego.Bool(true),
		},
		RequestTimeout:  githubcomnimblewaynimblego.Float(30000),
		SaveUserbrowser: githubcomnimblewaynimblego.Bool(false),
		Session: githubcomnimblewaynimblego.ExtractAsyncParamsSession{
			ID:                  githubcomnimblewaynimblego.String("id"),
			PrefetchUserbrowser: githubcomnimblewaynimblego.Bool(true),
			Retry:               githubcomnimblewaynimblego.Bool(true),
			Timeout:             githubcomnimblewaynimblego.Float(1),
		},
		Skill: githubcomnimblewaynimblego.ExtractAsyncParamsSkillUnion{
			OfString: githubcomnimblewaynimblego.String("dynamic-content"),
		},
		SkipUbct: githubcomnimblewaynimblego.Bool(false),
		State:    githubcomnimblewaynimblego.ExtractAsyncParamsStateCa,
		Tag:      githubcomnimblewaynimblego.String("campaign-2024-q1"),
		Template: githubcomnimblewaynimblego.ExtractAsyncParamsTemplate{
			Name: "x",
			Params: map[string]any{
				"foo": "bar",
			},
		},
		Type: githubcomnimblewaynimblego.String("generic"),
		UserbrowserCreationTemplateRendered: githubcomnimblewaynimblego.ExtractAsyncParamsUserbrowserCreationTemplateRendered{
			ID:                    "id",
			AllowedParameterNames: []string{"x"},
			RenderFlowRendered: []map[string]any{{
				"foo": "bar",
			}},
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

func TestExtractExtractWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Extract.Extract(context.TODO(), githubcomnimblewaynimblego.ExtractExtractParams{
		URL: "url",
		Browser: githubcomnimblewaynimblego.ExtractExtractParamsBrowserUnion{
			OfExtractExtractsBrowserString: githubcomnimblewaynimblego.String("chrome"),
		},
		BrowserActions: []githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionUnion{{
			OfExtractExtractsBrowserActionGotoAction: &githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionGotoAction{
				Goto: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionGotoActionGotoUnion{
					OfString: githubcomnimblewaynimblego.String("https://example.com/login"),
				},
			},
		}, {
			OfExtractExtractsBrowserActionWaitForElementAction: &githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionWaitForElementAction{
				WaitForElement: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementUnion{
					OfString: githubcomnimblewaynimblego.String("#login-form"),
				},
			},
		}, {
			OfExtractExtractsBrowserActionFillAction: &githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillAction{
				Fill: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillUnion{
					OfType: &githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillType{
						Selector: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillTypeSelectorUnion{
							OfString: githubcomnimblewaynimblego.String("#username"),
						},
						Value:          "user@example.com",
						ClickOnElement: githubcomnimblewaynimblego.Bool(true),
						Delay: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillTypeDelayUnion{
							OfFloat: githubcomnimblewaynimblego.Float(1000),
						},
						Mode:                  "type",
						MouseMovementStrategy: "linear",
						Required: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillTypeRequiredUnion{
							OfExtractExtractsBrowserActionFillActionFillTypeRequiredString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillTypeRequiredStringTrue),
						},
						Scroll: githubcomnimblewaynimblego.Bool(true),
						Skip: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillTypeSkipUnion{
							OfExtractExtractsBrowserActionFillActionFillTypeSkipString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillTypeSkipStringTrue),
						},
						Timeout: githubcomnimblewaynimblego.Float(0),
						TypingInterval: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillTypeTypingIntervalUnion{
							OfFloat: githubcomnimblewaynimblego.Float(1000),
						},
						TypingStrategy: "simple",
						Visible:        githubcomnimblewaynimblego.Bool(true),
					},
				},
			},
		}, {
			OfExtractExtractsBrowserActionFillAction: &githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillAction{
				Fill: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillUnion{
					OfType: &githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillType{
						Selector: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillTypeSelectorUnion{
							OfString: githubcomnimblewaynimblego.String("#password"),
						},
						Value:          "password123",
						ClickOnElement: githubcomnimblewaynimblego.Bool(true),
						Delay: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillTypeDelayUnion{
							OfFloat: githubcomnimblewaynimblego.Float(1000),
						},
						Mode:                  "type",
						MouseMovementStrategy: "linear",
						Required: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillTypeRequiredUnion{
							OfExtractExtractsBrowserActionFillActionFillTypeRequiredString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillTypeRequiredStringTrue),
						},
						Scroll: githubcomnimblewaynimblego.Bool(true),
						Skip: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillTypeSkipUnion{
							OfExtractExtractsBrowserActionFillActionFillTypeSkipString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillTypeSkipStringTrue),
						},
						Timeout: githubcomnimblewaynimblego.Float(0),
						TypingInterval: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionFillActionFillTypeTypingIntervalUnion{
							OfFloat: githubcomnimblewaynimblego.Float(1000),
						},
						TypingStrategy: "simple",
						Visible:        githubcomnimblewaynimblego.Bool(true),
					},
				},
			},
		}, {
			OfExtractExtractsBrowserActionClickAction: &githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionClickAction{
				Click: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionClickActionClickUnion{
					OfString: githubcomnimblewaynimblego.String("#submit"),
				},
			},
		}, {
			OfExtractExtractsBrowserActionScreenshotAction: &githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionScreenshotAction{
				Screenshot: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionScreenshotActionScreenshotUnion{
					OfExtractExtractsBrowserActionScreenshotActionScreenshotObject: &githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionScreenshotActionScreenshotObject{
						Format:   "png",
						FullPage: githubcomnimblewaynimblego.Bool(true),
						Quality:  githubcomnimblewaynimblego.Float(0),
						Required: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion{
							OfExtractExtractsBrowserActionScreenshotActionScreenshotObjectRequiredString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredStringTrue),
						},
						Skip: githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion{
							OfExtractExtractsBrowserActionScreenshotActionScreenshotObjectSkipString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipStringTrue),
						},
					},
				},
			},
		}},
		City:          githubcomnimblewaynimblego.String("Los Angeles"),
		ClientTimeout: githubcomnimblewaynimblego.Float(25000),
		ConsentHeader: githubcomnimblewaynimblego.Bool(true),
		Cookies: githubcomnimblewaynimblego.ExtractExtractParamsCookiesUnion{
			OfExtractExtractsCookiesArray: []githubcomnimblewaynimblego.ExtractExtractParamsCookiesArrayItem{{
				Creation:     githubcomnimblewaynimblego.String("creation"),
				Domain:       githubcomnimblewaynimblego.String("domain"),
				Expires:      githubcomnimblewaynimblego.String("expires"),
				Extensions:   []string{"string"},
				HostOnly:     githubcomnimblewaynimblego.Bool(true),
				HTTPOnly:     githubcomnimblewaynimblego.Bool(true),
				LastAccessed: githubcomnimblewaynimblego.String("lastAccessed"),
				MaxAge: githubcomnimblewaynimblego.ExtractExtractParamsCookiesArrayItemMaxAgeUnion{
					OfExtractExtractsCookiesArrayItemMaxAgeString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractExtractParamsCookiesArrayItemMaxAgeStringInfinity),
				},
				Name:          githubcomnimblewaynimblego.String("name"),
				Path:          githubcomnimblewaynimblego.String("path"),
				PathIsDefault: githubcomnimblewaynimblego.Bool(true),
				SameSite:      "strict",
				Secure:        githubcomnimblewaynimblego.Bool(true),
				Value:         githubcomnimblewaynimblego.String("value"),
			}},
		},
		Country:             githubcomnimblewaynimblego.ExtractExtractParamsCountryUs,
		Device:              githubcomnimblewaynimblego.ExtractExtractParamsDeviceDesktop,
		DisableIPCheck:      githubcomnimblewaynimblego.Bool(false),
		Driver:              githubcomnimblewaynimblego.ExtractExtractParamsDriverVx8,
		ExpectedStatusCodes: []int64{200, 201},
		Formats:             []string{"html"},
		Headers: map[string]githubcomnimblewaynimblego.ExtractExtractParamsHeaderUnion{
			"User-Agent": {
				OfString: githubcomnimblewaynimblego.String("CustomBot/1.0"),
			},
			"Accept-Language": {
				OfString: githubcomnimblewaynimblego.String("en-US"),
			},
		},
		Http2:  githubcomnimblewaynimblego.Bool(true),
		Ip6:    githubcomnimblewaynimblego.Bool(false),
		IsXhr:  githubcomnimblewaynimblego.Bool(true),
		Locale: githubcomnimblewaynimblego.ExtractExtractParamsLocaleEnUs,
		Metadata: githubcomnimblewaynimblego.ExtractExtractParamsMetadata{
			AccountName:         githubcomnimblewaynimblego.String("account_name"),
			APIType:             githubcomnimblewaynimblego.String("api_type"),
			CrawlDepth:          githubcomnimblewaynimblego.Int(-9007199254740991),
			CrawlID:             githubcomnimblewaynimblego.String("crawl_id"),
			DefinitionID:        githubcomnimblewaynimblego.Int(-9007199254740991),
			DefinitionName:      githubcomnimblewaynimblego.String("definition_name"),
			Endpoint:            githubcomnimblewaynimblego.String("endpoint"),
			ExecutionID:         githubcomnimblewaynimblego.String("execution_id"),
			FlowitTaskID:        githubcomnimblewaynimblego.String("flowit_task_id"),
			InputID:             githubcomnimblewaynimblego.String("input_id"),
			IsPublicWsa:         githubcomnimblewaynimblego.Bool(true),
			IsSitemap:           githubcomnimblewaynimblego.Bool(true),
			IsWsa:               githubcomnimblewaynimblego.Bool(true),
			ParserID:            githubcomnimblewaynimblego.String("parser_id"),
			PipelineExecutionID: githubcomnimblewaynimblego.Int(-9007199254740991),
			QueryTemplateID:     githubcomnimblewaynimblego.String("query_template_id"),
			Source:              githubcomnimblewaynimblego.String("source"),
			TemplateID:          githubcomnimblewaynimblego.Int(-9007199254740991),
			TemplateName:        githubcomnimblewaynimblego.String("template_name"),
			WsaID:               githubcomnimblewaynimblego.String("wsa_id"),
			WsaName:             githubcomnimblewaynimblego.String("wsa_name"),
			WsaVersion:          githubcomnimblewaynimblego.Float(0),
		},
		Method:     githubcomnimblewaynimblego.ExtractExtractParamsMethodGet,
		NativeMode: githubcomnimblewaynimblego.ExtractExtractParamsNativeModeRequester,
		NetworkCapture: []githubcomnimblewaynimblego.ExtractExtractParamsNetworkCapture{{
			Method: "GET",
			ResourceType: githubcomnimblewaynimblego.ExtractExtractParamsNetworkCaptureResourceTypeUnion{
				OfString: githubcomnimblewaynimblego.String("document"),
			},
			StatusCode: githubcomnimblewaynimblego.ExtractExtractParamsNetworkCaptureStatusCodeUnion{
				OfFloat: githubcomnimblewaynimblego.Float(100),
			},
			URL: githubcomnimblewaynimblego.ExtractExtractParamsNetworkCaptureURL{
				Value: "value",
				Type:  "exact",
			},
			Validation:                  githubcomnimblewaynimblego.Bool(true),
			WaitForRequestsCount:        githubcomnimblewaynimblego.Float(0),
			WaitForRequestsCountTimeout: githubcomnimblewaynimblego.Float(1),
		}},
		NoUserbrowser: githubcomnimblewaynimblego.Bool(false),
		Os:            githubcomnimblewaynimblego.ExtractExtractParamsOsWindows,
		Parse:         githubcomnimblewaynimblego.Bool(true),
		Parser: githubcomnimblewaynimblego.ExtractExtractParamsParserUnion{
			OfAnyMap: map[string]any{
				"myParser": "bar",
			},
		},
		ProxyProvider: githubcomnimblewaynimblego.ExtractExtractParamsProxyProviderBrightdata,
		ProxyProviders: map[string]float64{
			"brightdata": 70,
			"oxylabs":    30,
		},
		QueryTemplate: githubcomnimblewaynimblego.ExtractExtractParamsQueryTemplate{
			ID:      "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			APIType: "WEB",
			Pagination: githubcomnimblewaynimblego.ExtractExtractParamsQueryTemplatePaginationUnion{
				OfExtractExtractsQueryTemplatePaginationNextPageParams: &githubcomnimblewaynimblego.ExtractExtractParamsQueryTemplatePaginationNextPageParams{
					NextPageParams: map[string]any{
						"foo": "bar",
					},
				},
			},
			Params: map[string]any{
				"foo": "bar",
			},
		},
		RawHeaders:   githubcomnimblewaynimblego.Bool(true),
		ReferrerType: githubcomnimblewaynimblego.ExtractExtractParamsReferrerTypeRandom,
		Render:       githubcomnimblewaynimblego.Bool(true),
		RenderFlow: []map[string]any{{
			"wait": "bar",
		}, {
			"click": "bar",
		}},
		RenderOptions: githubcomnimblewaynimblego.ExtractExtractParamsRenderOptions{
			Adblock:        githubcomnimblewaynimblego.Bool(true),
			BlockedDomains: []string{"ads.example.com", "tracker.com"},
			BrowserEngine: githubcomnimblewaynimblego.ExtractExtractParamsRenderOptionsBrowserEngineUnion{
				OfExtractExtractsRenderOptionsBrowserEngineString: githubcomnimblewaynimblego.String("chrome"),
			},
			Cache:             githubcomnimblewaynimblego.Bool(false),
			ConnectorType:     "puppeteer",
			DisabledResources: []string{"image", "stylesheet"},
			Enable2captcha:    githubcomnimblewaynimblego.Bool(true),
			Extensions:        []string{"extension-id-1", "extension-id-2"},
			FingerprintID:     githubcomnimblewaynimblego.String("fp-abc123"),
			HackiumConfiguration: githubcomnimblewaynimblego.ExtractExtractParamsRenderOptionsHackiumConfiguration{
				CollectLogs:                 githubcomnimblewaynimblego.Bool(true),
				DoNotFixMathSalt:            githubcomnimblewaynimblego.Bool(true),
				EnableDocumentElementSpoof:  githubcomnimblewaynimblego.Bool(true),
				EnableDocumentHasFocus:      githubcomnimblewaynimblego.Bool(true),
				EnableFakeNavigationHistory: githubcomnimblewaynimblego.Bool(true),
				EnableKeyOrdering:           githubcomnimblewaynimblego.Bool(true),
				EnableSniffer:               githubcomnimblewaynimblego.Bool(true),
				EnableVerboseLogs:           githubcomnimblewaynimblego.Bool(true),
			},
			Headless:               githubcomnimblewaynimblego.Bool(true),
			IncludeIframes:         githubcomnimblewaynimblego.Bool(true),
			LoadLocalStorage:       githubcomnimblewaynimblego.Bool(true),
			LocalStorageKeysToLoad: []string{"authToken", "userId"},
			MouseStrategy:          "linear",
			NoAcceptEncoding:       githubcomnimblewaynimblego.Bool(true),
			OverridePermissions:    githubcomnimblewaynimblego.Bool(true),
			RandomHeaderOrder:      githubcomnimblewaynimblego.Bool(true),
			RenderType:             "domcontentloaded",
			StoreLocalStorage:      githubcomnimblewaynimblego.Bool(true),
			Timeout:                githubcomnimblewaynimblego.Float(30000),
			TypingInterval:         githubcomnimblewaynimblego.Float(100),
			TypingStrategy:         "simple",
			Userbrowser:            githubcomnimblewaynimblego.Bool(true),
			WaitUntil:              "networkidle2",
			WithPerformanceMetrics: githubcomnimblewaynimblego.Bool(true),
		},
		RequestTimeout:  githubcomnimblewaynimblego.Float(30000),
		SaveUserbrowser: githubcomnimblewaynimblego.Bool(false),
		Session: githubcomnimblewaynimblego.ExtractExtractParamsSession{
			ID:                  githubcomnimblewaynimblego.String("id"),
			PrefetchUserbrowser: githubcomnimblewaynimblego.Bool(true),
			Retry:               githubcomnimblewaynimblego.Bool(true),
			Timeout:             githubcomnimblewaynimblego.Float(1),
		},
		Skill: githubcomnimblewaynimblego.ExtractExtractParamsSkillUnion{
			OfString: githubcomnimblewaynimblego.String("dynamic-content"),
		},
		SkipUbct: githubcomnimblewaynimblego.Bool(false),
		State:    githubcomnimblewaynimblego.ExtractExtractParamsStateCa,
		Tag:      githubcomnimblewaynimblego.String("campaign-2024-q1"),
		Template: githubcomnimblewaynimblego.ExtractExtractParamsTemplate{
			Name: "x",
			Params: map[string]any{
				"foo": "bar",
			},
		},
		Type: githubcomnimblewaynimblego.String("generic"),
		UserbrowserCreationTemplateRendered: githubcomnimblewaynimblego.ExtractExtractParamsUserbrowserCreationTemplateRendered{
			ID:                    "id",
			AllowedParameterNames: []string{"x"},
			RenderFlowRendered: []map[string]any{{
				"foo": "bar",
			}},
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
