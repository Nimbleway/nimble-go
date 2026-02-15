// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomnimblewaynimblego

import (
	"encoding/json"

	"github.com/Nimbleway/nimble-go/internal/apijson"
	"github.com/Nimbleway/nimble-go/packages/param"
	"github.com/Nimbleway/nimble-go/packages/respjson"
	"github.com/Nimbleway/nimble-go/shared/constant"
)

type ExtractResponse struct {
	Data     ExtractResponseData     `json:"data,required"`
	Metadata ExtractResponseMetadata `json:"metadata,required"`
	// The status of the task.
	//
	// Any of "success", "skipped", "fatal", "error", "postponed", "ignored",
	// "rejected", "blocked".
	Status ExtractResponseStatus `json:"status,required"`
	// Unique identifier for the task.
	TaskID string `json:"task_id,required"`
	// The final URL.
	URL   string               `json:"url,required"`
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
	// The parsing results extracted from the HTML & network content.
	Parsing ExtractResponseDataParsingUnion `json:"parsing"`
	// The list of redirects that occurred during the task.
	Redirects []ExtractResponseDataRedirect `json:"redirects"`
	// The screenshots from browser actions taken during the task.
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
	Results       []ExtractResponseDataBrowserActionsResult `json:"results,required"`
	Success       bool                                      `json:"success,required"`
	TotalDuration float64                                   `json:"total_duration,required"`
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
	Duration float64 `json:"duration,required"`
	// Any of "goto", "wait", "wait_for_element", "wait_for_navigation", "click",
	// "fill", "press", "scroll", "auto_scroll", "screenshot", "get_cookies", "eval",
	// "fetch".
	Name string `json:"name,required"`
	// Any of "no-run", "in-progress", "done", "error", "skipped".
	Status string `json:"status,required"`
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
	Filter       ExtractResponseDataNetworkCaptureFilter   `json:"filter,required"`
	Results      []ExtractResponseDataNetworkCaptureResult `json:"results,required"`
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
	Validation           bool    `json:"validation,required"`
	WaitForRequestsCount float64 `json:"wait_for_requests_count,required"`
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
	Type  string `json:"type,required"`
	Value string `json:"value,required"`
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
	Request  ExtractResponseDataNetworkCaptureResultRequest  `json:"request,required"`
	Response ExtractResponseDataNetworkCaptureResultResponse `json:"response,required"`
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
	Headers map[string]string `json:"headers,required"`
	Method  string            `json:"method,required"`
	// Resource type for network capture filtering
	//
	// Any of "document", "stylesheet", "image", "media", "font", "script",
	// "texttrack", "xhr", "fetch", "prefetch", "eventsource", "websocket", "manifest",
	// "signedexchange", "ping", "cspviolationreport", "preflight", "other", "fedcm".
	ResourceType string `json:"resource_type,required"`
	URL          string `json:"url,required"`
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
	Body    string            `json:"body,required"`
	Headers map[string]string `json:"headers,required"`
	// Any of "none", "base64".
	Serialization string  `json:"serialization,required"`
	Status        float64 `json:"status,required"`
	StatusText    string  `json:"status_text,required"`
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
	Entities map[string]any   `json:"entities,required"`
	Status   constant.Success `json:"status,required"`
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
	Error  string         `json:"error,required"`
	Status constant.Error `json:"status,required"`
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
	StatusCode float64 `json:"status_code,required"`
	URL        string  `json:"url,required"`
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
	NextPageParams map[string]any `json:"next_page_params,required"`
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
	NextPageParams map[string]any `json:"next_page_params,required"`
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

// Response schema for map requests.
type MapResponse struct {
	// Array of mapped links with optional titles and descriptions.
	Links []MapResponseLink `json:"links,required"`
	// Indicates if the map request was successful.
	Success bool `json:"success,required"`
	// Unique identifier for the map task.
	TaskID string `json:"task_id,required"`
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
	URL         string `json:"url,required" format:"uri"`
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

type ExtractParams struct {
	// Target URL to scrape
	URL string `json:"url,required"`
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
	// Any of "html", "markdown".
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
	Name string `json:"name,omitzero,required"`
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
	OfExtractsBrowserActionAutoScrollAction        *ExtractParamsBrowserActionAutoScrollAction        `json:",omitzero,inline"`
	OfExtractsBrowserActionClickAction             *ExtractParamsBrowserActionClickAction             `json:",omitzero,inline"`
	OfExtractsBrowserActionEvalAction              *ExtractParamsBrowserActionEvalAction              `json:",omitzero,inline"`
	OfExtractsBrowserActionFetchAction             *ExtractParamsBrowserActionFetchAction             `json:",omitzero,inline"`
	OfExtractsBrowserActionFillAction              *ExtractParamsBrowserActionFillAction              `json:",omitzero,inline"`
	OfExtractsBrowserActionGetCookiesAction        *ExtractParamsBrowserActionGetCookiesAction        `json:",omitzero,inline"`
	OfExtractsBrowserActionGotoAction              *ExtractParamsBrowserActionGotoAction              `json:",omitzero,inline"`
	OfExtractsBrowserActionPressAction             *ExtractParamsBrowserActionPressAction             `json:",omitzero,inline"`
	OfExtractsBrowserActionScreenshotAction        *ExtractParamsBrowserActionScreenshotAction        `json:",omitzero,inline"`
	OfExtractsBrowserActionScrollAction            *ExtractParamsBrowserActionScrollAction            `json:",omitzero,inline"`
	OfExtractsBrowserActionWaitAction              *ExtractParamsBrowserActionWaitAction              `json:",omitzero,inline"`
	OfExtractsBrowserActionWaitForElementAction    *ExtractParamsBrowserActionWaitForElementAction    `json:",omitzero,inline"`
	OfExtractsBrowserActionWaitForNavigationAction *ExtractParamsBrowserActionWaitForNavigationAction `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionAutoScrollAction,
		u.OfExtractsBrowserActionClickAction,
		u.OfExtractsBrowserActionEvalAction,
		u.OfExtractsBrowserActionFetchAction,
		u.OfExtractsBrowserActionFillAction,
		u.OfExtractsBrowserActionGetCookiesAction,
		u.OfExtractsBrowserActionGotoAction,
		u.OfExtractsBrowserActionPressAction,
		u.OfExtractsBrowserActionScreenshotAction,
		u.OfExtractsBrowserActionScrollAction,
		u.OfExtractsBrowserActionWaitAction,
		u.OfExtractsBrowserActionWaitForElementAction,
		u.OfExtractsBrowserActionWaitForNavigationAction)
}
func (u *ExtractParamsBrowserActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionAutoScrollAction) {
		return u.OfExtractsBrowserActionAutoScrollAction
	} else if !param.IsOmitted(u.OfExtractsBrowserActionClickAction) {
		return u.OfExtractsBrowserActionClickAction
	} else if !param.IsOmitted(u.OfExtractsBrowserActionEvalAction) {
		return u.OfExtractsBrowserActionEvalAction
	} else if !param.IsOmitted(u.OfExtractsBrowserActionFetchAction) {
		return u.OfExtractsBrowserActionFetchAction
	} else if !param.IsOmitted(u.OfExtractsBrowserActionFillAction) {
		return u.OfExtractsBrowserActionFillAction
	} else if !param.IsOmitted(u.OfExtractsBrowserActionGetCookiesAction) {
		return u.OfExtractsBrowserActionGetCookiesAction
	} else if !param.IsOmitted(u.OfExtractsBrowserActionGotoAction) {
		return u.OfExtractsBrowserActionGotoAction
	} else if !param.IsOmitted(u.OfExtractsBrowserActionPressAction) {
		return u.OfExtractsBrowserActionPressAction
	} else if !param.IsOmitted(u.OfExtractsBrowserActionScreenshotAction) {
		return u.OfExtractsBrowserActionScreenshotAction
	} else if !param.IsOmitted(u.OfExtractsBrowserActionScrollAction) {
		return u.OfExtractsBrowserActionScrollAction
	} else if !param.IsOmitted(u.OfExtractsBrowserActionWaitAction) {
		return u.OfExtractsBrowserActionWaitAction
	} else if !param.IsOmitted(u.OfExtractsBrowserActionWaitForElementAction) {
		return u.OfExtractsBrowserActionWaitForElementAction
	} else if !param.IsOmitted(u.OfExtractsBrowserActionWaitForNavigationAction) {
		return u.OfExtractsBrowserActionWaitForNavigationAction
	}
	return nil
}

// Continuously scroll to load dynamic content
//
// The property AutoScroll is required.
type ExtractParamsBrowserActionAutoScrollAction struct {
	AutoScroll ExtractParamsBrowserActionAutoScrollActionAutoScrollUnion `json:"auto_scroll,omitzero,required"`
	paramObj
}

func (r ExtractParamsBrowserActionAutoScrollAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionAutoScrollAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionAutoScrollAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionAutoScrollActionAutoScrollUnion struct {
	OfBool                                                  param.Opt[bool]                                             `json:",omitzero,inline"`
	OfFloat                                                 param.Opt[float64]                                          `json:",omitzero,inline"`
	OfString                                                param.Opt[string]                                           `json:",omitzero,inline"`
	OfExtractsBrowserActionAutoScrollActionAutoScrollObject *ExtractParamsBrowserActionAutoScrollActionAutoScrollObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionAutoScrollActionAutoScrollUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfExtractsBrowserActionAutoScrollActionAutoScrollObject)
}
func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractsBrowserActionAutoScrollActionAutoScrollObject) {
		return u.OfExtractsBrowserActionAutoScrollActionAutoScrollObject
	}
	return nil
}

type ExtractParamsBrowserActionAutoScrollActionAutoScrollObject struct {
	StepSize param.Opt[float64] `json:"step_size,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	ClickSelector ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion `json:"click_selector,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Container ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion `json:"container,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	DelayAfterScroll ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion `json:"delay_after_scroll,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	IdleTimeout ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion `json:"idle_timeout,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	LoadingSelector ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion `json:"loading_selector,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	MaxDuration ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion `json:"max_duration,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	PauseOnSelector ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion `json:"pause_on_selector,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractParamsBrowserActionAutoScrollActionAutoScrollObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionAutoScrollActionAutoScrollObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionAutoScrollActionAutoScrollObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion) asAny() any {
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
type ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion) asAny() any {
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
type ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion) asAny() any {
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
type ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion) asAny() any {
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
type ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionAutoScrollActionAutoScrollObjectRequiredString)
	OfExtractsBrowserActionAutoScrollActionAutoScrollObjectRequiredString param.Opt[ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                param.Opt[bool]                                                                     `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionAutoScrollActionAutoScrollObjectRequiredString, u.OfBool)
}
func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionAutoScrollActionAutoScrollObjectRequiredString) {
		return &u.OfExtractsBrowserActionAutoScrollActionAutoScrollObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredString string

const (
	ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredStringTrue  ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredString = "true"
	ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredStringFalse ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionAutoScrollActionAutoScrollObjectSkipString)
	OfExtractsBrowserActionAutoScrollActionAutoScrollObjectSkipString param.Opt[ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                            param.Opt[bool]                                                                 `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionAutoScrollActionAutoScrollObjectSkipString, u.OfBool)
}
func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionAutoScrollActionAutoScrollObjectSkipString) {
		return &u.OfExtractsBrowserActionAutoScrollActionAutoScrollObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipString string

const (
	ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipStringTrue  ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipString = "true"
	ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipStringFalse ExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipString = "false"
)

// Click on an element by selector
//
// The property Click is required.
type ExtractParamsBrowserActionClickAction struct {
	Click ExtractParamsBrowserActionClickActionClickUnion `json:"click,omitzero,required"`
	paramObj
}

func (r ExtractParamsBrowserActionClickAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionClickAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionClickAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionClickActionClickUnion struct {
	OfString                                      param.Opt[string]                                 `json:",omitzero,inline"`
	OfStringArray                                 []string                                          `json:",omitzero,inline"`
	OfExtractsBrowserActionClickActionClickObject *ExtractParamsBrowserActionClickActionClickObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionClickActionClickUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray, u.OfExtractsBrowserActionClickActionClickObject)
}
func (u *ExtractParamsBrowserActionClickActionClickUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionClickActionClickUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfExtractsBrowserActionClickActionClickObject) {
		return u.OfExtractsBrowserActionClickActionClickObject
	}
	return nil
}

// The property Selector is required.
type ExtractParamsBrowserActionClickActionClickObject struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector ExtractParamsBrowserActionClickActionClickObjectSelectorUnion `json:"selector,omitzero,required"`
	Count    param.Opt[float64]                                            `json:"count,omitzero"`
	OffsetX  param.Opt[int64]                                              `json:"offset_x,omitzero"`
	OffsetY  param.Opt[int64]                                              `json:"offset_y,omitzero"`
	Scroll   param.Opt[bool]                                               `json:"scroll,omitzero"`
	Steps    param.Opt[float64]                                            `json:"steps,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay ExtractParamsBrowserActionClickActionClickObjectDelayUnion `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractParamsBrowserActionClickActionClickObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractParamsBrowserActionClickActionClickObjectSkipUnion `json:"skip,omitzero"`
	// Any of "linear", "ghost-cursor", "windmouse".
	Strategy string `json:"strategy,omitzero"`
	paramObj
}

func (r ExtractParamsBrowserActionClickActionClickObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionClickActionClickObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionClickActionClickObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractParamsBrowserActionClickActionClickObject](
		"strategy", "linear", "ghost-cursor", "windmouse",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionClickActionClickObjectSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionClickActionClickObjectSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractParamsBrowserActionClickActionClickObjectSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionClickActionClickObjectSelectorUnion) asAny() any {
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
type ExtractParamsBrowserActionClickActionClickObjectDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionClickActionClickObjectDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractParamsBrowserActionClickActionClickObjectDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionClickActionClickObjectDelayUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionClickActionClickObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionClickActionClickObjectRequiredString)
	OfExtractsBrowserActionClickActionClickObjectRequiredString param.Opt[ExtractParamsBrowserActionClickActionClickObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                      param.Opt[bool]                                                           `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionClickActionClickObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionClickActionClickObjectRequiredString, u.OfBool)
}
func (u *ExtractParamsBrowserActionClickActionClickObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionClickActionClickObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionClickActionClickObjectRequiredString) {
		return &u.OfExtractsBrowserActionClickActionClickObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionClickActionClickObjectRequiredString string

const (
	ExtractParamsBrowserActionClickActionClickObjectRequiredStringTrue  ExtractParamsBrowserActionClickActionClickObjectRequiredString = "true"
	ExtractParamsBrowserActionClickActionClickObjectRequiredStringFalse ExtractParamsBrowserActionClickActionClickObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionClickActionClickObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionClickActionClickObjectSkipString)
	OfExtractsBrowserActionClickActionClickObjectSkipString param.Opt[ExtractParamsBrowserActionClickActionClickObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                  param.Opt[bool]                                                       `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionClickActionClickObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionClickActionClickObjectSkipString, u.OfBool)
}
func (u *ExtractParamsBrowserActionClickActionClickObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionClickActionClickObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionClickActionClickObjectSkipString) {
		return &u.OfExtractsBrowserActionClickActionClickObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionClickActionClickObjectSkipString string

const (
	ExtractParamsBrowserActionClickActionClickObjectSkipStringTrue  ExtractParamsBrowserActionClickActionClickObjectSkipString = "true"
	ExtractParamsBrowserActionClickActionClickObjectSkipStringFalse ExtractParamsBrowserActionClickActionClickObjectSkipString = "false"
)

// Execute JavaScript code in page context
//
// The property Eval is required.
type ExtractParamsBrowserActionEvalAction struct {
	Eval ExtractParamsBrowserActionEvalActionEvalUnion `json:"eval,omitzero,required"`
	paramObj
}

func (r ExtractParamsBrowserActionEvalAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionEvalAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionEvalAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionEvalActionEvalUnion struct {
	OfString                                    param.Opt[string]                               `json:",omitzero,inline"`
	OfExtractsBrowserActionEvalActionEvalObject *ExtractParamsBrowserActionEvalActionEvalObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionEvalActionEvalUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractsBrowserActionEvalActionEvalObject)
}
func (u *ExtractParamsBrowserActionEvalActionEvalUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionEvalActionEvalUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractsBrowserActionEvalActionEvalObject) {
		return u.OfExtractsBrowserActionEvalActionEvalObject
	}
	return nil
}

// The property Code is required.
type ExtractParamsBrowserActionEvalActionEvalObject struct {
	Code string `json:"code,required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractParamsBrowserActionEvalActionEvalObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractParamsBrowserActionEvalActionEvalObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractParamsBrowserActionEvalActionEvalObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionEvalActionEvalObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionEvalActionEvalObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionEvalActionEvalObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionEvalActionEvalObjectRequiredString)
	OfExtractsBrowserActionEvalActionEvalObjectRequiredString param.Opt[ExtractParamsBrowserActionEvalActionEvalObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                    param.Opt[bool]                                                         `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionEvalActionEvalObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionEvalActionEvalObjectRequiredString, u.OfBool)
}
func (u *ExtractParamsBrowserActionEvalActionEvalObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionEvalActionEvalObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionEvalActionEvalObjectRequiredString) {
		return &u.OfExtractsBrowserActionEvalActionEvalObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionEvalActionEvalObjectRequiredString string

const (
	ExtractParamsBrowserActionEvalActionEvalObjectRequiredStringTrue  ExtractParamsBrowserActionEvalActionEvalObjectRequiredString = "true"
	ExtractParamsBrowserActionEvalActionEvalObjectRequiredStringFalse ExtractParamsBrowserActionEvalActionEvalObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionEvalActionEvalObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionEvalActionEvalObjectSkipString)
	OfExtractsBrowserActionEvalActionEvalObjectSkipString param.Opt[ExtractParamsBrowserActionEvalActionEvalObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                param.Opt[bool]                                                     `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionEvalActionEvalObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionEvalActionEvalObjectSkipString, u.OfBool)
}
func (u *ExtractParamsBrowserActionEvalActionEvalObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionEvalActionEvalObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionEvalActionEvalObjectSkipString) {
		return &u.OfExtractsBrowserActionEvalActionEvalObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionEvalActionEvalObjectSkipString string

const (
	ExtractParamsBrowserActionEvalActionEvalObjectSkipStringTrue  ExtractParamsBrowserActionEvalActionEvalObjectSkipString = "true"
	ExtractParamsBrowserActionEvalActionEvalObjectSkipStringFalse ExtractParamsBrowserActionEvalActionEvalObjectSkipString = "false"
)

// Make an HTTP request in browser context
//
// The property Fetch is required.
type ExtractParamsBrowserActionFetchAction struct {
	Fetch ExtractParamsBrowserActionFetchActionFetchUnion `json:"fetch,omitzero,required" format:"uri"`
	paramObj
}

func (r ExtractParamsBrowserActionFetchAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionFetchAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionFetchAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionFetchActionFetchUnion struct {
	OfString                                      param.Opt[string]                                 `json:",omitzero,inline"`
	OfExtractsBrowserActionFetchActionFetchObject *ExtractParamsBrowserActionFetchActionFetchObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionFetchActionFetchUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractsBrowserActionFetchActionFetchObject)
}
func (u *ExtractParamsBrowserActionFetchActionFetchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionFetchActionFetchUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractsBrowserActionFetchActionFetchObject) {
		return u.OfExtractsBrowserActionFetchActionFetchObject
	}
	return nil
}

// The property URL is required.
type ExtractParamsBrowserActionFetchActionFetchObject struct {
	URL  string            `json:"url,required" format:"uri"`
	Body param.Opt[string] `json:"body,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Headers map[string]string  `json:"headers,omitzero"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractParamsBrowserActionFetchActionFetchObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractParamsBrowserActionFetchActionFetchObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractParamsBrowserActionFetchActionFetchObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionFetchActionFetchObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionFetchActionFetchObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractParamsBrowserActionFetchActionFetchObject](
		"method", "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionFetchActionFetchObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionFetchActionFetchObjectRequiredString)
	OfExtractsBrowserActionFetchActionFetchObjectRequiredString param.Opt[ExtractParamsBrowserActionFetchActionFetchObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                      param.Opt[bool]                                                           `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionFetchActionFetchObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionFetchActionFetchObjectRequiredString, u.OfBool)
}
func (u *ExtractParamsBrowserActionFetchActionFetchObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionFetchActionFetchObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionFetchActionFetchObjectRequiredString) {
		return &u.OfExtractsBrowserActionFetchActionFetchObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionFetchActionFetchObjectRequiredString string

const (
	ExtractParamsBrowserActionFetchActionFetchObjectRequiredStringTrue  ExtractParamsBrowserActionFetchActionFetchObjectRequiredString = "true"
	ExtractParamsBrowserActionFetchActionFetchObjectRequiredStringFalse ExtractParamsBrowserActionFetchActionFetchObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionFetchActionFetchObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionFetchActionFetchObjectSkipString)
	OfExtractsBrowserActionFetchActionFetchObjectSkipString param.Opt[ExtractParamsBrowserActionFetchActionFetchObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                  param.Opt[bool]                                                       `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionFetchActionFetchObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionFetchActionFetchObjectSkipString, u.OfBool)
}
func (u *ExtractParamsBrowserActionFetchActionFetchObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionFetchActionFetchObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionFetchActionFetchObjectSkipString) {
		return &u.OfExtractsBrowserActionFetchActionFetchObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionFetchActionFetchObjectSkipString string

const (
	ExtractParamsBrowserActionFetchActionFetchObjectSkipStringTrue  ExtractParamsBrowserActionFetchActionFetchObjectSkipString = "true"
	ExtractParamsBrowserActionFetchActionFetchObjectSkipStringFalse ExtractParamsBrowserActionFetchActionFetchObjectSkipString = "false"
)

// Fill text into an input field
//
// The property Fill is required.
type ExtractParamsBrowserActionFillAction struct {
	// Fill options with mode-specific fields. Use "type" mode for behavioral typing
	// simulation, or "paste" mode for instant paste.
	Fill ExtractParamsBrowserActionFillActionFillUnion `json:"fill,omitzero,required"`
	paramObj
}

func (r ExtractParamsBrowserActionFillAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionFillAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionFillAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionFillActionFillUnion struct {
	OfType  *ExtractParamsBrowserActionFillActionFillType  `json:",omitzero,inline"`
	OfPaste *ExtractParamsBrowserActionFillActionFillPaste `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionFillActionFillUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfType, u.OfPaste)
}
func (u *ExtractParamsBrowserActionFillActionFillUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionFillActionFillUnion) asAny() any {
	if !param.IsOmitted(u.OfType) {
		return u.OfType
	} else if !param.IsOmitted(u.OfPaste) {
		return u.OfPaste
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractParamsBrowserActionFillActionFillUnion) GetMouseMovementStrategy() *string {
	if vt := u.OfType; vt != nil {
		return &vt.MouseMovementStrategy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractParamsBrowserActionFillActionFillUnion) GetTypingInterval() *ExtractParamsBrowserActionFillActionFillTypeTypingIntervalUnion {
	if vt := u.OfType; vt != nil {
		return &vt.TypingInterval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractParamsBrowserActionFillActionFillUnion) GetTypingStrategy() *string {
	if vt := u.OfType; vt != nil {
		return &vt.TypingStrategy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractParamsBrowserActionFillActionFillUnion) GetValue() *string {
	if vt := u.OfType; vt != nil {
		return (*string)(&vt.Value)
	} else if vt := u.OfPaste; vt != nil {
		return (*string)(&vt.Value)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractParamsBrowserActionFillActionFillUnion) GetClickOnElement() *bool {
	if vt := u.OfType; vt != nil && vt.ClickOnElement.Valid() {
		return &vt.ClickOnElement.Value
	} else if vt := u.OfPaste; vt != nil && vt.ClickOnElement.Valid() {
		return &vt.ClickOnElement.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractParamsBrowserActionFillActionFillUnion) GetMode() *string {
	if vt := u.OfType; vt != nil {
		return (*string)(&vt.Mode)
	} else if vt := u.OfPaste; vt != nil {
		return (*string)(&vt.Mode)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractParamsBrowserActionFillActionFillUnion) GetScroll() *bool {
	if vt := u.OfType; vt != nil && vt.Scroll.Valid() {
		return &vt.Scroll.Value
	} else if vt := u.OfPaste; vt != nil && vt.Scroll.Valid() {
		return &vt.Scroll.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractParamsBrowserActionFillActionFillUnion) GetTimeout() *float64 {
	if vt := u.OfType; vt != nil && vt.Timeout.Valid() {
		return &vt.Timeout.Value
	} else if vt := u.OfPaste; vt != nil && vt.Timeout.Valid() {
		return &vt.Timeout.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractParamsBrowserActionFillActionFillUnion) GetVisible() *bool {
	if vt := u.OfType; vt != nil && vt.Visible.Valid() {
		return &vt.Visible.Value
	} else if vt := u.OfPaste; vt != nil && vt.Visible.Valid() {
		return &vt.Visible.Value
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ExtractParamsBrowserActionFillActionFillUnion) GetSelector() (res extractParamsBrowserActionFillActionFillUnionSelector) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Selector.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Selector.asAny()
	}
	return
}

// Can have the runtime types [*string], [\*[]string]
type extractParamsBrowserActionFillActionFillUnionSelector struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extractParamsBrowserActionFillActionFillUnionSelector) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ExtractParamsBrowserActionFillActionFillUnion) GetDelay() (res extractParamsBrowserActionFillActionFillUnionDelay) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Delay.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Delay.asAny()
	}
	return
}

// Can have the runtime types [*float64], [*string]
type extractParamsBrowserActionFillActionFillUnionDelay struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *float64:
//	case *string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extractParamsBrowserActionFillActionFillUnionDelay) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ExtractParamsBrowserActionFillActionFillUnion) GetRequired() (res extractParamsBrowserActionFillActionFillUnionRequired) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Required.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Required.asAny()
	}
	return
}

// Can have the runtime types [*string], [*bool]
type extractParamsBrowserActionFillActionFillUnionRequired struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extractParamsBrowserActionFillActionFillUnionRequired) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ExtractParamsBrowserActionFillActionFillUnion) GetSkip() (res extractParamsBrowserActionFillActionFillUnionSkip) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Skip.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Skip.asAny()
	}
	return
}

// Can have the runtime types [*string], [*bool]
type extractParamsBrowserActionFillActionFillUnionSkip struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extractParamsBrowserActionFillActionFillUnionSkip) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[ExtractParamsBrowserActionFillActionFillUnion](
		"mode",
		apijson.Discriminator[ExtractParamsBrowserActionFillActionFillType]("type"),
		apijson.Discriminator[ExtractParamsBrowserActionFillActionFillPaste]("paste"),
	)
}

// The properties Selector, Value are required.
type ExtractParamsBrowserActionFillActionFillType struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector       ExtractParamsBrowserActionFillActionFillTypeSelectorUnion `json:"selector,omitzero,required"`
	Value          string                                                    `json:"value,required"`
	ClickOnElement param.Opt[bool]                                           `json:"click_on_element,omitzero"`
	Scroll         param.Opt[bool]                                           `json:"scroll,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay ExtractParamsBrowserActionFillActionFillTypeDelayUnion `json:"delay,omitzero"`
	// Any of "type".
	Mode string `json:"mode,omitzero"`
	// Any of "linear", "ghost-cursor", "windmouse".
	MouseMovementStrategy string `json:"mouse_movement_strategy,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractParamsBrowserActionFillActionFillTypeRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractParamsBrowserActionFillActionFillTypeSkipUnion `json:"skip,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	TypingInterval ExtractParamsBrowserActionFillActionFillTypeTypingIntervalUnion `json:"typing_interval,omitzero"`
	// Any of "simple", "distribution".
	TypingStrategy string `json:"typing_strategy,omitzero"`
	paramObj
}

func (r ExtractParamsBrowserActionFillActionFillType) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionFillActionFillType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionFillActionFillType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractParamsBrowserActionFillActionFillType](
		"mode", "type",
	)
	apijson.RegisterFieldValidator[ExtractParamsBrowserActionFillActionFillType](
		"mouse_movement_strategy", "linear", "ghost-cursor", "windmouse",
	)
	apijson.RegisterFieldValidator[ExtractParamsBrowserActionFillActionFillType](
		"typing_strategy", "simple", "distribution",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionFillActionFillTypeSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionFillActionFillTypeSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractParamsBrowserActionFillActionFillTypeSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionFillActionFillTypeSelectorUnion) asAny() any {
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
type ExtractParamsBrowserActionFillActionFillTypeDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionFillActionFillTypeDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractParamsBrowserActionFillActionFillTypeDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionFillActionFillTypeDelayUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionFillActionFillTypeRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionFillActionFillTypeRequiredString)
	OfExtractsBrowserActionFillActionFillTypeRequiredString param.Opt[ExtractParamsBrowserActionFillActionFillTypeRequiredString] `json:",omitzero,inline"`
	OfBool                                                  param.Opt[bool]                                                       `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionFillActionFillTypeRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionFillActionFillTypeRequiredString, u.OfBool)
}
func (u *ExtractParamsBrowserActionFillActionFillTypeRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionFillActionFillTypeRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionFillActionFillTypeRequiredString) {
		return &u.OfExtractsBrowserActionFillActionFillTypeRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionFillActionFillTypeRequiredString string

const (
	ExtractParamsBrowserActionFillActionFillTypeRequiredStringTrue  ExtractParamsBrowserActionFillActionFillTypeRequiredString = "true"
	ExtractParamsBrowserActionFillActionFillTypeRequiredStringFalse ExtractParamsBrowserActionFillActionFillTypeRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionFillActionFillTypeSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionFillActionFillTypeSkipString)
	OfExtractsBrowserActionFillActionFillTypeSkipString param.Opt[ExtractParamsBrowserActionFillActionFillTypeSkipString] `json:",omitzero,inline"`
	OfBool                                              param.Opt[bool]                                                   `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionFillActionFillTypeSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionFillActionFillTypeSkipString, u.OfBool)
}
func (u *ExtractParamsBrowserActionFillActionFillTypeSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionFillActionFillTypeSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionFillActionFillTypeSkipString) {
		return &u.OfExtractsBrowserActionFillActionFillTypeSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionFillActionFillTypeSkipString string

const (
	ExtractParamsBrowserActionFillActionFillTypeSkipStringTrue  ExtractParamsBrowserActionFillActionFillTypeSkipString = "true"
	ExtractParamsBrowserActionFillActionFillTypeSkipStringFalse ExtractParamsBrowserActionFillActionFillTypeSkipString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionFillActionFillTypeTypingIntervalUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionFillActionFillTypeTypingIntervalUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractParamsBrowserActionFillActionFillTypeTypingIntervalUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionFillActionFillTypeTypingIntervalUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// The properties Mode, Selector, Value are required.
type ExtractParamsBrowserActionFillActionFillPaste struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector       ExtractParamsBrowserActionFillActionFillPasteSelectorUnion `json:"selector,omitzero,required"`
	Value          string                                                     `json:"value,required"`
	ClickOnElement param.Opt[bool]                                            `json:"click_on_element,omitzero"`
	Scroll         param.Opt[bool]                                            `json:"scroll,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay ExtractParamsBrowserActionFillActionFillPasteDelayUnion `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractParamsBrowserActionFillActionFillPasteRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractParamsBrowserActionFillActionFillPasteSkipUnion `json:"skip,omitzero"`
	// This field can be elided, and will marshal its zero value as "paste".
	Mode constant.Paste `json:"mode,required"`
	paramObj
}

func (r ExtractParamsBrowserActionFillActionFillPaste) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionFillActionFillPaste
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionFillActionFillPaste) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionFillActionFillPasteSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionFillActionFillPasteSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractParamsBrowserActionFillActionFillPasteSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionFillActionFillPasteSelectorUnion) asAny() any {
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
type ExtractParamsBrowserActionFillActionFillPasteDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionFillActionFillPasteDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractParamsBrowserActionFillActionFillPasteDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionFillActionFillPasteDelayUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionFillActionFillPasteRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionFillActionFillPasteRequiredString)
	OfExtractsBrowserActionFillActionFillPasteRequiredString param.Opt[ExtractParamsBrowserActionFillActionFillPasteRequiredString] `json:",omitzero,inline"`
	OfBool                                                   param.Opt[bool]                                                        `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionFillActionFillPasteRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionFillActionFillPasteRequiredString, u.OfBool)
}
func (u *ExtractParamsBrowserActionFillActionFillPasteRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionFillActionFillPasteRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionFillActionFillPasteRequiredString) {
		return &u.OfExtractsBrowserActionFillActionFillPasteRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionFillActionFillPasteRequiredString string

const (
	ExtractParamsBrowserActionFillActionFillPasteRequiredStringTrue  ExtractParamsBrowserActionFillActionFillPasteRequiredString = "true"
	ExtractParamsBrowserActionFillActionFillPasteRequiredStringFalse ExtractParamsBrowserActionFillActionFillPasteRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionFillActionFillPasteSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionFillActionFillPasteSkipString)
	OfExtractsBrowserActionFillActionFillPasteSkipString param.Opt[ExtractParamsBrowserActionFillActionFillPasteSkipString] `json:",omitzero,inline"`
	OfBool                                               param.Opt[bool]                                                    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionFillActionFillPasteSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionFillActionFillPasteSkipString, u.OfBool)
}
func (u *ExtractParamsBrowserActionFillActionFillPasteSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionFillActionFillPasteSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionFillActionFillPasteSkipString) {
		return &u.OfExtractsBrowserActionFillActionFillPasteSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionFillActionFillPasteSkipString string

const (
	ExtractParamsBrowserActionFillActionFillPasteSkipStringTrue  ExtractParamsBrowserActionFillActionFillPasteSkipString = "true"
	ExtractParamsBrowserActionFillActionFillPasteSkipStringFalse ExtractParamsBrowserActionFillActionFillPasteSkipString = "false"
)

// Retrieve browser cookies
//
// The property GetCookies is required.
type ExtractParamsBrowserActionGetCookiesAction struct {
	GetCookies ExtractParamsBrowserActionGetCookiesActionGetCookiesUnion `json:"get_cookies,omitzero,required"`
	paramObj
}

func (r ExtractParamsBrowserActionGetCookiesAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionGetCookiesAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionGetCookiesAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionGetCookiesActionGetCookiesUnion struct {
	OfBool                                                  param.Opt[bool]                                             `json:",omitzero,inline"`
	OfExtractsBrowserActionGetCookiesActionGetCookiesObject *ExtractParamsBrowserActionGetCookiesActionGetCookiesObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionGetCookiesActionGetCookiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfExtractsBrowserActionGetCookiesActionGetCookiesObject)
}
func (u *ExtractParamsBrowserActionGetCookiesActionGetCookiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionGetCookiesActionGetCookiesUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfExtractsBrowserActionGetCookiesActionGetCookiesObject) {
		return u.OfExtractsBrowserActionGetCookiesActionGetCookiesObject
	}
	return nil
}

type ExtractParamsBrowserActionGetCookiesActionGetCookiesObject struct {
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip        ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion `json:"skip,omitzero"`
	ExtraFields map[string]any                                                      `json:"-"`
	paramObj
}

func (r ExtractParamsBrowserActionGetCookiesActionGetCookiesObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionGetCookiesActionGetCookiesObject
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ExtractParamsBrowserActionGetCookiesActionGetCookiesObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionGetCookiesActionGetCookiesObjectRequiredString)
	OfExtractsBrowserActionGetCookiesActionGetCookiesObjectRequiredString param.Opt[ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                param.Opt[bool]                                                                     `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionGetCookiesActionGetCookiesObjectRequiredString, u.OfBool)
}
func (u *ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionGetCookiesActionGetCookiesObjectRequiredString) {
		return &u.OfExtractsBrowserActionGetCookiesActionGetCookiesObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredString string

const (
	ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredStringTrue  ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredString = "true"
	ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredStringFalse ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionGetCookiesActionGetCookiesObjectSkipString)
	OfExtractsBrowserActionGetCookiesActionGetCookiesObjectSkipString param.Opt[ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                            param.Opt[bool]                                                                 `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionGetCookiesActionGetCookiesObjectSkipString, u.OfBool)
}
func (u *ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionGetCookiesActionGetCookiesObjectSkipString) {
		return &u.OfExtractsBrowserActionGetCookiesActionGetCookiesObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipString string

const (
	ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipStringTrue  ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipString = "true"
	ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipStringFalse ExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipString = "false"
)

// Navigate to a URL
//
// The property Goto is required.
type ExtractParamsBrowserActionGotoAction struct {
	Goto ExtractParamsBrowserActionGotoActionGotoUnion `json:"goto,omitzero,required" format:"uri"`
	paramObj
}

func (r ExtractParamsBrowserActionGotoAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionGotoAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionGotoAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionGotoActionGotoUnion struct {
	OfString                                    param.Opt[string]                               `json:",omitzero,inline"`
	OfExtractsBrowserActionGotoActionGotoObject *ExtractParamsBrowserActionGotoActionGotoObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionGotoActionGotoUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractsBrowserActionGotoActionGotoObject)
}
func (u *ExtractParamsBrowserActionGotoActionGotoUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionGotoActionGotoUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractsBrowserActionGotoActionGotoObject) {
		return u.OfExtractsBrowserActionGotoActionGotoObject
	}
	return nil
}

// The property URL is required.
type ExtractParamsBrowserActionGotoActionGotoObject struct {
	URL     string            `json:"url,required" format:"uri"`
	Referer param.Opt[string] `json:"referer,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractParamsBrowserActionGotoActionGotoObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractParamsBrowserActionGotoActionGotoObjectSkipUnion `json:"skip,omitzero"`
	// Any of "load", "domcontentloaded", "networkidle0", "networkidle2".
	WaitUntil string `json:"wait_until,omitzero"`
	paramObj
}

func (r ExtractParamsBrowserActionGotoActionGotoObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionGotoActionGotoObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionGotoActionGotoObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractParamsBrowserActionGotoActionGotoObject](
		"wait_until", "load", "domcontentloaded", "networkidle0", "networkidle2",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionGotoActionGotoObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionGotoActionGotoObjectRequiredString)
	OfExtractsBrowserActionGotoActionGotoObjectRequiredString param.Opt[ExtractParamsBrowserActionGotoActionGotoObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                    param.Opt[bool]                                                         `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionGotoActionGotoObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionGotoActionGotoObjectRequiredString, u.OfBool)
}
func (u *ExtractParamsBrowserActionGotoActionGotoObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionGotoActionGotoObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionGotoActionGotoObjectRequiredString) {
		return &u.OfExtractsBrowserActionGotoActionGotoObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionGotoActionGotoObjectRequiredString string

const (
	ExtractParamsBrowserActionGotoActionGotoObjectRequiredStringTrue  ExtractParamsBrowserActionGotoActionGotoObjectRequiredString = "true"
	ExtractParamsBrowserActionGotoActionGotoObjectRequiredStringFalse ExtractParamsBrowserActionGotoActionGotoObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionGotoActionGotoObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionGotoActionGotoObjectSkipString)
	OfExtractsBrowserActionGotoActionGotoObjectSkipString param.Opt[ExtractParamsBrowserActionGotoActionGotoObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                param.Opt[bool]                                                     `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionGotoActionGotoObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionGotoActionGotoObjectSkipString, u.OfBool)
}
func (u *ExtractParamsBrowserActionGotoActionGotoObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionGotoActionGotoObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionGotoActionGotoObjectSkipString) {
		return &u.OfExtractsBrowserActionGotoActionGotoObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionGotoActionGotoObjectSkipString string

const (
	ExtractParamsBrowserActionGotoActionGotoObjectSkipStringTrue  ExtractParamsBrowserActionGotoActionGotoObjectSkipString = "true"
	ExtractParamsBrowserActionGotoActionGotoObjectSkipStringFalse ExtractParamsBrowserActionGotoActionGotoObjectSkipString = "false"
)

// Press a keyboard key
//
// The property Press is required.
type ExtractParamsBrowserActionPressAction struct {
	Press ExtractParamsBrowserActionPressActionPressUnion `json:"press,omitzero,required"`
	paramObj
}

func (r ExtractParamsBrowserActionPressAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionPressAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionPressAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionPressActionPressUnion struct {
	OfString                                      param.Opt[string]                                 `json:",omitzero,inline"`
	OfExtractsBrowserActionPressActionPressObject *ExtractParamsBrowserActionPressActionPressObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionPressActionPressUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractsBrowserActionPressActionPressObject)
}
func (u *ExtractParamsBrowserActionPressActionPressUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionPressActionPressUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractsBrowserActionPressActionPressObject) {
		return u.OfExtractsBrowserActionPressActionPressObject
	}
	return nil
}

// The property Key is required.
type ExtractParamsBrowserActionPressActionPressObject struct {
	// Any of "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "Power", "Eject",
	// "Abort", "Help", "Backspace", "Tab", "Numpad5", "NumpadEnter", "Enter", "\r",
	// "\n", "ShiftLeft", "ShiftRight", "ControlLeft", "ControlRight", "AltLeft",
	// "AltRight", "Pause", "CapsLock", "Escape", "Convert", "NonConvert", "Space",
	// "Numpad9", "PageUp", "Numpad3", "PageDown", "End", "Numpad1", "Home", "Numpad7",
	// "ArrowLeft", "Numpad4", "Numpad8", "ArrowUp", "ArrowRight", "Numpad6",
	// "Numpad2", "ArrowDown", "Select", "Open", "PrintScreen", "Insert", "Numpad0",
	// "Delete", "NumpadDecimal", "Digit0", "Digit1", "Digit2", "Digit3", "Digit4",
	// "Digit5", "Digit6", "Digit7", "Digit8", "Digit9", "KeyA", "KeyB", "KeyC",
	// "KeyD", "KeyE", "KeyF", "KeyG", "KeyH", "KeyI", "KeyJ", "KeyK", "KeyL", "KeyM",
	// "KeyN", "KeyO", "KeyP", "KeyQ", "KeyR", "KeyS", "KeyT", "KeyU", "KeyV", "KeyW",
	// "KeyX", "KeyY", "KeyZ", "MetaLeft", "MetaRight", "ContextMenu",
	// "NumpadMultiply", "NumpadAdd", "NumpadSubtract", "NumpadDivide", "F1", "F2",
	// "F3", "F4", "F5", "F6", "F7", "F8", "F9", "F10", "F11", "F12", "F13", "F14",
	// "F15", "F16", "F17", "F18", "F19", "F20", "F21", "F22", "F23", "F24", "NumLock",
	// "ScrollLock", "AudioVolumeMute", "AudioVolumeDown", "AudioVolumeUp",
	// "MediaTrackNext", "MediaTrackPrevious", "MediaStop", "MediaPlayPause",
	// "Semicolon", "Equal", "NumpadEqual", "Comma", "Minus", "Period", "Slash",
	// "Backquote", "BracketLeft", "Backslash", "BracketRight", "Quote", "AltGraph",
	// "Props", "Cancel", "Clear", "Shift", "Control", "Alt", "Accept", "ModeChange", "
	// ", "Print", "Execute", "\u0000", "a", "b", "c", "d", "e", "f", "g", "h", "i",
	// "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y",
	// "z", "Meta", "\*", "+", "-", "/", ";", "=", ",", ".", "`", "[", "\\", "]", "'",
	// "Attn", "CrSel", "ExSel", "EraseEof", "Play", "ZoomOut", ")", "!", "@", "#",
	// "$", "%", "^", "&", "(", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K",
	// "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", ":",
	// "<", "\_", ">", "?", "~", "{", "|", "}", "\"", "SoftLeft", "SoftRight",
	// "Camera", "Call", "EndCall", "VolumeDown", "VolumeUp".
	Key string `json:"key,omitzero,required"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay ExtractParamsBrowserActionPressActionPressObjectDelayUnion `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractParamsBrowserActionPressActionPressObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractParamsBrowserActionPressActionPressObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractParamsBrowserActionPressActionPressObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionPressActionPressObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionPressActionPressObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractParamsBrowserActionPressActionPressObject](
		"key", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "Power", "Eject", "Abort", "Help", "Backspace", "Tab", "Numpad5", "NumpadEnter", "Enter", "\r", "\n", "ShiftLeft", "ShiftRight", "ControlLeft", "ControlRight", "AltLeft", "AltRight", "Pause", "CapsLock", "Escape", "Convert", "NonConvert", "Space", "Numpad9", "PageUp", "Numpad3", "PageDown", "End", "Numpad1", "Home", "Numpad7", "ArrowLeft", "Numpad4", "Numpad8", "ArrowUp", "ArrowRight", "Numpad6", "Numpad2", "ArrowDown", "Select", "Open", "PrintScreen", "Insert", "Numpad0", "Delete", "NumpadDecimal", "Digit0", "Digit1", "Digit2", "Digit3", "Digit4", "Digit5", "Digit6", "Digit7", "Digit8", "Digit9", "KeyA", "KeyB", "KeyC", "KeyD", "KeyE", "KeyF", "KeyG", "KeyH", "KeyI", "KeyJ", "KeyK", "KeyL", "KeyM", "KeyN", "KeyO", "KeyP", "KeyQ", "KeyR", "KeyS", "KeyT", "KeyU", "KeyV", "KeyW", "KeyX", "KeyY", "KeyZ", "MetaLeft", "MetaRight", "ContextMenu", "NumpadMultiply", "NumpadAdd", "NumpadSubtract", "NumpadDivide", "F1", "F2", "F3", "F4", "F5", "F6", "F7", "F8", "F9", "F10", "F11", "F12", "F13", "F14", "F15", "F16", "F17", "F18", "F19", "F20", "F21", "F22", "F23", "F24", "NumLock", "ScrollLock", "AudioVolumeMute", "AudioVolumeDown", "AudioVolumeUp", "MediaTrackNext", "MediaTrackPrevious", "MediaStop", "MediaPlayPause", "Semicolon", "Equal", "NumpadEqual", "Comma", "Minus", "Period", "Slash", "Backquote", "BracketLeft", "Backslash", "BracketRight", "Quote", "AltGraph", "Props", "Cancel", "Clear", "Shift", "Control", "Alt", "Accept", "ModeChange", " ", "Print", "Execute", "\u0000", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "Meta", "*", "+", "-", "/", ";", "=", ",", ".", "`", "[", "\\", "]", "'", "Attn", "CrSel", "ExSel", "EraseEof", "Play", "ZoomOut", ")", "!", "@", "#", "$", "%", "^", "&", "(", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", ":", "<", "_", ">", "?", "~", "{", "|", "}", "\"", "SoftLeft", "SoftRight", "Camera", "Call", "EndCall", "VolumeDown", "VolumeUp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionPressActionPressObjectDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionPressActionPressObjectDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractParamsBrowserActionPressActionPressObjectDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionPressActionPressObjectDelayUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionPressActionPressObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionPressActionPressObjectRequiredString)
	OfExtractsBrowserActionPressActionPressObjectRequiredString param.Opt[ExtractParamsBrowserActionPressActionPressObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                      param.Opt[bool]                                                           `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionPressActionPressObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionPressActionPressObjectRequiredString, u.OfBool)
}
func (u *ExtractParamsBrowserActionPressActionPressObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionPressActionPressObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionPressActionPressObjectRequiredString) {
		return &u.OfExtractsBrowserActionPressActionPressObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionPressActionPressObjectRequiredString string

const (
	ExtractParamsBrowserActionPressActionPressObjectRequiredStringTrue  ExtractParamsBrowserActionPressActionPressObjectRequiredString = "true"
	ExtractParamsBrowserActionPressActionPressObjectRequiredStringFalse ExtractParamsBrowserActionPressActionPressObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionPressActionPressObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionPressActionPressObjectSkipString)
	OfExtractsBrowserActionPressActionPressObjectSkipString param.Opt[ExtractParamsBrowserActionPressActionPressObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                  param.Opt[bool]                                                       `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionPressActionPressObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionPressActionPressObjectSkipString, u.OfBool)
}
func (u *ExtractParamsBrowserActionPressActionPressObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionPressActionPressObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionPressActionPressObjectSkipString) {
		return &u.OfExtractsBrowserActionPressActionPressObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionPressActionPressObjectSkipString string

const (
	ExtractParamsBrowserActionPressActionPressObjectSkipStringTrue  ExtractParamsBrowserActionPressActionPressObjectSkipString = "true"
	ExtractParamsBrowserActionPressActionPressObjectSkipStringFalse ExtractParamsBrowserActionPressActionPressObjectSkipString = "false"
)

// Capture a page screenshot
//
// The property Screenshot is required.
type ExtractParamsBrowserActionScreenshotAction struct {
	Screenshot ExtractParamsBrowserActionScreenshotActionScreenshotUnion `json:"screenshot,omitzero,required"`
	paramObj
}

func (r ExtractParamsBrowserActionScreenshotAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionScreenshotAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionScreenshotAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionScreenshotActionScreenshotUnion struct {
	OfBool                                                  param.Opt[bool]                                             `json:",omitzero,inline"`
	OfExtractsBrowserActionScreenshotActionScreenshotObject *ExtractParamsBrowserActionScreenshotActionScreenshotObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionScreenshotActionScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfExtractsBrowserActionScreenshotActionScreenshotObject)
}
func (u *ExtractParamsBrowserActionScreenshotActionScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionScreenshotActionScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfExtractsBrowserActionScreenshotActionScreenshotObject) {
		return u.OfExtractsBrowserActionScreenshotActionScreenshotObject
	}
	return nil
}

type ExtractParamsBrowserActionScreenshotActionScreenshotObject struct {
	FullPage param.Opt[bool]    `json:"full_page,omitzero"`
	Quality  param.Opt[float64] `json:"quality,omitzero"`
	// Any of "png", "jpeg", "webp".
	Format string `json:"format,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractParamsBrowserActionScreenshotActionScreenshotObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionScreenshotActionScreenshotObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionScreenshotActionScreenshotObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractParamsBrowserActionScreenshotActionScreenshotObject](
		"format", "png", "jpeg", "webp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionScreenshotActionScreenshotObjectRequiredString)
	OfExtractsBrowserActionScreenshotActionScreenshotObjectRequiredString param.Opt[ExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                param.Opt[bool]                                                                     `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionScreenshotActionScreenshotObjectRequiredString, u.OfBool)
}
func (u *ExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionScreenshotActionScreenshotObjectRequiredString) {
		return &u.OfExtractsBrowserActionScreenshotActionScreenshotObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredString string

const (
	ExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredStringTrue  ExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredString = "true"
	ExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredStringFalse ExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionScreenshotActionScreenshotObjectSkipString)
	OfExtractsBrowserActionScreenshotActionScreenshotObjectSkipString param.Opt[ExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                            param.Opt[bool]                                                                 `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionScreenshotActionScreenshotObjectSkipString, u.OfBool)
}
func (u *ExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionScreenshotActionScreenshotObjectSkipString) {
		return &u.OfExtractsBrowserActionScreenshotActionScreenshotObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipString string

const (
	ExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipStringTrue  ExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipString = "true"
	ExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipStringFalse ExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipString = "false"
)

// Scroll the page or an element
//
// The property Scroll is required.
type ExtractParamsBrowserActionScrollAction struct {
	Scroll ExtractParamsBrowserActionScrollActionScrollUnion `json:"scroll,omitzero,required"`
	paramObj
}

func (r ExtractParamsBrowserActionScrollAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionScrollAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionScrollAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionScrollActionScrollUnion struct {
	OfFloat                                         param.Opt[float64]                                  `json:",omitzero,inline"`
	OfString                                        param.Opt[string]                                   `json:",omitzero,inline"`
	OfExtractsBrowserActionScrollActionScrollObject *ExtractParamsBrowserActionScrollActionScrollObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionScrollActionScrollUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfExtractsBrowserActionScrollActionScrollObject)
}
func (u *ExtractParamsBrowserActionScrollActionScrollUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionScrollActionScrollUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractsBrowserActionScrollActionScrollObject) {
		return u.OfExtractsBrowserActionScrollActionScrollObject
	}
	return nil
}

type ExtractParamsBrowserActionScrollActionScrollObject struct {
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	X       param.Opt[float64] `json:"x,omitzero"`
	Y       param.Opt[float64] `json:"y,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Container ExtractParamsBrowserActionScrollActionScrollObjectContainerUnion `json:"container,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractParamsBrowserActionScrollActionScrollObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractParamsBrowserActionScrollActionScrollObjectSkipUnion `json:"skip,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	To ExtractParamsBrowserActionScrollActionScrollObjectToUnion `json:"to,omitzero"`
	paramObj
}

func (r ExtractParamsBrowserActionScrollActionScrollObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionScrollActionScrollObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionScrollActionScrollObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionScrollActionScrollObjectContainerUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionScrollActionScrollObjectContainerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractParamsBrowserActionScrollActionScrollObjectContainerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionScrollActionScrollObjectContainerUnion) asAny() any {
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
type ExtractParamsBrowserActionScrollActionScrollObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionScrollActionScrollObjectRequiredString)
	OfExtractsBrowserActionScrollActionScrollObjectRequiredString param.Opt[ExtractParamsBrowserActionScrollActionScrollObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                        param.Opt[bool]                                                             `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionScrollActionScrollObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionScrollActionScrollObjectRequiredString, u.OfBool)
}
func (u *ExtractParamsBrowserActionScrollActionScrollObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionScrollActionScrollObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionScrollActionScrollObjectRequiredString) {
		return &u.OfExtractsBrowserActionScrollActionScrollObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionScrollActionScrollObjectRequiredString string

const (
	ExtractParamsBrowserActionScrollActionScrollObjectRequiredStringTrue  ExtractParamsBrowserActionScrollActionScrollObjectRequiredString = "true"
	ExtractParamsBrowserActionScrollActionScrollObjectRequiredStringFalse ExtractParamsBrowserActionScrollActionScrollObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionScrollActionScrollObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionScrollActionScrollObjectSkipString)
	OfExtractsBrowserActionScrollActionScrollObjectSkipString param.Opt[ExtractParamsBrowserActionScrollActionScrollObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                    param.Opt[bool]                                                         `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionScrollActionScrollObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionScrollActionScrollObjectSkipString, u.OfBool)
}
func (u *ExtractParamsBrowserActionScrollActionScrollObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionScrollActionScrollObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionScrollActionScrollObjectSkipString) {
		return &u.OfExtractsBrowserActionScrollActionScrollObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionScrollActionScrollObjectSkipString string

const (
	ExtractParamsBrowserActionScrollActionScrollObjectSkipStringTrue  ExtractParamsBrowserActionScrollActionScrollObjectSkipString = "true"
	ExtractParamsBrowserActionScrollActionScrollObjectSkipStringFalse ExtractParamsBrowserActionScrollActionScrollObjectSkipString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionScrollActionScrollObjectToUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionScrollActionScrollObjectToUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractParamsBrowserActionScrollActionScrollObjectToUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionScrollActionScrollObjectToUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Wait for a specified duration
//
// The property Wait is required.
type ExtractParamsBrowserActionWaitAction struct {
	Wait ExtractParamsBrowserActionWaitActionWaitUnion `json:"wait,omitzero,required"`
	paramObj
}

func (r ExtractParamsBrowserActionWaitAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionWaitAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionWaitAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionWaitActionWaitUnion struct {
	OfFloat                                     param.Opt[float64]                              `json:",omitzero,inline"`
	OfString                                    param.Opt[string]                               `json:",omitzero,inline"`
	OfExtractsBrowserActionWaitActionWaitObject *ExtractParamsBrowserActionWaitActionWaitObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionWaitActionWaitUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfExtractsBrowserActionWaitActionWaitObject)
}
func (u *ExtractParamsBrowserActionWaitActionWaitUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionWaitActionWaitUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractsBrowserActionWaitActionWaitObject) {
		return u.OfExtractsBrowserActionWaitActionWaitObject
	}
	return nil
}

// The property Duration is required.
type ExtractParamsBrowserActionWaitActionWaitObject struct {
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Duration ExtractParamsBrowserActionWaitActionWaitObjectDurationUnion `json:"duration,omitzero,required"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractParamsBrowserActionWaitActionWaitObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractParamsBrowserActionWaitActionWaitObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractParamsBrowserActionWaitActionWaitObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionWaitActionWaitObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionWaitActionWaitObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionWaitActionWaitObjectDurationUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionWaitActionWaitObjectDurationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractParamsBrowserActionWaitActionWaitObjectDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionWaitActionWaitObjectDurationUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionWaitActionWaitObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionWaitActionWaitObjectRequiredString)
	OfExtractsBrowserActionWaitActionWaitObjectRequiredString param.Opt[ExtractParamsBrowserActionWaitActionWaitObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                    param.Opt[bool]                                                         `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionWaitActionWaitObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionWaitActionWaitObjectRequiredString, u.OfBool)
}
func (u *ExtractParamsBrowserActionWaitActionWaitObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionWaitActionWaitObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionWaitActionWaitObjectRequiredString) {
		return &u.OfExtractsBrowserActionWaitActionWaitObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionWaitActionWaitObjectRequiredString string

const (
	ExtractParamsBrowserActionWaitActionWaitObjectRequiredStringTrue  ExtractParamsBrowserActionWaitActionWaitObjectRequiredString = "true"
	ExtractParamsBrowserActionWaitActionWaitObjectRequiredStringFalse ExtractParamsBrowserActionWaitActionWaitObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionWaitActionWaitObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionWaitActionWaitObjectSkipString)
	OfExtractsBrowserActionWaitActionWaitObjectSkipString param.Opt[ExtractParamsBrowserActionWaitActionWaitObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                param.Opt[bool]                                                     `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionWaitActionWaitObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionWaitActionWaitObjectSkipString, u.OfBool)
}
func (u *ExtractParamsBrowserActionWaitActionWaitObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionWaitActionWaitObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionWaitActionWaitObjectSkipString) {
		return &u.OfExtractsBrowserActionWaitActionWaitObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionWaitActionWaitObjectSkipString string

const (
	ExtractParamsBrowserActionWaitActionWaitObjectSkipStringTrue  ExtractParamsBrowserActionWaitActionWaitObjectSkipString = "true"
	ExtractParamsBrowserActionWaitActionWaitObjectSkipStringFalse ExtractParamsBrowserActionWaitActionWaitObjectSkipString = "false"
)

// Wait for an element to appear or reach a specific state
//
// The property WaitForElement is required.
type ExtractParamsBrowserActionWaitForElementAction struct {
	WaitForElement ExtractParamsBrowserActionWaitForElementActionWaitForElementUnion `json:"wait_for_element,omitzero,required"`
	paramObj
}

func (r ExtractParamsBrowserActionWaitForElementAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionWaitForElementAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionWaitForElementAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionWaitForElementActionWaitForElementUnion struct {
	OfString                                                        param.Opt[string]                                                   `json:",omitzero,inline"`
	OfStringArray                                                   []string                                                            `json:",omitzero,inline"`
	OfExtractsBrowserActionWaitForElementActionWaitForElementObject *ExtractParamsBrowserActionWaitForElementActionWaitForElementObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionWaitForElementActionWaitForElementUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray, u.OfExtractsBrowserActionWaitForElementActionWaitForElementObject)
}
func (u *ExtractParamsBrowserActionWaitForElementActionWaitForElementUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionWaitForElementActionWaitForElementUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfExtractsBrowserActionWaitForElementActionWaitForElementObject) {
		return u.OfExtractsBrowserActionWaitForElementActionWaitForElementObject
	}
	return nil
}

// The property Selector is required.
type ExtractParamsBrowserActionWaitForElementActionWaitForElementObject struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion `json:"selector,omitzero,required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractParamsBrowserActionWaitForElementActionWaitForElementObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionWaitForElementActionWaitForElementObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionWaitForElementActionWaitForElementObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion) asAny() any {
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
type ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionWaitForElementActionWaitForElementObjectRequiredString)
	OfExtractsBrowserActionWaitForElementActionWaitForElementObjectRequiredString param.Opt[ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                        param.Opt[bool]                                                                             `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionWaitForElementActionWaitForElementObjectRequiredString, u.OfBool)
}
func (u *ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionWaitForElementActionWaitForElementObjectRequiredString) {
		return &u.OfExtractsBrowserActionWaitForElementActionWaitForElementObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredString string

const (
	ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredStringTrue  ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredString = "true"
	ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredStringFalse ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionWaitForElementActionWaitForElementObjectSkipString)
	OfExtractsBrowserActionWaitForElementActionWaitForElementObjectSkipString param.Opt[ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                    param.Opt[bool]                                                                         `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionWaitForElementActionWaitForElementObjectSkipString, u.OfBool)
}
func (u *ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionWaitForElementActionWaitForElementObjectSkipString) {
		return &u.OfExtractsBrowserActionWaitForElementActionWaitForElementObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipString string

const (
	ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipStringTrue  ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipString = "true"
	ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipStringFalse ExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipString = "false"
)

// Wait for page navigation to complete
//
// The property WaitForNavigation is required.
type ExtractParamsBrowserActionWaitForNavigationAction struct {
	WaitForNavigation ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion `json:"wait_for_navigation,omitzero,required"`
	paramObj
}

func (r ExtractParamsBrowserActionWaitForNavigationAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionWaitForNavigationAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionWaitForNavigationAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationString)
	OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationString param.Opt[string]                                                         `json:",omitzero,inline"`
	OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationObject *ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationString, u.OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationObject)
}
func (u *ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationString) {
		return &u.OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationString
	} else if !param.IsOmitted(u.OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationObject) {
		return u.OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationObject
	}
	return nil
}

type ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationString string

const (
	ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationStringLoad             ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationString = "load"
	ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationStringDomcontentloaded ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationString = "domcontentloaded"
	ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationStringNetworkidle0     ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationString = "networkidle0"
	ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationStringNetworkidle2     ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationString = "networkidle2"
)

// The property Navigation is required.
type ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObject struct {
	// Any of "load", "domcontentloaded", "networkidle0", "networkidle2".
	Navigation string `json:"navigation,omitzero,required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObject](
		"navigation", "load", "domcontentloaded", "networkidle0", "networkidle2",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString)
	OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString param.Opt[ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                              param.Opt[bool]                                                                                   `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString, u.OfBool)
}
func (u *ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString) {
		return &u.OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString string

const (
	ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredStringTrue  ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString = "true"
	ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredStringFalse ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString)
	OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString param.Opt[ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                          param.Opt[bool]                                                                               `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString, u.OfBool)
}
func (u *ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString) {
		return &u.OfExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString string

const (
	ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipStringTrue  ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString = "true"
	ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipStringFalse ExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString = "false"
)

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
	Value string `json:"value,required"`
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

type MapParams struct {
	// Url to map.
	URL string `json:"url,required"`
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
