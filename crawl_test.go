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

func TestCrawlListWithOptionalParams(t *testing.T) {
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
			Browser: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserUnion{
				OfCrawlRunsExtractOptionsBrowserString: githubcomnimblewaynimblego.String("chrome"),
			},
			BrowserActions: []githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionUnion{{
				OfCrawlRunsExtractOptionsBrowserActionGotoAction: &githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionGotoAction{
					Goto: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionGotoActionGotoUnion{
						OfString: githubcomnimblewaynimblego.String("https://example.com/login"),
					},
				},
			}, {
				OfCrawlRunsExtractOptionsBrowserActionWaitForElementAction: &githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionWaitForElementAction{
					WaitForElement: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementUnion{
						OfString: githubcomnimblewaynimblego.String("#login-form"),
					},
				},
			}, {
				OfCrawlRunsExtractOptionsBrowserActionFillAction: &githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillAction{
					Fill: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillUnion{
						OfType: &githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillType{
							Selector: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillTypeSelectorUnion{
								OfString: githubcomnimblewaynimblego.String("#username"),
							},
							Value:          "user@example.com",
							ClickOnElement: githubcomnimblewaynimblego.Bool(true),
							Delay: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillTypeDelayUnion{
								OfFloat: githubcomnimblewaynimblego.Float(1000),
							},
							Mode:                  "type",
							MouseMovementStrategy: "linear",
							Required: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillTypeRequiredUnion{
								OfCrawlRunsExtractOptionsBrowserActionFillActionFillTypeRequiredString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillTypeRequiredStringTrue),
							},
							Scroll: githubcomnimblewaynimblego.Bool(true),
							Skip: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillTypeSkipUnion{
								OfCrawlRunsExtractOptionsBrowserActionFillActionFillTypeSkipString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillTypeSkipStringTrue),
							},
							Timeout: githubcomnimblewaynimblego.Float(0),
							TypingInterval: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillTypeTypingIntervalUnion{
								OfFloat: githubcomnimblewaynimblego.Float(1000),
							},
							TypingStrategy: "simple",
							Visible:        githubcomnimblewaynimblego.Bool(true),
						},
					},
				},
			}, {
				OfCrawlRunsExtractOptionsBrowserActionFillAction: &githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillAction{
					Fill: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillUnion{
						OfType: &githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillType{
							Selector: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillTypeSelectorUnion{
								OfString: githubcomnimblewaynimblego.String("#password"),
							},
							Value:          "password123",
							ClickOnElement: githubcomnimblewaynimblego.Bool(true),
							Delay: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillTypeDelayUnion{
								OfFloat: githubcomnimblewaynimblego.Float(1000),
							},
							Mode:                  "type",
							MouseMovementStrategy: "linear",
							Required: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillTypeRequiredUnion{
								OfCrawlRunsExtractOptionsBrowserActionFillActionFillTypeRequiredString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillTypeRequiredStringTrue),
							},
							Scroll: githubcomnimblewaynimblego.Bool(true),
							Skip: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillTypeSkipUnion{
								OfCrawlRunsExtractOptionsBrowserActionFillActionFillTypeSkipString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillTypeSkipStringTrue),
							},
							Timeout: githubcomnimblewaynimblego.Float(0),
							TypingInterval: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionFillActionFillTypeTypingIntervalUnion{
								OfFloat: githubcomnimblewaynimblego.Float(1000),
							},
							TypingStrategy: "simple",
							Visible:        githubcomnimblewaynimblego.Bool(true),
						},
					},
				},
			}, {
				OfCrawlRunsExtractOptionsBrowserActionClickAction: &githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionClickAction{
					Click: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionClickActionClickUnion{
						OfString: githubcomnimblewaynimblego.String("#submit"),
					},
				},
			}, {
				OfCrawlRunsExtractOptionsBrowserActionScreenshotAction: &githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionScreenshotAction{
					Screenshot: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionScreenshotActionScreenshotUnion{
						OfCrawlRunsExtractOptionsBrowserActionScreenshotActionScreenshotObject: &githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionScreenshotActionScreenshotObject{
							Format:   "png",
							FullPage: githubcomnimblewaynimblego.Bool(true),
							Quality:  githubcomnimblewaynimblego.Float(0),
							Required: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredUnion{
								OfCrawlRunsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredStringTrue),
							},
							Skip: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipUnion{
								OfCrawlRunsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipStringTrue),
							},
						},
					},
				},
			}},
			City:          githubcomnimblewaynimblego.String("Los Angeles"),
			ClientTimeout: githubcomnimblewaynimblego.Float(25000),
			ConsentHeader: githubcomnimblewaynimblego.Bool(true),
			Cookies: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsCookiesUnion{
				OfCrawlRunsExtractOptionsCookiesArray: []githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsCookiesArrayItem{{
					Creation:     githubcomnimblewaynimblego.String("creation"),
					Domain:       githubcomnimblewaynimblego.String("domain"),
					Expires:      githubcomnimblewaynimblego.String("expires"),
					Extensions:   []string{"string"},
					HostOnly:     githubcomnimblewaynimblego.Bool(true),
					HTTPOnly:     githubcomnimblewaynimblego.Bool(true),
					LastAccessed: githubcomnimblewaynimblego.String("lastAccessed"),
					MaxAge: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsCookiesArrayItemMaxAgeUnion{
						OfCrawlRunsExtractOptionsCookiesArrayItemMaxAgeString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsCookiesArrayItemMaxAgeStringInfinity),
					},
					Name:          githubcomnimblewaynimblego.String("name"),
					Path:          githubcomnimblewaynimblego.String("path"),
					PathIsDefault: githubcomnimblewaynimblego.Bool(true),
					SameSite:      "strict",
					Secure:        githubcomnimblewaynimblego.Bool(true),
					Value:         githubcomnimblewaynimblego.String("value"),
				}},
			},
			Country:             githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsCountryUs,
			Device:              "desktop",
			DisableIPCheck:      githubcomnimblewaynimblego.Bool(false),
			Driver:              "vx8",
			ExpectedStatusCodes: []int64{200, 201},
			Formats:             []string{"html"},
			Headers: map[string]githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsHeaderUnion{
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
			Locale: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsLocaleEnUs,
			Metadata: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsMetadata{
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
			Method:     "GET",
			NativeMode: "requester",
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
			NoUserbrowser: githubcomnimblewaynimblego.Bool(false),
			Os:            "windows",
			Parse:         githubcomnimblewaynimblego.Bool(true),
			Parser: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsParserUnion{
				OfAnyMap: map[string]any{
					"myParser": "bar",
				},
			},
			ProxyProvider: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsProxyProviderBrightdata,
			ProxyProviders: map[string]float64{
				"brightdata": 70,
				"oxylabs":    30,
			},
			QueryTemplate: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsQueryTemplate{
				ID:      "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				APIType: "WEB",
				Pagination: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsQueryTemplatePaginationUnion{
					OfCrawlRunsExtractOptionsQueryTemplatePaginationNextPageParams: &githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsQueryTemplatePaginationNextPageParams{
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
			ReferrerType: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsReferrerTypeRandom,
			Render:       githubcomnimblewaynimblego.Bool(true),
			RenderFlow: []map[string]any{{
				"wait": "bar",
			}, {
				"click": "bar",
			}},
			RenderOptions: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsRenderOptions{
				Adblock:        githubcomnimblewaynimblego.Bool(true),
				BlockedDomains: []string{"ads.example.com", "tracker.com"},
				BrowserEngine: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsRenderOptionsBrowserEngineUnion{
					OfCrawlRunsExtractOptionsRenderOptionsBrowserEngineString: githubcomnimblewaynimblego.String("chrome"),
				},
				Cache:             githubcomnimblewaynimblego.Bool(false),
				ConnectorType:     "puppeteer",
				DisabledResources: []string{"image", "stylesheet"},
				Enable2captcha:    githubcomnimblewaynimblego.Bool(true),
				Extensions:        []string{"extension-id-1", "extension-id-2"},
				FingerprintID:     githubcomnimblewaynimblego.String("fp-abc123"),
				HackiumConfiguration: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsRenderOptionsHackiumConfiguration{
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
			Session: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsSession{
				ID:                  githubcomnimblewaynimblego.String("id"),
				PrefetchUserbrowser: githubcomnimblewaynimblego.Bool(true),
				Retry:               githubcomnimblewaynimblego.Bool(true),
				Timeout:             githubcomnimblewaynimblego.Float(1),
			},
			Skill: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsSkillUnion{
				OfString: githubcomnimblewaynimblego.String("dynamic-content"),
			},
			SkipUbct: githubcomnimblewaynimblego.Bool(false),
			State:    "CA",
			Tag:      githubcomnimblewaynimblego.String("campaign-2024-q1"),
			Template: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsTemplate{
				Name: "x",
				Params: map[string]any{
					"foo": "bar",
				},
			},
			Type: githubcomnimblewaynimblego.String("generic"),
			URL:  githubcomnimblewaynimblego.String("url"),
			UserbrowserCreationTemplateRendered: githubcomnimblewaynimblego.CrawlRunParamsExtractOptionsUserbrowserCreationTemplateRendered{
				ID:                    "id",
				AllowedParameterNames: []string{"x"},
				RenderFlowRendered: []map[string]any{{
					"foo": "bar",
				}},
			},
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
	_, err := client.Crawl.Terminate(context.TODO(), "123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
