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
		Status: githubcomnimblewaynimblego.CrawlListParamsStatusPending,
		Cursor: githubcomnimblewaynimblego.String("cursor"),
		Limit:  githubcomnimblewaynimblego.Int(10),
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestCrawlRootWithOptionalParams(t *testing.T) {
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
	_, err := client.Crawl.Root(context.TODO(), githubcomnimblewaynimblego.CrawlRootParams{
		URL:                "https://example.com",
		AllowExternalLinks: githubcomnimblewaynimblego.Bool(false),
		AllowSubdomains:    githubcomnimblewaynimblego.Bool(false),
		Callback: githubcomnimblewaynimblego.CrawlRootParamsCallbackUnion{
			OfCrawlRootsCallbackObject: &githubcomnimblewaynimblego.CrawlRootParamsCallbackObject{
				URL:    "https://example.com/webhook",
				Events: []string{"page"},
				Headers: map[string]any{
					"X-Custom-Header": "bar",
				},
				Metadata: map[string]any{
					"crawlId": "bar",
				},
			},
		},
		CrawlEntireDomain: githubcomnimblewaynimblego.Bool(false),
		ExcludePaths:      []string{"/exclude-this-path", "/and-this-path"},
		ExtractOptions: githubcomnimblewaynimblego.CrawlRootParamsExtractOptions{
			DebugOptions: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsDebugOptions{
				CollectHar: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsDebugOptionsCollectHarUnion{
					OfBool: githubcomnimblewaynimblego.Bool(true),
				},
				NoRetryMode: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsDebugOptionsNoRetryModeUnion{
					OfBool: githubcomnimblewaynimblego.Bool(true),
				},
				RecordScreen: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsDebugOptionsRecordScreenUnion{
					OfBool: githubcomnimblewaynimblego.Bool(true),
				},
				Redact: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsDebugOptionsRedactUnion{
					OfBool: githubcomnimblewaynimblego.Bool(true),
				},
				ShowCursor: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsDebugOptionsShowCursorUnion{
					OfBool: githubcomnimblewaynimblego.Bool(true),
				},
				SolveCaptcha: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsDebugOptionsSolveCaptchaUnion{
					OfBool: githubcomnimblewaynimblego.Bool(true),
				},
				Trace: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsDebugOptionsTraceUnion{
					OfBool: githubcomnimblewaynimblego.Bool(true),
				},
				UploadEngineLogs: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsDebugOptionsUploadEngineLogsUnion{
					OfBool: githubcomnimblewaynimblego.Bool(true),
				},
				Verbose: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsDebugOptionsVerboseUnion{
					OfBool: githubcomnimblewaynimblego.Bool(true),
				},
				WithProxyUsage: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsDebugOptionsWithProxyUsageUnion{
					OfBool: githubcomnimblewaynimblego.Bool(true),
				},
			},
			URL: "https://example.com/page",
			Browser: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsBrowserUnion{
				OfCrawlRootsExtractOptionsBrowserString: githubcomnimblewaynimblego.String("chrome"),
			},
			City:          githubcomnimblewaynimblego.String("Los Angeles"),
			ClientTimeout: githubcomnimblewaynimblego.Float(25000),
			ConsentHeader: githubcomnimblewaynimblego.Bool(true),
			Cookies: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsCookiesUnion{
				OfCrawlRootsExtractOptionsCookiesArray: []githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsCookiesArrayItem{{
					Creation:     githubcomnimblewaynimblego.String("creation"),
					Domain:       githubcomnimblewaynimblego.String("domain"),
					Expires:      githubcomnimblewaynimblego.String("expires"),
					Extensions:   []string{"string"},
					HostOnly:     githubcomnimblewaynimblego.Bool(true),
					HTTPOnly:     githubcomnimblewaynimblego.Bool(true),
					LastAccessed: githubcomnimblewaynimblego.String("lastAccessed"),
					MaxAge: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeUnion{
						OfCrawlRootsExtractOptionsCookiesArrayItemMaxAgeString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeStringInfinity),
					},
					Name:          githubcomnimblewaynimblego.String("name"),
					Path:          githubcomnimblewaynimblego.String("path"),
					PathIsDefault: githubcomnimblewaynimblego.Bool(true),
					SameSite:      "strict",
					Secure:        githubcomnimblewaynimblego.Bool(true),
					Value:         githubcomnimblewaynimblego.String("value"),
				}},
			},
			Country:        githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsCountryUs,
			Device:         "desktop",
			DisableIPCheck: githubcomnimblewaynimblego.Bool(false),
			Driver:         "vx8",
			DynamicParser: map[string]any{
				"myParser": "bar",
			},
			ExpectedStatusCodes: []int64{200, 201},
			ExportUserbrowser:   githubcomnimblewaynimblego.Bool(false),
			Format:              "json",
			Headers: map[string]githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsHeaderUnion{
				"User-Agent": {
					OfString: githubcomnimblewaynimblego.String("CustomBot/1.0"),
				},
				"Accept-Language": {
					OfString: githubcomnimblewaynimblego.String("en-US"),
				},
			},
			Http2:    githubcomnimblewaynimblego.Bool(true),
			Ip6:      githubcomnimblewaynimblego.Bool(false),
			IsXhr:    githubcomnimblewaynimblego.Bool(true),
			Locale:   githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsLocaleEnUs,
			Markdown: githubcomnimblewaynimblego.Bool(false),
			Metadata: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsMetadata{
				AccountName:         githubcomnimblewaynimblego.String("acme-corp"),
				DefinitionID:        githubcomnimblewaynimblego.Int(456),
				DefinitionName:      githubcomnimblewaynimblego.String("product-scraper"),
				Endpoint:            githubcomnimblewaynimblego.String("/api/v2/scrape"),
				ExecutionID:         githubcomnimblewaynimblego.String("exec-abc123"),
				FlowitTaskID:        githubcomnimblewaynimblego.String("task-xyz789"),
				InputID:             githubcomnimblewaynimblego.String("input-123"),
				PipelineExecutionID: githubcomnimblewaynimblego.Int(12345),
				QueryTemplateID:     githubcomnimblewaynimblego.String("template-qry-001"),
				Source:              githubcomnimblewaynimblego.String("web-app"),
				TemplateID:          githubcomnimblewaynimblego.Int(789),
				TemplateName:        githubcomnimblewaynimblego.String("e-commerce-template"),
			},
			Method:     "GET",
			NativeMode: "requester",
			NetworkCapture: []githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsNetworkCapture{{
				Method: "GET",
				ResourceType: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsNetworkCaptureResourceTypeUnion{
					OfString: githubcomnimblewaynimblego.String("document"),
				},
				StatusCode: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsNetworkCaptureStatusCodeUnion{
					OfFloat: githubcomnimblewaynimblego.Float(100),
				},
				URL: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsNetworkCaptureURL{
					Value: "value",
					Type:  "exact",
				},
				Validation:                  githubcomnimblewaynimblego.Bool(true),
				WaitForRequestsCount:        githubcomnimblewaynimblego.Float(0),
				WaitForRequestsCountTimeout: githubcomnimblewaynimblego.Float(1),
			}},
			NoHTML:        githubcomnimblewaynimblego.Bool(false),
			NoUserbrowser: githubcomnimblewaynimblego.Bool(false),
			Os:            "windows",
			Parse:         githubcomnimblewaynimblego.Bool(true),
			ParseOptions: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsParseOptions{
				MergeDynamic: githubcomnimblewaynimblego.Bool(true),
			},
			Parser: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsParserUnion{
				OfAnyMap: map[string]any{
					"myParser": "bar",
				},
			},
			ProxyProvider: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsProxyProviderBrightdata,
			ProxyProviders: map[string]float64{
				"brightdata": 70,
				"oxylabs":    30,
			},
			QueryTemplate: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsQueryTemplate{
				ID:      "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				APIType: "WEB",
				Pagination: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsQueryTemplatePaginationUnion{
					OfCrawlRootsExtractOptionsQueryTemplatePaginationNextPageParams: &githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsQueryTemplatePaginationNextPageParams{
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
			ReferrerType: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsReferrerTypeRandom,
			Render:       githubcomnimblewaynimblego.Bool(true),
			RenderFlow: []map[string]any{{
				"wait": "bar",
			}, {
				"click": "bar",
			}},
			RenderOptions: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsRenderOptions{
				Adblock:        githubcomnimblewaynimblego.Bool(true),
				BlockedDomains: []string{"ads.example.com", "tracker.com"},
				BrowserEngine: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineUnion{
					OfCrawlRootsExtractOptionsRenderOptionsBrowserEngineString: githubcomnimblewaynimblego.String("chrome"),
				},
				Cache:             githubcomnimblewaynimblego.Bool(false),
				ConnectorType:     "webit-cdp",
				DisabledResources: []string{"image", "stylesheet"},
				Enable2captcha:    githubcomnimblewaynimblego.Bool(true),
				Extensions:        []string{"extension-id-1", "extension-id-2"},
				FingerprintID:     githubcomnimblewaynimblego.String("fp-abc123"),
				HackiumConfiguration: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsRenderOptionsHackiumConfiguration{
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
				RenderType:             "load",
				StoreLocalStorage:      githubcomnimblewaynimblego.Bool(true),
				Timeout:                githubcomnimblewaynimblego.Float(30000),
				TypingInterval:         githubcomnimblewaynimblego.Float(100),
				TypingStrategy:         "simple",
				Userbrowser:            githubcomnimblewaynimblego.Bool(true),
				WaitUntil:              "networkidle2",
				WithPerformanceMetrics: githubcomnimblewaynimblego.Bool(true),
			},
			RequestTimeout:                githubcomnimblewaynimblego.Float(30000),
			ReturnResponseHeadersAsHeader: githubcomnimblewaynimblego.Bool(true),
			SaveUserbrowser:               githubcomnimblewaynimblego.Bool(false),
			Session: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsSession{
				ID:                  githubcomnimblewaynimblego.String("id"),
				PrefetchUserbrowser: githubcomnimblewaynimblego.Bool(true),
				Retry:               githubcomnimblewaynimblego.Bool(true),
				Timeout:             githubcomnimblewaynimblego.Float(1),
			},
			Skill: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsSkillUnion{
				OfString: githubcomnimblewaynimblego.String("dynamic-content"),
			},
			SkipUbct: githubcomnimblewaynimblego.Bool(false),
			State:    "CA",
			Tag:      githubcomnimblewaynimblego.String("campaign-2024-q1"),
			Template: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsTemplate{
				Name: "x",
				Params: map[string]any{
					"foo": "bar",
				},
			},
			Type: githubcomnimblewaynimblego.String("generic"),
			UserbrowserCreationTemplateRendered: githubcomnimblewaynimblego.CrawlRootParamsExtractOptionsUserbrowserCreationTemplateRendered{
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
		Sitemap:               githubcomnimblewaynimblego.CrawlRootParamsSitemapInclude,
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
