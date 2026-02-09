// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomnimblewaynimblego

import (
	"encoding/json"

	"github.com/Nimbleway/nimble-go/internal/apijson"
	"github.com/Nimbleway/nimble-go/packages/param"
	"github.com/Nimbleway/nimble-go/packages/respjson"
	"github.com/Nimbleway/nimble-go/shared/constant"
)

type AgentResponse struct {
	Data     AgentResponseData     `json:"data,required"`
	Metadata AgentResponseMetadata `json:"metadata,required"`
	// The status of the task.
	//
	// Any of "success", "skipped", "fatal", "error", "postponed", "ignored",
	// "rejected", "blocked".
	Status AgentResponseStatus `json:"status,required"`
	// Unique identifier for the task.
	TaskID string `json:"task_id,required"`
	// The final URL.
	URL   string             `json:"url,required"`
	Debug AgentResponseDebug `json:"debug"`
	// Pagination information if applicable.
	Pagination AgentResponsePaginationUnion `json:"pagination"`
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
func (r AgentResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResponseData struct {
	// The render flow browser actions status results.
	BrowserActions AgentResponseDataBrowserActions `json:"browser_actions"`
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
	NetworkCapture []AgentResponseDataNetworkCapture `json:"network_capture"`
	// The parsing results extracted from the HTML & network content.
	Parsing AgentResponseDataParsingUnion `json:"parsing"`
	// The list of redirects that occurred during the task.
	Redirects []AgentResponseDataRedirect `json:"redirects"`
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
func (r AgentResponseData) RawJSON() string { return r.JSON.raw }
func (r *AgentResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The render flow browser actions status results.
type AgentResponseDataBrowserActions struct {
	Results []AgentResponseDataBrowserActionsResultUnion `json:"results,required"`
	Success bool                                         `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentResponseDataBrowserActions) RawJSON() string { return r.JSON.raw }
func (r *AgentResponseDataBrowserActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentResponseDataBrowserActionsResultUnion contains all possible properties and
// values from [AgentResponseDataBrowserActionsResultObject],
// [AgentResponseDataBrowserActionsResultObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AgentResponseDataBrowserActionsResultUnion struct {
	// This field is from variant [AgentResponseDataBrowserActionsResultObject].
	Duration float64 `json:"duration"`
	// This field is from variant [AgentResponseDataBrowserActionsResultObject].
	Name string `json:"name"`
	// This field is from variant [AgentResponseDataBrowserActionsResultObject].
	Status string `json:"status"`
	// This field is from variant [AgentResponseDataBrowserActionsResultObject].
	Result any `json:"result"`
	// This field is from variant [AgentResponseDataBrowserActionsResultObject].
	Error string `json:"error"`
	JSON  struct {
		Duration respjson.Field
		Name     respjson.Field
		Status   respjson.Field
		Result   respjson.Field
		Error    respjson.Field
		raw      string
	} `json:"-"`
}

func (u AgentResponseDataBrowserActionsResultUnion) AsAgentResponseDataBrowserActionsResultObject() (v AgentResponseDataBrowserActionsResultObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentResponseDataBrowserActionsResultUnion) AsVariant2() (v AgentResponseDataBrowserActionsResultObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentResponseDataBrowserActionsResultUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentResponseDataBrowserActionsResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResponseDataBrowserActionsResultObject struct {
	Duration float64 `json:"duration,required"`
	Name     string  `json:"name,required"`
	// Any of "no-run", "in-progress", "done", "error", "skipped".
	Status string `json:"status,required"`
	Result any    `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentResponseDataBrowserActionsResultObject) RawJSON() string { return r.JSON.raw }
func (r *AgentResponseDataBrowserActionsResultObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResponseDataNetworkCapture struct {
	Filter       AgentResponseDataNetworkCaptureFilter   `json:"filter,required"`
	Results      []AgentResponseDataNetworkCaptureResult `json:"results,required"`
	ErrorMessage string                                  `json:"errorMessage"`
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
func (r AgentResponseDataNetworkCapture) RawJSON() string { return r.JSON.raw }
func (r *AgentResponseDataNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResponseDataNetworkCaptureFilter struct {
	Validation           bool    `json:"validation,required"`
	WaitForRequestsCount float64 `json:"wait_for_requests_count,required"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method"`
	// Resource type for network capture filtering
	ResourceType                AgentResponseDataNetworkCaptureFilterResourceTypeUnion `json:"resource_type"`
	StatusCode                  AgentResponseDataNetworkCaptureFilterStatusCodeUnion   `json:"status_code"`
	URL                         AgentResponseDataNetworkCaptureFilterURL               `json:"url"`
	WaitForRequestsCountTimeout float64                                                `json:"wait_for_requests_count_timeout"`
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
func (r AgentResponseDataNetworkCaptureFilter) RawJSON() string { return r.JSON.raw }
func (r *AgentResponseDataNetworkCaptureFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentResponseDataNetworkCaptureFilterResourceTypeUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAgentResponseDataNetworkCaptureFilterResourceTypeString
// OfAgentResponseDataNetworkCaptureFilterResourceTypeArrayItemArray]
type AgentResponseDataNetworkCaptureFilterResourceTypeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfAgentResponseDataNetworkCaptureFilterResourceTypeString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfAgentResponseDataNetworkCaptureFilterResourceTypeArrayItemArray []string `json:",inline"`
	JSON                                                              struct {
		OfAgentResponseDataNetworkCaptureFilterResourceTypeString         respjson.Field
		OfAgentResponseDataNetworkCaptureFilterResourceTypeArrayItemArray respjson.Field
		raw                                                               string
	} `json:"-"`
}

func (u AgentResponseDataNetworkCaptureFilterResourceTypeUnion) AsAgentResponseDataNetworkCaptureFilterResourceTypeString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentResponseDataNetworkCaptureFilterResourceTypeUnion) AsAgentResponseDataNetworkCaptureFilterResourceTypeArrayItemArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentResponseDataNetworkCaptureFilterResourceTypeUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentResponseDataNetworkCaptureFilterResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type for network capture filtering
type AgentResponseDataNetworkCaptureFilterResourceTypeString string

const (
	AgentResponseDataNetworkCaptureFilterResourceTypeStringDocument           AgentResponseDataNetworkCaptureFilterResourceTypeString = "document"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringStylesheet         AgentResponseDataNetworkCaptureFilterResourceTypeString = "stylesheet"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringImage              AgentResponseDataNetworkCaptureFilterResourceTypeString = "image"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringMedia              AgentResponseDataNetworkCaptureFilterResourceTypeString = "media"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringFont               AgentResponseDataNetworkCaptureFilterResourceTypeString = "font"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringScript             AgentResponseDataNetworkCaptureFilterResourceTypeString = "script"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringTexttrack          AgentResponseDataNetworkCaptureFilterResourceTypeString = "texttrack"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringXhr                AgentResponseDataNetworkCaptureFilterResourceTypeString = "xhr"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringFetch              AgentResponseDataNetworkCaptureFilterResourceTypeString = "fetch"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringPrefetch           AgentResponseDataNetworkCaptureFilterResourceTypeString = "prefetch"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringEventsource        AgentResponseDataNetworkCaptureFilterResourceTypeString = "eventsource"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringWebsocket          AgentResponseDataNetworkCaptureFilterResourceTypeString = "websocket"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringManifest           AgentResponseDataNetworkCaptureFilterResourceTypeString = "manifest"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringSignedexchange     AgentResponseDataNetworkCaptureFilterResourceTypeString = "signedexchange"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringPing               AgentResponseDataNetworkCaptureFilterResourceTypeString = "ping"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringCspviolationreport AgentResponseDataNetworkCaptureFilterResourceTypeString = "cspviolationreport"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringPreflight          AgentResponseDataNetworkCaptureFilterResourceTypeString = "preflight"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringOther              AgentResponseDataNetworkCaptureFilterResourceTypeString = "other"
	AgentResponseDataNetworkCaptureFilterResourceTypeStringFedcm              AgentResponseDataNetworkCaptureFilterResourceTypeString = "fedcm"
)

// AgentResponseDataNetworkCaptureFilterStatusCodeUnion contains all possible
// properties and values from [float64], [[]float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfFloatArray]
type AgentResponseDataNetworkCaptureFilterStatusCodeUnion struct {
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

func (u AgentResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloatArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentResponseDataNetworkCaptureFilterStatusCodeUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentResponseDataNetworkCaptureFilterStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResponseDataNetworkCaptureFilterURL struct {
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
func (r AgentResponseDataNetworkCaptureFilterURL) RawJSON() string { return r.JSON.raw }
func (r *AgentResponseDataNetworkCaptureFilterURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResponseDataNetworkCaptureResult struct {
	Request  AgentResponseDataNetworkCaptureResultRequest  `json:"request,required"`
	Response AgentResponseDataNetworkCaptureResultResponse `json:"response,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Request     respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentResponseDataNetworkCaptureResult) RawJSON() string { return r.JSON.raw }
func (r *AgentResponseDataNetworkCaptureResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResponseDataNetworkCaptureResultRequest struct {
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
func (r AgentResponseDataNetworkCaptureResultRequest) RawJSON() string { return r.JSON.raw }
func (r *AgentResponseDataNetworkCaptureResultRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResponseDataNetworkCaptureResultResponse struct {
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
func (r AgentResponseDataNetworkCaptureResultResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentResponseDataNetworkCaptureResultResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentResponseDataParsingUnion contains all possible properties and values from
// [AgentResponseDataParsingObject], [AgentResponseDataParsingObject],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAgentResponseDataParsingMapItem]
type AgentResponseDataParsingUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfAgentResponseDataParsingMapItem any `json:",inline"`
	// This field is from variant [AgentResponseDataParsingObject].
	Entities map[string]any `json:"entities"`
	Status   string         `json:"status"`
	// This field is from variant [AgentResponseDataParsingObject].
	Error string `json:"error"`
	JSON  struct {
		OfAgentResponseDataParsingMapItem respjson.Field
		Entities                          respjson.Field
		Status                            respjson.Field
		Error                             respjson.Field
		raw                               string
	} `json:"-"`
}

func (u AgentResponseDataParsingUnion) AsAgentResponseDataParsingObject() (v AgentResponseDataParsingObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentResponseDataParsingUnion) AsVariant2() (v AgentResponseDataParsingObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentResponseDataParsingUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentResponseDataParsingUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentResponseDataParsingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResponseDataParsingObject struct {
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
func (r AgentResponseDataParsingObject) RawJSON() string { return r.JSON.raw }
func (r *AgentResponseDataParsingObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResponseDataRedirect struct {
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
func (r AgentResponseDataRedirect) RawJSON() string { return r.JSON.raw }
func (r *AgentResponseDataRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResponseMetadata struct {
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
	// The identifier of the template used for the query.
	TemplateID string `json:"template_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Driver             respjson.Field
		LocalizationID     respjson.Field
		QueryDuration      respjson.Field
		QueryTime          respjson.Field
		ResponseParameters respjson.Field
		Tag                respjson.Field
		TemplateID         respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *AgentResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the task.
type AgentResponseStatus string

const (
	AgentResponseStatusSuccess   AgentResponseStatus = "success"
	AgentResponseStatusSkipped   AgentResponseStatus = "skipped"
	AgentResponseStatusFatal     AgentResponseStatus = "fatal"
	AgentResponseStatusError     AgentResponseStatus = "error"
	AgentResponseStatusPostponed AgentResponseStatus = "postponed"
	AgentResponseStatusIgnored   AgentResponseStatus = "ignored"
	AgentResponseStatusRejected  AgentResponseStatus = "rejected"
	AgentResponseStatusBlocked   AgentResponseStatus = "blocked"
)

type AgentResponseDebug struct {
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
func (r AgentResponseDebug) RawJSON() string { return r.JSON.raw }
func (r *AgentResponseDebug) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentResponsePaginationUnion contains all possible properties and values from
// [AgentResponsePaginationNextPageParams], [[]AgentResponsePaginationArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAgentResponsePaginationArray]
type AgentResponsePaginationUnion struct {
	// This field will be present if the value is a
	// [[]AgentResponsePaginationArrayItem] instead of an object.
	OfAgentResponsePaginationArray []AgentResponsePaginationArrayItem `json:",inline"`
	// This field is from variant [AgentResponsePaginationNextPageParams].
	NextPageParams map[string]any `json:"next_page_params"`
	JSON           struct {
		OfAgentResponsePaginationArray respjson.Field
		NextPageParams                 respjson.Field
		raw                            string
	} `json:"-"`
}

func (u AgentResponsePaginationUnion) AsAgentResponsePaginationNextPageParams() (v AgentResponsePaginationNextPageParams) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentResponsePaginationUnion) AsAgentResponsePaginationArray() (v []AgentResponsePaginationArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentResponsePaginationUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentResponsePaginationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResponsePaginationNextPageParams struct {
	NextPageParams map[string]any `json:"next_page_params,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentResponsePaginationNextPageParams) RawJSON() string { return r.JSON.raw }
func (r *AgentResponsePaginationNextPageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResponsePaginationArrayItem struct {
	NextPageParams map[string]any `json:"next_page_params,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentResponsePaginationArrayItem) RawJSON() string { return r.JSON.raw }
func (r *AgentResponsePaginationArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlResponse struct {
	ID           string                      `json:"id,required" format:"uuid"`
	AccountName  string                      `json:"account_name,required"`
	CrawlOptions CrawlResponseCrawlOptions   `json:"crawl_options,required"`
	CreatedAt    CrawlResponseCreatedAtUnion `json:"created_at,required"`
	// Any of "queued", "running", "succeeded", "failed", "canceled".
	Status         CrawlResponseStatus           `json:"status,required"`
	UpdatedAt      CrawlResponseUpdatedAtUnion   `json:"updated_at,required"`
	URL            string                        `json:"url,required" format:"uri"`
	Completed      float64                       `json:"completed"`
	CompletedAt    CrawlResponseCompletedAtUnion `json:"completed_at,nullable"`
	EncryptedToken string                        `json:"encrypted_token,nullable"`
	ExtractOptions map[string]any                `json:"extract_options,nullable"`
	Failed         float64                       `json:"failed"`
	Name           string                        `json:"name,nullable"`
	Pending        float64                       `json:"pending"`
	Tasks          []CrawlResponseTask           `json:"tasks"`
	Total          float64                       `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		AccountName    respjson.Field
		CrawlOptions   respjson.Field
		CreatedAt      respjson.Field
		Status         respjson.Field
		UpdatedAt      respjson.Field
		URL            respjson.Field
		Completed      respjson.Field
		CompletedAt    respjson.Field
		EncryptedToken respjson.Field
		ExtractOptions respjson.Field
		Failed         respjson.Field
		Name           respjson.Field
		Pending        respjson.Field
		Tasks          respjson.Field
		Total          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlResponse) RawJSON() string { return r.JSON.raw }
func (r *CrawlResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlResponseCrawlOptions struct {
	AllowExternalLinks    bool  `json:"allow_external_links,required"`
	AllowSubdomains       bool  `json:"allow_subdomains,required"`
	CrawlEntireDomain     bool  `json:"crawl_entire_domain,required"`
	IgnoreQueryParameters bool  `json:"ignore_query_parameters,required"`
	Limit                 int64 `json:"limit,required"`
	MaxDiscoveryDepth     int64 `json:"max_discovery_depth,required"`
	// Any of "skip", "include", "only".
	Sitemap      string                                 `json:"sitemap,required"`
	Callback     CrawlResponseCrawlOptionsCallbackUnion `json:"callback" format:"uri"`
	ExcludePaths []string                               `json:"exclude_paths"`
	IncludePaths []string                               `json:"include_paths"`
	ExtraFields  map[string]any                         `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowExternalLinks    respjson.Field
		AllowSubdomains       respjson.Field
		CrawlEntireDomain     respjson.Field
		IgnoreQueryParameters respjson.Field
		Limit                 respjson.Field
		MaxDiscoveryDepth     respjson.Field
		Sitemap               respjson.Field
		Callback              respjson.Field
		ExcludePaths          respjson.Field
		IncludePaths          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlResponseCrawlOptions) RawJSON() string { return r.JSON.raw }
func (r *CrawlResponseCrawlOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlResponseCrawlOptionsCallbackUnion contains all possible properties and
// values from [CrawlResponseCrawlOptionsCallbackObject], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type CrawlResponseCrawlOptionsCallbackUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [CrawlResponseCrawlOptionsCallbackObject].
	URL string `json:"url"`
	// This field is from variant [CrawlResponseCrawlOptionsCallbackObject].
	Events []string `json:"events"`
	// This field is from variant [CrawlResponseCrawlOptionsCallbackObject].
	Headers map[string]string `json:"headers"`
	// This field is from variant [CrawlResponseCrawlOptionsCallbackObject].
	Metadata map[string]any `json:"metadata"`
	JSON     struct {
		OfString respjson.Field
		URL      respjson.Field
		Events   respjson.Field
		Headers  respjson.Field
		Metadata respjson.Field
		raw      string
	} `json:"-"`
}

func (u CrawlResponseCrawlOptionsCallbackUnion) AsCrawlResponseCrawlOptionsCallbackObject() (v CrawlResponseCrawlOptionsCallbackObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlResponseCrawlOptionsCallbackUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlResponseCrawlOptionsCallbackUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlResponseCrawlOptionsCallbackUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlResponseCrawlOptionsCallbackObject struct {
	URL string `json:"url,required" format:"uri"`
	// Any of "started", "page", "completed", "failed".
	Events   []string          `json:"events"`
	Headers  map[string]string `json:"headers"`
	Metadata map[string]any    `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		Events      respjson.Field
		Headers     respjson.Field
		Metadata    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlResponseCrawlOptionsCallbackObject) RawJSON() string { return r.JSON.raw }
func (r *CrawlResponseCrawlOptionsCallbackObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlResponseCreatedAtUnion contains all possible properties and values from
// [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCrawlResponseCreatedAtMapItem]
type CrawlResponseCreatedAtUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfCrawlResponseCreatedAtMapItem any `json:",inline"`
	JSON                            struct {
		OfString                        respjson.Field
		OfCrawlResponseCreatedAtMapItem respjson.Field
		raw                             string
	} `json:"-"`
}

func (u CrawlResponseCreatedAtUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlResponseCreatedAtUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlResponseCreatedAtUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlResponseCreatedAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlResponseStatus string

const (
	CrawlResponseStatusQueued    CrawlResponseStatus = "queued"
	CrawlResponseStatusRunning   CrawlResponseStatus = "running"
	CrawlResponseStatusSucceeded CrawlResponseStatus = "succeeded"
	CrawlResponseStatusFailed    CrawlResponseStatus = "failed"
	CrawlResponseStatusCanceled  CrawlResponseStatus = "canceled"
)

// CrawlResponseUpdatedAtUnion contains all possible properties and values from
// [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCrawlResponseUpdatedAtMapItem]
type CrawlResponseUpdatedAtUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfCrawlResponseUpdatedAtMapItem any `json:",inline"`
	JSON                            struct {
		OfString                        respjson.Field
		OfCrawlResponseUpdatedAtMapItem respjson.Field
		raw                             string
	} `json:"-"`
}

func (u CrawlResponseUpdatedAtUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlResponseUpdatedAtUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlResponseUpdatedAtUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlResponseUpdatedAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlResponseCompletedAtUnion contains all possible properties and values from
// [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCrawlResponseCompletedAtMapItem]
type CrawlResponseCompletedAtUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfCrawlResponseCompletedAtMapItem any `json:",inline"`
	JSON                              struct {
		OfString                          respjson.Field
		OfCrawlResponseCompletedAtMapItem respjson.Field
		raw                               string
	} `json:"-"`
}

func (u CrawlResponseCompletedAtUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlResponseCompletedAtUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlResponseCompletedAtUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlResponseCompletedAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlResponseTask struct {
	CrawlID string `json:"crawl_id,required" format:"uuid"`
	// Any of "pending", "completed", "failed".
	Status      string                          `json:"status,required"`
	WebitTaskID string                          `json:"webit_task_id,required"`
	CreatedAt   CrawlResponseTaskCreatedAtUnion `json:"created_at"`
	UpdatedAt   CrawlResponseTaskUpdatedAtUnion `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CrawlID     respjson.Field
		Status      respjson.Field
		WebitTaskID respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlResponseTask) RawJSON() string { return r.JSON.raw }
func (r *CrawlResponseTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlResponseTaskCreatedAtUnion contains all possible properties and values from
// [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCrawlResponseTaskCreatedAtMapItem]
type CrawlResponseTaskCreatedAtUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfCrawlResponseTaskCreatedAtMapItem any `json:",inline"`
	JSON                                struct {
		OfString                            respjson.Field
		OfCrawlResponseTaskCreatedAtMapItem respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u CrawlResponseTaskCreatedAtUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlResponseTaskCreatedAtUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlResponseTaskCreatedAtUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlResponseTaskCreatedAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlResponseTaskUpdatedAtUnion contains all possible properties and values from
// [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCrawlResponseTaskUpdatedAtMapItem]
type CrawlResponseTaskUpdatedAtUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfCrawlResponseTaskUpdatedAtMapItem any `json:",inline"`
	JSON                                struct {
		OfString                            respjson.Field
		OfCrawlResponseTaskUpdatedAtMapItem respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u CrawlResponseTaskUpdatedAtUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlResponseTaskUpdatedAtUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlResponseTaskUpdatedAtUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlResponseTaskUpdatedAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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
	// The render flow browser actions status results.
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

// The render flow browser actions status results.
type ExtractResponseDataBrowserActions struct {
	Results []ExtractResponseDataBrowserActionsResultUnion `json:"results,required"`
	Success bool                                           `json:"success,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractResponseDataBrowserActions) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseDataBrowserActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractResponseDataBrowserActionsResultUnion contains all possible properties
// and values from [ExtractResponseDataBrowserActionsResultObject],
// [ExtractResponseDataBrowserActionsResultObject].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type ExtractResponseDataBrowserActionsResultUnion struct {
	// This field is from variant [ExtractResponseDataBrowserActionsResultObject].
	Duration float64 `json:"duration"`
	// This field is from variant [ExtractResponseDataBrowserActionsResultObject].
	Name string `json:"name"`
	// This field is from variant [ExtractResponseDataBrowserActionsResultObject].
	Status string `json:"status"`
	// This field is from variant [ExtractResponseDataBrowserActionsResultObject].
	Result any `json:"result"`
	// This field is from variant [ExtractResponseDataBrowserActionsResultObject].
	Error string `json:"error"`
	JSON  struct {
		Duration respjson.Field
		Name     respjson.Field
		Status   respjson.Field
		Result   respjson.Field
		Error    respjson.Field
		raw      string
	} `json:"-"`
}

func (u ExtractResponseDataBrowserActionsResultUnion) AsExtractResponseDataBrowserActionsResultObject() (v ExtractResponseDataBrowserActionsResultObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractResponseDataBrowserActionsResultUnion) AsVariant2() (v ExtractResponseDataBrowserActionsResultObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractResponseDataBrowserActionsResultUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractResponseDataBrowserActionsResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractResponseDataBrowserActionsResultObject struct {
	Duration float64 `json:"duration,required"`
	Name     string  `json:"name,required"`
	// Any of "no-run", "in-progress", "done", "error", "skipped".
	Status string `json:"status,required"`
	Result any    `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractResponseDataBrowserActionsResultObject) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseDataBrowserActionsResultObject) UnmarshalJSON(data []byte) error {
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
// [ExtractResponseDataParsingObject], [ExtractResponseDataParsingObject],
// [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractResponseDataParsingMapItem]
type ExtractResponseDataParsingUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExtractResponseDataParsingMapItem any `json:",inline"`
	// This field is from variant [ExtractResponseDataParsingObject].
	Entities map[string]any `json:"entities"`
	Status   string         `json:"status"`
	// This field is from variant [ExtractResponseDataParsingObject].
	Error string `json:"error"`
	JSON  struct {
		OfExtractResponseDataParsingMapItem respjson.Field
		Entities                            respjson.Field
		Status                              respjson.Field
		Error                               respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u ExtractResponseDataParsingUnion) AsExtractResponseDataParsingObject() (v ExtractResponseDataParsingObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractResponseDataParsingUnion) AsVariant2() (v ExtractResponseDataParsingObject) {
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

type ExtractResponseDataParsingObject struct {
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
func (r ExtractResponseDataParsingObject) RawJSON() string { return r.JSON.raw }
func (r *ExtractResponseDataParsingObject) UnmarshalJSON(data []byte) error {
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
	// The identifier of the template used for the query.
	TemplateID string `json:"template_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Driver             respjson.Field
		LocalizationID     respjson.Field
		QueryDuration      respjson.Field
		QueryTime          respjson.Field
		ResponseParameters respjson.Field
		Tag                respjson.Field
		TemplateID         respjson.Field
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

// Response model from SearchService with results and optional LLM answer.
//
// Note: request_id is always a valid UUID generated internally by the middleware,
// so no validation is needed.
type SearchResponse struct {
	// Unique identifier for this request (UUID)
	RequestID string                 `json:"request_id,required"`
	Results   []SearchResponseResult `json:"results,required"`
	// Number of results returned
	TotalResults int64  `json:"total_results,required"`
	Answer       string `json:"answer,nullable"`
	// Citations mapping citation markers to result indices
	AnswerCitations []SearchResponseAnswerCitation `json:"answer_citations,nullable"`
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
// platform-specific data in extra_fields and typed metadata.
type SearchResponseResult struct {
	Content     string `json:"content,required"`
	Description string `json:"description,required"`
	// Metadata for SERP-based search results (general, news, location).
	Metadata    SearchResponseResultMetadataUnion `json:"metadata,required"`
	Title       string                            `json:"title,required"`
	URL         string                            `json:"url,required"`
	ExtraFields map[string]any                    `json:"extra_fields,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Description respjson.Field
		Metadata    respjson.Field
		Title       respjson.Field
		URL         respjson.Field
		ExtraFields respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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
	Country    string `json:"country,required"`
	EntityType string `json:"entity_type,required"`
	Locale     string `json:"locale,required"`
	Position   int64  `json:"position,required"`
	Driver     string `json:"driver,nullable"`
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
	AgentName string `json:"agent_name,required"`
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
	Marker int64 `json:"marker,required"`
	// Zero-based index into the results array
	ResultIndex int64 `json:"result_index,required"`
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

type AgentParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Request
	// body for executing a WSA
	OfExtractTemplateBody *AgentParamsBodyExtractTemplateBody `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Request
	// body for executing a WSA
	OfAgentBody *AgentParamsBodyAgentBody `json:",inline"`

	paramObj
}

func (u AgentParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractTemplateBody, u.OfAgentBody)
}
func (r *AgentParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for executing a WSA
//
// The properties Params, Template are required.
type AgentParamsBodyExtractTemplateBody struct {
	Params       map[string]any  `json:"params,omitzero,required"`
	Template     string          `json:"template,required"`
	Localization param.Opt[bool] `json:"localization,omitzero"`
	paramObj
}

func (r AgentParamsBodyExtractTemplateBody) MarshalJSON() (data []byte, err error) {
	type shadow AgentParamsBodyExtractTemplateBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentParamsBodyExtractTemplateBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for executing a WSA
//
// The properties Agent, Params are required.
type AgentParamsBodyAgentBody struct {
	Agent        string          `json:"agent,required"`
	Params       map[string]any  `json:"params,omitzero,required"`
	Localization param.Opt[bool] `json:"localization,omitzero"`
	paramObj
}

func (r AgentParamsBodyAgentBody) MarshalJSON() (data []byte, err error) {
	type shadow AgentParamsBodyAgentBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentParamsBodyAgentBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlParams struct {
	// Url to crawl.
	URL string `json:"url,required"`
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
	Callback CrawlParamsCallbackUnion `json:"callback,omitzero" format:"uri"`
	// URL pathname regex patterns that exclude matching URLs from the crawl.
	ExcludePaths   []string                  `json:"exclude_paths,omitzero"`
	ExtractOptions CrawlParamsExtractOptions `json:"extract_options,omitzero"`
	// URL pathname regex patterns that include matching URLs in the crawl.
	IncludePaths []string `json:"include_paths,omitzero"`
	// Sitemap and other methods will be used together to find URLs.
	//
	// Any of "skip", "include", "only".
	Sitemap CrawlParamsSitemap `json:"sitemap,omitzero"`
	paramObj
}

func (r CrawlParams) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsCallbackUnion struct {
	OfCrawlsCallbackObject *CrawlParamsCallbackObject `json:",omitzero,inline"`
	OfString               param.Opt[string]          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsCallbackUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsCallbackObject, u.OfString)
}
func (u *CrawlParamsCallbackUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsCallbackUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsCallbackObject) {
		return u.OfCrawlsCallbackObject
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// The property URL is required.
type CrawlParamsCallbackObject struct {
	URL string `json:"url,required" format:"uri"`
	// Any of "started", "page", "completed", "failed".
	Events   []string          `json:"events,omitzero"`
	Headers  map[string]string `json:"headers,omitzero"`
	Metadata map[string]any    `json:"metadata,omitzero"`
	paramObj
}

func (r CrawlParamsCallbackObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsCallbackObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsCallbackObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlParamsExtractOptions struct {
	// City for geolocation
	City param.Opt[string] `json:"city,omitzero"`
	// Client-side timeout in milliseconds
	ClientTimeout param.Opt[float64] `json:"client_timeout,omitzero"`
	// Whether to automatically handle cookie consent headers
	ConsentHeader param.Opt[bool] `json:"consent_header,omitzero"`
	// Whether to disable IP address validation
	DisableIPCheck param.Opt[bool] `json:"disable_ip_check,omitzero"`
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
	// Whether to save the userbrowser session for reuse
	SaveUserbrowser param.Opt[bool] `json:"save_userbrowser,omitzero"`
	// Whether to skip userbrowser creation template processing
	SkipUbct param.Opt[bool] `json:"skip_ubct,omitzero"`
	// User-defined tag for request identification
	Tag param.Opt[string] `json:"tag,omitzero"`
	// Type of query or scraping template
	Type param.Opt[string] `json:"type,omitzero"`
	// Target URL to scrape
	URL param.Opt[string] `json:"url,omitzero" format:"uri"`
	// Browser type to emulate
	Browser CrawlParamsExtractOptionsBrowserUnion `json:"browser,omitzero"`
	// Browser cookies as array of cookie objects
	Cookies CrawlParamsExtractOptionsCookiesUnion `json:"cookies,omitzero"`
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
	Country CrawlParamsExtractOptionsCountry `json:"country,omitzero"`
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
	Headers map[string]CrawlParamsExtractOptionsHeaderUnion `json:"headers,omitzero"`
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
	Locale CrawlParamsExtractOptionsLocale `json:"locale,omitzero"`
	// Structured metadata about the request execution context
	Metadata CrawlParamsExtractOptionsMetadata `json:"metadata,omitzero"`
	// HTTP method for the request
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method,omitzero"`
	// Native execution mode
	//
	// Any of "requester", "apm", "direct".
	NativeMode string `json:"native_mode,omitzero"`
	// Filters for capturing network traffic
	NetworkCapture []CrawlParamsExtractOptionsNetworkCapture `json:"network_capture,omitzero"`
	// Operating system to emulate
	//
	// Any of "windows", "mac os", "linux", "android", "ios".
	Os string `json:"os,omitzero"`
	// Configuration options for parsing behavior
	ParseOptions CrawlParamsExtractOptionsParseOptions `json:"parse_options,omitzero"`
	// Custom parser configuration as a key-value map
	Parser CrawlParamsExtractOptionsParserUnion `json:"parser,omitzero"`
	// Proxy provider to use for the request
	//
	// Any of "brightdata", "oxylabs", "smartproxy", "proxit", "proxit_preprod",
	// "local", "rayobyte", "always", "oculusproxies", "froxy", "packetstream",
	// "911proxy", "direct911proxy", "thesocialproxy", "thesocialproxy2", "nimble-isp",
	// "nimble-isp-mobile", "proxit-linux", "proxit-macos", "proxit-windows",
	// "proxit-rental", "ipfoxy", "brightup", "research".
	ProxyProvider CrawlParamsExtractOptionsProxyProvider `json:"proxy_provider,omitzero"`
	// Weighted distribution of proxy providers
	ProxyProviders map[string]float64 `json:"proxy_providers,omitzero"`
	// Query template configuration for structured data extraction
	QueryTemplate CrawlParamsExtractOptionsQueryTemplate `json:"query_template,omitzero"`
	// Referrer policy for the request
	//
	// Any of "random", "no-referer", "same-origin", "google", "bing", "facebook",
	// "twitter", "instagram".
	ReferrerType CrawlParamsExtractOptionsReferrerType `json:"referrer_type,omitzero"`
	// Array of actions to perform during browser rendering
	RenderFlow    []map[string]any                       `json:"render_flow,omitzero"`
	RenderOptions CrawlParamsExtractOptionsRenderOptions `json:"render_options,omitzero"`
	Session       CrawlParamsExtractOptionsSession       `json:"session,omitzero"`
	// Skills or capabilities required for the request
	Skill CrawlParamsExtractOptionsSkillUnion `json:"skill,omitzero"`
	// US state for geolocation (only valid when country is US)
	//
	// Any of "AL", "AK", "AS", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA",
	// "GU", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI",
	// "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "MP",
	// "OH", "OK", "OR", "PA", "PR", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA",
	// "VI", "WA", "WV", "WI", "WY".
	State string `json:"state,omitzero"`
	// Userbrowser creation template configuration
	Template CrawlParamsExtractOptionsTemplate `json:"template,omitzero"`
	// Pre-rendered userbrowser creation template configuration
	UserbrowserCreationTemplateRendered CrawlParamsExtractOptionsUserbrowserCreationTemplateRendered `json:"userbrowser_creation_template_rendered,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptions) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlParamsExtractOptions](
		"device", "desktop", "mobile", "tablet",
	)
	apijson.RegisterFieldValidator[CrawlParamsExtractOptions](
		"driver", "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro",
	)
	apijson.RegisterFieldValidator[CrawlParamsExtractOptions](
		"format", "json", "html", "csv", "raw", "json-lines", "markdown",
	)
	apijson.RegisterFieldValidator[CrawlParamsExtractOptions](
		"method", "GET", "POST", "PUT", "PATCH", "DELETE",
	)
	apijson.RegisterFieldValidator[CrawlParamsExtractOptions](
		"native_mode", "requester", "apm", "direct",
	)
	apijson.RegisterFieldValidator[CrawlParamsExtractOptions](
		"os", "windows", "mac os", "linux", "android", "ios",
	)
	apijson.RegisterFieldValidator[CrawlParamsExtractOptions](
		"state", "AL", "AK", "AS", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA", "GU", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "MP", "OH", "OK", "OR", "PA", "PR", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "VI", "WA", "WV", "WI", "WY",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserString)
	OfCrawlsExtractOptionsBrowserString param.Opt[string]                       `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserObject *CrawlParamsExtractOptionsBrowserObject `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserString, u.OfCrawlsExtractOptionsBrowserObject)
}
func (u *CrawlParamsExtractOptionsBrowserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserString) {
		return &u.OfCrawlsExtractOptionsBrowserString
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserObject) {
		return u.OfCrawlsExtractOptionsBrowserObject
	}
	return nil
}

// Browser type to emulate
type CrawlParamsExtractOptionsBrowserString string

const (
	CrawlParamsExtractOptionsBrowserStringChrome  CrawlParamsExtractOptionsBrowserString = "chrome"
	CrawlParamsExtractOptionsBrowserStringFirefox CrawlParamsExtractOptionsBrowserString = "firefox"
)

// The property Name is required.
type CrawlParamsExtractOptionsBrowserObject struct {
	// Any of "chrome", "firefox".
	Name string `json:"name,omitzero,required"`
	// Specific browser version to emulate
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsBrowserObject](
		"name", "chrome", "firefox",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsCookiesUnion struct {
	OfCrawlsExtractOptionsCookiesArray []CrawlParamsExtractOptionsCookiesArrayItem `json:",omitzero,inline"`
	OfString                           param.Opt[string]                           `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsCookiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsCookiesArray, u.OfString)
}
func (u *CrawlParamsExtractOptionsCookiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsCookiesUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsCookiesArray) {
		return &u.OfCrawlsExtractOptionsCookiesArray
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type CrawlParamsExtractOptionsCookiesArrayItem struct {
	Creation      param.Opt[string]                                    `json:"creation,omitzero"`
	Domain        param.Opt[string]                                    `json:"domain,omitzero"`
	HostOnly      param.Opt[bool]                                      `json:"hostOnly,omitzero"`
	HTTPOnly      param.Opt[bool]                                      `json:"httpOnly,omitzero"`
	LastAccessed  param.Opt[string]                                    `json:"lastAccessed,omitzero"`
	Path          param.Opt[string]                                    `json:"path,omitzero"`
	PathIsDefault param.Opt[bool]                                      `json:"pathIsDefault,omitzero"`
	Expires       param.Opt[string]                                    `json:"expires,omitzero"`
	Name          param.Opt[string]                                    `json:"name,omitzero"`
	Secure        param.Opt[bool]                                      `json:"secure,omitzero"`
	Value         param.Opt[string]                                    `json:"value,omitzero"`
	Extensions    []string                                             `json:"extensions,omitzero"`
	MaxAge        CrawlParamsExtractOptionsCookiesArrayItemMaxAgeUnion `json:"maxAge,omitzero"`
	// Any of "strict", "lax", "none".
	SameSite    string         `json:"sameSite,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r CrawlParamsExtractOptionsCookiesArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsCookiesArrayItem
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *CrawlParamsExtractOptionsCookiesArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsCookiesArrayItem](
		"sameSite", "strict", "lax", "none",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsCookiesArrayItemMaxAgeUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsCookiesArrayItemMaxAgeString)
	OfCrawlsExtractOptionsCookiesArrayItemMaxAgeString param.Opt[CrawlParamsExtractOptionsCookiesArrayItemMaxAgeString] `json:",omitzero,inline"`
	OfFloat                                            param.Opt[float64]                                               `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsCookiesArrayItemMaxAgeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsCookiesArrayItemMaxAgeString, u.OfFloat)
}
func (u *CrawlParamsExtractOptionsCookiesArrayItemMaxAgeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsCookiesArrayItemMaxAgeUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsCookiesArrayItemMaxAgeString) {
		return &u.OfCrawlsExtractOptionsCookiesArrayItemMaxAgeString
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type CrawlParamsExtractOptionsCookiesArrayItemMaxAgeString string

const (
	CrawlParamsExtractOptionsCookiesArrayItemMaxAgeStringInfinity      CrawlParamsExtractOptionsCookiesArrayItemMaxAgeString = "Infinity"
	CrawlParamsExtractOptionsCookiesArrayItemMaxAgeStringMinusInfinity CrawlParamsExtractOptionsCookiesArrayItemMaxAgeString = "-Infinity"
)

// Country code for geolocation and proxy selection
type CrawlParamsExtractOptionsCountry string

const (
	CrawlParamsExtractOptionsCountryAd  CrawlParamsExtractOptionsCountry = "AD"
	CrawlParamsExtractOptionsCountryAe  CrawlParamsExtractOptionsCountry = "AE"
	CrawlParamsExtractOptionsCountryAf  CrawlParamsExtractOptionsCountry = "AF"
	CrawlParamsExtractOptionsCountryAg  CrawlParamsExtractOptionsCountry = "AG"
	CrawlParamsExtractOptionsCountryAI  CrawlParamsExtractOptionsCountry = "AI"
	CrawlParamsExtractOptionsCountryAl  CrawlParamsExtractOptionsCountry = "AL"
	CrawlParamsExtractOptionsCountryAm  CrawlParamsExtractOptionsCountry = "AM"
	CrawlParamsExtractOptionsCountryAo  CrawlParamsExtractOptionsCountry = "AO"
	CrawlParamsExtractOptionsCountryAq  CrawlParamsExtractOptionsCountry = "AQ"
	CrawlParamsExtractOptionsCountryAr  CrawlParamsExtractOptionsCountry = "AR"
	CrawlParamsExtractOptionsCountryAs  CrawlParamsExtractOptionsCountry = "AS"
	CrawlParamsExtractOptionsCountryAt  CrawlParamsExtractOptionsCountry = "AT"
	CrawlParamsExtractOptionsCountryAu  CrawlParamsExtractOptionsCountry = "AU"
	CrawlParamsExtractOptionsCountryAw  CrawlParamsExtractOptionsCountry = "AW"
	CrawlParamsExtractOptionsCountryAx  CrawlParamsExtractOptionsCountry = "AX"
	CrawlParamsExtractOptionsCountryAz  CrawlParamsExtractOptionsCountry = "AZ"
	CrawlParamsExtractOptionsCountryBa  CrawlParamsExtractOptionsCountry = "BA"
	CrawlParamsExtractOptionsCountryBb  CrawlParamsExtractOptionsCountry = "BB"
	CrawlParamsExtractOptionsCountryBd  CrawlParamsExtractOptionsCountry = "BD"
	CrawlParamsExtractOptionsCountryBe  CrawlParamsExtractOptionsCountry = "BE"
	CrawlParamsExtractOptionsCountryBf  CrawlParamsExtractOptionsCountry = "BF"
	CrawlParamsExtractOptionsCountryBg  CrawlParamsExtractOptionsCountry = "BG"
	CrawlParamsExtractOptionsCountryBh  CrawlParamsExtractOptionsCountry = "BH"
	CrawlParamsExtractOptionsCountryBi  CrawlParamsExtractOptionsCountry = "BI"
	CrawlParamsExtractOptionsCountryBj  CrawlParamsExtractOptionsCountry = "BJ"
	CrawlParamsExtractOptionsCountryBl  CrawlParamsExtractOptionsCountry = "BL"
	CrawlParamsExtractOptionsCountryBm  CrawlParamsExtractOptionsCountry = "BM"
	CrawlParamsExtractOptionsCountryBn  CrawlParamsExtractOptionsCountry = "BN"
	CrawlParamsExtractOptionsCountryBo  CrawlParamsExtractOptionsCountry = "BO"
	CrawlParamsExtractOptionsCountryBq  CrawlParamsExtractOptionsCountry = "BQ"
	CrawlParamsExtractOptionsCountryBr  CrawlParamsExtractOptionsCountry = "BR"
	CrawlParamsExtractOptionsCountryBs  CrawlParamsExtractOptionsCountry = "BS"
	CrawlParamsExtractOptionsCountryBt  CrawlParamsExtractOptionsCountry = "BT"
	CrawlParamsExtractOptionsCountryBv  CrawlParamsExtractOptionsCountry = "BV"
	CrawlParamsExtractOptionsCountryBw  CrawlParamsExtractOptionsCountry = "BW"
	CrawlParamsExtractOptionsCountryBy  CrawlParamsExtractOptionsCountry = "BY"
	CrawlParamsExtractOptionsCountryBz  CrawlParamsExtractOptionsCountry = "BZ"
	CrawlParamsExtractOptionsCountryCa  CrawlParamsExtractOptionsCountry = "CA"
	CrawlParamsExtractOptionsCountryCc  CrawlParamsExtractOptionsCountry = "CC"
	CrawlParamsExtractOptionsCountryCd  CrawlParamsExtractOptionsCountry = "CD"
	CrawlParamsExtractOptionsCountryCf  CrawlParamsExtractOptionsCountry = "CF"
	CrawlParamsExtractOptionsCountryCg  CrawlParamsExtractOptionsCountry = "CG"
	CrawlParamsExtractOptionsCountryCh  CrawlParamsExtractOptionsCountry = "CH"
	CrawlParamsExtractOptionsCountryCi  CrawlParamsExtractOptionsCountry = "CI"
	CrawlParamsExtractOptionsCountryCk  CrawlParamsExtractOptionsCountry = "CK"
	CrawlParamsExtractOptionsCountryCl  CrawlParamsExtractOptionsCountry = "CL"
	CrawlParamsExtractOptionsCountryCm  CrawlParamsExtractOptionsCountry = "CM"
	CrawlParamsExtractOptionsCountryCn  CrawlParamsExtractOptionsCountry = "CN"
	CrawlParamsExtractOptionsCountryCo  CrawlParamsExtractOptionsCountry = "CO"
	CrawlParamsExtractOptionsCountryCr  CrawlParamsExtractOptionsCountry = "CR"
	CrawlParamsExtractOptionsCountryCu  CrawlParamsExtractOptionsCountry = "CU"
	CrawlParamsExtractOptionsCountryCv  CrawlParamsExtractOptionsCountry = "CV"
	CrawlParamsExtractOptionsCountryCw  CrawlParamsExtractOptionsCountry = "CW"
	CrawlParamsExtractOptionsCountryCx  CrawlParamsExtractOptionsCountry = "CX"
	CrawlParamsExtractOptionsCountryCy  CrawlParamsExtractOptionsCountry = "CY"
	CrawlParamsExtractOptionsCountryCz  CrawlParamsExtractOptionsCountry = "CZ"
	CrawlParamsExtractOptionsCountryDe  CrawlParamsExtractOptionsCountry = "DE"
	CrawlParamsExtractOptionsCountryDj  CrawlParamsExtractOptionsCountry = "DJ"
	CrawlParamsExtractOptionsCountryDk  CrawlParamsExtractOptionsCountry = "DK"
	CrawlParamsExtractOptionsCountryDm  CrawlParamsExtractOptionsCountry = "DM"
	CrawlParamsExtractOptionsCountryDo  CrawlParamsExtractOptionsCountry = "DO"
	CrawlParamsExtractOptionsCountryDz  CrawlParamsExtractOptionsCountry = "DZ"
	CrawlParamsExtractOptionsCountryEc  CrawlParamsExtractOptionsCountry = "EC"
	CrawlParamsExtractOptionsCountryEe  CrawlParamsExtractOptionsCountry = "EE"
	CrawlParamsExtractOptionsCountryEg  CrawlParamsExtractOptionsCountry = "EG"
	CrawlParamsExtractOptionsCountryEh  CrawlParamsExtractOptionsCountry = "EH"
	CrawlParamsExtractOptionsCountryEr  CrawlParamsExtractOptionsCountry = "ER"
	CrawlParamsExtractOptionsCountryEs  CrawlParamsExtractOptionsCountry = "ES"
	CrawlParamsExtractOptionsCountryEt  CrawlParamsExtractOptionsCountry = "ET"
	CrawlParamsExtractOptionsCountryFi  CrawlParamsExtractOptionsCountry = "FI"
	CrawlParamsExtractOptionsCountryFj  CrawlParamsExtractOptionsCountry = "FJ"
	CrawlParamsExtractOptionsCountryFk  CrawlParamsExtractOptionsCountry = "FK"
	CrawlParamsExtractOptionsCountryFm  CrawlParamsExtractOptionsCountry = "FM"
	CrawlParamsExtractOptionsCountryFo  CrawlParamsExtractOptionsCountry = "FO"
	CrawlParamsExtractOptionsCountryFr  CrawlParamsExtractOptionsCountry = "FR"
	CrawlParamsExtractOptionsCountryGa  CrawlParamsExtractOptionsCountry = "GA"
	CrawlParamsExtractOptionsCountryGB  CrawlParamsExtractOptionsCountry = "GB"
	CrawlParamsExtractOptionsCountryGd  CrawlParamsExtractOptionsCountry = "GD"
	CrawlParamsExtractOptionsCountryGe  CrawlParamsExtractOptionsCountry = "GE"
	CrawlParamsExtractOptionsCountryGf  CrawlParamsExtractOptionsCountry = "GF"
	CrawlParamsExtractOptionsCountryGg  CrawlParamsExtractOptionsCountry = "GG"
	CrawlParamsExtractOptionsCountryGh  CrawlParamsExtractOptionsCountry = "GH"
	CrawlParamsExtractOptionsCountryGi  CrawlParamsExtractOptionsCountry = "GI"
	CrawlParamsExtractOptionsCountryGl  CrawlParamsExtractOptionsCountry = "GL"
	CrawlParamsExtractOptionsCountryGm  CrawlParamsExtractOptionsCountry = "GM"
	CrawlParamsExtractOptionsCountryGn  CrawlParamsExtractOptionsCountry = "GN"
	CrawlParamsExtractOptionsCountryGp  CrawlParamsExtractOptionsCountry = "GP"
	CrawlParamsExtractOptionsCountryGq  CrawlParamsExtractOptionsCountry = "GQ"
	CrawlParamsExtractOptionsCountryGr  CrawlParamsExtractOptionsCountry = "GR"
	CrawlParamsExtractOptionsCountryGs  CrawlParamsExtractOptionsCountry = "GS"
	CrawlParamsExtractOptionsCountryGt  CrawlParamsExtractOptionsCountry = "GT"
	CrawlParamsExtractOptionsCountryGu  CrawlParamsExtractOptionsCountry = "GU"
	CrawlParamsExtractOptionsCountryGw  CrawlParamsExtractOptionsCountry = "GW"
	CrawlParamsExtractOptionsCountryGy  CrawlParamsExtractOptionsCountry = "GY"
	CrawlParamsExtractOptionsCountryHk  CrawlParamsExtractOptionsCountry = "HK"
	CrawlParamsExtractOptionsCountryHm  CrawlParamsExtractOptionsCountry = "HM"
	CrawlParamsExtractOptionsCountryHn  CrawlParamsExtractOptionsCountry = "HN"
	CrawlParamsExtractOptionsCountryHr  CrawlParamsExtractOptionsCountry = "HR"
	CrawlParamsExtractOptionsCountryHt  CrawlParamsExtractOptionsCountry = "HT"
	CrawlParamsExtractOptionsCountryHu  CrawlParamsExtractOptionsCountry = "HU"
	CrawlParamsExtractOptionsCountryID  CrawlParamsExtractOptionsCountry = "ID"
	CrawlParamsExtractOptionsCountryIe  CrawlParamsExtractOptionsCountry = "IE"
	CrawlParamsExtractOptionsCountryIl  CrawlParamsExtractOptionsCountry = "IL"
	CrawlParamsExtractOptionsCountryIm  CrawlParamsExtractOptionsCountry = "IM"
	CrawlParamsExtractOptionsCountryIn  CrawlParamsExtractOptionsCountry = "IN"
	CrawlParamsExtractOptionsCountryIo  CrawlParamsExtractOptionsCountry = "IO"
	CrawlParamsExtractOptionsCountryIq  CrawlParamsExtractOptionsCountry = "IQ"
	CrawlParamsExtractOptionsCountryIr  CrawlParamsExtractOptionsCountry = "IR"
	CrawlParamsExtractOptionsCountryIs  CrawlParamsExtractOptionsCountry = "IS"
	CrawlParamsExtractOptionsCountryIt  CrawlParamsExtractOptionsCountry = "IT"
	CrawlParamsExtractOptionsCountryJe  CrawlParamsExtractOptionsCountry = "JE"
	CrawlParamsExtractOptionsCountryJm  CrawlParamsExtractOptionsCountry = "JM"
	CrawlParamsExtractOptionsCountryJo  CrawlParamsExtractOptionsCountry = "JO"
	CrawlParamsExtractOptionsCountryJp  CrawlParamsExtractOptionsCountry = "JP"
	CrawlParamsExtractOptionsCountryKe  CrawlParamsExtractOptionsCountry = "KE"
	CrawlParamsExtractOptionsCountryKg  CrawlParamsExtractOptionsCountry = "KG"
	CrawlParamsExtractOptionsCountryKh  CrawlParamsExtractOptionsCountry = "KH"
	CrawlParamsExtractOptionsCountryKi  CrawlParamsExtractOptionsCountry = "KI"
	CrawlParamsExtractOptionsCountryKm  CrawlParamsExtractOptionsCountry = "KM"
	CrawlParamsExtractOptionsCountryKn  CrawlParamsExtractOptionsCountry = "KN"
	CrawlParamsExtractOptionsCountryKp  CrawlParamsExtractOptionsCountry = "KP"
	CrawlParamsExtractOptionsCountryKr  CrawlParamsExtractOptionsCountry = "KR"
	CrawlParamsExtractOptionsCountryKw  CrawlParamsExtractOptionsCountry = "KW"
	CrawlParamsExtractOptionsCountryKy  CrawlParamsExtractOptionsCountry = "KY"
	CrawlParamsExtractOptionsCountryKz  CrawlParamsExtractOptionsCountry = "KZ"
	CrawlParamsExtractOptionsCountryLa  CrawlParamsExtractOptionsCountry = "LA"
	CrawlParamsExtractOptionsCountryLb  CrawlParamsExtractOptionsCountry = "LB"
	CrawlParamsExtractOptionsCountryLc  CrawlParamsExtractOptionsCountry = "LC"
	CrawlParamsExtractOptionsCountryLi  CrawlParamsExtractOptionsCountry = "LI"
	CrawlParamsExtractOptionsCountryLk  CrawlParamsExtractOptionsCountry = "LK"
	CrawlParamsExtractOptionsCountryLr  CrawlParamsExtractOptionsCountry = "LR"
	CrawlParamsExtractOptionsCountryLs  CrawlParamsExtractOptionsCountry = "LS"
	CrawlParamsExtractOptionsCountryLt  CrawlParamsExtractOptionsCountry = "LT"
	CrawlParamsExtractOptionsCountryLu  CrawlParamsExtractOptionsCountry = "LU"
	CrawlParamsExtractOptionsCountryLv  CrawlParamsExtractOptionsCountry = "LV"
	CrawlParamsExtractOptionsCountryLy  CrawlParamsExtractOptionsCountry = "LY"
	CrawlParamsExtractOptionsCountryMa  CrawlParamsExtractOptionsCountry = "MA"
	CrawlParamsExtractOptionsCountryMc  CrawlParamsExtractOptionsCountry = "MC"
	CrawlParamsExtractOptionsCountryMd  CrawlParamsExtractOptionsCountry = "MD"
	CrawlParamsExtractOptionsCountryMe  CrawlParamsExtractOptionsCountry = "ME"
	CrawlParamsExtractOptionsCountryMf  CrawlParamsExtractOptionsCountry = "MF"
	CrawlParamsExtractOptionsCountryMg  CrawlParamsExtractOptionsCountry = "MG"
	CrawlParamsExtractOptionsCountryMh  CrawlParamsExtractOptionsCountry = "MH"
	CrawlParamsExtractOptionsCountryMk  CrawlParamsExtractOptionsCountry = "MK"
	CrawlParamsExtractOptionsCountryMl  CrawlParamsExtractOptionsCountry = "ML"
	CrawlParamsExtractOptionsCountryMm  CrawlParamsExtractOptionsCountry = "MM"
	CrawlParamsExtractOptionsCountryMn  CrawlParamsExtractOptionsCountry = "MN"
	CrawlParamsExtractOptionsCountryMo  CrawlParamsExtractOptionsCountry = "MO"
	CrawlParamsExtractOptionsCountryMp  CrawlParamsExtractOptionsCountry = "MP"
	CrawlParamsExtractOptionsCountryMq  CrawlParamsExtractOptionsCountry = "MQ"
	CrawlParamsExtractOptionsCountryMr  CrawlParamsExtractOptionsCountry = "MR"
	CrawlParamsExtractOptionsCountryMs  CrawlParamsExtractOptionsCountry = "MS"
	CrawlParamsExtractOptionsCountryMt  CrawlParamsExtractOptionsCountry = "MT"
	CrawlParamsExtractOptionsCountryMu  CrawlParamsExtractOptionsCountry = "MU"
	CrawlParamsExtractOptionsCountryMv  CrawlParamsExtractOptionsCountry = "MV"
	CrawlParamsExtractOptionsCountryMw  CrawlParamsExtractOptionsCountry = "MW"
	CrawlParamsExtractOptionsCountryMx  CrawlParamsExtractOptionsCountry = "MX"
	CrawlParamsExtractOptionsCountryMy  CrawlParamsExtractOptionsCountry = "MY"
	CrawlParamsExtractOptionsCountryMz  CrawlParamsExtractOptionsCountry = "MZ"
	CrawlParamsExtractOptionsCountryNa  CrawlParamsExtractOptionsCountry = "NA"
	CrawlParamsExtractOptionsCountryNc  CrawlParamsExtractOptionsCountry = "NC"
	CrawlParamsExtractOptionsCountryNe  CrawlParamsExtractOptionsCountry = "NE"
	CrawlParamsExtractOptionsCountryNf  CrawlParamsExtractOptionsCountry = "NF"
	CrawlParamsExtractOptionsCountryNg  CrawlParamsExtractOptionsCountry = "NG"
	CrawlParamsExtractOptionsCountryNi  CrawlParamsExtractOptionsCountry = "NI"
	CrawlParamsExtractOptionsCountryNl  CrawlParamsExtractOptionsCountry = "NL"
	CrawlParamsExtractOptionsCountryNo  CrawlParamsExtractOptionsCountry = "NO"
	CrawlParamsExtractOptionsCountryNp  CrawlParamsExtractOptionsCountry = "NP"
	CrawlParamsExtractOptionsCountryNr  CrawlParamsExtractOptionsCountry = "NR"
	CrawlParamsExtractOptionsCountryNu  CrawlParamsExtractOptionsCountry = "NU"
	CrawlParamsExtractOptionsCountryNz  CrawlParamsExtractOptionsCountry = "NZ"
	CrawlParamsExtractOptionsCountryOm  CrawlParamsExtractOptionsCountry = "OM"
	CrawlParamsExtractOptionsCountryPa  CrawlParamsExtractOptionsCountry = "PA"
	CrawlParamsExtractOptionsCountryPe  CrawlParamsExtractOptionsCountry = "PE"
	CrawlParamsExtractOptionsCountryPf  CrawlParamsExtractOptionsCountry = "PF"
	CrawlParamsExtractOptionsCountryPg  CrawlParamsExtractOptionsCountry = "PG"
	CrawlParamsExtractOptionsCountryPh  CrawlParamsExtractOptionsCountry = "PH"
	CrawlParamsExtractOptionsCountryPk  CrawlParamsExtractOptionsCountry = "PK"
	CrawlParamsExtractOptionsCountryPl  CrawlParamsExtractOptionsCountry = "PL"
	CrawlParamsExtractOptionsCountryPm  CrawlParamsExtractOptionsCountry = "PM"
	CrawlParamsExtractOptionsCountryPn  CrawlParamsExtractOptionsCountry = "PN"
	CrawlParamsExtractOptionsCountryPr  CrawlParamsExtractOptionsCountry = "PR"
	CrawlParamsExtractOptionsCountryPs  CrawlParamsExtractOptionsCountry = "PS"
	CrawlParamsExtractOptionsCountryPt  CrawlParamsExtractOptionsCountry = "PT"
	CrawlParamsExtractOptionsCountryPw  CrawlParamsExtractOptionsCountry = "PW"
	CrawlParamsExtractOptionsCountryPy  CrawlParamsExtractOptionsCountry = "PY"
	CrawlParamsExtractOptionsCountryQa  CrawlParamsExtractOptionsCountry = "QA"
	CrawlParamsExtractOptionsCountryRe  CrawlParamsExtractOptionsCountry = "RE"
	CrawlParamsExtractOptionsCountryRo  CrawlParamsExtractOptionsCountry = "RO"
	CrawlParamsExtractOptionsCountryRs  CrawlParamsExtractOptionsCountry = "RS"
	CrawlParamsExtractOptionsCountryRu  CrawlParamsExtractOptionsCountry = "RU"
	CrawlParamsExtractOptionsCountryRw  CrawlParamsExtractOptionsCountry = "RW"
	CrawlParamsExtractOptionsCountrySa  CrawlParamsExtractOptionsCountry = "SA"
	CrawlParamsExtractOptionsCountrySb  CrawlParamsExtractOptionsCountry = "SB"
	CrawlParamsExtractOptionsCountrySc  CrawlParamsExtractOptionsCountry = "SC"
	CrawlParamsExtractOptionsCountrySd  CrawlParamsExtractOptionsCountry = "SD"
	CrawlParamsExtractOptionsCountrySe  CrawlParamsExtractOptionsCountry = "SE"
	CrawlParamsExtractOptionsCountrySg  CrawlParamsExtractOptionsCountry = "SG"
	CrawlParamsExtractOptionsCountrySh  CrawlParamsExtractOptionsCountry = "SH"
	CrawlParamsExtractOptionsCountrySi  CrawlParamsExtractOptionsCountry = "SI"
	CrawlParamsExtractOptionsCountrySj  CrawlParamsExtractOptionsCountry = "SJ"
	CrawlParamsExtractOptionsCountrySk  CrawlParamsExtractOptionsCountry = "SK"
	CrawlParamsExtractOptionsCountrySl  CrawlParamsExtractOptionsCountry = "SL"
	CrawlParamsExtractOptionsCountrySm  CrawlParamsExtractOptionsCountry = "SM"
	CrawlParamsExtractOptionsCountrySn  CrawlParamsExtractOptionsCountry = "SN"
	CrawlParamsExtractOptionsCountrySo  CrawlParamsExtractOptionsCountry = "SO"
	CrawlParamsExtractOptionsCountrySr  CrawlParamsExtractOptionsCountry = "SR"
	CrawlParamsExtractOptionsCountrySS  CrawlParamsExtractOptionsCountry = "SS"
	CrawlParamsExtractOptionsCountrySt  CrawlParamsExtractOptionsCountry = "ST"
	CrawlParamsExtractOptionsCountrySv  CrawlParamsExtractOptionsCountry = "SV"
	CrawlParamsExtractOptionsCountrySx  CrawlParamsExtractOptionsCountry = "SX"
	CrawlParamsExtractOptionsCountrySy  CrawlParamsExtractOptionsCountry = "SY"
	CrawlParamsExtractOptionsCountrySz  CrawlParamsExtractOptionsCountry = "SZ"
	CrawlParamsExtractOptionsCountryTc  CrawlParamsExtractOptionsCountry = "TC"
	CrawlParamsExtractOptionsCountryTd  CrawlParamsExtractOptionsCountry = "TD"
	CrawlParamsExtractOptionsCountryTf  CrawlParamsExtractOptionsCountry = "TF"
	CrawlParamsExtractOptionsCountryTg  CrawlParamsExtractOptionsCountry = "TG"
	CrawlParamsExtractOptionsCountryTh  CrawlParamsExtractOptionsCountry = "TH"
	CrawlParamsExtractOptionsCountryTj  CrawlParamsExtractOptionsCountry = "TJ"
	CrawlParamsExtractOptionsCountryTk  CrawlParamsExtractOptionsCountry = "TK"
	CrawlParamsExtractOptionsCountryTl  CrawlParamsExtractOptionsCountry = "TL"
	CrawlParamsExtractOptionsCountryTm  CrawlParamsExtractOptionsCountry = "TM"
	CrawlParamsExtractOptionsCountryTn  CrawlParamsExtractOptionsCountry = "TN"
	CrawlParamsExtractOptionsCountryTo  CrawlParamsExtractOptionsCountry = "TO"
	CrawlParamsExtractOptionsCountryTr  CrawlParamsExtractOptionsCountry = "TR"
	CrawlParamsExtractOptionsCountryTt  CrawlParamsExtractOptionsCountry = "TT"
	CrawlParamsExtractOptionsCountryTv  CrawlParamsExtractOptionsCountry = "TV"
	CrawlParamsExtractOptionsCountryTw  CrawlParamsExtractOptionsCountry = "TW"
	CrawlParamsExtractOptionsCountryTz  CrawlParamsExtractOptionsCountry = "TZ"
	CrawlParamsExtractOptionsCountryUa  CrawlParamsExtractOptionsCountry = "UA"
	CrawlParamsExtractOptionsCountryUg  CrawlParamsExtractOptionsCountry = "UG"
	CrawlParamsExtractOptionsCountryUm  CrawlParamsExtractOptionsCountry = "UM"
	CrawlParamsExtractOptionsCountryUs  CrawlParamsExtractOptionsCountry = "US"
	CrawlParamsExtractOptionsCountryUy  CrawlParamsExtractOptionsCountry = "UY"
	CrawlParamsExtractOptionsCountryUz  CrawlParamsExtractOptionsCountry = "UZ"
	CrawlParamsExtractOptionsCountryVa  CrawlParamsExtractOptionsCountry = "VA"
	CrawlParamsExtractOptionsCountryVc  CrawlParamsExtractOptionsCountry = "VC"
	CrawlParamsExtractOptionsCountryVe  CrawlParamsExtractOptionsCountry = "VE"
	CrawlParamsExtractOptionsCountryVg  CrawlParamsExtractOptionsCountry = "VG"
	CrawlParamsExtractOptionsCountryVi  CrawlParamsExtractOptionsCountry = "VI"
	CrawlParamsExtractOptionsCountryVn  CrawlParamsExtractOptionsCountry = "VN"
	CrawlParamsExtractOptionsCountryVu  CrawlParamsExtractOptionsCountry = "VU"
	CrawlParamsExtractOptionsCountryWf  CrawlParamsExtractOptionsCountry = "WF"
	CrawlParamsExtractOptionsCountryWs  CrawlParamsExtractOptionsCountry = "WS"
	CrawlParamsExtractOptionsCountryXk  CrawlParamsExtractOptionsCountry = "XK"
	CrawlParamsExtractOptionsCountryYe  CrawlParamsExtractOptionsCountry = "YE"
	CrawlParamsExtractOptionsCountryYt  CrawlParamsExtractOptionsCountry = "YT"
	CrawlParamsExtractOptionsCountryZa  CrawlParamsExtractOptionsCountry = "ZA"
	CrawlParamsExtractOptionsCountryZm  CrawlParamsExtractOptionsCountry = "ZM"
	CrawlParamsExtractOptionsCountryZw  CrawlParamsExtractOptionsCountry = "ZW"
	CrawlParamsExtractOptionsCountryAll CrawlParamsExtractOptionsCountry = "ALL"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsHeaderUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsHeaderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlParamsExtractOptionsHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsHeaderUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Locale for browser language and region settings
type CrawlParamsExtractOptionsLocale string

const (
	CrawlParamsExtractOptionsLocaleAaDj      CrawlParamsExtractOptionsLocale = "aa-DJ"
	CrawlParamsExtractOptionsLocaleAaEr      CrawlParamsExtractOptionsLocale = "aa-ER"
	CrawlParamsExtractOptionsLocaleAaEt      CrawlParamsExtractOptionsLocale = "aa-ET"
	CrawlParamsExtractOptionsLocaleAf        CrawlParamsExtractOptionsLocale = "af"
	CrawlParamsExtractOptionsLocaleAfNa      CrawlParamsExtractOptionsLocale = "af-NA"
	CrawlParamsExtractOptionsLocaleAfZa      CrawlParamsExtractOptionsLocale = "af-ZA"
	CrawlParamsExtractOptionsLocaleAk        CrawlParamsExtractOptionsLocale = "ak"
	CrawlParamsExtractOptionsLocaleAkGh      CrawlParamsExtractOptionsLocale = "ak-GH"
	CrawlParamsExtractOptionsLocaleAm        CrawlParamsExtractOptionsLocale = "am"
	CrawlParamsExtractOptionsLocaleAmEt      CrawlParamsExtractOptionsLocale = "am-ET"
	CrawlParamsExtractOptionsLocaleAnEs      CrawlParamsExtractOptionsLocale = "an-ES"
	CrawlParamsExtractOptionsLocaleAr        CrawlParamsExtractOptionsLocale = "ar"
	CrawlParamsExtractOptionsLocaleArAe      CrawlParamsExtractOptionsLocale = "ar-AE"
	CrawlParamsExtractOptionsLocaleArBh      CrawlParamsExtractOptionsLocale = "ar-BH"
	CrawlParamsExtractOptionsLocaleArDz      CrawlParamsExtractOptionsLocale = "ar-DZ"
	CrawlParamsExtractOptionsLocaleArEg      CrawlParamsExtractOptionsLocale = "ar-EG"
	CrawlParamsExtractOptionsLocaleArIn      CrawlParamsExtractOptionsLocale = "ar-IN"
	CrawlParamsExtractOptionsLocaleArIq      CrawlParamsExtractOptionsLocale = "ar-IQ"
	CrawlParamsExtractOptionsLocaleArJo      CrawlParamsExtractOptionsLocale = "ar-JO"
	CrawlParamsExtractOptionsLocaleArKw      CrawlParamsExtractOptionsLocale = "ar-KW"
	CrawlParamsExtractOptionsLocaleArLb      CrawlParamsExtractOptionsLocale = "ar-LB"
	CrawlParamsExtractOptionsLocaleArLy      CrawlParamsExtractOptionsLocale = "ar-LY"
	CrawlParamsExtractOptionsLocaleArMa      CrawlParamsExtractOptionsLocale = "ar-MA"
	CrawlParamsExtractOptionsLocaleArOm      CrawlParamsExtractOptionsLocale = "ar-OM"
	CrawlParamsExtractOptionsLocaleArQa      CrawlParamsExtractOptionsLocale = "ar-QA"
	CrawlParamsExtractOptionsLocaleArSa      CrawlParamsExtractOptionsLocale = "ar-SA"
	CrawlParamsExtractOptionsLocaleArSd      CrawlParamsExtractOptionsLocale = "ar-SD"
	CrawlParamsExtractOptionsLocaleArSy      CrawlParamsExtractOptionsLocale = "ar-SY"
	CrawlParamsExtractOptionsLocaleArTn      CrawlParamsExtractOptionsLocale = "ar-TN"
	CrawlParamsExtractOptionsLocaleArYe      CrawlParamsExtractOptionsLocale = "ar-YE"
	CrawlParamsExtractOptionsLocaleAs        CrawlParamsExtractOptionsLocale = "as"
	CrawlParamsExtractOptionsLocaleAsIn      CrawlParamsExtractOptionsLocale = "as-IN"
	CrawlParamsExtractOptionsLocaleAsa       CrawlParamsExtractOptionsLocale = "asa"
	CrawlParamsExtractOptionsLocaleAsaTz     CrawlParamsExtractOptionsLocale = "asa-TZ"
	CrawlParamsExtractOptionsLocaleAstEs     CrawlParamsExtractOptionsLocale = "ast-ES"
	CrawlParamsExtractOptionsLocaleAz        CrawlParamsExtractOptionsLocale = "az"
	CrawlParamsExtractOptionsLocaleAzAz      CrawlParamsExtractOptionsLocale = "az-AZ"
	CrawlParamsExtractOptionsLocaleAzCyrl    CrawlParamsExtractOptionsLocale = "az-Cyrl"
	CrawlParamsExtractOptionsLocaleAzCyrlAz  CrawlParamsExtractOptionsLocale = "az-Cyrl-AZ"
	CrawlParamsExtractOptionsLocaleAzLatn    CrawlParamsExtractOptionsLocale = "az-Latn"
	CrawlParamsExtractOptionsLocaleAzLatnAz  CrawlParamsExtractOptionsLocale = "az-Latn-AZ"
	CrawlParamsExtractOptionsLocaleBe        CrawlParamsExtractOptionsLocale = "be"
	CrawlParamsExtractOptionsLocaleBeBy      CrawlParamsExtractOptionsLocale = "be-BY"
	CrawlParamsExtractOptionsLocaleBem       CrawlParamsExtractOptionsLocale = "bem"
	CrawlParamsExtractOptionsLocaleBemZm     CrawlParamsExtractOptionsLocale = "bem-ZM"
	CrawlParamsExtractOptionsLocaleBerDz     CrawlParamsExtractOptionsLocale = "ber-DZ"
	CrawlParamsExtractOptionsLocaleBerMa     CrawlParamsExtractOptionsLocale = "ber-MA"
	CrawlParamsExtractOptionsLocaleBez       CrawlParamsExtractOptionsLocale = "bez"
	CrawlParamsExtractOptionsLocaleBezTz     CrawlParamsExtractOptionsLocale = "bez-TZ"
	CrawlParamsExtractOptionsLocaleBg        CrawlParamsExtractOptionsLocale = "bg"
	CrawlParamsExtractOptionsLocaleBgBg      CrawlParamsExtractOptionsLocale = "bg-BG"
	CrawlParamsExtractOptionsLocaleBhoIn     CrawlParamsExtractOptionsLocale = "bho-IN"
	CrawlParamsExtractOptionsLocaleBm        CrawlParamsExtractOptionsLocale = "bm"
	CrawlParamsExtractOptionsLocaleBmMl      CrawlParamsExtractOptionsLocale = "bm-ML"
	CrawlParamsExtractOptionsLocaleBn        CrawlParamsExtractOptionsLocale = "bn"
	CrawlParamsExtractOptionsLocaleBnBd      CrawlParamsExtractOptionsLocale = "bn-BD"
	CrawlParamsExtractOptionsLocaleBnIn      CrawlParamsExtractOptionsLocale = "bn-IN"
	CrawlParamsExtractOptionsLocaleBo        CrawlParamsExtractOptionsLocale = "bo"
	CrawlParamsExtractOptionsLocaleBoCn      CrawlParamsExtractOptionsLocale = "bo-CN"
	CrawlParamsExtractOptionsLocaleBoIn      CrawlParamsExtractOptionsLocale = "bo-IN"
	CrawlParamsExtractOptionsLocaleBrFr      CrawlParamsExtractOptionsLocale = "br-FR"
	CrawlParamsExtractOptionsLocaleBrxIn     CrawlParamsExtractOptionsLocale = "brx-IN"
	CrawlParamsExtractOptionsLocaleBs        CrawlParamsExtractOptionsLocale = "bs"
	CrawlParamsExtractOptionsLocaleBsBa      CrawlParamsExtractOptionsLocale = "bs-BA"
	CrawlParamsExtractOptionsLocaleBynEr     CrawlParamsExtractOptionsLocale = "byn-ER"
	CrawlParamsExtractOptionsLocaleCa        CrawlParamsExtractOptionsLocale = "ca"
	CrawlParamsExtractOptionsLocaleCaAd      CrawlParamsExtractOptionsLocale = "ca-AD"
	CrawlParamsExtractOptionsLocaleCaEs      CrawlParamsExtractOptionsLocale = "ca-ES"
	CrawlParamsExtractOptionsLocaleCaFr      CrawlParamsExtractOptionsLocale = "ca-FR"
	CrawlParamsExtractOptionsLocaleCaIt      CrawlParamsExtractOptionsLocale = "ca-IT"
	CrawlParamsExtractOptionsLocaleCgg       CrawlParamsExtractOptionsLocale = "cgg"
	CrawlParamsExtractOptionsLocaleCggUg     CrawlParamsExtractOptionsLocale = "cgg-UG"
	CrawlParamsExtractOptionsLocaleChr       CrawlParamsExtractOptionsLocale = "chr"
	CrawlParamsExtractOptionsLocaleChrUs     CrawlParamsExtractOptionsLocale = "chr-US"
	CrawlParamsExtractOptionsLocaleCrhUa     CrawlParamsExtractOptionsLocale = "crh-UA"
	CrawlParamsExtractOptionsLocaleCs        CrawlParamsExtractOptionsLocale = "cs"
	CrawlParamsExtractOptionsLocaleCsCz      CrawlParamsExtractOptionsLocale = "cs-CZ"
	CrawlParamsExtractOptionsLocaleCsbPl     CrawlParamsExtractOptionsLocale = "csb-PL"
	CrawlParamsExtractOptionsLocaleCvRu      CrawlParamsExtractOptionsLocale = "cv-RU"
	CrawlParamsExtractOptionsLocaleCy        CrawlParamsExtractOptionsLocale = "cy"
	CrawlParamsExtractOptionsLocaleCyGB      CrawlParamsExtractOptionsLocale = "cy-GB"
	CrawlParamsExtractOptionsLocaleDa        CrawlParamsExtractOptionsLocale = "da"
	CrawlParamsExtractOptionsLocaleDaDk      CrawlParamsExtractOptionsLocale = "da-DK"
	CrawlParamsExtractOptionsLocaleDav       CrawlParamsExtractOptionsLocale = "dav"
	CrawlParamsExtractOptionsLocaleDavKe     CrawlParamsExtractOptionsLocale = "dav-KE"
	CrawlParamsExtractOptionsLocaleDe        CrawlParamsExtractOptionsLocale = "de"
	CrawlParamsExtractOptionsLocaleDeAt      CrawlParamsExtractOptionsLocale = "de-AT"
	CrawlParamsExtractOptionsLocaleDeBe      CrawlParamsExtractOptionsLocale = "de-BE"
	CrawlParamsExtractOptionsLocaleDeCh      CrawlParamsExtractOptionsLocale = "de-CH"
	CrawlParamsExtractOptionsLocaleDeDe      CrawlParamsExtractOptionsLocale = "de-DE"
	CrawlParamsExtractOptionsLocaleDeLi      CrawlParamsExtractOptionsLocale = "de-LI"
	CrawlParamsExtractOptionsLocaleDeLu      CrawlParamsExtractOptionsLocale = "de-LU"
	CrawlParamsExtractOptionsLocaleDvMv      CrawlParamsExtractOptionsLocale = "dv-MV"
	CrawlParamsExtractOptionsLocaleDzBt      CrawlParamsExtractOptionsLocale = "dz-BT"
	CrawlParamsExtractOptionsLocaleEbu       CrawlParamsExtractOptionsLocale = "ebu"
	CrawlParamsExtractOptionsLocaleEbuKe     CrawlParamsExtractOptionsLocale = "ebu-KE"
	CrawlParamsExtractOptionsLocaleEe        CrawlParamsExtractOptionsLocale = "ee"
	CrawlParamsExtractOptionsLocaleEeGh      CrawlParamsExtractOptionsLocale = "ee-GH"
	CrawlParamsExtractOptionsLocaleEeTg      CrawlParamsExtractOptionsLocale = "ee-TG"
	CrawlParamsExtractOptionsLocaleEl        CrawlParamsExtractOptionsLocale = "el"
	CrawlParamsExtractOptionsLocaleElCy      CrawlParamsExtractOptionsLocale = "el-CY"
	CrawlParamsExtractOptionsLocaleElGr      CrawlParamsExtractOptionsLocale = "el-GR"
	CrawlParamsExtractOptionsLocaleEn        CrawlParamsExtractOptionsLocale = "en"
	CrawlParamsExtractOptionsLocaleEnAg      CrawlParamsExtractOptionsLocale = "en-AG"
	CrawlParamsExtractOptionsLocaleEnAs      CrawlParamsExtractOptionsLocale = "en-AS"
	CrawlParamsExtractOptionsLocaleEnAu      CrawlParamsExtractOptionsLocale = "en-AU"
	CrawlParamsExtractOptionsLocaleEnBe      CrawlParamsExtractOptionsLocale = "en-BE"
	CrawlParamsExtractOptionsLocaleEnBw      CrawlParamsExtractOptionsLocale = "en-BW"
	CrawlParamsExtractOptionsLocaleEnBz      CrawlParamsExtractOptionsLocale = "en-BZ"
	CrawlParamsExtractOptionsLocaleEnCa      CrawlParamsExtractOptionsLocale = "en-CA"
	CrawlParamsExtractOptionsLocaleEnDk      CrawlParamsExtractOptionsLocale = "en-DK"
	CrawlParamsExtractOptionsLocaleEnGB      CrawlParamsExtractOptionsLocale = "en-GB"
	CrawlParamsExtractOptionsLocaleEnGu      CrawlParamsExtractOptionsLocale = "en-GU"
	CrawlParamsExtractOptionsLocaleEnHk      CrawlParamsExtractOptionsLocale = "en-HK"
	CrawlParamsExtractOptionsLocaleEnIe      CrawlParamsExtractOptionsLocale = "en-IE"
	CrawlParamsExtractOptionsLocaleEnIn      CrawlParamsExtractOptionsLocale = "en-IN"
	CrawlParamsExtractOptionsLocaleEnJm      CrawlParamsExtractOptionsLocale = "en-JM"
	CrawlParamsExtractOptionsLocaleEnMh      CrawlParamsExtractOptionsLocale = "en-MH"
	CrawlParamsExtractOptionsLocaleEnMp      CrawlParamsExtractOptionsLocale = "en-MP"
	CrawlParamsExtractOptionsLocaleEnMt      CrawlParamsExtractOptionsLocale = "en-MT"
	CrawlParamsExtractOptionsLocaleEnMu      CrawlParamsExtractOptionsLocale = "en-MU"
	CrawlParamsExtractOptionsLocaleEnNa      CrawlParamsExtractOptionsLocale = "en-NA"
	CrawlParamsExtractOptionsLocaleEnNg      CrawlParamsExtractOptionsLocale = "en-NG"
	CrawlParamsExtractOptionsLocaleEnNz      CrawlParamsExtractOptionsLocale = "en-NZ"
	CrawlParamsExtractOptionsLocaleEnPh      CrawlParamsExtractOptionsLocale = "en-PH"
	CrawlParamsExtractOptionsLocaleEnPk      CrawlParamsExtractOptionsLocale = "en-PK"
	CrawlParamsExtractOptionsLocaleEnSg      CrawlParamsExtractOptionsLocale = "en-SG"
	CrawlParamsExtractOptionsLocaleEnTt      CrawlParamsExtractOptionsLocale = "en-TT"
	CrawlParamsExtractOptionsLocaleEnUm      CrawlParamsExtractOptionsLocale = "en-UM"
	CrawlParamsExtractOptionsLocaleEnUs      CrawlParamsExtractOptionsLocale = "en-US"
	CrawlParamsExtractOptionsLocaleEnVi      CrawlParamsExtractOptionsLocale = "en-VI"
	CrawlParamsExtractOptionsLocaleEnZa      CrawlParamsExtractOptionsLocale = "en-ZA"
	CrawlParamsExtractOptionsLocaleEnZm      CrawlParamsExtractOptionsLocale = "en-ZM"
	CrawlParamsExtractOptionsLocaleEnZw      CrawlParamsExtractOptionsLocale = "en-ZW"
	CrawlParamsExtractOptionsLocaleEo        CrawlParamsExtractOptionsLocale = "eo"
	CrawlParamsExtractOptionsLocaleEs        CrawlParamsExtractOptionsLocale = "es"
	CrawlParamsExtractOptionsLocaleEs419     CrawlParamsExtractOptionsLocale = "es-419"
	CrawlParamsExtractOptionsLocaleEsAr      CrawlParamsExtractOptionsLocale = "es-AR"
	CrawlParamsExtractOptionsLocaleEsBo      CrawlParamsExtractOptionsLocale = "es-BO"
	CrawlParamsExtractOptionsLocaleEsCl      CrawlParamsExtractOptionsLocale = "es-CL"
	CrawlParamsExtractOptionsLocaleEsCo      CrawlParamsExtractOptionsLocale = "es-CO"
	CrawlParamsExtractOptionsLocaleEsCr      CrawlParamsExtractOptionsLocale = "es-CR"
	CrawlParamsExtractOptionsLocaleEsCu      CrawlParamsExtractOptionsLocale = "es-CU"
	CrawlParamsExtractOptionsLocaleEsDo      CrawlParamsExtractOptionsLocale = "es-DO"
	CrawlParamsExtractOptionsLocaleEsEc      CrawlParamsExtractOptionsLocale = "es-EC"
	CrawlParamsExtractOptionsLocaleEsEs      CrawlParamsExtractOptionsLocale = "es-ES"
	CrawlParamsExtractOptionsLocaleEsGq      CrawlParamsExtractOptionsLocale = "es-GQ"
	CrawlParamsExtractOptionsLocaleEsGt      CrawlParamsExtractOptionsLocale = "es-GT"
	CrawlParamsExtractOptionsLocaleEsHn      CrawlParamsExtractOptionsLocale = "es-HN"
	CrawlParamsExtractOptionsLocaleEsMx      CrawlParamsExtractOptionsLocale = "es-MX"
	CrawlParamsExtractOptionsLocaleEsNi      CrawlParamsExtractOptionsLocale = "es-NI"
	CrawlParamsExtractOptionsLocaleEsPa      CrawlParamsExtractOptionsLocale = "es-PA"
	CrawlParamsExtractOptionsLocaleEsPe      CrawlParamsExtractOptionsLocale = "es-PE"
	CrawlParamsExtractOptionsLocaleEsPr      CrawlParamsExtractOptionsLocale = "es-PR"
	CrawlParamsExtractOptionsLocaleEsPy      CrawlParamsExtractOptionsLocale = "es-PY"
	CrawlParamsExtractOptionsLocaleEsSv      CrawlParamsExtractOptionsLocale = "es-SV"
	CrawlParamsExtractOptionsLocaleEsUs      CrawlParamsExtractOptionsLocale = "es-US"
	CrawlParamsExtractOptionsLocaleEsUy      CrawlParamsExtractOptionsLocale = "es-UY"
	CrawlParamsExtractOptionsLocaleEsVe      CrawlParamsExtractOptionsLocale = "es-VE"
	CrawlParamsExtractOptionsLocaleEt        CrawlParamsExtractOptionsLocale = "et"
	CrawlParamsExtractOptionsLocaleEtEe      CrawlParamsExtractOptionsLocale = "et-EE"
	CrawlParamsExtractOptionsLocaleEu        CrawlParamsExtractOptionsLocale = "eu"
	CrawlParamsExtractOptionsLocaleEuEs      CrawlParamsExtractOptionsLocale = "eu-ES"
	CrawlParamsExtractOptionsLocaleFa        CrawlParamsExtractOptionsLocale = "fa"
	CrawlParamsExtractOptionsLocaleFaAf      CrawlParamsExtractOptionsLocale = "fa-AF"
	CrawlParamsExtractOptionsLocaleFaIr      CrawlParamsExtractOptionsLocale = "fa-IR"
	CrawlParamsExtractOptionsLocaleFf        CrawlParamsExtractOptionsLocale = "ff"
	CrawlParamsExtractOptionsLocaleFfSn      CrawlParamsExtractOptionsLocale = "ff-SN"
	CrawlParamsExtractOptionsLocaleFi        CrawlParamsExtractOptionsLocale = "fi"
	CrawlParamsExtractOptionsLocaleFiFi      CrawlParamsExtractOptionsLocale = "fi-FI"
	CrawlParamsExtractOptionsLocaleFil       CrawlParamsExtractOptionsLocale = "fil"
	CrawlParamsExtractOptionsLocaleFilPh     CrawlParamsExtractOptionsLocale = "fil-PH"
	CrawlParamsExtractOptionsLocaleFo        CrawlParamsExtractOptionsLocale = "fo"
	CrawlParamsExtractOptionsLocaleFoFo      CrawlParamsExtractOptionsLocale = "fo-FO"
	CrawlParamsExtractOptionsLocaleFr        CrawlParamsExtractOptionsLocale = "fr"
	CrawlParamsExtractOptionsLocaleFrBe      CrawlParamsExtractOptionsLocale = "fr-BE"
	CrawlParamsExtractOptionsLocaleFrBf      CrawlParamsExtractOptionsLocale = "fr-BF"
	CrawlParamsExtractOptionsLocaleFrBi      CrawlParamsExtractOptionsLocale = "fr-BI"
	CrawlParamsExtractOptionsLocaleFrBj      CrawlParamsExtractOptionsLocale = "fr-BJ"
	CrawlParamsExtractOptionsLocaleFrBl      CrawlParamsExtractOptionsLocale = "fr-BL"
	CrawlParamsExtractOptionsLocaleFrCa      CrawlParamsExtractOptionsLocale = "fr-CA"
	CrawlParamsExtractOptionsLocaleFrCd      CrawlParamsExtractOptionsLocale = "fr-CD"
	CrawlParamsExtractOptionsLocaleFrCf      CrawlParamsExtractOptionsLocale = "fr-CF"
	CrawlParamsExtractOptionsLocaleFrCg      CrawlParamsExtractOptionsLocale = "fr-CG"
	CrawlParamsExtractOptionsLocaleFrCh      CrawlParamsExtractOptionsLocale = "fr-CH"
	CrawlParamsExtractOptionsLocaleFrCi      CrawlParamsExtractOptionsLocale = "fr-CI"
	CrawlParamsExtractOptionsLocaleFrCm      CrawlParamsExtractOptionsLocale = "fr-CM"
	CrawlParamsExtractOptionsLocaleFrDj      CrawlParamsExtractOptionsLocale = "fr-DJ"
	CrawlParamsExtractOptionsLocaleFrFr      CrawlParamsExtractOptionsLocale = "fr-FR"
	CrawlParamsExtractOptionsLocaleFrGa      CrawlParamsExtractOptionsLocale = "fr-GA"
	CrawlParamsExtractOptionsLocaleFrGn      CrawlParamsExtractOptionsLocale = "fr-GN"
	CrawlParamsExtractOptionsLocaleFrGp      CrawlParamsExtractOptionsLocale = "fr-GP"
	CrawlParamsExtractOptionsLocaleFrGq      CrawlParamsExtractOptionsLocale = "fr-GQ"
	CrawlParamsExtractOptionsLocaleFrKm      CrawlParamsExtractOptionsLocale = "fr-KM"
	CrawlParamsExtractOptionsLocaleFrLu      CrawlParamsExtractOptionsLocale = "fr-LU"
	CrawlParamsExtractOptionsLocaleFrMc      CrawlParamsExtractOptionsLocale = "fr-MC"
	CrawlParamsExtractOptionsLocaleFrMf      CrawlParamsExtractOptionsLocale = "fr-MF"
	CrawlParamsExtractOptionsLocaleFrMg      CrawlParamsExtractOptionsLocale = "fr-MG"
	CrawlParamsExtractOptionsLocaleFrMl      CrawlParamsExtractOptionsLocale = "fr-ML"
	CrawlParamsExtractOptionsLocaleFrMq      CrawlParamsExtractOptionsLocale = "fr-MQ"
	CrawlParamsExtractOptionsLocaleFrNe      CrawlParamsExtractOptionsLocale = "fr-NE"
	CrawlParamsExtractOptionsLocaleFrRe      CrawlParamsExtractOptionsLocale = "fr-RE"
	CrawlParamsExtractOptionsLocaleFrRw      CrawlParamsExtractOptionsLocale = "fr-RW"
	CrawlParamsExtractOptionsLocaleFrSn      CrawlParamsExtractOptionsLocale = "fr-SN"
	CrawlParamsExtractOptionsLocaleFrTd      CrawlParamsExtractOptionsLocale = "fr-TD"
	CrawlParamsExtractOptionsLocaleFrTg      CrawlParamsExtractOptionsLocale = "fr-TG"
	CrawlParamsExtractOptionsLocaleFurIt     CrawlParamsExtractOptionsLocale = "fur-IT"
	CrawlParamsExtractOptionsLocaleFyDe      CrawlParamsExtractOptionsLocale = "fy-DE"
	CrawlParamsExtractOptionsLocaleFyNl      CrawlParamsExtractOptionsLocale = "fy-NL"
	CrawlParamsExtractOptionsLocaleGa        CrawlParamsExtractOptionsLocale = "ga"
	CrawlParamsExtractOptionsLocaleGaIe      CrawlParamsExtractOptionsLocale = "ga-IE"
	CrawlParamsExtractOptionsLocaleGdGB      CrawlParamsExtractOptionsLocale = "gd-GB"
	CrawlParamsExtractOptionsLocaleGezEr     CrawlParamsExtractOptionsLocale = "gez-ER"
	CrawlParamsExtractOptionsLocaleGezEt     CrawlParamsExtractOptionsLocale = "gez-ET"
	CrawlParamsExtractOptionsLocaleGl        CrawlParamsExtractOptionsLocale = "gl"
	CrawlParamsExtractOptionsLocaleGlEs      CrawlParamsExtractOptionsLocale = "gl-ES"
	CrawlParamsExtractOptionsLocaleGsw       CrawlParamsExtractOptionsLocale = "gsw"
	CrawlParamsExtractOptionsLocaleGswCh     CrawlParamsExtractOptionsLocale = "gsw-CH"
	CrawlParamsExtractOptionsLocaleGu        CrawlParamsExtractOptionsLocale = "gu"
	CrawlParamsExtractOptionsLocaleGuIn      CrawlParamsExtractOptionsLocale = "gu-IN"
	CrawlParamsExtractOptionsLocaleGuz       CrawlParamsExtractOptionsLocale = "guz"
	CrawlParamsExtractOptionsLocaleGuzKe     CrawlParamsExtractOptionsLocale = "guz-KE"
	CrawlParamsExtractOptionsLocaleGv        CrawlParamsExtractOptionsLocale = "gv"
	CrawlParamsExtractOptionsLocaleGvGB      CrawlParamsExtractOptionsLocale = "gv-GB"
	CrawlParamsExtractOptionsLocaleHa        CrawlParamsExtractOptionsLocale = "ha"
	CrawlParamsExtractOptionsLocaleHaLatn    CrawlParamsExtractOptionsLocale = "ha-Latn"
	CrawlParamsExtractOptionsLocaleHaLatnGh  CrawlParamsExtractOptionsLocale = "ha-Latn-GH"
	CrawlParamsExtractOptionsLocaleHaLatnNe  CrawlParamsExtractOptionsLocale = "ha-Latn-NE"
	CrawlParamsExtractOptionsLocaleHaLatnNg  CrawlParamsExtractOptionsLocale = "ha-Latn-NG"
	CrawlParamsExtractOptionsLocaleHaNg      CrawlParamsExtractOptionsLocale = "ha-NG"
	CrawlParamsExtractOptionsLocaleHaw       CrawlParamsExtractOptionsLocale = "haw"
	CrawlParamsExtractOptionsLocaleHawUs     CrawlParamsExtractOptionsLocale = "haw-US"
	CrawlParamsExtractOptionsLocaleHe        CrawlParamsExtractOptionsLocale = "he"
	CrawlParamsExtractOptionsLocaleHeIl      CrawlParamsExtractOptionsLocale = "he-IL"
	CrawlParamsExtractOptionsLocaleHi        CrawlParamsExtractOptionsLocale = "hi"
	CrawlParamsExtractOptionsLocaleHiIn      CrawlParamsExtractOptionsLocale = "hi-IN"
	CrawlParamsExtractOptionsLocaleHneIn     CrawlParamsExtractOptionsLocale = "hne-IN"
	CrawlParamsExtractOptionsLocaleHr        CrawlParamsExtractOptionsLocale = "hr"
	CrawlParamsExtractOptionsLocaleHrHr      CrawlParamsExtractOptionsLocale = "hr-HR"
	CrawlParamsExtractOptionsLocaleHsbDe     CrawlParamsExtractOptionsLocale = "hsb-DE"
	CrawlParamsExtractOptionsLocaleHtHt      CrawlParamsExtractOptionsLocale = "ht-HT"
	CrawlParamsExtractOptionsLocaleHu        CrawlParamsExtractOptionsLocale = "hu"
	CrawlParamsExtractOptionsLocaleHuHu      CrawlParamsExtractOptionsLocale = "hu-HU"
	CrawlParamsExtractOptionsLocaleHy        CrawlParamsExtractOptionsLocale = "hy"
	CrawlParamsExtractOptionsLocaleHyAm      CrawlParamsExtractOptionsLocale = "hy-AM"
	CrawlParamsExtractOptionsLocaleID        CrawlParamsExtractOptionsLocale = "id"
	CrawlParamsExtractOptionsLocaleIDID      CrawlParamsExtractOptionsLocale = "id-ID"
	CrawlParamsExtractOptionsLocaleIg        CrawlParamsExtractOptionsLocale = "ig"
	CrawlParamsExtractOptionsLocaleIgNg      CrawlParamsExtractOptionsLocale = "ig-NG"
	CrawlParamsExtractOptionsLocaleIi        CrawlParamsExtractOptionsLocale = "ii"
	CrawlParamsExtractOptionsLocaleIiCn      CrawlParamsExtractOptionsLocale = "ii-CN"
	CrawlParamsExtractOptionsLocaleIkCa      CrawlParamsExtractOptionsLocale = "ik-CA"
	CrawlParamsExtractOptionsLocaleIs        CrawlParamsExtractOptionsLocale = "is"
	CrawlParamsExtractOptionsLocaleIsIs      CrawlParamsExtractOptionsLocale = "is-IS"
	CrawlParamsExtractOptionsLocaleIt        CrawlParamsExtractOptionsLocale = "it"
	CrawlParamsExtractOptionsLocaleItCh      CrawlParamsExtractOptionsLocale = "it-CH"
	CrawlParamsExtractOptionsLocaleItIt      CrawlParamsExtractOptionsLocale = "it-IT"
	CrawlParamsExtractOptionsLocaleIuCa      CrawlParamsExtractOptionsLocale = "iu-CA"
	CrawlParamsExtractOptionsLocaleIwIl      CrawlParamsExtractOptionsLocale = "iw-IL"
	CrawlParamsExtractOptionsLocaleJa        CrawlParamsExtractOptionsLocale = "ja"
	CrawlParamsExtractOptionsLocaleJaJp      CrawlParamsExtractOptionsLocale = "ja-JP"
	CrawlParamsExtractOptionsLocaleJmc       CrawlParamsExtractOptionsLocale = "jmc"
	CrawlParamsExtractOptionsLocaleJmcTz     CrawlParamsExtractOptionsLocale = "jmc-TZ"
	CrawlParamsExtractOptionsLocaleKa        CrawlParamsExtractOptionsLocale = "ka"
	CrawlParamsExtractOptionsLocaleKaGe      CrawlParamsExtractOptionsLocale = "ka-GE"
	CrawlParamsExtractOptionsLocaleKab       CrawlParamsExtractOptionsLocale = "kab"
	CrawlParamsExtractOptionsLocaleKabDz     CrawlParamsExtractOptionsLocale = "kab-DZ"
	CrawlParamsExtractOptionsLocaleKam       CrawlParamsExtractOptionsLocale = "kam"
	CrawlParamsExtractOptionsLocaleKamKe     CrawlParamsExtractOptionsLocale = "kam-KE"
	CrawlParamsExtractOptionsLocaleKde       CrawlParamsExtractOptionsLocale = "kde"
	CrawlParamsExtractOptionsLocaleKdeTz     CrawlParamsExtractOptionsLocale = "kde-TZ"
	CrawlParamsExtractOptionsLocaleKea       CrawlParamsExtractOptionsLocale = "kea"
	CrawlParamsExtractOptionsLocaleKeaCv     CrawlParamsExtractOptionsLocale = "kea-CV"
	CrawlParamsExtractOptionsLocaleKhq       CrawlParamsExtractOptionsLocale = "khq"
	CrawlParamsExtractOptionsLocaleKhqMl     CrawlParamsExtractOptionsLocale = "khq-ML"
	CrawlParamsExtractOptionsLocaleKi        CrawlParamsExtractOptionsLocale = "ki"
	CrawlParamsExtractOptionsLocaleKiKe      CrawlParamsExtractOptionsLocale = "ki-KE"
	CrawlParamsExtractOptionsLocaleKk        CrawlParamsExtractOptionsLocale = "kk"
	CrawlParamsExtractOptionsLocaleKkCyrl    CrawlParamsExtractOptionsLocale = "kk-Cyrl"
	CrawlParamsExtractOptionsLocaleKkCyrlKz  CrawlParamsExtractOptionsLocale = "kk-Cyrl-KZ"
	CrawlParamsExtractOptionsLocaleKkKz      CrawlParamsExtractOptionsLocale = "kk-KZ"
	CrawlParamsExtractOptionsLocaleKl        CrawlParamsExtractOptionsLocale = "kl"
	CrawlParamsExtractOptionsLocaleKlGl      CrawlParamsExtractOptionsLocale = "kl-GL"
	CrawlParamsExtractOptionsLocaleKln       CrawlParamsExtractOptionsLocale = "kln"
	CrawlParamsExtractOptionsLocaleKlnKe     CrawlParamsExtractOptionsLocale = "kln-KE"
	CrawlParamsExtractOptionsLocaleKm        CrawlParamsExtractOptionsLocale = "km"
	CrawlParamsExtractOptionsLocaleKmKh      CrawlParamsExtractOptionsLocale = "km-KH"
	CrawlParamsExtractOptionsLocaleKn        CrawlParamsExtractOptionsLocale = "kn"
	CrawlParamsExtractOptionsLocaleKnIn      CrawlParamsExtractOptionsLocale = "kn-IN"
	CrawlParamsExtractOptionsLocaleKo        CrawlParamsExtractOptionsLocale = "ko"
	CrawlParamsExtractOptionsLocaleKoKr      CrawlParamsExtractOptionsLocale = "ko-KR"
	CrawlParamsExtractOptionsLocaleKok       CrawlParamsExtractOptionsLocale = "kok"
	CrawlParamsExtractOptionsLocaleKokIn     CrawlParamsExtractOptionsLocale = "kok-IN"
	CrawlParamsExtractOptionsLocaleKsIn      CrawlParamsExtractOptionsLocale = "ks-IN"
	CrawlParamsExtractOptionsLocaleKuTr      CrawlParamsExtractOptionsLocale = "ku-TR"
	CrawlParamsExtractOptionsLocaleKw        CrawlParamsExtractOptionsLocale = "kw"
	CrawlParamsExtractOptionsLocaleKwGB      CrawlParamsExtractOptionsLocale = "kw-GB"
	CrawlParamsExtractOptionsLocaleKyKg      CrawlParamsExtractOptionsLocale = "ky-KG"
	CrawlParamsExtractOptionsLocaleLag       CrawlParamsExtractOptionsLocale = "lag"
	CrawlParamsExtractOptionsLocaleLagTz     CrawlParamsExtractOptionsLocale = "lag-TZ"
	CrawlParamsExtractOptionsLocaleLbLu      CrawlParamsExtractOptionsLocale = "lb-LU"
	CrawlParamsExtractOptionsLocaleLg        CrawlParamsExtractOptionsLocale = "lg"
	CrawlParamsExtractOptionsLocaleLgUg      CrawlParamsExtractOptionsLocale = "lg-UG"
	CrawlParamsExtractOptionsLocaleLiBe      CrawlParamsExtractOptionsLocale = "li-BE"
	CrawlParamsExtractOptionsLocaleLiNl      CrawlParamsExtractOptionsLocale = "li-NL"
	CrawlParamsExtractOptionsLocaleLijIt     CrawlParamsExtractOptionsLocale = "lij-IT"
	CrawlParamsExtractOptionsLocaleLoLa      CrawlParamsExtractOptionsLocale = "lo-LA"
	CrawlParamsExtractOptionsLocaleLt        CrawlParamsExtractOptionsLocale = "lt"
	CrawlParamsExtractOptionsLocaleLtLt      CrawlParamsExtractOptionsLocale = "lt-LT"
	CrawlParamsExtractOptionsLocaleLuo       CrawlParamsExtractOptionsLocale = "luo"
	CrawlParamsExtractOptionsLocaleLuoKe     CrawlParamsExtractOptionsLocale = "luo-KE"
	CrawlParamsExtractOptionsLocaleLuy       CrawlParamsExtractOptionsLocale = "luy"
	CrawlParamsExtractOptionsLocaleLuyKe     CrawlParamsExtractOptionsLocale = "luy-KE"
	CrawlParamsExtractOptionsLocaleLv        CrawlParamsExtractOptionsLocale = "lv"
	CrawlParamsExtractOptionsLocaleLvLv      CrawlParamsExtractOptionsLocale = "lv-LV"
	CrawlParamsExtractOptionsLocaleMagIn     CrawlParamsExtractOptionsLocale = "mag-IN"
	CrawlParamsExtractOptionsLocaleMaiIn     CrawlParamsExtractOptionsLocale = "mai-IN"
	CrawlParamsExtractOptionsLocaleMas       CrawlParamsExtractOptionsLocale = "mas"
	CrawlParamsExtractOptionsLocaleMasKe     CrawlParamsExtractOptionsLocale = "mas-KE"
	CrawlParamsExtractOptionsLocaleMasTz     CrawlParamsExtractOptionsLocale = "mas-TZ"
	CrawlParamsExtractOptionsLocaleMer       CrawlParamsExtractOptionsLocale = "mer"
	CrawlParamsExtractOptionsLocaleMerKe     CrawlParamsExtractOptionsLocale = "mer-KE"
	CrawlParamsExtractOptionsLocaleMfe       CrawlParamsExtractOptionsLocale = "mfe"
	CrawlParamsExtractOptionsLocaleMfeMu     CrawlParamsExtractOptionsLocale = "mfe-MU"
	CrawlParamsExtractOptionsLocaleMg        CrawlParamsExtractOptionsLocale = "mg"
	CrawlParamsExtractOptionsLocaleMgMg      CrawlParamsExtractOptionsLocale = "mg-MG"
	CrawlParamsExtractOptionsLocaleMhrRu     CrawlParamsExtractOptionsLocale = "mhr-RU"
	CrawlParamsExtractOptionsLocaleMiNz      CrawlParamsExtractOptionsLocale = "mi-NZ"
	CrawlParamsExtractOptionsLocaleMk        CrawlParamsExtractOptionsLocale = "mk"
	CrawlParamsExtractOptionsLocaleMkMk      CrawlParamsExtractOptionsLocale = "mk-MK"
	CrawlParamsExtractOptionsLocaleMl        CrawlParamsExtractOptionsLocale = "ml"
	CrawlParamsExtractOptionsLocaleMlIn      CrawlParamsExtractOptionsLocale = "ml-IN"
	CrawlParamsExtractOptionsLocaleMnMn      CrawlParamsExtractOptionsLocale = "mn-MN"
	CrawlParamsExtractOptionsLocaleMr        CrawlParamsExtractOptionsLocale = "mr"
	CrawlParamsExtractOptionsLocaleMrIn      CrawlParamsExtractOptionsLocale = "mr-IN"
	CrawlParamsExtractOptionsLocaleMs        CrawlParamsExtractOptionsLocale = "ms"
	CrawlParamsExtractOptionsLocaleMsBn      CrawlParamsExtractOptionsLocale = "ms-BN"
	CrawlParamsExtractOptionsLocaleMsMy      CrawlParamsExtractOptionsLocale = "ms-MY"
	CrawlParamsExtractOptionsLocaleMt        CrawlParamsExtractOptionsLocale = "mt"
	CrawlParamsExtractOptionsLocaleMtMt      CrawlParamsExtractOptionsLocale = "mt-MT"
	CrawlParamsExtractOptionsLocaleMy        CrawlParamsExtractOptionsLocale = "my"
	CrawlParamsExtractOptionsLocaleMyMm      CrawlParamsExtractOptionsLocale = "my-MM"
	CrawlParamsExtractOptionsLocaleNanTw     CrawlParamsExtractOptionsLocale = "nan-TW"
	CrawlParamsExtractOptionsLocaleNaq       CrawlParamsExtractOptionsLocale = "naq"
	CrawlParamsExtractOptionsLocaleNaqNa     CrawlParamsExtractOptionsLocale = "naq-NA"
	CrawlParamsExtractOptionsLocaleNb        CrawlParamsExtractOptionsLocale = "nb"
	CrawlParamsExtractOptionsLocaleNbNo      CrawlParamsExtractOptionsLocale = "nb-NO"
	CrawlParamsExtractOptionsLocaleNd        CrawlParamsExtractOptionsLocale = "nd"
	CrawlParamsExtractOptionsLocaleNdZw      CrawlParamsExtractOptionsLocale = "nd-ZW"
	CrawlParamsExtractOptionsLocaleNdsDe     CrawlParamsExtractOptionsLocale = "nds-DE"
	CrawlParamsExtractOptionsLocaleNdsNl     CrawlParamsExtractOptionsLocale = "nds-NL"
	CrawlParamsExtractOptionsLocaleNe        CrawlParamsExtractOptionsLocale = "ne"
	CrawlParamsExtractOptionsLocaleNeIn      CrawlParamsExtractOptionsLocale = "ne-IN"
	CrawlParamsExtractOptionsLocaleNeNp      CrawlParamsExtractOptionsLocale = "ne-NP"
	CrawlParamsExtractOptionsLocaleNl        CrawlParamsExtractOptionsLocale = "nl"
	CrawlParamsExtractOptionsLocaleNlAw      CrawlParamsExtractOptionsLocale = "nl-AW"
	CrawlParamsExtractOptionsLocaleNlBe      CrawlParamsExtractOptionsLocale = "nl-BE"
	CrawlParamsExtractOptionsLocaleNlNl      CrawlParamsExtractOptionsLocale = "nl-NL"
	CrawlParamsExtractOptionsLocaleNn        CrawlParamsExtractOptionsLocale = "nn"
	CrawlParamsExtractOptionsLocaleNnNo      CrawlParamsExtractOptionsLocale = "nn-NO"
	CrawlParamsExtractOptionsLocaleNrZa      CrawlParamsExtractOptionsLocale = "nr-ZA"
	CrawlParamsExtractOptionsLocaleNsoZa     CrawlParamsExtractOptionsLocale = "nso-ZA"
	CrawlParamsExtractOptionsLocaleNyn       CrawlParamsExtractOptionsLocale = "nyn"
	CrawlParamsExtractOptionsLocaleNynUg     CrawlParamsExtractOptionsLocale = "nyn-UG"
	CrawlParamsExtractOptionsLocaleOcFr      CrawlParamsExtractOptionsLocale = "oc-FR"
	CrawlParamsExtractOptionsLocaleOm        CrawlParamsExtractOptionsLocale = "om"
	CrawlParamsExtractOptionsLocaleOmEt      CrawlParamsExtractOptionsLocale = "om-ET"
	CrawlParamsExtractOptionsLocaleOmKe      CrawlParamsExtractOptionsLocale = "om-KE"
	CrawlParamsExtractOptionsLocaleOr        CrawlParamsExtractOptionsLocale = "or"
	CrawlParamsExtractOptionsLocaleOrIn      CrawlParamsExtractOptionsLocale = "or-IN"
	CrawlParamsExtractOptionsLocaleOsRu      CrawlParamsExtractOptionsLocale = "os-RU"
	CrawlParamsExtractOptionsLocalePa        CrawlParamsExtractOptionsLocale = "pa"
	CrawlParamsExtractOptionsLocalePaArab    CrawlParamsExtractOptionsLocale = "pa-Arab"
	CrawlParamsExtractOptionsLocalePaArabPk  CrawlParamsExtractOptionsLocale = "pa-Arab-PK"
	CrawlParamsExtractOptionsLocalePaGuru    CrawlParamsExtractOptionsLocale = "pa-Guru"
	CrawlParamsExtractOptionsLocalePaGuruIn  CrawlParamsExtractOptionsLocale = "pa-Guru-IN"
	CrawlParamsExtractOptionsLocalePaIn      CrawlParamsExtractOptionsLocale = "pa-IN"
	CrawlParamsExtractOptionsLocalePaPk      CrawlParamsExtractOptionsLocale = "pa-PK"
	CrawlParamsExtractOptionsLocalePapAn     CrawlParamsExtractOptionsLocale = "pap-AN"
	CrawlParamsExtractOptionsLocalePl        CrawlParamsExtractOptionsLocale = "pl"
	CrawlParamsExtractOptionsLocalePlPl      CrawlParamsExtractOptionsLocale = "pl-PL"
	CrawlParamsExtractOptionsLocalePs        CrawlParamsExtractOptionsLocale = "ps"
	CrawlParamsExtractOptionsLocalePsAf      CrawlParamsExtractOptionsLocale = "ps-AF"
	CrawlParamsExtractOptionsLocalePt        CrawlParamsExtractOptionsLocale = "pt"
	CrawlParamsExtractOptionsLocalePtBr      CrawlParamsExtractOptionsLocale = "pt-BR"
	CrawlParamsExtractOptionsLocalePtGw      CrawlParamsExtractOptionsLocale = "pt-GW"
	CrawlParamsExtractOptionsLocalePtMz      CrawlParamsExtractOptionsLocale = "pt-MZ"
	CrawlParamsExtractOptionsLocalePtPt      CrawlParamsExtractOptionsLocale = "pt-PT"
	CrawlParamsExtractOptionsLocaleRm        CrawlParamsExtractOptionsLocale = "rm"
	CrawlParamsExtractOptionsLocaleRmCh      CrawlParamsExtractOptionsLocale = "rm-CH"
	CrawlParamsExtractOptionsLocaleRo        CrawlParamsExtractOptionsLocale = "ro"
	CrawlParamsExtractOptionsLocaleRoMd      CrawlParamsExtractOptionsLocale = "ro-MD"
	CrawlParamsExtractOptionsLocaleRoRo      CrawlParamsExtractOptionsLocale = "ro-RO"
	CrawlParamsExtractOptionsLocaleRof       CrawlParamsExtractOptionsLocale = "rof"
	CrawlParamsExtractOptionsLocaleRofTz     CrawlParamsExtractOptionsLocale = "rof-TZ"
	CrawlParamsExtractOptionsLocaleRu        CrawlParamsExtractOptionsLocale = "ru"
	CrawlParamsExtractOptionsLocaleRuMd      CrawlParamsExtractOptionsLocale = "ru-MD"
	CrawlParamsExtractOptionsLocaleRuRu      CrawlParamsExtractOptionsLocale = "ru-RU"
	CrawlParamsExtractOptionsLocaleRuUa      CrawlParamsExtractOptionsLocale = "ru-UA"
	CrawlParamsExtractOptionsLocaleRw        CrawlParamsExtractOptionsLocale = "rw"
	CrawlParamsExtractOptionsLocaleRwRw      CrawlParamsExtractOptionsLocale = "rw-RW"
	CrawlParamsExtractOptionsLocaleRwk       CrawlParamsExtractOptionsLocale = "rwk"
	CrawlParamsExtractOptionsLocaleRwkTz     CrawlParamsExtractOptionsLocale = "rwk-TZ"
	CrawlParamsExtractOptionsLocaleSaIn      CrawlParamsExtractOptionsLocale = "sa-IN"
	CrawlParamsExtractOptionsLocaleSaq       CrawlParamsExtractOptionsLocale = "saq"
	CrawlParamsExtractOptionsLocaleSaqKe     CrawlParamsExtractOptionsLocale = "saq-KE"
	CrawlParamsExtractOptionsLocaleScIt      CrawlParamsExtractOptionsLocale = "sc-IT"
	CrawlParamsExtractOptionsLocaleSdIn      CrawlParamsExtractOptionsLocale = "sd-IN"
	CrawlParamsExtractOptionsLocaleSeNo      CrawlParamsExtractOptionsLocale = "se-NO"
	CrawlParamsExtractOptionsLocaleSeh       CrawlParamsExtractOptionsLocale = "seh"
	CrawlParamsExtractOptionsLocaleSehMz     CrawlParamsExtractOptionsLocale = "seh-MZ"
	CrawlParamsExtractOptionsLocaleSes       CrawlParamsExtractOptionsLocale = "ses"
	CrawlParamsExtractOptionsLocaleSesMl     CrawlParamsExtractOptionsLocale = "ses-ML"
	CrawlParamsExtractOptionsLocaleSg        CrawlParamsExtractOptionsLocale = "sg"
	CrawlParamsExtractOptionsLocaleSgCf      CrawlParamsExtractOptionsLocale = "sg-CF"
	CrawlParamsExtractOptionsLocaleShi       CrawlParamsExtractOptionsLocale = "shi"
	CrawlParamsExtractOptionsLocaleShiLatn   CrawlParamsExtractOptionsLocale = "shi-Latn"
	CrawlParamsExtractOptionsLocaleShiLatnMa CrawlParamsExtractOptionsLocale = "shi-Latn-MA"
	CrawlParamsExtractOptionsLocaleShiTfng   CrawlParamsExtractOptionsLocale = "shi-Tfng"
	CrawlParamsExtractOptionsLocaleShiTfngMa CrawlParamsExtractOptionsLocale = "shi-Tfng-MA"
	CrawlParamsExtractOptionsLocaleShsCa     CrawlParamsExtractOptionsLocale = "shs-CA"
	CrawlParamsExtractOptionsLocaleSi        CrawlParamsExtractOptionsLocale = "si"
	CrawlParamsExtractOptionsLocaleSiLk      CrawlParamsExtractOptionsLocale = "si-LK"
	CrawlParamsExtractOptionsLocaleSidEt     CrawlParamsExtractOptionsLocale = "sid-ET"
	CrawlParamsExtractOptionsLocaleSk        CrawlParamsExtractOptionsLocale = "sk"
	CrawlParamsExtractOptionsLocaleSkSk      CrawlParamsExtractOptionsLocale = "sk-SK"
	CrawlParamsExtractOptionsLocaleSl        CrawlParamsExtractOptionsLocale = "sl"
	CrawlParamsExtractOptionsLocaleSlSi      CrawlParamsExtractOptionsLocale = "sl-SI"
	CrawlParamsExtractOptionsLocaleSn        CrawlParamsExtractOptionsLocale = "sn"
	CrawlParamsExtractOptionsLocaleSnZw      CrawlParamsExtractOptionsLocale = "sn-ZW"
	CrawlParamsExtractOptionsLocaleSo        CrawlParamsExtractOptionsLocale = "so"
	CrawlParamsExtractOptionsLocaleSoDj      CrawlParamsExtractOptionsLocale = "so-DJ"
	CrawlParamsExtractOptionsLocaleSoEt      CrawlParamsExtractOptionsLocale = "so-ET"
	CrawlParamsExtractOptionsLocaleSoKe      CrawlParamsExtractOptionsLocale = "so-KE"
	CrawlParamsExtractOptionsLocaleSoSo      CrawlParamsExtractOptionsLocale = "so-SO"
	CrawlParamsExtractOptionsLocaleSq        CrawlParamsExtractOptionsLocale = "sq"
	CrawlParamsExtractOptionsLocaleSqAl      CrawlParamsExtractOptionsLocale = "sq-AL"
	CrawlParamsExtractOptionsLocaleSqMk      CrawlParamsExtractOptionsLocale = "sq-MK"
	CrawlParamsExtractOptionsLocaleSr        CrawlParamsExtractOptionsLocale = "sr"
	CrawlParamsExtractOptionsLocaleSrCyrl    CrawlParamsExtractOptionsLocale = "sr-Cyrl"
	CrawlParamsExtractOptionsLocaleSrCyrlBa  CrawlParamsExtractOptionsLocale = "sr-Cyrl-BA"
	CrawlParamsExtractOptionsLocaleSrCyrlMe  CrawlParamsExtractOptionsLocale = "sr-Cyrl-ME"
	CrawlParamsExtractOptionsLocaleSrCyrlRs  CrawlParamsExtractOptionsLocale = "sr-Cyrl-RS"
	CrawlParamsExtractOptionsLocaleSrLatn    CrawlParamsExtractOptionsLocale = "sr-Latn"
	CrawlParamsExtractOptionsLocaleSrLatnBa  CrawlParamsExtractOptionsLocale = "sr-Latn-BA"
	CrawlParamsExtractOptionsLocaleSrLatnMe  CrawlParamsExtractOptionsLocale = "sr-Latn-ME"
	CrawlParamsExtractOptionsLocaleSrLatnRs  CrawlParamsExtractOptionsLocale = "sr-Latn-RS"
	CrawlParamsExtractOptionsLocaleSrMe      CrawlParamsExtractOptionsLocale = "sr-ME"
	CrawlParamsExtractOptionsLocaleSrRs      CrawlParamsExtractOptionsLocale = "sr-RS"
	CrawlParamsExtractOptionsLocaleSSZa      CrawlParamsExtractOptionsLocale = "ss-ZA"
	CrawlParamsExtractOptionsLocaleStZa      CrawlParamsExtractOptionsLocale = "st-ZA"
	CrawlParamsExtractOptionsLocaleSv        CrawlParamsExtractOptionsLocale = "sv"
	CrawlParamsExtractOptionsLocaleSvFi      CrawlParamsExtractOptionsLocale = "sv-FI"
	CrawlParamsExtractOptionsLocaleSvSe      CrawlParamsExtractOptionsLocale = "sv-SE"
	CrawlParamsExtractOptionsLocaleSw        CrawlParamsExtractOptionsLocale = "sw"
	CrawlParamsExtractOptionsLocaleSwKe      CrawlParamsExtractOptionsLocale = "sw-KE"
	CrawlParamsExtractOptionsLocaleSwTz      CrawlParamsExtractOptionsLocale = "sw-TZ"
	CrawlParamsExtractOptionsLocaleTa        CrawlParamsExtractOptionsLocale = "ta"
	CrawlParamsExtractOptionsLocaleTaIn      CrawlParamsExtractOptionsLocale = "ta-IN"
	CrawlParamsExtractOptionsLocaleTaLk      CrawlParamsExtractOptionsLocale = "ta-LK"
	CrawlParamsExtractOptionsLocaleTe        CrawlParamsExtractOptionsLocale = "te"
	CrawlParamsExtractOptionsLocaleTeIn      CrawlParamsExtractOptionsLocale = "te-IN"
	CrawlParamsExtractOptionsLocaleTeo       CrawlParamsExtractOptionsLocale = "teo"
	CrawlParamsExtractOptionsLocaleTeoKe     CrawlParamsExtractOptionsLocale = "teo-KE"
	CrawlParamsExtractOptionsLocaleTeoUg     CrawlParamsExtractOptionsLocale = "teo-UG"
	CrawlParamsExtractOptionsLocaleTgTj      CrawlParamsExtractOptionsLocale = "tg-TJ"
	CrawlParamsExtractOptionsLocaleTh        CrawlParamsExtractOptionsLocale = "th"
	CrawlParamsExtractOptionsLocaleThTh      CrawlParamsExtractOptionsLocale = "th-TH"
	CrawlParamsExtractOptionsLocaleTi        CrawlParamsExtractOptionsLocale = "ti"
	CrawlParamsExtractOptionsLocaleTiEr      CrawlParamsExtractOptionsLocale = "ti-ER"
	CrawlParamsExtractOptionsLocaleTiEt      CrawlParamsExtractOptionsLocale = "ti-ET"
	CrawlParamsExtractOptionsLocaleTigEr     CrawlParamsExtractOptionsLocale = "tig-ER"
	CrawlParamsExtractOptionsLocaleTkTm      CrawlParamsExtractOptionsLocale = "tk-TM"
	CrawlParamsExtractOptionsLocaleTlPh      CrawlParamsExtractOptionsLocale = "tl-PH"
	CrawlParamsExtractOptionsLocaleTnZa      CrawlParamsExtractOptionsLocale = "tn-ZA"
	CrawlParamsExtractOptionsLocaleTo        CrawlParamsExtractOptionsLocale = "to"
	CrawlParamsExtractOptionsLocaleToTo      CrawlParamsExtractOptionsLocale = "to-TO"
	CrawlParamsExtractOptionsLocaleTr        CrawlParamsExtractOptionsLocale = "tr"
	CrawlParamsExtractOptionsLocaleTrCy      CrawlParamsExtractOptionsLocale = "tr-CY"
	CrawlParamsExtractOptionsLocaleTrTr      CrawlParamsExtractOptionsLocale = "tr-TR"
	CrawlParamsExtractOptionsLocaleTsZa      CrawlParamsExtractOptionsLocale = "ts-ZA"
	CrawlParamsExtractOptionsLocaleTtRu      CrawlParamsExtractOptionsLocale = "tt-RU"
	CrawlParamsExtractOptionsLocaleTzm       CrawlParamsExtractOptionsLocale = "tzm"
	CrawlParamsExtractOptionsLocaleTzmLatn   CrawlParamsExtractOptionsLocale = "tzm-Latn"
	CrawlParamsExtractOptionsLocaleTzmLatnMa CrawlParamsExtractOptionsLocale = "tzm-Latn-MA"
	CrawlParamsExtractOptionsLocaleUgCn      CrawlParamsExtractOptionsLocale = "ug-CN"
	CrawlParamsExtractOptionsLocaleUk        CrawlParamsExtractOptionsLocale = "uk"
	CrawlParamsExtractOptionsLocaleUkUa      CrawlParamsExtractOptionsLocale = "uk-UA"
	CrawlParamsExtractOptionsLocaleUnmUs     CrawlParamsExtractOptionsLocale = "unm-US"
	CrawlParamsExtractOptionsLocaleUr        CrawlParamsExtractOptionsLocale = "ur"
	CrawlParamsExtractOptionsLocaleUrIn      CrawlParamsExtractOptionsLocale = "ur-IN"
	CrawlParamsExtractOptionsLocaleUrPk      CrawlParamsExtractOptionsLocale = "ur-PK"
	CrawlParamsExtractOptionsLocaleUz        CrawlParamsExtractOptionsLocale = "uz"
	CrawlParamsExtractOptionsLocaleUzArab    CrawlParamsExtractOptionsLocale = "uz-Arab"
	CrawlParamsExtractOptionsLocaleUzArabAf  CrawlParamsExtractOptionsLocale = "uz-Arab-AF"
	CrawlParamsExtractOptionsLocaleUzCyrl    CrawlParamsExtractOptionsLocale = "uz-Cyrl"
	CrawlParamsExtractOptionsLocaleUzCyrlUz  CrawlParamsExtractOptionsLocale = "uz-Cyrl-UZ"
	CrawlParamsExtractOptionsLocaleUzLatn    CrawlParamsExtractOptionsLocale = "uz-Latn"
	CrawlParamsExtractOptionsLocaleUzLatnUz  CrawlParamsExtractOptionsLocale = "uz-Latn-UZ"
	CrawlParamsExtractOptionsLocaleUzUz      CrawlParamsExtractOptionsLocale = "uz-UZ"
	CrawlParamsExtractOptionsLocaleVeZa      CrawlParamsExtractOptionsLocale = "ve-ZA"
	CrawlParamsExtractOptionsLocaleVi        CrawlParamsExtractOptionsLocale = "vi"
	CrawlParamsExtractOptionsLocaleViVn      CrawlParamsExtractOptionsLocale = "vi-VN"
	CrawlParamsExtractOptionsLocaleVun       CrawlParamsExtractOptionsLocale = "vun"
	CrawlParamsExtractOptionsLocaleVunTz     CrawlParamsExtractOptionsLocale = "vun-TZ"
	CrawlParamsExtractOptionsLocaleWaBe      CrawlParamsExtractOptionsLocale = "wa-BE"
	CrawlParamsExtractOptionsLocaleWaeCh     CrawlParamsExtractOptionsLocale = "wae-CH"
	CrawlParamsExtractOptionsLocaleWalEt     CrawlParamsExtractOptionsLocale = "wal-ET"
	CrawlParamsExtractOptionsLocaleWoSn      CrawlParamsExtractOptionsLocale = "wo-SN"
	CrawlParamsExtractOptionsLocaleXhZa      CrawlParamsExtractOptionsLocale = "xh-ZA"
	CrawlParamsExtractOptionsLocaleXog       CrawlParamsExtractOptionsLocale = "xog"
	CrawlParamsExtractOptionsLocaleXogUg     CrawlParamsExtractOptionsLocale = "xog-UG"
	CrawlParamsExtractOptionsLocaleYiUs      CrawlParamsExtractOptionsLocale = "yi-US"
	CrawlParamsExtractOptionsLocaleYo        CrawlParamsExtractOptionsLocale = "yo"
	CrawlParamsExtractOptionsLocaleYoNg      CrawlParamsExtractOptionsLocale = "yo-NG"
	CrawlParamsExtractOptionsLocaleYueHk     CrawlParamsExtractOptionsLocale = "yue-HK"
	CrawlParamsExtractOptionsLocaleZh        CrawlParamsExtractOptionsLocale = "zh"
	CrawlParamsExtractOptionsLocaleZhCn      CrawlParamsExtractOptionsLocale = "zh-CN"
	CrawlParamsExtractOptionsLocaleZhHk      CrawlParamsExtractOptionsLocale = "zh-HK"
	CrawlParamsExtractOptionsLocaleZhHans    CrawlParamsExtractOptionsLocale = "zh-Hans"
	CrawlParamsExtractOptionsLocaleZhHansCn  CrawlParamsExtractOptionsLocale = "zh-Hans-CN"
	CrawlParamsExtractOptionsLocaleZhHansHk  CrawlParamsExtractOptionsLocale = "zh-Hans-HK"
	CrawlParamsExtractOptionsLocaleZhHansMo  CrawlParamsExtractOptionsLocale = "zh-Hans-MO"
	CrawlParamsExtractOptionsLocaleZhHansSg  CrawlParamsExtractOptionsLocale = "zh-Hans-SG"
	CrawlParamsExtractOptionsLocaleZhHant    CrawlParamsExtractOptionsLocale = "zh-Hant"
	CrawlParamsExtractOptionsLocaleZhHantHk  CrawlParamsExtractOptionsLocale = "zh-Hant-HK"
	CrawlParamsExtractOptionsLocaleZhHantMo  CrawlParamsExtractOptionsLocale = "zh-Hant-MO"
	CrawlParamsExtractOptionsLocaleZhHantTw  CrawlParamsExtractOptionsLocale = "zh-Hant-TW"
	CrawlParamsExtractOptionsLocaleZhSg      CrawlParamsExtractOptionsLocale = "zh-SG"
	CrawlParamsExtractOptionsLocaleZhTw      CrawlParamsExtractOptionsLocale = "zh-TW"
	CrawlParamsExtractOptionsLocaleZu        CrawlParamsExtractOptionsLocale = "zu"
	CrawlParamsExtractOptionsLocaleZuZa      CrawlParamsExtractOptionsLocale = "zu-ZA"
	CrawlParamsExtractOptionsLocaleAuto      CrawlParamsExtractOptionsLocale = "auto"
)

// Structured metadata about the request execution context
type CrawlParamsExtractOptionsMetadata struct {
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

func (r CrawlParamsExtractOptionsMetadata) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlParamsExtractOptionsNetworkCapture struct {
	Validation                  param.Opt[bool]    `json:"validation,omitzero"`
	WaitForRequestsCount        param.Opt[float64] `json:"wait_for_requests_count,omitzero"`
	WaitForRequestsCountTimeout param.Opt[float64] `json:"wait_for_requests_count_timeout,omitzero"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method,omitzero"`
	// Resource type for network capture filtering
	ResourceType CrawlParamsExtractOptionsNetworkCaptureResourceTypeUnion `json:"resource_type,omitzero"`
	StatusCode   CrawlParamsExtractOptionsNetworkCaptureStatusCodeUnion   `json:"status_code,omitzero"`
	URL          CrawlParamsExtractOptionsNetworkCaptureURL               `json:"url,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsNetworkCapture) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsNetworkCapture
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsNetworkCapture](
		"method", "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsNetworkCaptureResourceTypeUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsNetworkCaptureResourceTypeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlParamsExtractOptionsNetworkCaptureResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsNetworkCaptureResourceTypeUnion) asAny() any {
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
type CrawlParamsExtractOptionsNetworkCaptureStatusCodeUnion struct {
	OfFloat      param.Opt[float64] `json:",omitzero,inline"`
	OfFloatArray []float64          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsNetworkCaptureStatusCodeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfFloatArray)
}
func (u *CrawlParamsExtractOptionsNetworkCaptureStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsNetworkCaptureStatusCodeUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfFloatArray) {
		return &u.OfFloatArray
	}
	return nil
}

// The property Value is required.
type CrawlParamsExtractOptionsNetworkCaptureURL struct {
	Value string `json:"value,required"`
	// Any of "exact", "contains".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsNetworkCaptureURL) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsNetworkCaptureURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsNetworkCaptureURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsNetworkCaptureURL](
		"type", "exact", "contains",
	)
}

// Configuration options for parsing behavior
type CrawlParamsExtractOptionsParseOptions struct {
	// Whether to merge dynamic parsing results with static results
	MergeDynamic param.Opt[bool] `json:"merge_dynamic,omitzero"`
	ExtraFields  map[string]any  `json:"-"`
	paramObj
}

func (r CrawlParamsExtractOptionsParseOptions) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsParseOptions
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *CrawlParamsExtractOptionsParseOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsParserUnion struct {
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsParserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap, u.OfString)
}
func (u *CrawlParamsExtractOptionsParserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsParserUnion) asAny() any {
	if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Proxy provider to use for the request
type CrawlParamsExtractOptionsProxyProvider string

const (
	CrawlParamsExtractOptionsProxyProviderBrightdata      CrawlParamsExtractOptionsProxyProvider = "brightdata"
	CrawlParamsExtractOptionsProxyProviderOxylabs         CrawlParamsExtractOptionsProxyProvider = "oxylabs"
	CrawlParamsExtractOptionsProxyProviderSmartproxy      CrawlParamsExtractOptionsProxyProvider = "smartproxy"
	CrawlParamsExtractOptionsProxyProviderProxit          CrawlParamsExtractOptionsProxyProvider = "proxit"
	CrawlParamsExtractOptionsProxyProviderProxitPreprod   CrawlParamsExtractOptionsProxyProvider = "proxit_preprod"
	CrawlParamsExtractOptionsProxyProviderLocal           CrawlParamsExtractOptionsProxyProvider = "local"
	CrawlParamsExtractOptionsProxyProviderRayobyte        CrawlParamsExtractOptionsProxyProvider = "rayobyte"
	CrawlParamsExtractOptionsProxyProviderAlways          CrawlParamsExtractOptionsProxyProvider = "always"
	CrawlParamsExtractOptionsProxyProviderOculusproxies   CrawlParamsExtractOptionsProxyProvider = "oculusproxies"
	CrawlParamsExtractOptionsProxyProviderFroxy           CrawlParamsExtractOptionsProxyProvider = "froxy"
	CrawlParamsExtractOptionsProxyProviderPacketstream    CrawlParamsExtractOptionsProxyProvider = "packetstream"
	CrawlParamsExtractOptionsProxyProvider911proxy        CrawlParamsExtractOptionsProxyProvider = "911proxy"
	CrawlParamsExtractOptionsProxyProviderDirect911proxy  CrawlParamsExtractOptionsProxyProvider = "direct911proxy"
	CrawlParamsExtractOptionsProxyProviderThesocialproxy  CrawlParamsExtractOptionsProxyProvider = "thesocialproxy"
	CrawlParamsExtractOptionsProxyProviderThesocialproxy2 CrawlParamsExtractOptionsProxyProvider = "thesocialproxy2"
	CrawlParamsExtractOptionsProxyProviderNimbleIsp       CrawlParamsExtractOptionsProxyProvider = "nimble-isp"
	CrawlParamsExtractOptionsProxyProviderNimbleIspMobile CrawlParamsExtractOptionsProxyProvider = "nimble-isp-mobile"
	CrawlParamsExtractOptionsProxyProviderProxitLinux     CrawlParamsExtractOptionsProxyProvider = "proxit-linux"
	CrawlParamsExtractOptionsProxyProviderProxitMacos     CrawlParamsExtractOptionsProxyProvider = "proxit-macos"
	CrawlParamsExtractOptionsProxyProviderProxitWindows   CrawlParamsExtractOptionsProxyProvider = "proxit-windows"
	CrawlParamsExtractOptionsProxyProviderProxitRental    CrawlParamsExtractOptionsProxyProvider = "proxit-rental"
	CrawlParamsExtractOptionsProxyProviderIpfoxy          CrawlParamsExtractOptionsProxyProvider = "ipfoxy"
	CrawlParamsExtractOptionsProxyProviderBrightup        CrawlParamsExtractOptionsProxyProvider = "brightup"
	CrawlParamsExtractOptionsProxyProviderResearch        CrawlParamsExtractOptionsProxyProvider = "research"
)

// Query template configuration for structured data extraction
//
// The property ID is required.
type CrawlParamsExtractOptionsQueryTemplate struct {
	ID string `json:"id,required" format:"uuid"`
	// Any of "WEB", "SERP", "SOCIAL".
	APIType     string                                                `json:"api_type,omitzero"`
	Pagination  CrawlParamsExtractOptionsQueryTemplatePaginationUnion `json:"pagination,omitzero"`
	Params      map[string]any                                        `json:"params,omitzero"`
	ExtraFields map[string]any                                        `json:"-"`
	paramObj
}

func (r CrawlParamsExtractOptionsQueryTemplate) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsQueryTemplate
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *CrawlParamsExtractOptionsQueryTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsQueryTemplate](
		"api_type", "WEB", "SERP", "SOCIAL",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsQueryTemplatePaginationUnion struct {
	OfCrawlsExtractOptionsQueryTemplatePaginationNextPageParams *CrawlParamsExtractOptionsQueryTemplatePaginationNextPageParams `json:",omitzero,inline"`
	OfCrawlsExtractOptionsQueryTemplatePaginationArray          []CrawlParamsExtractOptionsQueryTemplatePaginationArrayItem     `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsQueryTemplatePaginationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsQueryTemplatePaginationNextPageParams, u.OfCrawlsExtractOptionsQueryTemplatePaginationArray)
}
func (u *CrawlParamsExtractOptionsQueryTemplatePaginationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsQueryTemplatePaginationUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsQueryTemplatePaginationNextPageParams) {
		return u.OfCrawlsExtractOptionsQueryTemplatePaginationNextPageParams
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsQueryTemplatePaginationArray) {
		return &u.OfCrawlsExtractOptionsQueryTemplatePaginationArray
	}
	return nil
}

// The property NextPageParams is required.
type CrawlParamsExtractOptionsQueryTemplatePaginationNextPageParams struct {
	NextPageParams map[string]any `json:"next_page_params,omitzero,required"`
	paramObj
}

func (r CrawlParamsExtractOptionsQueryTemplatePaginationNextPageParams) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsQueryTemplatePaginationNextPageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsQueryTemplatePaginationNextPageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property NextPageParams is required.
type CrawlParamsExtractOptionsQueryTemplatePaginationArrayItem struct {
	NextPageParams map[string]any `json:"next_page_params,omitzero,required"`
	paramObj
}

func (r CrawlParamsExtractOptionsQueryTemplatePaginationArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsQueryTemplatePaginationArrayItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsQueryTemplatePaginationArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Referrer policy for the request
type CrawlParamsExtractOptionsReferrerType string

const (
	CrawlParamsExtractOptionsReferrerTypeRandom     CrawlParamsExtractOptionsReferrerType = "random"
	CrawlParamsExtractOptionsReferrerTypeNoReferer  CrawlParamsExtractOptionsReferrerType = "no-referer"
	CrawlParamsExtractOptionsReferrerTypeSameOrigin CrawlParamsExtractOptionsReferrerType = "same-origin"
	CrawlParamsExtractOptionsReferrerTypeGoogle     CrawlParamsExtractOptionsReferrerType = "google"
	CrawlParamsExtractOptionsReferrerTypeBing       CrawlParamsExtractOptionsReferrerType = "bing"
	CrawlParamsExtractOptionsReferrerTypeFacebook   CrawlParamsExtractOptionsReferrerType = "facebook"
	CrawlParamsExtractOptionsReferrerTypeTwitter    CrawlParamsExtractOptionsReferrerType = "twitter"
	CrawlParamsExtractOptionsReferrerTypeInstagram  CrawlParamsExtractOptionsReferrerType = "instagram"
)

type CrawlParamsExtractOptionsRenderOptions struct {
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
	BrowserEngine CrawlParamsExtractOptionsRenderOptionsBrowserEngineUnion `json:"browser_engine,omitzero"`
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
	HackiumConfiguration CrawlParamsExtractOptionsRenderOptionsHackiumConfiguration `json:"hackium_configuration,omitzero"`
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

func (r CrawlParamsExtractOptionsRenderOptions) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsRenderOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsRenderOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsRenderOptions](
		"connector_type", "puppeteer", "puppeteer-cdp", "puppeteer-bidi", "webit-cdp", "playwright",
	)
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsRenderOptions](
		"mouse_strategy", "linear", "ghost-cursor", "windmouse",
	)
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsRenderOptions](
		"render_type", "domcontentloaded", "load", "idle0", "networkidle0", "idle2", "networkidle2", "navigate",
	)
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsRenderOptions](
		"typing_strategy", "simple", "distribution",
	)
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsRenderOptions](
		"wait_until", "load", "domcontentloaded", "idle0", "idle2", "networkidle0", "networkidle2", "navigate",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsRenderOptionsBrowserEngineUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsRenderOptionsBrowserEngineString)
	OfCrawlsExtractOptionsRenderOptionsBrowserEngineString param.Opt[string]  `json:",omitzero,inline"`
	OfFloatMap                                             map[string]float64 `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsRenderOptionsBrowserEngineUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsRenderOptionsBrowserEngineString, u.OfFloatMap)
}
func (u *CrawlParamsExtractOptionsRenderOptionsBrowserEngineUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsRenderOptionsBrowserEngineUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsRenderOptionsBrowserEngineString) {
		return &u.OfCrawlsExtractOptionsRenderOptionsBrowserEngineString
	} else if !param.IsOmitted(u.OfFloatMap) {
		return &u.OfFloatMap
	}
	return nil
}

type CrawlParamsExtractOptionsRenderOptionsBrowserEngineString string

const (
	CrawlParamsExtractOptionsRenderOptionsBrowserEngineStringChrome  CrawlParamsExtractOptionsRenderOptionsBrowserEngineString = "chrome"
	CrawlParamsExtractOptionsRenderOptionsBrowserEngineStringHackium CrawlParamsExtractOptionsRenderOptionsBrowserEngineString = "hackium"
	CrawlParamsExtractOptionsRenderOptionsBrowserEngineStringFirefox CrawlParamsExtractOptionsRenderOptionsBrowserEngineString = "firefox"
	CrawlParamsExtractOptionsRenderOptionsBrowserEngineStringHackfox CrawlParamsExtractOptionsRenderOptionsBrowserEngineString = "hackfox"
)

// Configuration for Hackium browser modifications
type CrawlParamsExtractOptionsRenderOptionsHackiumConfiguration struct {
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

func (r CrawlParamsExtractOptionsRenderOptionsHackiumConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsRenderOptionsHackiumConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsRenderOptionsHackiumConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlParamsExtractOptionsSession struct {
	ID                  param.Opt[string]  `json:"id,omitzero"`
	PrefetchUserbrowser param.Opt[bool]    `json:"prefetch_userbrowser,omitzero"`
	Retry               param.Opt[bool]    `json:"retry,omitzero"`
	Timeout             param.Opt[float64] `json:"timeout,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsSession) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsSession
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsSkillUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsSkillUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlParamsExtractOptionsSkillUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsSkillUnion) asAny() any {
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
type CrawlParamsExtractOptionsTemplate struct {
	Name   string         `json:"name,required"`
	Params map[string]any `json:"params,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsTemplate) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsTemplate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pre-rendered userbrowser creation template configuration
//
// The properties ID, AllowedParameterNames, RenderFlowRendered are required.
type CrawlParamsExtractOptionsUserbrowserCreationTemplateRendered struct {
	ID                    string           `json:"id,required"`
	AllowedParameterNames []string         `json:"allowed_parameter_names,omitzero,required"`
	RenderFlowRendered    []map[string]any `json:"render_flow_rendered,omitzero,required"`
	paramObj
}

func (r CrawlParamsExtractOptionsUserbrowserCreationTemplateRendered) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsUserbrowserCreationTemplateRendered
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsUserbrowserCreationTemplateRendered) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sitemap and other methods will be used together to find URLs.
type CrawlParamsSitemap string

const (
	CrawlParamsSitemapSkip    CrawlParamsSitemap = "skip"
	CrawlParamsSitemapInclude CrawlParamsSitemap = "include"
	CrawlParamsSitemapOnly    CrawlParamsSitemap = "only"
)

type ExtractParams struct {
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
	// Whether to save the userbrowser session for reuse
	SaveUserbrowser param.Opt[bool] `json:"save_userbrowser,omitzero"`
	// Whether to skip userbrowser creation template processing
	SkipUbct param.Opt[bool] `json:"skip_ubct,omitzero"`
	// User-defined tag for request identification
	Tag param.Opt[string] `json:"tag,omitzero"`
	// Type of query or scraping template
	Type param.Opt[string] `json:"type,omitzero"`
	// Browser type to emulate
	Browser ExtractParamsBrowserUnion `json:"browser,omitzero"`
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
	// Custom parser configuration as a key-value map
	DynamicParser map[string]any `json:"dynamic_parser,omitzero"`
	// Expected HTTP status codes for successful requests
	ExpectedStatusCodes []int64 `json:"expected_status_codes,omitzero"`
	// Response format
	//
	// Any of "json", "html", "csv", "raw", "json-lines", "markdown".
	Format ExtractParamsFormat `json:"format,omitzero"`
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
	// Structured metadata about the request execution context
	Metadata ExtractParamsMetadata `json:"metadata,omitzero"`
	// HTTP method for the request
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method ExtractParamsMethod `json:"method,omitzero"`
	// Native execution mode
	//
	// Any of "requester", "apm", "direct".
	NativeMode ExtractParamsNativeMode `json:"native_mode,omitzero"`
	// Filters for capturing network traffic
	NetworkCapture []ExtractParamsNetworkCapture `json:"network_capture,omitzero"`
	// Operating system to emulate
	//
	// Any of "windows", "mac os", "linux", "android", "ios".
	Os ExtractParamsOs `json:"os,omitzero"`
	// Configuration options for parsing behavior
	ParseOptions ExtractParamsParseOptions `json:"parse_options,omitzero"`
	// Custom parser configuration as a key-value map
	Parser ExtractParamsParserUnion `json:"parser,omitzero"`
	// Proxy provider to use for the request
	//
	// Any of "brightdata", "oxylabs", "smartproxy", "proxit", "proxit_preprod",
	// "local", "rayobyte", "always", "oculusproxies", "froxy", "packetstream",
	// "911proxy", "direct911proxy", "thesocialproxy", "thesocialproxy2", "nimble-isp",
	// "nimble-isp-mobile", "proxit-linux", "proxit-macos", "proxit-windows",
	// "proxit-rental", "ipfoxy", "brightup", "research".
	ProxyProvider ExtractParamsProxyProvider `json:"proxy_provider,omitzero"`
	// Weighted distribution of proxy providers
	ProxyProviders map[string]float64 `json:"proxy_providers,omitzero"`
	// Query template configuration for structured data extraction
	QueryTemplate ExtractParamsQueryTemplate `json:"query_template,omitzero"`
	// Referrer policy for the request
	//
	// Any of "random", "no-referer", "same-origin", "google", "bing", "facebook",
	// "twitter", "instagram".
	ReferrerType ExtractParamsReferrerType `json:"referrer_type,omitzero"`
	// Array of actions to perform during browser rendering
	RenderFlow    []map[string]any           `json:"render_flow,omitzero"`
	RenderOptions ExtractParamsRenderOptions `json:"render_options,omitzero"`
	Session       ExtractParamsSession       `json:"session,omitzero"`
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
	// Userbrowser creation template configuration
	Template ExtractParamsTemplate `json:"template,omitzero"`
	// Pre-rendered userbrowser creation template configuration
	UserbrowserCreationTemplateRendered ExtractParamsUserbrowserCreationTemplateRendered `json:"userbrowser_creation_template_rendered,omitzero"`
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

// Response format
type ExtractParamsFormat string

const (
	ExtractParamsFormatJson      ExtractParamsFormat = "json"
	ExtractParamsFormatHTML      ExtractParamsFormat = "html"
	ExtractParamsFormatCsv       ExtractParamsFormat = "csv"
	ExtractParamsFormatRaw       ExtractParamsFormat = "raw"
	ExtractParamsFormatJsonLines ExtractParamsFormat = "json-lines"
	ExtractParamsFormatMarkdown  ExtractParamsFormat = "markdown"
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

// Structured metadata about the request execution context
type ExtractParamsMetadata struct {
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

func (r ExtractParamsMetadata) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP method for the request
type ExtractParamsMethod string

const (
	ExtractParamsMethodGet    ExtractParamsMethod = "GET"
	ExtractParamsMethodPost   ExtractParamsMethod = "POST"
	ExtractParamsMethodPut    ExtractParamsMethod = "PUT"
	ExtractParamsMethodPatch  ExtractParamsMethod = "PATCH"
	ExtractParamsMethodDelete ExtractParamsMethod = "DELETE"
)

// Native execution mode
type ExtractParamsNativeMode string

const (
	ExtractParamsNativeModeRequester ExtractParamsNativeMode = "requester"
	ExtractParamsNativeModeApm       ExtractParamsNativeMode = "apm"
	ExtractParamsNativeModeDirect    ExtractParamsNativeMode = "direct"
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

// Configuration options for parsing behavior
type ExtractParamsParseOptions struct {
	// Whether to merge dynamic parsing results with static results
	MergeDynamic param.Opt[bool] `json:"merge_dynamic,omitzero"`
	ExtraFields  map[string]any  `json:"-"`
	paramObj
}

func (r ExtractParamsParseOptions) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsParseOptions
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ExtractParamsParseOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

// Proxy provider to use for the request
type ExtractParamsProxyProvider string

const (
	ExtractParamsProxyProviderBrightdata      ExtractParamsProxyProvider = "brightdata"
	ExtractParamsProxyProviderOxylabs         ExtractParamsProxyProvider = "oxylabs"
	ExtractParamsProxyProviderSmartproxy      ExtractParamsProxyProvider = "smartproxy"
	ExtractParamsProxyProviderProxit          ExtractParamsProxyProvider = "proxit"
	ExtractParamsProxyProviderProxitPreprod   ExtractParamsProxyProvider = "proxit_preprod"
	ExtractParamsProxyProviderLocal           ExtractParamsProxyProvider = "local"
	ExtractParamsProxyProviderRayobyte        ExtractParamsProxyProvider = "rayobyte"
	ExtractParamsProxyProviderAlways          ExtractParamsProxyProvider = "always"
	ExtractParamsProxyProviderOculusproxies   ExtractParamsProxyProvider = "oculusproxies"
	ExtractParamsProxyProviderFroxy           ExtractParamsProxyProvider = "froxy"
	ExtractParamsProxyProviderPacketstream    ExtractParamsProxyProvider = "packetstream"
	ExtractParamsProxyProvider911proxy        ExtractParamsProxyProvider = "911proxy"
	ExtractParamsProxyProviderDirect911proxy  ExtractParamsProxyProvider = "direct911proxy"
	ExtractParamsProxyProviderThesocialproxy  ExtractParamsProxyProvider = "thesocialproxy"
	ExtractParamsProxyProviderThesocialproxy2 ExtractParamsProxyProvider = "thesocialproxy2"
	ExtractParamsProxyProviderNimbleIsp       ExtractParamsProxyProvider = "nimble-isp"
	ExtractParamsProxyProviderNimbleIspMobile ExtractParamsProxyProvider = "nimble-isp-mobile"
	ExtractParamsProxyProviderProxitLinux     ExtractParamsProxyProvider = "proxit-linux"
	ExtractParamsProxyProviderProxitMacos     ExtractParamsProxyProvider = "proxit-macos"
	ExtractParamsProxyProviderProxitWindows   ExtractParamsProxyProvider = "proxit-windows"
	ExtractParamsProxyProviderProxitRental    ExtractParamsProxyProvider = "proxit-rental"
	ExtractParamsProxyProviderIpfoxy          ExtractParamsProxyProvider = "ipfoxy"
	ExtractParamsProxyProviderBrightup        ExtractParamsProxyProvider = "brightup"
	ExtractParamsProxyProviderResearch        ExtractParamsProxyProvider = "research"
)

// Query template configuration for structured data extraction
//
// The property ID is required.
type ExtractParamsQueryTemplate struct {
	ID string `json:"id,required" format:"uuid"`
	// Any of "WEB", "SERP", "SOCIAL".
	APIType     string                                    `json:"api_type,omitzero"`
	Pagination  ExtractParamsQueryTemplatePaginationUnion `json:"pagination,omitzero"`
	Params      map[string]any                            `json:"params,omitzero"`
	ExtraFields map[string]any                            `json:"-"`
	paramObj
}

func (r ExtractParamsQueryTemplate) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsQueryTemplate
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ExtractParamsQueryTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractParamsQueryTemplate](
		"api_type", "WEB", "SERP", "SOCIAL",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsQueryTemplatePaginationUnion struct {
	OfExtractsQueryTemplatePaginationNextPageParams *ExtractParamsQueryTemplatePaginationNextPageParams `json:",omitzero,inline"`
	OfExtractsQueryTemplatePaginationArray          []ExtractParamsQueryTemplatePaginationArrayItem     `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsQueryTemplatePaginationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsQueryTemplatePaginationNextPageParams, u.OfExtractsQueryTemplatePaginationArray)
}
func (u *ExtractParamsQueryTemplatePaginationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsQueryTemplatePaginationUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsQueryTemplatePaginationNextPageParams) {
		return u.OfExtractsQueryTemplatePaginationNextPageParams
	} else if !param.IsOmitted(u.OfExtractsQueryTemplatePaginationArray) {
		return &u.OfExtractsQueryTemplatePaginationArray
	}
	return nil
}

// The property NextPageParams is required.
type ExtractParamsQueryTemplatePaginationNextPageParams struct {
	NextPageParams map[string]any `json:"next_page_params,omitzero,required"`
	paramObj
}

func (r ExtractParamsQueryTemplatePaginationNextPageParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsQueryTemplatePaginationNextPageParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsQueryTemplatePaginationNextPageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property NextPageParams is required.
type ExtractParamsQueryTemplatePaginationArrayItem struct {
	NextPageParams map[string]any `json:"next_page_params,omitzero,required"`
	paramObj
}

func (r ExtractParamsQueryTemplatePaginationArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsQueryTemplatePaginationArrayItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsQueryTemplatePaginationArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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

type ExtractParamsRenderOptions struct {
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
	BrowserEngine ExtractParamsRenderOptionsBrowserEngineUnion `json:"browser_engine,omitzero"`
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
	HackiumConfiguration ExtractParamsRenderOptionsHackiumConfiguration `json:"hackium_configuration,omitzero"`
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

func (r ExtractParamsRenderOptions) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsRenderOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsRenderOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractParamsRenderOptions](
		"connector_type", "puppeteer", "puppeteer-cdp", "puppeteer-bidi", "webit-cdp", "playwright",
	)
	apijson.RegisterFieldValidator[ExtractParamsRenderOptions](
		"mouse_strategy", "linear", "ghost-cursor", "windmouse",
	)
	apijson.RegisterFieldValidator[ExtractParamsRenderOptions](
		"render_type", "domcontentloaded", "load", "idle0", "networkidle0", "idle2", "networkidle2", "navigate",
	)
	apijson.RegisterFieldValidator[ExtractParamsRenderOptions](
		"typing_strategy", "simple", "distribution",
	)
	apijson.RegisterFieldValidator[ExtractParamsRenderOptions](
		"wait_until", "load", "domcontentloaded", "idle0", "idle2", "networkidle0", "networkidle2", "navigate",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractParamsRenderOptionsBrowserEngineUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractsRenderOptionsBrowserEngineString)
	OfExtractsRenderOptionsBrowserEngineString param.Opt[string]  `json:",omitzero,inline"`
	OfFloatMap                                 map[string]float64 `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractParamsRenderOptionsBrowserEngineUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractsRenderOptionsBrowserEngineString, u.OfFloatMap)
}
func (u *ExtractParamsRenderOptionsBrowserEngineUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractParamsRenderOptionsBrowserEngineUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractsRenderOptionsBrowserEngineString) {
		return &u.OfExtractsRenderOptionsBrowserEngineString
	} else if !param.IsOmitted(u.OfFloatMap) {
		return &u.OfFloatMap
	}
	return nil
}

type ExtractParamsRenderOptionsBrowserEngineString string

const (
	ExtractParamsRenderOptionsBrowserEngineStringChrome  ExtractParamsRenderOptionsBrowserEngineString = "chrome"
	ExtractParamsRenderOptionsBrowserEngineStringHackium ExtractParamsRenderOptionsBrowserEngineString = "hackium"
	ExtractParamsRenderOptionsBrowserEngineStringFirefox ExtractParamsRenderOptionsBrowserEngineString = "firefox"
	ExtractParamsRenderOptionsBrowserEngineStringHackfox ExtractParamsRenderOptionsBrowserEngineString = "hackfox"
)

// Configuration for Hackium browser modifications
type ExtractParamsRenderOptionsHackiumConfiguration struct {
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

func (r ExtractParamsRenderOptionsHackiumConfiguration) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsRenderOptionsHackiumConfiguration
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsRenderOptionsHackiumConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

// Userbrowser creation template configuration
//
// The property Name is required.
type ExtractParamsTemplate struct {
	Name   string         `json:"name,required"`
	Params map[string]any `json:"params,omitzero"`
	paramObj
}

func (r ExtractParamsTemplate) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsTemplate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pre-rendered userbrowser creation template configuration
//
// The properties ID, AllowedParameterNames, RenderFlowRendered are required.
type ExtractParamsUserbrowserCreationTemplateRendered struct {
	ID                    string           `json:"id,required"`
	AllowedParameterNames []string         `json:"allowed_parameter_names,omitzero,required"`
	RenderFlowRendered    []map[string]any `json:"render_flow_rendered,omitzero,required"`
	paramObj
}

func (r ExtractParamsUserbrowserCreationTemplateRendered) MarshalJSON() (data []byte, err error) {
	type shadow ExtractParamsUserbrowserCreationTemplateRendered
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractParamsUserbrowserCreationTemplateRendered) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type SearchParams struct {
	// Search query string
	Query string `json:"query,required"`
	// Filter results before this date (format: YYYY-MM-DD or YYYY)
	EndDate param.Opt[string] `json:"end_date,omitzero"`
	// Filter results after this date (format: YYYY-MM-DD or YYYY)
	StartDate param.Opt[string] `json:"start_date,omitzero"`
	Country   param.Opt[string] `json:"country,omitzero"`
	// If True, fetches and extracts full page content for each search result. If
	// False, returns only metadata (title, snippet, URL)
	DeepSearch param.Opt[bool] `json:"deep_search,omitzero"`
	// Generate LLM answer summary based on search result snippets (works with both
	// deep_search=True and False)
	IncludeAnswer param.Opt[bool]   `json:"include_answer,omitzero"`
	Locale        param.Opt[string] `json:"locale,omitzero"`
	// Maximum number of subagents to execute in parallel for WSA focus modes
	// (shopping, social, geo). Ignored for traditional SERP focus modes. Default: 3,
	// Range: 1-5.
	MaxSubagents param.Opt[int64] `json:"max_subagents,omitzero"`
	// Maximum number of results to return (actual count may be less)
	NumResults param.Opt[int64] `json:"num_results,omitzero"`
	// Filter by content type (only supported with focus=general). Supports semantic
	// groups ('documents', 'spreadsheets', 'presentations') and specific formats
	// ('pdf', 'docx', 'xlsx', etc.)
	ContentType []string `json:"content_type,omitzero"`
	// List of domains to exclude from search results. Maximum 50 domains.
	ExcludeDomains []string `json:"exclude_domains,omitzero"`
	// List of domains to include in search results. Maximum 50 domains.
	IncludeDomains []string `json:"include_domains,omitzero"`
	// Enum representing the search engines supported by Nimble ⚠️ DEPRECATED: This
	// parameter is ignored. Use 'focus' parameter instead.
	//
	// Any of "google_search", "google_sge", "bing_search", "yandex_search".
	SearchEngine SearchParamsSearchEngine `json:"search_engine,omitzero"`
	// Time range filters passed to Webit SERP API as 'time' parameter.
	//
	// Any of "hour", "day", "week", "month", "year".
	TimeRange SearchParamsTimeRange `json:"time_range,omitzero"`
	// Output format: plain_text, markdown, or simplified_html
	//
	// Any of "plain_text", "markdown", "simplified_html".
	ParsingType SearchParamsParsingType `json:"parsing_type,omitzero"`
	// Search focus/specialization (general, news, or location)
	//
	// Any of "general", "news", "location", "coding", "academic", "geo", "shopping",
	// "social".
	Topic SearchParamsTopic `json:"topic,omitzero"`
	paramObj
}

func (r SearchParams) MarshalJSON() (data []byte, err error) {
	type shadow SearchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SearchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Output format: plain_text, markdown, or simplified_html
type SearchParamsParsingType string

const (
	SearchParamsParsingTypePlainText      SearchParamsParsingType = "plain_text"
	SearchParamsParsingTypeMarkdown       SearchParamsParsingType = "markdown"
	SearchParamsParsingTypeSimplifiedHTML SearchParamsParsingType = "simplified_html"
)

// Enum representing the search engines supported by Nimble ⚠️ DEPRECATED: This
// parameter is ignored. Use 'focus' parameter instead.
type SearchParamsSearchEngine string

const (
	SearchParamsSearchEngineGoogleSearch SearchParamsSearchEngine = "google_search"
	SearchParamsSearchEngineGoogleSge    SearchParamsSearchEngine = "google_sge"
	SearchParamsSearchEngineBingSearch   SearchParamsSearchEngine = "bing_search"
	SearchParamsSearchEngineYandexSearch SearchParamsSearchEngine = "yandex_search"
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

// Search focus/specialization (general, news, or location)
type SearchParamsTopic string

const (
	SearchParamsTopicGeneral  SearchParamsTopic = "general"
	SearchParamsTopicNews     SearchParamsTopic = "news"
	SearchParamsTopicLocation SearchParamsTopic = "location"
	SearchParamsTopicCoding   SearchParamsTopic = "coding"
	SearchParamsTopicAcademic SearchParamsTopic = "academic"
	SearchParamsTopicGeo      SearchParamsTopic = "geo"
	SearchParamsTopicShopping SearchParamsTopic = "shopping"
	SearchParamsTopicSocial   SearchParamsTopic = "social"
)
