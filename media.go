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

// MediaService contains methods and other services that help with interacting with
// the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMediaService] method instead.
type MediaService struct {
	Options []option.RequestOption
}

// NewMediaService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMediaService(opts ...option.RequestOption) (r MediaService) {
	r = MediaService{}
	r.Options = opts
	return
}

// Download media from a URL. Waits for the result before responding.
func (r *MediaService) Run(ctx context.Context, body MediaRunParams, opts ...option.RequestOption) (res *MediaRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/media"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Download media from a URL asynchronously. Returns a task ID immediately.
func (r *MediaService) RunAsync(ctx context.Context, body MediaRunAsyncParams, opts ...option.RequestOption) (res *MediaRunAsyncResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/media/async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type MediaRunResponse struct {
	Data     MediaRunResponseData     `json:"data" api:"required"`
	Metadata MediaRunResponseMetadata `json:"metadata" api:"required"`
	// The status of the task.
	//
	// Any of "success", "skipped", "fatal", "error", "postponed", "ignored",
	// "rejected", "blocked".
	Status MediaRunResponseStatus `json:"status" api:"required"`
	// Unique identifier for the task.
	TaskID string `json:"task_id" api:"required"`
	// The final URL.
	URL   string                `json:"url" api:"required"`
	Debug MediaRunResponseDebug `json:"debug"`
	// Pagination information if applicable.
	Pagination MediaRunResponsePaginationUnion `json:"pagination"`
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
func (r MediaRunResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaRunResponseData struct {
	// Browser actions execution results. Present only when browser_actions were
	// specified in the request.
	BrowserActions MediaRunResponseDataBrowserActions `json:"browser_actions"`
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
	NetworkCapture []MediaRunResponseDataNetworkCapture `json:"network_capture"`
	// Individual HTML content of each pagination page, before merging.
	PagesHTML []string `json:"pages_html"`
	// The parsing results extracted from the HTML & network content.
	Parsing MediaRunResponseDataParsingUnion `json:"parsing"`
	// The list of redirects that occurred during the task.
	Redirects []MediaRunResponseDataRedirect `json:"redirects"`
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
func (r MediaRunResponseData) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser actions execution results. Present only when browser_actions were
// specified in the request.
type MediaRunResponseDataBrowserActions struct {
	Results       []MediaRunResponseDataBrowserActionsResult `json:"results" api:"required"`
	Success       bool                                       `json:"success" api:"required"`
	TotalDuration float64                                    `json:"total_duration" api:"required"`
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
func (r MediaRunResponseDataBrowserActions) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponseDataBrowserActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaRunResponseDataBrowserActionsResult struct {
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
func (r MediaRunResponseDataBrowserActionsResult) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponseDataBrowserActionsResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaRunResponseDataNetworkCapture struct {
	Filter       MediaRunResponseDataNetworkCaptureFilter   `json:"filter" api:"required"`
	Results      []MediaRunResponseDataNetworkCaptureResult `json:"results" api:"required"`
	ErrorMessage string                                     `json:"errorMessage"`
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
func (r MediaRunResponseDataNetworkCapture) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponseDataNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaRunResponseDataNetworkCaptureFilter struct {
	Validation           bool    `json:"validation" api:"required"`
	WaitForRequestsCount float64 `json:"wait_for_requests_count" api:"required"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method"`
	// Resource type for network capture filtering
	ResourceType                MediaRunResponseDataNetworkCaptureFilterResourceTypeUnion `json:"resource_type"`
	StatusCode                  MediaRunResponseDataNetworkCaptureFilterStatusCodeUnion   `json:"status_code"`
	URL                         MediaRunResponseDataNetworkCaptureFilterURL               `json:"url"`
	WaitForRequestsCountTimeout float64                                                   `json:"wait_for_requests_count_timeout"`
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
func (r MediaRunResponseDataNetworkCaptureFilter) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponseDataNetworkCaptureFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MediaRunResponseDataNetworkCaptureFilterResourceTypeUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfMediaRunResponseDataNetworkCaptureFilterResourceTypeString
// OfMediaRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray]
type MediaRunResponseDataNetworkCaptureFilterResourceTypeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfMediaRunResponseDataNetworkCaptureFilterResourceTypeString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfMediaRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray []string `json:",inline"`
	JSON                                                                 struct {
		OfMediaRunResponseDataNetworkCaptureFilterResourceTypeString         respjson.Field
		OfMediaRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (u MediaRunResponseDataNetworkCaptureFilterResourceTypeUnion) AsMediaRunResponseDataNetworkCaptureFilterResourceTypeString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaRunResponseDataNetworkCaptureFilterResourceTypeUnion) AsMediaRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MediaRunResponseDataNetworkCaptureFilterResourceTypeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *MediaRunResponseDataNetworkCaptureFilterResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type for network capture filtering
type MediaRunResponseDataNetworkCaptureFilterResourceTypeString string

const (
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringDocument           MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "document"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringStylesheet         MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "stylesheet"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringImage              MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "image"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringMedia              MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "media"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringFont               MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "font"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringScript             MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "script"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringTexttrack          MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "texttrack"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringXhr                MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "xhr"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringFetch              MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "fetch"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringPrefetch           MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "prefetch"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringEventsource        MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "eventsource"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringWebsocket          MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "websocket"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringManifest           MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "manifest"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringSignedexchange     MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "signedexchange"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringPing               MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "ping"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringCspviolationreport MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "cspviolationreport"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringPreflight          MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "preflight"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringOther              MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "other"
	MediaRunResponseDataNetworkCaptureFilterResourceTypeStringFedcm              MediaRunResponseDataNetworkCaptureFilterResourceTypeString = "fedcm"
)

// MediaRunResponseDataNetworkCaptureFilterStatusCodeUnion contains all possible
// properties and values from [float64], [[]float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfFloatArray]
type MediaRunResponseDataNetworkCaptureFilterStatusCodeUnion struct {
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

func (u MediaRunResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaRunResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloatArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MediaRunResponseDataNetworkCaptureFilterStatusCodeUnion) RawJSON() string { return u.JSON.raw }

func (r *MediaRunResponseDataNetworkCaptureFilterStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaRunResponseDataNetworkCaptureFilterURL struct {
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
func (r MediaRunResponseDataNetworkCaptureFilterURL) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponseDataNetworkCaptureFilterURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaRunResponseDataNetworkCaptureResult struct {
	Request  MediaRunResponseDataNetworkCaptureResultRequest  `json:"request" api:"required"`
	Response MediaRunResponseDataNetworkCaptureResultResponse `json:"response" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Request     respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaRunResponseDataNetworkCaptureResult) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponseDataNetworkCaptureResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaRunResponseDataNetworkCaptureResultRequest struct {
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
func (r MediaRunResponseDataNetworkCaptureResultRequest) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponseDataNetworkCaptureResultRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaRunResponseDataNetworkCaptureResultResponse struct {
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
func (r MediaRunResponseDataNetworkCaptureResultResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponseDataNetworkCaptureResultResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MediaRunResponseDataParsingUnion contains all possible properties and values
// from [MediaRunResponseDataParsingParsingSuccessResult],
// [MediaRunResponseDataParsingParsingErrorResult], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfMediaRunResponseDataParsingMapItem]
type MediaRunResponseDataParsingUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfMediaRunResponseDataParsingMapItem any `json:",inline"`
	// This field is from variant [MediaRunResponseDataParsingParsingSuccessResult].
	Entities map[string]any `json:"entities"`
	Status   string         `json:"status"`
	// This field is from variant [MediaRunResponseDataParsingParsingErrorResult].
	Error string `json:"error"`
	JSON  struct {
		OfMediaRunResponseDataParsingMapItem respjson.Field
		Entities                             respjson.Field
		Status                               respjson.Field
		Error                                respjson.Field
		raw                                  string
	} `json:"-"`
}

func (u MediaRunResponseDataParsingUnion) AsMediaRunResponseDataParsingParsingSuccessResult() (v MediaRunResponseDataParsingParsingSuccessResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaRunResponseDataParsingUnion) AsMediaRunResponseDataParsingParsingErrorResult() (v MediaRunResponseDataParsingParsingErrorResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaRunResponseDataParsingUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MediaRunResponseDataParsingUnion) RawJSON() string { return u.JSON.raw }

func (r *MediaRunResponseDataParsingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaRunResponseDataParsingParsingSuccessResult struct {
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
func (r MediaRunResponseDataParsingParsingSuccessResult) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponseDataParsingParsingSuccessResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaRunResponseDataParsingParsingErrorResult struct {
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
func (r MediaRunResponseDataParsingParsingErrorResult) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponseDataParsingParsingErrorResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaRunResponseDataRedirect struct {
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
func (r MediaRunResponseDataRedirect) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponseDataRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaRunResponseMetadata struct {
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
func (r MediaRunResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the task.
type MediaRunResponseStatus string

const (
	MediaRunResponseStatusSuccess   MediaRunResponseStatus = "success"
	MediaRunResponseStatusSkipped   MediaRunResponseStatus = "skipped"
	MediaRunResponseStatusFatal     MediaRunResponseStatus = "fatal"
	MediaRunResponseStatusError     MediaRunResponseStatus = "error"
	MediaRunResponseStatusPostponed MediaRunResponseStatus = "postponed"
	MediaRunResponseStatusIgnored   MediaRunResponseStatus = "ignored"
	MediaRunResponseStatusRejected  MediaRunResponseStatus = "rejected"
	MediaRunResponseStatusBlocked   MediaRunResponseStatus = "blocked"
)

type MediaRunResponseDebug struct {
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
func (r MediaRunResponseDebug) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponseDebug) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// MediaRunResponsePaginationUnion contains all possible properties and values from
// [MediaRunResponsePaginationNextPageParams],
// [[]MediaRunResponsePaginationArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfMediaRunResponsePaginationArray]
type MediaRunResponsePaginationUnion struct {
	// This field will be present if the value is a
	// [[]MediaRunResponsePaginationArrayItem] instead of an object.
	OfMediaRunResponsePaginationArray []MediaRunResponsePaginationArrayItem `json:",inline"`
	// This field is from variant [MediaRunResponsePaginationNextPageParams].
	NextPageParams map[string]any `json:"next_page_params"`
	JSON           struct {
		OfMediaRunResponsePaginationArray respjson.Field
		NextPageParams                    respjson.Field
		raw                               string
	} `json:"-"`
}

func (u MediaRunResponsePaginationUnion) AsMediaRunResponsePaginationNextPageParams() (v MediaRunResponsePaginationNextPageParams) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u MediaRunResponsePaginationUnion) AsMediaRunResponsePaginationArray() (v []MediaRunResponsePaginationArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u MediaRunResponsePaginationUnion) RawJSON() string { return u.JSON.raw }

func (r *MediaRunResponsePaginationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaRunResponsePaginationNextPageParams struct {
	NextPageParams map[string]any `json:"next_page_params" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaRunResponsePaginationNextPageParams) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponsePaginationNextPageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaRunResponsePaginationArrayItem struct {
	NextPageParams map[string]any `json:"next_page_params" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaRunResponsePaginationArrayItem) RawJSON() string { return r.JSON.raw }
func (r *MediaRunResponsePaginationArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response when an async media download task is created successfully.
type MediaRunAsyncResponse struct {
	// Status indicating the async task was created successfully.
	Status constant.Success `json:"status" default:"success"`
	// The created async task details.
	Task MediaRunAsyncResponseTask `json:"task" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Task        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaRunAsyncResponse) RawJSON() string { return r.JSON.raw }
func (r *MediaRunAsyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The created async task details.
type MediaRunAsyncResponseTask struct {
	// Unique task identifier.
	ID    string `json:"id" api:"required"`
	Query any    `json:"_query" api:"required"`
	// Timestamp when the task was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Original input data for the task.
	Input any `json:"input" api:"required"`
	// Current state of the task.
	//
	// Any of "pending", "queued", "in_progress", "success", "error".
	State string `json:"state" api:"required"`
	// URL for checking the task status.
	StatusURL string `json:"status_url" api:"required" format:"uri"`
	// Account name that owns the task.
	AccountName string `json:"account_name"`
	// Any of "web", "serp", "ecommerce", "social", "media", "agent", "extract",
	// "fast-serp", "labs".
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
	// Queue name the task was submitted to.
	Queue string `json:"queue"`
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
		Queue       respjson.Field
		StatusCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaRunAsyncResponseTask) RawJSON() string { return r.JSON.raw }
func (r *MediaRunAsyncResponseTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MediaRunParams struct {
	URL               string                `json:"url" api:"required" format:"uri"`
	Country           param.Opt[string]     `json:"country,omitzero"`
	Locale            param.Opt[string]     `json:"locale,omitzero"`
	ExpectedMimeTypes []string              `json:"expected_mime_types,omitzero"`
	Storage           MediaRunParamsStorage `json:"storage,omitzero"`
	paramObj
}

func (r MediaRunParams) MarshalJSON() (data []byte, err error) {
	type shadow MediaRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type MediaRunParamsStorage struct {
	URL        string            `json:"url" api:"required"`
	ObjectName param.Opt[string] `json:"object_name,omitzero"`
	// Any of "s3", "gcs", "do", "oci".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r MediaRunParamsStorage) MarshalJSON() (data []byte, err error) {
	type shadow MediaRunParamsStorage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaRunParamsStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MediaRunParamsStorage](
		"type", "s3", "gcs", "do", "oci",
	)
}

type MediaRunAsyncParams struct {
	URL string `json:"url" api:"required" format:"uri"`
	// URL to call back when async operation completes
	CallbackURL param.Opt[string] `json:"callback_url,omitzero"`
	Country     param.Opt[string] `json:"country,omitzero"`
	Locale      param.Opt[string] `json:"locale,omitzero"`
	// Whether to compress stored data
	StorageCompress param.Opt[bool] `json:"storage_compress,omitzero"`
	// Custom name for the stored object
	StorageObjectName param.Opt[string] `json:"storage_object_name,omitzero"`
	// Type of storage to use for results
	StorageType param.Opt[string] `json:"storage_type,omitzero"`
	// URL for storage location
	StorageURL        param.Opt[string]          `json:"storage_url,omitzero"`
	ExpectedMimeTypes []string                   `json:"expected_mime_types,omitzero"`
	Storage           MediaRunAsyncParamsStorage `json:"storage,omitzero"`
	paramObj
}

func (r MediaRunAsyncParams) MarshalJSON() (data []byte, err error) {
	type shadow MediaRunAsyncParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaRunAsyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property URL is required.
type MediaRunAsyncParamsStorage struct {
	URL        string            `json:"url" api:"required"`
	ObjectName param.Opt[string] `json:"object_name,omitzero"`
	// Any of "s3", "gcs", "do", "oci".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r MediaRunAsyncParamsStorage) MarshalJSON() (data []byte, err error) {
	type shadow MediaRunAsyncParamsStorage
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaRunAsyncParamsStorage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[MediaRunAsyncParamsStorage](
		"type", "s3", "gcs", "do", "oci",
	)
}
