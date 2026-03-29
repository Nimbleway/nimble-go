// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomnimblewaynimblego

import (
	"encoding/json"

	"github.com/Nimbleway/nimble-go/internal/apijson"
	"github.com/Nimbleway/nimble-go/packages/param"
	"github.com/Nimbleway/nimble-go/packages/respjson"
	"github.com/Nimbleway/nimble-go/shared"
	"github.com/Nimbleway/nimble-go/shared/constant"
)

type ExtractResponse struct {
	Data     ExtractResponseData     `json:"data" api:"required"`
	Metadata ExtractResponseMetadata `json:"metadata" api:"required"`
	// The status of the task.
	//
	// Any of "success", "skipped", "fatal", "error", "postponed", "ignored",
	// "rejected", "blocked".
	Status ExtractResponseStatus `json:"status" api:"required"`
	// Unique identifier for the task.
	TaskID string `json:"task_id" api:"required"`
	// The final URL.
	URL   string               `json:"url" api:"required"`
	Debug ExtractResponseDebug `json:"debug"`
	// Pagination information if applicable.
	Pagination ExtractResponsePaginationUnion `json:"pagination"`
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
func (r ExtractResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractResponseData struct {
	// Browser actions execution results. Present only when browser_actions were
	// specified in the request.
	BrowserActions ExtractResponseDataBrowserActions `json:"browser_actions"`
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
	// The Markdown version of the HTML content.
	Markdown string `json:"markdown"`
	// The network capture data collected during the task.
	NetworkCapture []ExtractResponseDataNetworkCapture `json:"network_capture"`
	// Individual HTML content of each pagination page, before merging.
	PagesHTML []string `json:"pages_html"`
	// The parsing results extracted from the HTML & network content.
	Parsing ExtractResponseDataParsingUnion `json:"parsing"`
	// The list of redirects that occurred during the task.
	Redirects []ExtractResponseDataRedirect `json:"redirects"`
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
func (r ExtractResponseData) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser actions execution results. Present only when browser_actions were
// specified in the request.
type ExtractResponseDataBrowserActions struct {
	Results       []ExtractResponseDataBrowserActionsResult `json:"results" api:"required"`
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
func (r ExtractResponseDataBrowserActions) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseDataBrowserActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractResponseDataBrowserActionsResult struct {
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
func (r ExtractResponseDataBrowserActionsResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseDataBrowserActionsResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractResponseDataNetworkCapture struct {
	Filter       ExtractResponseDataNetworkCaptureFilter   `json:"filter" api:"required"`
	Results      []ExtractResponseDataNetworkCaptureResult `json:"results" api:"required"`
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
func (r ExtractResponseDataNetworkCapture) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseDataNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractResponseDataNetworkCaptureFilter struct {
	Validation           bool    `json:"validation" api:"required"`
	WaitForRequestsCount float64 `json:"wait_for_requests_count" api:"required"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method"`
	// Resource type for network capture filtering
	ResourceType                ExtractResponseDataNetworkCaptureFilterResourceTypeUnion `json:"resource_type"`
	StatusCode                  ExtractResponseDataNetworkCaptureFilterStatusCodeUnion   `json:"status_code"`
	URL                         ExtractResponseDataNetworkCaptureFilterURL               `json:"url"`
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
func (r ExtractResponseDataNetworkCaptureFilter) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseDataNetworkCaptureFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractResponseDataNetworkCaptureFilterResourceTypeUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractResponseDataNetworkCaptureFilterResourceTypeString
// OfExtractResponseDataNetworkCaptureFilterResourceTypeArrayItemArray]
type ExtractResponseDataNetworkCaptureFilterResourceTypeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfExtractResponseDataNetworkCaptureFilterResourceTypeString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfExtractResponseDataNetworkCaptureFilterResourceTypeArrayItemArray []string `json:",inline"`
	JSON                                                                struct {
		OfExtractResponseDataNetworkCaptureFilterResourceTypeString         respjson.Field
		OfExtractResponseDataNetworkCaptureFilterResourceTypeArrayItemArray respjson.Field
		raw                                                                 string
	} `json:"-"`
}

func (u ExtractResponseDataNetworkCaptureFilterResourceTypeUnion) AsExtractResponseDataNetworkCaptureFilterResourceTypeString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractResponseDataNetworkCaptureFilterResourceTypeUnion) AsExtractResponseDataNetworkCaptureFilterResourceTypeArrayItemArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractResponseDataNetworkCaptureFilterResourceTypeUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractResponseDataNetworkCaptureFilterResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type for network capture filtering
type ExtractResponseDataNetworkCaptureFilterResourceTypeString string

const (
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringDocument           ExtractResponseDataNetworkCaptureFilterResourceTypeString = "document"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringStylesheet         ExtractResponseDataNetworkCaptureFilterResourceTypeString = "stylesheet"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringImage              ExtractResponseDataNetworkCaptureFilterResourceTypeString = "image"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringMedia              ExtractResponseDataNetworkCaptureFilterResourceTypeString = "media"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringFont               ExtractResponseDataNetworkCaptureFilterResourceTypeString = "font"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringScript             ExtractResponseDataNetworkCaptureFilterResourceTypeString = "script"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringTexttrack          ExtractResponseDataNetworkCaptureFilterResourceTypeString = "texttrack"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringXhr                ExtractResponseDataNetworkCaptureFilterResourceTypeString = "xhr"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringFetch              ExtractResponseDataNetworkCaptureFilterResourceTypeString = "fetch"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringPrefetch           ExtractResponseDataNetworkCaptureFilterResourceTypeString = "prefetch"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringEventsource        ExtractResponseDataNetworkCaptureFilterResourceTypeString = "eventsource"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringWebsocket          ExtractResponseDataNetworkCaptureFilterResourceTypeString = "websocket"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringManifest           ExtractResponseDataNetworkCaptureFilterResourceTypeString = "manifest"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringSignedexchange     ExtractResponseDataNetworkCaptureFilterResourceTypeString = "signedexchange"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringPing               ExtractResponseDataNetworkCaptureFilterResourceTypeString = "ping"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringCspviolationreport ExtractResponseDataNetworkCaptureFilterResourceTypeString = "cspviolationreport"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringPreflight          ExtractResponseDataNetworkCaptureFilterResourceTypeString = "preflight"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringOther              ExtractResponseDataNetworkCaptureFilterResourceTypeString = "other"
	ExtractResponseDataNetworkCaptureFilterResourceTypeStringFedcm              ExtractResponseDataNetworkCaptureFilterResourceTypeString = "fedcm"
)

// ExtractResponseDataNetworkCaptureFilterStatusCodeUnion contains all possible
// properties and values from [float64], [[]float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfFloatArray]
type ExtractResponseDataNetworkCaptureFilterStatusCodeUnion struct {
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

func (u ExtractResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloatArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractResponseDataNetworkCaptureFilterStatusCodeUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractResponseDataNetworkCaptureFilterStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractResponseDataNetworkCaptureFilterURL struct {
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
func (r ExtractResponseDataNetworkCaptureFilterURL) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseDataNetworkCaptureFilterURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractResponseDataNetworkCaptureResult struct {
	Request  ExtractResponseDataNetworkCaptureResultRequest  `json:"request" api:"required"`
	Response ExtractResponseDataNetworkCaptureResultResponse `json:"response" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Request     respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractResponseDataNetworkCaptureResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseDataNetworkCaptureResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractResponseDataNetworkCaptureResultRequest struct {
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
func (r ExtractResponseDataNetworkCaptureResultRequest) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseDataNetworkCaptureResultRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractResponseDataNetworkCaptureResultResponse struct {
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
func (r ExtractResponseDataNetworkCaptureResultResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseDataNetworkCaptureResultResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractResponseDataParsingUnion contains all possible properties and values from
// [ExtractResponseDataParsingParsingSuccessResult],
// [ExtractResponseDataParsingParsingErrorResult], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractResponseDataParsingMapItem]
type ExtractResponseDataParsingUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExtractResponseDataParsingMapItem any `json:",inline"`
	// This field is from variant [ExtractResponseDataParsingParsingSuccessResult].
	Entities map[string]any `json:"entities"`
	Status   string         `json:"status"`
	// This field is from variant [ExtractResponseDataParsingParsingErrorResult].
	Error string `json:"error"`
	JSON  struct {
		OfExtractResponseDataParsingMapItem respjson.Field
		Entities                            respjson.Field
		Status                              respjson.Field
		Error                               respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u ExtractResponseDataParsingUnion) AsExtractResponseDataParsingParsingSuccessResult() (v ExtractResponseDataParsingParsingSuccessResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractResponseDataParsingUnion) AsExtractResponseDataParsingParsingErrorResult() (v ExtractResponseDataParsingParsingErrorResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractResponseDataParsingUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractResponseDataParsingUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractResponseDataParsingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractResponseDataParsingParsingSuccessResult struct {
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
func (r ExtractResponseDataParsingParsingSuccessResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseDataParsingParsingSuccessResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractResponseDataParsingParsingErrorResult struct {
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
func (r ExtractResponseDataParsingParsingErrorResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseDataParsingParsingErrorResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractResponseDataRedirect struct {
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
func (r ExtractResponseDataRedirect) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseDataRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractResponseMetadata struct {
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
func (r ExtractResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the task.
type ExtractResponseStatus string

const (
	ExtractResponseStatusSuccess   ExtractResponseStatus = "success"
	ExtractResponseStatusSkipped   ExtractResponseStatus = "skipped"
	ExtractResponseStatusFatal     ExtractResponseStatus = "fatal"
	ExtractResponseStatusError     ExtractResponseStatus = "error"
	ExtractResponseStatusPostponed ExtractResponseStatus = "postponed"
	ExtractResponseStatusIgnored   ExtractResponseStatus = "ignored"
	ExtractResponseStatusRejected  ExtractResponseStatus = "rejected"
	ExtractResponseStatusBlocked   ExtractResponseStatus = "blocked"
)

type ExtractResponseDebug struct {
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
func (r ExtractResponseDebug) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseDebug) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractResponsePaginationUnion contains all possible properties and values from
// [ExtractResponsePaginationNextPageParams],
// [[]ExtractResponsePaginationArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractResponsePaginationArray]
type ExtractResponsePaginationUnion struct {
	// This field will be present if the value is a
	// [[]ExtractResponsePaginationArrayItem] instead of an object.
	OfExtractResponsePaginationArray []ExtractResponsePaginationArrayItem `json:",inline"`
	// This field is from variant [ExtractResponsePaginationNextPageParams].
	NextPageParams map[string]any `json:"next_page_params"`
	JSON           struct {
		OfExtractResponsePaginationArray respjson.Field
		NextPageParams                   respjson.Field
		raw                              string
	} `json:"-"`
}

func (u ExtractResponsePaginationUnion) AsExtractResponsePaginationNextPageParams() (v ExtractResponsePaginationNextPageParams) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractResponsePaginationUnion) AsExtractResponsePaginationArray() (v []ExtractResponsePaginationArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractResponsePaginationUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractResponsePaginationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractResponsePaginationNextPageParams struct {
	NextPageParams map[string]any `json:"next_page_params" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractResponsePaginationNextPageParams) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponsePaginationNextPageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractResponsePaginationArrayItem struct {
	NextPageParams map[string]any `json:"next_page_params" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractResponsePaginationArrayItem) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponsePaginationArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response when an async extract task is created successfully.
type ExtractAsyncResponse struct {
	// Status indicating the async task was created successfully.
	Status constant.Success `json:"status" default:"success"`
	// The created async task details.
	Task ExtractAsyncResponseTask `json:"task" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Task        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractAsyncResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractAsyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The created async task details.
type ExtractAsyncResponseTask struct {
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
func (r ExtractAsyncResponseTask) RawJSON() string { return r.JSON.raw }
func (r *ExtractAsyncResponseTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response when a batch of extract tasks is created successfully.
type ExtractBatchResponse struct {
	// Unique identifier for the batch.
	BatchID string `json:"batch_id" api:"required"`
	// Number of tasks in the batch.
	BatchSize float64 `json:"batch_size" api:"required"`
	// List of created tasks.
	Tasks []ExtractBatchResponseTask `json:"tasks" api:"required"`
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
func (r ExtractBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractBatchResponseTask struct {
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
func (r ExtractBatchResponseTask) RawJSON() string { return r.JSON.raw }
func (r *ExtractBatchResponseTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response schema for map requests.
type MapResponse struct {
	// Array of mapped links with optional titles and descriptions.
	Links []MapResponseLink `json:"links" api:"required"`
	// Indicates if the map request was successful.
	Success bool `json:"success" api:"required"`
	// Unique identifier for the map task.
	TaskID string `json:"task_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Links       respjson.Field
		Success     respjson.Field
		TaskID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MapResponse) RawJSON() string { return r.JSON.raw }
func (r *MapResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MapResponseLink struct {
	URL         string `json:"url" api:"required" format:"uri"`
	Description string `json:"description"`
	Title       string `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Description respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MapResponseLink) RawJSON() string { return r.JSON.raw }
func (r *MapResponseLink) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response model from SearchService with results and optional LLM answer.
//
// Note: request_id is always a valid UUID generated internally by the middleware,
// so no validation is needed.
type SearchResponse struct {
	// Unique identifier for this request (UUID)
	RequestID string                 `json:"request_id" api:"required"`
	Results   []SearchResponseResult `json:"results" api:"required"`
	// Number of results returned
	TotalResults int64  `json:"total_results" api:"required"`
	Answer       string `json:"answer" api:"nullable"`
	// Citations mapping citation markers to result indices
	AnswerCitations []SearchResponseAnswerCitation `json:"answer_citations" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RequestID       respjson.Field
		Results         respjson.Field
		TotalResults    respjson.Field
		Answer          respjson.Field
		AnswerCitations respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponse) RawJSON() string { return r.JSON.raw }
func (r *SearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Unified result model for all search types (SERP and WSA).
//
// This model provides a consistent structure for search results, with
// platform-specific data in additional_data and typed metadata.
type SearchResponseResult struct {
	Content     string `json:"content" api:"required"`
	Description string `json:"description" api:"required"`
	// Metadata for SERP-based search results (general, news, location).
	Metadata SearchResponseResultMetadataUnion `json:"metadata" api:"required"`
	Title    string                            `json:"title" api:"required"`
	URL      string                            `json:"url" api:"required"`
	// Platform-specific fields (e.g., price, rating, publish_date). Omitted from
	// response when no extra data.
	AdditionalData map[string]any `json:"additional_data" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content        respjson.Field
		Description    respjson.Field
		Metadata       respjson.Field
		Title          respjson.Field
		URL            respjson.Field
		AdditionalData respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponseResult) RawJSON() string { return r.JSON.raw }
func (r *SearchResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SearchResponseResultMetadataUnion contains all possible properties and values
// from [SearchResponseResultMetadataSerpMetadata],
// [SearchResponseResultMetadataWsaMetadata].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SearchResponseResultMetadataUnion struct {
	// This field is from variant [SearchResponseResultMetadataSerpMetadata].
	Country string `json:"country"`
	// This field is from variant [SearchResponseResultMetadataSerpMetadata].
	EntityType string `json:"entity_type"`
	// This field is from variant [SearchResponseResultMetadataSerpMetadata].
	Locale string `json:"locale"`
	// This field is from variant [SearchResponseResultMetadataSerpMetadata].
	Position int64 `json:"position"`
	// This field is from variant [SearchResponseResultMetadataSerpMetadata].
	Driver string `json:"driver"`
	// This field is from variant [SearchResponseResultMetadataWsaMetadata].
	AgentName string `json:"agent_name"`
	JSON      struct {
		Country    respjson.Field
		EntityType respjson.Field
		Locale     respjson.Field
		Position   respjson.Field
		Driver     respjson.Field
		AgentName  respjson.Field
		raw        string
	} `json:"-"`
}

func (u SearchResponseResultMetadataUnion) AsSerpMetadata() (v SearchResponseResultMetadataSerpMetadata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SearchResponseResultMetadataUnion) AsWsaMetadata() (v SearchResponseResultMetadataWsaMetadata) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SearchResponseResultMetadataUnion) RawJSON() string { return u.JSON.raw }

func (r *SearchResponseResultMetadataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for SERP-based search results (general, news, location).
type SearchResponseResultMetadataSerpMetadata struct {
	Country    string `json:"country" api:"required"`
	EntityType string `json:"entity_type" api:"required"`
	Locale     string `json:"locale" api:"required"`
	Position   int64  `json:"position" api:"required"`
	Driver     string `json:"driver" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Country     respjson.Field
		EntityType  respjson.Field
		Locale      respjson.Field
		Position    respjson.Field
		Driver      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponseResultMetadataSerpMetadata) RawJSON() string { return r.JSON.raw }
func (r *SearchResponseResultMetadataSerpMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for WSA-based search results.
type SearchResponseResultMetadataWsaMetadata struct {
	AgentName string `json:"agent_name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgentName   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponseResultMetadataWsaMetadata) RawJSON() string { return r.JSON.raw }
func (r *SearchResponseResultMetadataWsaMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Citation model that maps citation markers to result indices.
type SearchResponseAnswerCitation struct {
	// Citation marker number (e.g., 1 for [1])
	Marker int64 `json:"marker" api:"required"`
	// Zero-based index into the results array
	ResultIndex int64 `json:"result_index" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Marker      respjson.Field
		ResultIndex respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SearchResponseAnswerCitation) RawJSON() string { return r.JSON.raw }
func (r *SearchResponseAnswerCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractParams struct {
	// Target URL to scrape
	URL string `json:"url" api:"required"`
	// City for geolocation
	City param.Opt[string] `json:"city,omitzero"`
	// Whether to automatically handle cookie consent headers
	ConsentHeader param.Opt[bool] `json:"consent_header,omitzero"`
	// Whether to use HTTP/2 protocol
	Http2 param.Opt[bool] `json:"http2,omitzero"`
	// Whether to emulate XMLHttpRequest behavior
	IsXhr param.Opt[bool] `json:"is_xhr,omitzero"`
	// Whether to parse the response content
	Parse param.Opt[bool] `json:"parse,omitzero"`
	// Whether to render JavaScript content using a browser
	Render param.Opt[bool] `json:"render,omitzero"`
	// Request timeout in milliseconds
	RequestTimeout param.Opt[float64] `json:"request_timeout,omitzero"`
	// User-defined tag for request identification
	Tag param.Opt[string] `json:"tag,omitzero"`
	// Browser type to emulate
	Browser ExtractParamsBrowserUnion `json:"browser,omitzero"`
	// Array of browser automation actions to execute sequentially
	BrowserActions []ExtractParamsBrowserActionUnion `json:"browser_actions,omitzero"`
	// Browser cookies as array of cookie objects
	Cookies ExtractParamsCookiesUnion `json:"cookies,omitzero"`
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
	Country ExtractParamsCountry `json:"country,omitzero"`
	// Device type for browser emulation
	//
	// Any of "desktop", "mobile", "tablet".
	Device ExtractParamsDevice `json:"device,omitzero"`
	// Browser driver to use
	//
	// Any of "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro".
	Driver ExtractParamsDriver `json:"driver,omitzero"`
	// Expected HTTP status codes for successful requests
	ExpectedStatusCodes []int64 `json:"expected_status_codes,omitzero"`
	// List of acceptable response formats in order of preference
	//
	// Any of "html", "markdown", "screenshot", "headers".
	Formats []string `json:"formats,omitzero"`
	// Custom HTTP headers to include in the request
	Headers map[string]ExtractParamsHeaderUnion `json:"headers,omitzero"`
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
	Locale ExtractParamsLocale `json:"locale,omitzero"`
	// HTTP method for the request
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method ExtractParamsMethod `json:"method,omitzero"`
	// Filters for capturing network traffic
	NetworkCapture []ExtractParamsNetworkCapture `json:"network_capture,omitzero"`
	// Operating system to emulate
	//
	// Any of "windows", "mac os", "linux", "android", "ios".
	Os ExtractParamsOs `json:"os,omitzero"`
	// Custom parser configuration as a key-value map
	Parser ExtractParamsParserUnion `json:"parser,omitzero"`
	// Referrer policy for the request
	//
	// Any of "random", "no-referer", "same-origin", "google", "bing", "facebook",
	// "twitter", "instagram".
	ReferrerType ExtractParamsReferrerType `json:"referrer_type,omitzero"`
	Session      ExtractParamsSession      `json:"session,omitzero"`
	// Skills or capabilities required for the request
	Skill ExtractParamsSkillUnion `json:"skill,omitzero"`
	// US state for geolocation (only valid when country is US)
	//
	// Any of "AL", "AK", "AS", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA",
	// "GU", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI",
	// "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "MP",
	// "OH", "OK", "OR", "PA", "PR", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA",
	// "VI", "WA", "WV", "WI", "WY".
	State ExtractParamsState `json:"state,omitzero"`
	paramObj
}

func (r ExtractParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserString)
	OfExtractsBrowserString param.Opt[string]           `json:",omitzero,inline"`
	OfExtractsBrowserObject *ExtractParamsBrowserObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserString, u.OfExtractsBrowserObject)
}
func (u *ExtractParamsBrowserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserString) {
		return &u.OfExtractsBrowserString
	} else if !param.IsOmitted(u.OfExtractsBrowserObject) {
		return u.OfExtractsBrowserObject
	}
	return nil
}

// Browser type to emulate
type ExtractParamsBrowserString string

const (
	ExtractParamsBrowserStringChrome  ExtractParamsBrowserString = "chrome"
	ExtractParamsBrowserStringFirefox ExtractParamsBrowserString = "firefox"
)

// The property Name is required.
type ExtractParamsBrowserObject struct {
	// Any of "chrome", "firefox".
	Name string `json:"name,omitzero" api:"required"`
	// Specific browser version to emulate
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r ExtractParamsBrowserObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractParamsBrowserObject](
		"name", "chrome", "firefox",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionUnion struct {
	OfAutoScrollAction        *shared.AutoScrollActionParam        `json:",omitzero,inline"`
	OfClickAction             *shared.ClickActionParam             `json:",omitzero,inline"`
	OfEvalAction              *shared.EvalActionParam              `json:",omitzero,inline"`
	OfFetchAction             *shared.FetchActionParam             `json:",omitzero,inline"`
	OfFillAction              *shared.FillActionParam              `json:",omitzero,inline"`
	OfGetCookiesAction        *shared.GetCookiesActionParam        `json:",omitzero,inline"`
	OfGotoAction              *shared.GotoActionParam              `json:",omitzero,inline"`
	OfPressAction             *shared.PressActionParam             `json:",omitzero,inline"`
	OfScreenshotAction        *shared.ScreenshotActionParam        `json:",omitzero,inline"`
	OfScrollAction            *shared.ScrollActionParam            `json:",omitzero,inline"`
	OfWaitAction              *shared.WaitActionParam              `json:",omitzero,inline"`
	OfWaitForElementAction    *shared.WaitForElementActionParam    `json:",omitzero,inline"`
	OfWaitForNavigationAction *shared.WaitForNavigationActionParam `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAutoScrollAction,
		u.OfClickAction,
		u.OfEvalAction,
		u.OfFetchAction,
		u.OfFillAction,
		u.OfGetCookiesAction,
		u.OfGotoAction,
		u.OfPressAction,
		u.OfScreenshotAction,
		u.OfScrollAction,
		u.OfWaitAction,
		u.OfWaitForElementAction,
		u.OfWaitForNavigationAction)
}
func (u *ExtractParamsBrowserActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionUnion) asAny() any {
	if !param.IsOmitted(u.OfAutoScrollAction) {
		return u.OfAutoScrollAction
	} else if !param.IsOmitted(u.OfClickAction) {
		return u.OfClickAction
	} else if !param.IsOmitted(u.OfEvalAction) {
		return u.OfEvalAction
	} else if !param.IsOmitted(u.OfFetchAction) {
		return u.OfFetchAction
	} else if !param.IsOmitted(u.OfFillAction) {
		return u.OfFillAction
	} else if !param.IsOmitted(u.OfGetCookiesAction) {
		return u.OfGetCookiesAction
	} else if !param.IsOmitted(u.OfGotoAction) {
		return u.OfGotoAction
	} else if !param.IsOmitted(u.OfPressAction) {
		return u.OfPressAction
	} else if !param.IsOmitted(u.OfScreenshotAction) {
		return u.OfScreenshotAction
	} else if !param.IsOmitted(u.OfScrollAction) {
		return u.OfScrollAction
	} else if !param.IsOmitted(u.OfWaitAction) {
		return u.OfWaitAction
	} else if !param.IsOmitted(u.OfWaitForElementAction) {
		return u.OfWaitForElementAction
	} else if !param.IsOmitted(u.OfWaitForNavigationAction) {
		return u.OfWaitForNavigationAction
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsCookiesUnion struct {
	OfExtractsCookiesArray []ExtractParamsCookiesArrayItem `json:",omitzero,inline"`
	OfString               param.Opt[string]               `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsCookiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsCookiesArray, u.OfString)
}
func (u *ExtractParamsCookiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsCookiesUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsCookiesArray) {
		return &u.OfExtractsCookiesArray
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type ExtractParamsCookiesArrayItem struct {
	Creation      param.Opt[string]                        `json:"creation,omitzero"`
	Domain        param.Opt[string]                        `json:"domain,omitzero"`
	HostOnly      param.Opt[bool]                          `json:"hostOnly,omitzero"`
	HTTPOnly      param.Opt[bool]                          `json:"httpOnly,omitzero"`
	LastAccessed  param.Opt[string]                        `json:"lastAccessed,omitzero"`
	Path          param.Opt[string]                        `json:"path,omitzero"`
	PathIsDefault param.Opt[bool]                          `json:"pathIsDefault,omitzero"`
	Expires       param.Opt[string]                        `json:"expires,omitzero"`
	Name          param.Opt[string]                        `json:"name,omitzero"`
	Secure        param.Opt[bool]                          `json:"secure,omitzero"`
	Value         param.Opt[string]                        `json:"value,omitzero"`
	Extensions    []string                                 `json:"extensions,omitzero"`
	MaxAge        ExtractParamsCookiesArrayItemMaxAgeUnion `json:"maxAge,omitzero"`
	// Any of "strict", "lax", "none".
	SameSite    string         `json:"sameSite,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r ExtractParamsCookiesArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsCookiesArrayItem
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ExtractParamsCookiesArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractParamsCookiesArrayItem](
		"sameSite", "strict", "lax", "none",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsCookiesArrayItemMaxAgeUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsCookiesArrayItemMaxAgeString)
	OfExtractsCookiesArrayItemMaxAgeString param.Opt[ExtractParamsCookiesArrayItemMaxAgeString] `json:",omitzero,inline"`
	OfFloat                                param.Opt[float64]                                   `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsCookiesArrayItemMaxAgeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsCookiesArrayItemMaxAgeString, u.OfFloat)
}
func (u *ExtractParamsCookiesArrayItemMaxAgeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsCookiesArrayItemMaxAgeUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsCookiesArrayItemMaxAgeString) {
		return &u.OfExtractsCookiesArrayItemMaxAgeString
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type ExtractParamsCookiesArrayItemMaxAgeString string

const (
	ExtractParamsCookiesArrayItemMaxAgeStringInfinity      ExtractParamsCookiesArrayItemMaxAgeString = "Infinity"
	ExtractParamsCookiesArrayItemMaxAgeStringMinusInfinity ExtractParamsCookiesArrayItemMaxAgeString = "-Infinity"
)

// Country code for geolocation and proxy selection
type ExtractParamsCountry string

const (
	ExtractParamsCountryAd  ExtractParamsCountry = "AD"
	ExtractParamsCountryAe  ExtractParamsCountry = "AE"
	ExtractParamsCountryAf  ExtractParamsCountry = "AF"
	ExtractParamsCountryAg  ExtractParamsCountry = "AG"
	ExtractParamsCountryAI  ExtractParamsCountry = "AI"
	ExtractParamsCountryAl  ExtractParamsCountry = "AL"
	ExtractParamsCountryAm  ExtractParamsCountry = "AM"
	ExtractParamsCountryAo  ExtractParamsCountry = "AO"
	ExtractParamsCountryAq  ExtractParamsCountry = "AQ"
	ExtractParamsCountryAr  ExtractParamsCountry = "AR"
	ExtractParamsCountryAs  ExtractParamsCountry = "AS"
	ExtractParamsCountryAt  ExtractParamsCountry = "AT"
	ExtractParamsCountryAu  ExtractParamsCountry = "AU"
	ExtractParamsCountryAw  ExtractParamsCountry = "AW"
	ExtractParamsCountryAx  ExtractParamsCountry = "AX"
	ExtractParamsCountryAz  ExtractParamsCountry = "AZ"
	ExtractParamsCountryBa  ExtractParamsCountry = "BA"
	ExtractParamsCountryBb  ExtractParamsCountry = "BB"
	ExtractParamsCountryBd  ExtractParamsCountry = "BD"
	ExtractParamsCountryBe  ExtractParamsCountry = "BE"
	ExtractParamsCountryBf  ExtractParamsCountry = "BF"
	ExtractParamsCountryBg  ExtractParamsCountry = "BG"
	ExtractParamsCountryBh  ExtractParamsCountry = "BH"
	ExtractParamsCountryBi  ExtractParamsCountry = "BI"
	ExtractParamsCountryBj  ExtractParamsCountry = "BJ"
	ExtractParamsCountryBl  ExtractParamsCountry = "BL"
	ExtractParamsCountryBm  ExtractParamsCountry = "BM"
	ExtractParamsCountryBn  ExtractParamsCountry = "BN"
	ExtractParamsCountryBo  ExtractParamsCountry = "BO"
	ExtractParamsCountryBq  ExtractParamsCountry = "BQ"
	ExtractParamsCountryBr  ExtractParamsCountry = "BR"
	ExtractParamsCountryBs  ExtractParamsCountry = "BS"
	ExtractParamsCountryBt  ExtractParamsCountry = "BT"
	ExtractParamsCountryBv  ExtractParamsCountry = "BV"
	ExtractParamsCountryBw  ExtractParamsCountry = "BW"
	ExtractParamsCountryBy  ExtractParamsCountry = "BY"
	ExtractParamsCountryBz  ExtractParamsCountry = "BZ"
	ExtractParamsCountryCa  ExtractParamsCountry = "CA"
	ExtractParamsCountryCc  ExtractParamsCountry = "CC"
	ExtractParamsCountryCd  ExtractParamsCountry = "CD"
	ExtractParamsCountryCf  ExtractParamsCountry = "CF"
	ExtractParamsCountryCg  ExtractParamsCountry = "CG"
	ExtractParamsCountryCh  ExtractParamsCountry = "CH"
	ExtractParamsCountryCi  ExtractParamsCountry = "CI"
	ExtractParamsCountryCk  ExtractParamsCountry = "CK"
	ExtractParamsCountryCl  ExtractParamsCountry = "CL"
	ExtractParamsCountryCm  ExtractParamsCountry = "CM"
	ExtractParamsCountryCn  ExtractParamsCountry = "CN"
	ExtractParamsCountryCo  ExtractParamsCountry = "CO"
	ExtractParamsCountryCr  ExtractParamsCountry = "CR"
	ExtractParamsCountryCu  ExtractParamsCountry = "CU"
	ExtractParamsCountryCv  ExtractParamsCountry = "CV"
	ExtractParamsCountryCw  ExtractParamsCountry = "CW"
	ExtractParamsCountryCx  ExtractParamsCountry = "CX"
	ExtractParamsCountryCy  ExtractParamsCountry = "CY"
	ExtractParamsCountryCz  ExtractParamsCountry = "CZ"
	ExtractParamsCountryDe  ExtractParamsCountry = "DE"
	ExtractParamsCountryDj  ExtractParamsCountry = "DJ"
	ExtractParamsCountryDk  ExtractParamsCountry = "DK"
	ExtractParamsCountryDm  ExtractParamsCountry = "DM"
	ExtractParamsCountryDo  ExtractParamsCountry = "DO"
	ExtractParamsCountryDz  ExtractParamsCountry = "DZ"
	ExtractParamsCountryEc  ExtractParamsCountry = "EC"
	ExtractParamsCountryEe  ExtractParamsCountry = "EE"
	ExtractParamsCountryEg  ExtractParamsCountry = "EG"
	ExtractParamsCountryEh  ExtractParamsCountry = "EH"
	ExtractParamsCountryEr  ExtractParamsCountry = "ER"
	ExtractParamsCountryEs  ExtractParamsCountry = "ES"
	ExtractParamsCountryEt  ExtractParamsCountry = "ET"
	ExtractParamsCountryFi  ExtractParamsCountry = "FI"
	ExtractParamsCountryFj  ExtractParamsCountry = "FJ"
	ExtractParamsCountryFk  ExtractParamsCountry = "FK"
	ExtractParamsCountryFm  ExtractParamsCountry = "FM"
	ExtractParamsCountryFo  ExtractParamsCountry = "FO"
	ExtractParamsCountryFr  ExtractParamsCountry = "FR"
	ExtractParamsCountryGa  ExtractParamsCountry = "GA"
	ExtractParamsCountryGB  ExtractParamsCountry = "GB"
	ExtractParamsCountryGd  ExtractParamsCountry = "GD"
	ExtractParamsCountryGe  ExtractParamsCountry = "GE"
	ExtractParamsCountryGf  ExtractParamsCountry = "GF"
	ExtractParamsCountryGg  ExtractParamsCountry = "GG"
	ExtractParamsCountryGh  ExtractParamsCountry = "GH"
	ExtractParamsCountryGi  ExtractParamsCountry = "GI"
	ExtractParamsCountryGl  ExtractParamsCountry = "GL"
	ExtractParamsCountryGm  ExtractParamsCountry = "GM"
	ExtractParamsCountryGn  ExtractParamsCountry = "GN"
	ExtractParamsCountryGp  ExtractParamsCountry = "GP"
	ExtractParamsCountryGq  ExtractParamsCountry = "GQ"
	ExtractParamsCountryGr  ExtractParamsCountry = "GR"
	ExtractParamsCountryGs  ExtractParamsCountry = "GS"
	ExtractParamsCountryGt  ExtractParamsCountry = "GT"
	ExtractParamsCountryGu  ExtractParamsCountry = "GU"
	ExtractParamsCountryGw  ExtractParamsCountry = "GW"
	ExtractParamsCountryGy  ExtractParamsCountry = "GY"
	ExtractParamsCountryHk  ExtractParamsCountry = "HK"
	ExtractParamsCountryHm  ExtractParamsCountry = "HM"
	ExtractParamsCountryHn  ExtractParamsCountry = "HN"
	ExtractParamsCountryHr  ExtractParamsCountry = "HR"
	ExtractParamsCountryHt  ExtractParamsCountry = "HT"
	ExtractParamsCountryHu  ExtractParamsCountry = "HU"
	ExtractParamsCountryID  ExtractParamsCountry = "ID"
	ExtractParamsCountryIe  ExtractParamsCountry = "IE"
	ExtractParamsCountryIl  ExtractParamsCountry = "IL"
	ExtractParamsCountryIm  ExtractParamsCountry = "IM"
	ExtractParamsCountryIn  ExtractParamsCountry = "IN"
	ExtractParamsCountryIo  ExtractParamsCountry = "IO"
	ExtractParamsCountryIq  ExtractParamsCountry = "IQ"
	ExtractParamsCountryIr  ExtractParamsCountry = "IR"
	ExtractParamsCountryIs  ExtractParamsCountry = "IS"
	ExtractParamsCountryIt  ExtractParamsCountry = "IT"
	ExtractParamsCountryJe  ExtractParamsCountry = "JE"
	ExtractParamsCountryJm  ExtractParamsCountry = "JM"
	ExtractParamsCountryJo  ExtractParamsCountry = "JO"
	ExtractParamsCountryJp  ExtractParamsCountry = "JP"
	ExtractParamsCountryKe  ExtractParamsCountry = "KE"
	ExtractParamsCountryKg  ExtractParamsCountry = "KG"
	ExtractParamsCountryKh  ExtractParamsCountry = "KH"
	ExtractParamsCountryKi  ExtractParamsCountry = "KI"
	ExtractParamsCountryKm  ExtractParamsCountry = "KM"
	ExtractParamsCountryKn  ExtractParamsCountry = "KN"
	ExtractParamsCountryKp  ExtractParamsCountry = "KP"
	ExtractParamsCountryKr  ExtractParamsCountry = "KR"
	ExtractParamsCountryKw  ExtractParamsCountry = "KW"
	ExtractParamsCountryKy  ExtractParamsCountry = "KY"
	ExtractParamsCountryKz  ExtractParamsCountry = "KZ"
	ExtractParamsCountryLa  ExtractParamsCountry = "LA"
	ExtractParamsCountryLb  ExtractParamsCountry = "LB"
	ExtractParamsCountryLc  ExtractParamsCountry = "LC"
	ExtractParamsCountryLi  ExtractParamsCountry = "LI"
	ExtractParamsCountryLk  ExtractParamsCountry = "LK"
	ExtractParamsCountryLr  ExtractParamsCountry = "LR"
	ExtractParamsCountryLs  ExtractParamsCountry = "LS"
	ExtractParamsCountryLt  ExtractParamsCountry = "LT"
	ExtractParamsCountryLu  ExtractParamsCountry = "LU"
	ExtractParamsCountryLv  ExtractParamsCountry = "LV"
	ExtractParamsCountryLy  ExtractParamsCountry = "LY"
	ExtractParamsCountryMa  ExtractParamsCountry = "MA"
	ExtractParamsCountryMc  ExtractParamsCountry = "MC"
	ExtractParamsCountryMd  ExtractParamsCountry = "MD"
	ExtractParamsCountryMe  ExtractParamsCountry = "ME"
	ExtractParamsCountryMf  ExtractParamsCountry = "MF"
	ExtractParamsCountryMg  ExtractParamsCountry = "MG"
	ExtractParamsCountryMh  ExtractParamsCountry = "MH"
	ExtractParamsCountryMk  ExtractParamsCountry = "MK"
	ExtractParamsCountryMl  ExtractParamsCountry = "ML"
	ExtractParamsCountryMm  ExtractParamsCountry = "MM"
	ExtractParamsCountryMn  ExtractParamsCountry = "MN"
	ExtractParamsCountryMo  ExtractParamsCountry = "MO"
	ExtractParamsCountryMp  ExtractParamsCountry = "MP"
	ExtractParamsCountryMq  ExtractParamsCountry = "MQ"
	ExtractParamsCountryMr  ExtractParamsCountry = "MR"
	ExtractParamsCountryMs  ExtractParamsCountry = "MS"
	ExtractParamsCountryMt  ExtractParamsCountry = "MT"
	ExtractParamsCountryMu  ExtractParamsCountry = "MU"
	ExtractParamsCountryMv  ExtractParamsCountry = "MV"
	ExtractParamsCountryMw  ExtractParamsCountry = "MW"
	ExtractParamsCountryMx  ExtractParamsCountry = "MX"
	ExtractParamsCountryMy  ExtractParamsCountry = "MY"
	ExtractParamsCountryMz  ExtractParamsCountry = "MZ"
	ExtractParamsCountryNa  ExtractParamsCountry = "NA"
	ExtractParamsCountryNc  ExtractParamsCountry = "NC"
	ExtractParamsCountryNe  ExtractParamsCountry = "NE"
	ExtractParamsCountryNf  ExtractParamsCountry = "NF"
	ExtractParamsCountryNg  ExtractParamsCountry = "NG"
	ExtractParamsCountryNi  ExtractParamsCountry = "NI"
	ExtractParamsCountryNl  ExtractParamsCountry = "NL"
	ExtractParamsCountryNo  ExtractParamsCountry = "NO"
	ExtractParamsCountryNp  ExtractParamsCountry = "NP"
	ExtractParamsCountryNr  ExtractParamsCountry = "NR"
	ExtractParamsCountryNu  ExtractParamsCountry = "NU"
	ExtractParamsCountryNz  ExtractParamsCountry = "NZ"
	ExtractParamsCountryOm  ExtractParamsCountry = "OM"
	ExtractParamsCountryPa  ExtractParamsCountry = "PA"
	ExtractParamsCountryPe  ExtractParamsCountry = "PE"
	ExtractParamsCountryPf  ExtractParamsCountry = "PF"
	ExtractParamsCountryPg  ExtractParamsCountry = "PG"
	ExtractParamsCountryPh  ExtractParamsCountry = "PH"
	ExtractParamsCountryPk  ExtractParamsCountry = "PK"
	ExtractParamsCountryPl  ExtractParamsCountry = "PL"
	ExtractParamsCountryPm  ExtractParamsCountry = "PM"
	ExtractParamsCountryPn  ExtractParamsCountry = "PN"
	ExtractParamsCountryPr  ExtractParamsCountry = "PR"
	ExtractParamsCountryPs  ExtractParamsCountry = "PS"
	ExtractParamsCountryPt  ExtractParamsCountry = "PT"
	ExtractParamsCountryPw  ExtractParamsCountry = "PW"
	ExtractParamsCountryPy  ExtractParamsCountry = "PY"
	ExtractParamsCountryQa  ExtractParamsCountry = "QA"
	ExtractParamsCountryRe  ExtractParamsCountry = "RE"
	ExtractParamsCountryRo  ExtractParamsCountry = "RO"
	ExtractParamsCountryRs  ExtractParamsCountry = "RS"
	ExtractParamsCountryRu  ExtractParamsCountry = "RU"
	ExtractParamsCountryRw  ExtractParamsCountry = "RW"
	ExtractParamsCountrySa  ExtractParamsCountry = "SA"
	ExtractParamsCountrySb  ExtractParamsCountry = "SB"
	ExtractParamsCountrySc  ExtractParamsCountry = "SC"
	ExtractParamsCountrySd  ExtractParamsCountry = "SD"
	ExtractParamsCountrySe  ExtractParamsCountry = "SE"
	ExtractParamsCountrySg  ExtractParamsCountry = "SG"
	ExtractParamsCountrySh  ExtractParamsCountry = "SH"
	ExtractParamsCountrySi  ExtractParamsCountry = "SI"
	ExtractParamsCountrySj  ExtractParamsCountry = "SJ"
	ExtractParamsCountrySk  ExtractParamsCountry = "SK"
	ExtractParamsCountrySl  ExtractParamsCountry = "SL"
	ExtractParamsCountrySm  ExtractParamsCountry = "SM"
	ExtractParamsCountrySn  ExtractParamsCountry = "SN"
	ExtractParamsCountrySo  ExtractParamsCountry = "SO"
	ExtractParamsCountrySr  ExtractParamsCountry = "SR"
	ExtractParamsCountrySS  ExtractParamsCountry = "SS"
	ExtractParamsCountrySt  ExtractParamsCountry = "ST"
	ExtractParamsCountrySv  ExtractParamsCountry = "SV"
	ExtractParamsCountrySx  ExtractParamsCountry = "SX"
	ExtractParamsCountrySy  ExtractParamsCountry = "SY"
	ExtractParamsCountrySz  ExtractParamsCountry = "SZ"
	ExtractParamsCountryTc  ExtractParamsCountry = "TC"
	ExtractParamsCountryTd  ExtractParamsCountry = "TD"
	ExtractParamsCountryTf  ExtractParamsCountry = "TF"
	ExtractParamsCountryTg  ExtractParamsCountry = "TG"
	ExtractParamsCountryTh  ExtractParamsCountry = "TH"
	ExtractParamsCountryTj  ExtractParamsCountry = "TJ"
	ExtractParamsCountryTk  ExtractParamsCountry = "TK"
	ExtractParamsCountryTl  ExtractParamsCountry = "TL"
	ExtractParamsCountryTm  ExtractParamsCountry = "TM"
	ExtractParamsCountryTn  ExtractParamsCountry = "TN"
	ExtractParamsCountryTo  ExtractParamsCountry = "TO"
	ExtractParamsCountryTr  ExtractParamsCountry = "TR"
	ExtractParamsCountryTt  ExtractParamsCountry = "TT"
	ExtractParamsCountryTv  ExtractParamsCountry = "TV"
	ExtractParamsCountryTw  ExtractParamsCountry = "TW"
	ExtractParamsCountryTz  ExtractParamsCountry = "TZ"
	ExtractParamsCountryUa  ExtractParamsCountry = "UA"
	ExtractParamsCountryUg  ExtractParamsCountry = "UG"
	ExtractParamsCountryUm  ExtractParamsCountry = "UM"
	ExtractParamsCountryUs  ExtractParamsCountry = "US"
	ExtractParamsCountryUy  ExtractParamsCountry = "UY"
	ExtractParamsCountryUz  ExtractParamsCountry = "UZ"
	ExtractParamsCountryVa  ExtractParamsCountry = "VA"
	ExtractParamsCountryVc  ExtractParamsCountry = "VC"
	ExtractParamsCountryVe  ExtractParamsCountry = "VE"
	ExtractParamsCountryVg  ExtractParamsCountry = "VG"
	ExtractParamsCountryVi  ExtractParamsCountry = "VI"
	ExtractParamsCountryVn  ExtractParamsCountry = "VN"
	ExtractParamsCountryVu  ExtractParamsCountry = "VU"
	ExtractParamsCountryWf  ExtractParamsCountry = "WF"
	ExtractParamsCountryWs  ExtractParamsCountry = "WS"
	ExtractParamsCountryXk  ExtractParamsCountry = "XK"
	ExtractParamsCountryYe  ExtractParamsCountry = "YE"
	ExtractParamsCountryYt  ExtractParamsCountry = "YT"
	ExtractParamsCountryZa  ExtractParamsCountry = "ZA"
	ExtractParamsCountryZm  ExtractParamsCountry = "ZM"
	ExtractParamsCountryZw  ExtractParamsCountry = "ZW"
	ExtractParamsCountryAll ExtractParamsCountry = "ALL"
)

// Device type for browser emulation
type ExtractParamsDevice string

const (
	ExtractParamsDeviceDesktop ExtractParamsDevice = "desktop"
	ExtractParamsDeviceMobile  ExtractParamsDevice = "mobile"
	ExtractParamsDeviceTablet  ExtractParamsDevice = "tablet"
)

// Browser driver to use
type ExtractParamsDriver string

const (
	ExtractParamsDriverVx6     ExtractParamsDriver = "vx6"
	ExtractParamsDriverVx8     ExtractParamsDriver = "vx8"
	ExtractParamsDriverVx8Pro  ExtractParamsDriver = "vx8-pro"
	ExtractParamsDriverVx10    ExtractParamsDriver = "vx10"
	ExtractParamsDriverVx10Pro ExtractParamsDriver = "vx10-pro"
	ExtractParamsDriverVx12    ExtractParamsDriver = "vx12"
	ExtractParamsDriverVx12Pro ExtractParamsDriver = "vx12-pro"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsHeaderUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsHeaderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractParamsHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsHeaderUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Locale for browser language and region settings
type ExtractParamsLocale string

const (
	ExtractParamsLocaleAaDj      ExtractParamsLocale = "aa-DJ"
	ExtractParamsLocaleAaEr      ExtractParamsLocale = "aa-ER"
	ExtractParamsLocaleAaEt      ExtractParamsLocale = "aa-ET"
	ExtractParamsLocaleAf        ExtractParamsLocale = "af"
	ExtractParamsLocaleAfNa      ExtractParamsLocale = "af-NA"
	ExtractParamsLocaleAfZa      ExtractParamsLocale = "af-ZA"
	ExtractParamsLocaleAk        ExtractParamsLocale = "ak"
	ExtractParamsLocaleAkGh      ExtractParamsLocale = "ak-GH"
	ExtractParamsLocaleAm        ExtractParamsLocale = "am"
	ExtractParamsLocaleAmEt      ExtractParamsLocale = "am-ET"
	ExtractParamsLocaleAnEs      ExtractParamsLocale = "an-ES"
	ExtractParamsLocaleAr        ExtractParamsLocale = "ar"
	ExtractParamsLocaleArAe      ExtractParamsLocale = "ar-AE"
	ExtractParamsLocaleArBh      ExtractParamsLocale = "ar-BH"
	ExtractParamsLocaleArDz      ExtractParamsLocale = "ar-DZ"
	ExtractParamsLocaleArEg      ExtractParamsLocale = "ar-EG"
	ExtractParamsLocaleArIn      ExtractParamsLocale = "ar-IN"
	ExtractParamsLocaleArIq      ExtractParamsLocale = "ar-IQ"
	ExtractParamsLocaleArJo      ExtractParamsLocale = "ar-JO"
	ExtractParamsLocaleArKw      ExtractParamsLocale = "ar-KW"
	ExtractParamsLocaleArLb      ExtractParamsLocale = "ar-LB"
	ExtractParamsLocaleArLy      ExtractParamsLocale = "ar-LY"
	ExtractParamsLocaleArMa      ExtractParamsLocale = "ar-MA"
	ExtractParamsLocaleArOm      ExtractParamsLocale = "ar-OM"
	ExtractParamsLocaleArQa      ExtractParamsLocale = "ar-QA"
	ExtractParamsLocaleArSa      ExtractParamsLocale = "ar-SA"
	ExtractParamsLocaleArSd      ExtractParamsLocale = "ar-SD"
	ExtractParamsLocaleArSy      ExtractParamsLocale = "ar-SY"
	ExtractParamsLocaleArTn      ExtractParamsLocale = "ar-TN"
	ExtractParamsLocaleArYe      ExtractParamsLocale = "ar-YE"
	ExtractParamsLocaleAs        ExtractParamsLocale = "as"
	ExtractParamsLocaleAsIn      ExtractParamsLocale = "as-IN"
	ExtractParamsLocaleAsa       ExtractParamsLocale = "asa"
	ExtractParamsLocaleAsaTz     ExtractParamsLocale = "asa-TZ"
	ExtractParamsLocaleAstEs     ExtractParamsLocale = "ast-ES"
	ExtractParamsLocaleAz        ExtractParamsLocale = "az"
	ExtractParamsLocaleAzAz      ExtractParamsLocale = "az-AZ"
	ExtractParamsLocaleAzCyrl    ExtractParamsLocale = "az-Cyrl"
	ExtractParamsLocaleAzCyrlAz  ExtractParamsLocale = "az-Cyrl-AZ"
	ExtractParamsLocaleAzLatn    ExtractParamsLocale = "az-Latn"
	ExtractParamsLocaleAzLatnAz  ExtractParamsLocale = "az-Latn-AZ"
	ExtractParamsLocaleBe        ExtractParamsLocale = "be"
	ExtractParamsLocaleBeBy      ExtractParamsLocale = "be-BY"
	ExtractParamsLocaleBem       ExtractParamsLocale = "bem"
	ExtractParamsLocaleBemZm     ExtractParamsLocale = "bem-ZM"
	ExtractParamsLocaleBerDz     ExtractParamsLocale = "ber-DZ"
	ExtractParamsLocaleBerMa     ExtractParamsLocale = "ber-MA"
	ExtractParamsLocaleBez       ExtractParamsLocale = "bez"
	ExtractParamsLocaleBezTz     ExtractParamsLocale = "bez-TZ"
	ExtractParamsLocaleBg        ExtractParamsLocale = "bg"
	ExtractParamsLocaleBgBg      ExtractParamsLocale = "bg-BG"
	ExtractParamsLocaleBhoIn     ExtractParamsLocale = "bho-IN"
	ExtractParamsLocaleBm        ExtractParamsLocale = "bm"
	ExtractParamsLocaleBmMl      ExtractParamsLocale = "bm-ML"
	ExtractParamsLocaleBn        ExtractParamsLocale = "bn"
	ExtractParamsLocaleBnBd      ExtractParamsLocale = "bn-BD"
	ExtractParamsLocaleBnIn      ExtractParamsLocale = "bn-IN"
	ExtractParamsLocaleBo        ExtractParamsLocale = "bo"
	ExtractParamsLocaleBoCn      ExtractParamsLocale = "bo-CN"
	ExtractParamsLocaleBoIn      ExtractParamsLocale = "bo-IN"
	ExtractParamsLocaleBrFr      ExtractParamsLocale = "br-FR"
	ExtractParamsLocaleBrxIn     ExtractParamsLocale = "brx-IN"
	ExtractParamsLocaleBs        ExtractParamsLocale = "bs"
	ExtractParamsLocaleBsBa      ExtractParamsLocale = "bs-BA"
	ExtractParamsLocaleBynEr     ExtractParamsLocale = "byn-ER"
	ExtractParamsLocaleCa        ExtractParamsLocale = "ca"
	ExtractParamsLocaleCaAd      ExtractParamsLocale = "ca-AD"
	ExtractParamsLocaleCaEs      ExtractParamsLocale = "ca-ES"
	ExtractParamsLocaleCaFr      ExtractParamsLocale = "ca-FR"
	ExtractParamsLocaleCaIt      ExtractParamsLocale = "ca-IT"
	ExtractParamsLocaleCgg       ExtractParamsLocale = "cgg"
	ExtractParamsLocaleCggUg     ExtractParamsLocale = "cgg-UG"
	ExtractParamsLocaleChr       ExtractParamsLocale = "chr"
	ExtractParamsLocaleChrUs     ExtractParamsLocale = "chr-US"
	ExtractParamsLocaleCrhUa     ExtractParamsLocale = "crh-UA"
	ExtractParamsLocaleCs        ExtractParamsLocale = "cs"
	ExtractParamsLocaleCsCz      ExtractParamsLocale = "cs-CZ"
	ExtractParamsLocaleCsbPl     ExtractParamsLocale = "csb-PL"
	ExtractParamsLocaleCvRu      ExtractParamsLocale = "cv-RU"
	ExtractParamsLocaleCy        ExtractParamsLocale = "cy"
	ExtractParamsLocaleCyGB      ExtractParamsLocale = "cy-GB"
	ExtractParamsLocaleDa        ExtractParamsLocale = "da"
	ExtractParamsLocaleDaDk      ExtractParamsLocale = "da-DK"
	ExtractParamsLocaleDav       ExtractParamsLocale = "dav"
	ExtractParamsLocaleDavKe     ExtractParamsLocale = "dav-KE"
	ExtractParamsLocaleDe        ExtractParamsLocale = "de"
	ExtractParamsLocaleDeAt      ExtractParamsLocale = "de-AT"
	ExtractParamsLocaleDeBe      ExtractParamsLocale = "de-BE"
	ExtractParamsLocaleDeCh      ExtractParamsLocale = "de-CH"
	ExtractParamsLocaleDeDe      ExtractParamsLocale = "de-DE"
	ExtractParamsLocaleDeLi      ExtractParamsLocale = "de-LI"
	ExtractParamsLocaleDeLu      ExtractParamsLocale = "de-LU"
	ExtractParamsLocaleDvMv      ExtractParamsLocale = "dv-MV"
	ExtractParamsLocaleDzBt      ExtractParamsLocale = "dz-BT"
	ExtractParamsLocaleEbu       ExtractParamsLocale = "ebu"
	ExtractParamsLocaleEbuKe     ExtractParamsLocale = "ebu-KE"
	ExtractParamsLocaleEe        ExtractParamsLocale = "ee"
	ExtractParamsLocaleEeGh      ExtractParamsLocale = "ee-GH"
	ExtractParamsLocaleEeTg      ExtractParamsLocale = "ee-TG"
	ExtractParamsLocaleEl        ExtractParamsLocale = "el"
	ExtractParamsLocaleElCy      ExtractParamsLocale = "el-CY"
	ExtractParamsLocaleElGr      ExtractParamsLocale = "el-GR"
	ExtractParamsLocaleEn        ExtractParamsLocale = "en"
	ExtractParamsLocaleEnAg      ExtractParamsLocale = "en-AG"
	ExtractParamsLocaleEnAs      ExtractParamsLocale = "en-AS"
	ExtractParamsLocaleEnAu      ExtractParamsLocale = "en-AU"
	ExtractParamsLocaleEnBe      ExtractParamsLocale = "en-BE"
	ExtractParamsLocaleEnBw      ExtractParamsLocale = "en-BW"
	ExtractParamsLocaleEnBz      ExtractParamsLocale = "en-BZ"
	ExtractParamsLocaleEnCa      ExtractParamsLocale = "en-CA"
	ExtractParamsLocaleEnDk      ExtractParamsLocale = "en-DK"
	ExtractParamsLocaleEnGB      ExtractParamsLocale = "en-GB"
	ExtractParamsLocaleEnGu      ExtractParamsLocale = "en-GU"
	ExtractParamsLocaleEnHk      ExtractParamsLocale = "en-HK"
	ExtractParamsLocaleEnIe      ExtractParamsLocale = "en-IE"
	ExtractParamsLocaleEnIn      ExtractParamsLocale = "en-IN"
	ExtractParamsLocaleEnJm      ExtractParamsLocale = "en-JM"
	ExtractParamsLocaleEnMh      ExtractParamsLocale = "en-MH"
	ExtractParamsLocaleEnMp      ExtractParamsLocale = "en-MP"
	ExtractParamsLocaleEnMt      ExtractParamsLocale = "en-MT"
	ExtractParamsLocaleEnMu      ExtractParamsLocale = "en-MU"
	ExtractParamsLocaleEnNa      ExtractParamsLocale = "en-NA"
	ExtractParamsLocaleEnNg      ExtractParamsLocale = "en-NG"
	ExtractParamsLocaleEnNz      ExtractParamsLocale = "en-NZ"
	ExtractParamsLocaleEnPh      ExtractParamsLocale = "en-PH"
	ExtractParamsLocaleEnPk      ExtractParamsLocale = "en-PK"
	ExtractParamsLocaleEnSg      ExtractParamsLocale = "en-SG"
	ExtractParamsLocaleEnTt      ExtractParamsLocale = "en-TT"
	ExtractParamsLocaleEnUm      ExtractParamsLocale = "en-UM"
	ExtractParamsLocaleEnUs      ExtractParamsLocale = "en-US"
	ExtractParamsLocaleEnVi      ExtractParamsLocale = "en-VI"
	ExtractParamsLocaleEnZa      ExtractParamsLocale = "en-ZA"
	ExtractParamsLocaleEnZm      ExtractParamsLocale = "en-ZM"
	ExtractParamsLocaleEnZw      ExtractParamsLocale = "en-ZW"
	ExtractParamsLocaleEo        ExtractParamsLocale = "eo"
	ExtractParamsLocaleEs        ExtractParamsLocale = "es"
	ExtractParamsLocaleEs419     ExtractParamsLocale = "es-419"
	ExtractParamsLocaleEsAr      ExtractParamsLocale = "es-AR"
	ExtractParamsLocaleEsBo      ExtractParamsLocale = "es-BO"
	ExtractParamsLocaleEsCl      ExtractParamsLocale = "es-CL"
	ExtractParamsLocaleEsCo      ExtractParamsLocale = "es-CO"
	ExtractParamsLocaleEsCr      ExtractParamsLocale = "es-CR"
	ExtractParamsLocaleEsCu      ExtractParamsLocale = "es-CU"
	ExtractParamsLocaleEsDo      ExtractParamsLocale = "es-DO"
	ExtractParamsLocaleEsEc      ExtractParamsLocale = "es-EC"
	ExtractParamsLocaleEsEs      ExtractParamsLocale = "es-ES"
	ExtractParamsLocaleEsGq      ExtractParamsLocale = "es-GQ"
	ExtractParamsLocaleEsGt      ExtractParamsLocale = "es-GT"
	ExtractParamsLocaleEsHn      ExtractParamsLocale = "es-HN"
	ExtractParamsLocaleEsMx      ExtractParamsLocale = "es-MX"
	ExtractParamsLocaleEsNi      ExtractParamsLocale = "es-NI"
	ExtractParamsLocaleEsPa      ExtractParamsLocale = "es-PA"
	ExtractParamsLocaleEsPe      ExtractParamsLocale = "es-PE"
	ExtractParamsLocaleEsPr      ExtractParamsLocale = "es-PR"
	ExtractParamsLocaleEsPy      ExtractParamsLocale = "es-PY"
	ExtractParamsLocaleEsSv      ExtractParamsLocale = "es-SV"
	ExtractParamsLocaleEsUs      ExtractParamsLocale = "es-US"
	ExtractParamsLocaleEsUy      ExtractParamsLocale = "es-UY"
	ExtractParamsLocaleEsVe      ExtractParamsLocale = "es-VE"
	ExtractParamsLocaleEt        ExtractParamsLocale = "et"
	ExtractParamsLocaleEtEe      ExtractParamsLocale = "et-EE"
	ExtractParamsLocaleEu        ExtractParamsLocale = "eu"
	ExtractParamsLocaleEuEs      ExtractParamsLocale = "eu-ES"
	ExtractParamsLocaleFa        ExtractParamsLocale = "fa"
	ExtractParamsLocaleFaAf      ExtractParamsLocale = "fa-AF"
	ExtractParamsLocaleFaIr      ExtractParamsLocale = "fa-IR"
	ExtractParamsLocaleFf        ExtractParamsLocale = "ff"
	ExtractParamsLocaleFfSn      ExtractParamsLocale = "ff-SN"
	ExtractParamsLocaleFi        ExtractParamsLocale = "fi"
	ExtractParamsLocaleFiFi      ExtractParamsLocale = "fi-FI"
	ExtractParamsLocaleFil       ExtractParamsLocale = "fil"
	ExtractParamsLocaleFilPh     ExtractParamsLocale = "fil-PH"
	ExtractParamsLocaleFo        ExtractParamsLocale = "fo"
	ExtractParamsLocaleFoFo      ExtractParamsLocale = "fo-FO"
	ExtractParamsLocaleFr        ExtractParamsLocale = "fr"
	ExtractParamsLocaleFrBe      ExtractParamsLocale = "fr-BE"
	ExtractParamsLocaleFrBf      ExtractParamsLocale = "fr-BF"
	ExtractParamsLocaleFrBi      ExtractParamsLocale = "fr-BI"
	ExtractParamsLocaleFrBj      ExtractParamsLocale = "fr-BJ"
	ExtractParamsLocaleFrBl      ExtractParamsLocale = "fr-BL"
	ExtractParamsLocaleFrCa      ExtractParamsLocale = "fr-CA"
	ExtractParamsLocaleFrCd      ExtractParamsLocale = "fr-CD"
	ExtractParamsLocaleFrCf      ExtractParamsLocale = "fr-CF"
	ExtractParamsLocaleFrCg      ExtractParamsLocale = "fr-CG"
	ExtractParamsLocaleFrCh      ExtractParamsLocale = "fr-CH"
	ExtractParamsLocaleFrCi      ExtractParamsLocale = "fr-CI"
	ExtractParamsLocaleFrCm      ExtractParamsLocale = "fr-CM"
	ExtractParamsLocaleFrDj      ExtractParamsLocale = "fr-DJ"
	ExtractParamsLocaleFrFr      ExtractParamsLocale = "fr-FR"
	ExtractParamsLocaleFrGa      ExtractParamsLocale = "fr-GA"
	ExtractParamsLocaleFrGn      ExtractParamsLocale = "fr-GN"
	ExtractParamsLocaleFrGp      ExtractParamsLocale = "fr-GP"
	ExtractParamsLocaleFrGq      ExtractParamsLocale = "fr-GQ"
	ExtractParamsLocaleFrKm      ExtractParamsLocale = "fr-KM"
	ExtractParamsLocaleFrLu      ExtractParamsLocale = "fr-LU"
	ExtractParamsLocaleFrMc      ExtractParamsLocale = "fr-MC"
	ExtractParamsLocaleFrMf      ExtractParamsLocale = "fr-MF"
	ExtractParamsLocaleFrMg      ExtractParamsLocale = "fr-MG"
	ExtractParamsLocaleFrMl      ExtractParamsLocale = "fr-ML"
	ExtractParamsLocaleFrMq      ExtractParamsLocale = "fr-MQ"
	ExtractParamsLocaleFrNe      ExtractParamsLocale = "fr-NE"
	ExtractParamsLocaleFrRe      ExtractParamsLocale = "fr-RE"
	ExtractParamsLocaleFrRw      ExtractParamsLocale = "fr-RW"
	ExtractParamsLocaleFrSn      ExtractParamsLocale = "fr-SN"
	ExtractParamsLocaleFrTd      ExtractParamsLocale = "fr-TD"
	ExtractParamsLocaleFrTg      ExtractParamsLocale = "fr-TG"
	ExtractParamsLocaleFurIt     ExtractParamsLocale = "fur-IT"
	ExtractParamsLocaleFyDe      ExtractParamsLocale = "fy-DE"
	ExtractParamsLocaleFyNl      ExtractParamsLocale = "fy-NL"
	ExtractParamsLocaleGa        ExtractParamsLocale = "ga"
	ExtractParamsLocaleGaIe      ExtractParamsLocale = "ga-IE"
	ExtractParamsLocaleGdGB      ExtractParamsLocale = "gd-GB"
	ExtractParamsLocaleGezEr     ExtractParamsLocale = "gez-ER"
	ExtractParamsLocaleGezEt     ExtractParamsLocale = "gez-ET"
	ExtractParamsLocaleGl        ExtractParamsLocale = "gl"
	ExtractParamsLocaleGlEs      ExtractParamsLocale = "gl-ES"
	ExtractParamsLocaleGsw       ExtractParamsLocale = "gsw"
	ExtractParamsLocaleGswCh     ExtractParamsLocale = "gsw-CH"
	ExtractParamsLocaleGu        ExtractParamsLocale = "gu"
	ExtractParamsLocaleGuIn      ExtractParamsLocale = "gu-IN"
	ExtractParamsLocaleGuz       ExtractParamsLocale = "guz"
	ExtractParamsLocaleGuzKe     ExtractParamsLocale = "guz-KE"
	ExtractParamsLocaleGv        ExtractParamsLocale = "gv"
	ExtractParamsLocaleGvGB      ExtractParamsLocale = "gv-GB"
	ExtractParamsLocaleHa        ExtractParamsLocale = "ha"
	ExtractParamsLocaleHaLatn    ExtractParamsLocale = "ha-Latn"
	ExtractParamsLocaleHaLatnGh  ExtractParamsLocale = "ha-Latn-GH"
	ExtractParamsLocaleHaLatnNe  ExtractParamsLocale = "ha-Latn-NE"
	ExtractParamsLocaleHaLatnNg  ExtractParamsLocale = "ha-Latn-NG"
	ExtractParamsLocaleHaNg      ExtractParamsLocale = "ha-NG"
	ExtractParamsLocaleHaw       ExtractParamsLocale = "haw"
	ExtractParamsLocaleHawUs     ExtractParamsLocale = "haw-US"
	ExtractParamsLocaleHe        ExtractParamsLocale = "he"
	ExtractParamsLocaleHeIl      ExtractParamsLocale = "he-IL"
	ExtractParamsLocaleHi        ExtractParamsLocale = "hi"
	ExtractParamsLocaleHiIn      ExtractParamsLocale = "hi-IN"
	ExtractParamsLocaleHneIn     ExtractParamsLocale = "hne-IN"
	ExtractParamsLocaleHr        ExtractParamsLocale = "hr"
	ExtractParamsLocaleHrHr      ExtractParamsLocale = "hr-HR"
	ExtractParamsLocaleHsbDe     ExtractParamsLocale = "hsb-DE"
	ExtractParamsLocaleHtHt      ExtractParamsLocale = "ht-HT"
	ExtractParamsLocaleHu        ExtractParamsLocale = "hu"
	ExtractParamsLocaleHuHu      ExtractParamsLocale = "hu-HU"
	ExtractParamsLocaleHy        ExtractParamsLocale = "hy"
	ExtractParamsLocaleHyAm      ExtractParamsLocale = "hy-AM"
	ExtractParamsLocaleID        ExtractParamsLocale = "id"
	ExtractParamsLocaleIDID      ExtractParamsLocale = "id-ID"
	ExtractParamsLocaleIg        ExtractParamsLocale = "ig"
	ExtractParamsLocaleIgNg      ExtractParamsLocale = "ig-NG"
	ExtractParamsLocaleIi        ExtractParamsLocale = "ii"
	ExtractParamsLocaleIiCn      ExtractParamsLocale = "ii-CN"
	ExtractParamsLocaleIkCa      ExtractParamsLocale = "ik-CA"
	ExtractParamsLocaleIs        ExtractParamsLocale = "is"
	ExtractParamsLocaleIsIs      ExtractParamsLocale = "is-IS"
	ExtractParamsLocaleIt        ExtractParamsLocale = "it"
	ExtractParamsLocaleItCh      ExtractParamsLocale = "it-CH"
	ExtractParamsLocaleItIt      ExtractParamsLocale = "it-IT"
	ExtractParamsLocaleIuCa      ExtractParamsLocale = "iu-CA"
	ExtractParamsLocaleIwIl      ExtractParamsLocale = "iw-IL"
	ExtractParamsLocaleJa        ExtractParamsLocale = "ja"
	ExtractParamsLocaleJaJp      ExtractParamsLocale = "ja-JP"
	ExtractParamsLocaleJmc       ExtractParamsLocale = "jmc"
	ExtractParamsLocaleJmcTz     ExtractParamsLocale = "jmc-TZ"
	ExtractParamsLocaleKa        ExtractParamsLocale = "ka"
	ExtractParamsLocaleKaGe      ExtractParamsLocale = "ka-GE"
	ExtractParamsLocaleKab       ExtractParamsLocale = "kab"
	ExtractParamsLocaleKabDz     ExtractParamsLocale = "kab-DZ"
	ExtractParamsLocaleKam       ExtractParamsLocale = "kam"
	ExtractParamsLocaleKamKe     ExtractParamsLocale = "kam-KE"
	ExtractParamsLocaleKde       ExtractParamsLocale = "kde"
	ExtractParamsLocaleKdeTz     ExtractParamsLocale = "kde-TZ"
	ExtractParamsLocaleKea       ExtractParamsLocale = "kea"
	ExtractParamsLocaleKeaCv     ExtractParamsLocale = "kea-CV"
	ExtractParamsLocaleKhq       ExtractParamsLocale = "khq"
	ExtractParamsLocaleKhqMl     ExtractParamsLocale = "khq-ML"
	ExtractParamsLocaleKi        ExtractParamsLocale = "ki"
	ExtractParamsLocaleKiKe      ExtractParamsLocale = "ki-KE"
	ExtractParamsLocaleKk        ExtractParamsLocale = "kk"
	ExtractParamsLocaleKkCyrl    ExtractParamsLocale = "kk-Cyrl"
	ExtractParamsLocaleKkCyrlKz  ExtractParamsLocale = "kk-Cyrl-KZ"
	ExtractParamsLocaleKkKz      ExtractParamsLocale = "kk-KZ"
	ExtractParamsLocaleKl        ExtractParamsLocale = "kl"
	ExtractParamsLocaleKlGl      ExtractParamsLocale = "kl-GL"
	ExtractParamsLocaleKln       ExtractParamsLocale = "kln"
	ExtractParamsLocaleKlnKe     ExtractParamsLocale = "kln-KE"
	ExtractParamsLocaleKm        ExtractParamsLocale = "km"
	ExtractParamsLocaleKmKh      ExtractParamsLocale = "km-KH"
	ExtractParamsLocaleKn        ExtractParamsLocale = "kn"
	ExtractParamsLocaleKnIn      ExtractParamsLocale = "kn-IN"
	ExtractParamsLocaleKo        ExtractParamsLocale = "ko"
	ExtractParamsLocaleKoKr      ExtractParamsLocale = "ko-KR"
	ExtractParamsLocaleKok       ExtractParamsLocale = "kok"
	ExtractParamsLocaleKokIn     ExtractParamsLocale = "kok-IN"
	ExtractParamsLocaleKsIn      ExtractParamsLocale = "ks-IN"
	ExtractParamsLocaleKuTr      ExtractParamsLocale = "ku-TR"
	ExtractParamsLocaleKw        ExtractParamsLocale = "kw"
	ExtractParamsLocaleKwGB      ExtractParamsLocale = "kw-GB"
	ExtractParamsLocaleKyKg      ExtractParamsLocale = "ky-KG"
	ExtractParamsLocaleLag       ExtractParamsLocale = "lag"
	ExtractParamsLocaleLagTz     ExtractParamsLocale = "lag-TZ"
	ExtractParamsLocaleLbLu      ExtractParamsLocale = "lb-LU"
	ExtractParamsLocaleLg        ExtractParamsLocale = "lg"
	ExtractParamsLocaleLgUg      ExtractParamsLocale = "lg-UG"
	ExtractParamsLocaleLiBe      ExtractParamsLocale = "li-BE"
	ExtractParamsLocaleLiNl      ExtractParamsLocale = "li-NL"
	ExtractParamsLocaleLijIt     ExtractParamsLocale = "lij-IT"
	ExtractParamsLocaleLoLa      ExtractParamsLocale = "lo-LA"
	ExtractParamsLocaleLt        ExtractParamsLocale = "lt"
	ExtractParamsLocaleLtLt      ExtractParamsLocale = "lt-LT"
	ExtractParamsLocaleLuo       ExtractParamsLocale = "luo"
	ExtractParamsLocaleLuoKe     ExtractParamsLocale = "luo-KE"
	ExtractParamsLocaleLuy       ExtractParamsLocale = "luy"
	ExtractParamsLocaleLuyKe     ExtractParamsLocale = "luy-KE"
	ExtractParamsLocaleLv        ExtractParamsLocale = "lv"
	ExtractParamsLocaleLvLv      ExtractParamsLocale = "lv-LV"
	ExtractParamsLocaleMagIn     ExtractParamsLocale = "mag-IN"
	ExtractParamsLocaleMaiIn     ExtractParamsLocale = "mai-IN"
	ExtractParamsLocaleMas       ExtractParamsLocale = "mas"
	ExtractParamsLocaleMasKe     ExtractParamsLocale = "mas-KE"
	ExtractParamsLocaleMasTz     ExtractParamsLocale = "mas-TZ"
	ExtractParamsLocaleMer       ExtractParamsLocale = "mer"
	ExtractParamsLocaleMerKe     ExtractParamsLocale = "mer-KE"
	ExtractParamsLocaleMfe       ExtractParamsLocale = "mfe"
	ExtractParamsLocaleMfeMu     ExtractParamsLocale = "mfe-MU"
	ExtractParamsLocaleMg        ExtractParamsLocale = "mg"
	ExtractParamsLocaleMgMg      ExtractParamsLocale = "mg-MG"
	ExtractParamsLocaleMhrRu     ExtractParamsLocale = "mhr-RU"
	ExtractParamsLocaleMiNz      ExtractParamsLocale = "mi-NZ"
	ExtractParamsLocaleMk        ExtractParamsLocale = "mk"
	ExtractParamsLocaleMkMk      ExtractParamsLocale = "mk-MK"
	ExtractParamsLocaleMl        ExtractParamsLocale = "ml"
	ExtractParamsLocaleMlIn      ExtractParamsLocale = "ml-IN"
	ExtractParamsLocaleMnMn      ExtractParamsLocale = "mn-MN"
	ExtractParamsLocaleMr        ExtractParamsLocale = "mr"
	ExtractParamsLocaleMrIn      ExtractParamsLocale = "mr-IN"
	ExtractParamsLocaleMs        ExtractParamsLocale = "ms"
	ExtractParamsLocaleMsBn      ExtractParamsLocale = "ms-BN"
	ExtractParamsLocaleMsMy      ExtractParamsLocale = "ms-MY"
	ExtractParamsLocaleMt        ExtractParamsLocale = "mt"
	ExtractParamsLocaleMtMt      ExtractParamsLocale = "mt-MT"
	ExtractParamsLocaleMy        ExtractParamsLocale = "my"
	ExtractParamsLocaleMyMm      ExtractParamsLocale = "my-MM"
	ExtractParamsLocaleNanTw     ExtractParamsLocale = "nan-TW"
	ExtractParamsLocaleNaq       ExtractParamsLocale = "naq"
	ExtractParamsLocaleNaqNa     ExtractParamsLocale = "naq-NA"
	ExtractParamsLocaleNb        ExtractParamsLocale = "nb"
	ExtractParamsLocaleNbNo      ExtractParamsLocale = "nb-NO"
	ExtractParamsLocaleNd        ExtractParamsLocale = "nd"
	ExtractParamsLocaleNdZw      ExtractParamsLocale = "nd-ZW"
	ExtractParamsLocaleNdsDe     ExtractParamsLocale = "nds-DE"
	ExtractParamsLocaleNdsNl     ExtractParamsLocale = "nds-NL"
	ExtractParamsLocaleNe        ExtractParamsLocale = "ne"
	ExtractParamsLocaleNeIn      ExtractParamsLocale = "ne-IN"
	ExtractParamsLocaleNeNp      ExtractParamsLocale = "ne-NP"
	ExtractParamsLocaleNl        ExtractParamsLocale = "nl"
	ExtractParamsLocaleNlAw      ExtractParamsLocale = "nl-AW"
	ExtractParamsLocaleNlBe      ExtractParamsLocale = "nl-BE"
	ExtractParamsLocaleNlNl      ExtractParamsLocale = "nl-NL"
	ExtractParamsLocaleNn        ExtractParamsLocale = "nn"
	ExtractParamsLocaleNnNo      ExtractParamsLocale = "nn-NO"
	ExtractParamsLocaleNrZa      ExtractParamsLocale = "nr-ZA"
	ExtractParamsLocaleNsoZa     ExtractParamsLocale = "nso-ZA"
	ExtractParamsLocaleNyn       ExtractParamsLocale = "nyn"
	ExtractParamsLocaleNynUg     ExtractParamsLocale = "nyn-UG"
	ExtractParamsLocaleOcFr      ExtractParamsLocale = "oc-FR"
	ExtractParamsLocaleOm        ExtractParamsLocale = "om"
	ExtractParamsLocaleOmEt      ExtractParamsLocale = "om-ET"
	ExtractParamsLocaleOmKe      ExtractParamsLocale = "om-KE"
	ExtractParamsLocaleOr        ExtractParamsLocale = "or"
	ExtractParamsLocaleOrIn      ExtractParamsLocale = "or-IN"
	ExtractParamsLocaleOsRu      ExtractParamsLocale = "os-RU"
	ExtractParamsLocalePa        ExtractParamsLocale = "pa"
	ExtractParamsLocalePaArab    ExtractParamsLocale = "pa-Arab"
	ExtractParamsLocalePaArabPk  ExtractParamsLocale = "pa-Arab-PK"
	ExtractParamsLocalePaGuru    ExtractParamsLocale = "pa-Guru"
	ExtractParamsLocalePaGuruIn  ExtractParamsLocale = "pa-Guru-IN"
	ExtractParamsLocalePaIn      ExtractParamsLocale = "pa-IN"
	ExtractParamsLocalePaPk      ExtractParamsLocale = "pa-PK"
	ExtractParamsLocalePapAn     ExtractParamsLocale = "pap-AN"
	ExtractParamsLocalePl        ExtractParamsLocale = "pl"
	ExtractParamsLocalePlPl      ExtractParamsLocale = "pl-PL"
	ExtractParamsLocalePs        ExtractParamsLocale = "ps"
	ExtractParamsLocalePsAf      ExtractParamsLocale = "ps-AF"
	ExtractParamsLocalePt        ExtractParamsLocale = "pt"
	ExtractParamsLocalePtBr      ExtractParamsLocale = "pt-BR"
	ExtractParamsLocalePtGw      ExtractParamsLocale = "pt-GW"
	ExtractParamsLocalePtMz      ExtractParamsLocale = "pt-MZ"
	ExtractParamsLocalePtPt      ExtractParamsLocale = "pt-PT"
	ExtractParamsLocaleRm        ExtractParamsLocale = "rm"
	ExtractParamsLocaleRmCh      ExtractParamsLocale = "rm-CH"
	ExtractParamsLocaleRo        ExtractParamsLocale = "ro"
	ExtractParamsLocaleRoMd      ExtractParamsLocale = "ro-MD"
	ExtractParamsLocaleRoRo      ExtractParamsLocale = "ro-RO"
	ExtractParamsLocaleRof       ExtractParamsLocale = "rof"
	ExtractParamsLocaleRofTz     ExtractParamsLocale = "rof-TZ"
	ExtractParamsLocaleRu        ExtractParamsLocale = "ru"
	ExtractParamsLocaleRuMd      ExtractParamsLocale = "ru-MD"
	ExtractParamsLocaleRuRu      ExtractParamsLocale = "ru-RU"
	ExtractParamsLocaleRuUa      ExtractParamsLocale = "ru-UA"
	ExtractParamsLocaleRw        ExtractParamsLocale = "rw"
	ExtractParamsLocaleRwRw      ExtractParamsLocale = "rw-RW"
	ExtractParamsLocaleRwk       ExtractParamsLocale = "rwk"
	ExtractParamsLocaleRwkTz     ExtractParamsLocale = "rwk-TZ"
	ExtractParamsLocaleSaIn      ExtractParamsLocale = "sa-IN"
	ExtractParamsLocaleSaq       ExtractParamsLocale = "saq"
	ExtractParamsLocaleSaqKe     ExtractParamsLocale = "saq-KE"
	ExtractParamsLocaleScIt      ExtractParamsLocale = "sc-IT"
	ExtractParamsLocaleSdIn      ExtractParamsLocale = "sd-IN"
	ExtractParamsLocaleSeNo      ExtractParamsLocale = "se-NO"
	ExtractParamsLocaleSeh       ExtractParamsLocale = "seh"
	ExtractParamsLocaleSehMz     ExtractParamsLocale = "seh-MZ"
	ExtractParamsLocaleSes       ExtractParamsLocale = "ses"
	ExtractParamsLocaleSesMl     ExtractParamsLocale = "ses-ML"
	ExtractParamsLocaleSg        ExtractParamsLocale = "sg"
	ExtractParamsLocaleSgCf      ExtractParamsLocale = "sg-CF"
	ExtractParamsLocaleShi       ExtractParamsLocale = "shi"
	ExtractParamsLocaleShiLatn   ExtractParamsLocale = "shi-Latn"
	ExtractParamsLocaleShiLatnMa ExtractParamsLocale = "shi-Latn-MA"
	ExtractParamsLocaleShiTfng   ExtractParamsLocale = "shi-Tfng"
	ExtractParamsLocaleShiTfngMa ExtractParamsLocale = "shi-Tfng-MA"
	ExtractParamsLocaleShsCa     ExtractParamsLocale = "shs-CA"
	ExtractParamsLocaleSi        ExtractParamsLocale = "si"
	ExtractParamsLocaleSiLk      ExtractParamsLocale = "si-LK"
	ExtractParamsLocaleSidEt     ExtractParamsLocale = "sid-ET"
	ExtractParamsLocaleSk        ExtractParamsLocale = "sk"
	ExtractParamsLocaleSkSk      ExtractParamsLocale = "sk-SK"
	ExtractParamsLocaleSl        ExtractParamsLocale = "sl"
	ExtractParamsLocaleSlSi      ExtractParamsLocale = "sl-SI"
	ExtractParamsLocaleSn        ExtractParamsLocale = "sn"
	ExtractParamsLocaleSnZw      ExtractParamsLocale = "sn-ZW"
	ExtractParamsLocaleSo        ExtractParamsLocale = "so"
	ExtractParamsLocaleSoDj      ExtractParamsLocale = "so-DJ"
	ExtractParamsLocaleSoEt      ExtractParamsLocale = "so-ET"
	ExtractParamsLocaleSoKe      ExtractParamsLocale = "so-KE"
	ExtractParamsLocaleSoSo      ExtractParamsLocale = "so-SO"
	ExtractParamsLocaleSq        ExtractParamsLocale = "sq"
	ExtractParamsLocaleSqAl      ExtractParamsLocale = "sq-AL"
	ExtractParamsLocaleSqMk      ExtractParamsLocale = "sq-MK"
	ExtractParamsLocaleSr        ExtractParamsLocale = "sr"
	ExtractParamsLocaleSrCyrl    ExtractParamsLocale = "sr-Cyrl"
	ExtractParamsLocaleSrCyrlBa  ExtractParamsLocale = "sr-Cyrl-BA"
	ExtractParamsLocaleSrCyrlMe  ExtractParamsLocale = "sr-Cyrl-ME"
	ExtractParamsLocaleSrCyrlRs  ExtractParamsLocale = "sr-Cyrl-RS"
	ExtractParamsLocaleSrLatn    ExtractParamsLocale = "sr-Latn"
	ExtractParamsLocaleSrLatnBa  ExtractParamsLocale = "sr-Latn-BA"
	ExtractParamsLocaleSrLatnMe  ExtractParamsLocale = "sr-Latn-ME"
	ExtractParamsLocaleSrLatnRs  ExtractParamsLocale = "sr-Latn-RS"
	ExtractParamsLocaleSrMe      ExtractParamsLocale = "sr-ME"
	ExtractParamsLocaleSrRs      ExtractParamsLocale = "sr-RS"
	ExtractParamsLocaleSSZa      ExtractParamsLocale = "ss-ZA"
	ExtractParamsLocaleStZa      ExtractParamsLocale = "st-ZA"
	ExtractParamsLocaleSv        ExtractParamsLocale = "sv"
	ExtractParamsLocaleSvFi      ExtractParamsLocale = "sv-FI"
	ExtractParamsLocaleSvSe      ExtractParamsLocale = "sv-SE"
	ExtractParamsLocaleSw        ExtractParamsLocale = "sw"
	ExtractParamsLocaleSwKe      ExtractParamsLocale = "sw-KE"
	ExtractParamsLocaleSwTz      ExtractParamsLocale = "sw-TZ"
	ExtractParamsLocaleTa        ExtractParamsLocale = "ta"
	ExtractParamsLocaleTaIn      ExtractParamsLocale = "ta-IN"
	ExtractParamsLocaleTaLk      ExtractParamsLocale = "ta-LK"
	ExtractParamsLocaleTe        ExtractParamsLocale = "te"
	ExtractParamsLocaleTeIn      ExtractParamsLocale = "te-IN"
	ExtractParamsLocaleTeo       ExtractParamsLocale = "teo"
	ExtractParamsLocaleTeoKe     ExtractParamsLocale = "teo-KE"
	ExtractParamsLocaleTeoUg     ExtractParamsLocale = "teo-UG"
	ExtractParamsLocaleTgTj      ExtractParamsLocale = "tg-TJ"
	ExtractParamsLocaleTh        ExtractParamsLocale = "th"
	ExtractParamsLocaleThTh      ExtractParamsLocale = "th-TH"
	ExtractParamsLocaleTi        ExtractParamsLocale = "ti"
	ExtractParamsLocaleTiEr      ExtractParamsLocale = "ti-ER"
	ExtractParamsLocaleTiEt      ExtractParamsLocale = "ti-ET"
	ExtractParamsLocaleTigEr     ExtractParamsLocale = "tig-ER"
	ExtractParamsLocaleTkTm      ExtractParamsLocale = "tk-TM"
	ExtractParamsLocaleTlPh      ExtractParamsLocale = "tl-PH"
	ExtractParamsLocaleTnZa      ExtractParamsLocale = "tn-ZA"
	ExtractParamsLocaleTo        ExtractParamsLocale = "to"
	ExtractParamsLocaleToTo      ExtractParamsLocale = "to-TO"
	ExtractParamsLocaleTr        ExtractParamsLocale = "tr"
	ExtractParamsLocaleTrCy      ExtractParamsLocale = "tr-CY"
	ExtractParamsLocaleTrTr      ExtractParamsLocale = "tr-TR"
	ExtractParamsLocaleTsZa      ExtractParamsLocale = "ts-ZA"
	ExtractParamsLocaleTtRu      ExtractParamsLocale = "tt-RU"
	ExtractParamsLocaleTzm       ExtractParamsLocale = "tzm"
	ExtractParamsLocaleTzmLatn   ExtractParamsLocale = "tzm-Latn"
	ExtractParamsLocaleTzmLatnMa ExtractParamsLocale = "tzm-Latn-MA"
	ExtractParamsLocaleUgCn      ExtractParamsLocale = "ug-CN"
	ExtractParamsLocaleUk        ExtractParamsLocale = "uk"
	ExtractParamsLocaleUkUa      ExtractParamsLocale = "uk-UA"
	ExtractParamsLocaleUnmUs     ExtractParamsLocale = "unm-US"
	ExtractParamsLocaleUr        ExtractParamsLocale = "ur"
	ExtractParamsLocaleUrIn      ExtractParamsLocale = "ur-IN"
	ExtractParamsLocaleUrPk      ExtractParamsLocale = "ur-PK"
	ExtractParamsLocaleUz        ExtractParamsLocale = "uz"
	ExtractParamsLocaleUzArab    ExtractParamsLocale = "uz-Arab"
	ExtractParamsLocaleUzArabAf  ExtractParamsLocale = "uz-Arab-AF"
	ExtractParamsLocaleUzCyrl    ExtractParamsLocale = "uz-Cyrl"
	ExtractParamsLocaleUzCyrlUz  ExtractParamsLocale = "uz-Cyrl-UZ"
	ExtractParamsLocaleUzLatn    ExtractParamsLocale = "uz-Latn"
	ExtractParamsLocaleUzLatnUz  ExtractParamsLocale = "uz-Latn-UZ"
	ExtractParamsLocaleUzUz      ExtractParamsLocale = "uz-UZ"
	ExtractParamsLocaleVeZa      ExtractParamsLocale = "ve-ZA"
	ExtractParamsLocaleVi        ExtractParamsLocale = "vi"
	ExtractParamsLocaleViVn      ExtractParamsLocale = "vi-VN"
	ExtractParamsLocaleVun       ExtractParamsLocale = "vun"
	ExtractParamsLocaleVunTz     ExtractParamsLocale = "vun-TZ"
	ExtractParamsLocaleWaBe      ExtractParamsLocale = "wa-BE"
	ExtractParamsLocaleWaeCh     ExtractParamsLocale = "wae-CH"
	ExtractParamsLocaleWalEt     ExtractParamsLocale = "wal-ET"
	ExtractParamsLocaleWoSn      ExtractParamsLocale = "wo-SN"
	ExtractParamsLocaleXhZa      ExtractParamsLocale = "xh-ZA"
	ExtractParamsLocaleXog       ExtractParamsLocale = "xog"
	ExtractParamsLocaleXogUg     ExtractParamsLocale = "xog-UG"
	ExtractParamsLocaleYiUs      ExtractParamsLocale = "yi-US"
	ExtractParamsLocaleYo        ExtractParamsLocale = "yo"
	ExtractParamsLocaleYoNg      ExtractParamsLocale = "yo-NG"
	ExtractParamsLocaleYueHk     ExtractParamsLocale = "yue-HK"
	ExtractParamsLocaleZh        ExtractParamsLocale = "zh"
	ExtractParamsLocaleZhCn      ExtractParamsLocale = "zh-CN"
	ExtractParamsLocaleZhHk      ExtractParamsLocale = "zh-HK"
	ExtractParamsLocaleZhHans    ExtractParamsLocale = "zh-Hans"
	ExtractParamsLocaleZhHansCn  ExtractParamsLocale = "zh-Hans-CN"
	ExtractParamsLocaleZhHansHk  ExtractParamsLocale = "zh-Hans-HK"
	ExtractParamsLocaleZhHansMo  ExtractParamsLocale = "zh-Hans-MO"
	ExtractParamsLocaleZhHansSg  ExtractParamsLocale = "zh-Hans-SG"
	ExtractParamsLocaleZhHant    ExtractParamsLocale = "zh-Hant"
	ExtractParamsLocaleZhHantHk  ExtractParamsLocale = "zh-Hant-HK"
	ExtractParamsLocaleZhHantMo  ExtractParamsLocale = "zh-Hant-MO"
	ExtractParamsLocaleZhHantTw  ExtractParamsLocale = "zh-Hant-TW"
	ExtractParamsLocaleZhSg      ExtractParamsLocale = "zh-SG"
	ExtractParamsLocaleZhTw      ExtractParamsLocale = "zh-TW"
	ExtractParamsLocaleZu        ExtractParamsLocale = "zu"
	ExtractParamsLocaleZuZa      ExtractParamsLocale = "zu-ZA"
	ExtractParamsLocaleAuto      ExtractParamsLocale = "auto"
)

// HTTP method for the request
type ExtractParamsMethod string

const (
	ExtractParamsMethodGet    ExtractParamsMethod = "GET"
	ExtractParamsMethodPost   ExtractParamsMethod = "POST"
	ExtractParamsMethodPut    ExtractParamsMethod = "PUT"
	ExtractParamsMethodPatch  ExtractParamsMethod = "PATCH"
	ExtractParamsMethodDelete ExtractParamsMethod = "DELETE"
)

type ExtractParamsNetworkCapture struct {
	Validation                  param.Opt[bool]    `json:"validation,omitzero"`
	WaitForRequestsCount        param.Opt[float64] `json:"wait_for_requests_count,omitzero"`
	WaitForRequestsCountTimeout param.Opt[float64] `json:"wait_for_requests_count_timeout,omitzero"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method,omitzero"`
	// Resource type for network capture filtering
	ResourceType ExtractParamsNetworkCaptureResourceTypeUnion `json:"resource_type,omitzero"`
	StatusCode   ExtractParamsNetworkCaptureStatusCodeUnion   `json:"status_code,omitzero"`
	URL          ExtractParamsNetworkCaptureURL               `json:"url,omitzero"`
	paramObj
}

func (r ExtractParamsNetworkCapture) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsNetworkCapture
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractParamsNetworkCapture](
		"method", "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsNetworkCaptureResourceTypeUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsNetworkCaptureResourceTypeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractParamsNetworkCaptureResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsNetworkCaptureResourceTypeUnion) asAny() any {
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
type ExtractParamsNetworkCaptureStatusCodeUnion struct {
	OfFloat      param.Opt[float64] `json:",omitzero,inline"`
	OfFloatArray []float64          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsNetworkCaptureStatusCodeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfFloatArray)
}
func (u *ExtractParamsNetworkCaptureStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsNetworkCaptureStatusCodeUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfFloatArray) {
		return &u.OfFloatArray
	}
	return nil
}

// The property Value is required.
type ExtractParamsNetworkCaptureURL struct {
	Value string `json:"value" api:"required"`
	// Any of "exact", "contains".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ExtractParamsNetworkCaptureURL) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsNetworkCaptureURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsNetworkCaptureURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractParamsNetworkCaptureURL](
		"type", "exact", "contains",
	)
}

// Operating system to emulate
type ExtractParamsOs string

const (
	ExtractParamsOsWindows ExtractParamsOs = "windows"
	ExtractParamsOsMacOs   ExtractParamsOs = "mac os"
	ExtractParamsOsLinux   ExtractParamsOs = "linux"
	ExtractParamsOsAndroid ExtractParamsOs = "android"
	ExtractParamsOsIos     ExtractParamsOs = "ios"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsParserUnion struct {
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsParserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap, u.OfString)
}
func (u *ExtractParamsParserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsParserUnion) asAny() any {
	if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Referrer policy for the request
type ExtractParamsReferrerType string

const (
	ExtractParamsReferrerTypeRandom     ExtractParamsReferrerType = "random"
	ExtractParamsReferrerTypeNoReferer  ExtractParamsReferrerType = "no-referer"
	ExtractParamsReferrerTypeSameOrigin ExtractParamsReferrerType = "same-origin"
	ExtractParamsReferrerTypeGoogle     ExtractParamsReferrerType = "google"
	ExtractParamsReferrerTypeBing       ExtractParamsReferrerType = "bing"
	ExtractParamsReferrerTypeFacebook   ExtractParamsReferrerType = "facebook"
	ExtractParamsReferrerTypeTwitter    ExtractParamsReferrerType = "twitter"
	ExtractParamsReferrerTypeInstagram  ExtractParamsReferrerType = "instagram"
)

type ExtractParamsSession struct {
	ID                  param.Opt[string]  `json:"id,omitzero"`
	PrefetchUserbrowser param.Opt[bool]    `json:"prefetch_userbrowser,omitzero"`
	Retry               param.Opt[bool]    `json:"retry,omitzero"`
	Timeout             param.Opt[float64] `json:"timeout,omitzero"`
	paramObj
}

func (r ExtractParamsSession) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsSession
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsSkillUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsSkillUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractParamsSkillUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsSkillUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// US state for geolocation (only valid when country is US)
type ExtractParamsState string

const (
	ExtractParamsStateAl ExtractParamsState = "AL"
	ExtractParamsStateAk ExtractParamsState = "AK"
	ExtractParamsStateAs ExtractParamsState = "AS"
	ExtractParamsStateAz ExtractParamsState = "AZ"
	ExtractParamsStateAr ExtractParamsState = "AR"
	ExtractParamsStateCa ExtractParamsState = "CA"
	ExtractParamsStateCo ExtractParamsState = "CO"
	ExtractParamsStateCt ExtractParamsState = "CT"
	ExtractParamsStateDe ExtractParamsState = "DE"
	ExtractParamsStateDc ExtractParamsState = "DC"
	ExtractParamsStateFl ExtractParamsState = "FL"
	ExtractParamsStateGa ExtractParamsState = "GA"
	ExtractParamsStateGu ExtractParamsState = "GU"
	ExtractParamsStateHi ExtractParamsState = "HI"
	ExtractParamsStateID ExtractParamsState = "ID"
	ExtractParamsStateIl ExtractParamsState = "IL"
	ExtractParamsStateIn ExtractParamsState = "IN"
	ExtractParamsStateIa ExtractParamsState = "IA"
	ExtractParamsStateKs ExtractParamsState = "KS"
	ExtractParamsStateKy ExtractParamsState = "KY"
	ExtractParamsStateLa ExtractParamsState = "LA"
	ExtractParamsStateMe ExtractParamsState = "ME"
	ExtractParamsStateMd ExtractParamsState = "MD"
	ExtractParamsStateMa ExtractParamsState = "MA"
	ExtractParamsStateMi ExtractParamsState = "MI"
	ExtractParamsStateMn ExtractParamsState = "MN"
	ExtractParamsStateMs ExtractParamsState = "MS"
	ExtractParamsStateMo ExtractParamsState = "MO"
	ExtractParamsStateMt ExtractParamsState = "MT"
	ExtractParamsStateNe ExtractParamsState = "NE"
	ExtractParamsStateNv ExtractParamsState = "NV"
	ExtractParamsStateNh ExtractParamsState = "NH"
	ExtractParamsStateNj ExtractParamsState = "NJ"
	ExtractParamsStateNm ExtractParamsState = "NM"
	ExtractParamsStateNy ExtractParamsState = "NY"
	ExtractParamsStateNc ExtractParamsState = "NC"
	ExtractParamsStateNd ExtractParamsState = "ND"
	ExtractParamsStateMp ExtractParamsState = "MP"
	ExtractParamsStateOh ExtractParamsState = "OH"
	ExtractParamsStateOk ExtractParamsState = "OK"
	ExtractParamsStateOr ExtractParamsState = "OR"
	ExtractParamsStatePa ExtractParamsState = "PA"
	ExtractParamsStatePr ExtractParamsState = "PR"
	ExtractParamsStateRi ExtractParamsState = "RI"
	ExtractParamsStateSc ExtractParamsState = "SC"
	ExtractParamsStateSd ExtractParamsState = "SD"
	ExtractParamsStateTn ExtractParamsState = "TN"
	ExtractParamsStateTx ExtractParamsState = "TX"
	ExtractParamsStateUt ExtractParamsState = "UT"
	ExtractParamsStateVt ExtractParamsState = "VT"
	ExtractParamsStateVa ExtractParamsState = "VA"
	ExtractParamsStateVi ExtractParamsState = "VI"
	ExtractParamsStateWa ExtractParamsState = "WA"
	ExtractParamsStateWv ExtractParamsState = "WV"
	ExtractParamsStateWi ExtractParamsState = "WI"
	ExtractParamsStateWy ExtractParamsState = "WY"
)

type ExtractAsyncParams struct {
	// Target URL to scrape
	URL string `json:"url" api:"required"`
	// URL to call back when async operation completes
	CallbackURL param.Opt[string] `json:"callback_url,omitzero"`
	// City for geolocation
	City param.Opt[string] `json:"city,omitzero"`
	// Whether to automatically handle cookie consent headers
	ConsentHeader param.Opt[bool] `json:"consent_header,omitzero"`
	// Whether to use HTTP/2 protocol
	Http2 param.Opt[bool] `json:"http2,omitzero"`
	// Whether to emulate XMLHttpRequest behavior
	IsXhr param.Opt[bool] `json:"is_xhr,omitzero"`
	// Whether to parse the response content
	Parse param.Opt[bool] `json:"parse,omitzero"`
	// Whether to render JavaScript content using a browser
	Render param.Opt[bool] `json:"render,omitzero"`
	// Request timeout in milliseconds
	RequestTimeout param.Opt[float64] `json:"request_timeout,omitzero"`
	// Whether to compress stored data
	StorageCompress param.Opt[bool] `json:"storage_compress,omitzero"`
	// Custom name for the stored object
	StorageObjectName param.Opt[string] `json:"storage_object_name,omitzero"`
	// Type of storage to use for results
	StorageType param.Opt[string] `json:"storage_type,omitzero"`
	// URL for storage location
	StorageURL param.Opt[string] `json:"storage_url,omitzero"`
	// User-defined tag for request identification
	Tag param.Opt[string] `json:"tag,omitzero"`
	// Browser type to emulate
	Browser ExtractAsyncParamsBrowserUnion `json:"browser,omitzero"`
	// Array of browser automation actions to execute sequentially
	BrowserActions []ExtractAsyncParamsBrowserActionUnion `json:"browser_actions,omitzero"`
	// Browser cookies as array of cookie objects
	Cookies ExtractAsyncParamsCookiesUnion `json:"cookies,omitzero"`
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
	Country ExtractAsyncParamsCountry `json:"country,omitzero"`
	// Device type for browser emulation
	//
	// Any of "desktop", "mobile", "tablet".
	Device ExtractAsyncParamsDevice `json:"device,omitzero"`
	// Browser driver to use
	//
	// Any of "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro".
	Driver ExtractAsyncParamsDriver `json:"driver,omitzero"`
	// Expected HTTP status codes for successful requests
	ExpectedStatusCodes []int64 `json:"expected_status_codes,omitzero"`
	// List of acceptable response formats in order of preference
	//
	// Any of "html", "markdown", "screenshot", "headers".
	Formats []string `json:"formats,omitzero"`
	// Custom HTTP headers to include in the request
	Headers map[string]ExtractAsyncParamsHeaderUnion `json:"headers,omitzero"`
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
	Locale ExtractAsyncParamsLocale `json:"locale,omitzero"`
	// HTTP method for the request
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method ExtractAsyncParamsMethod `json:"method,omitzero"`
	// Filters for capturing network traffic
	NetworkCapture []ExtractAsyncParamsNetworkCapture `json:"network_capture,omitzero"`
	// Operating system to emulate
	//
	// Any of "windows", "mac os", "linux", "android", "ios".
	Os ExtractAsyncParamsOs `json:"os,omitzero"`
	// Custom parser configuration as a key-value map
	Parser ExtractAsyncParamsParserUnion `json:"parser,omitzero"`
	// Referrer policy for the request
	//
	// Any of "random", "no-referer", "same-origin", "google", "bing", "facebook",
	// "twitter", "instagram".
	ReferrerType ExtractAsyncParamsReferrerType `json:"referrer_type,omitzero"`
	Session      ExtractAsyncParamsSession      `json:"session,omitzero"`
	// Skills or capabilities required for the request
	Skill ExtractAsyncParamsSkillUnion `json:"skill,omitzero"`
	// US state for geolocation (only valid when country is US)
	//
	// Any of "AL", "AK", "AS", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA",
	// "GU", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI",
	// "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "MP",
	// "OH", "OK", "OR", "PA", "PR", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA",
	// "VI", "WA", "WV", "WI", "WY".
	State ExtractAsyncParamsState `json:"state,omitzero"`
	paramObj
}

func (r ExtractAsyncParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserString)
	OfExtractAsyncsBrowserString param.Opt[string]                `json:",omitzero,inline"`
	OfExtractAsyncsBrowserObject *ExtractAsyncParamsBrowserObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserString, u.OfExtractAsyncsBrowserObject)
}
func (u *ExtractAsyncParamsBrowserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserString) {
		return &u.OfExtractAsyncsBrowserString
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserObject) {
		return u.OfExtractAsyncsBrowserObject
	}
	return nil
}

// Browser type to emulate
type ExtractAsyncParamsBrowserString string

const (
	ExtractAsyncParamsBrowserStringChrome  ExtractAsyncParamsBrowserString = "chrome"
	ExtractAsyncParamsBrowserStringFirefox ExtractAsyncParamsBrowserString = "firefox"
)

// The property Name is required.
type ExtractAsyncParamsBrowserObject struct {
	// Any of "chrome", "firefox".
	Name string `json:"name,omitzero" api:"required"`
	// Specific browser version to emulate
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r ExtractAsyncParamsBrowserObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractAsyncParamsBrowserObject](
		"name", "chrome", "firefox",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionUnion struct {
	OfAutoScrollAction        *shared.AutoScrollActionParam        `json:",omitzero,inline"`
	OfClickAction             *shared.ClickActionParam             `json:",omitzero,inline"`
	OfEvalAction              *shared.EvalActionParam              `json:",omitzero,inline"`
	OfFetchAction             *shared.FetchActionParam             `json:",omitzero,inline"`
	OfFillAction              *shared.FillActionParam              `json:",omitzero,inline"`
	OfGetCookiesAction        *shared.GetCookiesActionParam        `json:",omitzero,inline"`
	OfGotoAction              *shared.GotoActionParam              `json:",omitzero,inline"`
	OfPressAction             *shared.PressActionParam             `json:",omitzero,inline"`
	OfScreenshotAction        *shared.ScreenshotActionParam        `json:",omitzero,inline"`
	OfScrollAction            *shared.ScrollActionParam            `json:",omitzero,inline"`
	OfWaitAction              *shared.WaitActionParam              `json:",omitzero,inline"`
	OfWaitForElementAction    *shared.WaitForElementActionParam    `json:",omitzero,inline"`
	OfWaitForNavigationAction *shared.WaitForNavigationActionParam `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAutoScrollAction,
		u.OfClickAction,
		u.OfEvalAction,
		u.OfFetchAction,
		u.OfFillAction,
		u.OfGetCookiesAction,
		u.OfGotoAction,
		u.OfPressAction,
		u.OfScreenshotAction,
		u.OfScrollAction,
		u.OfWaitAction,
		u.OfWaitForElementAction,
		u.OfWaitForNavigationAction)
}
func (u *ExtractAsyncParamsBrowserActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionUnion) asAny() any {
	if !param.IsOmitted(u.OfAutoScrollAction) {
		return u.OfAutoScrollAction
	} else if !param.IsOmitted(u.OfClickAction) {
		return u.OfClickAction
	} else if !param.IsOmitted(u.OfEvalAction) {
		return u.OfEvalAction
	} else if !param.IsOmitted(u.OfFetchAction) {
		return u.OfFetchAction
	} else if !param.IsOmitted(u.OfFillAction) {
		return u.OfFillAction
	} else if !param.IsOmitted(u.OfGetCookiesAction) {
		return u.OfGetCookiesAction
	} else if !param.IsOmitted(u.OfGotoAction) {
		return u.OfGotoAction
	} else if !param.IsOmitted(u.OfPressAction) {
		return u.OfPressAction
	} else if !param.IsOmitted(u.OfScreenshotAction) {
		return u.OfScreenshotAction
	} else if !param.IsOmitted(u.OfScrollAction) {
		return u.OfScrollAction
	} else if !param.IsOmitted(u.OfWaitAction) {
		return u.OfWaitAction
	} else if !param.IsOmitted(u.OfWaitForElementAction) {
		return u.OfWaitForElementAction
	} else if !param.IsOmitted(u.OfWaitForNavigationAction) {
		return u.OfWaitForNavigationAction
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsCookiesUnion struct {
	OfExtractAsyncsCookiesArray []ExtractAsyncParamsCookiesArrayItem `json:",omitzero,inline"`
	OfString                    param.Opt[string]                    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsCookiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsCookiesArray, u.OfString)
}
func (u *ExtractAsyncParamsCookiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsCookiesUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsCookiesArray) {
		return &u.OfExtractAsyncsCookiesArray
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type ExtractAsyncParamsCookiesArrayItem struct {
	Creation      param.Opt[string]                             `json:"creation,omitzero"`
	Domain        param.Opt[string]                             `json:"domain,omitzero"`
	HostOnly      param.Opt[bool]                               `json:"hostOnly,omitzero"`
	HTTPOnly      param.Opt[bool]                               `json:"httpOnly,omitzero"`
	LastAccessed  param.Opt[string]                             `json:"lastAccessed,omitzero"`
	Path          param.Opt[string]                             `json:"path,omitzero"`
	PathIsDefault param.Opt[bool]                               `json:"pathIsDefault,omitzero"`
	Expires       param.Opt[string]                             `json:"expires,omitzero"`
	Name          param.Opt[string]                             `json:"name,omitzero"`
	Secure        param.Opt[bool]                               `json:"secure,omitzero"`
	Value         param.Opt[string]                             `json:"value,omitzero"`
	Extensions    []string                                      `json:"extensions,omitzero"`
	MaxAge        ExtractAsyncParamsCookiesArrayItemMaxAgeUnion `json:"maxAge,omitzero"`
	// Any of "strict", "lax", "none".
	SameSite    string         `json:"sameSite,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r ExtractAsyncParamsCookiesArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsCookiesArrayItem
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ExtractAsyncParamsCookiesArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractAsyncParamsCookiesArrayItem](
		"sameSite", "strict", "lax", "none",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsCookiesArrayItemMaxAgeUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsCookiesArrayItemMaxAgeString)
	OfExtractAsyncsCookiesArrayItemMaxAgeString param.Opt[ExtractAsyncParamsCookiesArrayItemMaxAgeString] `json:",omitzero,inline"`
	OfFloat                                     param.Opt[float64]                                        `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsCookiesArrayItemMaxAgeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsCookiesArrayItemMaxAgeString, u.OfFloat)
}
func (u *ExtractAsyncParamsCookiesArrayItemMaxAgeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsCookiesArrayItemMaxAgeUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsCookiesArrayItemMaxAgeString) {
		return &u.OfExtractAsyncsCookiesArrayItemMaxAgeString
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type ExtractAsyncParamsCookiesArrayItemMaxAgeString string

const (
	ExtractAsyncParamsCookiesArrayItemMaxAgeStringInfinity      ExtractAsyncParamsCookiesArrayItemMaxAgeString = "Infinity"
	ExtractAsyncParamsCookiesArrayItemMaxAgeStringMinusInfinity ExtractAsyncParamsCookiesArrayItemMaxAgeString = "-Infinity"
)

// Country code for geolocation and proxy selection
type ExtractAsyncParamsCountry string

const (
	ExtractAsyncParamsCountryAd  ExtractAsyncParamsCountry = "AD"
	ExtractAsyncParamsCountryAe  ExtractAsyncParamsCountry = "AE"
	ExtractAsyncParamsCountryAf  ExtractAsyncParamsCountry = "AF"
	ExtractAsyncParamsCountryAg  ExtractAsyncParamsCountry = "AG"
	ExtractAsyncParamsCountryAI  ExtractAsyncParamsCountry = "AI"
	ExtractAsyncParamsCountryAl  ExtractAsyncParamsCountry = "AL"
	ExtractAsyncParamsCountryAm  ExtractAsyncParamsCountry = "AM"
	ExtractAsyncParamsCountryAo  ExtractAsyncParamsCountry = "AO"
	ExtractAsyncParamsCountryAq  ExtractAsyncParamsCountry = "AQ"
	ExtractAsyncParamsCountryAr  ExtractAsyncParamsCountry = "AR"
	ExtractAsyncParamsCountryAs  ExtractAsyncParamsCountry = "AS"
	ExtractAsyncParamsCountryAt  ExtractAsyncParamsCountry = "AT"
	ExtractAsyncParamsCountryAu  ExtractAsyncParamsCountry = "AU"
	ExtractAsyncParamsCountryAw  ExtractAsyncParamsCountry = "AW"
	ExtractAsyncParamsCountryAx  ExtractAsyncParamsCountry = "AX"
	ExtractAsyncParamsCountryAz  ExtractAsyncParamsCountry = "AZ"
	ExtractAsyncParamsCountryBa  ExtractAsyncParamsCountry = "BA"
	ExtractAsyncParamsCountryBb  ExtractAsyncParamsCountry = "BB"
	ExtractAsyncParamsCountryBd  ExtractAsyncParamsCountry = "BD"
	ExtractAsyncParamsCountryBe  ExtractAsyncParamsCountry = "BE"
	ExtractAsyncParamsCountryBf  ExtractAsyncParamsCountry = "BF"
	ExtractAsyncParamsCountryBg  ExtractAsyncParamsCountry = "BG"
	ExtractAsyncParamsCountryBh  ExtractAsyncParamsCountry = "BH"
	ExtractAsyncParamsCountryBi  ExtractAsyncParamsCountry = "BI"
	ExtractAsyncParamsCountryBj  ExtractAsyncParamsCountry = "BJ"
	ExtractAsyncParamsCountryBl  ExtractAsyncParamsCountry = "BL"
	ExtractAsyncParamsCountryBm  ExtractAsyncParamsCountry = "BM"
	ExtractAsyncParamsCountryBn  ExtractAsyncParamsCountry = "BN"
	ExtractAsyncParamsCountryBo  ExtractAsyncParamsCountry = "BO"
	ExtractAsyncParamsCountryBq  ExtractAsyncParamsCountry = "BQ"
	ExtractAsyncParamsCountryBr  ExtractAsyncParamsCountry = "BR"
	ExtractAsyncParamsCountryBs  ExtractAsyncParamsCountry = "BS"
	ExtractAsyncParamsCountryBt  ExtractAsyncParamsCountry = "BT"
	ExtractAsyncParamsCountryBv  ExtractAsyncParamsCountry = "BV"
	ExtractAsyncParamsCountryBw  ExtractAsyncParamsCountry = "BW"
	ExtractAsyncParamsCountryBy  ExtractAsyncParamsCountry = "BY"
	ExtractAsyncParamsCountryBz  ExtractAsyncParamsCountry = "BZ"
	ExtractAsyncParamsCountryCa  ExtractAsyncParamsCountry = "CA"
	ExtractAsyncParamsCountryCc  ExtractAsyncParamsCountry = "CC"
	ExtractAsyncParamsCountryCd  ExtractAsyncParamsCountry = "CD"
	ExtractAsyncParamsCountryCf  ExtractAsyncParamsCountry = "CF"
	ExtractAsyncParamsCountryCg  ExtractAsyncParamsCountry = "CG"
	ExtractAsyncParamsCountryCh  ExtractAsyncParamsCountry = "CH"
	ExtractAsyncParamsCountryCi  ExtractAsyncParamsCountry = "CI"
	ExtractAsyncParamsCountryCk  ExtractAsyncParamsCountry = "CK"
	ExtractAsyncParamsCountryCl  ExtractAsyncParamsCountry = "CL"
	ExtractAsyncParamsCountryCm  ExtractAsyncParamsCountry = "CM"
	ExtractAsyncParamsCountryCn  ExtractAsyncParamsCountry = "CN"
	ExtractAsyncParamsCountryCo  ExtractAsyncParamsCountry = "CO"
	ExtractAsyncParamsCountryCr  ExtractAsyncParamsCountry = "CR"
	ExtractAsyncParamsCountryCu  ExtractAsyncParamsCountry = "CU"
	ExtractAsyncParamsCountryCv  ExtractAsyncParamsCountry = "CV"
	ExtractAsyncParamsCountryCw  ExtractAsyncParamsCountry = "CW"
	ExtractAsyncParamsCountryCx  ExtractAsyncParamsCountry = "CX"
	ExtractAsyncParamsCountryCy  ExtractAsyncParamsCountry = "CY"
	ExtractAsyncParamsCountryCz  ExtractAsyncParamsCountry = "CZ"
	ExtractAsyncParamsCountryDe  ExtractAsyncParamsCountry = "DE"
	ExtractAsyncParamsCountryDj  ExtractAsyncParamsCountry = "DJ"
	ExtractAsyncParamsCountryDk  ExtractAsyncParamsCountry = "DK"
	ExtractAsyncParamsCountryDm  ExtractAsyncParamsCountry = "DM"
	ExtractAsyncParamsCountryDo  ExtractAsyncParamsCountry = "DO"
	ExtractAsyncParamsCountryDz  ExtractAsyncParamsCountry = "DZ"
	ExtractAsyncParamsCountryEc  ExtractAsyncParamsCountry = "EC"
	ExtractAsyncParamsCountryEe  ExtractAsyncParamsCountry = "EE"
	ExtractAsyncParamsCountryEg  ExtractAsyncParamsCountry = "EG"
	ExtractAsyncParamsCountryEh  ExtractAsyncParamsCountry = "EH"
	ExtractAsyncParamsCountryEr  ExtractAsyncParamsCountry = "ER"
	ExtractAsyncParamsCountryEs  ExtractAsyncParamsCountry = "ES"
	ExtractAsyncParamsCountryEt  ExtractAsyncParamsCountry = "ET"
	ExtractAsyncParamsCountryFi  ExtractAsyncParamsCountry = "FI"
	ExtractAsyncParamsCountryFj  ExtractAsyncParamsCountry = "FJ"
	ExtractAsyncParamsCountryFk  ExtractAsyncParamsCountry = "FK"
	ExtractAsyncParamsCountryFm  ExtractAsyncParamsCountry = "FM"
	ExtractAsyncParamsCountryFo  ExtractAsyncParamsCountry = "FO"
	ExtractAsyncParamsCountryFr  ExtractAsyncParamsCountry = "FR"
	ExtractAsyncParamsCountryGa  ExtractAsyncParamsCountry = "GA"
	ExtractAsyncParamsCountryGB  ExtractAsyncParamsCountry = "GB"
	ExtractAsyncParamsCountryGd  ExtractAsyncParamsCountry = "GD"
	ExtractAsyncParamsCountryGe  ExtractAsyncParamsCountry = "GE"
	ExtractAsyncParamsCountryGf  ExtractAsyncParamsCountry = "GF"
	ExtractAsyncParamsCountryGg  ExtractAsyncParamsCountry = "GG"
	ExtractAsyncParamsCountryGh  ExtractAsyncParamsCountry = "GH"
	ExtractAsyncParamsCountryGi  ExtractAsyncParamsCountry = "GI"
	ExtractAsyncParamsCountryGl  ExtractAsyncParamsCountry = "GL"
	ExtractAsyncParamsCountryGm  ExtractAsyncParamsCountry = "GM"
	ExtractAsyncParamsCountryGn  ExtractAsyncParamsCountry = "GN"
	ExtractAsyncParamsCountryGp  ExtractAsyncParamsCountry = "GP"
	ExtractAsyncParamsCountryGq  ExtractAsyncParamsCountry = "GQ"
	ExtractAsyncParamsCountryGr  ExtractAsyncParamsCountry = "GR"
	ExtractAsyncParamsCountryGs  ExtractAsyncParamsCountry = "GS"
	ExtractAsyncParamsCountryGt  ExtractAsyncParamsCountry = "GT"
	ExtractAsyncParamsCountryGu  ExtractAsyncParamsCountry = "GU"
	ExtractAsyncParamsCountryGw  ExtractAsyncParamsCountry = "GW"
	ExtractAsyncParamsCountryGy  ExtractAsyncParamsCountry = "GY"
	ExtractAsyncParamsCountryHk  ExtractAsyncParamsCountry = "HK"
	ExtractAsyncParamsCountryHm  ExtractAsyncParamsCountry = "HM"
	ExtractAsyncParamsCountryHn  ExtractAsyncParamsCountry = "HN"
	ExtractAsyncParamsCountryHr  ExtractAsyncParamsCountry = "HR"
	ExtractAsyncParamsCountryHt  ExtractAsyncParamsCountry = "HT"
	ExtractAsyncParamsCountryHu  ExtractAsyncParamsCountry = "HU"
	ExtractAsyncParamsCountryID  ExtractAsyncParamsCountry = "ID"
	ExtractAsyncParamsCountryIe  ExtractAsyncParamsCountry = "IE"
	ExtractAsyncParamsCountryIl  ExtractAsyncParamsCountry = "IL"
	ExtractAsyncParamsCountryIm  ExtractAsyncParamsCountry = "IM"
	ExtractAsyncParamsCountryIn  ExtractAsyncParamsCountry = "IN"
	ExtractAsyncParamsCountryIo  ExtractAsyncParamsCountry = "IO"
	ExtractAsyncParamsCountryIq  ExtractAsyncParamsCountry = "IQ"
	ExtractAsyncParamsCountryIr  ExtractAsyncParamsCountry = "IR"
	ExtractAsyncParamsCountryIs  ExtractAsyncParamsCountry = "IS"
	ExtractAsyncParamsCountryIt  ExtractAsyncParamsCountry = "IT"
	ExtractAsyncParamsCountryJe  ExtractAsyncParamsCountry = "JE"
	ExtractAsyncParamsCountryJm  ExtractAsyncParamsCountry = "JM"
	ExtractAsyncParamsCountryJo  ExtractAsyncParamsCountry = "JO"
	ExtractAsyncParamsCountryJp  ExtractAsyncParamsCountry = "JP"
	ExtractAsyncParamsCountryKe  ExtractAsyncParamsCountry = "KE"
	ExtractAsyncParamsCountryKg  ExtractAsyncParamsCountry = "KG"
	ExtractAsyncParamsCountryKh  ExtractAsyncParamsCountry = "KH"
	ExtractAsyncParamsCountryKi  ExtractAsyncParamsCountry = "KI"
	ExtractAsyncParamsCountryKm  ExtractAsyncParamsCountry = "KM"
	ExtractAsyncParamsCountryKn  ExtractAsyncParamsCountry = "KN"
	ExtractAsyncParamsCountryKp  ExtractAsyncParamsCountry = "KP"
	ExtractAsyncParamsCountryKr  ExtractAsyncParamsCountry = "KR"
	ExtractAsyncParamsCountryKw  ExtractAsyncParamsCountry = "KW"
	ExtractAsyncParamsCountryKy  ExtractAsyncParamsCountry = "KY"
	ExtractAsyncParamsCountryKz  ExtractAsyncParamsCountry = "KZ"
	ExtractAsyncParamsCountryLa  ExtractAsyncParamsCountry = "LA"
	ExtractAsyncParamsCountryLb  ExtractAsyncParamsCountry = "LB"
	ExtractAsyncParamsCountryLc  ExtractAsyncParamsCountry = "LC"
	ExtractAsyncParamsCountryLi  ExtractAsyncParamsCountry = "LI"
	ExtractAsyncParamsCountryLk  ExtractAsyncParamsCountry = "LK"
	ExtractAsyncParamsCountryLr  ExtractAsyncParamsCountry = "LR"
	ExtractAsyncParamsCountryLs  ExtractAsyncParamsCountry = "LS"
	ExtractAsyncParamsCountryLt  ExtractAsyncParamsCountry = "LT"
	ExtractAsyncParamsCountryLu  ExtractAsyncParamsCountry = "LU"
	ExtractAsyncParamsCountryLv  ExtractAsyncParamsCountry = "LV"
	ExtractAsyncParamsCountryLy  ExtractAsyncParamsCountry = "LY"
	ExtractAsyncParamsCountryMa  ExtractAsyncParamsCountry = "MA"
	ExtractAsyncParamsCountryMc  ExtractAsyncParamsCountry = "MC"
	ExtractAsyncParamsCountryMd  ExtractAsyncParamsCountry = "MD"
	ExtractAsyncParamsCountryMe  ExtractAsyncParamsCountry = "ME"
	ExtractAsyncParamsCountryMf  ExtractAsyncParamsCountry = "MF"
	ExtractAsyncParamsCountryMg  ExtractAsyncParamsCountry = "MG"
	ExtractAsyncParamsCountryMh  ExtractAsyncParamsCountry = "MH"
	ExtractAsyncParamsCountryMk  ExtractAsyncParamsCountry = "MK"
	ExtractAsyncParamsCountryMl  ExtractAsyncParamsCountry = "ML"
	ExtractAsyncParamsCountryMm  ExtractAsyncParamsCountry = "MM"
	ExtractAsyncParamsCountryMn  ExtractAsyncParamsCountry = "MN"
	ExtractAsyncParamsCountryMo  ExtractAsyncParamsCountry = "MO"
	ExtractAsyncParamsCountryMp  ExtractAsyncParamsCountry = "MP"
	ExtractAsyncParamsCountryMq  ExtractAsyncParamsCountry = "MQ"
	ExtractAsyncParamsCountryMr  ExtractAsyncParamsCountry = "MR"
	ExtractAsyncParamsCountryMs  ExtractAsyncParamsCountry = "MS"
	ExtractAsyncParamsCountryMt  ExtractAsyncParamsCountry = "MT"
	ExtractAsyncParamsCountryMu  ExtractAsyncParamsCountry = "MU"
	ExtractAsyncParamsCountryMv  ExtractAsyncParamsCountry = "MV"
	ExtractAsyncParamsCountryMw  ExtractAsyncParamsCountry = "MW"
	ExtractAsyncParamsCountryMx  ExtractAsyncParamsCountry = "MX"
	ExtractAsyncParamsCountryMy  ExtractAsyncParamsCountry = "MY"
	ExtractAsyncParamsCountryMz  ExtractAsyncParamsCountry = "MZ"
	ExtractAsyncParamsCountryNa  ExtractAsyncParamsCountry = "NA"
	ExtractAsyncParamsCountryNc  ExtractAsyncParamsCountry = "NC"
	ExtractAsyncParamsCountryNe  ExtractAsyncParamsCountry = "NE"
	ExtractAsyncParamsCountryNf  ExtractAsyncParamsCountry = "NF"
	ExtractAsyncParamsCountryNg  ExtractAsyncParamsCountry = "NG"
	ExtractAsyncParamsCountryNi  ExtractAsyncParamsCountry = "NI"
	ExtractAsyncParamsCountryNl  ExtractAsyncParamsCountry = "NL"
	ExtractAsyncParamsCountryNo  ExtractAsyncParamsCountry = "NO"
	ExtractAsyncParamsCountryNp  ExtractAsyncParamsCountry = "NP"
	ExtractAsyncParamsCountryNr  ExtractAsyncParamsCountry = "NR"
	ExtractAsyncParamsCountryNu  ExtractAsyncParamsCountry = "NU"
	ExtractAsyncParamsCountryNz  ExtractAsyncParamsCountry = "NZ"
	ExtractAsyncParamsCountryOm  ExtractAsyncParamsCountry = "OM"
	ExtractAsyncParamsCountryPa  ExtractAsyncParamsCountry = "PA"
	ExtractAsyncParamsCountryPe  ExtractAsyncParamsCountry = "PE"
	ExtractAsyncParamsCountryPf  ExtractAsyncParamsCountry = "PF"
	ExtractAsyncParamsCountryPg  ExtractAsyncParamsCountry = "PG"
	ExtractAsyncParamsCountryPh  ExtractAsyncParamsCountry = "PH"
	ExtractAsyncParamsCountryPk  ExtractAsyncParamsCountry = "PK"
	ExtractAsyncParamsCountryPl  ExtractAsyncParamsCountry = "PL"
	ExtractAsyncParamsCountryPm  ExtractAsyncParamsCountry = "PM"
	ExtractAsyncParamsCountryPn  ExtractAsyncParamsCountry = "PN"
	ExtractAsyncParamsCountryPr  ExtractAsyncParamsCountry = "PR"
	ExtractAsyncParamsCountryPs  ExtractAsyncParamsCountry = "PS"
	ExtractAsyncParamsCountryPt  ExtractAsyncParamsCountry = "PT"
	ExtractAsyncParamsCountryPw  ExtractAsyncParamsCountry = "PW"
	ExtractAsyncParamsCountryPy  ExtractAsyncParamsCountry = "PY"
	ExtractAsyncParamsCountryQa  ExtractAsyncParamsCountry = "QA"
	ExtractAsyncParamsCountryRe  ExtractAsyncParamsCountry = "RE"
	ExtractAsyncParamsCountryRo  ExtractAsyncParamsCountry = "RO"
	ExtractAsyncParamsCountryRs  ExtractAsyncParamsCountry = "RS"
	ExtractAsyncParamsCountryRu  ExtractAsyncParamsCountry = "RU"
	ExtractAsyncParamsCountryRw  ExtractAsyncParamsCountry = "RW"
	ExtractAsyncParamsCountrySa  ExtractAsyncParamsCountry = "SA"
	ExtractAsyncParamsCountrySb  ExtractAsyncParamsCountry = "SB"
	ExtractAsyncParamsCountrySc  ExtractAsyncParamsCountry = "SC"
	ExtractAsyncParamsCountrySd  ExtractAsyncParamsCountry = "SD"
	ExtractAsyncParamsCountrySe  ExtractAsyncParamsCountry = "SE"
	ExtractAsyncParamsCountrySg  ExtractAsyncParamsCountry = "SG"
	ExtractAsyncParamsCountrySh  ExtractAsyncParamsCountry = "SH"
	ExtractAsyncParamsCountrySi  ExtractAsyncParamsCountry = "SI"
	ExtractAsyncParamsCountrySj  ExtractAsyncParamsCountry = "SJ"
	ExtractAsyncParamsCountrySk  ExtractAsyncParamsCountry = "SK"
	ExtractAsyncParamsCountrySl  ExtractAsyncParamsCountry = "SL"
	ExtractAsyncParamsCountrySm  ExtractAsyncParamsCountry = "SM"
	ExtractAsyncParamsCountrySn  ExtractAsyncParamsCountry = "SN"
	ExtractAsyncParamsCountrySo  ExtractAsyncParamsCountry = "SO"
	ExtractAsyncParamsCountrySr  ExtractAsyncParamsCountry = "SR"
	ExtractAsyncParamsCountrySS  ExtractAsyncParamsCountry = "SS"
	ExtractAsyncParamsCountrySt  ExtractAsyncParamsCountry = "ST"
	ExtractAsyncParamsCountrySv  ExtractAsyncParamsCountry = "SV"
	ExtractAsyncParamsCountrySx  ExtractAsyncParamsCountry = "SX"
	ExtractAsyncParamsCountrySy  ExtractAsyncParamsCountry = "SY"
	ExtractAsyncParamsCountrySz  ExtractAsyncParamsCountry = "SZ"
	ExtractAsyncParamsCountryTc  ExtractAsyncParamsCountry = "TC"
	ExtractAsyncParamsCountryTd  ExtractAsyncParamsCountry = "TD"
	ExtractAsyncParamsCountryTf  ExtractAsyncParamsCountry = "TF"
	ExtractAsyncParamsCountryTg  ExtractAsyncParamsCountry = "TG"
	ExtractAsyncParamsCountryTh  ExtractAsyncParamsCountry = "TH"
	ExtractAsyncParamsCountryTj  ExtractAsyncParamsCountry = "TJ"
	ExtractAsyncParamsCountryTk  ExtractAsyncParamsCountry = "TK"
	ExtractAsyncParamsCountryTl  ExtractAsyncParamsCountry = "TL"
	ExtractAsyncParamsCountryTm  ExtractAsyncParamsCountry = "TM"
	ExtractAsyncParamsCountryTn  ExtractAsyncParamsCountry = "TN"
	ExtractAsyncParamsCountryTo  ExtractAsyncParamsCountry = "TO"
	ExtractAsyncParamsCountryTr  ExtractAsyncParamsCountry = "TR"
	ExtractAsyncParamsCountryTt  ExtractAsyncParamsCountry = "TT"
	ExtractAsyncParamsCountryTv  ExtractAsyncParamsCountry = "TV"
	ExtractAsyncParamsCountryTw  ExtractAsyncParamsCountry = "TW"
	ExtractAsyncParamsCountryTz  ExtractAsyncParamsCountry = "TZ"
	ExtractAsyncParamsCountryUa  ExtractAsyncParamsCountry = "UA"
	ExtractAsyncParamsCountryUg  ExtractAsyncParamsCountry = "UG"
	ExtractAsyncParamsCountryUm  ExtractAsyncParamsCountry = "UM"
	ExtractAsyncParamsCountryUs  ExtractAsyncParamsCountry = "US"
	ExtractAsyncParamsCountryUy  ExtractAsyncParamsCountry = "UY"
	ExtractAsyncParamsCountryUz  ExtractAsyncParamsCountry = "UZ"
	ExtractAsyncParamsCountryVa  ExtractAsyncParamsCountry = "VA"
	ExtractAsyncParamsCountryVc  ExtractAsyncParamsCountry = "VC"
	ExtractAsyncParamsCountryVe  ExtractAsyncParamsCountry = "VE"
	ExtractAsyncParamsCountryVg  ExtractAsyncParamsCountry = "VG"
	ExtractAsyncParamsCountryVi  ExtractAsyncParamsCountry = "VI"
	ExtractAsyncParamsCountryVn  ExtractAsyncParamsCountry = "VN"
	ExtractAsyncParamsCountryVu  ExtractAsyncParamsCountry = "VU"
	ExtractAsyncParamsCountryWf  ExtractAsyncParamsCountry = "WF"
	ExtractAsyncParamsCountryWs  ExtractAsyncParamsCountry = "WS"
	ExtractAsyncParamsCountryXk  ExtractAsyncParamsCountry = "XK"
	ExtractAsyncParamsCountryYe  ExtractAsyncParamsCountry = "YE"
	ExtractAsyncParamsCountryYt  ExtractAsyncParamsCountry = "YT"
	ExtractAsyncParamsCountryZa  ExtractAsyncParamsCountry = "ZA"
	ExtractAsyncParamsCountryZm  ExtractAsyncParamsCountry = "ZM"
	ExtractAsyncParamsCountryZw  ExtractAsyncParamsCountry = "ZW"
	ExtractAsyncParamsCountryAll ExtractAsyncParamsCountry = "ALL"
)

// Device type for browser emulation
type ExtractAsyncParamsDevice string

const (
	ExtractAsyncParamsDeviceDesktop ExtractAsyncParamsDevice = "desktop"
	ExtractAsyncParamsDeviceMobile  ExtractAsyncParamsDevice = "mobile"
	ExtractAsyncParamsDeviceTablet  ExtractAsyncParamsDevice = "tablet"
)

// Browser driver to use
type ExtractAsyncParamsDriver string

const (
	ExtractAsyncParamsDriverVx6     ExtractAsyncParamsDriver = "vx6"
	ExtractAsyncParamsDriverVx8     ExtractAsyncParamsDriver = "vx8"
	ExtractAsyncParamsDriverVx8Pro  ExtractAsyncParamsDriver = "vx8-pro"
	ExtractAsyncParamsDriverVx10    ExtractAsyncParamsDriver = "vx10"
	ExtractAsyncParamsDriverVx10Pro ExtractAsyncParamsDriver = "vx10-pro"
	ExtractAsyncParamsDriverVx12    ExtractAsyncParamsDriver = "vx12"
	ExtractAsyncParamsDriverVx12Pro ExtractAsyncParamsDriver = "vx12-pro"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsHeaderUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsHeaderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractAsyncParamsHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsHeaderUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Locale for browser language and region settings
type ExtractAsyncParamsLocale string

const (
	ExtractAsyncParamsLocaleAaDj      ExtractAsyncParamsLocale = "aa-DJ"
	ExtractAsyncParamsLocaleAaEr      ExtractAsyncParamsLocale = "aa-ER"
	ExtractAsyncParamsLocaleAaEt      ExtractAsyncParamsLocale = "aa-ET"
	ExtractAsyncParamsLocaleAf        ExtractAsyncParamsLocale = "af"
	ExtractAsyncParamsLocaleAfNa      ExtractAsyncParamsLocale = "af-NA"
	ExtractAsyncParamsLocaleAfZa      ExtractAsyncParamsLocale = "af-ZA"
	ExtractAsyncParamsLocaleAk        ExtractAsyncParamsLocale = "ak"
	ExtractAsyncParamsLocaleAkGh      ExtractAsyncParamsLocale = "ak-GH"
	ExtractAsyncParamsLocaleAm        ExtractAsyncParamsLocale = "am"
	ExtractAsyncParamsLocaleAmEt      ExtractAsyncParamsLocale = "am-ET"
	ExtractAsyncParamsLocaleAnEs      ExtractAsyncParamsLocale = "an-ES"
	ExtractAsyncParamsLocaleAr        ExtractAsyncParamsLocale = "ar"
	ExtractAsyncParamsLocaleArAe      ExtractAsyncParamsLocale = "ar-AE"
	ExtractAsyncParamsLocaleArBh      ExtractAsyncParamsLocale = "ar-BH"
	ExtractAsyncParamsLocaleArDz      ExtractAsyncParamsLocale = "ar-DZ"
	ExtractAsyncParamsLocaleArEg      ExtractAsyncParamsLocale = "ar-EG"
	ExtractAsyncParamsLocaleArIn      ExtractAsyncParamsLocale = "ar-IN"
	ExtractAsyncParamsLocaleArIq      ExtractAsyncParamsLocale = "ar-IQ"
	ExtractAsyncParamsLocaleArJo      ExtractAsyncParamsLocale = "ar-JO"
	ExtractAsyncParamsLocaleArKw      ExtractAsyncParamsLocale = "ar-KW"
	ExtractAsyncParamsLocaleArLb      ExtractAsyncParamsLocale = "ar-LB"
	ExtractAsyncParamsLocaleArLy      ExtractAsyncParamsLocale = "ar-LY"
	ExtractAsyncParamsLocaleArMa      ExtractAsyncParamsLocale = "ar-MA"
	ExtractAsyncParamsLocaleArOm      ExtractAsyncParamsLocale = "ar-OM"
	ExtractAsyncParamsLocaleArQa      ExtractAsyncParamsLocale = "ar-QA"
	ExtractAsyncParamsLocaleArSa      ExtractAsyncParamsLocale = "ar-SA"
	ExtractAsyncParamsLocaleArSd      ExtractAsyncParamsLocale = "ar-SD"
	ExtractAsyncParamsLocaleArSy      ExtractAsyncParamsLocale = "ar-SY"
	ExtractAsyncParamsLocaleArTn      ExtractAsyncParamsLocale = "ar-TN"
	ExtractAsyncParamsLocaleArYe      ExtractAsyncParamsLocale = "ar-YE"
	ExtractAsyncParamsLocaleAs        ExtractAsyncParamsLocale = "as"
	ExtractAsyncParamsLocaleAsIn      ExtractAsyncParamsLocale = "as-IN"
	ExtractAsyncParamsLocaleAsa       ExtractAsyncParamsLocale = "asa"
	ExtractAsyncParamsLocaleAsaTz     ExtractAsyncParamsLocale = "asa-TZ"
	ExtractAsyncParamsLocaleAstEs     ExtractAsyncParamsLocale = "ast-ES"
	ExtractAsyncParamsLocaleAz        ExtractAsyncParamsLocale = "az"
	ExtractAsyncParamsLocaleAzAz      ExtractAsyncParamsLocale = "az-AZ"
	ExtractAsyncParamsLocaleAzCyrl    ExtractAsyncParamsLocale = "az-Cyrl"
	ExtractAsyncParamsLocaleAzCyrlAz  ExtractAsyncParamsLocale = "az-Cyrl-AZ"
	ExtractAsyncParamsLocaleAzLatn    ExtractAsyncParamsLocale = "az-Latn"
	ExtractAsyncParamsLocaleAzLatnAz  ExtractAsyncParamsLocale = "az-Latn-AZ"
	ExtractAsyncParamsLocaleBe        ExtractAsyncParamsLocale = "be"
	ExtractAsyncParamsLocaleBeBy      ExtractAsyncParamsLocale = "be-BY"
	ExtractAsyncParamsLocaleBem       ExtractAsyncParamsLocale = "bem"
	ExtractAsyncParamsLocaleBemZm     ExtractAsyncParamsLocale = "bem-ZM"
	ExtractAsyncParamsLocaleBerDz     ExtractAsyncParamsLocale = "ber-DZ"
	ExtractAsyncParamsLocaleBerMa     ExtractAsyncParamsLocale = "ber-MA"
	ExtractAsyncParamsLocaleBez       ExtractAsyncParamsLocale = "bez"
	ExtractAsyncParamsLocaleBezTz     ExtractAsyncParamsLocale = "bez-TZ"
	ExtractAsyncParamsLocaleBg        ExtractAsyncParamsLocale = "bg"
	ExtractAsyncParamsLocaleBgBg      ExtractAsyncParamsLocale = "bg-BG"
	ExtractAsyncParamsLocaleBhoIn     ExtractAsyncParamsLocale = "bho-IN"
	ExtractAsyncParamsLocaleBm        ExtractAsyncParamsLocale = "bm"
	ExtractAsyncParamsLocaleBmMl      ExtractAsyncParamsLocale = "bm-ML"
	ExtractAsyncParamsLocaleBn        ExtractAsyncParamsLocale = "bn"
	ExtractAsyncParamsLocaleBnBd      ExtractAsyncParamsLocale = "bn-BD"
	ExtractAsyncParamsLocaleBnIn      ExtractAsyncParamsLocale = "bn-IN"
	ExtractAsyncParamsLocaleBo        ExtractAsyncParamsLocale = "bo"
	ExtractAsyncParamsLocaleBoCn      ExtractAsyncParamsLocale = "bo-CN"
	ExtractAsyncParamsLocaleBoIn      ExtractAsyncParamsLocale = "bo-IN"
	ExtractAsyncParamsLocaleBrFr      ExtractAsyncParamsLocale = "br-FR"
	ExtractAsyncParamsLocaleBrxIn     ExtractAsyncParamsLocale = "brx-IN"
	ExtractAsyncParamsLocaleBs        ExtractAsyncParamsLocale = "bs"
	ExtractAsyncParamsLocaleBsBa      ExtractAsyncParamsLocale = "bs-BA"
	ExtractAsyncParamsLocaleBynEr     ExtractAsyncParamsLocale = "byn-ER"
	ExtractAsyncParamsLocaleCa        ExtractAsyncParamsLocale = "ca"
	ExtractAsyncParamsLocaleCaAd      ExtractAsyncParamsLocale = "ca-AD"
	ExtractAsyncParamsLocaleCaEs      ExtractAsyncParamsLocale = "ca-ES"
	ExtractAsyncParamsLocaleCaFr      ExtractAsyncParamsLocale = "ca-FR"
	ExtractAsyncParamsLocaleCaIt      ExtractAsyncParamsLocale = "ca-IT"
	ExtractAsyncParamsLocaleCgg       ExtractAsyncParamsLocale = "cgg"
	ExtractAsyncParamsLocaleCggUg     ExtractAsyncParamsLocale = "cgg-UG"
	ExtractAsyncParamsLocaleChr       ExtractAsyncParamsLocale = "chr"
	ExtractAsyncParamsLocaleChrUs     ExtractAsyncParamsLocale = "chr-US"
	ExtractAsyncParamsLocaleCrhUa     ExtractAsyncParamsLocale = "crh-UA"
	ExtractAsyncParamsLocaleCs        ExtractAsyncParamsLocale = "cs"
	ExtractAsyncParamsLocaleCsCz      ExtractAsyncParamsLocale = "cs-CZ"
	ExtractAsyncParamsLocaleCsbPl     ExtractAsyncParamsLocale = "csb-PL"
	ExtractAsyncParamsLocaleCvRu      ExtractAsyncParamsLocale = "cv-RU"
	ExtractAsyncParamsLocaleCy        ExtractAsyncParamsLocale = "cy"
	ExtractAsyncParamsLocaleCyGB      ExtractAsyncParamsLocale = "cy-GB"
	ExtractAsyncParamsLocaleDa        ExtractAsyncParamsLocale = "da"
	ExtractAsyncParamsLocaleDaDk      ExtractAsyncParamsLocale = "da-DK"
	ExtractAsyncParamsLocaleDav       ExtractAsyncParamsLocale = "dav"
	ExtractAsyncParamsLocaleDavKe     ExtractAsyncParamsLocale = "dav-KE"
	ExtractAsyncParamsLocaleDe        ExtractAsyncParamsLocale = "de"
	ExtractAsyncParamsLocaleDeAt      ExtractAsyncParamsLocale = "de-AT"
	ExtractAsyncParamsLocaleDeBe      ExtractAsyncParamsLocale = "de-BE"
	ExtractAsyncParamsLocaleDeCh      ExtractAsyncParamsLocale = "de-CH"
	ExtractAsyncParamsLocaleDeDe      ExtractAsyncParamsLocale = "de-DE"
	ExtractAsyncParamsLocaleDeLi      ExtractAsyncParamsLocale = "de-LI"
	ExtractAsyncParamsLocaleDeLu      ExtractAsyncParamsLocale = "de-LU"
	ExtractAsyncParamsLocaleDvMv      ExtractAsyncParamsLocale = "dv-MV"
	ExtractAsyncParamsLocaleDzBt      ExtractAsyncParamsLocale = "dz-BT"
	ExtractAsyncParamsLocaleEbu       ExtractAsyncParamsLocale = "ebu"
	ExtractAsyncParamsLocaleEbuKe     ExtractAsyncParamsLocale = "ebu-KE"
	ExtractAsyncParamsLocaleEe        ExtractAsyncParamsLocale = "ee"
	ExtractAsyncParamsLocaleEeGh      ExtractAsyncParamsLocale = "ee-GH"
	ExtractAsyncParamsLocaleEeTg      ExtractAsyncParamsLocale = "ee-TG"
	ExtractAsyncParamsLocaleEl        ExtractAsyncParamsLocale = "el"
	ExtractAsyncParamsLocaleElCy      ExtractAsyncParamsLocale = "el-CY"
	ExtractAsyncParamsLocaleElGr      ExtractAsyncParamsLocale = "el-GR"
	ExtractAsyncParamsLocaleEn        ExtractAsyncParamsLocale = "en"
	ExtractAsyncParamsLocaleEnAg      ExtractAsyncParamsLocale = "en-AG"
	ExtractAsyncParamsLocaleEnAs      ExtractAsyncParamsLocale = "en-AS"
	ExtractAsyncParamsLocaleEnAu      ExtractAsyncParamsLocale = "en-AU"
	ExtractAsyncParamsLocaleEnBe      ExtractAsyncParamsLocale = "en-BE"
	ExtractAsyncParamsLocaleEnBw      ExtractAsyncParamsLocale = "en-BW"
	ExtractAsyncParamsLocaleEnBz      ExtractAsyncParamsLocale = "en-BZ"
	ExtractAsyncParamsLocaleEnCa      ExtractAsyncParamsLocale = "en-CA"
	ExtractAsyncParamsLocaleEnDk      ExtractAsyncParamsLocale = "en-DK"
	ExtractAsyncParamsLocaleEnGB      ExtractAsyncParamsLocale = "en-GB"
	ExtractAsyncParamsLocaleEnGu      ExtractAsyncParamsLocale = "en-GU"
	ExtractAsyncParamsLocaleEnHk      ExtractAsyncParamsLocale = "en-HK"
	ExtractAsyncParamsLocaleEnIe      ExtractAsyncParamsLocale = "en-IE"
	ExtractAsyncParamsLocaleEnIn      ExtractAsyncParamsLocale = "en-IN"
	ExtractAsyncParamsLocaleEnJm      ExtractAsyncParamsLocale = "en-JM"
	ExtractAsyncParamsLocaleEnMh      ExtractAsyncParamsLocale = "en-MH"
	ExtractAsyncParamsLocaleEnMp      ExtractAsyncParamsLocale = "en-MP"
	ExtractAsyncParamsLocaleEnMt      ExtractAsyncParamsLocale = "en-MT"
	ExtractAsyncParamsLocaleEnMu      ExtractAsyncParamsLocale = "en-MU"
	ExtractAsyncParamsLocaleEnNa      ExtractAsyncParamsLocale = "en-NA"
	ExtractAsyncParamsLocaleEnNg      ExtractAsyncParamsLocale = "en-NG"
	ExtractAsyncParamsLocaleEnNz      ExtractAsyncParamsLocale = "en-NZ"
	ExtractAsyncParamsLocaleEnPh      ExtractAsyncParamsLocale = "en-PH"
	ExtractAsyncParamsLocaleEnPk      ExtractAsyncParamsLocale = "en-PK"
	ExtractAsyncParamsLocaleEnSg      ExtractAsyncParamsLocale = "en-SG"
	ExtractAsyncParamsLocaleEnTt      ExtractAsyncParamsLocale = "en-TT"
	ExtractAsyncParamsLocaleEnUm      ExtractAsyncParamsLocale = "en-UM"
	ExtractAsyncParamsLocaleEnUs      ExtractAsyncParamsLocale = "en-US"
	ExtractAsyncParamsLocaleEnVi      ExtractAsyncParamsLocale = "en-VI"
	ExtractAsyncParamsLocaleEnZa      ExtractAsyncParamsLocale = "en-ZA"
	ExtractAsyncParamsLocaleEnZm      ExtractAsyncParamsLocale = "en-ZM"
	ExtractAsyncParamsLocaleEnZw      ExtractAsyncParamsLocale = "en-ZW"
	ExtractAsyncParamsLocaleEo        ExtractAsyncParamsLocale = "eo"
	ExtractAsyncParamsLocaleEs        ExtractAsyncParamsLocale = "es"
	ExtractAsyncParamsLocaleEs419     ExtractAsyncParamsLocale = "es-419"
	ExtractAsyncParamsLocaleEsAr      ExtractAsyncParamsLocale = "es-AR"
	ExtractAsyncParamsLocaleEsBo      ExtractAsyncParamsLocale = "es-BO"
	ExtractAsyncParamsLocaleEsCl      ExtractAsyncParamsLocale = "es-CL"
	ExtractAsyncParamsLocaleEsCo      ExtractAsyncParamsLocale = "es-CO"
	ExtractAsyncParamsLocaleEsCr      ExtractAsyncParamsLocale = "es-CR"
	ExtractAsyncParamsLocaleEsCu      ExtractAsyncParamsLocale = "es-CU"
	ExtractAsyncParamsLocaleEsDo      ExtractAsyncParamsLocale = "es-DO"
	ExtractAsyncParamsLocaleEsEc      ExtractAsyncParamsLocale = "es-EC"
	ExtractAsyncParamsLocaleEsEs      ExtractAsyncParamsLocale = "es-ES"
	ExtractAsyncParamsLocaleEsGq      ExtractAsyncParamsLocale = "es-GQ"
	ExtractAsyncParamsLocaleEsGt      ExtractAsyncParamsLocale = "es-GT"
	ExtractAsyncParamsLocaleEsHn      ExtractAsyncParamsLocale = "es-HN"
	ExtractAsyncParamsLocaleEsMx      ExtractAsyncParamsLocale = "es-MX"
	ExtractAsyncParamsLocaleEsNi      ExtractAsyncParamsLocale = "es-NI"
	ExtractAsyncParamsLocaleEsPa      ExtractAsyncParamsLocale = "es-PA"
	ExtractAsyncParamsLocaleEsPe      ExtractAsyncParamsLocale = "es-PE"
	ExtractAsyncParamsLocaleEsPr      ExtractAsyncParamsLocale = "es-PR"
	ExtractAsyncParamsLocaleEsPy      ExtractAsyncParamsLocale = "es-PY"
	ExtractAsyncParamsLocaleEsSv      ExtractAsyncParamsLocale = "es-SV"
	ExtractAsyncParamsLocaleEsUs      ExtractAsyncParamsLocale = "es-US"
	ExtractAsyncParamsLocaleEsUy      ExtractAsyncParamsLocale = "es-UY"
	ExtractAsyncParamsLocaleEsVe      ExtractAsyncParamsLocale = "es-VE"
	ExtractAsyncParamsLocaleEt        ExtractAsyncParamsLocale = "et"
	ExtractAsyncParamsLocaleEtEe      ExtractAsyncParamsLocale = "et-EE"
	ExtractAsyncParamsLocaleEu        ExtractAsyncParamsLocale = "eu"
	ExtractAsyncParamsLocaleEuEs      ExtractAsyncParamsLocale = "eu-ES"
	ExtractAsyncParamsLocaleFa        ExtractAsyncParamsLocale = "fa"
	ExtractAsyncParamsLocaleFaAf      ExtractAsyncParamsLocale = "fa-AF"
	ExtractAsyncParamsLocaleFaIr      ExtractAsyncParamsLocale = "fa-IR"
	ExtractAsyncParamsLocaleFf        ExtractAsyncParamsLocale = "ff"
	ExtractAsyncParamsLocaleFfSn      ExtractAsyncParamsLocale = "ff-SN"
	ExtractAsyncParamsLocaleFi        ExtractAsyncParamsLocale = "fi"
	ExtractAsyncParamsLocaleFiFi      ExtractAsyncParamsLocale = "fi-FI"
	ExtractAsyncParamsLocaleFil       ExtractAsyncParamsLocale = "fil"
	ExtractAsyncParamsLocaleFilPh     ExtractAsyncParamsLocale = "fil-PH"
	ExtractAsyncParamsLocaleFo        ExtractAsyncParamsLocale = "fo"
	ExtractAsyncParamsLocaleFoFo      ExtractAsyncParamsLocale = "fo-FO"
	ExtractAsyncParamsLocaleFr        ExtractAsyncParamsLocale = "fr"
	ExtractAsyncParamsLocaleFrBe      ExtractAsyncParamsLocale = "fr-BE"
	ExtractAsyncParamsLocaleFrBf      ExtractAsyncParamsLocale = "fr-BF"
	ExtractAsyncParamsLocaleFrBi      ExtractAsyncParamsLocale = "fr-BI"
	ExtractAsyncParamsLocaleFrBj      ExtractAsyncParamsLocale = "fr-BJ"
	ExtractAsyncParamsLocaleFrBl      ExtractAsyncParamsLocale = "fr-BL"
	ExtractAsyncParamsLocaleFrCa      ExtractAsyncParamsLocale = "fr-CA"
	ExtractAsyncParamsLocaleFrCd      ExtractAsyncParamsLocale = "fr-CD"
	ExtractAsyncParamsLocaleFrCf      ExtractAsyncParamsLocale = "fr-CF"
	ExtractAsyncParamsLocaleFrCg      ExtractAsyncParamsLocale = "fr-CG"
	ExtractAsyncParamsLocaleFrCh      ExtractAsyncParamsLocale = "fr-CH"
	ExtractAsyncParamsLocaleFrCi      ExtractAsyncParamsLocale = "fr-CI"
	ExtractAsyncParamsLocaleFrCm      ExtractAsyncParamsLocale = "fr-CM"
	ExtractAsyncParamsLocaleFrDj      ExtractAsyncParamsLocale = "fr-DJ"
	ExtractAsyncParamsLocaleFrFr      ExtractAsyncParamsLocale = "fr-FR"
	ExtractAsyncParamsLocaleFrGa      ExtractAsyncParamsLocale = "fr-GA"
	ExtractAsyncParamsLocaleFrGn      ExtractAsyncParamsLocale = "fr-GN"
	ExtractAsyncParamsLocaleFrGp      ExtractAsyncParamsLocale = "fr-GP"
	ExtractAsyncParamsLocaleFrGq      ExtractAsyncParamsLocale = "fr-GQ"
	ExtractAsyncParamsLocaleFrKm      ExtractAsyncParamsLocale = "fr-KM"
	ExtractAsyncParamsLocaleFrLu      ExtractAsyncParamsLocale = "fr-LU"
	ExtractAsyncParamsLocaleFrMc      ExtractAsyncParamsLocale = "fr-MC"
	ExtractAsyncParamsLocaleFrMf      ExtractAsyncParamsLocale = "fr-MF"
	ExtractAsyncParamsLocaleFrMg      ExtractAsyncParamsLocale = "fr-MG"
	ExtractAsyncParamsLocaleFrMl      ExtractAsyncParamsLocale = "fr-ML"
	ExtractAsyncParamsLocaleFrMq      ExtractAsyncParamsLocale = "fr-MQ"
	ExtractAsyncParamsLocaleFrNe      ExtractAsyncParamsLocale = "fr-NE"
	ExtractAsyncParamsLocaleFrRe      ExtractAsyncParamsLocale = "fr-RE"
	ExtractAsyncParamsLocaleFrRw      ExtractAsyncParamsLocale = "fr-RW"
	ExtractAsyncParamsLocaleFrSn      ExtractAsyncParamsLocale = "fr-SN"
	ExtractAsyncParamsLocaleFrTd      ExtractAsyncParamsLocale = "fr-TD"
	ExtractAsyncParamsLocaleFrTg      ExtractAsyncParamsLocale = "fr-TG"
	ExtractAsyncParamsLocaleFurIt     ExtractAsyncParamsLocale = "fur-IT"
	ExtractAsyncParamsLocaleFyDe      ExtractAsyncParamsLocale = "fy-DE"
	ExtractAsyncParamsLocaleFyNl      ExtractAsyncParamsLocale = "fy-NL"
	ExtractAsyncParamsLocaleGa        ExtractAsyncParamsLocale = "ga"
	ExtractAsyncParamsLocaleGaIe      ExtractAsyncParamsLocale = "ga-IE"
	ExtractAsyncParamsLocaleGdGB      ExtractAsyncParamsLocale = "gd-GB"
	ExtractAsyncParamsLocaleGezEr     ExtractAsyncParamsLocale = "gez-ER"
	ExtractAsyncParamsLocaleGezEt     ExtractAsyncParamsLocale = "gez-ET"
	ExtractAsyncParamsLocaleGl        ExtractAsyncParamsLocale = "gl"
	ExtractAsyncParamsLocaleGlEs      ExtractAsyncParamsLocale = "gl-ES"
	ExtractAsyncParamsLocaleGsw       ExtractAsyncParamsLocale = "gsw"
	ExtractAsyncParamsLocaleGswCh     ExtractAsyncParamsLocale = "gsw-CH"
	ExtractAsyncParamsLocaleGu        ExtractAsyncParamsLocale = "gu"
	ExtractAsyncParamsLocaleGuIn      ExtractAsyncParamsLocale = "gu-IN"
	ExtractAsyncParamsLocaleGuz       ExtractAsyncParamsLocale = "guz"
	ExtractAsyncParamsLocaleGuzKe     ExtractAsyncParamsLocale = "guz-KE"
	ExtractAsyncParamsLocaleGv        ExtractAsyncParamsLocale = "gv"
	ExtractAsyncParamsLocaleGvGB      ExtractAsyncParamsLocale = "gv-GB"
	ExtractAsyncParamsLocaleHa        ExtractAsyncParamsLocale = "ha"
	ExtractAsyncParamsLocaleHaLatn    ExtractAsyncParamsLocale = "ha-Latn"
	ExtractAsyncParamsLocaleHaLatnGh  ExtractAsyncParamsLocale = "ha-Latn-GH"
	ExtractAsyncParamsLocaleHaLatnNe  ExtractAsyncParamsLocale = "ha-Latn-NE"
	ExtractAsyncParamsLocaleHaLatnNg  ExtractAsyncParamsLocale = "ha-Latn-NG"
	ExtractAsyncParamsLocaleHaNg      ExtractAsyncParamsLocale = "ha-NG"
	ExtractAsyncParamsLocaleHaw       ExtractAsyncParamsLocale = "haw"
	ExtractAsyncParamsLocaleHawUs     ExtractAsyncParamsLocale = "haw-US"
	ExtractAsyncParamsLocaleHe        ExtractAsyncParamsLocale = "he"
	ExtractAsyncParamsLocaleHeIl      ExtractAsyncParamsLocale = "he-IL"
	ExtractAsyncParamsLocaleHi        ExtractAsyncParamsLocale = "hi"
	ExtractAsyncParamsLocaleHiIn      ExtractAsyncParamsLocale = "hi-IN"
	ExtractAsyncParamsLocaleHneIn     ExtractAsyncParamsLocale = "hne-IN"
	ExtractAsyncParamsLocaleHr        ExtractAsyncParamsLocale = "hr"
	ExtractAsyncParamsLocaleHrHr      ExtractAsyncParamsLocale = "hr-HR"
	ExtractAsyncParamsLocaleHsbDe     ExtractAsyncParamsLocale = "hsb-DE"
	ExtractAsyncParamsLocaleHtHt      ExtractAsyncParamsLocale = "ht-HT"
	ExtractAsyncParamsLocaleHu        ExtractAsyncParamsLocale = "hu"
	ExtractAsyncParamsLocaleHuHu      ExtractAsyncParamsLocale = "hu-HU"
	ExtractAsyncParamsLocaleHy        ExtractAsyncParamsLocale = "hy"
	ExtractAsyncParamsLocaleHyAm      ExtractAsyncParamsLocale = "hy-AM"
	ExtractAsyncParamsLocaleID        ExtractAsyncParamsLocale = "id"
	ExtractAsyncParamsLocaleIDID      ExtractAsyncParamsLocale = "id-ID"
	ExtractAsyncParamsLocaleIg        ExtractAsyncParamsLocale = "ig"
	ExtractAsyncParamsLocaleIgNg      ExtractAsyncParamsLocale = "ig-NG"
	ExtractAsyncParamsLocaleIi        ExtractAsyncParamsLocale = "ii"
	ExtractAsyncParamsLocaleIiCn      ExtractAsyncParamsLocale = "ii-CN"
	ExtractAsyncParamsLocaleIkCa      ExtractAsyncParamsLocale = "ik-CA"
	ExtractAsyncParamsLocaleIs        ExtractAsyncParamsLocale = "is"
	ExtractAsyncParamsLocaleIsIs      ExtractAsyncParamsLocale = "is-IS"
	ExtractAsyncParamsLocaleIt        ExtractAsyncParamsLocale = "it"
	ExtractAsyncParamsLocaleItCh      ExtractAsyncParamsLocale = "it-CH"
	ExtractAsyncParamsLocaleItIt      ExtractAsyncParamsLocale = "it-IT"
	ExtractAsyncParamsLocaleIuCa      ExtractAsyncParamsLocale = "iu-CA"
	ExtractAsyncParamsLocaleIwIl      ExtractAsyncParamsLocale = "iw-IL"
	ExtractAsyncParamsLocaleJa        ExtractAsyncParamsLocale = "ja"
	ExtractAsyncParamsLocaleJaJp      ExtractAsyncParamsLocale = "ja-JP"
	ExtractAsyncParamsLocaleJmc       ExtractAsyncParamsLocale = "jmc"
	ExtractAsyncParamsLocaleJmcTz     ExtractAsyncParamsLocale = "jmc-TZ"
	ExtractAsyncParamsLocaleKa        ExtractAsyncParamsLocale = "ka"
	ExtractAsyncParamsLocaleKaGe      ExtractAsyncParamsLocale = "ka-GE"
	ExtractAsyncParamsLocaleKab       ExtractAsyncParamsLocale = "kab"
	ExtractAsyncParamsLocaleKabDz     ExtractAsyncParamsLocale = "kab-DZ"
	ExtractAsyncParamsLocaleKam       ExtractAsyncParamsLocale = "kam"
	ExtractAsyncParamsLocaleKamKe     ExtractAsyncParamsLocale = "kam-KE"
	ExtractAsyncParamsLocaleKde       ExtractAsyncParamsLocale = "kde"
	ExtractAsyncParamsLocaleKdeTz     ExtractAsyncParamsLocale = "kde-TZ"
	ExtractAsyncParamsLocaleKea       ExtractAsyncParamsLocale = "kea"
	ExtractAsyncParamsLocaleKeaCv     ExtractAsyncParamsLocale = "kea-CV"
	ExtractAsyncParamsLocaleKhq       ExtractAsyncParamsLocale = "khq"
	ExtractAsyncParamsLocaleKhqMl     ExtractAsyncParamsLocale = "khq-ML"
	ExtractAsyncParamsLocaleKi        ExtractAsyncParamsLocale = "ki"
	ExtractAsyncParamsLocaleKiKe      ExtractAsyncParamsLocale = "ki-KE"
	ExtractAsyncParamsLocaleKk        ExtractAsyncParamsLocale = "kk"
	ExtractAsyncParamsLocaleKkCyrl    ExtractAsyncParamsLocale = "kk-Cyrl"
	ExtractAsyncParamsLocaleKkCyrlKz  ExtractAsyncParamsLocale = "kk-Cyrl-KZ"
	ExtractAsyncParamsLocaleKkKz      ExtractAsyncParamsLocale = "kk-KZ"
	ExtractAsyncParamsLocaleKl        ExtractAsyncParamsLocale = "kl"
	ExtractAsyncParamsLocaleKlGl      ExtractAsyncParamsLocale = "kl-GL"
	ExtractAsyncParamsLocaleKln       ExtractAsyncParamsLocale = "kln"
	ExtractAsyncParamsLocaleKlnKe     ExtractAsyncParamsLocale = "kln-KE"
	ExtractAsyncParamsLocaleKm        ExtractAsyncParamsLocale = "km"
	ExtractAsyncParamsLocaleKmKh      ExtractAsyncParamsLocale = "km-KH"
	ExtractAsyncParamsLocaleKn        ExtractAsyncParamsLocale = "kn"
	ExtractAsyncParamsLocaleKnIn      ExtractAsyncParamsLocale = "kn-IN"
	ExtractAsyncParamsLocaleKo        ExtractAsyncParamsLocale = "ko"
	ExtractAsyncParamsLocaleKoKr      ExtractAsyncParamsLocale = "ko-KR"
	ExtractAsyncParamsLocaleKok       ExtractAsyncParamsLocale = "kok"
	ExtractAsyncParamsLocaleKokIn     ExtractAsyncParamsLocale = "kok-IN"
	ExtractAsyncParamsLocaleKsIn      ExtractAsyncParamsLocale = "ks-IN"
	ExtractAsyncParamsLocaleKuTr      ExtractAsyncParamsLocale = "ku-TR"
	ExtractAsyncParamsLocaleKw        ExtractAsyncParamsLocale = "kw"
	ExtractAsyncParamsLocaleKwGB      ExtractAsyncParamsLocale = "kw-GB"
	ExtractAsyncParamsLocaleKyKg      ExtractAsyncParamsLocale = "ky-KG"
	ExtractAsyncParamsLocaleLag       ExtractAsyncParamsLocale = "lag"
	ExtractAsyncParamsLocaleLagTz     ExtractAsyncParamsLocale = "lag-TZ"
	ExtractAsyncParamsLocaleLbLu      ExtractAsyncParamsLocale = "lb-LU"
	ExtractAsyncParamsLocaleLg        ExtractAsyncParamsLocale = "lg"
	ExtractAsyncParamsLocaleLgUg      ExtractAsyncParamsLocale = "lg-UG"
	ExtractAsyncParamsLocaleLiBe      ExtractAsyncParamsLocale = "li-BE"
	ExtractAsyncParamsLocaleLiNl      ExtractAsyncParamsLocale = "li-NL"
	ExtractAsyncParamsLocaleLijIt     ExtractAsyncParamsLocale = "lij-IT"
	ExtractAsyncParamsLocaleLoLa      ExtractAsyncParamsLocale = "lo-LA"
	ExtractAsyncParamsLocaleLt        ExtractAsyncParamsLocale = "lt"
	ExtractAsyncParamsLocaleLtLt      ExtractAsyncParamsLocale = "lt-LT"
	ExtractAsyncParamsLocaleLuo       ExtractAsyncParamsLocale = "luo"
	ExtractAsyncParamsLocaleLuoKe     ExtractAsyncParamsLocale = "luo-KE"
	ExtractAsyncParamsLocaleLuy       ExtractAsyncParamsLocale = "luy"
	ExtractAsyncParamsLocaleLuyKe     ExtractAsyncParamsLocale = "luy-KE"
	ExtractAsyncParamsLocaleLv        ExtractAsyncParamsLocale = "lv"
	ExtractAsyncParamsLocaleLvLv      ExtractAsyncParamsLocale = "lv-LV"
	ExtractAsyncParamsLocaleMagIn     ExtractAsyncParamsLocale = "mag-IN"
	ExtractAsyncParamsLocaleMaiIn     ExtractAsyncParamsLocale = "mai-IN"
	ExtractAsyncParamsLocaleMas       ExtractAsyncParamsLocale = "mas"
	ExtractAsyncParamsLocaleMasKe     ExtractAsyncParamsLocale = "mas-KE"
	ExtractAsyncParamsLocaleMasTz     ExtractAsyncParamsLocale = "mas-TZ"
	ExtractAsyncParamsLocaleMer       ExtractAsyncParamsLocale = "mer"
	ExtractAsyncParamsLocaleMerKe     ExtractAsyncParamsLocale = "mer-KE"
	ExtractAsyncParamsLocaleMfe       ExtractAsyncParamsLocale = "mfe"
	ExtractAsyncParamsLocaleMfeMu     ExtractAsyncParamsLocale = "mfe-MU"
	ExtractAsyncParamsLocaleMg        ExtractAsyncParamsLocale = "mg"
	ExtractAsyncParamsLocaleMgMg      ExtractAsyncParamsLocale = "mg-MG"
	ExtractAsyncParamsLocaleMhrRu     ExtractAsyncParamsLocale = "mhr-RU"
	ExtractAsyncParamsLocaleMiNz      ExtractAsyncParamsLocale = "mi-NZ"
	ExtractAsyncParamsLocaleMk        ExtractAsyncParamsLocale = "mk"
	ExtractAsyncParamsLocaleMkMk      ExtractAsyncParamsLocale = "mk-MK"
	ExtractAsyncParamsLocaleMl        ExtractAsyncParamsLocale = "ml"
	ExtractAsyncParamsLocaleMlIn      ExtractAsyncParamsLocale = "ml-IN"
	ExtractAsyncParamsLocaleMnMn      ExtractAsyncParamsLocale = "mn-MN"
	ExtractAsyncParamsLocaleMr        ExtractAsyncParamsLocale = "mr"
	ExtractAsyncParamsLocaleMrIn      ExtractAsyncParamsLocale = "mr-IN"
	ExtractAsyncParamsLocaleMs        ExtractAsyncParamsLocale = "ms"
	ExtractAsyncParamsLocaleMsBn      ExtractAsyncParamsLocale = "ms-BN"
	ExtractAsyncParamsLocaleMsMy      ExtractAsyncParamsLocale = "ms-MY"
	ExtractAsyncParamsLocaleMt        ExtractAsyncParamsLocale = "mt"
	ExtractAsyncParamsLocaleMtMt      ExtractAsyncParamsLocale = "mt-MT"
	ExtractAsyncParamsLocaleMy        ExtractAsyncParamsLocale = "my"
	ExtractAsyncParamsLocaleMyMm      ExtractAsyncParamsLocale = "my-MM"
	ExtractAsyncParamsLocaleNanTw     ExtractAsyncParamsLocale = "nan-TW"
	ExtractAsyncParamsLocaleNaq       ExtractAsyncParamsLocale = "naq"
	ExtractAsyncParamsLocaleNaqNa     ExtractAsyncParamsLocale = "naq-NA"
	ExtractAsyncParamsLocaleNb        ExtractAsyncParamsLocale = "nb"
	ExtractAsyncParamsLocaleNbNo      ExtractAsyncParamsLocale = "nb-NO"
	ExtractAsyncParamsLocaleNd        ExtractAsyncParamsLocale = "nd"
	ExtractAsyncParamsLocaleNdZw      ExtractAsyncParamsLocale = "nd-ZW"
	ExtractAsyncParamsLocaleNdsDe     ExtractAsyncParamsLocale = "nds-DE"
	ExtractAsyncParamsLocaleNdsNl     ExtractAsyncParamsLocale = "nds-NL"
	ExtractAsyncParamsLocaleNe        ExtractAsyncParamsLocale = "ne"
	ExtractAsyncParamsLocaleNeIn      ExtractAsyncParamsLocale = "ne-IN"
	ExtractAsyncParamsLocaleNeNp      ExtractAsyncParamsLocale = "ne-NP"
	ExtractAsyncParamsLocaleNl        ExtractAsyncParamsLocale = "nl"
	ExtractAsyncParamsLocaleNlAw      ExtractAsyncParamsLocale = "nl-AW"
	ExtractAsyncParamsLocaleNlBe      ExtractAsyncParamsLocale = "nl-BE"
	ExtractAsyncParamsLocaleNlNl      ExtractAsyncParamsLocale = "nl-NL"
	ExtractAsyncParamsLocaleNn        ExtractAsyncParamsLocale = "nn"
	ExtractAsyncParamsLocaleNnNo      ExtractAsyncParamsLocale = "nn-NO"
	ExtractAsyncParamsLocaleNrZa      ExtractAsyncParamsLocale = "nr-ZA"
	ExtractAsyncParamsLocaleNsoZa     ExtractAsyncParamsLocale = "nso-ZA"
	ExtractAsyncParamsLocaleNyn       ExtractAsyncParamsLocale = "nyn"
	ExtractAsyncParamsLocaleNynUg     ExtractAsyncParamsLocale = "nyn-UG"
	ExtractAsyncParamsLocaleOcFr      ExtractAsyncParamsLocale = "oc-FR"
	ExtractAsyncParamsLocaleOm        ExtractAsyncParamsLocale = "om"
	ExtractAsyncParamsLocaleOmEt      ExtractAsyncParamsLocale = "om-ET"
	ExtractAsyncParamsLocaleOmKe      ExtractAsyncParamsLocale = "om-KE"
	ExtractAsyncParamsLocaleOr        ExtractAsyncParamsLocale = "or"
	ExtractAsyncParamsLocaleOrIn      ExtractAsyncParamsLocale = "or-IN"
	ExtractAsyncParamsLocaleOsRu      ExtractAsyncParamsLocale = "os-RU"
	ExtractAsyncParamsLocalePa        ExtractAsyncParamsLocale = "pa"
	ExtractAsyncParamsLocalePaArab    ExtractAsyncParamsLocale = "pa-Arab"
	ExtractAsyncParamsLocalePaArabPk  ExtractAsyncParamsLocale = "pa-Arab-PK"
	ExtractAsyncParamsLocalePaGuru    ExtractAsyncParamsLocale = "pa-Guru"
	ExtractAsyncParamsLocalePaGuruIn  ExtractAsyncParamsLocale = "pa-Guru-IN"
	ExtractAsyncParamsLocalePaIn      ExtractAsyncParamsLocale = "pa-IN"
	ExtractAsyncParamsLocalePaPk      ExtractAsyncParamsLocale = "pa-PK"
	ExtractAsyncParamsLocalePapAn     ExtractAsyncParamsLocale = "pap-AN"
	ExtractAsyncParamsLocalePl        ExtractAsyncParamsLocale = "pl"
	ExtractAsyncParamsLocalePlPl      ExtractAsyncParamsLocale = "pl-PL"
	ExtractAsyncParamsLocalePs        ExtractAsyncParamsLocale = "ps"
	ExtractAsyncParamsLocalePsAf      ExtractAsyncParamsLocale = "ps-AF"
	ExtractAsyncParamsLocalePt        ExtractAsyncParamsLocale = "pt"
	ExtractAsyncParamsLocalePtBr      ExtractAsyncParamsLocale = "pt-BR"
	ExtractAsyncParamsLocalePtGw      ExtractAsyncParamsLocale = "pt-GW"
	ExtractAsyncParamsLocalePtMz      ExtractAsyncParamsLocale = "pt-MZ"
	ExtractAsyncParamsLocalePtPt      ExtractAsyncParamsLocale = "pt-PT"
	ExtractAsyncParamsLocaleRm        ExtractAsyncParamsLocale = "rm"
	ExtractAsyncParamsLocaleRmCh      ExtractAsyncParamsLocale = "rm-CH"
	ExtractAsyncParamsLocaleRo        ExtractAsyncParamsLocale = "ro"
	ExtractAsyncParamsLocaleRoMd      ExtractAsyncParamsLocale = "ro-MD"
	ExtractAsyncParamsLocaleRoRo      ExtractAsyncParamsLocale = "ro-RO"
	ExtractAsyncParamsLocaleRof       ExtractAsyncParamsLocale = "rof"
	ExtractAsyncParamsLocaleRofTz     ExtractAsyncParamsLocale = "rof-TZ"
	ExtractAsyncParamsLocaleRu        ExtractAsyncParamsLocale = "ru"
	ExtractAsyncParamsLocaleRuMd      ExtractAsyncParamsLocale = "ru-MD"
	ExtractAsyncParamsLocaleRuRu      ExtractAsyncParamsLocale = "ru-RU"
	ExtractAsyncParamsLocaleRuUa      ExtractAsyncParamsLocale = "ru-UA"
	ExtractAsyncParamsLocaleRw        ExtractAsyncParamsLocale = "rw"
	ExtractAsyncParamsLocaleRwRw      ExtractAsyncParamsLocale = "rw-RW"
	ExtractAsyncParamsLocaleRwk       ExtractAsyncParamsLocale = "rwk"
	ExtractAsyncParamsLocaleRwkTz     ExtractAsyncParamsLocale = "rwk-TZ"
	ExtractAsyncParamsLocaleSaIn      ExtractAsyncParamsLocale = "sa-IN"
	ExtractAsyncParamsLocaleSaq       ExtractAsyncParamsLocale = "saq"
	ExtractAsyncParamsLocaleSaqKe     ExtractAsyncParamsLocale = "saq-KE"
	ExtractAsyncParamsLocaleScIt      ExtractAsyncParamsLocale = "sc-IT"
	ExtractAsyncParamsLocaleSdIn      ExtractAsyncParamsLocale = "sd-IN"
	ExtractAsyncParamsLocaleSeNo      ExtractAsyncParamsLocale = "se-NO"
	ExtractAsyncParamsLocaleSeh       ExtractAsyncParamsLocale = "seh"
	ExtractAsyncParamsLocaleSehMz     ExtractAsyncParamsLocale = "seh-MZ"
	ExtractAsyncParamsLocaleSes       ExtractAsyncParamsLocale = "ses"
	ExtractAsyncParamsLocaleSesMl     ExtractAsyncParamsLocale = "ses-ML"
	ExtractAsyncParamsLocaleSg        ExtractAsyncParamsLocale = "sg"
	ExtractAsyncParamsLocaleSgCf      ExtractAsyncParamsLocale = "sg-CF"
	ExtractAsyncParamsLocaleShi       ExtractAsyncParamsLocale = "shi"
	ExtractAsyncParamsLocaleShiLatn   ExtractAsyncParamsLocale = "shi-Latn"
	ExtractAsyncParamsLocaleShiLatnMa ExtractAsyncParamsLocale = "shi-Latn-MA"
	ExtractAsyncParamsLocaleShiTfng   ExtractAsyncParamsLocale = "shi-Tfng"
	ExtractAsyncParamsLocaleShiTfngMa ExtractAsyncParamsLocale = "shi-Tfng-MA"
	ExtractAsyncParamsLocaleShsCa     ExtractAsyncParamsLocale = "shs-CA"
	ExtractAsyncParamsLocaleSi        ExtractAsyncParamsLocale = "si"
	ExtractAsyncParamsLocaleSiLk      ExtractAsyncParamsLocale = "si-LK"
	ExtractAsyncParamsLocaleSidEt     ExtractAsyncParamsLocale = "sid-ET"
	ExtractAsyncParamsLocaleSk        ExtractAsyncParamsLocale = "sk"
	ExtractAsyncParamsLocaleSkSk      ExtractAsyncParamsLocale = "sk-SK"
	ExtractAsyncParamsLocaleSl        ExtractAsyncParamsLocale = "sl"
	ExtractAsyncParamsLocaleSlSi      ExtractAsyncParamsLocale = "sl-SI"
	ExtractAsyncParamsLocaleSn        ExtractAsyncParamsLocale = "sn"
	ExtractAsyncParamsLocaleSnZw      ExtractAsyncParamsLocale = "sn-ZW"
	ExtractAsyncParamsLocaleSo        ExtractAsyncParamsLocale = "so"
	ExtractAsyncParamsLocaleSoDj      ExtractAsyncParamsLocale = "so-DJ"
	ExtractAsyncParamsLocaleSoEt      ExtractAsyncParamsLocale = "so-ET"
	ExtractAsyncParamsLocaleSoKe      ExtractAsyncParamsLocale = "so-KE"
	ExtractAsyncParamsLocaleSoSo      ExtractAsyncParamsLocale = "so-SO"
	ExtractAsyncParamsLocaleSq        ExtractAsyncParamsLocale = "sq"
	ExtractAsyncParamsLocaleSqAl      ExtractAsyncParamsLocale = "sq-AL"
	ExtractAsyncParamsLocaleSqMk      ExtractAsyncParamsLocale = "sq-MK"
	ExtractAsyncParamsLocaleSr        ExtractAsyncParamsLocale = "sr"
	ExtractAsyncParamsLocaleSrCyrl    ExtractAsyncParamsLocale = "sr-Cyrl"
	ExtractAsyncParamsLocaleSrCyrlBa  ExtractAsyncParamsLocale = "sr-Cyrl-BA"
	ExtractAsyncParamsLocaleSrCyrlMe  ExtractAsyncParamsLocale = "sr-Cyrl-ME"
	ExtractAsyncParamsLocaleSrCyrlRs  ExtractAsyncParamsLocale = "sr-Cyrl-RS"
	ExtractAsyncParamsLocaleSrLatn    ExtractAsyncParamsLocale = "sr-Latn"
	ExtractAsyncParamsLocaleSrLatnBa  ExtractAsyncParamsLocale = "sr-Latn-BA"
	ExtractAsyncParamsLocaleSrLatnMe  ExtractAsyncParamsLocale = "sr-Latn-ME"
	ExtractAsyncParamsLocaleSrLatnRs  ExtractAsyncParamsLocale = "sr-Latn-RS"
	ExtractAsyncParamsLocaleSrMe      ExtractAsyncParamsLocale = "sr-ME"
	ExtractAsyncParamsLocaleSrRs      ExtractAsyncParamsLocale = "sr-RS"
	ExtractAsyncParamsLocaleSSZa      ExtractAsyncParamsLocale = "ss-ZA"
	ExtractAsyncParamsLocaleStZa      ExtractAsyncParamsLocale = "st-ZA"
	ExtractAsyncParamsLocaleSv        ExtractAsyncParamsLocale = "sv"
	ExtractAsyncParamsLocaleSvFi      ExtractAsyncParamsLocale = "sv-FI"
	ExtractAsyncParamsLocaleSvSe      ExtractAsyncParamsLocale = "sv-SE"
	ExtractAsyncParamsLocaleSw        ExtractAsyncParamsLocale = "sw"
	ExtractAsyncParamsLocaleSwKe      ExtractAsyncParamsLocale = "sw-KE"
	ExtractAsyncParamsLocaleSwTz      ExtractAsyncParamsLocale = "sw-TZ"
	ExtractAsyncParamsLocaleTa        ExtractAsyncParamsLocale = "ta"
	ExtractAsyncParamsLocaleTaIn      ExtractAsyncParamsLocale = "ta-IN"
	ExtractAsyncParamsLocaleTaLk      ExtractAsyncParamsLocale = "ta-LK"
	ExtractAsyncParamsLocaleTe        ExtractAsyncParamsLocale = "te"
	ExtractAsyncParamsLocaleTeIn      ExtractAsyncParamsLocale = "te-IN"
	ExtractAsyncParamsLocaleTeo       ExtractAsyncParamsLocale = "teo"
	ExtractAsyncParamsLocaleTeoKe     ExtractAsyncParamsLocale = "teo-KE"
	ExtractAsyncParamsLocaleTeoUg     ExtractAsyncParamsLocale = "teo-UG"
	ExtractAsyncParamsLocaleTgTj      ExtractAsyncParamsLocale = "tg-TJ"
	ExtractAsyncParamsLocaleTh        ExtractAsyncParamsLocale = "th"
	ExtractAsyncParamsLocaleThTh      ExtractAsyncParamsLocale = "th-TH"
	ExtractAsyncParamsLocaleTi        ExtractAsyncParamsLocale = "ti"
	ExtractAsyncParamsLocaleTiEr      ExtractAsyncParamsLocale = "ti-ER"
	ExtractAsyncParamsLocaleTiEt      ExtractAsyncParamsLocale = "ti-ET"
	ExtractAsyncParamsLocaleTigEr     ExtractAsyncParamsLocale = "tig-ER"
	ExtractAsyncParamsLocaleTkTm      ExtractAsyncParamsLocale = "tk-TM"
	ExtractAsyncParamsLocaleTlPh      ExtractAsyncParamsLocale = "tl-PH"
	ExtractAsyncParamsLocaleTnZa      ExtractAsyncParamsLocale = "tn-ZA"
	ExtractAsyncParamsLocaleTo        ExtractAsyncParamsLocale = "to"
	ExtractAsyncParamsLocaleToTo      ExtractAsyncParamsLocale = "to-TO"
	ExtractAsyncParamsLocaleTr        ExtractAsyncParamsLocale = "tr"
	ExtractAsyncParamsLocaleTrCy      ExtractAsyncParamsLocale = "tr-CY"
	ExtractAsyncParamsLocaleTrTr      ExtractAsyncParamsLocale = "tr-TR"
	ExtractAsyncParamsLocaleTsZa      ExtractAsyncParamsLocale = "ts-ZA"
	ExtractAsyncParamsLocaleTtRu      ExtractAsyncParamsLocale = "tt-RU"
	ExtractAsyncParamsLocaleTzm       ExtractAsyncParamsLocale = "tzm"
	ExtractAsyncParamsLocaleTzmLatn   ExtractAsyncParamsLocale = "tzm-Latn"
	ExtractAsyncParamsLocaleTzmLatnMa ExtractAsyncParamsLocale = "tzm-Latn-MA"
	ExtractAsyncParamsLocaleUgCn      ExtractAsyncParamsLocale = "ug-CN"
	ExtractAsyncParamsLocaleUk        ExtractAsyncParamsLocale = "uk"
	ExtractAsyncParamsLocaleUkUa      ExtractAsyncParamsLocale = "uk-UA"
	ExtractAsyncParamsLocaleUnmUs     ExtractAsyncParamsLocale = "unm-US"
	ExtractAsyncParamsLocaleUr        ExtractAsyncParamsLocale = "ur"
	ExtractAsyncParamsLocaleUrIn      ExtractAsyncParamsLocale = "ur-IN"
	ExtractAsyncParamsLocaleUrPk      ExtractAsyncParamsLocale = "ur-PK"
	ExtractAsyncParamsLocaleUz        ExtractAsyncParamsLocale = "uz"
	ExtractAsyncParamsLocaleUzArab    ExtractAsyncParamsLocale = "uz-Arab"
	ExtractAsyncParamsLocaleUzArabAf  ExtractAsyncParamsLocale = "uz-Arab-AF"
	ExtractAsyncParamsLocaleUzCyrl    ExtractAsyncParamsLocale = "uz-Cyrl"
	ExtractAsyncParamsLocaleUzCyrlUz  ExtractAsyncParamsLocale = "uz-Cyrl-UZ"
	ExtractAsyncParamsLocaleUzLatn    ExtractAsyncParamsLocale = "uz-Latn"
	ExtractAsyncParamsLocaleUzLatnUz  ExtractAsyncParamsLocale = "uz-Latn-UZ"
	ExtractAsyncParamsLocaleUzUz      ExtractAsyncParamsLocale = "uz-UZ"
	ExtractAsyncParamsLocaleVeZa      ExtractAsyncParamsLocale = "ve-ZA"
	ExtractAsyncParamsLocaleVi        ExtractAsyncParamsLocale = "vi"
	ExtractAsyncParamsLocaleViVn      ExtractAsyncParamsLocale = "vi-VN"
	ExtractAsyncParamsLocaleVun       ExtractAsyncParamsLocale = "vun"
	ExtractAsyncParamsLocaleVunTz     ExtractAsyncParamsLocale = "vun-TZ"
	ExtractAsyncParamsLocaleWaBe      ExtractAsyncParamsLocale = "wa-BE"
	ExtractAsyncParamsLocaleWaeCh     ExtractAsyncParamsLocale = "wae-CH"
	ExtractAsyncParamsLocaleWalEt     ExtractAsyncParamsLocale = "wal-ET"
	ExtractAsyncParamsLocaleWoSn      ExtractAsyncParamsLocale = "wo-SN"
	ExtractAsyncParamsLocaleXhZa      ExtractAsyncParamsLocale = "xh-ZA"
	ExtractAsyncParamsLocaleXog       ExtractAsyncParamsLocale = "xog"
	ExtractAsyncParamsLocaleXogUg     ExtractAsyncParamsLocale = "xog-UG"
	ExtractAsyncParamsLocaleYiUs      ExtractAsyncParamsLocale = "yi-US"
	ExtractAsyncParamsLocaleYo        ExtractAsyncParamsLocale = "yo"
	ExtractAsyncParamsLocaleYoNg      ExtractAsyncParamsLocale = "yo-NG"
	ExtractAsyncParamsLocaleYueHk     ExtractAsyncParamsLocale = "yue-HK"
	ExtractAsyncParamsLocaleZh        ExtractAsyncParamsLocale = "zh"
	ExtractAsyncParamsLocaleZhCn      ExtractAsyncParamsLocale = "zh-CN"
	ExtractAsyncParamsLocaleZhHk      ExtractAsyncParamsLocale = "zh-HK"
	ExtractAsyncParamsLocaleZhHans    ExtractAsyncParamsLocale = "zh-Hans"
	ExtractAsyncParamsLocaleZhHansCn  ExtractAsyncParamsLocale = "zh-Hans-CN"
	ExtractAsyncParamsLocaleZhHansHk  ExtractAsyncParamsLocale = "zh-Hans-HK"
	ExtractAsyncParamsLocaleZhHansMo  ExtractAsyncParamsLocale = "zh-Hans-MO"
	ExtractAsyncParamsLocaleZhHansSg  ExtractAsyncParamsLocale = "zh-Hans-SG"
	ExtractAsyncParamsLocaleZhHant    ExtractAsyncParamsLocale = "zh-Hant"
	ExtractAsyncParamsLocaleZhHantHk  ExtractAsyncParamsLocale = "zh-Hant-HK"
	ExtractAsyncParamsLocaleZhHantMo  ExtractAsyncParamsLocale = "zh-Hant-MO"
	ExtractAsyncParamsLocaleZhHantTw  ExtractAsyncParamsLocale = "zh-Hant-TW"
	ExtractAsyncParamsLocaleZhSg      ExtractAsyncParamsLocale = "zh-SG"
	ExtractAsyncParamsLocaleZhTw      ExtractAsyncParamsLocale = "zh-TW"
	ExtractAsyncParamsLocaleZu        ExtractAsyncParamsLocale = "zu"
	ExtractAsyncParamsLocaleZuZa      ExtractAsyncParamsLocale = "zu-ZA"
	ExtractAsyncParamsLocaleAuto      ExtractAsyncParamsLocale = "auto"
)

// HTTP method for the request
type ExtractAsyncParamsMethod string

const (
	ExtractAsyncParamsMethodGet    ExtractAsyncParamsMethod = "GET"
	ExtractAsyncParamsMethodPost   ExtractAsyncParamsMethod = "POST"
	ExtractAsyncParamsMethodPut    ExtractAsyncParamsMethod = "PUT"
	ExtractAsyncParamsMethodPatch  ExtractAsyncParamsMethod = "PATCH"
	ExtractAsyncParamsMethodDelete ExtractAsyncParamsMethod = "DELETE"
)

type ExtractAsyncParamsNetworkCapture struct {
	Validation                  param.Opt[bool]    `json:"validation,omitzero"`
	WaitForRequestsCount        param.Opt[float64] `json:"wait_for_requests_count,omitzero"`
	WaitForRequestsCountTimeout param.Opt[float64] `json:"wait_for_requests_count_timeout,omitzero"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method,omitzero"`
	// Resource type for network capture filtering
	ResourceType ExtractAsyncParamsNetworkCaptureResourceTypeUnion `json:"resource_type,omitzero"`
	StatusCode   ExtractAsyncParamsNetworkCaptureStatusCodeUnion   `json:"status_code,omitzero"`
	URL          ExtractAsyncParamsNetworkCaptureURL               `json:"url,omitzero"`
	paramObj
}

func (r ExtractAsyncParamsNetworkCapture) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsNetworkCapture
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractAsyncParamsNetworkCapture](
		"method", "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsNetworkCaptureResourceTypeUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsNetworkCaptureResourceTypeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractAsyncParamsNetworkCaptureResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsNetworkCaptureResourceTypeUnion) asAny() any {
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
type ExtractAsyncParamsNetworkCaptureStatusCodeUnion struct {
	OfFloat      param.Opt[float64] `json:",omitzero,inline"`
	OfFloatArray []float64          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsNetworkCaptureStatusCodeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfFloatArray)
}
func (u *ExtractAsyncParamsNetworkCaptureStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsNetworkCaptureStatusCodeUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfFloatArray) {
		return &u.OfFloatArray
	}
	return nil
}

// The property Value is required.
type ExtractAsyncParamsNetworkCaptureURL struct {
	Value string `json:"value" api:"required"`
	// Any of "exact", "contains".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ExtractAsyncParamsNetworkCaptureURL) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsNetworkCaptureURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsNetworkCaptureURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractAsyncParamsNetworkCaptureURL](
		"type", "exact", "contains",
	)
}

// Operating system to emulate
type ExtractAsyncParamsOs string

const (
	ExtractAsyncParamsOsWindows ExtractAsyncParamsOs = "windows"
	ExtractAsyncParamsOsMacOs   ExtractAsyncParamsOs = "mac os"
	ExtractAsyncParamsOsLinux   ExtractAsyncParamsOs = "linux"
	ExtractAsyncParamsOsAndroid ExtractAsyncParamsOs = "android"
	ExtractAsyncParamsOsIos     ExtractAsyncParamsOs = "ios"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsParserUnion struct {
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsParserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap, u.OfString)
}
func (u *ExtractAsyncParamsParserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsParserUnion) asAny() any {
	if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Referrer policy for the request
type ExtractAsyncParamsReferrerType string

const (
	ExtractAsyncParamsReferrerTypeRandom     ExtractAsyncParamsReferrerType = "random"
	ExtractAsyncParamsReferrerTypeNoReferer  ExtractAsyncParamsReferrerType = "no-referer"
	ExtractAsyncParamsReferrerTypeSameOrigin ExtractAsyncParamsReferrerType = "same-origin"
	ExtractAsyncParamsReferrerTypeGoogle     ExtractAsyncParamsReferrerType = "google"
	ExtractAsyncParamsReferrerTypeBing       ExtractAsyncParamsReferrerType = "bing"
	ExtractAsyncParamsReferrerTypeFacebook   ExtractAsyncParamsReferrerType = "facebook"
	ExtractAsyncParamsReferrerTypeTwitter    ExtractAsyncParamsReferrerType = "twitter"
	ExtractAsyncParamsReferrerTypeInstagram  ExtractAsyncParamsReferrerType = "instagram"
)

type ExtractAsyncParamsSession struct {
	ID                  param.Opt[string]  `json:"id,omitzero"`
	PrefetchUserbrowser param.Opt[bool]    `json:"prefetch_userbrowser,omitzero"`
	Retry               param.Opt[bool]    `json:"retry,omitzero"`
	Timeout             param.Opt[float64] `json:"timeout,omitzero"`
	paramObj
}

func (r ExtractAsyncParamsSession) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsSession
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsSkillUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsSkillUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractAsyncParamsSkillUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsSkillUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// US state for geolocation (only valid when country is US)
type ExtractAsyncParamsState string

const (
	ExtractAsyncParamsStateAl ExtractAsyncParamsState = "AL"
	ExtractAsyncParamsStateAk ExtractAsyncParamsState = "AK"
	ExtractAsyncParamsStateAs ExtractAsyncParamsState = "AS"
	ExtractAsyncParamsStateAz ExtractAsyncParamsState = "AZ"
	ExtractAsyncParamsStateAr ExtractAsyncParamsState = "AR"
	ExtractAsyncParamsStateCa ExtractAsyncParamsState = "CA"
	ExtractAsyncParamsStateCo ExtractAsyncParamsState = "CO"
	ExtractAsyncParamsStateCt ExtractAsyncParamsState = "CT"
	ExtractAsyncParamsStateDe ExtractAsyncParamsState = "DE"
	ExtractAsyncParamsStateDc ExtractAsyncParamsState = "DC"
	ExtractAsyncParamsStateFl ExtractAsyncParamsState = "FL"
	ExtractAsyncParamsStateGa ExtractAsyncParamsState = "GA"
	ExtractAsyncParamsStateGu ExtractAsyncParamsState = "GU"
	ExtractAsyncParamsStateHi ExtractAsyncParamsState = "HI"
	ExtractAsyncParamsStateID ExtractAsyncParamsState = "ID"
	ExtractAsyncParamsStateIl ExtractAsyncParamsState = "IL"
	ExtractAsyncParamsStateIn ExtractAsyncParamsState = "IN"
	ExtractAsyncParamsStateIa ExtractAsyncParamsState = "IA"
	ExtractAsyncParamsStateKs ExtractAsyncParamsState = "KS"
	ExtractAsyncParamsStateKy ExtractAsyncParamsState = "KY"
	ExtractAsyncParamsStateLa ExtractAsyncParamsState = "LA"
	ExtractAsyncParamsStateMe ExtractAsyncParamsState = "ME"
	ExtractAsyncParamsStateMd ExtractAsyncParamsState = "MD"
	ExtractAsyncParamsStateMa ExtractAsyncParamsState = "MA"
	ExtractAsyncParamsStateMi ExtractAsyncParamsState = "MI"
	ExtractAsyncParamsStateMn ExtractAsyncParamsState = "MN"
	ExtractAsyncParamsStateMs ExtractAsyncParamsState = "MS"
	ExtractAsyncParamsStateMo ExtractAsyncParamsState = "MO"
	ExtractAsyncParamsStateMt ExtractAsyncParamsState = "MT"
	ExtractAsyncParamsStateNe ExtractAsyncParamsState = "NE"
	ExtractAsyncParamsStateNv ExtractAsyncParamsState = "NV"
	ExtractAsyncParamsStateNh ExtractAsyncParamsState = "NH"
	ExtractAsyncParamsStateNj ExtractAsyncParamsState = "NJ"
	ExtractAsyncParamsStateNm ExtractAsyncParamsState = "NM"
	ExtractAsyncParamsStateNy ExtractAsyncParamsState = "NY"
	ExtractAsyncParamsStateNc ExtractAsyncParamsState = "NC"
	ExtractAsyncParamsStateNd ExtractAsyncParamsState = "ND"
	ExtractAsyncParamsStateMp ExtractAsyncParamsState = "MP"
	ExtractAsyncParamsStateOh ExtractAsyncParamsState = "OH"
	ExtractAsyncParamsStateOk ExtractAsyncParamsState = "OK"
	ExtractAsyncParamsStateOr ExtractAsyncParamsState = "OR"
	ExtractAsyncParamsStatePa ExtractAsyncParamsState = "PA"
	ExtractAsyncParamsStatePr ExtractAsyncParamsState = "PR"
	ExtractAsyncParamsStateRi ExtractAsyncParamsState = "RI"
	ExtractAsyncParamsStateSc ExtractAsyncParamsState = "SC"
	ExtractAsyncParamsStateSd ExtractAsyncParamsState = "SD"
	ExtractAsyncParamsStateTn ExtractAsyncParamsState = "TN"
	ExtractAsyncParamsStateTx ExtractAsyncParamsState = "TX"
	ExtractAsyncParamsStateUt ExtractAsyncParamsState = "UT"
	ExtractAsyncParamsStateVt ExtractAsyncParamsState = "VT"
	ExtractAsyncParamsStateVa ExtractAsyncParamsState = "VA"
	ExtractAsyncParamsStateVi ExtractAsyncParamsState = "VI"
	ExtractAsyncParamsStateWa ExtractAsyncParamsState = "WA"
	ExtractAsyncParamsStateWv ExtractAsyncParamsState = "WV"
	ExtractAsyncParamsStateWi ExtractAsyncParamsState = "WI"
	ExtractAsyncParamsStateWy ExtractAsyncParamsState = "WY"
)

type ExtractBatchParams struct {
	// Array of extraction requests. Each object can include extraction parameters and
	// async/storage settings.
	Inputs []ExtractBatchParamsInput `json:"inputs,omitzero" api:"required"`
	// Shared parameters applied to the entire batch. Can include extraction parameters
	// and async/storage settings.
	SharedInputs ExtractBatchParamsSharedInputs `json:"shared_inputs,omitzero"`
	paramObj
}

func (r ExtractBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractBatchParamsInput struct {
	// URL to call back when async operation completes
	CallbackURL param.Opt[string] `json:"callback_url,omitzero"`
	// City for geolocation
	City param.Opt[string] `json:"city,omitzero"`
	// Whether to automatically handle cookie consent headers
	ConsentHeader param.Opt[bool] `json:"consent_header,omitzero"`
	// Whether to use HTTP/2 protocol
	Http2 param.Opt[bool] `json:"http2,omitzero"`
	// Whether to emulate XMLHttpRequest behavior
	IsXhr param.Opt[bool] `json:"is_xhr,omitzero"`
	// Whether to parse the response content
	Parse param.Opt[bool] `json:"parse,omitzero"`
	// Whether to render JavaScript content using a browser
	Render param.Opt[bool] `json:"render,omitzero"`
	// Request timeout in milliseconds
	RequestTimeout param.Opt[float64] `json:"request_timeout,omitzero"`
	// Whether to compress stored data
	StorageCompress param.Opt[bool] `json:"storage_compress,omitzero"`
	// Custom name for the stored object
	StorageObjectName param.Opt[string] `json:"storage_object_name,omitzero"`
	// Type of storage to use for results
	StorageType param.Opt[string] `json:"storage_type,omitzero"`
	// URL for storage location
	StorageURL param.Opt[string] `json:"storage_url,omitzero"`
	// User-defined tag for request identification
	Tag param.Opt[string] `json:"tag,omitzero"`
	// Target URL to scrape
	URL param.Opt[string] `json:"url,omitzero"`
	// Browser type to emulate
	Browser ExtractBatchParamsInputBrowserUnion `json:"browser,omitzero"`
	// Array of browser automation actions to execute sequentially
	BrowserActions []ExtractBatchParamsInputBrowserActionUnion `json:"browser_actions,omitzero"`
	// Browser cookies as array of cookie objects
	Cookies ExtractBatchParamsInputCookiesUnion `json:"cookies,omitzero"`
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
	Country ExtractBatchParamsInputCountry `json:"country,omitzero"`
	// Device type for browser emulation
	//
	// Any of "desktop", "mobile", "tablet".
	Device string `json:"device,omitzero"`
	// Browser driver to use
	//
	// Any of "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro".
	Driver string `json:"driver,omitzero"`
	// Expected HTTP status codes for successful requests
	ExpectedStatusCodes []int64 `json:"expected_status_codes,omitzero"`
	// List of acceptable response formats in order of preference
	//
	// Any of "html", "markdown", "screenshot", "headers".
	Formats []string `json:"formats,omitzero"`
	// Custom HTTP headers to include in the request
	Headers map[string]ExtractBatchParamsInputHeaderUnion `json:"headers,omitzero"`
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
	Locale ExtractBatchParamsInputLocale `json:"locale,omitzero"`
	// HTTP method for the request
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method,omitzero"`
	// Filters for capturing network traffic
	NetworkCapture []ExtractBatchParamsInputNetworkCapture `json:"network_capture,omitzero"`
	// Operating system to emulate
	//
	// Any of "windows", "mac os", "linux", "android", "ios".
	Os string `json:"os,omitzero"`
	// Custom parser configuration as a key-value map
	Parser ExtractBatchParamsInputParserUnion `json:"parser,omitzero"`
	// Referrer policy for the request
	//
	// Any of "random", "no-referer", "same-origin", "google", "bing", "facebook",
	// "twitter", "instagram".
	ReferrerType ExtractBatchParamsInputReferrerType `json:"referrer_type,omitzero"`
	Session      ExtractBatchParamsInputSession      `json:"session,omitzero"`
	// Skills or capabilities required for the request
	Skill ExtractBatchParamsInputSkillUnion `json:"skill,omitzero"`
	// US state for geolocation (only valid when country is US)
	//
	// Any of "AL", "AK", "AS", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA",
	// "GU", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI",
	// "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "MP",
	// "OH", "OK", "OR", "PA", "PR", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA",
	// "VI", "WA", "WV", "WI", "WY".
	State string `json:"state,omitzero"`
	paramObj
}

func (r ExtractBatchParamsInput) MarshalJSON() (data []byte, err error) {
	type shadow ExtractBatchParamsInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractBatchParamsInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractBatchParamsInput](
		"device", "desktop", "mobile", "tablet",
	)
	apijson.RegisterFieldValidator[ExtractBatchParamsInput](
		"driver", "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro",
	)
	apijson.RegisterFieldValidator[ExtractBatchParamsInput](
		"method", "GET", "POST", "PUT", "PATCH", "DELETE",
	)
	apijson.RegisterFieldValidator[ExtractBatchParamsInput](
		"os", "windows", "mac os", "linux", "android", "ios",
	)
	apijson.RegisterFieldValidator[ExtractBatchParamsInput](
		"state", "AL", "AK", "AS", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA", "GU", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "MP", "OH", "OK", "OR", "PA", "PR", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "VI", "WA", "WV", "WI", "WY",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsInputBrowserUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractBatchsInputBrowserString)
	OfExtractBatchsInputBrowserString param.Opt[string]                     `json:",omitzero,inline"`
	OfExtractBatchsInputBrowserObject *ExtractBatchParamsInputBrowserObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsInputBrowserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractBatchsInputBrowserString, u.OfExtractBatchsInputBrowserObject)
}
func (u *ExtractBatchParamsInputBrowserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsInputBrowserUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractBatchsInputBrowserString) {
		return &u.OfExtractBatchsInputBrowserString
	} else if !param.IsOmitted(u.OfExtractBatchsInputBrowserObject) {
		return u.OfExtractBatchsInputBrowserObject
	}
	return nil
}

// Browser type to emulate
type ExtractBatchParamsInputBrowserString string

const (
	ExtractBatchParamsInputBrowserStringChrome  ExtractBatchParamsInputBrowserString = "chrome"
	ExtractBatchParamsInputBrowserStringFirefox ExtractBatchParamsInputBrowserString = "firefox"
)

// The property Name is required.
type ExtractBatchParamsInputBrowserObject struct {
	// Any of "chrome", "firefox".
	Name string `json:"name,omitzero" api:"required"`
	// Specific browser version to emulate
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r ExtractBatchParamsInputBrowserObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractBatchParamsInputBrowserObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractBatchParamsInputBrowserObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractBatchParamsInputBrowserObject](
		"name", "chrome", "firefox",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsInputBrowserActionUnion struct {
	OfAutoScrollAction        *shared.AutoScrollActionParam        `json:",omitzero,inline"`
	OfClickAction             *shared.ClickActionParam             `json:",omitzero,inline"`
	OfEvalAction              *shared.EvalActionParam              `json:",omitzero,inline"`
	OfFetchAction             *shared.FetchActionParam             `json:",omitzero,inline"`
	OfFillAction              *shared.FillActionParam              `json:",omitzero,inline"`
	OfGetCookiesAction        *shared.GetCookiesActionParam        `json:",omitzero,inline"`
	OfGotoAction              *shared.GotoActionParam              `json:",omitzero,inline"`
	OfPressAction             *shared.PressActionParam             `json:",omitzero,inline"`
	OfScreenshotAction        *shared.ScreenshotActionParam        `json:",omitzero,inline"`
	OfScrollAction            *shared.ScrollActionParam            `json:",omitzero,inline"`
	OfWaitAction              *shared.WaitActionParam              `json:",omitzero,inline"`
	OfWaitForElementAction    *shared.WaitForElementActionParam    `json:",omitzero,inline"`
	OfWaitForNavigationAction *shared.WaitForNavigationActionParam `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsInputBrowserActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAutoScrollAction,
		u.OfClickAction,
		u.OfEvalAction,
		u.OfFetchAction,
		u.OfFillAction,
		u.OfGetCookiesAction,
		u.OfGotoAction,
		u.OfPressAction,
		u.OfScreenshotAction,
		u.OfScrollAction,
		u.OfWaitAction,
		u.OfWaitForElementAction,
		u.OfWaitForNavigationAction)
}
func (u *ExtractBatchParamsInputBrowserActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsInputBrowserActionUnion) asAny() any {
	if !param.IsOmitted(u.OfAutoScrollAction) {
		return u.OfAutoScrollAction
	} else if !param.IsOmitted(u.OfClickAction) {
		return u.OfClickAction
	} else if !param.IsOmitted(u.OfEvalAction) {
		return u.OfEvalAction
	} else if !param.IsOmitted(u.OfFetchAction) {
		return u.OfFetchAction
	} else if !param.IsOmitted(u.OfFillAction) {
		return u.OfFillAction
	} else if !param.IsOmitted(u.OfGetCookiesAction) {
		return u.OfGetCookiesAction
	} else if !param.IsOmitted(u.OfGotoAction) {
		return u.OfGotoAction
	} else if !param.IsOmitted(u.OfPressAction) {
		return u.OfPressAction
	} else if !param.IsOmitted(u.OfScreenshotAction) {
		return u.OfScreenshotAction
	} else if !param.IsOmitted(u.OfScrollAction) {
		return u.OfScrollAction
	} else if !param.IsOmitted(u.OfWaitAction) {
		return u.OfWaitAction
	} else if !param.IsOmitted(u.OfWaitForElementAction) {
		return u.OfWaitForElementAction
	} else if !param.IsOmitted(u.OfWaitForNavigationAction) {
		return u.OfWaitForNavigationAction
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsInputCookiesUnion struct {
	OfExtractBatchsInputCookiesArray []ExtractBatchParamsInputCookiesArrayItem `json:",omitzero,inline"`
	OfString                         param.Opt[string]                         `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsInputCookiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractBatchsInputCookiesArray, u.OfString)
}
func (u *ExtractBatchParamsInputCookiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsInputCookiesUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractBatchsInputCookiesArray) {
		return &u.OfExtractBatchsInputCookiesArray
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type ExtractBatchParamsInputCookiesArrayItem struct {
	Creation      param.Opt[string]                                  `json:"creation,omitzero"`
	Domain        param.Opt[string]                                  `json:"domain,omitzero"`
	HostOnly      param.Opt[bool]                                    `json:"hostOnly,omitzero"`
	HTTPOnly      param.Opt[bool]                                    `json:"httpOnly,omitzero"`
	LastAccessed  param.Opt[string]                                  `json:"lastAccessed,omitzero"`
	Path          param.Opt[string]                                  `json:"path,omitzero"`
	PathIsDefault param.Opt[bool]                                    `json:"pathIsDefault,omitzero"`
	Expires       param.Opt[string]                                  `json:"expires,omitzero"`
	Name          param.Opt[string]                                  `json:"name,omitzero"`
	Secure        param.Opt[bool]                                    `json:"secure,omitzero"`
	Value         param.Opt[string]                                  `json:"value,omitzero"`
	Extensions    []string                                           `json:"extensions,omitzero"`
	MaxAge        ExtractBatchParamsInputCookiesArrayItemMaxAgeUnion `json:"maxAge,omitzero"`
	// Any of "strict", "lax", "none".
	SameSite    string         `json:"sameSite,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r ExtractBatchParamsInputCookiesArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow ExtractBatchParamsInputCookiesArrayItem
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ExtractBatchParamsInputCookiesArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractBatchParamsInputCookiesArrayItem](
		"sameSite", "strict", "lax", "none",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsInputCookiesArrayItemMaxAgeUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractBatchsInputCookiesArrayItemMaxAgeString)
	OfExtractBatchsInputCookiesArrayItemMaxAgeString param.Opt[ExtractBatchParamsInputCookiesArrayItemMaxAgeString] `json:",omitzero,inline"`
	OfFloat                                          param.Opt[float64]                                             `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsInputCookiesArrayItemMaxAgeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractBatchsInputCookiesArrayItemMaxAgeString, u.OfFloat)
}
func (u *ExtractBatchParamsInputCookiesArrayItemMaxAgeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsInputCookiesArrayItemMaxAgeUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractBatchsInputCookiesArrayItemMaxAgeString) {
		return &u.OfExtractBatchsInputCookiesArrayItemMaxAgeString
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type ExtractBatchParamsInputCookiesArrayItemMaxAgeString string

const (
	ExtractBatchParamsInputCookiesArrayItemMaxAgeStringInfinity      ExtractBatchParamsInputCookiesArrayItemMaxAgeString = "Infinity"
	ExtractBatchParamsInputCookiesArrayItemMaxAgeStringMinusInfinity ExtractBatchParamsInputCookiesArrayItemMaxAgeString = "-Infinity"
)

// Country code for geolocation and proxy selection
type ExtractBatchParamsInputCountry string

const (
	ExtractBatchParamsInputCountryAd  ExtractBatchParamsInputCountry = "AD"
	ExtractBatchParamsInputCountryAe  ExtractBatchParamsInputCountry = "AE"
	ExtractBatchParamsInputCountryAf  ExtractBatchParamsInputCountry = "AF"
	ExtractBatchParamsInputCountryAg  ExtractBatchParamsInputCountry = "AG"
	ExtractBatchParamsInputCountryAI  ExtractBatchParamsInputCountry = "AI"
	ExtractBatchParamsInputCountryAl  ExtractBatchParamsInputCountry = "AL"
	ExtractBatchParamsInputCountryAm  ExtractBatchParamsInputCountry = "AM"
	ExtractBatchParamsInputCountryAo  ExtractBatchParamsInputCountry = "AO"
	ExtractBatchParamsInputCountryAq  ExtractBatchParamsInputCountry = "AQ"
	ExtractBatchParamsInputCountryAr  ExtractBatchParamsInputCountry = "AR"
	ExtractBatchParamsInputCountryAs  ExtractBatchParamsInputCountry = "AS"
	ExtractBatchParamsInputCountryAt  ExtractBatchParamsInputCountry = "AT"
	ExtractBatchParamsInputCountryAu  ExtractBatchParamsInputCountry = "AU"
	ExtractBatchParamsInputCountryAw  ExtractBatchParamsInputCountry = "AW"
	ExtractBatchParamsInputCountryAx  ExtractBatchParamsInputCountry = "AX"
	ExtractBatchParamsInputCountryAz  ExtractBatchParamsInputCountry = "AZ"
	ExtractBatchParamsInputCountryBa  ExtractBatchParamsInputCountry = "BA"
	ExtractBatchParamsInputCountryBb  ExtractBatchParamsInputCountry = "BB"
	ExtractBatchParamsInputCountryBd  ExtractBatchParamsInputCountry = "BD"
	ExtractBatchParamsInputCountryBe  ExtractBatchParamsInputCountry = "BE"
	ExtractBatchParamsInputCountryBf  ExtractBatchParamsInputCountry = "BF"
	ExtractBatchParamsInputCountryBg  ExtractBatchParamsInputCountry = "BG"
	ExtractBatchParamsInputCountryBh  ExtractBatchParamsInputCountry = "BH"
	ExtractBatchParamsInputCountryBi  ExtractBatchParamsInputCountry = "BI"
	ExtractBatchParamsInputCountryBj  ExtractBatchParamsInputCountry = "BJ"
	ExtractBatchParamsInputCountryBl  ExtractBatchParamsInputCountry = "BL"
	ExtractBatchParamsInputCountryBm  ExtractBatchParamsInputCountry = "BM"
	ExtractBatchParamsInputCountryBn  ExtractBatchParamsInputCountry = "BN"
	ExtractBatchParamsInputCountryBo  ExtractBatchParamsInputCountry = "BO"
	ExtractBatchParamsInputCountryBq  ExtractBatchParamsInputCountry = "BQ"
	ExtractBatchParamsInputCountryBr  ExtractBatchParamsInputCountry = "BR"
	ExtractBatchParamsInputCountryBs  ExtractBatchParamsInputCountry = "BS"
	ExtractBatchParamsInputCountryBt  ExtractBatchParamsInputCountry = "BT"
	ExtractBatchParamsInputCountryBv  ExtractBatchParamsInputCountry = "BV"
	ExtractBatchParamsInputCountryBw  ExtractBatchParamsInputCountry = "BW"
	ExtractBatchParamsInputCountryBy  ExtractBatchParamsInputCountry = "BY"
	ExtractBatchParamsInputCountryBz  ExtractBatchParamsInputCountry = "BZ"
	ExtractBatchParamsInputCountryCa  ExtractBatchParamsInputCountry = "CA"
	ExtractBatchParamsInputCountryCc  ExtractBatchParamsInputCountry = "CC"
	ExtractBatchParamsInputCountryCd  ExtractBatchParamsInputCountry = "CD"
	ExtractBatchParamsInputCountryCf  ExtractBatchParamsInputCountry = "CF"
	ExtractBatchParamsInputCountryCg  ExtractBatchParamsInputCountry = "CG"
	ExtractBatchParamsInputCountryCh  ExtractBatchParamsInputCountry = "CH"
	ExtractBatchParamsInputCountryCi  ExtractBatchParamsInputCountry = "CI"
	ExtractBatchParamsInputCountryCk  ExtractBatchParamsInputCountry = "CK"
	ExtractBatchParamsInputCountryCl  ExtractBatchParamsInputCountry = "CL"
	ExtractBatchParamsInputCountryCm  ExtractBatchParamsInputCountry = "CM"
	ExtractBatchParamsInputCountryCn  ExtractBatchParamsInputCountry = "CN"
	ExtractBatchParamsInputCountryCo  ExtractBatchParamsInputCountry = "CO"
	ExtractBatchParamsInputCountryCr  ExtractBatchParamsInputCountry = "CR"
	ExtractBatchParamsInputCountryCu  ExtractBatchParamsInputCountry = "CU"
	ExtractBatchParamsInputCountryCv  ExtractBatchParamsInputCountry = "CV"
	ExtractBatchParamsInputCountryCw  ExtractBatchParamsInputCountry = "CW"
	ExtractBatchParamsInputCountryCx  ExtractBatchParamsInputCountry = "CX"
	ExtractBatchParamsInputCountryCy  ExtractBatchParamsInputCountry = "CY"
	ExtractBatchParamsInputCountryCz  ExtractBatchParamsInputCountry = "CZ"
	ExtractBatchParamsInputCountryDe  ExtractBatchParamsInputCountry = "DE"
	ExtractBatchParamsInputCountryDj  ExtractBatchParamsInputCountry = "DJ"
	ExtractBatchParamsInputCountryDk  ExtractBatchParamsInputCountry = "DK"
	ExtractBatchParamsInputCountryDm  ExtractBatchParamsInputCountry = "DM"
	ExtractBatchParamsInputCountryDo  ExtractBatchParamsInputCountry = "DO"
	ExtractBatchParamsInputCountryDz  ExtractBatchParamsInputCountry = "DZ"
	ExtractBatchParamsInputCountryEc  ExtractBatchParamsInputCountry = "EC"
	ExtractBatchParamsInputCountryEe  ExtractBatchParamsInputCountry = "EE"
	ExtractBatchParamsInputCountryEg  ExtractBatchParamsInputCountry = "EG"
	ExtractBatchParamsInputCountryEh  ExtractBatchParamsInputCountry = "EH"
	ExtractBatchParamsInputCountryEr  ExtractBatchParamsInputCountry = "ER"
	ExtractBatchParamsInputCountryEs  ExtractBatchParamsInputCountry = "ES"
	ExtractBatchParamsInputCountryEt  ExtractBatchParamsInputCountry = "ET"
	ExtractBatchParamsInputCountryFi  ExtractBatchParamsInputCountry = "FI"
	ExtractBatchParamsInputCountryFj  ExtractBatchParamsInputCountry = "FJ"
	ExtractBatchParamsInputCountryFk  ExtractBatchParamsInputCountry = "FK"
	ExtractBatchParamsInputCountryFm  ExtractBatchParamsInputCountry = "FM"
	ExtractBatchParamsInputCountryFo  ExtractBatchParamsInputCountry = "FO"
	ExtractBatchParamsInputCountryFr  ExtractBatchParamsInputCountry = "FR"
	ExtractBatchParamsInputCountryGa  ExtractBatchParamsInputCountry = "GA"
	ExtractBatchParamsInputCountryGB  ExtractBatchParamsInputCountry = "GB"
	ExtractBatchParamsInputCountryGd  ExtractBatchParamsInputCountry = "GD"
	ExtractBatchParamsInputCountryGe  ExtractBatchParamsInputCountry = "GE"
	ExtractBatchParamsInputCountryGf  ExtractBatchParamsInputCountry = "GF"
	ExtractBatchParamsInputCountryGg  ExtractBatchParamsInputCountry = "GG"
	ExtractBatchParamsInputCountryGh  ExtractBatchParamsInputCountry = "GH"
	ExtractBatchParamsInputCountryGi  ExtractBatchParamsInputCountry = "GI"
	ExtractBatchParamsInputCountryGl  ExtractBatchParamsInputCountry = "GL"
	ExtractBatchParamsInputCountryGm  ExtractBatchParamsInputCountry = "GM"
	ExtractBatchParamsInputCountryGn  ExtractBatchParamsInputCountry = "GN"
	ExtractBatchParamsInputCountryGp  ExtractBatchParamsInputCountry = "GP"
	ExtractBatchParamsInputCountryGq  ExtractBatchParamsInputCountry = "GQ"
	ExtractBatchParamsInputCountryGr  ExtractBatchParamsInputCountry = "GR"
	ExtractBatchParamsInputCountryGs  ExtractBatchParamsInputCountry = "GS"
	ExtractBatchParamsInputCountryGt  ExtractBatchParamsInputCountry = "GT"
	ExtractBatchParamsInputCountryGu  ExtractBatchParamsInputCountry = "GU"
	ExtractBatchParamsInputCountryGw  ExtractBatchParamsInputCountry = "GW"
	ExtractBatchParamsInputCountryGy  ExtractBatchParamsInputCountry = "GY"
	ExtractBatchParamsInputCountryHk  ExtractBatchParamsInputCountry = "HK"
	ExtractBatchParamsInputCountryHm  ExtractBatchParamsInputCountry = "HM"
	ExtractBatchParamsInputCountryHn  ExtractBatchParamsInputCountry = "HN"
	ExtractBatchParamsInputCountryHr  ExtractBatchParamsInputCountry = "HR"
	ExtractBatchParamsInputCountryHt  ExtractBatchParamsInputCountry = "HT"
	ExtractBatchParamsInputCountryHu  ExtractBatchParamsInputCountry = "HU"
	ExtractBatchParamsInputCountryID  ExtractBatchParamsInputCountry = "ID"
	ExtractBatchParamsInputCountryIe  ExtractBatchParamsInputCountry = "IE"
	ExtractBatchParamsInputCountryIl  ExtractBatchParamsInputCountry = "IL"
	ExtractBatchParamsInputCountryIm  ExtractBatchParamsInputCountry = "IM"
	ExtractBatchParamsInputCountryIn  ExtractBatchParamsInputCountry = "IN"
	ExtractBatchParamsInputCountryIo  ExtractBatchParamsInputCountry = "IO"
	ExtractBatchParamsInputCountryIq  ExtractBatchParamsInputCountry = "IQ"
	ExtractBatchParamsInputCountryIr  ExtractBatchParamsInputCountry = "IR"
	ExtractBatchParamsInputCountryIs  ExtractBatchParamsInputCountry = "IS"
	ExtractBatchParamsInputCountryIt  ExtractBatchParamsInputCountry = "IT"
	ExtractBatchParamsInputCountryJe  ExtractBatchParamsInputCountry = "JE"
	ExtractBatchParamsInputCountryJm  ExtractBatchParamsInputCountry = "JM"
	ExtractBatchParamsInputCountryJo  ExtractBatchParamsInputCountry = "JO"
	ExtractBatchParamsInputCountryJp  ExtractBatchParamsInputCountry = "JP"
	ExtractBatchParamsInputCountryKe  ExtractBatchParamsInputCountry = "KE"
	ExtractBatchParamsInputCountryKg  ExtractBatchParamsInputCountry = "KG"
	ExtractBatchParamsInputCountryKh  ExtractBatchParamsInputCountry = "KH"
	ExtractBatchParamsInputCountryKi  ExtractBatchParamsInputCountry = "KI"
	ExtractBatchParamsInputCountryKm  ExtractBatchParamsInputCountry = "KM"
	ExtractBatchParamsInputCountryKn  ExtractBatchParamsInputCountry = "KN"
	ExtractBatchParamsInputCountryKp  ExtractBatchParamsInputCountry = "KP"
	ExtractBatchParamsInputCountryKr  ExtractBatchParamsInputCountry = "KR"
	ExtractBatchParamsInputCountryKw  ExtractBatchParamsInputCountry = "KW"
	ExtractBatchParamsInputCountryKy  ExtractBatchParamsInputCountry = "KY"
	ExtractBatchParamsInputCountryKz  ExtractBatchParamsInputCountry = "KZ"
	ExtractBatchParamsInputCountryLa  ExtractBatchParamsInputCountry = "LA"
	ExtractBatchParamsInputCountryLb  ExtractBatchParamsInputCountry = "LB"
	ExtractBatchParamsInputCountryLc  ExtractBatchParamsInputCountry = "LC"
	ExtractBatchParamsInputCountryLi  ExtractBatchParamsInputCountry = "LI"
	ExtractBatchParamsInputCountryLk  ExtractBatchParamsInputCountry = "LK"
	ExtractBatchParamsInputCountryLr  ExtractBatchParamsInputCountry = "LR"
	ExtractBatchParamsInputCountryLs  ExtractBatchParamsInputCountry = "LS"
	ExtractBatchParamsInputCountryLt  ExtractBatchParamsInputCountry = "LT"
	ExtractBatchParamsInputCountryLu  ExtractBatchParamsInputCountry = "LU"
	ExtractBatchParamsInputCountryLv  ExtractBatchParamsInputCountry = "LV"
	ExtractBatchParamsInputCountryLy  ExtractBatchParamsInputCountry = "LY"
	ExtractBatchParamsInputCountryMa  ExtractBatchParamsInputCountry = "MA"
	ExtractBatchParamsInputCountryMc  ExtractBatchParamsInputCountry = "MC"
	ExtractBatchParamsInputCountryMd  ExtractBatchParamsInputCountry = "MD"
	ExtractBatchParamsInputCountryMe  ExtractBatchParamsInputCountry = "ME"
	ExtractBatchParamsInputCountryMf  ExtractBatchParamsInputCountry = "MF"
	ExtractBatchParamsInputCountryMg  ExtractBatchParamsInputCountry = "MG"
	ExtractBatchParamsInputCountryMh  ExtractBatchParamsInputCountry = "MH"
	ExtractBatchParamsInputCountryMk  ExtractBatchParamsInputCountry = "MK"
	ExtractBatchParamsInputCountryMl  ExtractBatchParamsInputCountry = "ML"
	ExtractBatchParamsInputCountryMm  ExtractBatchParamsInputCountry = "MM"
	ExtractBatchParamsInputCountryMn  ExtractBatchParamsInputCountry = "MN"
	ExtractBatchParamsInputCountryMo  ExtractBatchParamsInputCountry = "MO"
	ExtractBatchParamsInputCountryMp  ExtractBatchParamsInputCountry = "MP"
	ExtractBatchParamsInputCountryMq  ExtractBatchParamsInputCountry = "MQ"
	ExtractBatchParamsInputCountryMr  ExtractBatchParamsInputCountry = "MR"
	ExtractBatchParamsInputCountryMs  ExtractBatchParamsInputCountry = "MS"
	ExtractBatchParamsInputCountryMt  ExtractBatchParamsInputCountry = "MT"
	ExtractBatchParamsInputCountryMu  ExtractBatchParamsInputCountry = "MU"
	ExtractBatchParamsInputCountryMv  ExtractBatchParamsInputCountry = "MV"
	ExtractBatchParamsInputCountryMw  ExtractBatchParamsInputCountry = "MW"
	ExtractBatchParamsInputCountryMx  ExtractBatchParamsInputCountry = "MX"
	ExtractBatchParamsInputCountryMy  ExtractBatchParamsInputCountry = "MY"
	ExtractBatchParamsInputCountryMz  ExtractBatchParamsInputCountry = "MZ"
	ExtractBatchParamsInputCountryNa  ExtractBatchParamsInputCountry = "NA"
	ExtractBatchParamsInputCountryNc  ExtractBatchParamsInputCountry = "NC"
	ExtractBatchParamsInputCountryNe  ExtractBatchParamsInputCountry = "NE"
	ExtractBatchParamsInputCountryNf  ExtractBatchParamsInputCountry = "NF"
	ExtractBatchParamsInputCountryNg  ExtractBatchParamsInputCountry = "NG"
	ExtractBatchParamsInputCountryNi  ExtractBatchParamsInputCountry = "NI"
	ExtractBatchParamsInputCountryNl  ExtractBatchParamsInputCountry = "NL"
	ExtractBatchParamsInputCountryNo  ExtractBatchParamsInputCountry = "NO"
	ExtractBatchParamsInputCountryNp  ExtractBatchParamsInputCountry = "NP"
	ExtractBatchParamsInputCountryNr  ExtractBatchParamsInputCountry = "NR"
	ExtractBatchParamsInputCountryNu  ExtractBatchParamsInputCountry = "NU"
	ExtractBatchParamsInputCountryNz  ExtractBatchParamsInputCountry = "NZ"
	ExtractBatchParamsInputCountryOm  ExtractBatchParamsInputCountry = "OM"
	ExtractBatchParamsInputCountryPa  ExtractBatchParamsInputCountry = "PA"
	ExtractBatchParamsInputCountryPe  ExtractBatchParamsInputCountry = "PE"
	ExtractBatchParamsInputCountryPf  ExtractBatchParamsInputCountry = "PF"
	ExtractBatchParamsInputCountryPg  ExtractBatchParamsInputCountry = "PG"
	ExtractBatchParamsInputCountryPh  ExtractBatchParamsInputCountry = "PH"
	ExtractBatchParamsInputCountryPk  ExtractBatchParamsInputCountry = "PK"
	ExtractBatchParamsInputCountryPl  ExtractBatchParamsInputCountry = "PL"
	ExtractBatchParamsInputCountryPm  ExtractBatchParamsInputCountry = "PM"
	ExtractBatchParamsInputCountryPn  ExtractBatchParamsInputCountry = "PN"
	ExtractBatchParamsInputCountryPr  ExtractBatchParamsInputCountry = "PR"
	ExtractBatchParamsInputCountryPs  ExtractBatchParamsInputCountry = "PS"
	ExtractBatchParamsInputCountryPt  ExtractBatchParamsInputCountry = "PT"
	ExtractBatchParamsInputCountryPw  ExtractBatchParamsInputCountry = "PW"
	ExtractBatchParamsInputCountryPy  ExtractBatchParamsInputCountry = "PY"
	ExtractBatchParamsInputCountryQa  ExtractBatchParamsInputCountry = "QA"
	ExtractBatchParamsInputCountryRe  ExtractBatchParamsInputCountry = "RE"
	ExtractBatchParamsInputCountryRo  ExtractBatchParamsInputCountry = "RO"
	ExtractBatchParamsInputCountryRs  ExtractBatchParamsInputCountry = "RS"
	ExtractBatchParamsInputCountryRu  ExtractBatchParamsInputCountry = "RU"
	ExtractBatchParamsInputCountryRw  ExtractBatchParamsInputCountry = "RW"
	ExtractBatchParamsInputCountrySa  ExtractBatchParamsInputCountry = "SA"
	ExtractBatchParamsInputCountrySb  ExtractBatchParamsInputCountry = "SB"
	ExtractBatchParamsInputCountrySc  ExtractBatchParamsInputCountry = "SC"
	ExtractBatchParamsInputCountrySd  ExtractBatchParamsInputCountry = "SD"
	ExtractBatchParamsInputCountrySe  ExtractBatchParamsInputCountry = "SE"
	ExtractBatchParamsInputCountrySg  ExtractBatchParamsInputCountry = "SG"
	ExtractBatchParamsInputCountrySh  ExtractBatchParamsInputCountry = "SH"
	ExtractBatchParamsInputCountrySi  ExtractBatchParamsInputCountry = "SI"
	ExtractBatchParamsInputCountrySj  ExtractBatchParamsInputCountry = "SJ"
	ExtractBatchParamsInputCountrySk  ExtractBatchParamsInputCountry = "SK"
	ExtractBatchParamsInputCountrySl  ExtractBatchParamsInputCountry = "SL"
	ExtractBatchParamsInputCountrySm  ExtractBatchParamsInputCountry = "SM"
	ExtractBatchParamsInputCountrySn  ExtractBatchParamsInputCountry = "SN"
	ExtractBatchParamsInputCountrySo  ExtractBatchParamsInputCountry = "SO"
	ExtractBatchParamsInputCountrySr  ExtractBatchParamsInputCountry = "SR"
	ExtractBatchParamsInputCountrySS  ExtractBatchParamsInputCountry = "SS"
	ExtractBatchParamsInputCountrySt  ExtractBatchParamsInputCountry = "ST"
	ExtractBatchParamsInputCountrySv  ExtractBatchParamsInputCountry = "SV"
	ExtractBatchParamsInputCountrySx  ExtractBatchParamsInputCountry = "SX"
	ExtractBatchParamsInputCountrySy  ExtractBatchParamsInputCountry = "SY"
	ExtractBatchParamsInputCountrySz  ExtractBatchParamsInputCountry = "SZ"
	ExtractBatchParamsInputCountryTc  ExtractBatchParamsInputCountry = "TC"
	ExtractBatchParamsInputCountryTd  ExtractBatchParamsInputCountry = "TD"
	ExtractBatchParamsInputCountryTf  ExtractBatchParamsInputCountry = "TF"
	ExtractBatchParamsInputCountryTg  ExtractBatchParamsInputCountry = "TG"
	ExtractBatchParamsInputCountryTh  ExtractBatchParamsInputCountry = "TH"
	ExtractBatchParamsInputCountryTj  ExtractBatchParamsInputCountry = "TJ"
	ExtractBatchParamsInputCountryTk  ExtractBatchParamsInputCountry = "TK"
	ExtractBatchParamsInputCountryTl  ExtractBatchParamsInputCountry = "TL"
	ExtractBatchParamsInputCountryTm  ExtractBatchParamsInputCountry = "TM"
	ExtractBatchParamsInputCountryTn  ExtractBatchParamsInputCountry = "TN"
	ExtractBatchParamsInputCountryTo  ExtractBatchParamsInputCountry = "TO"
	ExtractBatchParamsInputCountryTr  ExtractBatchParamsInputCountry = "TR"
	ExtractBatchParamsInputCountryTt  ExtractBatchParamsInputCountry = "TT"
	ExtractBatchParamsInputCountryTv  ExtractBatchParamsInputCountry = "TV"
	ExtractBatchParamsInputCountryTw  ExtractBatchParamsInputCountry = "TW"
	ExtractBatchParamsInputCountryTz  ExtractBatchParamsInputCountry = "TZ"
	ExtractBatchParamsInputCountryUa  ExtractBatchParamsInputCountry = "UA"
	ExtractBatchParamsInputCountryUg  ExtractBatchParamsInputCountry = "UG"
	ExtractBatchParamsInputCountryUm  ExtractBatchParamsInputCountry = "UM"
	ExtractBatchParamsInputCountryUs  ExtractBatchParamsInputCountry = "US"
	ExtractBatchParamsInputCountryUy  ExtractBatchParamsInputCountry = "UY"
	ExtractBatchParamsInputCountryUz  ExtractBatchParamsInputCountry = "UZ"
	ExtractBatchParamsInputCountryVa  ExtractBatchParamsInputCountry = "VA"
	ExtractBatchParamsInputCountryVc  ExtractBatchParamsInputCountry = "VC"
	ExtractBatchParamsInputCountryVe  ExtractBatchParamsInputCountry = "VE"
	ExtractBatchParamsInputCountryVg  ExtractBatchParamsInputCountry = "VG"
	ExtractBatchParamsInputCountryVi  ExtractBatchParamsInputCountry = "VI"
	ExtractBatchParamsInputCountryVn  ExtractBatchParamsInputCountry = "VN"
	ExtractBatchParamsInputCountryVu  ExtractBatchParamsInputCountry = "VU"
	ExtractBatchParamsInputCountryWf  ExtractBatchParamsInputCountry = "WF"
	ExtractBatchParamsInputCountryWs  ExtractBatchParamsInputCountry = "WS"
	ExtractBatchParamsInputCountryXk  ExtractBatchParamsInputCountry = "XK"
	ExtractBatchParamsInputCountryYe  ExtractBatchParamsInputCountry = "YE"
	ExtractBatchParamsInputCountryYt  ExtractBatchParamsInputCountry = "YT"
	ExtractBatchParamsInputCountryZa  ExtractBatchParamsInputCountry = "ZA"
	ExtractBatchParamsInputCountryZm  ExtractBatchParamsInputCountry = "ZM"
	ExtractBatchParamsInputCountryZw  ExtractBatchParamsInputCountry = "ZW"
	ExtractBatchParamsInputCountryAll ExtractBatchParamsInputCountry = "ALL"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsInputHeaderUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsInputHeaderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractBatchParamsInputHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsInputHeaderUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Locale for browser language and region settings
type ExtractBatchParamsInputLocale string

const (
	ExtractBatchParamsInputLocaleAaDj      ExtractBatchParamsInputLocale = "aa-DJ"
	ExtractBatchParamsInputLocaleAaEr      ExtractBatchParamsInputLocale = "aa-ER"
	ExtractBatchParamsInputLocaleAaEt      ExtractBatchParamsInputLocale = "aa-ET"
	ExtractBatchParamsInputLocaleAf        ExtractBatchParamsInputLocale = "af"
	ExtractBatchParamsInputLocaleAfNa      ExtractBatchParamsInputLocale = "af-NA"
	ExtractBatchParamsInputLocaleAfZa      ExtractBatchParamsInputLocale = "af-ZA"
	ExtractBatchParamsInputLocaleAk        ExtractBatchParamsInputLocale = "ak"
	ExtractBatchParamsInputLocaleAkGh      ExtractBatchParamsInputLocale = "ak-GH"
	ExtractBatchParamsInputLocaleAm        ExtractBatchParamsInputLocale = "am"
	ExtractBatchParamsInputLocaleAmEt      ExtractBatchParamsInputLocale = "am-ET"
	ExtractBatchParamsInputLocaleAnEs      ExtractBatchParamsInputLocale = "an-ES"
	ExtractBatchParamsInputLocaleAr        ExtractBatchParamsInputLocale = "ar"
	ExtractBatchParamsInputLocaleArAe      ExtractBatchParamsInputLocale = "ar-AE"
	ExtractBatchParamsInputLocaleArBh      ExtractBatchParamsInputLocale = "ar-BH"
	ExtractBatchParamsInputLocaleArDz      ExtractBatchParamsInputLocale = "ar-DZ"
	ExtractBatchParamsInputLocaleArEg      ExtractBatchParamsInputLocale = "ar-EG"
	ExtractBatchParamsInputLocaleArIn      ExtractBatchParamsInputLocale = "ar-IN"
	ExtractBatchParamsInputLocaleArIq      ExtractBatchParamsInputLocale = "ar-IQ"
	ExtractBatchParamsInputLocaleArJo      ExtractBatchParamsInputLocale = "ar-JO"
	ExtractBatchParamsInputLocaleArKw      ExtractBatchParamsInputLocale = "ar-KW"
	ExtractBatchParamsInputLocaleArLb      ExtractBatchParamsInputLocale = "ar-LB"
	ExtractBatchParamsInputLocaleArLy      ExtractBatchParamsInputLocale = "ar-LY"
	ExtractBatchParamsInputLocaleArMa      ExtractBatchParamsInputLocale = "ar-MA"
	ExtractBatchParamsInputLocaleArOm      ExtractBatchParamsInputLocale = "ar-OM"
	ExtractBatchParamsInputLocaleArQa      ExtractBatchParamsInputLocale = "ar-QA"
	ExtractBatchParamsInputLocaleArSa      ExtractBatchParamsInputLocale = "ar-SA"
	ExtractBatchParamsInputLocaleArSd      ExtractBatchParamsInputLocale = "ar-SD"
	ExtractBatchParamsInputLocaleArSy      ExtractBatchParamsInputLocale = "ar-SY"
	ExtractBatchParamsInputLocaleArTn      ExtractBatchParamsInputLocale = "ar-TN"
	ExtractBatchParamsInputLocaleArYe      ExtractBatchParamsInputLocale = "ar-YE"
	ExtractBatchParamsInputLocaleAs        ExtractBatchParamsInputLocale = "as"
	ExtractBatchParamsInputLocaleAsIn      ExtractBatchParamsInputLocale = "as-IN"
	ExtractBatchParamsInputLocaleAsa       ExtractBatchParamsInputLocale = "asa"
	ExtractBatchParamsInputLocaleAsaTz     ExtractBatchParamsInputLocale = "asa-TZ"
	ExtractBatchParamsInputLocaleAstEs     ExtractBatchParamsInputLocale = "ast-ES"
	ExtractBatchParamsInputLocaleAz        ExtractBatchParamsInputLocale = "az"
	ExtractBatchParamsInputLocaleAzAz      ExtractBatchParamsInputLocale = "az-AZ"
	ExtractBatchParamsInputLocaleAzCyrl    ExtractBatchParamsInputLocale = "az-Cyrl"
	ExtractBatchParamsInputLocaleAzCyrlAz  ExtractBatchParamsInputLocale = "az-Cyrl-AZ"
	ExtractBatchParamsInputLocaleAzLatn    ExtractBatchParamsInputLocale = "az-Latn"
	ExtractBatchParamsInputLocaleAzLatnAz  ExtractBatchParamsInputLocale = "az-Latn-AZ"
	ExtractBatchParamsInputLocaleBe        ExtractBatchParamsInputLocale = "be"
	ExtractBatchParamsInputLocaleBeBy      ExtractBatchParamsInputLocale = "be-BY"
	ExtractBatchParamsInputLocaleBem       ExtractBatchParamsInputLocale = "bem"
	ExtractBatchParamsInputLocaleBemZm     ExtractBatchParamsInputLocale = "bem-ZM"
	ExtractBatchParamsInputLocaleBerDz     ExtractBatchParamsInputLocale = "ber-DZ"
	ExtractBatchParamsInputLocaleBerMa     ExtractBatchParamsInputLocale = "ber-MA"
	ExtractBatchParamsInputLocaleBez       ExtractBatchParamsInputLocale = "bez"
	ExtractBatchParamsInputLocaleBezTz     ExtractBatchParamsInputLocale = "bez-TZ"
	ExtractBatchParamsInputLocaleBg        ExtractBatchParamsInputLocale = "bg"
	ExtractBatchParamsInputLocaleBgBg      ExtractBatchParamsInputLocale = "bg-BG"
	ExtractBatchParamsInputLocaleBhoIn     ExtractBatchParamsInputLocale = "bho-IN"
	ExtractBatchParamsInputLocaleBm        ExtractBatchParamsInputLocale = "bm"
	ExtractBatchParamsInputLocaleBmMl      ExtractBatchParamsInputLocale = "bm-ML"
	ExtractBatchParamsInputLocaleBn        ExtractBatchParamsInputLocale = "bn"
	ExtractBatchParamsInputLocaleBnBd      ExtractBatchParamsInputLocale = "bn-BD"
	ExtractBatchParamsInputLocaleBnIn      ExtractBatchParamsInputLocale = "bn-IN"
	ExtractBatchParamsInputLocaleBo        ExtractBatchParamsInputLocale = "bo"
	ExtractBatchParamsInputLocaleBoCn      ExtractBatchParamsInputLocale = "bo-CN"
	ExtractBatchParamsInputLocaleBoIn      ExtractBatchParamsInputLocale = "bo-IN"
	ExtractBatchParamsInputLocaleBrFr      ExtractBatchParamsInputLocale = "br-FR"
	ExtractBatchParamsInputLocaleBrxIn     ExtractBatchParamsInputLocale = "brx-IN"
	ExtractBatchParamsInputLocaleBs        ExtractBatchParamsInputLocale = "bs"
	ExtractBatchParamsInputLocaleBsBa      ExtractBatchParamsInputLocale = "bs-BA"
	ExtractBatchParamsInputLocaleBynEr     ExtractBatchParamsInputLocale = "byn-ER"
	ExtractBatchParamsInputLocaleCa        ExtractBatchParamsInputLocale = "ca"
	ExtractBatchParamsInputLocaleCaAd      ExtractBatchParamsInputLocale = "ca-AD"
	ExtractBatchParamsInputLocaleCaEs      ExtractBatchParamsInputLocale = "ca-ES"
	ExtractBatchParamsInputLocaleCaFr      ExtractBatchParamsInputLocale = "ca-FR"
	ExtractBatchParamsInputLocaleCaIt      ExtractBatchParamsInputLocale = "ca-IT"
	ExtractBatchParamsInputLocaleCgg       ExtractBatchParamsInputLocale = "cgg"
	ExtractBatchParamsInputLocaleCggUg     ExtractBatchParamsInputLocale = "cgg-UG"
	ExtractBatchParamsInputLocaleChr       ExtractBatchParamsInputLocale = "chr"
	ExtractBatchParamsInputLocaleChrUs     ExtractBatchParamsInputLocale = "chr-US"
	ExtractBatchParamsInputLocaleCrhUa     ExtractBatchParamsInputLocale = "crh-UA"
	ExtractBatchParamsInputLocaleCs        ExtractBatchParamsInputLocale = "cs"
	ExtractBatchParamsInputLocaleCsCz      ExtractBatchParamsInputLocale = "cs-CZ"
	ExtractBatchParamsInputLocaleCsbPl     ExtractBatchParamsInputLocale = "csb-PL"
	ExtractBatchParamsInputLocaleCvRu      ExtractBatchParamsInputLocale = "cv-RU"
	ExtractBatchParamsInputLocaleCy        ExtractBatchParamsInputLocale = "cy"
	ExtractBatchParamsInputLocaleCyGB      ExtractBatchParamsInputLocale = "cy-GB"
	ExtractBatchParamsInputLocaleDa        ExtractBatchParamsInputLocale = "da"
	ExtractBatchParamsInputLocaleDaDk      ExtractBatchParamsInputLocale = "da-DK"
	ExtractBatchParamsInputLocaleDav       ExtractBatchParamsInputLocale = "dav"
	ExtractBatchParamsInputLocaleDavKe     ExtractBatchParamsInputLocale = "dav-KE"
	ExtractBatchParamsInputLocaleDe        ExtractBatchParamsInputLocale = "de"
	ExtractBatchParamsInputLocaleDeAt      ExtractBatchParamsInputLocale = "de-AT"
	ExtractBatchParamsInputLocaleDeBe      ExtractBatchParamsInputLocale = "de-BE"
	ExtractBatchParamsInputLocaleDeCh      ExtractBatchParamsInputLocale = "de-CH"
	ExtractBatchParamsInputLocaleDeDe      ExtractBatchParamsInputLocale = "de-DE"
	ExtractBatchParamsInputLocaleDeLi      ExtractBatchParamsInputLocale = "de-LI"
	ExtractBatchParamsInputLocaleDeLu      ExtractBatchParamsInputLocale = "de-LU"
	ExtractBatchParamsInputLocaleDvMv      ExtractBatchParamsInputLocale = "dv-MV"
	ExtractBatchParamsInputLocaleDzBt      ExtractBatchParamsInputLocale = "dz-BT"
	ExtractBatchParamsInputLocaleEbu       ExtractBatchParamsInputLocale = "ebu"
	ExtractBatchParamsInputLocaleEbuKe     ExtractBatchParamsInputLocale = "ebu-KE"
	ExtractBatchParamsInputLocaleEe        ExtractBatchParamsInputLocale = "ee"
	ExtractBatchParamsInputLocaleEeGh      ExtractBatchParamsInputLocale = "ee-GH"
	ExtractBatchParamsInputLocaleEeTg      ExtractBatchParamsInputLocale = "ee-TG"
	ExtractBatchParamsInputLocaleEl        ExtractBatchParamsInputLocale = "el"
	ExtractBatchParamsInputLocaleElCy      ExtractBatchParamsInputLocale = "el-CY"
	ExtractBatchParamsInputLocaleElGr      ExtractBatchParamsInputLocale = "el-GR"
	ExtractBatchParamsInputLocaleEn        ExtractBatchParamsInputLocale = "en"
	ExtractBatchParamsInputLocaleEnAg      ExtractBatchParamsInputLocale = "en-AG"
	ExtractBatchParamsInputLocaleEnAs      ExtractBatchParamsInputLocale = "en-AS"
	ExtractBatchParamsInputLocaleEnAu      ExtractBatchParamsInputLocale = "en-AU"
	ExtractBatchParamsInputLocaleEnBe      ExtractBatchParamsInputLocale = "en-BE"
	ExtractBatchParamsInputLocaleEnBw      ExtractBatchParamsInputLocale = "en-BW"
	ExtractBatchParamsInputLocaleEnBz      ExtractBatchParamsInputLocale = "en-BZ"
	ExtractBatchParamsInputLocaleEnCa      ExtractBatchParamsInputLocale = "en-CA"
	ExtractBatchParamsInputLocaleEnDk      ExtractBatchParamsInputLocale = "en-DK"
	ExtractBatchParamsInputLocaleEnGB      ExtractBatchParamsInputLocale = "en-GB"
	ExtractBatchParamsInputLocaleEnGu      ExtractBatchParamsInputLocale = "en-GU"
	ExtractBatchParamsInputLocaleEnHk      ExtractBatchParamsInputLocale = "en-HK"
	ExtractBatchParamsInputLocaleEnIe      ExtractBatchParamsInputLocale = "en-IE"
	ExtractBatchParamsInputLocaleEnIn      ExtractBatchParamsInputLocale = "en-IN"
	ExtractBatchParamsInputLocaleEnJm      ExtractBatchParamsInputLocale = "en-JM"
	ExtractBatchParamsInputLocaleEnMh      ExtractBatchParamsInputLocale = "en-MH"
	ExtractBatchParamsInputLocaleEnMp      ExtractBatchParamsInputLocale = "en-MP"
	ExtractBatchParamsInputLocaleEnMt      ExtractBatchParamsInputLocale = "en-MT"
	ExtractBatchParamsInputLocaleEnMu      ExtractBatchParamsInputLocale = "en-MU"
	ExtractBatchParamsInputLocaleEnNa      ExtractBatchParamsInputLocale = "en-NA"
	ExtractBatchParamsInputLocaleEnNg      ExtractBatchParamsInputLocale = "en-NG"
	ExtractBatchParamsInputLocaleEnNz      ExtractBatchParamsInputLocale = "en-NZ"
	ExtractBatchParamsInputLocaleEnPh      ExtractBatchParamsInputLocale = "en-PH"
	ExtractBatchParamsInputLocaleEnPk      ExtractBatchParamsInputLocale = "en-PK"
	ExtractBatchParamsInputLocaleEnSg      ExtractBatchParamsInputLocale = "en-SG"
	ExtractBatchParamsInputLocaleEnTt      ExtractBatchParamsInputLocale = "en-TT"
	ExtractBatchParamsInputLocaleEnUm      ExtractBatchParamsInputLocale = "en-UM"
	ExtractBatchParamsInputLocaleEnUs      ExtractBatchParamsInputLocale = "en-US"
	ExtractBatchParamsInputLocaleEnVi      ExtractBatchParamsInputLocale = "en-VI"
	ExtractBatchParamsInputLocaleEnZa      ExtractBatchParamsInputLocale = "en-ZA"
	ExtractBatchParamsInputLocaleEnZm      ExtractBatchParamsInputLocale = "en-ZM"
	ExtractBatchParamsInputLocaleEnZw      ExtractBatchParamsInputLocale = "en-ZW"
	ExtractBatchParamsInputLocaleEo        ExtractBatchParamsInputLocale = "eo"
	ExtractBatchParamsInputLocaleEs        ExtractBatchParamsInputLocale = "es"
	ExtractBatchParamsInputLocaleEs419     ExtractBatchParamsInputLocale = "es-419"
	ExtractBatchParamsInputLocaleEsAr      ExtractBatchParamsInputLocale = "es-AR"
	ExtractBatchParamsInputLocaleEsBo      ExtractBatchParamsInputLocale = "es-BO"
	ExtractBatchParamsInputLocaleEsCl      ExtractBatchParamsInputLocale = "es-CL"
	ExtractBatchParamsInputLocaleEsCo      ExtractBatchParamsInputLocale = "es-CO"
	ExtractBatchParamsInputLocaleEsCr      ExtractBatchParamsInputLocale = "es-CR"
	ExtractBatchParamsInputLocaleEsCu      ExtractBatchParamsInputLocale = "es-CU"
	ExtractBatchParamsInputLocaleEsDo      ExtractBatchParamsInputLocale = "es-DO"
	ExtractBatchParamsInputLocaleEsEc      ExtractBatchParamsInputLocale = "es-EC"
	ExtractBatchParamsInputLocaleEsEs      ExtractBatchParamsInputLocale = "es-ES"
	ExtractBatchParamsInputLocaleEsGq      ExtractBatchParamsInputLocale = "es-GQ"
	ExtractBatchParamsInputLocaleEsGt      ExtractBatchParamsInputLocale = "es-GT"
	ExtractBatchParamsInputLocaleEsHn      ExtractBatchParamsInputLocale = "es-HN"
	ExtractBatchParamsInputLocaleEsMx      ExtractBatchParamsInputLocale = "es-MX"
	ExtractBatchParamsInputLocaleEsNi      ExtractBatchParamsInputLocale = "es-NI"
	ExtractBatchParamsInputLocaleEsPa      ExtractBatchParamsInputLocale = "es-PA"
	ExtractBatchParamsInputLocaleEsPe      ExtractBatchParamsInputLocale = "es-PE"
	ExtractBatchParamsInputLocaleEsPr      ExtractBatchParamsInputLocale = "es-PR"
	ExtractBatchParamsInputLocaleEsPy      ExtractBatchParamsInputLocale = "es-PY"
	ExtractBatchParamsInputLocaleEsSv      ExtractBatchParamsInputLocale = "es-SV"
	ExtractBatchParamsInputLocaleEsUs      ExtractBatchParamsInputLocale = "es-US"
	ExtractBatchParamsInputLocaleEsUy      ExtractBatchParamsInputLocale = "es-UY"
	ExtractBatchParamsInputLocaleEsVe      ExtractBatchParamsInputLocale = "es-VE"
	ExtractBatchParamsInputLocaleEt        ExtractBatchParamsInputLocale = "et"
	ExtractBatchParamsInputLocaleEtEe      ExtractBatchParamsInputLocale = "et-EE"
	ExtractBatchParamsInputLocaleEu        ExtractBatchParamsInputLocale = "eu"
	ExtractBatchParamsInputLocaleEuEs      ExtractBatchParamsInputLocale = "eu-ES"
	ExtractBatchParamsInputLocaleFa        ExtractBatchParamsInputLocale = "fa"
	ExtractBatchParamsInputLocaleFaAf      ExtractBatchParamsInputLocale = "fa-AF"
	ExtractBatchParamsInputLocaleFaIr      ExtractBatchParamsInputLocale = "fa-IR"
	ExtractBatchParamsInputLocaleFf        ExtractBatchParamsInputLocale = "ff"
	ExtractBatchParamsInputLocaleFfSn      ExtractBatchParamsInputLocale = "ff-SN"
	ExtractBatchParamsInputLocaleFi        ExtractBatchParamsInputLocale = "fi"
	ExtractBatchParamsInputLocaleFiFi      ExtractBatchParamsInputLocale = "fi-FI"
	ExtractBatchParamsInputLocaleFil       ExtractBatchParamsInputLocale = "fil"
	ExtractBatchParamsInputLocaleFilPh     ExtractBatchParamsInputLocale = "fil-PH"
	ExtractBatchParamsInputLocaleFo        ExtractBatchParamsInputLocale = "fo"
	ExtractBatchParamsInputLocaleFoFo      ExtractBatchParamsInputLocale = "fo-FO"
	ExtractBatchParamsInputLocaleFr        ExtractBatchParamsInputLocale = "fr"
	ExtractBatchParamsInputLocaleFrBe      ExtractBatchParamsInputLocale = "fr-BE"
	ExtractBatchParamsInputLocaleFrBf      ExtractBatchParamsInputLocale = "fr-BF"
	ExtractBatchParamsInputLocaleFrBi      ExtractBatchParamsInputLocale = "fr-BI"
	ExtractBatchParamsInputLocaleFrBj      ExtractBatchParamsInputLocale = "fr-BJ"
	ExtractBatchParamsInputLocaleFrBl      ExtractBatchParamsInputLocale = "fr-BL"
	ExtractBatchParamsInputLocaleFrCa      ExtractBatchParamsInputLocale = "fr-CA"
	ExtractBatchParamsInputLocaleFrCd      ExtractBatchParamsInputLocale = "fr-CD"
	ExtractBatchParamsInputLocaleFrCf      ExtractBatchParamsInputLocale = "fr-CF"
	ExtractBatchParamsInputLocaleFrCg      ExtractBatchParamsInputLocale = "fr-CG"
	ExtractBatchParamsInputLocaleFrCh      ExtractBatchParamsInputLocale = "fr-CH"
	ExtractBatchParamsInputLocaleFrCi      ExtractBatchParamsInputLocale = "fr-CI"
	ExtractBatchParamsInputLocaleFrCm      ExtractBatchParamsInputLocale = "fr-CM"
	ExtractBatchParamsInputLocaleFrDj      ExtractBatchParamsInputLocale = "fr-DJ"
	ExtractBatchParamsInputLocaleFrFr      ExtractBatchParamsInputLocale = "fr-FR"
	ExtractBatchParamsInputLocaleFrGa      ExtractBatchParamsInputLocale = "fr-GA"
	ExtractBatchParamsInputLocaleFrGn      ExtractBatchParamsInputLocale = "fr-GN"
	ExtractBatchParamsInputLocaleFrGp      ExtractBatchParamsInputLocale = "fr-GP"
	ExtractBatchParamsInputLocaleFrGq      ExtractBatchParamsInputLocale = "fr-GQ"
	ExtractBatchParamsInputLocaleFrKm      ExtractBatchParamsInputLocale = "fr-KM"
	ExtractBatchParamsInputLocaleFrLu      ExtractBatchParamsInputLocale = "fr-LU"
	ExtractBatchParamsInputLocaleFrMc      ExtractBatchParamsInputLocale = "fr-MC"
	ExtractBatchParamsInputLocaleFrMf      ExtractBatchParamsInputLocale = "fr-MF"
	ExtractBatchParamsInputLocaleFrMg      ExtractBatchParamsInputLocale = "fr-MG"
	ExtractBatchParamsInputLocaleFrMl      ExtractBatchParamsInputLocale = "fr-ML"
	ExtractBatchParamsInputLocaleFrMq      ExtractBatchParamsInputLocale = "fr-MQ"
	ExtractBatchParamsInputLocaleFrNe      ExtractBatchParamsInputLocale = "fr-NE"
	ExtractBatchParamsInputLocaleFrRe      ExtractBatchParamsInputLocale = "fr-RE"
	ExtractBatchParamsInputLocaleFrRw      ExtractBatchParamsInputLocale = "fr-RW"
	ExtractBatchParamsInputLocaleFrSn      ExtractBatchParamsInputLocale = "fr-SN"
	ExtractBatchParamsInputLocaleFrTd      ExtractBatchParamsInputLocale = "fr-TD"
	ExtractBatchParamsInputLocaleFrTg      ExtractBatchParamsInputLocale = "fr-TG"
	ExtractBatchParamsInputLocaleFurIt     ExtractBatchParamsInputLocale = "fur-IT"
	ExtractBatchParamsInputLocaleFyDe      ExtractBatchParamsInputLocale = "fy-DE"
	ExtractBatchParamsInputLocaleFyNl      ExtractBatchParamsInputLocale = "fy-NL"
	ExtractBatchParamsInputLocaleGa        ExtractBatchParamsInputLocale = "ga"
	ExtractBatchParamsInputLocaleGaIe      ExtractBatchParamsInputLocale = "ga-IE"
	ExtractBatchParamsInputLocaleGdGB      ExtractBatchParamsInputLocale = "gd-GB"
	ExtractBatchParamsInputLocaleGezEr     ExtractBatchParamsInputLocale = "gez-ER"
	ExtractBatchParamsInputLocaleGezEt     ExtractBatchParamsInputLocale = "gez-ET"
	ExtractBatchParamsInputLocaleGl        ExtractBatchParamsInputLocale = "gl"
	ExtractBatchParamsInputLocaleGlEs      ExtractBatchParamsInputLocale = "gl-ES"
	ExtractBatchParamsInputLocaleGsw       ExtractBatchParamsInputLocale = "gsw"
	ExtractBatchParamsInputLocaleGswCh     ExtractBatchParamsInputLocale = "gsw-CH"
	ExtractBatchParamsInputLocaleGu        ExtractBatchParamsInputLocale = "gu"
	ExtractBatchParamsInputLocaleGuIn      ExtractBatchParamsInputLocale = "gu-IN"
	ExtractBatchParamsInputLocaleGuz       ExtractBatchParamsInputLocale = "guz"
	ExtractBatchParamsInputLocaleGuzKe     ExtractBatchParamsInputLocale = "guz-KE"
	ExtractBatchParamsInputLocaleGv        ExtractBatchParamsInputLocale = "gv"
	ExtractBatchParamsInputLocaleGvGB      ExtractBatchParamsInputLocale = "gv-GB"
	ExtractBatchParamsInputLocaleHa        ExtractBatchParamsInputLocale = "ha"
	ExtractBatchParamsInputLocaleHaLatn    ExtractBatchParamsInputLocale = "ha-Latn"
	ExtractBatchParamsInputLocaleHaLatnGh  ExtractBatchParamsInputLocale = "ha-Latn-GH"
	ExtractBatchParamsInputLocaleHaLatnNe  ExtractBatchParamsInputLocale = "ha-Latn-NE"
	ExtractBatchParamsInputLocaleHaLatnNg  ExtractBatchParamsInputLocale = "ha-Latn-NG"
	ExtractBatchParamsInputLocaleHaNg      ExtractBatchParamsInputLocale = "ha-NG"
	ExtractBatchParamsInputLocaleHaw       ExtractBatchParamsInputLocale = "haw"
	ExtractBatchParamsInputLocaleHawUs     ExtractBatchParamsInputLocale = "haw-US"
	ExtractBatchParamsInputLocaleHe        ExtractBatchParamsInputLocale = "he"
	ExtractBatchParamsInputLocaleHeIl      ExtractBatchParamsInputLocale = "he-IL"
	ExtractBatchParamsInputLocaleHi        ExtractBatchParamsInputLocale = "hi"
	ExtractBatchParamsInputLocaleHiIn      ExtractBatchParamsInputLocale = "hi-IN"
	ExtractBatchParamsInputLocaleHneIn     ExtractBatchParamsInputLocale = "hne-IN"
	ExtractBatchParamsInputLocaleHr        ExtractBatchParamsInputLocale = "hr"
	ExtractBatchParamsInputLocaleHrHr      ExtractBatchParamsInputLocale = "hr-HR"
	ExtractBatchParamsInputLocaleHsbDe     ExtractBatchParamsInputLocale = "hsb-DE"
	ExtractBatchParamsInputLocaleHtHt      ExtractBatchParamsInputLocale = "ht-HT"
	ExtractBatchParamsInputLocaleHu        ExtractBatchParamsInputLocale = "hu"
	ExtractBatchParamsInputLocaleHuHu      ExtractBatchParamsInputLocale = "hu-HU"
	ExtractBatchParamsInputLocaleHy        ExtractBatchParamsInputLocale = "hy"
	ExtractBatchParamsInputLocaleHyAm      ExtractBatchParamsInputLocale = "hy-AM"
	ExtractBatchParamsInputLocaleID        ExtractBatchParamsInputLocale = "id"
	ExtractBatchParamsInputLocaleIDID      ExtractBatchParamsInputLocale = "id-ID"
	ExtractBatchParamsInputLocaleIg        ExtractBatchParamsInputLocale = "ig"
	ExtractBatchParamsInputLocaleIgNg      ExtractBatchParamsInputLocale = "ig-NG"
	ExtractBatchParamsInputLocaleIi        ExtractBatchParamsInputLocale = "ii"
	ExtractBatchParamsInputLocaleIiCn      ExtractBatchParamsInputLocale = "ii-CN"
	ExtractBatchParamsInputLocaleIkCa      ExtractBatchParamsInputLocale = "ik-CA"
	ExtractBatchParamsInputLocaleIs        ExtractBatchParamsInputLocale = "is"
	ExtractBatchParamsInputLocaleIsIs      ExtractBatchParamsInputLocale = "is-IS"
	ExtractBatchParamsInputLocaleIt        ExtractBatchParamsInputLocale = "it"
	ExtractBatchParamsInputLocaleItCh      ExtractBatchParamsInputLocale = "it-CH"
	ExtractBatchParamsInputLocaleItIt      ExtractBatchParamsInputLocale = "it-IT"
	ExtractBatchParamsInputLocaleIuCa      ExtractBatchParamsInputLocale = "iu-CA"
	ExtractBatchParamsInputLocaleIwIl      ExtractBatchParamsInputLocale = "iw-IL"
	ExtractBatchParamsInputLocaleJa        ExtractBatchParamsInputLocale = "ja"
	ExtractBatchParamsInputLocaleJaJp      ExtractBatchParamsInputLocale = "ja-JP"
	ExtractBatchParamsInputLocaleJmc       ExtractBatchParamsInputLocale = "jmc"
	ExtractBatchParamsInputLocaleJmcTz     ExtractBatchParamsInputLocale = "jmc-TZ"
	ExtractBatchParamsInputLocaleKa        ExtractBatchParamsInputLocale = "ka"
	ExtractBatchParamsInputLocaleKaGe      ExtractBatchParamsInputLocale = "ka-GE"
	ExtractBatchParamsInputLocaleKab       ExtractBatchParamsInputLocale = "kab"
	ExtractBatchParamsInputLocaleKabDz     ExtractBatchParamsInputLocale = "kab-DZ"
	ExtractBatchParamsInputLocaleKam       ExtractBatchParamsInputLocale = "kam"
	ExtractBatchParamsInputLocaleKamKe     ExtractBatchParamsInputLocale = "kam-KE"
	ExtractBatchParamsInputLocaleKde       ExtractBatchParamsInputLocale = "kde"
	ExtractBatchParamsInputLocaleKdeTz     ExtractBatchParamsInputLocale = "kde-TZ"
	ExtractBatchParamsInputLocaleKea       ExtractBatchParamsInputLocale = "kea"
	ExtractBatchParamsInputLocaleKeaCv     ExtractBatchParamsInputLocale = "kea-CV"
	ExtractBatchParamsInputLocaleKhq       ExtractBatchParamsInputLocale = "khq"
	ExtractBatchParamsInputLocaleKhqMl     ExtractBatchParamsInputLocale = "khq-ML"
	ExtractBatchParamsInputLocaleKi        ExtractBatchParamsInputLocale = "ki"
	ExtractBatchParamsInputLocaleKiKe      ExtractBatchParamsInputLocale = "ki-KE"
	ExtractBatchParamsInputLocaleKk        ExtractBatchParamsInputLocale = "kk"
	ExtractBatchParamsInputLocaleKkCyrl    ExtractBatchParamsInputLocale = "kk-Cyrl"
	ExtractBatchParamsInputLocaleKkCyrlKz  ExtractBatchParamsInputLocale = "kk-Cyrl-KZ"
	ExtractBatchParamsInputLocaleKkKz      ExtractBatchParamsInputLocale = "kk-KZ"
	ExtractBatchParamsInputLocaleKl        ExtractBatchParamsInputLocale = "kl"
	ExtractBatchParamsInputLocaleKlGl      ExtractBatchParamsInputLocale = "kl-GL"
	ExtractBatchParamsInputLocaleKln       ExtractBatchParamsInputLocale = "kln"
	ExtractBatchParamsInputLocaleKlnKe     ExtractBatchParamsInputLocale = "kln-KE"
	ExtractBatchParamsInputLocaleKm        ExtractBatchParamsInputLocale = "km"
	ExtractBatchParamsInputLocaleKmKh      ExtractBatchParamsInputLocale = "km-KH"
	ExtractBatchParamsInputLocaleKn        ExtractBatchParamsInputLocale = "kn"
	ExtractBatchParamsInputLocaleKnIn      ExtractBatchParamsInputLocale = "kn-IN"
	ExtractBatchParamsInputLocaleKo        ExtractBatchParamsInputLocale = "ko"
	ExtractBatchParamsInputLocaleKoKr      ExtractBatchParamsInputLocale = "ko-KR"
	ExtractBatchParamsInputLocaleKok       ExtractBatchParamsInputLocale = "kok"
	ExtractBatchParamsInputLocaleKokIn     ExtractBatchParamsInputLocale = "kok-IN"
	ExtractBatchParamsInputLocaleKsIn      ExtractBatchParamsInputLocale = "ks-IN"
	ExtractBatchParamsInputLocaleKuTr      ExtractBatchParamsInputLocale = "ku-TR"
	ExtractBatchParamsInputLocaleKw        ExtractBatchParamsInputLocale = "kw"
	ExtractBatchParamsInputLocaleKwGB      ExtractBatchParamsInputLocale = "kw-GB"
	ExtractBatchParamsInputLocaleKyKg      ExtractBatchParamsInputLocale = "ky-KG"
	ExtractBatchParamsInputLocaleLag       ExtractBatchParamsInputLocale = "lag"
	ExtractBatchParamsInputLocaleLagTz     ExtractBatchParamsInputLocale = "lag-TZ"
	ExtractBatchParamsInputLocaleLbLu      ExtractBatchParamsInputLocale = "lb-LU"
	ExtractBatchParamsInputLocaleLg        ExtractBatchParamsInputLocale = "lg"
	ExtractBatchParamsInputLocaleLgUg      ExtractBatchParamsInputLocale = "lg-UG"
	ExtractBatchParamsInputLocaleLiBe      ExtractBatchParamsInputLocale = "li-BE"
	ExtractBatchParamsInputLocaleLiNl      ExtractBatchParamsInputLocale = "li-NL"
	ExtractBatchParamsInputLocaleLijIt     ExtractBatchParamsInputLocale = "lij-IT"
	ExtractBatchParamsInputLocaleLoLa      ExtractBatchParamsInputLocale = "lo-LA"
	ExtractBatchParamsInputLocaleLt        ExtractBatchParamsInputLocale = "lt"
	ExtractBatchParamsInputLocaleLtLt      ExtractBatchParamsInputLocale = "lt-LT"
	ExtractBatchParamsInputLocaleLuo       ExtractBatchParamsInputLocale = "luo"
	ExtractBatchParamsInputLocaleLuoKe     ExtractBatchParamsInputLocale = "luo-KE"
	ExtractBatchParamsInputLocaleLuy       ExtractBatchParamsInputLocale = "luy"
	ExtractBatchParamsInputLocaleLuyKe     ExtractBatchParamsInputLocale = "luy-KE"
	ExtractBatchParamsInputLocaleLv        ExtractBatchParamsInputLocale = "lv"
	ExtractBatchParamsInputLocaleLvLv      ExtractBatchParamsInputLocale = "lv-LV"
	ExtractBatchParamsInputLocaleMagIn     ExtractBatchParamsInputLocale = "mag-IN"
	ExtractBatchParamsInputLocaleMaiIn     ExtractBatchParamsInputLocale = "mai-IN"
	ExtractBatchParamsInputLocaleMas       ExtractBatchParamsInputLocale = "mas"
	ExtractBatchParamsInputLocaleMasKe     ExtractBatchParamsInputLocale = "mas-KE"
	ExtractBatchParamsInputLocaleMasTz     ExtractBatchParamsInputLocale = "mas-TZ"
	ExtractBatchParamsInputLocaleMer       ExtractBatchParamsInputLocale = "mer"
	ExtractBatchParamsInputLocaleMerKe     ExtractBatchParamsInputLocale = "mer-KE"
	ExtractBatchParamsInputLocaleMfe       ExtractBatchParamsInputLocale = "mfe"
	ExtractBatchParamsInputLocaleMfeMu     ExtractBatchParamsInputLocale = "mfe-MU"
	ExtractBatchParamsInputLocaleMg        ExtractBatchParamsInputLocale = "mg"
	ExtractBatchParamsInputLocaleMgMg      ExtractBatchParamsInputLocale = "mg-MG"
	ExtractBatchParamsInputLocaleMhrRu     ExtractBatchParamsInputLocale = "mhr-RU"
	ExtractBatchParamsInputLocaleMiNz      ExtractBatchParamsInputLocale = "mi-NZ"
	ExtractBatchParamsInputLocaleMk        ExtractBatchParamsInputLocale = "mk"
	ExtractBatchParamsInputLocaleMkMk      ExtractBatchParamsInputLocale = "mk-MK"
	ExtractBatchParamsInputLocaleMl        ExtractBatchParamsInputLocale = "ml"
	ExtractBatchParamsInputLocaleMlIn      ExtractBatchParamsInputLocale = "ml-IN"
	ExtractBatchParamsInputLocaleMnMn      ExtractBatchParamsInputLocale = "mn-MN"
	ExtractBatchParamsInputLocaleMr        ExtractBatchParamsInputLocale = "mr"
	ExtractBatchParamsInputLocaleMrIn      ExtractBatchParamsInputLocale = "mr-IN"
	ExtractBatchParamsInputLocaleMs        ExtractBatchParamsInputLocale = "ms"
	ExtractBatchParamsInputLocaleMsBn      ExtractBatchParamsInputLocale = "ms-BN"
	ExtractBatchParamsInputLocaleMsMy      ExtractBatchParamsInputLocale = "ms-MY"
	ExtractBatchParamsInputLocaleMt        ExtractBatchParamsInputLocale = "mt"
	ExtractBatchParamsInputLocaleMtMt      ExtractBatchParamsInputLocale = "mt-MT"
	ExtractBatchParamsInputLocaleMy        ExtractBatchParamsInputLocale = "my"
	ExtractBatchParamsInputLocaleMyMm      ExtractBatchParamsInputLocale = "my-MM"
	ExtractBatchParamsInputLocaleNanTw     ExtractBatchParamsInputLocale = "nan-TW"
	ExtractBatchParamsInputLocaleNaq       ExtractBatchParamsInputLocale = "naq"
	ExtractBatchParamsInputLocaleNaqNa     ExtractBatchParamsInputLocale = "naq-NA"
	ExtractBatchParamsInputLocaleNb        ExtractBatchParamsInputLocale = "nb"
	ExtractBatchParamsInputLocaleNbNo      ExtractBatchParamsInputLocale = "nb-NO"
	ExtractBatchParamsInputLocaleNd        ExtractBatchParamsInputLocale = "nd"
	ExtractBatchParamsInputLocaleNdZw      ExtractBatchParamsInputLocale = "nd-ZW"
	ExtractBatchParamsInputLocaleNdsDe     ExtractBatchParamsInputLocale = "nds-DE"
	ExtractBatchParamsInputLocaleNdsNl     ExtractBatchParamsInputLocale = "nds-NL"
	ExtractBatchParamsInputLocaleNe        ExtractBatchParamsInputLocale = "ne"
	ExtractBatchParamsInputLocaleNeIn      ExtractBatchParamsInputLocale = "ne-IN"
	ExtractBatchParamsInputLocaleNeNp      ExtractBatchParamsInputLocale = "ne-NP"
	ExtractBatchParamsInputLocaleNl        ExtractBatchParamsInputLocale = "nl"
	ExtractBatchParamsInputLocaleNlAw      ExtractBatchParamsInputLocale = "nl-AW"
	ExtractBatchParamsInputLocaleNlBe      ExtractBatchParamsInputLocale = "nl-BE"
	ExtractBatchParamsInputLocaleNlNl      ExtractBatchParamsInputLocale = "nl-NL"
	ExtractBatchParamsInputLocaleNn        ExtractBatchParamsInputLocale = "nn"
	ExtractBatchParamsInputLocaleNnNo      ExtractBatchParamsInputLocale = "nn-NO"
	ExtractBatchParamsInputLocaleNrZa      ExtractBatchParamsInputLocale = "nr-ZA"
	ExtractBatchParamsInputLocaleNsoZa     ExtractBatchParamsInputLocale = "nso-ZA"
	ExtractBatchParamsInputLocaleNyn       ExtractBatchParamsInputLocale = "nyn"
	ExtractBatchParamsInputLocaleNynUg     ExtractBatchParamsInputLocale = "nyn-UG"
	ExtractBatchParamsInputLocaleOcFr      ExtractBatchParamsInputLocale = "oc-FR"
	ExtractBatchParamsInputLocaleOm        ExtractBatchParamsInputLocale = "om"
	ExtractBatchParamsInputLocaleOmEt      ExtractBatchParamsInputLocale = "om-ET"
	ExtractBatchParamsInputLocaleOmKe      ExtractBatchParamsInputLocale = "om-KE"
	ExtractBatchParamsInputLocaleOr        ExtractBatchParamsInputLocale = "or"
	ExtractBatchParamsInputLocaleOrIn      ExtractBatchParamsInputLocale = "or-IN"
	ExtractBatchParamsInputLocaleOsRu      ExtractBatchParamsInputLocale = "os-RU"
	ExtractBatchParamsInputLocalePa        ExtractBatchParamsInputLocale = "pa"
	ExtractBatchParamsInputLocalePaArab    ExtractBatchParamsInputLocale = "pa-Arab"
	ExtractBatchParamsInputLocalePaArabPk  ExtractBatchParamsInputLocale = "pa-Arab-PK"
	ExtractBatchParamsInputLocalePaGuru    ExtractBatchParamsInputLocale = "pa-Guru"
	ExtractBatchParamsInputLocalePaGuruIn  ExtractBatchParamsInputLocale = "pa-Guru-IN"
	ExtractBatchParamsInputLocalePaIn      ExtractBatchParamsInputLocale = "pa-IN"
	ExtractBatchParamsInputLocalePaPk      ExtractBatchParamsInputLocale = "pa-PK"
	ExtractBatchParamsInputLocalePapAn     ExtractBatchParamsInputLocale = "pap-AN"
	ExtractBatchParamsInputLocalePl        ExtractBatchParamsInputLocale = "pl"
	ExtractBatchParamsInputLocalePlPl      ExtractBatchParamsInputLocale = "pl-PL"
	ExtractBatchParamsInputLocalePs        ExtractBatchParamsInputLocale = "ps"
	ExtractBatchParamsInputLocalePsAf      ExtractBatchParamsInputLocale = "ps-AF"
	ExtractBatchParamsInputLocalePt        ExtractBatchParamsInputLocale = "pt"
	ExtractBatchParamsInputLocalePtBr      ExtractBatchParamsInputLocale = "pt-BR"
	ExtractBatchParamsInputLocalePtGw      ExtractBatchParamsInputLocale = "pt-GW"
	ExtractBatchParamsInputLocalePtMz      ExtractBatchParamsInputLocale = "pt-MZ"
	ExtractBatchParamsInputLocalePtPt      ExtractBatchParamsInputLocale = "pt-PT"
	ExtractBatchParamsInputLocaleRm        ExtractBatchParamsInputLocale = "rm"
	ExtractBatchParamsInputLocaleRmCh      ExtractBatchParamsInputLocale = "rm-CH"
	ExtractBatchParamsInputLocaleRo        ExtractBatchParamsInputLocale = "ro"
	ExtractBatchParamsInputLocaleRoMd      ExtractBatchParamsInputLocale = "ro-MD"
	ExtractBatchParamsInputLocaleRoRo      ExtractBatchParamsInputLocale = "ro-RO"
	ExtractBatchParamsInputLocaleRof       ExtractBatchParamsInputLocale = "rof"
	ExtractBatchParamsInputLocaleRofTz     ExtractBatchParamsInputLocale = "rof-TZ"
	ExtractBatchParamsInputLocaleRu        ExtractBatchParamsInputLocale = "ru"
	ExtractBatchParamsInputLocaleRuMd      ExtractBatchParamsInputLocale = "ru-MD"
	ExtractBatchParamsInputLocaleRuRu      ExtractBatchParamsInputLocale = "ru-RU"
	ExtractBatchParamsInputLocaleRuUa      ExtractBatchParamsInputLocale = "ru-UA"
	ExtractBatchParamsInputLocaleRw        ExtractBatchParamsInputLocale = "rw"
	ExtractBatchParamsInputLocaleRwRw      ExtractBatchParamsInputLocale = "rw-RW"
	ExtractBatchParamsInputLocaleRwk       ExtractBatchParamsInputLocale = "rwk"
	ExtractBatchParamsInputLocaleRwkTz     ExtractBatchParamsInputLocale = "rwk-TZ"
	ExtractBatchParamsInputLocaleSaIn      ExtractBatchParamsInputLocale = "sa-IN"
	ExtractBatchParamsInputLocaleSaq       ExtractBatchParamsInputLocale = "saq"
	ExtractBatchParamsInputLocaleSaqKe     ExtractBatchParamsInputLocale = "saq-KE"
	ExtractBatchParamsInputLocaleScIt      ExtractBatchParamsInputLocale = "sc-IT"
	ExtractBatchParamsInputLocaleSdIn      ExtractBatchParamsInputLocale = "sd-IN"
	ExtractBatchParamsInputLocaleSeNo      ExtractBatchParamsInputLocale = "se-NO"
	ExtractBatchParamsInputLocaleSeh       ExtractBatchParamsInputLocale = "seh"
	ExtractBatchParamsInputLocaleSehMz     ExtractBatchParamsInputLocale = "seh-MZ"
	ExtractBatchParamsInputLocaleSes       ExtractBatchParamsInputLocale = "ses"
	ExtractBatchParamsInputLocaleSesMl     ExtractBatchParamsInputLocale = "ses-ML"
	ExtractBatchParamsInputLocaleSg        ExtractBatchParamsInputLocale = "sg"
	ExtractBatchParamsInputLocaleSgCf      ExtractBatchParamsInputLocale = "sg-CF"
	ExtractBatchParamsInputLocaleShi       ExtractBatchParamsInputLocale = "shi"
	ExtractBatchParamsInputLocaleShiLatn   ExtractBatchParamsInputLocale = "shi-Latn"
	ExtractBatchParamsInputLocaleShiLatnMa ExtractBatchParamsInputLocale = "shi-Latn-MA"
	ExtractBatchParamsInputLocaleShiTfng   ExtractBatchParamsInputLocale = "shi-Tfng"
	ExtractBatchParamsInputLocaleShiTfngMa ExtractBatchParamsInputLocale = "shi-Tfng-MA"
	ExtractBatchParamsInputLocaleShsCa     ExtractBatchParamsInputLocale = "shs-CA"
	ExtractBatchParamsInputLocaleSi        ExtractBatchParamsInputLocale = "si"
	ExtractBatchParamsInputLocaleSiLk      ExtractBatchParamsInputLocale = "si-LK"
	ExtractBatchParamsInputLocaleSidEt     ExtractBatchParamsInputLocale = "sid-ET"
	ExtractBatchParamsInputLocaleSk        ExtractBatchParamsInputLocale = "sk"
	ExtractBatchParamsInputLocaleSkSk      ExtractBatchParamsInputLocale = "sk-SK"
	ExtractBatchParamsInputLocaleSl        ExtractBatchParamsInputLocale = "sl"
	ExtractBatchParamsInputLocaleSlSi      ExtractBatchParamsInputLocale = "sl-SI"
	ExtractBatchParamsInputLocaleSn        ExtractBatchParamsInputLocale = "sn"
	ExtractBatchParamsInputLocaleSnZw      ExtractBatchParamsInputLocale = "sn-ZW"
	ExtractBatchParamsInputLocaleSo        ExtractBatchParamsInputLocale = "so"
	ExtractBatchParamsInputLocaleSoDj      ExtractBatchParamsInputLocale = "so-DJ"
	ExtractBatchParamsInputLocaleSoEt      ExtractBatchParamsInputLocale = "so-ET"
	ExtractBatchParamsInputLocaleSoKe      ExtractBatchParamsInputLocale = "so-KE"
	ExtractBatchParamsInputLocaleSoSo      ExtractBatchParamsInputLocale = "so-SO"
	ExtractBatchParamsInputLocaleSq        ExtractBatchParamsInputLocale = "sq"
	ExtractBatchParamsInputLocaleSqAl      ExtractBatchParamsInputLocale = "sq-AL"
	ExtractBatchParamsInputLocaleSqMk      ExtractBatchParamsInputLocale = "sq-MK"
	ExtractBatchParamsInputLocaleSr        ExtractBatchParamsInputLocale = "sr"
	ExtractBatchParamsInputLocaleSrCyrl    ExtractBatchParamsInputLocale = "sr-Cyrl"
	ExtractBatchParamsInputLocaleSrCyrlBa  ExtractBatchParamsInputLocale = "sr-Cyrl-BA"
	ExtractBatchParamsInputLocaleSrCyrlMe  ExtractBatchParamsInputLocale = "sr-Cyrl-ME"
	ExtractBatchParamsInputLocaleSrCyrlRs  ExtractBatchParamsInputLocale = "sr-Cyrl-RS"
	ExtractBatchParamsInputLocaleSrLatn    ExtractBatchParamsInputLocale = "sr-Latn"
	ExtractBatchParamsInputLocaleSrLatnBa  ExtractBatchParamsInputLocale = "sr-Latn-BA"
	ExtractBatchParamsInputLocaleSrLatnMe  ExtractBatchParamsInputLocale = "sr-Latn-ME"
	ExtractBatchParamsInputLocaleSrLatnRs  ExtractBatchParamsInputLocale = "sr-Latn-RS"
	ExtractBatchParamsInputLocaleSrMe      ExtractBatchParamsInputLocale = "sr-ME"
	ExtractBatchParamsInputLocaleSrRs      ExtractBatchParamsInputLocale = "sr-RS"
	ExtractBatchParamsInputLocaleSSZa      ExtractBatchParamsInputLocale = "ss-ZA"
	ExtractBatchParamsInputLocaleStZa      ExtractBatchParamsInputLocale = "st-ZA"
	ExtractBatchParamsInputLocaleSv        ExtractBatchParamsInputLocale = "sv"
	ExtractBatchParamsInputLocaleSvFi      ExtractBatchParamsInputLocale = "sv-FI"
	ExtractBatchParamsInputLocaleSvSe      ExtractBatchParamsInputLocale = "sv-SE"
	ExtractBatchParamsInputLocaleSw        ExtractBatchParamsInputLocale = "sw"
	ExtractBatchParamsInputLocaleSwKe      ExtractBatchParamsInputLocale = "sw-KE"
	ExtractBatchParamsInputLocaleSwTz      ExtractBatchParamsInputLocale = "sw-TZ"
	ExtractBatchParamsInputLocaleTa        ExtractBatchParamsInputLocale = "ta"
	ExtractBatchParamsInputLocaleTaIn      ExtractBatchParamsInputLocale = "ta-IN"
	ExtractBatchParamsInputLocaleTaLk      ExtractBatchParamsInputLocale = "ta-LK"
	ExtractBatchParamsInputLocaleTe        ExtractBatchParamsInputLocale = "te"
	ExtractBatchParamsInputLocaleTeIn      ExtractBatchParamsInputLocale = "te-IN"
	ExtractBatchParamsInputLocaleTeo       ExtractBatchParamsInputLocale = "teo"
	ExtractBatchParamsInputLocaleTeoKe     ExtractBatchParamsInputLocale = "teo-KE"
	ExtractBatchParamsInputLocaleTeoUg     ExtractBatchParamsInputLocale = "teo-UG"
	ExtractBatchParamsInputLocaleTgTj      ExtractBatchParamsInputLocale = "tg-TJ"
	ExtractBatchParamsInputLocaleTh        ExtractBatchParamsInputLocale = "th"
	ExtractBatchParamsInputLocaleThTh      ExtractBatchParamsInputLocale = "th-TH"
	ExtractBatchParamsInputLocaleTi        ExtractBatchParamsInputLocale = "ti"
	ExtractBatchParamsInputLocaleTiEr      ExtractBatchParamsInputLocale = "ti-ER"
	ExtractBatchParamsInputLocaleTiEt      ExtractBatchParamsInputLocale = "ti-ET"
	ExtractBatchParamsInputLocaleTigEr     ExtractBatchParamsInputLocale = "tig-ER"
	ExtractBatchParamsInputLocaleTkTm      ExtractBatchParamsInputLocale = "tk-TM"
	ExtractBatchParamsInputLocaleTlPh      ExtractBatchParamsInputLocale = "tl-PH"
	ExtractBatchParamsInputLocaleTnZa      ExtractBatchParamsInputLocale = "tn-ZA"
	ExtractBatchParamsInputLocaleTo        ExtractBatchParamsInputLocale = "to"
	ExtractBatchParamsInputLocaleToTo      ExtractBatchParamsInputLocale = "to-TO"
	ExtractBatchParamsInputLocaleTr        ExtractBatchParamsInputLocale = "tr"
	ExtractBatchParamsInputLocaleTrCy      ExtractBatchParamsInputLocale = "tr-CY"
	ExtractBatchParamsInputLocaleTrTr      ExtractBatchParamsInputLocale = "tr-TR"
	ExtractBatchParamsInputLocaleTsZa      ExtractBatchParamsInputLocale = "ts-ZA"
	ExtractBatchParamsInputLocaleTtRu      ExtractBatchParamsInputLocale = "tt-RU"
	ExtractBatchParamsInputLocaleTzm       ExtractBatchParamsInputLocale = "tzm"
	ExtractBatchParamsInputLocaleTzmLatn   ExtractBatchParamsInputLocale = "tzm-Latn"
	ExtractBatchParamsInputLocaleTzmLatnMa ExtractBatchParamsInputLocale = "tzm-Latn-MA"
	ExtractBatchParamsInputLocaleUgCn      ExtractBatchParamsInputLocale = "ug-CN"
	ExtractBatchParamsInputLocaleUk        ExtractBatchParamsInputLocale = "uk"
	ExtractBatchParamsInputLocaleUkUa      ExtractBatchParamsInputLocale = "uk-UA"
	ExtractBatchParamsInputLocaleUnmUs     ExtractBatchParamsInputLocale = "unm-US"
	ExtractBatchParamsInputLocaleUr        ExtractBatchParamsInputLocale = "ur"
	ExtractBatchParamsInputLocaleUrIn      ExtractBatchParamsInputLocale = "ur-IN"
	ExtractBatchParamsInputLocaleUrPk      ExtractBatchParamsInputLocale = "ur-PK"
	ExtractBatchParamsInputLocaleUz        ExtractBatchParamsInputLocale = "uz"
	ExtractBatchParamsInputLocaleUzArab    ExtractBatchParamsInputLocale = "uz-Arab"
	ExtractBatchParamsInputLocaleUzArabAf  ExtractBatchParamsInputLocale = "uz-Arab-AF"
	ExtractBatchParamsInputLocaleUzCyrl    ExtractBatchParamsInputLocale = "uz-Cyrl"
	ExtractBatchParamsInputLocaleUzCyrlUz  ExtractBatchParamsInputLocale = "uz-Cyrl-UZ"
	ExtractBatchParamsInputLocaleUzLatn    ExtractBatchParamsInputLocale = "uz-Latn"
	ExtractBatchParamsInputLocaleUzLatnUz  ExtractBatchParamsInputLocale = "uz-Latn-UZ"
	ExtractBatchParamsInputLocaleUzUz      ExtractBatchParamsInputLocale = "uz-UZ"
	ExtractBatchParamsInputLocaleVeZa      ExtractBatchParamsInputLocale = "ve-ZA"
	ExtractBatchParamsInputLocaleVi        ExtractBatchParamsInputLocale = "vi"
	ExtractBatchParamsInputLocaleViVn      ExtractBatchParamsInputLocale = "vi-VN"
	ExtractBatchParamsInputLocaleVun       ExtractBatchParamsInputLocale = "vun"
	ExtractBatchParamsInputLocaleVunTz     ExtractBatchParamsInputLocale = "vun-TZ"
	ExtractBatchParamsInputLocaleWaBe      ExtractBatchParamsInputLocale = "wa-BE"
	ExtractBatchParamsInputLocaleWaeCh     ExtractBatchParamsInputLocale = "wae-CH"
	ExtractBatchParamsInputLocaleWalEt     ExtractBatchParamsInputLocale = "wal-ET"
	ExtractBatchParamsInputLocaleWoSn      ExtractBatchParamsInputLocale = "wo-SN"
	ExtractBatchParamsInputLocaleXhZa      ExtractBatchParamsInputLocale = "xh-ZA"
	ExtractBatchParamsInputLocaleXog       ExtractBatchParamsInputLocale = "xog"
	ExtractBatchParamsInputLocaleXogUg     ExtractBatchParamsInputLocale = "xog-UG"
	ExtractBatchParamsInputLocaleYiUs      ExtractBatchParamsInputLocale = "yi-US"
	ExtractBatchParamsInputLocaleYo        ExtractBatchParamsInputLocale = "yo"
	ExtractBatchParamsInputLocaleYoNg      ExtractBatchParamsInputLocale = "yo-NG"
	ExtractBatchParamsInputLocaleYueHk     ExtractBatchParamsInputLocale = "yue-HK"
	ExtractBatchParamsInputLocaleZh        ExtractBatchParamsInputLocale = "zh"
	ExtractBatchParamsInputLocaleZhCn      ExtractBatchParamsInputLocale = "zh-CN"
	ExtractBatchParamsInputLocaleZhHk      ExtractBatchParamsInputLocale = "zh-HK"
	ExtractBatchParamsInputLocaleZhHans    ExtractBatchParamsInputLocale = "zh-Hans"
	ExtractBatchParamsInputLocaleZhHansCn  ExtractBatchParamsInputLocale = "zh-Hans-CN"
	ExtractBatchParamsInputLocaleZhHansHk  ExtractBatchParamsInputLocale = "zh-Hans-HK"
	ExtractBatchParamsInputLocaleZhHansMo  ExtractBatchParamsInputLocale = "zh-Hans-MO"
	ExtractBatchParamsInputLocaleZhHansSg  ExtractBatchParamsInputLocale = "zh-Hans-SG"
	ExtractBatchParamsInputLocaleZhHant    ExtractBatchParamsInputLocale = "zh-Hant"
	ExtractBatchParamsInputLocaleZhHantHk  ExtractBatchParamsInputLocale = "zh-Hant-HK"
	ExtractBatchParamsInputLocaleZhHantMo  ExtractBatchParamsInputLocale = "zh-Hant-MO"
	ExtractBatchParamsInputLocaleZhHantTw  ExtractBatchParamsInputLocale = "zh-Hant-TW"
	ExtractBatchParamsInputLocaleZhSg      ExtractBatchParamsInputLocale = "zh-SG"
	ExtractBatchParamsInputLocaleZhTw      ExtractBatchParamsInputLocale = "zh-TW"
	ExtractBatchParamsInputLocaleZu        ExtractBatchParamsInputLocale = "zu"
	ExtractBatchParamsInputLocaleZuZa      ExtractBatchParamsInputLocale = "zu-ZA"
	ExtractBatchParamsInputLocaleAuto      ExtractBatchParamsInputLocale = "auto"
)

type ExtractBatchParamsInputNetworkCapture struct {
	Validation                  param.Opt[bool]    `json:"validation,omitzero"`
	WaitForRequestsCount        param.Opt[float64] `json:"wait_for_requests_count,omitzero"`
	WaitForRequestsCountTimeout param.Opt[float64] `json:"wait_for_requests_count_timeout,omitzero"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method,omitzero"`
	// Resource type for network capture filtering
	ResourceType ExtractBatchParamsInputNetworkCaptureResourceTypeUnion `json:"resource_type,omitzero"`
	StatusCode   ExtractBatchParamsInputNetworkCaptureStatusCodeUnion   `json:"status_code,omitzero"`
	URL          ExtractBatchParamsInputNetworkCaptureURL               `json:"url,omitzero"`
	paramObj
}

func (r ExtractBatchParamsInputNetworkCapture) MarshalJSON() (data []byte, err error) {
	type shadow ExtractBatchParamsInputNetworkCapture
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractBatchParamsInputNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractBatchParamsInputNetworkCapture](
		"method", "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsInputNetworkCaptureResourceTypeUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsInputNetworkCaptureResourceTypeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractBatchParamsInputNetworkCaptureResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsInputNetworkCaptureResourceTypeUnion) asAny() any {
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
type ExtractBatchParamsInputNetworkCaptureStatusCodeUnion struct {
	OfFloat      param.Opt[float64] `json:",omitzero,inline"`
	OfFloatArray []float64          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsInputNetworkCaptureStatusCodeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfFloatArray)
}
func (u *ExtractBatchParamsInputNetworkCaptureStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsInputNetworkCaptureStatusCodeUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfFloatArray) {
		return &u.OfFloatArray
	}
	return nil
}

// The property Value is required.
type ExtractBatchParamsInputNetworkCaptureURL struct {
	Value string `json:"value" api:"required"`
	// Any of "exact", "contains".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ExtractBatchParamsInputNetworkCaptureURL) MarshalJSON() (data []byte, err error) {
	type shadow ExtractBatchParamsInputNetworkCaptureURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractBatchParamsInputNetworkCaptureURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractBatchParamsInputNetworkCaptureURL](
		"type", "exact", "contains",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsInputParserUnion struct {
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsInputParserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap, u.OfString)
}
func (u *ExtractBatchParamsInputParserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsInputParserUnion) asAny() any {
	if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Referrer policy for the request
type ExtractBatchParamsInputReferrerType string

const (
	ExtractBatchParamsInputReferrerTypeRandom     ExtractBatchParamsInputReferrerType = "random"
	ExtractBatchParamsInputReferrerTypeNoReferer  ExtractBatchParamsInputReferrerType = "no-referer"
	ExtractBatchParamsInputReferrerTypeSameOrigin ExtractBatchParamsInputReferrerType = "same-origin"
	ExtractBatchParamsInputReferrerTypeGoogle     ExtractBatchParamsInputReferrerType = "google"
	ExtractBatchParamsInputReferrerTypeBing       ExtractBatchParamsInputReferrerType = "bing"
	ExtractBatchParamsInputReferrerTypeFacebook   ExtractBatchParamsInputReferrerType = "facebook"
	ExtractBatchParamsInputReferrerTypeTwitter    ExtractBatchParamsInputReferrerType = "twitter"
	ExtractBatchParamsInputReferrerTypeInstagram  ExtractBatchParamsInputReferrerType = "instagram"
)

type ExtractBatchParamsInputSession struct {
	ID                  param.Opt[string]  `json:"id,omitzero"`
	PrefetchUserbrowser param.Opt[bool]    `json:"prefetch_userbrowser,omitzero"`
	Retry               param.Opt[bool]    `json:"retry,omitzero"`
	Timeout             param.Opt[float64] `json:"timeout,omitzero"`
	paramObj
}

func (r ExtractBatchParamsInputSession) MarshalJSON() (data []byte, err error) {
	type shadow ExtractBatchParamsInputSession
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractBatchParamsInputSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsInputSkillUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsInputSkillUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractBatchParamsInputSkillUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsInputSkillUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Shared parameters applied to the entire batch. Can include extraction parameters
// and async/storage settings.
type ExtractBatchParamsSharedInputs struct {
	// URL to call back when async operation completes
	CallbackURL param.Opt[string] `json:"callback_url,omitzero"`
	// City for geolocation
	City param.Opt[string] `json:"city,omitzero"`
	// Whether to automatically handle cookie consent headers
	ConsentHeader param.Opt[bool] `json:"consent_header,omitzero"`
	// Whether to use HTTP/2 protocol
	Http2 param.Opt[bool] `json:"http2,omitzero"`
	// Whether to emulate XMLHttpRequest behavior
	IsXhr param.Opt[bool] `json:"is_xhr,omitzero"`
	// Whether to parse the response content
	Parse param.Opt[bool] `json:"parse,omitzero"`
	// Whether to render JavaScript content using a browser
	Render param.Opt[bool] `json:"render,omitzero"`
	// Request timeout in milliseconds
	RequestTimeout param.Opt[float64] `json:"request_timeout,omitzero"`
	// Whether to compress stored data
	StorageCompress param.Opt[bool] `json:"storage_compress,omitzero"`
	// Custom name for the stored object
	StorageObjectName param.Opt[string] `json:"storage_object_name,omitzero"`
	// Type of storage to use for results
	StorageType param.Opt[string] `json:"storage_type,omitzero"`
	// URL for storage location
	StorageURL param.Opt[string] `json:"storage_url,omitzero"`
	// User-defined tag for request identification
	Tag param.Opt[string] `json:"tag,omitzero"`
	// Target URL to scrape
	URL param.Opt[string] `json:"url,omitzero"`
	// Browser type to emulate
	Browser ExtractBatchParamsSharedInputsBrowserUnion `json:"browser,omitzero"`
	// Array of browser automation actions to execute sequentially
	BrowserActions []ExtractBatchParamsSharedInputsBrowserActionUnion `json:"browser_actions,omitzero"`
	// Browser cookies as array of cookie objects
	Cookies ExtractBatchParamsSharedInputsCookiesUnion `json:"cookies,omitzero"`
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
	Country ExtractBatchParamsSharedInputsCountry `json:"country,omitzero"`
	// Device type for browser emulation
	//
	// Any of "desktop", "mobile", "tablet".
	Device string `json:"device,omitzero"`
	// Browser driver to use
	//
	// Any of "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro".
	Driver string `json:"driver,omitzero"`
	// Expected HTTP status codes for successful requests
	ExpectedStatusCodes []int64 `json:"expected_status_codes,omitzero"`
	// List of acceptable response formats in order of preference
	//
	// Any of "html", "markdown", "screenshot", "headers".
	Formats []string `json:"formats,omitzero"`
	// Custom HTTP headers to include in the request
	Headers map[string]ExtractBatchParamsSharedInputsHeaderUnion `json:"headers,omitzero"`
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
	Locale ExtractBatchParamsSharedInputsLocale `json:"locale,omitzero"`
	// HTTP method for the request
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method,omitzero"`
	// Filters for capturing network traffic
	NetworkCapture []ExtractBatchParamsSharedInputsNetworkCapture `json:"network_capture,omitzero"`
	// Operating system to emulate
	//
	// Any of "windows", "mac os", "linux", "android", "ios".
	Os string `json:"os,omitzero"`
	// Custom parser configuration as a key-value map
	Parser ExtractBatchParamsSharedInputsParserUnion `json:"parser,omitzero"`
	// Referrer policy for the request
	//
	// Any of "random", "no-referer", "same-origin", "google", "bing", "facebook",
	// "twitter", "instagram".
	ReferrerType ExtractBatchParamsSharedInputsReferrerType `json:"referrer_type,omitzero"`
	Session      ExtractBatchParamsSharedInputsSession      `json:"session,omitzero"`
	// Skills or capabilities required for the request
	Skill ExtractBatchParamsSharedInputsSkillUnion `json:"skill,omitzero"`
	// US state for geolocation (only valid when country is US)
	//
	// Any of "AL", "AK", "AS", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA",
	// "GU", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI",
	// "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "MP",
	// "OH", "OK", "OR", "PA", "PR", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA",
	// "VI", "WA", "WV", "WI", "WY".
	State string `json:"state,omitzero"`
	paramObj
}

func (r ExtractBatchParamsSharedInputs) MarshalJSON() (data []byte, err error) {
	type shadow ExtractBatchParamsSharedInputs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractBatchParamsSharedInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractBatchParamsSharedInputs](
		"device", "desktop", "mobile", "tablet",
	)
	apijson.RegisterFieldValidator[ExtractBatchParamsSharedInputs](
		"driver", "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro",
	)
	apijson.RegisterFieldValidator[ExtractBatchParamsSharedInputs](
		"method", "GET", "POST", "PUT", "PATCH", "DELETE",
	)
	apijson.RegisterFieldValidator[ExtractBatchParamsSharedInputs](
		"os", "windows", "mac os", "linux", "android", "ios",
	)
	apijson.RegisterFieldValidator[ExtractBatchParamsSharedInputs](
		"state", "AL", "AK", "AS", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA", "GU", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "MP", "OH", "OK", "OR", "PA", "PR", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "VI", "WA", "WV", "WI", "WY",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsSharedInputsBrowserUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractBatchsSharedInputsBrowserString)
	OfExtractBatchsSharedInputsBrowserString param.Opt[string]                            `json:",omitzero,inline"`
	OfExtractBatchsSharedInputsBrowserObject *ExtractBatchParamsSharedInputsBrowserObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsSharedInputsBrowserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractBatchsSharedInputsBrowserString, u.OfExtractBatchsSharedInputsBrowserObject)
}
func (u *ExtractBatchParamsSharedInputsBrowserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsSharedInputsBrowserUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractBatchsSharedInputsBrowserString) {
		return &u.OfExtractBatchsSharedInputsBrowserString
	} else if !param.IsOmitted(u.OfExtractBatchsSharedInputsBrowserObject) {
		return u.OfExtractBatchsSharedInputsBrowserObject
	}
	return nil
}

// Browser type to emulate
type ExtractBatchParamsSharedInputsBrowserString string

const (
	ExtractBatchParamsSharedInputsBrowserStringChrome  ExtractBatchParamsSharedInputsBrowserString = "chrome"
	ExtractBatchParamsSharedInputsBrowserStringFirefox ExtractBatchParamsSharedInputsBrowserString = "firefox"
)

// The property Name is required.
type ExtractBatchParamsSharedInputsBrowserObject struct {
	// Any of "chrome", "firefox".
	Name string `json:"name,omitzero" api:"required"`
	// Specific browser version to emulate
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r ExtractBatchParamsSharedInputsBrowserObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractBatchParamsSharedInputsBrowserObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractBatchParamsSharedInputsBrowserObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractBatchParamsSharedInputsBrowserObject](
		"name", "chrome", "firefox",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsSharedInputsBrowserActionUnion struct {
	OfAutoScrollAction        *shared.AutoScrollActionParam        `json:",omitzero,inline"`
	OfClickAction             *shared.ClickActionParam             `json:",omitzero,inline"`
	OfEvalAction              *shared.EvalActionParam              `json:",omitzero,inline"`
	OfFetchAction             *shared.FetchActionParam             `json:",omitzero,inline"`
	OfFillAction              *shared.FillActionParam              `json:",omitzero,inline"`
	OfGetCookiesAction        *shared.GetCookiesActionParam        `json:",omitzero,inline"`
	OfGotoAction              *shared.GotoActionParam              `json:",omitzero,inline"`
	OfPressAction             *shared.PressActionParam             `json:",omitzero,inline"`
	OfScreenshotAction        *shared.ScreenshotActionParam        `json:",omitzero,inline"`
	OfScrollAction            *shared.ScrollActionParam            `json:",omitzero,inline"`
	OfWaitAction              *shared.WaitActionParam              `json:",omitzero,inline"`
	OfWaitForElementAction    *shared.WaitForElementActionParam    `json:",omitzero,inline"`
	OfWaitForNavigationAction *shared.WaitForNavigationActionParam `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsSharedInputsBrowserActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAutoScrollAction,
		u.OfClickAction,
		u.OfEvalAction,
		u.OfFetchAction,
		u.OfFillAction,
		u.OfGetCookiesAction,
		u.OfGotoAction,
		u.OfPressAction,
		u.OfScreenshotAction,
		u.OfScrollAction,
		u.OfWaitAction,
		u.OfWaitForElementAction,
		u.OfWaitForNavigationAction)
}
func (u *ExtractBatchParamsSharedInputsBrowserActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsSharedInputsBrowserActionUnion) asAny() any {
	if !param.IsOmitted(u.OfAutoScrollAction) {
		return u.OfAutoScrollAction
	} else if !param.IsOmitted(u.OfClickAction) {
		return u.OfClickAction
	} else if !param.IsOmitted(u.OfEvalAction) {
		return u.OfEvalAction
	} else if !param.IsOmitted(u.OfFetchAction) {
		return u.OfFetchAction
	} else if !param.IsOmitted(u.OfFillAction) {
		return u.OfFillAction
	} else if !param.IsOmitted(u.OfGetCookiesAction) {
		return u.OfGetCookiesAction
	} else if !param.IsOmitted(u.OfGotoAction) {
		return u.OfGotoAction
	} else if !param.IsOmitted(u.OfPressAction) {
		return u.OfPressAction
	} else if !param.IsOmitted(u.OfScreenshotAction) {
		return u.OfScreenshotAction
	} else if !param.IsOmitted(u.OfScrollAction) {
		return u.OfScrollAction
	} else if !param.IsOmitted(u.OfWaitAction) {
		return u.OfWaitAction
	} else if !param.IsOmitted(u.OfWaitForElementAction) {
		return u.OfWaitForElementAction
	} else if !param.IsOmitted(u.OfWaitForNavigationAction) {
		return u.OfWaitForNavigationAction
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsSharedInputsCookiesUnion struct {
	OfExtractBatchsSharedInputsCookiesArray []ExtractBatchParamsSharedInputsCookiesArrayItem `json:",omitzero,inline"`
	OfString                                param.Opt[string]                                `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsSharedInputsCookiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractBatchsSharedInputsCookiesArray, u.OfString)
}
func (u *ExtractBatchParamsSharedInputsCookiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsSharedInputsCookiesUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractBatchsSharedInputsCookiesArray) {
		return &u.OfExtractBatchsSharedInputsCookiesArray
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type ExtractBatchParamsSharedInputsCookiesArrayItem struct {
	Creation      param.Opt[string]                                         `json:"creation,omitzero"`
	Domain        param.Opt[string]                                         `json:"domain,omitzero"`
	HostOnly      param.Opt[bool]                                           `json:"hostOnly,omitzero"`
	HTTPOnly      param.Opt[bool]                                           `json:"httpOnly,omitzero"`
	LastAccessed  param.Opt[string]                                         `json:"lastAccessed,omitzero"`
	Path          param.Opt[string]                                         `json:"path,omitzero"`
	PathIsDefault param.Opt[bool]                                           `json:"pathIsDefault,omitzero"`
	Expires       param.Opt[string]                                         `json:"expires,omitzero"`
	Name          param.Opt[string]                                         `json:"name,omitzero"`
	Secure        param.Opt[bool]                                           `json:"secure,omitzero"`
	Value         param.Opt[string]                                         `json:"value,omitzero"`
	Extensions    []string                                                  `json:"extensions,omitzero"`
	MaxAge        ExtractBatchParamsSharedInputsCookiesArrayItemMaxAgeUnion `json:"maxAge,omitzero"`
	// Any of "strict", "lax", "none".
	SameSite    string         `json:"sameSite,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r ExtractBatchParamsSharedInputsCookiesArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow ExtractBatchParamsSharedInputsCookiesArrayItem
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ExtractBatchParamsSharedInputsCookiesArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractBatchParamsSharedInputsCookiesArrayItem](
		"sameSite", "strict", "lax", "none",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsSharedInputsCookiesArrayItemMaxAgeUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractBatchsSharedInputsCookiesArrayItemMaxAgeString)
	OfExtractBatchsSharedInputsCookiesArrayItemMaxAgeString param.Opt[ExtractBatchParamsSharedInputsCookiesArrayItemMaxAgeString] `json:",omitzero,inline"`
	OfFloat                                                 param.Opt[float64]                                                    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsSharedInputsCookiesArrayItemMaxAgeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractBatchsSharedInputsCookiesArrayItemMaxAgeString, u.OfFloat)
}
func (u *ExtractBatchParamsSharedInputsCookiesArrayItemMaxAgeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsSharedInputsCookiesArrayItemMaxAgeUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractBatchsSharedInputsCookiesArrayItemMaxAgeString) {
		return &u.OfExtractBatchsSharedInputsCookiesArrayItemMaxAgeString
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type ExtractBatchParamsSharedInputsCookiesArrayItemMaxAgeString string

const (
	ExtractBatchParamsSharedInputsCookiesArrayItemMaxAgeStringInfinity      ExtractBatchParamsSharedInputsCookiesArrayItemMaxAgeString = "Infinity"
	ExtractBatchParamsSharedInputsCookiesArrayItemMaxAgeStringMinusInfinity ExtractBatchParamsSharedInputsCookiesArrayItemMaxAgeString = "-Infinity"
)

// Country code for geolocation and proxy selection
type ExtractBatchParamsSharedInputsCountry string

const (
	ExtractBatchParamsSharedInputsCountryAd  ExtractBatchParamsSharedInputsCountry = "AD"
	ExtractBatchParamsSharedInputsCountryAe  ExtractBatchParamsSharedInputsCountry = "AE"
	ExtractBatchParamsSharedInputsCountryAf  ExtractBatchParamsSharedInputsCountry = "AF"
	ExtractBatchParamsSharedInputsCountryAg  ExtractBatchParamsSharedInputsCountry = "AG"
	ExtractBatchParamsSharedInputsCountryAI  ExtractBatchParamsSharedInputsCountry = "AI"
	ExtractBatchParamsSharedInputsCountryAl  ExtractBatchParamsSharedInputsCountry = "AL"
	ExtractBatchParamsSharedInputsCountryAm  ExtractBatchParamsSharedInputsCountry = "AM"
	ExtractBatchParamsSharedInputsCountryAo  ExtractBatchParamsSharedInputsCountry = "AO"
	ExtractBatchParamsSharedInputsCountryAq  ExtractBatchParamsSharedInputsCountry = "AQ"
	ExtractBatchParamsSharedInputsCountryAr  ExtractBatchParamsSharedInputsCountry = "AR"
	ExtractBatchParamsSharedInputsCountryAs  ExtractBatchParamsSharedInputsCountry = "AS"
	ExtractBatchParamsSharedInputsCountryAt  ExtractBatchParamsSharedInputsCountry = "AT"
	ExtractBatchParamsSharedInputsCountryAu  ExtractBatchParamsSharedInputsCountry = "AU"
	ExtractBatchParamsSharedInputsCountryAw  ExtractBatchParamsSharedInputsCountry = "AW"
	ExtractBatchParamsSharedInputsCountryAx  ExtractBatchParamsSharedInputsCountry = "AX"
	ExtractBatchParamsSharedInputsCountryAz  ExtractBatchParamsSharedInputsCountry = "AZ"
	ExtractBatchParamsSharedInputsCountryBa  ExtractBatchParamsSharedInputsCountry = "BA"
	ExtractBatchParamsSharedInputsCountryBb  ExtractBatchParamsSharedInputsCountry = "BB"
	ExtractBatchParamsSharedInputsCountryBd  ExtractBatchParamsSharedInputsCountry = "BD"
	ExtractBatchParamsSharedInputsCountryBe  ExtractBatchParamsSharedInputsCountry = "BE"
	ExtractBatchParamsSharedInputsCountryBf  ExtractBatchParamsSharedInputsCountry = "BF"
	ExtractBatchParamsSharedInputsCountryBg  ExtractBatchParamsSharedInputsCountry = "BG"
	ExtractBatchParamsSharedInputsCountryBh  ExtractBatchParamsSharedInputsCountry = "BH"
	ExtractBatchParamsSharedInputsCountryBi  ExtractBatchParamsSharedInputsCountry = "BI"
	ExtractBatchParamsSharedInputsCountryBj  ExtractBatchParamsSharedInputsCountry = "BJ"
	ExtractBatchParamsSharedInputsCountryBl  ExtractBatchParamsSharedInputsCountry = "BL"
	ExtractBatchParamsSharedInputsCountryBm  ExtractBatchParamsSharedInputsCountry = "BM"
	ExtractBatchParamsSharedInputsCountryBn  ExtractBatchParamsSharedInputsCountry = "BN"
	ExtractBatchParamsSharedInputsCountryBo  ExtractBatchParamsSharedInputsCountry = "BO"
	ExtractBatchParamsSharedInputsCountryBq  ExtractBatchParamsSharedInputsCountry = "BQ"
	ExtractBatchParamsSharedInputsCountryBr  ExtractBatchParamsSharedInputsCountry = "BR"
	ExtractBatchParamsSharedInputsCountryBs  ExtractBatchParamsSharedInputsCountry = "BS"
	ExtractBatchParamsSharedInputsCountryBt  ExtractBatchParamsSharedInputsCountry = "BT"
	ExtractBatchParamsSharedInputsCountryBv  ExtractBatchParamsSharedInputsCountry = "BV"
	ExtractBatchParamsSharedInputsCountryBw  ExtractBatchParamsSharedInputsCountry = "BW"
	ExtractBatchParamsSharedInputsCountryBy  ExtractBatchParamsSharedInputsCountry = "BY"
	ExtractBatchParamsSharedInputsCountryBz  ExtractBatchParamsSharedInputsCountry = "BZ"
	ExtractBatchParamsSharedInputsCountryCa  ExtractBatchParamsSharedInputsCountry = "CA"
	ExtractBatchParamsSharedInputsCountryCc  ExtractBatchParamsSharedInputsCountry = "CC"
	ExtractBatchParamsSharedInputsCountryCd  ExtractBatchParamsSharedInputsCountry = "CD"
	ExtractBatchParamsSharedInputsCountryCf  ExtractBatchParamsSharedInputsCountry = "CF"
	ExtractBatchParamsSharedInputsCountryCg  ExtractBatchParamsSharedInputsCountry = "CG"
	ExtractBatchParamsSharedInputsCountryCh  ExtractBatchParamsSharedInputsCountry = "CH"
	ExtractBatchParamsSharedInputsCountryCi  ExtractBatchParamsSharedInputsCountry = "CI"
	ExtractBatchParamsSharedInputsCountryCk  ExtractBatchParamsSharedInputsCountry = "CK"
	ExtractBatchParamsSharedInputsCountryCl  ExtractBatchParamsSharedInputsCountry = "CL"
	ExtractBatchParamsSharedInputsCountryCm  ExtractBatchParamsSharedInputsCountry = "CM"
	ExtractBatchParamsSharedInputsCountryCn  ExtractBatchParamsSharedInputsCountry = "CN"
	ExtractBatchParamsSharedInputsCountryCo  ExtractBatchParamsSharedInputsCountry = "CO"
	ExtractBatchParamsSharedInputsCountryCr  ExtractBatchParamsSharedInputsCountry = "CR"
	ExtractBatchParamsSharedInputsCountryCu  ExtractBatchParamsSharedInputsCountry = "CU"
	ExtractBatchParamsSharedInputsCountryCv  ExtractBatchParamsSharedInputsCountry = "CV"
	ExtractBatchParamsSharedInputsCountryCw  ExtractBatchParamsSharedInputsCountry = "CW"
	ExtractBatchParamsSharedInputsCountryCx  ExtractBatchParamsSharedInputsCountry = "CX"
	ExtractBatchParamsSharedInputsCountryCy  ExtractBatchParamsSharedInputsCountry = "CY"
	ExtractBatchParamsSharedInputsCountryCz  ExtractBatchParamsSharedInputsCountry = "CZ"
	ExtractBatchParamsSharedInputsCountryDe  ExtractBatchParamsSharedInputsCountry = "DE"
	ExtractBatchParamsSharedInputsCountryDj  ExtractBatchParamsSharedInputsCountry = "DJ"
	ExtractBatchParamsSharedInputsCountryDk  ExtractBatchParamsSharedInputsCountry = "DK"
	ExtractBatchParamsSharedInputsCountryDm  ExtractBatchParamsSharedInputsCountry = "DM"
	ExtractBatchParamsSharedInputsCountryDo  ExtractBatchParamsSharedInputsCountry = "DO"
	ExtractBatchParamsSharedInputsCountryDz  ExtractBatchParamsSharedInputsCountry = "DZ"
	ExtractBatchParamsSharedInputsCountryEc  ExtractBatchParamsSharedInputsCountry = "EC"
	ExtractBatchParamsSharedInputsCountryEe  ExtractBatchParamsSharedInputsCountry = "EE"
	ExtractBatchParamsSharedInputsCountryEg  ExtractBatchParamsSharedInputsCountry = "EG"
	ExtractBatchParamsSharedInputsCountryEh  ExtractBatchParamsSharedInputsCountry = "EH"
	ExtractBatchParamsSharedInputsCountryEr  ExtractBatchParamsSharedInputsCountry = "ER"
	ExtractBatchParamsSharedInputsCountryEs  ExtractBatchParamsSharedInputsCountry = "ES"
	ExtractBatchParamsSharedInputsCountryEt  ExtractBatchParamsSharedInputsCountry = "ET"
	ExtractBatchParamsSharedInputsCountryFi  ExtractBatchParamsSharedInputsCountry = "FI"
	ExtractBatchParamsSharedInputsCountryFj  ExtractBatchParamsSharedInputsCountry = "FJ"
	ExtractBatchParamsSharedInputsCountryFk  ExtractBatchParamsSharedInputsCountry = "FK"
	ExtractBatchParamsSharedInputsCountryFm  ExtractBatchParamsSharedInputsCountry = "FM"
	ExtractBatchParamsSharedInputsCountryFo  ExtractBatchParamsSharedInputsCountry = "FO"
	ExtractBatchParamsSharedInputsCountryFr  ExtractBatchParamsSharedInputsCountry = "FR"
	ExtractBatchParamsSharedInputsCountryGa  ExtractBatchParamsSharedInputsCountry = "GA"
	ExtractBatchParamsSharedInputsCountryGB  ExtractBatchParamsSharedInputsCountry = "GB"
	ExtractBatchParamsSharedInputsCountryGd  ExtractBatchParamsSharedInputsCountry = "GD"
	ExtractBatchParamsSharedInputsCountryGe  ExtractBatchParamsSharedInputsCountry = "GE"
	ExtractBatchParamsSharedInputsCountryGf  ExtractBatchParamsSharedInputsCountry = "GF"
	ExtractBatchParamsSharedInputsCountryGg  ExtractBatchParamsSharedInputsCountry = "GG"
	ExtractBatchParamsSharedInputsCountryGh  ExtractBatchParamsSharedInputsCountry = "GH"
	ExtractBatchParamsSharedInputsCountryGi  ExtractBatchParamsSharedInputsCountry = "GI"
	ExtractBatchParamsSharedInputsCountryGl  ExtractBatchParamsSharedInputsCountry = "GL"
	ExtractBatchParamsSharedInputsCountryGm  ExtractBatchParamsSharedInputsCountry = "GM"
	ExtractBatchParamsSharedInputsCountryGn  ExtractBatchParamsSharedInputsCountry = "GN"
	ExtractBatchParamsSharedInputsCountryGp  ExtractBatchParamsSharedInputsCountry = "GP"
	ExtractBatchParamsSharedInputsCountryGq  ExtractBatchParamsSharedInputsCountry = "GQ"
	ExtractBatchParamsSharedInputsCountryGr  ExtractBatchParamsSharedInputsCountry = "GR"
	ExtractBatchParamsSharedInputsCountryGs  ExtractBatchParamsSharedInputsCountry = "GS"
	ExtractBatchParamsSharedInputsCountryGt  ExtractBatchParamsSharedInputsCountry = "GT"
	ExtractBatchParamsSharedInputsCountryGu  ExtractBatchParamsSharedInputsCountry = "GU"
	ExtractBatchParamsSharedInputsCountryGw  ExtractBatchParamsSharedInputsCountry = "GW"
	ExtractBatchParamsSharedInputsCountryGy  ExtractBatchParamsSharedInputsCountry = "GY"
	ExtractBatchParamsSharedInputsCountryHk  ExtractBatchParamsSharedInputsCountry = "HK"
	ExtractBatchParamsSharedInputsCountryHm  ExtractBatchParamsSharedInputsCountry = "HM"
	ExtractBatchParamsSharedInputsCountryHn  ExtractBatchParamsSharedInputsCountry = "HN"
	ExtractBatchParamsSharedInputsCountryHr  ExtractBatchParamsSharedInputsCountry = "HR"
	ExtractBatchParamsSharedInputsCountryHt  ExtractBatchParamsSharedInputsCountry = "HT"
	ExtractBatchParamsSharedInputsCountryHu  ExtractBatchParamsSharedInputsCountry = "HU"
	ExtractBatchParamsSharedInputsCountryID  ExtractBatchParamsSharedInputsCountry = "ID"
	ExtractBatchParamsSharedInputsCountryIe  ExtractBatchParamsSharedInputsCountry = "IE"
	ExtractBatchParamsSharedInputsCountryIl  ExtractBatchParamsSharedInputsCountry = "IL"
	ExtractBatchParamsSharedInputsCountryIm  ExtractBatchParamsSharedInputsCountry = "IM"
	ExtractBatchParamsSharedInputsCountryIn  ExtractBatchParamsSharedInputsCountry = "IN"
	ExtractBatchParamsSharedInputsCountryIo  ExtractBatchParamsSharedInputsCountry = "IO"
	ExtractBatchParamsSharedInputsCountryIq  ExtractBatchParamsSharedInputsCountry = "IQ"
	ExtractBatchParamsSharedInputsCountryIr  ExtractBatchParamsSharedInputsCountry = "IR"
	ExtractBatchParamsSharedInputsCountryIs  ExtractBatchParamsSharedInputsCountry = "IS"
	ExtractBatchParamsSharedInputsCountryIt  ExtractBatchParamsSharedInputsCountry = "IT"
	ExtractBatchParamsSharedInputsCountryJe  ExtractBatchParamsSharedInputsCountry = "JE"
	ExtractBatchParamsSharedInputsCountryJm  ExtractBatchParamsSharedInputsCountry = "JM"
	ExtractBatchParamsSharedInputsCountryJo  ExtractBatchParamsSharedInputsCountry = "JO"
	ExtractBatchParamsSharedInputsCountryJp  ExtractBatchParamsSharedInputsCountry = "JP"
	ExtractBatchParamsSharedInputsCountryKe  ExtractBatchParamsSharedInputsCountry = "KE"
	ExtractBatchParamsSharedInputsCountryKg  ExtractBatchParamsSharedInputsCountry = "KG"
	ExtractBatchParamsSharedInputsCountryKh  ExtractBatchParamsSharedInputsCountry = "KH"
	ExtractBatchParamsSharedInputsCountryKi  ExtractBatchParamsSharedInputsCountry = "KI"
	ExtractBatchParamsSharedInputsCountryKm  ExtractBatchParamsSharedInputsCountry = "KM"
	ExtractBatchParamsSharedInputsCountryKn  ExtractBatchParamsSharedInputsCountry = "KN"
	ExtractBatchParamsSharedInputsCountryKp  ExtractBatchParamsSharedInputsCountry = "KP"
	ExtractBatchParamsSharedInputsCountryKr  ExtractBatchParamsSharedInputsCountry = "KR"
	ExtractBatchParamsSharedInputsCountryKw  ExtractBatchParamsSharedInputsCountry = "KW"
	ExtractBatchParamsSharedInputsCountryKy  ExtractBatchParamsSharedInputsCountry = "KY"
	ExtractBatchParamsSharedInputsCountryKz  ExtractBatchParamsSharedInputsCountry = "KZ"
	ExtractBatchParamsSharedInputsCountryLa  ExtractBatchParamsSharedInputsCountry = "LA"
	ExtractBatchParamsSharedInputsCountryLb  ExtractBatchParamsSharedInputsCountry = "LB"
	ExtractBatchParamsSharedInputsCountryLc  ExtractBatchParamsSharedInputsCountry = "LC"
	ExtractBatchParamsSharedInputsCountryLi  ExtractBatchParamsSharedInputsCountry = "LI"
	ExtractBatchParamsSharedInputsCountryLk  ExtractBatchParamsSharedInputsCountry = "LK"
	ExtractBatchParamsSharedInputsCountryLr  ExtractBatchParamsSharedInputsCountry = "LR"
	ExtractBatchParamsSharedInputsCountryLs  ExtractBatchParamsSharedInputsCountry = "LS"
	ExtractBatchParamsSharedInputsCountryLt  ExtractBatchParamsSharedInputsCountry = "LT"
	ExtractBatchParamsSharedInputsCountryLu  ExtractBatchParamsSharedInputsCountry = "LU"
	ExtractBatchParamsSharedInputsCountryLv  ExtractBatchParamsSharedInputsCountry = "LV"
	ExtractBatchParamsSharedInputsCountryLy  ExtractBatchParamsSharedInputsCountry = "LY"
	ExtractBatchParamsSharedInputsCountryMa  ExtractBatchParamsSharedInputsCountry = "MA"
	ExtractBatchParamsSharedInputsCountryMc  ExtractBatchParamsSharedInputsCountry = "MC"
	ExtractBatchParamsSharedInputsCountryMd  ExtractBatchParamsSharedInputsCountry = "MD"
	ExtractBatchParamsSharedInputsCountryMe  ExtractBatchParamsSharedInputsCountry = "ME"
	ExtractBatchParamsSharedInputsCountryMf  ExtractBatchParamsSharedInputsCountry = "MF"
	ExtractBatchParamsSharedInputsCountryMg  ExtractBatchParamsSharedInputsCountry = "MG"
	ExtractBatchParamsSharedInputsCountryMh  ExtractBatchParamsSharedInputsCountry = "MH"
	ExtractBatchParamsSharedInputsCountryMk  ExtractBatchParamsSharedInputsCountry = "MK"
	ExtractBatchParamsSharedInputsCountryMl  ExtractBatchParamsSharedInputsCountry = "ML"
	ExtractBatchParamsSharedInputsCountryMm  ExtractBatchParamsSharedInputsCountry = "MM"
	ExtractBatchParamsSharedInputsCountryMn  ExtractBatchParamsSharedInputsCountry = "MN"
	ExtractBatchParamsSharedInputsCountryMo  ExtractBatchParamsSharedInputsCountry = "MO"
	ExtractBatchParamsSharedInputsCountryMp  ExtractBatchParamsSharedInputsCountry = "MP"
	ExtractBatchParamsSharedInputsCountryMq  ExtractBatchParamsSharedInputsCountry = "MQ"
	ExtractBatchParamsSharedInputsCountryMr  ExtractBatchParamsSharedInputsCountry = "MR"
	ExtractBatchParamsSharedInputsCountryMs  ExtractBatchParamsSharedInputsCountry = "MS"
	ExtractBatchParamsSharedInputsCountryMt  ExtractBatchParamsSharedInputsCountry = "MT"
	ExtractBatchParamsSharedInputsCountryMu  ExtractBatchParamsSharedInputsCountry = "MU"
	ExtractBatchParamsSharedInputsCountryMv  ExtractBatchParamsSharedInputsCountry = "MV"
	ExtractBatchParamsSharedInputsCountryMw  ExtractBatchParamsSharedInputsCountry = "MW"
	ExtractBatchParamsSharedInputsCountryMx  ExtractBatchParamsSharedInputsCountry = "MX"
	ExtractBatchParamsSharedInputsCountryMy  ExtractBatchParamsSharedInputsCountry = "MY"
	ExtractBatchParamsSharedInputsCountryMz  ExtractBatchParamsSharedInputsCountry = "MZ"
	ExtractBatchParamsSharedInputsCountryNa  ExtractBatchParamsSharedInputsCountry = "NA"
	ExtractBatchParamsSharedInputsCountryNc  ExtractBatchParamsSharedInputsCountry = "NC"
	ExtractBatchParamsSharedInputsCountryNe  ExtractBatchParamsSharedInputsCountry = "NE"
	ExtractBatchParamsSharedInputsCountryNf  ExtractBatchParamsSharedInputsCountry = "NF"
	ExtractBatchParamsSharedInputsCountryNg  ExtractBatchParamsSharedInputsCountry = "NG"
	ExtractBatchParamsSharedInputsCountryNi  ExtractBatchParamsSharedInputsCountry = "NI"
	ExtractBatchParamsSharedInputsCountryNl  ExtractBatchParamsSharedInputsCountry = "NL"
	ExtractBatchParamsSharedInputsCountryNo  ExtractBatchParamsSharedInputsCountry = "NO"
	ExtractBatchParamsSharedInputsCountryNp  ExtractBatchParamsSharedInputsCountry = "NP"
	ExtractBatchParamsSharedInputsCountryNr  ExtractBatchParamsSharedInputsCountry = "NR"
	ExtractBatchParamsSharedInputsCountryNu  ExtractBatchParamsSharedInputsCountry = "NU"
	ExtractBatchParamsSharedInputsCountryNz  ExtractBatchParamsSharedInputsCountry = "NZ"
	ExtractBatchParamsSharedInputsCountryOm  ExtractBatchParamsSharedInputsCountry = "OM"
	ExtractBatchParamsSharedInputsCountryPa  ExtractBatchParamsSharedInputsCountry = "PA"
	ExtractBatchParamsSharedInputsCountryPe  ExtractBatchParamsSharedInputsCountry = "PE"
	ExtractBatchParamsSharedInputsCountryPf  ExtractBatchParamsSharedInputsCountry = "PF"
	ExtractBatchParamsSharedInputsCountryPg  ExtractBatchParamsSharedInputsCountry = "PG"
	ExtractBatchParamsSharedInputsCountryPh  ExtractBatchParamsSharedInputsCountry = "PH"
	ExtractBatchParamsSharedInputsCountryPk  ExtractBatchParamsSharedInputsCountry = "PK"
	ExtractBatchParamsSharedInputsCountryPl  ExtractBatchParamsSharedInputsCountry = "PL"
	ExtractBatchParamsSharedInputsCountryPm  ExtractBatchParamsSharedInputsCountry = "PM"
	ExtractBatchParamsSharedInputsCountryPn  ExtractBatchParamsSharedInputsCountry = "PN"
	ExtractBatchParamsSharedInputsCountryPr  ExtractBatchParamsSharedInputsCountry = "PR"
	ExtractBatchParamsSharedInputsCountryPs  ExtractBatchParamsSharedInputsCountry = "PS"
	ExtractBatchParamsSharedInputsCountryPt  ExtractBatchParamsSharedInputsCountry = "PT"
	ExtractBatchParamsSharedInputsCountryPw  ExtractBatchParamsSharedInputsCountry = "PW"
	ExtractBatchParamsSharedInputsCountryPy  ExtractBatchParamsSharedInputsCountry = "PY"
	ExtractBatchParamsSharedInputsCountryQa  ExtractBatchParamsSharedInputsCountry = "QA"
	ExtractBatchParamsSharedInputsCountryRe  ExtractBatchParamsSharedInputsCountry = "RE"
	ExtractBatchParamsSharedInputsCountryRo  ExtractBatchParamsSharedInputsCountry = "RO"
	ExtractBatchParamsSharedInputsCountryRs  ExtractBatchParamsSharedInputsCountry = "RS"
	ExtractBatchParamsSharedInputsCountryRu  ExtractBatchParamsSharedInputsCountry = "RU"
	ExtractBatchParamsSharedInputsCountryRw  ExtractBatchParamsSharedInputsCountry = "RW"
	ExtractBatchParamsSharedInputsCountrySa  ExtractBatchParamsSharedInputsCountry = "SA"
	ExtractBatchParamsSharedInputsCountrySb  ExtractBatchParamsSharedInputsCountry = "SB"
	ExtractBatchParamsSharedInputsCountrySc  ExtractBatchParamsSharedInputsCountry = "SC"
	ExtractBatchParamsSharedInputsCountrySd  ExtractBatchParamsSharedInputsCountry = "SD"
	ExtractBatchParamsSharedInputsCountrySe  ExtractBatchParamsSharedInputsCountry = "SE"
	ExtractBatchParamsSharedInputsCountrySg  ExtractBatchParamsSharedInputsCountry = "SG"
	ExtractBatchParamsSharedInputsCountrySh  ExtractBatchParamsSharedInputsCountry = "SH"
	ExtractBatchParamsSharedInputsCountrySi  ExtractBatchParamsSharedInputsCountry = "SI"
	ExtractBatchParamsSharedInputsCountrySj  ExtractBatchParamsSharedInputsCountry = "SJ"
	ExtractBatchParamsSharedInputsCountrySk  ExtractBatchParamsSharedInputsCountry = "SK"
	ExtractBatchParamsSharedInputsCountrySl  ExtractBatchParamsSharedInputsCountry = "SL"
	ExtractBatchParamsSharedInputsCountrySm  ExtractBatchParamsSharedInputsCountry = "SM"
	ExtractBatchParamsSharedInputsCountrySn  ExtractBatchParamsSharedInputsCountry = "SN"
	ExtractBatchParamsSharedInputsCountrySo  ExtractBatchParamsSharedInputsCountry = "SO"
	ExtractBatchParamsSharedInputsCountrySr  ExtractBatchParamsSharedInputsCountry = "SR"
	ExtractBatchParamsSharedInputsCountrySS  ExtractBatchParamsSharedInputsCountry = "SS"
	ExtractBatchParamsSharedInputsCountrySt  ExtractBatchParamsSharedInputsCountry = "ST"
	ExtractBatchParamsSharedInputsCountrySv  ExtractBatchParamsSharedInputsCountry = "SV"
	ExtractBatchParamsSharedInputsCountrySx  ExtractBatchParamsSharedInputsCountry = "SX"
	ExtractBatchParamsSharedInputsCountrySy  ExtractBatchParamsSharedInputsCountry = "SY"
	ExtractBatchParamsSharedInputsCountrySz  ExtractBatchParamsSharedInputsCountry = "SZ"
	ExtractBatchParamsSharedInputsCountryTc  ExtractBatchParamsSharedInputsCountry = "TC"
	ExtractBatchParamsSharedInputsCountryTd  ExtractBatchParamsSharedInputsCountry = "TD"
	ExtractBatchParamsSharedInputsCountryTf  ExtractBatchParamsSharedInputsCountry = "TF"
	ExtractBatchParamsSharedInputsCountryTg  ExtractBatchParamsSharedInputsCountry = "TG"
	ExtractBatchParamsSharedInputsCountryTh  ExtractBatchParamsSharedInputsCountry = "TH"
	ExtractBatchParamsSharedInputsCountryTj  ExtractBatchParamsSharedInputsCountry = "TJ"
	ExtractBatchParamsSharedInputsCountryTk  ExtractBatchParamsSharedInputsCountry = "TK"
	ExtractBatchParamsSharedInputsCountryTl  ExtractBatchParamsSharedInputsCountry = "TL"
	ExtractBatchParamsSharedInputsCountryTm  ExtractBatchParamsSharedInputsCountry = "TM"
	ExtractBatchParamsSharedInputsCountryTn  ExtractBatchParamsSharedInputsCountry = "TN"
	ExtractBatchParamsSharedInputsCountryTo  ExtractBatchParamsSharedInputsCountry = "TO"
	ExtractBatchParamsSharedInputsCountryTr  ExtractBatchParamsSharedInputsCountry = "TR"
	ExtractBatchParamsSharedInputsCountryTt  ExtractBatchParamsSharedInputsCountry = "TT"
	ExtractBatchParamsSharedInputsCountryTv  ExtractBatchParamsSharedInputsCountry = "TV"
	ExtractBatchParamsSharedInputsCountryTw  ExtractBatchParamsSharedInputsCountry = "TW"
	ExtractBatchParamsSharedInputsCountryTz  ExtractBatchParamsSharedInputsCountry = "TZ"
	ExtractBatchParamsSharedInputsCountryUa  ExtractBatchParamsSharedInputsCountry = "UA"
	ExtractBatchParamsSharedInputsCountryUg  ExtractBatchParamsSharedInputsCountry = "UG"
	ExtractBatchParamsSharedInputsCountryUm  ExtractBatchParamsSharedInputsCountry = "UM"
	ExtractBatchParamsSharedInputsCountryUs  ExtractBatchParamsSharedInputsCountry = "US"
	ExtractBatchParamsSharedInputsCountryUy  ExtractBatchParamsSharedInputsCountry = "UY"
	ExtractBatchParamsSharedInputsCountryUz  ExtractBatchParamsSharedInputsCountry = "UZ"
	ExtractBatchParamsSharedInputsCountryVa  ExtractBatchParamsSharedInputsCountry = "VA"
	ExtractBatchParamsSharedInputsCountryVc  ExtractBatchParamsSharedInputsCountry = "VC"
	ExtractBatchParamsSharedInputsCountryVe  ExtractBatchParamsSharedInputsCountry = "VE"
	ExtractBatchParamsSharedInputsCountryVg  ExtractBatchParamsSharedInputsCountry = "VG"
	ExtractBatchParamsSharedInputsCountryVi  ExtractBatchParamsSharedInputsCountry = "VI"
	ExtractBatchParamsSharedInputsCountryVn  ExtractBatchParamsSharedInputsCountry = "VN"
	ExtractBatchParamsSharedInputsCountryVu  ExtractBatchParamsSharedInputsCountry = "VU"
	ExtractBatchParamsSharedInputsCountryWf  ExtractBatchParamsSharedInputsCountry = "WF"
	ExtractBatchParamsSharedInputsCountryWs  ExtractBatchParamsSharedInputsCountry = "WS"
	ExtractBatchParamsSharedInputsCountryXk  ExtractBatchParamsSharedInputsCountry = "XK"
	ExtractBatchParamsSharedInputsCountryYe  ExtractBatchParamsSharedInputsCountry = "YE"
	ExtractBatchParamsSharedInputsCountryYt  ExtractBatchParamsSharedInputsCountry = "YT"
	ExtractBatchParamsSharedInputsCountryZa  ExtractBatchParamsSharedInputsCountry = "ZA"
	ExtractBatchParamsSharedInputsCountryZm  ExtractBatchParamsSharedInputsCountry = "ZM"
	ExtractBatchParamsSharedInputsCountryZw  ExtractBatchParamsSharedInputsCountry = "ZW"
	ExtractBatchParamsSharedInputsCountryAll ExtractBatchParamsSharedInputsCountry = "ALL"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsSharedInputsHeaderUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsSharedInputsHeaderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractBatchParamsSharedInputsHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsSharedInputsHeaderUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Locale for browser language and region settings
type ExtractBatchParamsSharedInputsLocale string

const (
	ExtractBatchParamsSharedInputsLocaleAaDj      ExtractBatchParamsSharedInputsLocale = "aa-DJ"
	ExtractBatchParamsSharedInputsLocaleAaEr      ExtractBatchParamsSharedInputsLocale = "aa-ER"
	ExtractBatchParamsSharedInputsLocaleAaEt      ExtractBatchParamsSharedInputsLocale = "aa-ET"
	ExtractBatchParamsSharedInputsLocaleAf        ExtractBatchParamsSharedInputsLocale = "af"
	ExtractBatchParamsSharedInputsLocaleAfNa      ExtractBatchParamsSharedInputsLocale = "af-NA"
	ExtractBatchParamsSharedInputsLocaleAfZa      ExtractBatchParamsSharedInputsLocale = "af-ZA"
	ExtractBatchParamsSharedInputsLocaleAk        ExtractBatchParamsSharedInputsLocale = "ak"
	ExtractBatchParamsSharedInputsLocaleAkGh      ExtractBatchParamsSharedInputsLocale = "ak-GH"
	ExtractBatchParamsSharedInputsLocaleAm        ExtractBatchParamsSharedInputsLocale = "am"
	ExtractBatchParamsSharedInputsLocaleAmEt      ExtractBatchParamsSharedInputsLocale = "am-ET"
	ExtractBatchParamsSharedInputsLocaleAnEs      ExtractBatchParamsSharedInputsLocale = "an-ES"
	ExtractBatchParamsSharedInputsLocaleAr        ExtractBatchParamsSharedInputsLocale = "ar"
	ExtractBatchParamsSharedInputsLocaleArAe      ExtractBatchParamsSharedInputsLocale = "ar-AE"
	ExtractBatchParamsSharedInputsLocaleArBh      ExtractBatchParamsSharedInputsLocale = "ar-BH"
	ExtractBatchParamsSharedInputsLocaleArDz      ExtractBatchParamsSharedInputsLocale = "ar-DZ"
	ExtractBatchParamsSharedInputsLocaleArEg      ExtractBatchParamsSharedInputsLocale = "ar-EG"
	ExtractBatchParamsSharedInputsLocaleArIn      ExtractBatchParamsSharedInputsLocale = "ar-IN"
	ExtractBatchParamsSharedInputsLocaleArIq      ExtractBatchParamsSharedInputsLocale = "ar-IQ"
	ExtractBatchParamsSharedInputsLocaleArJo      ExtractBatchParamsSharedInputsLocale = "ar-JO"
	ExtractBatchParamsSharedInputsLocaleArKw      ExtractBatchParamsSharedInputsLocale = "ar-KW"
	ExtractBatchParamsSharedInputsLocaleArLb      ExtractBatchParamsSharedInputsLocale = "ar-LB"
	ExtractBatchParamsSharedInputsLocaleArLy      ExtractBatchParamsSharedInputsLocale = "ar-LY"
	ExtractBatchParamsSharedInputsLocaleArMa      ExtractBatchParamsSharedInputsLocale = "ar-MA"
	ExtractBatchParamsSharedInputsLocaleArOm      ExtractBatchParamsSharedInputsLocale = "ar-OM"
	ExtractBatchParamsSharedInputsLocaleArQa      ExtractBatchParamsSharedInputsLocale = "ar-QA"
	ExtractBatchParamsSharedInputsLocaleArSa      ExtractBatchParamsSharedInputsLocale = "ar-SA"
	ExtractBatchParamsSharedInputsLocaleArSd      ExtractBatchParamsSharedInputsLocale = "ar-SD"
	ExtractBatchParamsSharedInputsLocaleArSy      ExtractBatchParamsSharedInputsLocale = "ar-SY"
	ExtractBatchParamsSharedInputsLocaleArTn      ExtractBatchParamsSharedInputsLocale = "ar-TN"
	ExtractBatchParamsSharedInputsLocaleArYe      ExtractBatchParamsSharedInputsLocale = "ar-YE"
	ExtractBatchParamsSharedInputsLocaleAs        ExtractBatchParamsSharedInputsLocale = "as"
	ExtractBatchParamsSharedInputsLocaleAsIn      ExtractBatchParamsSharedInputsLocale = "as-IN"
	ExtractBatchParamsSharedInputsLocaleAsa       ExtractBatchParamsSharedInputsLocale = "asa"
	ExtractBatchParamsSharedInputsLocaleAsaTz     ExtractBatchParamsSharedInputsLocale = "asa-TZ"
	ExtractBatchParamsSharedInputsLocaleAstEs     ExtractBatchParamsSharedInputsLocale = "ast-ES"
	ExtractBatchParamsSharedInputsLocaleAz        ExtractBatchParamsSharedInputsLocale = "az"
	ExtractBatchParamsSharedInputsLocaleAzAz      ExtractBatchParamsSharedInputsLocale = "az-AZ"
	ExtractBatchParamsSharedInputsLocaleAzCyrl    ExtractBatchParamsSharedInputsLocale = "az-Cyrl"
	ExtractBatchParamsSharedInputsLocaleAzCyrlAz  ExtractBatchParamsSharedInputsLocale = "az-Cyrl-AZ"
	ExtractBatchParamsSharedInputsLocaleAzLatn    ExtractBatchParamsSharedInputsLocale = "az-Latn"
	ExtractBatchParamsSharedInputsLocaleAzLatnAz  ExtractBatchParamsSharedInputsLocale = "az-Latn-AZ"
	ExtractBatchParamsSharedInputsLocaleBe        ExtractBatchParamsSharedInputsLocale = "be"
	ExtractBatchParamsSharedInputsLocaleBeBy      ExtractBatchParamsSharedInputsLocale = "be-BY"
	ExtractBatchParamsSharedInputsLocaleBem       ExtractBatchParamsSharedInputsLocale = "bem"
	ExtractBatchParamsSharedInputsLocaleBemZm     ExtractBatchParamsSharedInputsLocale = "bem-ZM"
	ExtractBatchParamsSharedInputsLocaleBerDz     ExtractBatchParamsSharedInputsLocale = "ber-DZ"
	ExtractBatchParamsSharedInputsLocaleBerMa     ExtractBatchParamsSharedInputsLocale = "ber-MA"
	ExtractBatchParamsSharedInputsLocaleBez       ExtractBatchParamsSharedInputsLocale = "bez"
	ExtractBatchParamsSharedInputsLocaleBezTz     ExtractBatchParamsSharedInputsLocale = "bez-TZ"
	ExtractBatchParamsSharedInputsLocaleBg        ExtractBatchParamsSharedInputsLocale = "bg"
	ExtractBatchParamsSharedInputsLocaleBgBg      ExtractBatchParamsSharedInputsLocale = "bg-BG"
	ExtractBatchParamsSharedInputsLocaleBhoIn     ExtractBatchParamsSharedInputsLocale = "bho-IN"
	ExtractBatchParamsSharedInputsLocaleBm        ExtractBatchParamsSharedInputsLocale = "bm"
	ExtractBatchParamsSharedInputsLocaleBmMl      ExtractBatchParamsSharedInputsLocale = "bm-ML"
	ExtractBatchParamsSharedInputsLocaleBn        ExtractBatchParamsSharedInputsLocale = "bn"
	ExtractBatchParamsSharedInputsLocaleBnBd      ExtractBatchParamsSharedInputsLocale = "bn-BD"
	ExtractBatchParamsSharedInputsLocaleBnIn      ExtractBatchParamsSharedInputsLocale = "bn-IN"
	ExtractBatchParamsSharedInputsLocaleBo        ExtractBatchParamsSharedInputsLocale = "bo"
	ExtractBatchParamsSharedInputsLocaleBoCn      ExtractBatchParamsSharedInputsLocale = "bo-CN"
	ExtractBatchParamsSharedInputsLocaleBoIn      ExtractBatchParamsSharedInputsLocale = "bo-IN"
	ExtractBatchParamsSharedInputsLocaleBrFr      ExtractBatchParamsSharedInputsLocale = "br-FR"
	ExtractBatchParamsSharedInputsLocaleBrxIn     ExtractBatchParamsSharedInputsLocale = "brx-IN"
	ExtractBatchParamsSharedInputsLocaleBs        ExtractBatchParamsSharedInputsLocale = "bs"
	ExtractBatchParamsSharedInputsLocaleBsBa      ExtractBatchParamsSharedInputsLocale = "bs-BA"
	ExtractBatchParamsSharedInputsLocaleBynEr     ExtractBatchParamsSharedInputsLocale = "byn-ER"
	ExtractBatchParamsSharedInputsLocaleCa        ExtractBatchParamsSharedInputsLocale = "ca"
	ExtractBatchParamsSharedInputsLocaleCaAd      ExtractBatchParamsSharedInputsLocale = "ca-AD"
	ExtractBatchParamsSharedInputsLocaleCaEs      ExtractBatchParamsSharedInputsLocale = "ca-ES"
	ExtractBatchParamsSharedInputsLocaleCaFr      ExtractBatchParamsSharedInputsLocale = "ca-FR"
	ExtractBatchParamsSharedInputsLocaleCaIt      ExtractBatchParamsSharedInputsLocale = "ca-IT"
	ExtractBatchParamsSharedInputsLocaleCgg       ExtractBatchParamsSharedInputsLocale = "cgg"
	ExtractBatchParamsSharedInputsLocaleCggUg     ExtractBatchParamsSharedInputsLocale = "cgg-UG"
	ExtractBatchParamsSharedInputsLocaleChr       ExtractBatchParamsSharedInputsLocale = "chr"
	ExtractBatchParamsSharedInputsLocaleChrUs     ExtractBatchParamsSharedInputsLocale = "chr-US"
	ExtractBatchParamsSharedInputsLocaleCrhUa     ExtractBatchParamsSharedInputsLocale = "crh-UA"
	ExtractBatchParamsSharedInputsLocaleCs        ExtractBatchParamsSharedInputsLocale = "cs"
	ExtractBatchParamsSharedInputsLocaleCsCz      ExtractBatchParamsSharedInputsLocale = "cs-CZ"
	ExtractBatchParamsSharedInputsLocaleCsbPl     ExtractBatchParamsSharedInputsLocale = "csb-PL"
	ExtractBatchParamsSharedInputsLocaleCvRu      ExtractBatchParamsSharedInputsLocale = "cv-RU"
	ExtractBatchParamsSharedInputsLocaleCy        ExtractBatchParamsSharedInputsLocale = "cy"
	ExtractBatchParamsSharedInputsLocaleCyGB      ExtractBatchParamsSharedInputsLocale = "cy-GB"
	ExtractBatchParamsSharedInputsLocaleDa        ExtractBatchParamsSharedInputsLocale = "da"
	ExtractBatchParamsSharedInputsLocaleDaDk      ExtractBatchParamsSharedInputsLocale = "da-DK"
	ExtractBatchParamsSharedInputsLocaleDav       ExtractBatchParamsSharedInputsLocale = "dav"
	ExtractBatchParamsSharedInputsLocaleDavKe     ExtractBatchParamsSharedInputsLocale = "dav-KE"
	ExtractBatchParamsSharedInputsLocaleDe        ExtractBatchParamsSharedInputsLocale = "de"
	ExtractBatchParamsSharedInputsLocaleDeAt      ExtractBatchParamsSharedInputsLocale = "de-AT"
	ExtractBatchParamsSharedInputsLocaleDeBe      ExtractBatchParamsSharedInputsLocale = "de-BE"
	ExtractBatchParamsSharedInputsLocaleDeCh      ExtractBatchParamsSharedInputsLocale = "de-CH"
	ExtractBatchParamsSharedInputsLocaleDeDe      ExtractBatchParamsSharedInputsLocale = "de-DE"
	ExtractBatchParamsSharedInputsLocaleDeLi      ExtractBatchParamsSharedInputsLocale = "de-LI"
	ExtractBatchParamsSharedInputsLocaleDeLu      ExtractBatchParamsSharedInputsLocale = "de-LU"
	ExtractBatchParamsSharedInputsLocaleDvMv      ExtractBatchParamsSharedInputsLocale = "dv-MV"
	ExtractBatchParamsSharedInputsLocaleDzBt      ExtractBatchParamsSharedInputsLocale = "dz-BT"
	ExtractBatchParamsSharedInputsLocaleEbu       ExtractBatchParamsSharedInputsLocale = "ebu"
	ExtractBatchParamsSharedInputsLocaleEbuKe     ExtractBatchParamsSharedInputsLocale = "ebu-KE"
	ExtractBatchParamsSharedInputsLocaleEe        ExtractBatchParamsSharedInputsLocale = "ee"
	ExtractBatchParamsSharedInputsLocaleEeGh      ExtractBatchParamsSharedInputsLocale = "ee-GH"
	ExtractBatchParamsSharedInputsLocaleEeTg      ExtractBatchParamsSharedInputsLocale = "ee-TG"
	ExtractBatchParamsSharedInputsLocaleEl        ExtractBatchParamsSharedInputsLocale = "el"
	ExtractBatchParamsSharedInputsLocaleElCy      ExtractBatchParamsSharedInputsLocale = "el-CY"
	ExtractBatchParamsSharedInputsLocaleElGr      ExtractBatchParamsSharedInputsLocale = "el-GR"
	ExtractBatchParamsSharedInputsLocaleEn        ExtractBatchParamsSharedInputsLocale = "en"
	ExtractBatchParamsSharedInputsLocaleEnAg      ExtractBatchParamsSharedInputsLocale = "en-AG"
	ExtractBatchParamsSharedInputsLocaleEnAs      ExtractBatchParamsSharedInputsLocale = "en-AS"
	ExtractBatchParamsSharedInputsLocaleEnAu      ExtractBatchParamsSharedInputsLocale = "en-AU"
	ExtractBatchParamsSharedInputsLocaleEnBe      ExtractBatchParamsSharedInputsLocale = "en-BE"
	ExtractBatchParamsSharedInputsLocaleEnBw      ExtractBatchParamsSharedInputsLocale = "en-BW"
	ExtractBatchParamsSharedInputsLocaleEnBz      ExtractBatchParamsSharedInputsLocale = "en-BZ"
	ExtractBatchParamsSharedInputsLocaleEnCa      ExtractBatchParamsSharedInputsLocale = "en-CA"
	ExtractBatchParamsSharedInputsLocaleEnDk      ExtractBatchParamsSharedInputsLocale = "en-DK"
	ExtractBatchParamsSharedInputsLocaleEnGB      ExtractBatchParamsSharedInputsLocale = "en-GB"
	ExtractBatchParamsSharedInputsLocaleEnGu      ExtractBatchParamsSharedInputsLocale = "en-GU"
	ExtractBatchParamsSharedInputsLocaleEnHk      ExtractBatchParamsSharedInputsLocale = "en-HK"
	ExtractBatchParamsSharedInputsLocaleEnIe      ExtractBatchParamsSharedInputsLocale = "en-IE"
	ExtractBatchParamsSharedInputsLocaleEnIn      ExtractBatchParamsSharedInputsLocale = "en-IN"
	ExtractBatchParamsSharedInputsLocaleEnJm      ExtractBatchParamsSharedInputsLocale = "en-JM"
	ExtractBatchParamsSharedInputsLocaleEnMh      ExtractBatchParamsSharedInputsLocale = "en-MH"
	ExtractBatchParamsSharedInputsLocaleEnMp      ExtractBatchParamsSharedInputsLocale = "en-MP"
	ExtractBatchParamsSharedInputsLocaleEnMt      ExtractBatchParamsSharedInputsLocale = "en-MT"
	ExtractBatchParamsSharedInputsLocaleEnMu      ExtractBatchParamsSharedInputsLocale = "en-MU"
	ExtractBatchParamsSharedInputsLocaleEnNa      ExtractBatchParamsSharedInputsLocale = "en-NA"
	ExtractBatchParamsSharedInputsLocaleEnNg      ExtractBatchParamsSharedInputsLocale = "en-NG"
	ExtractBatchParamsSharedInputsLocaleEnNz      ExtractBatchParamsSharedInputsLocale = "en-NZ"
	ExtractBatchParamsSharedInputsLocaleEnPh      ExtractBatchParamsSharedInputsLocale = "en-PH"
	ExtractBatchParamsSharedInputsLocaleEnPk      ExtractBatchParamsSharedInputsLocale = "en-PK"
	ExtractBatchParamsSharedInputsLocaleEnSg      ExtractBatchParamsSharedInputsLocale = "en-SG"
	ExtractBatchParamsSharedInputsLocaleEnTt      ExtractBatchParamsSharedInputsLocale = "en-TT"
	ExtractBatchParamsSharedInputsLocaleEnUm      ExtractBatchParamsSharedInputsLocale = "en-UM"
	ExtractBatchParamsSharedInputsLocaleEnUs      ExtractBatchParamsSharedInputsLocale = "en-US"
	ExtractBatchParamsSharedInputsLocaleEnVi      ExtractBatchParamsSharedInputsLocale = "en-VI"
	ExtractBatchParamsSharedInputsLocaleEnZa      ExtractBatchParamsSharedInputsLocale = "en-ZA"
	ExtractBatchParamsSharedInputsLocaleEnZm      ExtractBatchParamsSharedInputsLocale = "en-ZM"
	ExtractBatchParamsSharedInputsLocaleEnZw      ExtractBatchParamsSharedInputsLocale = "en-ZW"
	ExtractBatchParamsSharedInputsLocaleEo        ExtractBatchParamsSharedInputsLocale = "eo"
	ExtractBatchParamsSharedInputsLocaleEs        ExtractBatchParamsSharedInputsLocale = "es"
	ExtractBatchParamsSharedInputsLocaleEs419     ExtractBatchParamsSharedInputsLocale = "es-419"
	ExtractBatchParamsSharedInputsLocaleEsAr      ExtractBatchParamsSharedInputsLocale = "es-AR"
	ExtractBatchParamsSharedInputsLocaleEsBo      ExtractBatchParamsSharedInputsLocale = "es-BO"
	ExtractBatchParamsSharedInputsLocaleEsCl      ExtractBatchParamsSharedInputsLocale = "es-CL"
	ExtractBatchParamsSharedInputsLocaleEsCo      ExtractBatchParamsSharedInputsLocale = "es-CO"
	ExtractBatchParamsSharedInputsLocaleEsCr      ExtractBatchParamsSharedInputsLocale = "es-CR"
	ExtractBatchParamsSharedInputsLocaleEsCu      ExtractBatchParamsSharedInputsLocale = "es-CU"
	ExtractBatchParamsSharedInputsLocaleEsDo      ExtractBatchParamsSharedInputsLocale = "es-DO"
	ExtractBatchParamsSharedInputsLocaleEsEc      ExtractBatchParamsSharedInputsLocale = "es-EC"
	ExtractBatchParamsSharedInputsLocaleEsEs      ExtractBatchParamsSharedInputsLocale = "es-ES"
	ExtractBatchParamsSharedInputsLocaleEsGq      ExtractBatchParamsSharedInputsLocale = "es-GQ"
	ExtractBatchParamsSharedInputsLocaleEsGt      ExtractBatchParamsSharedInputsLocale = "es-GT"
	ExtractBatchParamsSharedInputsLocaleEsHn      ExtractBatchParamsSharedInputsLocale = "es-HN"
	ExtractBatchParamsSharedInputsLocaleEsMx      ExtractBatchParamsSharedInputsLocale = "es-MX"
	ExtractBatchParamsSharedInputsLocaleEsNi      ExtractBatchParamsSharedInputsLocale = "es-NI"
	ExtractBatchParamsSharedInputsLocaleEsPa      ExtractBatchParamsSharedInputsLocale = "es-PA"
	ExtractBatchParamsSharedInputsLocaleEsPe      ExtractBatchParamsSharedInputsLocale = "es-PE"
	ExtractBatchParamsSharedInputsLocaleEsPr      ExtractBatchParamsSharedInputsLocale = "es-PR"
	ExtractBatchParamsSharedInputsLocaleEsPy      ExtractBatchParamsSharedInputsLocale = "es-PY"
	ExtractBatchParamsSharedInputsLocaleEsSv      ExtractBatchParamsSharedInputsLocale = "es-SV"
	ExtractBatchParamsSharedInputsLocaleEsUs      ExtractBatchParamsSharedInputsLocale = "es-US"
	ExtractBatchParamsSharedInputsLocaleEsUy      ExtractBatchParamsSharedInputsLocale = "es-UY"
	ExtractBatchParamsSharedInputsLocaleEsVe      ExtractBatchParamsSharedInputsLocale = "es-VE"
	ExtractBatchParamsSharedInputsLocaleEt        ExtractBatchParamsSharedInputsLocale = "et"
	ExtractBatchParamsSharedInputsLocaleEtEe      ExtractBatchParamsSharedInputsLocale = "et-EE"
	ExtractBatchParamsSharedInputsLocaleEu        ExtractBatchParamsSharedInputsLocale = "eu"
	ExtractBatchParamsSharedInputsLocaleEuEs      ExtractBatchParamsSharedInputsLocale = "eu-ES"
	ExtractBatchParamsSharedInputsLocaleFa        ExtractBatchParamsSharedInputsLocale = "fa"
	ExtractBatchParamsSharedInputsLocaleFaAf      ExtractBatchParamsSharedInputsLocale = "fa-AF"
	ExtractBatchParamsSharedInputsLocaleFaIr      ExtractBatchParamsSharedInputsLocale = "fa-IR"
	ExtractBatchParamsSharedInputsLocaleFf        ExtractBatchParamsSharedInputsLocale = "ff"
	ExtractBatchParamsSharedInputsLocaleFfSn      ExtractBatchParamsSharedInputsLocale = "ff-SN"
	ExtractBatchParamsSharedInputsLocaleFi        ExtractBatchParamsSharedInputsLocale = "fi"
	ExtractBatchParamsSharedInputsLocaleFiFi      ExtractBatchParamsSharedInputsLocale = "fi-FI"
	ExtractBatchParamsSharedInputsLocaleFil       ExtractBatchParamsSharedInputsLocale = "fil"
	ExtractBatchParamsSharedInputsLocaleFilPh     ExtractBatchParamsSharedInputsLocale = "fil-PH"
	ExtractBatchParamsSharedInputsLocaleFo        ExtractBatchParamsSharedInputsLocale = "fo"
	ExtractBatchParamsSharedInputsLocaleFoFo      ExtractBatchParamsSharedInputsLocale = "fo-FO"
	ExtractBatchParamsSharedInputsLocaleFr        ExtractBatchParamsSharedInputsLocale = "fr"
	ExtractBatchParamsSharedInputsLocaleFrBe      ExtractBatchParamsSharedInputsLocale = "fr-BE"
	ExtractBatchParamsSharedInputsLocaleFrBf      ExtractBatchParamsSharedInputsLocale = "fr-BF"
	ExtractBatchParamsSharedInputsLocaleFrBi      ExtractBatchParamsSharedInputsLocale = "fr-BI"
	ExtractBatchParamsSharedInputsLocaleFrBj      ExtractBatchParamsSharedInputsLocale = "fr-BJ"
	ExtractBatchParamsSharedInputsLocaleFrBl      ExtractBatchParamsSharedInputsLocale = "fr-BL"
	ExtractBatchParamsSharedInputsLocaleFrCa      ExtractBatchParamsSharedInputsLocale = "fr-CA"
	ExtractBatchParamsSharedInputsLocaleFrCd      ExtractBatchParamsSharedInputsLocale = "fr-CD"
	ExtractBatchParamsSharedInputsLocaleFrCf      ExtractBatchParamsSharedInputsLocale = "fr-CF"
	ExtractBatchParamsSharedInputsLocaleFrCg      ExtractBatchParamsSharedInputsLocale = "fr-CG"
	ExtractBatchParamsSharedInputsLocaleFrCh      ExtractBatchParamsSharedInputsLocale = "fr-CH"
	ExtractBatchParamsSharedInputsLocaleFrCi      ExtractBatchParamsSharedInputsLocale = "fr-CI"
	ExtractBatchParamsSharedInputsLocaleFrCm      ExtractBatchParamsSharedInputsLocale = "fr-CM"
	ExtractBatchParamsSharedInputsLocaleFrDj      ExtractBatchParamsSharedInputsLocale = "fr-DJ"
	ExtractBatchParamsSharedInputsLocaleFrFr      ExtractBatchParamsSharedInputsLocale = "fr-FR"
	ExtractBatchParamsSharedInputsLocaleFrGa      ExtractBatchParamsSharedInputsLocale = "fr-GA"
	ExtractBatchParamsSharedInputsLocaleFrGn      ExtractBatchParamsSharedInputsLocale = "fr-GN"
	ExtractBatchParamsSharedInputsLocaleFrGp      ExtractBatchParamsSharedInputsLocale = "fr-GP"
	ExtractBatchParamsSharedInputsLocaleFrGq      ExtractBatchParamsSharedInputsLocale = "fr-GQ"
	ExtractBatchParamsSharedInputsLocaleFrKm      ExtractBatchParamsSharedInputsLocale = "fr-KM"
	ExtractBatchParamsSharedInputsLocaleFrLu      ExtractBatchParamsSharedInputsLocale = "fr-LU"
	ExtractBatchParamsSharedInputsLocaleFrMc      ExtractBatchParamsSharedInputsLocale = "fr-MC"
	ExtractBatchParamsSharedInputsLocaleFrMf      ExtractBatchParamsSharedInputsLocale = "fr-MF"
	ExtractBatchParamsSharedInputsLocaleFrMg      ExtractBatchParamsSharedInputsLocale = "fr-MG"
	ExtractBatchParamsSharedInputsLocaleFrMl      ExtractBatchParamsSharedInputsLocale = "fr-ML"
	ExtractBatchParamsSharedInputsLocaleFrMq      ExtractBatchParamsSharedInputsLocale = "fr-MQ"
	ExtractBatchParamsSharedInputsLocaleFrNe      ExtractBatchParamsSharedInputsLocale = "fr-NE"
	ExtractBatchParamsSharedInputsLocaleFrRe      ExtractBatchParamsSharedInputsLocale = "fr-RE"
	ExtractBatchParamsSharedInputsLocaleFrRw      ExtractBatchParamsSharedInputsLocale = "fr-RW"
	ExtractBatchParamsSharedInputsLocaleFrSn      ExtractBatchParamsSharedInputsLocale = "fr-SN"
	ExtractBatchParamsSharedInputsLocaleFrTd      ExtractBatchParamsSharedInputsLocale = "fr-TD"
	ExtractBatchParamsSharedInputsLocaleFrTg      ExtractBatchParamsSharedInputsLocale = "fr-TG"
	ExtractBatchParamsSharedInputsLocaleFurIt     ExtractBatchParamsSharedInputsLocale = "fur-IT"
	ExtractBatchParamsSharedInputsLocaleFyDe      ExtractBatchParamsSharedInputsLocale = "fy-DE"
	ExtractBatchParamsSharedInputsLocaleFyNl      ExtractBatchParamsSharedInputsLocale = "fy-NL"
	ExtractBatchParamsSharedInputsLocaleGa        ExtractBatchParamsSharedInputsLocale = "ga"
	ExtractBatchParamsSharedInputsLocaleGaIe      ExtractBatchParamsSharedInputsLocale = "ga-IE"
	ExtractBatchParamsSharedInputsLocaleGdGB      ExtractBatchParamsSharedInputsLocale = "gd-GB"
	ExtractBatchParamsSharedInputsLocaleGezEr     ExtractBatchParamsSharedInputsLocale = "gez-ER"
	ExtractBatchParamsSharedInputsLocaleGezEt     ExtractBatchParamsSharedInputsLocale = "gez-ET"
	ExtractBatchParamsSharedInputsLocaleGl        ExtractBatchParamsSharedInputsLocale = "gl"
	ExtractBatchParamsSharedInputsLocaleGlEs      ExtractBatchParamsSharedInputsLocale = "gl-ES"
	ExtractBatchParamsSharedInputsLocaleGsw       ExtractBatchParamsSharedInputsLocale = "gsw"
	ExtractBatchParamsSharedInputsLocaleGswCh     ExtractBatchParamsSharedInputsLocale = "gsw-CH"
	ExtractBatchParamsSharedInputsLocaleGu        ExtractBatchParamsSharedInputsLocale = "gu"
	ExtractBatchParamsSharedInputsLocaleGuIn      ExtractBatchParamsSharedInputsLocale = "gu-IN"
	ExtractBatchParamsSharedInputsLocaleGuz       ExtractBatchParamsSharedInputsLocale = "guz"
	ExtractBatchParamsSharedInputsLocaleGuzKe     ExtractBatchParamsSharedInputsLocale = "guz-KE"
	ExtractBatchParamsSharedInputsLocaleGv        ExtractBatchParamsSharedInputsLocale = "gv"
	ExtractBatchParamsSharedInputsLocaleGvGB      ExtractBatchParamsSharedInputsLocale = "gv-GB"
	ExtractBatchParamsSharedInputsLocaleHa        ExtractBatchParamsSharedInputsLocale = "ha"
	ExtractBatchParamsSharedInputsLocaleHaLatn    ExtractBatchParamsSharedInputsLocale = "ha-Latn"
	ExtractBatchParamsSharedInputsLocaleHaLatnGh  ExtractBatchParamsSharedInputsLocale = "ha-Latn-GH"
	ExtractBatchParamsSharedInputsLocaleHaLatnNe  ExtractBatchParamsSharedInputsLocale = "ha-Latn-NE"
	ExtractBatchParamsSharedInputsLocaleHaLatnNg  ExtractBatchParamsSharedInputsLocale = "ha-Latn-NG"
	ExtractBatchParamsSharedInputsLocaleHaNg      ExtractBatchParamsSharedInputsLocale = "ha-NG"
	ExtractBatchParamsSharedInputsLocaleHaw       ExtractBatchParamsSharedInputsLocale = "haw"
	ExtractBatchParamsSharedInputsLocaleHawUs     ExtractBatchParamsSharedInputsLocale = "haw-US"
	ExtractBatchParamsSharedInputsLocaleHe        ExtractBatchParamsSharedInputsLocale = "he"
	ExtractBatchParamsSharedInputsLocaleHeIl      ExtractBatchParamsSharedInputsLocale = "he-IL"
	ExtractBatchParamsSharedInputsLocaleHi        ExtractBatchParamsSharedInputsLocale = "hi"
	ExtractBatchParamsSharedInputsLocaleHiIn      ExtractBatchParamsSharedInputsLocale = "hi-IN"
	ExtractBatchParamsSharedInputsLocaleHneIn     ExtractBatchParamsSharedInputsLocale = "hne-IN"
	ExtractBatchParamsSharedInputsLocaleHr        ExtractBatchParamsSharedInputsLocale = "hr"
	ExtractBatchParamsSharedInputsLocaleHrHr      ExtractBatchParamsSharedInputsLocale = "hr-HR"
	ExtractBatchParamsSharedInputsLocaleHsbDe     ExtractBatchParamsSharedInputsLocale = "hsb-DE"
	ExtractBatchParamsSharedInputsLocaleHtHt      ExtractBatchParamsSharedInputsLocale = "ht-HT"
	ExtractBatchParamsSharedInputsLocaleHu        ExtractBatchParamsSharedInputsLocale = "hu"
	ExtractBatchParamsSharedInputsLocaleHuHu      ExtractBatchParamsSharedInputsLocale = "hu-HU"
	ExtractBatchParamsSharedInputsLocaleHy        ExtractBatchParamsSharedInputsLocale = "hy"
	ExtractBatchParamsSharedInputsLocaleHyAm      ExtractBatchParamsSharedInputsLocale = "hy-AM"
	ExtractBatchParamsSharedInputsLocaleID        ExtractBatchParamsSharedInputsLocale = "id"
	ExtractBatchParamsSharedInputsLocaleIDID      ExtractBatchParamsSharedInputsLocale = "id-ID"
	ExtractBatchParamsSharedInputsLocaleIg        ExtractBatchParamsSharedInputsLocale = "ig"
	ExtractBatchParamsSharedInputsLocaleIgNg      ExtractBatchParamsSharedInputsLocale = "ig-NG"
	ExtractBatchParamsSharedInputsLocaleIi        ExtractBatchParamsSharedInputsLocale = "ii"
	ExtractBatchParamsSharedInputsLocaleIiCn      ExtractBatchParamsSharedInputsLocale = "ii-CN"
	ExtractBatchParamsSharedInputsLocaleIkCa      ExtractBatchParamsSharedInputsLocale = "ik-CA"
	ExtractBatchParamsSharedInputsLocaleIs        ExtractBatchParamsSharedInputsLocale = "is"
	ExtractBatchParamsSharedInputsLocaleIsIs      ExtractBatchParamsSharedInputsLocale = "is-IS"
	ExtractBatchParamsSharedInputsLocaleIt        ExtractBatchParamsSharedInputsLocale = "it"
	ExtractBatchParamsSharedInputsLocaleItCh      ExtractBatchParamsSharedInputsLocale = "it-CH"
	ExtractBatchParamsSharedInputsLocaleItIt      ExtractBatchParamsSharedInputsLocale = "it-IT"
	ExtractBatchParamsSharedInputsLocaleIuCa      ExtractBatchParamsSharedInputsLocale = "iu-CA"
	ExtractBatchParamsSharedInputsLocaleIwIl      ExtractBatchParamsSharedInputsLocale = "iw-IL"
	ExtractBatchParamsSharedInputsLocaleJa        ExtractBatchParamsSharedInputsLocale = "ja"
	ExtractBatchParamsSharedInputsLocaleJaJp      ExtractBatchParamsSharedInputsLocale = "ja-JP"
	ExtractBatchParamsSharedInputsLocaleJmc       ExtractBatchParamsSharedInputsLocale = "jmc"
	ExtractBatchParamsSharedInputsLocaleJmcTz     ExtractBatchParamsSharedInputsLocale = "jmc-TZ"
	ExtractBatchParamsSharedInputsLocaleKa        ExtractBatchParamsSharedInputsLocale = "ka"
	ExtractBatchParamsSharedInputsLocaleKaGe      ExtractBatchParamsSharedInputsLocale = "ka-GE"
	ExtractBatchParamsSharedInputsLocaleKab       ExtractBatchParamsSharedInputsLocale = "kab"
	ExtractBatchParamsSharedInputsLocaleKabDz     ExtractBatchParamsSharedInputsLocale = "kab-DZ"
	ExtractBatchParamsSharedInputsLocaleKam       ExtractBatchParamsSharedInputsLocale = "kam"
	ExtractBatchParamsSharedInputsLocaleKamKe     ExtractBatchParamsSharedInputsLocale = "kam-KE"
	ExtractBatchParamsSharedInputsLocaleKde       ExtractBatchParamsSharedInputsLocale = "kde"
	ExtractBatchParamsSharedInputsLocaleKdeTz     ExtractBatchParamsSharedInputsLocale = "kde-TZ"
	ExtractBatchParamsSharedInputsLocaleKea       ExtractBatchParamsSharedInputsLocale = "kea"
	ExtractBatchParamsSharedInputsLocaleKeaCv     ExtractBatchParamsSharedInputsLocale = "kea-CV"
	ExtractBatchParamsSharedInputsLocaleKhq       ExtractBatchParamsSharedInputsLocale = "khq"
	ExtractBatchParamsSharedInputsLocaleKhqMl     ExtractBatchParamsSharedInputsLocale = "khq-ML"
	ExtractBatchParamsSharedInputsLocaleKi        ExtractBatchParamsSharedInputsLocale = "ki"
	ExtractBatchParamsSharedInputsLocaleKiKe      ExtractBatchParamsSharedInputsLocale = "ki-KE"
	ExtractBatchParamsSharedInputsLocaleKk        ExtractBatchParamsSharedInputsLocale = "kk"
	ExtractBatchParamsSharedInputsLocaleKkCyrl    ExtractBatchParamsSharedInputsLocale = "kk-Cyrl"
	ExtractBatchParamsSharedInputsLocaleKkCyrlKz  ExtractBatchParamsSharedInputsLocale = "kk-Cyrl-KZ"
	ExtractBatchParamsSharedInputsLocaleKkKz      ExtractBatchParamsSharedInputsLocale = "kk-KZ"
	ExtractBatchParamsSharedInputsLocaleKl        ExtractBatchParamsSharedInputsLocale = "kl"
	ExtractBatchParamsSharedInputsLocaleKlGl      ExtractBatchParamsSharedInputsLocale = "kl-GL"
	ExtractBatchParamsSharedInputsLocaleKln       ExtractBatchParamsSharedInputsLocale = "kln"
	ExtractBatchParamsSharedInputsLocaleKlnKe     ExtractBatchParamsSharedInputsLocale = "kln-KE"
	ExtractBatchParamsSharedInputsLocaleKm        ExtractBatchParamsSharedInputsLocale = "km"
	ExtractBatchParamsSharedInputsLocaleKmKh      ExtractBatchParamsSharedInputsLocale = "km-KH"
	ExtractBatchParamsSharedInputsLocaleKn        ExtractBatchParamsSharedInputsLocale = "kn"
	ExtractBatchParamsSharedInputsLocaleKnIn      ExtractBatchParamsSharedInputsLocale = "kn-IN"
	ExtractBatchParamsSharedInputsLocaleKo        ExtractBatchParamsSharedInputsLocale = "ko"
	ExtractBatchParamsSharedInputsLocaleKoKr      ExtractBatchParamsSharedInputsLocale = "ko-KR"
	ExtractBatchParamsSharedInputsLocaleKok       ExtractBatchParamsSharedInputsLocale = "kok"
	ExtractBatchParamsSharedInputsLocaleKokIn     ExtractBatchParamsSharedInputsLocale = "kok-IN"
	ExtractBatchParamsSharedInputsLocaleKsIn      ExtractBatchParamsSharedInputsLocale = "ks-IN"
	ExtractBatchParamsSharedInputsLocaleKuTr      ExtractBatchParamsSharedInputsLocale = "ku-TR"
	ExtractBatchParamsSharedInputsLocaleKw        ExtractBatchParamsSharedInputsLocale = "kw"
	ExtractBatchParamsSharedInputsLocaleKwGB      ExtractBatchParamsSharedInputsLocale = "kw-GB"
	ExtractBatchParamsSharedInputsLocaleKyKg      ExtractBatchParamsSharedInputsLocale = "ky-KG"
	ExtractBatchParamsSharedInputsLocaleLag       ExtractBatchParamsSharedInputsLocale = "lag"
	ExtractBatchParamsSharedInputsLocaleLagTz     ExtractBatchParamsSharedInputsLocale = "lag-TZ"
	ExtractBatchParamsSharedInputsLocaleLbLu      ExtractBatchParamsSharedInputsLocale = "lb-LU"
	ExtractBatchParamsSharedInputsLocaleLg        ExtractBatchParamsSharedInputsLocale = "lg"
	ExtractBatchParamsSharedInputsLocaleLgUg      ExtractBatchParamsSharedInputsLocale = "lg-UG"
	ExtractBatchParamsSharedInputsLocaleLiBe      ExtractBatchParamsSharedInputsLocale = "li-BE"
	ExtractBatchParamsSharedInputsLocaleLiNl      ExtractBatchParamsSharedInputsLocale = "li-NL"
	ExtractBatchParamsSharedInputsLocaleLijIt     ExtractBatchParamsSharedInputsLocale = "lij-IT"
	ExtractBatchParamsSharedInputsLocaleLoLa      ExtractBatchParamsSharedInputsLocale = "lo-LA"
	ExtractBatchParamsSharedInputsLocaleLt        ExtractBatchParamsSharedInputsLocale = "lt"
	ExtractBatchParamsSharedInputsLocaleLtLt      ExtractBatchParamsSharedInputsLocale = "lt-LT"
	ExtractBatchParamsSharedInputsLocaleLuo       ExtractBatchParamsSharedInputsLocale = "luo"
	ExtractBatchParamsSharedInputsLocaleLuoKe     ExtractBatchParamsSharedInputsLocale = "luo-KE"
	ExtractBatchParamsSharedInputsLocaleLuy       ExtractBatchParamsSharedInputsLocale = "luy"
	ExtractBatchParamsSharedInputsLocaleLuyKe     ExtractBatchParamsSharedInputsLocale = "luy-KE"
	ExtractBatchParamsSharedInputsLocaleLv        ExtractBatchParamsSharedInputsLocale = "lv"
	ExtractBatchParamsSharedInputsLocaleLvLv      ExtractBatchParamsSharedInputsLocale = "lv-LV"
	ExtractBatchParamsSharedInputsLocaleMagIn     ExtractBatchParamsSharedInputsLocale = "mag-IN"
	ExtractBatchParamsSharedInputsLocaleMaiIn     ExtractBatchParamsSharedInputsLocale = "mai-IN"
	ExtractBatchParamsSharedInputsLocaleMas       ExtractBatchParamsSharedInputsLocale = "mas"
	ExtractBatchParamsSharedInputsLocaleMasKe     ExtractBatchParamsSharedInputsLocale = "mas-KE"
	ExtractBatchParamsSharedInputsLocaleMasTz     ExtractBatchParamsSharedInputsLocale = "mas-TZ"
	ExtractBatchParamsSharedInputsLocaleMer       ExtractBatchParamsSharedInputsLocale = "mer"
	ExtractBatchParamsSharedInputsLocaleMerKe     ExtractBatchParamsSharedInputsLocale = "mer-KE"
	ExtractBatchParamsSharedInputsLocaleMfe       ExtractBatchParamsSharedInputsLocale = "mfe"
	ExtractBatchParamsSharedInputsLocaleMfeMu     ExtractBatchParamsSharedInputsLocale = "mfe-MU"
	ExtractBatchParamsSharedInputsLocaleMg        ExtractBatchParamsSharedInputsLocale = "mg"
	ExtractBatchParamsSharedInputsLocaleMgMg      ExtractBatchParamsSharedInputsLocale = "mg-MG"
	ExtractBatchParamsSharedInputsLocaleMhrRu     ExtractBatchParamsSharedInputsLocale = "mhr-RU"
	ExtractBatchParamsSharedInputsLocaleMiNz      ExtractBatchParamsSharedInputsLocale = "mi-NZ"
	ExtractBatchParamsSharedInputsLocaleMk        ExtractBatchParamsSharedInputsLocale = "mk"
	ExtractBatchParamsSharedInputsLocaleMkMk      ExtractBatchParamsSharedInputsLocale = "mk-MK"
	ExtractBatchParamsSharedInputsLocaleMl        ExtractBatchParamsSharedInputsLocale = "ml"
	ExtractBatchParamsSharedInputsLocaleMlIn      ExtractBatchParamsSharedInputsLocale = "ml-IN"
	ExtractBatchParamsSharedInputsLocaleMnMn      ExtractBatchParamsSharedInputsLocale = "mn-MN"
	ExtractBatchParamsSharedInputsLocaleMr        ExtractBatchParamsSharedInputsLocale = "mr"
	ExtractBatchParamsSharedInputsLocaleMrIn      ExtractBatchParamsSharedInputsLocale = "mr-IN"
	ExtractBatchParamsSharedInputsLocaleMs        ExtractBatchParamsSharedInputsLocale = "ms"
	ExtractBatchParamsSharedInputsLocaleMsBn      ExtractBatchParamsSharedInputsLocale = "ms-BN"
	ExtractBatchParamsSharedInputsLocaleMsMy      ExtractBatchParamsSharedInputsLocale = "ms-MY"
	ExtractBatchParamsSharedInputsLocaleMt        ExtractBatchParamsSharedInputsLocale = "mt"
	ExtractBatchParamsSharedInputsLocaleMtMt      ExtractBatchParamsSharedInputsLocale = "mt-MT"
	ExtractBatchParamsSharedInputsLocaleMy        ExtractBatchParamsSharedInputsLocale = "my"
	ExtractBatchParamsSharedInputsLocaleMyMm      ExtractBatchParamsSharedInputsLocale = "my-MM"
	ExtractBatchParamsSharedInputsLocaleNanTw     ExtractBatchParamsSharedInputsLocale = "nan-TW"
	ExtractBatchParamsSharedInputsLocaleNaq       ExtractBatchParamsSharedInputsLocale = "naq"
	ExtractBatchParamsSharedInputsLocaleNaqNa     ExtractBatchParamsSharedInputsLocale = "naq-NA"
	ExtractBatchParamsSharedInputsLocaleNb        ExtractBatchParamsSharedInputsLocale = "nb"
	ExtractBatchParamsSharedInputsLocaleNbNo      ExtractBatchParamsSharedInputsLocale = "nb-NO"
	ExtractBatchParamsSharedInputsLocaleNd        ExtractBatchParamsSharedInputsLocale = "nd"
	ExtractBatchParamsSharedInputsLocaleNdZw      ExtractBatchParamsSharedInputsLocale = "nd-ZW"
	ExtractBatchParamsSharedInputsLocaleNdsDe     ExtractBatchParamsSharedInputsLocale = "nds-DE"
	ExtractBatchParamsSharedInputsLocaleNdsNl     ExtractBatchParamsSharedInputsLocale = "nds-NL"
	ExtractBatchParamsSharedInputsLocaleNe        ExtractBatchParamsSharedInputsLocale = "ne"
	ExtractBatchParamsSharedInputsLocaleNeIn      ExtractBatchParamsSharedInputsLocale = "ne-IN"
	ExtractBatchParamsSharedInputsLocaleNeNp      ExtractBatchParamsSharedInputsLocale = "ne-NP"
	ExtractBatchParamsSharedInputsLocaleNl        ExtractBatchParamsSharedInputsLocale = "nl"
	ExtractBatchParamsSharedInputsLocaleNlAw      ExtractBatchParamsSharedInputsLocale = "nl-AW"
	ExtractBatchParamsSharedInputsLocaleNlBe      ExtractBatchParamsSharedInputsLocale = "nl-BE"
	ExtractBatchParamsSharedInputsLocaleNlNl      ExtractBatchParamsSharedInputsLocale = "nl-NL"
	ExtractBatchParamsSharedInputsLocaleNn        ExtractBatchParamsSharedInputsLocale = "nn"
	ExtractBatchParamsSharedInputsLocaleNnNo      ExtractBatchParamsSharedInputsLocale = "nn-NO"
	ExtractBatchParamsSharedInputsLocaleNrZa      ExtractBatchParamsSharedInputsLocale = "nr-ZA"
	ExtractBatchParamsSharedInputsLocaleNsoZa     ExtractBatchParamsSharedInputsLocale = "nso-ZA"
	ExtractBatchParamsSharedInputsLocaleNyn       ExtractBatchParamsSharedInputsLocale = "nyn"
	ExtractBatchParamsSharedInputsLocaleNynUg     ExtractBatchParamsSharedInputsLocale = "nyn-UG"
	ExtractBatchParamsSharedInputsLocaleOcFr      ExtractBatchParamsSharedInputsLocale = "oc-FR"
	ExtractBatchParamsSharedInputsLocaleOm        ExtractBatchParamsSharedInputsLocale = "om"
	ExtractBatchParamsSharedInputsLocaleOmEt      ExtractBatchParamsSharedInputsLocale = "om-ET"
	ExtractBatchParamsSharedInputsLocaleOmKe      ExtractBatchParamsSharedInputsLocale = "om-KE"
	ExtractBatchParamsSharedInputsLocaleOr        ExtractBatchParamsSharedInputsLocale = "or"
	ExtractBatchParamsSharedInputsLocaleOrIn      ExtractBatchParamsSharedInputsLocale = "or-IN"
	ExtractBatchParamsSharedInputsLocaleOsRu      ExtractBatchParamsSharedInputsLocale = "os-RU"
	ExtractBatchParamsSharedInputsLocalePa        ExtractBatchParamsSharedInputsLocale = "pa"
	ExtractBatchParamsSharedInputsLocalePaArab    ExtractBatchParamsSharedInputsLocale = "pa-Arab"
	ExtractBatchParamsSharedInputsLocalePaArabPk  ExtractBatchParamsSharedInputsLocale = "pa-Arab-PK"
	ExtractBatchParamsSharedInputsLocalePaGuru    ExtractBatchParamsSharedInputsLocale = "pa-Guru"
	ExtractBatchParamsSharedInputsLocalePaGuruIn  ExtractBatchParamsSharedInputsLocale = "pa-Guru-IN"
	ExtractBatchParamsSharedInputsLocalePaIn      ExtractBatchParamsSharedInputsLocale = "pa-IN"
	ExtractBatchParamsSharedInputsLocalePaPk      ExtractBatchParamsSharedInputsLocale = "pa-PK"
	ExtractBatchParamsSharedInputsLocalePapAn     ExtractBatchParamsSharedInputsLocale = "pap-AN"
	ExtractBatchParamsSharedInputsLocalePl        ExtractBatchParamsSharedInputsLocale = "pl"
	ExtractBatchParamsSharedInputsLocalePlPl      ExtractBatchParamsSharedInputsLocale = "pl-PL"
	ExtractBatchParamsSharedInputsLocalePs        ExtractBatchParamsSharedInputsLocale = "ps"
	ExtractBatchParamsSharedInputsLocalePsAf      ExtractBatchParamsSharedInputsLocale = "ps-AF"
	ExtractBatchParamsSharedInputsLocalePt        ExtractBatchParamsSharedInputsLocale = "pt"
	ExtractBatchParamsSharedInputsLocalePtBr      ExtractBatchParamsSharedInputsLocale = "pt-BR"
	ExtractBatchParamsSharedInputsLocalePtGw      ExtractBatchParamsSharedInputsLocale = "pt-GW"
	ExtractBatchParamsSharedInputsLocalePtMz      ExtractBatchParamsSharedInputsLocale = "pt-MZ"
	ExtractBatchParamsSharedInputsLocalePtPt      ExtractBatchParamsSharedInputsLocale = "pt-PT"
	ExtractBatchParamsSharedInputsLocaleRm        ExtractBatchParamsSharedInputsLocale = "rm"
	ExtractBatchParamsSharedInputsLocaleRmCh      ExtractBatchParamsSharedInputsLocale = "rm-CH"
	ExtractBatchParamsSharedInputsLocaleRo        ExtractBatchParamsSharedInputsLocale = "ro"
	ExtractBatchParamsSharedInputsLocaleRoMd      ExtractBatchParamsSharedInputsLocale = "ro-MD"
	ExtractBatchParamsSharedInputsLocaleRoRo      ExtractBatchParamsSharedInputsLocale = "ro-RO"
	ExtractBatchParamsSharedInputsLocaleRof       ExtractBatchParamsSharedInputsLocale = "rof"
	ExtractBatchParamsSharedInputsLocaleRofTz     ExtractBatchParamsSharedInputsLocale = "rof-TZ"
	ExtractBatchParamsSharedInputsLocaleRu        ExtractBatchParamsSharedInputsLocale = "ru"
	ExtractBatchParamsSharedInputsLocaleRuMd      ExtractBatchParamsSharedInputsLocale = "ru-MD"
	ExtractBatchParamsSharedInputsLocaleRuRu      ExtractBatchParamsSharedInputsLocale = "ru-RU"
	ExtractBatchParamsSharedInputsLocaleRuUa      ExtractBatchParamsSharedInputsLocale = "ru-UA"
	ExtractBatchParamsSharedInputsLocaleRw        ExtractBatchParamsSharedInputsLocale = "rw"
	ExtractBatchParamsSharedInputsLocaleRwRw      ExtractBatchParamsSharedInputsLocale = "rw-RW"
	ExtractBatchParamsSharedInputsLocaleRwk       ExtractBatchParamsSharedInputsLocale = "rwk"
	ExtractBatchParamsSharedInputsLocaleRwkTz     ExtractBatchParamsSharedInputsLocale = "rwk-TZ"
	ExtractBatchParamsSharedInputsLocaleSaIn      ExtractBatchParamsSharedInputsLocale = "sa-IN"
	ExtractBatchParamsSharedInputsLocaleSaq       ExtractBatchParamsSharedInputsLocale = "saq"
	ExtractBatchParamsSharedInputsLocaleSaqKe     ExtractBatchParamsSharedInputsLocale = "saq-KE"
	ExtractBatchParamsSharedInputsLocaleScIt      ExtractBatchParamsSharedInputsLocale = "sc-IT"
	ExtractBatchParamsSharedInputsLocaleSdIn      ExtractBatchParamsSharedInputsLocale = "sd-IN"
	ExtractBatchParamsSharedInputsLocaleSeNo      ExtractBatchParamsSharedInputsLocale = "se-NO"
	ExtractBatchParamsSharedInputsLocaleSeh       ExtractBatchParamsSharedInputsLocale = "seh"
	ExtractBatchParamsSharedInputsLocaleSehMz     ExtractBatchParamsSharedInputsLocale = "seh-MZ"
	ExtractBatchParamsSharedInputsLocaleSes       ExtractBatchParamsSharedInputsLocale = "ses"
	ExtractBatchParamsSharedInputsLocaleSesMl     ExtractBatchParamsSharedInputsLocale = "ses-ML"
	ExtractBatchParamsSharedInputsLocaleSg        ExtractBatchParamsSharedInputsLocale = "sg"
	ExtractBatchParamsSharedInputsLocaleSgCf      ExtractBatchParamsSharedInputsLocale = "sg-CF"
	ExtractBatchParamsSharedInputsLocaleShi       ExtractBatchParamsSharedInputsLocale = "shi"
	ExtractBatchParamsSharedInputsLocaleShiLatn   ExtractBatchParamsSharedInputsLocale = "shi-Latn"
	ExtractBatchParamsSharedInputsLocaleShiLatnMa ExtractBatchParamsSharedInputsLocale = "shi-Latn-MA"
	ExtractBatchParamsSharedInputsLocaleShiTfng   ExtractBatchParamsSharedInputsLocale = "shi-Tfng"
	ExtractBatchParamsSharedInputsLocaleShiTfngMa ExtractBatchParamsSharedInputsLocale = "shi-Tfng-MA"
	ExtractBatchParamsSharedInputsLocaleShsCa     ExtractBatchParamsSharedInputsLocale = "shs-CA"
	ExtractBatchParamsSharedInputsLocaleSi        ExtractBatchParamsSharedInputsLocale = "si"
	ExtractBatchParamsSharedInputsLocaleSiLk      ExtractBatchParamsSharedInputsLocale = "si-LK"
	ExtractBatchParamsSharedInputsLocaleSidEt     ExtractBatchParamsSharedInputsLocale = "sid-ET"
	ExtractBatchParamsSharedInputsLocaleSk        ExtractBatchParamsSharedInputsLocale = "sk"
	ExtractBatchParamsSharedInputsLocaleSkSk      ExtractBatchParamsSharedInputsLocale = "sk-SK"
	ExtractBatchParamsSharedInputsLocaleSl        ExtractBatchParamsSharedInputsLocale = "sl"
	ExtractBatchParamsSharedInputsLocaleSlSi      ExtractBatchParamsSharedInputsLocale = "sl-SI"
	ExtractBatchParamsSharedInputsLocaleSn        ExtractBatchParamsSharedInputsLocale = "sn"
	ExtractBatchParamsSharedInputsLocaleSnZw      ExtractBatchParamsSharedInputsLocale = "sn-ZW"
	ExtractBatchParamsSharedInputsLocaleSo        ExtractBatchParamsSharedInputsLocale = "so"
	ExtractBatchParamsSharedInputsLocaleSoDj      ExtractBatchParamsSharedInputsLocale = "so-DJ"
	ExtractBatchParamsSharedInputsLocaleSoEt      ExtractBatchParamsSharedInputsLocale = "so-ET"
	ExtractBatchParamsSharedInputsLocaleSoKe      ExtractBatchParamsSharedInputsLocale = "so-KE"
	ExtractBatchParamsSharedInputsLocaleSoSo      ExtractBatchParamsSharedInputsLocale = "so-SO"
	ExtractBatchParamsSharedInputsLocaleSq        ExtractBatchParamsSharedInputsLocale = "sq"
	ExtractBatchParamsSharedInputsLocaleSqAl      ExtractBatchParamsSharedInputsLocale = "sq-AL"
	ExtractBatchParamsSharedInputsLocaleSqMk      ExtractBatchParamsSharedInputsLocale = "sq-MK"
	ExtractBatchParamsSharedInputsLocaleSr        ExtractBatchParamsSharedInputsLocale = "sr"
	ExtractBatchParamsSharedInputsLocaleSrCyrl    ExtractBatchParamsSharedInputsLocale = "sr-Cyrl"
	ExtractBatchParamsSharedInputsLocaleSrCyrlBa  ExtractBatchParamsSharedInputsLocale = "sr-Cyrl-BA"
	ExtractBatchParamsSharedInputsLocaleSrCyrlMe  ExtractBatchParamsSharedInputsLocale = "sr-Cyrl-ME"
	ExtractBatchParamsSharedInputsLocaleSrCyrlRs  ExtractBatchParamsSharedInputsLocale = "sr-Cyrl-RS"
	ExtractBatchParamsSharedInputsLocaleSrLatn    ExtractBatchParamsSharedInputsLocale = "sr-Latn"
	ExtractBatchParamsSharedInputsLocaleSrLatnBa  ExtractBatchParamsSharedInputsLocale = "sr-Latn-BA"
	ExtractBatchParamsSharedInputsLocaleSrLatnMe  ExtractBatchParamsSharedInputsLocale = "sr-Latn-ME"
	ExtractBatchParamsSharedInputsLocaleSrLatnRs  ExtractBatchParamsSharedInputsLocale = "sr-Latn-RS"
	ExtractBatchParamsSharedInputsLocaleSrMe      ExtractBatchParamsSharedInputsLocale = "sr-ME"
	ExtractBatchParamsSharedInputsLocaleSrRs      ExtractBatchParamsSharedInputsLocale = "sr-RS"
	ExtractBatchParamsSharedInputsLocaleSSZa      ExtractBatchParamsSharedInputsLocale = "ss-ZA"
	ExtractBatchParamsSharedInputsLocaleStZa      ExtractBatchParamsSharedInputsLocale = "st-ZA"
	ExtractBatchParamsSharedInputsLocaleSv        ExtractBatchParamsSharedInputsLocale = "sv"
	ExtractBatchParamsSharedInputsLocaleSvFi      ExtractBatchParamsSharedInputsLocale = "sv-FI"
	ExtractBatchParamsSharedInputsLocaleSvSe      ExtractBatchParamsSharedInputsLocale = "sv-SE"
	ExtractBatchParamsSharedInputsLocaleSw        ExtractBatchParamsSharedInputsLocale = "sw"
	ExtractBatchParamsSharedInputsLocaleSwKe      ExtractBatchParamsSharedInputsLocale = "sw-KE"
	ExtractBatchParamsSharedInputsLocaleSwTz      ExtractBatchParamsSharedInputsLocale = "sw-TZ"
	ExtractBatchParamsSharedInputsLocaleTa        ExtractBatchParamsSharedInputsLocale = "ta"
	ExtractBatchParamsSharedInputsLocaleTaIn      ExtractBatchParamsSharedInputsLocale = "ta-IN"
	ExtractBatchParamsSharedInputsLocaleTaLk      ExtractBatchParamsSharedInputsLocale = "ta-LK"
	ExtractBatchParamsSharedInputsLocaleTe        ExtractBatchParamsSharedInputsLocale = "te"
	ExtractBatchParamsSharedInputsLocaleTeIn      ExtractBatchParamsSharedInputsLocale = "te-IN"
	ExtractBatchParamsSharedInputsLocaleTeo       ExtractBatchParamsSharedInputsLocale = "teo"
	ExtractBatchParamsSharedInputsLocaleTeoKe     ExtractBatchParamsSharedInputsLocale = "teo-KE"
	ExtractBatchParamsSharedInputsLocaleTeoUg     ExtractBatchParamsSharedInputsLocale = "teo-UG"
	ExtractBatchParamsSharedInputsLocaleTgTj      ExtractBatchParamsSharedInputsLocale = "tg-TJ"
	ExtractBatchParamsSharedInputsLocaleTh        ExtractBatchParamsSharedInputsLocale = "th"
	ExtractBatchParamsSharedInputsLocaleThTh      ExtractBatchParamsSharedInputsLocale = "th-TH"
	ExtractBatchParamsSharedInputsLocaleTi        ExtractBatchParamsSharedInputsLocale = "ti"
	ExtractBatchParamsSharedInputsLocaleTiEr      ExtractBatchParamsSharedInputsLocale = "ti-ER"
	ExtractBatchParamsSharedInputsLocaleTiEt      ExtractBatchParamsSharedInputsLocale = "ti-ET"
	ExtractBatchParamsSharedInputsLocaleTigEr     ExtractBatchParamsSharedInputsLocale = "tig-ER"
	ExtractBatchParamsSharedInputsLocaleTkTm      ExtractBatchParamsSharedInputsLocale = "tk-TM"
	ExtractBatchParamsSharedInputsLocaleTlPh      ExtractBatchParamsSharedInputsLocale = "tl-PH"
	ExtractBatchParamsSharedInputsLocaleTnZa      ExtractBatchParamsSharedInputsLocale = "tn-ZA"
	ExtractBatchParamsSharedInputsLocaleTo        ExtractBatchParamsSharedInputsLocale = "to"
	ExtractBatchParamsSharedInputsLocaleToTo      ExtractBatchParamsSharedInputsLocale = "to-TO"
	ExtractBatchParamsSharedInputsLocaleTr        ExtractBatchParamsSharedInputsLocale = "tr"
	ExtractBatchParamsSharedInputsLocaleTrCy      ExtractBatchParamsSharedInputsLocale = "tr-CY"
	ExtractBatchParamsSharedInputsLocaleTrTr      ExtractBatchParamsSharedInputsLocale = "tr-TR"
	ExtractBatchParamsSharedInputsLocaleTsZa      ExtractBatchParamsSharedInputsLocale = "ts-ZA"
	ExtractBatchParamsSharedInputsLocaleTtRu      ExtractBatchParamsSharedInputsLocale = "tt-RU"
	ExtractBatchParamsSharedInputsLocaleTzm       ExtractBatchParamsSharedInputsLocale = "tzm"
	ExtractBatchParamsSharedInputsLocaleTzmLatn   ExtractBatchParamsSharedInputsLocale = "tzm-Latn"
	ExtractBatchParamsSharedInputsLocaleTzmLatnMa ExtractBatchParamsSharedInputsLocale = "tzm-Latn-MA"
	ExtractBatchParamsSharedInputsLocaleUgCn      ExtractBatchParamsSharedInputsLocale = "ug-CN"
	ExtractBatchParamsSharedInputsLocaleUk        ExtractBatchParamsSharedInputsLocale = "uk"
	ExtractBatchParamsSharedInputsLocaleUkUa      ExtractBatchParamsSharedInputsLocale = "uk-UA"
	ExtractBatchParamsSharedInputsLocaleUnmUs     ExtractBatchParamsSharedInputsLocale = "unm-US"
	ExtractBatchParamsSharedInputsLocaleUr        ExtractBatchParamsSharedInputsLocale = "ur"
	ExtractBatchParamsSharedInputsLocaleUrIn      ExtractBatchParamsSharedInputsLocale = "ur-IN"
	ExtractBatchParamsSharedInputsLocaleUrPk      ExtractBatchParamsSharedInputsLocale = "ur-PK"
	ExtractBatchParamsSharedInputsLocaleUz        ExtractBatchParamsSharedInputsLocale = "uz"
	ExtractBatchParamsSharedInputsLocaleUzArab    ExtractBatchParamsSharedInputsLocale = "uz-Arab"
	ExtractBatchParamsSharedInputsLocaleUzArabAf  ExtractBatchParamsSharedInputsLocale = "uz-Arab-AF"
	ExtractBatchParamsSharedInputsLocaleUzCyrl    ExtractBatchParamsSharedInputsLocale = "uz-Cyrl"
	ExtractBatchParamsSharedInputsLocaleUzCyrlUz  ExtractBatchParamsSharedInputsLocale = "uz-Cyrl-UZ"
	ExtractBatchParamsSharedInputsLocaleUzLatn    ExtractBatchParamsSharedInputsLocale = "uz-Latn"
	ExtractBatchParamsSharedInputsLocaleUzLatnUz  ExtractBatchParamsSharedInputsLocale = "uz-Latn-UZ"
	ExtractBatchParamsSharedInputsLocaleUzUz      ExtractBatchParamsSharedInputsLocale = "uz-UZ"
	ExtractBatchParamsSharedInputsLocaleVeZa      ExtractBatchParamsSharedInputsLocale = "ve-ZA"
	ExtractBatchParamsSharedInputsLocaleVi        ExtractBatchParamsSharedInputsLocale = "vi"
	ExtractBatchParamsSharedInputsLocaleViVn      ExtractBatchParamsSharedInputsLocale = "vi-VN"
	ExtractBatchParamsSharedInputsLocaleVun       ExtractBatchParamsSharedInputsLocale = "vun"
	ExtractBatchParamsSharedInputsLocaleVunTz     ExtractBatchParamsSharedInputsLocale = "vun-TZ"
	ExtractBatchParamsSharedInputsLocaleWaBe      ExtractBatchParamsSharedInputsLocale = "wa-BE"
	ExtractBatchParamsSharedInputsLocaleWaeCh     ExtractBatchParamsSharedInputsLocale = "wae-CH"
	ExtractBatchParamsSharedInputsLocaleWalEt     ExtractBatchParamsSharedInputsLocale = "wal-ET"
	ExtractBatchParamsSharedInputsLocaleWoSn      ExtractBatchParamsSharedInputsLocale = "wo-SN"
	ExtractBatchParamsSharedInputsLocaleXhZa      ExtractBatchParamsSharedInputsLocale = "xh-ZA"
	ExtractBatchParamsSharedInputsLocaleXog       ExtractBatchParamsSharedInputsLocale = "xog"
	ExtractBatchParamsSharedInputsLocaleXogUg     ExtractBatchParamsSharedInputsLocale = "xog-UG"
	ExtractBatchParamsSharedInputsLocaleYiUs      ExtractBatchParamsSharedInputsLocale = "yi-US"
	ExtractBatchParamsSharedInputsLocaleYo        ExtractBatchParamsSharedInputsLocale = "yo"
	ExtractBatchParamsSharedInputsLocaleYoNg      ExtractBatchParamsSharedInputsLocale = "yo-NG"
	ExtractBatchParamsSharedInputsLocaleYueHk     ExtractBatchParamsSharedInputsLocale = "yue-HK"
	ExtractBatchParamsSharedInputsLocaleZh        ExtractBatchParamsSharedInputsLocale = "zh"
	ExtractBatchParamsSharedInputsLocaleZhCn      ExtractBatchParamsSharedInputsLocale = "zh-CN"
	ExtractBatchParamsSharedInputsLocaleZhHk      ExtractBatchParamsSharedInputsLocale = "zh-HK"
	ExtractBatchParamsSharedInputsLocaleZhHans    ExtractBatchParamsSharedInputsLocale = "zh-Hans"
	ExtractBatchParamsSharedInputsLocaleZhHansCn  ExtractBatchParamsSharedInputsLocale = "zh-Hans-CN"
	ExtractBatchParamsSharedInputsLocaleZhHansHk  ExtractBatchParamsSharedInputsLocale = "zh-Hans-HK"
	ExtractBatchParamsSharedInputsLocaleZhHansMo  ExtractBatchParamsSharedInputsLocale = "zh-Hans-MO"
	ExtractBatchParamsSharedInputsLocaleZhHansSg  ExtractBatchParamsSharedInputsLocale = "zh-Hans-SG"
	ExtractBatchParamsSharedInputsLocaleZhHant    ExtractBatchParamsSharedInputsLocale = "zh-Hant"
	ExtractBatchParamsSharedInputsLocaleZhHantHk  ExtractBatchParamsSharedInputsLocale = "zh-Hant-HK"
	ExtractBatchParamsSharedInputsLocaleZhHantMo  ExtractBatchParamsSharedInputsLocale = "zh-Hant-MO"
	ExtractBatchParamsSharedInputsLocaleZhHantTw  ExtractBatchParamsSharedInputsLocale = "zh-Hant-TW"
	ExtractBatchParamsSharedInputsLocaleZhSg      ExtractBatchParamsSharedInputsLocale = "zh-SG"
	ExtractBatchParamsSharedInputsLocaleZhTw      ExtractBatchParamsSharedInputsLocale = "zh-TW"
	ExtractBatchParamsSharedInputsLocaleZu        ExtractBatchParamsSharedInputsLocale = "zu"
	ExtractBatchParamsSharedInputsLocaleZuZa      ExtractBatchParamsSharedInputsLocale = "zu-ZA"
	ExtractBatchParamsSharedInputsLocaleAuto      ExtractBatchParamsSharedInputsLocale = "auto"
)

type ExtractBatchParamsSharedInputsNetworkCapture struct {
	Validation                  param.Opt[bool]    `json:"validation,omitzero"`
	WaitForRequestsCount        param.Opt[float64] `json:"wait_for_requests_count,omitzero"`
	WaitForRequestsCountTimeout param.Opt[float64] `json:"wait_for_requests_count_timeout,omitzero"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method,omitzero"`
	// Resource type for network capture filtering
	ResourceType ExtractBatchParamsSharedInputsNetworkCaptureResourceTypeUnion `json:"resource_type,omitzero"`
	StatusCode   ExtractBatchParamsSharedInputsNetworkCaptureStatusCodeUnion   `json:"status_code,omitzero"`
	URL          ExtractBatchParamsSharedInputsNetworkCaptureURL               `json:"url,omitzero"`
	paramObj
}

func (r ExtractBatchParamsSharedInputsNetworkCapture) MarshalJSON() (data []byte, err error) {
	type shadow ExtractBatchParamsSharedInputsNetworkCapture
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractBatchParamsSharedInputsNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractBatchParamsSharedInputsNetworkCapture](
		"method", "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsSharedInputsNetworkCaptureResourceTypeUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsSharedInputsNetworkCaptureResourceTypeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractBatchParamsSharedInputsNetworkCaptureResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsSharedInputsNetworkCaptureResourceTypeUnion) asAny() any {
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
type ExtractBatchParamsSharedInputsNetworkCaptureStatusCodeUnion struct {
	OfFloat      param.Opt[float64] `json:",omitzero,inline"`
	OfFloatArray []float64          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsSharedInputsNetworkCaptureStatusCodeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfFloatArray)
}
func (u *ExtractBatchParamsSharedInputsNetworkCaptureStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsSharedInputsNetworkCaptureStatusCodeUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfFloatArray) {
		return &u.OfFloatArray
	}
	return nil
}

// The property Value is required.
type ExtractBatchParamsSharedInputsNetworkCaptureURL struct {
	Value string `json:"value" api:"required"`
	// Any of "exact", "contains".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ExtractBatchParamsSharedInputsNetworkCaptureURL) MarshalJSON() (data []byte, err error) {
	type shadow ExtractBatchParamsSharedInputsNetworkCaptureURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractBatchParamsSharedInputsNetworkCaptureURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractBatchParamsSharedInputsNetworkCaptureURL](
		"type", "exact", "contains",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsSharedInputsParserUnion struct {
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsSharedInputsParserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap, u.OfString)
}
func (u *ExtractBatchParamsSharedInputsParserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsSharedInputsParserUnion) asAny() any {
	if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Referrer policy for the request
type ExtractBatchParamsSharedInputsReferrerType string

const (
	ExtractBatchParamsSharedInputsReferrerTypeRandom     ExtractBatchParamsSharedInputsReferrerType = "random"
	ExtractBatchParamsSharedInputsReferrerTypeNoReferer  ExtractBatchParamsSharedInputsReferrerType = "no-referer"
	ExtractBatchParamsSharedInputsReferrerTypeSameOrigin ExtractBatchParamsSharedInputsReferrerType = "same-origin"
	ExtractBatchParamsSharedInputsReferrerTypeGoogle     ExtractBatchParamsSharedInputsReferrerType = "google"
	ExtractBatchParamsSharedInputsReferrerTypeBing       ExtractBatchParamsSharedInputsReferrerType = "bing"
	ExtractBatchParamsSharedInputsReferrerTypeFacebook   ExtractBatchParamsSharedInputsReferrerType = "facebook"
	ExtractBatchParamsSharedInputsReferrerTypeTwitter    ExtractBatchParamsSharedInputsReferrerType = "twitter"
	ExtractBatchParamsSharedInputsReferrerTypeInstagram  ExtractBatchParamsSharedInputsReferrerType = "instagram"
)

type ExtractBatchParamsSharedInputsSession struct {
	ID                  param.Opt[string]  `json:"id,omitzero"`
	PrefetchUserbrowser param.Opt[bool]    `json:"prefetch_userbrowser,omitzero"`
	Retry               param.Opt[bool]    `json:"retry,omitzero"`
	Timeout             param.Opt[float64] `json:"timeout,omitzero"`
	paramObj
}

func (r ExtractBatchParamsSharedInputsSession) MarshalJSON() (data []byte, err error) {
	type shadow ExtractBatchParamsSharedInputsSession
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractBatchParamsSharedInputsSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsSharedInputsSkillUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsSharedInputsSkillUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractBatchParamsSharedInputsSkillUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsSharedInputsSkillUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

type MapParams struct {
	// Url to map.
	URL string `json:"url" api:"required"`
	// Maximum number of links to return.
	Limit param.Opt[int64] `json:"limit,omitzero"`
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
	// "YT", "ZA", "ZM", "ZW".
	Country MapParamsCountry `json:"country,omitzero"`
	// Includes subdomains of the main domain in the mapping process.
	//
	// Any of "domain", "subdomain", "all".
	DomainFilter MapParamsDomainFilter `json:"domain_filter,omitzero"`
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
	Locale MapParamsLocale `json:"locale,omitzero"`
	// Sitemap and other methods will be used together to find URLs.
	//
	// Any of "skip", "include", "only".
	Sitemap MapParamsSitemap `json:"sitemap,omitzero"`
	paramObj
}

func (r MapParams) MarshalJSON() (data []byte, err error) {
	type shadow MapParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MapParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Country code for geolocation and proxy selection
type MapParamsCountry string

const (
	MapParamsCountryAd MapParamsCountry = "AD"
	MapParamsCountryAe MapParamsCountry = "AE"
	MapParamsCountryAf MapParamsCountry = "AF"
	MapParamsCountryAg MapParamsCountry = "AG"
	MapParamsCountryAI MapParamsCountry = "AI"
	MapParamsCountryAl MapParamsCountry = "AL"
	MapParamsCountryAm MapParamsCountry = "AM"
	MapParamsCountryAo MapParamsCountry = "AO"
	MapParamsCountryAq MapParamsCountry = "AQ"
	MapParamsCountryAr MapParamsCountry = "AR"
	MapParamsCountryAs MapParamsCountry = "AS"
	MapParamsCountryAt MapParamsCountry = "AT"
	MapParamsCountryAu MapParamsCountry = "AU"
	MapParamsCountryAw MapParamsCountry = "AW"
	MapParamsCountryAx MapParamsCountry = "AX"
	MapParamsCountryAz MapParamsCountry = "AZ"
	MapParamsCountryBa MapParamsCountry = "BA"
	MapParamsCountryBb MapParamsCountry = "BB"
	MapParamsCountryBd MapParamsCountry = "BD"
	MapParamsCountryBe MapParamsCountry = "BE"
	MapParamsCountryBf MapParamsCountry = "BF"
	MapParamsCountryBg MapParamsCountry = "BG"
	MapParamsCountryBh MapParamsCountry = "BH"
	MapParamsCountryBi MapParamsCountry = "BI"
	MapParamsCountryBj MapParamsCountry = "BJ"
	MapParamsCountryBl MapParamsCountry = "BL"
	MapParamsCountryBm MapParamsCountry = "BM"
	MapParamsCountryBn MapParamsCountry = "BN"
	MapParamsCountryBo MapParamsCountry = "BO"
	MapParamsCountryBq MapParamsCountry = "BQ"
	MapParamsCountryBr MapParamsCountry = "BR"
	MapParamsCountryBs MapParamsCountry = "BS"
	MapParamsCountryBt MapParamsCountry = "BT"
	MapParamsCountryBv MapParamsCountry = "BV"
	MapParamsCountryBw MapParamsCountry = "BW"
	MapParamsCountryBy MapParamsCountry = "BY"
	MapParamsCountryBz MapParamsCountry = "BZ"
	MapParamsCountryCa MapParamsCountry = "CA"
	MapParamsCountryCc MapParamsCountry = "CC"
	MapParamsCountryCd MapParamsCountry = "CD"
	MapParamsCountryCf MapParamsCountry = "CF"
	MapParamsCountryCg MapParamsCountry = "CG"
	MapParamsCountryCh MapParamsCountry = "CH"
	MapParamsCountryCi MapParamsCountry = "CI"
	MapParamsCountryCk MapParamsCountry = "CK"
	MapParamsCountryCl MapParamsCountry = "CL"
	MapParamsCountryCm MapParamsCountry = "CM"
	MapParamsCountryCn MapParamsCountry = "CN"
	MapParamsCountryCo MapParamsCountry = "CO"
	MapParamsCountryCr MapParamsCountry = "CR"
	MapParamsCountryCu MapParamsCountry = "CU"
	MapParamsCountryCv MapParamsCountry = "CV"
	MapParamsCountryCw MapParamsCountry = "CW"
	MapParamsCountryCx MapParamsCountry = "CX"
	MapParamsCountryCy MapParamsCountry = "CY"
	MapParamsCountryCz MapParamsCountry = "CZ"
	MapParamsCountryDe MapParamsCountry = "DE"
	MapParamsCountryDj MapParamsCountry = "DJ"
	MapParamsCountryDk MapParamsCountry = "DK"
	MapParamsCountryDm MapParamsCountry = "DM"
	MapParamsCountryDo MapParamsCountry = "DO"
	MapParamsCountryDz MapParamsCountry = "DZ"
	MapParamsCountryEc MapParamsCountry = "EC"
	MapParamsCountryEe MapParamsCountry = "EE"
	MapParamsCountryEg MapParamsCountry = "EG"
	MapParamsCountryEh MapParamsCountry = "EH"
	MapParamsCountryEr MapParamsCountry = "ER"
	MapParamsCountryEs MapParamsCountry = "ES"
	MapParamsCountryEt MapParamsCountry = "ET"
	MapParamsCountryFi MapParamsCountry = "FI"
	MapParamsCountryFj MapParamsCountry = "FJ"
	MapParamsCountryFk MapParamsCountry = "FK"
	MapParamsCountryFm MapParamsCountry = "FM"
	MapParamsCountryFo MapParamsCountry = "FO"
	MapParamsCountryFr MapParamsCountry = "FR"
	MapParamsCountryGa MapParamsCountry = "GA"
	MapParamsCountryGB MapParamsCountry = "GB"
	MapParamsCountryGd MapParamsCountry = "GD"
	MapParamsCountryGe MapParamsCountry = "GE"
	MapParamsCountryGf MapParamsCountry = "GF"
	MapParamsCountryGg MapParamsCountry = "GG"
	MapParamsCountryGh MapParamsCountry = "GH"
	MapParamsCountryGi MapParamsCountry = "GI"
	MapParamsCountryGl MapParamsCountry = "GL"
	MapParamsCountryGm MapParamsCountry = "GM"
	MapParamsCountryGn MapParamsCountry = "GN"
	MapParamsCountryGp MapParamsCountry = "GP"
	MapParamsCountryGq MapParamsCountry = "GQ"
	MapParamsCountryGr MapParamsCountry = "GR"
	MapParamsCountryGs MapParamsCountry = "GS"
	MapParamsCountryGt MapParamsCountry = "GT"
	MapParamsCountryGu MapParamsCountry = "GU"
	MapParamsCountryGw MapParamsCountry = "GW"
	MapParamsCountryGy MapParamsCountry = "GY"
	MapParamsCountryHk MapParamsCountry = "HK"
	MapParamsCountryHm MapParamsCountry = "HM"
	MapParamsCountryHn MapParamsCountry = "HN"
	MapParamsCountryHr MapParamsCountry = "HR"
	MapParamsCountryHt MapParamsCountry = "HT"
	MapParamsCountryHu MapParamsCountry = "HU"
	MapParamsCountryID MapParamsCountry = "ID"
	MapParamsCountryIe MapParamsCountry = "IE"
	MapParamsCountryIl MapParamsCountry = "IL"
	MapParamsCountryIm MapParamsCountry = "IM"
	MapParamsCountryIn MapParamsCountry = "IN"
	MapParamsCountryIo MapParamsCountry = "IO"
	MapParamsCountryIq MapParamsCountry = "IQ"
	MapParamsCountryIr MapParamsCountry = "IR"
	MapParamsCountryIs MapParamsCountry = "IS"
	MapParamsCountryIt MapParamsCountry = "IT"
	MapParamsCountryJe MapParamsCountry = "JE"
	MapParamsCountryJm MapParamsCountry = "JM"
	MapParamsCountryJo MapParamsCountry = "JO"
	MapParamsCountryJp MapParamsCountry = "JP"
	MapParamsCountryKe MapParamsCountry = "KE"
	MapParamsCountryKg MapParamsCountry = "KG"
	MapParamsCountryKh MapParamsCountry = "KH"
	MapParamsCountryKi MapParamsCountry = "KI"
	MapParamsCountryKm MapParamsCountry = "KM"
	MapParamsCountryKn MapParamsCountry = "KN"
	MapParamsCountryKp MapParamsCountry = "KP"
	MapParamsCountryKr MapParamsCountry = "KR"
	MapParamsCountryKw MapParamsCountry = "KW"
	MapParamsCountryKy MapParamsCountry = "KY"
	MapParamsCountryKz MapParamsCountry = "KZ"
	MapParamsCountryLa MapParamsCountry = "LA"
	MapParamsCountryLb MapParamsCountry = "LB"
	MapParamsCountryLc MapParamsCountry = "LC"
	MapParamsCountryLi MapParamsCountry = "LI"
	MapParamsCountryLk MapParamsCountry = "LK"
	MapParamsCountryLr MapParamsCountry = "LR"
	MapParamsCountryLs MapParamsCountry = "LS"
	MapParamsCountryLt MapParamsCountry = "LT"
	MapParamsCountryLu MapParamsCountry = "LU"
	MapParamsCountryLv MapParamsCountry = "LV"
	MapParamsCountryLy MapParamsCountry = "LY"
	MapParamsCountryMa MapParamsCountry = "MA"
	MapParamsCountryMc MapParamsCountry = "MC"
	MapParamsCountryMd MapParamsCountry = "MD"
	MapParamsCountryMe MapParamsCountry = "ME"
	MapParamsCountryMf MapParamsCountry = "MF"
	MapParamsCountryMg MapParamsCountry = "MG"
	MapParamsCountryMh MapParamsCountry = "MH"
	MapParamsCountryMk MapParamsCountry = "MK"
	MapParamsCountryMl MapParamsCountry = "ML"
	MapParamsCountryMm MapParamsCountry = "MM"
	MapParamsCountryMn MapParamsCountry = "MN"
	MapParamsCountryMo MapParamsCountry = "MO"
	MapParamsCountryMp MapParamsCountry = "MP"
	MapParamsCountryMq MapParamsCountry = "MQ"
	MapParamsCountryMr MapParamsCountry = "MR"
	MapParamsCountryMs MapParamsCountry = "MS"
	MapParamsCountryMt MapParamsCountry = "MT"
	MapParamsCountryMu MapParamsCountry = "MU"
	MapParamsCountryMv MapParamsCountry = "MV"
	MapParamsCountryMw MapParamsCountry = "MW"
	MapParamsCountryMx MapParamsCountry = "MX"
	MapParamsCountryMy MapParamsCountry = "MY"
	MapParamsCountryMz MapParamsCountry = "MZ"
	MapParamsCountryNa MapParamsCountry = "NA"
	MapParamsCountryNc MapParamsCountry = "NC"
	MapParamsCountryNe MapParamsCountry = "NE"
	MapParamsCountryNf MapParamsCountry = "NF"
	MapParamsCountryNg MapParamsCountry = "NG"
	MapParamsCountryNi MapParamsCountry = "NI"
	MapParamsCountryNl MapParamsCountry = "NL"
	MapParamsCountryNo MapParamsCountry = "NO"
	MapParamsCountryNp MapParamsCountry = "NP"
	MapParamsCountryNr MapParamsCountry = "NR"
	MapParamsCountryNu MapParamsCountry = "NU"
	MapParamsCountryNz MapParamsCountry = "NZ"
	MapParamsCountryOm MapParamsCountry = "OM"
	MapParamsCountryPa MapParamsCountry = "PA"
	MapParamsCountryPe MapParamsCountry = "PE"
	MapParamsCountryPf MapParamsCountry = "PF"
	MapParamsCountryPg MapParamsCountry = "PG"
	MapParamsCountryPh MapParamsCountry = "PH"
	MapParamsCountryPk MapParamsCountry = "PK"
	MapParamsCountryPl MapParamsCountry = "PL"
	MapParamsCountryPm MapParamsCountry = "PM"
	MapParamsCountryPn MapParamsCountry = "PN"
	MapParamsCountryPr MapParamsCountry = "PR"
	MapParamsCountryPs MapParamsCountry = "PS"
	MapParamsCountryPt MapParamsCountry = "PT"
	MapParamsCountryPw MapParamsCountry = "PW"
	MapParamsCountryPy MapParamsCountry = "PY"
	MapParamsCountryQa MapParamsCountry = "QA"
	MapParamsCountryRe MapParamsCountry = "RE"
	MapParamsCountryRo MapParamsCountry = "RO"
	MapParamsCountryRs MapParamsCountry = "RS"
	MapParamsCountryRu MapParamsCountry = "RU"
	MapParamsCountryRw MapParamsCountry = "RW"
	MapParamsCountrySa MapParamsCountry = "SA"
	MapParamsCountrySb MapParamsCountry = "SB"
	MapParamsCountrySc MapParamsCountry = "SC"
	MapParamsCountrySd MapParamsCountry = "SD"
	MapParamsCountrySe MapParamsCountry = "SE"
	MapParamsCountrySg MapParamsCountry = "SG"
	MapParamsCountrySh MapParamsCountry = "SH"
	MapParamsCountrySi MapParamsCountry = "SI"
	MapParamsCountrySj MapParamsCountry = "SJ"
	MapParamsCountrySk MapParamsCountry = "SK"
	MapParamsCountrySl MapParamsCountry = "SL"
	MapParamsCountrySm MapParamsCountry = "SM"
	MapParamsCountrySn MapParamsCountry = "SN"
	MapParamsCountrySo MapParamsCountry = "SO"
	MapParamsCountrySr MapParamsCountry = "SR"
	MapParamsCountrySS MapParamsCountry = "SS"
	MapParamsCountrySt MapParamsCountry = "ST"
	MapParamsCountrySv MapParamsCountry = "SV"
	MapParamsCountrySx MapParamsCountry = "SX"
	MapParamsCountrySy MapParamsCountry = "SY"
	MapParamsCountrySz MapParamsCountry = "SZ"
	MapParamsCountryTc MapParamsCountry = "TC"
	MapParamsCountryTd MapParamsCountry = "TD"
	MapParamsCountryTf MapParamsCountry = "TF"
	MapParamsCountryTg MapParamsCountry = "TG"
	MapParamsCountryTh MapParamsCountry = "TH"
	MapParamsCountryTj MapParamsCountry = "TJ"
	MapParamsCountryTk MapParamsCountry = "TK"
	MapParamsCountryTl MapParamsCountry = "TL"
	MapParamsCountryTm MapParamsCountry = "TM"
	MapParamsCountryTn MapParamsCountry = "TN"
	MapParamsCountryTo MapParamsCountry = "TO"
	MapParamsCountryTr MapParamsCountry = "TR"
	MapParamsCountryTt MapParamsCountry = "TT"
	MapParamsCountryTv MapParamsCountry = "TV"
	MapParamsCountryTw MapParamsCountry = "TW"
	MapParamsCountryTz MapParamsCountry = "TZ"
	MapParamsCountryUa MapParamsCountry = "UA"
	MapParamsCountryUg MapParamsCountry = "UG"
	MapParamsCountryUm MapParamsCountry = "UM"
	MapParamsCountryUs MapParamsCountry = "US"
	MapParamsCountryUy MapParamsCountry = "UY"
	MapParamsCountryUz MapParamsCountry = "UZ"
	MapParamsCountryVa MapParamsCountry = "VA"
	MapParamsCountryVc MapParamsCountry = "VC"
	MapParamsCountryVe MapParamsCountry = "VE"
	MapParamsCountryVg MapParamsCountry = "VG"
	MapParamsCountryVi MapParamsCountry = "VI"
	MapParamsCountryVn MapParamsCountry = "VN"
	MapParamsCountryVu MapParamsCountry = "VU"
	MapParamsCountryWf MapParamsCountry = "WF"
	MapParamsCountryWs MapParamsCountry = "WS"
	MapParamsCountryXk MapParamsCountry = "XK"
	MapParamsCountryYe MapParamsCountry = "YE"
	MapParamsCountryYt MapParamsCountry = "YT"
	MapParamsCountryZa MapParamsCountry = "ZA"
	MapParamsCountryZm MapParamsCountry = "ZM"
	MapParamsCountryZw MapParamsCountry = "ZW"
)

// Includes subdomains of the main domain in the mapping process.
type MapParamsDomainFilter string

const (
	MapParamsDomainFilterDomain    MapParamsDomainFilter = "domain"
	MapParamsDomainFilterSubdomain MapParamsDomainFilter = "subdomain"
	MapParamsDomainFilterAll       MapParamsDomainFilter = "all"
)

// Locale for browser language and region settings
type MapParamsLocale string

const (
	MapParamsLocaleAaDj      MapParamsLocale = "aa-DJ"
	MapParamsLocaleAaEr      MapParamsLocale = "aa-ER"
	MapParamsLocaleAaEt      MapParamsLocale = "aa-ET"
	MapParamsLocaleAf        MapParamsLocale = "af"
	MapParamsLocaleAfNa      MapParamsLocale = "af-NA"
	MapParamsLocaleAfZa      MapParamsLocale = "af-ZA"
	MapParamsLocaleAk        MapParamsLocale = "ak"
	MapParamsLocaleAkGh      MapParamsLocale = "ak-GH"
	MapParamsLocaleAm        MapParamsLocale = "am"
	MapParamsLocaleAmEt      MapParamsLocale = "am-ET"
	MapParamsLocaleAnEs      MapParamsLocale = "an-ES"
	MapParamsLocaleAr        MapParamsLocale = "ar"
	MapParamsLocaleArAe      MapParamsLocale = "ar-AE"
	MapParamsLocaleArBh      MapParamsLocale = "ar-BH"
	MapParamsLocaleArDz      MapParamsLocale = "ar-DZ"
	MapParamsLocaleArEg      MapParamsLocale = "ar-EG"
	MapParamsLocaleArIn      MapParamsLocale = "ar-IN"
	MapParamsLocaleArIq      MapParamsLocale = "ar-IQ"
	MapParamsLocaleArJo      MapParamsLocale = "ar-JO"
	MapParamsLocaleArKw      MapParamsLocale = "ar-KW"
	MapParamsLocaleArLb      MapParamsLocale = "ar-LB"
	MapParamsLocaleArLy      MapParamsLocale = "ar-LY"
	MapParamsLocaleArMa      MapParamsLocale = "ar-MA"
	MapParamsLocaleArOm      MapParamsLocale = "ar-OM"
	MapParamsLocaleArQa      MapParamsLocale = "ar-QA"
	MapParamsLocaleArSa      MapParamsLocale = "ar-SA"
	MapParamsLocaleArSd      MapParamsLocale = "ar-SD"
	MapParamsLocaleArSy      MapParamsLocale = "ar-SY"
	MapParamsLocaleArTn      MapParamsLocale = "ar-TN"
	MapParamsLocaleArYe      MapParamsLocale = "ar-YE"
	MapParamsLocaleAs        MapParamsLocale = "as"
	MapParamsLocaleAsIn      MapParamsLocale = "as-IN"
	MapParamsLocaleAsa       MapParamsLocale = "asa"
	MapParamsLocaleAsaTz     MapParamsLocale = "asa-TZ"
	MapParamsLocaleAstEs     MapParamsLocale = "ast-ES"
	MapParamsLocaleAz        MapParamsLocale = "az"
	MapParamsLocaleAzAz      MapParamsLocale = "az-AZ"
	MapParamsLocaleAzCyrl    MapParamsLocale = "az-Cyrl"
	MapParamsLocaleAzCyrlAz  MapParamsLocale = "az-Cyrl-AZ"
	MapParamsLocaleAzLatn    MapParamsLocale = "az-Latn"
	MapParamsLocaleAzLatnAz  MapParamsLocale = "az-Latn-AZ"
	MapParamsLocaleBe        MapParamsLocale = "be"
	MapParamsLocaleBeBy      MapParamsLocale = "be-BY"
	MapParamsLocaleBem       MapParamsLocale = "bem"
	MapParamsLocaleBemZm     MapParamsLocale = "bem-ZM"
	MapParamsLocaleBerDz     MapParamsLocale = "ber-DZ"
	MapParamsLocaleBerMa     MapParamsLocale = "ber-MA"
	MapParamsLocaleBez       MapParamsLocale = "bez"
	MapParamsLocaleBezTz     MapParamsLocale = "bez-TZ"
	MapParamsLocaleBg        MapParamsLocale = "bg"
	MapParamsLocaleBgBg      MapParamsLocale = "bg-BG"
	MapParamsLocaleBhoIn     MapParamsLocale = "bho-IN"
	MapParamsLocaleBm        MapParamsLocale = "bm"
	MapParamsLocaleBmMl      MapParamsLocale = "bm-ML"
	MapParamsLocaleBn        MapParamsLocale = "bn"
	MapParamsLocaleBnBd      MapParamsLocale = "bn-BD"
	MapParamsLocaleBnIn      MapParamsLocale = "bn-IN"
	MapParamsLocaleBo        MapParamsLocale = "bo"
	MapParamsLocaleBoCn      MapParamsLocale = "bo-CN"
	MapParamsLocaleBoIn      MapParamsLocale = "bo-IN"
	MapParamsLocaleBrFr      MapParamsLocale = "br-FR"
	MapParamsLocaleBrxIn     MapParamsLocale = "brx-IN"
	MapParamsLocaleBs        MapParamsLocale = "bs"
	MapParamsLocaleBsBa      MapParamsLocale = "bs-BA"
	MapParamsLocaleBynEr     MapParamsLocale = "byn-ER"
	MapParamsLocaleCa        MapParamsLocale = "ca"
	MapParamsLocaleCaAd      MapParamsLocale = "ca-AD"
	MapParamsLocaleCaEs      MapParamsLocale = "ca-ES"
	MapParamsLocaleCaFr      MapParamsLocale = "ca-FR"
	MapParamsLocaleCaIt      MapParamsLocale = "ca-IT"
	MapParamsLocaleCgg       MapParamsLocale = "cgg"
	MapParamsLocaleCggUg     MapParamsLocale = "cgg-UG"
	MapParamsLocaleChr       MapParamsLocale = "chr"
	MapParamsLocaleChrUs     MapParamsLocale = "chr-US"
	MapParamsLocaleCrhUa     MapParamsLocale = "crh-UA"
	MapParamsLocaleCs        MapParamsLocale = "cs"
	MapParamsLocaleCsCz      MapParamsLocale = "cs-CZ"
	MapParamsLocaleCsbPl     MapParamsLocale = "csb-PL"
	MapParamsLocaleCvRu      MapParamsLocale = "cv-RU"
	MapParamsLocaleCy        MapParamsLocale = "cy"
	MapParamsLocaleCyGB      MapParamsLocale = "cy-GB"
	MapParamsLocaleDa        MapParamsLocale = "da"
	MapParamsLocaleDaDk      MapParamsLocale = "da-DK"
	MapParamsLocaleDav       MapParamsLocale = "dav"
	MapParamsLocaleDavKe     MapParamsLocale = "dav-KE"
	MapParamsLocaleDe        MapParamsLocale = "de"
	MapParamsLocaleDeAt      MapParamsLocale = "de-AT"
	MapParamsLocaleDeBe      MapParamsLocale = "de-BE"
	MapParamsLocaleDeCh      MapParamsLocale = "de-CH"
	MapParamsLocaleDeDe      MapParamsLocale = "de-DE"
	MapParamsLocaleDeLi      MapParamsLocale = "de-LI"
	MapParamsLocaleDeLu      MapParamsLocale = "de-LU"
	MapParamsLocaleDvMv      MapParamsLocale = "dv-MV"
	MapParamsLocaleDzBt      MapParamsLocale = "dz-BT"
	MapParamsLocaleEbu       MapParamsLocale = "ebu"
	MapParamsLocaleEbuKe     MapParamsLocale = "ebu-KE"
	MapParamsLocaleEe        MapParamsLocale = "ee"
	MapParamsLocaleEeGh      MapParamsLocale = "ee-GH"
	MapParamsLocaleEeTg      MapParamsLocale = "ee-TG"
	MapParamsLocaleEl        MapParamsLocale = "el"
	MapParamsLocaleElCy      MapParamsLocale = "el-CY"
	MapParamsLocaleElGr      MapParamsLocale = "el-GR"
	MapParamsLocaleEn        MapParamsLocale = "en"
	MapParamsLocaleEnAg      MapParamsLocale = "en-AG"
	MapParamsLocaleEnAs      MapParamsLocale = "en-AS"
	MapParamsLocaleEnAu      MapParamsLocale = "en-AU"
	MapParamsLocaleEnBe      MapParamsLocale = "en-BE"
	MapParamsLocaleEnBw      MapParamsLocale = "en-BW"
	MapParamsLocaleEnBz      MapParamsLocale = "en-BZ"
	MapParamsLocaleEnCa      MapParamsLocale = "en-CA"
	MapParamsLocaleEnDk      MapParamsLocale = "en-DK"
	MapParamsLocaleEnGB      MapParamsLocale = "en-GB"
	MapParamsLocaleEnGu      MapParamsLocale = "en-GU"
	MapParamsLocaleEnHk      MapParamsLocale = "en-HK"
	MapParamsLocaleEnIe      MapParamsLocale = "en-IE"
	MapParamsLocaleEnIn      MapParamsLocale = "en-IN"
	MapParamsLocaleEnJm      MapParamsLocale = "en-JM"
	MapParamsLocaleEnMh      MapParamsLocale = "en-MH"
	MapParamsLocaleEnMp      MapParamsLocale = "en-MP"
	MapParamsLocaleEnMt      MapParamsLocale = "en-MT"
	MapParamsLocaleEnMu      MapParamsLocale = "en-MU"
	MapParamsLocaleEnNa      MapParamsLocale = "en-NA"
	MapParamsLocaleEnNg      MapParamsLocale = "en-NG"
	MapParamsLocaleEnNz      MapParamsLocale = "en-NZ"
	MapParamsLocaleEnPh      MapParamsLocale = "en-PH"
	MapParamsLocaleEnPk      MapParamsLocale = "en-PK"
	MapParamsLocaleEnSg      MapParamsLocale = "en-SG"
	MapParamsLocaleEnTt      MapParamsLocale = "en-TT"
	MapParamsLocaleEnUm      MapParamsLocale = "en-UM"
	MapParamsLocaleEnUs      MapParamsLocale = "en-US"
	MapParamsLocaleEnVi      MapParamsLocale = "en-VI"
	MapParamsLocaleEnZa      MapParamsLocale = "en-ZA"
	MapParamsLocaleEnZm      MapParamsLocale = "en-ZM"
	MapParamsLocaleEnZw      MapParamsLocale = "en-ZW"
	MapParamsLocaleEo        MapParamsLocale = "eo"
	MapParamsLocaleEs        MapParamsLocale = "es"
	MapParamsLocaleEs419     MapParamsLocale = "es-419"
	MapParamsLocaleEsAr      MapParamsLocale = "es-AR"
	MapParamsLocaleEsBo      MapParamsLocale = "es-BO"
	MapParamsLocaleEsCl      MapParamsLocale = "es-CL"
	MapParamsLocaleEsCo      MapParamsLocale = "es-CO"
	MapParamsLocaleEsCr      MapParamsLocale = "es-CR"
	MapParamsLocaleEsCu      MapParamsLocale = "es-CU"
	MapParamsLocaleEsDo      MapParamsLocale = "es-DO"
	MapParamsLocaleEsEc      MapParamsLocale = "es-EC"
	MapParamsLocaleEsEs      MapParamsLocale = "es-ES"
	MapParamsLocaleEsGq      MapParamsLocale = "es-GQ"
	MapParamsLocaleEsGt      MapParamsLocale = "es-GT"
	MapParamsLocaleEsHn      MapParamsLocale = "es-HN"
	MapParamsLocaleEsMx      MapParamsLocale = "es-MX"
	MapParamsLocaleEsNi      MapParamsLocale = "es-NI"
	MapParamsLocaleEsPa      MapParamsLocale = "es-PA"
	MapParamsLocaleEsPe      MapParamsLocale = "es-PE"
	MapParamsLocaleEsPr      MapParamsLocale = "es-PR"
	MapParamsLocaleEsPy      MapParamsLocale = "es-PY"
	MapParamsLocaleEsSv      MapParamsLocale = "es-SV"
	MapParamsLocaleEsUs      MapParamsLocale = "es-US"
	MapParamsLocaleEsUy      MapParamsLocale = "es-UY"
	MapParamsLocaleEsVe      MapParamsLocale = "es-VE"
	MapParamsLocaleEt        MapParamsLocale = "et"
	MapParamsLocaleEtEe      MapParamsLocale = "et-EE"
	MapParamsLocaleEu        MapParamsLocale = "eu"
	MapParamsLocaleEuEs      MapParamsLocale = "eu-ES"
	MapParamsLocaleFa        MapParamsLocale = "fa"
	MapParamsLocaleFaAf      MapParamsLocale = "fa-AF"
	MapParamsLocaleFaIr      MapParamsLocale = "fa-IR"
	MapParamsLocaleFf        MapParamsLocale = "ff"
	MapParamsLocaleFfSn      MapParamsLocale = "ff-SN"
	MapParamsLocaleFi        MapParamsLocale = "fi"
	MapParamsLocaleFiFi      MapParamsLocale = "fi-FI"
	MapParamsLocaleFil       MapParamsLocale = "fil"
	MapParamsLocaleFilPh     MapParamsLocale = "fil-PH"
	MapParamsLocaleFo        MapParamsLocale = "fo"
	MapParamsLocaleFoFo      MapParamsLocale = "fo-FO"
	MapParamsLocaleFr        MapParamsLocale = "fr"
	MapParamsLocaleFrBe      MapParamsLocale = "fr-BE"
	MapParamsLocaleFrBf      MapParamsLocale = "fr-BF"
	MapParamsLocaleFrBi      MapParamsLocale = "fr-BI"
	MapParamsLocaleFrBj      MapParamsLocale = "fr-BJ"
	MapParamsLocaleFrBl      MapParamsLocale = "fr-BL"
	MapParamsLocaleFrCa      MapParamsLocale = "fr-CA"
	MapParamsLocaleFrCd      MapParamsLocale = "fr-CD"
	MapParamsLocaleFrCf      MapParamsLocale = "fr-CF"
	MapParamsLocaleFrCg      MapParamsLocale = "fr-CG"
	MapParamsLocaleFrCh      MapParamsLocale = "fr-CH"
	MapParamsLocaleFrCi      MapParamsLocale = "fr-CI"
	MapParamsLocaleFrCm      MapParamsLocale = "fr-CM"
	MapParamsLocaleFrDj      MapParamsLocale = "fr-DJ"
	MapParamsLocaleFrFr      MapParamsLocale = "fr-FR"
	MapParamsLocaleFrGa      MapParamsLocale = "fr-GA"
	MapParamsLocaleFrGn      MapParamsLocale = "fr-GN"
	MapParamsLocaleFrGp      MapParamsLocale = "fr-GP"
	MapParamsLocaleFrGq      MapParamsLocale = "fr-GQ"
	MapParamsLocaleFrKm      MapParamsLocale = "fr-KM"
	MapParamsLocaleFrLu      MapParamsLocale = "fr-LU"
	MapParamsLocaleFrMc      MapParamsLocale = "fr-MC"
	MapParamsLocaleFrMf      MapParamsLocale = "fr-MF"
	MapParamsLocaleFrMg      MapParamsLocale = "fr-MG"
	MapParamsLocaleFrMl      MapParamsLocale = "fr-ML"
	MapParamsLocaleFrMq      MapParamsLocale = "fr-MQ"
	MapParamsLocaleFrNe      MapParamsLocale = "fr-NE"
	MapParamsLocaleFrRe      MapParamsLocale = "fr-RE"
	MapParamsLocaleFrRw      MapParamsLocale = "fr-RW"
	MapParamsLocaleFrSn      MapParamsLocale = "fr-SN"
	MapParamsLocaleFrTd      MapParamsLocale = "fr-TD"
	MapParamsLocaleFrTg      MapParamsLocale = "fr-TG"
	MapParamsLocaleFurIt     MapParamsLocale = "fur-IT"
	MapParamsLocaleFyDe      MapParamsLocale = "fy-DE"
	MapParamsLocaleFyNl      MapParamsLocale = "fy-NL"
	MapParamsLocaleGa        MapParamsLocale = "ga"
	MapParamsLocaleGaIe      MapParamsLocale = "ga-IE"
	MapParamsLocaleGdGB      MapParamsLocale = "gd-GB"
	MapParamsLocaleGezEr     MapParamsLocale = "gez-ER"
	MapParamsLocaleGezEt     MapParamsLocale = "gez-ET"
	MapParamsLocaleGl        MapParamsLocale = "gl"
	MapParamsLocaleGlEs      MapParamsLocale = "gl-ES"
	MapParamsLocaleGsw       MapParamsLocale = "gsw"
	MapParamsLocaleGswCh     MapParamsLocale = "gsw-CH"
	MapParamsLocaleGu        MapParamsLocale = "gu"
	MapParamsLocaleGuIn      MapParamsLocale = "gu-IN"
	MapParamsLocaleGuz       MapParamsLocale = "guz"
	MapParamsLocaleGuzKe     MapParamsLocale = "guz-KE"
	MapParamsLocaleGv        MapParamsLocale = "gv"
	MapParamsLocaleGvGB      MapParamsLocale = "gv-GB"
	MapParamsLocaleHa        MapParamsLocale = "ha"
	MapParamsLocaleHaLatn    MapParamsLocale = "ha-Latn"
	MapParamsLocaleHaLatnGh  MapParamsLocale = "ha-Latn-GH"
	MapParamsLocaleHaLatnNe  MapParamsLocale = "ha-Latn-NE"
	MapParamsLocaleHaLatnNg  MapParamsLocale = "ha-Latn-NG"
	MapParamsLocaleHaNg      MapParamsLocale = "ha-NG"
	MapParamsLocaleHaw       MapParamsLocale = "haw"
	MapParamsLocaleHawUs     MapParamsLocale = "haw-US"
	MapParamsLocaleHe        MapParamsLocale = "he"
	MapParamsLocaleHeIl      MapParamsLocale = "he-IL"
	MapParamsLocaleHi        MapParamsLocale = "hi"
	MapParamsLocaleHiIn      MapParamsLocale = "hi-IN"
	MapParamsLocaleHneIn     MapParamsLocale = "hne-IN"
	MapParamsLocaleHr        MapParamsLocale = "hr"
	MapParamsLocaleHrHr      MapParamsLocale = "hr-HR"
	MapParamsLocaleHsbDe     MapParamsLocale = "hsb-DE"
	MapParamsLocaleHtHt      MapParamsLocale = "ht-HT"
	MapParamsLocaleHu        MapParamsLocale = "hu"
	MapParamsLocaleHuHu      MapParamsLocale = "hu-HU"
	MapParamsLocaleHy        MapParamsLocale = "hy"
	MapParamsLocaleHyAm      MapParamsLocale = "hy-AM"
	MapParamsLocaleID        MapParamsLocale = "id"
	MapParamsLocaleIDID      MapParamsLocale = "id-ID"
	MapParamsLocaleIg        MapParamsLocale = "ig"
	MapParamsLocaleIgNg      MapParamsLocale = "ig-NG"
	MapParamsLocaleIi        MapParamsLocale = "ii"
	MapParamsLocaleIiCn      MapParamsLocale = "ii-CN"
	MapParamsLocaleIkCa      MapParamsLocale = "ik-CA"
	MapParamsLocaleIs        MapParamsLocale = "is"
	MapParamsLocaleIsIs      MapParamsLocale = "is-IS"
	MapParamsLocaleIt        MapParamsLocale = "it"
	MapParamsLocaleItCh      MapParamsLocale = "it-CH"
	MapParamsLocaleItIt      MapParamsLocale = "it-IT"
	MapParamsLocaleIuCa      MapParamsLocale = "iu-CA"
	MapParamsLocaleIwIl      MapParamsLocale = "iw-IL"
	MapParamsLocaleJa        MapParamsLocale = "ja"
	MapParamsLocaleJaJp      MapParamsLocale = "ja-JP"
	MapParamsLocaleJmc       MapParamsLocale = "jmc"
	MapParamsLocaleJmcTz     MapParamsLocale = "jmc-TZ"
	MapParamsLocaleKa        MapParamsLocale = "ka"
	MapParamsLocaleKaGe      MapParamsLocale = "ka-GE"
	MapParamsLocaleKab       MapParamsLocale = "kab"
	MapParamsLocaleKabDz     MapParamsLocale = "kab-DZ"
	MapParamsLocaleKam       MapParamsLocale = "kam"
	MapParamsLocaleKamKe     MapParamsLocale = "kam-KE"
	MapParamsLocaleKde       MapParamsLocale = "kde"
	MapParamsLocaleKdeTz     MapParamsLocale = "kde-TZ"
	MapParamsLocaleKea       MapParamsLocale = "kea"
	MapParamsLocaleKeaCv     MapParamsLocale = "kea-CV"
	MapParamsLocaleKhq       MapParamsLocale = "khq"
	MapParamsLocaleKhqMl     MapParamsLocale = "khq-ML"
	MapParamsLocaleKi        MapParamsLocale = "ki"
	MapParamsLocaleKiKe      MapParamsLocale = "ki-KE"
	MapParamsLocaleKk        MapParamsLocale = "kk"
	MapParamsLocaleKkCyrl    MapParamsLocale = "kk-Cyrl"
	MapParamsLocaleKkCyrlKz  MapParamsLocale = "kk-Cyrl-KZ"
	MapParamsLocaleKkKz      MapParamsLocale = "kk-KZ"
	MapParamsLocaleKl        MapParamsLocale = "kl"
	MapParamsLocaleKlGl      MapParamsLocale = "kl-GL"
	MapParamsLocaleKln       MapParamsLocale = "kln"
	MapParamsLocaleKlnKe     MapParamsLocale = "kln-KE"
	MapParamsLocaleKm        MapParamsLocale = "km"
	MapParamsLocaleKmKh      MapParamsLocale = "km-KH"
	MapParamsLocaleKn        MapParamsLocale = "kn"
	MapParamsLocaleKnIn      MapParamsLocale = "kn-IN"
	MapParamsLocaleKo        MapParamsLocale = "ko"
	MapParamsLocaleKoKr      MapParamsLocale = "ko-KR"
	MapParamsLocaleKok       MapParamsLocale = "kok"
	MapParamsLocaleKokIn     MapParamsLocale = "kok-IN"
	MapParamsLocaleKsIn      MapParamsLocale = "ks-IN"
	MapParamsLocaleKuTr      MapParamsLocale = "ku-TR"
	MapParamsLocaleKw        MapParamsLocale = "kw"
	MapParamsLocaleKwGB      MapParamsLocale = "kw-GB"
	MapParamsLocaleKyKg      MapParamsLocale = "ky-KG"
	MapParamsLocaleLag       MapParamsLocale = "lag"
	MapParamsLocaleLagTz     MapParamsLocale = "lag-TZ"
	MapParamsLocaleLbLu      MapParamsLocale = "lb-LU"
	MapParamsLocaleLg        MapParamsLocale = "lg"
	MapParamsLocaleLgUg      MapParamsLocale = "lg-UG"
	MapParamsLocaleLiBe      MapParamsLocale = "li-BE"
	MapParamsLocaleLiNl      MapParamsLocale = "li-NL"
	MapParamsLocaleLijIt     MapParamsLocale = "lij-IT"
	MapParamsLocaleLoLa      MapParamsLocale = "lo-LA"
	MapParamsLocaleLt        MapParamsLocale = "lt"
	MapParamsLocaleLtLt      MapParamsLocale = "lt-LT"
	MapParamsLocaleLuo       MapParamsLocale = "luo"
	MapParamsLocaleLuoKe     MapParamsLocale = "luo-KE"
	MapParamsLocaleLuy       MapParamsLocale = "luy"
	MapParamsLocaleLuyKe     MapParamsLocale = "luy-KE"
	MapParamsLocaleLv        MapParamsLocale = "lv"
	MapParamsLocaleLvLv      MapParamsLocale = "lv-LV"
	MapParamsLocaleMagIn     MapParamsLocale = "mag-IN"
	MapParamsLocaleMaiIn     MapParamsLocale = "mai-IN"
	MapParamsLocaleMas       MapParamsLocale = "mas"
	MapParamsLocaleMasKe     MapParamsLocale = "mas-KE"
	MapParamsLocaleMasTz     MapParamsLocale = "mas-TZ"
	MapParamsLocaleMer       MapParamsLocale = "mer"
	MapParamsLocaleMerKe     MapParamsLocale = "mer-KE"
	MapParamsLocaleMfe       MapParamsLocale = "mfe"
	MapParamsLocaleMfeMu     MapParamsLocale = "mfe-MU"
	MapParamsLocaleMg        MapParamsLocale = "mg"
	MapParamsLocaleMgMg      MapParamsLocale = "mg-MG"
	MapParamsLocaleMhrRu     MapParamsLocale = "mhr-RU"
	MapParamsLocaleMiNz      MapParamsLocale = "mi-NZ"
	MapParamsLocaleMk        MapParamsLocale = "mk"
	MapParamsLocaleMkMk      MapParamsLocale = "mk-MK"
	MapParamsLocaleMl        MapParamsLocale = "ml"
	MapParamsLocaleMlIn      MapParamsLocale = "ml-IN"
	MapParamsLocaleMnMn      MapParamsLocale = "mn-MN"
	MapParamsLocaleMr        MapParamsLocale = "mr"
	MapParamsLocaleMrIn      MapParamsLocale = "mr-IN"
	MapParamsLocaleMs        MapParamsLocale = "ms"
	MapParamsLocaleMsBn      MapParamsLocale = "ms-BN"
	MapParamsLocaleMsMy      MapParamsLocale = "ms-MY"
	MapParamsLocaleMt        MapParamsLocale = "mt"
	MapParamsLocaleMtMt      MapParamsLocale = "mt-MT"
	MapParamsLocaleMy        MapParamsLocale = "my"
	MapParamsLocaleMyMm      MapParamsLocale = "my-MM"
	MapParamsLocaleNanTw     MapParamsLocale = "nan-TW"
	MapParamsLocaleNaq       MapParamsLocale = "naq"
	MapParamsLocaleNaqNa     MapParamsLocale = "naq-NA"
	MapParamsLocaleNb        MapParamsLocale = "nb"
	MapParamsLocaleNbNo      MapParamsLocale = "nb-NO"
	MapParamsLocaleNd        MapParamsLocale = "nd"
	MapParamsLocaleNdZw      MapParamsLocale = "nd-ZW"
	MapParamsLocaleNdsDe     MapParamsLocale = "nds-DE"
	MapParamsLocaleNdsNl     MapParamsLocale = "nds-NL"
	MapParamsLocaleNe        MapParamsLocale = "ne"
	MapParamsLocaleNeIn      MapParamsLocale = "ne-IN"
	MapParamsLocaleNeNp      MapParamsLocale = "ne-NP"
	MapParamsLocaleNl        MapParamsLocale = "nl"
	MapParamsLocaleNlAw      MapParamsLocale = "nl-AW"
	MapParamsLocaleNlBe      MapParamsLocale = "nl-BE"
	MapParamsLocaleNlNl      MapParamsLocale = "nl-NL"
	MapParamsLocaleNn        MapParamsLocale = "nn"
	MapParamsLocaleNnNo      MapParamsLocale = "nn-NO"
	MapParamsLocaleNrZa      MapParamsLocale = "nr-ZA"
	MapParamsLocaleNsoZa     MapParamsLocale = "nso-ZA"
	MapParamsLocaleNyn       MapParamsLocale = "nyn"
	MapParamsLocaleNynUg     MapParamsLocale = "nyn-UG"
	MapParamsLocaleOcFr      MapParamsLocale = "oc-FR"
	MapParamsLocaleOm        MapParamsLocale = "om"
	MapParamsLocaleOmEt      MapParamsLocale = "om-ET"
	MapParamsLocaleOmKe      MapParamsLocale = "om-KE"
	MapParamsLocaleOr        MapParamsLocale = "or"
	MapParamsLocaleOrIn      MapParamsLocale = "or-IN"
	MapParamsLocaleOsRu      MapParamsLocale = "os-RU"
	MapParamsLocalePa        MapParamsLocale = "pa"
	MapParamsLocalePaArab    MapParamsLocale = "pa-Arab"
	MapParamsLocalePaArabPk  MapParamsLocale = "pa-Arab-PK"
	MapParamsLocalePaGuru    MapParamsLocale = "pa-Guru"
	MapParamsLocalePaGuruIn  MapParamsLocale = "pa-Guru-IN"
	MapParamsLocalePaIn      MapParamsLocale = "pa-IN"
	MapParamsLocalePaPk      MapParamsLocale = "pa-PK"
	MapParamsLocalePapAn     MapParamsLocale = "pap-AN"
	MapParamsLocalePl        MapParamsLocale = "pl"
	MapParamsLocalePlPl      MapParamsLocale = "pl-PL"
	MapParamsLocalePs        MapParamsLocale = "ps"
	MapParamsLocalePsAf      MapParamsLocale = "ps-AF"
	MapParamsLocalePt        MapParamsLocale = "pt"
	MapParamsLocalePtBr      MapParamsLocale = "pt-BR"
	MapParamsLocalePtGw      MapParamsLocale = "pt-GW"
	MapParamsLocalePtMz      MapParamsLocale = "pt-MZ"
	MapParamsLocalePtPt      MapParamsLocale = "pt-PT"
	MapParamsLocaleRm        MapParamsLocale = "rm"
	MapParamsLocaleRmCh      MapParamsLocale = "rm-CH"
	MapParamsLocaleRo        MapParamsLocale = "ro"
	MapParamsLocaleRoMd      MapParamsLocale = "ro-MD"
	MapParamsLocaleRoRo      MapParamsLocale = "ro-RO"
	MapParamsLocaleRof       MapParamsLocale = "rof"
	MapParamsLocaleRofTz     MapParamsLocale = "rof-TZ"
	MapParamsLocaleRu        MapParamsLocale = "ru"
	MapParamsLocaleRuMd      MapParamsLocale = "ru-MD"
	MapParamsLocaleRuRu      MapParamsLocale = "ru-RU"
	MapParamsLocaleRuUa      MapParamsLocale = "ru-UA"
	MapParamsLocaleRw        MapParamsLocale = "rw"
	MapParamsLocaleRwRw      MapParamsLocale = "rw-RW"
	MapParamsLocaleRwk       MapParamsLocale = "rwk"
	MapParamsLocaleRwkTz     MapParamsLocale = "rwk-TZ"
	MapParamsLocaleSaIn      MapParamsLocale = "sa-IN"
	MapParamsLocaleSaq       MapParamsLocale = "saq"
	MapParamsLocaleSaqKe     MapParamsLocale = "saq-KE"
	MapParamsLocaleScIt      MapParamsLocale = "sc-IT"
	MapParamsLocaleSdIn      MapParamsLocale = "sd-IN"
	MapParamsLocaleSeNo      MapParamsLocale = "se-NO"
	MapParamsLocaleSeh       MapParamsLocale = "seh"
	MapParamsLocaleSehMz     MapParamsLocale = "seh-MZ"
	MapParamsLocaleSes       MapParamsLocale = "ses"
	MapParamsLocaleSesMl     MapParamsLocale = "ses-ML"
	MapParamsLocaleSg        MapParamsLocale = "sg"
	MapParamsLocaleSgCf      MapParamsLocale = "sg-CF"
	MapParamsLocaleShi       MapParamsLocale = "shi"
	MapParamsLocaleShiLatn   MapParamsLocale = "shi-Latn"
	MapParamsLocaleShiLatnMa MapParamsLocale = "shi-Latn-MA"
	MapParamsLocaleShiTfng   MapParamsLocale = "shi-Tfng"
	MapParamsLocaleShiTfngMa MapParamsLocale = "shi-Tfng-MA"
	MapParamsLocaleShsCa     MapParamsLocale = "shs-CA"
	MapParamsLocaleSi        MapParamsLocale = "si"
	MapParamsLocaleSiLk      MapParamsLocale = "si-LK"
	MapParamsLocaleSidEt     MapParamsLocale = "sid-ET"
	MapParamsLocaleSk        MapParamsLocale = "sk"
	MapParamsLocaleSkSk      MapParamsLocale = "sk-SK"
	MapParamsLocaleSl        MapParamsLocale = "sl"
	MapParamsLocaleSlSi      MapParamsLocale = "sl-SI"
	MapParamsLocaleSn        MapParamsLocale = "sn"
	MapParamsLocaleSnZw      MapParamsLocale = "sn-ZW"
	MapParamsLocaleSo        MapParamsLocale = "so"
	MapParamsLocaleSoDj      MapParamsLocale = "so-DJ"
	MapParamsLocaleSoEt      MapParamsLocale = "so-ET"
	MapParamsLocaleSoKe      MapParamsLocale = "so-KE"
	MapParamsLocaleSoSo      MapParamsLocale = "so-SO"
	MapParamsLocaleSq        MapParamsLocale = "sq"
	MapParamsLocaleSqAl      MapParamsLocale = "sq-AL"
	MapParamsLocaleSqMk      MapParamsLocale = "sq-MK"
	MapParamsLocaleSr        MapParamsLocale = "sr"
	MapParamsLocaleSrCyrl    MapParamsLocale = "sr-Cyrl"
	MapParamsLocaleSrCyrlBa  MapParamsLocale = "sr-Cyrl-BA"
	MapParamsLocaleSrCyrlMe  MapParamsLocale = "sr-Cyrl-ME"
	MapParamsLocaleSrCyrlRs  MapParamsLocale = "sr-Cyrl-RS"
	MapParamsLocaleSrLatn    MapParamsLocale = "sr-Latn"
	MapParamsLocaleSrLatnBa  MapParamsLocale = "sr-Latn-BA"
	MapParamsLocaleSrLatnMe  MapParamsLocale = "sr-Latn-ME"
	MapParamsLocaleSrLatnRs  MapParamsLocale = "sr-Latn-RS"
	MapParamsLocaleSrMe      MapParamsLocale = "sr-ME"
	MapParamsLocaleSrRs      MapParamsLocale = "sr-RS"
	MapParamsLocaleSSZa      MapParamsLocale = "ss-ZA"
	MapParamsLocaleStZa      MapParamsLocale = "st-ZA"
	MapParamsLocaleSv        MapParamsLocale = "sv"
	MapParamsLocaleSvFi      MapParamsLocale = "sv-FI"
	MapParamsLocaleSvSe      MapParamsLocale = "sv-SE"
	MapParamsLocaleSw        MapParamsLocale = "sw"
	MapParamsLocaleSwKe      MapParamsLocale = "sw-KE"
	MapParamsLocaleSwTz      MapParamsLocale = "sw-TZ"
	MapParamsLocaleTa        MapParamsLocale = "ta"
	MapParamsLocaleTaIn      MapParamsLocale = "ta-IN"
	MapParamsLocaleTaLk      MapParamsLocale = "ta-LK"
	MapParamsLocaleTe        MapParamsLocale = "te"
	MapParamsLocaleTeIn      MapParamsLocale = "te-IN"
	MapParamsLocaleTeo       MapParamsLocale = "teo"
	MapParamsLocaleTeoKe     MapParamsLocale = "teo-KE"
	MapParamsLocaleTeoUg     MapParamsLocale = "teo-UG"
	MapParamsLocaleTgTj      MapParamsLocale = "tg-TJ"
	MapParamsLocaleTh        MapParamsLocale = "th"
	MapParamsLocaleThTh      MapParamsLocale = "th-TH"
	MapParamsLocaleTi        MapParamsLocale = "ti"
	MapParamsLocaleTiEr      MapParamsLocale = "ti-ER"
	MapParamsLocaleTiEt      MapParamsLocale = "ti-ET"
	MapParamsLocaleTigEr     MapParamsLocale = "tig-ER"
	MapParamsLocaleTkTm      MapParamsLocale = "tk-TM"
	MapParamsLocaleTlPh      MapParamsLocale = "tl-PH"
	MapParamsLocaleTnZa      MapParamsLocale = "tn-ZA"
	MapParamsLocaleTo        MapParamsLocale = "to"
	MapParamsLocaleToTo      MapParamsLocale = "to-TO"
	MapParamsLocaleTr        MapParamsLocale = "tr"
	MapParamsLocaleTrCy      MapParamsLocale = "tr-CY"
	MapParamsLocaleTrTr      MapParamsLocale = "tr-TR"
	MapParamsLocaleTsZa      MapParamsLocale = "ts-ZA"
	MapParamsLocaleTtRu      MapParamsLocale = "tt-RU"
	MapParamsLocaleTzm       MapParamsLocale = "tzm"
	MapParamsLocaleTzmLatn   MapParamsLocale = "tzm-Latn"
	MapParamsLocaleTzmLatnMa MapParamsLocale = "tzm-Latn-MA"
	MapParamsLocaleUgCn      MapParamsLocale = "ug-CN"
	MapParamsLocaleUk        MapParamsLocale = "uk"
	MapParamsLocaleUkUa      MapParamsLocale = "uk-UA"
	MapParamsLocaleUnmUs     MapParamsLocale = "unm-US"
	MapParamsLocaleUr        MapParamsLocale = "ur"
	MapParamsLocaleUrIn      MapParamsLocale = "ur-IN"
	MapParamsLocaleUrPk      MapParamsLocale = "ur-PK"
	MapParamsLocaleUz        MapParamsLocale = "uz"
	MapParamsLocaleUzArab    MapParamsLocale = "uz-Arab"
	MapParamsLocaleUzArabAf  MapParamsLocale = "uz-Arab-AF"
	MapParamsLocaleUzCyrl    MapParamsLocale = "uz-Cyrl"
	MapParamsLocaleUzCyrlUz  MapParamsLocale = "uz-Cyrl-UZ"
	MapParamsLocaleUzLatn    MapParamsLocale = "uz-Latn"
	MapParamsLocaleUzLatnUz  MapParamsLocale = "uz-Latn-UZ"
	MapParamsLocaleUzUz      MapParamsLocale = "uz-UZ"
	MapParamsLocaleVeZa      MapParamsLocale = "ve-ZA"
	MapParamsLocaleVi        MapParamsLocale = "vi"
	MapParamsLocaleViVn      MapParamsLocale = "vi-VN"
	MapParamsLocaleVun       MapParamsLocale = "vun"
	MapParamsLocaleVunTz     MapParamsLocale = "vun-TZ"
	MapParamsLocaleWaBe      MapParamsLocale = "wa-BE"
	MapParamsLocaleWaeCh     MapParamsLocale = "wae-CH"
	MapParamsLocaleWalEt     MapParamsLocale = "wal-ET"
	MapParamsLocaleWoSn      MapParamsLocale = "wo-SN"
	MapParamsLocaleXhZa      MapParamsLocale = "xh-ZA"
	MapParamsLocaleXog       MapParamsLocale = "xog"
	MapParamsLocaleXogUg     MapParamsLocale = "xog-UG"
	MapParamsLocaleYiUs      MapParamsLocale = "yi-US"
	MapParamsLocaleYo        MapParamsLocale = "yo"
	MapParamsLocaleYoNg      MapParamsLocale = "yo-NG"
	MapParamsLocaleYueHk     MapParamsLocale = "yue-HK"
	MapParamsLocaleZh        MapParamsLocale = "zh"
	MapParamsLocaleZhCn      MapParamsLocale = "zh-CN"
	MapParamsLocaleZhHk      MapParamsLocale = "zh-HK"
	MapParamsLocaleZhHans    MapParamsLocale = "zh-Hans"
	MapParamsLocaleZhHansCn  MapParamsLocale = "zh-Hans-CN"
	MapParamsLocaleZhHansHk  MapParamsLocale = "zh-Hans-HK"
	MapParamsLocaleZhHansMo  MapParamsLocale = "zh-Hans-MO"
	MapParamsLocaleZhHansSg  MapParamsLocale = "zh-Hans-SG"
	MapParamsLocaleZhHant    MapParamsLocale = "zh-Hant"
	MapParamsLocaleZhHantHk  MapParamsLocale = "zh-Hant-HK"
	MapParamsLocaleZhHantMo  MapParamsLocale = "zh-Hant-MO"
	MapParamsLocaleZhHantTw  MapParamsLocale = "zh-Hant-TW"
	MapParamsLocaleZhSg      MapParamsLocale = "zh-SG"
	MapParamsLocaleZhTw      MapParamsLocale = "zh-TW"
	MapParamsLocaleZu        MapParamsLocale = "zu"
	MapParamsLocaleZuZa      MapParamsLocale = "zu-ZA"
	MapParamsLocaleAuto      MapParamsLocale = "auto"
)

// Sitemap and other methods will be used together to find URLs.
type MapParamsSitemap string

const (
	MapParamsSitemapSkip    MapParamsSitemap = "skip"
	MapParamsSitemapInclude MapParamsSitemap = "include"
	MapParamsSitemapOnly    MapParamsSitemap = "only"
)

type SearchParams struct {
	// Search query string
	Query string `json:"query" api:"required"`
	// Deprecated. Use search_depth instead. true maps to 'deep', false maps to 'lite'.
	DeepSearch param.Opt[bool] `json:"deep_search,omitzero"`
	// Filter results before this date (format: YYYY-MM-DD or YYYY)
	EndDate param.Opt[string] `json:"end_date,omitzero"`
	// Filter results after this date (format: YYYY-MM-DD or YYYY)
	StartDate param.Opt[string] `json:"start_date,omitzero"`
	// Country code for geo-targeted results (e.g., 'US', 'GB', 'IL')
	Country param.Opt[string] `json:"country,omitzero"`
	// Generate an LLM-powered answer summary based on search result snippets.
	IncludeAnswer param.Opt[bool] `json:"include_answer,omitzero"`
	// Language/locale code (e.g., 'en', 'fr', 'de')
	Locale param.Opt[string] `json:"locale,omitzero"`
	// Maximum number of results to return. Actual count may be lower depending on
	// availability.
	MaxResults param.Opt[int64] `json:"max_results,omitzero"`
	// Maximum number of subagents to execute in parallel for WSA focus modes
	// (shopping, social, geo). Ignored for SERP focus modes.
	MaxSubagents param.Opt[int64] `json:"max_subagents,omitzero"`
	// Filter by content type (only supported with focus=general). Supports semantic
	// groups ('documents', 'spreadsheets', 'presentations') and specific formats
	// ('pdf', 'docx', 'xlsx', etc.)
	ContentType []string `json:"content_type,omitzero"`
	// List of domains to exclude from search results. Maximum 50 domains.
	ExcludeDomains []string `json:"exclude_domains,omitzero"`
	// List of domains to include in search results. Maximum 50 domains.
	IncludeDomains []string `json:"include_domains,omitzero"`
	// Controls content richness and latency of search results.
	//
	//   - lite: Token-efficient metadata for high-volume pipelines (title, URL,
	//     description only)
	//   - fast: Rich content (~2K chars) optimized for AI agents
	//   - deep: Full page content via Webit scraping for comprehensive analysis
	//
	// Any of "lite", "fast", "deep".
	SearchDepth SearchParamsSearchDepth `json:"search_depth,omitzero"`
	// Time range filters passed to Webit SERP API as 'time' parameter.
	//
	// Any of "hour", "day", "week", "month", "year".
	TimeRange SearchParamsTimeRange `json:"time_range,omitzero"`
	// Search focus mode (e.g., 'general', 'news', 'shopping') or a list of explicit
	// subagent names (e.g., ['amazon_serp', 'target_serp'])
	Focus SearchParamsFocusUnion `json:"focus,omitzero"`
	// Output format: plain_text, markdown, or simplified_html
	//
	// Any of "plain_text", "markdown", "simplified_html".
	OutputFormat SearchParamsOutputFormat `json:"output_format,omitzero"`
	paramObj
}

func (r SearchParams) MarshalJSON() (data []byte, err error) {
	type shadow SearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SearchParamsFocusUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u SearchParamsFocusUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *SearchParamsFocusUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SearchParamsFocusUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Output format: plain_text, markdown, or simplified_html
type SearchParamsOutputFormat string

const (
	SearchParamsOutputFormatPlainText      SearchParamsOutputFormat = "plain_text"
	SearchParamsOutputFormatMarkdown       SearchParamsOutputFormat = "markdown"
	SearchParamsOutputFormatSimplifiedHTML SearchParamsOutputFormat = "simplified_html"
)

// Controls content richness and latency of search results.
//
//   - lite: Token-efficient metadata for high-volume pipelines (title, URL,
//     description only)
//   - fast: Rich content (~2K chars) optimized for AI agents
//   - deep: Full page content via Webit scraping for comprehensive analysis
type SearchParamsSearchDepth string

const (
	SearchParamsSearchDepthLite SearchParamsSearchDepth = "lite"
	SearchParamsSearchDepthFast SearchParamsSearchDepth = "fast"
	SearchParamsSearchDepthDeep SearchParamsSearchDepth = "deep"
)

// Time range filters passed to Webit SERP API as 'time' parameter.
type SearchParamsTimeRange string

const (
	SearchParamsTimeRangeHour  SearchParamsTimeRange = "hour"
	SearchParamsTimeRangeDay   SearchParamsTimeRange = "day"
	SearchParamsTimeRangeWeek  SearchParamsTimeRange = "week"
	SearchParamsTimeRangeMonth SearchParamsTimeRange = "month"
	SearchParamsTimeRangeYear  SearchParamsTimeRange = "year"
)
