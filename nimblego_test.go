// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nimblego_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/nimbleway-go"
	"github.com/stainless-sdks/nimbleway-go/internal/testutil"
	"github.com/stainless-sdks/nimbleway-go/option"
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
	client := nimblego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Extract(context.TODO(), nimblego.ExtractParams{
		DebugOptions: nimblego.ExtractParamsDebugOptions{
			CollectHar: nimblego.ExtractParamsDebugOptionsCollectHarUnion{
				OfBool: nimblego.Bool(true),
			},
			NoRetryMode: nimblego.ExtractParamsDebugOptionsNoRetryModeUnion{
				OfBool: nimblego.Bool(true),
			},
			RecordScreen: nimblego.ExtractParamsDebugOptionsRecordScreenUnion{
				OfBool: nimblego.Bool(true),
			},
			Redact: nimblego.ExtractParamsDebugOptionsRedactUnion{
				OfBool: nimblego.Bool(true),
			},
			ShowCursor: nimblego.ExtractParamsDebugOptionsShowCursorUnion{
				OfBool: nimblego.Bool(true),
			},
			SolveCaptcha: nimblego.ExtractParamsDebugOptionsSolveCaptchaUnion{
				OfBool: nimblego.Bool(true),
			},
			Trace: nimblego.ExtractParamsDebugOptionsTraceUnion{
				OfBool: nimblego.Bool(true),
			},
			UploadEngineLogs: nimblego.ExtractParamsDebugOptionsUploadEngineLogsUnion{
				OfBool: nimblego.Bool(true),
			},
			Verbose: nimblego.ExtractParamsDebugOptionsVerboseUnion{
				OfBool: nimblego.Bool(true),
			},
			WithProxyUsage: nimblego.ExtractParamsDebugOptionsWithProxyUsageUnion{
				OfBool: nimblego.Bool(true),
			},
		},
		URL: "https://example.com/page",
		Browser: nimblego.ExtractParamsBrowserUnion{
			OfExtractsBrowserString: nimblego.String("chrome"),
		},
		City:          nimblego.String("Los Angeles"),
		ClientTimeout: nimblego.Float(25000),
		ConsentHeader: nimblego.Bool(true),
		Cookies: nimblego.ExtractParamsCookiesUnion{
			OfExtractsCookiesArray: []nimblego.ExtractParamsCookiesArrayItem{{
				Creation:     nimblego.String("creation"),
				Domain:       nimblego.String("domain"),
				Expires:      nimblego.String("expires"),
				Extensions:   []string{"string"},
				HostOnly:     nimblego.Bool(true),
				HTTPOnly:     nimblego.Bool(true),
				LastAccessed: nimblego.String("lastAccessed"),
				MaxAge: nimblego.ExtractParamsCookiesArrayItemMaxAgeUnion{
					OfExtractsCookiesArrayItemMaxAgeString: nimblego.Opt(nimblego.ExtractParamsCookiesArrayItemMaxAgeStringInfinity),
				},
				Name:          nimblego.String("name"),
				Path:          nimblego.String("path"),
				PathIsDefault: nimblego.Bool(true),
				SameSite:      "strict",
				Secure:        nimblego.Bool(true),
				Value:         nimblego.String("value"),
			}},
		},
		Country:        nimblego.ExtractParamsCountryUs,
		Device:         nimblego.ExtractParamsDeviceDesktop,
		DisableIPCheck: nimblego.Bool(false),
		Driver:         nimblego.ExtractParamsDriverVx8,
		DynamicParser: map[string]any{
			"myParser": "bar",
		},
		ExpectedStatusCodes: []int64{200, 201},
		ExportUserbrowser:   nimblego.Bool(false),
		Format:              nimblego.ExtractParamsFormatJson,
		Headers: map[string]nimblego.ExtractParamsHeaderUnion{
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
		Locale:   nimblego.ExtractParamsLocaleEnUs,
		Markdown: nimblego.Bool(false),
		Metadata: nimblego.ExtractParamsMetadata{
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
		Method:     nimblego.ExtractParamsMethodGet,
		NativeMode: nimblego.ExtractParamsNativeModeRequester,
		NetworkCapture: []nimblego.ExtractParamsNetworkCapture{{
			Method: "GET",
			ResourceType: nimblego.ExtractParamsNetworkCaptureResourceTypeUnion{
				OfString: nimblego.String("document"),
			},
			StatusCode: nimblego.ExtractParamsNetworkCaptureStatusCodeUnion{
				OfFloat: nimblego.Float(100),
			},
			URL: nimblego.ExtractParamsNetworkCaptureURL{
				Value: "value",
				Type:  "exact",
			},
			Validation:                  nimblego.Bool(true),
			WaitForRequestsCount:        nimblego.Float(0),
			WaitForRequestsCountTimeout: nimblego.Float(1),
		}},
		NoHTML:        nimblego.Bool(false),
		NoUserbrowser: nimblego.Bool(false),
		Os:            nimblego.ExtractParamsOsWindows,
		Parse:         nimblego.Bool(true),
		ParseOptions: nimblego.ExtractParamsParseOptions{
			MergeDynamic: nimblego.Bool(true),
		},
		Parser: nimblego.ExtractParamsParserUnion{
			OfAnyMap: map[string]any{
				"myParser": "bar",
			},
		},
		ProxyProvider: nimblego.ExtractParamsProxyProviderBrightdata,
		ProxyProviders: map[string]float64{
			"brightdata": 70,
			"oxylabs":    30,
		},
		QueryTemplate: nimblego.ExtractParamsQueryTemplate{
			ID:      "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			APIType: "WEB",
			Pagination: nimblego.ExtractParamsQueryTemplatePaginationUnion{
				OfExtractsQueryTemplatePaginationNextPageParams: &nimblego.ExtractParamsQueryTemplatePaginationNextPageParams{
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
		ReferrerType: nimblego.ExtractParamsReferrerTypeRandom,
		Render:       nimblego.Bool(true),
		RenderFlow: []map[string]any{{
			"wait": "bar",
		}, {
			"click": "bar",
		}},
		RenderOptions: nimblego.ExtractParamsRenderOptions{
			Adblock:        nimblego.Bool(true),
			BlockedDomains: []string{"ads.example.com", "tracker.com"},
			BrowserEngine: nimblego.ExtractParamsRenderOptionsBrowserEngineUnion{
				OfExtractsRenderOptionsBrowserEngineString: nimblego.String("chrome"),
			},
			Cache:             nimblego.Bool(false),
			ConnectorType:     "webit-cdp",
			DisabledResources: []string{"image", "stylesheet"},
			Enable2captcha:    nimblego.Bool(true),
			Extensions:        []string{"extension-id-1", "extension-id-2"},
			FingerprintID:     nimblego.String("fp-abc123"),
			HackiumConfiguration: nimblego.ExtractParamsRenderOptionsHackiumConfiguration{
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
		Session: nimblego.ExtractParamsSession{
			ID:                  nimblego.String("id"),
			PrefetchUserbrowser: nimblego.Bool(true),
			Retry:               nimblego.Bool(true),
			Timeout:             nimblego.Float(1),
		},
		Skill: nimblego.ExtractParamsSkillUnion{
			OfString: nimblego.String("dynamic-content"),
		},
		SkipUbct: nimblego.Bool(false),
		State:    nimblego.ExtractParamsStateCa,
		Tag:      nimblego.String("campaign-2024-q1"),
		Template: nimblego.ExtractParamsTemplate{
			Name: "x",
			Params: map[string]any{
				"foo": "bar",
			},
		},
		Type: nimblego.String("generic"),
		UserbrowserCreationTemplateRendered: nimblego.ExtractParamsUserbrowserCreationTemplateRendered{
			ID:                    "id",
			AllowedParameterNames: []string{"x"},
			RenderFlowRendered: []map[string]any{{
				"foo": "bar",
			}},
		},
	})
	if err != nil {
		var apierr *nimblego.Error
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
	client := nimblego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.ExtractTemplate(context.TODO(), nimblego.ExtractTemplateParams{
		Params: map[string]any{
			"foo": "bar",
		},
		Template: "template",
	})
	if err != nil {
		var apierr *nimblego.Error
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
	client := nimblego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Map(context.TODO(), nimblego.MapParams{
		URL:          "https://example.com",
		Country:      nimblego.MapParamsCountryUs,
		DomainFilter: nimblego.MapParamsDomainFilterAll,
		Limit:        nimblego.Int(1000),
		Locale:       nimblego.MapParamsLocaleEnUs,
		Sitemap:      nimblego.MapParamsSitemapInclude,
	})
	if err != nil {
		var apierr *nimblego.Error
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
	client := nimblego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Search(context.TODO(), nimblego.SearchParams{
		Query:          "x",
		ContentType:    []string{"string"},
		Country:        nimblego.String("country"),
		DeepSearch:     nimblego.Bool(true),
		EndDate:        nimblego.String("end_date"),
		ExcludeDomains: []string{"string"},
		IncludeAnswer:  nimblego.Bool(true),
		IncludeDomains: []string{"string"},
		Locale:         nimblego.String("locale"),
		MaxSubagents:   nimblego.Int(1),
		NumResults:     nimblego.Int(1),
		ParsingType:    nimblego.SearchParamsParsingTypePlainText,
		SearchEngine:   nimblego.SearchParamsSearchEngineGoogleSearch,
		StartDate:      nimblego.String("start_date"),
		TimeRange:      nimblego.SearchParamsTimeRangeHour,
		Topic:          nimblego.SearchParamsTopicGeneral,
	})
	if err != nil {
		var apierr *nimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
