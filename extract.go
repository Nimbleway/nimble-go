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

// ExtractService contains methods and other services that help with interacting
// with the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtractService] method instead.
type ExtractService struct {
	Options []option.RequestOption
}

// NewExtractService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewExtractService(opts ...option.RequestOption) (r ExtractService) {
	r = ExtractService{}
	r.Options = opts
	return
}

// Extract Async Endpoint
func (r *ExtractService) Async(ctx context.Context, body ExtractAsyncParams, opts ...option.RequestOption) (res *ExtractAsyncResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/extract/async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Extract
func (r *ExtractService) Run(ctx context.Context, body ExtractRunParams, opts ...option.RequestOption) (res *ExtractRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/extract"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Response when an async extract task is created successfully.
type ExtractAsyncResponse struct {
	// Status indicating the async task was created successfully.
	Status constant.Success `json:"status,required"`
	// The created async task details.
	Task ExtractAsyncResponseTask `json:"task,required"`
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
	ID    string `json:"id,required"`
	Query any    `json:"_query,required"`
	// Timestamp when the task was created.
	CreatedAt string `json:"created_at,required"`
	// Original input data for the task.
	Input any `json:"input,required"`
	// Current state of the task.
	//
	// Any of "pending", "success", "error".
	State string `json:"state,required"`
	// URL for checking the task status.
	StatusURL string `json:"status_url,required" format:"uri"`
	// Account name that owns the task.
	AccountName string `json:"account_name"`
	// Any of "web", "serp", "ecommerce", "social".
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

type ExtractRunResponse struct {
	Data     ExtractRunResponseData     `json:"data,required"`
	Metadata ExtractRunResponseMetadata `json:"metadata,required"`
	// The status of the task.
	//
	// Any of "success", "skipped", "fatal", "error", "postponed", "ignored",
	// "rejected", "blocked".
	Status ExtractRunResponseStatus `json:"status,required"`
	// Unique identifier for the task.
	TaskID string `json:"task_id,required"`
	// The final URL.
	URL   string                  `json:"url,required"`
	Debug ExtractRunResponseDebug `json:"debug"`
	// Pagination information if applicable.
	Pagination ExtractRunResponsePaginationUnion `json:"pagination"`
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
func (r ExtractRunResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseData struct {
	// Browser actions execution results. Present only when browser_actions were
	// specified in the request.
	BrowserActions ExtractRunResponseDataBrowserActions `json:"browser_actions"`
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
	NetworkCapture []ExtractRunResponseDataNetworkCapture `json:"network_capture"`
	// The parsing results extracted from the HTML & network content.
	Parsing ExtractRunResponseDataParsingUnion `json:"parsing"`
	// The list of redirects that occurred during the task.
	Redirects []ExtractRunResponseDataRedirect `json:"redirects"`
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
func (r ExtractRunResponseData) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser actions execution results. Present only when browser_actions were
// specified in the request.
type ExtractRunResponseDataBrowserActions struct {
	Results       []ExtractRunResponseDataBrowserActionsResult `json:"results,required"`
	Success       bool                                         `json:"success,required"`
	TotalDuration float64                                      `json:"total_duration,required"`
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
func (r ExtractRunResponseDataBrowserActions) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDataBrowserActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseDataBrowserActionsResult struct {
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
func (r ExtractRunResponseDataBrowserActionsResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDataBrowserActionsResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseDataNetworkCapture struct {
	Filter       ExtractRunResponseDataNetworkCaptureFilter   `json:"filter,required"`
	Results      []ExtractRunResponseDataNetworkCaptureResult `json:"results,required"`
	ErrorMessage string                                       `json:"errorMessage"`
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
func (r ExtractRunResponseDataNetworkCapture) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDataNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseDataNetworkCaptureFilter struct {
	Validation           bool    `json:"validation,required"`
	WaitForRequestsCount float64 `json:"wait_for_requests_count,required"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method"`
	// Resource type for network capture filtering
	ResourceType                ExtractRunResponseDataNetworkCaptureFilterResourceTypeUnion `json:"resource_type"`
	StatusCode                  ExtractRunResponseDataNetworkCaptureFilterStatusCodeUnion   `json:"status_code"`
	URL                         ExtractRunResponseDataNetworkCaptureFilterURL               `json:"url"`
	WaitForRequestsCountTimeout float64                                                     `json:"wait_for_requests_count_timeout"`
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
func (r ExtractRunResponseDataNetworkCaptureFilter) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDataNetworkCaptureFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractRunResponseDataNetworkCaptureFilterResourceTypeUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractRunResponseDataNetworkCaptureFilterResourceTypeString
// OfExtractRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray]
type ExtractRunResponseDataNetworkCaptureFilterResourceTypeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfExtractRunResponseDataNetworkCaptureFilterResourceTypeString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfExtractRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray []string `json:",inline"`
	JSON                                                                   struct {
		OfExtractRunResponseDataNetworkCaptureFilterResourceTypeString         respjson.Field
		OfExtractRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray respjson.Field
		raw                                                                    string
	} `json:"-"`
}

func (u ExtractRunResponseDataNetworkCaptureFilterResourceTypeUnion) AsExtractRunResponseDataNetworkCaptureFilterResourceTypeString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractRunResponseDataNetworkCaptureFilterResourceTypeUnion) AsExtractRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractRunResponseDataNetworkCaptureFilterResourceTypeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ExtractRunResponseDataNetworkCaptureFilterResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type for network capture filtering
type ExtractRunResponseDataNetworkCaptureFilterResourceTypeString string

const (
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringDocument           ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "document"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringStylesheet         ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "stylesheet"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringImage              ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "image"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringMedia              ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "media"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringFont               ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "font"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringScript             ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "script"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringTexttrack          ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "texttrack"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringXhr                ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "xhr"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringFetch              ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "fetch"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringPrefetch           ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "prefetch"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringEventsource        ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "eventsource"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringWebsocket          ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "websocket"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringManifest           ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "manifest"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringSignedexchange     ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "signedexchange"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringPing               ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "ping"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringCspviolationreport ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "cspviolationreport"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringPreflight          ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "preflight"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringOther              ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "other"
	ExtractRunResponseDataNetworkCaptureFilterResourceTypeStringFedcm              ExtractRunResponseDataNetworkCaptureFilterResourceTypeString = "fedcm"
)

// ExtractRunResponseDataNetworkCaptureFilterStatusCodeUnion contains all possible
// properties and values from [float64], [[]float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfFloatArray]
type ExtractRunResponseDataNetworkCaptureFilterStatusCodeUnion struct {
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

func (u ExtractRunResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractRunResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloatArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractRunResponseDataNetworkCaptureFilterStatusCodeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ExtractRunResponseDataNetworkCaptureFilterStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseDataNetworkCaptureFilterURL struct {
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
func (r ExtractRunResponseDataNetworkCaptureFilterURL) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDataNetworkCaptureFilterURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseDataNetworkCaptureResult struct {
	Request  ExtractRunResponseDataNetworkCaptureResultRequest  `json:"request,required"`
	Response ExtractRunResponseDataNetworkCaptureResultResponse `json:"response,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Request     respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractRunResponseDataNetworkCaptureResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDataNetworkCaptureResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseDataNetworkCaptureResultRequest struct {
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
func (r ExtractRunResponseDataNetworkCaptureResultRequest) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDataNetworkCaptureResultRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseDataNetworkCaptureResultResponse struct {
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
func (r ExtractRunResponseDataNetworkCaptureResultResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDataNetworkCaptureResultResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractRunResponseDataParsingUnion contains all possible properties and values
// from [ExtractRunResponseDataParsingParsingSuccessResult],
// [ExtractRunResponseDataParsingParsingErrorResult], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractRunResponseDataParsingMapItem]
type ExtractRunResponseDataParsingUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExtractRunResponseDataParsingMapItem any `json:",inline"`
	// This field is from variant [ExtractRunResponseDataParsingParsingSuccessResult].
	Entities map[string]any `json:"entities"`
	Status   string         `json:"status"`
	// This field is from variant [ExtractRunResponseDataParsingParsingErrorResult].
	Error string `json:"error"`
	JSON  struct {
		OfExtractRunResponseDataParsingMapItem respjson.Field
		Entities                               respjson.Field
		Status                                 respjson.Field
		Error                                  respjson.Field
		raw                                    string
	} `json:"-"`
}

func (u ExtractRunResponseDataParsingUnion) AsExtractRunResponseDataParsingParsingSuccessResult() (v ExtractRunResponseDataParsingParsingSuccessResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractRunResponseDataParsingUnion) AsExtractRunResponseDataParsingParsingErrorResult() (v ExtractRunResponseDataParsingParsingErrorResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractRunResponseDataParsingUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractRunResponseDataParsingUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractRunResponseDataParsingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseDataParsingParsingSuccessResult struct {
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
func (r ExtractRunResponseDataParsingParsingSuccessResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDataParsingParsingSuccessResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseDataParsingParsingErrorResult struct {
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
func (r ExtractRunResponseDataParsingParsingErrorResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDataParsingParsingErrorResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseDataRedirect struct {
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
func (r ExtractRunResponseDataRedirect) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDataRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseMetadata struct {
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
func (r ExtractRunResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the task.
type ExtractRunResponseStatus string

const (
	ExtractRunResponseStatusSuccess   ExtractRunResponseStatus = "success"
	ExtractRunResponseStatusSkipped   ExtractRunResponseStatus = "skipped"
	ExtractRunResponseStatusFatal     ExtractRunResponseStatus = "fatal"
	ExtractRunResponseStatusError     ExtractRunResponseStatus = "error"
	ExtractRunResponseStatusPostponed ExtractRunResponseStatus = "postponed"
	ExtractRunResponseStatusIgnored   ExtractRunResponseStatus = "ignored"
	ExtractRunResponseStatusRejected  ExtractRunResponseStatus = "rejected"
	ExtractRunResponseStatusBlocked   ExtractRunResponseStatus = "blocked"
)

type ExtractRunResponseDebug struct {
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
func (r ExtractRunResponseDebug) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDebug) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractRunResponsePaginationUnion contains all possible properties and values
// from [ExtractRunResponsePaginationNextPageParams],
// [[]ExtractRunResponsePaginationArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractRunResponsePaginationArray]
type ExtractRunResponsePaginationUnion struct {
	// This field will be present if the value is a
	// [[]ExtractRunResponsePaginationArrayItem] instead of an object.
	OfExtractRunResponsePaginationArray []ExtractRunResponsePaginationArrayItem `json:",inline"`
	// This field is from variant [ExtractRunResponsePaginationNextPageParams].
	NextPageParams map[string]any `json:"next_page_params"`
	JSON           struct {
		OfExtractRunResponsePaginationArray respjson.Field
		NextPageParams                      respjson.Field
		raw                                 string
	} `json:"-"`
}

func (u ExtractRunResponsePaginationUnion) AsExtractRunResponsePaginationNextPageParams() (v ExtractRunResponsePaginationNextPageParams) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractRunResponsePaginationUnion) AsExtractRunResponsePaginationArray() (v []ExtractRunResponsePaginationArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractRunResponsePaginationUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractRunResponsePaginationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponsePaginationNextPageParams struct {
	NextPageParams map[string]any `json:"next_page_params,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractRunResponsePaginationNextPageParams) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponsePaginationNextPageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponsePaginationArrayItem struct {
	NextPageParams map[string]any `json:"next_page_params,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractRunResponsePaginationArrayItem) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponsePaginationArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractAsyncParams struct {
	// Target URL to scrape
	URL string `json:"url,required"`
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
	// Any of "html", "markdown".
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
	Name string `json:"name,omitzero,required"`
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
	OfExtractAsyncsBrowserActionAutoScrollAction        *ExtractAsyncParamsBrowserActionAutoScrollAction        `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionClickAction             *ExtractAsyncParamsBrowserActionClickAction             `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionEvalAction              *ExtractAsyncParamsBrowserActionEvalAction              `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionFetchAction             *ExtractAsyncParamsBrowserActionFetchAction             `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionFillAction              *ExtractAsyncParamsBrowserActionFillAction              `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionGetCookiesAction        *ExtractAsyncParamsBrowserActionGetCookiesAction        `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionGotoAction              *ExtractAsyncParamsBrowserActionGotoAction              `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionPressAction             *ExtractAsyncParamsBrowserActionPressAction             `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionScreenshotAction        *ExtractAsyncParamsBrowserActionScreenshotAction        `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionScrollAction            *ExtractAsyncParamsBrowserActionScrollAction            `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionWaitAction              *ExtractAsyncParamsBrowserActionWaitAction              `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionWaitForElementAction    *ExtractAsyncParamsBrowserActionWaitForElementAction    `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionWaitForNavigationAction *ExtractAsyncParamsBrowserActionWaitForNavigationAction `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionAutoScrollAction,
		u.OfExtractAsyncsBrowserActionClickAction,
		u.OfExtractAsyncsBrowserActionEvalAction,
		u.OfExtractAsyncsBrowserActionFetchAction,
		u.OfExtractAsyncsBrowserActionFillAction,
		u.OfExtractAsyncsBrowserActionGetCookiesAction,
		u.OfExtractAsyncsBrowserActionGotoAction,
		u.OfExtractAsyncsBrowserActionPressAction,
		u.OfExtractAsyncsBrowserActionScreenshotAction,
		u.OfExtractAsyncsBrowserActionScrollAction,
		u.OfExtractAsyncsBrowserActionWaitAction,
		u.OfExtractAsyncsBrowserActionWaitForElementAction,
		u.OfExtractAsyncsBrowserActionWaitForNavigationAction)
}
func (u *ExtractAsyncParamsBrowserActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionAutoScrollAction) {
		return u.OfExtractAsyncsBrowserActionAutoScrollAction
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionClickAction) {
		return u.OfExtractAsyncsBrowserActionClickAction
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionEvalAction) {
		return u.OfExtractAsyncsBrowserActionEvalAction
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionFetchAction) {
		return u.OfExtractAsyncsBrowserActionFetchAction
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionFillAction) {
		return u.OfExtractAsyncsBrowserActionFillAction
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionGetCookiesAction) {
		return u.OfExtractAsyncsBrowserActionGetCookiesAction
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionGotoAction) {
		return u.OfExtractAsyncsBrowserActionGotoAction
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionPressAction) {
		return u.OfExtractAsyncsBrowserActionPressAction
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionScreenshotAction) {
		return u.OfExtractAsyncsBrowserActionScreenshotAction
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionScrollAction) {
		return u.OfExtractAsyncsBrowserActionScrollAction
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionWaitAction) {
		return u.OfExtractAsyncsBrowserActionWaitAction
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionWaitForElementAction) {
		return u.OfExtractAsyncsBrowserActionWaitForElementAction
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionWaitForNavigationAction) {
		return u.OfExtractAsyncsBrowserActionWaitForNavigationAction
	}
	return nil
}

// Continuously scroll to load dynamic content
//
// The property AutoScroll is required.
type ExtractAsyncParamsBrowserActionAutoScrollAction struct {
	AutoScroll ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollUnion `json:"auto_scroll,omitzero,required"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionAutoScrollAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionAutoScrollAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionAutoScrollAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollUnion struct {
	OfBool                                                       param.Opt[bool]                                                  `json:",omitzero,inline"`
	OfFloat                                                      param.Opt[float64]                                               `json:",omitzero,inline"`
	OfString                                                     param.Opt[string]                                                `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionAutoScrollActionAutoScrollObject *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfExtractAsyncsBrowserActionAutoScrollActionAutoScrollObject)
}
func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionAutoScrollActionAutoScrollObject) {
		return u.OfExtractAsyncsBrowserActionAutoScrollActionAutoScrollObject
	}
	return nil
}

type ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObject struct {
	StepSize param.Opt[float64] `json:"step_size,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	ClickSelector ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion `json:"click_selector,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Container ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion `json:"container,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	DelayAfterScroll ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion `json:"delay_after_scroll,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	IdleTimeout ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion `json:"idle_timeout,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	LoadingSelector ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion `json:"loading_selector,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	MaxDuration ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion `json:"max_duration,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	PauseOnSelector ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion `json:"pause_on_selector,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionAutoScrollActionAutoScrollObjectRequiredString)
	OfExtractAsyncsBrowserActionAutoScrollActionAutoScrollObjectRequiredString param.Opt[ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                     param.Opt[bool]                                                                          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionAutoScrollActionAutoScrollObjectRequiredString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionAutoScrollActionAutoScrollObjectRequiredString) {
		return &u.OfExtractAsyncsBrowserActionAutoScrollActionAutoScrollObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredString string

const (
	ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredStringTrue  ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredString = "true"
	ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredStringFalse ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionAutoScrollActionAutoScrollObjectSkipString)
	OfExtractAsyncsBrowserActionAutoScrollActionAutoScrollObjectSkipString param.Opt[ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                 param.Opt[bool]                                                                      `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionAutoScrollActionAutoScrollObjectSkipString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionAutoScrollActionAutoScrollObjectSkipString) {
		return &u.OfExtractAsyncsBrowserActionAutoScrollActionAutoScrollObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectSkipString string

const (
	ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectSkipStringTrue  ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectSkipString = "true"
	ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectSkipStringFalse ExtractAsyncParamsBrowserActionAutoScrollActionAutoScrollObjectSkipString = "false"
)

// Click on an element by selector
//
// The property Click is required.
type ExtractAsyncParamsBrowserActionClickAction struct {
	Click ExtractAsyncParamsBrowserActionClickActionClickUnion `json:"click,omitzero,required"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionClickAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionClickAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionClickAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionClickActionClickUnion struct {
	OfString                                           param.Opt[string]                                      `json:",omitzero,inline"`
	OfStringArray                                      []string                                               `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionClickActionClickObject *ExtractAsyncParamsBrowserActionClickActionClickObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionClickActionClickUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray, u.OfExtractAsyncsBrowserActionClickActionClickObject)
}
func (u *ExtractAsyncParamsBrowserActionClickActionClickUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionClickActionClickUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionClickActionClickObject) {
		return u.OfExtractAsyncsBrowserActionClickActionClickObject
	}
	return nil
}

// The property Selector is required.
type ExtractAsyncParamsBrowserActionClickActionClickObject struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector ExtractAsyncParamsBrowserActionClickActionClickObjectSelectorUnion `json:"selector,omitzero,required"`
	Count    param.Opt[float64]                                                 `json:"count,omitzero"`
	OffsetX  param.Opt[int64]                                                   `json:"offset_x,omitzero"`
	OffsetY  param.Opt[int64]                                                   `json:"offset_y,omitzero"`
	Scroll   param.Opt[bool]                                                    `json:"scroll,omitzero"`
	Steps    param.Opt[float64]                                                 `json:"steps,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay ExtractAsyncParamsBrowserActionClickActionClickObjectDelayUnion `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractAsyncParamsBrowserActionClickActionClickObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractAsyncParamsBrowserActionClickActionClickObjectSkipUnion `json:"skip,omitzero"`
	// Any of "linear", "ghost-cursor", "windmouse".
	Strategy string `json:"strategy,omitzero"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionClickActionClickObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionClickActionClickObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionClickActionClickObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractAsyncParamsBrowserActionClickActionClickObject](
		"strategy", "linear", "ghost-cursor", "windmouse",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionClickActionClickObjectSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionClickActionClickObjectSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractAsyncParamsBrowserActionClickActionClickObjectSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionClickActionClickObjectSelectorUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionClickActionClickObjectDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionClickActionClickObjectDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractAsyncParamsBrowserActionClickActionClickObjectDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionClickActionClickObjectDelayUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionClickActionClickObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionClickActionClickObjectRequiredString)
	OfExtractAsyncsBrowserActionClickActionClickObjectRequiredString param.Opt[ExtractAsyncParamsBrowserActionClickActionClickObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                           param.Opt[bool]                                                                `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionClickActionClickObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionClickActionClickObjectRequiredString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionClickActionClickObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionClickActionClickObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionClickActionClickObjectRequiredString) {
		return &u.OfExtractAsyncsBrowserActionClickActionClickObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionClickActionClickObjectRequiredString string

const (
	ExtractAsyncParamsBrowserActionClickActionClickObjectRequiredStringTrue  ExtractAsyncParamsBrowserActionClickActionClickObjectRequiredString = "true"
	ExtractAsyncParamsBrowserActionClickActionClickObjectRequiredStringFalse ExtractAsyncParamsBrowserActionClickActionClickObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionClickActionClickObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionClickActionClickObjectSkipString)
	OfExtractAsyncsBrowserActionClickActionClickObjectSkipString param.Opt[ExtractAsyncParamsBrowserActionClickActionClickObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                       param.Opt[bool]                                                            `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionClickActionClickObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionClickActionClickObjectSkipString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionClickActionClickObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionClickActionClickObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionClickActionClickObjectSkipString) {
		return &u.OfExtractAsyncsBrowserActionClickActionClickObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionClickActionClickObjectSkipString string

const (
	ExtractAsyncParamsBrowserActionClickActionClickObjectSkipStringTrue  ExtractAsyncParamsBrowserActionClickActionClickObjectSkipString = "true"
	ExtractAsyncParamsBrowserActionClickActionClickObjectSkipStringFalse ExtractAsyncParamsBrowserActionClickActionClickObjectSkipString = "false"
)

// Execute JavaScript code in page context
//
// The property Eval is required.
type ExtractAsyncParamsBrowserActionEvalAction struct {
	Eval ExtractAsyncParamsBrowserActionEvalActionEvalUnion `json:"eval,omitzero,required"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionEvalAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionEvalAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionEvalAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionEvalActionEvalUnion struct {
	OfString                                         param.Opt[string]                                    `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionEvalActionEvalObject *ExtractAsyncParamsBrowserActionEvalActionEvalObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionEvalActionEvalUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractAsyncsBrowserActionEvalActionEvalObject)
}
func (u *ExtractAsyncParamsBrowserActionEvalActionEvalUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionEvalActionEvalUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionEvalActionEvalObject) {
		return u.OfExtractAsyncsBrowserActionEvalActionEvalObject
	}
	return nil
}

// The property Code is required.
type ExtractAsyncParamsBrowserActionEvalActionEvalObject struct {
	Code string `json:"code,required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractAsyncParamsBrowserActionEvalActionEvalObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractAsyncParamsBrowserActionEvalActionEvalObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionEvalActionEvalObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionEvalActionEvalObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionEvalActionEvalObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionEvalActionEvalObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionEvalActionEvalObjectRequiredString)
	OfExtractAsyncsBrowserActionEvalActionEvalObjectRequiredString param.Opt[ExtractAsyncParamsBrowserActionEvalActionEvalObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                         param.Opt[bool]                                                              `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionEvalActionEvalObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionEvalActionEvalObjectRequiredString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionEvalActionEvalObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionEvalActionEvalObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionEvalActionEvalObjectRequiredString) {
		return &u.OfExtractAsyncsBrowserActionEvalActionEvalObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionEvalActionEvalObjectRequiredString string

const (
	ExtractAsyncParamsBrowserActionEvalActionEvalObjectRequiredStringTrue  ExtractAsyncParamsBrowserActionEvalActionEvalObjectRequiredString = "true"
	ExtractAsyncParamsBrowserActionEvalActionEvalObjectRequiredStringFalse ExtractAsyncParamsBrowserActionEvalActionEvalObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionEvalActionEvalObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionEvalActionEvalObjectSkipString)
	OfExtractAsyncsBrowserActionEvalActionEvalObjectSkipString param.Opt[ExtractAsyncParamsBrowserActionEvalActionEvalObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                     param.Opt[bool]                                                          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionEvalActionEvalObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionEvalActionEvalObjectSkipString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionEvalActionEvalObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionEvalActionEvalObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionEvalActionEvalObjectSkipString) {
		return &u.OfExtractAsyncsBrowserActionEvalActionEvalObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionEvalActionEvalObjectSkipString string

const (
	ExtractAsyncParamsBrowserActionEvalActionEvalObjectSkipStringTrue  ExtractAsyncParamsBrowserActionEvalActionEvalObjectSkipString = "true"
	ExtractAsyncParamsBrowserActionEvalActionEvalObjectSkipStringFalse ExtractAsyncParamsBrowserActionEvalActionEvalObjectSkipString = "false"
)

// Make an HTTP request in browser context
//
// The property Fetch is required.
type ExtractAsyncParamsBrowserActionFetchAction struct {
	Fetch ExtractAsyncParamsBrowserActionFetchActionFetchUnion `json:"fetch,omitzero,required" format:"uri"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionFetchAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionFetchAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionFetchAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionFetchActionFetchUnion struct {
	OfString                                           param.Opt[string]                                      `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionFetchActionFetchObject *ExtractAsyncParamsBrowserActionFetchActionFetchObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionFetchActionFetchUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractAsyncsBrowserActionFetchActionFetchObject)
}
func (u *ExtractAsyncParamsBrowserActionFetchActionFetchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionFetchActionFetchUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionFetchActionFetchObject) {
		return u.OfExtractAsyncsBrowserActionFetchActionFetchObject
	}
	return nil
}

// The property URL is required.
type ExtractAsyncParamsBrowserActionFetchActionFetchObject struct {
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
	Required ExtractAsyncParamsBrowserActionFetchActionFetchObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractAsyncParamsBrowserActionFetchActionFetchObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionFetchActionFetchObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionFetchActionFetchObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionFetchActionFetchObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractAsyncParamsBrowserActionFetchActionFetchObject](
		"method", "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionFetchActionFetchObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionFetchActionFetchObjectRequiredString)
	OfExtractAsyncsBrowserActionFetchActionFetchObjectRequiredString param.Opt[ExtractAsyncParamsBrowserActionFetchActionFetchObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                           param.Opt[bool]                                                                `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionFetchActionFetchObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionFetchActionFetchObjectRequiredString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionFetchActionFetchObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionFetchActionFetchObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionFetchActionFetchObjectRequiredString) {
		return &u.OfExtractAsyncsBrowserActionFetchActionFetchObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionFetchActionFetchObjectRequiredString string

const (
	ExtractAsyncParamsBrowserActionFetchActionFetchObjectRequiredStringTrue  ExtractAsyncParamsBrowserActionFetchActionFetchObjectRequiredString = "true"
	ExtractAsyncParamsBrowserActionFetchActionFetchObjectRequiredStringFalse ExtractAsyncParamsBrowserActionFetchActionFetchObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionFetchActionFetchObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionFetchActionFetchObjectSkipString)
	OfExtractAsyncsBrowserActionFetchActionFetchObjectSkipString param.Opt[ExtractAsyncParamsBrowserActionFetchActionFetchObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                       param.Opt[bool]                                                            `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionFetchActionFetchObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionFetchActionFetchObjectSkipString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionFetchActionFetchObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionFetchActionFetchObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionFetchActionFetchObjectSkipString) {
		return &u.OfExtractAsyncsBrowserActionFetchActionFetchObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionFetchActionFetchObjectSkipString string

const (
	ExtractAsyncParamsBrowserActionFetchActionFetchObjectSkipStringTrue  ExtractAsyncParamsBrowserActionFetchActionFetchObjectSkipString = "true"
	ExtractAsyncParamsBrowserActionFetchActionFetchObjectSkipStringFalse ExtractAsyncParamsBrowserActionFetchActionFetchObjectSkipString = "false"
)

// Fill text into an input field
//
// The property Fill is required.
type ExtractAsyncParamsBrowserActionFillAction struct {
	// Fill options with mode-specific fields. Use "type" mode for behavioral typing
	// simulation, or "paste" mode for instant paste.
	Fill ExtractAsyncParamsBrowserActionFillActionFillUnion `json:"fill,omitzero,required"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionFillAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionFillAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionFillAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionFillActionFillUnion struct {
	OfType  *ExtractAsyncParamsBrowserActionFillActionFillType  `json:",omitzero,inline"`
	OfPaste *ExtractAsyncParamsBrowserActionFillActionFillPaste `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionFillActionFillUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfType, u.OfPaste)
}
func (u *ExtractAsyncParamsBrowserActionFillActionFillUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionFillActionFillUnion) asAny() any {
	if !param.IsOmitted(u.OfType) {
		return u.OfType
	} else if !param.IsOmitted(u.OfPaste) {
		return u.OfPaste
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractAsyncParamsBrowserActionFillActionFillUnion) GetMouseMovementStrategy() *string {
	if vt := u.OfType; vt != nil {
		return &vt.MouseMovementStrategy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractAsyncParamsBrowserActionFillActionFillUnion) GetTypingInterval() *ExtractAsyncParamsBrowserActionFillActionFillTypeTypingIntervalUnion {
	if vt := u.OfType; vt != nil {
		return &vt.TypingInterval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractAsyncParamsBrowserActionFillActionFillUnion) GetTypingStrategy() *string {
	if vt := u.OfType; vt != nil {
		return &vt.TypingStrategy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractAsyncParamsBrowserActionFillActionFillUnion) GetValue() *string {
	if vt := u.OfType; vt != nil {
		return (*string)(&vt.Value)
	} else if vt := u.OfPaste; vt != nil {
		return (*string)(&vt.Value)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractAsyncParamsBrowserActionFillActionFillUnion) GetClickOnElement() *bool {
	if vt := u.OfType; vt != nil && vt.ClickOnElement.Valid() {
		return &vt.ClickOnElement.Value
	} else if vt := u.OfPaste; vt != nil && vt.ClickOnElement.Valid() {
		return &vt.ClickOnElement.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractAsyncParamsBrowserActionFillActionFillUnion) GetMode() *string {
	if vt := u.OfType; vt != nil {
		return (*string)(&vt.Mode)
	} else if vt := u.OfPaste; vt != nil {
		return (*string)(&vt.Mode)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractAsyncParamsBrowserActionFillActionFillUnion) GetScroll() *bool {
	if vt := u.OfType; vt != nil && vt.Scroll.Valid() {
		return &vt.Scroll.Value
	} else if vt := u.OfPaste; vt != nil && vt.Scroll.Valid() {
		return &vt.Scroll.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractAsyncParamsBrowserActionFillActionFillUnion) GetTimeout() *float64 {
	if vt := u.OfType; vt != nil && vt.Timeout.Valid() {
		return &vt.Timeout.Value
	} else if vt := u.OfPaste; vt != nil && vt.Timeout.Valid() {
		return &vt.Timeout.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractAsyncParamsBrowserActionFillActionFillUnion) GetVisible() *bool {
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
func (u ExtractAsyncParamsBrowserActionFillActionFillUnion) GetSelector() (res extractAsyncParamsBrowserActionFillActionFillUnionSelector) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Selector.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Selector.asAny()
	}
	return
}

// Can have the runtime types [*string], [\*[]string]
type extractAsyncParamsBrowserActionFillActionFillUnionSelector struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extractAsyncParamsBrowserActionFillActionFillUnionSelector) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ExtractAsyncParamsBrowserActionFillActionFillUnion) GetDelay() (res extractAsyncParamsBrowserActionFillActionFillUnionDelay) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Delay.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Delay.asAny()
	}
	return
}

// Can have the runtime types [*float64], [*string]
type extractAsyncParamsBrowserActionFillActionFillUnionDelay struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *float64:
//	case *string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extractAsyncParamsBrowserActionFillActionFillUnionDelay) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ExtractAsyncParamsBrowserActionFillActionFillUnion) GetRequired() (res extractAsyncParamsBrowserActionFillActionFillUnionRequired) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Required.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Required.asAny()
	}
	return
}

// Can have the runtime types [*string], [*bool]
type extractAsyncParamsBrowserActionFillActionFillUnionRequired struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extractAsyncParamsBrowserActionFillActionFillUnionRequired) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ExtractAsyncParamsBrowserActionFillActionFillUnion) GetSkip() (res extractAsyncParamsBrowserActionFillActionFillUnionSkip) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Skip.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Skip.asAny()
	}
	return
}

// Can have the runtime types [*string], [*bool]
type extractAsyncParamsBrowserActionFillActionFillUnionSkip struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extractAsyncParamsBrowserActionFillActionFillUnionSkip) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[ExtractAsyncParamsBrowserActionFillActionFillUnion](
		"mode",
		apijson.Discriminator[ExtractAsyncParamsBrowserActionFillActionFillType]("type"),
		apijson.Discriminator[ExtractAsyncParamsBrowserActionFillActionFillPaste]("paste"),
	)
}

// The properties Selector, Value are required.
type ExtractAsyncParamsBrowserActionFillActionFillType struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector       ExtractAsyncParamsBrowserActionFillActionFillTypeSelectorUnion `json:"selector,omitzero,required"`
	Value          string                                                         `json:"value,required"`
	ClickOnElement param.Opt[bool]                                                `json:"click_on_element,omitzero"`
	Scroll         param.Opt[bool]                                                `json:"scroll,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay ExtractAsyncParamsBrowserActionFillActionFillTypeDelayUnion `json:"delay,omitzero"`
	// Any of "type".
	Mode string `json:"mode,omitzero"`
	// Any of "linear", "ghost-cursor", "windmouse".
	MouseMovementStrategy string `json:"mouse_movement_strategy,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractAsyncParamsBrowserActionFillActionFillTypeRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractAsyncParamsBrowserActionFillActionFillTypeSkipUnion `json:"skip,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	TypingInterval ExtractAsyncParamsBrowserActionFillActionFillTypeTypingIntervalUnion `json:"typing_interval,omitzero"`
	// Any of "simple", "distribution".
	TypingStrategy string `json:"typing_strategy,omitzero"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionFillActionFillType) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionFillActionFillType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionFillActionFillType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractAsyncParamsBrowserActionFillActionFillType](
		"mode", "type",
	)
	apijson.RegisterFieldValidator[ExtractAsyncParamsBrowserActionFillActionFillType](
		"mouse_movement_strategy", "linear", "ghost-cursor", "windmouse",
	)
	apijson.RegisterFieldValidator[ExtractAsyncParamsBrowserActionFillActionFillType](
		"typing_strategy", "simple", "distribution",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionFillActionFillTypeSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionFillActionFillTypeSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractAsyncParamsBrowserActionFillActionFillTypeSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionFillActionFillTypeSelectorUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionFillActionFillTypeDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionFillActionFillTypeDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractAsyncParamsBrowserActionFillActionFillTypeDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionFillActionFillTypeDelayUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionFillActionFillTypeRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionFillActionFillTypeRequiredString)
	OfExtractAsyncsBrowserActionFillActionFillTypeRequiredString param.Opt[ExtractAsyncParamsBrowserActionFillActionFillTypeRequiredString] `json:",omitzero,inline"`
	OfBool                                                       param.Opt[bool]                                                            `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionFillActionFillTypeRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionFillActionFillTypeRequiredString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionFillActionFillTypeRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionFillActionFillTypeRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionFillActionFillTypeRequiredString) {
		return &u.OfExtractAsyncsBrowserActionFillActionFillTypeRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionFillActionFillTypeRequiredString string

const (
	ExtractAsyncParamsBrowserActionFillActionFillTypeRequiredStringTrue  ExtractAsyncParamsBrowserActionFillActionFillTypeRequiredString = "true"
	ExtractAsyncParamsBrowserActionFillActionFillTypeRequiredStringFalse ExtractAsyncParamsBrowserActionFillActionFillTypeRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionFillActionFillTypeSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionFillActionFillTypeSkipString)
	OfExtractAsyncsBrowserActionFillActionFillTypeSkipString param.Opt[ExtractAsyncParamsBrowserActionFillActionFillTypeSkipString] `json:",omitzero,inline"`
	OfBool                                                   param.Opt[bool]                                                        `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionFillActionFillTypeSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionFillActionFillTypeSkipString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionFillActionFillTypeSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionFillActionFillTypeSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionFillActionFillTypeSkipString) {
		return &u.OfExtractAsyncsBrowserActionFillActionFillTypeSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionFillActionFillTypeSkipString string

const (
	ExtractAsyncParamsBrowserActionFillActionFillTypeSkipStringTrue  ExtractAsyncParamsBrowserActionFillActionFillTypeSkipString = "true"
	ExtractAsyncParamsBrowserActionFillActionFillTypeSkipStringFalse ExtractAsyncParamsBrowserActionFillActionFillTypeSkipString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionFillActionFillTypeTypingIntervalUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionFillActionFillTypeTypingIntervalUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractAsyncParamsBrowserActionFillActionFillTypeTypingIntervalUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionFillActionFillTypeTypingIntervalUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// The properties Mode, Selector, Value are required.
type ExtractAsyncParamsBrowserActionFillActionFillPaste struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector       ExtractAsyncParamsBrowserActionFillActionFillPasteSelectorUnion `json:"selector,omitzero,required"`
	Value          string                                                          `json:"value,required"`
	ClickOnElement param.Opt[bool]                                                 `json:"click_on_element,omitzero"`
	Scroll         param.Opt[bool]                                                 `json:"scroll,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay ExtractAsyncParamsBrowserActionFillActionFillPasteDelayUnion `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractAsyncParamsBrowserActionFillActionFillPasteRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractAsyncParamsBrowserActionFillActionFillPasteSkipUnion `json:"skip,omitzero"`
	// This field can be elided, and will marshal its zero value as "paste".
	Mode constant.Paste `json:"mode,required"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionFillActionFillPaste) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionFillActionFillPaste
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionFillActionFillPaste) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionFillActionFillPasteSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionFillActionFillPasteSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractAsyncParamsBrowserActionFillActionFillPasteSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionFillActionFillPasteSelectorUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionFillActionFillPasteDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionFillActionFillPasteDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractAsyncParamsBrowserActionFillActionFillPasteDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionFillActionFillPasteDelayUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionFillActionFillPasteRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionFillActionFillPasteRequiredString)
	OfExtractAsyncsBrowserActionFillActionFillPasteRequiredString param.Opt[ExtractAsyncParamsBrowserActionFillActionFillPasteRequiredString] `json:",omitzero,inline"`
	OfBool                                                        param.Opt[bool]                                                             `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionFillActionFillPasteRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionFillActionFillPasteRequiredString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionFillActionFillPasteRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionFillActionFillPasteRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionFillActionFillPasteRequiredString) {
		return &u.OfExtractAsyncsBrowserActionFillActionFillPasteRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionFillActionFillPasteRequiredString string

const (
	ExtractAsyncParamsBrowserActionFillActionFillPasteRequiredStringTrue  ExtractAsyncParamsBrowserActionFillActionFillPasteRequiredString = "true"
	ExtractAsyncParamsBrowserActionFillActionFillPasteRequiredStringFalse ExtractAsyncParamsBrowserActionFillActionFillPasteRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionFillActionFillPasteSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionFillActionFillPasteSkipString)
	OfExtractAsyncsBrowserActionFillActionFillPasteSkipString param.Opt[ExtractAsyncParamsBrowserActionFillActionFillPasteSkipString] `json:",omitzero,inline"`
	OfBool                                                    param.Opt[bool]                                                         `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionFillActionFillPasteSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionFillActionFillPasteSkipString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionFillActionFillPasteSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionFillActionFillPasteSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionFillActionFillPasteSkipString) {
		return &u.OfExtractAsyncsBrowserActionFillActionFillPasteSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionFillActionFillPasteSkipString string

const (
	ExtractAsyncParamsBrowserActionFillActionFillPasteSkipStringTrue  ExtractAsyncParamsBrowserActionFillActionFillPasteSkipString = "true"
	ExtractAsyncParamsBrowserActionFillActionFillPasteSkipStringFalse ExtractAsyncParamsBrowserActionFillActionFillPasteSkipString = "false"
)

// Retrieve browser cookies
//
// The property GetCookies is required.
type ExtractAsyncParamsBrowserActionGetCookiesAction struct {
	GetCookies ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesUnion `json:"get_cookies,omitzero,required"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionGetCookiesAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionGetCookiesAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionGetCookiesAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesUnion struct {
	OfBool                                                       param.Opt[bool]                                                  `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionGetCookiesActionGetCookiesObject *ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfExtractAsyncsBrowserActionGetCookiesActionGetCookiesObject)
}
func (u *ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionGetCookiesActionGetCookiesObject) {
		return u.OfExtractAsyncsBrowserActionGetCookiesActionGetCookiesObject
	}
	return nil
}

type ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObject struct {
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip        ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion `json:"skip,omitzero"`
	ExtraFields map[string]any                                                           `json:"-"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObject
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionGetCookiesActionGetCookiesObjectRequiredString)
	OfExtractAsyncsBrowserActionGetCookiesActionGetCookiesObjectRequiredString param.Opt[ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                     param.Opt[bool]                                                                          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionGetCookiesActionGetCookiesObjectRequiredString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionGetCookiesActionGetCookiesObjectRequiredString) {
		return &u.OfExtractAsyncsBrowserActionGetCookiesActionGetCookiesObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredString string

const (
	ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredStringTrue  ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredString = "true"
	ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredStringFalse ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionGetCookiesActionGetCookiesObjectSkipString)
	OfExtractAsyncsBrowserActionGetCookiesActionGetCookiesObjectSkipString param.Opt[ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                 param.Opt[bool]                                                                      `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionGetCookiesActionGetCookiesObjectSkipString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionGetCookiesActionGetCookiesObjectSkipString) {
		return &u.OfExtractAsyncsBrowserActionGetCookiesActionGetCookiesObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectSkipString string

const (
	ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectSkipStringTrue  ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectSkipString = "true"
	ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectSkipStringFalse ExtractAsyncParamsBrowserActionGetCookiesActionGetCookiesObjectSkipString = "false"
)

// Navigate to a URL
//
// The property Goto is required.
type ExtractAsyncParamsBrowserActionGotoAction struct {
	Goto ExtractAsyncParamsBrowserActionGotoActionGotoUnion `json:"goto,omitzero,required" format:"uri"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionGotoAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionGotoAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionGotoAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionGotoActionGotoUnion struct {
	OfString                                         param.Opt[string]                                    `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionGotoActionGotoObject *ExtractAsyncParamsBrowserActionGotoActionGotoObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionGotoActionGotoUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractAsyncsBrowserActionGotoActionGotoObject)
}
func (u *ExtractAsyncParamsBrowserActionGotoActionGotoUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionGotoActionGotoUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionGotoActionGotoObject) {
		return u.OfExtractAsyncsBrowserActionGotoActionGotoObject
	}
	return nil
}

// The property URL is required.
type ExtractAsyncParamsBrowserActionGotoActionGotoObject struct {
	URL     string            `json:"url,required" format:"uri"`
	Referer param.Opt[string] `json:"referer,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractAsyncParamsBrowserActionGotoActionGotoObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractAsyncParamsBrowserActionGotoActionGotoObjectSkipUnion `json:"skip,omitzero"`
	// Any of "load", "domcontentloaded", "networkidle0", "networkidle2".
	WaitUntil string `json:"wait_until,omitzero"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionGotoActionGotoObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionGotoActionGotoObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionGotoActionGotoObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractAsyncParamsBrowserActionGotoActionGotoObject](
		"wait_until", "load", "domcontentloaded", "networkidle0", "networkidle2",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionGotoActionGotoObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionGotoActionGotoObjectRequiredString)
	OfExtractAsyncsBrowserActionGotoActionGotoObjectRequiredString param.Opt[ExtractAsyncParamsBrowserActionGotoActionGotoObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                         param.Opt[bool]                                                              `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionGotoActionGotoObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionGotoActionGotoObjectRequiredString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionGotoActionGotoObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionGotoActionGotoObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionGotoActionGotoObjectRequiredString) {
		return &u.OfExtractAsyncsBrowserActionGotoActionGotoObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionGotoActionGotoObjectRequiredString string

const (
	ExtractAsyncParamsBrowserActionGotoActionGotoObjectRequiredStringTrue  ExtractAsyncParamsBrowserActionGotoActionGotoObjectRequiredString = "true"
	ExtractAsyncParamsBrowserActionGotoActionGotoObjectRequiredStringFalse ExtractAsyncParamsBrowserActionGotoActionGotoObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionGotoActionGotoObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionGotoActionGotoObjectSkipString)
	OfExtractAsyncsBrowserActionGotoActionGotoObjectSkipString param.Opt[ExtractAsyncParamsBrowserActionGotoActionGotoObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                     param.Opt[bool]                                                          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionGotoActionGotoObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionGotoActionGotoObjectSkipString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionGotoActionGotoObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionGotoActionGotoObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionGotoActionGotoObjectSkipString) {
		return &u.OfExtractAsyncsBrowserActionGotoActionGotoObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionGotoActionGotoObjectSkipString string

const (
	ExtractAsyncParamsBrowserActionGotoActionGotoObjectSkipStringTrue  ExtractAsyncParamsBrowserActionGotoActionGotoObjectSkipString = "true"
	ExtractAsyncParamsBrowserActionGotoActionGotoObjectSkipStringFalse ExtractAsyncParamsBrowserActionGotoActionGotoObjectSkipString = "false"
)

// Press a keyboard key
//
// The property Press is required.
type ExtractAsyncParamsBrowserActionPressAction struct {
	Press ExtractAsyncParamsBrowserActionPressActionPressUnion `json:"press,omitzero,required"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionPressAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionPressAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionPressAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionPressActionPressUnion struct {
	OfString                                           param.Opt[string]                                      `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionPressActionPressObject *ExtractAsyncParamsBrowserActionPressActionPressObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionPressActionPressUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractAsyncsBrowserActionPressActionPressObject)
}
func (u *ExtractAsyncParamsBrowserActionPressActionPressUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionPressActionPressUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionPressActionPressObject) {
		return u.OfExtractAsyncsBrowserActionPressActionPressObject
	}
	return nil
}

// The property Key is required.
type ExtractAsyncParamsBrowserActionPressActionPressObject struct {
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
	Delay ExtractAsyncParamsBrowserActionPressActionPressObjectDelayUnion `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractAsyncParamsBrowserActionPressActionPressObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractAsyncParamsBrowserActionPressActionPressObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionPressActionPressObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionPressActionPressObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionPressActionPressObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractAsyncParamsBrowserActionPressActionPressObject](
		"key", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "Power", "Eject", "Abort", "Help", "Backspace", "Tab", "Numpad5", "NumpadEnter", "Enter", "\r", "\n", "ShiftLeft", "ShiftRight", "ControlLeft", "ControlRight", "AltLeft", "AltRight", "Pause", "CapsLock", "Escape", "Convert", "NonConvert", "Space", "Numpad9", "PageUp", "Numpad3", "PageDown", "End", "Numpad1", "Home", "Numpad7", "ArrowLeft", "Numpad4", "Numpad8", "ArrowUp", "ArrowRight", "Numpad6", "Numpad2", "ArrowDown", "Select", "Open", "PrintScreen", "Insert", "Numpad0", "Delete", "NumpadDecimal", "Digit0", "Digit1", "Digit2", "Digit3", "Digit4", "Digit5", "Digit6", "Digit7", "Digit8", "Digit9", "KeyA", "KeyB", "KeyC", "KeyD", "KeyE", "KeyF", "KeyG", "KeyH", "KeyI", "KeyJ", "KeyK", "KeyL", "KeyM", "KeyN", "KeyO", "KeyP", "KeyQ", "KeyR", "KeyS", "KeyT", "KeyU", "KeyV", "KeyW", "KeyX", "KeyY", "KeyZ", "MetaLeft", "MetaRight", "ContextMenu", "NumpadMultiply", "NumpadAdd", "NumpadSubtract", "NumpadDivide", "F1", "F2", "F3", "F4", "F5", "F6", "F7", "F8", "F9", "F10", "F11", "F12", "F13", "F14", "F15", "F16", "F17", "F18", "F19", "F20", "F21", "F22", "F23", "F24", "NumLock", "ScrollLock", "AudioVolumeMute", "AudioVolumeDown", "AudioVolumeUp", "MediaTrackNext", "MediaTrackPrevious", "MediaStop", "MediaPlayPause", "Semicolon", "Equal", "NumpadEqual", "Comma", "Minus", "Period", "Slash", "Backquote", "BracketLeft", "Backslash", "BracketRight", "Quote", "AltGraph", "Props", "Cancel", "Clear", "Shift", "Control", "Alt", "Accept", "ModeChange", " ", "Print", "Execute", "\u0000", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "Meta", "*", "+", "-", "/", ";", "=", ",", ".", "`", "[", "\\", "]", "'", "Attn", "CrSel", "ExSel", "EraseEof", "Play", "ZoomOut", ")", "!", "@", "#", "$", "%", "^", "&", "(", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", ":", "<", "_", ">", "?", "~", "{", "|", "}", "\"", "SoftLeft", "SoftRight", "Camera", "Call", "EndCall", "VolumeDown", "VolumeUp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionPressActionPressObjectDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionPressActionPressObjectDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractAsyncParamsBrowserActionPressActionPressObjectDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionPressActionPressObjectDelayUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionPressActionPressObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionPressActionPressObjectRequiredString)
	OfExtractAsyncsBrowserActionPressActionPressObjectRequiredString param.Opt[ExtractAsyncParamsBrowserActionPressActionPressObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                           param.Opt[bool]                                                                `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionPressActionPressObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionPressActionPressObjectRequiredString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionPressActionPressObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionPressActionPressObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionPressActionPressObjectRequiredString) {
		return &u.OfExtractAsyncsBrowserActionPressActionPressObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionPressActionPressObjectRequiredString string

const (
	ExtractAsyncParamsBrowserActionPressActionPressObjectRequiredStringTrue  ExtractAsyncParamsBrowserActionPressActionPressObjectRequiredString = "true"
	ExtractAsyncParamsBrowserActionPressActionPressObjectRequiredStringFalse ExtractAsyncParamsBrowserActionPressActionPressObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionPressActionPressObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionPressActionPressObjectSkipString)
	OfExtractAsyncsBrowserActionPressActionPressObjectSkipString param.Opt[ExtractAsyncParamsBrowserActionPressActionPressObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                       param.Opt[bool]                                                            `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionPressActionPressObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionPressActionPressObjectSkipString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionPressActionPressObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionPressActionPressObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionPressActionPressObjectSkipString) {
		return &u.OfExtractAsyncsBrowserActionPressActionPressObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionPressActionPressObjectSkipString string

const (
	ExtractAsyncParamsBrowserActionPressActionPressObjectSkipStringTrue  ExtractAsyncParamsBrowserActionPressActionPressObjectSkipString = "true"
	ExtractAsyncParamsBrowserActionPressActionPressObjectSkipStringFalse ExtractAsyncParamsBrowserActionPressActionPressObjectSkipString = "false"
)

// Capture a page screenshot
//
// The property Screenshot is required.
type ExtractAsyncParamsBrowserActionScreenshotAction struct {
	Screenshot ExtractAsyncParamsBrowserActionScreenshotActionScreenshotUnion `json:"screenshot,omitzero,required"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionScreenshotAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionScreenshotAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionScreenshotAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionScreenshotActionScreenshotUnion struct {
	OfBool                                                       param.Opt[bool]                                                  `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionScreenshotActionScreenshotObject *ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionScreenshotActionScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfExtractAsyncsBrowserActionScreenshotActionScreenshotObject)
}
func (u *ExtractAsyncParamsBrowserActionScreenshotActionScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionScreenshotActionScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionScreenshotActionScreenshotObject) {
		return u.OfExtractAsyncsBrowserActionScreenshotActionScreenshotObject
	}
	return nil
}

type ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObject struct {
	FullPage param.Opt[bool]    `json:"full_page,omitzero"`
	Quality  param.Opt[float64] `json:"quality,omitzero"`
	// Any of "png", "jpeg", "webp".
	Format string `json:"format,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObject](
		"format", "png", "jpeg", "webp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionScreenshotActionScreenshotObjectRequiredString)
	OfExtractAsyncsBrowserActionScreenshotActionScreenshotObjectRequiredString param.Opt[ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                     param.Opt[bool]                                                                          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionScreenshotActionScreenshotObjectRequiredString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionScreenshotActionScreenshotObjectRequiredString) {
		return &u.OfExtractAsyncsBrowserActionScreenshotActionScreenshotObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectRequiredString string

const (
	ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectRequiredStringTrue  ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectRequiredString = "true"
	ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectRequiredStringFalse ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionScreenshotActionScreenshotObjectSkipString)
	OfExtractAsyncsBrowserActionScreenshotActionScreenshotObjectSkipString param.Opt[ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                 param.Opt[bool]                                                                      `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionScreenshotActionScreenshotObjectSkipString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionScreenshotActionScreenshotObjectSkipString) {
		return &u.OfExtractAsyncsBrowserActionScreenshotActionScreenshotObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectSkipString string

const (
	ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectSkipStringTrue  ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectSkipString = "true"
	ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectSkipStringFalse ExtractAsyncParamsBrowserActionScreenshotActionScreenshotObjectSkipString = "false"
)

// Scroll the page or an element
//
// The property Scroll is required.
type ExtractAsyncParamsBrowserActionScrollAction struct {
	Scroll ExtractAsyncParamsBrowserActionScrollActionScrollUnion `json:"scroll,omitzero,required"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionScrollAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionScrollAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionScrollAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionScrollActionScrollUnion struct {
	OfFloat                                              param.Opt[float64]                                       `json:",omitzero,inline"`
	OfString                                             param.Opt[string]                                        `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionScrollActionScrollObject *ExtractAsyncParamsBrowserActionScrollActionScrollObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionScrollActionScrollUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfExtractAsyncsBrowserActionScrollActionScrollObject)
}
func (u *ExtractAsyncParamsBrowserActionScrollActionScrollUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionScrollActionScrollUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionScrollActionScrollObject) {
		return u.OfExtractAsyncsBrowserActionScrollActionScrollObject
	}
	return nil
}

type ExtractAsyncParamsBrowserActionScrollActionScrollObject struct {
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	X       param.Opt[float64] `json:"x,omitzero"`
	Y       param.Opt[float64] `json:"y,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Container ExtractAsyncParamsBrowserActionScrollActionScrollObjectContainerUnion `json:"container,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractAsyncParamsBrowserActionScrollActionScrollObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractAsyncParamsBrowserActionScrollActionScrollObjectSkipUnion `json:"skip,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	To ExtractAsyncParamsBrowserActionScrollActionScrollObjectToUnion `json:"to,omitzero"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionScrollActionScrollObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionScrollActionScrollObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionScrollActionScrollObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionScrollActionScrollObjectContainerUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionScrollActionScrollObjectContainerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractAsyncParamsBrowserActionScrollActionScrollObjectContainerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionScrollActionScrollObjectContainerUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionScrollActionScrollObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionScrollActionScrollObjectRequiredString)
	OfExtractAsyncsBrowserActionScrollActionScrollObjectRequiredString param.Opt[ExtractAsyncParamsBrowserActionScrollActionScrollObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                             param.Opt[bool]                                                                  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionScrollActionScrollObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionScrollActionScrollObjectRequiredString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionScrollActionScrollObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionScrollActionScrollObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionScrollActionScrollObjectRequiredString) {
		return &u.OfExtractAsyncsBrowserActionScrollActionScrollObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionScrollActionScrollObjectRequiredString string

const (
	ExtractAsyncParamsBrowserActionScrollActionScrollObjectRequiredStringTrue  ExtractAsyncParamsBrowserActionScrollActionScrollObjectRequiredString = "true"
	ExtractAsyncParamsBrowserActionScrollActionScrollObjectRequiredStringFalse ExtractAsyncParamsBrowserActionScrollActionScrollObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionScrollActionScrollObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionScrollActionScrollObjectSkipString)
	OfExtractAsyncsBrowserActionScrollActionScrollObjectSkipString param.Opt[ExtractAsyncParamsBrowserActionScrollActionScrollObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                         param.Opt[bool]                                                              `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionScrollActionScrollObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionScrollActionScrollObjectSkipString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionScrollActionScrollObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionScrollActionScrollObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionScrollActionScrollObjectSkipString) {
		return &u.OfExtractAsyncsBrowserActionScrollActionScrollObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionScrollActionScrollObjectSkipString string

const (
	ExtractAsyncParamsBrowserActionScrollActionScrollObjectSkipStringTrue  ExtractAsyncParamsBrowserActionScrollActionScrollObjectSkipString = "true"
	ExtractAsyncParamsBrowserActionScrollActionScrollObjectSkipStringFalse ExtractAsyncParamsBrowserActionScrollActionScrollObjectSkipString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionScrollActionScrollObjectToUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionScrollActionScrollObjectToUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractAsyncParamsBrowserActionScrollActionScrollObjectToUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionScrollActionScrollObjectToUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionWaitAction struct {
	Wait ExtractAsyncParamsBrowserActionWaitActionWaitUnion `json:"wait,omitzero,required"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionWaitAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionWaitAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionWaitAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionWaitActionWaitUnion struct {
	OfFloat                                          param.Opt[float64]                                   `json:",omitzero,inline"`
	OfString                                         param.Opt[string]                                    `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionWaitActionWaitObject *ExtractAsyncParamsBrowserActionWaitActionWaitObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionWaitActionWaitUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfExtractAsyncsBrowserActionWaitActionWaitObject)
}
func (u *ExtractAsyncParamsBrowserActionWaitActionWaitUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionWaitActionWaitUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionWaitActionWaitObject) {
		return u.OfExtractAsyncsBrowserActionWaitActionWaitObject
	}
	return nil
}

// The property Duration is required.
type ExtractAsyncParamsBrowserActionWaitActionWaitObject struct {
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Duration ExtractAsyncParamsBrowserActionWaitActionWaitObjectDurationUnion `json:"duration,omitzero,required"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractAsyncParamsBrowserActionWaitActionWaitObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractAsyncParamsBrowserActionWaitActionWaitObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionWaitActionWaitObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionWaitActionWaitObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionWaitActionWaitObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionWaitActionWaitObjectDurationUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionWaitActionWaitObjectDurationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractAsyncParamsBrowserActionWaitActionWaitObjectDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionWaitActionWaitObjectDurationUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionWaitActionWaitObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionWaitActionWaitObjectRequiredString)
	OfExtractAsyncsBrowserActionWaitActionWaitObjectRequiredString param.Opt[ExtractAsyncParamsBrowserActionWaitActionWaitObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                         param.Opt[bool]                                                              `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionWaitActionWaitObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionWaitActionWaitObjectRequiredString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionWaitActionWaitObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionWaitActionWaitObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionWaitActionWaitObjectRequiredString) {
		return &u.OfExtractAsyncsBrowserActionWaitActionWaitObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionWaitActionWaitObjectRequiredString string

const (
	ExtractAsyncParamsBrowserActionWaitActionWaitObjectRequiredStringTrue  ExtractAsyncParamsBrowserActionWaitActionWaitObjectRequiredString = "true"
	ExtractAsyncParamsBrowserActionWaitActionWaitObjectRequiredStringFalse ExtractAsyncParamsBrowserActionWaitActionWaitObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionWaitActionWaitObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionWaitActionWaitObjectSkipString)
	OfExtractAsyncsBrowserActionWaitActionWaitObjectSkipString param.Opt[ExtractAsyncParamsBrowserActionWaitActionWaitObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                     param.Opt[bool]                                                          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionWaitActionWaitObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionWaitActionWaitObjectSkipString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionWaitActionWaitObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionWaitActionWaitObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionWaitActionWaitObjectSkipString) {
		return &u.OfExtractAsyncsBrowserActionWaitActionWaitObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionWaitActionWaitObjectSkipString string

const (
	ExtractAsyncParamsBrowserActionWaitActionWaitObjectSkipStringTrue  ExtractAsyncParamsBrowserActionWaitActionWaitObjectSkipString = "true"
	ExtractAsyncParamsBrowserActionWaitActionWaitObjectSkipStringFalse ExtractAsyncParamsBrowserActionWaitActionWaitObjectSkipString = "false"
)

// Wait for an element to appear or reach a specific state
//
// The property WaitForElement is required.
type ExtractAsyncParamsBrowserActionWaitForElementAction struct {
	WaitForElement ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementUnion `json:"wait_for_element,omitzero,required"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionWaitForElementAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionWaitForElementAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionWaitForElementAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementUnion struct {
	OfString                                                             param.Opt[string]                                                        `json:",omitzero,inline"`
	OfStringArray                                                        []string                                                                 `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionWaitForElementActionWaitForElementObject *ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray, u.OfExtractAsyncsBrowserActionWaitForElementActionWaitForElementObject)
}
func (u *ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionWaitForElementActionWaitForElementObject) {
		return u.OfExtractAsyncsBrowserActionWaitForElementActionWaitForElementObject
	}
	return nil
}

// The property Selector is required.
type ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObject struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion `json:"selector,omitzero,required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion) asAny() any {
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
type ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionWaitForElementActionWaitForElementObjectRequiredString)
	OfExtractAsyncsBrowserActionWaitForElementActionWaitForElementObjectRequiredString param.Opt[ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                             param.Opt[bool]                                                                                  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionWaitForElementActionWaitForElementObjectRequiredString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionWaitForElementActionWaitForElementObjectRequiredString) {
		return &u.OfExtractAsyncsBrowserActionWaitForElementActionWaitForElementObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredString string

const (
	ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredStringTrue  ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredString = "true"
	ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredStringFalse ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionWaitForElementActionWaitForElementObjectSkipString)
	OfExtractAsyncsBrowserActionWaitForElementActionWaitForElementObjectSkipString param.Opt[ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                         param.Opt[bool]                                                                              `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionWaitForElementActionWaitForElementObjectSkipString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionWaitForElementActionWaitForElementObjectSkipString) {
		return &u.OfExtractAsyncsBrowserActionWaitForElementActionWaitForElementObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectSkipString string

const (
	ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectSkipStringTrue  ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectSkipString = "true"
	ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectSkipStringFalse ExtractAsyncParamsBrowserActionWaitForElementActionWaitForElementObjectSkipString = "false"
)

// Wait for page navigation to complete
//
// The property WaitForNavigation is required.
type ExtractAsyncParamsBrowserActionWaitForNavigationAction struct {
	WaitForNavigation ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion `json:"wait_for_navigation,omitzero,required"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionWaitForNavigationAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionWaitForNavigationAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionWaitForNavigationAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationString)
	OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationString param.Opt[string]                                                              `json:",omitzero,inline"`
	OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationObject *ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationString, u.OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationObject)
}
func (u *ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationString) {
		return &u.OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationString
	} else if !param.IsOmitted(u.OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationObject) {
		return u.OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationObject
	}
	return nil
}

type ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationString string

const (
	ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationStringLoad             ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationString = "load"
	ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationStringDomcontentloaded ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationString = "domcontentloaded"
	ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationStringNetworkidle0     ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationString = "networkidle0"
	ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationStringNetworkidle2     ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationString = "networkidle2"
)

// The property Navigation is required.
type ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObject struct {
	// Any of "load", "domcontentloaded", "networkidle0", "networkidle2".
	Navigation string `json:"navigation,omitzero,required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObject](
		"navigation", "load", "domcontentloaded", "networkidle0", "networkidle2",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString)
	OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString param.Opt[ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                                   param.Opt[bool]                                                                                        `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString) {
		return &u.OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString string

const (
	ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredStringTrue  ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString = "true"
	ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredStringFalse ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString)
	OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString param.Opt[ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                               param.Opt[bool]                                                                                    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString, u.OfBool)
}
func (u *ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString) {
		return &u.OfExtractAsyncsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString string

const (
	ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipStringTrue  ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString = "true"
	ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipStringFalse ExtractAsyncParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString = "false"
)

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
	Value string `json:"value,required"`
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

type ExtractRunParams struct {
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
	Browser ExtractRunParamsBrowserUnion `json:"browser,omitzero"`
	// Array of browser automation actions to execute sequentially
	BrowserActions []ExtractRunParamsBrowserActionUnion `json:"browser_actions,omitzero"`
	// Browser cookies as array of cookie objects
	Cookies ExtractRunParamsCookiesUnion `json:"cookies,omitzero"`
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
	Country ExtractRunParamsCountry `json:"country,omitzero"`
	// Device type for browser emulation
	//
	// Any of "desktop", "mobile", "tablet".
	Device ExtractRunParamsDevice `json:"device,omitzero"`
	// Browser driver to use
	//
	// Any of "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro".
	Driver ExtractRunParamsDriver `json:"driver,omitzero"`
	// Expected HTTP status codes for successful requests
	ExpectedStatusCodes []int64 `json:"expected_status_codes,omitzero"`
	// List of acceptable response formats in order of preference
	//
	// Any of "html", "markdown".
	Formats []string `json:"formats,omitzero"`
	// Custom HTTP headers to include in the request
	Headers map[string]ExtractRunParamsHeaderUnion `json:"headers,omitzero"`
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
	Locale ExtractRunParamsLocale `json:"locale,omitzero"`
	// HTTP method for the request
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method ExtractRunParamsMethod `json:"method,omitzero"`
	// Filters for capturing network traffic
	NetworkCapture []ExtractRunParamsNetworkCapture `json:"network_capture,omitzero"`
	// Operating system to emulate
	//
	// Any of "windows", "mac os", "linux", "android", "ios".
	Os ExtractRunParamsOs `json:"os,omitzero"`
	// Custom parser configuration as a key-value map
	Parser ExtractRunParamsParserUnion `json:"parser,omitzero"`
	// Referrer policy for the request
	//
	// Any of "random", "no-referer", "same-origin", "google", "bing", "facebook",
	// "twitter", "instagram".
	ReferrerType ExtractRunParamsReferrerType `json:"referrer_type,omitzero"`
	Session      ExtractRunParamsSession      `json:"session,omitzero"`
	// Skills or capabilities required for the request
	Skill ExtractRunParamsSkillUnion `json:"skill,omitzero"`
	// US state for geolocation (only valid when country is US)
	//
	// Any of "AL", "AK", "AS", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA",
	// "GU", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI",
	// "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "MP",
	// "OH", "OK", "OR", "PA", "PR", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA",
	// "VI", "WA", "WV", "WI", "WY".
	State ExtractRunParamsState `json:"state,omitzero"`
	paramObj
}

func (r ExtractRunParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserString)
	OfExtractRunsBrowserString param.Opt[string]              `json:",omitzero,inline"`
	OfExtractRunsBrowserObject *ExtractRunParamsBrowserObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserString, u.OfExtractRunsBrowserObject)
}
func (u *ExtractRunParamsBrowserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserString) {
		return &u.OfExtractRunsBrowserString
	} else if !param.IsOmitted(u.OfExtractRunsBrowserObject) {
		return u.OfExtractRunsBrowserObject
	}
	return nil
}

// Browser type to emulate
type ExtractRunParamsBrowserString string

const (
	ExtractRunParamsBrowserStringChrome  ExtractRunParamsBrowserString = "chrome"
	ExtractRunParamsBrowserStringFirefox ExtractRunParamsBrowserString = "firefox"
)

// The property Name is required.
type ExtractRunParamsBrowserObject struct {
	// Any of "chrome", "firefox".
	Name string `json:"name,omitzero,required"`
	// Specific browser version to emulate
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r ExtractRunParamsBrowserObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractRunParamsBrowserObject](
		"name", "chrome", "firefox",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionUnion struct {
	OfExtractRunsBrowserActionAutoScrollAction        *ExtractRunParamsBrowserActionAutoScrollAction        `json:",omitzero,inline"`
	OfExtractRunsBrowserActionClickAction             *ExtractRunParamsBrowserActionClickAction             `json:",omitzero,inline"`
	OfExtractRunsBrowserActionEvalAction              *ExtractRunParamsBrowserActionEvalAction              `json:",omitzero,inline"`
	OfExtractRunsBrowserActionFetchAction             *ExtractRunParamsBrowserActionFetchAction             `json:",omitzero,inline"`
	OfExtractRunsBrowserActionFillAction              *ExtractRunParamsBrowserActionFillAction              `json:",omitzero,inline"`
	OfExtractRunsBrowserActionGetCookiesAction        *ExtractRunParamsBrowserActionGetCookiesAction        `json:",omitzero,inline"`
	OfExtractRunsBrowserActionGotoAction              *ExtractRunParamsBrowserActionGotoAction              `json:",omitzero,inline"`
	OfExtractRunsBrowserActionPressAction             *ExtractRunParamsBrowserActionPressAction             `json:",omitzero,inline"`
	OfExtractRunsBrowserActionScreenshotAction        *ExtractRunParamsBrowserActionScreenshotAction        `json:",omitzero,inline"`
	OfExtractRunsBrowserActionScrollAction            *ExtractRunParamsBrowserActionScrollAction            `json:",omitzero,inline"`
	OfExtractRunsBrowserActionWaitAction              *ExtractRunParamsBrowserActionWaitAction              `json:",omitzero,inline"`
	OfExtractRunsBrowserActionWaitForElementAction    *ExtractRunParamsBrowserActionWaitForElementAction    `json:",omitzero,inline"`
	OfExtractRunsBrowserActionWaitForNavigationAction *ExtractRunParamsBrowserActionWaitForNavigationAction `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionAutoScrollAction,
		u.OfExtractRunsBrowserActionClickAction,
		u.OfExtractRunsBrowserActionEvalAction,
		u.OfExtractRunsBrowserActionFetchAction,
		u.OfExtractRunsBrowserActionFillAction,
		u.OfExtractRunsBrowserActionGetCookiesAction,
		u.OfExtractRunsBrowserActionGotoAction,
		u.OfExtractRunsBrowserActionPressAction,
		u.OfExtractRunsBrowserActionScreenshotAction,
		u.OfExtractRunsBrowserActionScrollAction,
		u.OfExtractRunsBrowserActionWaitAction,
		u.OfExtractRunsBrowserActionWaitForElementAction,
		u.OfExtractRunsBrowserActionWaitForNavigationAction)
}
func (u *ExtractRunParamsBrowserActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionAutoScrollAction) {
		return u.OfExtractRunsBrowserActionAutoScrollAction
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionClickAction) {
		return u.OfExtractRunsBrowserActionClickAction
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionEvalAction) {
		return u.OfExtractRunsBrowserActionEvalAction
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionFetchAction) {
		return u.OfExtractRunsBrowserActionFetchAction
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionFillAction) {
		return u.OfExtractRunsBrowserActionFillAction
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionGetCookiesAction) {
		return u.OfExtractRunsBrowserActionGetCookiesAction
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionGotoAction) {
		return u.OfExtractRunsBrowserActionGotoAction
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionPressAction) {
		return u.OfExtractRunsBrowserActionPressAction
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionScreenshotAction) {
		return u.OfExtractRunsBrowserActionScreenshotAction
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionScrollAction) {
		return u.OfExtractRunsBrowserActionScrollAction
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionWaitAction) {
		return u.OfExtractRunsBrowserActionWaitAction
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionWaitForElementAction) {
		return u.OfExtractRunsBrowserActionWaitForElementAction
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionWaitForNavigationAction) {
		return u.OfExtractRunsBrowserActionWaitForNavigationAction
	}
	return nil
}

// Continuously scroll to load dynamic content
//
// The property AutoScroll is required.
type ExtractRunParamsBrowserActionAutoScrollAction struct {
	AutoScroll ExtractRunParamsBrowserActionAutoScrollActionAutoScrollUnion `json:"auto_scroll,omitzero,required"`
	paramObj
}

func (r ExtractRunParamsBrowserActionAutoScrollAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionAutoScrollAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionAutoScrollAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionAutoScrollActionAutoScrollUnion struct {
	OfBool                                                     param.Opt[bool]                                                `json:",omitzero,inline"`
	OfFloat                                                    param.Opt[float64]                                             `json:",omitzero,inline"`
	OfString                                                   param.Opt[string]                                              `json:",omitzero,inline"`
	OfExtractRunsBrowserActionAutoScrollActionAutoScrollObject *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionAutoScrollActionAutoScrollUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfExtractRunsBrowserActionAutoScrollActionAutoScrollObject)
}
func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionAutoScrollActionAutoScrollObject) {
		return u.OfExtractRunsBrowserActionAutoScrollActionAutoScrollObject
	}
	return nil
}

type ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObject struct {
	StepSize param.Opt[float64] `json:"step_size,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	ClickSelector ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion `json:"click_selector,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Container ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion `json:"container,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	DelayAfterScroll ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion `json:"delay_after_scroll,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	IdleTimeout ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion `json:"idle_timeout,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	LoadingSelector ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion `json:"loading_selector,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	MaxDuration ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion `json:"max_duration,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	PauseOnSelector ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion `json:"pause_on_selector,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion) asAny() any {
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
type ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion) asAny() any {
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
type ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion) asAny() any {
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
type ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion) asAny() any {
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
type ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion) asAny() any {
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
type ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion) asAny() any {
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
type ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion) asAny() any {
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
type ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionAutoScrollActionAutoScrollObjectRequiredString)
	OfExtractRunsBrowserActionAutoScrollActionAutoScrollObjectRequiredString param.Opt[ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                   param.Opt[bool]                                                                        `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionAutoScrollActionAutoScrollObjectRequiredString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionAutoScrollActionAutoScrollObjectRequiredString) {
		return &u.OfExtractRunsBrowserActionAutoScrollActionAutoScrollObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredString string

const (
	ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredStringTrue  ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredString = "true"
	ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredStringFalse ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionAutoScrollActionAutoScrollObjectSkipString)
	OfExtractRunsBrowserActionAutoScrollActionAutoScrollObjectSkipString param.Opt[ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                               param.Opt[bool]                                                                    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionAutoScrollActionAutoScrollObjectSkipString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionAutoScrollActionAutoScrollObjectSkipString) {
		return &u.OfExtractRunsBrowserActionAutoScrollActionAutoScrollObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectSkipString string

const (
	ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectSkipStringTrue  ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectSkipString = "true"
	ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectSkipStringFalse ExtractRunParamsBrowserActionAutoScrollActionAutoScrollObjectSkipString = "false"
)

// Click on an element by selector
//
// The property Click is required.
type ExtractRunParamsBrowserActionClickAction struct {
	Click ExtractRunParamsBrowserActionClickActionClickUnion `json:"click,omitzero,required"`
	paramObj
}

func (r ExtractRunParamsBrowserActionClickAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionClickAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionClickAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionClickActionClickUnion struct {
	OfString                                         param.Opt[string]                                    `json:",omitzero,inline"`
	OfStringArray                                    []string                                             `json:",omitzero,inline"`
	OfExtractRunsBrowserActionClickActionClickObject *ExtractRunParamsBrowserActionClickActionClickObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionClickActionClickUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray, u.OfExtractRunsBrowserActionClickActionClickObject)
}
func (u *ExtractRunParamsBrowserActionClickActionClickUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionClickActionClickUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionClickActionClickObject) {
		return u.OfExtractRunsBrowserActionClickActionClickObject
	}
	return nil
}

// The property Selector is required.
type ExtractRunParamsBrowserActionClickActionClickObject struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector ExtractRunParamsBrowserActionClickActionClickObjectSelectorUnion `json:"selector,omitzero,required"`
	Count    param.Opt[float64]                                               `json:"count,omitzero"`
	OffsetX  param.Opt[int64]                                                 `json:"offset_x,omitzero"`
	OffsetY  param.Opt[int64]                                                 `json:"offset_y,omitzero"`
	Scroll   param.Opt[bool]                                                  `json:"scroll,omitzero"`
	Steps    param.Opt[float64]                                               `json:"steps,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay ExtractRunParamsBrowserActionClickActionClickObjectDelayUnion `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractRunParamsBrowserActionClickActionClickObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractRunParamsBrowserActionClickActionClickObjectSkipUnion `json:"skip,omitzero"`
	// Any of "linear", "ghost-cursor", "windmouse".
	Strategy string `json:"strategy,omitzero"`
	paramObj
}

func (r ExtractRunParamsBrowserActionClickActionClickObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionClickActionClickObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionClickActionClickObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractRunParamsBrowserActionClickActionClickObject](
		"strategy", "linear", "ghost-cursor", "windmouse",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionClickActionClickObjectSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionClickActionClickObjectSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractRunParamsBrowserActionClickActionClickObjectSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionClickActionClickObjectSelectorUnion) asAny() any {
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
type ExtractRunParamsBrowserActionClickActionClickObjectDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionClickActionClickObjectDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractRunParamsBrowserActionClickActionClickObjectDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionClickActionClickObjectDelayUnion) asAny() any {
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
type ExtractRunParamsBrowserActionClickActionClickObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionClickActionClickObjectRequiredString)
	OfExtractRunsBrowserActionClickActionClickObjectRequiredString param.Opt[ExtractRunParamsBrowserActionClickActionClickObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                         param.Opt[bool]                                                              `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionClickActionClickObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionClickActionClickObjectRequiredString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionClickActionClickObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionClickActionClickObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionClickActionClickObjectRequiredString) {
		return &u.OfExtractRunsBrowserActionClickActionClickObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionClickActionClickObjectRequiredString string

const (
	ExtractRunParamsBrowserActionClickActionClickObjectRequiredStringTrue  ExtractRunParamsBrowserActionClickActionClickObjectRequiredString = "true"
	ExtractRunParamsBrowserActionClickActionClickObjectRequiredStringFalse ExtractRunParamsBrowserActionClickActionClickObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionClickActionClickObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionClickActionClickObjectSkipString)
	OfExtractRunsBrowserActionClickActionClickObjectSkipString param.Opt[ExtractRunParamsBrowserActionClickActionClickObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                     param.Opt[bool]                                                          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionClickActionClickObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionClickActionClickObjectSkipString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionClickActionClickObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionClickActionClickObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionClickActionClickObjectSkipString) {
		return &u.OfExtractRunsBrowserActionClickActionClickObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionClickActionClickObjectSkipString string

const (
	ExtractRunParamsBrowserActionClickActionClickObjectSkipStringTrue  ExtractRunParamsBrowserActionClickActionClickObjectSkipString = "true"
	ExtractRunParamsBrowserActionClickActionClickObjectSkipStringFalse ExtractRunParamsBrowserActionClickActionClickObjectSkipString = "false"
)

// Execute JavaScript code in page context
//
// The property Eval is required.
type ExtractRunParamsBrowserActionEvalAction struct {
	Eval ExtractRunParamsBrowserActionEvalActionEvalUnion `json:"eval,omitzero,required"`
	paramObj
}

func (r ExtractRunParamsBrowserActionEvalAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionEvalAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionEvalAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionEvalActionEvalUnion struct {
	OfString                                       param.Opt[string]                                  `json:",omitzero,inline"`
	OfExtractRunsBrowserActionEvalActionEvalObject *ExtractRunParamsBrowserActionEvalActionEvalObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionEvalActionEvalUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractRunsBrowserActionEvalActionEvalObject)
}
func (u *ExtractRunParamsBrowserActionEvalActionEvalUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionEvalActionEvalUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionEvalActionEvalObject) {
		return u.OfExtractRunsBrowserActionEvalActionEvalObject
	}
	return nil
}

// The property Code is required.
type ExtractRunParamsBrowserActionEvalActionEvalObject struct {
	Code string `json:"code,required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractRunParamsBrowserActionEvalActionEvalObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractRunParamsBrowserActionEvalActionEvalObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractRunParamsBrowserActionEvalActionEvalObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionEvalActionEvalObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionEvalActionEvalObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionEvalActionEvalObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionEvalActionEvalObjectRequiredString)
	OfExtractRunsBrowserActionEvalActionEvalObjectRequiredString param.Opt[ExtractRunParamsBrowserActionEvalActionEvalObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                       param.Opt[bool]                                                            `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionEvalActionEvalObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionEvalActionEvalObjectRequiredString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionEvalActionEvalObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionEvalActionEvalObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionEvalActionEvalObjectRequiredString) {
		return &u.OfExtractRunsBrowserActionEvalActionEvalObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionEvalActionEvalObjectRequiredString string

const (
	ExtractRunParamsBrowserActionEvalActionEvalObjectRequiredStringTrue  ExtractRunParamsBrowserActionEvalActionEvalObjectRequiredString = "true"
	ExtractRunParamsBrowserActionEvalActionEvalObjectRequiredStringFalse ExtractRunParamsBrowserActionEvalActionEvalObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionEvalActionEvalObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionEvalActionEvalObjectSkipString)
	OfExtractRunsBrowserActionEvalActionEvalObjectSkipString param.Opt[ExtractRunParamsBrowserActionEvalActionEvalObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                   param.Opt[bool]                                                        `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionEvalActionEvalObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionEvalActionEvalObjectSkipString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionEvalActionEvalObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionEvalActionEvalObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionEvalActionEvalObjectSkipString) {
		return &u.OfExtractRunsBrowserActionEvalActionEvalObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionEvalActionEvalObjectSkipString string

const (
	ExtractRunParamsBrowserActionEvalActionEvalObjectSkipStringTrue  ExtractRunParamsBrowserActionEvalActionEvalObjectSkipString = "true"
	ExtractRunParamsBrowserActionEvalActionEvalObjectSkipStringFalse ExtractRunParamsBrowserActionEvalActionEvalObjectSkipString = "false"
)

// Make an HTTP request in browser context
//
// The property Fetch is required.
type ExtractRunParamsBrowserActionFetchAction struct {
	Fetch ExtractRunParamsBrowserActionFetchActionFetchUnion `json:"fetch,omitzero,required" format:"uri"`
	paramObj
}

func (r ExtractRunParamsBrowserActionFetchAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionFetchAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionFetchAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionFetchActionFetchUnion struct {
	OfString                                         param.Opt[string]                                    `json:",omitzero,inline"`
	OfExtractRunsBrowserActionFetchActionFetchObject *ExtractRunParamsBrowserActionFetchActionFetchObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionFetchActionFetchUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractRunsBrowserActionFetchActionFetchObject)
}
func (u *ExtractRunParamsBrowserActionFetchActionFetchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionFetchActionFetchUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionFetchActionFetchObject) {
		return u.OfExtractRunsBrowserActionFetchActionFetchObject
	}
	return nil
}

// The property URL is required.
type ExtractRunParamsBrowserActionFetchActionFetchObject struct {
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
	Required ExtractRunParamsBrowserActionFetchActionFetchObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractRunParamsBrowserActionFetchActionFetchObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractRunParamsBrowserActionFetchActionFetchObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionFetchActionFetchObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionFetchActionFetchObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractRunParamsBrowserActionFetchActionFetchObject](
		"method", "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionFetchActionFetchObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionFetchActionFetchObjectRequiredString)
	OfExtractRunsBrowserActionFetchActionFetchObjectRequiredString param.Opt[ExtractRunParamsBrowserActionFetchActionFetchObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                         param.Opt[bool]                                                              `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionFetchActionFetchObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionFetchActionFetchObjectRequiredString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionFetchActionFetchObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionFetchActionFetchObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionFetchActionFetchObjectRequiredString) {
		return &u.OfExtractRunsBrowserActionFetchActionFetchObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionFetchActionFetchObjectRequiredString string

const (
	ExtractRunParamsBrowserActionFetchActionFetchObjectRequiredStringTrue  ExtractRunParamsBrowserActionFetchActionFetchObjectRequiredString = "true"
	ExtractRunParamsBrowserActionFetchActionFetchObjectRequiredStringFalse ExtractRunParamsBrowserActionFetchActionFetchObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionFetchActionFetchObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionFetchActionFetchObjectSkipString)
	OfExtractRunsBrowserActionFetchActionFetchObjectSkipString param.Opt[ExtractRunParamsBrowserActionFetchActionFetchObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                     param.Opt[bool]                                                          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionFetchActionFetchObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionFetchActionFetchObjectSkipString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionFetchActionFetchObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionFetchActionFetchObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionFetchActionFetchObjectSkipString) {
		return &u.OfExtractRunsBrowserActionFetchActionFetchObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionFetchActionFetchObjectSkipString string

const (
	ExtractRunParamsBrowserActionFetchActionFetchObjectSkipStringTrue  ExtractRunParamsBrowserActionFetchActionFetchObjectSkipString = "true"
	ExtractRunParamsBrowserActionFetchActionFetchObjectSkipStringFalse ExtractRunParamsBrowserActionFetchActionFetchObjectSkipString = "false"
)

// Fill text into an input field
//
// The property Fill is required.
type ExtractRunParamsBrowserActionFillAction struct {
	// Fill options with mode-specific fields. Use "type" mode for behavioral typing
	// simulation, or "paste" mode for instant paste.
	Fill ExtractRunParamsBrowserActionFillActionFillUnion `json:"fill,omitzero,required"`
	paramObj
}

func (r ExtractRunParamsBrowserActionFillAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionFillAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionFillAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionFillActionFillUnion struct {
	OfType  *ExtractRunParamsBrowserActionFillActionFillType  `json:",omitzero,inline"`
	OfPaste *ExtractRunParamsBrowserActionFillActionFillPaste `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionFillActionFillUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfType, u.OfPaste)
}
func (u *ExtractRunParamsBrowserActionFillActionFillUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionFillActionFillUnion) asAny() any {
	if !param.IsOmitted(u.OfType) {
		return u.OfType
	} else if !param.IsOmitted(u.OfPaste) {
		return u.OfPaste
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractRunParamsBrowserActionFillActionFillUnion) GetMouseMovementStrategy() *string {
	if vt := u.OfType; vt != nil {
		return &vt.MouseMovementStrategy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractRunParamsBrowserActionFillActionFillUnion) GetTypingInterval() *ExtractRunParamsBrowserActionFillActionFillTypeTypingIntervalUnion {
	if vt := u.OfType; vt != nil {
		return &vt.TypingInterval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractRunParamsBrowserActionFillActionFillUnion) GetTypingStrategy() *string {
	if vt := u.OfType; vt != nil {
		return &vt.TypingStrategy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractRunParamsBrowserActionFillActionFillUnion) GetValue() *string {
	if vt := u.OfType; vt != nil {
		return (*string)(&vt.Value)
	} else if vt := u.OfPaste; vt != nil {
		return (*string)(&vt.Value)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractRunParamsBrowserActionFillActionFillUnion) GetClickOnElement() *bool {
	if vt := u.OfType; vt != nil && vt.ClickOnElement.Valid() {
		return &vt.ClickOnElement.Value
	} else if vt := u.OfPaste; vt != nil && vt.ClickOnElement.Valid() {
		return &vt.ClickOnElement.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractRunParamsBrowserActionFillActionFillUnion) GetMode() *string {
	if vt := u.OfType; vt != nil {
		return (*string)(&vt.Mode)
	} else if vt := u.OfPaste; vt != nil {
		return (*string)(&vt.Mode)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractRunParamsBrowserActionFillActionFillUnion) GetScroll() *bool {
	if vt := u.OfType; vt != nil && vt.Scroll.Valid() {
		return &vt.Scroll.Value
	} else if vt := u.OfPaste; vt != nil && vt.Scroll.Valid() {
		return &vt.Scroll.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractRunParamsBrowserActionFillActionFillUnion) GetTimeout() *float64 {
	if vt := u.OfType; vt != nil && vt.Timeout.Valid() {
		return &vt.Timeout.Value
	} else if vt := u.OfPaste; vt != nil && vt.Timeout.Valid() {
		return &vt.Timeout.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractRunParamsBrowserActionFillActionFillUnion) GetVisible() *bool {
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
func (u ExtractRunParamsBrowserActionFillActionFillUnion) GetSelector() (res extractRunParamsBrowserActionFillActionFillUnionSelector) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Selector.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Selector.asAny()
	}
	return
}

// Can have the runtime types [*string], [\*[]string]
type extractRunParamsBrowserActionFillActionFillUnionSelector struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extractRunParamsBrowserActionFillActionFillUnionSelector) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ExtractRunParamsBrowserActionFillActionFillUnion) GetDelay() (res extractRunParamsBrowserActionFillActionFillUnionDelay) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Delay.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Delay.asAny()
	}
	return
}

// Can have the runtime types [*float64], [*string]
type extractRunParamsBrowserActionFillActionFillUnionDelay struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *float64:
//	case *string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extractRunParamsBrowserActionFillActionFillUnionDelay) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ExtractRunParamsBrowserActionFillActionFillUnion) GetRequired() (res extractRunParamsBrowserActionFillActionFillUnionRequired) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Required.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Required.asAny()
	}
	return
}

// Can have the runtime types [*string], [*bool]
type extractRunParamsBrowserActionFillActionFillUnionRequired struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extractRunParamsBrowserActionFillActionFillUnionRequired) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ExtractRunParamsBrowserActionFillActionFillUnion) GetSkip() (res extractRunParamsBrowserActionFillActionFillUnionSkip) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Skip.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Skip.asAny()
	}
	return
}

// Can have the runtime types [*string], [*bool]
type extractRunParamsBrowserActionFillActionFillUnionSkip struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extractRunParamsBrowserActionFillActionFillUnionSkip) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[ExtractRunParamsBrowserActionFillActionFillUnion](
		"mode",
		apijson.Discriminator[ExtractRunParamsBrowserActionFillActionFillType]("type"),
		apijson.Discriminator[ExtractRunParamsBrowserActionFillActionFillPaste]("paste"),
	)
}

// The properties Selector, Value are required.
type ExtractRunParamsBrowserActionFillActionFillType struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector       ExtractRunParamsBrowserActionFillActionFillTypeSelectorUnion `json:"selector,omitzero,required"`
	Value          string                                                       `json:"value,required"`
	ClickOnElement param.Opt[bool]                                              `json:"click_on_element,omitzero"`
	Scroll         param.Opt[bool]                                              `json:"scroll,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay ExtractRunParamsBrowserActionFillActionFillTypeDelayUnion `json:"delay,omitzero"`
	// Any of "type".
	Mode string `json:"mode,omitzero"`
	// Any of "linear", "ghost-cursor", "windmouse".
	MouseMovementStrategy string `json:"mouse_movement_strategy,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractRunParamsBrowserActionFillActionFillTypeRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractRunParamsBrowserActionFillActionFillTypeSkipUnion `json:"skip,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	TypingInterval ExtractRunParamsBrowserActionFillActionFillTypeTypingIntervalUnion `json:"typing_interval,omitzero"`
	// Any of "simple", "distribution".
	TypingStrategy string `json:"typing_strategy,omitzero"`
	paramObj
}

func (r ExtractRunParamsBrowserActionFillActionFillType) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionFillActionFillType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionFillActionFillType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractRunParamsBrowserActionFillActionFillType](
		"mode", "type",
	)
	apijson.RegisterFieldValidator[ExtractRunParamsBrowserActionFillActionFillType](
		"mouse_movement_strategy", "linear", "ghost-cursor", "windmouse",
	)
	apijson.RegisterFieldValidator[ExtractRunParamsBrowserActionFillActionFillType](
		"typing_strategy", "simple", "distribution",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionFillActionFillTypeSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionFillActionFillTypeSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractRunParamsBrowserActionFillActionFillTypeSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionFillActionFillTypeSelectorUnion) asAny() any {
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
type ExtractRunParamsBrowserActionFillActionFillTypeDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionFillActionFillTypeDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractRunParamsBrowserActionFillActionFillTypeDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionFillActionFillTypeDelayUnion) asAny() any {
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
type ExtractRunParamsBrowserActionFillActionFillTypeRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionFillActionFillTypeRequiredString)
	OfExtractRunsBrowserActionFillActionFillTypeRequiredString param.Opt[ExtractRunParamsBrowserActionFillActionFillTypeRequiredString] `json:",omitzero,inline"`
	OfBool                                                     param.Opt[bool]                                                          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionFillActionFillTypeRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionFillActionFillTypeRequiredString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionFillActionFillTypeRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionFillActionFillTypeRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionFillActionFillTypeRequiredString) {
		return &u.OfExtractRunsBrowserActionFillActionFillTypeRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionFillActionFillTypeRequiredString string

const (
	ExtractRunParamsBrowserActionFillActionFillTypeRequiredStringTrue  ExtractRunParamsBrowserActionFillActionFillTypeRequiredString = "true"
	ExtractRunParamsBrowserActionFillActionFillTypeRequiredStringFalse ExtractRunParamsBrowserActionFillActionFillTypeRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionFillActionFillTypeSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionFillActionFillTypeSkipString)
	OfExtractRunsBrowserActionFillActionFillTypeSkipString param.Opt[ExtractRunParamsBrowserActionFillActionFillTypeSkipString] `json:",omitzero,inline"`
	OfBool                                                 param.Opt[bool]                                                      `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionFillActionFillTypeSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionFillActionFillTypeSkipString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionFillActionFillTypeSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionFillActionFillTypeSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionFillActionFillTypeSkipString) {
		return &u.OfExtractRunsBrowserActionFillActionFillTypeSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionFillActionFillTypeSkipString string

const (
	ExtractRunParamsBrowserActionFillActionFillTypeSkipStringTrue  ExtractRunParamsBrowserActionFillActionFillTypeSkipString = "true"
	ExtractRunParamsBrowserActionFillActionFillTypeSkipStringFalse ExtractRunParamsBrowserActionFillActionFillTypeSkipString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionFillActionFillTypeTypingIntervalUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionFillActionFillTypeTypingIntervalUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractRunParamsBrowserActionFillActionFillTypeTypingIntervalUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionFillActionFillTypeTypingIntervalUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// The properties Mode, Selector, Value are required.
type ExtractRunParamsBrowserActionFillActionFillPaste struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector       ExtractRunParamsBrowserActionFillActionFillPasteSelectorUnion `json:"selector,omitzero,required"`
	Value          string                                                        `json:"value,required"`
	ClickOnElement param.Opt[bool]                                               `json:"click_on_element,omitzero"`
	Scroll         param.Opt[bool]                                               `json:"scroll,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay ExtractRunParamsBrowserActionFillActionFillPasteDelayUnion `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractRunParamsBrowserActionFillActionFillPasteRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractRunParamsBrowserActionFillActionFillPasteSkipUnion `json:"skip,omitzero"`
	// This field can be elided, and will marshal its zero value as "paste".
	Mode constant.Paste `json:"mode,required"`
	paramObj
}

func (r ExtractRunParamsBrowserActionFillActionFillPaste) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionFillActionFillPaste
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionFillActionFillPaste) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionFillActionFillPasteSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionFillActionFillPasteSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractRunParamsBrowserActionFillActionFillPasteSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionFillActionFillPasteSelectorUnion) asAny() any {
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
type ExtractRunParamsBrowserActionFillActionFillPasteDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionFillActionFillPasteDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractRunParamsBrowserActionFillActionFillPasteDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionFillActionFillPasteDelayUnion) asAny() any {
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
type ExtractRunParamsBrowserActionFillActionFillPasteRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionFillActionFillPasteRequiredString)
	OfExtractRunsBrowserActionFillActionFillPasteRequiredString param.Opt[ExtractRunParamsBrowserActionFillActionFillPasteRequiredString] `json:",omitzero,inline"`
	OfBool                                                      param.Opt[bool]                                                           `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionFillActionFillPasteRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionFillActionFillPasteRequiredString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionFillActionFillPasteRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionFillActionFillPasteRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionFillActionFillPasteRequiredString) {
		return &u.OfExtractRunsBrowserActionFillActionFillPasteRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionFillActionFillPasteRequiredString string

const (
	ExtractRunParamsBrowserActionFillActionFillPasteRequiredStringTrue  ExtractRunParamsBrowserActionFillActionFillPasteRequiredString = "true"
	ExtractRunParamsBrowserActionFillActionFillPasteRequiredStringFalse ExtractRunParamsBrowserActionFillActionFillPasteRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionFillActionFillPasteSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionFillActionFillPasteSkipString)
	OfExtractRunsBrowserActionFillActionFillPasteSkipString param.Opt[ExtractRunParamsBrowserActionFillActionFillPasteSkipString] `json:",omitzero,inline"`
	OfBool                                                  param.Opt[bool]                                                       `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionFillActionFillPasteSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionFillActionFillPasteSkipString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionFillActionFillPasteSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionFillActionFillPasteSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionFillActionFillPasteSkipString) {
		return &u.OfExtractRunsBrowserActionFillActionFillPasteSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionFillActionFillPasteSkipString string

const (
	ExtractRunParamsBrowserActionFillActionFillPasteSkipStringTrue  ExtractRunParamsBrowserActionFillActionFillPasteSkipString = "true"
	ExtractRunParamsBrowserActionFillActionFillPasteSkipStringFalse ExtractRunParamsBrowserActionFillActionFillPasteSkipString = "false"
)

// Retrieve browser cookies
//
// The property GetCookies is required.
type ExtractRunParamsBrowserActionGetCookiesAction struct {
	GetCookies ExtractRunParamsBrowserActionGetCookiesActionGetCookiesUnion `json:"get_cookies,omitzero,required"`
	paramObj
}

func (r ExtractRunParamsBrowserActionGetCookiesAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionGetCookiesAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionGetCookiesAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionGetCookiesActionGetCookiesUnion struct {
	OfBool                                                     param.Opt[bool]                                                `json:",omitzero,inline"`
	OfExtractRunsBrowserActionGetCookiesActionGetCookiesObject *ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionGetCookiesActionGetCookiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfExtractRunsBrowserActionGetCookiesActionGetCookiesObject)
}
func (u *ExtractRunParamsBrowserActionGetCookiesActionGetCookiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionGetCookiesActionGetCookiesUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionGetCookiesActionGetCookiesObject) {
		return u.OfExtractRunsBrowserActionGetCookiesActionGetCookiesObject
	}
	return nil
}

type ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObject struct {
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip        ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion `json:"skip,omitzero"`
	ExtraFields map[string]any                                                         `json:"-"`
	paramObj
}

func (r ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObject
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionGetCookiesActionGetCookiesObjectRequiredString)
	OfExtractRunsBrowserActionGetCookiesActionGetCookiesObjectRequiredString param.Opt[ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                   param.Opt[bool]                                                                        `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionGetCookiesActionGetCookiesObjectRequiredString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionGetCookiesActionGetCookiesObjectRequiredString) {
		return &u.OfExtractRunsBrowserActionGetCookiesActionGetCookiesObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredString string

const (
	ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredStringTrue  ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredString = "true"
	ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredStringFalse ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionGetCookiesActionGetCookiesObjectSkipString)
	OfExtractRunsBrowserActionGetCookiesActionGetCookiesObjectSkipString param.Opt[ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                               param.Opt[bool]                                                                    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionGetCookiesActionGetCookiesObjectSkipString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionGetCookiesActionGetCookiesObjectSkipString) {
		return &u.OfExtractRunsBrowserActionGetCookiesActionGetCookiesObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectSkipString string

const (
	ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectSkipStringTrue  ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectSkipString = "true"
	ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectSkipStringFalse ExtractRunParamsBrowserActionGetCookiesActionGetCookiesObjectSkipString = "false"
)

// Navigate to a URL
//
// The property Goto is required.
type ExtractRunParamsBrowserActionGotoAction struct {
	Goto ExtractRunParamsBrowserActionGotoActionGotoUnion `json:"goto,omitzero,required" format:"uri"`
	paramObj
}

func (r ExtractRunParamsBrowserActionGotoAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionGotoAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionGotoAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionGotoActionGotoUnion struct {
	OfString                                       param.Opt[string]                                  `json:",omitzero,inline"`
	OfExtractRunsBrowserActionGotoActionGotoObject *ExtractRunParamsBrowserActionGotoActionGotoObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionGotoActionGotoUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractRunsBrowserActionGotoActionGotoObject)
}
func (u *ExtractRunParamsBrowserActionGotoActionGotoUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionGotoActionGotoUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionGotoActionGotoObject) {
		return u.OfExtractRunsBrowserActionGotoActionGotoObject
	}
	return nil
}

// The property URL is required.
type ExtractRunParamsBrowserActionGotoActionGotoObject struct {
	URL     string            `json:"url,required" format:"uri"`
	Referer param.Opt[string] `json:"referer,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractRunParamsBrowserActionGotoActionGotoObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractRunParamsBrowserActionGotoActionGotoObjectSkipUnion `json:"skip,omitzero"`
	// Any of "load", "domcontentloaded", "networkidle0", "networkidle2".
	WaitUntil string `json:"wait_until,omitzero"`
	paramObj
}

func (r ExtractRunParamsBrowserActionGotoActionGotoObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionGotoActionGotoObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionGotoActionGotoObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractRunParamsBrowserActionGotoActionGotoObject](
		"wait_until", "load", "domcontentloaded", "networkidle0", "networkidle2",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionGotoActionGotoObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionGotoActionGotoObjectRequiredString)
	OfExtractRunsBrowserActionGotoActionGotoObjectRequiredString param.Opt[ExtractRunParamsBrowserActionGotoActionGotoObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                       param.Opt[bool]                                                            `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionGotoActionGotoObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionGotoActionGotoObjectRequiredString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionGotoActionGotoObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionGotoActionGotoObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionGotoActionGotoObjectRequiredString) {
		return &u.OfExtractRunsBrowserActionGotoActionGotoObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionGotoActionGotoObjectRequiredString string

const (
	ExtractRunParamsBrowserActionGotoActionGotoObjectRequiredStringTrue  ExtractRunParamsBrowserActionGotoActionGotoObjectRequiredString = "true"
	ExtractRunParamsBrowserActionGotoActionGotoObjectRequiredStringFalse ExtractRunParamsBrowserActionGotoActionGotoObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionGotoActionGotoObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionGotoActionGotoObjectSkipString)
	OfExtractRunsBrowserActionGotoActionGotoObjectSkipString param.Opt[ExtractRunParamsBrowserActionGotoActionGotoObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                   param.Opt[bool]                                                        `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionGotoActionGotoObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionGotoActionGotoObjectSkipString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionGotoActionGotoObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionGotoActionGotoObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionGotoActionGotoObjectSkipString) {
		return &u.OfExtractRunsBrowserActionGotoActionGotoObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionGotoActionGotoObjectSkipString string

const (
	ExtractRunParamsBrowserActionGotoActionGotoObjectSkipStringTrue  ExtractRunParamsBrowserActionGotoActionGotoObjectSkipString = "true"
	ExtractRunParamsBrowserActionGotoActionGotoObjectSkipStringFalse ExtractRunParamsBrowserActionGotoActionGotoObjectSkipString = "false"
)

// Press a keyboard key
//
// The property Press is required.
type ExtractRunParamsBrowserActionPressAction struct {
	Press ExtractRunParamsBrowserActionPressActionPressUnion `json:"press,omitzero,required"`
	paramObj
}

func (r ExtractRunParamsBrowserActionPressAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionPressAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionPressAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionPressActionPressUnion struct {
	OfString                                         param.Opt[string]                                    `json:",omitzero,inline"`
	OfExtractRunsBrowserActionPressActionPressObject *ExtractRunParamsBrowserActionPressActionPressObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionPressActionPressUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractRunsBrowserActionPressActionPressObject)
}
func (u *ExtractRunParamsBrowserActionPressActionPressUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionPressActionPressUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionPressActionPressObject) {
		return u.OfExtractRunsBrowserActionPressActionPressObject
	}
	return nil
}

// The property Key is required.
type ExtractRunParamsBrowserActionPressActionPressObject struct {
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
	Delay ExtractRunParamsBrowserActionPressActionPressObjectDelayUnion `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractRunParamsBrowserActionPressActionPressObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractRunParamsBrowserActionPressActionPressObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractRunParamsBrowserActionPressActionPressObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionPressActionPressObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionPressActionPressObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractRunParamsBrowserActionPressActionPressObject](
		"key", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "Power", "Eject", "Abort", "Help", "Backspace", "Tab", "Numpad5", "NumpadEnter", "Enter", "\r", "\n", "ShiftLeft", "ShiftRight", "ControlLeft", "ControlRight", "AltLeft", "AltRight", "Pause", "CapsLock", "Escape", "Convert", "NonConvert", "Space", "Numpad9", "PageUp", "Numpad3", "PageDown", "End", "Numpad1", "Home", "Numpad7", "ArrowLeft", "Numpad4", "Numpad8", "ArrowUp", "ArrowRight", "Numpad6", "Numpad2", "ArrowDown", "Select", "Open", "PrintScreen", "Insert", "Numpad0", "Delete", "NumpadDecimal", "Digit0", "Digit1", "Digit2", "Digit3", "Digit4", "Digit5", "Digit6", "Digit7", "Digit8", "Digit9", "KeyA", "KeyB", "KeyC", "KeyD", "KeyE", "KeyF", "KeyG", "KeyH", "KeyI", "KeyJ", "KeyK", "KeyL", "KeyM", "KeyN", "KeyO", "KeyP", "KeyQ", "KeyR", "KeyS", "KeyT", "KeyU", "KeyV", "KeyW", "KeyX", "KeyY", "KeyZ", "MetaLeft", "MetaRight", "ContextMenu", "NumpadMultiply", "NumpadAdd", "NumpadSubtract", "NumpadDivide", "F1", "F2", "F3", "F4", "F5", "F6", "F7", "F8", "F9", "F10", "F11", "F12", "F13", "F14", "F15", "F16", "F17", "F18", "F19", "F20", "F21", "F22", "F23", "F24", "NumLock", "ScrollLock", "AudioVolumeMute", "AudioVolumeDown", "AudioVolumeUp", "MediaTrackNext", "MediaTrackPrevious", "MediaStop", "MediaPlayPause", "Semicolon", "Equal", "NumpadEqual", "Comma", "Minus", "Period", "Slash", "Backquote", "BracketLeft", "Backslash", "BracketRight", "Quote", "AltGraph", "Props", "Cancel", "Clear", "Shift", "Control", "Alt", "Accept", "ModeChange", " ", "Print", "Execute", "\u0000", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "Meta", "*", "+", "-", "/", ";", "=", ",", ".", "`", "[", "\\", "]", "'", "Attn", "CrSel", "ExSel", "EraseEof", "Play", "ZoomOut", ")", "!", "@", "#", "$", "%", "^", "&", "(", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", ":", "<", "_", ">", "?", "~", "{", "|", "}", "\"", "SoftLeft", "SoftRight", "Camera", "Call", "EndCall", "VolumeDown", "VolumeUp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionPressActionPressObjectDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionPressActionPressObjectDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractRunParamsBrowserActionPressActionPressObjectDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionPressActionPressObjectDelayUnion) asAny() any {
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
type ExtractRunParamsBrowserActionPressActionPressObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionPressActionPressObjectRequiredString)
	OfExtractRunsBrowserActionPressActionPressObjectRequiredString param.Opt[ExtractRunParamsBrowserActionPressActionPressObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                         param.Opt[bool]                                                              `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionPressActionPressObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionPressActionPressObjectRequiredString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionPressActionPressObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionPressActionPressObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionPressActionPressObjectRequiredString) {
		return &u.OfExtractRunsBrowserActionPressActionPressObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionPressActionPressObjectRequiredString string

const (
	ExtractRunParamsBrowserActionPressActionPressObjectRequiredStringTrue  ExtractRunParamsBrowserActionPressActionPressObjectRequiredString = "true"
	ExtractRunParamsBrowserActionPressActionPressObjectRequiredStringFalse ExtractRunParamsBrowserActionPressActionPressObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionPressActionPressObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionPressActionPressObjectSkipString)
	OfExtractRunsBrowserActionPressActionPressObjectSkipString param.Opt[ExtractRunParamsBrowserActionPressActionPressObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                     param.Opt[bool]                                                          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionPressActionPressObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionPressActionPressObjectSkipString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionPressActionPressObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionPressActionPressObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionPressActionPressObjectSkipString) {
		return &u.OfExtractRunsBrowserActionPressActionPressObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionPressActionPressObjectSkipString string

const (
	ExtractRunParamsBrowserActionPressActionPressObjectSkipStringTrue  ExtractRunParamsBrowserActionPressActionPressObjectSkipString = "true"
	ExtractRunParamsBrowserActionPressActionPressObjectSkipStringFalse ExtractRunParamsBrowserActionPressActionPressObjectSkipString = "false"
)

// Capture a page screenshot
//
// The property Screenshot is required.
type ExtractRunParamsBrowserActionScreenshotAction struct {
	Screenshot ExtractRunParamsBrowserActionScreenshotActionScreenshotUnion `json:"screenshot,omitzero,required"`
	paramObj
}

func (r ExtractRunParamsBrowserActionScreenshotAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionScreenshotAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionScreenshotAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionScreenshotActionScreenshotUnion struct {
	OfBool                                                     param.Opt[bool]                                                `json:",omitzero,inline"`
	OfExtractRunsBrowserActionScreenshotActionScreenshotObject *ExtractRunParamsBrowserActionScreenshotActionScreenshotObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionScreenshotActionScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfExtractRunsBrowserActionScreenshotActionScreenshotObject)
}
func (u *ExtractRunParamsBrowserActionScreenshotActionScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionScreenshotActionScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionScreenshotActionScreenshotObject) {
		return u.OfExtractRunsBrowserActionScreenshotActionScreenshotObject
	}
	return nil
}

type ExtractRunParamsBrowserActionScreenshotActionScreenshotObject struct {
	FullPage param.Opt[bool]    `json:"full_page,omitzero"`
	Quality  param.Opt[float64] `json:"quality,omitzero"`
	// Any of "png", "jpeg", "webp".
	Format string `json:"format,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractRunParamsBrowserActionScreenshotActionScreenshotObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionScreenshotActionScreenshotObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionScreenshotActionScreenshotObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractRunParamsBrowserActionScreenshotActionScreenshotObject](
		"format", "png", "jpeg", "webp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionScreenshotActionScreenshotObjectRequiredString)
	OfExtractRunsBrowserActionScreenshotActionScreenshotObjectRequiredString param.Opt[ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                   param.Opt[bool]                                                                        `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionScreenshotActionScreenshotObjectRequiredString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionScreenshotActionScreenshotObjectRequiredString) {
		return &u.OfExtractRunsBrowserActionScreenshotActionScreenshotObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectRequiredString string

const (
	ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectRequiredStringTrue  ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectRequiredString = "true"
	ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectRequiredStringFalse ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionScreenshotActionScreenshotObjectSkipString)
	OfExtractRunsBrowserActionScreenshotActionScreenshotObjectSkipString param.Opt[ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                               param.Opt[bool]                                                                    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionScreenshotActionScreenshotObjectSkipString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionScreenshotActionScreenshotObjectSkipString) {
		return &u.OfExtractRunsBrowserActionScreenshotActionScreenshotObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectSkipString string

const (
	ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectSkipStringTrue  ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectSkipString = "true"
	ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectSkipStringFalse ExtractRunParamsBrowserActionScreenshotActionScreenshotObjectSkipString = "false"
)

// Scroll the page or an element
//
// The property Scroll is required.
type ExtractRunParamsBrowserActionScrollAction struct {
	Scroll ExtractRunParamsBrowserActionScrollActionScrollUnion `json:"scroll,omitzero,required"`
	paramObj
}

func (r ExtractRunParamsBrowserActionScrollAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionScrollAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionScrollAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionScrollActionScrollUnion struct {
	OfFloat                                            param.Opt[float64]                                     `json:",omitzero,inline"`
	OfString                                           param.Opt[string]                                      `json:",omitzero,inline"`
	OfExtractRunsBrowserActionScrollActionScrollObject *ExtractRunParamsBrowserActionScrollActionScrollObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionScrollActionScrollUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfExtractRunsBrowserActionScrollActionScrollObject)
}
func (u *ExtractRunParamsBrowserActionScrollActionScrollUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionScrollActionScrollUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionScrollActionScrollObject) {
		return u.OfExtractRunsBrowserActionScrollActionScrollObject
	}
	return nil
}

type ExtractRunParamsBrowserActionScrollActionScrollObject struct {
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	X       param.Opt[float64] `json:"x,omitzero"`
	Y       param.Opt[float64] `json:"y,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Container ExtractRunParamsBrowserActionScrollActionScrollObjectContainerUnion `json:"container,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractRunParamsBrowserActionScrollActionScrollObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractRunParamsBrowserActionScrollActionScrollObjectSkipUnion `json:"skip,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	To ExtractRunParamsBrowserActionScrollActionScrollObjectToUnion `json:"to,omitzero"`
	paramObj
}

func (r ExtractRunParamsBrowserActionScrollActionScrollObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionScrollActionScrollObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionScrollActionScrollObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionScrollActionScrollObjectContainerUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionScrollActionScrollObjectContainerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractRunParamsBrowserActionScrollActionScrollObjectContainerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionScrollActionScrollObjectContainerUnion) asAny() any {
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
type ExtractRunParamsBrowserActionScrollActionScrollObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionScrollActionScrollObjectRequiredString)
	OfExtractRunsBrowserActionScrollActionScrollObjectRequiredString param.Opt[ExtractRunParamsBrowserActionScrollActionScrollObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                           param.Opt[bool]                                                                `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionScrollActionScrollObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionScrollActionScrollObjectRequiredString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionScrollActionScrollObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionScrollActionScrollObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionScrollActionScrollObjectRequiredString) {
		return &u.OfExtractRunsBrowserActionScrollActionScrollObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionScrollActionScrollObjectRequiredString string

const (
	ExtractRunParamsBrowserActionScrollActionScrollObjectRequiredStringTrue  ExtractRunParamsBrowserActionScrollActionScrollObjectRequiredString = "true"
	ExtractRunParamsBrowserActionScrollActionScrollObjectRequiredStringFalse ExtractRunParamsBrowserActionScrollActionScrollObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionScrollActionScrollObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionScrollActionScrollObjectSkipString)
	OfExtractRunsBrowserActionScrollActionScrollObjectSkipString param.Opt[ExtractRunParamsBrowserActionScrollActionScrollObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                       param.Opt[bool]                                                            `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionScrollActionScrollObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionScrollActionScrollObjectSkipString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionScrollActionScrollObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionScrollActionScrollObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionScrollActionScrollObjectSkipString) {
		return &u.OfExtractRunsBrowserActionScrollActionScrollObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionScrollActionScrollObjectSkipString string

const (
	ExtractRunParamsBrowserActionScrollActionScrollObjectSkipStringTrue  ExtractRunParamsBrowserActionScrollActionScrollObjectSkipString = "true"
	ExtractRunParamsBrowserActionScrollActionScrollObjectSkipStringFalse ExtractRunParamsBrowserActionScrollActionScrollObjectSkipString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionScrollActionScrollObjectToUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionScrollActionScrollObjectToUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractRunParamsBrowserActionScrollActionScrollObjectToUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionScrollActionScrollObjectToUnion) asAny() any {
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
type ExtractRunParamsBrowserActionWaitAction struct {
	Wait ExtractRunParamsBrowserActionWaitActionWaitUnion `json:"wait,omitzero,required"`
	paramObj
}

func (r ExtractRunParamsBrowserActionWaitAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionWaitAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionWaitAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionWaitActionWaitUnion struct {
	OfFloat                                        param.Opt[float64]                                 `json:",omitzero,inline"`
	OfString                                       param.Opt[string]                                  `json:",omitzero,inline"`
	OfExtractRunsBrowserActionWaitActionWaitObject *ExtractRunParamsBrowserActionWaitActionWaitObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionWaitActionWaitUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfExtractRunsBrowserActionWaitActionWaitObject)
}
func (u *ExtractRunParamsBrowserActionWaitActionWaitUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionWaitActionWaitUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionWaitActionWaitObject) {
		return u.OfExtractRunsBrowserActionWaitActionWaitObject
	}
	return nil
}

// The property Duration is required.
type ExtractRunParamsBrowserActionWaitActionWaitObject struct {
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Duration ExtractRunParamsBrowserActionWaitActionWaitObjectDurationUnion `json:"duration,omitzero,required"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractRunParamsBrowserActionWaitActionWaitObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractRunParamsBrowserActionWaitActionWaitObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractRunParamsBrowserActionWaitActionWaitObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionWaitActionWaitObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionWaitActionWaitObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionWaitActionWaitObjectDurationUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionWaitActionWaitObjectDurationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractRunParamsBrowserActionWaitActionWaitObjectDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionWaitActionWaitObjectDurationUnion) asAny() any {
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
type ExtractRunParamsBrowserActionWaitActionWaitObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionWaitActionWaitObjectRequiredString)
	OfExtractRunsBrowserActionWaitActionWaitObjectRequiredString param.Opt[ExtractRunParamsBrowserActionWaitActionWaitObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                       param.Opt[bool]                                                            `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionWaitActionWaitObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionWaitActionWaitObjectRequiredString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionWaitActionWaitObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionWaitActionWaitObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionWaitActionWaitObjectRequiredString) {
		return &u.OfExtractRunsBrowserActionWaitActionWaitObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionWaitActionWaitObjectRequiredString string

const (
	ExtractRunParamsBrowserActionWaitActionWaitObjectRequiredStringTrue  ExtractRunParamsBrowserActionWaitActionWaitObjectRequiredString = "true"
	ExtractRunParamsBrowserActionWaitActionWaitObjectRequiredStringFalse ExtractRunParamsBrowserActionWaitActionWaitObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionWaitActionWaitObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionWaitActionWaitObjectSkipString)
	OfExtractRunsBrowserActionWaitActionWaitObjectSkipString param.Opt[ExtractRunParamsBrowserActionWaitActionWaitObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                   param.Opt[bool]                                                        `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionWaitActionWaitObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionWaitActionWaitObjectSkipString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionWaitActionWaitObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionWaitActionWaitObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionWaitActionWaitObjectSkipString) {
		return &u.OfExtractRunsBrowserActionWaitActionWaitObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionWaitActionWaitObjectSkipString string

const (
	ExtractRunParamsBrowserActionWaitActionWaitObjectSkipStringTrue  ExtractRunParamsBrowserActionWaitActionWaitObjectSkipString = "true"
	ExtractRunParamsBrowserActionWaitActionWaitObjectSkipStringFalse ExtractRunParamsBrowserActionWaitActionWaitObjectSkipString = "false"
)

// Wait for an element to appear or reach a specific state
//
// The property WaitForElement is required.
type ExtractRunParamsBrowserActionWaitForElementAction struct {
	WaitForElement ExtractRunParamsBrowserActionWaitForElementActionWaitForElementUnion `json:"wait_for_element,omitzero,required"`
	paramObj
}

func (r ExtractRunParamsBrowserActionWaitForElementAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionWaitForElementAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionWaitForElementAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionWaitForElementActionWaitForElementUnion struct {
	OfString                                                           param.Opt[string]                                                      `json:",omitzero,inline"`
	OfStringArray                                                      []string                                                               `json:",omitzero,inline"`
	OfExtractRunsBrowserActionWaitForElementActionWaitForElementObject *ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionWaitForElementActionWaitForElementUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray, u.OfExtractRunsBrowserActionWaitForElementActionWaitForElementObject)
}
func (u *ExtractRunParamsBrowserActionWaitForElementActionWaitForElementUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionWaitForElementActionWaitForElementUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionWaitForElementActionWaitForElementObject) {
		return u.OfExtractRunsBrowserActionWaitForElementActionWaitForElementObject
	}
	return nil
}

// The property Selector is required.
type ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObject struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion `json:"selector,omitzero,required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion) asAny() any {
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
type ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionWaitForElementActionWaitForElementObjectRequiredString)
	OfExtractRunsBrowserActionWaitForElementActionWaitForElementObjectRequiredString param.Opt[ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                           param.Opt[bool]                                                                                `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionWaitForElementActionWaitForElementObjectRequiredString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionWaitForElementActionWaitForElementObjectRequiredString) {
		return &u.OfExtractRunsBrowserActionWaitForElementActionWaitForElementObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredString string

const (
	ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredStringTrue  ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredString = "true"
	ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredStringFalse ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionWaitForElementActionWaitForElementObjectSkipString)
	OfExtractRunsBrowserActionWaitForElementActionWaitForElementObjectSkipString param.Opt[ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                       param.Opt[bool]                                                                            `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionWaitForElementActionWaitForElementObjectSkipString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionWaitForElementActionWaitForElementObjectSkipString) {
		return &u.OfExtractRunsBrowserActionWaitForElementActionWaitForElementObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectSkipString string

const (
	ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectSkipStringTrue  ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectSkipString = "true"
	ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectSkipStringFalse ExtractRunParamsBrowserActionWaitForElementActionWaitForElementObjectSkipString = "false"
)

// Wait for page navigation to complete
//
// The property WaitForNavigation is required.
type ExtractRunParamsBrowserActionWaitForNavigationAction struct {
	WaitForNavigation ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion `json:"wait_for_navigation,omitzero,required"`
	paramObj
}

func (r ExtractRunParamsBrowserActionWaitForNavigationAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionWaitForNavigationAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionWaitForNavigationAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationString)
	OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationString param.Opt[string]                                                            `json:",omitzero,inline"`
	OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationObject *ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationString, u.OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationObject)
}
func (u *ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationString) {
		return &u.OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationString
	} else if !param.IsOmitted(u.OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationObject) {
		return u.OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationObject
	}
	return nil
}

type ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationString string

const (
	ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationStringLoad             ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationString = "load"
	ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationStringDomcontentloaded ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationString = "domcontentloaded"
	ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationStringNetworkidle0     ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationString = "networkidle0"
	ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationStringNetworkidle2     ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationString = "networkidle2"
)

// The property Navigation is required.
type ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObject struct {
	// Any of "load", "domcontentloaded", "networkidle0", "networkidle2".
	Navigation string `json:"navigation,omitzero,required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObject](
		"navigation", "load", "domcontentloaded", "networkidle0", "networkidle2",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString)
	OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString param.Opt[ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                                 param.Opt[bool]                                                                                      `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString) {
		return &u.OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString string

const (
	ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredStringTrue  ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString = "true"
	ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredStringFalse ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString)
	OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString param.Opt[ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                             param.Opt[bool]                                                                                  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString, u.OfBool)
}
func (u *ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString) {
		return &u.OfExtractRunsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString string

const (
	ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipStringTrue  ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString = "true"
	ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipStringFalse ExtractRunParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsCookiesUnion struct {
	OfExtractRunsCookiesArray []ExtractRunParamsCookiesArrayItem `json:",omitzero,inline"`
	OfString                  param.Opt[string]                  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsCookiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsCookiesArray, u.OfString)
}
func (u *ExtractRunParamsCookiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsCookiesUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsCookiesArray) {
		return &u.OfExtractRunsCookiesArray
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type ExtractRunParamsCookiesArrayItem struct {
	Creation      param.Opt[string]                           `json:"creation,omitzero"`
	Domain        param.Opt[string]                           `json:"domain,omitzero"`
	HostOnly      param.Opt[bool]                             `json:"hostOnly,omitzero"`
	HTTPOnly      param.Opt[bool]                             `json:"httpOnly,omitzero"`
	LastAccessed  param.Opt[string]                           `json:"lastAccessed,omitzero"`
	Path          param.Opt[string]                           `json:"path,omitzero"`
	PathIsDefault param.Opt[bool]                             `json:"pathIsDefault,omitzero"`
	Expires       param.Opt[string]                           `json:"expires,omitzero"`
	Name          param.Opt[string]                           `json:"name,omitzero"`
	Secure        param.Opt[bool]                             `json:"secure,omitzero"`
	Value         param.Opt[string]                           `json:"value,omitzero"`
	Extensions    []string                                    `json:"extensions,omitzero"`
	MaxAge        ExtractRunParamsCookiesArrayItemMaxAgeUnion `json:"maxAge,omitzero"`
	// Any of "strict", "lax", "none".
	SameSite    string         `json:"sameSite,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r ExtractRunParamsCookiesArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsCookiesArrayItem
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ExtractRunParamsCookiesArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractRunParamsCookiesArrayItem](
		"sameSite", "strict", "lax", "none",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsCookiesArrayItemMaxAgeUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractRunsCookiesArrayItemMaxAgeString)
	OfExtractRunsCookiesArrayItemMaxAgeString param.Opt[ExtractRunParamsCookiesArrayItemMaxAgeString] `json:",omitzero,inline"`
	OfFloat                                   param.Opt[float64]                                      `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsCookiesArrayItemMaxAgeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractRunsCookiesArrayItemMaxAgeString, u.OfFloat)
}
func (u *ExtractRunParamsCookiesArrayItemMaxAgeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsCookiesArrayItemMaxAgeUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractRunsCookiesArrayItemMaxAgeString) {
		return &u.OfExtractRunsCookiesArrayItemMaxAgeString
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type ExtractRunParamsCookiesArrayItemMaxAgeString string

const (
	ExtractRunParamsCookiesArrayItemMaxAgeStringInfinity      ExtractRunParamsCookiesArrayItemMaxAgeString = "Infinity"
	ExtractRunParamsCookiesArrayItemMaxAgeStringMinusInfinity ExtractRunParamsCookiesArrayItemMaxAgeString = "-Infinity"
)

// Country code for geolocation and proxy selection
type ExtractRunParamsCountry string

const (
	ExtractRunParamsCountryAd  ExtractRunParamsCountry = "AD"
	ExtractRunParamsCountryAe  ExtractRunParamsCountry = "AE"
	ExtractRunParamsCountryAf  ExtractRunParamsCountry = "AF"
	ExtractRunParamsCountryAg  ExtractRunParamsCountry = "AG"
	ExtractRunParamsCountryAI  ExtractRunParamsCountry = "AI"
	ExtractRunParamsCountryAl  ExtractRunParamsCountry = "AL"
	ExtractRunParamsCountryAm  ExtractRunParamsCountry = "AM"
	ExtractRunParamsCountryAo  ExtractRunParamsCountry = "AO"
	ExtractRunParamsCountryAq  ExtractRunParamsCountry = "AQ"
	ExtractRunParamsCountryAr  ExtractRunParamsCountry = "AR"
	ExtractRunParamsCountryAs  ExtractRunParamsCountry = "AS"
	ExtractRunParamsCountryAt  ExtractRunParamsCountry = "AT"
	ExtractRunParamsCountryAu  ExtractRunParamsCountry = "AU"
	ExtractRunParamsCountryAw  ExtractRunParamsCountry = "AW"
	ExtractRunParamsCountryAx  ExtractRunParamsCountry = "AX"
	ExtractRunParamsCountryAz  ExtractRunParamsCountry = "AZ"
	ExtractRunParamsCountryBa  ExtractRunParamsCountry = "BA"
	ExtractRunParamsCountryBb  ExtractRunParamsCountry = "BB"
	ExtractRunParamsCountryBd  ExtractRunParamsCountry = "BD"
	ExtractRunParamsCountryBe  ExtractRunParamsCountry = "BE"
	ExtractRunParamsCountryBf  ExtractRunParamsCountry = "BF"
	ExtractRunParamsCountryBg  ExtractRunParamsCountry = "BG"
	ExtractRunParamsCountryBh  ExtractRunParamsCountry = "BH"
	ExtractRunParamsCountryBi  ExtractRunParamsCountry = "BI"
	ExtractRunParamsCountryBj  ExtractRunParamsCountry = "BJ"
	ExtractRunParamsCountryBl  ExtractRunParamsCountry = "BL"
	ExtractRunParamsCountryBm  ExtractRunParamsCountry = "BM"
	ExtractRunParamsCountryBn  ExtractRunParamsCountry = "BN"
	ExtractRunParamsCountryBo  ExtractRunParamsCountry = "BO"
	ExtractRunParamsCountryBq  ExtractRunParamsCountry = "BQ"
	ExtractRunParamsCountryBr  ExtractRunParamsCountry = "BR"
	ExtractRunParamsCountryBs  ExtractRunParamsCountry = "BS"
	ExtractRunParamsCountryBt  ExtractRunParamsCountry = "BT"
	ExtractRunParamsCountryBv  ExtractRunParamsCountry = "BV"
	ExtractRunParamsCountryBw  ExtractRunParamsCountry = "BW"
	ExtractRunParamsCountryBy  ExtractRunParamsCountry = "BY"
	ExtractRunParamsCountryBz  ExtractRunParamsCountry = "BZ"
	ExtractRunParamsCountryCa  ExtractRunParamsCountry = "CA"
	ExtractRunParamsCountryCc  ExtractRunParamsCountry = "CC"
	ExtractRunParamsCountryCd  ExtractRunParamsCountry = "CD"
	ExtractRunParamsCountryCf  ExtractRunParamsCountry = "CF"
	ExtractRunParamsCountryCg  ExtractRunParamsCountry = "CG"
	ExtractRunParamsCountryCh  ExtractRunParamsCountry = "CH"
	ExtractRunParamsCountryCi  ExtractRunParamsCountry = "CI"
	ExtractRunParamsCountryCk  ExtractRunParamsCountry = "CK"
	ExtractRunParamsCountryCl  ExtractRunParamsCountry = "CL"
	ExtractRunParamsCountryCm  ExtractRunParamsCountry = "CM"
	ExtractRunParamsCountryCn  ExtractRunParamsCountry = "CN"
	ExtractRunParamsCountryCo  ExtractRunParamsCountry = "CO"
	ExtractRunParamsCountryCr  ExtractRunParamsCountry = "CR"
	ExtractRunParamsCountryCu  ExtractRunParamsCountry = "CU"
	ExtractRunParamsCountryCv  ExtractRunParamsCountry = "CV"
	ExtractRunParamsCountryCw  ExtractRunParamsCountry = "CW"
	ExtractRunParamsCountryCx  ExtractRunParamsCountry = "CX"
	ExtractRunParamsCountryCy  ExtractRunParamsCountry = "CY"
	ExtractRunParamsCountryCz  ExtractRunParamsCountry = "CZ"
	ExtractRunParamsCountryDe  ExtractRunParamsCountry = "DE"
	ExtractRunParamsCountryDj  ExtractRunParamsCountry = "DJ"
	ExtractRunParamsCountryDk  ExtractRunParamsCountry = "DK"
	ExtractRunParamsCountryDm  ExtractRunParamsCountry = "DM"
	ExtractRunParamsCountryDo  ExtractRunParamsCountry = "DO"
	ExtractRunParamsCountryDz  ExtractRunParamsCountry = "DZ"
	ExtractRunParamsCountryEc  ExtractRunParamsCountry = "EC"
	ExtractRunParamsCountryEe  ExtractRunParamsCountry = "EE"
	ExtractRunParamsCountryEg  ExtractRunParamsCountry = "EG"
	ExtractRunParamsCountryEh  ExtractRunParamsCountry = "EH"
	ExtractRunParamsCountryEr  ExtractRunParamsCountry = "ER"
	ExtractRunParamsCountryEs  ExtractRunParamsCountry = "ES"
	ExtractRunParamsCountryEt  ExtractRunParamsCountry = "ET"
	ExtractRunParamsCountryFi  ExtractRunParamsCountry = "FI"
	ExtractRunParamsCountryFj  ExtractRunParamsCountry = "FJ"
	ExtractRunParamsCountryFk  ExtractRunParamsCountry = "FK"
	ExtractRunParamsCountryFm  ExtractRunParamsCountry = "FM"
	ExtractRunParamsCountryFo  ExtractRunParamsCountry = "FO"
	ExtractRunParamsCountryFr  ExtractRunParamsCountry = "FR"
	ExtractRunParamsCountryGa  ExtractRunParamsCountry = "GA"
	ExtractRunParamsCountryGB  ExtractRunParamsCountry = "GB"
	ExtractRunParamsCountryGd  ExtractRunParamsCountry = "GD"
	ExtractRunParamsCountryGe  ExtractRunParamsCountry = "GE"
	ExtractRunParamsCountryGf  ExtractRunParamsCountry = "GF"
	ExtractRunParamsCountryGg  ExtractRunParamsCountry = "GG"
	ExtractRunParamsCountryGh  ExtractRunParamsCountry = "GH"
	ExtractRunParamsCountryGi  ExtractRunParamsCountry = "GI"
	ExtractRunParamsCountryGl  ExtractRunParamsCountry = "GL"
	ExtractRunParamsCountryGm  ExtractRunParamsCountry = "GM"
	ExtractRunParamsCountryGn  ExtractRunParamsCountry = "GN"
	ExtractRunParamsCountryGp  ExtractRunParamsCountry = "GP"
	ExtractRunParamsCountryGq  ExtractRunParamsCountry = "GQ"
	ExtractRunParamsCountryGr  ExtractRunParamsCountry = "GR"
	ExtractRunParamsCountryGs  ExtractRunParamsCountry = "GS"
	ExtractRunParamsCountryGt  ExtractRunParamsCountry = "GT"
	ExtractRunParamsCountryGu  ExtractRunParamsCountry = "GU"
	ExtractRunParamsCountryGw  ExtractRunParamsCountry = "GW"
	ExtractRunParamsCountryGy  ExtractRunParamsCountry = "GY"
	ExtractRunParamsCountryHk  ExtractRunParamsCountry = "HK"
	ExtractRunParamsCountryHm  ExtractRunParamsCountry = "HM"
	ExtractRunParamsCountryHn  ExtractRunParamsCountry = "HN"
	ExtractRunParamsCountryHr  ExtractRunParamsCountry = "HR"
	ExtractRunParamsCountryHt  ExtractRunParamsCountry = "HT"
	ExtractRunParamsCountryHu  ExtractRunParamsCountry = "HU"
	ExtractRunParamsCountryID  ExtractRunParamsCountry = "ID"
	ExtractRunParamsCountryIe  ExtractRunParamsCountry = "IE"
	ExtractRunParamsCountryIl  ExtractRunParamsCountry = "IL"
	ExtractRunParamsCountryIm  ExtractRunParamsCountry = "IM"
	ExtractRunParamsCountryIn  ExtractRunParamsCountry = "IN"
	ExtractRunParamsCountryIo  ExtractRunParamsCountry = "IO"
	ExtractRunParamsCountryIq  ExtractRunParamsCountry = "IQ"
	ExtractRunParamsCountryIr  ExtractRunParamsCountry = "IR"
	ExtractRunParamsCountryIs  ExtractRunParamsCountry = "IS"
	ExtractRunParamsCountryIt  ExtractRunParamsCountry = "IT"
	ExtractRunParamsCountryJe  ExtractRunParamsCountry = "JE"
	ExtractRunParamsCountryJm  ExtractRunParamsCountry = "JM"
	ExtractRunParamsCountryJo  ExtractRunParamsCountry = "JO"
	ExtractRunParamsCountryJp  ExtractRunParamsCountry = "JP"
	ExtractRunParamsCountryKe  ExtractRunParamsCountry = "KE"
	ExtractRunParamsCountryKg  ExtractRunParamsCountry = "KG"
	ExtractRunParamsCountryKh  ExtractRunParamsCountry = "KH"
	ExtractRunParamsCountryKi  ExtractRunParamsCountry = "KI"
	ExtractRunParamsCountryKm  ExtractRunParamsCountry = "KM"
	ExtractRunParamsCountryKn  ExtractRunParamsCountry = "KN"
	ExtractRunParamsCountryKp  ExtractRunParamsCountry = "KP"
	ExtractRunParamsCountryKr  ExtractRunParamsCountry = "KR"
	ExtractRunParamsCountryKw  ExtractRunParamsCountry = "KW"
	ExtractRunParamsCountryKy  ExtractRunParamsCountry = "KY"
	ExtractRunParamsCountryKz  ExtractRunParamsCountry = "KZ"
	ExtractRunParamsCountryLa  ExtractRunParamsCountry = "LA"
	ExtractRunParamsCountryLb  ExtractRunParamsCountry = "LB"
	ExtractRunParamsCountryLc  ExtractRunParamsCountry = "LC"
	ExtractRunParamsCountryLi  ExtractRunParamsCountry = "LI"
	ExtractRunParamsCountryLk  ExtractRunParamsCountry = "LK"
	ExtractRunParamsCountryLr  ExtractRunParamsCountry = "LR"
	ExtractRunParamsCountryLs  ExtractRunParamsCountry = "LS"
	ExtractRunParamsCountryLt  ExtractRunParamsCountry = "LT"
	ExtractRunParamsCountryLu  ExtractRunParamsCountry = "LU"
	ExtractRunParamsCountryLv  ExtractRunParamsCountry = "LV"
	ExtractRunParamsCountryLy  ExtractRunParamsCountry = "LY"
	ExtractRunParamsCountryMa  ExtractRunParamsCountry = "MA"
	ExtractRunParamsCountryMc  ExtractRunParamsCountry = "MC"
	ExtractRunParamsCountryMd  ExtractRunParamsCountry = "MD"
	ExtractRunParamsCountryMe  ExtractRunParamsCountry = "ME"
	ExtractRunParamsCountryMf  ExtractRunParamsCountry = "MF"
	ExtractRunParamsCountryMg  ExtractRunParamsCountry = "MG"
	ExtractRunParamsCountryMh  ExtractRunParamsCountry = "MH"
	ExtractRunParamsCountryMk  ExtractRunParamsCountry = "MK"
	ExtractRunParamsCountryMl  ExtractRunParamsCountry = "ML"
	ExtractRunParamsCountryMm  ExtractRunParamsCountry = "MM"
	ExtractRunParamsCountryMn  ExtractRunParamsCountry = "MN"
	ExtractRunParamsCountryMo  ExtractRunParamsCountry = "MO"
	ExtractRunParamsCountryMp  ExtractRunParamsCountry = "MP"
	ExtractRunParamsCountryMq  ExtractRunParamsCountry = "MQ"
	ExtractRunParamsCountryMr  ExtractRunParamsCountry = "MR"
	ExtractRunParamsCountryMs  ExtractRunParamsCountry = "MS"
	ExtractRunParamsCountryMt  ExtractRunParamsCountry = "MT"
	ExtractRunParamsCountryMu  ExtractRunParamsCountry = "MU"
	ExtractRunParamsCountryMv  ExtractRunParamsCountry = "MV"
	ExtractRunParamsCountryMw  ExtractRunParamsCountry = "MW"
	ExtractRunParamsCountryMx  ExtractRunParamsCountry = "MX"
	ExtractRunParamsCountryMy  ExtractRunParamsCountry = "MY"
	ExtractRunParamsCountryMz  ExtractRunParamsCountry = "MZ"
	ExtractRunParamsCountryNa  ExtractRunParamsCountry = "NA"
	ExtractRunParamsCountryNc  ExtractRunParamsCountry = "NC"
	ExtractRunParamsCountryNe  ExtractRunParamsCountry = "NE"
	ExtractRunParamsCountryNf  ExtractRunParamsCountry = "NF"
	ExtractRunParamsCountryNg  ExtractRunParamsCountry = "NG"
	ExtractRunParamsCountryNi  ExtractRunParamsCountry = "NI"
	ExtractRunParamsCountryNl  ExtractRunParamsCountry = "NL"
	ExtractRunParamsCountryNo  ExtractRunParamsCountry = "NO"
	ExtractRunParamsCountryNp  ExtractRunParamsCountry = "NP"
	ExtractRunParamsCountryNr  ExtractRunParamsCountry = "NR"
	ExtractRunParamsCountryNu  ExtractRunParamsCountry = "NU"
	ExtractRunParamsCountryNz  ExtractRunParamsCountry = "NZ"
	ExtractRunParamsCountryOm  ExtractRunParamsCountry = "OM"
	ExtractRunParamsCountryPa  ExtractRunParamsCountry = "PA"
	ExtractRunParamsCountryPe  ExtractRunParamsCountry = "PE"
	ExtractRunParamsCountryPf  ExtractRunParamsCountry = "PF"
	ExtractRunParamsCountryPg  ExtractRunParamsCountry = "PG"
	ExtractRunParamsCountryPh  ExtractRunParamsCountry = "PH"
	ExtractRunParamsCountryPk  ExtractRunParamsCountry = "PK"
	ExtractRunParamsCountryPl  ExtractRunParamsCountry = "PL"
	ExtractRunParamsCountryPm  ExtractRunParamsCountry = "PM"
	ExtractRunParamsCountryPn  ExtractRunParamsCountry = "PN"
	ExtractRunParamsCountryPr  ExtractRunParamsCountry = "PR"
	ExtractRunParamsCountryPs  ExtractRunParamsCountry = "PS"
	ExtractRunParamsCountryPt  ExtractRunParamsCountry = "PT"
	ExtractRunParamsCountryPw  ExtractRunParamsCountry = "PW"
	ExtractRunParamsCountryPy  ExtractRunParamsCountry = "PY"
	ExtractRunParamsCountryQa  ExtractRunParamsCountry = "QA"
	ExtractRunParamsCountryRe  ExtractRunParamsCountry = "RE"
	ExtractRunParamsCountryRo  ExtractRunParamsCountry = "RO"
	ExtractRunParamsCountryRs  ExtractRunParamsCountry = "RS"
	ExtractRunParamsCountryRu  ExtractRunParamsCountry = "RU"
	ExtractRunParamsCountryRw  ExtractRunParamsCountry = "RW"
	ExtractRunParamsCountrySa  ExtractRunParamsCountry = "SA"
	ExtractRunParamsCountrySb  ExtractRunParamsCountry = "SB"
	ExtractRunParamsCountrySc  ExtractRunParamsCountry = "SC"
	ExtractRunParamsCountrySd  ExtractRunParamsCountry = "SD"
	ExtractRunParamsCountrySe  ExtractRunParamsCountry = "SE"
	ExtractRunParamsCountrySg  ExtractRunParamsCountry = "SG"
	ExtractRunParamsCountrySh  ExtractRunParamsCountry = "SH"
	ExtractRunParamsCountrySi  ExtractRunParamsCountry = "SI"
	ExtractRunParamsCountrySj  ExtractRunParamsCountry = "SJ"
	ExtractRunParamsCountrySk  ExtractRunParamsCountry = "SK"
	ExtractRunParamsCountrySl  ExtractRunParamsCountry = "SL"
	ExtractRunParamsCountrySm  ExtractRunParamsCountry = "SM"
	ExtractRunParamsCountrySn  ExtractRunParamsCountry = "SN"
	ExtractRunParamsCountrySo  ExtractRunParamsCountry = "SO"
	ExtractRunParamsCountrySr  ExtractRunParamsCountry = "SR"
	ExtractRunParamsCountrySS  ExtractRunParamsCountry = "SS"
	ExtractRunParamsCountrySt  ExtractRunParamsCountry = "ST"
	ExtractRunParamsCountrySv  ExtractRunParamsCountry = "SV"
	ExtractRunParamsCountrySx  ExtractRunParamsCountry = "SX"
	ExtractRunParamsCountrySy  ExtractRunParamsCountry = "SY"
	ExtractRunParamsCountrySz  ExtractRunParamsCountry = "SZ"
	ExtractRunParamsCountryTc  ExtractRunParamsCountry = "TC"
	ExtractRunParamsCountryTd  ExtractRunParamsCountry = "TD"
	ExtractRunParamsCountryTf  ExtractRunParamsCountry = "TF"
	ExtractRunParamsCountryTg  ExtractRunParamsCountry = "TG"
	ExtractRunParamsCountryTh  ExtractRunParamsCountry = "TH"
	ExtractRunParamsCountryTj  ExtractRunParamsCountry = "TJ"
	ExtractRunParamsCountryTk  ExtractRunParamsCountry = "TK"
	ExtractRunParamsCountryTl  ExtractRunParamsCountry = "TL"
	ExtractRunParamsCountryTm  ExtractRunParamsCountry = "TM"
	ExtractRunParamsCountryTn  ExtractRunParamsCountry = "TN"
	ExtractRunParamsCountryTo  ExtractRunParamsCountry = "TO"
	ExtractRunParamsCountryTr  ExtractRunParamsCountry = "TR"
	ExtractRunParamsCountryTt  ExtractRunParamsCountry = "TT"
	ExtractRunParamsCountryTv  ExtractRunParamsCountry = "TV"
	ExtractRunParamsCountryTw  ExtractRunParamsCountry = "TW"
	ExtractRunParamsCountryTz  ExtractRunParamsCountry = "TZ"
	ExtractRunParamsCountryUa  ExtractRunParamsCountry = "UA"
	ExtractRunParamsCountryUg  ExtractRunParamsCountry = "UG"
	ExtractRunParamsCountryUm  ExtractRunParamsCountry = "UM"
	ExtractRunParamsCountryUs  ExtractRunParamsCountry = "US"
	ExtractRunParamsCountryUy  ExtractRunParamsCountry = "UY"
	ExtractRunParamsCountryUz  ExtractRunParamsCountry = "UZ"
	ExtractRunParamsCountryVa  ExtractRunParamsCountry = "VA"
	ExtractRunParamsCountryVc  ExtractRunParamsCountry = "VC"
	ExtractRunParamsCountryVe  ExtractRunParamsCountry = "VE"
	ExtractRunParamsCountryVg  ExtractRunParamsCountry = "VG"
	ExtractRunParamsCountryVi  ExtractRunParamsCountry = "VI"
	ExtractRunParamsCountryVn  ExtractRunParamsCountry = "VN"
	ExtractRunParamsCountryVu  ExtractRunParamsCountry = "VU"
	ExtractRunParamsCountryWf  ExtractRunParamsCountry = "WF"
	ExtractRunParamsCountryWs  ExtractRunParamsCountry = "WS"
	ExtractRunParamsCountryXk  ExtractRunParamsCountry = "XK"
	ExtractRunParamsCountryYe  ExtractRunParamsCountry = "YE"
	ExtractRunParamsCountryYt  ExtractRunParamsCountry = "YT"
	ExtractRunParamsCountryZa  ExtractRunParamsCountry = "ZA"
	ExtractRunParamsCountryZm  ExtractRunParamsCountry = "ZM"
	ExtractRunParamsCountryZw  ExtractRunParamsCountry = "ZW"
	ExtractRunParamsCountryAll ExtractRunParamsCountry = "ALL"
)

// Device type for browser emulation
type ExtractRunParamsDevice string

const (
	ExtractRunParamsDeviceDesktop ExtractRunParamsDevice = "desktop"
	ExtractRunParamsDeviceMobile  ExtractRunParamsDevice = "mobile"
	ExtractRunParamsDeviceTablet  ExtractRunParamsDevice = "tablet"
)

// Browser driver to use
type ExtractRunParamsDriver string

const (
	ExtractRunParamsDriverVx6     ExtractRunParamsDriver = "vx6"
	ExtractRunParamsDriverVx8     ExtractRunParamsDriver = "vx8"
	ExtractRunParamsDriverVx8Pro  ExtractRunParamsDriver = "vx8-pro"
	ExtractRunParamsDriverVx10    ExtractRunParamsDriver = "vx10"
	ExtractRunParamsDriverVx10Pro ExtractRunParamsDriver = "vx10-pro"
	ExtractRunParamsDriverVx12    ExtractRunParamsDriver = "vx12"
	ExtractRunParamsDriverVx12Pro ExtractRunParamsDriver = "vx12-pro"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsHeaderUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsHeaderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractRunParamsHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsHeaderUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Locale for browser language and region settings
type ExtractRunParamsLocale string

const (
	ExtractRunParamsLocaleAaDj      ExtractRunParamsLocale = "aa-DJ"
	ExtractRunParamsLocaleAaEr      ExtractRunParamsLocale = "aa-ER"
	ExtractRunParamsLocaleAaEt      ExtractRunParamsLocale = "aa-ET"
	ExtractRunParamsLocaleAf        ExtractRunParamsLocale = "af"
	ExtractRunParamsLocaleAfNa      ExtractRunParamsLocale = "af-NA"
	ExtractRunParamsLocaleAfZa      ExtractRunParamsLocale = "af-ZA"
	ExtractRunParamsLocaleAk        ExtractRunParamsLocale = "ak"
	ExtractRunParamsLocaleAkGh      ExtractRunParamsLocale = "ak-GH"
	ExtractRunParamsLocaleAm        ExtractRunParamsLocale = "am"
	ExtractRunParamsLocaleAmEt      ExtractRunParamsLocale = "am-ET"
	ExtractRunParamsLocaleAnEs      ExtractRunParamsLocale = "an-ES"
	ExtractRunParamsLocaleAr        ExtractRunParamsLocale = "ar"
	ExtractRunParamsLocaleArAe      ExtractRunParamsLocale = "ar-AE"
	ExtractRunParamsLocaleArBh      ExtractRunParamsLocale = "ar-BH"
	ExtractRunParamsLocaleArDz      ExtractRunParamsLocale = "ar-DZ"
	ExtractRunParamsLocaleArEg      ExtractRunParamsLocale = "ar-EG"
	ExtractRunParamsLocaleArIn      ExtractRunParamsLocale = "ar-IN"
	ExtractRunParamsLocaleArIq      ExtractRunParamsLocale = "ar-IQ"
	ExtractRunParamsLocaleArJo      ExtractRunParamsLocale = "ar-JO"
	ExtractRunParamsLocaleArKw      ExtractRunParamsLocale = "ar-KW"
	ExtractRunParamsLocaleArLb      ExtractRunParamsLocale = "ar-LB"
	ExtractRunParamsLocaleArLy      ExtractRunParamsLocale = "ar-LY"
	ExtractRunParamsLocaleArMa      ExtractRunParamsLocale = "ar-MA"
	ExtractRunParamsLocaleArOm      ExtractRunParamsLocale = "ar-OM"
	ExtractRunParamsLocaleArQa      ExtractRunParamsLocale = "ar-QA"
	ExtractRunParamsLocaleArSa      ExtractRunParamsLocale = "ar-SA"
	ExtractRunParamsLocaleArSd      ExtractRunParamsLocale = "ar-SD"
	ExtractRunParamsLocaleArSy      ExtractRunParamsLocale = "ar-SY"
	ExtractRunParamsLocaleArTn      ExtractRunParamsLocale = "ar-TN"
	ExtractRunParamsLocaleArYe      ExtractRunParamsLocale = "ar-YE"
	ExtractRunParamsLocaleAs        ExtractRunParamsLocale = "as"
	ExtractRunParamsLocaleAsIn      ExtractRunParamsLocale = "as-IN"
	ExtractRunParamsLocaleAsa       ExtractRunParamsLocale = "asa"
	ExtractRunParamsLocaleAsaTz     ExtractRunParamsLocale = "asa-TZ"
	ExtractRunParamsLocaleAstEs     ExtractRunParamsLocale = "ast-ES"
	ExtractRunParamsLocaleAz        ExtractRunParamsLocale = "az"
	ExtractRunParamsLocaleAzAz      ExtractRunParamsLocale = "az-AZ"
	ExtractRunParamsLocaleAzCyrl    ExtractRunParamsLocale = "az-Cyrl"
	ExtractRunParamsLocaleAzCyrlAz  ExtractRunParamsLocale = "az-Cyrl-AZ"
	ExtractRunParamsLocaleAzLatn    ExtractRunParamsLocale = "az-Latn"
	ExtractRunParamsLocaleAzLatnAz  ExtractRunParamsLocale = "az-Latn-AZ"
	ExtractRunParamsLocaleBe        ExtractRunParamsLocale = "be"
	ExtractRunParamsLocaleBeBy      ExtractRunParamsLocale = "be-BY"
	ExtractRunParamsLocaleBem       ExtractRunParamsLocale = "bem"
	ExtractRunParamsLocaleBemZm     ExtractRunParamsLocale = "bem-ZM"
	ExtractRunParamsLocaleBerDz     ExtractRunParamsLocale = "ber-DZ"
	ExtractRunParamsLocaleBerMa     ExtractRunParamsLocale = "ber-MA"
	ExtractRunParamsLocaleBez       ExtractRunParamsLocale = "bez"
	ExtractRunParamsLocaleBezTz     ExtractRunParamsLocale = "bez-TZ"
	ExtractRunParamsLocaleBg        ExtractRunParamsLocale = "bg"
	ExtractRunParamsLocaleBgBg      ExtractRunParamsLocale = "bg-BG"
	ExtractRunParamsLocaleBhoIn     ExtractRunParamsLocale = "bho-IN"
	ExtractRunParamsLocaleBm        ExtractRunParamsLocale = "bm"
	ExtractRunParamsLocaleBmMl      ExtractRunParamsLocale = "bm-ML"
	ExtractRunParamsLocaleBn        ExtractRunParamsLocale = "bn"
	ExtractRunParamsLocaleBnBd      ExtractRunParamsLocale = "bn-BD"
	ExtractRunParamsLocaleBnIn      ExtractRunParamsLocale = "bn-IN"
	ExtractRunParamsLocaleBo        ExtractRunParamsLocale = "bo"
	ExtractRunParamsLocaleBoCn      ExtractRunParamsLocale = "bo-CN"
	ExtractRunParamsLocaleBoIn      ExtractRunParamsLocale = "bo-IN"
	ExtractRunParamsLocaleBrFr      ExtractRunParamsLocale = "br-FR"
	ExtractRunParamsLocaleBrxIn     ExtractRunParamsLocale = "brx-IN"
	ExtractRunParamsLocaleBs        ExtractRunParamsLocale = "bs"
	ExtractRunParamsLocaleBsBa      ExtractRunParamsLocale = "bs-BA"
	ExtractRunParamsLocaleBynEr     ExtractRunParamsLocale = "byn-ER"
	ExtractRunParamsLocaleCa        ExtractRunParamsLocale = "ca"
	ExtractRunParamsLocaleCaAd      ExtractRunParamsLocale = "ca-AD"
	ExtractRunParamsLocaleCaEs      ExtractRunParamsLocale = "ca-ES"
	ExtractRunParamsLocaleCaFr      ExtractRunParamsLocale = "ca-FR"
	ExtractRunParamsLocaleCaIt      ExtractRunParamsLocale = "ca-IT"
	ExtractRunParamsLocaleCgg       ExtractRunParamsLocale = "cgg"
	ExtractRunParamsLocaleCggUg     ExtractRunParamsLocale = "cgg-UG"
	ExtractRunParamsLocaleChr       ExtractRunParamsLocale = "chr"
	ExtractRunParamsLocaleChrUs     ExtractRunParamsLocale = "chr-US"
	ExtractRunParamsLocaleCrhUa     ExtractRunParamsLocale = "crh-UA"
	ExtractRunParamsLocaleCs        ExtractRunParamsLocale = "cs"
	ExtractRunParamsLocaleCsCz      ExtractRunParamsLocale = "cs-CZ"
	ExtractRunParamsLocaleCsbPl     ExtractRunParamsLocale = "csb-PL"
	ExtractRunParamsLocaleCvRu      ExtractRunParamsLocale = "cv-RU"
	ExtractRunParamsLocaleCy        ExtractRunParamsLocale = "cy"
	ExtractRunParamsLocaleCyGB      ExtractRunParamsLocale = "cy-GB"
	ExtractRunParamsLocaleDa        ExtractRunParamsLocale = "da"
	ExtractRunParamsLocaleDaDk      ExtractRunParamsLocale = "da-DK"
	ExtractRunParamsLocaleDav       ExtractRunParamsLocale = "dav"
	ExtractRunParamsLocaleDavKe     ExtractRunParamsLocale = "dav-KE"
	ExtractRunParamsLocaleDe        ExtractRunParamsLocale = "de"
	ExtractRunParamsLocaleDeAt      ExtractRunParamsLocale = "de-AT"
	ExtractRunParamsLocaleDeBe      ExtractRunParamsLocale = "de-BE"
	ExtractRunParamsLocaleDeCh      ExtractRunParamsLocale = "de-CH"
	ExtractRunParamsLocaleDeDe      ExtractRunParamsLocale = "de-DE"
	ExtractRunParamsLocaleDeLi      ExtractRunParamsLocale = "de-LI"
	ExtractRunParamsLocaleDeLu      ExtractRunParamsLocale = "de-LU"
	ExtractRunParamsLocaleDvMv      ExtractRunParamsLocale = "dv-MV"
	ExtractRunParamsLocaleDzBt      ExtractRunParamsLocale = "dz-BT"
	ExtractRunParamsLocaleEbu       ExtractRunParamsLocale = "ebu"
	ExtractRunParamsLocaleEbuKe     ExtractRunParamsLocale = "ebu-KE"
	ExtractRunParamsLocaleEe        ExtractRunParamsLocale = "ee"
	ExtractRunParamsLocaleEeGh      ExtractRunParamsLocale = "ee-GH"
	ExtractRunParamsLocaleEeTg      ExtractRunParamsLocale = "ee-TG"
	ExtractRunParamsLocaleEl        ExtractRunParamsLocale = "el"
	ExtractRunParamsLocaleElCy      ExtractRunParamsLocale = "el-CY"
	ExtractRunParamsLocaleElGr      ExtractRunParamsLocale = "el-GR"
	ExtractRunParamsLocaleEn        ExtractRunParamsLocale = "en"
	ExtractRunParamsLocaleEnAg      ExtractRunParamsLocale = "en-AG"
	ExtractRunParamsLocaleEnAs      ExtractRunParamsLocale = "en-AS"
	ExtractRunParamsLocaleEnAu      ExtractRunParamsLocale = "en-AU"
	ExtractRunParamsLocaleEnBe      ExtractRunParamsLocale = "en-BE"
	ExtractRunParamsLocaleEnBw      ExtractRunParamsLocale = "en-BW"
	ExtractRunParamsLocaleEnBz      ExtractRunParamsLocale = "en-BZ"
	ExtractRunParamsLocaleEnCa      ExtractRunParamsLocale = "en-CA"
	ExtractRunParamsLocaleEnDk      ExtractRunParamsLocale = "en-DK"
	ExtractRunParamsLocaleEnGB      ExtractRunParamsLocale = "en-GB"
	ExtractRunParamsLocaleEnGu      ExtractRunParamsLocale = "en-GU"
	ExtractRunParamsLocaleEnHk      ExtractRunParamsLocale = "en-HK"
	ExtractRunParamsLocaleEnIe      ExtractRunParamsLocale = "en-IE"
	ExtractRunParamsLocaleEnIn      ExtractRunParamsLocale = "en-IN"
	ExtractRunParamsLocaleEnJm      ExtractRunParamsLocale = "en-JM"
	ExtractRunParamsLocaleEnMh      ExtractRunParamsLocale = "en-MH"
	ExtractRunParamsLocaleEnMp      ExtractRunParamsLocale = "en-MP"
	ExtractRunParamsLocaleEnMt      ExtractRunParamsLocale = "en-MT"
	ExtractRunParamsLocaleEnMu      ExtractRunParamsLocale = "en-MU"
	ExtractRunParamsLocaleEnNa      ExtractRunParamsLocale = "en-NA"
	ExtractRunParamsLocaleEnNg      ExtractRunParamsLocale = "en-NG"
	ExtractRunParamsLocaleEnNz      ExtractRunParamsLocale = "en-NZ"
	ExtractRunParamsLocaleEnPh      ExtractRunParamsLocale = "en-PH"
	ExtractRunParamsLocaleEnPk      ExtractRunParamsLocale = "en-PK"
	ExtractRunParamsLocaleEnSg      ExtractRunParamsLocale = "en-SG"
	ExtractRunParamsLocaleEnTt      ExtractRunParamsLocale = "en-TT"
	ExtractRunParamsLocaleEnUm      ExtractRunParamsLocale = "en-UM"
	ExtractRunParamsLocaleEnUs      ExtractRunParamsLocale = "en-US"
	ExtractRunParamsLocaleEnVi      ExtractRunParamsLocale = "en-VI"
	ExtractRunParamsLocaleEnZa      ExtractRunParamsLocale = "en-ZA"
	ExtractRunParamsLocaleEnZm      ExtractRunParamsLocale = "en-ZM"
	ExtractRunParamsLocaleEnZw      ExtractRunParamsLocale = "en-ZW"
	ExtractRunParamsLocaleEo        ExtractRunParamsLocale = "eo"
	ExtractRunParamsLocaleEs        ExtractRunParamsLocale = "es"
	ExtractRunParamsLocaleEs419     ExtractRunParamsLocale = "es-419"
	ExtractRunParamsLocaleEsAr      ExtractRunParamsLocale = "es-AR"
	ExtractRunParamsLocaleEsBo      ExtractRunParamsLocale = "es-BO"
	ExtractRunParamsLocaleEsCl      ExtractRunParamsLocale = "es-CL"
	ExtractRunParamsLocaleEsCo      ExtractRunParamsLocale = "es-CO"
	ExtractRunParamsLocaleEsCr      ExtractRunParamsLocale = "es-CR"
	ExtractRunParamsLocaleEsCu      ExtractRunParamsLocale = "es-CU"
	ExtractRunParamsLocaleEsDo      ExtractRunParamsLocale = "es-DO"
	ExtractRunParamsLocaleEsEc      ExtractRunParamsLocale = "es-EC"
	ExtractRunParamsLocaleEsEs      ExtractRunParamsLocale = "es-ES"
	ExtractRunParamsLocaleEsGq      ExtractRunParamsLocale = "es-GQ"
	ExtractRunParamsLocaleEsGt      ExtractRunParamsLocale = "es-GT"
	ExtractRunParamsLocaleEsHn      ExtractRunParamsLocale = "es-HN"
	ExtractRunParamsLocaleEsMx      ExtractRunParamsLocale = "es-MX"
	ExtractRunParamsLocaleEsNi      ExtractRunParamsLocale = "es-NI"
	ExtractRunParamsLocaleEsPa      ExtractRunParamsLocale = "es-PA"
	ExtractRunParamsLocaleEsPe      ExtractRunParamsLocale = "es-PE"
	ExtractRunParamsLocaleEsPr      ExtractRunParamsLocale = "es-PR"
	ExtractRunParamsLocaleEsPy      ExtractRunParamsLocale = "es-PY"
	ExtractRunParamsLocaleEsSv      ExtractRunParamsLocale = "es-SV"
	ExtractRunParamsLocaleEsUs      ExtractRunParamsLocale = "es-US"
	ExtractRunParamsLocaleEsUy      ExtractRunParamsLocale = "es-UY"
	ExtractRunParamsLocaleEsVe      ExtractRunParamsLocale = "es-VE"
	ExtractRunParamsLocaleEt        ExtractRunParamsLocale = "et"
	ExtractRunParamsLocaleEtEe      ExtractRunParamsLocale = "et-EE"
	ExtractRunParamsLocaleEu        ExtractRunParamsLocale = "eu"
	ExtractRunParamsLocaleEuEs      ExtractRunParamsLocale = "eu-ES"
	ExtractRunParamsLocaleFa        ExtractRunParamsLocale = "fa"
	ExtractRunParamsLocaleFaAf      ExtractRunParamsLocale = "fa-AF"
	ExtractRunParamsLocaleFaIr      ExtractRunParamsLocale = "fa-IR"
	ExtractRunParamsLocaleFf        ExtractRunParamsLocale = "ff"
	ExtractRunParamsLocaleFfSn      ExtractRunParamsLocale = "ff-SN"
	ExtractRunParamsLocaleFi        ExtractRunParamsLocale = "fi"
	ExtractRunParamsLocaleFiFi      ExtractRunParamsLocale = "fi-FI"
	ExtractRunParamsLocaleFil       ExtractRunParamsLocale = "fil"
	ExtractRunParamsLocaleFilPh     ExtractRunParamsLocale = "fil-PH"
	ExtractRunParamsLocaleFo        ExtractRunParamsLocale = "fo"
	ExtractRunParamsLocaleFoFo      ExtractRunParamsLocale = "fo-FO"
	ExtractRunParamsLocaleFr        ExtractRunParamsLocale = "fr"
	ExtractRunParamsLocaleFrBe      ExtractRunParamsLocale = "fr-BE"
	ExtractRunParamsLocaleFrBf      ExtractRunParamsLocale = "fr-BF"
	ExtractRunParamsLocaleFrBi      ExtractRunParamsLocale = "fr-BI"
	ExtractRunParamsLocaleFrBj      ExtractRunParamsLocale = "fr-BJ"
	ExtractRunParamsLocaleFrBl      ExtractRunParamsLocale = "fr-BL"
	ExtractRunParamsLocaleFrCa      ExtractRunParamsLocale = "fr-CA"
	ExtractRunParamsLocaleFrCd      ExtractRunParamsLocale = "fr-CD"
	ExtractRunParamsLocaleFrCf      ExtractRunParamsLocale = "fr-CF"
	ExtractRunParamsLocaleFrCg      ExtractRunParamsLocale = "fr-CG"
	ExtractRunParamsLocaleFrCh      ExtractRunParamsLocale = "fr-CH"
	ExtractRunParamsLocaleFrCi      ExtractRunParamsLocale = "fr-CI"
	ExtractRunParamsLocaleFrCm      ExtractRunParamsLocale = "fr-CM"
	ExtractRunParamsLocaleFrDj      ExtractRunParamsLocale = "fr-DJ"
	ExtractRunParamsLocaleFrFr      ExtractRunParamsLocale = "fr-FR"
	ExtractRunParamsLocaleFrGa      ExtractRunParamsLocale = "fr-GA"
	ExtractRunParamsLocaleFrGn      ExtractRunParamsLocale = "fr-GN"
	ExtractRunParamsLocaleFrGp      ExtractRunParamsLocale = "fr-GP"
	ExtractRunParamsLocaleFrGq      ExtractRunParamsLocale = "fr-GQ"
	ExtractRunParamsLocaleFrKm      ExtractRunParamsLocale = "fr-KM"
	ExtractRunParamsLocaleFrLu      ExtractRunParamsLocale = "fr-LU"
	ExtractRunParamsLocaleFrMc      ExtractRunParamsLocale = "fr-MC"
	ExtractRunParamsLocaleFrMf      ExtractRunParamsLocale = "fr-MF"
	ExtractRunParamsLocaleFrMg      ExtractRunParamsLocale = "fr-MG"
	ExtractRunParamsLocaleFrMl      ExtractRunParamsLocale = "fr-ML"
	ExtractRunParamsLocaleFrMq      ExtractRunParamsLocale = "fr-MQ"
	ExtractRunParamsLocaleFrNe      ExtractRunParamsLocale = "fr-NE"
	ExtractRunParamsLocaleFrRe      ExtractRunParamsLocale = "fr-RE"
	ExtractRunParamsLocaleFrRw      ExtractRunParamsLocale = "fr-RW"
	ExtractRunParamsLocaleFrSn      ExtractRunParamsLocale = "fr-SN"
	ExtractRunParamsLocaleFrTd      ExtractRunParamsLocale = "fr-TD"
	ExtractRunParamsLocaleFrTg      ExtractRunParamsLocale = "fr-TG"
	ExtractRunParamsLocaleFurIt     ExtractRunParamsLocale = "fur-IT"
	ExtractRunParamsLocaleFyDe      ExtractRunParamsLocale = "fy-DE"
	ExtractRunParamsLocaleFyNl      ExtractRunParamsLocale = "fy-NL"
	ExtractRunParamsLocaleGa        ExtractRunParamsLocale = "ga"
	ExtractRunParamsLocaleGaIe      ExtractRunParamsLocale = "ga-IE"
	ExtractRunParamsLocaleGdGB      ExtractRunParamsLocale = "gd-GB"
	ExtractRunParamsLocaleGezEr     ExtractRunParamsLocale = "gez-ER"
	ExtractRunParamsLocaleGezEt     ExtractRunParamsLocale = "gez-ET"
	ExtractRunParamsLocaleGl        ExtractRunParamsLocale = "gl"
	ExtractRunParamsLocaleGlEs      ExtractRunParamsLocale = "gl-ES"
	ExtractRunParamsLocaleGsw       ExtractRunParamsLocale = "gsw"
	ExtractRunParamsLocaleGswCh     ExtractRunParamsLocale = "gsw-CH"
	ExtractRunParamsLocaleGu        ExtractRunParamsLocale = "gu"
	ExtractRunParamsLocaleGuIn      ExtractRunParamsLocale = "gu-IN"
	ExtractRunParamsLocaleGuz       ExtractRunParamsLocale = "guz"
	ExtractRunParamsLocaleGuzKe     ExtractRunParamsLocale = "guz-KE"
	ExtractRunParamsLocaleGv        ExtractRunParamsLocale = "gv"
	ExtractRunParamsLocaleGvGB      ExtractRunParamsLocale = "gv-GB"
	ExtractRunParamsLocaleHa        ExtractRunParamsLocale = "ha"
	ExtractRunParamsLocaleHaLatn    ExtractRunParamsLocale = "ha-Latn"
	ExtractRunParamsLocaleHaLatnGh  ExtractRunParamsLocale = "ha-Latn-GH"
	ExtractRunParamsLocaleHaLatnNe  ExtractRunParamsLocale = "ha-Latn-NE"
	ExtractRunParamsLocaleHaLatnNg  ExtractRunParamsLocale = "ha-Latn-NG"
	ExtractRunParamsLocaleHaNg      ExtractRunParamsLocale = "ha-NG"
	ExtractRunParamsLocaleHaw       ExtractRunParamsLocale = "haw"
	ExtractRunParamsLocaleHawUs     ExtractRunParamsLocale = "haw-US"
	ExtractRunParamsLocaleHe        ExtractRunParamsLocale = "he"
	ExtractRunParamsLocaleHeIl      ExtractRunParamsLocale = "he-IL"
	ExtractRunParamsLocaleHi        ExtractRunParamsLocale = "hi"
	ExtractRunParamsLocaleHiIn      ExtractRunParamsLocale = "hi-IN"
	ExtractRunParamsLocaleHneIn     ExtractRunParamsLocale = "hne-IN"
	ExtractRunParamsLocaleHr        ExtractRunParamsLocale = "hr"
	ExtractRunParamsLocaleHrHr      ExtractRunParamsLocale = "hr-HR"
	ExtractRunParamsLocaleHsbDe     ExtractRunParamsLocale = "hsb-DE"
	ExtractRunParamsLocaleHtHt      ExtractRunParamsLocale = "ht-HT"
	ExtractRunParamsLocaleHu        ExtractRunParamsLocale = "hu"
	ExtractRunParamsLocaleHuHu      ExtractRunParamsLocale = "hu-HU"
	ExtractRunParamsLocaleHy        ExtractRunParamsLocale = "hy"
	ExtractRunParamsLocaleHyAm      ExtractRunParamsLocale = "hy-AM"
	ExtractRunParamsLocaleID        ExtractRunParamsLocale = "id"
	ExtractRunParamsLocaleIDID      ExtractRunParamsLocale = "id-ID"
	ExtractRunParamsLocaleIg        ExtractRunParamsLocale = "ig"
	ExtractRunParamsLocaleIgNg      ExtractRunParamsLocale = "ig-NG"
	ExtractRunParamsLocaleIi        ExtractRunParamsLocale = "ii"
	ExtractRunParamsLocaleIiCn      ExtractRunParamsLocale = "ii-CN"
	ExtractRunParamsLocaleIkCa      ExtractRunParamsLocale = "ik-CA"
	ExtractRunParamsLocaleIs        ExtractRunParamsLocale = "is"
	ExtractRunParamsLocaleIsIs      ExtractRunParamsLocale = "is-IS"
	ExtractRunParamsLocaleIt        ExtractRunParamsLocale = "it"
	ExtractRunParamsLocaleItCh      ExtractRunParamsLocale = "it-CH"
	ExtractRunParamsLocaleItIt      ExtractRunParamsLocale = "it-IT"
	ExtractRunParamsLocaleIuCa      ExtractRunParamsLocale = "iu-CA"
	ExtractRunParamsLocaleIwIl      ExtractRunParamsLocale = "iw-IL"
	ExtractRunParamsLocaleJa        ExtractRunParamsLocale = "ja"
	ExtractRunParamsLocaleJaJp      ExtractRunParamsLocale = "ja-JP"
	ExtractRunParamsLocaleJmc       ExtractRunParamsLocale = "jmc"
	ExtractRunParamsLocaleJmcTz     ExtractRunParamsLocale = "jmc-TZ"
	ExtractRunParamsLocaleKa        ExtractRunParamsLocale = "ka"
	ExtractRunParamsLocaleKaGe      ExtractRunParamsLocale = "ka-GE"
	ExtractRunParamsLocaleKab       ExtractRunParamsLocale = "kab"
	ExtractRunParamsLocaleKabDz     ExtractRunParamsLocale = "kab-DZ"
	ExtractRunParamsLocaleKam       ExtractRunParamsLocale = "kam"
	ExtractRunParamsLocaleKamKe     ExtractRunParamsLocale = "kam-KE"
	ExtractRunParamsLocaleKde       ExtractRunParamsLocale = "kde"
	ExtractRunParamsLocaleKdeTz     ExtractRunParamsLocale = "kde-TZ"
	ExtractRunParamsLocaleKea       ExtractRunParamsLocale = "kea"
	ExtractRunParamsLocaleKeaCv     ExtractRunParamsLocale = "kea-CV"
	ExtractRunParamsLocaleKhq       ExtractRunParamsLocale = "khq"
	ExtractRunParamsLocaleKhqMl     ExtractRunParamsLocale = "khq-ML"
	ExtractRunParamsLocaleKi        ExtractRunParamsLocale = "ki"
	ExtractRunParamsLocaleKiKe      ExtractRunParamsLocale = "ki-KE"
	ExtractRunParamsLocaleKk        ExtractRunParamsLocale = "kk"
	ExtractRunParamsLocaleKkCyrl    ExtractRunParamsLocale = "kk-Cyrl"
	ExtractRunParamsLocaleKkCyrlKz  ExtractRunParamsLocale = "kk-Cyrl-KZ"
	ExtractRunParamsLocaleKkKz      ExtractRunParamsLocale = "kk-KZ"
	ExtractRunParamsLocaleKl        ExtractRunParamsLocale = "kl"
	ExtractRunParamsLocaleKlGl      ExtractRunParamsLocale = "kl-GL"
	ExtractRunParamsLocaleKln       ExtractRunParamsLocale = "kln"
	ExtractRunParamsLocaleKlnKe     ExtractRunParamsLocale = "kln-KE"
	ExtractRunParamsLocaleKm        ExtractRunParamsLocale = "km"
	ExtractRunParamsLocaleKmKh      ExtractRunParamsLocale = "km-KH"
	ExtractRunParamsLocaleKn        ExtractRunParamsLocale = "kn"
	ExtractRunParamsLocaleKnIn      ExtractRunParamsLocale = "kn-IN"
	ExtractRunParamsLocaleKo        ExtractRunParamsLocale = "ko"
	ExtractRunParamsLocaleKoKr      ExtractRunParamsLocale = "ko-KR"
	ExtractRunParamsLocaleKok       ExtractRunParamsLocale = "kok"
	ExtractRunParamsLocaleKokIn     ExtractRunParamsLocale = "kok-IN"
	ExtractRunParamsLocaleKsIn      ExtractRunParamsLocale = "ks-IN"
	ExtractRunParamsLocaleKuTr      ExtractRunParamsLocale = "ku-TR"
	ExtractRunParamsLocaleKw        ExtractRunParamsLocale = "kw"
	ExtractRunParamsLocaleKwGB      ExtractRunParamsLocale = "kw-GB"
	ExtractRunParamsLocaleKyKg      ExtractRunParamsLocale = "ky-KG"
	ExtractRunParamsLocaleLag       ExtractRunParamsLocale = "lag"
	ExtractRunParamsLocaleLagTz     ExtractRunParamsLocale = "lag-TZ"
	ExtractRunParamsLocaleLbLu      ExtractRunParamsLocale = "lb-LU"
	ExtractRunParamsLocaleLg        ExtractRunParamsLocale = "lg"
	ExtractRunParamsLocaleLgUg      ExtractRunParamsLocale = "lg-UG"
	ExtractRunParamsLocaleLiBe      ExtractRunParamsLocale = "li-BE"
	ExtractRunParamsLocaleLiNl      ExtractRunParamsLocale = "li-NL"
	ExtractRunParamsLocaleLijIt     ExtractRunParamsLocale = "lij-IT"
	ExtractRunParamsLocaleLoLa      ExtractRunParamsLocale = "lo-LA"
	ExtractRunParamsLocaleLt        ExtractRunParamsLocale = "lt"
	ExtractRunParamsLocaleLtLt      ExtractRunParamsLocale = "lt-LT"
	ExtractRunParamsLocaleLuo       ExtractRunParamsLocale = "luo"
	ExtractRunParamsLocaleLuoKe     ExtractRunParamsLocale = "luo-KE"
	ExtractRunParamsLocaleLuy       ExtractRunParamsLocale = "luy"
	ExtractRunParamsLocaleLuyKe     ExtractRunParamsLocale = "luy-KE"
	ExtractRunParamsLocaleLv        ExtractRunParamsLocale = "lv"
	ExtractRunParamsLocaleLvLv      ExtractRunParamsLocale = "lv-LV"
	ExtractRunParamsLocaleMagIn     ExtractRunParamsLocale = "mag-IN"
	ExtractRunParamsLocaleMaiIn     ExtractRunParamsLocale = "mai-IN"
	ExtractRunParamsLocaleMas       ExtractRunParamsLocale = "mas"
	ExtractRunParamsLocaleMasKe     ExtractRunParamsLocale = "mas-KE"
	ExtractRunParamsLocaleMasTz     ExtractRunParamsLocale = "mas-TZ"
	ExtractRunParamsLocaleMer       ExtractRunParamsLocale = "mer"
	ExtractRunParamsLocaleMerKe     ExtractRunParamsLocale = "mer-KE"
	ExtractRunParamsLocaleMfe       ExtractRunParamsLocale = "mfe"
	ExtractRunParamsLocaleMfeMu     ExtractRunParamsLocale = "mfe-MU"
	ExtractRunParamsLocaleMg        ExtractRunParamsLocale = "mg"
	ExtractRunParamsLocaleMgMg      ExtractRunParamsLocale = "mg-MG"
	ExtractRunParamsLocaleMhrRu     ExtractRunParamsLocale = "mhr-RU"
	ExtractRunParamsLocaleMiNz      ExtractRunParamsLocale = "mi-NZ"
	ExtractRunParamsLocaleMk        ExtractRunParamsLocale = "mk"
	ExtractRunParamsLocaleMkMk      ExtractRunParamsLocale = "mk-MK"
	ExtractRunParamsLocaleMl        ExtractRunParamsLocale = "ml"
	ExtractRunParamsLocaleMlIn      ExtractRunParamsLocale = "ml-IN"
	ExtractRunParamsLocaleMnMn      ExtractRunParamsLocale = "mn-MN"
	ExtractRunParamsLocaleMr        ExtractRunParamsLocale = "mr"
	ExtractRunParamsLocaleMrIn      ExtractRunParamsLocale = "mr-IN"
	ExtractRunParamsLocaleMs        ExtractRunParamsLocale = "ms"
	ExtractRunParamsLocaleMsBn      ExtractRunParamsLocale = "ms-BN"
	ExtractRunParamsLocaleMsMy      ExtractRunParamsLocale = "ms-MY"
	ExtractRunParamsLocaleMt        ExtractRunParamsLocale = "mt"
	ExtractRunParamsLocaleMtMt      ExtractRunParamsLocale = "mt-MT"
	ExtractRunParamsLocaleMy        ExtractRunParamsLocale = "my"
	ExtractRunParamsLocaleMyMm      ExtractRunParamsLocale = "my-MM"
	ExtractRunParamsLocaleNanTw     ExtractRunParamsLocale = "nan-TW"
	ExtractRunParamsLocaleNaq       ExtractRunParamsLocale = "naq"
	ExtractRunParamsLocaleNaqNa     ExtractRunParamsLocale = "naq-NA"
	ExtractRunParamsLocaleNb        ExtractRunParamsLocale = "nb"
	ExtractRunParamsLocaleNbNo      ExtractRunParamsLocale = "nb-NO"
	ExtractRunParamsLocaleNd        ExtractRunParamsLocale = "nd"
	ExtractRunParamsLocaleNdZw      ExtractRunParamsLocale = "nd-ZW"
	ExtractRunParamsLocaleNdsDe     ExtractRunParamsLocale = "nds-DE"
	ExtractRunParamsLocaleNdsNl     ExtractRunParamsLocale = "nds-NL"
	ExtractRunParamsLocaleNe        ExtractRunParamsLocale = "ne"
	ExtractRunParamsLocaleNeIn      ExtractRunParamsLocale = "ne-IN"
	ExtractRunParamsLocaleNeNp      ExtractRunParamsLocale = "ne-NP"
	ExtractRunParamsLocaleNl        ExtractRunParamsLocale = "nl"
	ExtractRunParamsLocaleNlAw      ExtractRunParamsLocale = "nl-AW"
	ExtractRunParamsLocaleNlBe      ExtractRunParamsLocale = "nl-BE"
	ExtractRunParamsLocaleNlNl      ExtractRunParamsLocale = "nl-NL"
	ExtractRunParamsLocaleNn        ExtractRunParamsLocale = "nn"
	ExtractRunParamsLocaleNnNo      ExtractRunParamsLocale = "nn-NO"
	ExtractRunParamsLocaleNrZa      ExtractRunParamsLocale = "nr-ZA"
	ExtractRunParamsLocaleNsoZa     ExtractRunParamsLocale = "nso-ZA"
	ExtractRunParamsLocaleNyn       ExtractRunParamsLocale = "nyn"
	ExtractRunParamsLocaleNynUg     ExtractRunParamsLocale = "nyn-UG"
	ExtractRunParamsLocaleOcFr      ExtractRunParamsLocale = "oc-FR"
	ExtractRunParamsLocaleOm        ExtractRunParamsLocale = "om"
	ExtractRunParamsLocaleOmEt      ExtractRunParamsLocale = "om-ET"
	ExtractRunParamsLocaleOmKe      ExtractRunParamsLocale = "om-KE"
	ExtractRunParamsLocaleOr        ExtractRunParamsLocale = "or"
	ExtractRunParamsLocaleOrIn      ExtractRunParamsLocale = "or-IN"
	ExtractRunParamsLocaleOsRu      ExtractRunParamsLocale = "os-RU"
	ExtractRunParamsLocalePa        ExtractRunParamsLocale = "pa"
	ExtractRunParamsLocalePaArab    ExtractRunParamsLocale = "pa-Arab"
	ExtractRunParamsLocalePaArabPk  ExtractRunParamsLocale = "pa-Arab-PK"
	ExtractRunParamsLocalePaGuru    ExtractRunParamsLocale = "pa-Guru"
	ExtractRunParamsLocalePaGuruIn  ExtractRunParamsLocale = "pa-Guru-IN"
	ExtractRunParamsLocalePaIn      ExtractRunParamsLocale = "pa-IN"
	ExtractRunParamsLocalePaPk      ExtractRunParamsLocale = "pa-PK"
	ExtractRunParamsLocalePapAn     ExtractRunParamsLocale = "pap-AN"
	ExtractRunParamsLocalePl        ExtractRunParamsLocale = "pl"
	ExtractRunParamsLocalePlPl      ExtractRunParamsLocale = "pl-PL"
	ExtractRunParamsLocalePs        ExtractRunParamsLocale = "ps"
	ExtractRunParamsLocalePsAf      ExtractRunParamsLocale = "ps-AF"
	ExtractRunParamsLocalePt        ExtractRunParamsLocale = "pt"
	ExtractRunParamsLocalePtBr      ExtractRunParamsLocale = "pt-BR"
	ExtractRunParamsLocalePtGw      ExtractRunParamsLocale = "pt-GW"
	ExtractRunParamsLocalePtMz      ExtractRunParamsLocale = "pt-MZ"
	ExtractRunParamsLocalePtPt      ExtractRunParamsLocale = "pt-PT"
	ExtractRunParamsLocaleRm        ExtractRunParamsLocale = "rm"
	ExtractRunParamsLocaleRmCh      ExtractRunParamsLocale = "rm-CH"
	ExtractRunParamsLocaleRo        ExtractRunParamsLocale = "ro"
	ExtractRunParamsLocaleRoMd      ExtractRunParamsLocale = "ro-MD"
	ExtractRunParamsLocaleRoRo      ExtractRunParamsLocale = "ro-RO"
	ExtractRunParamsLocaleRof       ExtractRunParamsLocale = "rof"
	ExtractRunParamsLocaleRofTz     ExtractRunParamsLocale = "rof-TZ"
	ExtractRunParamsLocaleRu        ExtractRunParamsLocale = "ru"
	ExtractRunParamsLocaleRuMd      ExtractRunParamsLocale = "ru-MD"
	ExtractRunParamsLocaleRuRu      ExtractRunParamsLocale = "ru-RU"
	ExtractRunParamsLocaleRuUa      ExtractRunParamsLocale = "ru-UA"
	ExtractRunParamsLocaleRw        ExtractRunParamsLocale = "rw"
	ExtractRunParamsLocaleRwRw      ExtractRunParamsLocale = "rw-RW"
	ExtractRunParamsLocaleRwk       ExtractRunParamsLocale = "rwk"
	ExtractRunParamsLocaleRwkTz     ExtractRunParamsLocale = "rwk-TZ"
	ExtractRunParamsLocaleSaIn      ExtractRunParamsLocale = "sa-IN"
	ExtractRunParamsLocaleSaq       ExtractRunParamsLocale = "saq"
	ExtractRunParamsLocaleSaqKe     ExtractRunParamsLocale = "saq-KE"
	ExtractRunParamsLocaleScIt      ExtractRunParamsLocale = "sc-IT"
	ExtractRunParamsLocaleSdIn      ExtractRunParamsLocale = "sd-IN"
	ExtractRunParamsLocaleSeNo      ExtractRunParamsLocale = "se-NO"
	ExtractRunParamsLocaleSeh       ExtractRunParamsLocale = "seh"
	ExtractRunParamsLocaleSehMz     ExtractRunParamsLocale = "seh-MZ"
	ExtractRunParamsLocaleSes       ExtractRunParamsLocale = "ses"
	ExtractRunParamsLocaleSesMl     ExtractRunParamsLocale = "ses-ML"
	ExtractRunParamsLocaleSg        ExtractRunParamsLocale = "sg"
	ExtractRunParamsLocaleSgCf      ExtractRunParamsLocale = "sg-CF"
	ExtractRunParamsLocaleShi       ExtractRunParamsLocale = "shi"
	ExtractRunParamsLocaleShiLatn   ExtractRunParamsLocale = "shi-Latn"
	ExtractRunParamsLocaleShiLatnMa ExtractRunParamsLocale = "shi-Latn-MA"
	ExtractRunParamsLocaleShiTfng   ExtractRunParamsLocale = "shi-Tfng"
	ExtractRunParamsLocaleShiTfngMa ExtractRunParamsLocale = "shi-Tfng-MA"
	ExtractRunParamsLocaleShsCa     ExtractRunParamsLocale = "shs-CA"
	ExtractRunParamsLocaleSi        ExtractRunParamsLocale = "si"
	ExtractRunParamsLocaleSiLk      ExtractRunParamsLocale = "si-LK"
	ExtractRunParamsLocaleSidEt     ExtractRunParamsLocale = "sid-ET"
	ExtractRunParamsLocaleSk        ExtractRunParamsLocale = "sk"
	ExtractRunParamsLocaleSkSk      ExtractRunParamsLocale = "sk-SK"
	ExtractRunParamsLocaleSl        ExtractRunParamsLocale = "sl"
	ExtractRunParamsLocaleSlSi      ExtractRunParamsLocale = "sl-SI"
	ExtractRunParamsLocaleSn        ExtractRunParamsLocale = "sn"
	ExtractRunParamsLocaleSnZw      ExtractRunParamsLocale = "sn-ZW"
	ExtractRunParamsLocaleSo        ExtractRunParamsLocale = "so"
	ExtractRunParamsLocaleSoDj      ExtractRunParamsLocale = "so-DJ"
	ExtractRunParamsLocaleSoEt      ExtractRunParamsLocale = "so-ET"
	ExtractRunParamsLocaleSoKe      ExtractRunParamsLocale = "so-KE"
	ExtractRunParamsLocaleSoSo      ExtractRunParamsLocale = "so-SO"
	ExtractRunParamsLocaleSq        ExtractRunParamsLocale = "sq"
	ExtractRunParamsLocaleSqAl      ExtractRunParamsLocale = "sq-AL"
	ExtractRunParamsLocaleSqMk      ExtractRunParamsLocale = "sq-MK"
	ExtractRunParamsLocaleSr        ExtractRunParamsLocale = "sr"
	ExtractRunParamsLocaleSrCyrl    ExtractRunParamsLocale = "sr-Cyrl"
	ExtractRunParamsLocaleSrCyrlBa  ExtractRunParamsLocale = "sr-Cyrl-BA"
	ExtractRunParamsLocaleSrCyrlMe  ExtractRunParamsLocale = "sr-Cyrl-ME"
	ExtractRunParamsLocaleSrCyrlRs  ExtractRunParamsLocale = "sr-Cyrl-RS"
	ExtractRunParamsLocaleSrLatn    ExtractRunParamsLocale = "sr-Latn"
	ExtractRunParamsLocaleSrLatnBa  ExtractRunParamsLocale = "sr-Latn-BA"
	ExtractRunParamsLocaleSrLatnMe  ExtractRunParamsLocale = "sr-Latn-ME"
	ExtractRunParamsLocaleSrLatnRs  ExtractRunParamsLocale = "sr-Latn-RS"
	ExtractRunParamsLocaleSrMe      ExtractRunParamsLocale = "sr-ME"
	ExtractRunParamsLocaleSrRs      ExtractRunParamsLocale = "sr-RS"
	ExtractRunParamsLocaleSSZa      ExtractRunParamsLocale = "ss-ZA"
	ExtractRunParamsLocaleStZa      ExtractRunParamsLocale = "st-ZA"
	ExtractRunParamsLocaleSv        ExtractRunParamsLocale = "sv"
	ExtractRunParamsLocaleSvFi      ExtractRunParamsLocale = "sv-FI"
	ExtractRunParamsLocaleSvSe      ExtractRunParamsLocale = "sv-SE"
	ExtractRunParamsLocaleSw        ExtractRunParamsLocale = "sw"
	ExtractRunParamsLocaleSwKe      ExtractRunParamsLocale = "sw-KE"
	ExtractRunParamsLocaleSwTz      ExtractRunParamsLocale = "sw-TZ"
	ExtractRunParamsLocaleTa        ExtractRunParamsLocale = "ta"
	ExtractRunParamsLocaleTaIn      ExtractRunParamsLocale = "ta-IN"
	ExtractRunParamsLocaleTaLk      ExtractRunParamsLocale = "ta-LK"
	ExtractRunParamsLocaleTe        ExtractRunParamsLocale = "te"
	ExtractRunParamsLocaleTeIn      ExtractRunParamsLocale = "te-IN"
	ExtractRunParamsLocaleTeo       ExtractRunParamsLocale = "teo"
	ExtractRunParamsLocaleTeoKe     ExtractRunParamsLocale = "teo-KE"
	ExtractRunParamsLocaleTeoUg     ExtractRunParamsLocale = "teo-UG"
	ExtractRunParamsLocaleTgTj      ExtractRunParamsLocale = "tg-TJ"
	ExtractRunParamsLocaleTh        ExtractRunParamsLocale = "th"
	ExtractRunParamsLocaleThTh      ExtractRunParamsLocale = "th-TH"
	ExtractRunParamsLocaleTi        ExtractRunParamsLocale = "ti"
	ExtractRunParamsLocaleTiEr      ExtractRunParamsLocale = "ti-ER"
	ExtractRunParamsLocaleTiEt      ExtractRunParamsLocale = "ti-ET"
	ExtractRunParamsLocaleTigEr     ExtractRunParamsLocale = "tig-ER"
	ExtractRunParamsLocaleTkTm      ExtractRunParamsLocale = "tk-TM"
	ExtractRunParamsLocaleTlPh      ExtractRunParamsLocale = "tl-PH"
	ExtractRunParamsLocaleTnZa      ExtractRunParamsLocale = "tn-ZA"
	ExtractRunParamsLocaleTo        ExtractRunParamsLocale = "to"
	ExtractRunParamsLocaleToTo      ExtractRunParamsLocale = "to-TO"
	ExtractRunParamsLocaleTr        ExtractRunParamsLocale = "tr"
	ExtractRunParamsLocaleTrCy      ExtractRunParamsLocale = "tr-CY"
	ExtractRunParamsLocaleTrTr      ExtractRunParamsLocale = "tr-TR"
	ExtractRunParamsLocaleTsZa      ExtractRunParamsLocale = "ts-ZA"
	ExtractRunParamsLocaleTtRu      ExtractRunParamsLocale = "tt-RU"
	ExtractRunParamsLocaleTzm       ExtractRunParamsLocale = "tzm"
	ExtractRunParamsLocaleTzmLatn   ExtractRunParamsLocale = "tzm-Latn"
	ExtractRunParamsLocaleTzmLatnMa ExtractRunParamsLocale = "tzm-Latn-MA"
	ExtractRunParamsLocaleUgCn      ExtractRunParamsLocale = "ug-CN"
	ExtractRunParamsLocaleUk        ExtractRunParamsLocale = "uk"
	ExtractRunParamsLocaleUkUa      ExtractRunParamsLocale = "uk-UA"
	ExtractRunParamsLocaleUnmUs     ExtractRunParamsLocale = "unm-US"
	ExtractRunParamsLocaleUr        ExtractRunParamsLocale = "ur"
	ExtractRunParamsLocaleUrIn      ExtractRunParamsLocale = "ur-IN"
	ExtractRunParamsLocaleUrPk      ExtractRunParamsLocale = "ur-PK"
	ExtractRunParamsLocaleUz        ExtractRunParamsLocale = "uz"
	ExtractRunParamsLocaleUzArab    ExtractRunParamsLocale = "uz-Arab"
	ExtractRunParamsLocaleUzArabAf  ExtractRunParamsLocale = "uz-Arab-AF"
	ExtractRunParamsLocaleUzCyrl    ExtractRunParamsLocale = "uz-Cyrl"
	ExtractRunParamsLocaleUzCyrlUz  ExtractRunParamsLocale = "uz-Cyrl-UZ"
	ExtractRunParamsLocaleUzLatn    ExtractRunParamsLocale = "uz-Latn"
	ExtractRunParamsLocaleUzLatnUz  ExtractRunParamsLocale = "uz-Latn-UZ"
	ExtractRunParamsLocaleUzUz      ExtractRunParamsLocale = "uz-UZ"
	ExtractRunParamsLocaleVeZa      ExtractRunParamsLocale = "ve-ZA"
	ExtractRunParamsLocaleVi        ExtractRunParamsLocale = "vi"
	ExtractRunParamsLocaleViVn      ExtractRunParamsLocale = "vi-VN"
	ExtractRunParamsLocaleVun       ExtractRunParamsLocale = "vun"
	ExtractRunParamsLocaleVunTz     ExtractRunParamsLocale = "vun-TZ"
	ExtractRunParamsLocaleWaBe      ExtractRunParamsLocale = "wa-BE"
	ExtractRunParamsLocaleWaeCh     ExtractRunParamsLocale = "wae-CH"
	ExtractRunParamsLocaleWalEt     ExtractRunParamsLocale = "wal-ET"
	ExtractRunParamsLocaleWoSn      ExtractRunParamsLocale = "wo-SN"
	ExtractRunParamsLocaleXhZa      ExtractRunParamsLocale = "xh-ZA"
	ExtractRunParamsLocaleXog       ExtractRunParamsLocale = "xog"
	ExtractRunParamsLocaleXogUg     ExtractRunParamsLocale = "xog-UG"
	ExtractRunParamsLocaleYiUs      ExtractRunParamsLocale = "yi-US"
	ExtractRunParamsLocaleYo        ExtractRunParamsLocale = "yo"
	ExtractRunParamsLocaleYoNg      ExtractRunParamsLocale = "yo-NG"
	ExtractRunParamsLocaleYueHk     ExtractRunParamsLocale = "yue-HK"
	ExtractRunParamsLocaleZh        ExtractRunParamsLocale = "zh"
	ExtractRunParamsLocaleZhCn      ExtractRunParamsLocale = "zh-CN"
	ExtractRunParamsLocaleZhHk      ExtractRunParamsLocale = "zh-HK"
	ExtractRunParamsLocaleZhHans    ExtractRunParamsLocale = "zh-Hans"
	ExtractRunParamsLocaleZhHansCn  ExtractRunParamsLocale = "zh-Hans-CN"
	ExtractRunParamsLocaleZhHansHk  ExtractRunParamsLocale = "zh-Hans-HK"
	ExtractRunParamsLocaleZhHansMo  ExtractRunParamsLocale = "zh-Hans-MO"
	ExtractRunParamsLocaleZhHansSg  ExtractRunParamsLocale = "zh-Hans-SG"
	ExtractRunParamsLocaleZhHant    ExtractRunParamsLocale = "zh-Hant"
	ExtractRunParamsLocaleZhHantHk  ExtractRunParamsLocale = "zh-Hant-HK"
	ExtractRunParamsLocaleZhHantMo  ExtractRunParamsLocale = "zh-Hant-MO"
	ExtractRunParamsLocaleZhHantTw  ExtractRunParamsLocale = "zh-Hant-TW"
	ExtractRunParamsLocaleZhSg      ExtractRunParamsLocale = "zh-SG"
	ExtractRunParamsLocaleZhTw      ExtractRunParamsLocale = "zh-TW"
	ExtractRunParamsLocaleZu        ExtractRunParamsLocale = "zu"
	ExtractRunParamsLocaleZuZa      ExtractRunParamsLocale = "zu-ZA"
	ExtractRunParamsLocaleAuto      ExtractRunParamsLocale = "auto"
)

// HTTP method for the request
type ExtractRunParamsMethod string

const (
	ExtractRunParamsMethodGet    ExtractRunParamsMethod = "GET"
	ExtractRunParamsMethodPost   ExtractRunParamsMethod = "POST"
	ExtractRunParamsMethodPut    ExtractRunParamsMethod = "PUT"
	ExtractRunParamsMethodPatch  ExtractRunParamsMethod = "PATCH"
	ExtractRunParamsMethodDelete ExtractRunParamsMethod = "DELETE"
)

type ExtractRunParamsNetworkCapture struct {
	Validation                  param.Opt[bool]    `json:"validation,omitzero"`
	WaitForRequestsCount        param.Opt[float64] `json:"wait_for_requests_count,omitzero"`
	WaitForRequestsCountTimeout param.Opt[float64] `json:"wait_for_requests_count_timeout,omitzero"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method,omitzero"`
	// Resource type for network capture filtering
	ResourceType ExtractRunParamsNetworkCaptureResourceTypeUnion `json:"resource_type,omitzero"`
	StatusCode   ExtractRunParamsNetworkCaptureStatusCodeUnion   `json:"status_code,omitzero"`
	URL          ExtractRunParamsNetworkCaptureURL               `json:"url,omitzero"`
	paramObj
}

func (r ExtractRunParamsNetworkCapture) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsNetworkCapture
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractRunParamsNetworkCapture](
		"method", "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsNetworkCaptureResourceTypeUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsNetworkCaptureResourceTypeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractRunParamsNetworkCaptureResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsNetworkCaptureResourceTypeUnion) asAny() any {
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
type ExtractRunParamsNetworkCaptureStatusCodeUnion struct {
	OfFloat      param.Opt[float64] `json:",omitzero,inline"`
	OfFloatArray []float64          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsNetworkCaptureStatusCodeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfFloatArray)
}
func (u *ExtractRunParamsNetworkCaptureStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsNetworkCaptureStatusCodeUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfFloatArray) {
		return &u.OfFloatArray
	}
	return nil
}

// The property Value is required.
type ExtractRunParamsNetworkCaptureURL struct {
	Value string `json:"value,required"`
	// Any of "exact", "contains".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ExtractRunParamsNetworkCaptureURL) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsNetworkCaptureURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsNetworkCaptureURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractRunParamsNetworkCaptureURL](
		"type", "exact", "contains",
	)
}

// Operating system to emulate
type ExtractRunParamsOs string

const (
	ExtractRunParamsOsWindows ExtractRunParamsOs = "windows"
	ExtractRunParamsOsMacOs   ExtractRunParamsOs = "mac os"
	ExtractRunParamsOsLinux   ExtractRunParamsOs = "linux"
	ExtractRunParamsOsAndroid ExtractRunParamsOs = "android"
	ExtractRunParamsOsIos     ExtractRunParamsOs = "ios"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsParserUnion struct {
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsParserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap, u.OfString)
}
func (u *ExtractRunParamsParserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsParserUnion) asAny() any {
	if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Referrer policy for the request
type ExtractRunParamsReferrerType string

const (
	ExtractRunParamsReferrerTypeRandom     ExtractRunParamsReferrerType = "random"
	ExtractRunParamsReferrerTypeNoReferer  ExtractRunParamsReferrerType = "no-referer"
	ExtractRunParamsReferrerTypeSameOrigin ExtractRunParamsReferrerType = "same-origin"
	ExtractRunParamsReferrerTypeGoogle     ExtractRunParamsReferrerType = "google"
	ExtractRunParamsReferrerTypeBing       ExtractRunParamsReferrerType = "bing"
	ExtractRunParamsReferrerTypeFacebook   ExtractRunParamsReferrerType = "facebook"
	ExtractRunParamsReferrerTypeTwitter    ExtractRunParamsReferrerType = "twitter"
	ExtractRunParamsReferrerTypeInstagram  ExtractRunParamsReferrerType = "instagram"
)

type ExtractRunParamsSession struct {
	ID                  param.Opt[string]  `json:"id,omitzero"`
	PrefetchUserbrowser param.Opt[bool]    `json:"prefetch_userbrowser,omitzero"`
	Retry               param.Opt[bool]    `json:"retry,omitzero"`
	Timeout             param.Opt[float64] `json:"timeout,omitzero"`
	paramObj
}

func (r ExtractRunParamsSession) MarshalJSON() (data []byte, err error) {
	type shadow ExtractRunParamsSession
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractRunParamsSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsSkillUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsSkillUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractRunParamsSkillUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsSkillUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// US state for geolocation (only valid when country is US)
type ExtractRunParamsState string

const (
	ExtractRunParamsStateAl ExtractRunParamsState = "AL"
	ExtractRunParamsStateAk ExtractRunParamsState = "AK"
	ExtractRunParamsStateAs ExtractRunParamsState = "AS"
	ExtractRunParamsStateAz ExtractRunParamsState = "AZ"
	ExtractRunParamsStateAr ExtractRunParamsState = "AR"
	ExtractRunParamsStateCa ExtractRunParamsState = "CA"
	ExtractRunParamsStateCo ExtractRunParamsState = "CO"
	ExtractRunParamsStateCt ExtractRunParamsState = "CT"
	ExtractRunParamsStateDe ExtractRunParamsState = "DE"
	ExtractRunParamsStateDc ExtractRunParamsState = "DC"
	ExtractRunParamsStateFl ExtractRunParamsState = "FL"
	ExtractRunParamsStateGa ExtractRunParamsState = "GA"
	ExtractRunParamsStateGu ExtractRunParamsState = "GU"
	ExtractRunParamsStateHi ExtractRunParamsState = "HI"
	ExtractRunParamsStateID ExtractRunParamsState = "ID"
	ExtractRunParamsStateIl ExtractRunParamsState = "IL"
	ExtractRunParamsStateIn ExtractRunParamsState = "IN"
	ExtractRunParamsStateIa ExtractRunParamsState = "IA"
	ExtractRunParamsStateKs ExtractRunParamsState = "KS"
	ExtractRunParamsStateKy ExtractRunParamsState = "KY"
	ExtractRunParamsStateLa ExtractRunParamsState = "LA"
	ExtractRunParamsStateMe ExtractRunParamsState = "ME"
	ExtractRunParamsStateMd ExtractRunParamsState = "MD"
	ExtractRunParamsStateMa ExtractRunParamsState = "MA"
	ExtractRunParamsStateMi ExtractRunParamsState = "MI"
	ExtractRunParamsStateMn ExtractRunParamsState = "MN"
	ExtractRunParamsStateMs ExtractRunParamsState = "MS"
	ExtractRunParamsStateMo ExtractRunParamsState = "MO"
	ExtractRunParamsStateMt ExtractRunParamsState = "MT"
	ExtractRunParamsStateNe ExtractRunParamsState = "NE"
	ExtractRunParamsStateNv ExtractRunParamsState = "NV"
	ExtractRunParamsStateNh ExtractRunParamsState = "NH"
	ExtractRunParamsStateNj ExtractRunParamsState = "NJ"
	ExtractRunParamsStateNm ExtractRunParamsState = "NM"
	ExtractRunParamsStateNy ExtractRunParamsState = "NY"
	ExtractRunParamsStateNc ExtractRunParamsState = "NC"
	ExtractRunParamsStateNd ExtractRunParamsState = "ND"
	ExtractRunParamsStateMp ExtractRunParamsState = "MP"
	ExtractRunParamsStateOh ExtractRunParamsState = "OH"
	ExtractRunParamsStateOk ExtractRunParamsState = "OK"
	ExtractRunParamsStateOr ExtractRunParamsState = "OR"
	ExtractRunParamsStatePa ExtractRunParamsState = "PA"
	ExtractRunParamsStatePr ExtractRunParamsState = "PR"
	ExtractRunParamsStateRi ExtractRunParamsState = "RI"
	ExtractRunParamsStateSc ExtractRunParamsState = "SC"
	ExtractRunParamsStateSd ExtractRunParamsState = "SD"
	ExtractRunParamsStateTn ExtractRunParamsState = "TN"
	ExtractRunParamsStateTx ExtractRunParamsState = "TX"
	ExtractRunParamsStateUt ExtractRunParamsState = "UT"
	ExtractRunParamsStateVt ExtractRunParamsState = "VT"
	ExtractRunParamsStateVa ExtractRunParamsState = "VA"
	ExtractRunParamsStateVi ExtractRunParamsState = "VI"
	ExtractRunParamsStateWa ExtractRunParamsState = "WA"
	ExtractRunParamsStateWv ExtractRunParamsState = "WV"
	ExtractRunParamsStateWi ExtractRunParamsState = "WI"
	ExtractRunParamsStateWy ExtractRunParamsState = "WY"
)
