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
	"github.com/Nimbleway/nimble-go/shared"
	"github.com/Nimbleway/nimble-go/shared/constant"
)

// ExtractService contains methods and other services that help with interacting
// with the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtractService] method instead.
type ExtractService struct {
	Options   []option.RequestOption
	Templates ExtractTemplateService
}

// NewExtractService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewExtractService(opts ...option.RequestOption) (r ExtractService) {
	r = ExtractService{}
	r.Options = opts
	r.Templates = NewExtractTemplateService(opts...)
	return
}

// Extract Async Endpoint
func (r *ExtractService) Async(ctx context.Context, body ExtractAsyncParams, opts ...option.RequestOption) (res *ExtractAsyncResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/extract/async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Extract Batch Endpoint
func (r *ExtractService) Batch(ctx context.Context, body ExtractBatchParams, opts ...option.RequestOption) (res *ExtractBatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/extract/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Extract
func (r *ExtractService) Run(ctx context.Context, body ExtractRunParams, opts ...option.RequestOption) (res *ExtractRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/extract"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
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
	BatchID string `json:"batch_id" api:"nullable"`
	// URL for downloading the task results.
	DownloadURL string `json:"download_url" api:"nullable" format:"uri"`
	// Error message if the task failed.
	Error string `json:"error" api:"nullable"`
	// Classification of the error type.
	ErrorType string `json:"error_type" api:"nullable"`
	// Timestamp when the task was last modified.
	ModifiedAt string `json:"modified_at"`
	// Storage location of the output data.
	OutputURL string `json:"output_url" api:"nullable"`
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
	BatchID string `json:"batch_id" api:"nullable"`
	// URL for downloading the task results.
	DownloadURL string `json:"download_url" api:"nullable" format:"uri"`
	// Error message if the task failed.
	Error string `json:"error" api:"nullable"`
	// Classification of the error type.
	ErrorType string `json:"error_type" api:"nullable"`
	// Timestamp when the task was last modified.
	ModifiedAt string `json:"modified_at"`
	// Storage location of the output data.
	OutputURL string `json:"output_url" api:"nullable"`
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
func (r ExtractBatchResponseTask) RawJSON() string { return r.JSON.raw }
func (r *ExtractBatchResponseTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponse struct {
	Data     ExtractRunResponseData     `json:"data" api:"required"`
	Metadata ExtractRunResponseMetadata `json:"metadata" api:"required"`
	// The status of the task.
	//
	// Any of "success", "skipped", "fatal", "error", "postponed", "ignored",
	// "rejected", "blocked".
	Status ExtractRunResponseStatus `json:"status" api:"required"`
	// Unique identifier for the task.
	TaskID string `json:"task_id" api:"required"`
	// The final URL.
	URL   string                  `json:"url" api:"required"`
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
	// List of all unique URLs found on the page.
	Links []string `json:"links"`
	// The Markdown version of the HTML content.
	Markdown string `json:"markdown"`
	// The network capture data collected during the task.
	NetworkCapture []ExtractRunResponseDataNetworkCapture `json:"network_capture"`
	// Individual HTML content of each pagination page, before merging.
	PagesHTML []string `json:"pages_html"`
	// The parsing results extracted from the HTML & network content.
	Parsing ExtractRunResponseDataParsingUnion `json:"parsing"`
	// The list of redirects that occurred during the task.
	Redirects []ExtractRunResponseDataRedirect `json:"redirects"`
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
func (r ExtractRunResponseData) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser actions execution results. Present only when browser_actions were
// specified in the request.
type ExtractRunResponseDataBrowserActions struct {
	Results       []ExtractRunResponseDataBrowserActionsResult `json:"results" api:"required"`
	Success       bool                                         `json:"success" api:"required"`
	TotalDuration float64                                      `json:"total_duration" api:"required"`
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
func (r ExtractRunResponseDataBrowserActionsResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDataBrowserActionsResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseDataNetworkCapture struct {
	Filter       ExtractRunResponseDataNetworkCaptureFilter   `json:"filter" api:"required"`
	Results      []ExtractRunResponseDataNetworkCaptureResult `json:"results" api:"required"`
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
	Validation           bool    `json:"validation" api:"required"`
	WaitForRequestsCount float64 `json:"wait_for_requests_count" api:"required"`
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
func (r ExtractRunResponseDataNetworkCaptureFilterURL) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDataNetworkCaptureFilterURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseDataNetworkCaptureResult struct {
	Request  ExtractRunResponseDataNetworkCaptureResultRequest  `json:"request" api:"required"`
	Response ExtractRunResponseDataNetworkCaptureResultResponse `json:"response" api:"required"`
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
func (r ExtractRunResponseDataNetworkCaptureResultRequest) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDataNetworkCaptureResultRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseDataNetworkCaptureResultResponse struct {
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
func (r ExtractRunResponseDataParsingParsingSuccessResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDataParsingParsingSuccessResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseDataParsingParsingErrorResult struct {
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
func (r ExtractRunResponseDataParsingParsingErrorResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractRunResponseDataParsingParsingErrorResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractRunResponseDataRedirect struct {
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
	NextPageParams map[string]any `json:"next_page_params" api:"required"`
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
	NextPageParams map[string]any `json:"next_page_params" api:"required"`
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
	// Overall deadline in milliseconds for a realtime request. Clamped to the account
	// total timeout — it can shorten the deadline but never extend it. Has no effect
	// on async requests.
	RealtimeTotalTimeout param.Opt[float64] `json:"realtime_total_timeout,omitzero"`
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
	// Custom flow for the optimization engine: maps candidate names to the number of
	// attempts to spend on each candidate before advancing (0 skips it). Key order
	// defines the flow order. Providing it opts the request into 'auto' driver
	// selection.
	AutoDriverConfiguration map[string]int64 `json:"auto_driver_configuration,omitzero"`
	// Request body for POST, PUT, PATCH methods
	Body any `json:"body,omitzero"`
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
	// Any of "auto", "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro",
	// "media-vx6", "fast-vx6".
	Driver ExtractAsyncParamsDriver `json:"driver,omitzero"`
	// Expected HTTP status codes for successful requests
	ExpectedStatusCodes []int64 `json:"expected_status_codes,omitzero"`
	// List of acceptable response formats in order of preference
	//
	// Any of "html", "markdown", "screenshot", "headers", "links".
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
	// Selects which markdown conversion strategy to use. "full_page" converts the
	// entire HTML page. "main_content" uses Mozilla Readability to extract the main
	// article content before converting.
	//
	// Any of "full_page", "main_content".
	MarkdownBackend ExtractAsyncParamsMarkdownBackend `json:"markdown_backend,omitzero"`
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
	// Whether to render JavaScript content using a browser
	Render  ExtractAsyncParamsRenderUnion `json:"render,omitzero"`
	Session ExtractAsyncParamsSession     `json:"session,omitzero"`
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
	ExtractAsyncParamsDriverAuto     ExtractAsyncParamsDriver = "auto"
	ExtractAsyncParamsDriverVx6      ExtractAsyncParamsDriver = "vx6"
	ExtractAsyncParamsDriverVx8      ExtractAsyncParamsDriver = "vx8"
	ExtractAsyncParamsDriverVx8Pro   ExtractAsyncParamsDriver = "vx8-pro"
	ExtractAsyncParamsDriverVx10     ExtractAsyncParamsDriver = "vx10"
	ExtractAsyncParamsDriverVx10Pro  ExtractAsyncParamsDriver = "vx10-pro"
	ExtractAsyncParamsDriverVx12     ExtractAsyncParamsDriver = "vx12"
	ExtractAsyncParamsDriverVx12Pro  ExtractAsyncParamsDriver = "vx12-pro"
	ExtractAsyncParamsDriverMediaVx6 ExtractAsyncParamsDriver = "media-vx6"
	ExtractAsyncParamsDriverFastVx6  ExtractAsyncParamsDriver = "fast-vx6"
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

// Selects which markdown conversion strategy to use. "full_page" converts the
// entire HTML page. "main_content" uses Mozilla Readability to extract the main
// article content before converting.
type ExtractAsyncParamsMarkdownBackend string

const (
	ExtractAsyncParamsMarkdownBackendFullPage    ExtractAsyncParamsMarkdownBackend = "full_page"
	ExtractAsyncParamsMarkdownBackendMainContent ExtractAsyncParamsMarkdownBackend = "main_content"
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractAsyncParamsRenderUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Construct this variant with constant.ValueOf[constant.Auto]()
	OfAuto constant.Auto `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractAsyncParamsRenderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfAuto)
}
func (u *ExtractAsyncParamsRenderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractAsyncParamsRenderUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAuto) {
		return &u.OfAuto
	}
	return nil
}

type ExtractAsyncParamsSession struct {
	ID                  param.Opt[string]  `json:"id,omitzero"`
	PrefetchUserbrowser param.Opt[bool]    `json:"prefetch_userbrowser,omitzero"`
	RenewOnBlocked      param.Opt[bool]    `json:"renew_on_blocked,omitzero"`
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
	// Overall deadline in milliseconds for a realtime request. Clamped to the account
	// total timeout — it can shorten the deadline but never extend it. Has no effect
	// on async requests.
	RealtimeTotalTimeout param.Opt[float64] `json:"realtime_total_timeout,omitzero"`
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
	// Custom flow for the optimization engine: maps candidate names to the number of
	// attempts to spend on each candidate before advancing (0 skips it). Key order
	// defines the flow order. Providing it opts the request into 'auto' driver
	// selection.
	AutoDriverConfiguration map[string]int64 `json:"auto_driver_configuration,omitzero"`
	// Request body for POST, PUT, PATCH methods
	Body any `json:"body,omitzero"`
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
	// Browser driver to use. Use 'auto' to let the engine select the candidate config
	// per domain.
	//
	// Any of "auto", "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro",
	// "media-vx6", "fast-vx6".
	Driver string `json:"driver,omitzero"`
	// Expected HTTP status codes for successful requests
	ExpectedStatusCodes []int64 `json:"expected_status_codes,omitzero"`
	// List of acceptable response formats in order of preference
	//
	// Any of "html", "markdown", "screenshot", "headers", "links".
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
	// Selects which markdown conversion strategy to use. "full_page" converts the
	// entire HTML page. "main_content" uses Mozilla Readability to extract the main
	// article content before converting.
	//
	// Any of "full_page", "main_content".
	MarkdownBackend string `json:"markdown_backend,omitzero"`
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
	// Whether to render JavaScript content using a browser. Use 'auto' to let the
	// engine select the candidate config per domain.
	Render  ExtractBatchParamsInputRenderUnion `json:"render,omitzero"`
	Session ExtractBatchParamsInputSession     `json:"session,omitzero"`
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
		"driver", "auto", "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro", "media-vx6", "fast-vx6",
	)
	apijson.RegisterFieldValidator[ExtractBatchParamsInput](
		"markdown_backend", "full_page", "main_content",
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsInputRenderUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Construct this variant with constant.ValueOf[constant.Auto]()
	OfAuto constant.Auto `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsInputRenderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfAuto)
}
func (u *ExtractBatchParamsInputRenderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsInputRenderUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAuto) {
		return &u.OfAuto
	}
	return nil
}

type ExtractBatchParamsInputSession struct {
	ID                  param.Opt[string]  `json:"id,omitzero"`
	PrefetchUserbrowser param.Opt[bool]    `json:"prefetch_userbrowser,omitzero"`
	RenewOnBlocked      param.Opt[bool]    `json:"renew_on_blocked,omitzero"`
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
	// Overall deadline in milliseconds for a realtime request. Clamped to the account
	// total timeout — it can shorten the deadline but never extend it. Has no effect
	// on async requests.
	RealtimeTotalTimeout param.Opt[float64] `json:"realtime_total_timeout,omitzero"`
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
	// Custom flow for the optimization engine: maps candidate names to the number of
	// attempts to spend on each candidate before advancing (0 skips it). Key order
	// defines the flow order. Providing it opts the request into 'auto' driver
	// selection.
	AutoDriverConfiguration map[string]int64 `json:"auto_driver_configuration,omitzero"`
	// Request body for POST, PUT, PATCH methods
	Body any `json:"body,omitzero"`
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
	// Browser driver to use. Use 'auto' to let the engine select the candidate config
	// per domain.
	//
	// Any of "auto", "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro",
	// "media-vx6", "fast-vx6".
	Driver string `json:"driver,omitzero"`
	// Expected HTTP status codes for successful requests
	ExpectedStatusCodes []int64 `json:"expected_status_codes,omitzero"`
	// List of acceptable response formats in order of preference
	//
	// Any of "html", "markdown", "screenshot", "headers", "links".
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
	// Selects which markdown conversion strategy to use. "full_page" converts the
	// entire HTML page. "main_content" uses Mozilla Readability to extract the main
	// article content before converting.
	//
	// Any of "full_page", "main_content".
	MarkdownBackend string `json:"markdown_backend,omitzero"`
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
	// Whether to render JavaScript content using a browser. Use 'auto' to let the
	// engine select the candidate config per domain.
	Render  ExtractBatchParamsSharedInputsRenderUnion `json:"render,omitzero"`
	Session ExtractBatchParamsSharedInputsSession     `json:"session,omitzero"`
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
		"driver", "auto", "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro", "media-vx6", "fast-vx6",
	)
	apijson.RegisterFieldValidator[ExtractBatchParamsSharedInputs](
		"markdown_backend", "full_page", "main_content",
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractBatchParamsSharedInputsRenderUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Construct this variant with constant.ValueOf[constant.Auto]()
	OfAuto constant.Auto `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractBatchParamsSharedInputsRenderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfAuto)
}
func (u *ExtractBatchParamsSharedInputsRenderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractBatchParamsSharedInputsRenderUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAuto) {
		return &u.OfAuto
	}
	return nil
}

type ExtractBatchParamsSharedInputsSession struct {
	ID                  param.Opt[string]  `json:"id,omitzero"`
	PrefetchUserbrowser param.Opt[bool]    `json:"prefetch_userbrowser,omitzero"`
	RenewOnBlocked      param.Opt[bool]    `json:"renew_on_blocked,omitzero"`
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

type ExtractRunParams struct {
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
	// Overall deadline in milliseconds for a realtime request. Clamped to the account
	// total timeout — it can shorten the deadline but never extend it. Has no effect
	// on async requests.
	RealtimeTotalTimeout param.Opt[float64] `json:"realtime_total_timeout,omitzero"`
	// Request timeout in milliseconds
	RequestTimeout param.Opt[float64] `json:"request_timeout,omitzero"`
	// User-defined tag for request identification
	Tag param.Opt[string] `json:"tag,omitzero"`
	// Custom flow for the optimization engine: maps candidate names to the number of
	// attempts to spend on each candidate before advancing (0 skips it). Key order
	// defines the flow order. Providing it opts the request into 'auto' driver
	// selection.
	AutoDriverConfiguration map[string]int64 `json:"auto_driver_configuration,omitzero"`
	// Request body for POST, PUT, PATCH methods
	Body any `json:"body,omitzero"`
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
	// Any of "auto", "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro",
	// "media-vx6", "fast-vx6".
	Driver ExtractRunParamsDriver `json:"driver,omitzero"`
	// Expected HTTP status codes for successful requests
	ExpectedStatusCodes []int64 `json:"expected_status_codes,omitzero"`
	// List of acceptable response formats in order of preference
	//
	// Any of "html", "markdown", "screenshot", "headers", "links".
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
	// Selects which markdown conversion strategy to use. "full_page" converts the
	// entire HTML page. "main_content" uses Mozilla Readability to extract the main
	// article content before converting.
	//
	// Any of "full_page", "main_content".
	MarkdownBackend ExtractRunParamsMarkdownBackend `json:"markdown_backend,omitzero"`
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
	// Whether to render JavaScript content using a browser
	Render  ExtractRunParamsRenderUnion `json:"render,omitzero"`
	Session ExtractRunParamsSession     `json:"session,omitzero"`
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
	Name string `json:"name,omitzero" api:"required"`
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

func (u ExtractRunParamsBrowserActionUnion) MarshalJSON() ([]byte, error) {
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
func (u *ExtractRunParamsBrowserActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsBrowserActionUnion) asAny() any {
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
	ExtractRunParamsDriverAuto     ExtractRunParamsDriver = "auto"
	ExtractRunParamsDriverVx6      ExtractRunParamsDriver = "vx6"
	ExtractRunParamsDriverVx8      ExtractRunParamsDriver = "vx8"
	ExtractRunParamsDriverVx8Pro   ExtractRunParamsDriver = "vx8-pro"
	ExtractRunParamsDriverVx10     ExtractRunParamsDriver = "vx10"
	ExtractRunParamsDriverVx10Pro  ExtractRunParamsDriver = "vx10-pro"
	ExtractRunParamsDriverVx12     ExtractRunParamsDriver = "vx12"
	ExtractRunParamsDriverVx12Pro  ExtractRunParamsDriver = "vx12-pro"
	ExtractRunParamsDriverMediaVx6 ExtractRunParamsDriver = "media-vx6"
	ExtractRunParamsDriverFastVx6  ExtractRunParamsDriver = "fast-vx6"
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

// Selects which markdown conversion strategy to use. "full_page" converts the
// entire HTML page. "main_content" uses Mozilla Readability to extract the main
// article content before converting.
type ExtractRunParamsMarkdownBackend string

const (
	ExtractRunParamsMarkdownBackendFullPage    ExtractRunParamsMarkdownBackend = "full_page"
	ExtractRunParamsMarkdownBackendMainContent ExtractRunParamsMarkdownBackend = "main_content"
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
	Value string `json:"value" api:"required"`
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractRunParamsRenderUnion struct {
	OfBool param.Opt[bool] `json:",omitzero,inline"`
	// Construct this variant with constant.ValueOf[constant.Auto]()
	OfAuto constant.Auto `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractRunParamsRenderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfAuto)
}
func (u *ExtractRunParamsRenderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractRunParamsRenderUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfAuto) {
		return &u.OfAuto
	}
	return nil
}

type ExtractRunParamsSession struct {
	ID                  param.Opt[string]  `json:"id,omitzero"`
	PrefetchUserbrowser param.Opt[bool]    `json:"prefetch_userbrowser,omitzero"`
	RenewOnBlocked      param.Opt[bool]    `json:"renew_on_blocked,omitzero"`
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
