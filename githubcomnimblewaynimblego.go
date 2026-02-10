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
	// Browser actions execution results. Present only when browser_actions were
	// specified in the request.
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

// Browser actions execution results. Present only when browser_actions were
// specified in the request.
type AgentResponseDataBrowserActions struct {
	Results       []AgentResponseDataBrowserActionsResult `json:"results,required"`
	Success       bool                                    `json:"success,required"`
	TotalDuration float64                                 `json:"total_duration,required"`
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
func (r AgentResponseDataBrowserActions) RawJSON() string { return r.JSON.raw }
func (r *AgentResponseDataBrowserActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentResponseDataBrowserActionsResult struct {
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
func (r AgentResponseDataBrowserActionsResult) RawJSON() string { return r.JSON.raw }
func (r *AgentResponseDataBrowserActionsResult) UnmarshalJSON(data []byte) error {
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
	// Array of browser automation actions to execute sequentially
	BrowserActions []CrawlParamsExtractOptionsBrowserActionUnion `json:"browser_actions,omitzero"`
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
	// Expected HTTP status codes for successful requests
	ExpectedStatusCodes []int64 `json:"expected_status_codes,omitzero"`
	// List of acceptable response formats in order of preference
	//
	// Any of "html", "markdown".
	Formats []string `json:"formats,omitzero"`
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
type CrawlParamsExtractOptionsBrowserActionUnion struct {
	OfCrawlsExtractOptionsBrowserActionAutoScrollAction        *CrawlParamsExtractOptionsBrowserActionAutoScrollAction        `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionClickAction             *CrawlParamsExtractOptionsBrowserActionClickAction             `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionEvalAction              *CrawlParamsExtractOptionsBrowserActionEvalAction              `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionFetchAction             *CrawlParamsExtractOptionsBrowserActionFetchAction             `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionFillAction              *CrawlParamsExtractOptionsBrowserActionFillAction              `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionGetCookiesAction        *CrawlParamsExtractOptionsBrowserActionGetCookiesAction        `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionGotoAction              *CrawlParamsExtractOptionsBrowserActionGotoAction              `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionPressAction             *CrawlParamsExtractOptionsBrowserActionPressAction             `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionScreenshotAction        *CrawlParamsExtractOptionsBrowserActionScreenshotAction        `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionScrollAction            *CrawlParamsExtractOptionsBrowserActionScrollAction            `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionWaitAction              *CrawlParamsExtractOptionsBrowserActionWaitAction              `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionWaitForElementAction    *CrawlParamsExtractOptionsBrowserActionWaitForElementAction    `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionWaitForNavigationAction *CrawlParamsExtractOptionsBrowserActionWaitForNavigationAction `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionAutoScrollAction,
		u.OfCrawlsExtractOptionsBrowserActionClickAction,
		u.OfCrawlsExtractOptionsBrowserActionEvalAction,
		u.OfCrawlsExtractOptionsBrowserActionFetchAction,
		u.OfCrawlsExtractOptionsBrowserActionFillAction,
		u.OfCrawlsExtractOptionsBrowserActionGetCookiesAction,
		u.OfCrawlsExtractOptionsBrowserActionGotoAction,
		u.OfCrawlsExtractOptionsBrowserActionPressAction,
		u.OfCrawlsExtractOptionsBrowserActionScreenshotAction,
		u.OfCrawlsExtractOptionsBrowserActionScrollAction,
		u.OfCrawlsExtractOptionsBrowserActionWaitAction,
		u.OfCrawlsExtractOptionsBrowserActionWaitForElementAction,
		u.OfCrawlsExtractOptionsBrowserActionWaitForNavigationAction)
}
func (u *CrawlParamsExtractOptionsBrowserActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionAutoScrollAction) {
		return u.OfCrawlsExtractOptionsBrowserActionAutoScrollAction
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionClickAction) {
		return u.OfCrawlsExtractOptionsBrowserActionClickAction
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionEvalAction) {
		return u.OfCrawlsExtractOptionsBrowserActionEvalAction
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionFetchAction) {
		return u.OfCrawlsExtractOptionsBrowserActionFetchAction
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionFillAction) {
		return u.OfCrawlsExtractOptionsBrowserActionFillAction
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionGetCookiesAction) {
		return u.OfCrawlsExtractOptionsBrowserActionGetCookiesAction
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionGotoAction) {
		return u.OfCrawlsExtractOptionsBrowserActionGotoAction
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionPressAction) {
		return u.OfCrawlsExtractOptionsBrowserActionPressAction
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionScreenshotAction) {
		return u.OfCrawlsExtractOptionsBrowserActionScreenshotAction
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionScrollAction) {
		return u.OfCrawlsExtractOptionsBrowserActionScrollAction
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionWaitAction) {
		return u.OfCrawlsExtractOptionsBrowserActionWaitAction
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionWaitForElementAction) {
		return u.OfCrawlsExtractOptionsBrowserActionWaitForElementAction
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionWaitForNavigationAction) {
		return u.OfCrawlsExtractOptionsBrowserActionWaitForNavigationAction
	}
	return nil
}

// Continuously scroll to load dynamic content
//
// The property AutoScroll is required.
type CrawlParamsExtractOptionsBrowserActionAutoScrollAction struct {
	AutoScroll CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollUnion `json:"auto_scroll,omitzero,required"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionAutoScrollAction) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionAutoScrollAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionAutoScrollAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollUnion struct {
	OfBool                                                              param.Opt[bool]                                                         `json:",omitzero,inline"`
	OfFloat                                                             param.Opt[float64]                                                      `json:",omitzero,inline"`
	OfString                                                            param.Opt[string]                                                       `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionAutoScrollActionAutoScrollObject *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObject `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfCrawlsExtractOptionsBrowserActionAutoScrollActionAutoScrollObject)
}
func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionAutoScrollActionAutoScrollObject) {
		return u.OfCrawlsExtractOptionsBrowserActionAutoScrollActionAutoScrollObject
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObject struct {
	StepSize param.Opt[float64] `json:"step_size,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	ClickSelector CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion `json:"click_selector,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Container CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion `json:"container,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	DelayAfterScroll CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion `json:"delay_after_scroll,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	IdleTimeout CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion `json:"idle_timeout,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	LoadingSelector CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion `json:"loading_selector,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	MaxDuration CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion `json:"max_duration,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	PauseOnSelector CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion `json:"pause_on_selector,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectRequiredString)
	OfCrawlsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectRequiredString param.Opt[CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                            param.Opt[bool]                                                                                 `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectRequiredString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectRequiredString) {
		return &u.OfCrawlsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectRequiredString string

const (
	CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectRequiredStringTrue  CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectRequiredString = "true"
	CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectRequiredStringFalse CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectSkipString)
	OfCrawlsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectSkipString param.Opt[CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                        param.Opt[bool]                                                                             `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectSkipString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectSkipString) {
		return &u.OfCrawlsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectSkipString string

const (
	CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectSkipStringTrue  CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectSkipString = "true"
	CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectSkipStringFalse CrawlParamsExtractOptionsBrowserActionAutoScrollActionAutoScrollObjectSkipString = "false"
)

// Click on an element by selector
//
// The property Click is required.
type CrawlParamsExtractOptionsBrowserActionClickAction struct {
	Click CrawlParamsExtractOptionsBrowserActionClickActionClickUnion `json:"click,omitzero,required"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionClickAction) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionClickAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionClickAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionClickActionClickUnion struct {
	OfString                                                  param.Opt[string]                                             `json:",omitzero,inline"`
	OfStringArray                                             []string                                                      `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionClickActionClickObject *CrawlParamsExtractOptionsBrowserActionClickActionClickObject `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionClickActionClickUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray, u.OfCrawlsExtractOptionsBrowserActionClickActionClickObject)
}
func (u *CrawlParamsExtractOptionsBrowserActionClickActionClickUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionClickActionClickUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionClickActionClickObject) {
		return u.OfCrawlsExtractOptionsBrowserActionClickActionClickObject
	}
	return nil
}

// The property Selector is required.
type CrawlParamsExtractOptionsBrowserActionClickActionClickObject struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector CrawlParamsExtractOptionsBrowserActionClickActionClickObjectSelectorUnion `json:"selector,omitzero,required"`
	Count    param.Opt[float64]                                                        `json:"count,omitzero"`
	OffsetX  param.Opt[int64]                                                          `json:"offset_x,omitzero"`
	OffsetY  param.Opt[int64]                                                          `json:"offset_y,omitzero"`
	Scroll   param.Opt[bool]                                                           `json:"scroll,omitzero"`
	Steps    param.Opt[float64]                                                        `json:"steps,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay CrawlParamsExtractOptionsBrowserActionClickActionClickObjectDelayUnion `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required CrawlParamsExtractOptionsBrowserActionClickActionClickObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip CrawlParamsExtractOptionsBrowserActionClickActionClickObjectSkipUnion `json:"skip,omitzero"`
	// Any of "linear", "ghost-cursor", "windmouse".
	Strategy string `json:"strategy,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionClickActionClickObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionClickActionClickObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionClickActionClickObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsBrowserActionClickActionClickObject](
		"strategy", "linear", "ghost-cursor", "windmouse",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionClickActionClickObjectSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionClickActionClickObjectSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlParamsExtractOptionsBrowserActionClickActionClickObjectSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionClickActionClickObjectSelectorUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionClickActionClickObjectDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionClickActionClickObjectDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *CrawlParamsExtractOptionsBrowserActionClickActionClickObjectDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionClickActionClickObjectDelayUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionClickActionClickObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionClickActionClickObjectRequiredString)
	OfCrawlsExtractOptionsBrowserActionClickActionClickObjectRequiredString param.Opt[CrawlParamsExtractOptionsBrowserActionClickActionClickObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                  param.Opt[bool]                                                                       `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionClickActionClickObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionClickActionClickObjectRequiredString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionClickActionClickObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionClickActionClickObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionClickActionClickObjectRequiredString) {
		return &u.OfCrawlsExtractOptionsBrowserActionClickActionClickObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionClickActionClickObjectRequiredString string

const (
	CrawlParamsExtractOptionsBrowserActionClickActionClickObjectRequiredStringTrue  CrawlParamsExtractOptionsBrowserActionClickActionClickObjectRequiredString = "true"
	CrawlParamsExtractOptionsBrowserActionClickActionClickObjectRequiredStringFalse CrawlParamsExtractOptionsBrowserActionClickActionClickObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionClickActionClickObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionClickActionClickObjectSkipString)
	OfCrawlsExtractOptionsBrowserActionClickActionClickObjectSkipString param.Opt[CrawlParamsExtractOptionsBrowserActionClickActionClickObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                              param.Opt[bool]                                                                   `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionClickActionClickObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionClickActionClickObjectSkipString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionClickActionClickObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionClickActionClickObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionClickActionClickObjectSkipString) {
		return &u.OfCrawlsExtractOptionsBrowserActionClickActionClickObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionClickActionClickObjectSkipString string

const (
	CrawlParamsExtractOptionsBrowserActionClickActionClickObjectSkipStringTrue  CrawlParamsExtractOptionsBrowserActionClickActionClickObjectSkipString = "true"
	CrawlParamsExtractOptionsBrowserActionClickActionClickObjectSkipStringFalse CrawlParamsExtractOptionsBrowserActionClickActionClickObjectSkipString = "false"
)

// Execute JavaScript code in page context
//
// The property Eval is required.
type CrawlParamsExtractOptionsBrowserActionEvalAction struct {
	Eval CrawlParamsExtractOptionsBrowserActionEvalActionEvalUnion `json:"eval,omitzero,required"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionEvalAction) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionEvalAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionEvalAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionEvalActionEvalUnion struct {
	OfString                                                param.Opt[string]                                           `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionEvalActionEvalObject *CrawlParamsExtractOptionsBrowserActionEvalActionEvalObject `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionEvalActionEvalUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfCrawlsExtractOptionsBrowserActionEvalActionEvalObject)
}
func (u *CrawlParamsExtractOptionsBrowserActionEvalActionEvalUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionEvalActionEvalUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionEvalActionEvalObject) {
		return u.OfCrawlsExtractOptionsBrowserActionEvalActionEvalObject
	}
	return nil
}

// The property Code is required.
type CrawlParamsExtractOptionsBrowserActionEvalActionEvalObject struct {
	Code string `json:"code,required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionEvalActionEvalObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionEvalActionEvalObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionEvalActionEvalObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionEvalActionEvalObjectRequiredString)
	OfCrawlsExtractOptionsBrowserActionEvalActionEvalObjectRequiredString param.Opt[CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                param.Opt[bool]                                                                     `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionEvalActionEvalObjectRequiredString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionEvalActionEvalObjectRequiredString) {
		return &u.OfCrawlsExtractOptionsBrowserActionEvalActionEvalObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectRequiredString string

const (
	CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectRequiredStringTrue  CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectRequiredString = "true"
	CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectRequiredStringFalse CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionEvalActionEvalObjectSkipString)
	OfCrawlsExtractOptionsBrowserActionEvalActionEvalObjectSkipString param.Opt[CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                            param.Opt[bool]                                                                 `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionEvalActionEvalObjectSkipString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionEvalActionEvalObjectSkipString) {
		return &u.OfCrawlsExtractOptionsBrowserActionEvalActionEvalObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectSkipString string

const (
	CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectSkipStringTrue  CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectSkipString = "true"
	CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectSkipStringFalse CrawlParamsExtractOptionsBrowserActionEvalActionEvalObjectSkipString = "false"
)

// Make an HTTP request in browser context
//
// The property Fetch is required.
type CrawlParamsExtractOptionsBrowserActionFetchAction struct {
	Fetch CrawlParamsExtractOptionsBrowserActionFetchActionFetchUnion `json:"fetch,omitzero,required" format:"uri"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionFetchAction) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionFetchAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionFetchAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionFetchActionFetchUnion struct {
	OfString                                                  param.Opt[string]                                             `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionFetchActionFetchObject *CrawlParamsExtractOptionsBrowserActionFetchActionFetchObject `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionFetchActionFetchUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfCrawlsExtractOptionsBrowserActionFetchActionFetchObject)
}
func (u *CrawlParamsExtractOptionsBrowserActionFetchActionFetchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionFetchActionFetchUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionFetchActionFetchObject) {
		return u.OfCrawlsExtractOptionsBrowserActionFetchActionFetchObject
	}
	return nil
}

// The property URL is required.
type CrawlParamsExtractOptionsBrowserActionFetchActionFetchObject struct {
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
	Required CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionFetchActionFetchObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionFetchActionFetchObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionFetchActionFetchObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsBrowserActionFetchActionFetchObject](
		"method", "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionFetchActionFetchObjectRequiredString)
	OfCrawlsExtractOptionsBrowserActionFetchActionFetchObjectRequiredString param.Opt[CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                  param.Opt[bool]                                                                       `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionFetchActionFetchObjectRequiredString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionFetchActionFetchObjectRequiredString) {
		return &u.OfCrawlsExtractOptionsBrowserActionFetchActionFetchObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectRequiredString string

const (
	CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectRequiredStringTrue  CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectRequiredString = "true"
	CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectRequiredStringFalse CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionFetchActionFetchObjectSkipString)
	OfCrawlsExtractOptionsBrowserActionFetchActionFetchObjectSkipString param.Opt[CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                              param.Opt[bool]                                                                   `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionFetchActionFetchObjectSkipString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionFetchActionFetchObjectSkipString) {
		return &u.OfCrawlsExtractOptionsBrowserActionFetchActionFetchObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectSkipString string

const (
	CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectSkipStringTrue  CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectSkipString = "true"
	CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectSkipStringFalse CrawlParamsExtractOptionsBrowserActionFetchActionFetchObjectSkipString = "false"
)

// Fill text into an input field
//
// The property Fill is required.
type CrawlParamsExtractOptionsBrowserActionFillAction struct {
	// Fill options with mode-specific fields. Use "type" mode for behavioral typing
	// simulation, or "paste" mode for instant paste.
	Fill CrawlParamsExtractOptionsBrowserActionFillActionFillUnion `json:"fill,omitzero,required"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionFillAction) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionFillAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionFillAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionFillActionFillUnion struct {
	OfType  *CrawlParamsExtractOptionsBrowserActionFillActionFillType  `json:",omitzero,inline"`
	OfPaste *CrawlParamsExtractOptionsBrowserActionFillActionFillPaste `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionFillActionFillUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfType, u.OfPaste)
}
func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillUnion) asAny() any {
	if !param.IsOmitted(u.OfType) {
		return u.OfType
	} else if !param.IsOmitted(u.OfPaste) {
		return u.OfPaste
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CrawlParamsExtractOptionsBrowserActionFillActionFillUnion) GetMouseMovementStrategy() *string {
	if vt := u.OfType; vt != nil {
		return &vt.MouseMovementStrategy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CrawlParamsExtractOptionsBrowserActionFillActionFillUnion) GetTypingInterval() *CrawlParamsExtractOptionsBrowserActionFillActionFillTypeTypingIntervalUnion {
	if vt := u.OfType; vt != nil {
		return &vt.TypingInterval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CrawlParamsExtractOptionsBrowserActionFillActionFillUnion) GetTypingStrategy() *string {
	if vt := u.OfType; vt != nil {
		return &vt.TypingStrategy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CrawlParamsExtractOptionsBrowserActionFillActionFillUnion) GetValue() *string {
	if vt := u.OfType; vt != nil {
		return (*string)(&vt.Value)
	} else if vt := u.OfPaste; vt != nil {
		return (*string)(&vt.Value)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CrawlParamsExtractOptionsBrowserActionFillActionFillUnion) GetClickOnElement() *bool {
	if vt := u.OfType; vt != nil && vt.ClickOnElement.Valid() {
		return &vt.ClickOnElement.Value
	} else if vt := u.OfPaste; vt != nil && vt.ClickOnElement.Valid() {
		return &vt.ClickOnElement.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CrawlParamsExtractOptionsBrowserActionFillActionFillUnion) GetMode() *string {
	if vt := u.OfType; vt != nil {
		return (*string)(&vt.Mode)
	} else if vt := u.OfPaste; vt != nil {
		return (*string)(&vt.Mode)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CrawlParamsExtractOptionsBrowserActionFillActionFillUnion) GetScroll() *bool {
	if vt := u.OfType; vt != nil && vt.Scroll.Valid() {
		return &vt.Scroll.Value
	} else if vt := u.OfPaste; vt != nil && vt.Scroll.Valid() {
		return &vt.Scroll.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CrawlParamsExtractOptionsBrowserActionFillActionFillUnion) GetTimeout() *float64 {
	if vt := u.OfType; vt != nil && vt.Timeout.Valid() {
		return &vt.Timeout.Value
	} else if vt := u.OfPaste; vt != nil && vt.Timeout.Valid() {
		return &vt.Timeout.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CrawlParamsExtractOptionsBrowserActionFillActionFillUnion) GetVisible() *bool {
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
func (u CrawlParamsExtractOptionsBrowserActionFillActionFillUnion) GetSelector() (res crawlParamsExtractOptionsBrowserActionFillActionFillUnionSelector) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Selector.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Selector.asAny()
	}
	return
}

// Can have the runtime types [*string], [\*[]string]
type crawlParamsExtractOptionsBrowserActionFillActionFillUnionSelector struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u crawlParamsExtractOptionsBrowserActionFillActionFillUnionSelector) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u CrawlParamsExtractOptionsBrowserActionFillActionFillUnion) GetDelay() (res crawlParamsExtractOptionsBrowserActionFillActionFillUnionDelay) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Delay.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Delay.asAny()
	}
	return
}

// Can have the runtime types [*float64], [*string]
type crawlParamsExtractOptionsBrowserActionFillActionFillUnionDelay struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *float64:
//	case *string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u crawlParamsExtractOptionsBrowserActionFillActionFillUnionDelay) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u CrawlParamsExtractOptionsBrowserActionFillActionFillUnion) GetRequired() (res crawlParamsExtractOptionsBrowserActionFillActionFillUnionRequired) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Required.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Required.asAny()
	}
	return
}

// Can have the runtime types [*string], [*bool]
type crawlParamsExtractOptionsBrowserActionFillActionFillUnionRequired struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u crawlParamsExtractOptionsBrowserActionFillActionFillUnionRequired) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u CrawlParamsExtractOptionsBrowserActionFillActionFillUnion) GetSkip() (res crawlParamsExtractOptionsBrowserActionFillActionFillUnionSkip) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Skip.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Skip.asAny()
	}
	return
}

// Can have the runtime types [*string], [*bool]
type crawlParamsExtractOptionsBrowserActionFillActionFillUnionSkip struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u crawlParamsExtractOptionsBrowserActionFillActionFillUnionSkip) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[CrawlParamsExtractOptionsBrowserActionFillActionFillUnion](
		"mode",
		apijson.Discriminator[CrawlParamsExtractOptionsBrowserActionFillActionFillType]("type"),
		apijson.Discriminator[CrawlParamsExtractOptionsBrowserActionFillActionFillPaste]("paste"),
	)
}

// The properties Selector, Value are required.
type CrawlParamsExtractOptionsBrowserActionFillActionFillType struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector       CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSelectorUnion `json:"selector,omitzero,required"`
	Value          string                                                                `json:"value,required"`
	ClickOnElement param.Opt[bool]                                                       `json:"click_on_element,omitzero"`
	Scroll         param.Opt[bool]                                                       `json:"scroll,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay CrawlParamsExtractOptionsBrowserActionFillActionFillTypeDelayUnion `json:"delay,omitzero"`
	// Any of "type".
	Mode string `json:"mode,omitzero"`
	// Any of "linear", "ghost-cursor", "windmouse".
	MouseMovementStrategy string `json:"mouse_movement_strategy,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required CrawlParamsExtractOptionsBrowserActionFillActionFillTypeRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSkipUnion `json:"skip,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	TypingInterval CrawlParamsExtractOptionsBrowserActionFillActionFillTypeTypingIntervalUnion `json:"typing_interval,omitzero"`
	// Any of "simple", "distribution".
	TypingStrategy string `json:"typing_strategy,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionFillActionFillType) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionFillActionFillType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionFillActionFillType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsBrowserActionFillActionFillType](
		"mode", "type",
	)
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsBrowserActionFillActionFillType](
		"mouse_movement_strategy", "linear", "ghost-cursor", "windmouse",
	)
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsBrowserActionFillActionFillType](
		"typing_strategy", "simple", "distribution",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSelectorUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionFillActionFillTypeDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionFillActionFillTypeDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillTypeDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillTypeDelayUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionFillActionFillTypeRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionFillActionFillTypeRequiredString)
	OfCrawlsExtractOptionsBrowserActionFillActionFillTypeRequiredString param.Opt[CrawlParamsExtractOptionsBrowserActionFillActionFillTypeRequiredString] `json:",omitzero,inline"`
	OfBool                                                              param.Opt[bool]                                                                   `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionFillActionFillTypeRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionFillActionFillTypeRequiredString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillTypeRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillTypeRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionFillActionFillTypeRequiredString) {
		return &u.OfCrawlsExtractOptionsBrowserActionFillActionFillTypeRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionFillActionFillTypeRequiredString string

const (
	CrawlParamsExtractOptionsBrowserActionFillActionFillTypeRequiredStringTrue  CrawlParamsExtractOptionsBrowserActionFillActionFillTypeRequiredString = "true"
	CrawlParamsExtractOptionsBrowserActionFillActionFillTypeRequiredStringFalse CrawlParamsExtractOptionsBrowserActionFillActionFillTypeRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionFillActionFillTypeSkipString)
	OfCrawlsExtractOptionsBrowserActionFillActionFillTypeSkipString param.Opt[CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSkipString] `json:",omitzero,inline"`
	OfBool                                                          param.Opt[bool]                                                               `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionFillActionFillTypeSkipString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionFillActionFillTypeSkipString) {
		return &u.OfCrawlsExtractOptionsBrowserActionFillActionFillTypeSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSkipString string

const (
	CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSkipStringTrue  CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSkipString = "true"
	CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSkipStringFalse CrawlParamsExtractOptionsBrowserActionFillActionFillTypeSkipString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionFillActionFillTypeTypingIntervalUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionFillActionFillTypeTypingIntervalUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillTypeTypingIntervalUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillTypeTypingIntervalUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// The properties Mode, Selector, Value are required.
type CrawlParamsExtractOptionsBrowserActionFillActionFillPaste struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector       CrawlParamsExtractOptionsBrowserActionFillActionFillPasteSelectorUnion `json:"selector,omitzero,required"`
	Value          string                                                                 `json:"value,required"`
	ClickOnElement param.Opt[bool]                                                        `json:"click_on_element,omitzero"`
	Scroll         param.Opt[bool]                                                        `json:"scroll,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay CrawlParamsExtractOptionsBrowserActionFillActionFillPasteDelayUnion `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required CrawlParamsExtractOptionsBrowserActionFillActionFillPasteRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip CrawlParamsExtractOptionsBrowserActionFillActionFillPasteSkipUnion `json:"skip,omitzero"`
	// This field can be elided, and will marshal its zero value as "paste".
	Mode constant.Paste `json:"mode,required"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionFillActionFillPaste) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionFillActionFillPaste
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionFillActionFillPaste) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionFillActionFillPasteSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionFillActionFillPasteSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillPasteSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillPasteSelectorUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionFillActionFillPasteDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionFillActionFillPasteDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillPasteDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillPasteDelayUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionFillActionFillPasteRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionFillActionFillPasteRequiredString)
	OfCrawlsExtractOptionsBrowserActionFillActionFillPasteRequiredString param.Opt[CrawlParamsExtractOptionsBrowserActionFillActionFillPasteRequiredString] `json:",omitzero,inline"`
	OfBool                                                               param.Opt[bool]                                                                    `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionFillActionFillPasteRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionFillActionFillPasteRequiredString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillPasteRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillPasteRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionFillActionFillPasteRequiredString) {
		return &u.OfCrawlsExtractOptionsBrowserActionFillActionFillPasteRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionFillActionFillPasteRequiredString string

const (
	CrawlParamsExtractOptionsBrowserActionFillActionFillPasteRequiredStringTrue  CrawlParamsExtractOptionsBrowserActionFillActionFillPasteRequiredString = "true"
	CrawlParamsExtractOptionsBrowserActionFillActionFillPasteRequiredStringFalse CrawlParamsExtractOptionsBrowserActionFillActionFillPasteRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionFillActionFillPasteSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionFillActionFillPasteSkipString)
	OfCrawlsExtractOptionsBrowserActionFillActionFillPasteSkipString param.Opt[CrawlParamsExtractOptionsBrowserActionFillActionFillPasteSkipString] `json:",omitzero,inline"`
	OfBool                                                           param.Opt[bool]                                                                `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionFillActionFillPasteSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionFillActionFillPasteSkipString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillPasteSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionFillActionFillPasteSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionFillActionFillPasteSkipString) {
		return &u.OfCrawlsExtractOptionsBrowserActionFillActionFillPasteSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionFillActionFillPasteSkipString string

const (
	CrawlParamsExtractOptionsBrowserActionFillActionFillPasteSkipStringTrue  CrawlParamsExtractOptionsBrowserActionFillActionFillPasteSkipString = "true"
	CrawlParamsExtractOptionsBrowserActionFillActionFillPasteSkipStringFalse CrawlParamsExtractOptionsBrowserActionFillActionFillPasteSkipString = "false"
)

// Retrieve browser cookies
//
// The property GetCookies is required.
type CrawlParamsExtractOptionsBrowserActionGetCookiesAction struct {
	GetCookies CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesUnion `json:"get_cookies,omitzero,required"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionGetCookiesAction) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionGetCookiesAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionGetCookiesAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesUnion struct {
	OfBool                                                              param.Opt[bool]                                                         `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionGetCookiesActionGetCookiesObject *CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObject `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfCrawlsExtractOptionsBrowserActionGetCookiesActionGetCookiesObject)
}
func (u *CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionGetCookiesActionGetCookiesObject) {
		return u.OfCrawlsExtractOptionsBrowserActionGetCookiesActionGetCookiesObject
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObject struct {
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip        CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion `json:"skip,omitzero"`
	ExtraFields map[string]any                                                                  `json:"-"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObject
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectRequiredString)
	OfCrawlsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectRequiredString param.Opt[CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                            param.Opt[bool]                                                                                 `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectRequiredString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectRequiredString) {
		return &u.OfCrawlsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectRequiredString string

const (
	CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectRequiredStringTrue  CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectRequiredString = "true"
	CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectRequiredStringFalse CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectSkipString)
	OfCrawlsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectSkipString param.Opt[CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                        param.Opt[bool]                                                                             `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectSkipString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectSkipString) {
		return &u.OfCrawlsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectSkipString string

const (
	CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectSkipStringTrue  CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectSkipString = "true"
	CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectSkipStringFalse CrawlParamsExtractOptionsBrowserActionGetCookiesActionGetCookiesObjectSkipString = "false"
)

// Navigate to a URL
//
// The property Goto is required.
type CrawlParamsExtractOptionsBrowserActionGotoAction struct {
	Goto CrawlParamsExtractOptionsBrowserActionGotoActionGotoUnion `json:"goto,omitzero,required" format:"uri"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionGotoAction) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionGotoAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionGotoAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionGotoActionGotoUnion struct {
	OfString                                                param.Opt[string]                                           `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionGotoActionGotoObject *CrawlParamsExtractOptionsBrowserActionGotoActionGotoObject `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionGotoActionGotoUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfCrawlsExtractOptionsBrowserActionGotoActionGotoObject)
}
func (u *CrawlParamsExtractOptionsBrowserActionGotoActionGotoUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionGotoActionGotoUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionGotoActionGotoObject) {
		return u.OfCrawlsExtractOptionsBrowserActionGotoActionGotoObject
	}
	return nil
}

// The property URL is required.
type CrawlParamsExtractOptionsBrowserActionGotoActionGotoObject struct {
	URL     string            `json:"url,required" format:"uri"`
	Referer param.Opt[string] `json:"referer,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectSkipUnion `json:"skip,omitzero"`
	// Any of "load", "domcontentloaded", "networkidle0", "networkidle2".
	WaitUntil string `json:"wait_until,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionGotoActionGotoObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionGotoActionGotoObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionGotoActionGotoObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsBrowserActionGotoActionGotoObject](
		"wait_until", "load", "domcontentloaded", "networkidle0", "networkidle2",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionGotoActionGotoObjectRequiredString)
	OfCrawlsExtractOptionsBrowserActionGotoActionGotoObjectRequiredString param.Opt[CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                param.Opt[bool]                                                                     `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionGotoActionGotoObjectRequiredString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionGotoActionGotoObjectRequiredString) {
		return &u.OfCrawlsExtractOptionsBrowserActionGotoActionGotoObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectRequiredString string

const (
	CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectRequiredStringTrue  CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectRequiredString = "true"
	CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectRequiredStringFalse CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionGotoActionGotoObjectSkipString)
	OfCrawlsExtractOptionsBrowserActionGotoActionGotoObjectSkipString param.Opt[CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                            param.Opt[bool]                                                                 `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionGotoActionGotoObjectSkipString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionGotoActionGotoObjectSkipString) {
		return &u.OfCrawlsExtractOptionsBrowserActionGotoActionGotoObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectSkipString string

const (
	CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectSkipStringTrue  CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectSkipString = "true"
	CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectSkipStringFalse CrawlParamsExtractOptionsBrowserActionGotoActionGotoObjectSkipString = "false"
)

// Press a keyboard key
//
// The property Press is required.
type CrawlParamsExtractOptionsBrowserActionPressAction struct {
	Press CrawlParamsExtractOptionsBrowserActionPressActionPressUnion `json:"press,omitzero,required"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionPressAction) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionPressAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionPressAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionPressActionPressUnion struct {
	OfString                                                  param.Opt[string]                                             `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionPressActionPressObject *CrawlParamsExtractOptionsBrowserActionPressActionPressObject `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionPressActionPressUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfCrawlsExtractOptionsBrowserActionPressActionPressObject)
}
func (u *CrawlParamsExtractOptionsBrowserActionPressActionPressUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionPressActionPressUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionPressActionPressObject) {
		return u.OfCrawlsExtractOptionsBrowserActionPressActionPressObject
	}
	return nil
}

// The property Key is required.
type CrawlParamsExtractOptionsBrowserActionPressActionPressObject struct {
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
	Delay CrawlParamsExtractOptionsBrowserActionPressActionPressObjectDelayUnion `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required CrawlParamsExtractOptionsBrowserActionPressActionPressObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip CrawlParamsExtractOptionsBrowserActionPressActionPressObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionPressActionPressObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionPressActionPressObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionPressActionPressObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsBrowserActionPressActionPressObject](
		"key", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "Power", "Eject", "Abort", "Help", "Backspace", "Tab", "Numpad5", "NumpadEnter", "Enter", "\r", "\n", "ShiftLeft", "ShiftRight", "ControlLeft", "ControlRight", "AltLeft", "AltRight", "Pause", "CapsLock", "Escape", "Convert", "NonConvert", "Space", "Numpad9", "PageUp", "Numpad3", "PageDown", "End", "Numpad1", "Home", "Numpad7", "ArrowLeft", "Numpad4", "Numpad8", "ArrowUp", "ArrowRight", "Numpad6", "Numpad2", "ArrowDown", "Select", "Open", "PrintScreen", "Insert", "Numpad0", "Delete", "NumpadDecimal", "Digit0", "Digit1", "Digit2", "Digit3", "Digit4", "Digit5", "Digit6", "Digit7", "Digit8", "Digit9", "KeyA", "KeyB", "KeyC", "KeyD", "KeyE", "KeyF", "KeyG", "KeyH", "KeyI", "KeyJ", "KeyK", "KeyL", "KeyM", "KeyN", "KeyO", "KeyP", "KeyQ", "KeyR", "KeyS", "KeyT", "KeyU", "KeyV", "KeyW", "KeyX", "KeyY", "KeyZ", "MetaLeft", "MetaRight", "ContextMenu", "NumpadMultiply", "NumpadAdd", "NumpadSubtract", "NumpadDivide", "F1", "F2", "F3", "F4", "F5", "F6", "F7", "F8", "F9", "F10", "F11", "F12", "F13", "F14", "F15", "F16", "F17", "F18", "F19", "F20", "F21", "F22", "F23", "F24", "NumLock", "ScrollLock", "AudioVolumeMute", "AudioVolumeDown", "AudioVolumeUp", "MediaTrackNext", "MediaTrackPrevious", "MediaStop", "MediaPlayPause", "Semicolon", "Equal", "NumpadEqual", "Comma", "Minus", "Period", "Slash", "Backquote", "BracketLeft", "Backslash", "BracketRight", "Quote", "AltGraph", "Props", "Cancel", "Clear", "Shift", "Control", "Alt", "Accept", "ModeChange", " ", "Print", "Execute", "\u0000", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "Meta", "*", "+", "-", "/", ";", "=", ",", ".", "`", "[", "\\", "]", "'", "Attn", "CrSel", "ExSel", "EraseEof", "Play", "ZoomOut", ")", "!", "@", "#", "$", "%", "^", "&", "(", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", ":", "<", "_", ">", "?", "~", "{", "|", "}", "\"", "SoftLeft", "SoftRight", "Camera", "Call", "EndCall", "VolumeDown", "VolumeUp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionPressActionPressObjectDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionPressActionPressObjectDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *CrawlParamsExtractOptionsBrowserActionPressActionPressObjectDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionPressActionPressObjectDelayUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionPressActionPressObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionPressActionPressObjectRequiredString)
	OfCrawlsExtractOptionsBrowserActionPressActionPressObjectRequiredString param.Opt[CrawlParamsExtractOptionsBrowserActionPressActionPressObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                  param.Opt[bool]                                                                       `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionPressActionPressObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionPressActionPressObjectRequiredString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionPressActionPressObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionPressActionPressObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionPressActionPressObjectRequiredString) {
		return &u.OfCrawlsExtractOptionsBrowserActionPressActionPressObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionPressActionPressObjectRequiredString string

const (
	CrawlParamsExtractOptionsBrowserActionPressActionPressObjectRequiredStringTrue  CrawlParamsExtractOptionsBrowserActionPressActionPressObjectRequiredString = "true"
	CrawlParamsExtractOptionsBrowserActionPressActionPressObjectRequiredStringFalse CrawlParamsExtractOptionsBrowserActionPressActionPressObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionPressActionPressObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionPressActionPressObjectSkipString)
	OfCrawlsExtractOptionsBrowserActionPressActionPressObjectSkipString param.Opt[CrawlParamsExtractOptionsBrowserActionPressActionPressObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                              param.Opt[bool]                                                                   `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionPressActionPressObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionPressActionPressObjectSkipString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionPressActionPressObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionPressActionPressObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionPressActionPressObjectSkipString) {
		return &u.OfCrawlsExtractOptionsBrowserActionPressActionPressObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionPressActionPressObjectSkipString string

const (
	CrawlParamsExtractOptionsBrowserActionPressActionPressObjectSkipStringTrue  CrawlParamsExtractOptionsBrowserActionPressActionPressObjectSkipString = "true"
	CrawlParamsExtractOptionsBrowserActionPressActionPressObjectSkipStringFalse CrawlParamsExtractOptionsBrowserActionPressActionPressObjectSkipString = "false"
)

// Capture a page screenshot
//
// The property Screenshot is required.
type CrawlParamsExtractOptionsBrowserActionScreenshotAction struct {
	Screenshot CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotUnion `json:"screenshot,omitzero,required"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionScreenshotAction) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionScreenshotAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionScreenshotAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotUnion struct {
	OfBool                                                              param.Opt[bool]                                                         `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObject *CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObject `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObject)
}
func (u *CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObject) {
		return u.OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObject
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObject struct {
	FullPage param.Opt[bool]    `json:"full_page,omitzero"`
	Quality  param.Opt[float64] `json:"quality,omitzero"`
	// Any of "png", "jpeg", "webp".
	Format string `json:"format,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObject](
		"format", "png", "jpeg", "webp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredString)
	OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredString param.Opt[CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                            param.Opt[bool]                                                                                 `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredString) {
		return &u.OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredString string

const (
	CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredStringTrue  CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredString = "true"
	CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredStringFalse CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipString)
	OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipString param.Opt[CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                        param.Opt[bool]                                                                             `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipString) {
		return &u.OfCrawlsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipString string

const (
	CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipStringTrue  CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipString = "true"
	CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipStringFalse CrawlParamsExtractOptionsBrowserActionScreenshotActionScreenshotObjectSkipString = "false"
)

// Scroll the page or an element
//
// The property Scroll is required.
type CrawlParamsExtractOptionsBrowserActionScrollAction struct {
	Scroll CrawlParamsExtractOptionsBrowserActionScrollActionScrollUnion `json:"scroll,omitzero,required"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionScrollAction) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionScrollAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionScrollAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionScrollActionScrollUnion struct {
	OfFloat                                                     param.Opt[float64]                                              `json:",omitzero,inline"`
	OfString                                                    param.Opt[string]                                               `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionScrollActionScrollObject *CrawlParamsExtractOptionsBrowserActionScrollActionScrollObject `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionScrollActionScrollUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfCrawlsExtractOptionsBrowserActionScrollActionScrollObject)
}
func (u *CrawlParamsExtractOptionsBrowserActionScrollActionScrollUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionScrollActionScrollUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionScrollActionScrollObject) {
		return u.OfCrawlsExtractOptionsBrowserActionScrollActionScrollObject
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionScrollActionScrollObject struct {
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	X       param.Opt[float64] `json:"x,omitzero"`
	Y       param.Opt[float64] `json:"y,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Container CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectContainerUnion `json:"container,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectSkipUnion `json:"skip,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	To CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectToUnion `json:"to,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionScrollActionScrollObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionScrollActionScrollObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionScrollActionScrollObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectContainerUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectContainerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectContainerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectContainerUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionScrollActionScrollObjectRequiredString)
	OfCrawlsExtractOptionsBrowserActionScrollActionScrollObjectRequiredString param.Opt[CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                    param.Opt[bool]                                                                         `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionScrollActionScrollObjectRequiredString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionScrollActionScrollObjectRequiredString) {
		return &u.OfCrawlsExtractOptionsBrowserActionScrollActionScrollObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectRequiredString string

const (
	CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectRequiredStringTrue  CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectRequiredString = "true"
	CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectRequiredStringFalse CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionScrollActionScrollObjectSkipString)
	OfCrawlsExtractOptionsBrowserActionScrollActionScrollObjectSkipString param.Opt[CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                param.Opt[bool]                                                                     `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionScrollActionScrollObjectSkipString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionScrollActionScrollObjectSkipString) {
		return &u.OfCrawlsExtractOptionsBrowserActionScrollActionScrollObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectSkipString string

const (
	CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectSkipStringTrue  CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectSkipString = "true"
	CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectSkipStringFalse CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectSkipString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectToUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectToUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectToUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionScrollActionScrollObjectToUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionWaitAction struct {
	Wait CrawlParamsExtractOptionsBrowserActionWaitActionWaitUnion `json:"wait,omitzero,required"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionWaitAction) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionWaitAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionWaitAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionWaitActionWaitUnion struct {
	OfFloat                                                 param.Opt[float64]                                          `json:",omitzero,inline"`
	OfString                                                param.Opt[string]                                           `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionWaitActionWaitObject *CrawlParamsExtractOptionsBrowserActionWaitActionWaitObject `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionWaitActionWaitUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfCrawlsExtractOptionsBrowserActionWaitActionWaitObject)
}
func (u *CrawlParamsExtractOptionsBrowserActionWaitActionWaitUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionWaitActionWaitUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionWaitActionWaitObject) {
		return u.OfCrawlsExtractOptionsBrowserActionWaitActionWaitObject
	}
	return nil
}

// The property Duration is required.
type CrawlParamsExtractOptionsBrowserActionWaitActionWaitObject struct {
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Duration CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectDurationUnion `json:"duration,omitzero,required"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionWaitActionWaitObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionWaitActionWaitObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionWaitActionWaitObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectDurationUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectDurationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectDurationUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionWaitActionWaitObjectRequiredString)
	OfCrawlsExtractOptionsBrowserActionWaitActionWaitObjectRequiredString param.Opt[CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                param.Opt[bool]                                                                     `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionWaitActionWaitObjectRequiredString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionWaitActionWaitObjectRequiredString) {
		return &u.OfCrawlsExtractOptionsBrowserActionWaitActionWaitObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectRequiredString string

const (
	CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectRequiredStringTrue  CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectRequiredString = "true"
	CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectRequiredStringFalse CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionWaitActionWaitObjectSkipString)
	OfCrawlsExtractOptionsBrowserActionWaitActionWaitObjectSkipString param.Opt[CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                            param.Opt[bool]                                                                 `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionWaitActionWaitObjectSkipString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionWaitActionWaitObjectSkipString) {
		return &u.OfCrawlsExtractOptionsBrowserActionWaitActionWaitObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectSkipString string

const (
	CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectSkipStringTrue  CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectSkipString = "true"
	CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectSkipStringFalse CrawlParamsExtractOptionsBrowserActionWaitActionWaitObjectSkipString = "false"
)

// Wait for an element to appear or reach a specific state
//
// The property WaitForElement is required.
type CrawlParamsExtractOptionsBrowserActionWaitForElementAction struct {
	WaitForElement CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementUnion `json:"wait_for_element,omitzero,required"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionWaitForElementAction) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionWaitForElementAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionWaitForElementAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementUnion struct {
	OfString                                                                    param.Opt[string]                                                               `json:",omitzero,inline"`
	OfStringArray                                                               []string                                                                        `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionWaitForElementActionWaitForElementObject *CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObject `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray, u.OfCrawlsExtractOptionsBrowserActionWaitForElementActionWaitForElementObject)
}
func (u *CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionWaitForElementActionWaitForElementObject) {
		return u.OfCrawlsExtractOptionsBrowserActionWaitForElementActionWaitForElementObject
	}
	return nil
}

// The property Selector is required.
type CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObject struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion `json:"selector,omitzero,required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion) asAny() any {
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
type CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectRequiredString)
	OfCrawlsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectRequiredString param.Opt[CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                                    param.Opt[bool]                                                                                         `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectRequiredString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectRequiredString) {
		return &u.OfCrawlsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectRequiredString string

const (
	CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectRequiredStringTrue  CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectRequiredString = "true"
	CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectRequiredStringFalse CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSkipString)
	OfCrawlsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSkipString param.Opt[CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                                param.Opt[bool]                                                                                     `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSkipString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSkipString) {
		return &u.OfCrawlsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSkipString string

const (
	CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSkipStringTrue  CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSkipString = "true"
	CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSkipStringFalse CrawlParamsExtractOptionsBrowserActionWaitForElementActionWaitForElementObjectSkipString = "false"
)

// Wait for page navigation to complete
//
// The property WaitForNavigation is required.
type CrawlParamsExtractOptionsBrowserActionWaitForNavigationAction struct {
	WaitForNavigation CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationUnion `json:"wait_for_navigation,omitzero,required"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionWaitForNavigationAction) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionWaitForNavigationAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionWaitForNavigationAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationString)
	OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationString param.Opt[string]                                                                     `json:",omitzero,inline"`
	OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObject *CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObject `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationString, u.OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObject)
}
func (u *CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationString) {
		return &u.OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationString
	} else if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObject) {
		return u.OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObject
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationString string

const (
	CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationStringLoad             CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationString = "load"
	CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationStringDomcontentloaded CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationString = "domcontentloaded"
	CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationStringNetworkidle0     CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationString = "networkidle0"
	CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationStringNetworkidle2     CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationString = "networkidle2"
)

// The property Navigation is required.
type CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObject struct {
	// Any of "load", "domcontentloaded", "networkidle0", "networkidle2".
	Navigation string `json:"navigation,omitzero,required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObject](
		"navigation", "load", "domcontentloaded", "networkidle0", "networkidle2",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString)
	OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString param.Opt[CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                                          param.Opt[bool]                                                                                               `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString) {
		return &u.OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString string

const (
	CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredStringTrue  CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString = "true"
	CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredStringFalse CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString)
	OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString param.Opt[CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                                      param.Opt[bool]                                                                                           `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString, u.OfBool)
}
func (u *CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString) {
		return &u.OfCrawlsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString string

const (
	CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipStringTrue  CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString = "true"
	CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipStringFalse CrawlParamsExtractOptionsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString = "false"
)

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
