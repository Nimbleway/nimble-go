// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nimblego_test

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
	client := nimblego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Crawl.List(context.TODO(), nimblego.CrawlListParams{
		Status: nimblego.CrawlListParamsStatusPending,
		Cursor: nimblego.String("cursor"),
		Limit:  nimblego.Int(10),
	})
	if err != nil {
		var apierr *nimblego.Error
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
	client := nimblego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Crawl.Root(context.TODO(), nimblego.CrawlRootParams{
		URL:                "https://example.com",
		AllowExternalLinks: nimblego.Bool(false),
		AllowSubdomains:    nimblego.Bool(false),
		Callback: nimblego.CrawlRootParamsCallbackUnion{
			OfCrawlRootsCallbackObject: &nimblego.CrawlRootParamsCallbackObject{
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
		CrawlEntireDomain: nimblego.Bool(false),
		ExcludePaths:      []string{"/exclude-this-path", "/and-this-path"},
		ExtractOptions: nimblego.CrawlRootParamsExtractOptions{
			DebugOptions: nimblego.CrawlRootParamsExtractOptionsDebugOptions{
				CollectHar: nimblego.CrawlRootParamsExtractOptionsDebugOptionsCollectHarUnion{
					OfBool: nimblego.Bool(true),
				},
				NoRetryMode: nimblego.CrawlRootParamsExtractOptionsDebugOptionsNoRetryModeUnion{
					OfBool: nimblego.Bool(true),
				},
				RecordScreen: nimblego.CrawlRootParamsExtractOptionsDebugOptionsRecordScreenUnion{
					OfBool: nimblego.Bool(true),
				},
				Redact: nimblego.CrawlRootParamsExtractOptionsDebugOptionsRedactUnion{
					OfBool: nimblego.Bool(true),
				},
				ShowCursor: nimblego.CrawlRootParamsExtractOptionsDebugOptionsShowCursorUnion{
					OfBool: nimblego.Bool(true),
				},
				SolveCaptcha: nimblego.CrawlRootParamsExtractOptionsDebugOptionsSolveCaptchaUnion{
					OfBool: nimblego.Bool(true),
				},
				Trace: nimblego.CrawlRootParamsExtractOptionsDebugOptionsTraceUnion{
					OfBool: nimblego.Bool(true),
				},
				UploadEngineLogs: nimblego.CrawlRootParamsExtractOptionsDebugOptionsUploadEngineLogsUnion{
					OfBool: nimblego.Bool(true),
				},
				Verbose: nimblego.CrawlRootParamsExtractOptionsDebugOptionsVerboseUnion{
					OfBool: nimblego.Bool(true),
				},
				WithProxyUsage: nimblego.CrawlRootParamsExtractOptionsDebugOptionsWithProxyUsageUnion{
					OfBool: nimblego.Bool(true),
				},
			},
			URL: "https://example.com/page",
			Browser: nimblego.CrawlRootParamsExtractOptionsBrowserUnion{
				OfCrawlRootsExtractOptionsBrowserString: nimblego.String("chrome"),
			},
			City:          nimblego.String("Los Angeles"),
			ClientTimeout: nimblego.Float(25000),
			ConsentHeader: nimblego.Bool(true),
			Cookies: nimblego.CrawlRootParamsExtractOptionsCookiesUnion{
				OfCrawlRootsExtractOptionsCookiesArray: []nimblego.CrawlRootParamsExtractOptionsCookiesArrayItem{{
					Creation:     nimblego.String("creation"),
					Domain:       nimblego.String("domain"),
					Expires:      nimblego.String("expires"),
					Extensions:   []string{"string"},
					HostOnly:     nimblego.Bool(true),
					HTTPOnly:     nimblego.Bool(true),
					LastAccessed: nimblego.String("lastAccessed"),
					MaxAge: nimblego.CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeUnion{
						OfCrawlRootsExtractOptionsCookiesArrayItemMaxAgeString: nimblego.Opt(nimblego.CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeStringInfinity),
					},
					Name:          nimblego.String("name"),
					Path:          nimblego.String("path"),
					PathIsDefault: nimblego.Bool(true),
					SameSite:      "strict",
					Secure:        nimblego.Bool(true),
					Value:         nimblego.String("value"),
				}},
			},
			Country:        nimblego.CrawlRootParamsExtractOptionsCountryUs,
			Device:         "desktop",
			DisableIPCheck: nimblego.Bool(false),
			Driver:         "vx8",
			DynamicParser: map[string]any{
				"myParser": "bar",
			},
			ExpectedStatusCodes: []int64{200, 201},
			ExportUserbrowser:   nimblego.Bool(false),
			Format:              "json",
			Headers: map[string]nimblego.CrawlRootParamsExtractOptionsHeaderUnion{
				"User-Agent": {
					OfString: nimblego.String("CustomBot/1.0"),
				},
				"Accept-Language": {
					OfString: nimblego.String("en-US"),
				},
			},
			Http2:    nimblego.Bool(true),
			Ip6:      nimblego.Bool(false),
			IsXhr:    nimblego.Bool(true),
			Locale:   nimblego.CrawlRootParamsExtractOptionsLocaleEnUs,
			Markdown: nimblego.Bool(false),
			Metadata: nimblego.CrawlRootParamsExtractOptionsMetadata{
				AccountName:         nimblego.String("acme-corp"),
				DefinitionID:        nimblego.Int(456),
				DefinitionName:      nimblego.String("product-scraper"),
				Endpoint:            nimblego.String("/api/v2/scrape"),
				ExecutionID:         nimblego.String("exec-abc123"),
				FlowitTaskID:        nimblego.String("task-xyz789"),
				InputID:             nimblego.String("input-123"),
				PipelineExecutionID: nimblego.Int(12345),
				QueryTemplateID:     nimblego.String("template-qry-001"),
				Source:              nimblego.String("web-app"),
				TemplateID:          nimblego.Int(789),
				TemplateName:        nimblego.String("e-commerce-template"),
			},
			Method:     "GET",
			NativeMode: "requester",
			NetworkCapture: []nimblego.CrawlRootParamsExtractOptionsNetworkCapture{{
				Method: "GET",
				ResourceType: nimblego.CrawlRootParamsExtractOptionsNetworkCaptureResourceTypeUnion{
					OfString: nimblego.String("document"),
				},
				StatusCode: nimblego.CrawlRootParamsExtractOptionsNetworkCaptureStatusCodeUnion{
					OfFloat: nimblego.Float(100),
				},
				URL: nimblego.CrawlRootParamsExtractOptionsNetworkCaptureURL{
					Value: "value",
					Type:  "exact",
				},
				Validation:                  nimblego.Bool(true),
				WaitForRequestsCount:        nimblego.Float(0),
				WaitForRequestsCountTimeout: nimblego.Float(1),
			}},
			NoHTML:        nimblego.Bool(false),
			NoUserbrowser: nimblego.Bool(false),
			Os:            "windows",
			Parse:         nimblego.Bool(true),
			ParseOptions: nimblego.CrawlRootParamsExtractOptionsParseOptions{
				MergeDynamic: nimblego.Bool(true),
			},
			Parser: nimblego.CrawlRootParamsExtractOptionsParserUnion{
				OfAnyMap: map[string]any{
					"myParser": "bar",
				},
			},
			ProxyProvider: nimblego.CrawlRootParamsExtractOptionsProxyProviderBrightdata,
			ProxyProviders: map[string]float64{
				"brightdata": 70,
				"oxylabs":    30,
			},
			QueryTemplate: nimblego.CrawlRootParamsExtractOptionsQueryTemplate{
				ID:      "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				APIType: "WEB",
				Pagination: nimblego.CrawlRootParamsExtractOptionsQueryTemplatePaginationUnion{
					OfCrawlRootsExtractOptionsQueryTemplatePaginationNextPageParams: &nimblego.CrawlRootParamsExtractOptionsQueryTemplatePaginationNextPageParams{
						NextPageParams: map[string]any{
							"foo": "bar",
						},
					},
				},
				Params: map[string]any{
					"foo": "bar",
				},
			},
			RawHeaders:   nimblego.Bool(true),
			ReferrerType: nimblego.CrawlRootParamsExtractOptionsReferrerTypeRandom,
			Render:       nimblego.Bool(true),
			RenderFlow: []map[string]any{{
				"wait": "bar",
			}, {
				"click": "bar",
			}},
			RenderOptions: nimblego.CrawlRootParamsExtractOptionsRenderOptions{
				Adblock:        nimblego.Bool(true),
				BlockedDomains: []string{"ads.example.com", "tracker.com"},
				BrowserEngine: nimblego.CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineUnion{
					OfCrawlRootsExtractOptionsRenderOptionsBrowserEngineString: nimblego.String("chrome"),
				},
				Cache:             nimblego.Bool(false),
				ConnectorType:     "webit-cdp",
				DisabledResources: []string{"image", "stylesheet"},
				Enable2captcha:    nimblego.Bool(true),
				Extensions:        []string{"extension-id-1", "extension-id-2"},
				FingerprintID:     nimblego.String("fp-abc123"),
				HackiumConfiguration: nimblego.CrawlRootParamsExtractOptionsRenderOptionsHackiumConfiguration{
					CollectLogs:                 nimblego.Bool(true),
					DoNotFixMathSalt:            nimblego.Bool(true),
					EnableDocumentElementSpoof:  nimblego.Bool(true),
					EnableDocumentHasFocus:      nimblego.Bool(true),
					EnableFakeNavigationHistory: nimblego.Bool(true),
					EnableKeyOrdering:           nimblego.Bool(true),
					EnableSniffer:               nimblego.Bool(true),
					EnableVerboseLogs:           nimblego.Bool(true),
				},
				Headless:               nimblego.Bool(true),
				IncludeIframes:         nimblego.Bool(true),
				LoadLocalStorage:       nimblego.Bool(true),
				LocalStorageKeysToLoad: []string{"authToken", "userId"},
				MouseStrategy:          "linear",
				NoAcceptEncoding:       nimblego.Bool(true),
				OverridePermissions:    nimblego.Bool(true),
				RandomHeaderOrder:      nimblego.Bool(true),
				RenderType:             "load",
				StoreLocalStorage:      nimblego.Bool(true),
				Timeout:                nimblego.Float(30000),
				TypingInterval:         nimblego.Float(100),
				TypingStrategy:         "simple",
				Userbrowser:            nimblego.Bool(true),
				WaitUntil:              "networkidle2",
				WithPerformanceMetrics: nimblego.Bool(true),
			},
			RequestTimeout:                nimblego.Float(30000),
			ReturnResponseHeadersAsHeader: nimblego.Bool(true),
			SaveUserbrowser:               nimblego.Bool(false),
			Session: nimblego.CrawlRootParamsExtractOptionsSession{
				ID:                  nimblego.String("id"),
				PrefetchUserbrowser: nimblego.Bool(true),
				Retry:               nimblego.Bool(true),
				Timeout:             nimblego.Float(1),
			},
			Skill: nimblego.CrawlRootParamsExtractOptionsSkillUnion{
				OfString: nimblego.String("dynamic-content"),
			},
			SkipUbct: nimblego.Bool(false),
			State:    "CA",
			Tag:      nimblego.String("campaign-2024-q1"),
			Template: nimblego.CrawlRootParamsExtractOptionsTemplate{
				Name: "x",
				Params: map[string]any{
					"foo": "bar",
				},
			},
			Type: nimblego.String("generic"),
			UserbrowserCreationTemplateRendered: nimblego.CrawlRootParamsExtractOptionsUserbrowserCreationTemplateRendered{
				ID:                    "id",
				AllowedParameterNames: []string{"x"},
				RenderFlowRendered: []map[string]any{{
					"foo": "bar",
				}},
			},
		},
		IgnoreQueryParameters: nimblego.Bool(false),
		IncludePaths:          []string{"/include-this-path", "/and-this-path"},
		Limit:                 nimblego.Int(100),
		MaxDiscoveryDepth:     nimblego.Int(3),
		Name:                  nimblego.String("The best crawl ever"),
		Sitemap:               nimblego.CrawlRootParamsSitemapInclude,
	})
	if err != nil {
		var apierr *nimblego.Error
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
	client := nimblego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Crawl.Status(context.TODO(), "123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		var apierr *nimblego.Error
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
	client := nimblego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Crawl.Terminate(context.TODO(), "123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		var apierr *nimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
