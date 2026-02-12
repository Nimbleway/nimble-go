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
func (r *ExtractService) Extract(ctx context.Context, body ExtractExtractParams, opts ...option.RequestOption) (res *ExtractExtractResponse, err error) {
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

type ExtractExtractResponse struct {
	Data     ExtractExtractResponseData     `json:"data,required"`
	Metadata ExtractExtractResponseMetadata `json:"metadata,required"`
	// The status of the task.
	//
	// Any of "success", "skipped", "fatal", "error", "postponed", "ignored",
	// "rejected", "blocked".
	Status ExtractExtractResponseStatus `json:"status,required"`
	// Unique identifier for the task.
	TaskID string `json:"task_id,required"`
	// The final URL.
	URL   string                      `json:"url,required"`
	Debug ExtractExtractResponseDebug `json:"debug"`
	// Pagination information if applicable.
	Pagination ExtractExtractResponsePaginationUnion `json:"pagination"`
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
func (r ExtractExtractResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractExtractResponseData struct {
	// Browser actions execution results. Present only when browser_actions were
	// specified in the request.
	BrowserActions ExtractExtractResponseDataBrowserActions `json:"browser_actions"`
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
	NetworkCapture []ExtractExtractResponseDataNetworkCapture `json:"network_capture"`
	// The parsing results extracted from the HTML & network content.
	Parsing ExtractExtractResponseDataParsingUnion `json:"parsing"`
	// The list of redirects that occurred during the task.
	Redirects []ExtractExtractResponseDataRedirect `json:"redirects"`
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
func (r ExtractExtractResponseData) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser actions execution results. Present only when browser_actions were
// specified in the request.
type ExtractExtractResponseDataBrowserActions struct {
	Results       []ExtractExtractResponseDataBrowserActionsResult `json:"results,required"`
	Success       bool                                             `json:"success,required"`
	TotalDuration float64                                          `json:"total_duration,required"`
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
func (r ExtractExtractResponseDataBrowserActions) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponseDataBrowserActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractExtractResponseDataBrowserActionsResult struct {
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
func (r ExtractExtractResponseDataBrowserActionsResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponseDataBrowserActionsResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractExtractResponseDataNetworkCapture struct {
	Filter       ExtractExtractResponseDataNetworkCaptureFilter   `json:"filter,required"`
	Results      []ExtractExtractResponseDataNetworkCaptureResult `json:"results,required"`
	ErrorMessage string                                           `json:"errorMessage"`
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
func (r ExtractExtractResponseDataNetworkCapture) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponseDataNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractExtractResponseDataNetworkCaptureFilter struct {
	Validation           bool    `json:"validation,required"`
	WaitForRequestsCount float64 `json:"wait_for_requests_count,required"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method"`
	// Resource type for network capture filtering
	ResourceType                ExtractExtractResponseDataNetworkCaptureFilterResourceTypeUnion `json:"resource_type"`
	StatusCode                  ExtractExtractResponseDataNetworkCaptureFilterStatusCodeUnion   `json:"status_code"`
	URL                         ExtractExtractResponseDataNetworkCaptureFilterURL               `json:"url"`
	WaitForRequestsCountTimeout float64                                                         `json:"wait_for_requests_count_timeout"`
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
func (r ExtractExtractResponseDataNetworkCaptureFilter) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponseDataNetworkCaptureFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractExtractResponseDataNetworkCaptureFilterResourceTypeUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfExtractExtractResponseDataNetworkCaptureFilterResourceTypeString
// OfExtractExtractResponseDataNetworkCaptureFilterResourceTypeArrayItemArray]
type ExtractExtractResponseDataNetworkCaptureFilterResourceTypeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfExtractExtractResponseDataNetworkCaptureFilterResourceTypeString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfExtractExtractResponseDataNetworkCaptureFilterResourceTypeArrayItemArray []string `json:",inline"`
	JSON                                                                       struct {
		OfExtractExtractResponseDataNetworkCaptureFilterResourceTypeString         respjson.Field
		OfExtractExtractResponseDataNetworkCaptureFilterResourceTypeArrayItemArray respjson.Field
		raw                                                                        string
	} `json:"-"`
}

func (u ExtractExtractResponseDataNetworkCaptureFilterResourceTypeUnion) AsExtractExtractResponseDataNetworkCaptureFilterResourceTypeString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractExtractResponseDataNetworkCaptureFilterResourceTypeUnion) AsExtractExtractResponseDataNetworkCaptureFilterResourceTypeArrayItemArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractExtractResponseDataNetworkCaptureFilterResourceTypeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ExtractExtractResponseDataNetworkCaptureFilterResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type for network capture filtering
type ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString string

const (
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringDocument           ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "document"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringStylesheet         ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "stylesheet"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringImage              ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "image"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringMedia              ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "media"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringFont               ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "font"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringScript             ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "script"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringTexttrack          ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "texttrack"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringXhr                ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "xhr"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringFetch              ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "fetch"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringPrefetch           ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "prefetch"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringEventsource        ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "eventsource"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringWebsocket          ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "websocket"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringManifest           ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "manifest"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringSignedexchange     ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "signedexchange"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringPing               ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "ping"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringCspviolationreport ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "cspviolationreport"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringPreflight          ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "preflight"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringOther              ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "other"
	ExtractExtractResponseDataNetworkCaptureFilterResourceTypeStringFedcm              ExtractExtractResponseDataNetworkCaptureFilterResourceTypeString = "fedcm"
)

// ExtractExtractResponseDataNetworkCaptureFilterStatusCodeUnion contains all
// possible properties and values from [float64], [[]float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfFloatArray]
type ExtractExtractResponseDataNetworkCaptureFilterStatusCodeUnion struct {
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

func (u ExtractExtractResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractExtractResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloatArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractExtractResponseDataNetworkCaptureFilterStatusCodeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ExtractExtractResponseDataNetworkCaptureFilterStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractExtractResponseDataNetworkCaptureFilterURL struct {
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
func (r ExtractExtractResponseDataNetworkCaptureFilterURL) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponseDataNetworkCaptureFilterURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractExtractResponseDataNetworkCaptureResult struct {
	Request  ExtractExtractResponseDataNetworkCaptureResultRequest  `json:"request,required"`
	Response ExtractExtractResponseDataNetworkCaptureResultResponse `json:"response,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Request     respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractExtractResponseDataNetworkCaptureResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponseDataNetworkCaptureResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractExtractResponseDataNetworkCaptureResultRequest struct {
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
func (r ExtractExtractResponseDataNetworkCaptureResultRequest) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponseDataNetworkCaptureResultRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractExtractResponseDataNetworkCaptureResultResponse struct {
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
func (r ExtractExtractResponseDataNetworkCaptureResultResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponseDataNetworkCaptureResultResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractExtractResponseDataParsingUnion contains all possible properties and
// values from [ExtractExtractResponseDataParsingParsingSuccessResult],
// [ExtractExtractResponseDataParsingParsingErrorResult], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractExtractResponseDataParsingMapItem]
type ExtractExtractResponseDataParsingUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExtractExtractResponseDataParsingMapItem any `json:",inline"`
	// This field is from variant
	// [ExtractExtractResponseDataParsingParsingSuccessResult].
	Entities map[string]any `json:"entities"`
	Status   string         `json:"status"`
	// This field is from variant
	// [ExtractExtractResponseDataParsingParsingErrorResult].
	Error string `json:"error"`
	JSON  struct {
		OfExtractExtractResponseDataParsingMapItem respjson.Field
		Entities                                   respjson.Field
		Status                                     respjson.Field
		Error                                      respjson.Field
		raw                                        string
	} `json:"-"`
}

func (u ExtractExtractResponseDataParsingUnion) AsExtractExtractResponseDataParsingParsingSuccessResult() (v ExtractExtractResponseDataParsingParsingSuccessResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractExtractResponseDataParsingUnion) AsExtractExtractResponseDataParsingParsingErrorResult() (v ExtractExtractResponseDataParsingParsingErrorResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractExtractResponseDataParsingUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractExtractResponseDataParsingUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractExtractResponseDataParsingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractExtractResponseDataParsingParsingSuccessResult struct {
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
func (r ExtractExtractResponseDataParsingParsingSuccessResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponseDataParsingParsingSuccessResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractExtractResponseDataParsingParsingErrorResult struct {
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
func (r ExtractExtractResponseDataParsingParsingErrorResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponseDataParsingParsingErrorResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractExtractResponseDataRedirect struct {
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
func (r ExtractExtractResponseDataRedirect) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponseDataRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractExtractResponseMetadata struct {
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
func (r ExtractExtractResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the task.
type ExtractExtractResponseStatus string

const (
	ExtractExtractResponseStatusSuccess   ExtractExtractResponseStatus = "success"
	ExtractExtractResponseStatusSkipped   ExtractExtractResponseStatus = "skipped"
	ExtractExtractResponseStatusFatal     ExtractExtractResponseStatus = "fatal"
	ExtractExtractResponseStatusError     ExtractExtractResponseStatus = "error"
	ExtractExtractResponseStatusPostponed ExtractExtractResponseStatus = "postponed"
	ExtractExtractResponseStatusIgnored   ExtractExtractResponseStatus = "ignored"
	ExtractExtractResponseStatusRejected  ExtractExtractResponseStatus = "rejected"
	ExtractExtractResponseStatusBlocked   ExtractExtractResponseStatus = "blocked"
)

type ExtractExtractResponseDebug struct {
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
func (r ExtractExtractResponseDebug) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponseDebug) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractExtractResponsePaginationUnion contains all possible properties and
// values from [ExtractExtractResponsePaginationNextPageParams],
// [[]ExtractExtractResponsePaginationArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractExtractResponsePaginationArray]
type ExtractExtractResponsePaginationUnion struct {
	// This field will be present if the value is a
	// [[]ExtractExtractResponsePaginationArrayItem] instead of an object.
	OfExtractExtractResponsePaginationArray []ExtractExtractResponsePaginationArrayItem `json:",inline"`
	// This field is from variant [ExtractExtractResponsePaginationNextPageParams].
	NextPageParams map[string]any `json:"next_page_params"`
	JSON           struct {
		OfExtractExtractResponsePaginationArray respjson.Field
		NextPageParams                          respjson.Field
		raw                                     string
	} `json:"-"`
}

func (u ExtractExtractResponsePaginationUnion) AsExtractExtractResponsePaginationNextPageParams() (v ExtractExtractResponsePaginationNextPageParams) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractExtractResponsePaginationUnion) AsExtractExtractResponsePaginationArray() (v []ExtractExtractResponsePaginationArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractExtractResponsePaginationUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractExtractResponsePaginationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractExtractResponsePaginationNextPageParams struct {
	NextPageParams map[string]any `json:"next_page_params,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractExtractResponsePaginationNextPageParams) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponsePaginationNextPageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractExtractResponsePaginationArrayItem struct {
	NextPageParams map[string]any `json:"next_page_params,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractExtractResponsePaginationArrayItem) RawJSON() string { return r.JSON.raw }
func (r *ExtractExtractResponsePaginationArrayItem) UnmarshalJSON(data []byte) error {
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

type ExtractExtractParams struct {
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
	Browser ExtractExtractParamsBrowserUnion `json:"browser,omitzero"`
	// Array of browser automation actions to execute sequentially
	BrowserActions []ExtractExtractParamsBrowserActionUnion `json:"browser_actions,omitzero"`
	// Browser cookies as array of cookie objects
	Cookies ExtractExtractParamsCookiesUnion `json:"cookies,omitzero"`
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
	Country ExtractExtractParamsCountry `json:"country,omitzero"`
	// Device type for browser emulation
	//
	// Any of "desktop", "mobile", "tablet".
	Device ExtractExtractParamsDevice `json:"device,omitzero"`
	// Browser driver to use
	//
	// Any of "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro".
	Driver ExtractExtractParamsDriver `json:"driver,omitzero"`
	// Expected HTTP status codes for successful requests
	ExpectedStatusCodes []int64 `json:"expected_status_codes,omitzero"`
	// List of acceptable response formats in order of preference
	//
	// Any of "html", "markdown".
	Formats []string `json:"formats,omitzero"`
	// Custom HTTP headers to include in the request
	Headers map[string]ExtractExtractParamsHeaderUnion `json:"headers,omitzero"`
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
	Locale ExtractExtractParamsLocale `json:"locale,omitzero"`
	// HTTP method for the request
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method ExtractExtractParamsMethod `json:"method,omitzero"`
	// Filters for capturing network traffic
	NetworkCapture []ExtractExtractParamsNetworkCapture `json:"network_capture,omitzero"`
	// Operating system to emulate
	//
	// Any of "windows", "mac os", "linux", "android", "ios".
	Os ExtractExtractParamsOs `json:"os,omitzero"`
	// Custom parser configuration as a key-value map
	Parser ExtractExtractParamsParserUnion `json:"parser,omitzero"`
	// Referrer policy for the request
	//
	// Any of "random", "no-referer", "same-origin", "google", "bing", "facebook",
	// "twitter", "instagram".
	ReferrerType ExtractExtractParamsReferrerType `json:"referrer_type,omitzero"`
	Session      ExtractExtractParamsSession      `json:"session,omitzero"`
	// Skills or capabilities required for the request
	Skill ExtractExtractParamsSkillUnion `json:"skill,omitzero"`
	// US state for geolocation (only valid when country is US)
	//
	// Any of "AL", "AK", "AS", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA",
	// "GU", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI",
	// "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "MP",
	// "OH", "OK", "OR", "PA", "PR", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA",
	// "VI", "WA", "WV", "WI", "WY".
	State ExtractExtractParamsState `json:"state,omitzero"`
	paramObj
}

func (r ExtractExtractParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserString)
	OfExtractExtractsBrowserString param.Opt[string]                  `json:",omitzero,inline"`
	OfExtractExtractsBrowserObject *ExtractExtractParamsBrowserObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserString, u.OfExtractExtractsBrowserObject)
}
func (u *ExtractExtractParamsBrowserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserString) {
		return &u.OfExtractExtractsBrowserString
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserObject) {
		return u.OfExtractExtractsBrowserObject
	}
	return nil
}

// Browser type to emulate
type ExtractExtractParamsBrowserString string

const (
	ExtractExtractParamsBrowserStringChrome  ExtractExtractParamsBrowserString = "chrome"
	ExtractExtractParamsBrowserStringFirefox ExtractExtractParamsBrowserString = "firefox"
)

// The property Name is required.
type ExtractExtractParamsBrowserObject struct {
	// Any of "chrome", "firefox".
	Name string `json:"name,omitzero,required"`
	// Specific browser version to emulate
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r ExtractExtractParamsBrowserObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractExtractParamsBrowserObject](
		"name", "chrome", "firefox",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionUnion struct {
	OfExtractExtractsBrowserActionAutoScrollAction        *ExtractExtractParamsBrowserActionAutoScrollAction        `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionClickAction             *ExtractExtractParamsBrowserActionClickAction             `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionEvalAction              *ExtractExtractParamsBrowserActionEvalAction              `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionFetchAction             *ExtractExtractParamsBrowserActionFetchAction             `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionFillAction              *ExtractExtractParamsBrowserActionFillAction              `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionGetCookiesAction        *ExtractExtractParamsBrowserActionGetCookiesAction        `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionGotoAction              *ExtractExtractParamsBrowserActionGotoAction              `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionPressAction             *ExtractExtractParamsBrowserActionPressAction             `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionScreenshotAction        *ExtractExtractParamsBrowserActionScreenshotAction        `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionScrollAction            *ExtractExtractParamsBrowserActionScrollAction            `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionWaitAction              *ExtractExtractParamsBrowserActionWaitAction              `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionWaitForElementAction    *ExtractExtractParamsBrowserActionWaitForElementAction    `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionWaitForNavigationAction *ExtractExtractParamsBrowserActionWaitForNavigationAction `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionAutoScrollAction,
		u.OfExtractExtractsBrowserActionClickAction,
		u.OfExtractExtractsBrowserActionEvalAction,
		u.OfExtractExtractsBrowserActionFetchAction,
		u.OfExtractExtractsBrowserActionFillAction,
		u.OfExtractExtractsBrowserActionGetCookiesAction,
		u.OfExtractExtractsBrowserActionGotoAction,
		u.OfExtractExtractsBrowserActionPressAction,
		u.OfExtractExtractsBrowserActionScreenshotAction,
		u.OfExtractExtractsBrowserActionScrollAction,
		u.OfExtractExtractsBrowserActionWaitAction,
		u.OfExtractExtractsBrowserActionWaitForElementAction,
		u.OfExtractExtractsBrowserActionWaitForNavigationAction)
}
func (u *ExtractExtractParamsBrowserActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionAutoScrollAction) {
		return u.OfExtractExtractsBrowserActionAutoScrollAction
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionClickAction) {
		return u.OfExtractExtractsBrowserActionClickAction
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionEvalAction) {
		return u.OfExtractExtractsBrowserActionEvalAction
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionFetchAction) {
		return u.OfExtractExtractsBrowserActionFetchAction
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionFillAction) {
		return u.OfExtractExtractsBrowserActionFillAction
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionGetCookiesAction) {
		return u.OfExtractExtractsBrowserActionGetCookiesAction
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionGotoAction) {
		return u.OfExtractExtractsBrowserActionGotoAction
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionPressAction) {
		return u.OfExtractExtractsBrowserActionPressAction
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionScreenshotAction) {
		return u.OfExtractExtractsBrowserActionScreenshotAction
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionScrollAction) {
		return u.OfExtractExtractsBrowserActionScrollAction
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionWaitAction) {
		return u.OfExtractExtractsBrowserActionWaitAction
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionWaitForElementAction) {
		return u.OfExtractExtractsBrowserActionWaitForElementAction
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionWaitForNavigationAction) {
		return u.OfExtractExtractsBrowserActionWaitForNavigationAction
	}
	return nil
}

// Continuously scroll to load dynamic content
//
// The property AutoScroll is required.
type ExtractExtractParamsBrowserActionAutoScrollAction struct {
	AutoScroll ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollUnion `json:"auto_scroll,omitzero,required"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionAutoScrollAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionAutoScrollAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionAutoScrollAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollUnion struct {
	OfBool                                                         param.Opt[bool]                                                    `json:",omitzero,inline"`
	OfFloat                                                        param.Opt[float64]                                                 `json:",omitzero,inline"`
	OfString                                                       param.Opt[string]                                                  `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionAutoScrollActionAutoScrollObject *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfFloat, u.OfString, u.OfExtractExtractsBrowserActionAutoScrollActionAutoScrollObject)
}
func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionAutoScrollActionAutoScrollObject) {
		return u.OfExtractExtractsBrowserActionAutoScrollActionAutoScrollObject
	}
	return nil
}

type ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObject struct {
	StepSize param.Opt[float64] `json:"step_size,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	ClickSelector ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion `json:"click_selector,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Container ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion `json:"container,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	DelayAfterScroll ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion `json:"delay_after_scroll,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	IdleTimeout ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion `json:"idle_timeout,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	LoadingSelector ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion `json:"loading_selector,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	MaxDuration ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion `json:"max_duration,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	PauseOnSelector ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion `json:"pause_on_selector,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectClickSelectorUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectContainerUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectDelayAfterScrollUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectIdleTimeoutUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectLoadingSelectorUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectMaxDurationUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectPauseOnSelectorUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionAutoScrollActionAutoScrollObjectRequiredString)
	OfExtractExtractsBrowserActionAutoScrollActionAutoScrollObjectRequiredString param.Opt[ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                       param.Opt[bool]                                                                            `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionAutoScrollActionAutoScrollObjectRequiredString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionAutoScrollActionAutoScrollObjectRequiredString) {
		return &u.OfExtractExtractsBrowserActionAutoScrollActionAutoScrollObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredString string

const (
	ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredStringTrue  ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredString = "true"
	ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredStringFalse ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionAutoScrollActionAutoScrollObjectSkipString)
	OfExtractExtractsBrowserActionAutoScrollActionAutoScrollObjectSkipString param.Opt[ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                   param.Opt[bool]                                                                        `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionAutoScrollActionAutoScrollObjectSkipString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionAutoScrollActionAutoScrollObjectSkipString) {
		return &u.OfExtractExtractsBrowserActionAutoScrollActionAutoScrollObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipString string

const (
	ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipStringTrue  ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipString = "true"
	ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipStringFalse ExtractExtractParamsBrowserActionAutoScrollActionAutoScrollObjectSkipString = "false"
)

// Click on an element by selector
//
// The property Click is required.
type ExtractExtractParamsBrowserActionClickAction struct {
	Click ExtractExtractParamsBrowserActionClickActionClickUnion `json:"click,omitzero,required"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionClickAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionClickAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionClickAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionClickActionClickUnion struct {
	OfString                                             param.Opt[string]                                        `json:",omitzero,inline"`
	OfStringArray                                        []string                                                 `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionClickActionClickObject *ExtractExtractParamsBrowserActionClickActionClickObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionClickActionClickUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray, u.OfExtractExtractsBrowserActionClickActionClickObject)
}
func (u *ExtractExtractParamsBrowserActionClickActionClickUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionClickActionClickUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionClickActionClickObject) {
		return u.OfExtractExtractsBrowserActionClickActionClickObject
	}
	return nil
}

// The property Selector is required.
type ExtractExtractParamsBrowserActionClickActionClickObject struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector ExtractExtractParamsBrowserActionClickActionClickObjectSelectorUnion `json:"selector,omitzero,required"`
	Count    param.Opt[float64]                                                   `json:"count,omitzero"`
	OffsetX  param.Opt[int64]                                                     `json:"offset_x,omitzero"`
	OffsetY  param.Opt[int64]                                                     `json:"offset_y,omitzero"`
	Scroll   param.Opt[bool]                                                      `json:"scroll,omitzero"`
	Steps    param.Opt[float64]                                                   `json:"steps,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay ExtractExtractParamsBrowserActionClickActionClickObjectDelayUnion `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractExtractParamsBrowserActionClickActionClickObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractExtractParamsBrowserActionClickActionClickObjectSkipUnion `json:"skip,omitzero"`
	// Any of "linear", "ghost-cursor", "windmouse".
	Strategy string `json:"strategy,omitzero"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionClickActionClickObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionClickActionClickObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionClickActionClickObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractExtractParamsBrowserActionClickActionClickObject](
		"strategy", "linear", "ghost-cursor", "windmouse",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionClickActionClickObjectSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionClickActionClickObjectSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractExtractParamsBrowserActionClickActionClickObjectSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionClickActionClickObjectSelectorUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionClickActionClickObjectDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionClickActionClickObjectDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractExtractParamsBrowserActionClickActionClickObjectDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionClickActionClickObjectDelayUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionClickActionClickObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionClickActionClickObjectRequiredString)
	OfExtractExtractsBrowserActionClickActionClickObjectRequiredString param.Opt[ExtractExtractParamsBrowserActionClickActionClickObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                             param.Opt[bool]                                                                  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionClickActionClickObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionClickActionClickObjectRequiredString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionClickActionClickObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionClickActionClickObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionClickActionClickObjectRequiredString) {
		return &u.OfExtractExtractsBrowserActionClickActionClickObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionClickActionClickObjectRequiredString string

const (
	ExtractExtractParamsBrowserActionClickActionClickObjectRequiredStringTrue  ExtractExtractParamsBrowserActionClickActionClickObjectRequiredString = "true"
	ExtractExtractParamsBrowserActionClickActionClickObjectRequiredStringFalse ExtractExtractParamsBrowserActionClickActionClickObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionClickActionClickObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionClickActionClickObjectSkipString)
	OfExtractExtractsBrowserActionClickActionClickObjectSkipString param.Opt[ExtractExtractParamsBrowserActionClickActionClickObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                         param.Opt[bool]                                                              `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionClickActionClickObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionClickActionClickObjectSkipString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionClickActionClickObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionClickActionClickObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionClickActionClickObjectSkipString) {
		return &u.OfExtractExtractsBrowserActionClickActionClickObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionClickActionClickObjectSkipString string

const (
	ExtractExtractParamsBrowserActionClickActionClickObjectSkipStringTrue  ExtractExtractParamsBrowserActionClickActionClickObjectSkipString = "true"
	ExtractExtractParamsBrowserActionClickActionClickObjectSkipStringFalse ExtractExtractParamsBrowserActionClickActionClickObjectSkipString = "false"
)

// Execute JavaScript code in page context
//
// The property Eval is required.
type ExtractExtractParamsBrowserActionEvalAction struct {
	Eval ExtractExtractParamsBrowserActionEvalActionEvalUnion `json:"eval,omitzero,required"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionEvalAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionEvalAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionEvalAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionEvalActionEvalUnion struct {
	OfString                                           param.Opt[string]                                      `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionEvalActionEvalObject *ExtractExtractParamsBrowserActionEvalActionEvalObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionEvalActionEvalUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractExtractsBrowserActionEvalActionEvalObject)
}
func (u *ExtractExtractParamsBrowserActionEvalActionEvalUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionEvalActionEvalUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionEvalActionEvalObject) {
		return u.OfExtractExtractsBrowserActionEvalActionEvalObject
	}
	return nil
}

// The property Code is required.
type ExtractExtractParamsBrowserActionEvalActionEvalObject struct {
	Code string `json:"code,required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractExtractParamsBrowserActionEvalActionEvalObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractExtractParamsBrowserActionEvalActionEvalObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionEvalActionEvalObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionEvalActionEvalObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionEvalActionEvalObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionEvalActionEvalObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionEvalActionEvalObjectRequiredString)
	OfExtractExtractsBrowserActionEvalActionEvalObjectRequiredString param.Opt[ExtractExtractParamsBrowserActionEvalActionEvalObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                           param.Opt[bool]                                                                `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionEvalActionEvalObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionEvalActionEvalObjectRequiredString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionEvalActionEvalObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionEvalActionEvalObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionEvalActionEvalObjectRequiredString) {
		return &u.OfExtractExtractsBrowserActionEvalActionEvalObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionEvalActionEvalObjectRequiredString string

const (
	ExtractExtractParamsBrowserActionEvalActionEvalObjectRequiredStringTrue  ExtractExtractParamsBrowserActionEvalActionEvalObjectRequiredString = "true"
	ExtractExtractParamsBrowserActionEvalActionEvalObjectRequiredStringFalse ExtractExtractParamsBrowserActionEvalActionEvalObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionEvalActionEvalObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionEvalActionEvalObjectSkipString)
	OfExtractExtractsBrowserActionEvalActionEvalObjectSkipString param.Opt[ExtractExtractParamsBrowserActionEvalActionEvalObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                       param.Opt[bool]                                                            `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionEvalActionEvalObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionEvalActionEvalObjectSkipString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionEvalActionEvalObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionEvalActionEvalObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionEvalActionEvalObjectSkipString) {
		return &u.OfExtractExtractsBrowserActionEvalActionEvalObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionEvalActionEvalObjectSkipString string

const (
	ExtractExtractParamsBrowserActionEvalActionEvalObjectSkipStringTrue  ExtractExtractParamsBrowserActionEvalActionEvalObjectSkipString = "true"
	ExtractExtractParamsBrowserActionEvalActionEvalObjectSkipStringFalse ExtractExtractParamsBrowserActionEvalActionEvalObjectSkipString = "false"
)

// Make an HTTP request in browser context
//
// The property Fetch is required.
type ExtractExtractParamsBrowserActionFetchAction struct {
	Fetch ExtractExtractParamsBrowserActionFetchActionFetchUnion `json:"fetch,omitzero,required" format:"uri"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionFetchAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionFetchAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionFetchAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionFetchActionFetchUnion struct {
	OfString                                             param.Opt[string]                                        `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionFetchActionFetchObject *ExtractExtractParamsBrowserActionFetchActionFetchObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionFetchActionFetchUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractExtractsBrowserActionFetchActionFetchObject)
}
func (u *ExtractExtractParamsBrowserActionFetchActionFetchUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionFetchActionFetchUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionFetchActionFetchObject) {
		return u.OfExtractExtractsBrowserActionFetchActionFetchObject
	}
	return nil
}

// The property URL is required.
type ExtractExtractParamsBrowserActionFetchActionFetchObject struct {
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
	Required ExtractExtractParamsBrowserActionFetchActionFetchObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractExtractParamsBrowserActionFetchActionFetchObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionFetchActionFetchObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionFetchActionFetchObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionFetchActionFetchObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractExtractParamsBrowserActionFetchActionFetchObject](
		"method", "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionFetchActionFetchObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionFetchActionFetchObjectRequiredString)
	OfExtractExtractsBrowserActionFetchActionFetchObjectRequiredString param.Opt[ExtractExtractParamsBrowserActionFetchActionFetchObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                             param.Opt[bool]                                                                  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionFetchActionFetchObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionFetchActionFetchObjectRequiredString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionFetchActionFetchObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionFetchActionFetchObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionFetchActionFetchObjectRequiredString) {
		return &u.OfExtractExtractsBrowserActionFetchActionFetchObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionFetchActionFetchObjectRequiredString string

const (
	ExtractExtractParamsBrowserActionFetchActionFetchObjectRequiredStringTrue  ExtractExtractParamsBrowserActionFetchActionFetchObjectRequiredString = "true"
	ExtractExtractParamsBrowserActionFetchActionFetchObjectRequiredStringFalse ExtractExtractParamsBrowserActionFetchActionFetchObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionFetchActionFetchObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionFetchActionFetchObjectSkipString)
	OfExtractExtractsBrowserActionFetchActionFetchObjectSkipString param.Opt[ExtractExtractParamsBrowserActionFetchActionFetchObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                         param.Opt[bool]                                                              `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionFetchActionFetchObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionFetchActionFetchObjectSkipString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionFetchActionFetchObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionFetchActionFetchObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionFetchActionFetchObjectSkipString) {
		return &u.OfExtractExtractsBrowserActionFetchActionFetchObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionFetchActionFetchObjectSkipString string

const (
	ExtractExtractParamsBrowserActionFetchActionFetchObjectSkipStringTrue  ExtractExtractParamsBrowserActionFetchActionFetchObjectSkipString = "true"
	ExtractExtractParamsBrowserActionFetchActionFetchObjectSkipStringFalse ExtractExtractParamsBrowserActionFetchActionFetchObjectSkipString = "false"
)

// Fill text into an input field
//
// The property Fill is required.
type ExtractExtractParamsBrowserActionFillAction struct {
	// Fill options with mode-specific fields. Use "type" mode for behavioral typing
	// simulation, or "paste" mode for instant paste.
	Fill ExtractExtractParamsBrowserActionFillActionFillUnion `json:"fill,omitzero,required"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionFillAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionFillAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionFillAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionFillActionFillUnion struct {
	OfType  *ExtractExtractParamsBrowserActionFillActionFillType  `json:",omitzero,inline"`
	OfPaste *ExtractExtractParamsBrowserActionFillActionFillPaste `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionFillActionFillUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfType, u.OfPaste)
}
func (u *ExtractExtractParamsBrowserActionFillActionFillUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionFillActionFillUnion) asAny() any {
	if !param.IsOmitted(u.OfType) {
		return u.OfType
	} else if !param.IsOmitted(u.OfPaste) {
		return u.OfPaste
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractExtractParamsBrowserActionFillActionFillUnion) GetMouseMovementStrategy() *string {
	if vt := u.OfType; vt != nil {
		return &vt.MouseMovementStrategy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractExtractParamsBrowserActionFillActionFillUnion) GetTypingInterval() *ExtractExtractParamsBrowserActionFillActionFillTypeTypingIntervalUnion {
	if vt := u.OfType; vt != nil {
		return &vt.TypingInterval
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractExtractParamsBrowserActionFillActionFillUnion) GetTypingStrategy() *string {
	if vt := u.OfType; vt != nil {
		return &vt.TypingStrategy
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractExtractParamsBrowserActionFillActionFillUnion) GetValue() *string {
	if vt := u.OfType; vt != nil {
		return (*string)(&vt.Value)
	} else if vt := u.OfPaste; vt != nil {
		return (*string)(&vt.Value)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractExtractParamsBrowserActionFillActionFillUnion) GetClickOnElement() *bool {
	if vt := u.OfType; vt != nil && vt.ClickOnElement.Valid() {
		return &vt.ClickOnElement.Value
	} else if vt := u.OfPaste; vt != nil && vt.ClickOnElement.Valid() {
		return &vt.ClickOnElement.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractExtractParamsBrowserActionFillActionFillUnion) GetMode() *string {
	if vt := u.OfType; vt != nil {
		return (*string)(&vt.Mode)
	} else if vt := u.OfPaste; vt != nil {
		return (*string)(&vt.Mode)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractExtractParamsBrowserActionFillActionFillUnion) GetScroll() *bool {
	if vt := u.OfType; vt != nil && vt.Scroll.Valid() {
		return &vt.Scroll.Value
	} else if vt := u.OfPaste; vt != nil && vt.Scroll.Valid() {
		return &vt.Scroll.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractExtractParamsBrowserActionFillActionFillUnion) GetTimeout() *float64 {
	if vt := u.OfType; vt != nil && vt.Timeout.Valid() {
		return &vt.Timeout.Value
	} else if vt := u.OfPaste; vt != nil && vt.Timeout.Valid() {
		return &vt.Timeout.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u ExtractExtractParamsBrowserActionFillActionFillUnion) GetVisible() *bool {
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
func (u ExtractExtractParamsBrowserActionFillActionFillUnion) GetSelector() (res extractExtractParamsBrowserActionFillActionFillUnionSelector) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Selector.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Selector.asAny()
	}
	return
}

// Can have the runtime types [*string], [\*[]string]
type extractExtractParamsBrowserActionFillActionFillUnionSelector struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *[]string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extractExtractParamsBrowserActionFillActionFillUnionSelector) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ExtractExtractParamsBrowserActionFillActionFillUnion) GetDelay() (res extractExtractParamsBrowserActionFillActionFillUnionDelay) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Delay.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Delay.asAny()
	}
	return
}

// Can have the runtime types [*float64], [*string]
type extractExtractParamsBrowserActionFillActionFillUnionDelay struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *float64:
//	case *string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extractExtractParamsBrowserActionFillActionFillUnionDelay) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ExtractExtractParamsBrowserActionFillActionFillUnion) GetRequired() (res extractExtractParamsBrowserActionFillActionFillUnionRequired) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Required.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Required.asAny()
	}
	return
}

// Can have the runtime types [*string], [*bool]
type extractExtractParamsBrowserActionFillActionFillUnionRequired struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extractExtractParamsBrowserActionFillActionFillUnionRequired) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u ExtractExtractParamsBrowserActionFillActionFillUnion) GetSkip() (res extractExtractParamsBrowserActionFillActionFillUnionSkip) {
	if vt := u.OfType; vt != nil {
		res.any = vt.Skip.asAny()
	} else if vt := u.OfPaste; vt != nil {
		res.any = vt.Skip.asAny()
	}
	return
}

// Can have the runtime types [*string], [*bool]
type extractExtractParamsBrowserActionFillActionFillUnionSkip struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *bool:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u extractExtractParamsBrowserActionFillActionFillUnionSkip) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[ExtractExtractParamsBrowserActionFillActionFillUnion](
		"mode",
		apijson.Discriminator[ExtractExtractParamsBrowserActionFillActionFillType]("type"),
		apijson.Discriminator[ExtractExtractParamsBrowserActionFillActionFillPaste]("paste"),
	)
}

// The properties Selector, Value are required.
type ExtractExtractParamsBrowserActionFillActionFillType struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector       ExtractExtractParamsBrowserActionFillActionFillTypeSelectorUnion `json:"selector,omitzero,required"`
	Value          string                                                           `json:"value,required"`
	ClickOnElement param.Opt[bool]                                                  `json:"click_on_element,omitzero"`
	Scroll         param.Opt[bool]                                                  `json:"scroll,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay ExtractExtractParamsBrowserActionFillActionFillTypeDelayUnion `json:"delay,omitzero"`
	// Any of "type".
	Mode string `json:"mode,omitzero"`
	// Any of "linear", "ghost-cursor", "windmouse".
	MouseMovementStrategy string `json:"mouse_movement_strategy,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractExtractParamsBrowserActionFillActionFillTypeRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractExtractParamsBrowserActionFillActionFillTypeSkipUnion `json:"skip,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	TypingInterval ExtractExtractParamsBrowserActionFillActionFillTypeTypingIntervalUnion `json:"typing_interval,omitzero"`
	// Any of "simple", "distribution".
	TypingStrategy string `json:"typing_strategy,omitzero"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionFillActionFillType) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionFillActionFillType
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionFillActionFillType) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractExtractParamsBrowserActionFillActionFillType](
		"mode", "type",
	)
	apijson.RegisterFieldValidator[ExtractExtractParamsBrowserActionFillActionFillType](
		"mouse_movement_strategy", "linear", "ghost-cursor", "windmouse",
	)
	apijson.RegisterFieldValidator[ExtractExtractParamsBrowserActionFillActionFillType](
		"typing_strategy", "simple", "distribution",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionFillActionFillTypeSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionFillActionFillTypeSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractExtractParamsBrowserActionFillActionFillTypeSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionFillActionFillTypeSelectorUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionFillActionFillTypeDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionFillActionFillTypeDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractExtractParamsBrowserActionFillActionFillTypeDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionFillActionFillTypeDelayUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionFillActionFillTypeRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionFillActionFillTypeRequiredString)
	OfExtractExtractsBrowserActionFillActionFillTypeRequiredString param.Opt[ExtractExtractParamsBrowserActionFillActionFillTypeRequiredString] `json:",omitzero,inline"`
	OfBool                                                         param.Opt[bool]                                                              `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionFillActionFillTypeRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionFillActionFillTypeRequiredString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionFillActionFillTypeRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionFillActionFillTypeRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionFillActionFillTypeRequiredString) {
		return &u.OfExtractExtractsBrowserActionFillActionFillTypeRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionFillActionFillTypeRequiredString string

const (
	ExtractExtractParamsBrowserActionFillActionFillTypeRequiredStringTrue  ExtractExtractParamsBrowserActionFillActionFillTypeRequiredString = "true"
	ExtractExtractParamsBrowserActionFillActionFillTypeRequiredStringFalse ExtractExtractParamsBrowserActionFillActionFillTypeRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionFillActionFillTypeSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionFillActionFillTypeSkipString)
	OfExtractExtractsBrowserActionFillActionFillTypeSkipString param.Opt[ExtractExtractParamsBrowserActionFillActionFillTypeSkipString] `json:",omitzero,inline"`
	OfBool                                                     param.Opt[bool]                                                          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionFillActionFillTypeSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionFillActionFillTypeSkipString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionFillActionFillTypeSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionFillActionFillTypeSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionFillActionFillTypeSkipString) {
		return &u.OfExtractExtractsBrowserActionFillActionFillTypeSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionFillActionFillTypeSkipString string

const (
	ExtractExtractParamsBrowserActionFillActionFillTypeSkipStringTrue  ExtractExtractParamsBrowserActionFillActionFillTypeSkipString = "true"
	ExtractExtractParamsBrowserActionFillActionFillTypeSkipStringFalse ExtractExtractParamsBrowserActionFillActionFillTypeSkipString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionFillActionFillTypeTypingIntervalUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionFillActionFillTypeTypingIntervalUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractExtractParamsBrowserActionFillActionFillTypeTypingIntervalUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionFillActionFillTypeTypingIntervalUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// The properties Mode, Selector, Value are required.
type ExtractExtractParamsBrowserActionFillActionFillPaste struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector       ExtractExtractParamsBrowserActionFillActionFillPasteSelectorUnion `json:"selector,omitzero,required"`
	Value          string                                                            `json:"value,required"`
	ClickOnElement param.Opt[bool]                                                   `json:"click_on_element,omitzero"`
	Scroll         param.Opt[bool]                                                   `json:"scroll,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Delay ExtractExtractParamsBrowserActionFillActionFillPasteDelayUnion `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractExtractParamsBrowserActionFillActionFillPasteRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractExtractParamsBrowserActionFillActionFillPasteSkipUnion `json:"skip,omitzero"`
	// This field can be elided, and will marshal its zero value as "paste".
	Mode constant.Paste `json:"mode,required"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionFillActionFillPaste) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionFillActionFillPaste
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionFillActionFillPaste) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionFillActionFillPasteSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionFillActionFillPasteSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractExtractParamsBrowserActionFillActionFillPasteSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionFillActionFillPasteSelectorUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionFillActionFillPasteDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionFillActionFillPasteDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractExtractParamsBrowserActionFillActionFillPasteDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionFillActionFillPasteDelayUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionFillActionFillPasteRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionFillActionFillPasteRequiredString)
	OfExtractExtractsBrowserActionFillActionFillPasteRequiredString param.Opt[ExtractExtractParamsBrowserActionFillActionFillPasteRequiredString] `json:",omitzero,inline"`
	OfBool                                                          param.Opt[bool]                                                               `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionFillActionFillPasteRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionFillActionFillPasteRequiredString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionFillActionFillPasteRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionFillActionFillPasteRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionFillActionFillPasteRequiredString) {
		return &u.OfExtractExtractsBrowserActionFillActionFillPasteRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionFillActionFillPasteRequiredString string

const (
	ExtractExtractParamsBrowserActionFillActionFillPasteRequiredStringTrue  ExtractExtractParamsBrowserActionFillActionFillPasteRequiredString = "true"
	ExtractExtractParamsBrowserActionFillActionFillPasteRequiredStringFalse ExtractExtractParamsBrowserActionFillActionFillPasteRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionFillActionFillPasteSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionFillActionFillPasteSkipString)
	OfExtractExtractsBrowserActionFillActionFillPasteSkipString param.Opt[ExtractExtractParamsBrowserActionFillActionFillPasteSkipString] `json:",omitzero,inline"`
	OfBool                                                      param.Opt[bool]                                                           `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionFillActionFillPasteSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionFillActionFillPasteSkipString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionFillActionFillPasteSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionFillActionFillPasteSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionFillActionFillPasteSkipString) {
		return &u.OfExtractExtractsBrowserActionFillActionFillPasteSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionFillActionFillPasteSkipString string

const (
	ExtractExtractParamsBrowserActionFillActionFillPasteSkipStringTrue  ExtractExtractParamsBrowserActionFillActionFillPasteSkipString = "true"
	ExtractExtractParamsBrowserActionFillActionFillPasteSkipStringFalse ExtractExtractParamsBrowserActionFillActionFillPasteSkipString = "false"
)

// Retrieve browser cookies
//
// The property GetCookies is required.
type ExtractExtractParamsBrowserActionGetCookiesAction struct {
	GetCookies ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesUnion `json:"get_cookies,omitzero,required"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionGetCookiesAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionGetCookiesAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionGetCookiesAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesUnion struct {
	OfBool                                                         param.Opt[bool]                                                    `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionGetCookiesActionGetCookiesObject *ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfExtractExtractsBrowserActionGetCookiesActionGetCookiesObject)
}
func (u *ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionGetCookiesActionGetCookiesObject) {
		return u.OfExtractExtractsBrowserActionGetCookiesActionGetCookiesObject
	}
	return nil
}

type ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObject struct {
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip        ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion `json:"skip,omitzero"`
	ExtraFields map[string]any                                                             `json:"-"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObject
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionGetCookiesActionGetCookiesObjectRequiredString)
	OfExtractExtractsBrowserActionGetCookiesActionGetCookiesObjectRequiredString param.Opt[ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                       param.Opt[bool]                                                                            `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionGetCookiesActionGetCookiesObjectRequiredString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionGetCookiesActionGetCookiesObjectRequiredString) {
		return &u.OfExtractExtractsBrowserActionGetCookiesActionGetCookiesObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredString string

const (
	ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredStringTrue  ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredString = "true"
	ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredStringFalse ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionGetCookiesActionGetCookiesObjectSkipString)
	OfExtractExtractsBrowserActionGetCookiesActionGetCookiesObjectSkipString param.Opt[ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                   param.Opt[bool]                                                                        `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionGetCookiesActionGetCookiesObjectSkipString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionGetCookiesActionGetCookiesObjectSkipString) {
		return &u.OfExtractExtractsBrowserActionGetCookiesActionGetCookiesObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipString string

const (
	ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipStringTrue  ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipString = "true"
	ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipStringFalse ExtractExtractParamsBrowserActionGetCookiesActionGetCookiesObjectSkipString = "false"
)

// Navigate to a URL
//
// The property Goto is required.
type ExtractExtractParamsBrowserActionGotoAction struct {
	Goto ExtractExtractParamsBrowserActionGotoActionGotoUnion `json:"goto,omitzero,required" format:"uri"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionGotoAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionGotoAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionGotoAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionGotoActionGotoUnion struct {
	OfString                                           param.Opt[string]                                      `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionGotoActionGotoObject *ExtractExtractParamsBrowserActionGotoActionGotoObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionGotoActionGotoUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractExtractsBrowserActionGotoActionGotoObject)
}
func (u *ExtractExtractParamsBrowserActionGotoActionGotoUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionGotoActionGotoUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionGotoActionGotoObject) {
		return u.OfExtractExtractsBrowserActionGotoActionGotoObject
	}
	return nil
}

// The property URL is required.
type ExtractExtractParamsBrowserActionGotoActionGotoObject struct {
	URL     string            `json:"url,required" format:"uri"`
	Referer param.Opt[string] `json:"referer,omitzero"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractExtractParamsBrowserActionGotoActionGotoObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractExtractParamsBrowserActionGotoActionGotoObjectSkipUnion `json:"skip,omitzero"`
	// Any of "load", "domcontentloaded", "networkidle0", "networkidle2".
	WaitUntil string `json:"wait_until,omitzero"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionGotoActionGotoObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionGotoActionGotoObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionGotoActionGotoObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractExtractParamsBrowserActionGotoActionGotoObject](
		"wait_until", "load", "domcontentloaded", "networkidle0", "networkidle2",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionGotoActionGotoObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionGotoActionGotoObjectRequiredString)
	OfExtractExtractsBrowserActionGotoActionGotoObjectRequiredString param.Opt[ExtractExtractParamsBrowserActionGotoActionGotoObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                           param.Opt[bool]                                                                `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionGotoActionGotoObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionGotoActionGotoObjectRequiredString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionGotoActionGotoObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionGotoActionGotoObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionGotoActionGotoObjectRequiredString) {
		return &u.OfExtractExtractsBrowserActionGotoActionGotoObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionGotoActionGotoObjectRequiredString string

const (
	ExtractExtractParamsBrowserActionGotoActionGotoObjectRequiredStringTrue  ExtractExtractParamsBrowserActionGotoActionGotoObjectRequiredString = "true"
	ExtractExtractParamsBrowserActionGotoActionGotoObjectRequiredStringFalse ExtractExtractParamsBrowserActionGotoActionGotoObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionGotoActionGotoObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionGotoActionGotoObjectSkipString)
	OfExtractExtractsBrowserActionGotoActionGotoObjectSkipString param.Opt[ExtractExtractParamsBrowserActionGotoActionGotoObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                       param.Opt[bool]                                                            `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionGotoActionGotoObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionGotoActionGotoObjectSkipString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionGotoActionGotoObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionGotoActionGotoObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionGotoActionGotoObjectSkipString) {
		return &u.OfExtractExtractsBrowserActionGotoActionGotoObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionGotoActionGotoObjectSkipString string

const (
	ExtractExtractParamsBrowserActionGotoActionGotoObjectSkipStringTrue  ExtractExtractParamsBrowserActionGotoActionGotoObjectSkipString = "true"
	ExtractExtractParamsBrowserActionGotoActionGotoObjectSkipStringFalse ExtractExtractParamsBrowserActionGotoActionGotoObjectSkipString = "false"
)

// Press a keyboard key
//
// The property Press is required.
type ExtractExtractParamsBrowserActionPressAction struct {
	Press ExtractExtractParamsBrowserActionPressActionPressUnion `json:"press,omitzero,required"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionPressAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionPressAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionPressAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionPressActionPressUnion struct {
	OfString                                             param.Opt[string]                                        `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionPressActionPressObject *ExtractExtractParamsBrowserActionPressActionPressObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionPressActionPressUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfExtractExtractsBrowserActionPressActionPressObject)
}
func (u *ExtractExtractParamsBrowserActionPressActionPressUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionPressActionPressUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionPressActionPressObject) {
		return u.OfExtractExtractsBrowserActionPressActionPressObject
	}
	return nil
}

// The property Key is required.
type ExtractExtractParamsBrowserActionPressActionPressObject struct {
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
	Delay ExtractExtractParamsBrowserActionPressActionPressObjectDelayUnion `json:"delay,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractExtractParamsBrowserActionPressActionPressObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractExtractParamsBrowserActionPressActionPressObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionPressActionPressObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionPressActionPressObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionPressActionPressObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractExtractParamsBrowserActionPressActionPressObject](
		"key", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "Power", "Eject", "Abort", "Help", "Backspace", "Tab", "Numpad5", "NumpadEnter", "Enter", "\r", "\n", "ShiftLeft", "ShiftRight", "ControlLeft", "ControlRight", "AltLeft", "AltRight", "Pause", "CapsLock", "Escape", "Convert", "NonConvert", "Space", "Numpad9", "PageUp", "Numpad3", "PageDown", "End", "Numpad1", "Home", "Numpad7", "ArrowLeft", "Numpad4", "Numpad8", "ArrowUp", "ArrowRight", "Numpad6", "Numpad2", "ArrowDown", "Select", "Open", "PrintScreen", "Insert", "Numpad0", "Delete", "NumpadDecimal", "Digit0", "Digit1", "Digit2", "Digit3", "Digit4", "Digit5", "Digit6", "Digit7", "Digit8", "Digit9", "KeyA", "KeyB", "KeyC", "KeyD", "KeyE", "KeyF", "KeyG", "KeyH", "KeyI", "KeyJ", "KeyK", "KeyL", "KeyM", "KeyN", "KeyO", "KeyP", "KeyQ", "KeyR", "KeyS", "KeyT", "KeyU", "KeyV", "KeyW", "KeyX", "KeyY", "KeyZ", "MetaLeft", "MetaRight", "ContextMenu", "NumpadMultiply", "NumpadAdd", "NumpadSubtract", "NumpadDivide", "F1", "F2", "F3", "F4", "F5", "F6", "F7", "F8", "F9", "F10", "F11", "F12", "F13", "F14", "F15", "F16", "F17", "F18", "F19", "F20", "F21", "F22", "F23", "F24", "NumLock", "ScrollLock", "AudioVolumeMute", "AudioVolumeDown", "AudioVolumeUp", "MediaTrackNext", "MediaTrackPrevious", "MediaStop", "MediaPlayPause", "Semicolon", "Equal", "NumpadEqual", "Comma", "Minus", "Period", "Slash", "Backquote", "BracketLeft", "Backslash", "BracketRight", "Quote", "AltGraph", "Props", "Cancel", "Clear", "Shift", "Control", "Alt", "Accept", "ModeChange", " ", "Print", "Execute", "\u0000", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "Meta", "*", "+", "-", "/", ";", "=", ",", ".", "`", "[", "\\", "]", "'", "Attn", "CrSel", "ExSel", "EraseEof", "Play", "ZoomOut", ")", "!", "@", "#", "$", "%", "^", "&", "(", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", ":", "<", "_", ">", "?", "~", "{", "|", "}", "\"", "SoftLeft", "SoftRight", "Camera", "Call", "EndCall", "VolumeDown", "VolumeUp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionPressActionPressObjectDelayUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionPressActionPressObjectDelayUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractExtractParamsBrowserActionPressActionPressObjectDelayUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionPressActionPressObjectDelayUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionPressActionPressObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionPressActionPressObjectRequiredString)
	OfExtractExtractsBrowserActionPressActionPressObjectRequiredString param.Opt[ExtractExtractParamsBrowserActionPressActionPressObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                             param.Opt[bool]                                                                  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionPressActionPressObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionPressActionPressObjectRequiredString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionPressActionPressObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionPressActionPressObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionPressActionPressObjectRequiredString) {
		return &u.OfExtractExtractsBrowserActionPressActionPressObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionPressActionPressObjectRequiredString string

const (
	ExtractExtractParamsBrowserActionPressActionPressObjectRequiredStringTrue  ExtractExtractParamsBrowserActionPressActionPressObjectRequiredString = "true"
	ExtractExtractParamsBrowserActionPressActionPressObjectRequiredStringFalse ExtractExtractParamsBrowserActionPressActionPressObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionPressActionPressObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionPressActionPressObjectSkipString)
	OfExtractExtractsBrowserActionPressActionPressObjectSkipString param.Opt[ExtractExtractParamsBrowserActionPressActionPressObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                         param.Opt[bool]                                                              `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionPressActionPressObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionPressActionPressObjectSkipString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionPressActionPressObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionPressActionPressObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionPressActionPressObjectSkipString) {
		return &u.OfExtractExtractsBrowserActionPressActionPressObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionPressActionPressObjectSkipString string

const (
	ExtractExtractParamsBrowserActionPressActionPressObjectSkipStringTrue  ExtractExtractParamsBrowserActionPressActionPressObjectSkipString = "true"
	ExtractExtractParamsBrowserActionPressActionPressObjectSkipStringFalse ExtractExtractParamsBrowserActionPressActionPressObjectSkipString = "false"
)

// Capture a page screenshot
//
// The property Screenshot is required.
type ExtractExtractParamsBrowserActionScreenshotAction struct {
	Screenshot ExtractExtractParamsBrowserActionScreenshotActionScreenshotUnion `json:"screenshot,omitzero,required"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionScreenshotAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionScreenshotAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionScreenshotAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionScreenshotActionScreenshotUnion struct {
	OfBool                                                         param.Opt[bool]                                                    `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionScreenshotActionScreenshotObject *ExtractExtractParamsBrowserActionScreenshotActionScreenshotObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionScreenshotActionScreenshotUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfBool, u.OfExtractExtractsBrowserActionScreenshotActionScreenshotObject)
}
func (u *ExtractExtractParamsBrowserActionScreenshotActionScreenshotUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionScreenshotActionScreenshotUnion) asAny() any {
	if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionScreenshotActionScreenshotObject) {
		return u.OfExtractExtractsBrowserActionScreenshotActionScreenshotObject
	}
	return nil
}

type ExtractExtractParamsBrowserActionScreenshotActionScreenshotObject struct {
	FullPage param.Opt[bool]    `json:"full_page,omitzero"`
	Quality  param.Opt[float64] `json:"quality,omitzero"`
	// Any of "png", "jpeg", "webp".
	Format string `json:"format,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionScreenshotActionScreenshotObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionScreenshotActionScreenshotObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionScreenshotActionScreenshotObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractExtractParamsBrowserActionScreenshotActionScreenshotObject](
		"format", "png", "jpeg", "webp",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionScreenshotActionScreenshotObjectRequiredString)
	OfExtractExtractsBrowserActionScreenshotActionScreenshotObjectRequiredString param.Opt[ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                       param.Opt[bool]                                                                            `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionScreenshotActionScreenshotObjectRequiredString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionScreenshotActionScreenshotObjectRequiredString) {
		return &u.OfExtractExtractsBrowserActionScreenshotActionScreenshotObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredString string

const (
	ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredStringTrue  ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredString = "true"
	ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredStringFalse ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionScreenshotActionScreenshotObjectSkipString)
	OfExtractExtractsBrowserActionScreenshotActionScreenshotObjectSkipString param.Opt[ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                   param.Opt[bool]                                                                        `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionScreenshotActionScreenshotObjectSkipString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionScreenshotActionScreenshotObjectSkipString) {
		return &u.OfExtractExtractsBrowserActionScreenshotActionScreenshotObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipString string

const (
	ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipStringTrue  ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipString = "true"
	ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipStringFalse ExtractExtractParamsBrowserActionScreenshotActionScreenshotObjectSkipString = "false"
)

// Scroll the page or an element
//
// The property Scroll is required.
type ExtractExtractParamsBrowserActionScrollAction struct {
	Scroll ExtractExtractParamsBrowserActionScrollActionScrollUnion `json:"scroll,omitzero,required"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionScrollAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionScrollAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionScrollAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionScrollActionScrollUnion struct {
	OfFloat                                                param.Opt[float64]                                         `json:",omitzero,inline"`
	OfString                                               param.Opt[string]                                          `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionScrollActionScrollObject *ExtractExtractParamsBrowserActionScrollActionScrollObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionScrollActionScrollUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfExtractExtractsBrowserActionScrollActionScrollObject)
}
func (u *ExtractExtractParamsBrowserActionScrollActionScrollUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionScrollActionScrollUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionScrollActionScrollObject) {
		return u.OfExtractExtractsBrowserActionScrollActionScrollObject
	}
	return nil
}

type ExtractExtractParamsBrowserActionScrollActionScrollObject struct {
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	X       param.Opt[float64] `json:"x,omitzero"`
	Y       param.Opt[float64] `json:"y,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Container ExtractExtractParamsBrowserActionScrollActionScrollObjectContainerUnion `json:"container,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractExtractParamsBrowserActionScrollActionScrollObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractExtractParamsBrowserActionScrollActionScrollObjectSkipUnion `json:"skip,omitzero"`
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	To ExtractExtractParamsBrowserActionScrollActionScrollObjectToUnion `json:"to,omitzero"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionScrollActionScrollObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionScrollActionScrollObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionScrollActionScrollObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionScrollActionScrollObjectContainerUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionScrollActionScrollObjectContainerUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractExtractParamsBrowserActionScrollActionScrollObjectContainerUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionScrollActionScrollObjectContainerUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionScrollActionScrollObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionScrollActionScrollObjectRequiredString)
	OfExtractExtractsBrowserActionScrollActionScrollObjectRequiredString param.Opt[ExtractExtractParamsBrowserActionScrollActionScrollObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                               param.Opt[bool]                                                                    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionScrollActionScrollObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionScrollActionScrollObjectRequiredString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionScrollActionScrollObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionScrollActionScrollObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionScrollActionScrollObjectRequiredString) {
		return &u.OfExtractExtractsBrowserActionScrollActionScrollObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionScrollActionScrollObjectRequiredString string

const (
	ExtractExtractParamsBrowserActionScrollActionScrollObjectRequiredStringTrue  ExtractExtractParamsBrowserActionScrollActionScrollObjectRequiredString = "true"
	ExtractExtractParamsBrowserActionScrollActionScrollObjectRequiredStringFalse ExtractExtractParamsBrowserActionScrollActionScrollObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionScrollActionScrollObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionScrollActionScrollObjectSkipString)
	OfExtractExtractsBrowserActionScrollActionScrollObjectSkipString param.Opt[ExtractExtractParamsBrowserActionScrollActionScrollObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                           param.Opt[bool]                                                                `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionScrollActionScrollObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionScrollActionScrollObjectSkipString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionScrollActionScrollObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionScrollActionScrollObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionScrollActionScrollObjectSkipString) {
		return &u.OfExtractExtractsBrowserActionScrollActionScrollObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionScrollActionScrollObjectSkipString string

const (
	ExtractExtractParamsBrowserActionScrollActionScrollObjectSkipStringTrue  ExtractExtractParamsBrowserActionScrollActionScrollObjectSkipString = "true"
	ExtractExtractParamsBrowserActionScrollActionScrollObjectSkipStringFalse ExtractExtractParamsBrowserActionScrollActionScrollObjectSkipString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionScrollActionScrollObjectToUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionScrollActionScrollObjectToUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractExtractParamsBrowserActionScrollActionScrollObjectToUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionScrollActionScrollObjectToUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionWaitAction struct {
	Wait ExtractExtractParamsBrowserActionWaitActionWaitUnion `json:"wait,omitzero,required"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionWaitAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionWaitAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionWaitAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionWaitActionWaitUnion struct {
	OfFloat                                            param.Opt[float64]                                     `json:",omitzero,inline"`
	OfString                                           param.Opt[string]                                      `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionWaitActionWaitObject *ExtractExtractParamsBrowserActionWaitActionWaitObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionWaitActionWaitUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString, u.OfExtractExtractsBrowserActionWaitActionWaitObject)
}
func (u *ExtractExtractParamsBrowserActionWaitActionWaitUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionWaitActionWaitUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionWaitActionWaitObject) {
		return u.OfExtractExtractsBrowserActionWaitActionWaitObject
	}
	return nil
}

// The property Duration is required.
type ExtractExtractParamsBrowserActionWaitActionWaitObject struct {
	// Duration value that accepts various formats. Supports: number (ms), string
	// ("1000"), or string with unit ("2s", "500ms", "2m", "1h")
	Duration ExtractExtractParamsBrowserActionWaitActionWaitObjectDurationUnion `json:"duration,omitzero,required"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractExtractParamsBrowserActionWaitActionWaitObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractExtractParamsBrowserActionWaitActionWaitObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionWaitActionWaitObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionWaitActionWaitObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionWaitActionWaitObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionWaitActionWaitObjectDurationUnion struct {
	OfFloat  param.Opt[float64] `json:",omitzero,inline"`
	OfString param.Opt[string]  `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionWaitActionWaitObjectDurationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfString)
}
func (u *ExtractExtractParamsBrowserActionWaitActionWaitObjectDurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionWaitActionWaitObjectDurationUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionWaitActionWaitObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionWaitActionWaitObjectRequiredString)
	OfExtractExtractsBrowserActionWaitActionWaitObjectRequiredString param.Opt[ExtractExtractParamsBrowserActionWaitActionWaitObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                           param.Opt[bool]                                                                `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionWaitActionWaitObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionWaitActionWaitObjectRequiredString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionWaitActionWaitObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionWaitActionWaitObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionWaitActionWaitObjectRequiredString) {
		return &u.OfExtractExtractsBrowserActionWaitActionWaitObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionWaitActionWaitObjectRequiredString string

const (
	ExtractExtractParamsBrowserActionWaitActionWaitObjectRequiredStringTrue  ExtractExtractParamsBrowserActionWaitActionWaitObjectRequiredString = "true"
	ExtractExtractParamsBrowserActionWaitActionWaitObjectRequiredStringFalse ExtractExtractParamsBrowserActionWaitActionWaitObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionWaitActionWaitObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionWaitActionWaitObjectSkipString)
	OfExtractExtractsBrowserActionWaitActionWaitObjectSkipString param.Opt[ExtractExtractParamsBrowserActionWaitActionWaitObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                       param.Opt[bool]                                                            `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionWaitActionWaitObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionWaitActionWaitObjectSkipString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionWaitActionWaitObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionWaitActionWaitObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionWaitActionWaitObjectSkipString) {
		return &u.OfExtractExtractsBrowserActionWaitActionWaitObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionWaitActionWaitObjectSkipString string

const (
	ExtractExtractParamsBrowserActionWaitActionWaitObjectSkipStringTrue  ExtractExtractParamsBrowserActionWaitActionWaitObjectSkipString = "true"
	ExtractExtractParamsBrowserActionWaitActionWaitObjectSkipStringFalse ExtractExtractParamsBrowserActionWaitActionWaitObjectSkipString = "false"
)

// Wait for an element to appear or reach a specific state
//
// The property WaitForElement is required.
type ExtractExtractParamsBrowserActionWaitForElementAction struct {
	WaitForElement ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementUnion `json:"wait_for_element,omitzero,required"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionWaitForElementAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionWaitForElementAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionWaitForElementAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementUnion struct {
	OfString                                                               param.Opt[string]                                                          `json:",omitzero,inline"`
	OfStringArray                                                          []string                                                                   `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionWaitForElementActionWaitForElementObject *ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray, u.OfExtractExtractsBrowserActionWaitForElementActionWaitForElementObject)
}
func (u *ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionWaitForElementActionWaitForElementObject) {
		return u.OfExtractExtractsBrowserActionWaitForElementActionWaitForElementObject
	}
	return nil
}

// The property Selector is required.
type ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObject struct {
	// CSS selector or array of alternative selectors. Use an array when you have
	// multiple possible selectors for the same element.
	Selector ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion `json:"selector,omitzero,required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	Visible param.Opt[bool]    `json:"visible,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSelectorUnion) asAny() any {
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
type ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionWaitForElementActionWaitForElementObjectRequiredString)
	OfExtractExtractsBrowserActionWaitForElementActionWaitForElementObjectRequiredString param.Opt[ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                               param.Opt[bool]                                                                                    `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionWaitForElementActionWaitForElementObjectRequiredString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionWaitForElementActionWaitForElementObjectRequiredString) {
		return &u.OfExtractExtractsBrowserActionWaitForElementActionWaitForElementObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredString string

const (
	ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredStringTrue  ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredString = "true"
	ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredStringFalse ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionWaitForElementActionWaitForElementObjectSkipString)
	OfExtractExtractsBrowserActionWaitForElementActionWaitForElementObjectSkipString param.Opt[ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                           param.Opt[bool]                                                                                `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionWaitForElementActionWaitForElementObjectSkipString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionWaitForElementActionWaitForElementObjectSkipString) {
		return &u.OfExtractExtractsBrowserActionWaitForElementActionWaitForElementObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipString string

const (
	ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipStringTrue  ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipString = "true"
	ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipStringFalse ExtractExtractParamsBrowserActionWaitForElementActionWaitForElementObjectSkipString = "false"
)

// Wait for page navigation to complete
//
// The property WaitForNavigation is required.
type ExtractExtractParamsBrowserActionWaitForNavigationAction struct {
	WaitForNavigation ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion `json:"wait_for_navigation,omitzero,required"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionWaitForNavigationAction) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionWaitForNavigationAction
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionWaitForNavigationAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationString)
	OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationString param.Opt[string]                                                                `json:",omitzero,inline"`
	OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationObject *ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObject `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationString, u.OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationObject)
}
func (u *ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationString) {
		return &u.OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationString
	} else if !param.IsOmitted(u.OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationObject) {
		return u.OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationObject
	}
	return nil
}

type ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationString string

const (
	ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationStringLoad             ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationString = "load"
	ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationStringDomcontentloaded ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationString = "domcontentloaded"
	ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationStringNetworkidle0     ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationString = "networkidle0"
	ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationStringNetworkidle2     ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationString = "networkidle2"
)

// The property Navigation is required.
type ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObject struct {
	// Any of "load", "domcontentloaded", "networkidle0", "networkidle2".
	Navigation string `json:"navigation,omitzero,required"`
	// Timeout in milliseconds. Set to 0 for infinite timeout (no timeout). Default:
	// 15000ms.
	Timeout param.Opt[float64] `json:"timeout,omitzero"`
	// Whether this action is required. If true, pipeline stops on failure. Accepts
	// boolean or string "true"/"false". Default: true.
	Required ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion `json:"required,omitzero"`
	// Whether to skip this action. Accepts boolean or string "true"/"false". Default:
	// false.
	Skip ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion `json:"skip,omitzero"`
	paramObj
}

func (r ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObject) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObject](
		"navigation", "load", "domcontentloaded", "networkidle0", "networkidle2",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString)
	OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString param.Opt[ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString] `json:",omitzero,inline"`
	OfBool                                                                                     param.Opt[bool]                                                                                          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString) {
		return &u.OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString string

const (
	ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredStringTrue  ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString = "true"
	ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredStringFalse ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectRequiredString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString)
	OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString param.Opt[ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString] `json:",omitzero,inline"`
	OfBool                                                                                 param.Opt[bool]                                                                                      `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString, u.OfBool)
}
func (u *ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString) {
		return &u.OfExtractExtractsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

type ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString string

const (
	ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipStringTrue  ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString = "true"
	ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipStringFalse ExtractExtractParamsBrowserActionWaitForNavigationActionWaitForNavigationObjectSkipString = "false"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsCookiesUnion struct {
	OfExtractExtractsCookiesArray []ExtractExtractParamsCookiesArrayItem `json:",omitzero,inline"`
	OfString                      param.Opt[string]                      `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsCookiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsCookiesArray, u.OfString)
}
func (u *ExtractExtractParamsCookiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsCookiesUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsCookiesArray) {
		return &u.OfExtractExtractsCookiesArray
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type ExtractExtractParamsCookiesArrayItem struct {
	Creation      param.Opt[string]                               `json:"creation,omitzero"`
	Domain        param.Opt[string]                               `json:"domain,omitzero"`
	HostOnly      param.Opt[bool]                                 `json:"hostOnly,omitzero"`
	HTTPOnly      param.Opt[bool]                                 `json:"httpOnly,omitzero"`
	LastAccessed  param.Opt[string]                               `json:"lastAccessed,omitzero"`
	Path          param.Opt[string]                               `json:"path,omitzero"`
	PathIsDefault param.Opt[bool]                                 `json:"pathIsDefault,omitzero"`
	Expires       param.Opt[string]                               `json:"expires,omitzero"`
	Name          param.Opt[string]                               `json:"name,omitzero"`
	Secure        param.Opt[bool]                                 `json:"secure,omitzero"`
	Value         param.Opt[string]                               `json:"value,omitzero"`
	Extensions    []string                                        `json:"extensions,omitzero"`
	MaxAge        ExtractExtractParamsCookiesArrayItemMaxAgeUnion `json:"maxAge,omitzero"`
	// Any of "strict", "lax", "none".
	SameSite    string         `json:"sameSite,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r ExtractExtractParamsCookiesArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsCookiesArrayItem
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *ExtractExtractParamsCookiesArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractExtractParamsCookiesArrayItem](
		"sameSite", "strict", "lax", "none",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsCookiesArrayItemMaxAgeUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfExtractExtractsCookiesArrayItemMaxAgeString)
	OfExtractExtractsCookiesArrayItemMaxAgeString param.Opt[ExtractExtractParamsCookiesArrayItemMaxAgeString] `json:",omitzero,inline"`
	OfFloat                                       param.Opt[float64]                                          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsCookiesArrayItemMaxAgeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfExtractExtractsCookiesArrayItemMaxAgeString, u.OfFloat)
}
func (u *ExtractExtractParamsCookiesArrayItemMaxAgeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsCookiesArrayItemMaxAgeUnion) asAny() any {
	if !param.IsOmitted(u.OfExtractExtractsCookiesArrayItemMaxAgeString) {
		return &u.OfExtractExtractsCookiesArrayItemMaxAgeString
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type ExtractExtractParamsCookiesArrayItemMaxAgeString string

const (
	ExtractExtractParamsCookiesArrayItemMaxAgeStringInfinity      ExtractExtractParamsCookiesArrayItemMaxAgeString = "Infinity"
	ExtractExtractParamsCookiesArrayItemMaxAgeStringMinusInfinity ExtractExtractParamsCookiesArrayItemMaxAgeString = "-Infinity"
)

// Country code for geolocation and proxy selection
type ExtractExtractParamsCountry string

const (
	ExtractExtractParamsCountryAd  ExtractExtractParamsCountry = "AD"
	ExtractExtractParamsCountryAe  ExtractExtractParamsCountry = "AE"
	ExtractExtractParamsCountryAf  ExtractExtractParamsCountry = "AF"
	ExtractExtractParamsCountryAg  ExtractExtractParamsCountry = "AG"
	ExtractExtractParamsCountryAI  ExtractExtractParamsCountry = "AI"
	ExtractExtractParamsCountryAl  ExtractExtractParamsCountry = "AL"
	ExtractExtractParamsCountryAm  ExtractExtractParamsCountry = "AM"
	ExtractExtractParamsCountryAo  ExtractExtractParamsCountry = "AO"
	ExtractExtractParamsCountryAq  ExtractExtractParamsCountry = "AQ"
	ExtractExtractParamsCountryAr  ExtractExtractParamsCountry = "AR"
	ExtractExtractParamsCountryAs  ExtractExtractParamsCountry = "AS"
	ExtractExtractParamsCountryAt  ExtractExtractParamsCountry = "AT"
	ExtractExtractParamsCountryAu  ExtractExtractParamsCountry = "AU"
	ExtractExtractParamsCountryAw  ExtractExtractParamsCountry = "AW"
	ExtractExtractParamsCountryAx  ExtractExtractParamsCountry = "AX"
	ExtractExtractParamsCountryAz  ExtractExtractParamsCountry = "AZ"
	ExtractExtractParamsCountryBa  ExtractExtractParamsCountry = "BA"
	ExtractExtractParamsCountryBb  ExtractExtractParamsCountry = "BB"
	ExtractExtractParamsCountryBd  ExtractExtractParamsCountry = "BD"
	ExtractExtractParamsCountryBe  ExtractExtractParamsCountry = "BE"
	ExtractExtractParamsCountryBf  ExtractExtractParamsCountry = "BF"
	ExtractExtractParamsCountryBg  ExtractExtractParamsCountry = "BG"
	ExtractExtractParamsCountryBh  ExtractExtractParamsCountry = "BH"
	ExtractExtractParamsCountryBi  ExtractExtractParamsCountry = "BI"
	ExtractExtractParamsCountryBj  ExtractExtractParamsCountry = "BJ"
	ExtractExtractParamsCountryBl  ExtractExtractParamsCountry = "BL"
	ExtractExtractParamsCountryBm  ExtractExtractParamsCountry = "BM"
	ExtractExtractParamsCountryBn  ExtractExtractParamsCountry = "BN"
	ExtractExtractParamsCountryBo  ExtractExtractParamsCountry = "BO"
	ExtractExtractParamsCountryBq  ExtractExtractParamsCountry = "BQ"
	ExtractExtractParamsCountryBr  ExtractExtractParamsCountry = "BR"
	ExtractExtractParamsCountryBs  ExtractExtractParamsCountry = "BS"
	ExtractExtractParamsCountryBt  ExtractExtractParamsCountry = "BT"
	ExtractExtractParamsCountryBv  ExtractExtractParamsCountry = "BV"
	ExtractExtractParamsCountryBw  ExtractExtractParamsCountry = "BW"
	ExtractExtractParamsCountryBy  ExtractExtractParamsCountry = "BY"
	ExtractExtractParamsCountryBz  ExtractExtractParamsCountry = "BZ"
	ExtractExtractParamsCountryCa  ExtractExtractParamsCountry = "CA"
	ExtractExtractParamsCountryCc  ExtractExtractParamsCountry = "CC"
	ExtractExtractParamsCountryCd  ExtractExtractParamsCountry = "CD"
	ExtractExtractParamsCountryCf  ExtractExtractParamsCountry = "CF"
	ExtractExtractParamsCountryCg  ExtractExtractParamsCountry = "CG"
	ExtractExtractParamsCountryCh  ExtractExtractParamsCountry = "CH"
	ExtractExtractParamsCountryCi  ExtractExtractParamsCountry = "CI"
	ExtractExtractParamsCountryCk  ExtractExtractParamsCountry = "CK"
	ExtractExtractParamsCountryCl  ExtractExtractParamsCountry = "CL"
	ExtractExtractParamsCountryCm  ExtractExtractParamsCountry = "CM"
	ExtractExtractParamsCountryCn  ExtractExtractParamsCountry = "CN"
	ExtractExtractParamsCountryCo  ExtractExtractParamsCountry = "CO"
	ExtractExtractParamsCountryCr  ExtractExtractParamsCountry = "CR"
	ExtractExtractParamsCountryCu  ExtractExtractParamsCountry = "CU"
	ExtractExtractParamsCountryCv  ExtractExtractParamsCountry = "CV"
	ExtractExtractParamsCountryCw  ExtractExtractParamsCountry = "CW"
	ExtractExtractParamsCountryCx  ExtractExtractParamsCountry = "CX"
	ExtractExtractParamsCountryCy  ExtractExtractParamsCountry = "CY"
	ExtractExtractParamsCountryCz  ExtractExtractParamsCountry = "CZ"
	ExtractExtractParamsCountryDe  ExtractExtractParamsCountry = "DE"
	ExtractExtractParamsCountryDj  ExtractExtractParamsCountry = "DJ"
	ExtractExtractParamsCountryDk  ExtractExtractParamsCountry = "DK"
	ExtractExtractParamsCountryDm  ExtractExtractParamsCountry = "DM"
	ExtractExtractParamsCountryDo  ExtractExtractParamsCountry = "DO"
	ExtractExtractParamsCountryDz  ExtractExtractParamsCountry = "DZ"
	ExtractExtractParamsCountryEc  ExtractExtractParamsCountry = "EC"
	ExtractExtractParamsCountryEe  ExtractExtractParamsCountry = "EE"
	ExtractExtractParamsCountryEg  ExtractExtractParamsCountry = "EG"
	ExtractExtractParamsCountryEh  ExtractExtractParamsCountry = "EH"
	ExtractExtractParamsCountryEr  ExtractExtractParamsCountry = "ER"
	ExtractExtractParamsCountryEs  ExtractExtractParamsCountry = "ES"
	ExtractExtractParamsCountryEt  ExtractExtractParamsCountry = "ET"
	ExtractExtractParamsCountryFi  ExtractExtractParamsCountry = "FI"
	ExtractExtractParamsCountryFj  ExtractExtractParamsCountry = "FJ"
	ExtractExtractParamsCountryFk  ExtractExtractParamsCountry = "FK"
	ExtractExtractParamsCountryFm  ExtractExtractParamsCountry = "FM"
	ExtractExtractParamsCountryFo  ExtractExtractParamsCountry = "FO"
	ExtractExtractParamsCountryFr  ExtractExtractParamsCountry = "FR"
	ExtractExtractParamsCountryGa  ExtractExtractParamsCountry = "GA"
	ExtractExtractParamsCountryGB  ExtractExtractParamsCountry = "GB"
	ExtractExtractParamsCountryGd  ExtractExtractParamsCountry = "GD"
	ExtractExtractParamsCountryGe  ExtractExtractParamsCountry = "GE"
	ExtractExtractParamsCountryGf  ExtractExtractParamsCountry = "GF"
	ExtractExtractParamsCountryGg  ExtractExtractParamsCountry = "GG"
	ExtractExtractParamsCountryGh  ExtractExtractParamsCountry = "GH"
	ExtractExtractParamsCountryGi  ExtractExtractParamsCountry = "GI"
	ExtractExtractParamsCountryGl  ExtractExtractParamsCountry = "GL"
	ExtractExtractParamsCountryGm  ExtractExtractParamsCountry = "GM"
	ExtractExtractParamsCountryGn  ExtractExtractParamsCountry = "GN"
	ExtractExtractParamsCountryGp  ExtractExtractParamsCountry = "GP"
	ExtractExtractParamsCountryGq  ExtractExtractParamsCountry = "GQ"
	ExtractExtractParamsCountryGr  ExtractExtractParamsCountry = "GR"
	ExtractExtractParamsCountryGs  ExtractExtractParamsCountry = "GS"
	ExtractExtractParamsCountryGt  ExtractExtractParamsCountry = "GT"
	ExtractExtractParamsCountryGu  ExtractExtractParamsCountry = "GU"
	ExtractExtractParamsCountryGw  ExtractExtractParamsCountry = "GW"
	ExtractExtractParamsCountryGy  ExtractExtractParamsCountry = "GY"
	ExtractExtractParamsCountryHk  ExtractExtractParamsCountry = "HK"
	ExtractExtractParamsCountryHm  ExtractExtractParamsCountry = "HM"
	ExtractExtractParamsCountryHn  ExtractExtractParamsCountry = "HN"
	ExtractExtractParamsCountryHr  ExtractExtractParamsCountry = "HR"
	ExtractExtractParamsCountryHt  ExtractExtractParamsCountry = "HT"
	ExtractExtractParamsCountryHu  ExtractExtractParamsCountry = "HU"
	ExtractExtractParamsCountryID  ExtractExtractParamsCountry = "ID"
	ExtractExtractParamsCountryIe  ExtractExtractParamsCountry = "IE"
	ExtractExtractParamsCountryIl  ExtractExtractParamsCountry = "IL"
	ExtractExtractParamsCountryIm  ExtractExtractParamsCountry = "IM"
	ExtractExtractParamsCountryIn  ExtractExtractParamsCountry = "IN"
	ExtractExtractParamsCountryIo  ExtractExtractParamsCountry = "IO"
	ExtractExtractParamsCountryIq  ExtractExtractParamsCountry = "IQ"
	ExtractExtractParamsCountryIr  ExtractExtractParamsCountry = "IR"
	ExtractExtractParamsCountryIs  ExtractExtractParamsCountry = "IS"
	ExtractExtractParamsCountryIt  ExtractExtractParamsCountry = "IT"
	ExtractExtractParamsCountryJe  ExtractExtractParamsCountry = "JE"
	ExtractExtractParamsCountryJm  ExtractExtractParamsCountry = "JM"
	ExtractExtractParamsCountryJo  ExtractExtractParamsCountry = "JO"
	ExtractExtractParamsCountryJp  ExtractExtractParamsCountry = "JP"
	ExtractExtractParamsCountryKe  ExtractExtractParamsCountry = "KE"
	ExtractExtractParamsCountryKg  ExtractExtractParamsCountry = "KG"
	ExtractExtractParamsCountryKh  ExtractExtractParamsCountry = "KH"
	ExtractExtractParamsCountryKi  ExtractExtractParamsCountry = "KI"
	ExtractExtractParamsCountryKm  ExtractExtractParamsCountry = "KM"
	ExtractExtractParamsCountryKn  ExtractExtractParamsCountry = "KN"
	ExtractExtractParamsCountryKp  ExtractExtractParamsCountry = "KP"
	ExtractExtractParamsCountryKr  ExtractExtractParamsCountry = "KR"
	ExtractExtractParamsCountryKw  ExtractExtractParamsCountry = "KW"
	ExtractExtractParamsCountryKy  ExtractExtractParamsCountry = "KY"
	ExtractExtractParamsCountryKz  ExtractExtractParamsCountry = "KZ"
	ExtractExtractParamsCountryLa  ExtractExtractParamsCountry = "LA"
	ExtractExtractParamsCountryLb  ExtractExtractParamsCountry = "LB"
	ExtractExtractParamsCountryLc  ExtractExtractParamsCountry = "LC"
	ExtractExtractParamsCountryLi  ExtractExtractParamsCountry = "LI"
	ExtractExtractParamsCountryLk  ExtractExtractParamsCountry = "LK"
	ExtractExtractParamsCountryLr  ExtractExtractParamsCountry = "LR"
	ExtractExtractParamsCountryLs  ExtractExtractParamsCountry = "LS"
	ExtractExtractParamsCountryLt  ExtractExtractParamsCountry = "LT"
	ExtractExtractParamsCountryLu  ExtractExtractParamsCountry = "LU"
	ExtractExtractParamsCountryLv  ExtractExtractParamsCountry = "LV"
	ExtractExtractParamsCountryLy  ExtractExtractParamsCountry = "LY"
	ExtractExtractParamsCountryMa  ExtractExtractParamsCountry = "MA"
	ExtractExtractParamsCountryMc  ExtractExtractParamsCountry = "MC"
	ExtractExtractParamsCountryMd  ExtractExtractParamsCountry = "MD"
	ExtractExtractParamsCountryMe  ExtractExtractParamsCountry = "ME"
	ExtractExtractParamsCountryMf  ExtractExtractParamsCountry = "MF"
	ExtractExtractParamsCountryMg  ExtractExtractParamsCountry = "MG"
	ExtractExtractParamsCountryMh  ExtractExtractParamsCountry = "MH"
	ExtractExtractParamsCountryMk  ExtractExtractParamsCountry = "MK"
	ExtractExtractParamsCountryMl  ExtractExtractParamsCountry = "ML"
	ExtractExtractParamsCountryMm  ExtractExtractParamsCountry = "MM"
	ExtractExtractParamsCountryMn  ExtractExtractParamsCountry = "MN"
	ExtractExtractParamsCountryMo  ExtractExtractParamsCountry = "MO"
	ExtractExtractParamsCountryMp  ExtractExtractParamsCountry = "MP"
	ExtractExtractParamsCountryMq  ExtractExtractParamsCountry = "MQ"
	ExtractExtractParamsCountryMr  ExtractExtractParamsCountry = "MR"
	ExtractExtractParamsCountryMs  ExtractExtractParamsCountry = "MS"
	ExtractExtractParamsCountryMt  ExtractExtractParamsCountry = "MT"
	ExtractExtractParamsCountryMu  ExtractExtractParamsCountry = "MU"
	ExtractExtractParamsCountryMv  ExtractExtractParamsCountry = "MV"
	ExtractExtractParamsCountryMw  ExtractExtractParamsCountry = "MW"
	ExtractExtractParamsCountryMx  ExtractExtractParamsCountry = "MX"
	ExtractExtractParamsCountryMy  ExtractExtractParamsCountry = "MY"
	ExtractExtractParamsCountryMz  ExtractExtractParamsCountry = "MZ"
	ExtractExtractParamsCountryNa  ExtractExtractParamsCountry = "NA"
	ExtractExtractParamsCountryNc  ExtractExtractParamsCountry = "NC"
	ExtractExtractParamsCountryNe  ExtractExtractParamsCountry = "NE"
	ExtractExtractParamsCountryNf  ExtractExtractParamsCountry = "NF"
	ExtractExtractParamsCountryNg  ExtractExtractParamsCountry = "NG"
	ExtractExtractParamsCountryNi  ExtractExtractParamsCountry = "NI"
	ExtractExtractParamsCountryNl  ExtractExtractParamsCountry = "NL"
	ExtractExtractParamsCountryNo  ExtractExtractParamsCountry = "NO"
	ExtractExtractParamsCountryNp  ExtractExtractParamsCountry = "NP"
	ExtractExtractParamsCountryNr  ExtractExtractParamsCountry = "NR"
	ExtractExtractParamsCountryNu  ExtractExtractParamsCountry = "NU"
	ExtractExtractParamsCountryNz  ExtractExtractParamsCountry = "NZ"
	ExtractExtractParamsCountryOm  ExtractExtractParamsCountry = "OM"
	ExtractExtractParamsCountryPa  ExtractExtractParamsCountry = "PA"
	ExtractExtractParamsCountryPe  ExtractExtractParamsCountry = "PE"
	ExtractExtractParamsCountryPf  ExtractExtractParamsCountry = "PF"
	ExtractExtractParamsCountryPg  ExtractExtractParamsCountry = "PG"
	ExtractExtractParamsCountryPh  ExtractExtractParamsCountry = "PH"
	ExtractExtractParamsCountryPk  ExtractExtractParamsCountry = "PK"
	ExtractExtractParamsCountryPl  ExtractExtractParamsCountry = "PL"
	ExtractExtractParamsCountryPm  ExtractExtractParamsCountry = "PM"
	ExtractExtractParamsCountryPn  ExtractExtractParamsCountry = "PN"
	ExtractExtractParamsCountryPr  ExtractExtractParamsCountry = "PR"
	ExtractExtractParamsCountryPs  ExtractExtractParamsCountry = "PS"
	ExtractExtractParamsCountryPt  ExtractExtractParamsCountry = "PT"
	ExtractExtractParamsCountryPw  ExtractExtractParamsCountry = "PW"
	ExtractExtractParamsCountryPy  ExtractExtractParamsCountry = "PY"
	ExtractExtractParamsCountryQa  ExtractExtractParamsCountry = "QA"
	ExtractExtractParamsCountryRe  ExtractExtractParamsCountry = "RE"
	ExtractExtractParamsCountryRo  ExtractExtractParamsCountry = "RO"
	ExtractExtractParamsCountryRs  ExtractExtractParamsCountry = "RS"
	ExtractExtractParamsCountryRu  ExtractExtractParamsCountry = "RU"
	ExtractExtractParamsCountryRw  ExtractExtractParamsCountry = "RW"
	ExtractExtractParamsCountrySa  ExtractExtractParamsCountry = "SA"
	ExtractExtractParamsCountrySb  ExtractExtractParamsCountry = "SB"
	ExtractExtractParamsCountrySc  ExtractExtractParamsCountry = "SC"
	ExtractExtractParamsCountrySd  ExtractExtractParamsCountry = "SD"
	ExtractExtractParamsCountrySe  ExtractExtractParamsCountry = "SE"
	ExtractExtractParamsCountrySg  ExtractExtractParamsCountry = "SG"
	ExtractExtractParamsCountrySh  ExtractExtractParamsCountry = "SH"
	ExtractExtractParamsCountrySi  ExtractExtractParamsCountry = "SI"
	ExtractExtractParamsCountrySj  ExtractExtractParamsCountry = "SJ"
	ExtractExtractParamsCountrySk  ExtractExtractParamsCountry = "SK"
	ExtractExtractParamsCountrySl  ExtractExtractParamsCountry = "SL"
	ExtractExtractParamsCountrySm  ExtractExtractParamsCountry = "SM"
	ExtractExtractParamsCountrySn  ExtractExtractParamsCountry = "SN"
	ExtractExtractParamsCountrySo  ExtractExtractParamsCountry = "SO"
	ExtractExtractParamsCountrySr  ExtractExtractParamsCountry = "SR"
	ExtractExtractParamsCountrySS  ExtractExtractParamsCountry = "SS"
	ExtractExtractParamsCountrySt  ExtractExtractParamsCountry = "ST"
	ExtractExtractParamsCountrySv  ExtractExtractParamsCountry = "SV"
	ExtractExtractParamsCountrySx  ExtractExtractParamsCountry = "SX"
	ExtractExtractParamsCountrySy  ExtractExtractParamsCountry = "SY"
	ExtractExtractParamsCountrySz  ExtractExtractParamsCountry = "SZ"
	ExtractExtractParamsCountryTc  ExtractExtractParamsCountry = "TC"
	ExtractExtractParamsCountryTd  ExtractExtractParamsCountry = "TD"
	ExtractExtractParamsCountryTf  ExtractExtractParamsCountry = "TF"
	ExtractExtractParamsCountryTg  ExtractExtractParamsCountry = "TG"
	ExtractExtractParamsCountryTh  ExtractExtractParamsCountry = "TH"
	ExtractExtractParamsCountryTj  ExtractExtractParamsCountry = "TJ"
	ExtractExtractParamsCountryTk  ExtractExtractParamsCountry = "TK"
	ExtractExtractParamsCountryTl  ExtractExtractParamsCountry = "TL"
	ExtractExtractParamsCountryTm  ExtractExtractParamsCountry = "TM"
	ExtractExtractParamsCountryTn  ExtractExtractParamsCountry = "TN"
	ExtractExtractParamsCountryTo  ExtractExtractParamsCountry = "TO"
	ExtractExtractParamsCountryTr  ExtractExtractParamsCountry = "TR"
	ExtractExtractParamsCountryTt  ExtractExtractParamsCountry = "TT"
	ExtractExtractParamsCountryTv  ExtractExtractParamsCountry = "TV"
	ExtractExtractParamsCountryTw  ExtractExtractParamsCountry = "TW"
	ExtractExtractParamsCountryTz  ExtractExtractParamsCountry = "TZ"
	ExtractExtractParamsCountryUa  ExtractExtractParamsCountry = "UA"
	ExtractExtractParamsCountryUg  ExtractExtractParamsCountry = "UG"
	ExtractExtractParamsCountryUm  ExtractExtractParamsCountry = "UM"
	ExtractExtractParamsCountryUs  ExtractExtractParamsCountry = "US"
	ExtractExtractParamsCountryUy  ExtractExtractParamsCountry = "UY"
	ExtractExtractParamsCountryUz  ExtractExtractParamsCountry = "UZ"
	ExtractExtractParamsCountryVa  ExtractExtractParamsCountry = "VA"
	ExtractExtractParamsCountryVc  ExtractExtractParamsCountry = "VC"
	ExtractExtractParamsCountryVe  ExtractExtractParamsCountry = "VE"
	ExtractExtractParamsCountryVg  ExtractExtractParamsCountry = "VG"
	ExtractExtractParamsCountryVi  ExtractExtractParamsCountry = "VI"
	ExtractExtractParamsCountryVn  ExtractExtractParamsCountry = "VN"
	ExtractExtractParamsCountryVu  ExtractExtractParamsCountry = "VU"
	ExtractExtractParamsCountryWf  ExtractExtractParamsCountry = "WF"
	ExtractExtractParamsCountryWs  ExtractExtractParamsCountry = "WS"
	ExtractExtractParamsCountryXk  ExtractExtractParamsCountry = "XK"
	ExtractExtractParamsCountryYe  ExtractExtractParamsCountry = "YE"
	ExtractExtractParamsCountryYt  ExtractExtractParamsCountry = "YT"
	ExtractExtractParamsCountryZa  ExtractExtractParamsCountry = "ZA"
	ExtractExtractParamsCountryZm  ExtractExtractParamsCountry = "ZM"
	ExtractExtractParamsCountryZw  ExtractExtractParamsCountry = "ZW"
	ExtractExtractParamsCountryAll ExtractExtractParamsCountry = "ALL"
)

// Device type for browser emulation
type ExtractExtractParamsDevice string

const (
	ExtractExtractParamsDeviceDesktop ExtractExtractParamsDevice = "desktop"
	ExtractExtractParamsDeviceMobile  ExtractExtractParamsDevice = "mobile"
	ExtractExtractParamsDeviceTablet  ExtractExtractParamsDevice = "tablet"
)

// Browser driver to use
type ExtractExtractParamsDriver string

const (
	ExtractExtractParamsDriverVx6     ExtractExtractParamsDriver = "vx6"
	ExtractExtractParamsDriverVx8     ExtractExtractParamsDriver = "vx8"
	ExtractExtractParamsDriverVx8Pro  ExtractExtractParamsDriver = "vx8-pro"
	ExtractExtractParamsDriverVx10    ExtractExtractParamsDriver = "vx10"
	ExtractExtractParamsDriverVx10Pro ExtractExtractParamsDriver = "vx10-pro"
	ExtractExtractParamsDriverVx12    ExtractExtractParamsDriver = "vx12"
	ExtractExtractParamsDriverVx12Pro ExtractExtractParamsDriver = "vx12-pro"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsHeaderUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsHeaderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractExtractParamsHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsHeaderUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Locale for browser language and region settings
type ExtractExtractParamsLocale string

const (
	ExtractExtractParamsLocaleAaDj      ExtractExtractParamsLocale = "aa-DJ"
	ExtractExtractParamsLocaleAaEr      ExtractExtractParamsLocale = "aa-ER"
	ExtractExtractParamsLocaleAaEt      ExtractExtractParamsLocale = "aa-ET"
	ExtractExtractParamsLocaleAf        ExtractExtractParamsLocale = "af"
	ExtractExtractParamsLocaleAfNa      ExtractExtractParamsLocale = "af-NA"
	ExtractExtractParamsLocaleAfZa      ExtractExtractParamsLocale = "af-ZA"
	ExtractExtractParamsLocaleAk        ExtractExtractParamsLocale = "ak"
	ExtractExtractParamsLocaleAkGh      ExtractExtractParamsLocale = "ak-GH"
	ExtractExtractParamsLocaleAm        ExtractExtractParamsLocale = "am"
	ExtractExtractParamsLocaleAmEt      ExtractExtractParamsLocale = "am-ET"
	ExtractExtractParamsLocaleAnEs      ExtractExtractParamsLocale = "an-ES"
	ExtractExtractParamsLocaleAr        ExtractExtractParamsLocale = "ar"
	ExtractExtractParamsLocaleArAe      ExtractExtractParamsLocale = "ar-AE"
	ExtractExtractParamsLocaleArBh      ExtractExtractParamsLocale = "ar-BH"
	ExtractExtractParamsLocaleArDz      ExtractExtractParamsLocale = "ar-DZ"
	ExtractExtractParamsLocaleArEg      ExtractExtractParamsLocale = "ar-EG"
	ExtractExtractParamsLocaleArIn      ExtractExtractParamsLocale = "ar-IN"
	ExtractExtractParamsLocaleArIq      ExtractExtractParamsLocale = "ar-IQ"
	ExtractExtractParamsLocaleArJo      ExtractExtractParamsLocale = "ar-JO"
	ExtractExtractParamsLocaleArKw      ExtractExtractParamsLocale = "ar-KW"
	ExtractExtractParamsLocaleArLb      ExtractExtractParamsLocale = "ar-LB"
	ExtractExtractParamsLocaleArLy      ExtractExtractParamsLocale = "ar-LY"
	ExtractExtractParamsLocaleArMa      ExtractExtractParamsLocale = "ar-MA"
	ExtractExtractParamsLocaleArOm      ExtractExtractParamsLocale = "ar-OM"
	ExtractExtractParamsLocaleArQa      ExtractExtractParamsLocale = "ar-QA"
	ExtractExtractParamsLocaleArSa      ExtractExtractParamsLocale = "ar-SA"
	ExtractExtractParamsLocaleArSd      ExtractExtractParamsLocale = "ar-SD"
	ExtractExtractParamsLocaleArSy      ExtractExtractParamsLocale = "ar-SY"
	ExtractExtractParamsLocaleArTn      ExtractExtractParamsLocale = "ar-TN"
	ExtractExtractParamsLocaleArYe      ExtractExtractParamsLocale = "ar-YE"
	ExtractExtractParamsLocaleAs        ExtractExtractParamsLocale = "as"
	ExtractExtractParamsLocaleAsIn      ExtractExtractParamsLocale = "as-IN"
	ExtractExtractParamsLocaleAsa       ExtractExtractParamsLocale = "asa"
	ExtractExtractParamsLocaleAsaTz     ExtractExtractParamsLocale = "asa-TZ"
	ExtractExtractParamsLocaleAstEs     ExtractExtractParamsLocale = "ast-ES"
	ExtractExtractParamsLocaleAz        ExtractExtractParamsLocale = "az"
	ExtractExtractParamsLocaleAzAz      ExtractExtractParamsLocale = "az-AZ"
	ExtractExtractParamsLocaleAzCyrl    ExtractExtractParamsLocale = "az-Cyrl"
	ExtractExtractParamsLocaleAzCyrlAz  ExtractExtractParamsLocale = "az-Cyrl-AZ"
	ExtractExtractParamsLocaleAzLatn    ExtractExtractParamsLocale = "az-Latn"
	ExtractExtractParamsLocaleAzLatnAz  ExtractExtractParamsLocale = "az-Latn-AZ"
	ExtractExtractParamsLocaleBe        ExtractExtractParamsLocale = "be"
	ExtractExtractParamsLocaleBeBy      ExtractExtractParamsLocale = "be-BY"
	ExtractExtractParamsLocaleBem       ExtractExtractParamsLocale = "bem"
	ExtractExtractParamsLocaleBemZm     ExtractExtractParamsLocale = "bem-ZM"
	ExtractExtractParamsLocaleBerDz     ExtractExtractParamsLocale = "ber-DZ"
	ExtractExtractParamsLocaleBerMa     ExtractExtractParamsLocale = "ber-MA"
	ExtractExtractParamsLocaleBez       ExtractExtractParamsLocale = "bez"
	ExtractExtractParamsLocaleBezTz     ExtractExtractParamsLocale = "bez-TZ"
	ExtractExtractParamsLocaleBg        ExtractExtractParamsLocale = "bg"
	ExtractExtractParamsLocaleBgBg      ExtractExtractParamsLocale = "bg-BG"
	ExtractExtractParamsLocaleBhoIn     ExtractExtractParamsLocale = "bho-IN"
	ExtractExtractParamsLocaleBm        ExtractExtractParamsLocale = "bm"
	ExtractExtractParamsLocaleBmMl      ExtractExtractParamsLocale = "bm-ML"
	ExtractExtractParamsLocaleBn        ExtractExtractParamsLocale = "bn"
	ExtractExtractParamsLocaleBnBd      ExtractExtractParamsLocale = "bn-BD"
	ExtractExtractParamsLocaleBnIn      ExtractExtractParamsLocale = "bn-IN"
	ExtractExtractParamsLocaleBo        ExtractExtractParamsLocale = "bo"
	ExtractExtractParamsLocaleBoCn      ExtractExtractParamsLocale = "bo-CN"
	ExtractExtractParamsLocaleBoIn      ExtractExtractParamsLocale = "bo-IN"
	ExtractExtractParamsLocaleBrFr      ExtractExtractParamsLocale = "br-FR"
	ExtractExtractParamsLocaleBrxIn     ExtractExtractParamsLocale = "brx-IN"
	ExtractExtractParamsLocaleBs        ExtractExtractParamsLocale = "bs"
	ExtractExtractParamsLocaleBsBa      ExtractExtractParamsLocale = "bs-BA"
	ExtractExtractParamsLocaleBynEr     ExtractExtractParamsLocale = "byn-ER"
	ExtractExtractParamsLocaleCa        ExtractExtractParamsLocale = "ca"
	ExtractExtractParamsLocaleCaAd      ExtractExtractParamsLocale = "ca-AD"
	ExtractExtractParamsLocaleCaEs      ExtractExtractParamsLocale = "ca-ES"
	ExtractExtractParamsLocaleCaFr      ExtractExtractParamsLocale = "ca-FR"
	ExtractExtractParamsLocaleCaIt      ExtractExtractParamsLocale = "ca-IT"
	ExtractExtractParamsLocaleCgg       ExtractExtractParamsLocale = "cgg"
	ExtractExtractParamsLocaleCggUg     ExtractExtractParamsLocale = "cgg-UG"
	ExtractExtractParamsLocaleChr       ExtractExtractParamsLocale = "chr"
	ExtractExtractParamsLocaleChrUs     ExtractExtractParamsLocale = "chr-US"
	ExtractExtractParamsLocaleCrhUa     ExtractExtractParamsLocale = "crh-UA"
	ExtractExtractParamsLocaleCs        ExtractExtractParamsLocale = "cs"
	ExtractExtractParamsLocaleCsCz      ExtractExtractParamsLocale = "cs-CZ"
	ExtractExtractParamsLocaleCsbPl     ExtractExtractParamsLocale = "csb-PL"
	ExtractExtractParamsLocaleCvRu      ExtractExtractParamsLocale = "cv-RU"
	ExtractExtractParamsLocaleCy        ExtractExtractParamsLocale = "cy"
	ExtractExtractParamsLocaleCyGB      ExtractExtractParamsLocale = "cy-GB"
	ExtractExtractParamsLocaleDa        ExtractExtractParamsLocale = "da"
	ExtractExtractParamsLocaleDaDk      ExtractExtractParamsLocale = "da-DK"
	ExtractExtractParamsLocaleDav       ExtractExtractParamsLocale = "dav"
	ExtractExtractParamsLocaleDavKe     ExtractExtractParamsLocale = "dav-KE"
	ExtractExtractParamsLocaleDe        ExtractExtractParamsLocale = "de"
	ExtractExtractParamsLocaleDeAt      ExtractExtractParamsLocale = "de-AT"
	ExtractExtractParamsLocaleDeBe      ExtractExtractParamsLocale = "de-BE"
	ExtractExtractParamsLocaleDeCh      ExtractExtractParamsLocale = "de-CH"
	ExtractExtractParamsLocaleDeDe      ExtractExtractParamsLocale = "de-DE"
	ExtractExtractParamsLocaleDeLi      ExtractExtractParamsLocale = "de-LI"
	ExtractExtractParamsLocaleDeLu      ExtractExtractParamsLocale = "de-LU"
	ExtractExtractParamsLocaleDvMv      ExtractExtractParamsLocale = "dv-MV"
	ExtractExtractParamsLocaleDzBt      ExtractExtractParamsLocale = "dz-BT"
	ExtractExtractParamsLocaleEbu       ExtractExtractParamsLocale = "ebu"
	ExtractExtractParamsLocaleEbuKe     ExtractExtractParamsLocale = "ebu-KE"
	ExtractExtractParamsLocaleEe        ExtractExtractParamsLocale = "ee"
	ExtractExtractParamsLocaleEeGh      ExtractExtractParamsLocale = "ee-GH"
	ExtractExtractParamsLocaleEeTg      ExtractExtractParamsLocale = "ee-TG"
	ExtractExtractParamsLocaleEl        ExtractExtractParamsLocale = "el"
	ExtractExtractParamsLocaleElCy      ExtractExtractParamsLocale = "el-CY"
	ExtractExtractParamsLocaleElGr      ExtractExtractParamsLocale = "el-GR"
	ExtractExtractParamsLocaleEn        ExtractExtractParamsLocale = "en"
	ExtractExtractParamsLocaleEnAg      ExtractExtractParamsLocale = "en-AG"
	ExtractExtractParamsLocaleEnAs      ExtractExtractParamsLocale = "en-AS"
	ExtractExtractParamsLocaleEnAu      ExtractExtractParamsLocale = "en-AU"
	ExtractExtractParamsLocaleEnBe      ExtractExtractParamsLocale = "en-BE"
	ExtractExtractParamsLocaleEnBw      ExtractExtractParamsLocale = "en-BW"
	ExtractExtractParamsLocaleEnBz      ExtractExtractParamsLocale = "en-BZ"
	ExtractExtractParamsLocaleEnCa      ExtractExtractParamsLocale = "en-CA"
	ExtractExtractParamsLocaleEnDk      ExtractExtractParamsLocale = "en-DK"
	ExtractExtractParamsLocaleEnGB      ExtractExtractParamsLocale = "en-GB"
	ExtractExtractParamsLocaleEnGu      ExtractExtractParamsLocale = "en-GU"
	ExtractExtractParamsLocaleEnHk      ExtractExtractParamsLocale = "en-HK"
	ExtractExtractParamsLocaleEnIe      ExtractExtractParamsLocale = "en-IE"
	ExtractExtractParamsLocaleEnIn      ExtractExtractParamsLocale = "en-IN"
	ExtractExtractParamsLocaleEnJm      ExtractExtractParamsLocale = "en-JM"
	ExtractExtractParamsLocaleEnMh      ExtractExtractParamsLocale = "en-MH"
	ExtractExtractParamsLocaleEnMp      ExtractExtractParamsLocale = "en-MP"
	ExtractExtractParamsLocaleEnMt      ExtractExtractParamsLocale = "en-MT"
	ExtractExtractParamsLocaleEnMu      ExtractExtractParamsLocale = "en-MU"
	ExtractExtractParamsLocaleEnNa      ExtractExtractParamsLocale = "en-NA"
	ExtractExtractParamsLocaleEnNg      ExtractExtractParamsLocale = "en-NG"
	ExtractExtractParamsLocaleEnNz      ExtractExtractParamsLocale = "en-NZ"
	ExtractExtractParamsLocaleEnPh      ExtractExtractParamsLocale = "en-PH"
	ExtractExtractParamsLocaleEnPk      ExtractExtractParamsLocale = "en-PK"
	ExtractExtractParamsLocaleEnSg      ExtractExtractParamsLocale = "en-SG"
	ExtractExtractParamsLocaleEnTt      ExtractExtractParamsLocale = "en-TT"
	ExtractExtractParamsLocaleEnUm      ExtractExtractParamsLocale = "en-UM"
	ExtractExtractParamsLocaleEnUs      ExtractExtractParamsLocale = "en-US"
	ExtractExtractParamsLocaleEnVi      ExtractExtractParamsLocale = "en-VI"
	ExtractExtractParamsLocaleEnZa      ExtractExtractParamsLocale = "en-ZA"
	ExtractExtractParamsLocaleEnZm      ExtractExtractParamsLocale = "en-ZM"
	ExtractExtractParamsLocaleEnZw      ExtractExtractParamsLocale = "en-ZW"
	ExtractExtractParamsLocaleEo        ExtractExtractParamsLocale = "eo"
	ExtractExtractParamsLocaleEs        ExtractExtractParamsLocale = "es"
	ExtractExtractParamsLocaleEs419     ExtractExtractParamsLocale = "es-419"
	ExtractExtractParamsLocaleEsAr      ExtractExtractParamsLocale = "es-AR"
	ExtractExtractParamsLocaleEsBo      ExtractExtractParamsLocale = "es-BO"
	ExtractExtractParamsLocaleEsCl      ExtractExtractParamsLocale = "es-CL"
	ExtractExtractParamsLocaleEsCo      ExtractExtractParamsLocale = "es-CO"
	ExtractExtractParamsLocaleEsCr      ExtractExtractParamsLocale = "es-CR"
	ExtractExtractParamsLocaleEsCu      ExtractExtractParamsLocale = "es-CU"
	ExtractExtractParamsLocaleEsDo      ExtractExtractParamsLocale = "es-DO"
	ExtractExtractParamsLocaleEsEc      ExtractExtractParamsLocale = "es-EC"
	ExtractExtractParamsLocaleEsEs      ExtractExtractParamsLocale = "es-ES"
	ExtractExtractParamsLocaleEsGq      ExtractExtractParamsLocale = "es-GQ"
	ExtractExtractParamsLocaleEsGt      ExtractExtractParamsLocale = "es-GT"
	ExtractExtractParamsLocaleEsHn      ExtractExtractParamsLocale = "es-HN"
	ExtractExtractParamsLocaleEsMx      ExtractExtractParamsLocale = "es-MX"
	ExtractExtractParamsLocaleEsNi      ExtractExtractParamsLocale = "es-NI"
	ExtractExtractParamsLocaleEsPa      ExtractExtractParamsLocale = "es-PA"
	ExtractExtractParamsLocaleEsPe      ExtractExtractParamsLocale = "es-PE"
	ExtractExtractParamsLocaleEsPr      ExtractExtractParamsLocale = "es-PR"
	ExtractExtractParamsLocaleEsPy      ExtractExtractParamsLocale = "es-PY"
	ExtractExtractParamsLocaleEsSv      ExtractExtractParamsLocale = "es-SV"
	ExtractExtractParamsLocaleEsUs      ExtractExtractParamsLocale = "es-US"
	ExtractExtractParamsLocaleEsUy      ExtractExtractParamsLocale = "es-UY"
	ExtractExtractParamsLocaleEsVe      ExtractExtractParamsLocale = "es-VE"
	ExtractExtractParamsLocaleEt        ExtractExtractParamsLocale = "et"
	ExtractExtractParamsLocaleEtEe      ExtractExtractParamsLocale = "et-EE"
	ExtractExtractParamsLocaleEu        ExtractExtractParamsLocale = "eu"
	ExtractExtractParamsLocaleEuEs      ExtractExtractParamsLocale = "eu-ES"
	ExtractExtractParamsLocaleFa        ExtractExtractParamsLocale = "fa"
	ExtractExtractParamsLocaleFaAf      ExtractExtractParamsLocale = "fa-AF"
	ExtractExtractParamsLocaleFaIr      ExtractExtractParamsLocale = "fa-IR"
	ExtractExtractParamsLocaleFf        ExtractExtractParamsLocale = "ff"
	ExtractExtractParamsLocaleFfSn      ExtractExtractParamsLocale = "ff-SN"
	ExtractExtractParamsLocaleFi        ExtractExtractParamsLocale = "fi"
	ExtractExtractParamsLocaleFiFi      ExtractExtractParamsLocale = "fi-FI"
	ExtractExtractParamsLocaleFil       ExtractExtractParamsLocale = "fil"
	ExtractExtractParamsLocaleFilPh     ExtractExtractParamsLocale = "fil-PH"
	ExtractExtractParamsLocaleFo        ExtractExtractParamsLocale = "fo"
	ExtractExtractParamsLocaleFoFo      ExtractExtractParamsLocale = "fo-FO"
	ExtractExtractParamsLocaleFr        ExtractExtractParamsLocale = "fr"
	ExtractExtractParamsLocaleFrBe      ExtractExtractParamsLocale = "fr-BE"
	ExtractExtractParamsLocaleFrBf      ExtractExtractParamsLocale = "fr-BF"
	ExtractExtractParamsLocaleFrBi      ExtractExtractParamsLocale = "fr-BI"
	ExtractExtractParamsLocaleFrBj      ExtractExtractParamsLocale = "fr-BJ"
	ExtractExtractParamsLocaleFrBl      ExtractExtractParamsLocale = "fr-BL"
	ExtractExtractParamsLocaleFrCa      ExtractExtractParamsLocale = "fr-CA"
	ExtractExtractParamsLocaleFrCd      ExtractExtractParamsLocale = "fr-CD"
	ExtractExtractParamsLocaleFrCf      ExtractExtractParamsLocale = "fr-CF"
	ExtractExtractParamsLocaleFrCg      ExtractExtractParamsLocale = "fr-CG"
	ExtractExtractParamsLocaleFrCh      ExtractExtractParamsLocale = "fr-CH"
	ExtractExtractParamsLocaleFrCi      ExtractExtractParamsLocale = "fr-CI"
	ExtractExtractParamsLocaleFrCm      ExtractExtractParamsLocale = "fr-CM"
	ExtractExtractParamsLocaleFrDj      ExtractExtractParamsLocale = "fr-DJ"
	ExtractExtractParamsLocaleFrFr      ExtractExtractParamsLocale = "fr-FR"
	ExtractExtractParamsLocaleFrGa      ExtractExtractParamsLocale = "fr-GA"
	ExtractExtractParamsLocaleFrGn      ExtractExtractParamsLocale = "fr-GN"
	ExtractExtractParamsLocaleFrGp      ExtractExtractParamsLocale = "fr-GP"
	ExtractExtractParamsLocaleFrGq      ExtractExtractParamsLocale = "fr-GQ"
	ExtractExtractParamsLocaleFrKm      ExtractExtractParamsLocale = "fr-KM"
	ExtractExtractParamsLocaleFrLu      ExtractExtractParamsLocale = "fr-LU"
	ExtractExtractParamsLocaleFrMc      ExtractExtractParamsLocale = "fr-MC"
	ExtractExtractParamsLocaleFrMf      ExtractExtractParamsLocale = "fr-MF"
	ExtractExtractParamsLocaleFrMg      ExtractExtractParamsLocale = "fr-MG"
	ExtractExtractParamsLocaleFrMl      ExtractExtractParamsLocale = "fr-ML"
	ExtractExtractParamsLocaleFrMq      ExtractExtractParamsLocale = "fr-MQ"
	ExtractExtractParamsLocaleFrNe      ExtractExtractParamsLocale = "fr-NE"
	ExtractExtractParamsLocaleFrRe      ExtractExtractParamsLocale = "fr-RE"
	ExtractExtractParamsLocaleFrRw      ExtractExtractParamsLocale = "fr-RW"
	ExtractExtractParamsLocaleFrSn      ExtractExtractParamsLocale = "fr-SN"
	ExtractExtractParamsLocaleFrTd      ExtractExtractParamsLocale = "fr-TD"
	ExtractExtractParamsLocaleFrTg      ExtractExtractParamsLocale = "fr-TG"
	ExtractExtractParamsLocaleFurIt     ExtractExtractParamsLocale = "fur-IT"
	ExtractExtractParamsLocaleFyDe      ExtractExtractParamsLocale = "fy-DE"
	ExtractExtractParamsLocaleFyNl      ExtractExtractParamsLocale = "fy-NL"
	ExtractExtractParamsLocaleGa        ExtractExtractParamsLocale = "ga"
	ExtractExtractParamsLocaleGaIe      ExtractExtractParamsLocale = "ga-IE"
	ExtractExtractParamsLocaleGdGB      ExtractExtractParamsLocale = "gd-GB"
	ExtractExtractParamsLocaleGezEr     ExtractExtractParamsLocale = "gez-ER"
	ExtractExtractParamsLocaleGezEt     ExtractExtractParamsLocale = "gez-ET"
	ExtractExtractParamsLocaleGl        ExtractExtractParamsLocale = "gl"
	ExtractExtractParamsLocaleGlEs      ExtractExtractParamsLocale = "gl-ES"
	ExtractExtractParamsLocaleGsw       ExtractExtractParamsLocale = "gsw"
	ExtractExtractParamsLocaleGswCh     ExtractExtractParamsLocale = "gsw-CH"
	ExtractExtractParamsLocaleGu        ExtractExtractParamsLocale = "gu"
	ExtractExtractParamsLocaleGuIn      ExtractExtractParamsLocale = "gu-IN"
	ExtractExtractParamsLocaleGuz       ExtractExtractParamsLocale = "guz"
	ExtractExtractParamsLocaleGuzKe     ExtractExtractParamsLocale = "guz-KE"
	ExtractExtractParamsLocaleGv        ExtractExtractParamsLocale = "gv"
	ExtractExtractParamsLocaleGvGB      ExtractExtractParamsLocale = "gv-GB"
	ExtractExtractParamsLocaleHa        ExtractExtractParamsLocale = "ha"
	ExtractExtractParamsLocaleHaLatn    ExtractExtractParamsLocale = "ha-Latn"
	ExtractExtractParamsLocaleHaLatnGh  ExtractExtractParamsLocale = "ha-Latn-GH"
	ExtractExtractParamsLocaleHaLatnNe  ExtractExtractParamsLocale = "ha-Latn-NE"
	ExtractExtractParamsLocaleHaLatnNg  ExtractExtractParamsLocale = "ha-Latn-NG"
	ExtractExtractParamsLocaleHaNg      ExtractExtractParamsLocale = "ha-NG"
	ExtractExtractParamsLocaleHaw       ExtractExtractParamsLocale = "haw"
	ExtractExtractParamsLocaleHawUs     ExtractExtractParamsLocale = "haw-US"
	ExtractExtractParamsLocaleHe        ExtractExtractParamsLocale = "he"
	ExtractExtractParamsLocaleHeIl      ExtractExtractParamsLocale = "he-IL"
	ExtractExtractParamsLocaleHi        ExtractExtractParamsLocale = "hi"
	ExtractExtractParamsLocaleHiIn      ExtractExtractParamsLocale = "hi-IN"
	ExtractExtractParamsLocaleHneIn     ExtractExtractParamsLocale = "hne-IN"
	ExtractExtractParamsLocaleHr        ExtractExtractParamsLocale = "hr"
	ExtractExtractParamsLocaleHrHr      ExtractExtractParamsLocale = "hr-HR"
	ExtractExtractParamsLocaleHsbDe     ExtractExtractParamsLocale = "hsb-DE"
	ExtractExtractParamsLocaleHtHt      ExtractExtractParamsLocale = "ht-HT"
	ExtractExtractParamsLocaleHu        ExtractExtractParamsLocale = "hu"
	ExtractExtractParamsLocaleHuHu      ExtractExtractParamsLocale = "hu-HU"
	ExtractExtractParamsLocaleHy        ExtractExtractParamsLocale = "hy"
	ExtractExtractParamsLocaleHyAm      ExtractExtractParamsLocale = "hy-AM"
	ExtractExtractParamsLocaleID        ExtractExtractParamsLocale = "id"
	ExtractExtractParamsLocaleIDID      ExtractExtractParamsLocale = "id-ID"
	ExtractExtractParamsLocaleIg        ExtractExtractParamsLocale = "ig"
	ExtractExtractParamsLocaleIgNg      ExtractExtractParamsLocale = "ig-NG"
	ExtractExtractParamsLocaleIi        ExtractExtractParamsLocale = "ii"
	ExtractExtractParamsLocaleIiCn      ExtractExtractParamsLocale = "ii-CN"
	ExtractExtractParamsLocaleIkCa      ExtractExtractParamsLocale = "ik-CA"
	ExtractExtractParamsLocaleIs        ExtractExtractParamsLocale = "is"
	ExtractExtractParamsLocaleIsIs      ExtractExtractParamsLocale = "is-IS"
	ExtractExtractParamsLocaleIt        ExtractExtractParamsLocale = "it"
	ExtractExtractParamsLocaleItCh      ExtractExtractParamsLocale = "it-CH"
	ExtractExtractParamsLocaleItIt      ExtractExtractParamsLocale = "it-IT"
	ExtractExtractParamsLocaleIuCa      ExtractExtractParamsLocale = "iu-CA"
	ExtractExtractParamsLocaleIwIl      ExtractExtractParamsLocale = "iw-IL"
	ExtractExtractParamsLocaleJa        ExtractExtractParamsLocale = "ja"
	ExtractExtractParamsLocaleJaJp      ExtractExtractParamsLocale = "ja-JP"
	ExtractExtractParamsLocaleJmc       ExtractExtractParamsLocale = "jmc"
	ExtractExtractParamsLocaleJmcTz     ExtractExtractParamsLocale = "jmc-TZ"
	ExtractExtractParamsLocaleKa        ExtractExtractParamsLocale = "ka"
	ExtractExtractParamsLocaleKaGe      ExtractExtractParamsLocale = "ka-GE"
	ExtractExtractParamsLocaleKab       ExtractExtractParamsLocale = "kab"
	ExtractExtractParamsLocaleKabDz     ExtractExtractParamsLocale = "kab-DZ"
	ExtractExtractParamsLocaleKam       ExtractExtractParamsLocale = "kam"
	ExtractExtractParamsLocaleKamKe     ExtractExtractParamsLocale = "kam-KE"
	ExtractExtractParamsLocaleKde       ExtractExtractParamsLocale = "kde"
	ExtractExtractParamsLocaleKdeTz     ExtractExtractParamsLocale = "kde-TZ"
	ExtractExtractParamsLocaleKea       ExtractExtractParamsLocale = "kea"
	ExtractExtractParamsLocaleKeaCv     ExtractExtractParamsLocale = "kea-CV"
	ExtractExtractParamsLocaleKhq       ExtractExtractParamsLocale = "khq"
	ExtractExtractParamsLocaleKhqMl     ExtractExtractParamsLocale = "khq-ML"
	ExtractExtractParamsLocaleKi        ExtractExtractParamsLocale = "ki"
	ExtractExtractParamsLocaleKiKe      ExtractExtractParamsLocale = "ki-KE"
	ExtractExtractParamsLocaleKk        ExtractExtractParamsLocale = "kk"
	ExtractExtractParamsLocaleKkCyrl    ExtractExtractParamsLocale = "kk-Cyrl"
	ExtractExtractParamsLocaleKkCyrlKz  ExtractExtractParamsLocale = "kk-Cyrl-KZ"
	ExtractExtractParamsLocaleKkKz      ExtractExtractParamsLocale = "kk-KZ"
	ExtractExtractParamsLocaleKl        ExtractExtractParamsLocale = "kl"
	ExtractExtractParamsLocaleKlGl      ExtractExtractParamsLocale = "kl-GL"
	ExtractExtractParamsLocaleKln       ExtractExtractParamsLocale = "kln"
	ExtractExtractParamsLocaleKlnKe     ExtractExtractParamsLocale = "kln-KE"
	ExtractExtractParamsLocaleKm        ExtractExtractParamsLocale = "km"
	ExtractExtractParamsLocaleKmKh      ExtractExtractParamsLocale = "km-KH"
	ExtractExtractParamsLocaleKn        ExtractExtractParamsLocale = "kn"
	ExtractExtractParamsLocaleKnIn      ExtractExtractParamsLocale = "kn-IN"
	ExtractExtractParamsLocaleKo        ExtractExtractParamsLocale = "ko"
	ExtractExtractParamsLocaleKoKr      ExtractExtractParamsLocale = "ko-KR"
	ExtractExtractParamsLocaleKok       ExtractExtractParamsLocale = "kok"
	ExtractExtractParamsLocaleKokIn     ExtractExtractParamsLocale = "kok-IN"
	ExtractExtractParamsLocaleKsIn      ExtractExtractParamsLocale = "ks-IN"
	ExtractExtractParamsLocaleKuTr      ExtractExtractParamsLocale = "ku-TR"
	ExtractExtractParamsLocaleKw        ExtractExtractParamsLocale = "kw"
	ExtractExtractParamsLocaleKwGB      ExtractExtractParamsLocale = "kw-GB"
	ExtractExtractParamsLocaleKyKg      ExtractExtractParamsLocale = "ky-KG"
	ExtractExtractParamsLocaleLag       ExtractExtractParamsLocale = "lag"
	ExtractExtractParamsLocaleLagTz     ExtractExtractParamsLocale = "lag-TZ"
	ExtractExtractParamsLocaleLbLu      ExtractExtractParamsLocale = "lb-LU"
	ExtractExtractParamsLocaleLg        ExtractExtractParamsLocale = "lg"
	ExtractExtractParamsLocaleLgUg      ExtractExtractParamsLocale = "lg-UG"
	ExtractExtractParamsLocaleLiBe      ExtractExtractParamsLocale = "li-BE"
	ExtractExtractParamsLocaleLiNl      ExtractExtractParamsLocale = "li-NL"
	ExtractExtractParamsLocaleLijIt     ExtractExtractParamsLocale = "lij-IT"
	ExtractExtractParamsLocaleLoLa      ExtractExtractParamsLocale = "lo-LA"
	ExtractExtractParamsLocaleLt        ExtractExtractParamsLocale = "lt"
	ExtractExtractParamsLocaleLtLt      ExtractExtractParamsLocale = "lt-LT"
	ExtractExtractParamsLocaleLuo       ExtractExtractParamsLocale = "luo"
	ExtractExtractParamsLocaleLuoKe     ExtractExtractParamsLocale = "luo-KE"
	ExtractExtractParamsLocaleLuy       ExtractExtractParamsLocale = "luy"
	ExtractExtractParamsLocaleLuyKe     ExtractExtractParamsLocale = "luy-KE"
	ExtractExtractParamsLocaleLv        ExtractExtractParamsLocale = "lv"
	ExtractExtractParamsLocaleLvLv      ExtractExtractParamsLocale = "lv-LV"
	ExtractExtractParamsLocaleMagIn     ExtractExtractParamsLocale = "mag-IN"
	ExtractExtractParamsLocaleMaiIn     ExtractExtractParamsLocale = "mai-IN"
	ExtractExtractParamsLocaleMas       ExtractExtractParamsLocale = "mas"
	ExtractExtractParamsLocaleMasKe     ExtractExtractParamsLocale = "mas-KE"
	ExtractExtractParamsLocaleMasTz     ExtractExtractParamsLocale = "mas-TZ"
	ExtractExtractParamsLocaleMer       ExtractExtractParamsLocale = "mer"
	ExtractExtractParamsLocaleMerKe     ExtractExtractParamsLocale = "mer-KE"
	ExtractExtractParamsLocaleMfe       ExtractExtractParamsLocale = "mfe"
	ExtractExtractParamsLocaleMfeMu     ExtractExtractParamsLocale = "mfe-MU"
	ExtractExtractParamsLocaleMg        ExtractExtractParamsLocale = "mg"
	ExtractExtractParamsLocaleMgMg      ExtractExtractParamsLocale = "mg-MG"
	ExtractExtractParamsLocaleMhrRu     ExtractExtractParamsLocale = "mhr-RU"
	ExtractExtractParamsLocaleMiNz      ExtractExtractParamsLocale = "mi-NZ"
	ExtractExtractParamsLocaleMk        ExtractExtractParamsLocale = "mk"
	ExtractExtractParamsLocaleMkMk      ExtractExtractParamsLocale = "mk-MK"
	ExtractExtractParamsLocaleMl        ExtractExtractParamsLocale = "ml"
	ExtractExtractParamsLocaleMlIn      ExtractExtractParamsLocale = "ml-IN"
	ExtractExtractParamsLocaleMnMn      ExtractExtractParamsLocale = "mn-MN"
	ExtractExtractParamsLocaleMr        ExtractExtractParamsLocale = "mr"
	ExtractExtractParamsLocaleMrIn      ExtractExtractParamsLocale = "mr-IN"
	ExtractExtractParamsLocaleMs        ExtractExtractParamsLocale = "ms"
	ExtractExtractParamsLocaleMsBn      ExtractExtractParamsLocale = "ms-BN"
	ExtractExtractParamsLocaleMsMy      ExtractExtractParamsLocale = "ms-MY"
	ExtractExtractParamsLocaleMt        ExtractExtractParamsLocale = "mt"
	ExtractExtractParamsLocaleMtMt      ExtractExtractParamsLocale = "mt-MT"
	ExtractExtractParamsLocaleMy        ExtractExtractParamsLocale = "my"
	ExtractExtractParamsLocaleMyMm      ExtractExtractParamsLocale = "my-MM"
	ExtractExtractParamsLocaleNanTw     ExtractExtractParamsLocale = "nan-TW"
	ExtractExtractParamsLocaleNaq       ExtractExtractParamsLocale = "naq"
	ExtractExtractParamsLocaleNaqNa     ExtractExtractParamsLocale = "naq-NA"
	ExtractExtractParamsLocaleNb        ExtractExtractParamsLocale = "nb"
	ExtractExtractParamsLocaleNbNo      ExtractExtractParamsLocale = "nb-NO"
	ExtractExtractParamsLocaleNd        ExtractExtractParamsLocale = "nd"
	ExtractExtractParamsLocaleNdZw      ExtractExtractParamsLocale = "nd-ZW"
	ExtractExtractParamsLocaleNdsDe     ExtractExtractParamsLocale = "nds-DE"
	ExtractExtractParamsLocaleNdsNl     ExtractExtractParamsLocale = "nds-NL"
	ExtractExtractParamsLocaleNe        ExtractExtractParamsLocale = "ne"
	ExtractExtractParamsLocaleNeIn      ExtractExtractParamsLocale = "ne-IN"
	ExtractExtractParamsLocaleNeNp      ExtractExtractParamsLocale = "ne-NP"
	ExtractExtractParamsLocaleNl        ExtractExtractParamsLocale = "nl"
	ExtractExtractParamsLocaleNlAw      ExtractExtractParamsLocale = "nl-AW"
	ExtractExtractParamsLocaleNlBe      ExtractExtractParamsLocale = "nl-BE"
	ExtractExtractParamsLocaleNlNl      ExtractExtractParamsLocale = "nl-NL"
	ExtractExtractParamsLocaleNn        ExtractExtractParamsLocale = "nn"
	ExtractExtractParamsLocaleNnNo      ExtractExtractParamsLocale = "nn-NO"
	ExtractExtractParamsLocaleNrZa      ExtractExtractParamsLocale = "nr-ZA"
	ExtractExtractParamsLocaleNsoZa     ExtractExtractParamsLocale = "nso-ZA"
	ExtractExtractParamsLocaleNyn       ExtractExtractParamsLocale = "nyn"
	ExtractExtractParamsLocaleNynUg     ExtractExtractParamsLocale = "nyn-UG"
	ExtractExtractParamsLocaleOcFr      ExtractExtractParamsLocale = "oc-FR"
	ExtractExtractParamsLocaleOm        ExtractExtractParamsLocale = "om"
	ExtractExtractParamsLocaleOmEt      ExtractExtractParamsLocale = "om-ET"
	ExtractExtractParamsLocaleOmKe      ExtractExtractParamsLocale = "om-KE"
	ExtractExtractParamsLocaleOr        ExtractExtractParamsLocale = "or"
	ExtractExtractParamsLocaleOrIn      ExtractExtractParamsLocale = "or-IN"
	ExtractExtractParamsLocaleOsRu      ExtractExtractParamsLocale = "os-RU"
	ExtractExtractParamsLocalePa        ExtractExtractParamsLocale = "pa"
	ExtractExtractParamsLocalePaArab    ExtractExtractParamsLocale = "pa-Arab"
	ExtractExtractParamsLocalePaArabPk  ExtractExtractParamsLocale = "pa-Arab-PK"
	ExtractExtractParamsLocalePaGuru    ExtractExtractParamsLocale = "pa-Guru"
	ExtractExtractParamsLocalePaGuruIn  ExtractExtractParamsLocale = "pa-Guru-IN"
	ExtractExtractParamsLocalePaIn      ExtractExtractParamsLocale = "pa-IN"
	ExtractExtractParamsLocalePaPk      ExtractExtractParamsLocale = "pa-PK"
	ExtractExtractParamsLocalePapAn     ExtractExtractParamsLocale = "pap-AN"
	ExtractExtractParamsLocalePl        ExtractExtractParamsLocale = "pl"
	ExtractExtractParamsLocalePlPl      ExtractExtractParamsLocale = "pl-PL"
	ExtractExtractParamsLocalePs        ExtractExtractParamsLocale = "ps"
	ExtractExtractParamsLocalePsAf      ExtractExtractParamsLocale = "ps-AF"
	ExtractExtractParamsLocalePt        ExtractExtractParamsLocale = "pt"
	ExtractExtractParamsLocalePtBr      ExtractExtractParamsLocale = "pt-BR"
	ExtractExtractParamsLocalePtGw      ExtractExtractParamsLocale = "pt-GW"
	ExtractExtractParamsLocalePtMz      ExtractExtractParamsLocale = "pt-MZ"
	ExtractExtractParamsLocalePtPt      ExtractExtractParamsLocale = "pt-PT"
	ExtractExtractParamsLocaleRm        ExtractExtractParamsLocale = "rm"
	ExtractExtractParamsLocaleRmCh      ExtractExtractParamsLocale = "rm-CH"
	ExtractExtractParamsLocaleRo        ExtractExtractParamsLocale = "ro"
	ExtractExtractParamsLocaleRoMd      ExtractExtractParamsLocale = "ro-MD"
	ExtractExtractParamsLocaleRoRo      ExtractExtractParamsLocale = "ro-RO"
	ExtractExtractParamsLocaleRof       ExtractExtractParamsLocale = "rof"
	ExtractExtractParamsLocaleRofTz     ExtractExtractParamsLocale = "rof-TZ"
	ExtractExtractParamsLocaleRu        ExtractExtractParamsLocale = "ru"
	ExtractExtractParamsLocaleRuMd      ExtractExtractParamsLocale = "ru-MD"
	ExtractExtractParamsLocaleRuRu      ExtractExtractParamsLocale = "ru-RU"
	ExtractExtractParamsLocaleRuUa      ExtractExtractParamsLocale = "ru-UA"
	ExtractExtractParamsLocaleRw        ExtractExtractParamsLocale = "rw"
	ExtractExtractParamsLocaleRwRw      ExtractExtractParamsLocale = "rw-RW"
	ExtractExtractParamsLocaleRwk       ExtractExtractParamsLocale = "rwk"
	ExtractExtractParamsLocaleRwkTz     ExtractExtractParamsLocale = "rwk-TZ"
	ExtractExtractParamsLocaleSaIn      ExtractExtractParamsLocale = "sa-IN"
	ExtractExtractParamsLocaleSaq       ExtractExtractParamsLocale = "saq"
	ExtractExtractParamsLocaleSaqKe     ExtractExtractParamsLocale = "saq-KE"
	ExtractExtractParamsLocaleScIt      ExtractExtractParamsLocale = "sc-IT"
	ExtractExtractParamsLocaleSdIn      ExtractExtractParamsLocale = "sd-IN"
	ExtractExtractParamsLocaleSeNo      ExtractExtractParamsLocale = "se-NO"
	ExtractExtractParamsLocaleSeh       ExtractExtractParamsLocale = "seh"
	ExtractExtractParamsLocaleSehMz     ExtractExtractParamsLocale = "seh-MZ"
	ExtractExtractParamsLocaleSes       ExtractExtractParamsLocale = "ses"
	ExtractExtractParamsLocaleSesMl     ExtractExtractParamsLocale = "ses-ML"
	ExtractExtractParamsLocaleSg        ExtractExtractParamsLocale = "sg"
	ExtractExtractParamsLocaleSgCf      ExtractExtractParamsLocale = "sg-CF"
	ExtractExtractParamsLocaleShi       ExtractExtractParamsLocale = "shi"
	ExtractExtractParamsLocaleShiLatn   ExtractExtractParamsLocale = "shi-Latn"
	ExtractExtractParamsLocaleShiLatnMa ExtractExtractParamsLocale = "shi-Latn-MA"
	ExtractExtractParamsLocaleShiTfng   ExtractExtractParamsLocale = "shi-Tfng"
	ExtractExtractParamsLocaleShiTfngMa ExtractExtractParamsLocale = "shi-Tfng-MA"
	ExtractExtractParamsLocaleShsCa     ExtractExtractParamsLocale = "shs-CA"
	ExtractExtractParamsLocaleSi        ExtractExtractParamsLocale = "si"
	ExtractExtractParamsLocaleSiLk      ExtractExtractParamsLocale = "si-LK"
	ExtractExtractParamsLocaleSidEt     ExtractExtractParamsLocale = "sid-ET"
	ExtractExtractParamsLocaleSk        ExtractExtractParamsLocale = "sk"
	ExtractExtractParamsLocaleSkSk      ExtractExtractParamsLocale = "sk-SK"
	ExtractExtractParamsLocaleSl        ExtractExtractParamsLocale = "sl"
	ExtractExtractParamsLocaleSlSi      ExtractExtractParamsLocale = "sl-SI"
	ExtractExtractParamsLocaleSn        ExtractExtractParamsLocale = "sn"
	ExtractExtractParamsLocaleSnZw      ExtractExtractParamsLocale = "sn-ZW"
	ExtractExtractParamsLocaleSo        ExtractExtractParamsLocale = "so"
	ExtractExtractParamsLocaleSoDj      ExtractExtractParamsLocale = "so-DJ"
	ExtractExtractParamsLocaleSoEt      ExtractExtractParamsLocale = "so-ET"
	ExtractExtractParamsLocaleSoKe      ExtractExtractParamsLocale = "so-KE"
	ExtractExtractParamsLocaleSoSo      ExtractExtractParamsLocale = "so-SO"
	ExtractExtractParamsLocaleSq        ExtractExtractParamsLocale = "sq"
	ExtractExtractParamsLocaleSqAl      ExtractExtractParamsLocale = "sq-AL"
	ExtractExtractParamsLocaleSqMk      ExtractExtractParamsLocale = "sq-MK"
	ExtractExtractParamsLocaleSr        ExtractExtractParamsLocale = "sr"
	ExtractExtractParamsLocaleSrCyrl    ExtractExtractParamsLocale = "sr-Cyrl"
	ExtractExtractParamsLocaleSrCyrlBa  ExtractExtractParamsLocale = "sr-Cyrl-BA"
	ExtractExtractParamsLocaleSrCyrlMe  ExtractExtractParamsLocale = "sr-Cyrl-ME"
	ExtractExtractParamsLocaleSrCyrlRs  ExtractExtractParamsLocale = "sr-Cyrl-RS"
	ExtractExtractParamsLocaleSrLatn    ExtractExtractParamsLocale = "sr-Latn"
	ExtractExtractParamsLocaleSrLatnBa  ExtractExtractParamsLocale = "sr-Latn-BA"
	ExtractExtractParamsLocaleSrLatnMe  ExtractExtractParamsLocale = "sr-Latn-ME"
	ExtractExtractParamsLocaleSrLatnRs  ExtractExtractParamsLocale = "sr-Latn-RS"
	ExtractExtractParamsLocaleSrMe      ExtractExtractParamsLocale = "sr-ME"
	ExtractExtractParamsLocaleSrRs      ExtractExtractParamsLocale = "sr-RS"
	ExtractExtractParamsLocaleSSZa      ExtractExtractParamsLocale = "ss-ZA"
	ExtractExtractParamsLocaleStZa      ExtractExtractParamsLocale = "st-ZA"
	ExtractExtractParamsLocaleSv        ExtractExtractParamsLocale = "sv"
	ExtractExtractParamsLocaleSvFi      ExtractExtractParamsLocale = "sv-FI"
	ExtractExtractParamsLocaleSvSe      ExtractExtractParamsLocale = "sv-SE"
	ExtractExtractParamsLocaleSw        ExtractExtractParamsLocale = "sw"
	ExtractExtractParamsLocaleSwKe      ExtractExtractParamsLocale = "sw-KE"
	ExtractExtractParamsLocaleSwTz      ExtractExtractParamsLocale = "sw-TZ"
	ExtractExtractParamsLocaleTa        ExtractExtractParamsLocale = "ta"
	ExtractExtractParamsLocaleTaIn      ExtractExtractParamsLocale = "ta-IN"
	ExtractExtractParamsLocaleTaLk      ExtractExtractParamsLocale = "ta-LK"
	ExtractExtractParamsLocaleTe        ExtractExtractParamsLocale = "te"
	ExtractExtractParamsLocaleTeIn      ExtractExtractParamsLocale = "te-IN"
	ExtractExtractParamsLocaleTeo       ExtractExtractParamsLocale = "teo"
	ExtractExtractParamsLocaleTeoKe     ExtractExtractParamsLocale = "teo-KE"
	ExtractExtractParamsLocaleTeoUg     ExtractExtractParamsLocale = "teo-UG"
	ExtractExtractParamsLocaleTgTj      ExtractExtractParamsLocale = "tg-TJ"
	ExtractExtractParamsLocaleTh        ExtractExtractParamsLocale = "th"
	ExtractExtractParamsLocaleThTh      ExtractExtractParamsLocale = "th-TH"
	ExtractExtractParamsLocaleTi        ExtractExtractParamsLocale = "ti"
	ExtractExtractParamsLocaleTiEr      ExtractExtractParamsLocale = "ti-ER"
	ExtractExtractParamsLocaleTiEt      ExtractExtractParamsLocale = "ti-ET"
	ExtractExtractParamsLocaleTigEr     ExtractExtractParamsLocale = "tig-ER"
	ExtractExtractParamsLocaleTkTm      ExtractExtractParamsLocale = "tk-TM"
	ExtractExtractParamsLocaleTlPh      ExtractExtractParamsLocale = "tl-PH"
	ExtractExtractParamsLocaleTnZa      ExtractExtractParamsLocale = "tn-ZA"
	ExtractExtractParamsLocaleTo        ExtractExtractParamsLocale = "to"
	ExtractExtractParamsLocaleToTo      ExtractExtractParamsLocale = "to-TO"
	ExtractExtractParamsLocaleTr        ExtractExtractParamsLocale = "tr"
	ExtractExtractParamsLocaleTrCy      ExtractExtractParamsLocale = "tr-CY"
	ExtractExtractParamsLocaleTrTr      ExtractExtractParamsLocale = "tr-TR"
	ExtractExtractParamsLocaleTsZa      ExtractExtractParamsLocale = "ts-ZA"
	ExtractExtractParamsLocaleTtRu      ExtractExtractParamsLocale = "tt-RU"
	ExtractExtractParamsLocaleTzm       ExtractExtractParamsLocale = "tzm"
	ExtractExtractParamsLocaleTzmLatn   ExtractExtractParamsLocale = "tzm-Latn"
	ExtractExtractParamsLocaleTzmLatnMa ExtractExtractParamsLocale = "tzm-Latn-MA"
	ExtractExtractParamsLocaleUgCn      ExtractExtractParamsLocale = "ug-CN"
	ExtractExtractParamsLocaleUk        ExtractExtractParamsLocale = "uk"
	ExtractExtractParamsLocaleUkUa      ExtractExtractParamsLocale = "uk-UA"
	ExtractExtractParamsLocaleUnmUs     ExtractExtractParamsLocale = "unm-US"
	ExtractExtractParamsLocaleUr        ExtractExtractParamsLocale = "ur"
	ExtractExtractParamsLocaleUrIn      ExtractExtractParamsLocale = "ur-IN"
	ExtractExtractParamsLocaleUrPk      ExtractExtractParamsLocale = "ur-PK"
	ExtractExtractParamsLocaleUz        ExtractExtractParamsLocale = "uz"
	ExtractExtractParamsLocaleUzArab    ExtractExtractParamsLocale = "uz-Arab"
	ExtractExtractParamsLocaleUzArabAf  ExtractExtractParamsLocale = "uz-Arab-AF"
	ExtractExtractParamsLocaleUzCyrl    ExtractExtractParamsLocale = "uz-Cyrl"
	ExtractExtractParamsLocaleUzCyrlUz  ExtractExtractParamsLocale = "uz-Cyrl-UZ"
	ExtractExtractParamsLocaleUzLatn    ExtractExtractParamsLocale = "uz-Latn"
	ExtractExtractParamsLocaleUzLatnUz  ExtractExtractParamsLocale = "uz-Latn-UZ"
	ExtractExtractParamsLocaleUzUz      ExtractExtractParamsLocale = "uz-UZ"
	ExtractExtractParamsLocaleVeZa      ExtractExtractParamsLocale = "ve-ZA"
	ExtractExtractParamsLocaleVi        ExtractExtractParamsLocale = "vi"
	ExtractExtractParamsLocaleViVn      ExtractExtractParamsLocale = "vi-VN"
	ExtractExtractParamsLocaleVun       ExtractExtractParamsLocale = "vun"
	ExtractExtractParamsLocaleVunTz     ExtractExtractParamsLocale = "vun-TZ"
	ExtractExtractParamsLocaleWaBe      ExtractExtractParamsLocale = "wa-BE"
	ExtractExtractParamsLocaleWaeCh     ExtractExtractParamsLocale = "wae-CH"
	ExtractExtractParamsLocaleWalEt     ExtractExtractParamsLocale = "wal-ET"
	ExtractExtractParamsLocaleWoSn      ExtractExtractParamsLocale = "wo-SN"
	ExtractExtractParamsLocaleXhZa      ExtractExtractParamsLocale = "xh-ZA"
	ExtractExtractParamsLocaleXog       ExtractExtractParamsLocale = "xog"
	ExtractExtractParamsLocaleXogUg     ExtractExtractParamsLocale = "xog-UG"
	ExtractExtractParamsLocaleYiUs      ExtractExtractParamsLocale = "yi-US"
	ExtractExtractParamsLocaleYo        ExtractExtractParamsLocale = "yo"
	ExtractExtractParamsLocaleYoNg      ExtractExtractParamsLocale = "yo-NG"
	ExtractExtractParamsLocaleYueHk     ExtractExtractParamsLocale = "yue-HK"
	ExtractExtractParamsLocaleZh        ExtractExtractParamsLocale = "zh"
	ExtractExtractParamsLocaleZhCn      ExtractExtractParamsLocale = "zh-CN"
	ExtractExtractParamsLocaleZhHk      ExtractExtractParamsLocale = "zh-HK"
	ExtractExtractParamsLocaleZhHans    ExtractExtractParamsLocale = "zh-Hans"
	ExtractExtractParamsLocaleZhHansCn  ExtractExtractParamsLocale = "zh-Hans-CN"
	ExtractExtractParamsLocaleZhHansHk  ExtractExtractParamsLocale = "zh-Hans-HK"
	ExtractExtractParamsLocaleZhHansMo  ExtractExtractParamsLocale = "zh-Hans-MO"
	ExtractExtractParamsLocaleZhHansSg  ExtractExtractParamsLocale = "zh-Hans-SG"
	ExtractExtractParamsLocaleZhHant    ExtractExtractParamsLocale = "zh-Hant"
	ExtractExtractParamsLocaleZhHantHk  ExtractExtractParamsLocale = "zh-Hant-HK"
	ExtractExtractParamsLocaleZhHantMo  ExtractExtractParamsLocale = "zh-Hant-MO"
	ExtractExtractParamsLocaleZhHantTw  ExtractExtractParamsLocale = "zh-Hant-TW"
	ExtractExtractParamsLocaleZhSg      ExtractExtractParamsLocale = "zh-SG"
	ExtractExtractParamsLocaleZhTw      ExtractExtractParamsLocale = "zh-TW"
	ExtractExtractParamsLocaleZu        ExtractExtractParamsLocale = "zu"
	ExtractExtractParamsLocaleZuZa      ExtractExtractParamsLocale = "zu-ZA"
	ExtractExtractParamsLocaleAuto      ExtractExtractParamsLocale = "auto"
)

// HTTP method for the request
type ExtractExtractParamsMethod string

const (
	ExtractExtractParamsMethodGet    ExtractExtractParamsMethod = "GET"
	ExtractExtractParamsMethodPost   ExtractExtractParamsMethod = "POST"
	ExtractExtractParamsMethodPut    ExtractExtractParamsMethod = "PUT"
	ExtractExtractParamsMethodPatch  ExtractExtractParamsMethod = "PATCH"
	ExtractExtractParamsMethodDelete ExtractExtractParamsMethod = "DELETE"
)

type ExtractExtractParamsNetworkCapture struct {
	Validation                  param.Opt[bool]    `json:"validation,omitzero"`
	WaitForRequestsCount        param.Opt[float64] `json:"wait_for_requests_count,omitzero"`
	WaitForRequestsCountTimeout param.Opt[float64] `json:"wait_for_requests_count_timeout,omitzero"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method,omitzero"`
	// Resource type for network capture filtering
	ResourceType ExtractExtractParamsNetworkCaptureResourceTypeUnion `json:"resource_type,omitzero"`
	StatusCode   ExtractExtractParamsNetworkCaptureStatusCodeUnion   `json:"status_code,omitzero"`
	URL          ExtractExtractParamsNetworkCaptureURL               `json:"url,omitzero"`
	paramObj
}

func (r ExtractExtractParamsNetworkCapture) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsNetworkCapture
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractExtractParamsNetworkCapture](
		"method", "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsNetworkCaptureResourceTypeUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsNetworkCaptureResourceTypeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractExtractParamsNetworkCaptureResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsNetworkCaptureResourceTypeUnion) asAny() any {
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
type ExtractExtractParamsNetworkCaptureStatusCodeUnion struct {
	OfFloat      param.Opt[float64] `json:",omitzero,inline"`
	OfFloatArray []float64          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsNetworkCaptureStatusCodeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfFloatArray)
}
func (u *ExtractExtractParamsNetworkCaptureStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsNetworkCaptureStatusCodeUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfFloatArray) {
		return &u.OfFloatArray
	}
	return nil
}

// The property Value is required.
type ExtractExtractParamsNetworkCaptureURL struct {
	Value string `json:"value,required"`
	// Any of "exact", "contains".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ExtractExtractParamsNetworkCaptureURL) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsNetworkCaptureURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsNetworkCaptureURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractExtractParamsNetworkCaptureURL](
		"type", "exact", "contains",
	)
}

// Operating system to emulate
type ExtractExtractParamsOs string

const (
	ExtractExtractParamsOsWindows ExtractExtractParamsOs = "windows"
	ExtractExtractParamsOsMacOs   ExtractExtractParamsOs = "mac os"
	ExtractExtractParamsOsLinux   ExtractExtractParamsOs = "linux"
	ExtractExtractParamsOsAndroid ExtractExtractParamsOs = "android"
	ExtractExtractParamsOsIos     ExtractExtractParamsOs = "ios"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsParserUnion struct {
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsParserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap, u.OfString)
}
func (u *ExtractExtractParamsParserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsParserUnion) asAny() any {
	if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Referrer policy for the request
type ExtractExtractParamsReferrerType string

const (
	ExtractExtractParamsReferrerTypeRandom     ExtractExtractParamsReferrerType = "random"
	ExtractExtractParamsReferrerTypeNoReferer  ExtractExtractParamsReferrerType = "no-referer"
	ExtractExtractParamsReferrerTypeSameOrigin ExtractExtractParamsReferrerType = "same-origin"
	ExtractExtractParamsReferrerTypeGoogle     ExtractExtractParamsReferrerType = "google"
	ExtractExtractParamsReferrerTypeBing       ExtractExtractParamsReferrerType = "bing"
	ExtractExtractParamsReferrerTypeFacebook   ExtractExtractParamsReferrerType = "facebook"
	ExtractExtractParamsReferrerTypeTwitter    ExtractExtractParamsReferrerType = "twitter"
	ExtractExtractParamsReferrerTypeInstagram  ExtractExtractParamsReferrerType = "instagram"
)

type ExtractExtractParamsSession struct {
	ID                  param.Opt[string]  `json:"id,omitzero"`
	PrefetchUserbrowser param.Opt[bool]    `json:"prefetch_userbrowser,omitzero"`
	Retry               param.Opt[bool]    `json:"retry,omitzero"`
	Timeout             param.Opt[float64] `json:"timeout,omitzero"`
	paramObj
}

func (r ExtractExtractParamsSession) MarshalJSON() (data []byte, err error) {
	type shadow ExtractExtractParamsSession
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractExtractParamsSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ExtractExtractParamsSkillUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u ExtractExtractParamsSkillUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *ExtractExtractParamsSkillUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ExtractExtractParamsSkillUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// US state for geolocation (only valid when country is US)
type ExtractExtractParamsState string

const (
	ExtractExtractParamsStateAl ExtractExtractParamsState = "AL"
	ExtractExtractParamsStateAk ExtractExtractParamsState = "AK"
	ExtractExtractParamsStateAs ExtractExtractParamsState = "AS"
	ExtractExtractParamsStateAz ExtractExtractParamsState = "AZ"
	ExtractExtractParamsStateAr ExtractExtractParamsState = "AR"
	ExtractExtractParamsStateCa ExtractExtractParamsState = "CA"
	ExtractExtractParamsStateCo ExtractExtractParamsState = "CO"
	ExtractExtractParamsStateCt ExtractExtractParamsState = "CT"
	ExtractExtractParamsStateDe ExtractExtractParamsState = "DE"
	ExtractExtractParamsStateDc ExtractExtractParamsState = "DC"
	ExtractExtractParamsStateFl ExtractExtractParamsState = "FL"
	ExtractExtractParamsStateGa ExtractExtractParamsState = "GA"
	ExtractExtractParamsStateGu ExtractExtractParamsState = "GU"
	ExtractExtractParamsStateHi ExtractExtractParamsState = "HI"
	ExtractExtractParamsStateID ExtractExtractParamsState = "ID"
	ExtractExtractParamsStateIl ExtractExtractParamsState = "IL"
	ExtractExtractParamsStateIn ExtractExtractParamsState = "IN"
	ExtractExtractParamsStateIa ExtractExtractParamsState = "IA"
	ExtractExtractParamsStateKs ExtractExtractParamsState = "KS"
	ExtractExtractParamsStateKy ExtractExtractParamsState = "KY"
	ExtractExtractParamsStateLa ExtractExtractParamsState = "LA"
	ExtractExtractParamsStateMe ExtractExtractParamsState = "ME"
	ExtractExtractParamsStateMd ExtractExtractParamsState = "MD"
	ExtractExtractParamsStateMa ExtractExtractParamsState = "MA"
	ExtractExtractParamsStateMi ExtractExtractParamsState = "MI"
	ExtractExtractParamsStateMn ExtractExtractParamsState = "MN"
	ExtractExtractParamsStateMs ExtractExtractParamsState = "MS"
	ExtractExtractParamsStateMo ExtractExtractParamsState = "MO"
	ExtractExtractParamsStateMt ExtractExtractParamsState = "MT"
	ExtractExtractParamsStateNe ExtractExtractParamsState = "NE"
	ExtractExtractParamsStateNv ExtractExtractParamsState = "NV"
	ExtractExtractParamsStateNh ExtractExtractParamsState = "NH"
	ExtractExtractParamsStateNj ExtractExtractParamsState = "NJ"
	ExtractExtractParamsStateNm ExtractExtractParamsState = "NM"
	ExtractExtractParamsStateNy ExtractExtractParamsState = "NY"
	ExtractExtractParamsStateNc ExtractExtractParamsState = "NC"
	ExtractExtractParamsStateNd ExtractExtractParamsState = "ND"
	ExtractExtractParamsStateMp ExtractExtractParamsState = "MP"
	ExtractExtractParamsStateOh ExtractExtractParamsState = "OH"
	ExtractExtractParamsStateOk ExtractExtractParamsState = "OK"
	ExtractExtractParamsStateOr ExtractExtractParamsState = "OR"
	ExtractExtractParamsStatePa ExtractExtractParamsState = "PA"
	ExtractExtractParamsStatePr ExtractExtractParamsState = "PR"
	ExtractExtractParamsStateRi ExtractExtractParamsState = "RI"
	ExtractExtractParamsStateSc ExtractExtractParamsState = "SC"
	ExtractExtractParamsStateSd ExtractExtractParamsState = "SD"
	ExtractExtractParamsStateTn ExtractExtractParamsState = "TN"
	ExtractExtractParamsStateTx ExtractExtractParamsState = "TX"
	ExtractExtractParamsStateUt ExtractExtractParamsState = "UT"
	ExtractExtractParamsStateVt ExtractExtractParamsState = "VT"
	ExtractExtractParamsStateVa ExtractExtractParamsState = "VA"
	ExtractExtractParamsStateVi ExtractExtractParamsState = "VI"
	ExtractExtractParamsStateWa ExtractExtractParamsState = "WA"
	ExtractExtractParamsStateWv ExtractExtractParamsState = "WV"
	ExtractExtractParamsStateWi ExtractExtractParamsState = "WI"
	ExtractExtractParamsStateWy ExtractExtractParamsState = "WY"
)
