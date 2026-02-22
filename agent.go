// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomnimblewaynimblego

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/Nimbleway/nimble-go/internal/apijson"
	"github.com/Nimbleway/nimble-go/internal/apiquery"
	"github.com/Nimbleway/nimble-go/internal/requestconfig"
	"github.com/Nimbleway/nimble-go/option"
	"github.com/Nimbleway/nimble-go/packages/param"
	"github.com/Nimbleway/nimble-go/packages/respjson"
	"github.com/Nimbleway/nimble-go/shared/constant"
)

// AgentService contains methods and other services that help with interacting with
// the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentService] method instead.
type AgentService struct {
	Options []option.RequestOption
}

// NewAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentService(opts ...option.RequestOption) (r AgentService) {
	r = AgentService{}
	r.Options = opts
	return
}

// List Templates
func (r *AgentService) List(ctx context.Context, query AgentListParams, opts ...option.RequestOption) (res *[]AgentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Execute WSA Async Endpoint
func (r *AgentService) Async(ctx context.Context, body AgentAsyncParams, opts ...option.RequestOption) (res *AgentAsyncResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/agents/async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get Template
func (r *AgentService) Get(ctx context.Context, templateName string, opts ...option.RequestOption) (res *AgentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateName == "" {
		err = errors.New("missing required template_name parameter")
		return
	}
	path := fmt.Sprintf("v1/agents/%s", templateName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Execute WSA Realtime Endpoint
func (r *AgentService) Run(ctx context.Context, body AgentRunParams, opts ...option.RequestOption) (res *AgentRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/agents/run"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AgentListResponse struct {
	DisplayName string `json:"display_name,required"`
	IsPublic    bool   `json:"is_public,required"`
	Name        string `json:"name,required"`
	Description string `json:"description,nullable"`
	Domain      string `json:"domain,nullable"`
	EntityType  string `json:"entity_type,nullable"`
	ManagedBy   string `json:"managed_by,nullable"`
	Vertical    string `json:"vertical,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		IsPublic    respjson.Field
		Name        respjson.Field
		Description respjson.Field
		Domain      respjson.Field
		EntityType  respjson.Field
		ManagedBy   respjson.Field
		Vertical    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentAsyncResponse struct {
	Status constant.Success `json:"status,required"`
	Task   map[string]any   `json:"task,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Task        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentAsyncResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentAsyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGetResponse struct {
	DisplayName     string                          `json:"display_name,required"`
	IsPublic        bool                            `json:"is_public,required"`
	Name            string                          `json:"name,required"`
	Description     string                          `json:"description,nullable"`
	Domain          string                          `json:"domain,nullable"`
	EntityType      string                          `json:"entity_type,nullable"`
	FeatureFlags    AgentGetResponseFeatureFlags    `json:"feature_flags"`
	InputProperties []AgentGetResponseInputProperty `json:"input_properties,nullable"`
	ManagedBy       string                          `json:"managed_by,nullable"`
	OutputSchema    map[string]any                  `json:"output_schema,nullable"`
	Vertical        string                          `json:"vertical,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName     respjson.Field
		IsPublic        respjson.Field
		Name            respjson.Field
		Description     respjson.Field
		Domain          respjson.Field
		EntityType      respjson.Field
		FeatureFlags    respjson.Field
		InputProperties respjson.Field
		ManagedBy       respjson.Field
		OutputSchema    respjson.Field
		Vertical        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGetResponseFeatureFlags struct {
	IsLocalizationSupported bool `json:"is_localization_supported"`
	IsPaginationSupported   bool `json:"is_pagination_supported"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsLocalizationSupported respjson.Field
		IsPaginationSupported   respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGetResponseFeatureFlags) RawJSON() string { return r.JSON.raw }
func (r *AgentGetResponseFeatureFlags) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGetResponseInputProperty struct {
	Default             string   `json:"default,nullable"`
	Description         string   `json:"description,nullable"`
	Examples            []string `json:"examples,nullable"`
	IsLocalizationParam bool     `json:"is_localization_param"`
	Name                string   `json:"name"`
	Required            bool     `json:"required"`
	Rules               []string `json:"rules,nullable"`
	Type                string   `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default             respjson.Field
		Description         respjson.Field
		Examples            respjson.Field
		IsLocalizationParam respjson.Field
		Name                respjson.Field
		Required            respjson.Field
		Rules               respjson.Field
		Type                respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGetResponseInputProperty) RawJSON() string { return r.JSON.raw }
func (r *AgentGetResponseInputProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResponse struct {
	Data     AgentRunResponseData     `json:"data,required"`
	Metadata AgentRunResponseMetadata `json:"metadata,required"`
	// The status of the task.
	//
	// Any of "success", "skipped", "fatal", "error", "postponed", "ignored",
	// "rejected", "blocked".
	Status AgentRunResponseStatus `json:"status,required"`
	// Unique identifier for the task.
	TaskID string `json:"task_id,required"`
	// The final URL.
	URL   string                `json:"url,required"`
	Debug AgentRunResponseDebug `json:"debug"`
	// Pagination information if applicable.
	Pagination AgentRunResponsePaginationUnion `json:"pagination"`
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
func (r AgentRunResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResponseData struct {
	// Browser actions execution results. Present only when browser_actions were
	// specified in the request.
	BrowserActions AgentRunResponseDataBrowserActions `json:"browser_actions"`
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
	NetworkCapture []AgentRunResponseDataNetworkCapture `json:"network_capture"`
	// The parsing results extracted from the HTML & network content.
	Parsing AgentRunResponseDataParsingUnion `json:"parsing"`
	// The list of redirects that occurred during the task.
	Redirects []AgentRunResponseDataRedirect `json:"redirects"`
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
		Parsing        respjson.Field
		Redirects      respjson.Field
		Screenshots    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResponseData) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser actions execution results. Present only when browser_actions were
// specified in the request.
type AgentRunResponseDataBrowserActions struct {
	Results       []AgentRunResponseDataBrowserActionsResult `json:"results,required"`
	Success       bool                                       `json:"success,required"`
	TotalDuration float64                                    `json:"total_duration,required"`
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
func (r AgentRunResponseDataBrowserActions) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponseDataBrowserActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResponseDataBrowserActionsResult struct {
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
func (r AgentRunResponseDataBrowserActionsResult) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponseDataBrowserActionsResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResponseDataNetworkCapture struct {
	Filter       AgentRunResponseDataNetworkCaptureFilter   `json:"filter,required"`
	Results      []AgentRunResponseDataNetworkCaptureResult `json:"results,required"`
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
func (r AgentRunResponseDataNetworkCapture) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponseDataNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResponseDataNetworkCaptureFilter struct {
	Validation           bool    `json:"validation,required"`
	WaitForRequestsCount float64 `json:"wait_for_requests_count,required"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method"`
	// Resource type for network capture filtering
	ResourceType                AgentRunResponseDataNetworkCaptureFilterResourceTypeUnion `json:"resource_type"`
	StatusCode                  AgentRunResponseDataNetworkCaptureFilterStatusCodeUnion   `json:"status_code"`
	URL                         AgentRunResponseDataNetworkCaptureFilterURL               `json:"url"`
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
func (r AgentRunResponseDataNetworkCaptureFilter) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponseDataNetworkCaptureFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentRunResponseDataNetworkCaptureFilterResourceTypeUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAgentRunResponseDataNetworkCaptureFilterResourceTypeString
// OfAgentRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray]
type AgentRunResponseDataNetworkCaptureFilterResourceTypeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfAgentRunResponseDataNetworkCaptureFilterResourceTypeString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfAgentRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray []string `json:",inline"`
	JSON                                                                 struct {
		OfAgentRunResponseDataNetworkCaptureFilterResourceTypeString         respjson.Field
		OfAgentRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray respjson.Field
		raw                                                                  string
	} `json:"-"`
}

func (u AgentRunResponseDataNetworkCaptureFilterResourceTypeUnion) AsAgentRunResponseDataNetworkCaptureFilterResourceTypeString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentRunResponseDataNetworkCaptureFilterResourceTypeUnion) AsAgentRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentRunResponseDataNetworkCaptureFilterResourceTypeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AgentRunResponseDataNetworkCaptureFilterResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type for network capture filtering
type AgentRunResponseDataNetworkCaptureFilterResourceTypeString string

const (
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringDocument           AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "document"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringStylesheet         AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "stylesheet"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringImage              AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "image"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringMedia              AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "media"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringFont               AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "font"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringScript             AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "script"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringTexttrack          AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "texttrack"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringXhr                AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "xhr"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringFetch              AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "fetch"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringPrefetch           AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "prefetch"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringEventsource        AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "eventsource"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringWebsocket          AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "websocket"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringManifest           AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "manifest"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringSignedexchange     AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "signedexchange"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringPing               AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "ping"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringCspviolationreport AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "cspviolationreport"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringPreflight          AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "preflight"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringOther              AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "other"
	AgentRunResponseDataNetworkCaptureFilterResourceTypeStringFedcm              AgentRunResponseDataNetworkCaptureFilterResourceTypeString = "fedcm"
)

// AgentRunResponseDataNetworkCaptureFilterStatusCodeUnion contains all possible
// properties and values from [float64], [[]float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfFloatArray]
type AgentRunResponseDataNetworkCaptureFilterStatusCodeUnion struct {
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

func (u AgentRunResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentRunResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloatArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentRunResponseDataNetworkCaptureFilterStatusCodeUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentRunResponseDataNetworkCaptureFilterStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResponseDataNetworkCaptureFilterURL struct {
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
func (r AgentRunResponseDataNetworkCaptureFilterURL) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponseDataNetworkCaptureFilterURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResponseDataNetworkCaptureResult struct {
	Request  AgentRunResponseDataNetworkCaptureResultRequest  `json:"request,required"`
	Response AgentRunResponseDataNetworkCaptureResultResponse `json:"response,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Request     respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResponseDataNetworkCaptureResult) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponseDataNetworkCaptureResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResponseDataNetworkCaptureResultRequest struct {
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
func (r AgentRunResponseDataNetworkCaptureResultRequest) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponseDataNetworkCaptureResultRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResponseDataNetworkCaptureResultResponse struct {
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
func (r AgentRunResponseDataNetworkCaptureResultResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponseDataNetworkCaptureResultResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentRunResponseDataParsingUnion contains all possible properties and values
// from [AgentRunResponseDataParsingParsingSuccessResult],
// [AgentRunResponseDataParsingParsingErrorResult], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAgentRunResponseDataParsingMapItem]
type AgentRunResponseDataParsingUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfAgentRunResponseDataParsingMapItem any `json:",inline"`
	// This field is from variant [AgentRunResponseDataParsingParsingSuccessResult].
	Entities map[string]any `json:"entities"`
	Status   string         `json:"status"`
	// This field is from variant [AgentRunResponseDataParsingParsingErrorResult].
	Error string `json:"error"`
	JSON  struct {
		OfAgentRunResponseDataParsingMapItem respjson.Field
		Entities                             respjson.Field
		Status                               respjson.Field
		Error                                respjson.Field
		raw                                  string
	} `json:"-"`
}

func (u AgentRunResponseDataParsingUnion) AsAgentRunResponseDataParsingParsingSuccessResult() (v AgentRunResponseDataParsingParsingSuccessResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentRunResponseDataParsingUnion) AsAgentRunResponseDataParsingParsingErrorResult() (v AgentRunResponseDataParsingParsingErrorResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentRunResponseDataParsingUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentRunResponseDataParsingUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentRunResponseDataParsingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResponseDataParsingParsingSuccessResult struct {
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
func (r AgentRunResponseDataParsingParsingSuccessResult) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponseDataParsingParsingSuccessResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResponseDataParsingParsingErrorResult struct {
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
func (r AgentRunResponseDataParsingParsingErrorResult) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponseDataParsingParsingErrorResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResponseDataRedirect struct {
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
func (r AgentRunResponseDataRedirect) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponseDataRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResponseMetadata struct {
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
func (r AgentRunResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the task.
type AgentRunResponseStatus string

const (
	AgentRunResponseStatusSuccess   AgentRunResponseStatus = "success"
	AgentRunResponseStatusSkipped   AgentRunResponseStatus = "skipped"
	AgentRunResponseStatusFatal     AgentRunResponseStatus = "fatal"
	AgentRunResponseStatusError     AgentRunResponseStatus = "error"
	AgentRunResponseStatusPostponed AgentRunResponseStatus = "postponed"
	AgentRunResponseStatusIgnored   AgentRunResponseStatus = "ignored"
	AgentRunResponseStatusRejected  AgentRunResponseStatus = "rejected"
	AgentRunResponseStatusBlocked   AgentRunResponseStatus = "blocked"
)

type AgentRunResponseDebug struct {
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
func (r AgentRunResponseDebug) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponseDebug) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentRunResponsePaginationUnion contains all possible properties and values from
// [AgentRunResponsePaginationNextPageParams],
// [[]AgentRunResponsePaginationArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAgentRunResponsePaginationArray]
type AgentRunResponsePaginationUnion struct {
	// This field will be present if the value is a
	// [[]AgentRunResponsePaginationArrayItem] instead of an object.
	OfAgentRunResponsePaginationArray []AgentRunResponsePaginationArrayItem `json:",inline"`
	// This field is from variant [AgentRunResponsePaginationNextPageParams].
	NextPageParams map[string]any `json:"next_page_params"`
	JSON           struct {
		OfAgentRunResponsePaginationArray respjson.Field
		NextPageParams                    respjson.Field
		raw                               string
	} `json:"-"`
}

func (u AgentRunResponsePaginationUnion) AsAgentRunResponsePaginationNextPageParams() (v AgentRunResponsePaginationNextPageParams) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentRunResponsePaginationUnion) AsAgentRunResponsePaginationArray() (v []AgentRunResponsePaginationArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentRunResponsePaginationUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentRunResponsePaginationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResponsePaginationNextPageParams struct {
	NextPageParams map[string]any `json:"next_page_params,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResponsePaginationNextPageParams) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponsePaginationNextPageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResponsePaginationArrayItem struct {
	NextPageParams map[string]any `json:"next_page_params,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResponsePaginationArrayItem) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponsePaginationArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentListParams struct {
	// Number of results per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination offset
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	// Filter public templates by attribution
	//
	// Any of "nimble", "community".
	ManagedBy AgentListParamsManagedBy `query:"managed_by,omitzero" json:"-"`
	// Filter by privacy level
	//
	// Any of "public", "private", "all".
	Privacy AgentListParamsPrivacy `query:"privacy,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentListParams]'s query parameters as `url.Values`.
func (r AgentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter public templates by attribution
type AgentListParamsManagedBy string

const (
	AgentListParamsManagedByNimble    AgentListParamsManagedBy = "nimble"
	AgentListParamsManagedByCommunity AgentListParamsManagedBy = "community"
)

// Filter by privacy level
type AgentListParamsPrivacy string

const (
	AgentListParamsPrivacyPublic  AgentListParamsPrivacy = "public"
	AgentListParamsPrivacyPrivate AgentListParamsPrivacy = "private"
	AgentListParamsPrivacyAll     AgentListParamsPrivacy = "all"
)

type AgentAsyncParams struct {
	Agent  string         `json:"agent,required"`
	Params map[string]any `json:"params,omitzero,required"`
	// URL to call back when async operation completes
	CallbackURL  param.Opt[string] `json:"callback_url,omitzero"`
	Localization param.Opt[bool]   `json:"localization,omitzero"`
	// Whether to compress stored data
	StorageCompress param.Opt[bool] `json:"storage_compress,omitzero"`
	// Custom name for the stored object
	StorageObjectName param.Opt[string] `json:"storage_object_name,omitzero"`
	// Type of storage to use for results
	StorageType param.Opt[string] `json:"storage_type,omitzero"`
	// URL for storage location
	StorageURL param.Opt[string] `json:"storage_url,omitzero"`
	paramObj
}

func (r AgentAsyncParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentAsyncParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentAsyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunParams struct {
	Agent        string          `json:"agent,required"`
	Params       map[string]any  `json:"params,omitzero,required"`
	Localization param.Opt[bool] `json:"localization,omitzero"`
	paramObj
}

func (r AgentRunParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
