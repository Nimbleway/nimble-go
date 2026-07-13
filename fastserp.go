// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomnimblewaynimblego

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/Nimbleway/nimble-go/internal/apijson"
	"github.com/Nimbleway/nimble-go/internal/requestconfig"
	"github.com/Nimbleway/nimble-go/option"
	"github.com/Nimbleway/nimble-go/packages/param"
	"github.com/Nimbleway/nimble-go/packages/respjson"
	"github.com/Nimbleway/nimble-go/shared/constant"
)

// FastSerpService contains methods and other services that help with interacting
// with the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFastSerpService] method instead.
type FastSerpService struct {
	Options []option.RequestOption
}

// NewFastSerpService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFastSerpService(opts ...option.RequestOption) (r FastSerpService) {
	r = FastSerpService{}
	r.Options = opts
	return
}

// Fast SERP
func (r *FastSerpService) Run(ctx context.Context, body FastSerpRunParams, opts ...option.RequestOption) (res *FastSerpRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/fast-serp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type FastSerpRunResponse struct {
	Data     FastSerpRunResponseData     `json:"data" api:"required"`
	Metadata FastSerpRunResponseMetadata `json:"metadata" api:"required"`
	// The status of the task.
	//
	// Any of "success", "skipped", "fatal", "error", "postponed", "ignored",
	// "rejected", "blocked".
	Status FastSerpRunResponseStatus `json:"status" api:"required"`
	// Unique identifier for the task.
	TaskID string `json:"task_id" api:"required"`
	// The final URL.
	URL   string                   `json:"url" api:"required"`
	Debug FastSerpRunResponseDebug `json:"debug"`
	// Pagination information if applicable.
	Pagination FastSerpRunResponsePaginationUnion `json:"pagination"`
	// The HTTP status code of the task.
	StatusCode float64 `json:"status_code"`
	// List of warnings generated during the task.
	Warnings []string `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Metadata    respjson.Field
		Status      respjson.Field
		TaskID      respjson.Field
		URL         respjson.Field
		Debug       respjson.Field
		Pagination  respjson.Field
		StatusCode  respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponse) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FastSerpRunResponseData struct {
	// Browser actions execution results. Present only when browser_actions were
	// specified in the request.
	BrowserActions FastSerpRunResponseDataBrowserActions `json:"browser_actions"`
	// The cookies collected from browser actions during the task.
	Cookies []any `json:"cookies"`
	// The evaluation results from browser actions during the task.
	Eval []any `json:"eval"`
	// The http requests from browser actions made during the task.
	Fetch []any `json:"fetch"`
	// The headers received during the task.
	Headers map[string]string `json:"headers"`
	// The HTML content of the page.
	HTML string `json:"html"`
	// List of all unique URLs found on the page.
	Links []string `json:"links"`
	// The Markdown version of the HTML content.
	Markdown string `json:"markdown"`
	// The network capture data collected during the task.
	NetworkCapture []FastSerpRunResponseDataNetworkCapture `json:"network_capture"`
	// Individual HTML content of each pagination page, before merging.
	PagesHTML []string `json:"pages_html"`
	// The parsing results extracted from the HTML & network content.
	Parsing FastSerpRunResponseDataParsingUnion `json:"parsing"`
	// The list of redirects that occurred during the task.
	Redirects []FastSerpRunResponseDataRedirect `json:"redirects"`
	// Screenshots taken during the task, from browser actions, or the screenshot
	// format.
	Screenshots []any `json:"screenshots"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrowserActions respjson.Field
		Cookies        respjson.Field
		Eval           respjson.Field
		Fetch          respjson.Field
		Headers        respjson.Field
		HTML           respjson.Field
		Links          respjson.Field
		Markdown       respjson.Field
		NetworkCapture respjson.Field
		PagesHTML      respjson.Field
		Parsing        respjson.Field
		Redirects      respjson.Field
		Screenshots    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponseData) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser actions execution results. Present only when browser_actions were
// specified in the request.
type FastSerpRunResponseDataBrowserActions struct {
	Results       []FastSerpRunResponseDataBrowserActionsResult `json:"results" api:"required"`
	Success       bool                                          `json:"success" api:"required"`
	TotalDuration float64                                       `json:"total_duration" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results       respjson.Field
		Success       respjson.Field
		TotalDuration respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponseDataBrowserActions) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponseDataBrowserActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FastSerpRunResponseDataBrowserActionsResult struct {
	Duration float64 `json:"duration" api:"required"`
	// Any of "goto", "wait", "wait_for_element", "wait_for_navigation", "click",
	// "fill", "press", "scroll", "auto_scroll", "screenshot", "get_cookies", "eval",
	// "fetch".
	Name string `json:"name" api:"required"`
	// Any of "no-run", "in-progress", "done", "error", "skipped".
	Status string `json:"status" api:"required"`
	Error  string `json:"error"`
	Result any    `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		Error       respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponseDataBrowserActionsResult) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponseDataBrowserActionsResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FastSerpRunResponseDataNetworkCapture struct {
	Filter       FastSerpRunResponseDataNetworkCaptureFilter   `json:"filter" api:"required"`
	Results      []FastSerpRunResponseDataNetworkCaptureResult `json:"results" api:"required"`
	ErrorMessage string                                        `json:"errorMessage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter       respjson.Field
		Results      respjson.Field
		ErrorMessage respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponseDataNetworkCapture) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponseDataNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FastSerpRunResponseDataNetworkCaptureFilter struct {
	Validation           bool    `json:"validation" api:"required"`
	WaitForRequestsCount float64 `json:"wait_for_requests_count" api:"required"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method"`
	// Resource type for network capture filtering
	ResourceType                FastSerpRunResponseDataNetworkCaptureFilterResourceTypeUnion `json:"resource_type"`
	StatusCode                  FastSerpRunResponseDataNetworkCaptureFilterStatusCodeUnion   `json:"status_code"`
	URL                         FastSerpRunResponseDataNetworkCaptureFilterURL               `json:"url"`
	WaitForRequestsCountTimeout float64                                                      `json:"wait_for_requests_count_timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Validation                  respjson.Field
		WaitForRequestsCount        respjson.Field
		Method                      respjson.Field
		ResourceType                respjson.Field
		StatusCode                  respjson.Field
		URL                         respjson.Field
		WaitForRequestsCountTimeout respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponseDataNetworkCaptureFilter) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponseDataNetworkCaptureFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FastSerpRunResponseDataNetworkCaptureFilterResourceTypeUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFastSerpRunResponseDataNetworkCaptureFilterResourceTypeString
// OfFastSerpRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray]
type FastSerpRunResponseDataNetworkCaptureFilterResourceTypeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfFastSerpRunResponseDataNetworkCaptureFilterResourceTypeString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfFastSerpRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray []string `json:",inline"`
	JSON                                                                    struct {
		OfFastSerpRunResponseDataNetworkCaptureFilterResourceTypeString         respjson.Field
		OfFastSerpRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray respjson.Field
		raw                                                                     string
	} `json:"-"`
}

func (u FastSerpRunResponseDataNetworkCaptureFilterResourceTypeUnion) AsFastSerpRunResponseDataNetworkCaptureFilterResourceTypeString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FastSerpRunResponseDataNetworkCaptureFilterResourceTypeUnion) AsFastSerpRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FastSerpRunResponseDataNetworkCaptureFilterResourceTypeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FastSerpRunResponseDataNetworkCaptureFilterResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type for network capture filtering
type FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString string

const (
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringDocument           FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "document"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringStylesheet         FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "stylesheet"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringImage              FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "image"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringMedia              FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "media"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringFont               FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "font"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringScript             FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "script"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringTexttrack          FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "texttrack"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringXhr                FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "xhr"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringFetch              FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "fetch"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringPrefetch           FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "prefetch"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringEventsource        FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "eventsource"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringWebsocket          FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "websocket"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringManifest           FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "manifest"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringSignedexchange     FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "signedexchange"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringPing               FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "ping"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringCspviolationreport FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "cspviolationreport"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringPreflight          FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "preflight"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringOther              FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "other"
	FastSerpRunResponseDataNetworkCaptureFilterResourceTypeStringFedcm              FastSerpRunResponseDataNetworkCaptureFilterResourceTypeString = "fedcm"
)

// FastSerpRunResponseDataNetworkCaptureFilterStatusCodeUnion contains all possible
// properties and values from [float64], [[]float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfFloatArray]
type FastSerpRunResponseDataNetworkCaptureFilterStatusCodeUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [[]float64] instead of an object.
	OfFloatArray []float64 `json:",inline"`
	JSON         struct {
		OfFloat      respjson.Field
		OfFloatArray respjson.Field
		raw          string
	} `json:"-"`
}

func (u FastSerpRunResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FastSerpRunResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloatArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FastSerpRunResponseDataNetworkCaptureFilterStatusCodeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *FastSerpRunResponseDataNetworkCaptureFilterStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FastSerpRunResponseDataNetworkCaptureFilterURL struct {
	// Any of "exact", "contains".
	Type  string `json:"type" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponseDataNetworkCaptureFilterURL) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponseDataNetworkCaptureFilterURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FastSerpRunResponseDataNetworkCaptureResult struct {
	Request  FastSerpRunResponseDataNetworkCaptureResultRequest  `json:"request" api:"required"`
	Response FastSerpRunResponseDataNetworkCaptureResultResponse `json:"response" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Request     respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponseDataNetworkCaptureResult) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponseDataNetworkCaptureResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FastSerpRunResponseDataNetworkCaptureResultRequest struct {
	Headers map[string]string `json:"headers" api:"required"`
	Method  string            `json:"method" api:"required"`
	// Resource type for network capture filtering
	//
	// Any of "document", "stylesheet", "image", "media", "font", "script",
	// "texttrack", "xhr", "fetch", "prefetch", "eventsource", "websocket", "manifest",
	// "signedexchange", "ping", "cspviolationreport", "preflight", "other", "fedcm".
	ResourceType string `json:"resource_type" api:"required"`
	URL          string `json:"url" api:"required"`
	Body         string `json:"body"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers      respjson.Field
		Method       respjson.Field
		ResourceType respjson.Field
		URL          respjson.Field
		Body         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponseDataNetworkCaptureResultRequest) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponseDataNetworkCaptureResultRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FastSerpRunResponseDataNetworkCaptureResultResponse struct {
	Body    string            `json:"body" api:"required"`
	Headers map[string]string `json:"headers" api:"required"`
	// Any of "none", "base64".
	Serialization string  `json:"serialization" api:"required"`
	Status        float64 `json:"status" api:"required"`
	StatusText    string  `json:"status_text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body          respjson.Field
		Headers       respjson.Field
		Serialization respjson.Field
		Status        respjson.Field
		StatusText    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponseDataNetworkCaptureResultResponse) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponseDataNetworkCaptureResultResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FastSerpRunResponseDataParsingUnion contains all possible properties and values
// from [FastSerpRunResponseDataParsingParsingSuccessResult],
// [FastSerpRunResponseDataParsingParsingErrorResult], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFastSerpRunResponseDataParsingMapItem]
type FastSerpRunResponseDataParsingUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfFastSerpRunResponseDataParsingMapItem any `json:",inline"`
	// This field is from variant [FastSerpRunResponseDataParsingParsingSuccessResult].
	Entities map[string]any `json:"entities"`
	Status   string         `json:"status"`
	// This field is from variant [FastSerpRunResponseDataParsingParsingErrorResult].
	Error string `json:"error"`
	JSON  struct {
		OfFastSerpRunResponseDataParsingMapItem respjson.Field
		Entities                                respjson.Field
		Status                                  respjson.Field
		Error                                   respjson.Field
		raw                                     string
	} `json:"-"`
}

func (u FastSerpRunResponseDataParsingUnion) AsFastSerpRunResponseDataParsingParsingSuccessResult() (v FastSerpRunResponseDataParsingParsingSuccessResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FastSerpRunResponseDataParsingUnion) AsFastSerpRunResponseDataParsingParsingErrorResult() (v FastSerpRunResponseDataParsingParsingErrorResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FastSerpRunResponseDataParsingUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FastSerpRunResponseDataParsingUnion) RawJSON() string { return u.JSON.raw }

func (r *FastSerpRunResponseDataParsingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FastSerpRunResponseDataParsingParsingSuccessResult struct {
	Entities map[string]any   `json:"entities" api:"required"`
	Status   constant.Success `json:"status" default:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponseDataParsingParsingSuccessResult) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponseDataParsingParsingSuccessResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FastSerpRunResponseDataParsingParsingErrorResult struct {
	Error  string         `json:"error" api:"required"`
	Status constant.Error `json:"status" default:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponseDataParsingParsingErrorResult) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponseDataParsingParsingErrorResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FastSerpRunResponseDataRedirect struct {
	StatusCode float64 `json:"status_code" api:"required"`
	URL        string  `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StatusCode  respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponseDataRedirect) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponseDataRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FastSerpRunResponseMetadata struct {
	// The name of the agent used for the query.
	Agent string `json:"agent"`
	// The driver used for the task.
	Driver string `json:"driver"`
	// The localization identifier for the query.
	LocalizationID string `json:"localization_id"`
	// The duration in milliseconds of the query processing.
	QueryDuration float64 `json:"query_duration"`
	// The time when the query was received.
	QueryTime string `json:"query_time"`
	// Additional response parameters.
	ResponseParameters any `json:"response_parameters"`
	// A tag associated with the query.
	Tag string `json:"tag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent              respjson.Field
		Driver             respjson.Field
		LocalizationID     respjson.Field
		QueryDuration      respjson.Field
		QueryTime          respjson.Field
		ResponseParameters respjson.Field
		Tag                respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the task.
type FastSerpRunResponseStatus string

const (
	FastSerpRunResponseStatusSuccess   FastSerpRunResponseStatus = "success"
	FastSerpRunResponseStatusSkipped   FastSerpRunResponseStatus = "skipped"
	FastSerpRunResponseStatusFatal     FastSerpRunResponseStatus = "fatal"
	FastSerpRunResponseStatusError     FastSerpRunResponseStatus = "error"
	FastSerpRunResponseStatusPostponed FastSerpRunResponseStatus = "postponed"
	FastSerpRunResponseStatusIgnored   FastSerpRunResponseStatus = "ignored"
	FastSerpRunResponseStatusRejected  FastSerpRunResponseStatus = "rejected"
	FastSerpRunResponseStatusBlocked   FastSerpRunResponseStatus = "blocked"
)

type FastSerpRunResponseDebug struct {
	// Performance metrics collected during the task.
	PerformanceMetrics map[string]float64 `json:"performance_metrics"`
	// Total bytes used by the proxy during the task.
	ProxyTotalBytesUsage float64 `json:"proxy_total_bytes_usage"`
	// The transformed output after applying any transformations.
	TransformedOutput any `json:"transformed_output"`
	// The userbrowser instance using during the task.
	Userbrowser any `json:"userbrowser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PerformanceMetrics   respjson.Field
		ProxyTotalBytesUsage respjson.Field
		TransformedOutput    respjson.Field
		Userbrowser          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponseDebug) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponseDebug) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// FastSerpRunResponsePaginationUnion contains all possible properties and values
// from [FastSerpRunResponsePaginationNextPageParams],
// [[]FastSerpRunResponsePaginationArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFastSerpRunResponsePaginationArray]
type FastSerpRunResponsePaginationUnion struct {
	// This field will be present if the value is a
	// [[]FastSerpRunResponsePaginationArrayItem] instead of an object.
	OfFastSerpRunResponsePaginationArray []FastSerpRunResponsePaginationArrayItem `json:",inline"`
	// This field is from variant [FastSerpRunResponsePaginationNextPageParams].
	NextPageParams map[string]any `json:"next_page_params"`
	JSON           struct {
		OfFastSerpRunResponsePaginationArray respjson.Field
		NextPageParams                       respjson.Field
		raw                                  string
	} `json:"-"`
}

func (u FastSerpRunResponsePaginationUnion) AsFastSerpRunResponsePaginationNextPageParams() (v FastSerpRunResponsePaginationNextPageParams) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u FastSerpRunResponsePaginationUnion) AsFastSerpRunResponsePaginationArray() (v []FastSerpRunResponsePaginationArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u FastSerpRunResponsePaginationUnion) RawJSON() string { return u.JSON.raw }

func (r *FastSerpRunResponsePaginationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FastSerpRunResponsePaginationNextPageParams struct {
	NextPageParams map[string]any `json:"next_page_params" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponsePaginationNextPageParams) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponsePaginationNextPageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FastSerpRunResponsePaginationArrayItem struct {
	NextPageParams map[string]any `json:"next_page_params" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FastSerpRunResponsePaginationArrayItem) RawJSON() string { return r.JSON.raw }
func (r *FastSerpRunResponsePaginationArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FastSerpRunParams struct {
	// The search engine to query.
	//
	// Any of "google_search", "google_sge", "google_aio", "google_maps_search",
	// "google_maps_reviews", "google_maps_place", "google_news", "google_images",
	// "bing_search", "yandex_search".
	SearchEngine FastSerpRunParamsSearchEngine `json:"search_engine,omitzero" api:"required"`
	// ISO Alpha-2 country code used to access the target search engine (e.g. US, DE,
	// GB).
	Country param.Opt[string] `json:"country,omitzero"`
	// Top-level domain for the search engine (e.g. "com", "co.uk", "de").
	Domain param.Opt[string] `json:"domain,omitzero"`
	// Locale used for the search request.
	Locale param.Opt[string] `json:"locale,omitzero"`
	// Geo-location for the search (canonical Google location name).
	Location param.Opt[string] `json:"location,omitzero"`
	// Number of results to return (1–100).
	NumResults param.Opt[int64] `json:"num_results,omitzero"`
	// The result page number for pagination.
	Page param.Opt[int64] `json:"page,omitzero"`
	// When true, the SERP response is parsed into structured JSON.
	Parse param.Opt[bool] `json:"parse,omitzero"`
	// The search keyword or phrase to query.
	Query param.Opt[string] `json:"query,omitzero"`
	// Whether to render the page in a browser before extracting.
	Render param.Opt[bool] `json:"render,omitzero"`
	// When true, disables Google result filtering (filter=0) so omitted/duplicate and
	// highly similar pages are also returned. Applies to Google search engines.
	ShowHiddenResults param.Opt[bool] `json:"show_hidden_results,omitzero"`
	// Device type used for the search request.
	//
	// Any of "desktop", "mobile".
	Device FastSerpRunParamsDevice `json:"device,omitzero"`
	paramObj
}

func (r FastSerpRunParams) MarshalJSON() (data []byte, err error) {
	type shadow FastSerpRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FastSerpRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The search engine to query.
type FastSerpRunParamsSearchEngine string

const (
	FastSerpRunParamsSearchEngineGoogleSearch      FastSerpRunParamsSearchEngine = "google_search"
	FastSerpRunParamsSearchEngineGoogleSge         FastSerpRunParamsSearchEngine = "google_sge"
	FastSerpRunParamsSearchEngineGoogleAio         FastSerpRunParamsSearchEngine = "google_aio"
	FastSerpRunParamsSearchEngineGoogleMapsSearch  FastSerpRunParamsSearchEngine = "google_maps_search"
	FastSerpRunParamsSearchEngineGoogleMapsReviews FastSerpRunParamsSearchEngine = "google_maps_reviews"
	FastSerpRunParamsSearchEngineGoogleMapsPlace   FastSerpRunParamsSearchEngine = "google_maps_place"
	FastSerpRunParamsSearchEngineGoogleNews        FastSerpRunParamsSearchEngine = "google_news"
	FastSerpRunParamsSearchEngineGoogleImages      FastSerpRunParamsSearchEngine = "google_images"
	FastSerpRunParamsSearchEngineBingSearch        FastSerpRunParamsSearchEngine = "bing_search"
	FastSerpRunParamsSearchEngineYandexSearch      FastSerpRunParamsSearchEngine = "yandex_search"
)

// Device type used for the search request.
type FastSerpRunParamsDevice string

const (
	FastSerpRunParamsDeviceDesktop FastSerpRunParamsDevice = "desktop"
	FastSerpRunParamsDeviceMobile  FastSerpRunParamsDevice = "mobile"
)
