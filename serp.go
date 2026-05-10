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

// SerpService contains methods and other services that help with interacting with
// the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSerpService] method instead.
type SerpService struct {
	Options []option.RequestOption
}

// NewSerpService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSerpService(opts ...option.RequestOption) (r SerpService) {
	r = SerpService{}
	r.Options = opts
	return
}

// SERP
func (r *SerpService) Run(ctx context.Context, body SerpRunParams, opts ...option.RequestOption) (res *SerpRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/serp"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// SERP Async Endpoint
func (r *SerpService) RunAsync(ctx context.Context, body SerpRunAsyncParams, opts ...option.RequestOption) (res *SerpRunAsyncResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/serp/async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// SERP Batch Endpoint
func (r *SerpService) RunBatch(ctx context.Context, body SerpRunBatchParams, opts ...option.RequestOption) (res *SerpRunBatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/serp/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SerpRunResponse struct {
	Data     SerpRunResponseData     `json:"data" api:"required"`
	Metadata SerpRunResponseMetadata `json:"metadata" api:"required"`
	// The status of the task.
	//
	// Any of "success", "skipped", "fatal", "error", "postponed", "ignored",
	// "rejected", "blocked".
	Status SerpRunResponseStatus `json:"status" api:"required"`
	// Unique identifier for the task.
	TaskID string `json:"task_id" api:"required"`
	// The final URL.
	URL   string               `json:"url" api:"required"`
	Debug SerpRunResponseDebug `json:"debug"`
	// Pagination information if applicable.
	Pagination SerpRunResponsePaginationUnion `json:"pagination"`
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
func (r SerpRunResponse) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunResponseData struct {
	// Browser actions execution results. Present only when browser_actions were
	// specified in the request.
	BrowserActions SerpRunResponseDataBrowserActions `json:"browser_actions"`
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
	NetworkCapture []SerpRunResponseDataNetworkCapture `json:"network_capture"`
	// Individual HTML content of each pagination page, before merging.
	PagesHTML []string `json:"pages_html"`
	// The parsing results extracted from the HTML & network content.
	Parsing SerpRunResponseDataParsingUnion `json:"parsing"`
	// The list of redirects that occurred during the task.
	Redirects []SerpRunResponseDataRedirect `json:"redirects"`
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
func (r SerpRunResponseData) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser actions execution results. Present only when browser_actions were
// specified in the request.
type SerpRunResponseDataBrowserActions struct {
	Results       []SerpRunResponseDataBrowserActionsResult `json:"results" api:"required"`
	Success       bool                                      `json:"success" api:"required"`
	TotalDuration float64                                   `json:"total_duration" api:"required"`
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
func (r SerpRunResponseDataBrowserActions) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponseDataBrowserActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunResponseDataBrowserActionsResult struct {
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
func (r SerpRunResponseDataBrowserActionsResult) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponseDataBrowserActionsResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunResponseDataNetworkCapture struct {
	Filter       SerpRunResponseDataNetworkCaptureFilter   `json:"filter" api:"required"`
	Results      []SerpRunResponseDataNetworkCaptureResult `json:"results" api:"required"`
	ErrorMessage string                                    `json:"errorMessage"`
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
func (r SerpRunResponseDataNetworkCapture) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponseDataNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunResponseDataNetworkCaptureFilter struct {
	Validation           bool    `json:"validation" api:"required"`
	WaitForRequestsCount float64 `json:"wait_for_requests_count" api:"required"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method"`
	// Resource type for network capture filtering
	ResourceType                SerpRunResponseDataNetworkCaptureFilterResourceTypeUnion `json:"resource_type"`
	StatusCode                  SerpRunResponseDataNetworkCaptureFilterStatusCodeUnion   `json:"status_code"`
	URL                         SerpRunResponseDataNetworkCaptureFilterURL               `json:"url"`
	WaitForRequestsCountTimeout float64                                                  `json:"wait_for_requests_count_timeout"`
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
func (r SerpRunResponseDataNetworkCaptureFilter) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponseDataNetworkCaptureFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SerpRunResponseDataNetworkCaptureFilterResourceTypeUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfSerpRunResponseDataNetworkCaptureFilterResourceTypeString
// OfSerpRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray]
type SerpRunResponseDataNetworkCaptureFilterResourceTypeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfSerpRunResponseDataNetworkCaptureFilterResourceTypeString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfSerpRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray []string `json:",inline"`
	JSON                                                                struct {
		OfSerpRunResponseDataNetworkCaptureFilterResourceTypeString         respjson.Field
		OfSerpRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray respjson.Field
		raw                                                                 string
	} `json:"-"`
}

func (u SerpRunResponseDataNetworkCaptureFilterResourceTypeUnion) AsSerpRunResponseDataNetworkCaptureFilterResourceTypeString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SerpRunResponseDataNetworkCaptureFilterResourceTypeUnion) AsSerpRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SerpRunResponseDataNetworkCaptureFilterResourceTypeUnion) RawJSON() string { return u.JSON.raw }

func (r *SerpRunResponseDataNetworkCaptureFilterResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type for network capture filtering
type SerpRunResponseDataNetworkCaptureFilterResourceTypeString string

const (
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringDocument           SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "document"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringStylesheet         SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "stylesheet"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringImage              SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "image"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringMedia              SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "media"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringFont               SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "font"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringScript             SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "script"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringTexttrack          SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "texttrack"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringXhr                SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "xhr"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringFetch              SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "fetch"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringPrefetch           SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "prefetch"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringEventsource        SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "eventsource"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringWebsocket          SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "websocket"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringManifest           SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "manifest"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringSignedexchange     SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "signedexchange"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringPing               SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "ping"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringCspviolationreport SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "cspviolationreport"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringPreflight          SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "preflight"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringOther              SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "other"
	SerpRunResponseDataNetworkCaptureFilterResourceTypeStringFedcm              SerpRunResponseDataNetworkCaptureFilterResourceTypeString = "fedcm"
)

// SerpRunResponseDataNetworkCaptureFilterStatusCodeUnion contains all possible
// properties and values from [float64], [[]float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfFloatArray]
type SerpRunResponseDataNetworkCaptureFilterStatusCodeUnion struct {
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

func (u SerpRunResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SerpRunResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloatArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SerpRunResponseDataNetworkCaptureFilterStatusCodeUnion) RawJSON() string { return u.JSON.raw }

func (r *SerpRunResponseDataNetworkCaptureFilterStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunResponseDataNetworkCaptureFilterURL struct {
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
func (r SerpRunResponseDataNetworkCaptureFilterURL) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponseDataNetworkCaptureFilterURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunResponseDataNetworkCaptureResult struct {
	Request  SerpRunResponseDataNetworkCaptureResultRequest  `json:"request" api:"required"`
	Response SerpRunResponseDataNetworkCaptureResultResponse `json:"response" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Request     respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SerpRunResponseDataNetworkCaptureResult) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponseDataNetworkCaptureResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunResponseDataNetworkCaptureResultRequest struct {
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
func (r SerpRunResponseDataNetworkCaptureResultRequest) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponseDataNetworkCaptureResultRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunResponseDataNetworkCaptureResultResponse struct {
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
func (r SerpRunResponseDataNetworkCaptureResultResponse) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponseDataNetworkCaptureResultResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SerpRunResponseDataParsingUnion contains all possible properties and values from
// [SerpRunResponseDataParsingParsingSuccessResult],
// [SerpRunResponseDataParsingParsingErrorResult], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfSerpRunResponseDataParsingMapItem]
type SerpRunResponseDataParsingUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfSerpRunResponseDataParsingMapItem any `json:",inline"`
	// This field is from variant [SerpRunResponseDataParsingParsingSuccessResult].
	Entities map[string]any `json:"entities"`
	Status   string         `json:"status"`
	// This field is from variant [SerpRunResponseDataParsingParsingErrorResult].
	Error string `json:"error"`
	JSON  struct {
		OfSerpRunResponseDataParsingMapItem respjson.Field
		Entities                            respjson.Field
		Status                              respjson.Field
		Error                               respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u SerpRunResponseDataParsingUnion) AsSerpRunResponseDataParsingParsingSuccessResult() (v SerpRunResponseDataParsingParsingSuccessResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SerpRunResponseDataParsingUnion) AsSerpRunResponseDataParsingParsingErrorResult() (v SerpRunResponseDataParsingParsingErrorResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SerpRunResponseDataParsingUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SerpRunResponseDataParsingUnion) RawJSON() string { return u.JSON.raw }

func (r *SerpRunResponseDataParsingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunResponseDataParsingParsingSuccessResult struct {
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
func (r SerpRunResponseDataParsingParsingSuccessResult) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponseDataParsingParsingSuccessResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunResponseDataParsingParsingErrorResult struct {
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
func (r SerpRunResponseDataParsingParsingErrorResult) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponseDataParsingParsingErrorResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunResponseDataRedirect struct {
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
func (r SerpRunResponseDataRedirect) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponseDataRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunResponseMetadata struct {
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
func (r SerpRunResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the task.
type SerpRunResponseStatus string

const (
	SerpRunResponseStatusSuccess   SerpRunResponseStatus = "success"
	SerpRunResponseStatusSkipped   SerpRunResponseStatus = "skipped"
	SerpRunResponseStatusFatal     SerpRunResponseStatus = "fatal"
	SerpRunResponseStatusError     SerpRunResponseStatus = "error"
	SerpRunResponseStatusPostponed SerpRunResponseStatus = "postponed"
	SerpRunResponseStatusIgnored   SerpRunResponseStatus = "ignored"
	SerpRunResponseStatusRejected  SerpRunResponseStatus = "rejected"
	SerpRunResponseStatusBlocked   SerpRunResponseStatus = "blocked"
)

type SerpRunResponseDebug struct {
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
func (r SerpRunResponseDebug) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponseDebug) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SerpRunResponsePaginationUnion contains all possible properties and values from
// [SerpRunResponsePaginationNextPageParams],
// [[]SerpRunResponsePaginationArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfSerpRunResponsePaginationArray]
type SerpRunResponsePaginationUnion struct {
	// This field will be present if the value is a
	// [[]SerpRunResponsePaginationArrayItem] instead of an object.
	OfSerpRunResponsePaginationArray []SerpRunResponsePaginationArrayItem `json:",inline"`
	// This field is from variant [SerpRunResponsePaginationNextPageParams].
	NextPageParams map[string]any `json:"next_page_params"`
	JSON           struct {
		OfSerpRunResponsePaginationArray respjson.Field
		NextPageParams                   respjson.Field
		raw                              string
	} `json:"-"`
}

func (u SerpRunResponsePaginationUnion) AsSerpRunResponsePaginationNextPageParams() (v SerpRunResponsePaginationNextPageParams) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SerpRunResponsePaginationUnion) AsSerpRunResponsePaginationArray() (v []SerpRunResponsePaginationArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SerpRunResponsePaginationUnion) RawJSON() string { return u.JSON.raw }

func (r *SerpRunResponsePaginationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunResponsePaginationNextPageParams struct {
	NextPageParams map[string]any `json:"next_page_params" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SerpRunResponsePaginationNextPageParams) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponsePaginationNextPageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunResponsePaginationArrayItem struct {
	NextPageParams map[string]any `json:"next_page_params" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SerpRunResponsePaginationArrayItem) RawJSON() string { return r.JSON.raw }
func (r *SerpRunResponsePaginationArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response when an async SERP task is created successfully.
type SerpRunAsyncResponse struct {
	// Status indicating the async SERP task was created successfully.
	Status constant.Success `json:"status" default:"success"`
	// The created async task details.
	Task SerpRunAsyncResponseTask `json:"task" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Task        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SerpRunAsyncResponse) RawJSON() string { return r.JSON.raw }
func (r *SerpRunAsyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The created async task details.
type SerpRunAsyncResponseTask struct {
	// Unique task identifier.
	ID    string `json:"id" api:"required"`
	Query any    `json:"_query" api:"required"`
	// Timestamp when the task was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Original input data for the task.
	Input any `json:"input" api:"required"`
	// Current state of the task.
	//
	// Any of "pending", "success", "error".
	State string `json:"state" api:"required"`
	// URL for checking the task status.
	StatusURL string `json:"status_url" api:"required" format:"uri"`
	// Account name that owns the task.
	AccountName string `json:"account_name"`
	// Any of "web", "serp", "ecommerce", "social", "media", "agent", "extract".
	APIType string `json:"api_type"`
	// Batch ID if this task is part of a batch.
	BatchID string `json:"batch_id"`
	// URL for downloading the task results.
	DownloadURL string `json:"download_url" format:"uri"`
	// Error message if the task failed.
	Error string `json:"error"`
	// Classification of the error type.
	ErrorType string `json:"error_type"`
	// Timestamp when the task was last modified.
	ModifiedAt string `json:"modified_at"`
	// Storage location of the output data.
	OutputURL string `json:"output_url"`
	// HTTP status code from the task execution.
	StatusCode float64 `json:"status_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Query       respjson.Field
		CreatedAt   respjson.Field
		Input       respjson.Field
		State       respjson.Field
		StatusURL   respjson.Field
		AccountName respjson.Field
		APIType     respjson.Field
		BatchID     respjson.Field
		DownloadURL respjson.Field
		Error       respjson.Field
		ErrorType   respjson.Field
		ModifiedAt  respjson.Field
		OutputURL   respjson.Field
		StatusCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SerpRunAsyncResponseTask) RawJSON() string { return r.JSON.raw }
func (r *SerpRunAsyncResponseTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response when a batch of SERP tasks is created successfully.
type SerpRunBatchResponse struct {
	// Unique identifier for the batch.
	BatchID string `json:"batch_id" api:"required"`
	// Number of tasks in the batch.
	BatchSize float64 `json:"batch_size" api:"required"`
	// List of created tasks.
	Tasks []SerpRunBatchResponseTask `json:"tasks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BatchID     respjson.Field
		BatchSize   respjson.Field
		Tasks       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SerpRunBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *SerpRunBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunBatchResponseTask struct {
	// Unique task identifier.
	ID    string `json:"id" api:"required"`
	Query any    `json:"_query" api:"required"`
	// Timestamp when the task was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Original input data for the task.
	Input any `json:"input" api:"required"`
	// Current state of the task.
	//
	// Any of "pending", "success", "error".
	State string `json:"state" api:"required"`
	// URL for checking the task status.
	StatusURL string `json:"status_url" api:"required" format:"uri"`
	// Account name that owns the task.
	AccountName string `json:"account_name"`
	// Any of "web", "serp", "ecommerce", "social", "media", "agent", "extract".
	APIType string `json:"api_type"`
	// Batch ID if this task is part of a batch.
	BatchID string `json:"batch_id"`
	// URL for downloading the task results.
	DownloadURL string `json:"download_url" format:"uri"`
	// Error message if the task failed.
	Error string `json:"error"`
	// Classification of the error type.
	ErrorType string `json:"error_type"`
	// Timestamp when the task was last modified.
	ModifiedAt string `json:"modified_at"`
	// Storage location of the output data.
	OutputURL string `json:"output_url"`
	// HTTP status code from the task execution.
	StatusCode float64 `json:"status_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Query       respjson.Field
		CreatedAt   respjson.Field
		Input       respjson.Field
		State       respjson.Field
		StatusURL   respjson.Field
		AccountName respjson.Field
		APIType     respjson.Field
		BatchID     respjson.Field
		DownloadURL respjson.Field
		Error       respjson.Field
		ErrorType   respjson.Field
		ModifiedAt  respjson.Field
		OutputURL   respjson.Field
		StatusCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SerpRunBatchResponseTask) RawJSON() string { return r.JSON.raw }
func (r *SerpRunBatchResponseTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunParams struct {
	// The search engine to query.
	//
	// Any of "google_search", "google_sge", "google_aio", "google_maps_search",
	// "google_maps_reviews", "google_maps_place", "google_news", "google_images",
	// "bing_search", "yandex_search".
	SearchEngine SerpRunParamsSearchEngine `json:"search_engine,omitzero" api:"required"`
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
	// Device type used for the search request.
	//
	// Any of "desktop", "mobile".
	Device SerpRunParamsDevice `json:"device,omitzero"`
	paramObj
}

func (r SerpRunParams) MarshalJSON() (data []byte, err error) {
	type shadow SerpRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SerpRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The search engine to query.
type SerpRunParamsSearchEngine string

const (
	SerpRunParamsSearchEngineGoogleSearch      SerpRunParamsSearchEngine = "google_search"
	SerpRunParamsSearchEngineGoogleSge         SerpRunParamsSearchEngine = "google_sge"
	SerpRunParamsSearchEngineGoogleAio         SerpRunParamsSearchEngine = "google_aio"
	SerpRunParamsSearchEngineGoogleMapsSearch  SerpRunParamsSearchEngine = "google_maps_search"
	SerpRunParamsSearchEngineGoogleMapsReviews SerpRunParamsSearchEngine = "google_maps_reviews"
	SerpRunParamsSearchEngineGoogleMapsPlace   SerpRunParamsSearchEngine = "google_maps_place"
	SerpRunParamsSearchEngineGoogleNews        SerpRunParamsSearchEngine = "google_news"
	SerpRunParamsSearchEngineGoogleImages      SerpRunParamsSearchEngine = "google_images"
	SerpRunParamsSearchEngineBingSearch        SerpRunParamsSearchEngine = "bing_search"
	SerpRunParamsSearchEngineYandexSearch      SerpRunParamsSearchEngine = "yandex_search"
)

// Device type used for the search request.
type SerpRunParamsDevice string

const (
	SerpRunParamsDeviceDesktop SerpRunParamsDevice = "desktop"
	SerpRunParamsDeviceMobile  SerpRunParamsDevice = "mobile"
)

type SerpRunAsyncParams struct {
	// The search engine to query.
	//
	// Any of "google_search", "google_sge", "google_aio", "google_maps_search",
	// "google_maps_reviews", "google_maps_place", "google_news", "google_images",
	// "bing_search", "yandex_search".
	SearchEngine SerpRunAsyncParamsSearchEngine `json:"search_engine,omitzero" api:"required"`
	// URL to call back when async operation completes
	CallbackURL param.Opt[string] `json:"callback_url,omitzero"`
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
	// Whether to compress stored data
	StorageCompress param.Opt[bool] `json:"storage_compress,omitzero"`
	// Custom name for the stored object
	StorageObjectName param.Opt[string] `json:"storage_object_name,omitzero"`
	// Type of storage to use for results
	StorageType param.Opt[string] `json:"storage_type,omitzero"`
	// URL for storage location
	StorageURL param.Opt[string] `json:"storage_url,omitzero"`
	// Device type used for the search request.
	//
	// Any of "desktop", "mobile".
	Device SerpRunAsyncParamsDevice `json:"device,omitzero"`
	paramObj
}

func (r SerpRunAsyncParams) MarshalJSON() (data []byte, err error) {
	type shadow SerpRunAsyncParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SerpRunAsyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The search engine to query.
type SerpRunAsyncParamsSearchEngine string

const (
	SerpRunAsyncParamsSearchEngineGoogleSearch      SerpRunAsyncParamsSearchEngine = "google_search"
	SerpRunAsyncParamsSearchEngineGoogleSge         SerpRunAsyncParamsSearchEngine = "google_sge"
	SerpRunAsyncParamsSearchEngineGoogleAio         SerpRunAsyncParamsSearchEngine = "google_aio"
	SerpRunAsyncParamsSearchEngineGoogleMapsSearch  SerpRunAsyncParamsSearchEngine = "google_maps_search"
	SerpRunAsyncParamsSearchEngineGoogleMapsReviews SerpRunAsyncParamsSearchEngine = "google_maps_reviews"
	SerpRunAsyncParamsSearchEngineGoogleMapsPlace   SerpRunAsyncParamsSearchEngine = "google_maps_place"
	SerpRunAsyncParamsSearchEngineGoogleNews        SerpRunAsyncParamsSearchEngine = "google_news"
	SerpRunAsyncParamsSearchEngineGoogleImages      SerpRunAsyncParamsSearchEngine = "google_images"
	SerpRunAsyncParamsSearchEngineBingSearch        SerpRunAsyncParamsSearchEngine = "bing_search"
	SerpRunAsyncParamsSearchEngineYandexSearch      SerpRunAsyncParamsSearchEngine = "yandex_search"
)

// Device type used for the search request.
type SerpRunAsyncParamsDevice string

const (
	SerpRunAsyncParamsDeviceDesktop SerpRunAsyncParamsDevice = "desktop"
	SerpRunAsyncParamsDeviceMobile  SerpRunAsyncParamsDevice = "mobile"
)

type SerpRunBatchParams struct {
	// Array of SERP requests. Each object can include search parameters and
	// async/storage settings.
	Inputs []SerpRunBatchParamsInput `json:"inputs,omitzero" api:"required"`
	// Shared parameters applied to the entire batch. Can include search parameters and
	// async/storage settings.
	SharedInputs SerpRunBatchParamsSharedInputs `json:"shared_inputs,omitzero"`
	paramObj
}

func (r SerpRunBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow SerpRunBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SerpRunBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SerpRunBatchParamsInput struct {
	// URL to call back when async operation completes
	CallbackURL param.Opt[string] `json:"callback_url,omitzero"`
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
	// Whether to compress stored data
	StorageCompress param.Opt[bool] `json:"storage_compress,omitzero"`
	// Custom name for the stored object
	StorageObjectName param.Opt[string] `json:"storage_object_name,omitzero"`
	// Type of storage to use for results
	StorageType param.Opt[string] `json:"storage_type,omitzero"`
	// URL for storage location
	StorageURL param.Opt[string] `json:"storage_url,omitzero"`
	// Device type used for the search request.
	//
	// Any of "desktop", "mobile".
	Device string `json:"device,omitzero"`
	// The search engine to query.
	//
	// Any of "google_search", "google_sge", "google_aio", "google_maps_search",
	// "google_maps_reviews", "google_maps_place", "google_news", "google_images",
	// "bing_search", "yandex_search".
	SearchEngine string `json:"search_engine,omitzero"`
	paramObj
}

func (r SerpRunBatchParamsInput) MarshalJSON() (data []byte, err error) {
	type shadow SerpRunBatchParamsInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SerpRunBatchParamsInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SerpRunBatchParamsInput](
		"device", "desktop", "mobile",
	)
	apijson.RegisterFieldValidator[SerpRunBatchParamsInput](
		"search_engine", "google_search", "google_sge", "google_aio", "google_maps_search", "google_maps_reviews", "google_maps_place", "google_news", "google_images", "bing_search", "yandex_search",
	)
}

// Shared parameters applied to the entire batch. Can include search parameters and
// async/storage settings.
type SerpRunBatchParamsSharedInputs struct {
	// URL to call back when async operation completes
	CallbackURL param.Opt[string] `json:"callback_url,omitzero"`
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
	// Whether to compress stored data
	StorageCompress param.Opt[bool] `json:"storage_compress,omitzero"`
	// Custom name for the stored object
	StorageObjectName param.Opt[string] `json:"storage_object_name,omitzero"`
	// Type of storage to use for results
	StorageType param.Opt[string] `json:"storage_type,omitzero"`
	// URL for storage location
	StorageURL param.Opt[string] `json:"storage_url,omitzero"`
	// Device type used for the search request.
	//
	// Any of "desktop", "mobile".
	Device string `json:"device,omitzero"`
	// The search engine to query.
	//
	// Any of "google_search", "google_sge", "google_aio", "google_maps_search",
	// "google_maps_reviews", "google_maps_place", "google_news", "google_images",
	// "bing_search", "yandex_search".
	SearchEngine string `json:"search_engine,omitzero"`
	paramObj
}

func (r SerpRunBatchParamsSharedInputs) MarshalJSON() (data []byte, err error) {
	type shadow SerpRunBatchParamsSharedInputs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SerpRunBatchParamsSharedInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SerpRunBatchParamsSharedInputs](
		"device", "desktop", "mobile",
	)
	apijson.RegisterFieldValidator[SerpRunBatchParamsSharedInputs](
		"search_engine", "google_search", "google_sge", "google_aio", "google_maps_search", "google_maps_reviews", "google_maps_place", "google_news", "google_images", "bing_search", "yandex_search",
	)
}
