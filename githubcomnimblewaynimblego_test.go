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

func TestAgentWithOptionalParams(t *testing.T) {
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
	_, err := client.Agent(context.TODO(), githubcomnimblewaynimblego.AgentParams{
		OfExtractTemplateBody: &githubcomnimblewaynimblego.AgentParamsBodyExtractTemplateBody{
			Params: map[string]any{
				"foo": "bar",
			},
			Template:     "template",
			Localization: githubcomnimblewaynimblego.Bool(true),
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

func TestCrawlWithOptionalParams(t *testing.T) {
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
	_, err := client.Crawl(context.TODO(), githubcomnimblewaynimblego.CrawlParams{
		URL:                "url",
		AllowExternalLinks: githubcomnimblewaynimblego.Bool(false),
		AllowSubdomains:    githubcomnimblewaynimblego.Bool(false),
		Callback: githubcomnimblewaynimblego.CrawlParamsCallbackUnion{
			OfCrawlsCallbackObject: &githubcomnimblewaynimblego.CrawlParamsCallbackObject{
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
		ExtractOptions: githubcomnimblewaynimblego.CrawlParamsExtractOptions{
			Browser: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserUnion{
				OfCrawlsExtractOptionsBrowserString: githubcomnimblewaynimblego.String("chrome"),
			},
			BrowserActions: []githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionUnion{{
				OfCrawlsExtractOptionsBrowserActionGotoAction: &githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionGotoAction{
					Goto: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionGotoActionGotoUnion{
						OfString: githubcomnimblewaynimblego.String("https://example.com/login"),
					},
				},
			}, {
				OfCrawlsExtractOptionsBrowserActionWaitForElementAction: &githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionWaitForElementAction{
					WaitForElement: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementUnion{
						OfString: githubcomnimblewaynimblego.String("#login-form"),
					},
				},
			}, {
				OfCrawlsExtractOptionsBrowserActionFillAction: &githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillAction{
					Fill: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillUnion{
						OfType: &githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillType{
							Selector: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSelectorUnion{
								OfString: githubcomnimblewaynimblego.String("#username"),
							},
							Value:          "user@example.com",
							ClickOnElement: githubcomnimblewaynimblego.Bool(true),
							Delay: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillTypeDelayUnion{
								OfFloat: githubcomnimblewaynimblego.Float(1000),
							},
							Mode:                  "type",
							MouseMovementStrategy: "linear",
							Required: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillTypeRequiredUnion{
								OfCrawlsExtractOptionsBrowserActionFillActionFillTypeRequiredString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillTypeRequiredStringTrue),
							},
							Scroll: githubcomnimblewaynimblego.Bool(true),
							Skip: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSkipUnion{
								OfCrawlsExtractOptionsBrowserActionFillActionFillTypeSkipString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSkipStringTrue),
							},
							Timeout: githubcomnimblewaynimblego.Float(0),
							TypingInterval: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillTypeTypingIntervalUnion{
								OfFloat: githubcomnimblewaynimblego.Float(1000),
							},
							TypingStrategy: "simple",
							Visible:        githubcomnimblewaynimblego.Bool(true),
						},
					},
				},
			}, {
				OfCrawlsExtractOptionsBrowserActionFillAction: &githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillAction{
					Fill: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillUnion{
						OfType: &githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillType{
							Selector: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSelectorUnion{
								OfString: githubcomnimblewaynimblego.String("#password"),
							},
							Value:          "password123",
							ClickOnElement: githubcomnimblewaynimblego.Bool(true),
							Delay: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillTypeDelayUnion{
								OfFloat: githubcomnimblewaynimblego.Float(1000),
							},
							Mode:                  "type",
							MouseMovementStrategy: "linear",
							Required: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillTypeRequiredUnion{
								OfCrawlsExtractOptionsBrowserActionFillActionFillTypeRequiredString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillTypeRequiredStringTrue),
							},
							Scroll: githubcomnimblewaynimblego.Bool(true),
							Skip: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSkipUnion{
								OfCrawlsExtractOptionsBrowserActionFillActionFillTypeSkipString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSkipStringTrue),
							},
							Timeout: githubcomnimblewaynimblego.Float(0),
							TypingInterval: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionFillActionFillTypeTypingIntervalUnion{
								OfFloat: githubcomnimblewaynimblego.Float(1000),
							},
							TypingStrategy: "simple",
							Visible:        githubcomnimblewaynimblego.Bool(true),
						},
					},
				},
			}, {
				OfCrawlsExtractOptionsBrowserActionClickAction: &githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionClickAction{
					Click: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionClickActionClickUnion{
						OfString: githubcomnimblewaynimblego.String("#submit"),
					},
				},
			}, {
				OfCrawlsExtractOptionsBrowserActionScreenshotAction: &githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionScreenshotAction{
					Screenshot: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotUnion{
						OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObject: &githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObject{
							Format:   "png",
							FullPage: githubcomnimblewaynimblego.Bool(true),
							Quality:  githubcomnimblewaynimblego.Float(0),
							Required: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredUnion{
								OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredStringTrue),
							},
							Skip: githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipUnion{
								OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipStringTrue),
							},
						},
					},
				},
			}},
			City:          githubcomnimblewaynimblego.String("Los Angeles"),
			ClientTimeout: githubcomnimblewaynimblego.Float(25000),
			ConsentHeader: githubcomnimblewaynimblego.Bool(true),
			Cookies: githubcomnimblewaynimblego.CrawlParamsExtractOptionsCookiesUnion{
				OfCrawlsExtractOptionsCookiesArray: []githubcomnimblewaynimblego.CrawlParamsExtractOptionsCookiesArrayItem{{
					Creation:     githubcomnimblewaynimblego.String("creation"),
					Domain:       githubcomnimblewaynimblego.String("domain"),
					Expires:      githubcomnimblewaynimblego.String("expires"),
					Extensions:   []string{"string"},
					HostOnly:     githubcomnimblewaynimblego.Bool(true),
					HTTPOnly:     githubcomnimblewaynimblego.Bool(true),
					LastAccessed: githubcomnimblewaynimblego.String("lastAccessed"),
					MaxAge: githubcomnimblewaynimblego.CrawlParamsExtractOptionsCookiesArrayItemMaxAgeUnion{
						OfCrawlsExtractOptionsCookiesArrayItemMaxAgeString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.CrawlParamsExtractOptionsCookiesArrayItemMaxAgeStringInfinity),
					},
					Name:          githubcomnimblewaynimblego.String("name"),
					Path:          githubcomnimblewaynimblego.String("path"),
					PathIsDefault: githubcomnimblewaynimblego.Bool(true),
					SameSite:      "strict",
					Secure:        githubcomnimblewaynimblego.Bool(true),
					Value:         githubcomnimblewaynimblego.String("value"),
				}},
			},
			Country:             githubcomnimblewaynimblego.CrawlParamsExtractOptionsCountryUs,
			Device:              "desktop",
			DisableIPCheck:      githubcomnimblewaynimblego.Bool(false),
			Driver:              "vx8",
			ExpectedStatusCodes: []int64{200, 201},
			Formats:             []string{"html"},
			Headers: map[string]githubcomnimblewaynimblego.CrawlParamsExtractOptionsHeaderUnion{
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
			Locale: githubcomnimblewaynimblego.CrawlParamsExtractOptionsLocaleEnUs,
			Metadata: githubcomnimblewaynimblego.CrawlParamsExtractOptionsMetadata{
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
			NetworkCapture: []githubcomnimblewaynimblego.CrawlParamsExtractOptionsNetworkCapture{{
				Method: "GET",
				ResourceType: githubcomnimblewaynimblego.CrawlParamsExtractOptionsNetworkCaptureResourceTypeUnion{
					OfString: githubcomnimblewaynimblego.String("document"),
				},
				StatusCode: githubcomnimblewaynimblego.CrawlParamsExtractOptionsNetworkCaptureStatusCodeUnion{
					OfFloat: githubcomnimblewaynimblego.Float(100),
				},
				URL: githubcomnimblewaynimblego.CrawlParamsExtractOptionsNetworkCaptureURL{
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
			Parser: githubcomnimblewaynimblego.CrawlParamsExtractOptionsParserUnion{
				OfAnyMap: map[string]any{
					"myParser": "bar",
				},
			},
			ProxyProvider: githubcomnimblewaynimblego.CrawlParamsExtractOptionsProxyProviderBrightdata,
			ProxyProviders: map[string]float64{
				"brightdata": 70,
				"oxylabs":    30,
			},
			QueryTemplate: githubcomnimblewaynimblego.CrawlParamsExtractOptionsQueryTemplate{
				ID:      "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
				APIType: "WEB",
				Pagination: githubcomnimblewaynimblego.CrawlParamsExtractOptionsQueryTemplatePaginationUnion{
					OfCrawlsExtractOptionsQueryTemplatePaginationNextPageParams: &githubcomnimblewaynimblego.CrawlParamsExtractOptionsQueryTemplatePaginationNextPageParams{
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
			ReferrerType: githubcomnimblewaynimblego.CrawlParamsExtractOptionsReferrerTypeRandom,
			Render:       githubcomnimblewaynimblego.Bool(true),
			RenderFlow: []map[string]any{{
				"wait": "bar",
			}, {
				"click": "bar",
			}},
			RenderOptions: githubcomnimblewaynimblego.CrawlParamsExtractOptionsRenderOptions{
				Adblock:        githubcomnimblewaynimblego.Bool(true),
				BlockedDomains: []string{"ads.example.com", "tracker.com"},
				BrowserEngine: githubcomnimblewaynimblego.CrawlParamsExtractOptionsRenderOptionsBrowserEngineUnion{
					OfCrawlsExtractOptionsRenderOptionsBrowserEngineString: githubcomnimblewaynimblego.String("chrome"),
				},
				Cache:             githubcomnimblewaynimblego.Bool(false),
				ConnectorType:     "puppeteer",
				DisabledResources: []string{"image", "stylesheet"},
				Enable2captcha:    githubcomnimblewaynimblego.Bool(true),
				Extensions:        []string{"extension-id-1", "extension-id-2"},
				FingerprintID:     githubcomnimblewaynimblego.String("fp-abc123"),
				HackiumConfiguration: githubcomnimblewaynimblego.CrawlParamsExtractOptionsRenderOptionsHackiumConfiguration{
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
			Session: githubcomnimblewaynimblego.CrawlParamsExtractOptionsSession{
				ID:                  githubcomnimblewaynimblego.String("id"),
				PrefetchUserbrowser: githubcomnimblewaynimblego.Bool(true),
				Retry:               githubcomnimblewaynimblego.Bool(true),
				Timeout:             githubcomnimblewaynimblego.Float(1),
			},
			Skill: githubcomnimblewaynimblego.CrawlParamsExtractOptionsSkillUnion{
				OfString: githubcomnimblewaynimblego.String("dynamic-content"),
			},
			SkipUbct: githubcomnimblewaynimblego.Bool(false),
			State:    "CA",
			Tag:      githubcomnimblewaynimblego.String("campaign-2024-q1"),
			Template: githubcomnimblewaynimblego.CrawlParamsExtractOptionsTemplate{
				Name: "x",
				Params: map[string]any{
					"foo": "bar",
				},
			},
			Type: githubcomnimblewaynimblego.String("generic"),
			URL:  githubcomnimblewaynimblego.String("https://example.com/page"),
			UserbrowserCreationTemplateRendered: githubcomnimblewaynimblego.CrawlParamsExtractOptionsUserbrowserCreationTemplateRendered{
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
		Sitemap:               githubcomnimblewaynimblego.CrawlParamsSitemapInclude,
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestExtractWithOptionalParams(t *testing.T) {
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
	_, err := client.Extract(context.TODO(), githubcomnimblewaynimblego.ExtractParams{
		URL: "https://example.com/page",
		Browser: githubcomnimblewaynimblego.ExtractParamsBrowserUnion{
			OfExtractsBrowserString: githubcomnimblewaynimblego.String("chrome"),
		},
		BrowserActions: []githubcomnimblewaynimblego.ExtractParamsBrowserActionUnion{{
			OfExtractsBrowserActionGotoAction: &githubcomnimblewaynimblego.ExtractParamsBrowserActionGotoAction{
				Goto: githubcomnimblewaynimblego.ExtractParamsBrowserActionGotoActionGotoUnion{
					OfString: githubcomnimblewaynimblego.String("https://example.com/login"),
				},
			},
		}, {
			OfExtractsBrowserActionWaitForElementAction: &githubcomnimblewaynimblego.ExtractParamsBrowserActionWaitForElementAction{
				WaitForElement: githubcomnimblewaynimblego.ExtractParamsBrowserActionWaitForElementActionWaitForElementUnion{
					OfString: githubcomnimblewaynimblego.String("#login-form"),
				},
			},
		}, {
			OfExtractsBrowserActionFillAction: &githubcomnimblewaynimblego.ExtractParamsBrowserActionFillAction{
				Fill: githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillUnion{
					OfType: &githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillType{
						Selector: githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillTypeSelectorUnion{
							OfString: githubcomnimblewaynimblego.String("#username"),
						},
						Value:          "user@example.com",
						ClickOnElement: githubcomnimblewaynimblego.Bool(true),
						Delay: githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillTypeDelayUnion{
							OfFloat: githubcomnimblewaynimblego.Float(1000),
						},
						Mode:                  "type",
						MouseMovementStrategy: "linear",
						Required: githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillTypeRequiredUnion{
							OfExtractsBrowserActionFillActionFillTypeRequiredString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillTypeRequiredStringTrue),
						},
						Scroll: githubcomnimblewaynimblego.Bool(true),
						Skip: githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillTypeSkipUnion{
							OfExtractsBrowserActionFillActionFillTypeSkipString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillTypeSkipStringTrue),
						},
						Timeout: githubcomnimblewaynimblego.Float(0),
						TypingInterval: githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillTypeTypingIntervalUnion{
							OfFloat: githubcomnimblewaynimblego.Float(1000),
						},
						TypingStrategy: "simple",
						Visible:        githubcomnimblewaynimblego.Bool(true),
					},
				},
			},
		}, {
			OfExtractsBrowserActionFillAction: &githubcomnimblewaynimblego.ExtractParamsBrowserActionFillAction{
				Fill: githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillUnion{
					OfType: &githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillType{
						Selector: githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillTypeSelectorUnion{
							OfString: githubcomnimblewaynimblego.String("#password"),
						},
						Value:          "password123",
						ClickOnElement: githubcomnimblewaynimblego.Bool(true),
						Delay: githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillTypeDelayUnion{
							OfFloat: githubcomnimblewaynimblego.Float(1000),
						},
						Mode:                  "type",
						MouseMovementStrategy: "linear",
						Required: githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillTypeRequiredUnion{
							OfExtractsBrowserActionFillActionFillTypeRequiredString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillTypeRequiredStringTrue),
						},
						Scroll: githubcomnimblewaynimblego.Bool(true),
						Skip: githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillTypeSkipUnion{
							OfExtractsBrowserActionFillActionFillTypeSkipString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillTypeSkipStringTrue),
						},
						Timeout: githubcomnimblewaynimblego.Float(0),
						TypingInterval: githubcomnimblewaynimblego.ExtractParamsBrowserActionFillActionFillTypeTypingIntervalUnion{
							OfFloat: githubcomnimblewaynimblego.Float(1000),
						},
						TypingStrategy: "simple",
						Visible:        githubcomnimblewaynimblego.Bool(true),
					},
				},
			},
		}, {
			OfExtractsBrowserActionClickAction: &githubcomnimblewaynimblego.ExtractParamsBrowserActionClickAction{
				Click: githubcomnimblewaynimblego.ExtractParamsBrowserActionClickActionClickUnion{
					OfString: githubcomnimblewaynimblego.String("#submit"),
				},
			},
		}, {
			OfExtractsBrowserActionScreenshotAction: &githubcomnimblewaynimblego.ExtractParamsBrowserActionScreenshotAction{
				Screenshot: githubcomnimblewaynimblego.ExtractParamsBrowserActionScreenshotActionScreenshotUnion{
					OfExtractsBrowserActionScreenshotActionScreenshotObject: &githubcomnimblewaynimblego.ExtractParamsBrowserActionScreenshotActionScreenshotObject{
						Format:   "png",
						FullPage: githubcomnimblewaynimblego.Bool(true),
						Quality:  githubcomnimblewaynimblego.Float(0),
						Required: githubcomnimblewaynimblego.ExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion{
							OfExtractsBrowserActionScreenshotActionScreenshotObjectRequiredString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredStringTrue),
						},
						Skip: githubcomnimblewaynimblego.ExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion{
							OfExtractsBrowserActionScreenshotActionScreenshotObjectSkipString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipStringTrue),
						},
					},
				},
			},
		}},
		City:          githubcomnimblewaynimblego.String("Los Angeles"),
		ClientTimeout: githubcomnimblewaynimblego.Float(25000),
		ConsentHeader: githubcomnimblewaynimblego.Bool(true),
		Cookies: githubcomnimblewaynimblego.ExtractParamsCookiesUnion{
			OfExtractsCookiesArray: []githubcomnimblewaynimblego.ExtractParamsCookiesArrayItem{{
				Creation:     githubcomnimblewaynimblego.String("creation"),
				Domain:       githubcomnimblewaynimblego.String("domain"),
				Expires:      githubcomnimblewaynimblego.String("expires"),
				Extensions:   []string{"string"},
				HostOnly:     githubcomnimblewaynimblego.Bool(true),
				HTTPOnly:     githubcomnimblewaynimblego.Bool(true),
				LastAccessed: githubcomnimblewaynimblego.String("lastAccessed"),
				MaxAge: githubcomnimblewaynimblego.ExtractParamsCookiesArrayItemMaxAgeUnion{
					OfExtractsCookiesArrayItemMaxAgeString: githubcomnimblewaynimblego.Opt(githubcomnimblewaynimblego.ExtractParamsCookiesArrayItemMaxAgeStringInfinity),
				},
				Name:          githubcomnimblewaynimblego.String("name"),
				Path:          githubcomnimblewaynimblego.String("path"),
				PathIsDefault: githubcomnimblewaynimblego.Bool(true),
				SameSite:      "strict",
				Secure:        githubcomnimblewaynimblego.Bool(true),
				Value:         githubcomnimblewaynimblego.String("value"),
			}},
		},
		Country:             githubcomnimblewaynimblego.ExtractParamsCountryUs,
		Device:              githubcomnimblewaynimblego.ExtractParamsDeviceDesktop,
		DisableIPCheck:      githubcomnimblewaynimblego.Bool(false),
		Driver:              githubcomnimblewaynimblego.ExtractParamsDriverVx8,
		ExpectedStatusCodes: []int64{200, 201},
		Formats:             []string{"html"},
		Headers: map[string]githubcomnimblewaynimblego.ExtractParamsHeaderUnion{
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
		Locale: githubcomnimblewaynimblego.ExtractParamsLocaleEnUs,
		Metadata: githubcomnimblewaynimblego.ExtractParamsMetadata{
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
		Method:     githubcomnimblewaynimblego.ExtractParamsMethodGet,
		NativeMode: githubcomnimblewaynimblego.ExtractParamsNativeModeRequester,
		NetworkCapture: []githubcomnimblewaynimblego.ExtractParamsNetworkCapture{{
			Method: "GET",
			ResourceType: githubcomnimblewaynimblego.ExtractParamsNetworkCaptureResourceTypeUnion{
				OfString: githubcomnimblewaynimblego.String("document"),
			},
			StatusCode: githubcomnimblewaynimblego.ExtractParamsNetworkCaptureStatusCodeUnion{
				OfFloat: githubcomnimblewaynimblego.Float(100),
			},
			URL: githubcomnimblewaynimblego.ExtractParamsNetworkCaptureURL{
				Value: "value",
				Type:  "exact",
			},
			Validation:                  githubcomnimblewaynimblego.Bool(true),
			WaitForRequestsCount:        githubcomnimblewaynimblego.Float(0),
			WaitForRequestsCountTimeout: githubcomnimblewaynimblego.Float(1),
		}},
		NoUserbrowser: githubcomnimblewaynimblego.Bool(false),
		Os:            githubcomnimblewaynimblego.ExtractParamsOsWindows,
		Parse:         githubcomnimblewaynimblego.Bool(true),
		Parser: githubcomnimblewaynimblego.ExtractParamsParserUnion{
			OfAnyMap: map[string]any{
				"myParser": "bar",
			},
		},
		ProxyProvider: githubcomnimblewaynimblego.ExtractParamsProxyProviderBrightdata,
		ProxyProviders: map[string]float64{
			"brightdata": 70,
			"oxylabs":    30,
		},
		QueryTemplate: githubcomnimblewaynimblego.ExtractParamsQueryTemplate{
			ID:      "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
			APIType: "WEB",
			Pagination: githubcomnimblewaynimblego.ExtractParamsQueryTemplatePaginationUnion{
				OfExtractsQueryTemplatePaginationNextPageParams: &githubcomnimblewaynimblego.ExtractParamsQueryTemplatePaginationNextPageParams{
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
		ReferrerType: githubcomnimblewaynimblego.ExtractParamsReferrerTypeRandom,
		Render:       githubcomnimblewaynimblego.Bool(true),
		RenderFlow: []map[string]any{{
			"wait": "bar",
		}, {
			"click": "bar",
		}},
		RenderOptions: githubcomnimblewaynimblego.ExtractParamsRenderOptions{
			Adblock:        githubcomnimblewaynimblego.Bool(true),
			BlockedDomains: []string{"ads.example.com", "tracker.com"},
			BrowserEngine: githubcomnimblewaynimblego.ExtractParamsRenderOptionsBrowserEngineUnion{
				OfExtractsRenderOptionsBrowserEngineString: githubcomnimblewaynimblego.String("chrome"),
			},
			Cache:             githubcomnimblewaynimblego.Bool(false),
			ConnectorType:     "puppeteer",
			DisabledResources: []string{"image", "stylesheet"},
			Enable2captcha:    githubcomnimblewaynimblego.Bool(true),
			Extensions:        []string{"extension-id-1", "extension-id-2"},
			FingerprintID:     githubcomnimblewaynimblego.String("fp-abc123"),
			HackiumConfiguration: githubcomnimblewaynimblego.ExtractParamsRenderOptionsHackiumConfiguration{
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
		Session: githubcomnimblewaynimblego.ExtractParamsSession{
			ID:                  githubcomnimblewaynimblego.String("id"),
			PrefetchUserbrowser: githubcomnimblewaynimblego.Bool(true),
			Retry:               githubcomnimblewaynimblego.Bool(true),
			Timeout:             githubcomnimblewaynimblego.Float(1),
		},
		Skill: githubcomnimblewaynimblego.ExtractParamsSkillUnion{
			OfString: githubcomnimblewaynimblego.String("dynamic-content"),
		},
		SkipUbct: githubcomnimblewaynimblego.Bool(false),
		State:    githubcomnimblewaynimblego.ExtractParamsStateCa,
		Tag:      githubcomnimblewaynimblego.String("campaign-2024-q1"),
		Template: githubcomnimblewaynimblego.ExtractParamsTemplate{
			Name: "x",
			Params: map[string]any{
				"foo": "bar",
			},
		},
		Type: githubcomnimblewaynimblego.String("generic"),
		UserbrowserCreationTemplateRendered: githubcomnimblewaynimblego.ExtractParamsUserbrowserCreationTemplateRendered{
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

func TestMapWithOptionalParams(t *testing.T) {
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
	_, err := client.Search(context.TODO(), githubcomnimblewaynimblego.SearchParams{
		Query:          "x",
		ContentType:    []string{"string"},
		Country:        githubcomnimblewaynimblego.String("country"),
		DeepSearch:     githubcomnimblewaynimblego.Bool(true),
		EndDate:        githubcomnimblewaynimblego.String("end_date"),
		ExcludeDomains: []string{"string"},
		IncludeAnswer:  githubcomnimblewaynimblego.Bool(true),
		IncludeDomains: []string{"string"},
		Locale:         githubcomnimblewaynimblego.String("locale"),
		MaxSubagents:   githubcomnimblewaynimblego.Int(1),
		NumResults:     githubcomnimblewaynimblego.Int(1),
		ParsingType:    githubcomnimblewaynimblego.SearchParamsParsingTypePlainText,
		SearchEngine:   githubcomnimblewaynimblego.SearchParamsSearchEngineGoogleSearch,
		StartDate:      githubcomnimblewaynimblego.String("start_date"),
		TimeRange:      githubcomnimblewaynimblego.SearchParamsTimeRangeHour,
		Topic:          githubcomnimblewaynimblego.SearchParamsTopicGeneral,
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
