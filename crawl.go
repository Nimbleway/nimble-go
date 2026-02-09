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
)

// CrawlService contains methods and other services that help with interacting with
// the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCrawlService] method instead.
type CrawlService struct {
	Options []option.RequestOption
}

// NewCrawlService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewCrawlService(opts ...option.RequestOption) (r CrawlService) {
	r = CrawlService{}
	r.Options = opts
	return
}

// Crawl by Filter
func (r *CrawlService) List(ctx context.Context, query CrawlListParams, opts ...option.RequestOption) (res *CrawlListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/crawl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get crawl data
func (r *CrawlService) Status(ctx context.Context, id string, opts ...option.RequestOption) (res *CrawlStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/crawl/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Cancel Crawl
func (r *CrawlService) Terminate(ctx context.Context, id string, opts ...option.RequestOption) (res *CrawlTerminateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v1/crawl/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Successful get crawl response
type CrawlListResponse struct {
	Data       []CrawlListResponseData     `json:"data,required"`
	Pagination CrawlListResponsePagination `json:"pagination,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlListResponse) RawJSON() string { return r.JSON.raw }
func (r *CrawlListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlListResponseData struct {
	ID           string                              `json:"id,required" format:"uuid"`
	AccountName  string                              `json:"account_name,required"`
	CrawlOptions CrawlListResponseDataCrawlOptions   `json:"crawl_options,required"`
	CreatedAt    CrawlListResponseDataCreatedAtUnion `json:"created_at,required"`
	// Any of "queued", "running", "succeeded", "failed", "canceled".
	Status         string                                `json:"status,required"`
	UpdatedAt      CrawlListResponseDataUpdatedAtUnion   `json:"updated_at,required"`
	URL            string                                `json:"url,required" format:"uri"`
	Completed      float64                               `json:"completed"`
	CompletedAt    CrawlListResponseDataCompletedAtUnion `json:"completed_at,nullable"`
	EncryptedToken string                                `json:"encrypted_token,nullable"`
	ExtractOptions map[string]any                        `json:"extract_options,nullable"`
	Failed         float64                               `json:"failed"`
	Name           string                                `json:"name,nullable"`
	Pending        float64                               `json:"pending"`
	Tasks          []CrawlListResponseDataTask           `json:"tasks"`
	Total          float64                               `json:"total"`
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
func (r CrawlListResponseData) RawJSON() string { return r.JSON.raw }
func (r *CrawlListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlListResponseDataCrawlOptions struct {
	AllowExternalLinks    bool  `json:"allow_external_links,required"`
	AllowSubdomains       bool  `json:"allow_subdomains,required"`
	CrawlEntireDomain     bool  `json:"crawl_entire_domain,required"`
	IgnoreQueryParameters bool  `json:"ignore_query_parameters,required"`
	Limit                 int64 `json:"limit,required"`
	MaxDiscoveryDepth     int64 `json:"max_discovery_depth,required"`
	// Any of "skip", "include", "only".
	Sitemap      string                                         `json:"sitemap,required"`
	Callback     CrawlListResponseDataCrawlOptionsCallbackUnion `json:"callback" format:"uri"`
	ExcludePaths []string                                       `json:"exclude_paths"`
	IncludePaths []string                                       `json:"include_paths"`
	ExtraFields  map[string]any                                 `json:",extras"`
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
func (r CrawlListResponseDataCrawlOptions) RawJSON() string { return r.JSON.raw }
func (r *CrawlListResponseDataCrawlOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlListResponseDataCrawlOptionsCallbackUnion contains all possible properties
// and values from [CrawlListResponseDataCrawlOptionsCallbackObject], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type CrawlListResponseDataCrawlOptionsCallbackUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [CrawlListResponseDataCrawlOptionsCallbackObject].
	URL string `json:"url"`
	// This field is from variant [CrawlListResponseDataCrawlOptionsCallbackObject].
	Events []string `json:"events"`
	// This field is from variant [CrawlListResponseDataCrawlOptionsCallbackObject].
	Headers map[string]string `json:"headers"`
	// This field is from variant [CrawlListResponseDataCrawlOptionsCallbackObject].
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

func (u CrawlListResponseDataCrawlOptionsCallbackUnion) AsCrawlListResponseDataCrawlOptionsCallbackObject() (v CrawlListResponseDataCrawlOptionsCallbackObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlListResponseDataCrawlOptionsCallbackUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlListResponseDataCrawlOptionsCallbackUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlListResponseDataCrawlOptionsCallbackUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlListResponseDataCrawlOptionsCallbackObject struct {
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
func (r CrawlListResponseDataCrawlOptionsCallbackObject) RawJSON() string { return r.JSON.raw }
func (r *CrawlListResponseDataCrawlOptionsCallbackObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlListResponseDataCreatedAtUnion contains all possible properties and values
// from [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCrawlListResponseDataCreatedAtMapItem]
type CrawlListResponseDataCreatedAtUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfCrawlListResponseDataCreatedAtMapItem any `json:",inline"`
	JSON                                    struct {
		OfString                                respjson.Field
		OfCrawlListResponseDataCreatedAtMapItem respjson.Field
		raw                                     string
	} `json:"-"`
}

func (u CrawlListResponseDataCreatedAtUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlListResponseDataCreatedAtUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlListResponseDataCreatedAtUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlListResponseDataCreatedAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlListResponseDataUpdatedAtUnion contains all possible properties and values
// from [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCrawlListResponseDataUpdatedAtMapItem]
type CrawlListResponseDataUpdatedAtUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfCrawlListResponseDataUpdatedAtMapItem any `json:",inline"`
	JSON                                    struct {
		OfString                                respjson.Field
		OfCrawlListResponseDataUpdatedAtMapItem respjson.Field
		raw                                     string
	} `json:"-"`
}

func (u CrawlListResponseDataUpdatedAtUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlListResponseDataUpdatedAtUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlListResponseDataUpdatedAtUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlListResponseDataUpdatedAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlListResponseDataCompletedAtUnion contains all possible properties and
// values from [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCrawlListResponseDataCompletedAtMapItem]
type CrawlListResponseDataCompletedAtUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfCrawlListResponseDataCompletedAtMapItem any `json:",inline"`
	JSON                                      struct {
		OfString                                  respjson.Field
		OfCrawlListResponseDataCompletedAtMapItem respjson.Field
		raw                                       string
	} `json:"-"`
}

func (u CrawlListResponseDataCompletedAtUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlListResponseDataCompletedAtUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlListResponseDataCompletedAtUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlListResponseDataCompletedAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlListResponseDataTask struct {
	CrawlID string `json:"crawl_id,required" format:"uuid"`
	// Any of "pending", "completed", "failed".
	Status      string                                  `json:"status,required"`
	WebitTaskID string                                  `json:"webit_task_id,required"`
	CreatedAt   CrawlListResponseDataTaskCreatedAtUnion `json:"created_at"`
	UpdatedAt   CrawlListResponseDataTaskUpdatedAtUnion `json:"updated_at"`
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
func (r CrawlListResponseDataTask) RawJSON() string { return r.JSON.raw }
func (r *CrawlListResponseDataTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlListResponseDataTaskCreatedAtUnion contains all possible properties and
// values from [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCrawlListResponseDataTaskCreatedAtMapItem]
type CrawlListResponseDataTaskCreatedAtUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfCrawlListResponseDataTaskCreatedAtMapItem any `json:",inline"`
	JSON                                        struct {
		OfString                                    respjson.Field
		OfCrawlListResponseDataTaskCreatedAtMapItem respjson.Field
		raw                                         string
	} `json:"-"`
}

func (u CrawlListResponseDataTaskCreatedAtUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlListResponseDataTaskCreatedAtUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlListResponseDataTaskCreatedAtUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlListResponseDataTaskCreatedAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlListResponseDataTaskUpdatedAtUnion contains all possible properties and
// values from [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCrawlListResponseDataTaskUpdatedAtMapItem]
type CrawlListResponseDataTaskUpdatedAtUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfCrawlListResponseDataTaskUpdatedAtMapItem any `json:",inline"`
	JSON                                        struct {
		OfString                                    respjson.Field
		OfCrawlListResponseDataTaskUpdatedAtMapItem respjson.Field
		raw                                         string
	} `json:"-"`
}

func (u CrawlListResponseDataTaskUpdatedAtUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlListResponseDataTaskUpdatedAtUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlListResponseDataTaskUpdatedAtUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlListResponseDataTaskUpdatedAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlListResponsePagination struct {
	HasNext    bool    `json:"has_next,required"`
	NextCursor string  `json:"next_cursor,nullable"`
	Total      float64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNext     respjson.Field
		NextCursor  respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *CrawlListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlStatusResponse map[string]any

type CrawlTerminateResponse struct {
	// Any of "canceled".
	Status CrawlTerminateResponseStatus `json:"status,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlTerminateResponse) RawJSON() string { return r.JSON.raw }
func (r *CrawlTerminateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlTerminateResponseStatus string

const (
	CrawlTerminateResponseStatusCanceled CrawlTerminateResponseStatus = "canceled"
)

type CrawlListParams struct {
	// Filter crawls by their status.
	//
	// Any of "queued", "running", "succeeded", "failed", "canceled".
	Status CrawlListParamsStatus `query:"status,omitzero,required" json:"-"`
	// Cursor for pagination.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Number of crawls to return per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [CrawlListParams]'s query parameters as `url.Values`.
func (r CrawlListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter crawls by their status.
type CrawlListParamsStatus string

const (
	CrawlListParamsStatusQueued    CrawlListParamsStatus = "queued"
	CrawlListParamsStatusRunning   CrawlListParamsStatus = "running"
	CrawlListParamsStatusSucceeded CrawlListParamsStatus = "succeeded"
	CrawlListParamsStatusFailed    CrawlListParamsStatus = "failed"
	CrawlListParamsStatusCanceled  CrawlListParamsStatus = "canceled"
)
