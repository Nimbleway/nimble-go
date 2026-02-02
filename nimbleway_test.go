// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nimbleway_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Nimbleway/nimbleway-go"
	"github.com/Nimbleway/nimbleway-go/internal/testutil"
	"github.com/Nimbleway/nimbleway-go/option"
)

func TestExtractWithOptionalParams(t *testing.T) {
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
	_, err := client.Extract(context.TODO(), nimbleway.ExtractParams{
		DebugOptions: nimbleway.ExtractParamsDebugOptions{
			CollectHar: nimbleway.ExtractParamsDebugOptionsCollectHarUnion{
				OfBool: nimbleway.Bool(true),
			},
			NoRetryMode: nimbleway.ExtractParamsDebugOptionsNoRetryModeUnion{
				OfBool: nimbleway.Bool(true),
			},
			RecordScreen: nimbleway.ExtractParamsDebugOptionsRecordScreenUnion{
				OfBool: nimbleway.Bool(true),
			},
			Redact: nimbleway.ExtractParamsDebugOptionsRedactUnion{
				OfBool: nimbleway.Bool(true),
			},
			ShowCursor: nimbleway.ExtractParamsDebugOptionsShowCursorUnion{
				OfBool: nimbleway.Bool(true),
			},
			SolveCaptcha: nimbleway.ExtractParamsDebugOptionsSolveCaptchaUnion{
				OfBool: nimbleway.Bool(true),
			},
			Trace: nimbleway.ExtractParamsDebugOptionsTraceUnion{
				OfBool: nimbleway.Bool(true),
			},
			UploadEngineLogs: nimbleway.ExtractParamsDebugOptionsUploadEngineLogsUnion{
				OfBool: nimbleway.Bool(true),
			},
			Verbose: nimbleway.ExtractParamsDebugOptionsVerboseUnion{
				OfBool: nimbleway.Bool(true),
			},
			WithProxyUsage: nimbleway.ExtractParamsDebugOptionsWithProxyUsageUnion{
				OfBool: nimbleway.Bool(true),
			},
		},
		URL: "https://example.com/page",
		Browser: nimbleway.ExtractParamsBrowserUnion{
			OfExtractsBrowserString: nimbleway.String("chrome"),
		},
		City:          nimbleway.String("Los Angeles"),
		ClientTimeout: nimbleway.Float(25000),
		ConsentHeader: nimbleway.Bool(true),
		Cookies: nimbleway.ExtractParamsCookiesUnion{
			OfExtractsCookiesArray: []nimbleway.ExtractParamsCookiesArrayItem{{
				Creation:     nimbleway.String("creation"),
				Domain:       nimbleway.String("domain"),
				Expires:      nimbleway.String("expires"),
				Extensions:   []string{"string"},
				HostOnly:     nimbleway.Bool(true),
				HTTPOnly:     nimbleway.Bool(true),
				LastAccessed: nimbleway.String("lastAccessed"),
				MaxAge: nimbleway.ExtractParamsCookiesArrayItemMaxAgeUnion{
					OfExtractsCookiesArrayItemMaxAgeString: nimbleway.Opt(nimbleway.ExtractParamsCookiesArrayItemMaxAgeStringInfinity),
				},
				Name:          nimbleway.String("name"),
				Path:          nimbleway.String("path"),
				PathIsDefault: nimbleway.Bool(true),
				SameSite:      "strict",
				Secure:        nimbleway.Bool(true),
				Value:         nimbleway.String("value"),
			}},
		},
		Country:        nimbleway.ExtractParamsCountryUs,
		Device:         nimbleway.ExtractParamsDeviceDesktop,
		DisableIPCheck: nimbleway.Bool(false),
		Driver:         nimbleway.ExtractParamsDriverVx8,
		DynamicParser: map[string]any{
			"myParser": "bar",
		},
		ExpectedStatusCodes: []int64{200, 201},
		ExportUserbrowser:   nimbleway.Bool(false),
		Format:              nimbleway.ExtractParamsFormatJson,
		Headers: map[string]nimbleway.ExtractParamsHeaderUnion{
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
		Locale:   nimbleway.ExtractParamsLocaleEnUs,
		Markdown: nimbleway.Bool(false),
		Metadata: nimbleway.ExtractParamsMetadata{
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
		Method:     nimbleway.ExtractParamsMethodGet,
		NativeMode: nimbleway.ExtractParamsNativeModeRequester,
		NetworkCapture: []nimbleway.ExtractParamsNetworkCapture{{
			Method: "GET",
			ResourceType: nimbleway.ExtractParamsNetworkCaptureResourceTypeUnion{
				OfString: nimbleway.String("document"),
			},
			StatusCode: nimbleway.ExtractParamsNetworkCaptureStatusCodeUnion{
				OfFloat: nimbleway.Float(100),
			},
			URL: nimbleway.ExtractParamsNetworkCaptureURL{
				Value: "value",
				Type:  "exact",
			},
			Validation:                  nimbleway.Bool(true),
			WaitForRequestsCount:        nimbleway.Float(0),
			WaitForRequestsCountTimeout: nimbleway.Float(1),
		}},
		NoHTML:        nimbleway.Bool(false),
		NoUserbrowser: nimbleway.Bool(false),
		Os:            nimbleway.ExtractParamsOsWindows,
		Parse:         nimbleway.Bool(true),
		ParseOptions: nimbleway.ExtractParamsParseOptions{
			MergeDynamic: nimbleway.Bool(true),
		},
		Parser: nimbleway.ExtractParamsParserUnion{
			OfAnyMap: map[string]any{
				"myParser": "bar",
			},
		},
		ProxyProvider: nimbleway.ExtractParamsProxyProviderBrightdata,
		ProxyProviders: map[string]float64{
			"brightdata": 70,
			"oxylabs":    30,
		},
		QueryTemplate: nimbleway.ExtractParamsQueryTemplate{
			ID:      "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			APIType: "WEB",
			Pagination: nimbleway.ExtractParamsQueryTemplatePaginationUnion{
				OfExtractsQueryTemplatePaginationNextPageParams: &nimbleway.ExtractParamsQueryTemplatePaginationNextPageParams{
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
		ReferrerType: nimbleway.ExtractParamsReferrerTypeRandom,
		Render:       nimbleway.Bool(true),
		RenderFlow: []map[string]any{{
			"wait": "bar",
		}, {
			"click": "bar",
		}},
		RenderOptions: nimbleway.ExtractParamsRenderOptions{
			Adblock:        nimbleway.Bool(true),
			BlockedDomains: []string{"ads.example.com", "tracker.com"},
			BrowserEngine: nimbleway.ExtractParamsRenderOptionsBrowserEngineUnion{
				OfExtractsRenderOptionsBrowserEngineString: nimbleway.String("chrome"),
			},
			Cache:             nimbleway.Bool(false),
			ConnectorType:     "webit-cdp",
			DisabledResources: []string{"image", "stylesheet"},
			Enable2captcha:    nimbleway.Bool(true),
			Extensions:        []string{"extension-id-1", "extension-id-2"},
			FingerprintID:     nimbleway.String("fp-abc123"),
			HackiumConfiguration: nimbleway.ExtractParamsRenderOptionsHackiumConfiguration{
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
		Session: nimbleway.ExtractParamsSession{
			ID:                  nimbleway.String("id"),
			PrefetchUserbrowser: nimbleway.Bool(true),
			Retry:               nimbleway.Bool(true),
			Timeout:             nimbleway.Float(1),
		},
		Skill: nimbleway.ExtractParamsSkillUnion{
			OfString: nimbleway.String("dynamic-content"),
		},
		SkipUbct: nimbleway.Bool(false),
		State:    nimbleway.ExtractParamsStateCa,
		Tag:      nimbleway.String("campaign-2024-q1"),
		Template: nimbleway.ExtractParamsTemplate{
			Name: "x",
			Params: map[string]any{
				"foo": "bar",
			},
		},
		Type: nimbleway.String("generic"),
		UserbrowserCreationTemplateRendered: nimbleway.ExtractParamsUserbrowserCreationTemplateRendered{
			ID:                    "id",
			AllowedParameterNames: []string{"x"},
			RenderFlowRendered: []map[string]any{{
				"foo": "bar",
			}},
		},
	})
	if err != nil {
		var apierr *nimbleway.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtractTemplate(t *testing.T) {
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
	_, err := client.ExtractTemplate(context.TODO(), nimbleway.ExtractTemplateParams{
		Params: map[string]any{
			"foo": "bar",
		},
		Template: "template",
	})
	if err != nil {
		var apierr *nimbleway.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMapWithOptionalParams(t *testing.T) {
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
	_, err := client.Map(context.TODO(), nimbleway.MapParams{
		URL:          "https://example.com",
		Country:      nimbleway.MapParamsCountryUs,
		DomainFilter: nimbleway.MapParamsDomainFilterAll,
		Limit:        nimbleway.Int(1000),
		Locale:       nimbleway.MapParamsLocaleEnUs,
		Sitemap:      nimbleway.MapParamsSitemapInclude,
	})
	if err != nil {
		var apierr *nimbleway.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSearchWithOptionalParams(t *testing.T) {
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
	_, err := client.Search(context.TODO(), nimbleway.SearchParams{
		Query:          "x",
		ContentType:    []string{"string"},
		Country:        nimbleway.String("country"),
		DeepSearch:     nimbleway.Bool(true),
		EndDate:        nimbleway.String("end_date"),
		ExcludeDomains: []string{"string"},
		IncludeAnswer:  nimbleway.Bool(true),
		IncludeDomains: []string{"string"},
		Locale:         nimbleway.String("locale"),
		MaxSubagents:   nimbleway.Int(1),
		NumResults:     nimbleway.Int(1),
		ParsingType:    nimbleway.SearchParamsParsingTypePlainText,
		SearchEngine:   nimbleway.SearchParamsSearchEngineGoogleSearch,
		StartDate:      nimbleway.String("start_date"),
		TimeRange:      nimbleway.SearchParamsTimeRangeHour,
		Topic:          nimbleway.SearchParamsTopicGeneral,
	})
	if err != nil {
		var apierr *nimbleway.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
