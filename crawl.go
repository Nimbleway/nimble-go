// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package nimbleway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/nimbleway-go/internal/apijson"
	"github.com/stainless-sdks/nimbleway-go/internal/requestconfig"
	"github.com/stainless-sdks/nimbleway-go/option"
	"github.com/stainless-sdks/nimbleway-go/packages/param"
	"github.com/stainless-sdks/nimbleway-go/packages/respjson"
)

// CrawlService contains methods and other services that help with interacting with
// the nimbleway API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCrawlService] method instead.
type CrawlService struct {
	Options []option.RequestOption
}

// NewCrawlService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCrawlService(opts ...option.RequestOption) (r CrawlService) {
	r = CrawlService{}
	r.Options = opts
	return
}

// Create crawl task
func (r *CrawlService) Root(ctx context.Context, body CrawlRootParams, opts ...option.RequestOption) (res *CrawlRootResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/crawl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get crawl data
func (r *CrawlService) Status(ctx context.Context, id string, opts ...option.RequestOption) (res *CrawlStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/crawl/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Cancel crawl task
func (r *CrawlService) Terminate(ctx context.Context, id string, opts ...option.RequestOption) (res *CrawlTerminateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/crawl/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type CrawlRootResponse struct {
	ID  string `json:"id,required" format:"uuid"`
	URL string `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlRootResponse) RawJSON() string { return r.JSON.raw }
func (r *CrawlRootResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlStatusResponse struct {
	ID          string                    `json:"id,required" format:"uuid"`
	AccountName string                    `json:"account_name,required"`
	Completed   float64                   `json:"completed,required"`
	CreatedAt   string                    `json:"created_at,required"`
	Status      bool                      `json:"status,required"`
	Tasks       []CrawlStatusResponseTask `json:"tasks,required"`
	Total       float64                   `json:"total,required"`
	CompletedAt string                    `json:"completed_at"`
	Name        string                    `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AccountName respjson.Field
		Completed   respjson.Field
		CreatedAt   respjson.Field
		Status      respjson.Field
		Tasks       respjson.Field
		Total       respjson.Field
		CompletedAt respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *CrawlStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlStatusResponseTask struct {
	ID  string `json:"id,required" format:"uuid"`
	URL string `json:"url,required" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlStatusResponseTask) RawJSON() string { return r.JSON.raw }
func (r *CrawlStatusResponseTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlTerminateResponse struct {
	// Any of "canceled".
	Status CrawlTerminateResponseStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlTerminateResponse) RawJSON() string { return r.JSON.raw }
func (r *CrawlTerminateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlTerminateResponseStatus string

const (
	CrawlTerminateResponseStatusCanceled CrawlTerminateResponseStatus = "canceled"
)

type CrawlRootParams struct {
	// Url to crawl.
	URL string `json:"url,required" format:"uri"`
	// Allows the crawler to follow links to external websites.
	AllowExternalLinks param.Opt[bool] `json:"allow_external_links,omitzero"`
	// Allows the crawler to follow links to subdomains of the main domain.
	AllowSubdomains param.Opt[bool] `json:"allow_subdomains,omitzero"`
	// Allows the crawler to follow internal links to sibling or parent URLs, not just
	// child paths.
	CrawlEntireDomain param.Opt[bool] `json:"crawl_entire_domain,omitzero"`
	// Do not re-scrape the same path with different (or none) query parameters.
	IgnoreQueryParameters param.Opt[bool] `json:"ignore_query_parameters,omitzero"`
	// Maximum number of pages to crawl.
	Limit param.Opt[int64] `json:"limit,omitzero"`
	// Maximum depth to crawl based on discovery order.
	MaxDiscoveryDepth param.Opt[int64] `json:"max_discovery_depth,omitzero"`
	// Name of the crawl.
	Name param.Opt[string] `json:"name,omitzero"`
	// Webhook configuration for receiving crawl results.
	Callback CrawlRootParamsCallbackUnion `json:"callback,omitzero" format:"uri"`
	// URL pathname regex patterns that exclude matching URLs from the crawl.
	ExcludePaths   []string                      `json:"exclude_paths,omitzero"`
	ExtractOptions CrawlRootParamsExtractOptions `json:"extract_options,omitzero"`
	// URL pathname regex patterns that include matching URLs in the crawl.
	IncludePaths []string `json:"include_paths,omitzero"`
	// Sitemap and other methods will be used together to find URLs.
	//
	// Any of "skip", "include", "only".
	Sitemap CrawlRootParamsSitemap `json:"sitemap,omitzero"`
	paramObj
}

func (r CrawlRootParams) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRootParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsCallbackUnion struct {
	OfCrawlRootsCallbackObject *CrawlRootParamsCallbackObject `json:",omitzero,inline"`
	OfString                   param.Opt[string]              `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsCallbackUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlRootsCallbackObject, u.OfString)
}
func (u *CrawlRootParamsCallbackUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsCallbackUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlRootsCallbackObject) {
		return u.OfCrawlRootsCallbackObject
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// The property URL is required.
type CrawlRootParamsCallbackObject struct {
	// Webhook URL to receive crawl results.
	URL string `json:"url,required" format:"uri"`
	// Type of events that should be sent to the webhook URL. (default: all)
	//
	// Any of "completed", "page", "failed", "started".
	Events []string `json:"events,omitzero"`
	// Headers to send to the webhook URL.
	Headers map[string]any `json:"headers,omitzero"`
	// Custom metadata that will be included in all webhook payloads for this crawl.
	Metadata map[string]any `json:"metadata,omitzero"`
	paramObj
}

func (r CrawlRootParamsCallbackObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsCallbackObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRootParamsCallbackObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties DebugOptions, URL are required.
type CrawlRootParamsExtractOptions struct {
	// Debug and troubleshooting options for the request
	DebugOptions CrawlRootParamsExtractOptionsDebugOptions `json:"debug_options,omitzero,required"`
	// Target URL to scrape
	URL string `json:"url,required" format:"uri"`
	// City for geolocation
	City param.Opt[string] `json:"city,omitzero"`
	// Client-side timeout in milliseconds
	ClientTimeout param.Opt[float64] `json:"client_timeout,omitzero"`
	// Whether to automatically handle cookie consent headers
	ConsentHeader param.Opt[bool] `json:"consent_header,omitzero"`
	// Whether to disable IP address validation
	DisableIPCheck param.Opt[bool] `json:"disable_ip_check,omitzero"`
	// Whether to export the userbrowser session
	ExportUserbrowser param.Opt[bool] `json:"export_userbrowser,omitzero"`
	// Whether to use HTTP/2 protocol
	Http2 param.Opt[bool] `json:"http2,omitzero"`
	// Whether to use IPv6 for the request
	Ip6 param.Opt[bool] `json:"ip6,omitzero"`
	// Whether to emulate XMLHttpRequest behavior
	IsXhr param.Opt[bool] `json:"is_xhr,omitzero"`
	// Whether to return response in Markdown format
	Markdown param.Opt[bool] `json:"markdown,omitzero"`
	// Whether to exclude HTML from the response
	NoHTML param.Opt[bool] `json:"no_html,omitzero"`
	// Whether to disable browser-based rendering
	NoUserbrowser param.Opt[bool] `json:"no_userbrowser,omitzero"`
	// Whether to parse the response content
	Parse param.Opt[bool] `json:"parse,omitzero"`
	// Whether to return raw HTTP headers in response
	RawHeaders param.Opt[bool] `json:"raw_headers,omitzero"`
	// Whether to render JavaScript content using a browser
	Render param.Opt[bool] `json:"render,omitzero"`
	// Request timeout in milliseconds
	RequestTimeout param.Opt[float64] `json:"request_timeout,omitzero"`
	// Whether to return response headers in HTTP headers
	ReturnResponseHeadersAsHeader param.Opt[bool] `json:"return_response_headers_as_header,omitzero"`
	// Whether to save the userbrowser session for reuse
	SaveUserbrowser param.Opt[bool] `json:"save_userbrowser,omitzero"`
	// Whether to skip userbrowser creation template processing
	SkipUbct param.Opt[bool] `json:"skip_ubct,omitzero"`
	// User-defined tag for request identification
	Tag param.Opt[string] `json:"tag,omitzero"`
	// Type of query or scraping template
	Type param.Opt[string] `json:"type,omitzero"`
	// Browser type to emulate
	Browser CrawlRootParamsExtractOptionsBrowserUnion `json:"browser,omitzero"`
	// Browser cookies as array of cookie objects
	Cookies CrawlRootParamsExtractOptionsCookiesUnion `json:"cookies,omitzero"`
	// Country code for geolocation and proxy selection
	//
	// Any of "AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS", "AT",
	// "AU", "AW", "AX", "AZ", "BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI", "BJ",
	// "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY", "BZ", "CA",
	// "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN", "CO", "CR", "CU",
	// "CV", "CW", "CX", "CY", "CZ", "DE", "DJ", "DK", "DM", "DO", "DZ", "EC", "EE",
	// "EG", "EH", "ER", "ES", "ET", "FI", "FJ", "FK", "FM", "FO", "FR", "GA", "GB",
	// "GD", "GE", "GF", "GG", "GH", "GI", "GL", "GM", "GN", "GP", "GQ", "GR", "GS",
	// "GT", "GU", "GW", "GY", "HK", "HM", "HN", "HR", "HT", "HU", "ID", "IE", "IL",
	// "IM", "IN", "IO", "IQ", "IR", "IS", "IT", "JE", "JM", "JO", "JP", "KE", "KG",
	// "KH", "KI", "KM", "KN", "KP", "KR", "KW", "KY", "KZ", "LA", "LB", "LC", "LI",
	// "LK", "LR", "LS", "LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MF", "MG",
	// "MH", "MK", "ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV",
	// "MW", "MX", "MY", "MZ", "NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP",
	// "NR", "NU", "NZ", "OM", "PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM", "PN",
	// "PR", "PS", "PT", "PW", "PY", "QA", "RE", "RO", "RS", "RU", "RW", "SA", "SB",
	// "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM", "SN", "SO", "SR",
	// "SS", "ST", "SV", "SX", "SY", "SZ", "TC", "TD", "TF", "TG", "TH", "TJ", "TK",
	// "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW", "TZ", "UA", "UG", "UM", "US",
	// "UY", "UZ", "VA", "VC", "VE", "VG", "VI", "VN", "VU", "WF", "WS", "XK", "YE",
	// "YT", "ZA", "ZM", "ZW", "ALL".
	Country CrawlRootParamsExtractOptionsCountry `json:"country,omitzero"`
	// Device type for browser emulation
	//
	// Any of "desktop", "mobile", "tablet".
	Device string `json:"device,omitzero"`
	// Browser driver to use
	//
	// Any of "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro".
	Driver string `json:"driver,omitzero"`
	// Custom parser configuration as a key-value map
	DynamicParser map[string]any `json:"dynamic_parser,omitzero"`
	// Expected HTTP status codes for successful requests
	ExpectedStatusCodes []int64 `json:"expected_status_codes,omitzero"`
	// Response format
	//
	// Any of "json", "html", "csv", "raw", "json-lines", "markdown".
	Format string `json:"format,omitzero"`
	// Custom HTTP headers to include in the request
	Headers map[string]CrawlRootParamsExtractOptionsHeaderUnion `json:"headers,omitzero"`
	// Locale for browser language and region settings
	//
	// Any of "aa-DJ", "aa-ER", "aa-ET", "af", "af-NA", "af-ZA", "ak", "ak-GH", "am",
	// "am-ET", "an-ES", "ar", "ar-AE", "ar-BH", "ar-DZ", "ar-EG", "ar-IN", "ar-IQ",
	// "ar-JO", "ar-KW", "ar-LB", "ar-LY", "ar-MA", "ar-OM", "ar-QA", "ar-SA", "ar-SD",
	// "ar-SY", "ar-TN", "ar-YE", "as", "as-IN", "asa", "asa-TZ", "ast-ES", "az",
	// "az-AZ", "az-Cyrl", "az-Cyrl-AZ", "az-Latn", "az-Latn-AZ", "be", "be-BY", "bem",
	// "bem-ZM", "ber-DZ", "ber-MA", "bez", "bez-TZ", "bg", "bg-BG", "bho-IN", "bm",
	// "bm-ML", "bn", "bn-BD", "bn-IN", "bo", "bo-CN", "bo-IN", "br-FR", "brx-IN",
	// "bs", "bs-BA", "byn-ER", "ca", "ca-AD", "ca-ES", "ca-FR", "ca-IT", "cgg",
	// "cgg-UG", "chr", "chr-US", "crh-UA", "cs", "cs-CZ", "csb-PL", "cv-RU", "cy",
	// "cy-GB", "da", "da-DK", "dav", "dav-KE", "de", "de-AT", "de-BE", "de-CH",
	// "de-DE", "de-LI", "de-LU", "dv-MV", "dz-BT", "ebu", "ebu-KE", "ee", "ee-GH",
	// "ee-TG", "el", "el-CY", "el-GR", "en", "en-AG", "en-AS", "en-AU", "en-BE",
	// "en-BW", "en-BZ", "en-CA", "en-DK", "en-GB", "en-GU", "en-HK", "en-IE", "en-IN",
	// "en-JM", "en-MH", "en-MP", "en-MT", "en-MU", "en-NA", "en-NG", "en-NZ", "en-PH",
	// "en-PK", "en-SG", "en-TT", "en-UM", "en-US", "en-VI", "en-ZA", "en-ZM", "en-ZW",
	// "eo", "es", "es-419", "es-AR", "es-BO", "es-CL", "es-CO", "es-CR", "es-CU",
	// "es-DO", "es-EC", "es-ES", "es-GQ", "es-GT", "es-HN", "es-MX", "es-NI", "es-PA",
	// "es-PE", "es-PR", "es-PY", "es-SV", "es-US", "es-UY", "es-VE", "et", "et-EE",
	// "eu", "eu-ES", "fa", "fa-AF", "fa-IR", "ff", "ff-SN", "fi", "fi-FI", "fil",
	// "fil-PH", "fo", "fo-FO", "fr", "fr-BE", "fr-BF", "fr-BI", "fr-BJ", "fr-BL",
	// "fr-CA", "fr-CD", "fr-CF", "fr-CG", "fr-CH", "fr-CI", "fr-CM", "fr-DJ", "fr-FR",
	// "fr-GA", "fr-GN", "fr-GP", "fr-GQ", "fr-KM", "fr-LU", "fr-MC", "fr-MF", "fr-MG",
	// "fr-ML", "fr-MQ", "fr-NE", "fr-RE", "fr-RW", "fr-SN", "fr-TD", "fr-TG",
	// "fur-IT", "fy-DE", "fy-NL", "ga", "ga-IE", "gd-GB", "gez-ER", "gez-ET", "gl",
	// "gl-ES", "gsw", "gsw-CH", "gu", "gu-IN", "guz", "guz-KE", "gv", "gv-GB", "ha",
	// "ha-Latn", "ha-Latn-GH", "ha-Latn-NE", "ha-Latn-NG", "ha-NG", "haw", "haw-US",
	// "he", "he-IL", "hi", "hi-IN", "hne-IN", "hr", "hr-HR", "hsb-DE", "ht-HT", "hu",
	// "hu-HU", "hy", "hy-AM", "id", "id-ID", "ig", "ig-NG", "ii", "ii-CN", "ik-CA",
	// "is", "is-IS", "it", "it-CH", "it-IT", "iu-CA", "iw-IL", "ja", "ja-JP", "jmc",
	// "jmc-TZ", "ka", "ka-GE", "kab", "kab-DZ", "kam", "kam-KE", "kde", "kde-TZ",
	// "kea", "kea-CV", "khq", "khq-ML", "ki", "ki-KE", "kk", "kk-Cyrl", "kk-Cyrl-KZ",
	// "kk-KZ", "kl", "kl-GL", "kln", "kln-KE", "km", "km-KH", "kn", "kn-IN", "ko",
	// "ko-KR", "kok", "kok-IN", "ks-IN", "ku-TR", "kw", "kw-GB", "ky-KG", "lag",
	// "lag-TZ", "lb-LU", "lg", "lg-UG", "li-BE", "li-NL", "lij-IT", "lo-LA", "lt",
	// "lt-LT", "luo", "luo-KE", "luy", "luy-KE", "lv", "lv-LV", "mag-IN", "mai-IN",
	// "mas", "mas-KE", "mas-TZ", "mer", "mer-KE", "mfe", "mfe-MU", "mg", "mg-MG",
	// "mhr-RU", "mi-NZ", "mk", "mk-MK", "ml", "ml-IN", "mn-MN", "mr", "mr-IN", "ms",
	// "ms-BN", "ms-MY", "mt", "mt-MT", "my", "my-MM", "nan-TW", "naq", "naq-NA", "nb",
	// "nb-NO", "nd", "nd-ZW", "nds-DE", "nds-NL", "ne", "ne-IN", "ne-NP", "nl",
	// "nl-AW", "nl-BE", "nl-NL", "nn", "nn-NO", "nr-ZA", "nso-ZA", "nyn", "nyn-UG",
	// "oc-FR", "om", "om-ET", "om-KE", "or", "or-IN", "os-RU", "pa", "pa-Arab",
	// "pa-Arab-PK", "pa-Guru", "pa-Guru-IN", "pa-IN", "pa-PK", "pap-AN", "pl",
	// "pl-PL", "ps", "ps-AF", "pt", "pt-BR", "pt-GW", "pt-MZ", "pt-PT", "rm", "rm-CH",
	// "ro", "ro-MD", "ro-RO", "rof", "rof-TZ", "ru", "ru-MD", "ru-RU", "ru-UA", "rw",
	// "rw-RW", "rwk", "rwk-TZ", "sa-IN", "saq", "saq-KE", "sc-IT", "sd-IN", "se-NO",
	// "seh", "seh-MZ", "ses", "ses-ML", "sg", "sg-CF", "shi", "shi-Latn",
	// "shi-Latn-MA", "shi-Tfng", "shi-Tfng-MA", "shs-CA", "si", "si-LK", "sid-ET",
	// "sk", "sk-SK", "sl", "sl-SI", "sn", "sn-ZW", "so", "so-DJ", "so-ET", "so-KE",
	// "so-SO", "sq", "sq-AL", "sq-MK", "sr", "sr-Cyrl", "sr-Cyrl-BA", "sr-Cyrl-ME",
	// "sr-Cyrl-RS", "sr-Latn", "sr-Latn-BA", "sr-Latn-ME", "sr-Latn-RS", "sr-ME",
	// "sr-RS", "ss-ZA", "st-ZA", "sv", "sv-FI", "sv-SE", "sw", "sw-KE", "sw-TZ", "ta",
	// "ta-IN", "ta-LK", "te", "te-IN", "teo", "teo-KE", "teo-UG", "tg-TJ", "th",
	// "th-TH", "ti", "ti-ER", "ti-ET", "tig-ER", "tk-TM", "tl-PH", "tn-ZA", "to",
	// "to-TO", "tr", "tr-CY", "tr-TR", "ts-ZA", "tt-RU", "tzm", "tzm-Latn",
	// "tzm-Latn-MA", "ug-CN", "uk", "uk-UA", "unm-US", "ur", "ur-IN", "ur-PK", "uz",
	// "uz-Arab", "uz-Arab-AF", "uz-Cyrl", "uz-Cyrl-UZ", "uz-Latn", "uz-Latn-UZ",
	// "uz-UZ", "ve-ZA", "vi", "vi-VN", "vun", "vun-TZ", "wa-BE", "wae-CH", "wal-ET",
	// "wo-SN", "xh-ZA", "xog", "xog-UG", "yi-US", "yo", "yo-NG", "yue-HK", "zh",
	// "zh-CN", "zh-HK", "zh-Hans", "zh-Hans-CN", "zh-Hans-HK", "zh-Hans-MO",
	// "zh-Hans-SG", "zh-Hant", "zh-Hant-HK", "zh-Hant-MO", "zh-Hant-TW", "zh-SG",
	// "zh-TW", "zu", "zu-ZA", "auto".
	Locale CrawlRootParamsExtractOptionsLocale `json:"locale,omitzero"`
	// Structured metadata about the request execution context
	Metadata CrawlRootParamsExtractOptionsMetadata `json:"metadata,omitzero"`
	// HTTP method for the request
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method,omitzero"`
	// Native execution mode
	//
	// Any of "requester", "apm", "direct".
	NativeMode string `json:"native_mode,omitzero"`
	// Filters for capturing network traffic
	NetworkCapture []CrawlRootParamsExtractOptionsNetworkCapture `json:"network_capture,omitzero"`
	// Operating system to emulate
	//
	// Any of "windows", "mac os", "linux", "android", "ios".
	Os string `json:"os,omitzero"`
	// Configuration options for parsing behavior
	ParseOptions CrawlRootParamsExtractOptionsParseOptions `json:"parse_options,omitzero"`
	// Custom parser configuration as a key-value map
	Parser CrawlRootParamsExtractOptionsParserUnion `json:"parser,omitzero"`
	// Proxy provider to use for the request
	//
	// Any of "brightdata", "oxylabs", "smartproxy", "proxit", "proxit_preprod",
	// "local", "rayobyte", "always", "oculusproxies", "froxy", "packetstream",
	// "911proxy", "direct911proxy", "thesocialproxy", "thesocialproxy2", "nimble-isp",
	// "nimble-isp-mobile", "proxit-linux", "proxit-macos", "proxit-windows",
	// "proxit-rental", "ipfoxy", "brightup", "research".
	ProxyProvider CrawlRootParamsExtractOptionsProxyProvider `json:"proxy_provider,omitzero"`
	// Weighted distribution of proxy providers
	ProxyProviders map[string]float64 `json:"proxy_providers,omitzero"`
	// Query template configuration for structured data extraction
	QueryTemplate CrawlRootParamsExtractOptionsQueryTemplate `json:"query_template,omitzero"`
	// Referrer policy for the request
	//
	// Any of "random", "no-referer", "same-origin", "google", "bing", "facebook",
	// "twitter", "instagram".
	ReferrerType CrawlRootParamsExtractOptionsReferrerType `json:"referrer_type,omitzero"`
	// Array of actions to perform during browser rendering
	RenderFlow    []map[string]any                           `json:"render_flow,omitzero"`
	RenderOptions CrawlRootParamsExtractOptionsRenderOptions `json:"render_options,omitzero"`
	Session       CrawlRootParamsExtractOptionsSession       `json:"session,omitzero"`
	// Skills or capabilities required for the request
	Skill CrawlRootParamsExtractOptionsSkillUnion `json:"skill,omitzero"`
	// US state for geolocation (only valid when country is US)
	//
	// Any of "AL", "AK", "AS", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA",
	// "GU", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI",
	// "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "MP",
	// "OH", "OK", "OR", "PA", "PR", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA",
	// "VI", "WA", "WV", "WI", "WY".
	State string `json:"state,omitzero"`
	// Userbrowser creation template configuration
	Template CrawlRootParamsExtractOptionsTemplate `json:"template,omitzero"`
	// Pre-rendered userbrowser creation template configuration
	UserbrowserCreationTemplateRendered CrawlRootParamsExtractOptionsUserbrowserCreationTemplateRendered `json:"userbrowser_creation_template_rendered,omitzero"`
	paramObj
}

func (r CrawlRootParamsExtractOptions) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsExtractOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRootParamsExtractOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptions](
		"device", "desktop", "mobile", "tablet",
	)
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptions](
		"driver", "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro",
	)
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptions](
		"format", "json", "html", "csv", "raw", "json-lines", "markdown",
	)
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptions](
		"method", "GET", "POST", "PUT", "PATCH", "DELETE",
	)
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptions](
		"native_mode", "requester", "apm", "direct",
	)
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptions](
		"os", "windows", "mac os", "linux", "android", "ios",
	)
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptions](
		"state", "AL", "AK", "AS", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA", "GU", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "MP", "OH", "OK", "OR", "PA", "PR", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "VI", "WA", "WV", "WI", "WY",
	)
}

// Debug and troubleshooting options for the request
type CrawlRootParamsExtractOptionsDebugOptions struct {
	CollectHar       CrawlRootParamsExtractOptionsDebugOptionsCollectHarUnion       `json:"collect_har,omitzero"`
	NoRetryMode      CrawlRootParamsExtractOptionsDebugOptionsNoRetryModeUnion      `json:"no_retry_mode,omitzero"`
	RecordScreen     CrawlRootParamsExtractOptionsDebugOptionsRecordScreenUnion     `json:"record_screen,omitzero"`
	Redact           CrawlRootParamsExtractOptionsDebugOptionsRedactUnion           `json:"redact,omitzero"`
	ShowCursor       CrawlRootParamsExtractOptionsDebugOptionsShowCursorUnion       `json:"show_cursor,omitzero"`
	SolveCaptcha     CrawlRootParamsExtractOptionsDebugOptionsSolveCaptchaUnion     `json:"solve_captcha,omitzero"`
	Trace            CrawlRootParamsExtractOptionsDebugOptionsTraceUnion            `json:"trace,omitzero"`
	UploadEngineLogs CrawlRootParamsExtractOptionsDebugOptionsUploadEngineLogsUnion `json:"upload_engine_logs,omitzero"`
	Verbose          CrawlRootParamsExtractOptionsDebugOptionsVerboseUnion          `json:"verbose,omitzero"`
	WithProxyUsage   CrawlRootParamsExtractOptionsDebugOptionsWithProxyUsageUnion   `json:"with_proxy_usage,omitzero"`
	paramObj
}

func (r CrawlRootParamsExtractOptionsDebugOptions) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsExtractOptionsDebugOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRootParamsExtractOptionsDebugOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsDebugOptionsCollectHarUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlRootsExtractOptionsDebugOptionsCollectHarString)
	OfCrawlRootsExtractOptionsDebugOptionsCollectHarString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsDebugOptionsCollectHarUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfCrawlRootsExtractOptionsDebugOptionsCollectHarString)
}
func (u *CrawlRootParamsExtractOptionsDebugOptionsCollectHarUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsDebugOptionsCollectHarUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfCrawlRootsExtractOptionsDebugOptionsCollectHarString) {
		return &u.OfCrawlRootsExtractOptionsDebugOptionsCollectHarString
	}
	return nil
}

type CrawlRootParamsExtractOptionsDebugOptionsCollectHarString string

const (
	CrawlRootParamsExtractOptionsDebugOptionsCollectHarStringNever   CrawlRootParamsExtractOptionsDebugOptionsCollectHarString = "never"
	CrawlRootParamsExtractOptionsDebugOptionsCollectHarStringOnError CrawlRootParamsExtractOptionsDebugOptionsCollectHarString = "on-error"
	CrawlRootParamsExtractOptionsDebugOptionsCollectHarStringAlways  CrawlRootParamsExtractOptionsDebugOptionsCollectHarString = "always"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsDebugOptionsNoRetryModeUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlRootsExtractOptionsDebugOptionsNoRetryModeString)
	OfCrawlRootsExtractOptionsDebugOptionsNoRetryModeString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsDebugOptionsNoRetryModeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfCrawlRootsExtractOptionsDebugOptionsNoRetryModeString)
}
func (u *CrawlRootParamsExtractOptionsDebugOptionsNoRetryModeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsDebugOptionsNoRetryModeUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfCrawlRootsExtractOptionsDebugOptionsNoRetryModeString) {
		return &u.OfCrawlRootsExtractOptionsDebugOptionsNoRetryModeString
	}
	return nil
}

type CrawlRootParamsExtractOptionsDebugOptionsNoRetryModeString string

const (
	CrawlRootParamsExtractOptionsDebugOptionsNoRetryModeStringNever  CrawlRootParamsExtractOptionsDebugOptionsNoRetryModeString = "never"
	CrawlRootParamsExtractOptionsDebugOptionsNoRetryModeStringAlways CrawlRootParamsExtractOptionsDebugOptionsNoRetryModeString = "always"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsDebugOptionsRecordScreenUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlRootsExtractOptionsDebugOptionsRecordScreenString)
	OfCrawlRootsExtractOptionsDebugOptionsRecordScreenString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsDebugOptionsRecordScreenUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfCrawlRootsExtractOptionsDebugOptionsRecordScreenString)
}
func (u *CrawlRootParamsExtractOptionsDebugOptionsRecordScreenUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsDebugOptionsRecordScreenUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfCrawlRootsExtractOptionsDebugOptionsRecordScreenString) {
		return &u.OfCrawlRootsExtractOptionsDebugOptionsRecordScreenString
	}
	return nil
}

type CrawlRootParamsExtractOptionsDebugOptionsRecordScreenString string

const (
	CrawlRootParamsExtractOptionsDebugOptionsRecordScreenStringNever   CrawlRootParamsExtractOptionsDebugOptionsRecordScreenString = "never"
	CrawlRootParamsExtractOptionsDebugOptionsRecordScreenStringOnError CrawlRootParamsExtractOptionsDebugOptionsRecordScreenString = "on-error"
	CrawlRootParamsExtractOptionsDebugOptionsRecordScreenStringAlways  CrawlRootParamsExtractOptionsDebugOptionsRecordScreenString = "always"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsDebugOptionsRedactUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlRootsExtractOptionsDebugOptionsRedactString)
	OfCrawlRootsExtractOptionsDebugOptionsRedactString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsDebugOptionsRedactUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfCrawlRootsExtractOptionsDebugOptionsRedactString)
}
func (u *CrawlRootParamsExtractOptionsDebugOptionsRedactUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsDebugOptionsRedactUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfCrawlRootsExtractOptionsDebugOptionsRedactString) {
		return &u.OfCrawlRootsExtractOptionsDebugOptionsRedactString
	}
	return nil
}

type CrawlRootParamsExtractOptionsDebugOptionsRedactString string

const (
	CrawlRootParamsExtractOptionsDebugOptionsRedactStringNever  CrawlRootParamsExtractOptionsDebugOptionsRedactString = "never"
	CrawlRootParamsExtractOptionsDebugOptionsRedactStringAlways CrawlRootParamsExtractOptionsDebugOptionsRedactString = "always"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsDebugOptionsShowCursorUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlRootsExtractOptionsDebugOptionsShowCursorString)
	OfCrawlRootsExtractOptionsDebugOptionsShowCursorString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsDebugOptionsShowCursorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfCrawlRootsExtractOptionsDebugOptionsShowCursorString)
}
func (u *CrawlRootParamsExtractOptionsDebugOptionsShowCursorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsDebugOptionsShowCursorUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfCrawlRootsExtractOptionsDebugOptionsShowCursorString) {
		return &u.OfCrawlRootsExtractOptionsDebugOptionsShowCursorString
	}
	return nil
}

type CrawlRootParamsExtractOptionsDebugOptionsShowCursorString string

const (
	CrawlRootParamsExtractOptionsDebugOptionsShowCursorStringNever  CrawlRootParamsExtractOptionsDebugOptionsShowCursorString = "never"
	CrawlRootParamsExtractOptionsDebugOptionsShowCursorStringAlways CrawlRootParamsExtractOptionsDebugOptionsShowCursorString = "always"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsDebugOptionsSolveCaptchaUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlRootsExtractOptionsDebugOptionsSolveCaptchaString)
	OfCrawlRootsExtractOptionsDebugOptionsSolveCaptchaString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsDebugOptionsSolveCaptchaUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfCrawlRootsExtractOptionsDebugOptionsSolveCaptchaString)
}
func (u *CrawlRootParamsExtractOptionsDebugOptionsSolveCaptchaUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsDebugOptionsSolveCaptchaUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfCrawlRootsExtractOptionsDebugOptionsSolveCaptchaString) {
		return &u.OfCrawlRootsExtractOptionsDebugOptionsSolveCaptchaString
	}
	return nil
}

type CrawlRootParamsExtractOptionsDebugOptionsSolveCaptchaString string

const (
	CrawlRootParamsExtractOptionsDebugOptionsSolveCaptchaStringNever  CrawlRootParamsExtractOptionsDebugOptionsSolveCaptchaString = "never"
	CrawlRootParamsExtractOptionsDebugOptionsSolveCaptchaStringAlways CrawlRootParamsExtractOptionsDebugOptionsSolveCaptchaString = "always"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsDebugOptionsTraceUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlRootsExtractOptionsDebugOptionsTraceString)
	OfCrawlRootsExtractOptionsDebugOptionsTraceString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsDebugOptionsTraceUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfCrawlRootsExtractOptionsDebugOptionsTraceString)
}
func (u *CrawlRootParamsExtractOptionsDebugOptionsTraceUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsDebugOptionsTraceUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfCrawlRootsExtractOptionsDebugOptionsTraceString) {
		return &u.OfCrawlRootsExtractOptionsDebugOptionsTraceString
	}
	return nil
}

type CrawlRootParamsExtractOptionsDebugOptionsTraceString string

const (
	CrawlRootParamsExtractOptionsDebugOptionsTraceStringNever   CrawlRootParamsExtractOptionsDebugOptionsTraceString = "never"
	CrawlRootParamsExtractOptionsDebugOptionsTraceStringOnError CrawlRootParamsExtractOptionsDebugOptionsTraceString = "on-error"
	CrawlRootParamsExtractOptionsDebugOptionsTraceStringAlways  CrawlRootParamsExtractOptionsDebugOptionsTraceString = "always"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsDebugOptionsUploadEngineLogsUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlRootsExtractOptionsDebugOptionsUploadEngineLogsString)
	OfCrawlRootsExtractOptionsDebugOptionsUploadEngineLogsString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsDebugOptionsUploadEngineLogsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfCrawlRootsExtractOptionsDebugOptionsUploadEngineLogsString)
}
func (u *CrawlRootParamsExtractOptionsDebugOptionsUploadEngineLogsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsDebugOptionsUploadEngineLogsUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfCrawlRootsExtractOptionsDebugOptionsUploadEngineLogsString) {
		return &u.OfCrawlRootsExtractOptionsDebugOptionsUploadEngineLogsString
	}
	return nil
}

type CrawlRootParamsExtractOptionsDebugOptionsUploadEngineLogsString string

const (
	CrawlRootParamsExtractOptionsDebugOptionsUploadEngineLogsStringNever   CrawlRootParamsExtractOptionsDebugOptionsUploadEngineLogsString = "never"
	CrawlRootParamsExtractOptionsDebugOptionsUploadEngineLogsStringOnError CrawlRootParamsExtractOptionsDebugOptionsUploadEngineLogsString = "on-error"
	CrawlRootParamsExtractOptionsDebugOptionsUploadEngineLogsStringAlways  CrawlRootParamsExtractOptionsDebugOptionsUploadEngineLogsString = "always"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsDebugOptionsVerboseUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlRootsExtractOptionsDebugOptionsVerboseString)
	OfCrawlRootsExtractOptionsDebugOptionsVerboseString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsDebugOptionsVerboseUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfCrawlRootsExtractOptionsDebugOptionsVerboseString)
}
func (u *CrawlRootParamsExtractOptionsDebugOptionsVerboseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsDebugOptionsVerboseUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfCrawlRootsExtractOptionsDebugOptionsVerboseString) {
		return &u.OfCrawlRootsExtractOptionsDebugOptionsVerboseString
	}
	return nil
}

type CrawlRootParamsExtractOptionsDebugOptionsVerboseString string

const (
	CrawlRootParamsExtractOptionsDebugOptionsVerboseStringNever  CrawlRootParamsExtractOptionsDebugOptionsVerboseString = "never"
	CrawlRootParamsExtractOptionsDebugOptionsVerboseStringAlways CrawlRootParamsExtractOptionsDebugOptionsVerboseString = "always"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsDebugOptionsWithProxyUsageUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlRootsExtractOptionsDebugOptionsWithProxyUsageString)
	OfCrawlRootsExtractOptionsDebugOptionsWithProxyUsageString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsDebugOptionsWithProxyUsageUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfCrawlRootsExtractOptionsDebugOptionsWithProxyUsageString)
}
func (u *CrawlRootParamsExtractOptionsDebugOptionsWithProxyUsageUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsDebugOptionsWithProxyUsageUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfCrawlRootsExtractOptionsDebugOptionsWithProxyUsageString) {
		return &u.OfCrawlRootsExtractOptionsDebugOptionsWithProxyUsageString
	}
	return nil
}

type CrawlRootParamsExtractOptionsDebugOptionsWithProxyUsageString string

const (
	CrawlRootParamsExtractOptionsDebugOptionsWithProxyUsageStringNever  CrawlRootParamsExtractOptionsDebugOptionsWithProxyUsageString = "never"
	CrawlRootParamsExtractOptionsDebugOptionsWithProxyUsageStringAlways CrawlRootParamsExtractOptionsDebugOptionsWithProxyUsageString = "always"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsBrowserUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlRootsExtractOptionsBrowserString)
	OfCrawlRootsExtractOptionsBrowserString param.Opt[string]                           `json:",omitzero,inline"`
	OfCrawlRootsExtractOptionsBrowserObject *CrawlRootParamsExtractOptionsBrowserObject `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsBrowserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlRootsExtractOptionsBrowserString, u.OfCrawlRootsExtractOptionsBrowserObject)
}
func (u *CrawlRootParamsExtractOptionsBrowserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsBrowserUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlRootsExtractOptionsBrowserString) {
		return &u.OfCrawlRootsExtractOptionsBrowserString
	} else if !param.IsOmitted(u.OfCrawlRootsExtractOptionsBrowserObject) {
		return u.OfCrawlRootsExtractOptionsBrowserObject
	}
	return nil
}

// Browser type to emulate
type CrawlRootParamsExtractOptionsBrowserString string

const (
	CrawlRootParamsExtractOptionsBrowserStringChrome  CrawlRootParamsExtractOptionsBrowserString = "chrome"
	CrawlRootParamsExtractOptionsBrowserStringFirefox CrawlRootParamsExtractOptionsBrowserString = "firefox"
)

// The property Name is required.
type CrawlRootParamsExtractOptionsBrowserObject struct {
	// Any of "chrome", "firefox".
	Name string `json:"name,omitzero,required"`
	// Specific browser version to emulate
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r CrawlRootParamsExtractOptionsBrowserObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsExtractOptionsBrowserObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRootParamsExtractOptionsBrowserObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptionsBrowserObject](
		"name", "chrome", "firefox",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsCookiesUnion struct {
	OfCrawlRootsExtractOptionsCookiesArray []CrawlRootParamsExtractOptionsCookiesArrayItem `json:",omitzero,inline"`
	OfString                               param.Opt[string]                               `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsCookiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlRootsExtractOptionsCookiesArray, u.OfString)
}
func (u *CrawlRootParamsExtractOptionsCookiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsCookiesUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlRootsExtractOptionsCookiesArray) {
		return &u.OfCrawlRootsExtractOptionsCookiesArray
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type CrawlRootParamsExtractOptionsCookiesArrayItem struct {
	Creation      param.Opt[string]                                        `json:"creation,omitzero"`
	Domain        param.Opt[string]                                        `json:"domain,omitzero"`
	HostOnly      param.Opt[bool]                                          `json:"hostOnly,omitzero"`
	HTTPOnly      param.Opt[bool]                                          `json:"httpOnly,omitzero"`
	LastAccessed  param.Opt[string]                                        `json:"lastAccessed,omitzero"`
	Path          param.Opt[string]                                        `json:"path,omitzero"`
	PathIsDefault param.Opt[bool]                                          `json:"pathIsDefault,omitzero"`
	Expires       param.Opt[string]                                        `json:"expires,omitzero"`
	Name          param.Opt[string]                                        `json:"name,omitzero"`
	Secure        param.Opt[bool]                                          `json:"secure,omitzero"`
	Value         param.Opt[string]                                        `json:"value,omitzero"`
	Extensions    []string                                                 `json:"extensions,omitzero"`
	MaxAge        CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeUnion `json:"maxAge,omitzero"`
	// Any of "strict", "lax", "none".
	SameSite    string         `json:"sameSite,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r CrawlRootParamsExtractOptionsCookiesArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsExtractOptionsCookiesArrayItem
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *CrawlRootParamsExtractOptionsCookiesArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptionsCookiesArrayItem](
		"sameSite", "strict", "lax", "none",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlRootsExtractOptionsCookiesArrayItemMaxAgeString)
	OfCrawlRootsExtractOptionsCookiesArrayItemMaxAgeString param.Opt[CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeString] `json:",omitzero,inline"`
	OfFloat                                                param.Opt[float64]                                                   `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlRootsExtractOptionsCookiesArrayItemMaxAgeString, u.OfFloat)
}
func (u *CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlRootsExtractOptionsCookiesArrayItemMaxAgeString) {
		return &u.OfCrawlRootsExtractOptionsCookiesArrayItemMaxAgeString
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeString string

const (
	CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeStringInfinity      CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeString = "Infinity"
	CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeStringMinusInfinity CrawlRootParamsExtractOptionsCookiesArrayItemMaxAgeString = "-Infinity"
)

// Country code for geolocation and proxy selection
type CrawlRootParamsExtractOptionsCountry string

const (
	CrawlRootParamsExtractOptionsCountryAd  CrawlRootParamsExtractOptionsCountry = "AD"
	CrawlRootParamsExtractOptionsCountryAe  CrawlRootParamsExtractOptionsCountry = "AE"
	CrawlRootParamsExtractOptionsCountryAf  CrawlRootParamsExtractOptionsCountry = "AF"
	CrawlRootParamsExtractOptionsCountryAg  CrawlRootParamsExtractOptionsCountry = "AG"
	CrawlRootParamsExtractOptionsCountryAI  CrawlRootParamsExtractOptionsCountry = "AI"
	CrawlRootParamsExtractOptionsCountryAl  CrawlRootParamsExtractOptionsCountry = "AL"
	CrawlRootParamsExtractOptionsCountryAm  CrawlRootParamsExtractOptionsCountry = "AM"
	CrawlRootParamsExtractOptionsCountryAo  CrawlRootParamsExtractOptionsCountry = "AO"
	CrawlRootParamsExtractOptionsCountryAq  CrawlRootParamsExtractOptionsCountry = "AQ"
	CrawlRootParamsExtractOptionsCountryAr  CrawlRootParamsExtractOptionsCountry = "AR"
	CrawlRootParamsExtractOptionsCountryAs  CrawlRootParamsExtractOptionsCountry = "AS"
	CrawlRootParamsExtractOptionsCountryAt  CrawlRootParamsExtractOptionsCountry = "AT"
	CrawlRootParamsExtractOptionsCountryAu  CrawlRootParamsExtractOptionsCountry = "AU"
	CrawlRootParamsExtractOptionsCountryAw  CrawlRootParamsExtractOptionsCountry = "AW"
	CrawlRootParamsExtractOptionsCountryAx  CrawlRootParamsExtractOptionsCountry = "AX"
	CrawlRootParamsExtractOptionsCountryAz  CrawlRootParamsExtractOptionsCountry = "AZ"
	CrawlRootParamsExtractOptionsCountryBa  CrawlRootParamsExtractOptionsCountry = "BA"
	CrawlRootParamsExtractOptionsCountryBb  CrawlRootParamsExtractOptionsCountry = "BB"
	CrawlRootParamsExtractOptionsCountryBd  CrawlRootParamsExtractOptionsCountry = "BD"
	CrawlRootParamsExtractOptionsCountryBe  CrawlRootParamsExtractOptionsCountry = "BE"
	CrawlRootParamsExtractOptionsCountryBf  CrawlRootParamsExtractOptionsCountry = "BF"
	CrawlRootParamsExtractOptionsCountryBg  CrawlRootParamsExtractOptionsCountry = "BG"
	CrawlRootParamsExtractOptionsCountryBh  CrawlRootParamsExtractOptionsCountry = "BH"
	CrawlRootParamsExtractOptionsCountryBi  CrawlRootParamsExtractOptionsCountry = "BI"
	CrawlRootParamsExtractOptionsCountryBj  CrawlRootParamsExtractOptionsCountry = "BJ"
	CrawlRootParamsExtractOptionsCountryBl  CrawlRootParamsExtractOptionsCountry = "BL"
	CrawlRootParamsExtractOptionsCountryBm  CrawlRootParamsExtractOptionsCountry = "BM"
	CrawlRootParamsExtractOptionsCountryBn  CrawlRootParamsExtractOptionsCountry = "BN"
	CrawlRootParamsExtractOptionsCountryBo  CrawlRootParamsExtractOptionsCountry = "BO"
	CrawlRootParamsExtractOptionsCountryBq  CrawlRootParamsExtractOptionsCountry = "BQ"
	CrawlRootParamsExtractOptionsCountryBr  CrawlRootParamsExtractOptionsCountry = "BR"
	CrawlRootParamsExtractOptionsCountryBs  CrawlRootParamsExtractOptionsCountry = "BS"
	CrawlRootParamsExtractOptionsCountryBt  CrawlRootParamsExtractOptionsCountry = "BT"
	CrawlRootParamsExtractOptionsCountryBv  CrawlRootParamsExtractOptionsCountry = "BV"
	CrawlRootParamsExtractOptionsCountryBw  CrawlRootParamsExtractOptionsCountry = "BW"
	CrawlRootParamsExtractOptionsCountryBy  CrawlRootParamsExtractOptionsCountry = "BY"
	CrawlRootParamsExtractOptionsCountryBz  CrawlRootParamsExtractOptionsCountry = "BZ"
	CrawlRootParamsExtractOptionsCountryCa  CrawlRootParamsExtractOptionsCountry = "CA"
	CrawlRootParamsExtractOptionsCountryCc  CrawlRootParamsExtractOptionsCountry = "CC"
	CrawlRootParamsExtractOptionsCountryCd  CrawlRootParamsExtractOptionsCountry = "CD"
	CrawlRootParamsExtractOptionsCountryCf  CrawlRootParamsExtractOptionsCountry = "CF"
	CrawlRootParamsExtractOptionsCountryCg  CrawlRootParamsExtractOptionsCountry = "CG"
	CrawlRootParamsExtractOptionsCountryCh  CrawlRootParamsExtractOptionsCountry = "CH"
	CrawlRootParamsExtractOptionsCountryCi  CrawlRootParamsExtractOptionsCountry = "CI"
	CrawlRootParamsExtractOptionsCountryCk  CrawlRootParamsExtractOptionsCountry = "CK"
	CrawlRootParamsExtractOptionsCountryCl  CrawlRootParamsExtractOptionsCountry = "CL"
	CrawlRootParamsExtractOptionsCountryCm  CrawlRootParamsExtractOptionsCountry = "CM"
	CrawlRootParamsExtractOptionsCountryCn  CrawlRootParamsExtractOptionsCountry = "CN"
	CrawlRootParamsExtractOptionsCountryCo  CrawlRootParamsExtractOptionsCountry = "CO"
	CrawlRootParamsExtractOptionsCountryCr  CrawlRootParamsExtractOptionsCountry = "CR"
	CrawlRootParamsExtractOptionsCountryCu  CrawlRootParamsExtractOptionsCountry = "CU"
	CrawlRootParamsExtractOptionsCountryCv  CrawlRootParamsExtractOptionsCountry = "CV"
	CrawlRootParamsExtractOptionsCountryCw  CrawlRootParamsExtractOptionsCountry = "CW"
	CrawlRootParamsExtractOptionsCountryCx  CrawlRootParamsExtractOptionsCountry = "CX"
	CrawlRootParamsExtractOptionsCountryCy  CrawlRootParamsExtractOptionsCountry = "CY"
	CrawlRootParamsExtractOptionsCountryCz  CrawlRootParamsExtractOptionsCountry = "CZ"
	CrawlRootParamsExtractOptionsCountryDe  CrawlRootParamsExtractOptionsCountry = "DE"
	CrawlRootParamsExtractOptionsCountryDj  CrawlRootParamsExtractOptionsCountry = "DJ"
	CrawlRootParamsExtractOptionsCountryDk  CrawlRootParamsExtractOptionsCountry = "DK"
	CrawlRootParamsExtractOptionsCountryDm  CrawlRootParamsExtractOptionsCountry = "DM"
	CrawlRootParamsExtractOptionsCountryDo  CrawlRootParamsExtractOptionsCountry = "DO"
	CrawlRootParamsExtractOptionsCountryDz  CrawlRootParamsExtractOptionsCountry = "DZ"
	CrawlRootParamsExtractOptionsCountryEc  CrawlRootParamsExtractOptionsCountry = "EC"
	CrawlRootParamsExtractOptionsCountryEe  CrawlRootParamsExtractOptionsCountry = "EE"
	CrawlRootParamsExtractOptionsCountryEg  CrawlRootParamsExtractOptionsCountry = "EG"
	CrawlRootParamsExtractOptionsCountryEh  CrawlRootParamsExtractOptionsCountry = "EH"
	CrawlRootParamsExtractOptionsCountryEr  CrawlRootParamsExtractOptionsCountry = "ER"
	CrawlRootParamsExtractOptionsCountryEs  CrawlRootParamsExtractOptionsCountry = "ES"
	CrawlRootParamsExtractOptionsCountryEt  CrawlRootParamsExtractOptionsCountry = "ET"
	CrawlRootParamsExtractOptionsCountryFi  CrawlRootParamsExtractOptionsCountry = "FI"
	CrawlRootParamsExtractOptionsCountryFj  CrawlRootParamsExtractOptionsCountry = "FJ"
	CrawlRootParamsExtractOptionsCountryFk  CrawlRootParamsExtractOptionsCountry = "FK"
	CrawlRootParamsExtractOptionsCountryFm  CrawlRootParamsExtractOptionsCountry = "FM"
	CrawlRootParamsExtractOptionsCountryFo  CrawlRootParamsExtractOptionsCountry = "FO"
	CrawlRootParamsExtractOptionsCountryFr  CrawlRootParamsExtractOptionsCountry = "FR"
	CrawlRootParamsExtractOptionsCountryGa  CrawlRootParamsExtractOptionsCountry = "GA"
	CrawlRootParamsExtractOptionsCountryGB  CrawlRootParamsExtractOptionsCountry = "GB"
	CrawlRootParamsExtractOptionsCountryGd  CrawlRootParamsExtractOptionsCountry = "GD"
	CrawlRootParamsExtractOptionsCountryGe  CrawlRootParamsExtractOptionsCountry = "GE"
	CrawlRootParamsExtractOptionsCountryGf  CrawlRootParamsExtractOptionsCountry = "GF"
	CrawlRootParamsExtractOptionsCountryGg  CrawlRootParamsExtractOptionsCountry = "GG"
	CrawlRootParamsExtractOptionsCountryGh  CrawlRootParamsExtractOptionsCountry = "GH"
	CrawlRootParamsExtractOptionsCountryGi  CrawlRootParamsExtractOptionsCountry = "GI"
	CrawlRootParamsExtractOptionsCountryGl  CrawlRootParamsExtractOptionsCountry = "GL"
	CrawlRootParamsExtractOptionsCountryGm  CrawlRootParamsExtractOptionsCountry = "GM"
	CrawlRootParamsExtractOptionsCountryGn  CrawlRootParamsExtractOptionsCountry = "GN"
	CrawlRootParamsExtractOptionsCountryGp  CrawlRootParamsExtractOptionsCountry = "GP"
	CrawlRootParamsExtractOptionsCountryGq  CrawlRootParamsExtractOptionsCountry = "GQ"
	CrawlRootParamsExtractOptionsCountryGr  CrawlRootParamsExtractOptionsCountry = "GR"
	CrawlRootParamsExtractOptionsCountryGs  CrawlRootParamsExtractOptionsCountry = "GS"
	CrawlRootParamsExtractOptionsCountryGt  CrawlRootParamsExtractOptionsCountry = "GT"
	CrawlRootParamsExtractOptionsCountryGu  CrawlRootParamsExtractOptionsCountry = "GU"
	CrawlRootParamsExtractOptionsCountryGw  CrawlRootParamsExtractOptionsCountry = "GW"
	CrawlRootParamsExtractOptionsCountryGy  CrawlRootParamsExtractOptionsCountry = "GY"
	CrawlRootParamsExtractOptionsCountryHk  CrawlRootParamsExtractOptionsCountry = "HK"
	CrawlRootParamsExtractOptionsCountryHm  CrawlRootParamsExtractOptionsCountry = "HM"
	CrawlRootParamsExtractOptionsCountryHn  CrawlRootParamsExtractOptionsCountry = "HN"
	CrawlRootParamsExtractOptionsCountryHr  CrawlRootParamsExtractOptionsCountry = "HR"
	CrawlRootParamsExtractOptionsCountryHt  CrawlRootParamsExtractOptionsCountry = "HT"
	CrawlRootParamsExtractOptionsCountryHu  CrawlRootParamsExtractOptionsCountry = "HU"
	CrawlRootParamsExtractOptionsCountryID  CrawlRootParamsExtractOptionsCountry = "ID"
	CrawlRootParamsExtractOptionsCountryIe  CrawlRootParamsExtractOptionsCountry = "IE"
	CrawlRootParamsExtractOptionsCountryIl  CrawlRootParamsExtractOptionsCountry = "IL"
	CrawlRootParamsExtractOptionsCountryIm  CrawlRootParamsExtractOptionsCountry = "IM"
	CrawlRootParamsExtractOptionsCountryIn  CrawlRootParamsExtractOptionsCountry = "IN"
	CrawlRootParamsExtractOptionsCountryIo  CrawlRootParamsExtractOptionsCountry = "IO"
	CrawlRootParamsExtractOptionsCountryIq  CrawlRootParamsExtractOptionsCountry = "IQ"
	CrawlRootParamsExtractOptionsCountryIr  CrawlRootParamsExtractOptionsCountry = "IR"
	CrawlRootParamsExtractOptionsCountryIs  CrawlRootParamsExtractOptionsCountry = "IS"
	CrawlRootParamsExtractOptionsCountryIt  CrawlRootParamsExtractOptionsCountry = "IT"
	CrawlRootParamsExtractOptionsCountryJe  CrawlRootParamsExtractOptionsCountry = "JE"
	CrawlRootParamsExtractOptionsCountryJm  CrawlRootParamsExtractOptionsCountry = "JM"
	CrawlRootParamsExtractOptionsCountryJo  CrawlRootParamsExtractOptionsCountry = "JO"
	CrawlRootParamsExtractOptionsCountryJp  CrawlRootParamsExtractOptionsCountry = "JP"
	CrawlRootParamsExtractOptionsCountryKe  CrawlRootParamsExtractOptionsCountry = "KE"
	CrawlRootParamsExtractOptionsCountryKg  CrawlRootParamsExtractOptionsCountry = "KG"
	CrawlRootParamsExtractOptionsCountryKh  CrawlRootParamsExtractOptionsCountry = "KH"
	CrawlRootParamsExtractOptionsCountryKi  CrawlRootParamsExtractOptionsCountry = "KI"
	CrawlRootParamsExtractOptionsCountryKm  CrawlRootParamsExtractOptionsCountry = "KM"
	CrawlRootParamsExtractOptionsCountryKn  CrawlRootParamsExtractOptionsCountry = "KN"
	CrawlRootParamsExtractOptionsCountryKp  CrawlRootParamsExtractOptionsCountry = "KP"
	CrawlRootParamsExtractOptionsCountryKr  CrawlRootParamsExtractOptionsCountry = "KR"
	CrawlRootParamsExtractOptionsCountryKw  CrawlRootParamsExtractOptionsCountry = "KW"
	CrawlRootParamsExtractOptionsCountryKy  CrawlRootParamsExtractOptionsCountry = "KY"
	CrawlRootParamsExtractOptionsCountryKz  CrawlRootParamsExtractOptionsCountry = "KZ"
	CrawlRootParamsExtractOptionsCountryLa  CrawlRootParamsExtractOptionsCountry = "LA"
	CrawlRootParamsExtractOptionsCountryLb  CrawlRootParamsExtractOptionsCountry = "LB"
	CrawlRootParamsExtractOptionsCountryLc  CrawlRootParamsExtractOptionsCountry = "LC"
	CrawlRootParamsExtractOptionsCountryLi  CrawlRootParamsExtractOptionsCountry = "LI"
	CrawlRootParamsExtractOptionsCountryLk  CrawlRootParamsExtractOptionsCountry = "LK"
	CrawlRootParamsExtractOptionsCountryLr  CrawlRootParamsExtractOptionsCountry = "LR"
	CrawlRootParamsExtractOptionsCountryLs  CrawlRootParamsExtractOptionsCountry = "LS"
	CrawlRootParamsExtractOptionsCountryLt  CrawlRootParamsExtractOptionsCountry = "LT"
	CrawlRootParamsExtractOptionsCountryLu  CrawlRootParamsExtractOptionsCountry = "LU"
	CrawlRootParamsExtractOptionsCountryLv  CrawlRootParamsExtractOptionsCountry = "LV"
	CrawlRootParamsExtractOptionsCountryLy  CrawlRootParamsExtractOptionsCountry = "LY"
	CrawlRootParamsExtractOptionsCountryMa  CrawlRootParamsExtractOptionsCountry = "MA"
	CrawlRootParamsExtractOptionsCountryMc  CrawlRootParamsExtractOptionsCountry = "MC"
	CrawlRootParamsExtractOptionsCountryMd  CrawlRootParamsExtractOptionsCountry = "MD"
	CrawlRootParamsExtractOptionsCountryMe  CrawlRootParamsExtractOptionsCountry = "ME"
	CrawlRootParamsExtractOptionsCountryMf  CrawlRootParamsExtractOptionsCountry = "MF"
	CrawlRootParamsExtractOptionsCountryMg  CrawlRootParamsExtractOptionsCountry = "MG"
	CrawlRootParamsExtractOptionsCountryMh  CrawlRootParamsExtractOptionsCountry = "MH"
	CrawlRootParamsExtractOptionsCountryMk  CrawlRootParamsExtractOptionsCountry = "MK"
	CrawlRootParamsExtractOptionsCountryMl  CrawlRootParamsExtractOptionsCountry = "ML"
	CrawlRootParamsExtractOptionsCountryMm  CrawlRootParamsExtractOptionsCountry = "MM"
	CrawlRootParamsExtractOptionsCountryMn  CrawlRootParamsExtractOptionsCountry = "MN"
	CrawlRootParamsExtractOptionsCountryMo  CrawlRootParamsExtractOptionsCountry = "MO"
	CrawlRootParamsExtractOptionsCountryMp  CrawlRootParamsExtractOptionsCountry = "MP"
	CrawlRootParamsExtractOptionsCountryMq  CrawlRootParamsExtractOptionsCountry = "MQ"
	CrawlRootParamsExtractOptionsCountryMr  CrawlRootParamsExtractOptionsCountry = "MR"
	CrawlRootParamsExtractOptionsCountryMs  CrawlRootParamsExtractOptionsCountry = "MS"
	CrawlRootParamsExtractOptionsCountryMt  CrawlRootParamsExtractOptionsCountry = "MT"
	CrawlRootParamsExtractOptionsCountryMu  CrawlRootParamsExtractOptionsCountry = "MU"
	CrawlRootParamsExtractOptionsCountryMv  CrawlRootParamsExtractOptionsCountry = "MV"
	CrawlRootParamsExtractOptionsCountryMw  CrawlRootParamsExtractOptionsCountry = "MW"
	CrawlRootParamsExtractOptionsCountryMx  CrawlRootParamsExtractOptionsCountry = "MX"
	CrawlRootParamsExtractOptionsCountryMy  CrawlRootParamsExtractOptionsCountry = "MY"
	CrawlRootParamsExtractOptionsCountryMz  CrawlRootParamsExtractOptionsCountry = "MZ"
	CrawlRootParamsExtractOptionsCountryNa  CrawlRootParamsExtractOptionsCountry = "NA"
	CrawlRootParamsExtractOptionsCountryNc  CrawlRootParamsExtractOptionsCountry = "NC"
	CrawlRootParamsExtractOptionsCountryNe  CrawlRootParamsExtractOptionsCountry = "NE"
	CrawlRootParamsExtractOptionsCountryNf  CrawlRootParamsExtractOptionsCountry = "NF"
	CrawlRootParamsExtractOptionsCountryNg  CrawlRootParamsExtractOptionsCountry = "NG"
	CrawlRootParamsExtractOptionsCountryNi  CrawlRootParamsExtractOptionsCountry = "NI"
	CrawlRootParamsExtractOptionsCountryNl  CrawlRootParamsExtractOptionsCountry = "NL"
	CrawlRootParamsExtractOptionsCountryNo  CrawlRootParamsExtractOptionsCountry = "NO"
	CrawlRootParamsExtractOptionsCountryNp  CrawlRootParamsExtractOptionsCountry = "NP"
	CrawlRootParamsExtractOptionsCountryNr  CrawlRootParamsExtractOptionsCountry = "NR"
	CrawlRootParamsExtractOptionsCountryNu  CrawlRootParamsExtractOptionsCountry = "NU"
	CrawlRootParamsExtractOptionsCountryNz  CrawlRootParamsExtractOptionsCountry = "NZ"
	CrawlRootParamsExtractOptionsCountryOm  CrawlRootParamsExtractOptionsCountry = "OM"
	CrawlRootParamsExtractOptionsCountryPa  CrawlRootParamsExtractOptionsCountry = "PA"
	CrawlRootParamsExtractOptionsCountryPe  CrawlRootParamsExtractOptionsCountry = "PE"
	CrawlRootParamsExtractOptionsCountryPf  CrawlRootParamsExtractOptionsCountry = "PF"
	CrawlRootParamsExtractOptionsCountryPg  CrawlRootParamsExtractOptionsCountry = "PG"
	CrawlRootParamsExtractOptionsCountryPh  CrawlRootParamsExtractOptionsCountry = "PH"
	CrawlRootParamsExtractOptionsCountryPk  CrawlRootParamsExtractOptionsCountry = "PK"
	CrawlRootParamsExtractOptionsCountryPl  CrawlRootParamsExtractOptionsCountry = "PL"
	CrawlRootParamsExtractOptionsCountryPm  CrawlRootParamsExtractOptionsCountry = "PM"
	CrawlRootParamsExtractOptionsCountryPn  CrawlRootParamsExtractOptionsCountry = "PN"
	CrawlRootParamsExtractOptionsCountryPr  CrawlRootParamsExtractOptionsCountry = "PR"
	CrawlRootParamsExtractOptionsCountryPs  CrawlRootParamsExtractOptionsCountry = "PS"
	CrawlRootParamsExtractOptionsCountryPt  CrawlRootParamsExtractOptionsCountry = "PT"
	CrawlRootParamsExtractOptionsCountryPw  CrawlRootParamsExtractOptionsCountry = "PW"
	CrawlRootParamsExtractOptionsCountryPy  CrawlRootParamsExtractOptionsCountry = "PY"
	CrawlRootParamsExtractOptionsCountryQa  CrawlRootParamsExtractOptionsCountry = "QA"
	CrawlRootParamsExtractOptionsCountryRe  CrawlRootParamsExtractOptionsCountry = "RE"
	CrawlRootParamsExtractOptionsCountryRo  CrawlRootParamsExtractOptionsCountry = "RO"
	CrawlRootParamsExtractOptionsCountryRs  CrawlRootParamsExtractOptionsCountry = "RS"
	CrawlRootParamsExtractOptionsCountryRu  CrawlRootParamsExtractOptionsCountry = "RU"
	CrawlRootParamsExtractOptionsCountryRw  CrawlRootParamsExtractOptionsCountry = "RW"
	CrawlRootParamsExtractOptionsCountrySa  CrawlRootParamsExtractOptionsCountry = "SA"
	CrawlRootParamsExtractOptionsCountrySb  CrawlRootParamsExtractOptionsCountry = "SB"
	CrawlRootParamsExtractOptionsCountrySc  CrawlRootParamsExtractOptionsCountry = "SC"
	CrawlRootParamsExtractOptionsCountrySd  CrawlRootParamsExtractOptionsCountry = "SD"
	CrawlRootParamsExtractOptionsCountrySe  CrawlRootParamsExtractOptionsCountry = "SE"
	CrawlRootParamsExtractOptionsCountrySg  CrawlRootParamsExtractOptionsCountry = "SG"
	CrawlRootParamsExtractOptionsCountrySh  CrawlRootParamsExtractOptionsCountry = "SH"
	CrawlRootParamsExtractOptionsCountrySi  CrawlRootParamsExtractOptionsCountry = "SI"
	CrawlRootParamsExtractOptionsCountrySj  CrawlRootParamsExtractOptionsCountry = "SJ"
	CrawlRootParamsExtractOptionsCountrySk  CrawlRootParamsExtractOptionsCountry = "SK"
	CrawlRootParamsExtractOptionsCountrySl  CrawlRootParamsExtractOptionsCountry = "SL"
	CrawlRootParamsExtractOptionsCountrySm  CrawlRootParamsExtractOptionsCountry = "SM"
	CrawlRootParamsExtractOptionsCountrySn  CrawlRootParamsExtractOptionsCountry = "SN"
	CrawlRootParamsExtractOptionsCountrySo  CrawlRootParamsExtractOptionsCountry = "SO"
	CrawlRootParamsExtractOptionsCountrySr  CrawlRootParamsExtractOptionsCountry = "SR"
	CrawlRootParamsExtractOptionsCountrySS  CrawlRootParamsExtractOptionsCountry = "SS"
	CrawlRootParamsExtractOptionsCountrySt  CrawlRootParamsExtractOptionsCountry = "ST"
	CrawlRootParamsExtractOptionsCountrySv  CrawlRootParamsExtractOptionsCountry = "SV"
	CrawlRootParamsExtractOptionsCountrySx  CrawlRootParamsExtractOptionsCountry = "SX"
	CrawlRootParamsExtractOptionsCountrySy  CrawlRootParamsExtractOptionsCountry = "SY"
	CrawlRootParamsExtractOptionsCountrySz  CrawlRootParamsExtractOptionsCountry = "SZ"
	CrawlRootParamsExtractOptionsCountryTc  CrawlRootParamsExtractOptionsCountry = "TC"
	CrawlRootParamsExtractOptionsCountryTd  CrawlRootParamsExtractOptionsCountry = "TD"
	CrawlRootParamsExtractOptionsCountryTf  CrawlRootParamsExtractOptionsCountry = "TF"
	CrawlRootParamsExtractOptionsCountryTg  CrawlRootParamsExtractOptionsCountry = "TG"
	CrawlRootParamsExtractOptionsCountryTh  CrawlRootParamsExtractOptionsCountry = "TH"
	CrawlRootParamsExtractOptionsCountryTj  CrawlRootParamsExtractOptionsCountry = "TJ"
	CrawlRootParamsExtractOptionsCountryTk  CrawlRootParamsExtractOptionsCountry = "TK"
	CrawlRootParamsExtractOptionsCountryTl  CrawlRootParamsExtractOptionsCountry = "TL"
	CrawlRootParamsExtractOptionsCountryTm  CrawlRootParamsExtractOptionsCountry = "TM"
	CrawlRootParamsExtractOptionsCountryTn  CrawlRootParamsExtractOptionsCountry = "TN"
	CrawlRootParamsExtractOptionsCountryTo  CrawlRootParamsExtractOptionsCountry = "TO"
	CrawlRootParamsExtractOptionsCountryTr  CrawlRootParamsExtractOptionsCountry = "TR"
	CrawlRootParamsExtractOptionsCountryTt  CrawlRootParamsExtractOptionsCountry = "TT"
	CrawlRootParamsExtractOptionsCountryTv  CrawlRootParamsExtractOptionsCountry = "TV"
	CrawlRootParamsExtractOptionsCountryTw  CrawlRootParamsExtractOptionsCountry = "TW"
	CrawlRootParamsExtractOptionsCountryTz  CrawlRootParamsExtractOptionsCountry = "TZ"
	CrawlRootParamsExtractOptionsCountryUa  CrawlRootParamsExtractOptionsCountry = "UA"
	CrawlRootParamsExtractOptionsCountryUg  CrawlRootParamsExtractOptionsCountry = "UG"
	CrawlRootParamsExtractOptionsCountryUm  CrawlRootParamsExtractOptionsCountry = "UM"
	CrawlRootParamsExtractOptionsCountryUs  CrawlRootParamsExtractOptionsCountry = "US"
	CrawlRootParamsExtractOptionsCountryUy  CrawlRootParamsExtractOptionsCountry = "UY"
	CrawlRootParamsExtractOptionsCountryUz  CrawlRootParamsExtractOptionsCountry = "UZ"
	CrawlRootParamsExtractOptionsCountryVa  CrawlRootParamsExtractOptionsCountry = "VA"
	CrawlRootParamsExtractOptionsCountryVc  CrawlRootParamsExtractOptionsCountry = "VC"
	CrawlRootParamsExtractOptionsCountryVe  CrawlRootParamsExtractOptionsCountry = "VE"
	CrawlRootParamsExtractOptionsCountryVg  CrawlRootParamsExtractOptionsCountry = "VG"
	CrawlRootParamsExtractOptionsCountryVi  CrawlRootParamsExtractOptionsCountry = "VI"
	CrawlRootParamsExtractOptionsCountryVn  CrawlRootParamsExtractOptionsCountry = "VN"
	CrawlRootParamsExtractOptionsCountryVu  CrawlRootParamsExtractOptionsCountry = "VU"
	CrawlRootParamsExtractOptionsCountryWf  CrawlRootParamsExtractOptionsCountry = "WF"
	CrawlRootParamsExtractOptionsCountryWs  CrawlRootParamsExtractOptionsCountry = "WS"
	CrawlRootParamsExtractOptionsCountryXk  CrawlRootParamsExtractOptionsCountry = "XK"
	CrawlRootParamsExtractOptionsCountryYe  CrawlRootParamsExtractOptionsCountry = "YE"
	CrawlRootParamsExtractOptionsCountryYt  CrawlRootParamsExtractOptionsCountry = "YT"
	CrawlRootParamsExtractOptionsCountryZa  CrawlRootParamsExtractOptionsCountry = "ZA"
	CrawlRootParamsExtractOptionsCountryZm  CrawlRootParamsExtractOptionsCountry = "ZM"
	CrawlRootParamsExtractOptionsCountryZw  CrawlRootParamsExtractOptionsCountry = "ZW"
	CrawlRootParamsExtractOptionsCountryAll CrawlRootParamsExtractOptionsCountry = "ALL"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsHeaderUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsHeaderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlRootParamsExtractOptionsHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsHeaderUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Locale for browser language and region settings
type CrawlRootParamsExtractOptionsLocale string

const (
	CrawlRootParamsExtractOptionsLocaleAaDj      CrawlRootParamsExtractOptionsLocale = "aa-DJ"
	CrawlRootParamsExtractOptionsLocaleAaEr      CrawlRootParamsExtractOptionsLocale = "aa-ER"
	CrawlRootParamsExtractOptionsLocaleAaEt      CrawlRootParamsExtractOptionsLocale = "aa-ET"
	CrawlRootParamsExtractOptionsLocaleAf        CrawlRootParamsExtractOptionsLocale = "af"
	CrawlRootParamsExtractOptionsLocaleAfNa      CrawlRootParamsExtractOptionsLocale = "af-NA"
	CrawlRootParamsExtractOptionsLocaleAfZa      CrawlRootParamsExtractOptionsLocale = "af-ZA"
	CrawlRootParamsExtractOptionsLocaleAk        CrawlRootParamsExtractOptionsLocale = "ak"
	CrawlRootParamsExtractOptionsLocaleAkGh      CrawlRootParamsExtractOptionsLocale = "ak-GH"
	CrawlRootParamsExtractOptionsLocaleAm        CrawlRootParamsExtractOptionsLocale = "am"
	CrawlRootParamsExtractOptionsLocaleAmEt      CrawlRootParamsExtractOptionsLocale = "am-ET"
	CrawlRootParamsExtractOptionsLocaleAnEs      CrawlRootParamsExtractOptionsLocale = "an-ES"
	CrawlRootParamsExtractOptionsLocaleAr        CrawlRootParamsExtractOptionsLocale = "ar"
	CrawlRootParamsExtractOptionsLocaleArAe      CrawlRootParamsExtractOptionsLocale = "ar-AE"
	CrawlRootParamsExtractOptionsLocaleArBh      CrawlRootParamsExtractOptionsLocale = "ar-BH"
	CrawlRootParamsExtractOptionsLocaleArDz      CrawlRootParamsExtractOptionsLocale = "ar-DZ"
	CrawlRootParamsExtractOptionsLocaleArEg      CrawlRootParamsExtractOptionsLocale = "ar-EG"
	CrawlRootParamsExtractOptionsLocaleArIn      CrawlRootParamsExtractOptionsLocale = "ar-IN"
	CrawlRootParamsExtractOptionsLocaleArIq      CrawlRootParamsExtractOptionsLocale = "ar-IQ"
	CrawlRootParamsExtractOptionsLocaleArJo      CrawlRootParamsExtractOptionsLocale = "ar-JO"
	CrawlRootParamsExtractOptionsLocaleArKw      CrawlRootParamsExtractOptionsLocale = "ar-KW"
	CrawlRootParamsExtractOptionsLocaleArLb      CrawlRootParamsExtractOptionsLocale = "ar-LB"
	CrawlRootParamsExtractOptionsLocaleArLy      CrawlRootParamsExtractOptionsLocale = "ar-LY"
	CrawlRootParamsExtractOptionsLocaleArMa      CrawlRootParamsExtractOptionsLocale = "ar-MA"
	CrawlRootParamsExtractOptionsLocaleArOm      CrawlRootParamsExtractOptionsLocale = "ar-OM"
	CrawlRootParamsExtractOptionsLocaleArQa      CrawlRootParamsExtractOptionsLocale = "ar-QA"
	CrawlRootParamsExtractOptionsLocaleArSa      CrawlRootParamsExtractOptionsLocale = "ar-SA"
	CrawlRootParamsExtractOptionsLocaleArSd      CrawlRootParamsExtractOptionsLocale = "ar-SD"
	CrawlRootParamsExtractOptionsLocaleArSy      CrawlRootParamsExtractOptionsLocale = "ar-SY"
	CrawlRootParamsExtractOptionsLocaleArTn      CrawlRootParamsExtractOptionsLocale = "ar-TN"
	CrawlRootParamsExtractOptionsLocaleArYe      CrawlRootParamsExtractOptionsLocale = "ar-YE"
	CrawlRootParamsExtractOptionsLocaleAs        CrawlRootParamsExtractOptionsLocale = "as"
	CrawlRootParamsExtractOptionsLocaleAsIn      CrawlRootParamsExtractOptionsLocale = "as-IN"
	CrawlRootParamsExtractOptionsLocaleAsa       CrawlRootParamsExtractOptionsLocale = "asa"
	CrawlRootParamsExtractOptionsLocaleAsaTz     CrawlRootParamsExtractOptionsLocale = "asa-TZ"
	CrawlRootParamsExtractOptionsLocaleAstEs     CrawlRootParamsExtractOptionsLocale = "ast-ES"
	CrawlRootParamsExtractOptionsLocaleAz        CrawlRootParamsExtractOptionsLocale = "az"
	CrawlRootParamsExtractOptionsLocaleAzAz      CrawlRootParamsExtractOptionsLocale = "az-AZ"
	CrawlRootParamsExtractOptionsLocaleAzCyrl    CrawlRootParamsExtractOptionsLocale = "az-Cyrl"
	CrawlRootParamsExtractOptionsLocaleAzCyrlAz  CrawlRootParamsExtractOptionsLocale = "az-Cyrl-AZ"
	CrawlRootParamsExtractOptionsLocaleAzLatn    CrawlRootParamsExtractOptionsLocale = "az-Latn"
	CrawlRootParamsExtractOptionsLocaleAzLatnAz  CrawlRootParamsExtractOptionsLocale = "az-Latn-AZ"
	CrawlRootParamsExtractOptionsLocaleBe        CrawlRootParamsExtractOptionsLocale = "be"
	CrawlRootParamsExtractOptionsLocaleBeBy      CrawlRootParamsExtractOptionsLocale = "be-BY"
	CrawlRootParamsExtractOptionsLocaleBem       CrawlRootParamsExtractOptionsLocale = "bem"
	CrawlRootParamsExtractOptionsLocaleBemZm     CrawlRootParamsExtractOptionsLocale = "bem-ZM"
	CrawlRootParamsExtractOptionsLocaleBerDz     CrawlRootParamsExtractOptionsLocale = "ber-DZ"
	CrawlRootParamsExtractOptionsLocaleBerMa     CrawlRootParamsExtractOptionsLocale = "ber-MA"
	CrawlRootParamsExtractOptionsLocaleBez       CrawlRootParamsExtractOptionsLocale = "bez"
	CrawlRootParamsExtractOptionsLocaleBezTz     CrawlRootParamsExtractOptionsLocale = "bez-TZ"
	CrawlRootParamsExtractOptionsLocaleBg        CrawlRootParamsExtractOptionsLocale = "bg"
	CrawlRootParamsExtractOptionsLocaleBgBg      CrawlRootParamsExtractOptionsLocale = "bg-BG"
	CrawlRootParamsExtractOptionsLocaleBhoIn     CrawlRootParamsExtractOptionsLocale = "bho-IN"
	CrawlRootParamsExtractOptionsLocaleBm        CrawlRootParamsExtractOptionsLocale = "bm"
	CrawlRootParamsExtractOptionsLocaleBmMl      CrawlRootParamsExtractOptionsLocale = "bm-ML"
	CrawlRootParamsExtractOptionsLocaleBn        CrawlRootParamsExtractOptionsLocale = "bn"
	CrawlRootParamsExtractOptionsLocaleBnBd      CrawlRootParamsExtractOptionsLocale = "bn-BD"
	CrawlRootParamsExtractOptionsLocaleBnIn      CrawlRootParamsExtractOptionsLocale = "bn-IN"
	CrawlRootParamsExtractOptionsLocaleBo        CrawlRootParamsExtractOptionsLocale = "bo"
	CrawlRootParamsExtractOptionsLocaleBoCn      CrawlRootParamsExtractOptionsLocale = "bo-CN"
	CrawlRootParamsExtractOptionsLocaleBoIn      CrawlRootParamsExtractOptionsLocale = "bo-IN"
	CrawlRootParamsExtractOptionsLocaleBrFr      CrawlRootParamsExtractOptionsLocale = "br-FR"
	CrawlRootParamsExtractOptionsLocaleBrxIn     CrawlRootParamsExtractOptionsLocale = "brx-IN"
	CrawlRootParamsExtractOptionsLocaleBs        CrawlRootParamsExtractOptionsLocale = "bs"
	CrawlRootParamsExtractOptionsLocaleBsBa      CrawlRootParamsExtractOptionsLocale = "bs-BA"
	CrawlRootParamsExtractOptionsLocaleBynEr     CrawlRootParamsExtractOptionsLocale = "byn-ER"
	CrawlRootParamsExtractOptionsLocaleCa        CrawlRootParamsExtractOptionsLocale = "ca"
	CrawlRootParamsExtractOptionsLocaleCaAd      CrawlRootParamsExtractOptionsLocale = "ca-AD"
	CrawlRootParamsExtractOptionsLocaleCaEs      CrawlRootParamsExtractOptionsLocale = "ca-ES"
	CrawlRootParamsExtractOptionsLocaleCaFr      CrawlRootParamsExtractOptionsLocale = "ca-FR"
	CrawlRootParamsExtractOptionsLocaleCaIt      CrawlRootParamsExtractOptionsLocale = "ca-IT"
	CrawlRootParamsExtractOptionsLocaleCgg       CrawlRootParamsExtractOptionsLocale = "cgg"
	CrawlRootParamsExtractOptionsLocaleCggUg     CrawlRootParamsExtractOptionsLocale = "cgg-UG"
	CrawlRootParamsExtractOptionsLocaleChr       CrawlRootParamsExtractOptionsLocale = "chr"
	CrawlRootParamsExtractOptionsLocaleChrUs     CrawlRootParamsExtractOptionsLocale = "chr-US"
	CrawlRootParamsExtractOptionsLocaleCrhUa     CrawlRootParamsExtractOptionsLocale = "crh-UA"
	CrawlRootParamsExtractOptionsLocaleCs        CrawlRootParamsExtractOptionsLocale = "cs"
	CrawlRootParamsExtractOptionsLocaleCsCz      CrawlRootParamsExtractOptionsLocale = "cs-CZ"
	CrawlRootParamsExtractOptionsLocaleCsbPl     CrawlRootParamsExtractOptionsLocale = "csb-PL"
	CrawlRootParamsExtractOptionsLocaleCvRu      CrawlRootParamsExtractOptionsLocale = "cv-RU"
	CrawlRootParamsExtractOptionsLocaleCy        CrawlRootParamsExtractOptionsLocale = "cy"
	CrawlRootParamsExtractOptionsLocaleCyGB      CrawlRootParamsExtractOptionsLocale = "cy-GB"
	CrawlRootParamsExtractOptionsLocaleDa        CrawlRootParamsExtractOptionsLocale = "da"
	CrawlRootParamsExtractOptionsLocaleDaDk      CrawlRootParamsExtractOptionsLocale = "da-DK"
	CrawlRootParamsExtractOptionsLocaleDav       CrawlRootParamsExtractOptionsLocale = "dav"
	CrawlRootParamsExtractOptionsLocaleDavKe     CrawlRootParamsExtractOptionsLocale = "dav-KE"
	CrawlRootParamsExtractOptionsLocaleDe        CrawlRootParamsExtractOptionsLocale = "de"
	CrawlRootParamsExtractOptionsLocaleDeAt      CrawlRootParamsExtractOptionsLocale = "de-AT"
	CrawlRootParamsExtractOptionsLocaleDeBe      CrawlRootParamsExtractOptionsLocale = "de-BE"
	CrawlRootParamsExtractOptionsLocaleDeCh      CrawlRootParamsExtractOptionsLocale = "de-CH"
	CrawlRootParamsExtractOptionsLocaleDeDe      CrawlRootParamsExtractOptionsLocale = "de-DE"
	CrawlRootParamsExtractOptionsLocaleDeLi      CrawlRootParamsExtractOptionsLocale = "de-LI"
	CrawlRootParamsExtractOptionsLocaleDeLu      CrawlRootParamsExtractOptionsLocale = "de-LU"
	CrawlRootParamsExtractOptionsLocaleDvMv      CrawlRootParamsExtractOptionsLocale = "dv-MV"
	CrawlRootParamsExtractOptionsLocaleDzBt      CrawlRootParamsExtractOptionsLocale = "dz-BT"
	CrawlRootParamsExtractOptionsLocaleEbu       CrawlRootParamsExtractOptionsLocale = "ebu"
	CrawlRootParamsExtractOptionsLocaleEbuKe     CrawlRootParamsExtractOptionsLocale = "ebu-KE"
	CrawlRootParamsExtractOptionsLocaleEe        CrawlRootParamsExtractOptionsLocale = "ee"
	CrawlRootParamsExtractOptionsLocaleEeGh      CrawlRootParamsExtractOptionsLocale = "ee-GH"
	CrawlRootParamsExtractOptionsLocaleEeTg      CrawlRootParamsExtractOptionsLocale = "ee-TG"
	CrawlRootParamsExtractOptionsLocaleEl        CrawlRootParamsExtractOptionsLocale = "el"
	CrawlRootParamsExtractOptionsLocaleElCy      CrawlRootParamsExtractOptionsLocale = "el-CY"
	CrawlRootParamsExtractOptionsLocaleElGr      CrawlRootParamsExtractOptionsLocale = "el-GR"
	CrawlRootParamsExtractOptionsLocaleEn        CrawlRootParamsExtractOptionsLocale = "en"
	CrawlRootParamsExtractOptionsLocaleEnAg      CrawlRootParamsExtractOptionsLocale = "en-AG"
	CrawlRootParamsExtractOptionsLocaleEnAs      CrawlRootParamsExtractOptionsLocale = "en-AS"
	CrawlRootParamsExtractOptionsLocaleEnAu      CrawlRootParamsExtractOptionsLocale = "en-AU"
	CrawlRootParamsExtractOptionsLocaleEnBe      CrawlRootParamsExtractOptionsLocale = "en-BE"
	CrawlRootParamsExtractOptionsLocaleEnBw      CrawlRootParamsExtractOptionsLocale = "en-BW"
	CrawlRootParamsExtractOptionsLocaleEnBz      CrawlRootParamsExtractOptionsLocale = "en-BZ"
	CrawlRootParamsExtractOptionsLocaleEnCa      CrawlRootParamsExtractOptionsLocale = "en-CA"
	CrawlRootParamsExtractOptionsLocaleEnDk      CrawlRootParamsExtractOptionsLocale = "en-DK"
	CrawlRootParamsExtractOptionsLocaleEnGB      CrawlRootParamsExtractOptionsLocale = "en-GB"
	CrawlRootParamsExtractOptionsLocaleEnGu      CrawlRootParamsExtractOptionsLocale = "en-GU"
	CrawlRootParamsExtractOptionsLocaleEnHk      CrawlRootParamsExtractOptionsLocale = "en-HK"
	CrawlRootParamsExtractOptionsLocaleEnIe      CrawlRootParamsExtractOptionsLocale = "en-IE"
	CrawlRootParamsExtractOptionsLocaleEnIn      CrawlRootParamsExtractOptionsLocale = "en-IN"
	CrawlRootParamsExtractOptionsLocaleEnJm      CrawlRootParamsExtractOptionsLocale = "en-JM"
	CrawlRootParamsExtractOptionsLocaleEnMh      CrawlRootParamsExtractOptionsLocale = "en-MH"
	CrawlRootParamsExtractOptionsLocaleEnMp      CrawlRootParamsExtractOptionsLocale = "en-MP"
	CrawlRootParamsExtractOptionsLocaleEnMt      CrawlRootParamsExtractOptionsLocale = "en-MT"
	CrawlRootParamsExtractOptionsLocaleEnMu      CrawlRootParamsExtractOptionsLocale = "en-MU"
	CrawlRootParamsExtractOptionsLocaleEnNa      CrawlRootParamsExtractOptionsLocale = "en-NA"
	CrawlRootParamsExtractOptionsLocaleEnNg      CrawlRootParamsExtractOptionsLocale = "en-NG"
	CrawlRootParamsExtractOptionsLocaleEnNz      CrawlRootParamsExtractOptionsLocale = "en-NZ"
	CrawlRootParamsExtractOptionsLocaleEnPh      CrawlRootParamsExtractOptionsLocale = "en-PH"
	CrawlRootParamsExtractOptionsLocaleEnPk      CrawlRootParamsExtractOptionsLocale = "en-PK"
	CrawlRootParamsExtractOptionsLocaleEnSg      CrawlRootParamsExtractOptionsLocale = "en-SG"
	CrawlRootParamsExtractOptionsLocaleEnTt      CrawlRootParamsExtractOptionsLocale = "en-TT"
	CrawlRootParamsExtractOptionsLocaleEnUm      CrawlRootParamsExtractOptionsLocale = "en-UM"
	CrawlRootParamsExtractOptionsLocaleEnUs      CrawlRootParamsExtractOptionsLocale = "en-US"
	CrawlRootParamsExtractOptionsLocaleEnVi      CrawlRootParamsExtractOptionsLocale = "en-VI"
	CrawlRootParamsExtractOptionsLocaleEnZa      CrawlRootParamsExtractOptionsLocale = "en-ZA"
	CrawlRootParamsExtractOptionsLocaleEnZm      CrawlRootParamsExtractOptionsLocale = "en-ZM"
	CrawlRootParamsExtractOptionsLocaleEnZw      CrawlRootParamsExtractOptionsLocale = "en-ZW"
	CrawlRootParamsExtractOptionsLocaleEo        CrawlRootParamsExtractOptionsLocale = "eo"
	CrawlRootParamsExtractOptionsLocaleEs        CrawlRootParamsExtractOptionsLocale = "es"
	CrawlRootParamsExtractOptionsLocaleEs419     CrawlRootParamsExtractOptionsLocale = "es-419"
	CrawlRootParamsExtractOptionsLocaleEsAr      CrawlRootParamsExtractOptionsLocale = "es-AR"
	CrawlRootParamsExtractOptionsLocaleEsBo      CrawlRootParamsExtractOptionsLocale = "es-BO"
	CrawlRootParamsExtractOptionsLocaleEsCl      CrawlRootParamsExtractOptionsLocale = "es-CL"
	CrawlRootParamsExtractOptionsLocaleEsCo      CrawlRootParamsExtractOptionsLocale = "es-CO"
	CrawlRootParamsExtractOptionsLocaleEsCr      CrawlRootParamsExtractOptionsLocale = "es-CR"
	CrawlRootParamsExtractOptionsLocaleEsCu      CrawlRootParamsExtractOptionsLocale = "es-CU"
	CrawlRootParamsExtractOptionsLocaleEsDo      CrawlRootParamsExtractOptionsLocale = "es-DO"
	CrawlRootParamsExtractOptionsLocaleEsEc      CrawlRootParamsExtractOptionsLocale = "es-EC"
	CrawlRootParamsExtractOptionsLocaleEsEs      CrawlRootParamsExtractOptionsLocale = "es-ES"
	CrawlRootParamsExtractOptionsLocaleEsGq      CrawlRootParamsExtractOptionsLocale = "es-GQ"
	CrawlRootParamsExtractOptionsLocaleEsGt      CrawlRootParamsExtractOptionsLocale = "es-GT"
	CrawlRootParamsExtractOptionsLocaleEsHn      CrawlRootParamsExtractOptionsLocale = "es-HN"
	CrawlRootParamsExtractOptionsLocaleEsMx      CrawlRootParamsExtractOptionsLocale = "es-MX"
	CrawlRootParamsExtractOptionsLocaleEsNi      CrawlRootParamsExtractOptionsLocale = "es-NI"
	CrawlRootParamsExtractOptionsLocaleEsPa      CrawlRootParamsExtractOptionsLocale = "es-PA"
	CrawlRootParamsExtractOptionsLocaleEsPe      CrawlRootParamsExtractOptionsLocale = "es-PE"
	CrawlRootParamsExtractOptionsLocaleEsPr      CrawlRootParamsExtractOptionsLocale = "es-PR"
	CrawlRootParamsExtractOptionsLocaleEsPy      CrawlRootParamsExtractOptionsLocale = "es-PY"
	CrawlRootParamsExtractOptionsLocaleEsSv      CrawlRootParamsExtractOptionsLocale = "es-SV"
	CrawlRootParamsExtractOptionsLocaleEsUs      CrawlRootParamsExtractOptionsLocale = "es-US"
	CrawlRootParamsExtractOptionsLocaleEsUy      CrawlRootParamsExtractOptionsLocale = "es-UY"
	CrawlRootParamsExtractOptionsLocaleEsVe      CrawlRootParamsExtractOptionsLocale = "es-VE"
	CrawlRootParamsExtractOptionsLocaleEt        CrawlRootParamsExtractOptionsLocale = "et"
	CrawlRootParamsExtractOptionsLocaleEtEe      CrawlRootParamsExtractOptionsLocale = "et-EE"
	CrawlRootParamsExtractOptionsLocaleEu        CrawlRootParamsExtractOptionsLocale = "eu"
	CrawlRootParamsExtractOptionsLocaleEuEs      CrawlRootParamsExtractOptionsLocale = "eu-ES"
	CrawlRootParamsExtractOptionsLocaleFa        CrawlRootParamsExtractOptionsLocale = "fa"
	CrawlRootParamsExtractOptionsLocaleFaAf      CrawlRootParamsExtractOptionsLocale = "fa-AF"
	CrawlRootParamsExtractOptionsLocaleFaIr      CrawlRootParamsExtractOptionsLocale = "fa-IR"
	CrawlRootParamsExtractOptionsLocaleFf        CrawlRootParamsExtractOptionsLocale = "ff"
	CrawlRootParamsExtractOptionsLocaleFfSn      CrawlRootParamsExtractOptionsLocale = "ff-SN"
	CrawlRootParamsExtractOptionsLocaleFi        CrawlRootParamsExtractOptionsLocale = "fi"
	CrawlRootParamsExtractOptionsLocaleFiFi      CrawlRootParamsExtractOptionsLocale = "fi-FI"
	CrawlRootParamsExtractOptionsLocaleFil       CrawlRootParamsExtractOptionsLocale = "fil"
	CrawlRootParamsExtractOptionsLocaleFilPh     CrawlRootParamsExtractOptionsLocale = "fil-PH"
	CrawlRootParamsExtractOptionsLocaleFo        CrawlRootParamsExtractOptionsLocale = "fo"
	CrawlRootParamsExtractOptionsLocaleFoFo      CrawlRootParamsExtractOptionsLocale = "fo-FO"
	CrawlRootParamsExtractOptionsLocaleFr        CrawlRootParamsExtractOptionsLocale = "fr"
	CrawlRootParamsExtractOptionsLocaleFrBe      CrawlRootParamsExtractOptionsLocale = "fr-BE"
	CrawlRootParamsExtractOptionsLocaleFrBf      CrawlRootParamsExtractOptionsLocale = "fr-BF"
	CrawlRootParamsExtractOptionsLocaleFrBi      CrawlRootParamsExtractOptionsLocale = "fr-BI"
	CrawlRootParamsExtractOptionsLocaleFrBj      CrawlRootParamsExtractOptionsLocale = "fr-BJ"
	CrawlRootParamsExtractOptionsLocaleFrBl      CrawlRootParamsExtractOptionsLocale = "fr-BL"
	CrawlRootParamsExtractOptionsLocaleFrCa      CrawlRootParamsExtractOptionsLocale = "fr-CA"
	CrawlRootParamsExtractOptionsLocaleFrCd      CrawlRootParamsExtractOptionsLocale = "fr-CD"
	CrawlRootParamsExtractOptionsLocaleFrCf      CrawlRootParamsExtractOptionsLocale = "fr-CF"
	CrawlRootParamsExtractOptionsLocaleFrCg      CrawlRootParamsExtractOptionsLocale = "fr-CG"
	CrawlRootParamsExtractOptionsLocaleFrCh      CrawlRootParamsExtractOptionsLocale = "fr-CH"
	CrawlRootParamsExtractOptionsLocaleFrCi      CrawlRootParamsExtractOptionsLocale = "fr-CI"
	CrawlRootParamsExtractOptionsLocaleFrCm      CrawlRootParamsExtractOptionsLocale = "fr-CM"
	CrawlRootParamsExtractOptionsLocaleFrDj      CrawlRootParamsExtractOptionsLocale = "fr-DJ"
	CrawlRootParamsExtractOptionsLocaleFrFr      CrawlRootParamsExtractOptionsLocale = "fr-FR"
	CrawlRootParamsExtractOptionsLocaleFrGa      CrawlRootParamsExtractOptionsLocale = "fr-GA"
	CrawlRootParamsExtractOptionsLocaleFrGn      CrawlRootParamsExtractOptionsLocale = "fr-GN"
	CrawlRootParamsExtractOptionsLocaleFrGp      CrawlRootParamsExtractOptionsLocale = "fr-GP"
	CrawlRootParamsExtractOptionsLocaleFrGq      CrawlRootParamsExtractOptionsLocale = "fr-GQ"
	CrawlRootParamsExtractOptionsLocaleFrKm      CrawlRootParamsExtractOptionsLocale = "fr-KM"
	CrawlRootParamsExtractOptionsLocaleFrLu      CrawlRootParamsExtractOptionsLocale = "fr-LU"
	CrawlRootParamsExtractOptionsLocaleFrMc      CrawlRootParamsExtractOptionsLocale = "fr-MC"
	CrawlRootParamsExtractOptionsLocaleFrMf      CrawlRootParamsExtractOptionsLocale = "fr-MF"
	CrawlRootParamsExtractOptionsLocaleFrMg      CrawlRootParamsExtractOptionsLocale = "fr-MG"
	CrawlRootParamsExtractOptionsLocaleFrMl      CrawlRootParamsExtractOptionsLocale = "fr-ML"
	CrawlRootParamsExtractOptionsLocaleFrMq      CrawlRootParamsExtractOptionsLocale = "fr-MQ"
	CrawlRootParamsExtractOptionsLocaleFrNe      CrawlRootParamsExtractOptionsLocale = "fr-NE"
	CrawlRootParamsExtractOptionsLocaleFrRe      CrawlRootParamsExtractOptionsLocale = "fr-RE"
	CrawlRootParamsExtractOptionsLocaleFrRw      CrawlRootParamsExtractOptionsLocale = "fr-RW"
	CrawlRootParamsExtractOptionsLocaleFrSn      CrawlRootParamsExtractOptionsLocale = "fr-SN"
	CrawlRootParamsExtractOptionsLocaleFrTd      CrawlRootParamsExtractOptionsLocale = "fr-TD"
	CrawlRootParamsExtractOptionsLocaleFrTg      CrawlRootParamsExtractOptionsLocale = "fr-TG"
	CrawlRootParamsExtractOptionsLocaleFurIt     CrawlRootParamsExtractOptionsLocale = "fur-IT"
	CrawlRootParamsExtractOptionsLocaleFyDe      CrawlRootParamsExtractOptionsLocale = "fy-DE"
	CrawlRootParamsExtractOptionsLocaleFyNl      CrawlRootParamsExtractOptionsLocale = "fy-NL"
	CrawlRootParamsExtractOptionsLocaleGa        CrawlRootParamsExtractOptionsLocale = "ga"
	CrawlRootParamsExtractOptionsLocaleGaIe      CrawlRootParamsExtractOptionsLocale = "ga-IE"
	CrawlRootParamsExtractOptionsLocaleGdGB      CrawlRootParamsExtractOptionsLocale = "gd-GB"
	CrawlRootParamsExtractOptionsLocaleGezEr     CrawlRootParamsExtractOptionsLocale = "gez-ER"
	CrawlRootParamsExtractOptionsLocaleGezEt     CrawlRootParamsExtractOptionsLocale = "gez-ET"
	CrawlRootParamsExtractOptionsLocaleGl        CrawlRootParamsExtractOptionsLocale = "gl"
	CrawlRootParamsExtractOptionsLocaleGlEs      CrawlRootParamsExtractOptionsLocale = "gl-ES"
	CrawlRootParamsExtractOptionsLocaleGsw       CrawlRootParamsExtractOptionsLocale = "gsw"
	CrawlRootParamsExtractOptionsLocaleGswCh     CrawlRootParamsExtractOptionsLocale = "gsw-CH"
	CrawlRootParamsExtractOptionsLocaleGu        CrawlRootParamsExtractOptionsLocale = "gu"
	CrawlRootParamsExtractOptionsLocaleGuIn      CrawlRootParamsExtractOptionsLocale = "gu-IN"
	CrawlRootParamsExtractOptionsLocaleGuz       CrawlRootParamsExtractOptionsLocale = "guz"
	CrawlRootParamsExtractOptionsLocaleGuzKe     CrawlRootParamsExtractOptionsLocale = "guz-KE"
	CrawlRootParamsExtractOptionsLocaleGv        CrawlRootParamsExtractOptionsLocale = "gv"
	CrawlRootParamsExtractOptionsLocaleGvGB      CrawlRootParamsExtractOptionsLocale = "gv-GB"
	CrawlRootParamsExtractOptionsLocaleHa        CrawlRootParamsExtractOptionsLocale = "ha"
	CrawlRootParamsExtractOptionsLocaleHaLatn    CrawlRootParamsExtractOptionsLocale = "ha-Latn"
	CrawlRootParamsExtractOptionsLocaleHaLatnGh  CrawlRootParamsExtractOptionsLocale = "ha-Latn-GH"
	CrawlRootParamsExtractOptionsLocaleHaLatnNe  CrawlRootParamsExtractOptionsLocale = "ha-Latn-NE"
	CrawlRootParamsExtractOptionsLocaleHaLatnNg  CrawlRootParamsExtractOptionsLocale = "ha-Latn-NG"
	CrawlRootParamsExtractOptionsLocaleHaNg      CrawlRootParamsExtractOptionsLocale = "ha-NG"
	CrawlRootParamsExtractOptionsLocaleHaw       CrawlRootParamsExtractOptionsLocale = "haw"
	CrawlRootParamsExtractOptionsLocaleHawUs     CrawlRootParamsExtractOptionsLocale = "haw-US"
	CrawlRootParamsExtractOptionsLocaleHe        CrawlRootParamsExtractOptionsLocale = "he"
	CrawlRootParamsExtractOptionsLocaleHeIl      CrawlRootParamsExtractOptionsLocale = "he-IL"
	CrawlRootParamsExtractOptionsLocaleHi        CrawlRootParamsExtractOptionsLocale = "hi"
	CrawlRootParamsExtractOptionsLocaleHiIn      CrawlRootParamsExtractOptionsLocale = "hi-IN"
	CrawlRootParamsExtractOptionsLocaleHneIn     CrawlRootParamsExtractOptionsLocale = "hne-IN"
	CrawlRootParamsExtractOptionsLocaleHr        CrawlRootParamsExtractOptionsLocale = "hr"
	CrawlRootParamsExtractOptionsLocaleHrHr      CrawlRootParamsExtractOptionsLocale = "hr-HR"
	CrawlRootParamsExtractOptionsLocaleHsbDe     CrawlRootParamsExtractOptionsLocale = "hsb-DE"
	CrawlRootParamsExtractOptionsLocaleHtHt      CrawlRootParamsExtractOptionsLocale = "ht-HT"
	CrawlRootParamsExtractOptionsLocaleHu        CrawlRootParamsExtractOptionsLocale = "hu"
	CrawlRootParamsExtractOptionsLocaleHuHu      CrawlRootParamsExtractOptionsLocale = "hu-HU"
	CrawlRootParamsExtractOptionsLocaleHy        CrawlRootParamsExtractOptionsLocale = "hy"
	CrawlRootParamsExtractOptionsLocaleHyAm      CrawlRootParamsExtractOptionsLocale = "hy-AM"
	CrawlRootParamsExtractOptionsLocaleID        CrawlRootParamsExtractOptionsLocale = "id"
	CrawlRootParamsExtractOptionsLocaleIDID      CrawlRootParamsExtractOptionsLocale = "id-ID"
	CrawlRootParamsExtractOptionsLocaleIg        CrawlRootParamsExtractOptionsLocale = "ig"
	CrawlRootParamsExtractOptionsLocaleIgNg      CrawlRootParamsExtractOptionsLocale = "ig-NG"
	CrawlRootParamsExtractOptionsLocaleIi        CrawlRootParamsExtractOptionsLocale = "ii"
	CrawlRootParamsExtractOptionsLocaleIiCn      CrawlRootParamsExtractOptionsLocale = "ii-CN"
	CrawlRootParamsExtractOptionsLocaleIkCa      CrawlRootParamsExtractOptionsLocale = "ik-CA"
	CrawlRootParamsExtractOptionsLocaleIs        CrawlRootParamsExtractOptionsLocale = "is"
	CrawlRootParamsExtractOptionsLocaleIsIs      CrawlRootParamsExtractOptionsLocale = "is-IS"
	CrawlRootParamsExtractOptionsLocaleIt        CrawlRootParamsExtractOptionsLocale = "it"
	CrawlRootParamsExtractOptionsLocaleItCh      CrawlRootParamsExtractOptionsLocale = "it-CH"
	CrawlRootParamsExtractOptionsLocaleItIt      CrawlRootParamsExtractOptionsLocale = "it-IT"
	CrawlRootParamsExtractOptionsLocaleIuCa      CrawlRootParamsExtractOptionsLocale = "iu-CA"
	CrawlRootParamsExtractOptionsLocaleIwIl      CrawlRootParamsExtractOptionsLocale = "iw-IL"
	CrawlRootParamsExtractOptionsLocaleJa        CrawlRootParamsExtractOptionsLocale = "ja"
	CrawlRootParamsExtractOptionsLocaleJaJp      CrawlRootParamsExtractOptionsLocale = "ja-JP"
	CrawlRootParamsExtractOptionsLocaleJmc       CrawlRootParamsExtractOptionsLocale = "jmc"
	CrawlRootParamsExtractOptionsLocaleJmcTz     CrawlRootParamsExtractOptionsLocale = "jmc-TZ"
	CrawlRootParamsExtractOptionsLocaleKa        CrawlRootParamsExtractOptionsLocale = "ka"
	CrawlRootParamsExtractOptionsLocaleKaGe      CrawlRootParamsExtractOptionsLocale = "ka-GE"
	CrawlRootParamsExtractOptionsLocaleKab       CrawlRootParamsExtractOptionsLocale = "kab"
	CrawlRootParamsExtractOptionsLocaleKabDz     CrawlRootParamsExtractOptionsLocale = "kab-DZ"
	CrawlRootParamsExtractOptionsLocaleKam       CrawlRootParamsExtractOptionsLocale = "kam"
	CrawlRootParamsExtractOptionsLocaleKamKe     CrawlRootParamsExtractOptionsLocale = "kam-KE"
	CrawlRootParamsExtractOptionsLocaleKde       CrawlRootParamsExtractOptionsLocale = "kde"
	CrawlRootParamsExtractOptionsLocaleKdeTz     CrawlRootParamsExtractOptionsLocale = "kde-TZ"
	CrawlRootParamsExtractOptionsLocaleKea       CrawlRootParamsExtractOptionsLocale = "kea"
	CrawlRootParamsExtractOptionsLocaleKeaCv     CrawlRootParamsExtractOptionsLocale = "kea-CV"
	CrawlRootParamsExtractOptionsLocaleKhq       CrawlRootParamsExtractOptionsLocale = "khq"
	CrawlRootParamsExtractOptionsLocaleKhqMl     CrawlRootParamsExtractOptionsLocale = "khq-ML"
	CrawlRootParamsExtractOptionsLocaleKi        CrawlRootParamsExtractOptionsLocale = "ki"
	CrawlRootParamsExtractOptionsLocaleKiKe      CrawlRootParamsExtractOptionsLocale = "ki-KE"
	CrawlRootParamsExtractOptionsLocaleKk        CrawlRootParamsExtractOptionsLocale = "kk"
	CrawlRootParamsExtractOptionsLocaleKkCyrl    CrawlRootParamsExtractOptionsLocale = "kk-Cyrl"
	CrawlRootParamsExtractOptionsLocaleKkCyrlKz  CrawlRootParamsExtractOptionsLocale = "kk-Cyrl-KZ"
	CrawlRootParamsExtractOptionsLocaleKkKz      CrawlRootParamsExtractOptionsLocale = "kk-KZ"
	CrawlRootParamsExtractOptionsLocaleKl        CrawlRootParamsExtractOptionsLocale = "kl"
	CrawlRootParamsExtractOptionsLocaleKlGl      CrawlRootParamsExtractOptionsLocale = "kl-GL"
	CrawlRootParamsExtractOptionsLocaleKln       CrawlRootParamsExtractOptionsLocale = "kln"
	CrawlRootParamsExtractOptionsLocaleKlnKe     CrawlRootParamsExtractOptionsLocale = "kln-KE"
	CrawlRootParamsExtractOptionsLocaleKm        CrawlRootParamsExtractOptionsLocale = "km"
	CrawlRootParamsExtractOptionsLocaleKmKh      CrawlRootParamsExtractOptionsLocale = "km-KH"
	CrawlRootParamsExtractOptionsLocaleKn        CrawlRootParamsExtractOptionsLocale = "kn"
	CrawlRootParamsExtractOptionsLocaleKnIn      CrawlRootParamsExtractOptionsLocale = "kn-IN"
	CrawlRootParamsExtractOptionsLocaleKo        CrawlRootParamsExtractOptionsLocale = "ko"
	CrawlRootParamsExtractOptionsLocaleKoKr      CrawlRootParamsExtractOptionsLocale = "ko-KR"
	CrawlRootParamsExtractOptionsLocaleKok       CrawlRootParamsExtractOptionsLocale = "kok"
	CrawlRootParamsExtractOptionsLocaleKokIn     CrawlRootParamsExtractOptionsLocale = "kok-IN"
	CrawlRootParamsExtractOptionsLocaleKsIn      CrawlRootParamsExtractOptionsLocale = "ks-IN"
	CrawlRootParamsExtractOptionsLocaleKuTr      CrawlRootParamsExtractOptionsLocale = "ku-TR"
	CrawlRootParamsExtractOptionsLocaleKw        CrawlRootParamsExtractOptionsLocale = "kw"
	CrawlRootParamsExtractOptionsLocaleKwGB      CrawlRootParamsExtractOptionsLocale = "kw-GB"
	CrawlRootParamsExtractOptionsLocaleKyKg      CrawlRootParamsExtractOptionsLocale = "ky-KG"
	CrawlRootParamsExtractOptionsLocaleLag       CrawlRootParamsExtractOptionsLocale = "lag"
	CrawlRootParamsExtractOptionsLocaleLagTz     CrawlRootParamsExtractOptionsLocale = "lag-TZ"
	CrawlRootParamsExtractOptionsLocaleLbLu      CrawlRootParamsExtractOptionsLocale = "lb-LU"
	CrawlRootParamsExtractOptionsLocaleLg        CrawlRootParamsExtractOptionsLocale = "lg"
	CrawlRootParamsExtractOptionsLocaleLgUg      CrawlRootParamsExtractOptionsLocale = "lg-UG"
	CrawlRootParamsExtractOptionsLocaleLiBe      CrawlRootParamsExtractOptionsLocale = "li-BE"
	CrawlRootParamsExtractOptionsLocaleLiNl      CrawlRootParamsExtractOptionsLocale = "li-NL"
	CrawlRootParamsExtractOptionsLocaleLijIt     CrawlRootParamsExtractOptionsLocale = "lij-IT"
	CrawlRootParamsExtractOptionsLocaleLoLa      CrawlRootParamsExtractOptionsLocale = "lo-LA"
	CrawlRootParamsExtractOptionsLocaleLt        CrawlRootParamsExtractOptionsLocale = "lt"
	CrawlRootParamsExtractOptionsLocaleLtLt      CrawlRootParamsExtractOptionsLocale = "lt-LT"
	CrawlRootParamsExtractOptionsLocaleLuo       CrawlRootParamsExtractOptionsLocale = "luo"
	CrawlRootParamsExtractOptionsLocaleLuoKe     CrawlRootParamsExtractOptionsLocale = "luo-KE"
	CrawlRootParamsExtractOptionsLocaleLuy       CrawlRootParamsExtractOptionsLocale = "luy"
	CrawlRootParamsExtractOptionsLocaleLuyKe     CrawlRootParamsExtractOptionsLocale = "luy-KE"
	CrawlRootParamsExtractOptionsLocaleLv        CrawlRootParamsExtractOptionsLocale = "lv"
	CrawlRootParamsExtractOptionsLocaleLvLv      CrawlRootParamsExtractOptionsLocale = "lv-LV"
	CrawlRootParamsExtractOptionsLocaleMagIn     CrawlRootParamsExtractOptionsLocale = "mag-IN"
	CrawlRootParamsExtractOptionsLocaleMaiIn     CrawlRootParamsExtractOptionsLocale = "mai-IN"
	CrawlRootParamsExtractOptionsLocaleMas       CrawlRootParamsExtractOptionsLocale = "mas"
	CrawlRootParamsExtractOptionsLocaleMasKe     CrawlRootParamsExtractOptionsLocale = "mas-KE"
	CrawlRootParamsExtractOptionsLocaleMasTz     CrawlRootParamsExtractOptionsLocale = "mas-TZ"
	CrawlRootParamsExtractOptionsLocaleMer       CrawlRootParamsExtractOptionsLocale = "mer"
	CrawlRootParamsExtractOptionsLocaleMerKe     CrawlRootParamsExtractOptionsLocale = "mer-KE"
	CrawlRootParamsExtractOptionsLocaleMfe       CrawlRootParamsExtractOptionsLocale = "mfe"
	CrawlRootParamsExtractOptionsLocaleMfeMu     CrawlRootParamsExtractOptionsLocale = "mfe-MU"
	CrawlRootParamsExtractOptionsLocaleMg        CrawlRootParamsExtractOptionsLocale = "mg"
	CrawlRootParamsExtractOptionsLocaleMgMg      CrawlRootParamsExtractOptionsLocale = "mg-MG"
	CrawlRootParamsExtractOptionsLocaleMhrRu     CrawlRootParamsExtractOptionsLocale = "mhr-RU"
	CrawlRootParamsExtractOptionsLocaleMiNz      CrawlRootParamsExtractOptionsLocale = "mi-NZ"
	CrawlRootParamsExtractOptionsLocaleMk        CrawlRootParamsExtractOptionsLocale = "mk"
	CrawlRootParamsExtractOptionsLocaleMkMk      CrawlRootParamsExtractOptionsLocale = "mk-MK"
	CrawlRootParamsExtractOptionsLocaleMl        CrawlRootParamsExtractOptionsLocale = "ml"
	CrawlRootParamsExtractOptionsLocaleMlIn      CrawlRootParamsExtractOptionsLocale = "ml-IN"
	CrawlRootParamsExtractOptionsLocaleMnMn      CrawlRootParamsExtractOptionsLocale = "mn-MN"
	CrawlRootParamsExtractOptionsLocaleMr        CrawlRootParamsExtractOptionsLocale = "mr"
	CrawlRootParamsExtractOptionsLocaleMrIn      CrawlRootParamsExtractOptionsLocale = "mr-IN"
	CrawlRootParamsExtractOptionsLocaleMs        CrawlRootParamsExtractOptionsLocale = "ms"
	CrawlRootParamsExtractOptionsLocaleMsBn      CrawlRootParamsExtractOptionsLocale = "ms-BN"
	CrawlRootParamsExtractOptionsLocaleMsMy      CrawlRootParamsExtractOptionsLocale = "ms-MY"
	CrawlRootParamsExtractOptionsLocaleMt        CrawlRootParamsExtractOptionsLocale = "mt"
	CrawlRootParamsExtractOptionsLocaleMtMt      CrawlRootParamsExtractOptionsLocale = "mt-MT"
	CrawlRootParamsExtractOptionsLocaleMy        CrawlRootParamsExtractOptionsLocale = "my"
	CrawlRootParamsExtractOptionsLocaleMyMm      CrawlRootParamsExtractOptionsLocale = "my-MM"
	CrawlRootParamsExtractOptionsLocaleNanTw     CrawlRootParamsExtractOptionsLocale = "nan-TW"
	CrawlRootParamsExtractOptionsLocaleNaq       CrawlRootParamsExtractOptionsLocale = "naq"
	CrawlRootParamsExtractOptionsLocaleNaqNa     CrawlRootParamsExtractOptionsLocale = "naq-NA"
	CrawlRootParamsExtractOptionsLocaleNb        CrawlRootParamsExtractOptionsLocale = "nb"
	CrawlRootParamsExtractOptionsLocaleNbNo      CrawlRootParamsExtractOptionsLocale = "nb-NO"
	CrawlRootParamsExtractOptionsLocaleNd        CrawlRootParamsExtractOptionsLocale = "nd"
	CrawlRootParamsExtractOptionsLocaleNdZw      CrawlRootParamsExtractOptionsLocale = "nd-ZW"
	CrawlRootParamsExtractOptionsLocaleNdsDe     CrawlRootParamsExtractOptionsLocale = "nds-DE"
	CrawlRootParamsExtractOptionsLocaleNdsNl     CrawlRootParamsExtractOptionsLocale = "nds-NL"
	CrawlRootParamsExtractOptionsLocaleNe        CrawlRootParamsExtractOptionsLocale = "ne"
	CrawlRootParamsExtractOptionsLocaleNeIn      CrawlRootParamsExtractOptionsLocale = "ne-IN"
	CrawlRootParamsExtractOptionsLocaleNeNp      CrawlRootParamsExtractOptionsLocale = "ne-NP"
	CrawlRootParamsExtractOptionsLocaleNl        CrawlRootParamsExtractOptionsLocale = "nl"
	CrawlRootParamsExtractOptionsLocaleNlAw      CrawlRootParamsExtractOptionsLocale = "nl-AW"
	CrawlRootParamsExtractOptionsLocaleNlBe      CrawlRootParamsExtractOptionsLocale = "nl-BE"
	CrawlRootParamsExtractOptionsLocaleNlNl      CrawlRootParamsExtractOptionsLocale = "nl-NL"
	CrawlRootParamsExtractOptionsLocaleNn        CrawlRootParamsExtractOptionsLocale = "nn"
	CrawlRootParamsExtractOptionsLocaleNnNo      CrawlRootParamsExtractOptionsLocale = "nn-NO"
	CrawlRootParamsExtractOptionsLocaleNrZa      CrawlRootParamsExtractOptionsLocale = "nr-ZA"
	CrawlRootParamsExtractOptionsLocaleNsoZa     CrawlRootParamsExtractOptionsLocale = "nso-ZA"
	CrawlRootParamsExtractOptionsLocaleNyn       CrawlRootParamsExtractOptionsLocale = "nyn"
	CrawlRootParamsExtractOptionsLocaleNynUg     CrawlRootParamsExtractOptionsLocale = "nyn-UG"
	CrawlRootParamsExtractOptionsLocaleOcFr      CrawlRootParamsExtractOptionsLocale = "oc-FR"
	CrawlRootParamsExtractOptionsLocaleOm        CrawlRootParamsExtractOptionsLocale = "om"
	CrawlRootParamsExtractOptionsLocaleOmEt      CrawlRootParamsExtractOptionsLocale = "om-ET"
	CrawlRootParamsExtractOptionsLocaleOmKe      CrawlRootParamsExtractOptionsLocale = "om-KE"
	CrawlRootParamsExtractOptionsLocaleOr        CrawlRootParamsExtractOptionsLocale = "or"
	CrawlRootParamsExtractOptionsLocaleOrIn      CrawlRootParamsExtractOptionsLocale = "or-IN"
	CrawlRootParamsExtractOptionsLocaleOsRu      CrawlRootParamsExtractOptionsLocale = "os-RU"
	CrawlRootParamsExtractOptionsLocalePa        CrawlRootParamsExtractOptionsLocale = "pa"
	CrawlRootParamsExtractOptionsLocalePaArab    CrawlRootParamsExtractOptionsLocale = "pa-Arab"
	CrawlRootParamsExtractOptionsLocalePaArabPk  CrawlRootParamsExtractOptionsLocale = "pa-Arab-PK"
	CrawlRootParamsExtractOptionsLocalePaGuru    CrawlRootParamsExtractOptionsLocale = "pa-Guru"
	CrawlRootParamsExtractOptionsLocalePaGuruIn  CrawlRootParamsExtractOptionsLocale = "pa-Guru-IN"
	CrawlRootParamsExtractOptionsLocalePaIn      CrawlRootParamsExtractOptionsLocale = "pa-IN"
	CrawlRootParamsExtractOptionsLocalePaPk      CrawlRootParamsExtractOptionsLocale = "pa-PK"
	CrawlRootParamsExtractOptionsLocalePapAn     CrawlRootParamsExtractOptionsLocale = "pap-AN"
	CrawlRootParamsExtractOptionsLocalePl        CrawlRootParamsExtractOptionsLocale = "pl"
	CrawlRootParamsExtractOptionsLocalePlPl      CrawlRootParamsExtractOptionsLocale = "pl-PL"
	CrawlRootParamsExtractOptionsLocalePs        CrawlRootParamsExtractOptionsLocale = "ps"
	CrawlRootParamsExtractOptionsLocalePsAf      CrawlRootParamsExtractOptionsLocale = "ps-AF"
	CrawlRootParamsExtractOptionsLocalePt        CrawlRootParamsExtractOptionsLocale = "pt"
	CrawlRootParamsExtractOptionsLocalePtBr      CrawlRootParamsExtractOptionsLocale = "pt-BR"
	CrawlRootParamsExtractOptionsLocalePtGw      CrawlRootParamsExtractOptionsLocale = "pt-GW"
	CrawlRootParamsExtractOptionsLocalePtMz      CrawlRootParamsExtractOptionsLocale = "pt-MZ"
	CrawlRootParamsExtractOptionsLocalePtPt      CrawlRootParamsExtractOptionsLocale = "pt-PT"
	CrawlRootParamsExtractOptionsLocaleRm        CrawlRootParamsExtractOptionsLocale = "rm"
	CrawlRootParamsExtractOptionsLocaleRmCh      CrawlRootParamsExtractOptionsLocale = "rm-CH"
	CrawlRootParamsExtractOptionsLocaleRo        CrawlRootParamsExtractOptionsLocale = "ro"
	CrawlRootParamsExtractOptionsLocaleRoMd      CrawlRootParamsExtractOptionsLocale = "ro-MD"
	CrawlRootParamsExtractOptionsLocaleRoRo      CrawlRootParamsExtractOptionsLocale = "ro-RO"
	CrawlRootParamsExtractOptionsLocaleRof       CrawlRootParamsExtractOptionsLocale = "rof"
	CrawlRootParamsExtractOptionsLocaleRofTz     CrawlRootParamsExtractOptionsLocale = "rof-TZ"
	CrawlRootParamsExtractOptionsLocaleRu        CrawlRootParamsExtractOptionsLocale = "ru"
	CrawlRootParamsExtractOptionsLocaleRuMd      CrawlRootParamsExtractOptionsLocale = "ru-MD"
	CrawlRootParamsExtractOptionsLocaleRuRu      CrawlRootParamsExtractOptionsLocale = "ru-RU"
	CrawlRootParamsExtractOptionsLocaleRuUa      CrawlRootParamsExtractOptionsLocale = "ru-UA"
	CrawlRootParamsExtractOptionsLocaleRw        CrawlRootParamsExtractOptionsLocale = "rw"
	CrawlRootParamsExtractOptionsLocaleRwRw      CrawlRootParamsExtractOptionsLocale = "rw-RW"
	CrawlRootParamsExtractOptionsLocaleRwk       CrawlRootParamsExtractOptionsLocale = "rwk"
	CrawlRootParamsExtractOptionsLocaleRwkTz     CrawlRootParamsExtractOptionsLocale = "rwk-TZ"
	CrawlRootParamsExtractOptionsLocaleSaIn      CrawlRootParamsExtractOptionsLocale = "sa-IN"
	CrawlRootParamsExtractOptionsLocaleSaq       CrawlRootParamsExtractOptionsLocale = "saq"
	CrawlRootParamsExtractOptionsLocaleSaqKe     CrawlRootParamsExtractOptionsLocale = "saq-KE"
	CrawlRootParamsExtractOptionsLocaleScIt      CrawlRootParamsExtractOptionsLocale = "sc-IT"
	CrawlRootParamsExtractOptionsLocaleSdIn      CrawlRootParamsExtractOptionsLocale = "sd-IN"
	CrawlRootParamsExtractOptionsLocaleSeNo      CrawlRootParamsExtractOptionsLocale = "se-NO"
	CrawlRootParamsExtractOptionsLocaleSeh       CrawlRootParamsExtractOptionsLocale = "seh"
	CrawlRootParamsExtractOptionsLocaleSehMz     CrawlRootParamsExtractOptionsLocale = "seh-MZ"
	CrawlRootParamsExtractOptionsLocaleSes       CrawlRootParamsExtractOptionsLocale = "ses"
	CrawlRootParamsExtractOptionsLocaleSesMl     CrawlRootParamsExtractOptionsLocale = "ses-ML"
	CrawlRootParamsExtractOptionsLocaleSg        CrawlRootParamsExtractOptionsLocale = "sg"
	CrawlRootParamsExtractOptionsLocaleSgCf      CrawlRootParamsExtractOptionsLocale = "sg-CF"
	CrawlRootParamsExtractOptionsLocaleShi       CrawlRootParamsExtractOptionsLocale = "shi"
	CrawlRootParamsExtractOptionsLocaleShiLatn   CrawlRootParamsExtractOptionsLocale = "shi-Latn"
	CrawlRootParamsExtractOptionsLocaleShiLatnMa CrawlRootParamsExtractOptionsLocale = "shi-Latn-MA"
	CrawlRootParamsExtractOptionsLocaleShiTfng   CrawlRootParamsExtractOptionsLocale = "shi-Tfng"
	CrawlRootParamsExtractOptionsLocaleShiTfngMa CrawlRootParamsExtractOptionsLocale = "shi-Tfng-MA"
	CrawlRootParamsExtractOptionsLocaleShsCa     CrawlRootParamsExtractOptionsLocale = "shs-CA"
	CrawlRootParamsExtractOptionsLocaleSi        CrawlRootParamsExtractOptionsLocale = "si"
	CrawlRootParamsExtractOptionsLocaleSiLk      CrawlRootParamsExtractOptionsLocale = "si-LK"
	CrawlRootParamsExtractOptionsLocaleSidEt     CrawlRootParamsExtractOptionsLocale = "sid-ET"
	CrawlRootParamsExtractOptionsLocaleSk        CrawlRootParamsExtractOptionsLocale = "sk"
	CrawlRootParamsExtractOptionsLocaleSkSk      CrawlRootParamsExtractOptionsLocale = "sk-SK"
	CrawlRootParamsExtractOptionsLocaleSl        CrawlRootParamsExtractOptionsLocale = "sl"
	CrawlRootParamsExtractOptionsLocaleSlSi      CrawlRootParamsExtractOptionsLocale = "sl-SI"
	CrawlRootParamsExtractOptionsLocaleSn        CrawlRootParamsExtractOptionsLocale = "sn"
	CrawlRootParamsExtractOptionsLocaleSnZw      CrawlRootParamsExtractOptionsLocale = "sn-ZW"
	CrawlRootParamsExtractOptionsLocaleSo        CrawlRootParamsExtractOptionsLocale = "so"
	CrawlRootParamsExtractOptionsLocaleSoDj      CrawlRootParamsExtractOptionsLocale = "so-DJ"
	CrawlRootParamsExtractOptionsLocaleSoEt      CrawlRootParamsExtractOptionsLocale = "so-ET"
	CrawlRootParamsExtractOptionsLocaleSoKe      CrawlRootParamsExtractOptionsLocale = "so-KE"
	CrawlRootParamsExtractOptionsLocaleSoSo      CrawlRootParamsExtractOptionsLocale = "so-SO"
	CrawlRootParamsExtractOptionsLocaleSq        CrawlRootParamsExtractOptionsLocale = "sq"
	CrawlRootParamsExtractOptionsLocaleSqAl      CrawlRootParamsExtractOptionsLocale = "sq-AL"
	CrawlRootParamsExtractOptionsLocaleSqMk      CrawlRootParamsExtractOptionsLocale = "sq-MK"
	CrawlRootParamsExtractOptionsLocaleSr        CrawlRootParamsExtractOptionsLocale = "sr"
	CrawlRootParamsExtractOptionsLocaleSrCyrl    CrawlRootParamsExtractOptionsLocale = "sr-Cyrl"
	CrawlRootParamsExtractOptionsLocaleSrCyrlBa  CrawlRootParamsExtractOptionsLocale = "sr-Cyrl-BA"
	CrawlRootParamsExtractOptionsLocaleSrCyrlMe  CrawlRootParamsExtractOptionsLocale = "sr-Cyrl-ME"
	CrawlRootParamsExtractOptionsLocaleSrCyrlRs  CrawlRootParamsExtractOptionsLocale = "sr-Cyrl-RS"
	CrawlRootParamsExtractOptionsLocaleSrLatn    CrawlRootParamsExtractOptionsLocale = "sr-Latn"
	CrawlRootParamsExtractOptionsLocaleSrLatnBa  CrawlRootParamsExtractOptionsLocale = "sr-Latn-BA"
	CrawlRootParamsExtractOptionsLocaleSrLatnMe  CrawlRootParamsExtractOptionsLocale = "sr-Latn-ME"
	CrawlRootParamsExtractOptionsLocaleSrLatnRs  CrawlRootParamsExtractOptionsLocale = "sr-Latn-RS"
	CrawlRootParamsExtractOptionsLocaleSrMe      CrawlRootParamsExtractOptionsLocale = "sr-ME"
	CrawlRootParamsExtractOptionsLocaleSrRs      CrawlRootParamsExtractOptionsLocale = "sr-RS"
	CrawlRootParamsExtractOptionsLocaleSSZa      CrawlRootParamsExtractOptionsLocale = "ss-ZA"
	CrawlRootParamsExtractOptionsLocaleStZa      CrawlRootParamsExtractOptionsLocale = "st-ZA"
	CrawlRootParamsExtractOptionsLocaleSv        CrawlRootParamsExtractOptionsLocale = "sv"
	CrawlRootParamsExtractOptionsLocaleSvFi      CrawlRootParamsExtractOptionsLocale = "sv-FI"
	CrawlRootParamsExtractOptionsLocaleSvSe      CrawlRootParamsExtractOptionsLocale = "sv-SE"
	CrawlRootParamsExtractOptionsLocaleSw        CrawlRootParamsExtractOptionsLocale = "sw"
	CrawlRootParamsExtractOptionsLocaleSwKe      CrawlRootParamsExtractOptionsLocale = "sw-KE"
	CrawlRootParamsExtractOptionsLocaleSwTz      CrawlRootParamsExtractOptionsLocale = "sw-TZ"
	CrawlRootParamsExtractOptionsLocaleTa        CrawlRootParamsExtractOptionsLocale = "ta"
	CrawlRootParamsExtractOptionsLocaleTaIn      CrawlRootParamsExtractOptionsLocale = "ta-IN"
	CrawlRootParamsExtractOptionsLocaleTaLk      CrawlRootParamsExtractOptionsLocale = "ta-LK"
	CrawlRootParamsExtractOptionsLocaleTe        CrawlRootParamsExtractOptionsLocale = "te"
	CrawlRootParamsExtractOptionsLocaleTeIn      CrawlRootParamsExtractOptionsLocale = "te-IN"
	CrawlRootParamsExtractOptionsLocaleTeo       CrawlRootParamsExtractOptionsLocale = "teo"
	CrawlRootParamsExtractOptionsLocaleTeoKe     CrawlRootParamsExtractOptionsLocale = "teo-KE"
	CrawlRootParamsExtractOptionsLocaleTeoUg     CrawlRootParamsExtractOptionsLocale = "teo-UG"
	CrawlRootParamsExtractOptionsLocaleTgTj      CrawlRootParamsExtractOptionsLocale = "tg-TJ"
	CrawlRootParamsExtractOptionsLocaleTh        CrawlRootParamsExtractOptionsLocale = "th"
	CrawlRootParamsExtractOptionsLocaleThTh      CrawlRootParamsExtractOptionsLocale = "th-TH"
	CrawlRootParamsExtractOptionsLocaleTi        CrawlRootParamsExtractOptionsLocale = "ti"
	CrawlRootParamsExtractOptionsLocaleTiEr      CrawlRootParamsExtractOptionsLocale = "ti-ER"
	CrawlRootParamsExtractOptionsLocaleTiEt      CrawlRootParamsExtractOptionsLocale = "ti-ET"
	CrawlRootParamsExtractOptionsLocaleTigEr     CrawlRootParamsExtractOptionsLocale = "tig-ER"
	CrawlRootParamsExtractOptionsLocaleTkTm      CrawlRootParamsExtractOptionsLocale = "tk-TM"
	CrawlRootParamsExtractOptionsLocaleTlPh      CrawlRootParamsExtractOptionsLocale = "tl-PH"
	CrawlRootParamsExtractOptionsLocaleTnZa      CrawlRootParamsExtractOptionsLocale = "tn-ZA"
	CrawlRootParamsExtractOptionsLocaleTo        CrawlRootParamsExtractOptionsLocale = "to"
	CrawlRootParamsExtractOptionsLocaleToTo      CrawlRootParamsExtractOptionsLocale = "to-TO"
	CrawlRootParamsExtractOptionsLocaleTr        CrawlRootParamsExtractOptionsLocale = "tr"
	CrawlRootParamsExtractOptionsLocaleTrCy      CrawlRootParamsExtractOptionsLocale = "tr-CY"
	CrawlRootParamsExtractOptionsLocaleTrTr      CrawlRootParamsExtractOptionsLocale = "tr-TR"
	CrawlRootParamsExtractOptionsLocaleTsZa      CrawlRootParamsExtractOptionsLocale = "ts-ZA"
	CrawlRootParamsExtractOptionsLocaleTtRu      CrawlRootParamsExtractOptionsLocale = "tt-RU"
	CrawlRootParamsExtractOptionsLocaleTzm       CrawlRootParamsExtractOptionsLocale = "tzm"
	CrawlRootParamsExtractOptionsLocaleTzmLatn   CrawlRootParamsExtractOptionsLocale = "tzm-Latn"
	CrawlRootParamsExtractOptionsLocaleTzmLatnMa CrawlRootParamsExtractOptionsLocale = "tzm-Latn-MA"
	CrawlRootParamsExtractOptionsLocaleUgCn      CrawlRootParamsExtractOptionsLocale = "ug-CN"
	CrawlRootParamsExtractOptionsLocaleUk        CrawlRootParamsExtractOptionsLocale = "uk"
	CrawlRootParamsExtractOptionsLocaleUkUa      CrawlRootParamsExtractOptionsLocale = "uk-UA"
	CrawlRootParamsExtractOptionsLocaleUnmUs     CrawlRootParamsExtractOptionsLocale = "unm-US"
	CrawlRootParamsExtractOptionsLocaleUr        CrawlRootParamsExtractOptionsLocale = "ur"
	CrawlRootParamsExtractOptionsLocaleUrIn      CrawlRootParamsExtractOptionsLocale = "ur-IN"
	CrawlRootParamsExtractOptionsLocaleUrPk      CrawlRootParamsExtractOptionsLocale = "ur-PK"
	CrawlRootParamsExtractOptionsLocaleUz        CrawlRootParamsExtractOptionsLocale = "uz"
	CrawlRootParamsExtractOptionsLocaleUzArab    CrawlRootParamsExtractOptionsLocale = "uz-Arab"
	CrawlRootParamsExtractOptionsLocaleUzArabAf  CrawlRootParamsExtractOptionsLocale = "uz-Arab-AF"
	CrawlRootParamsExtractOptionsLocaleUzCyrl    CrawlRootParamsExtractOptionsLocale = "uz-Cyrl"
	CrawlRootParamsExtractOptionsLocaleUzCyrlUz  CrawlRootParamsExtractOptionsLocale = "uz-Cyrl-UZ"
	CrawlRootParamsExtractOptionsLocaleUzLatn    CrawlRootParamsExtractOptionsLocale = "uz-Latn"
	CrawlRootParamsExtractOptionsLocaleUzLatnUz  CrawlRootParamsExtractOptionsLocale = "uz-Latn-UZ"
	CrawlRootParamsExtractOptionsLocaleUzUz      CrawlRootParamsExtractOptionsLocale = "uz-UZ"
	CrawlRootParamsExtractOptionsLocaleVeZa      CrawlRootParamsExtractOptionsLocale = "ve-ZA"
	CrawlRootParamsExtractOptionsLocaleVi        CrawlRootParamsExtractOptionsLocale = "vi"
	CrawlRootParamsExtractOptionsLocaleViVn      CrawlRootParamsExtractOptionsLocale = "vi-VN"
	CrawlRootParamsExtractOptionsLocaleVun       CrawlRootParamsExtractOptionsLocale = "vun"
	CrawlRootParamsExtractOptionsLocaleVunTz     CrawlRootParamsExtractOptionsLocale = "vun-TZ"
	CrawlRootParamsExtractOptionsLocaleWaBe      CrawlRootParamsExtractOptionsLocale = "wa-BE"
	CrawlRootParamsExtractOptionsLocaleWaeCh     CrawlRootParamsExtractOptionsLocale = "wae-CH"
	CrawlRootParamsExtractOptionsLocaleWalEt     CrawlRootParamsExtractOptionsLocale = "wal-ET"
	CrawlRootParamsExtractOptionsLocaleWoSn      CrawlRootParamsExtractOptionsLocale = "wo-SN"
	CrawlRootParamsExtractOptionsLocaleXhZa      CrawlRootParamsExtractOptionsLocale = "xh-ZA"
	CrawlRootParamsExtractOptionsLocaleXog       CrawlRootParamsExtractOptionsLocale = "xog"
	CrawlRootParamsExtractOptionsLocaleXogUg     CrawlRootParamsExtractOptionsLocale = "xog-UG"
	CrawlRootParamsExtractOptionsLocaleYiUs      CrawlRootParamsExtractOptionsLocale = "yi-US"
	CrawlRootParamsExtractOptionsLocaleYo        CrawlRootParamsExtractOptionsLocale = "yo"
	CrawlRootParamsExtractOptionsLocaleYoNg      CrawlRootParamsExtractOptionsLocale = "yo-NG"
	CrawlRootParamsExtractOptionsLocaleYueHk     CrawlRootParamsExtractOptionsLocale = "yue-HK"
	CrawlRootParamsExtractOptionsLocaleZh        CrawlRootParamsExtractOptionsLocale = "zh"
	CrawlRootParamsExtractOptionsLocaleZhCn      CrawlRootParamsExtractOptionsLocale = "zh-CN"
	CrawlRootParamsExtractOptionsLocaleZhHk      CrawlRootParamsExtractOptionsLocale = "zh-HK"
	CrawlRootParamsExtractOptionsLocaleZhHans    CrawlRootParamsExtractOptionsLocale = "zh-Hans"
	CrawlRootParamsExtractOptionsLocaleZhHansCn  CrawlRootParamsExtractOptionsLocale = "zh-Hans-CN"
	CrawlRootParamsExtractOptionsLocaleZhHansHk  CrawlRootParamsExtractOptionsLocale = "zh-Hans-HK"
	CrawlRootParamsExtractOptionsLocaleZhHansMo  CrawlRootParamsExtractOptionsLocale = "zh-Hans-MO"
	CrawlRootParamsExtractOptionsLocaleZhHansSg  CrawlRootParamsExtractOptionsLocale = "zh-Hans-SG"
	CrawlRootParamsExtractOptionsLocaleZhHant    CrawlRootParamsExtractOptionsLocale = "zh-Hant"
	CrawlRootParamsExtractOptionsLocaleZhHantHk  CrawlRootParamsExtractOptionsLocale = "zh-Hant-HK"
	CrawlRootParamsExtractOptionsLocaleZhHantMo  CrawlRootParamsExtractOptionsLocale = "zh-Hant-MO"
	CrawlRootParamsExtractOptionsLocaleZhHantTw  CrawlRootParamsExtractOptionsLocale = "zh-Hant-TW"
	CrawlRootParamsExtractOptionsLocaleZhSg      CrawlRootParamsExtractOptionsLocale = "zh-SG"
	CrawlRootParamsExtractOptionsLocaleZhTw      CrawlRootParamsExtractOptionsLocale = "zh-TW"
	CrawlRootParamsExtractOptionsLocaleZu        CrawlRootParamsExtractOptionsLocale = "zu"
	CrawlRootParamsExtractOptionsLocaleZuZa      CrawlRootParamsExtractOptionsLocale = "zu-ZA"
	CrawlRootParamsExtractOptionsLocaleAuto      CrawlRootParamsExtractOptionsLocale = "auto"
)

// Structured metadata about the request execution context
type CrawlRootParamsExtractOptionsMetadata struct {
	// Account name associated with the request
	AccountName param.Opt[string] `json:"account_name,omitzero"`
	// Definition identifier
	DefinitionID param.Opt[int64] `json:"definition_id,omitzero"`
	// Name of the definition
	DefinitionName param.Opt[string] `json:"definition_name,omitzero"`
	// API endpoint being called
	Endpoint param.Opt[string] `json:"endpoint,omitzero"`
	// Unique identifier for this execution
	ExecutionID param.Opt[string] `json:"execution_id,omitzero"`
	// FlowIt task identifier
	FlowitTaskID param.Opt[string] `json:"flowit_task_id,omitzero"`
	// Input data identifier
	InputID param.Opt[string] `json:"input_id,omitzero"`
	// Identifier for the pipeline execution
	PipelineExecutionID param.Opt[int64] `json:"pipeline_execution_id,omitzero"`
	// Query template identifier
	QueryTemplateID param.Opt[string] `json:"query_template_id,omitzero"`
	// Source system or application making the request
	Source param.Opt[string] `json:"source,omitzero"`
	// Template identifier
	TemplateID param.Opt[int64] `json:"template_id,omitzero"`
	// Name of the template
	TemplateName param.Opt[string] `json:"template_name,omitzero"`
	paramObj
}

func (r CrawlRootParamsExtractOptionsMetadata) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsExtractOptionsMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRootParamsExtractOptionsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlRootParamsExtractOptionsNetworkCapture struct {
	Validation                  param.Opt[bool]    `json:"validation,omitzero"`
	WaitForRequestsCount        param.Opt[float64] `json:"wait_for_requests_count,omitzero"`
	WaitForRequestsCountTimeout param.Opt[float64] `json:"wait_for_requests_count_timeout,omitzero"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method,omitzero"`
	// Resource type for network capture filtering
	ResourceType CrawlRootParamsExtractOptionsNetworkCaptureResourceTypeUnion `json:"resource_type,omitzero"`
	StatusCode   CrawlRootParamsExtractOptionsNetworkCaptureStatusCodeUnion   `json:"status_code,omitzero"`
	URL          CrawlRootParamsExtractOptionsNetworkCaptureURL               `json:"url,omitzero"`
	paramObj
}

func (r CrawlRootParamsExtractOptionsNetworkCapture) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsExtractOptionsNetworkCapture
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRootParamsExtractOptionsNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptionsNetworkCapture](
		"method", "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsNetworkCaptureResourceTypeUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsNetworkCaptureResourceTypeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlRootParamsExtractOptionsNetworkCaptureResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsNetworkCaptureResourceTypeUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsNetworkCaptureStatusCodeUnion struct {
	OfFloat      param.Opt[float64] `json:",omitzero,inline"`
	OfFloatArray []float64          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsNetworkCaptureStatusCodeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfFloatArray)
}
func (u *CrawlRootParamsExtractOptionsNetworkCaptureStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsNetworkCaptureStatusCodeUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfFloatArray) {
		return &u.OfFloatArray
	}
	return nil
}

// The property Value is required.
type CrawlRootParamsExtractOptionsNetworkCaptureURL struct {
	Value string `json:"value,required"`
	// Any of "exact", "contains".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r CrawlRootParamsExtractOptionsNetworkCaptureURL) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsExtractOptionsNetworkCaptureURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRootParamsExtractOptionsNetworkCaptureURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptionsNetworkCaptureURL](
		"type", "exact", "contains",
	)
}

// Configuration options for parsing behavior
type CrawlRootParamsExtractOptionsParseOptions struct {
	// Whether to merge dynamic parsing results with static results
	MergeDynamic param.Opt[bool] `json:"merge_dynamic,omitzero"`
	ExtraFields  map[string]any  `json:"-"`
	paramObj
}

func (r CrawlRootParamsExtractOptionsParseOptions) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsExtractOptionsParseOptions
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *CrawlRootParamsExtractOptionsParseOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsParserUnion struct {
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsParserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap, u.OfString)
}
func (u *CrawlRootParamsExtractOptionsParserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsParserUnion) asAny() any {
	if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Proxy provider to use for the request
type CrawlRootParamsExtractOptionsProxyProvider string

const (
	CrawlRootParamsExtractOptionsProxyProviderBrightdata      CrawlRootParamsExtractOptionsProxyProvider = "brightdata"
	CrawlRootParamsExtractOptionsProxyProviderOxylabs         CrawlRootParamsExtractOptionsProxyProvider = "oxylabs"
	CrawlRootParamsExtractOptionsProxyProviderSmartproxy      CrawlRootParamsExtractOptionsProxyProvider = "smartproxy"
	CrawlRootParamsExtractOptionsProxyProviderProxit          CrawlRootParamsExtractOptionsProxyProvider = "proxit"
	CrawlRootParamsExtractOptionsProxyProviderProxitPreprod   CrawlRootParamsExtractOptionsProxyProvider = "proxit_preprod"
	CrawlRootParamsExtractOptionsProxyProviderLocal           CrawlRootParamsExtractOptionsProxyProvider = "local"
	CrawlRootParamsExtractOptionsProxyProviderRayobyte        CrawlRootParamsExtractOptionsProxyProvider = "rayobyte"
	CrawlRootParamsExtractOptionsProxyProviderAlways          CrawlRootParamsExtractOptionsProxyProvider = "always"
	CrawlRootParamsExtractOptionsProxyProviderOculusproxies   CrawlRootParamsExtractOptionsProxyProvider = "oculusproxies"
	CrawlRootParamsExtractOptionsProxyProviderFroxy           CrawlRootParamsExtractOptionsProxyProvider = "froxy"
	CrawlRootParamsExtractOptionsProxyProviderPacketstream    CrawlRootParamsExtractOptionsProxyProvider = "packetstream"
	CrawlRootParamsExtractOptionsProxyProvider911proxy        CrawlRootParamsExtractOptionsProxyProvider = "911proxy"
	CrawlRootParamsExtractOptionsProxyProviderDirect911proxy  CrawlRootParamsExtractOptionsProxyProvider = "direct911proxy"
	CrawlRootParamsExtractOptionsProxyProviderThesocialproxy  CrawlRootParamsExtractOptionsProxyProvider = "thesocialproxy"
	CrawlRootParamsExtractOptionsProxyProviderThesocialproxy2 CrawlRootParamsExtractOptionsProxyProvider = "thesocialproxy2"
	CrawlRootParamsExtractOptionsProxyProviderNimbleIsp       CrawlRootParamsExtractOptionsProxyProvider = "nimble-isp"
	CrawlRootParamsExtractOptionsProxyProviderNimbleIspMobile CrawlRootParamsExtractOptionsProxyProvider = "nimble-isp-mobile"
	CrawlRootParamsExtractOptionsProxyProviderProxitLinux     CrawlRootParamsExtractOptionsProxyProvider = "proxit-linux"
	CrawlRootParamsExtractOptionsProxyProviderProxitMacos     CrawlRootParamsExtractOptionsProxyProvider = "proxit-macos"
	CrawlRootParamsExtractOptionsProxyProviderProxitWindows   CrawlRootParamsExtractOptionsProxyProvider = "proxit-windows"
	CrawlRootParamsExtractOptionsProxyProviderProxitRental    CrawlRootParamsExtractOptionsProxyProvider = "proxit-rental"
	CrawlRootParamsExtractOptionsProxyProviderIpfoxy          CrawlRootParamsExtractOptionsProxyProvider = "ipfoxy"
	CrawlRootParamsExtractOptionsProxyProviderBrightup        CrawlRootParamsExtractOptionsProxyProvider = "brightup"
	CrawlRootParamsExtractOptionsProxyProviderResearch        CrawlRootParamsExtractOptionsProxyProvider = "research"
)

// Query template configuration for structured data extraction
//
// The property ID is required.
type CrawlRootParamsExtractOptionsQueryTemplate struct {
	ID string `json:"id,required" format:"uuid"`
	// Any of "WEB", "SERP", "SOCIAL".
	APIType     string                                                    `json:"api_type,omitzero"`
	Pagination  CrawlRootParamsExtractOptionsQueryTemplatePaginationUnion `json:"pagination,omitzero"`
	Params      map[string]any                                            `json:"params,omitzero"`
	ExtraFields map[string]any                                            `json:"-"`
	paramObj
}

func (r CrawlRootParamsExtractOptionsQueryTemplate) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsExtractOptionsQueryTemplate
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *CrawlRootParamsExtractOptionsQueryTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptionsQueryTemplate](
		"api_type", "WEB", "SERP", "SOCIAL",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsQueryTemplatePaginationUnion struct {
	OfCrawlRootsExtractOptionsQueryTemplatePaginationNextPageParams *CrawlRootParamsExtractOptionsQueryTemplatePaginationNextPageParams `json:",omitzero,inline"`
	OfCrawlRootsExtractOptionsQueryTemplatePaginationArray          []CrawlRootParamsExtractOptionsQueryTemplatePaginationArrayItem     `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsQueryTemplatePaginationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlRootsExtractOptionsQueryTemplatePaginationNextPageParams, u.OfCrawlRootsExtractOptionsQueryTemplatePaginationArray)
}
func (u *CrawlRootParamsExtractOptionsQueryTemplatePaginationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsQueryTemplatePaginationUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlRootsExtractOptionsQueryTemplatePaginationNextPageParams) {
		return u.OfCrawlRootsExtractOptionsQueryTemplatePaginationNextPageParams
	} else if !param.IsOmitted(u.OfCrawlRootsExtractOptionsQueryTemplatePaginationArray) {
		return &u.OfCrawlRootsExtractOptionsQueryTemplatePaginationArray
	}
	return nil
}

// The property NextPageParams is required.
type CrawlRootParamsExtractOptionsQueryTemplatePaginationNextPageParams struct {
	NextPageParams map[string]any `json:"next_page_params,omitzero,required"`
	paramObj
}

func (r CrawlRootParamsExtractOptionsQueryTemplatePaginationNextPageParams) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsExtractOptionsQueryTemplatePaginationNextPageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRootParamsExtractOptionsQueryTemplatePaginationNextPageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property NextPageParams is required.
type CrawlRootParamsExtractOptionsQueryTemplatePaginationArrayItem struct {
	NextPageParams map[string]any `json:"next_page_params,omitzero,required"`
	paramObj
}

func (r CrawlRootParamsExtractOptionsQueryTemplatePaginationArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsExtractOptionsQueryTemplatePaginationArrayItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRootParamsExtractOptionsQueryTemplatePaginationArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Referrer policy for the request
type CrawlRootParamsExtractOptionsReferrerType string

const (
	CrawlRootParamsExtractOptionsReferrerTypeRandom     CrawlRootParamsExtractOptionsReferrerType = "random"
	CrawlRootParamsExtractOptionsReferrerTypeNoReferer  CrawlRootParamsExtractOptionsReferrerType = "no-referer"
	CrawlRootParamsExtractOptionsReferrerTypeSameOrigin CrawlRootParamsExtractOptionsReferrerType = "same-origin"
	CrawlRootParamsExtractOptionsReferrerTypeGoogle     CrawlRootParamsExtractOptionsReferrerType = "google"
	CrawlRootParamsExtractOptionsReferrerTypeBing       CrawlRootParamsExtractOptionsReferrerType = "bing"
	CrawlRootParamsExtractOptionsReferrerTypeFacebook   CrawlRootParamsExtractOptionsReferrerType = "facebook"
	CrawlRootParamsExtractOptionsReferrerTypeTwitter    CrawlRootParamsExtractOptionsReferrerType = "twitter"
	CrawlRootParamsExtractOptionsReferrerTypeInstagram  CrawlRootParamsExtractOptionsReferrerType = "instagram"
)

type CrawlRootParamsExtractOptionsRenderOptions struct {
	// Whether to enable ad blocking
	Adblock param.Opt[bool] `json:"adblock,omitzero"`
	// Whether to enable browser caching
	Cache param.Opt[bool] `json:"cache,omitzero"`
	// Whether to use 2Captcha service for solving captchas
	Enable2captcha param.Opt[bool] `json:"enable_2captcha,omitzero"`
	// Fingerprint identifier for browser customization
	FingerprintID param.Opt[string] `json:"fingerprint_id,omitzero"`
	// Whether to run browser in headless mode
	Headless param.Opt[bool] `json:"headless,omitzero"`
	// Whether to include iframe content in the result
	IncludeIframes param.Opt[bool] `json:"include_iframes,omitzero"`
	// Whether to load previously stored localStorage data
	LoadLocalStorage param.Opt[bool] `json:"load_local_storage,omitzero"`
	// Disable content encoding to avoid cached responses
	NoAcceptEncoding param.Opt[bool] `json:"no_accept_encoding,omitzero"`
	// Whether to override default browser permissions
	OverridePermissions param.Opt[bool] `json:"override_permissions,omitzero"`
	// Whether to randomize HTTP header order
	RandomHeaderOrder param.Opt[bool] `json:"random_header_order,omitzero"`
	// Whether to store localStorage data for future sessions
	StoreLocalStorage param.Opt[bool] `json:"store_local_storage,omitzero"`
	// Maximum time in milliseconds to wait for page render
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Interval in milliseconds between key presses
	TypingInterval param.Opt[float64] `json:"typing_interval,omitzero"`
	// Whether to use a persistent browser session
	Userbrowser param.Opt[bool] `json:"userbrowser,omitzero"`
	// Whether to collect performance metrics during rendering
	WithPerformanceMetrics param.Opt[bool] `json:"with_performance_metrics,omitzero"`
	// Domains to block from loading
	BlockedDomains []string `json:"blocked_domains,omitzero"`
	// Browser engine to use, or weighted distribution of engines
	BrowserEngine CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineUnion `json:"browser_engine,omitzero"`
	// Type of browser connector to use
	//
	// Any of "puppeteer", "puppeteer-cdp", "puppeteer-bidi", "webit-cdp",
	// "playwright".
	ConnectorType string `json:"connector_type,omitzero"`
	// Types of resources to block from loading
	//
	// Any of "other", "document", "stylesheet", "image", "media", "font", "script",
	// "texttrack", "xhr", "fetch", "eventsource", "websocket", "manifest",
	// "signedexchange", "ping", "cspviolationreport", "prefetch", "preflight",
	// "fedcm".
	DisabledResources []string `json:"disabled_resources,omitzero"`
	// Browser extensions to load
	Extensions []string `json:"extensions,omitzero"`
	// Configuration for Hackium browser modifications
	HackiumConfiguration CrawlRootParamsExtractOptionsRenderOptionsHackiumConfiguration `json:"hackium_configuration,omitzero"`
	// Specific localStorage keys to load
	LocalStorageKeysToLoad []string `json:"local_storage_keys_to_load,omitzero"`
	// Strategy for simulating mouse movements
	//
	// Any of "linear", "ghost-cursor", "windmouse".
	MouseStrategy string `json:"mouse_strategy,omitzero"`
	// Type of render completion to wait for
	//
	// Any of "domcontentloaded", "load", "idle0", "networkidle0", "idle2",
	// "networkidle2", "navigate".
	RenderType string `json:"render_type,omitzero"`
	// Strategy for simulating keyboard typing
	//
	// Any of "simple", "distribution".
	TypingStrategy string `json:"typing_strategy,omitzero"`
	// Browser event to wait for before considering page loaded
	//
	// Any of "load", "domcontentloaded", "idle0", "idle2", "networkidle0",
	// "networkidle2", "navigate".
	WaitUntil string `json:"wait_until,omitzero"`
	paramObj
}

func (r CrawlRootParamsExtractOptionsRenderOptions) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsExtractOptionsRenderOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRootParamsExtractOptionsRenderOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptionsRenderOptions](
		"connector_type", "puppeteer", "puppeteer-cdp", "puppeteer-bidi", "webit-cdp", "playwright",
	)
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptionsRenderOptions](
		"mouse_strategy", "linear", "ghost-cursor", "windmouse",
	)
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptionsRenderOptions](
		"render_type", "domcontentloaded", "load", "idle0", "networkidle0", "idle2", "networkidle2", "navigate",
	)
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptionsRenderOptions](
		"typing_strategy", "simple", "distribution",
	)
	apijson.RegisterFieldValidator[CrawlRootParamsExtractOptionsRenderOptions](
		"wait_until", "load", "domcontentloaded", "idle0", "idle2", "networkidle0", "networkidle2", "navigate",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlRootsExtractOptionsRenderOptionsBrowserEngineString)
	OfCrawlRootsExtractOptionsRenderOptionsBrowserEngineString param.Opt[string]  `json:",omitzero,inline"`
	OfFloatMap                                                 map[string]float64 `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlRootsExtractOptionsRenderOptionsBrowserEngineString, u.OfFloatMap)
}
func (u *CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlRootsExtractOptionsRenderOptionsBrowserEngineString) {
		return &u.OfCrawlRootsExtractOptionsRenderOptionsBrowserEngineString
	} else if !param.IsOmitted(u.OfFloatMap) {
		return &u.OfFloatMap
	}
	return nil
}

type CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineString string

const (
	CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineStringChrome  CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineString = "chrome"
	CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineStringHackium CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineString = "hackium"
	CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineStringFirefox CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineString = "firefox"
	CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineStringHackfox CrawlRootParamsExtractOptionsRenderOptionsBrowserEngineString = "hackfox"
)

// Configuration for Hackium browser modifications
type CrawlRootParamsExtractOptionsRenderOptionsHackiumConfiguration struct {
	CollectLogs                 param.Opt[bool] `json:"collect_logs,omitzero"`
	DoNotFixMathSalt            param.Opt[bool] `json:"do_not_fix_math_salt,omitzero"`
	EnableDocumentElementSpoof  param.Opt[bool] `json:"enable_document_element_spoof,omitzero"`
	EnableDocumentHasFocus      param.Opt[bool] `json:"enable_document_has_focus,omitzero"`
	EnableFakeNavigationHistory param.Opt[bool] `json:"enable_fake_navigation_history,omitzero"`
	EnableKeyOrdering           param.Opt[bool] `json:"enable_key_ordering,omitzero"`
	EnableSniffer               param.Opt[bool] `json:"enable_sniffer,omitzero"`
	EnableVerboseLogs           param.Opt[bool] `json:"enable_verbose_logs,omitzero"`
	paramObj
}

func (r CrawlRootParamsExtractOptionsRenderOptionsHackiumConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsExtractOptionsRenderOptionsHackiumConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRootParamsExtractOptionsRenderOptionsHackiumConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlRootParamsExtractOptionsSession struct {
	ID                  param.Opt[string]  `json:"id,omitzero"`
	PrefetchUserbrowser param.Opt[bool]    `json:"prefetch_userbrowser,omitzero"`
	Retry               param.Opt[bool]    `json:"retry,omitzero"`
	Timeout             param.Opt[float64] `json:"timeout,omitzero"`
	paramObj
}

func (r CrawlRootParamsExtractOptionsSession) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsExtractOptionsSession
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRootParamsExtractOptionsSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRootParamsExtractOptionsSkillUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRootParamsExtractOptionsSkillUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlRootParamsExtractOptionsSkillUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRootParamsExtractOptionsSkillUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Userbrowser creation template configuration
//
// The property Name is required.
type CrawlRootParamsExtractOptionsTemplate struct {
	Name   string         `json:"name,required"`
	Params map[string]any `json:"params,omitzero"`
	paramObj
}

func (r CrawlRootParamsExtractOptionsTemplate) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsExtractOptionsTemplate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRootParamsExtractOptionsTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pre-rendered userbrowser creation template configuration
//
// The properties ID, AllowedParameterNames, RenderFlowRendered are required.
type CrawlRootParamsExtractOptionsUserbrowserCreationTemplateRendered struct {
	ID                    string           `json:"id,required"`
	AllowedParameterNames []string         `json:"allowed_parameter_names,omitzero,required"`
	RenderFlowRendered    []map[string]any `json:"render_flow_rendered,omitzero,required"`
	paramObj
}

func (r CrawlRootParamsExtractOptionsUserbrowserCreationTemplateRendered) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRootParamsExtractOptionsUserbrowserCreationTemplateRendered
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRootParamsExtractOptionsUserbrowserCreationTemplateRendered) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sitemap and other methods will be used together to find URLs.
type CrawlRootParamsSitemap string

const (
	CrawlRootParamsSitemapSkip    CrawlRootParamsSitemap = "skip"
	CrawlRootParamsSitemapInclude CrawlRootParamsSitemap = "include"
	CrawlRootParamsSitemapOnly    CrawlRootParamsSitemap = "only"
)
