// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nimbleway_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/nimbleway-go"
	"github.com/stainless-sdks/nimbleway-go/internal/testutil"
	"github.com/stainless-sdks/nimbleway-go/option"
)

func TestCrawlRootWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	_, err := client.Crawl.Root(context.TODO(), nimbleway.CrawlRootParams{
		URL:                "https://example.com",
		AllowExternalLinks: nimbleway.Bool(false),
		AllowSubdomains:    nimbleway.Bool(false),
		Callback: nimbleway.CrawlRootParamsCallbackUnion{
			OfCrawlRootsCallbackObject: &nimbleway.CrawlRootParamsCallbackObject{
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
		CrawlEntireDomain: nimbleway.Bool(false),
		ExcludePaths:      []string{"/exclude-this-path", "/and-this-path"},
		ExtractOptions: nimbleway.CrawlRootParamsExtractOptions{
			DebugOptions: nimbleway.CrawlRootParamsExtractOptionsDebugOptions{
				CollectHar: nimbleway.CrawlRootParamsExtractOptionsDebugOptionsCollectHarUnion{
					OfBool: nimbleway.Bool(true),
				},
				NoRetryMode: nimbleway.CrawlRootParamsExtractOptionsDebugOptionsNoRetryModeUnion{
					OfBool: nimbleway.Bool(true),
				},
				RecordScreen: nimbleway.CrawlRootParamsExtractOptionsDebugOptionsRecordScreenUnion{
					OfBool: nimbleway.Bool(true),
				},
				Redact: nimbleway.CrawlRootParamsExtractOptionsDebugOptionsRedactUnion{
					OfBool: nimbleway.Bool(true),
				},
				ShowCursor: nimbleway.CrawlRootParamsExtractOptionsDebugOptionsShowCursorUnion{
					OfBool: nimbleway.Bool(true),
				},
				SolveCaptcha: nimbleway.CrawlRootParamsExtractOptionsDebugOptionsSolveCaptchaUnion{
					OfBool: nimbleway.Bool(true),
				},
				Trace: nimbleway.CrawlRootParamsExtractOptionsDebugOptionsTraceUnion{
					OfBool: nimbleway.Bool(true),
				},
				UploadEngineLogs: nimbleway.CrawlRootParamsExtractOptionsDebugOptionsUploadEngineLogsUnion{
					OfBool: nimbleway.Bool(true),
				},
				Verbose: nimbleway.CrawlRootParamsExtractOptionsDebugOptionsVerboseUnion{
					OfBool: nimbleway.Bool(true),
				},
				WithProxyUsage: nimbleway.CrawlRootParamsExtractOptionsDebugOptionsWithProxyUsageUnion{
					OfBool: nimbleway.Bool(true),
				},
			},
			URL: "https://example.com/page",
			Browser: nimbleway.CrawlRootParamsExtractOptionsBrowserUnion{
				OfCrawlRootsExtractOptionsBrowserString: nimbleway.String("chrome"),
			},
			City:          nimbleway.String("Los Angeles"),
			ClientTimeout: nimbleway.Float(25000),
			ConsentHeader: nimbleway.Bool(true),
			Cookies: nimbleway.CrawlRootParamsExtractOptionsCookiesUnion{
				OfCrawlRootsExtractOptionsCookiesArray: []nimbleway.CrawlRootParamsExtractOptionsCookiesArrayItem{{
					Creation:     nimbleway.String("creation"),
					Domain:       nimbleway.String("domain"),
					Expires:      nimbleway.String("expires"),
					Extensions:   []string{"string"},
					HostOnly:     nimbleway.Bool(true),
					HTTPOnly:     nimbleway.Bool(true),
					LastAccessed: nimbleway.String("lastAccessed"),
					MaxAge: nimbleway.CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeUnion{
						OfCrawlRootsExtractOptionsCookiesArrayItemMaxAgeString: nimbleway.Opt(nimbleway.CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeStringInfinity),
					},
					Name:          nimbleway.String("name"),
					Path:          nimbleway.String("path"),
					PathIsDefault: nimbleway.Bool(true),
					SameSite:      "strict",
					Secure:        nimbleway.Bool(true),
					Value:         nimbleway.String("value"),
				}},
			},
			Country:        nimbleway.CrawlRootParamsExtractOptionsCountryUs,
			Device:         "desktop",
			DisableIPCheck: nimbleway.Bool(false),
			Driver:         "vx8",
			DynamicParser: map[string]any{
				"myParser": "bar",
			},
			ExpectedStatusCodes: []int64{200, 201},
			ExportUserbrowser:   nimbleway.Bool(false),
			Format:              "json",
			Headers: map[string]nimbleway.CrawlRootParamsExtractOptionsHeaderUnion{
				"User-Agent": {
					OfString: nimbleway.String("CustomBot/1.0"),
				},
				"Accept-Language": {
					OfString: nimbleway.String("en-US"),
				},
			},
			Http2:    nimbleway.Bool(true),
			Ip6:      nimbleway.Bool(false),
			IsXhr:    nimbleway.Bool(true),
			Locale:   nimbleway.CrawlRootParamsExtractOptionsLocaleEnUs,
			Markdown: nimbleway.Bool(false),
			Metadata: nimbleway.CrawlRootParamsExtractOptionsMetadata{
				AccountName:         nimbleway.String("acme-corp"),
				DefinitionID:        nimbleway.Int(456),
				DefinitionName:      nimbleway.String("product-scraper"),
				Endpoint:            nimbleway.String("/api/v2/scrape"),
				ExecutionID:         nimbleway.String("exec-abc123"),
				FlowitTaskID:        nimbleway.String("task-xyz789"),
				InputID:             nimbleway.String("input-123"),
				PipelineExecutionID: nimbleway.Int(12345),
				QueryTemplateID:     nimbleway.String("template-qry-001"),
				Source:              nimbleway.String("web-app"),
				TemplateID:          nimbleway.Int(789),
				TemplateName:        nimbleway.String("e-commerce-template"),
			},
			Method:     "GET",
			NativeMode: "requester",
			NetworkCapture: []nimbleway.CrawlRootParamsExtractOptionsNetworkCapture{{
				Method: "GET",
				ResourceType: nimbleway.CrawlRootParamsExtractOptionsNetworkCaptureResourceTypeUnion{
					OfString: nimbleway.String("document"),
				},
				StatusCode: nimbleway.CrawlRootParamsExtractOptionsNetworkCaptureStatusCodeUnion{
					OfFloat: nimbleway.Float(100),
				},
				URL: nimbleway.CrawlRootParamsExtractOptionsNetworkCaptureURL{
					Value: "value",
					Type:  "exact",
				},
				Validation:                  nimbleway.Bool(true),
				WaitForRequestsCount:        nimbleway.Float(0),
				WaitForRequestsCountTimeout: nimbleway.Float(1),
			}},
			NoHTML:        nimbleway.Bool(false),
			NoUserbrowser: nimbleway.Bool(false),
			Os:            "windows",
			Parse:         nimbleway.Bool(true),
			ParseOptions: nimbleway.CrawlRootParamsExtractOptionsParseOptions{
				MergeDynamic: nimbleway.Bool(true),
			},
			Parser: nimbleway.CrawlRootParamsExtractOptionsParserUnion{
				OfAnyMap: map[string]any{
					"myParser": "bar",
				},
			},
			ProxyProvider: nimbleway.CrawlRootParamsExtractOptionsProxyProviderBrightdata,
			ProxyProviders: map[string]float64{
				"brightdata": 70,
				"oxylabs":    30,
			},
			QueryTemplate: nimbleway.CrawlRootParamsExtractOptionsQueryTemplate{
				ID:      "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				APIType: "WEB",
				Pagination: nimbleway.CrawlRootParamsExtractOptionsQueryTemplatePaginationUnion{
					OfCrawlRootsExtractOptionsQueryTemplatePaginationNextPageParams: &nimbleway.CrawlRootParamsExtractOptionsQueryTemplatePaginationNextPageParams{
						NextPageParams: map[string]any{
							"foo": "bar",
						},
					},
				},
				Params: map[string]any{
					"foo": "bar",
				},
			},
			RawHeaders:   nimbleway.Bool(true),
			ReferrerType: nimbleway.CrawlRootParamsExtractOptionsReferrerTypeRandom,
			Render:       nimbleway.Bool(true),
			RenderFlow: []map[string]any{{
				"wait": "bar",
			}, {
				"click": "bar",
			}},
			RenderOptions: nimbleway.CrawlRootParamsExtractOptionsRenderOptions{
				Adblock:        nimbleway.Bool(true),
				BlockedDomains: []string{"ads.example.com", "tracker.com"},
				BrowserEngine: nimbleway.CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineUnion{
					OfCrawlRootsExtractOptionsRenderOptionsBrowserEngineString: nimbleway.String("chrome"),
				},
				Cache:             nimbleway.Bool(false),
				ConnectorType:     "webit-cdp",
				DisabledResources: []string{"image", "stylesheet"},
				Enable2captcha:    nimbleway.Bool(true),
				Extensions:        []string{"extension-id-1", "extension-id-2"},
				FingerprintID:     nimbleway.String("fp-abc123"),
				HackiumConfiguration: nimbleway.CrawlRootParamsExtractOptionsRenderOptionsHackiumConfiguration{
					CollectLogs:                 nimbleway.Bool(true),
					DoNotFixMathSalt:            nimbleway.Bool(true),
					EnableDocumentElementSpoof:  nimbleway.Bool(true),
					EnableDocumentHasFocus:      nimbleway.Bool(true),
					EnableFakeNavigationHistory: nimbleway.Bool(true),
					EnableKeyOrdering:           nimbleway.Bool(true),
					EnableSniffer:               nimbleway.Bool(true),
					EnableVerboseLogs:           nimbleway.Bool(true),
				},
				Headless:               nimbleway.Bool(true),
				IncludeIframes:         nimbleway.Bool(true),
				LoadLocalStorage:       nimbleway.Bool(true),
				LocalStorageKeysToLoad: []string{"authToken", "userId"},
				MouseStrategy:          "linear",
				NoAcceptEncoding:       nimbleway.Bool(true),
				OverridePermissions:    nimbleway.Bool(true),
				RandomHeaderOrder:      nimbleway.Bool(true),
				RenderType:             "load",
				StoreLocalStorage:      nimbleway.Bool(true),
				Timeout:                nimbleway.Float(30000),
				TypingInterval:         nimbleway.Float(100),
				TypingStrategy:         "simple",
				Userbrowser:            nimbleway.Bool(true),
				WaitUntil:              "networkidle2",
				WithPerformanceMetrics: nimbleway.Bool(true),
			},
			RequestTimeout:                nimbleway.Float(30000),
			ReturnResponseHeadersAsHeader: nimbleway.Bool(true),
			SaveUserbrowser:               nimbleway.Bool(false),
			Session: nimbleway.CrawlRootParamsExtractOptionsSession{
				ID:                  nimbleway.String("id"),
				PrefetchUserbrowser: nimbleway.Bool(true),
				Retry:               nimbleway.Bool(true),
				Timeout:             nimbleway.Float(1),
			},
			Skill: nimbleway.CrawlRootParamsExtractOptionsSkillUnion{
				OfString: nimbleway.String("dynamic-content"),
			},
			SkipUbct: nimbleway.Bool(false),
			State:    "CA",
			Tag:      nimbleway.String("campaign-2024-q1"),
			Template: nimbleway.CrawlRootParamsExtractOptionsTemplate{
				Name: "x",
				Params: map[string]any{
					"foo": "bar",
				},
			},
			Type: nimbleway.String("generic"),
			UserbrowserCreationTemplateRendered: nimbleway.CrawlRootParamsExtractOptionsUserbrowserCreationTemplateRendered{
				ID:                    "id",
				AllowedParameterNames: []string{"x"},
				RenderFlowRendered: []map[string]any{{
					"foo": "bar",
				}},
			},
		},
		IgnoreQueryParameters: nimbleway.Bool(false),
		IncludePaths:          []string{"/include-this-path", "/and-this-path"},
		Limit:                 nimbleway.Int(100),
		MaxDiscoveryDepth:     nimbleway.Int(3),
		Name:                  nimbleway.String("The best crawl ever"),
		Sitemap:               nimbleway.CrawlRootParamsSitemapInclude,
	})
	if err != nil {
		var apierr *nimbleway.Error
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
	client := nimbleway.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Crawl.Status(context.TODO(), "123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		var apierr *nimbleway.Error
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
	client := nimbleway.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Crawl.Terminate(context.TODO(), "123e4567-e89b-12d3-a456-426614174000")
	if err != nil {
		var apierr *nimbleway.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
