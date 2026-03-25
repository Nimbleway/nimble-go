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
	"github.com/Nimbleway/nimble-go/shared"
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
	return res, err
}

// Create crawl task
func (r *CrawlService) Run(ctx context.Context, body CrawlRunParams, opts ...option.RequestOption) (res *CrawlRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/crawl"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get crawl data
func (r *CrawlService) Status(ctx context.Context, id string, opts ...option.RequestOption) (res *CrawlStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/crawl/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Cancel Crawl
func (r *CrawlService) Terminate(ctx context.Context, id string, opts ...option.RequestOption) (res *CrawlTerminateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/crawl/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Successful get crawl response
type CrawlListResponse struct {
	Data       []CrawlListResponseData     `json:"data" api:"required"`
	Pagination CrawlListResponsePagination `json:"pagination" api:"required"`
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

// Crawl API response
type CrawlListResponseData struct {
	AccountName  string                              `json:"account_name" api:"required"`
	CrawlID      string                              `json:"crawl_id" api:"required" format:"uuid"`
	CrawlOptions CrawlListResponseDataCrawlOptions   `json:"crawl_options" api:"required"`
	CreatedAt    CrawlListResponseDataCreatedAtUnion `json:"created_at" api:"required"`
	// Any of "queued", "running", "succeeded", "failed", "canceled".
	Status         string                                `json:"status" api:"required"`
	UpdatedAt      CrawlListResponseDataUpdatedAtUnion   `json:"updated_at" api:"required"`
	URL            string                                `json:"url" api:"required" format:"uri"`
	Completed      float64                               `json:"completed"`
	CompletedAt    CrawlListResponseDataCompletedAtUnion `json:"completed_at" api:"nullable"`
	ExtractOptions map[string]any                        `json:"extract_options" api:"nullable"`
	Failed         float64                               `json:"failed"`
	Name           string                                `json:"name" api:"nullable"`
	Pending        float64                               `json:"pending"`
	Tasks          []CrawlListResponseDataTask           `json:"tasks"`
	Total          float64                               `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountName    respjson.Field
		CrawlID        respjson.Field
		CrawlOptions   respjson.Field
		CreatedAt      respjson.Field
		Status         respjson.Field
		UpdatedAt      respjson.Field
		URL            respjson.Field
		Completed      respjson.Field
		CompletedAt    respjson.Field
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
	AllowExternalLinks    bool  `json:"allow_external_links" api:"required"`
	AllowSubdomains       bool  `json:"allow_subdomains" api:"required"`
	CrawlEntireDomain     bool  `json:"crawl_entire_domain" api:"required"`
	IgnoreQueryParameters bool  `json:"ignore_query_parameters" api:"required"`
	Limit                 int64 `json:"limit" api:"required"`
	MaxDiscoveryDepth     int64 `json:"max_discovery_depth" api:"required"`
	// Any of "skip", "include", "only".
	Sitemap      string                                         `json:"sitemap" api:"required"`
	Callback     CrawlListResponseDataCrawlOptionsCallbackUnion `json:"callback" format:"uri"`
	ExcludePaths []string                                       `json:"exclude_paths"`
	IncludePaths []string                                       `json:"include_paths"`
	ExtraFields  map[string]any                                 `json:"" api:"extrafields"`
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
	URL string `json:"url" api:"required" format:"uri"`
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
	// Any of "pending", "completed", "failed".
	Status    string `json:"status" api:"required"`
	TaskID    string `json:"task_id" api:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		TaskID      respjson.Field
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

type CrawlListResponsePagination struct {
	HasNext    bool    `json:"has_next" api:"required"`
	NextCursor string  `json:"next_cursor" api:"nullable"`
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

// Crawl API response
type CrawlRunResponse struct {
	AccountName  string                         `json:"account_name" api:"required"`
	CrawlID      string                         `json:"crawl_id" api:"required" format:"uuid"`
	CrawlOptions CrawlRunResponseCrawlOptions   `json:"crawl_options" api:"required"`
	CreatedAt    CrawlRunResponseCreatedAtUnion `json:"created_at" api:"required"`
	// Any of "queued", "running", "succeeded", "failed", "canceled".
	Status         CrawlRunResponseStatus           `json:"status" api:"required"`
	UpdatedAt      CrawlRunResponseUpdatedAtUnion   `json:"updated_at" api:"required"`
	URL            string                           `json:"url" api:"required" format:"uri"`
	Completed      float64                          `json:"completed"`
	CompletedAt    CrawlRunResponseCompletedAtUnion `json:"completed_at" api:"nullable"`
	ExtractOptions map[string]any                   `json:"extract_options" api:"nullable"`
	Failed         float64                          `json:"failed"`
	Name           string                           `json:"name" api:"nullable"`
	Pending        float64                          `json:"pending"`
	Tasks          []CrawlRunResponseTask           `json:"tasks"`
	Total          float64                          `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountName    respjson.Field
		CrawlID        respjson.Field
		CrawlOptions   respjson.Field
		CreatedAt      respjson.Field
		Status         respjson.Field
		UpdatedAt      respjson.Field
		URL            respjson.Field
		Completed      respjson.Field
		CompletedAt    respjson.Field
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
func (r CrawlRunResponse) RawJSON() string { return r.JSON.raw }
func (r *CrawlRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlRunResponseCrawlOptions struct {
	AllowExternalLinks    bool  `json:"allow_external_links" api:"required"`
	AllowSubdomains       bool  `json:"allow_subdomains" api:"required"`
	CrawlEntireDomain     bool  `json:"crawl_entire_domain" api:"required"`
	IgnoreQueryParameters bool  `json:"ignore_query_parameters" api:"required"`
	Limit                 int64 `json:"limit" api:"required"`
	MaxDiscoveryDepth     int64 `json:"max_discovery_depth" api:"required"`
	// Any of "skip", "include", "only".
	Sitemap      string                                    `json:"sitemap" api:"required"`
	Callback     CrawlRunResponseCrawlOptionsCallbackUnion `json:"callback" format:"uri"`
	ExcludePaths []string                                  `json:"exclude_paths"`
	IncludePaths []string                                  `json:"include_paths"`
	ExtraFields  map[string]any                            `json:"" api:"extrafields"`
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
func (r CrawlRunResponseCrawlOptions) RawJSON() string { return r.JSON.raw }
func (r *CrawlRunResponseCrawlOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlRunResponseCrawlOptionsCallbackUnion contains all possible properties and
// values from [CrawlRunResponseCrawlOptionsCallbackObject], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type CrawlRunResponseCrawlOptionsCallbackUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [CrawlRunResponseCrawlOptionsCallbackObject].
	URL string `json:"url"`
	// This field is from variant [CrawlRunResponseCrawlOptionsCallbackObject].
	Events []string `json:"events"`
	// This field is from variant [CrawlRunResponseCrawlOptionsCallbackObject].
	Headers map[string]string `json:"headers"`
	// This field is from variant [CrawlRunResponseCrawlOptionsCallbackObject].
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

func (u CrawlRunResponseCrawlOptionsCallbackUnion) AsCrawlRunResponseCrawlOptionsCallbackObject() (v CrawlRunResponseCrawlOptionsCallbackObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlRunResponseCrawlOptionsCallbackUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlRunResponseCrawlOptionsCallbackUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlRunResponseCrawlOptionsCallbackUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlRunResponseCrawlOptionsCallbackObject struct {
	URL string `json:"url" api:"required" format:"uri"`
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
func (r CrawlRunResponseCrawlOptionsCallbackObject) RawJSON() string { return r.JSON.raw }
func (r *CrawlRunResponseCrawlOptionsCallbackObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlRunResponseCreatedAtUnion contains all possible properties and values from
// [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCrawlRunResponseCreatedAtMapItem]
type CrawlRunResponseCreatedAtUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfCrawlRunResponseCreatedAtMapItem any `json:",inline"`
	JSON                               struct {
		OfString                           respjson.Field
		OfCrawlRunResponseCreatedAtMapItem respjson.Field
		raw                                string
	} `json:"-"`
}

func (u CrawlRunResponseCreatedAtUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlRunResponseCreatedAtUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlRunResponseCreatedAtUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlRunResponseCreatedAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlRunResponseStatus string

const (
	CrawlRunResponseStatusQueued    CrawlRunResponseStatus = "queued"
	CrawlRunResponseStatusRunning   CrawlRunResponseStatus = "running"
	CrawlRunResponseStatusSucceeded CrawlRunResponseStatus = "succeeded"
	CrawlRunResponseStatusFailed    CrawlRunResponseStatus = "failed"
	CrawlRunResponseStatusCanceled  CrawlRunResponseStatus = "canceled"
)

// CrawlRunResponseUpdatedAtUnion contains all possible properties and values from
// [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCrawlRunResponseUpdatedAtMapItem]
type CrawlRunResponseUpdatedAtUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfCrawlRunResponseUpdatedAtMapItem any `json:",inline"`
	JSON                               struct {
		OfString                           respjson.Field
		OfCrawlRunResponseUpdatedAtMapItem respjson.Field
		raw                                string
	} `json:"-"`
}

func (u CrawlRunResponseUpdatedAtUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlRunResponseUpdatedAtUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlRunResponseUpdatedAtUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlRunResponseUpdatedAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlRunResponseCompletedAtUnion contains all possible properties and values
// from [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCrawlRunResponseCompletedAtMapItem]
type CrawlRunResponseCompletedAtUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfCrawlRunResponseCompletedAtMapItem any `json:",inline"`
	JSON                                 struct {
		OfString                             respjson.Field
		OfCrawlRunResponseCompletedAtMapItem respjson.Field
		raw                                  string
	} `json:"-"`
}

func (u CrawlRunResponseCompletedAtUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlRunResponseCompletedAtUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlRunResponseCompletedAtUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlRunResponseCompletedAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlRunResponseTask struct {
	// Any of "pending", "completed", "failed".
	Status    string `json:"status" api:"required"`
	TaskID    string `json:"task_id" api:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		TaskID      respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlRunResponseTask) RawJSON() string { return r.JSON.raw }
func (r *CrawlRunResponseTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Crawl API response
type CrawlStatusResponse struct {
	AccountName  string                            `json:"account_name" api:"required"`
	CrawlID      string                            `json:"crawl_id" api:"required" format:"uuid"`
	CrawlOptions CrawlStatusResponseCrawlOptions   `json:"crawl_options" api:"required"`
	CreatedAt    CrawlStatusResponseCreatedAtUnion `json:"created_at" api:"required"`
	// Any of "queued", "running", "succeeded", "failed", "canceled".
	Status         CrawlStatusResponseStatus           `json:"status" api:"required"`
	UpdatedAt      CrawlStatusResponseUpdatedAtUnion   `json:"updated_at" api:"required"`
	URL            string                              `json:"url" api:"required" format:"uri"`
	Completed      float64                             `json:"completed"`
	CompletedAt    CrawlStatusResponseCompletedAtUnion `json:"completed_at" api:"nullable"`
	ExtractOptions map[string]any                      `json:"extract_options" api:"nullable"`
	Failed         float64                             `json:"failed"`
	Name           string                              `json:"name" api:"nullable"`
	Pending        float64                             `json:"pending"`
	Tasks          []CrawlStatusResponseTask           `json:"tasks"`
	Total          float64                             `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountName    respjson.Field
		CrawlID        respjson.Field
		CrawlOptions   respjson.Field
		CreatedAt      respjson.Field
		Status         respjson.Field
		UpdatedAt      respjson.Field
		URL            respjson.Field
		Completed      respjson.Field
		CompletedAt    respjson.Field
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
func (r CrawlStatusResponse) RawJSON() string { return r.JSON.raw }
func (r *CrawlStatusResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlStatusResponseCrawlOptions struct {
	AllowExternalLinks    bool  `json:"allow_external_links" api:"required"`
	AllowSubdomains       bool  `json:"allow_subdomains" api:"required"`
	CrawlEntireDomain     bool  `json:"crawl_entire_domain" api:"required"`
	IgnoreQueryParameters bool  `json:"ignore_query_parameters" api:"required"`
	Limit                 int64 `json:"limit" api:"required"`
	MaxDiscoveryDepth     int64 `json:"max_discovery_depth" api:"required"`
	// Any of "skip", "include", "only".
	Sitemap      string                                       `json:"sitemap" api:"required"`
	Callback     CrawlStatusResponseCrawlOptionsCallbackUnion `json:"callback" format:"uri"`
	ExcludePaths []string                                     `json:"exclude_paths"`
	IncludePaths []string                                     `json:"include_paths"`
	ExtraFields  map[string]any                               `json:"" api:"extrafields"`
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
func (r CrawlStatusResponseCrawlOptions) RawJSON() string { return r.JSON.raw }
func (r *CrawlStatusResponseCrawlOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlStatusResponseCrawlOptionsCallbackUnion contains all possible properties
// and values from [CrawlStatusResponseCrawlOptionsCallbackObject], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString]
type CrawlStatusResponseCrawlOptionsCallbackUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field is from variant [CrawlStatusResponseCrawlOptionsCallbackObject].
	URL string `json:"url"`
	// This field is from variant [CrawlStatusResponseCrawlOptionsCallbackObject].
	Events []string `json:"events"`
	// This field is from variant [CrawlStatusResponseCrawlOptionsCallbackObject].
	Headers map[string]string `json:"headers"`
	// This field is from variant [CrawlStatusResponseCrawlOptionsCallbackObject].
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

func (u CrawlStatusResponseCrawlOptionsCallbackUnion) AsCrawlStatusResponseCrawlOptionsCallbackObject() (v CrawlStatusResponseCrawlOptionsCallbackObject) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlStatusResponseCrawlOptionsCallbackUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlStatusResponseCrawlOptionsCallbackUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlStatusResponseCrawlOptionsCallbackUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlStatusResponseCrawlOptionsCallbackObject struct {
	URL string `json:"url" api:"required" format:"uri"`
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
func (r CrawlStatusResponseCrawlOptionsCallbackObject) RawJSON() string { return r.JSON.raw }
func (r *CrawlStatusResponseCrawlOptionsCallbackObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlStatusResponseCreatedAtUnion contains all possible properties and values
// from [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCrawlStatusResponseCreatedAtMapItem]
type CrawlStatusResponseCreatedAtUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfCrawlStatusResponseCreatedAtMapItem any `json:",inline"`
	JSON                                  struct {
		OfString                              respjson.Field
		OfCrawlStatusResponseCreatedAtMapItem respjson.Field
		raw                                   string
	} `json:"-"`
}

func (u CrawlStatusResponseCreatedAtUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlStatusResponseCreatedAtUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlStatusResponseCreatedAtUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlStatusResponseCreatedAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlStatusResponseStatus string

const (
	CrawlStatusResponseStatusQueued    CrawlStatusResponseStatus = "queued"
	CrawlStatusResponseStatusRunning   CrawlStatusResponseStatus = "running"
	CrawlStatusResponseStatusSucceeded CrawlStatusResponseStatus = "succeeded"
	CrawlStatusResponseStatusFailed    CrawlStatusResponseStatus = "failed"
	CrawlStatusResponseStatusCanceled  CrawlStatusResponseStatus = "canceled"
)

// CrawlStatusResponseUpdatedAtUnion contains all possible properties and values
// from [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCrawlStatusResponseUpdatedAtMapItem]
type CrawlStatusResponseUpdatedAtUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfCrawlStatusResponseUpdatedAtMapItem any `json:",inline"`
	JSON                                  struct {
		OfString                              respjson.Field
		OfCrawlStatusResponseUpdatedAtMapItem respjson.Field
		raw                                   string
	} `json:"-"`
}

func (u CrawlStatusResponseUpdatedAtUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlStatusResponseUpdatedAtUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlStatusResponseUpdatedAtUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlStatusResponseUpdatedAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// CrawlStatusResponseCompletedAtUnion contains all possible properties and values
// from [string], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfCrawlStatusResponseCompletedAtMapItem]
type CrawlStatusResponseCompletedAtUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfCrawlStatusResponseCompletedAtMapItem any `json:",inline"`
	JSON                                    struct {
		OfString                                respjson.Field
		OfCrawlStatusResponseCompletedAtMapItem respjson.Field
		raw                                     string
	} `json:"-"`
}

func (u CrawlStatusResponseCompletedAtUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CrawlStatusResponseCompletedAtUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CrawlStatusResponseCompletedAtUnion) RawJSON() string { return u.JSON.raw }

func (r *CrawlStatusResponseCompletedAtUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlStatusResponseTask struct {
	// Any of "pending", "completed", "failed".
	Status    string `json:"status" api:"required"`
	TaskID    string `json:"task_id" api:"required"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		TaskID      respjson.Field
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlStatusResponseTask) RawJSON() string { return r.JSON.raw }
func (r *CrawlStatusResponseTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlTerminateResponse struct {
	// Any of "canceled".
	Status CrawlTerminateResponseStatus `json:"status" api:"required"`
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
	// Cursor for pagination.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Number of crawls to return per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Filter crawls by their status.
	//
	// Any of "queued", "running", "succeeded", "failed", "canceled", "all".
	Status CrawlListParamsStatus `query:"status,omitzero" json:"-"`
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
	CrawlListParamsStatusAll       CrawlListParamsStatus = "all"
)

type CrawlRunParams struct {
	// Url to crawl.
	URL string `json:"url" api:"required"`
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
	Callback CrawlRunParamsCallbackUnion `json:"callback,omitzero" format:"uri"`
	// URL pathname regex patterns that exclude matching URLs from the crawl.
	ExcludePaths   []string                     `json:"exclude_paths,omitzero"`
	ExtractOptions CrawlRunParamsExtractOptions `json:"extract_options,omitzero"`
	// URL pathname regex patterns that include matching URLs in the crawl.
	IncludePaths []string `json:"include_paths,omitzero"`
	// Sitemap and other methods will be used together to find URLs.
	//
	// Any of "skip", "include", "only".
	Sitemap CrawlRunParamsSitemap `json:"sitemap,omitzero"`
	paramObj
}

func (r CrawlRunParams) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRunParamsCallbackUnion struct {
	OfCrawlRunsCallbackObject *CrawlRunParamsCallbackObject `json:",omitzero,inline"`
	OfString                  param.Opt[string]             `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRunParamsCallbackUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlRunsCallbackObject, u.OfString)
}
func (u *CrawlRunParamsCallbackUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRunParamsCallbackUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlRunsCallbackObject) {
		return u.OfCrawlRunsCallbackObject
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// The property URL is required.
type CrawlRunParamsCallbackObject struct {
	URL string `json:"url" api:"required" format:"uri"`
	// Any of "started", "page", "completed", "failed".
	Events   []string          `json:"events,omitzero"`
	Headers  map[string]string `json:"headers,omitzero"`
	Metadata map[string]any    `json:"metadata,omitzero"`
	paramObj
}

func (r CrawlRunParamsCallbackObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRunParamsCallbackObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRunParamsCallbackObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlRunParamsExtractOptions struct {
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
	// Target URL to scrape
	URL param.Opt[string] `json:"url,omitzero"`
	// Browser type to emulate
	Browser CrawlRunParamsExtractOptionsBrowserUnion `json:"browser,omitzero"`
	// Array of browser automation actions to execute sequentially
	BrowserActions []CrawlRunParamsExtractOptionsBrowserActionUnion `json:"browser_actions,omitzero"`
	// Browser cookies as array of cookie objects
	Cookies CrawlRunParamsExtractOptionsCookiesUnion `json:"cookies,omitzero"`
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
	Country CrawlRunParamsExtractOptionsCountry `json:"country,omitzero"`
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
	Headers map[string]CrawlRunParamsExtractOptionsHeaderUnion `json:"headers,omitzero"`
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
	Locale CrawlRunParamsExtractOptionsLocale `json:"locale,omitzero"`
	// HTTP method for the request
	//
	// Any of "GET", "POST", "PUT", "PATCH", "DELETE".
	Method string `json:"method,omitzero"`
	// Filters for capturing network traffic
	NetworkCapture []CrawlRunParamsExtractOptionsNetworkCapture `json:"network_capture,omitzero"`
	// Operating system to emulate
	//
	// Any of "windows", "mac os", "linux", "android", "ios".
	Os string `json:"os,omitzero"`
	// Custom parser configuration as a key-value map
	Parser CrawlRunParamsExtractOptionsParserUnion `json:"parser,omitzero"`
	// Referrer policy for the request
	//
	// Any of "random", "no-referer", "same-origin", "google", "bing", "facebook",
	// "twitter", "instagram".
	ReferrerType CrawlRunParamsExtractOptionsReferrerType `json:"referrer_type,omitzero"`
	Session      CrawlRunParamsExtractOptionsSession      `json:"session,omitzero"`
	// Skills or capabilities required for the request
	Skill CrawlRunParamsExtractOptionsSkillUnion `json:"skill,omitzero"`
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

func (r CrawlRunParamsExtractOptions) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRunParamsExtractOptions
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRunParamsExtractOptions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlRunParamsExtractOptions](
		"device", "desktop", "mobile", "tablet",
	)
	apijson.RegisterFieldValidator[CrawlRunParamsExtractOptions](
		"driver", "vx6", "vx8", "vx8-pro", "vx10", "vx10-pro", "vx12", "vx12-pro",
	)
	apijson.RegisterFieldValidator[CrawlRunParamsExtractOptions](
		"method", "GET", "POST", "PUT", "PATCH", "DELETE",
	)
	apijson.RegisterFieldValidator[CrawlRunParamsExtractOptions](
		"os", "windows", "mac os", "linux", "android", "ios",
	)
	apijson.RegisterFieldValidator[CrawlRunParamsExtractOptions](
		"state", "AL", "AK", "AS", "AZ", "AR", "CA", "CO", "CT", "DE", "DC", "FL", "GA", "GU", "HI", "ID", "IL", "IN", "IA", "KS", "KY", "LA", "ME", "MD", "MA", "MI", "MN", "MS", "MO", "MT", "NE", "NV", "NH", "NJ", "NM", "NY", "NC", "ND", "MP", "OH", "OK", "OR", "PA", "PR", "RI", "SC", "SD", "TN", "TX", "UT", "VT", "VA", "VI", "WA", "WV", "WI", "WY",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRunParamsExtractOptionsBrowserUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlRunsExtractOptionsBrowserString)
	OfCrawlRunsExtractOptionsBrowserString param.Opt[string]                          `json:",omitzero,inline"`
	OfCrawlRunsExtractOptionsBrowserObject *CrawlRunParamsExtractOptionsBrowserObject `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRunParamsExtractOptionsBrowserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlRunsExtractOptionsBrowserString, u.OfCrawlRunsExtractOptionsBrowserObject)
}
func (u *CrawlRunParamsExtractOptionsBrowserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRunParamsExtractOptionsBrowserUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlRunsExtractOptionsBrowserString) {
		return &u.OfCrawlRunsExtractOptionsBrowserString
	} else if !param.IsOmitted(u.OfCrawlRunsExtractOptionsBrowserObject) {
		return u.OfCrawlRunsExtractOptionsBrowserObject
	}
	return nil
}

// Browser type to emulate
type CrawlRunParamsExtractOptionsBrowserString string

const (
	CrawlRunParamsExtractOptionsBrowserStringChrome  CrawlRunParamsExtractOptionsBrowserString = "chrome"
	CrawlRunParamsExtractOptionsBrowserStringFirefox CrawlRunParamsExtractOptionsBrowserString = "firefox"
)

// The property Name is required.
type CrawlRunParamsExtractOptionsBrowserObject struct {
	// Any of "chrome", "firefox".
	Name string `json:"name,omitzero" api:"required"`
	// Specific browser version to emulate
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r CrawlRunParamsExtractOptionsBrowserObject) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRunParamsExtractOptionsBrowserObject
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRunParamsExtractOptionsBrowserObject) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlRunParamsExtractOptionsBrowserObject](
		"name", "chrome", "firefox",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRunParamsExtractOptionsBrowserActionUnion struct {
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

func (u CrawlRunParamsExtractOptionsBrowserActionUnion) MarshalJSON() ([]byte, error) {
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
func (u *CrawlRunParamsExtractOptionsBrowserActionUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRunParamsExtractOptionsBrowserActionUnion) asAny() any {
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
type CrawlRunParamsExtractOptionsCookiesUnion struct {
	OfCrawlRunsExtractOptionsCookiesArray []CrawlRunParamsExtractOptionsCookiesArrayItem `json:",omitzero,inline"`
	OfString                              param.Opt[string]                              `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRunParamsExtractOptionsCookiesUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlRunsExtractOptionsCookiesArray, u.OfString)
}
func (u *CrawlRunParamsExtractOptionsCookiesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRunParamsExtractOptionsCookiesUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlRunsExtractOptionsCookiesArray) {
		return &u.OfCrawlRunsExtractOptionsCookiesArray
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

type CrawlRunParamsExtractOptionsCookiesArrayItem struct {
	Creation      param.Opt[string]                                       `json:"creation,omitzero"`
	Domain        param.Opt[string]                                       `json:"domain,omitzero"`
	HostOnly      param.Opt[bool]                                         `json:"hostOnly,omitzero"`
	HTTPOnly      param.Opt[bool]                                         `json:"httpOnly,omitzero"`
	LastAccessed  param.Opt[string]                                       `json:"lastAccessed,omitzero"`
	Path          param.Opt[string]                                       `json:"path,omitzero"`
	PathIsDefault param.Opt[bool]                                         `json:"pathIsDefault,omitzero"`
	Expires       param.Opt[string]                                       `json:"expires,omitzero"`
	Name          param.Opt[string]                                       `json:"name,omitzero"`
	Secure        param.Opt[bool]                                         `json:"secure,omitzero"`
	Value         param.Opt[string]                                       `json:"value,omitzero"`
	Extensions    []string                                                `json:"extensions,omitzero"`
	MaxAge        CrawlRunParamsExtractOptionsCookiesArrayItemMaxAgeUnion `json:"maxAge,omitzero"`
	// Any of "strict", "lax", "none".
	SameSite    string         `json:"sameSite,omitzero"`
	ExtraFields map[string]any `json:"-"`
	paramObj
}

func (r CrawlRunParamsExtractOptionsCookiesArrayItem) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRunParamsExtractOptionsCookiesArrayItem
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *CrawlRunParamsExtractOptionsCookiesArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlRunParamsExtractOptionsCookiesArrayItem](
		"sameSite", "strict", "lax", "none",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRunParamsExtractOptionsCookiesArrayItemMaxAgeUnion struct {
	// Check if union is this variant with
	// !param.IsOmitted(union.OfCrawlRunsExtractOptionsCookiesArrayItemMaxAgeString)
	OfCrawlRunsExtractOptionsCookiesArrayItemMaxAgeString param.Opt[CrawlRunParamsExtractOptionsCookiesArrayItemMaxAgeString] `json:",omitzero,inline"`
	OfFloat                                               param.Opt[float64]                                                  `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRunParamsExtractOptionsCookiesArrayItemMaxAgeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCrawlRunsExtractOptionsCookiesArrayItemMaxAgeString, u.OfFloat)
}
func (u *CrawlRunParamsExtractOptionsCookiesArrayItemMaxAgeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRunParamsExtractOptionsCookiesArrayItemMaxAgeUnion) asAny() any {
	if !param.IsOmitted(u.OfCrawlRunsExtractOptionsCookiesArrayItemMaxAgeString) {
		return &u.OfCrawlRunsExtractOptionsCookiesArrayItemMaxAgeString
	} else if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	}
	return nil
}

type CrawlRunParamsExtractOptionsCookiesArrayItemMaxAgeString string

const (
	CrawlRunParamsExtractOptionsCookiesArrayItemMaxAgeStringInfinity      CrawlRunParamsExtractOptionsCookiesArrayItemMaxAgeString = "Infinity"
	CrawlRunParamsExtractOptionsCookiesArrayItemMaxAgeStringMinusInfinity CrawlRunParamsExtractOptionsCookiesArrayItemMaxAgeString = "-Infinity"
)

// Country code for geolocation and proxy selection
type CrawlRunParamsExtractOptionsCountry string

const (
	CrawlRunParamsExtractOptionsCountryAd  CrawlRunParamsExtractOptionsCountry = "AD"
	CrawlRunParamsExtractOptionsCountryAe  CrawlRunParamsExtractOptionsCountry = "AE"
	CrawlRunParamsExtractOptionsCountryAf  CrawlRunParamsExtractOptionsCountry = "AF"
	CrawlRunParamsExtractOptionsCountryAg  CrawlRunParamsExtractOptionsCountry = "AG"
	CrawlRunParamsExtractOptionsCountryAI  CrawlRunParamsExtractOptionsCountry = "AI"
	CrawlRunParamsExtractOptionsCountryAl  CrawlRunParamsExtractOptionsCountry = "AL"
	CrawlRunParamsExtractOptionsCountryAm  CrawlRunParamsExtractOptionsCountry = "AM"
	CrawlRunParamsExtractOptionsCountryAo  CrawlRunParamsExtractOptionsCountry = "AO"
	CrawlRunParamsExtractOptionsCountryAq  CrawlRunParamsExtractOptionsCountry = "AQ"
	CrawlRunParamsExtractOptionsCountryAr  CrawlRunParamsExtractOptionsCountry = "AR"
	CrawlRunParamsExtractOptionsCountryAs  CrawlRunParamsExtractOptionsCountry = "AS"
	CrawlRunParamsExtractOptionsCountryAt  CrawlRunParamsExtractOptionsCountry = "AT"
	CrawlRunParamsExtractOptionsCountryAu  CrawlRunParamsExtractOptionsCountry = "AU"
	CrawlRunParamsExtractOptionsCountryAw  CrawlRunParamsExtractOptionsCountry = "AW"
	CrawlRunParamsExtractOptionsCountryAx  CrawlRunParamsExtractOptionsCountry = "AX"
	CrawlRunParamsExtractOptionsCountryAz  CrawlRunParamsExtractOptionsCountry = "AZ"
	CrawlRunParamsExtractOptionsCountryBa  CrawlRunParamsExtractOptionsCountry = "BA"
	CrawlRunParamsExtractOptionsCountryBb  CrawlRunParamsExtractOptionsCountry = "BB"
	CrawlRunParamsExtractOptionsCountryBd  CrawlRunParamsExtractOptionsCountry = "BD"
	CrawlRunParamsExtractOptionsCountryBe  CrawlRunParamsExtractOptionsCountry = "BE"
	CrawlRunParamsExtractOptionsCountryBf  CrawlRunParamsExtractOptionsCountry = "BF"
	CrawlRunParamsExtractOptionsCountryBg  CrawlRunParamsExtractOptionsCountry = "BG"
	CrawlRunParamsExtractOptionsCountryBh  CrawlRunParamsExtractOptionsCountry = "BH"
	CrawlRunParamsExtractOptionsCountryBi  CrawlRunParamsExtractOptionsCountry = "BI"
	CrawlRunParamsExtractOptionsCountryBj  CrawlRunParamsExtractOptionsCountry = "BJ"
	CrawlRunParamsExtractOptionsCountryBl  CrawlRunParamsExtractOptionsCountry = "BL"
	CrawlRunParamsExtractOptionsCountryBm  CrawlRunParamsExtractOptionsCountry = "BM"
	CrawlRunParamsExtractOptionsCountryBn  CrawlRunParamsExtractOptionsCountry = "BN"
	CrawlRunParamsExtractOptionsCountryBo  CrawlRunParamsExtractOptionsCountry = "BO"
	CrawlRunParamsExtractOptionsCountryBq  CrawlRunParamsExtractOptionsCountry = "BQ"
	CrawlRunParamsExtractOptionsCountryBr  CrawlRunParamsExtractOptionsCountry = "BR"
	CrawlRunParamsExtractOptionsCountryBs  CrawlRunParamsExtractOptionsCountry = "BS"
	CrawlRunParamsExtractOptionsCountryBt  CrawlRunParamsExtractOptionsCountry = "BT"
	CrawlRunParamsExtractOptionsCountryBv  CrawlRunParamsExtractOptionsCountry = "BV"
	CrawlRunParamsExtractOptionsCountryBw  CrawlRunParamsExtractOptionsCountry = "BW"
	CrawlRunParamsExtractOptionsCountryBy  CrawlRunParamsExtractOptionsCountry = "BY"
	CrawlRunParamsExtractOptionsCountryBz  CrawlRunParamsExtractOptionsCountry = "BZ"
	CrawlRunParamsExtractOptionsCountryCa  CrawlRunParamsExtractOptionsCountry = "CA"
	CrawlRunParamsExtractOptionsCountryCc  CrawlRunParamsExtractOptionsCountry = "CC"
	CrawlRunParamsExtractOptionsCountryCd  CrawlRunParamsExtractOptionsCountry = "CD"
	CrawlRunParamsExtractOptionsCountryCf  CrawlRunParamsExtractOptionsCountry = "CF"
	CrawlRunParamsExtractOptionsCountryCg  CrawlRunParamsExtractOptionsCountry = "CG"
	CrawlRunParamsExtractOptionsCountryCh  CrawlRunParamsExtractOptionsCountry = "CH"
	CrawlRunParamsExtractOptionsCountryCi  CrawlRunParamsExtractOptionsCountry = "CI"
	CrawlRunParamsExtractOptionsCountryCk  CrawlRunParamsExtractOptionsCountry = "CK"
	CrawlRunParamsExtractOptionsCountryCl  CrawlRunParamsExtractOptionsCountry = "CL"
	CrawlRunParamsExtractOptionsCountryCm  CrawlRunParamsExtractOptionsCountry = "CM"
	CrawlRunParamsExtractOptionsCountryCn  CrawlRunParamsExtractOptionsCountry = "CN"
	CrawlRunParamsExtractOptionsCountryCo  CrawlRunParamsExtractOptionsCountry = "CO"
	CrawlRunParamsExtractOptionsCountryCr  CrawlRunParamsExtractOptionsCountry = "CR"
	CrawlRunParamsExtractOptionsCountryCu  CrawlRunParamsExtractOptionsCountry = "CU"
	CrawlRunParamsExtractOptionsCountryCv  CrawlRunParamsExtractOptionsCountry = "CV"
	CrawlRunParamsExtractOptionsCountryCw  CrawlRunParamsExtractOptionsCountry = "CW"
	CrawlRunParamsExtractOptionsCountryCx  CrawlRunParamsExtractOptionsCountry = "CX"
	CrawlRunParamsExtractOptionsCountryCy  CrawlRunParamsExtractOptionsCountry = "CY"
	CrawlRunParamsExtractOptionsCountryCz  CrawlRunParamsExtractOptionsCountry = "CZ"
	CrawlRunParamsExtractOptionsCountryDe  CrawlRunParamsExtractOptionsCountry = "DE"
	CrawlRunParamsExtractOptionsCountryDj  CrawlRunParamsExtractOptionsCountry = "DJ"
	CrawlRunParamsExtractOptionsCountryDk  CrawlRunParamsExtractOptionsCountry = "DK"
	CrawlRunParamsExtractOptionsCountryDm  CrawlRunParamsExtractOptionsCountry = "DM"
	CrawlRunParamsExtractOptionsCountryDo  CrawlRunParamsExtractOptionsCountry = "DO"
	CrawlRunParamsExtractOptionsCountryDz  CrawlRunParamsExtractOptionsCountry = "DZ"
	CrawlRunParamsExtractOptionsCountryEc  CrawlRunParamsExtractOptionsCountry = "EC"
	CrawlRunParamsExtractOptionsCountryEe  CrawlRunParamsExtractOptionsCountry = "EE"
	CrawlRunParamsExtractOptionsCountryEg  CrawlRunParamsExtractOptionsCountry = "EG"
	CrawlRunParamsExtractOptionsCountryEh  CrawlRunParamsExtractOptionsCountry = "EH"
	CrawlRunParamsExtractOptionsCountryEr  CrawlRunParamsExtractOptionsCountry = "ER"
	CrawlRunParamsExtractOptionsCountryEs  CrawlRunParamsExtractOptionsCountry = "ES"
	CrawlRunParamsExtractOptionsCountryEt  CrawlRunParamsExtractOptionsCountry = "ET"
	CrawlRunParamsExtractOptionsCountryFi  CrawlRunParamsExtractOptionsCountry = "FI"
	CrawlRunParamsExtractOptionsCountryFj  CrawlRunParamsExtractOptionsCountry = "FJ"
	CrawlRunParamsExtractOptionsCountryFk  CrawlRunParamsExtractOptionsCountry = "FK"
	CrawlRunParamsExtractOptionsCountryFm  CrawlRunParamsExtractOptionsCountry = "FM"
	CrawlRunParamsExtractOptionsCountryFo  CrawlRunParamsExtractOptionsCountry = "FO"
	CrawlRunParamsExtractOptionsCountryFr  CrawlRunParamsExtractOptionsCountry = "FR"
	CrawlRunParamsExtractOptionsCountryGa  CrawlRunParamsExtractOptionsCountry = "GA"
	CrawlRunParamsExtractOptionsCountryGB  CrawlRunParamsExtractOptionsCountry = "GB"
	CrawlRunParamsExtractOptionsCountryGd  CrawlRunParamsExtractOptionsCountry = "GD"
	CrawlRunParamsExtractOptionsCountryGe  CrawlRunParamsExtractOptionsCountry = "GE"
	CrawlRunParamsExtractOptionsCountryGf  CrawlRunParamsExtractOptionsCountry = "GF"
	CrawlRunParamsExtractOptionsCountryGg  CrawlRunParamsExtractOptionsCountry = "GG"
	CrawlRunParamsExtractOptionsCountryGh  CrawlRunParamsExtractOptionsCountry = "GH"
	CrawlRunParamsExtractOptionsCountryGi  CrawlRunParamsExtractOptionsCountry = "GI"
	CrawlRunParamsExtractOptionsCountryGl  CrawlRunParamsExtractOptionsCountry = "GL"
	CrawlRunParamsExtractOptionsCountryGm  CrawlRunParamsExtractOptionsCountry = "GM"
	CrawlRunParamsExtractOptionsCountryGn  CrawlRunParamsExtractOptionsCountry = "GN"
	CrawlRunParamsExtractOptionsCountryGp  CrawlRunParamsExtractOptionsCountry = "GP"
	CrawlRunParamsExtractOptionsCountryGq  CrawlRunParamsExtractOptionsCountry = "GQ"
	CrawlRunParamsExtractOptionsCountryGr  CrawlRunParamsExtractOptionsCountry = "GR"
	CrawlRunParamsExtractOptionsCountryGs  CrawlRunParamsExtractOptionsCountry = "GS"
	CrawlRunParamsExtractOptionsCountryGt  CrawlRunParamsExtractOptionsCountry = "GT"
	CrawlRunParamsExtractOptionsCountryGu  CrawlRunParamsExtractOptionsCountry = "GU"
	CrawlRunParamsExtractOptionsCountryGw  CrawlRunParamsExtractOptionsCountry = "GW"
	CrawlRunParamsExtractOptionsCountryGy  CrawlRunParamsExtractOptionsCountry = "GY"
	CrawlRunParamsExtractOptionsCountryHk  CrawlRunParamsExtractOptionsCountry = "HK"
	CrawlRunParamsExtractOptionsCountryHm  CrawlRunParamsExtractOptionsCountry = "HM"
	CrawlRunParamsExtractOptionsCountryHn  CrawlRunParamsExtractOptionsCountry = "HN"
	CrawlRunParamsExtractOptionsCountryHr  CrawlRunParamsExtractOptionsCountry = "HR"
	CrawlRunParamsExtractOptionsCountryHt  CrawlRunParamsExtractOptionsCountry = "HT"
	CrawlRunParamsExtractOptionsCountryHu  CrawlRunParamsExtractOptionsCountry = "HU"
	CrawlRunParamsExtractOptionsCountryID  CrawlRunParamsExtractOptionsCountry = "ID"
	CrawlRunParamsExtractOptionsCountryIe  CrawlRunParamsExtractOptionsCountry = "IE"
	CrawlRunParamsExtractOptionsCountryIl  CrawlRunParamsExtractOptionsCountry = "IL"
	CrawlRunParamsExtractOptionsCountryIm  CrawlRunParamsExtractOptionsCountry = "IM"
	CrawlRunParamsExtractOptionsCountryIn  CrawlRunParamsExtractOptionsCountry = "IN"
	CrawlRunParamsExtractOptionsCountryIo  CrawlRunParamsExtractOptionsCountry = "IO"
	CrawlRunParamsExtractOptionsCountryIq  CrawlRunParamsExtractOptionsCountry = "IQ"
	CrawlRunParamsExtractOptionsCountryIr  CrawlRunParamsExtractOptionsCountry = "IR"
	CrawlRunParamsExtractOptionsCountryIs  CrawlRunParamsExtractOptionsCountry = "IS"
	CrawlRunParamsExtractOptionsCountryIt  CrawlRunParamsExtractOptionsCountry = "IT"
	CrawlRunParamsExtractOptionsCountryJe  CrawlRunParamsExtractOptionsCountry = "JE"
	CrawlRunParamsExtractOptionsCountryJm  CrawlRunParamsExtractOptionsCountry = "JM"
	CrawlRunParamsExtractOptionsCountryJo  CrawlRunParamsExtractOptionsCountry = "JO"
	CrawlRunParamsExtractOptionsCountryJp  CrawlRunParamsExtractOptionsCountry = "JP"
	CrawlRunParamsExtractOptionsCountryKe  CrawlRunParamsExtractOptionsCountry = "KE"
	CrawlRunParamsExtractOptionsCountryKg  CrawlRunParamsExtractOptionsCountry = "KG"
	CrawlRunParamsExtractOptionsCountryKh  CrawlRunParamsExtractOptionsCountry = "KH"
	CrawlRunParamsExtractOptionsCountryKi  CrawlRunParamsExtractOptionsCountry = "KI"
	CrawlRunParamsExtractOptionsCountryKm  CrawlRunParamsExtractOptionsCountry = "KM"
	CrawlRunParamsExtractOptionsCountryKn  CrawlRunParamsExtractOptionsCountry = "KN"
	CrawlRunParamsExtractOptionsCountryKp  CrawlRunParamsExtractOptionsCountry = "KP"
	CrawlRunParamsExtractOptionsCountryKr  CrawlRunParamsExtractOptionsCountry = "KR"
	CrawlRunParamsExtractOptionsCountryKw  CrawlRunParamsExtractOptionsCountry = "KW"
	CrawlRunParamsExtractOptionsCountryKy  CrawlRunParamsExtractOptionsCountry = "KY"
	CrawlRunParamsExtractOptionsCountryKz  CrawlRunParamsExtractOptionsCountry = "KZ"
	CrawlRunParamsExtractOptionsCountryLa  CrawlRunParamsExtractOptionsCountry = "LA"
	CrawlRunParamsExtractOptionsCountryLb  CrawlRunParamsExtractOptionsCountry = "LB"
	CrawlRunParamsExtractOptionsCountryLc  CrawlRunParamsExtractOptionsCountry = "LC"
	CrawlRunParamsExtractOptionsCountryLi  CrawlRunParamsExtractOptionsCountry = "LI"
	CrawlRunParamsExtractOptionsCountryLk  CrawlRunParamsExtractOptionsCountry = "LK"
	CrawlRunParamsExtractOptionsCountryLr  CrawlRunParamsExtractOptionsCountry = "LR"
	CrawlRunParamsExtractOptionsCountryLs  CrawlRunParamsExtractOptionsCountry = "LS"
	CrawlRunParamsExtractOptionsCountryLt  CrawlRunParamsExtractOptionsCountry = "LT"
	CrawlRunParamsExtractOptionsCountryLu  CrawlRunParamsExtractOptionsCountry = "LU"
	CrawlRunParamsExtractOptionsCountryLv  CrawlRunParamsExtractOptionsCountry = "LV"
	CrawlRunParamsExtractOptionsCountryLy  CrawlRunParamsExtractOptionsCountry = "LY"
	CrawlRunParamsExtractOptionsCountryMa  CrawlRunParamsExtractOptionsCountry = "MA"
	CrawlRunParamsExtractOptionsCountryMc  CrawlRunParamsExtractOptionsCountry = "MC"
	CrawlRunParamsExtractOptionsCountryMd  CrawlRunParamsExtractOptionsCountry = "MD"
	CrawlRunParamsExtractOptionsCountryMe  CrawlRunParamsExtractOptionsCountry = "ME"
	CrawlRunParamsExtractOptionsCountryMf  CrawlRunParamsExtractOptionsCountry = "MF"
	CrawlRunParamsExtractOptionsCountryMg  CrawlRunParamsExtractOptionsCountry = "MG"
	CrawlRunParamsExtractOptionsCountryMh  CrawlRunParamsExtractOptionsCountry = "MH"
	CrawlRunParamsExtractOptionsCountryMk  CrawlRunParamsExtractOptionsCountry = "MK"
	CrawlRunParamsExtractOptionsCountryMl  CrawlRunParamsExtractOptionsCountry = "ML"
	CrawlRunParamsExtractOptionsCountryMm  CrawlRunParamsExtractOptionsCountry = "MM"
	CrawlRunParamsExtractOptionsCountryMn  CrawlRunParamsExtractOptionsCountry = "MN"
	CrawlRunParamsExtractOptionsCountryMo  CrawlRunParamsExtractOptionsCountry = "MO"
	CrawlRunParamsExtractOptionsCountryMp  CrawlRunParamsExtractOptionsCountry = "MP"
	CrawlRunParamsExtractOptionsCountryMq  CrawlRunParamsExtractOptionsCountry = "MQ"
	CrawlRunParamsExtractOptionsCountryMr  CrawlRunParamsExtractOptionsCountry = "MR"
	CrawlRunParamsExtractOptionsCountryMs  CrawlRunParamsExtractOptionsCountry = "MS"
	CrawlRunParamsExtractOptionsCountryMt  CrawlRunParamsExtractOptionsCountry = "MT"
	CrawlRunParamsExtractOptionsCountryMu  CrawlRunParamsExtractOptionsCountry = "MU"
	CrawlRunParamsExtractOptionsCountryMv  CrawlRunParamsExtractOptionsCountry = "MV"
	CrawlRunParamsExtractOptionsCountryMw  CrawlRunParamsExtractOptionsCountry = "MW"
	CrawlRunParamsExtractOptionsCountryMx  CrawlRunParamsExtractOptionsCountry = "MX"
	CrawlRunParamsExtractOptionsCountryMy  CrawlRunParamsExtractOptionsCountry = "MY"
	CrawlRunParamsExtractOptionsCountryMz  CrawlRunParamsExtractOptionsCountry = "MZ"
	CrawlRunParamsExtractOptionsCountryNa  CrawlRunParamsExtractOptionsCountry = "NA"
	CrawlRunParamsExtractOptionsCountryNc  CrawlRunParamsExtractOptionsCountry = "NC"
	CrawlRunParamsExtractOptionsCountryNe  CrawlRunParamsExtractOptionsCountry = "NE"
	CrawlRunParamsExtractOptionsCountryNf  CrawlRunParamsExtractOptionsCountry = "NF"
	CrawlRunParamsExtractOptionsCountryNg  CrawlRunParamsExtractOptionsCountry = "NG"
	CrawlRunParamsExtractOptionsCountryNi  CrawlRunParamsExtractOptionsCountry = "NI"
	CrawlRunParamsExtractOptionsCountryNl  CrawlRunParamsExtractOptionsCountry = "NL"
	CrawlRunParamsExtractOptionsCountryNo  CrawlRunParamsExtractOptionsCountry = "NO"
	CrawlRunParamsExtractOptionsCountryNp  CrawlRunParamsExtractOptionsCountry = "NP"
	CrawlRunParamsExtractOptionsCountryNr  CrawlRunParamsExtractOptionsCountry = "NR"
	CrawlRunParamsExtractOptionsCountryNu  CrawlRunParamsExtractOptionsCountry = "NU"
	CrawlRunParamsExtractOptionsCountryNz  CrawlRunParamsExtractOptionsCountry = "NZ"
	CrawlRunParamsExtractOptionsCountryOm  CrawlRunParamsExtractOptionsCountry = "OM"
	CrawlRunParamsExtractOptionsCountryPa  CrawlRunParamsExtractOptionsCountry = "PA"
	CrawlRunParamsExtractOptionsCountryPe  CrawlRunParamsExtractOptionsCountry = "PE"
	CrawlRunParamsExtractOptionsCountryPf  CrawlRunParamsExtractOptionsCountry = "PF"
	CrawlRunParamsExtractOptionsCountryPg  CrawlRunParamsExtractOptionsCountry = "PG"
	CrawlRunParamsExtractOptionsCountryPh  CrawlRunParamsExtractOptionsCountry = "PH"
	CrawlRunParamsExtractOptionsCountryPk  CrawlRunParamsExtractOptionsCountry = "PK"
	CrawlRunParamsExtractOptionsCountryPl  CrawlRunParamsExtractOptionsCountry = "PL"
	CrawlRunParamsExtractOptionsCountryPm  CrawlRunParamsExtractOptionsCountry = "PM"
	CrawlRunParamsExtractOptionsCountryPn  CrawlRunParamsExtractOptionsCountry = "PN"
	CrawlRunParamsExtractOptionsCountryPr  CrawlRunParamsExtractOptionsCountry = "PR"
	CrawlRunParamsExtractOptionsCountryPs  CrawlRunParamsExtractOptionsCountry = "PS"
	CrawlRunParamsExtractOptionsCountryPt  CrawlRunParamsExtractOptionsCountry = "PT"
	CrawlRunParamsExtractOptionsCountryPw  CrawlRunParamsExtractOptionsCountry = "PW"
	CrawlRunParamsExtractOptionsCountryPy  CrawlRunParamsExtractOptionsCountry = "PY"
	CrawlRunParamsExtractOptionsCountryQa  CrawlRunParamsExtractOptionsCountry = "QA"
	CrawlRunParamsExtractOptionsCountryRe  CrawlRunParamsExtractOptionsCountry = "RE"
	CrawlRunParamsExtractOptionsCountryRo  CrawlRunParamsExtractOptionsCountry = "RO"
	CrawlRunParamsExtractOptionsCountryRs  CrawlRunParamsExtractOptionsCountry = "RS"
	CrawlRunParamsExtractOptionsCountryRu  CrawlRunParamsExtractOptionsCountry = "RU"
	CrawlRunParamsExtractOptionsCountryRw  CrawlRunParamsExtractOptionsCountry = "RW"
	CrawlRunParamsExtractOptionsCountrySa  CrawlRunParamsExtractOptionsCountry = "SA"
	CrawlRunParamsExtractOptionsCountrySb  CrawlRunParamsExtractOptionsCountry = "SB"
	CrawlRunParamsExtractOptionsCountrySc  CrawlRunParamsExtractOptionsCountry = "SC"
	CrawlRunParamsExtractOptionsCountrySd  CrawlRunParamsExtractOptionsCountry = "SD"
	CrawlRunParamsExtractOptionsCountrySe  CrawlRunParamsExtractOptionsCountry = "SE"
	CrawlRunParamsExtractOptionsCountrySg  CrawlRunParamsExtractOptionsCountry = "SG"
	CrawlRunParamsExtractOptionsCountrySh  CrawlRunParamsExtractOptionsCountry = "SH"
	CrawlRunParamsExtractOptionsCountrySi  CrawlRunParamsExtractOptionsCountry = "SI"
	CrawlRunParamsExtractOptionsCountrySj  CrawlRunParamsExtractOptionsCountry = "SJ"
	CrawlRunParamsExtractOptionsCountrySk  CrawlRunParamsExtractOptionsCountry = "SK"
	CrawlRunParamsExtractOptionsCountrySl  CrawlRunParamsExtractOptionsCountry = "SL"
	CrawlRunParamsExtractOptionsCountrySm  CrawlRunParamsExtractOptionsCountry = "SM"
	CrawlRunParamsExtractOptionsCountrySn  CrawlRunParamsExtractOptionsCountry = "SN"
	CrawlRunParamsExtractOptionsCountrySo  CrawlRunParamsExtractOptionsCountry = "SO"
	CrawlRunParamsExtractOptionsCountrySr  CrawlRunParamsExtractOptionsCountry = "SR"
	CrawlRunParamsExtractOptionsCountrySS  CrawlRunParamsExtractOptionsCountry = "SS"
	CrawlRunParamsExtractOptionsCountrySt  CrawlRunParamsExtractOptionsCountry = "ST"
	CrawlRunParamsExtractOptionsCountrySv  CrawlRunParamsExtractOptionsCountry = "SV"
	CrawlRunParamsExtractOptionsCountrySx  CrawlRunParamsExtractOptionsCountry = "SX"
	CrawlRunParamsExtractOptionsCountrySy  CrawlRunParamsExtractOptionsCountry = "SY"
	CrawlRunParamsExtractOptionsCountrySz  CrawlRunParamsExtractOptionsCountry = "SZ"
	CrawlRunParamsExtractOptionsCountryTc  CrawlRunParamsExtractOptionsCountry = "TC"
	CrawlRunParamsExtractOptionsCountryTd  CrawlRunParamsExtractOptionsCountry = "TD"
	CrawlRunParamsExtractOptionsCountryTf  CrawlRunParamsExtractOptionsCountry = "TF"
	CrawlRunParamsExtractOptionsCountryTg  CrawlRunParamsExtractOptionsCountry = "TG"
	CrawlRunParamsExtractOptionsCountryTh  CrawlRunParamsExtractOptionsCountry = "TH"
	CrawlRunParamsExtractOptionsCountryTj  CrawlRunParamsExtractOptionsCountry = "TJ"
	CrawlRunParamsExtractOptionsCountryTk  CrawlRunParamsExtractOptionsCountry = "TK"
	CrawlRunParamsExtractOptionsCountryTl  CrawlRunParamsExtractOptionsCountry = "TL"
	CrawlRunParamsExtractOptionsCountryTm  CrawlRunParamsExtractOptionsCountry = "TM"
	CrawlRunParamsExtractOptionsCountryTn  CrawlRunParamsExtractOptionsCountry = "TN"
	CrawlRunParamsExtractOptionsCountryTo  CrawlRunParamsExtractOptionsCountry = "TO"
	CrawlRunParamsExtractOptionsCountryTr  CrawlRunParamsExtractOptionsCountry = "TR"
	CrawlRunParamsExtractOptionsCountryTt  CrawlRunParamsExtractOptionsCountry = "TT"
	CrawlRunParamsExtractOptionsCountryTv  CrawlRunParamsExtractOptionsCountry = "TV"
	CrawlRunParamsExtractOptionsCountryTw  CrawlRunParamsExtractOptionsCountry = "TW"
	CrawlRunParamsExtractOptionsCountryTz  CrawlRunParamsExtractOptionsCountry = "TZ"
	CrawlRunParamsExtractOptionsCountryUa  CrawlRunParamsExtractOptionsCountry = "UA"
	CrawlRunParamsExtractOptionsCountryUg  CrawlRunParamsExtractOptionsCountry = "UG"
	CrawlRunParamsExtractOptionsCountryUm  CrawlRunParamsExtractOptionsCountry = "UM"
	CrawlRunParamsExtractOptionsCountryUs  CrawlRunParamsExtractOptionsCountry = "US"
	CrawlRunParamsExtractOptionsCountryUy  CrawlRunParamsExtractOptionsCountry = "UY"
	CrawlRunParamsExtractOptionsCountryUz  CrawlRunParamsExtractOptionsCountry = "UZ"
	CrawlRunParamsExtractOptionsCountryVa  CrawlRunParamsExtractOptionsCountry = "VA"
	CrawlRunParamsExtractOptionsCountryVc  CrawlRunParamsExtractOptionsCountry = "VC"
	CrawlRunParamsExtractOptionsCountryVe  CrawlRunParamsExtractOptionsCountry = "VE"
	CrawlRunParamsExtractOptionsCountryVg  CrawlRunParamsExtractOptionsCountry = "VG"
	CrawlRunParamsExtractOptionsCountryVi  CrawlRunParamsExtractOptionsCountry = "VI"
	CrawlRunParamsExtractOptionsCountryVn  CrawlRunParamsExtractOptionsCountry = "VN"
	CrawlRunParamsExtractOptionsCountryVu  CrawlRunParamsExtractOptionsCountry = "VU"
	CrawlRunParamsExtractOptionsCountryWf  CrawlRunParamsExtractOptionsCountry = "WF"
	CrawlRunParamsExtractOptionsCountryWs  CrawlRunParamsExtractOptionsCountry = "WS"
	CrawlRunParamsExtractOptionsCountryXk  CrawlRunParamsExtractOptionsCountry = "XK"
	CrawlRunParamsExtractOptionsCountryYe  CrawlRunParamsExtractOptionsCountry = "YE"
	CrawlRunParamsExtractOptionsCountryYt  CrawlRunParamsExtractOptionsCountry = "YT"
	CrawlRunParamsExtractOptionsCountryZa  CrawlRunParamsExtractOptionsCountry = "ZA"
	CrawlRunParamsExtractOptionsCountryZm  CrawlRunParamsExtractOptionsCountry = "ZM"
	CrawlRunParamsExtractOptionsCountryZw  CrawlRunParamsExtractOptionsCountry = "ZW"
	CrawlRunParamsExtractOptionsCountryAll CrawlRunParamsExtractOptionsCountry = "ALL"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRunParamsExtractOptionsHeaderUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRunParamsExtractOptionsHeaderUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlRunParamsExtractOptionsHeaderUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRunParamsExtractOptionsHeaderUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Locale for browser language and region settings
type CrawlRunParamsExtractOptionsLocale string

const (
	CrawlRunParamsExtractOptionsLocaleAaDj      CrawlRunParamsExtractOptionsLocale = "aa-DJ"
	CrawlRunParamsExtractOptionsLocaleAaEr      CrawlRunParamsExtractOptionsLocale = "aa-ER"
	CrawlRunParamsExtractOptionsLocaleAaEt      CrawlRunParamsExtractOptionsLocale = "aa-ET"
	CrawlRunParamsExtractOptionsLocaleAf        CrawlRunParamsExtractOptionsLocale = "af"
	CrawlRunParamsExtractOptionsLocaleAfNa      CrawlRunParamsExtractOptionsLocale = "af-NA"
	CrawlRunParamsExtractOptionsLocaleAfZa      CrawlRunParamsExtractOptionsLocale = "af-ZA"
	CrawlRunParamsExtractOptionsLocaleAk        CrawlRunParamsExtractOptionsLocale = "ak"
	CrawlRunParamsExtractOptionsLocaleAkGh      CrawlRunParamsExtractOptionsLocale = "ak-GH"
	CrawlRunParamsExtractOptionsLocaleAm        CrawlRunParamsExtractOptionsLocale = "am"
	CrawlRunParamsExtractOptionsLocaleAmEt      CrawlRunParamsExtractOptionsLocale = "am-ET"
	CrawlRunParamsExtractOptionsLocaleAnEs      CrawlRunParamsExtractOptionsLocale = "an-ES"
	CrawlRunParamsExtractOptionsLocaleAr        CrawlRunParamsExtractOptionsLocale = "ar"
	CrawlRunParamsExtractOptionsLocaleArAe      CrawlRunParamsExtractOptionsLocale = "ar-AE"
	CrawlRunParamsExtractOptionsLocaleArBh      CrawlRunParamsExtractOptionsLocale = "ar-BH"
	CrawlRunParamsExtractOptionsLocaleArDz      CrawlRunParamsExtractOptionsLocale = "ar-DZ"
	CrawlRunParamsExtractOptionsLocaleArEg      CrawlRunParamsExtractOptionsLocale = "ar-EG"
	CrawlRunParamsExtractOptionsLocaleArIn      CrawlRunParamsExtractOptionsLocale = "ar-IN"
	CrawlRunParamsExtractOptionsLocaleArIq      CrawlRunParamsExtractOptionsLocale = "ar-IQ"
	CrawlRunParamsExtractOptionsLocaleArJo      CrawlRunParamsExtractOptionsLocale = "ar-JO"
	CrawlRunParamsExtractOptionsLocaleArKw      CrawlRunParamsExtractOptionsLocale = "ar-KW"
	CrawlRunParamsExtractOptionsLocaleArLb      CrawlRunParamsExtractOptionsLocale = "ar-LB"
	CrawlRunParamsExtractOptionsLocaleArLy      CrawlRunParamsExtractOptionsLocale = "ar-LY"
	CrawlRunParamsExtractOptionsLocaleArMa      CrawlRunParamsExtractOptionsLocale = "ar-MA"
	CrawlRunParamsExtractOptionsLocaleArOm      CrawlRunParamsExtractOptionsLocale = "ar-OM"
	CrawlRunParamsExtractOptionsLocaleArQa      CrawlRunParamsExtractOptionsLocale = "ar-QA"
	CrawlRunParamsExtractOptionsLocaleArSa      CrawlRunParamsExtractOptionsLocale = "ar-SA"
	CrawlRunParamsExtractOptionsLocaleArSd      CrawlRunParamsExtractOptionsLocale = "ar-SD"
	CrawlRunParamsExtractOptionsLocaleArSy      CrawlRunParamsExtractOptionsLocale = "ar-SY"
	CrawlRunParamsExtractOptionsLocaleArTn      CrawlRunParamsExtractOptionsLocale = "ar-TN"
	CrawlRunParamsExtractOptionsLocaleArYe      CrawlRunParamsExtractOptionsLocale = "ar-YE"
	CrawlRunParamsExtractOptionsLocaleAs        CrawlRunParamsExtractOptionsLocale = "as"
	CrawlRunParamsExtractOptionsLocaleAsIn      CrawlRunParamsExtractOptionsLocale = "as-IN"
	CrawlRunParamsExtractOptionsLocaleAsa       CrawlRunParamsExtractOptionsLocale = "asa"
	CrawlRunParamsExtractOptionsLocaleAsaTz     CrawlRunParamsExtractOptionsLocale = "asa-TZ"
	CrawlRunParamsExtractOptionsLocaleAstEs     CrawlRunParamsExtractOptionsLocale = "ast-ES"
	CrawlRunParamsExtractOptionsLocaleAz        CrawlRunParamsExtractOptionsLocale = "az"
	CrawlRunParamsExtractOptionsLocaleAzAz      CrawlRunParamsExtractOptionsLocale = "az-AZ"
	CrawlRunParamsExtractOptionsLocaleAzCyrl    CrawlRunParamsExtractOptionsLocale = "az-Cyrl"
	CrawlRunParamsExtractOptionsLocaleAzCyrlAz  CrawlRunParamsExtractOptionsLocale = "az-Cyrl-AZ"
	CrawlRunParamsExtractOptionsLocaleAzLatn    CrawlRunParamsExtractOptionsLocale = "az-Latn"
	CrawlRunParamsExtractOptionsLocaleAzLatnAz  CrawlRunParamsExtractOptionsLocale = "az-Latn-AZ"
	CrawlRunParamsExtractOptionsLocaleBe        CrawlRunParamsExtractOptionsLocale = "be"
	CrawlRunParamsExtractOptionsLocaleBeBy      CrawlRunParamsExtractOptionsLocale = "be-BY"
	CrawlRunParamsExtractOptionsLocaleBem       CrawlRunParamsExtractOptionsLocale = "bem"
	CrawlRunParamsExtractOptionsLocaleBemZm     CrawlRunParamsExtractOptionsLocale = "bem-ZM"
	CrawlRunParamsExtractOptionsLocaleBerDz     CrawlRunParamsExtractOptionsLocale = "ber-DZ"
	CrawlRunParamsExtractOptionsLocaleBerMa     CrawlRunParamsExtractOptionsLocale = "ber-MA"
	CrawlRunParamsExtractOptionsLocaleBez       CrawlRunParamsExtractOptionsLocale = "bez"
	CrawlRunParamsExtractOptionsLocaleBezTz     CrawlRunParamsExtractOptionsLocale = "bez-TZ"
	CrawlRunParamsExtractOptionsLocaleBg        CrawlRunParamsExtractOptionsLocale = "bg"
	CrawlRunParamsExtractOptionsLocaleBgBg      CrawlRunParamsExtractOptionsLocale = "bg-BG"
	CrawlRunParamsExtractOptionsLocaleBhoIn     CrawlRunParamsExtractOptionsLocale = "bho-IN"
	CrawlRunParamsExtractOptionsLocaleBm        CrawlRunParamsExtractOptionsLocale = "bm"
	CrawlRunParamsExtractOptionsLocaleBmMl      CrawlRunParamsExtractOptionsLocale = "bm-ML"
	CrawlRunParamsExtractOptionsLocaleBn        CrawlRunParamsExtractOptionsLocale = "bn"
	CrawlRunParamsExtractOptionsLocaleBnBd      CrawlRunParamsExtractOptionsLocale = "bn-BD"
	CrawlRunParamsExtractOptionsLocaleBnIn      CrawlRunParamsExtractOptionsLocale = "bn-IN"
	CrawlRunParamsExtractOptionsLocaleBo        CrawlRunParamsExtractOptionsLocale = "bo"
	CrawlRunParamsExtractOptionsLocaleBoCn      CrawlRunParamsExtractOptionsLocale = "bo-CN"
	CrawlRunParamsExtractOptionsLocaleBoIn      CrawlRunParamsExtractOptionsLocale = "bo-IN"
	CrawlRunParamsExtractOptionsLocaleBrFr      CrawlRunParamsExtractOptionsLocale = "br-FR"
	CrawlRunParamsExtractOptionsLocaleBrxIn     CrawlRunParamsExtractOptionsLocale = "brx-IN"
	CrawlRunParamsExtractOptionsLocaleBs        CrawlRunParamsExtractOptionsLocale = "bs"
	CrawlRunParamsExtractOptionsLocaleBsBa      CrawlRunParamsExtractOptionsLocale = "bs-BA"
	CrawlRunParamsExtractOptionsLocaleBynEr     CrawlRunParamsExtractOptionsLocale = "byn-ER"
	CrawlRunParamsExtractOptionsLocaleCa        CrawlRunParamsExtractOptionsLocale = "ca"
	CrawlRunParamsExtractOptionsLocaleCaAd      CrawlRunParamsExtractOptionsLocale = "ca-AD"
	CrawlRunParamsExtractOptionsLocaleCaEs      CrawlRunParamsExtractOptionsLocale = "ca-ES"
	CrawlRunParamsExtractOptionsLocaleCaFr      CrawlRunParamsExtractOptionsLocale = "ca-FR"
	CrawlRunParamsExtractOptionsLocaleCaIt      CrawlRunParamsExtractOptionsLocale = "ca-IT"
	CrawlRunParamsExtractOptionsLocaleCgg       CrawlRunParamsExtractOptionsLocale = "cgg"
	CrawlRunParamsExtractOptionsLocaleCggUg     CrawlRunParamsExtractOptionsLocale = "cgg-UG"
	CrawlRunParamsExtractOptionsLocaleChr       CrawlRunParamsExtractOptionsLocale = "chr"
	CrawlRunParamsExtractOptionsLocaleChrUs     CrawlRunParamsExtractOptionsLocale = "chr-US"
	CrawlRunParamsExtractOptionsLocaleCrhUa     CrawlRunParamsExtractOptionsLocale = "crh-UA"
	CrawlRunParamsExtractOptionsLocaleCs        CrawlRunParamsExtractOptionsLocale = "cs"
	CrawlRunParamsExtractOptionsLocaleCsCz      CrawlRunParamsExtractOptionsLocale = "cs-CZ"
	CrawlRunParamsExtractOptionsLocaleCsbPl     CrawlRunParamsExtractOptionsLocale = "csb-PL"
	CrawlRunParamsExtractOptionsLocaleCvRu      CrawlRunParamsExtractOptionsLocale = "cv-RU"
	CrawlRunParamsExtractOptionsLocaleCy        CrawlRunParamsExtractOptionsLocale = "cy"
	CrawlRunParamsExtractOptionsLocaleCyGB      CrawlRunParamsExtractOptionsLocale = "cy-GB"
	CrawlRunParamsExtractOptionsLocaleDa        CrawlRunParamsExtractOptionsLocale = "da"
	CrawlRunParamsExtractOptionsLocaleDaDk      CrawlRunParamsExtractOptionsLocale = "da-DK"
	CrawlRunParamsExtractOptionsLocaleDav       CrawlRunParamsExtractOptionsLocale = "dav"
	CrawlRunParamsExtractOptionsLocaleDavKe     CrawlRunParamsExtractOptionsLocale = "dav-KE"
	CrawlRunParamsExtractOptionsLocaleDe        CrawlRunParamsExtractOptionsLocale = "de"
	CrawlRunParamsExtractOptionsLocaleDeAt      CrawlRunParamsExtractOptionsLocale = "de-AT"
	CrawlRunParamsExtractOptionsLocaleDeBe      CrawlRunParamsExtractOptionsLocale = "de-BE"
	CrawlRunParamsExtractOptionsLocaleDeCh      CrawlRunParamsExtractOptionsLocale = "de-CH"
	CrawlRunParamsExtractOptionsLocaleDeDe      CrawlRunParamsExtractOptionsLocale = "de-DE"
	CrawlRunParamsExtractOptionsLocaleDeLi      CrawlRunParamsExtractOptionsLocale = "de-LI"
	CrawlRunParamsExtractOptionsLocaleDeLu      CrawlRunParamsExtractOptionsLocale = "de-LU"
	CrawlRunParamsExtractOptionsLocaleDvMv      CrawlRunParamsExtractOptionsLocale = "dv-MV"
	CrawlRunParamsExtractOptionsLocaleDzBt      CrawlRunParamsExtractOptionsLocale = "dz-BT"
	CrawlRunParamsExtractOptionsLocaleEbu       CrawlRunParamsExtractOptionsLocale = "ebu"
	CrawlRunParamsExtractOptionsLocaleEbuKe     CrawlRunParamsExtractOptionsLocale = "ebu-KE"
	CrawlRunParamsExtractOptionsLocaleEe        CrawlRunParamsExtractOptionsLocale = "ee"
	CrawlRunParamsExtractOptionsLocaleEeGh      CrawlRunParamsExtractOptionsLocale = "ee-GH"
	CrawlRunParamsExtractOptionsLocaleEeTg      CrawlRunParamsExtractOptionsLocale = "ee-TG"
	CrawlRunParamsExtractOptionsLocaleEl        CrawlRunParamsExtractOptionsLocale = "el"
	CrawlRunParamsExtractOptionsLocaleElCy      CrawlRunParamsExtractOptionsLocale = "el-CY"
	CrawlRunParamsExtractOptionsLocaleElGr      CrawlRunParamsExtractOptionsLocale = "el-GR"
	CrawlRunParamsExtractOptionsLocaleEn        CrawlRunParamsExtractOptionsLocale = "en"
	CrawlRunParamsExtractOptionsLocaleEnAg      CrawlRunParamsExtractOptionsLocale = "en-AG"
	CrawlRunParamsExtractOptionsLocaleEnAs      CrawlRunParamsExtractOptionsLocale = "en-AS"
	CrawlRunParamsExtractOptionsLocaleEnAu      CrawlRunParamsExtractOptionsLocale = "en-AU"
	CrawlRunParamsExtractOptionsLocaleEnBe      CrawlRunParamsExtractOptionsLocale = "en-BE"
	CrawlRunParamsExtractOptionsLocaleEnBw      CrawlRunParamsExtractOptionsLocale = "en-BW"
	CrawlRunParamsExtractOptionsLocaleEnBz      CrawlRunParamsExtractOptionsLocale = "en-BZ"
	CrawlRunParamsExtractOptionsLocaleEnCa      CrawlRunParamsExtractOptionsLocale = "en-CA"
	CrawlRunParamsExtractOptionsLocaleEnDk      CrawlRunParamsExtractOptionsLocale = "en-DK"
	CrawlRunParamsExtractOptionsLocaleEnGB      CrawlRunParamsExtractOptionsLocale = "en-GB"
	CrawlRunParamsExtractOptionsLocaleEnGu      CrawlRunParamsExtractOptionsLocale = "en-GU"
	CrawlRunParamsExtractOptionsLocaleEnHk      CrawlRunParamsExtractOptionsLocale = "en-HK"
	CrawlRunParamsExtractOptionsLocaleEnIe      CrawlRunParamsExtractOptionsLocale = "en-IE"
	CrawlRunParamsExtractOptionsLocaleEnIn      CrawlRunParamsExtractOptionsLocale = "en-IN"
	CrawlRunParamsExtractOptionsLocaleEnJm      CrawlRunParamsExtractOptionsLocale = "en-JM"
	CrawlRunParamsExtractOptionsLocaleEnMh      CrawlRunParamsExtractOptionsLocale = "en-MH"
	CrawlRunParamsExtractOptionsLocaleEnMp      CrawlRunParamsExtractOptionsLocale = "en-MP"
	CrawlRunParamsExtractOptionsLocaleEnMt      CrawlRunParamsExtractOptionsLocale = "en-MT"
	CrawlRunParamsExtractOptionsLocaleEnMu      CrawlRunParamsExtractOptionsLocale = "en-MU"
	CrawlRunParamsExtractOptionsLocaleEnNa      CrawlRunParamsExtractOptionsLocale = "en-NA"
	CrawlRunParamsExtractOptionsLocaleEnNg      CrawlRunParamsExtractOptionsLocale = "en-NG"
	CrawlRunParamsExtractOptionsLocaleEnNz      CrawlRunParamsExtractOptionsLocale = "en-NZ"
	CrawlRunParamsExtractOptionsLocaleEnPh      CrawlRunParamsExtractOptionsLocale = "en-PH"
	CrawlRunParamsExtractOptionsLocaleEnPk      CrawlRunParamsExtractOptionsLocale = "en-PK"
	CrawlRunParamsExtractOptionsLocaleEnSg      CrawlRunParamsExtractOptionsLocale = "en-SG"
	CrawlRunParamsExtractOptionsLocaleEnTt      CrawlRunParamsExtractOptionsLocale = "en-TT"
	CrawlRunParamsExtractOptionsLocaleEnUm      CrawlRunParamsExtractOptionsLocale = "en-UM"
	CrawlRunParamsExtractOptionsLocaleEnUs      CrawlRunParamsExtractOptionsLocale = "en-US"
	CrawlRunParamsExtractOptionsLocaleEnVi      CrawlRunParamsExtractOptionsLocale = "en-VI"
	CrawlRunParamsExtractOptionsLocaleEnZa      CrawlRunParamsExtractOptionsLocale = "en-ZA"
	CrawlRunParamsExtractOptionsLocaleEnZm      CrawlRunParamsExtractOptionsLocale = "en-ZM"
	CrawlRunParamsExtractOptionsLocaleEnZw      CrawlRunParamsExtractOptionsLocale = "en-ZW"
	CrawlRunParamsExtractOptionsLocaleEo        CrawlRunParamsExtractOptionsLocale = "eo"
	CrawlRunParamsExtractOptionsLocaleEs        CrawlRunParamsExtractOptionsLocale = "es"
	CrawlRunParamsExtractOptionsLocaleEs419     CrawlRunParamsExtractOptionsLocale = "es-419"
	CrawlRunParamsExtractOptionsLocaleEsAr      CrawlRunParamsExtractOptionsLocale = "es-AR"
	CrawlRunParamsExtractOptionsLocaleEsBo      CrawlRunParamsExtractOptionsLocale = "es-BO"
	CrawlRunParamsExtractOptionsLocaleEsCl      CrawlRunParamsExtractOptionsLocale = "es-CL"
	CrawlRunParamsExtractOptionsLocaleEsCo      CrawlRunParamsExtractOptionsLocale = "es-CO"
	CrawlRunParamsExtractOptionsLocaleEsCr      CrawlRunParamsExtractOptionsLocale = "es-CR"
	CrawlRunParamsExtractOptionsLocaleEsCu      CrawlRunParamsExtractOptionsLocale = "es-CU"
	CrawlRunParamsExtractOptionsLocaleEsDo      CrawlRunParamsExtractOptionsLocale = "es-DO"
	CrawlRunParamsExtractOptionsLocaleEsEc      CrawlRunParamsExtractOptionsLocale = "es-EC"
	CrawlRunParamsExtractOptionsLocaleEsEs      CrawlRunParamsExtractOptionsLocale = "es-ES"
	CrawlRunParamsExtractOptionsLocaleEsGq      CrawlRunParamsExtractOptionsLocale = "es-GQ"
	CrawlRunParamsExtractOptionsLocaleEsGt      CrawlRunParamsExtractOptionsLocale = "es-GT"
	CrawlRunParamsExtractOptionsLocaleEsHn      CrawlRunParamsExtractOptionsLocale = "es-HN"
	CrawlRunParamsExtractOptionsLocaleEsMx      CrawlRunParamsExtractOptionsLocale = "es-MX"
	CrawlRunParamsExtractOptionsLocaleEsNi      CrawlRunParamsExtractOptionsLocale = "es-NI"
	CrawlRunParamsExtractOptionsLocaleEsPa      CrawlRunParamsExtractOptionsLocale = "es-PA"
	CrawlRunParamsExtractOptionsLocaleEsPe      CrawlRunParamsExtractOptionsLocale = "es-PE"
	CrawlRunParamsExtractOptionsLocaleEsPr      CrawlRunParamsExtractOptionsLocale = "es-PR"
	CrawlRunParamsExtractOptionsLocaleEsPy      CrawlRunParamsExtractOptionsLocale = "es-PY"
	CrawlRunParamsExtractOptionsLocaleEsSv      CrawlRunParamsExtractOptionsLocale = "es-SV"
	CrawlRunParamsExtractOptionsLocaleEsUs      CrawlRunParamsExtractOptionsLocale = "es-US"
	CrawlRunParamsExtractOptionsLocaleEsUy      CrawlRunParamsExtractOptionsLocale = "es-UY"
	CrawlRunParamsExtractOptionsLocaleEsVe      CrawlRunParamsExtractOptionsLocale = "es-VE"
	CrawlRunParamsExtractOptionsLocaleEt        CrawlRunParamsExtractOptionsLocale = "et"
	CrawlRunParamsExtractOptionsLocaleEtEe      CrawlRunParamsExtractOptionsLocale = "et-EE"
	CrawlRunParamsExtractOptionsLocaleEu        CrawlRunParamsExtractOptionsLocale = "eu"
	CrawlRunParamsExtractOptionsLocaleEuEs      CrawlRunParamsExtractOptionsLocale = "eu-ES"
	CrawlRunParamsExtractOptionsLocaleFa        CrawlRunParamsExtractOptionsLocale = "fa"
	CrawlRunParamsExtractOptionsLocaleFaAf      CrawlRunParamsExtractOptionsLocale = "fa-AF"
	CrawlRunParamsExtractOptionsLocaleFaIr      CrawlRunParamsExtractOptionsLocale = "fa-IR"
	CrawlRunParamsExtractOptionsLocaleFf        CrawlRunParamsExtractOptionsLocale = "ff"
	CrawlRunParamsExtractOptionsLocaleFfSn      CrawlRunParamsExtractOptionsLocale = "ff-SN"
	CrawlRunParamsExtractOptionsLocaleFi        CrawlRunParamsExtractOptionsLocale = "fi"
	CrawlRunParamsExtractOptionsLocaleFiFi      CrawlRunParamsExtractOptionsLocale = "fi-FI"
	CrawlRunParamsExtractOptionsLocaleFil       CrawlRunParamsExtractOptionsLocale = "fil"
	CrawlRunParamsExtractOptionsLocaleFilPh     CrawlRunParamsExtractOptionsLocale = "fil-PH"
	CrawlRunParamsExtractOptionsLocaleFo        CrawlRunParamsExtractOptionsLocale = "fo"
	CrawlRunParamsExtractOptionsLocaleFoFo      CrawlRunParamsExtractOptionsLocale = "fo-FO"
	CrawlRunParamsExtractOptionsLocaleFr        CrawlRunParamsExtractOptionsLocale = "fr"
	CrawlRunParamsExtractOptionsLocaleFrBe      CrawlRunParamsExtractOptionsLocale = "fr-BE"
	CrawlRunParamsExtractOptionsLocaleFrBf      CrawlRunParamsExtractOptionsLocale = "fr-BF"
	CrawlRunParamsExtractOptionsLocaleFrBi      CrawlRunParamsExtractOptionsLocale = "fr-BI"
	CrawlRunParamsExtractOptionsLocaleFrBj      CrawlRunParamsExtractOptionsLocale = "fr-BJ"
	CrawlRunParamsExtractOptionsLocaleFrBl      CrawlRunParamsExtractOptionsLocale = "fr-BL"
	CrawlRunParamsExtractOptionsLocaleFrCa      CrawlRunParamsExtractOptionsLocale = "fr-CA"
	CrawlRunParamsExtractOptionsLocaleFrCd      CrawlRunParamsExtractOptionsLocale = "fr-CD"
	CrawlRunParamsExtractOptionsLocaleFrCf      CrawlRunParamsExtractOptionsLocale = "fr-CF"
	CrawlRunParamsExtractOptionsLocaleFrCg      CrawlRunParamsExtractOptionsLocale = "fr-CG"
	CrawlRunParamsExtractOptionsLocaleFrCh      CrawlRunParamsExtractOptionsLocale = "fr-CH"
	CrawlRunParamsExtractOptionsLocaleFrCi      CrawlRunParamsExtractOptionsLocale = "fr-CI"
	CrawlRunParamsExtractOptionsLocaleFrCm      CrawlRunParamsExtractOptionsLocale = "fr-CM"
	CrawlRunParamsExtractOptionsLocaleFrDj      CrawlRunParamsExtractOptionsLocale = "fr-DJ"
	CrawlRunParamsExtractOptionsLocaleFrFr      CrawlRunParamsExtractOptionsLocale = "fr-FR"
	CrawlRunParamsExtractOptionsLocaleFrGa      CrawlRunParamsExtractOptionsLocale = "fr-GA"
	CrawlRunParamsExtractOptionsLocaleFrGn      CrawlRunParamsExtractOptionsLocale = "fr-GN"
	CrawlRunParamsExtractOptionsLocaleFrGp      CrawlRunParamsExtractOptionsLocale = "fr-GP"
	CrawlRunParamsExtractOptionsLocaleFrGq      CrawlRunParamsExtractOptionsLocale = "fr-GQ"
	CrawlRunParamsExtractOptionsLocaleFrKm      CrawlRunParamsExtractOptionsLocale = "fr-KM"
	CrawlRunParamsExtractOptionsLocaleFrLu      CrawlRunParamsExtractOptionsLocale = "fr-LU"
	CrawlRunParamsExtractOptionsLocaleFrMc      CrawlRunParamsExtractOptionsLocale = "fr-MC"
	CrawlRunParamsExtractOptionsLocaleFrMf      CrawlRunParamsExtractOptionsLocale = "fr-MF"
	CrawlRunParamsExtractOptionsLocaleFrMg      CrawlRunParamsExtractOptionsLocale = "fr-MG"
	CrawlRunParamsExtractOptionsLocaleFrMl      CrawlRunParamsExtractOptionsLocale = "fr-ML"
	CrawlRunParamsExtractOptionsLocaleFrMq      CrawlRunParamsExtractOptionsLocale = "fr-MQ"
	CrawlRunParamsExtractOptionsLocaleFrNe      CrawlRunParamsExtractOptionsLocale = "fr-NE"
	CrawlRunParamsExtractOptionsLocaleFrRe      CrawlRunParamsExtractOptionsLocale = "fr-RE"
	CrawlRunParamsExtractOptionsLocaleFrRw      CrawlRunParamsExtractOptionsLocale = "fr-RW"
	CrawlRunParamsExtractOptionsLocaleFrSn      CrawlRunParamsExtractOptionsLocale = "fr-SN"
	CrawlRunParamsExtractOptionsLocaleFrTd      CrawlRunParamsExtractOptionsLocale = "fr-TD"
	CrawlRunParamsExtractOptionsLocaleFrTg      CrawlRunParamsExtractOptionsLocale = "fr-TG"
	CrawlRunParamsExtractOptionsLocaleFurIt     CrawlRunParamsExtractOptionsLocale = "fur-IT"
	CrawlRunParamsExtractOptionsLocaleFyDe      CrawlRunParamsExtractOptionsLocale = "fy-DE"
	CrawlRunParamsExtractOptionsLocaleFyNl      CrawlRunParamsExtractOptionsLocale = "fy-NL"
	CrawlRunParamsExtractOptionsLocaleGa        CrawlRunParamsExtractOptionsLocale = "ga"
	CrawlRunParamsExtractOptionsLocaleGaIe      CrawlRunParamsExtractOptionsLocale = "ga-IE"
	CrawlRunParamsExtractOptionsLocaleGdGB      CrawlRunParamsExtractOptionsLocale = "gd-GB"
	CrawlRunParamsExtractOptionsLocaleGezEr     CrawlRunParamsExtractOptionsLocale = "gez-ER"
	CrawlRunParamsExtractOptionsLocaleGezEt     CrawlRunParamsExtractOptionsLocale = "gez-ET"
	CrawlRunParamsExtractOptionsLocaleGl        CrawlRunParamsExtractOptionsLocale = "gl"
	CrawlRunParamsExtractOptionsLocaleGlEs      CrawlRunParamsExtractOptionsLocale = "gl-ES"
	CrawlRunParamsExtractOptionsLocaleGsw       CrawlRunParamsExtractOptionsLocale = "gsw"
	CrawlRunParamsExtractOptionsLocaleGswCh     CrawlRunParamsExtractOptionsLocale = "gsw-CH"
	CrawlRunParamsExtractOptionsLocaleGu        CrawlRunParamsExtractOptionsLocale = "gu"
	CrawlRunParamsExtractOptionsLocaleGuIn      CrawlRunParamsExtractOptionsLocale = "gu-IN"
	CrawlRunParamsExtractOptionsLocaleGuz       CrawlRunParamsExtractOptionsLocale = "guz"
	CrawlRunParamsExtractOptionsLocaleGuzKe     CrawlRunParamsExtractOptionsLocale = "guz-KE"
	CrawlRunParamsExtractOptionsLocaleGv        CrawlRunParamsExtractOptionsLocale = "gv"
	CrawlRunParamsExtractOptionsLocaleGvGB      CrawlRunParamsExtractOptionsLocale = "gv-GB"
	CrawlRunParamsExtractOptionsLocaleHa        CrawlRunParamsExtractOptionsLocale = "ha"
	CrawlRunParamsExtractOptionsLocaleHaLatn    CrawlRunParamsExtractOptionsLocale = "ha-Latn"
	CrawlRunParamsExtractOptionsLocaleHaLatnGh  CrawlRunParamsExtractOptionsLocale = "ha-Latn-GH"
	CrawlRunParamsExtractOptionsLocaleHaLatnNe  CrawlRunParamsExtractOptionsLocale = "ha-Latn-NE"
	CrawlRunParamsExtractOptionsLocaleHaLatnNg  CrawlRunParamsExtractOptionsLocale = "ha-Latn-NG"
	CrawlRunParamsExtractOptionsLocaleHaNg      CrawlRunParamsExtractOptionsLocale = "ha-NG"
	CrawlRunParamsExtractOptionsLocaleHaw       CrawlRunParamsExtractOptionsLocale = "haw"
	CrawlRunParamsExtractOptionsLocaleHawUs     CrawlRunParamsExtractOptionsLocale = "haw-US"
	CrawlRunParamsExtractOptionsLocaleHe        CrawlRunParamsExtractOptionsLocale = "he"
	CrawlRunParamsExtractOptionsLocaleHeIl      CrawlRunParamsExtractOptionsLocale = "he-IL"
	CrawlRunParamsExtractOptionsLocaleHi        CrawlRunParamsExtractOptionsLocale = "hi"
	CrawlRunParamsExtractOptionsLocaleHiIn      CrawlRunParamsExtractOptionsLocale = "hi-IN"
	CrawlRunParamsExtractOptionsLocaleHneIn     CrawlRunParamsExtractOptionsLocale = "hne-IN"
	CrawlRunParamsExtractOptionsLocaleHr        CrawlRunParamsExtractOptionsLocale = "hr"
	CrawlRunParamsExtractOptionsLocaleHrHr      CrawlRunParamsExtractOptionsLocale = "hr-HR"
	CrawlRunParamsExtractOptionsLocaleHsbDe     CrawlRunParamsExtractOptionsLocale = "hsb-DE"
	CrawlRunParamsExtractOptionsLocaleHtHt      CrawlRunParamsExtractOptionsLocale = "ht-HT"
	CrawlRunParamsExtractOptionsLocaleHu        CrawlRunParamsExtractOptionsLocale = "hu"
	CrawlRunParamsExtractOptionsLocaleHuHu      CrawlRunParamsExtractOptionsLocale = "hu-HU"
	CrawlRunParamsExtractOptionsLocaleHy        CrawlRunParamsExtractOptionsLocale = "hy"
	CrawlRunParamsExtractOptionsLocaleHyAm      CrawlRunParamsExtractOptionsLocale = "hy-AM"
	CrawlRunParamsExtractOptionsLocaleID        CrawlRunParamsExtractOptionsLocale = "id"
	CrawlRunParamsExtractOptionsLocaleIDID      CrawlRunParamsExtractOptionsLocale = "id-ID"
	CrawlRunParamsExtractOptionsLocaleIg        CrawlRunParamsExtractOptionsLocale = "ig"
	CrawlRunParamsExtractOptionsLocaleIgNg      CrawlRunParamsExtractOptionsLocale = "ig-NG"
	CrawlRunParamsExtractOptionsLocaleIi        CrawlRunParamsExtractOptionsLocale = "ii"
	CrawlRunParamsExtractOptionsLocaleIiCn      CrawlRunParamsExtractOptionsLocale = "ii-CN"
	CrawlRunParamsExtractOptionsLocaleIkCa      CrawlRunParamsExtractOptionsLocale = "ik-CA"
	CrawlRunParamsExtractOptionsLocaleIs        CrawlRunParamsExtractOptionsLocale = "is"
	CrawlRunParamsExtractOptionsLocaleIsIs      CrawlRunParamsExtractOptionsLocale = "is-IS"
	CrawlRunParamsExtractOptionsLocaleIt        CrawlRunParamsExtractOptionsLocale = "it"
	CrawlRunParamsExtractOptionsLocaleItCh      CrawlRunParamsExtractOptionsLocale = "it-CH"
	CrawlRunParamsExtractOptionsLocaleItIt      CrawlRunParamsExtractOptionsLocale = "it-IT"
	CrawlRunParamsExtractOptionsLocaleIuCa      CrawlRunParamsExtractOptionsLocale = "iu-CA"
	CrawlRunParamsExtractOptionsLocaleIwIl      CrawlRunParamsExtractOptionsLocale = "iw-IL"
	CrawlRunParamsExtractOptionsLocaleJa        CrawlRunParamsExtractOptionsLocale = "ja"
	CrawlRunParamsExtractOptionsLocaleJaJp      CrawlRunParamsExtractOptionsLocale = "ja-JP"
	CrawlRunParamsExtractOptionsLocaleJmc       CrawlRunParamsExtractOptionsLocale = "jmc"
	CrawlRunParamsExtractOptionsLocaleJmcTz     CrawlRunParamsExtractOptionsLocale = "jmc-TZ"
	CrawlRunParamsExtractOptionsLocaleKa        CrawlRunParamsExtractOptionsLocale = "ka"
	CrawlRunParamsExtractOptionsLocaleKaGe      CrawlRunParamsExtractOptionsLocale = "ka-GE"
	CrawlRunParamsExtractOptionsLocaleKab       CrawlRunParamsExtractOptionsLocale = "kab"
	CrawlRunParamsExtractOptionsLocaleKabDz     CrawlRunParamsExtractOptionsLocale = "kab-DZ"
	CrawlRunParamsExtractOptionsLocaleKam       CrawlRunParamsExtractOptionsLocale = "kam"
	CrawlRunParamsExtractOptionsLocaleKamKe     CrawlRunParamsExtractOptionsLocale = "kam-KE"
	CrawlRunParamsExtractOptionsLocaleKde       CrawlRunParamsExtractOptionsLocale = "kde"
	CrawlRunParamsExtractOptionsLocaleKdeTz     CrawlRunParamsExtractOptionsLocale = "kde-TZ"
	CrawlRunParamsExtractOptionsLocaleKea       CrawlRunParamsExtractOptionsLocale = "kea"
	CrawlRunParamsExtractOptionsLocaleKeaCv     CrawlRunParamsExtractOptionsLocale = "kea-CV"
	CrawlRunParamsExtractOptionsLocaleKhq       CrawlRunParamsExtractOptionsLocale = "khq"
	CrawlRunParamsExtractOptionsLocaleKhqMl     CrawlRunParamsExtractOptionsLocale = "khq-ML"
	CrawlRunParamsExtractOptionsLocaleKi        CrawlRunParamsExtractOptionsLocale = "ki"
	CrawlRunParamsExtractOptionsLocaleKiKe      CrawlRunParamsExtractOptionsLocale = "ki-KE"
	CrawlRunParamsExtractOptionsLocaleKk        CrawlRunParamsExtractOptionsLocale = "kk"
	CrawlRunParamsExtractOptionsLocaleKkCyrl    CrawlRunParamsExtractOptionsLocale = "kk-Cyrl"
	CrawlRunParamsExtractOptionsLocaleKkCyrlKz  CrawlRunParamsExtractOptionsLocale = "kk-Cyrl-KZ"
	CrawlRunParamsExtractOptionsLocaleKkKz      CrawlRunParamsExtractOptionsLocale = "kk-KZ"
	CrawlRunParamsExtractOptionsLocaleKl        CrawlRunParamsExtractOptionsLocale = "kl"
	CrawlRunParamsExtractOptionsLocaleKlGl      CrawlRunParamsExtractOptionsLocale = "kl-GL"
	CrawlRunParamsExtractOptionsLocaleKln       CrawlRunParamsExtractOptionsLocale = "kln"
	CrawlRunParamsExtractOptionsLocaleKlnKe     CrawlRunParamsExtractOptionsLocale = "kln-KE"
	CrawlRunParamsExtractOptionsLocaleKm        CrawlRunParamsExtractOptionsLocale = "km"
	CrawlRunParamsExtractOptionsLocaleKmKh      CrawlRunParamsExtractOptionsLocale = "km-KH"
	CrawlRunParamsExtractOptionsLocaleKn        CrawlRunParamsExtractOptionsLocale = "kn"
	CrawlRunParamsExtractOptionsLocaleKnIn      CrawlRunParamsExtractOptionsLocale = "kn-IN"
	CrawlRunParamsExtractOptionsLocaleKo        CrawlRunParamsExtractOptionsLocale = "ko"
	CrawlRunParamsExtractOptionsLocaleKoKr      CrawlRunParamsExtractOptionsLocale = "ko-KR"
	CrawlRunParamsExtractOptionsLocaleKok       CrawlRunParamsExtractOptionsLocale = "kok"
	CrawlRunParamsExtractOptionsLocaleKokIn     CrawlRunParamsExtractOptionsLocale = "kok-IN"
	CrawlRunParamsExtractOptionsLocaleKsIn      CrawlRunParamsExtractOptionsLocale = "ks-IN"
	CrawlRunParamsExtractOptionsLocaleKuTr      CrawlRunParamsExtractOptionsLocale = "ku-TR"
	CrawlRunParamsExtractOptionsLocaleKw        CrawlRunParamsExtractOptionsLocale = "kw"
	CrawlRunParamsExtractOptionsLocaleKwGB      CrawlRunParamsExtractOptionsLocale = "kw-GB"
	CrawlRunParamsExtractOptionsLocaleKyKg      CrawlRunParamsExtractOptionsLocale = "ky-KG"
	CrawlRunParamsExtractOptionsLocaleLag       CrawlRunParamsExtractOptionsLocale = "lag"
	CrawlRunParamsExtractOptionsLocaleLagTz     CrawlRunParamsExtractOptionsLocale = "lag-TZ"
	CrawlRunParamsExtractOptionsLocaleLbLu      CrawlRunParamsExtractOptionsLocale = "lb-LU"
	CrawlRunParamsExtractOptionsLocaleLg        CrawlRunParamsExtractOptionsLocale = "lg"
	CrawlRunParamsExtractOptionsLocaleLgUg      CrawlRunParamsExtractOptionsLocale = "lg-UG"
	CrawlRunParamsExtractOptionsLocaleLiBe      CrawlRunParamsExtractOptionsLocale = "li-BE"
	CrawlRunParamsExtractOptionsLocaleLiNl      CrawlRunParamsExtractOptionsLocale = "li-NL"
	CrawlRunParamsExtractOptionsLocaleLijIt     CrawlRunParamsExtractOptionsLocale = "lij-IT"
	CrawlRunParamsExtractOptionsLocaleLoLa      CrawlRunParamsExtractOptionsLocale = "lo-LA"
	CrawlRunParamsExtractOptionsLocaleLt        CrawlRunParamsExtractOptionsLocale = "lt"
	CrawlRunParamsExtractOptionsLocaleLtLt      CrawlRunParamsExtractOptionsLocale = "lt-LT"
	CrawlRunParamsExtractOptionsLocaleLuo       CrawlRunParamsExtractOptionsLocale = "luo"
	CrawlRunParamsExtractOptionsLocaleLuoKe     CrawlRunParamsExtractOptionsLocale = "luo-KE"
	CrawlRunParamsExtractOptionsLocaleLuy       CrawlRunParamsExtractOptionsLocale = "luy"
	CrawlRunParamsExtractOptionsLocaleLuyKe     CrawlRunParamsExtractOptionsLocale = "luy-KE"
	CrawlRunParamsExtractOptionsLocaleLv        CrawlRunParamsExtractOptionsLocale = "lv"
	CrawlRunParamsExtractOptionsLocaleLvLv      CrawlRunParamsExtractOptionsLocale = "lv-LV"
	CrawlRunParamsExtractOptionsLocaleMagIn     CrawlRunParamsExtractOptionsLocale = "mag-IN"
	CrawlRunParamsExtractOptionsLocaleMaiIn     CrawlRunParamsExtractOptionsLocale = "mai-IN"
	CrawlRunParamsExtractOptionsLocaleMas       CrawlRunParamsExtractOptionsLocale = "mas"
	CrawlRunParamsExtractOptionsLocaleMasKe     CrawlRunParamsExtractOptionsLocale = "mas-KE"
	CrawlRunParamsExtractOptionsLocaleMasTz     CrawlRunParamsExtractOptionsLocale = "mas-TZ"
	CrawlRunParamsExtractOptionsLocaleMer       CrawlRunParamsExtractOptionsLocale = "mer"
	CrawlRunParamsExtractOptionsLocaleMerKe     CrawlRunParamsExtractOptionsLocale = "mer-KE"
	CrawlRunParamsExtractOptionsLocaleMfe       CrawlRunParamsExtractOptionsLocale = "mfe"
	CrawlRunParamsExtractOptionsLocaleMfeMu     CrawlRunParamsExtractOptionsLocale = "mfe-MU"
	CrawlRunParamsExtractOptionsLocaleMg        CrawlRunParamsExtractOptionsLocale = "mg"
	CrawlRunParamsExtractOptionsLocaleMgMg      CrawlRunParamsExtractOptionsLocale = "mg-MG"
	CrawlRunParamsExtractOptionsLocaleMhrRu     CrawlRunParamsExtractOptionsLocale = "mhr-RU"
	CrawlRunParamsExtractOptionsLocaleMiNz      CrawlRunParamsExtractOptionsLocale = "mi-NZ"
	CrawlRunParamsExtractOptionsLocaleMk        CrawlRunParamsExtractOptionsLocale = "mk"
	CrawlRunParamsExtractOptionsLocaleMkMk      CrawlRunParamsExtractOptionsLocale = "mk-MK"
	CrawlRunParamsExtractOptionsLocaleMl        CrawlRunParamsExtractOptionsLocale = "ml"
	CrawlRunParamsExtractOptionsLocaleMlIn      CrawlRunParamsExtractOptionsLocale = "ml-IN"
	CrawlRunParamsExtractOptionsLocaleMnMn      CrawlRunParamsExtractOptionsLocale = "mn-MN"
	CrawlRunParamsExtractOptionsLocaleMr        CrawlRunParamsExtractOptionsLocale = "mr"
	CrawlRunParamsExtractOptionsLocaleMrIn      CrawlRunParamsExtractOptionsLocale = "mr-IN"
	CrawlRunParamsExtractOptionsLocaleMs        CrawlRunParamsExtractOptionsLocale = "ms"
	CrawlRunParamsExtractOptionsLocaleMsBn      CrawlRunParamsExtractOptionsLocale = "ms-BN"
	CrawlRunParamsExtractOptionsLocaleMsMy      CrawlRunParamsExtractOptionsLocale = "ms-MY"
	CrawlRunParamsExtractOptionsLocaleMt        CrawlRunParamsExtractOptionsLocale = "mt"
	CrawlRunParamsExtractOptionsLocaleMtMt      CrawlRunParamsExtractOptionsLocale = "mt-MT"
	CrawlRunParamsExtractOptionsLocaleMy        CrawlRunParamsExtractOptionsLocale = "my"
	CrawlRunParamsExtractOptionsLocaleMyMm      CrawlRunParamsExtractOptionsLocale = "my-MM"
	CrawlRunParamsExtractOptionsLocaleNanTw     CrawlRunParamsExtractOptionsLocale = "nan-TW"
	CrawlRunParamsExtractOptionsLocaleNaq       CrawlRunParamsExtractOptionsLocale = "naq"
	CrawlRunParamsExtractOptionsLocaleNaqNa     CrawlRunParamsExtractOptionsLocale = "naq-NA"
	CrawlRunParamsExtractOptionsLocaleNb        CrawlRunParamsExtractOptionsLocale = "nb"
	CrawlRunParamsExtractOptionsLocaleNbNo      CrawlRunParamsExtractOptionsLocale = "nb-NO"
	CrawlRunParamsExtractOptionsLocaleNd        CrawlRunParamsExtractOptionsLocale = "nd"
	CrawlRunParamsExtractOptionsLocaleNdZw      CrawlRunParamsExtractOptionsLocale = "nd-ZW"
	CrawlRunParamsExtractOptionsLocaleNdsDe     CrawlRunParamsExtractOptionsLocale = "nds-DE"
	CrawlRunParamsExtractOptionsLocaleNdsNl     CrawlRunParamsExtractOptionsLocale = "nds-NL"
	CrawlRunParamsExtractOptionsLocaleNe        CrawlRunParamsExtractOptionsLocale = "ne"
	CrawlRunParamsExtractOptionsLocaleNeIn      CrawlRunParamsExtractOptionsLocale = "ne-IN"
	CrawlRunParamsExtractOptionsLocaleNeNp      CrawlRunParamsExtractOptionsLocale = "ne-NP"
	CrawlRunParamsExtractOptionsLocaleNl        CrawlRunParamsExtractOptionsLocale = "nl"
	CrawlRunParamsExtractOptionsLocaleNlAw      CrawlRunParamsExtractOptionsLocale = "nl-AW"
	CrawlRunParamsExtractOptionsLocaleNlBe      CrawlRunParamsExtractOptionsLocale = "nl-BE"
	CrawlRunParamsExtractOptionsLocaleNlNl      CrawlRunParamsExtractOptionsLocale = "nl-NL"
	CrawlRunParamsExtractOptionsLocaleNn        CrawlRunParamsExtractOptionsLocale = "nn"
	CrawlRunParamsExtractOptionsLocaleNnNo      CrawlRunParamsExtractOptionsLocale = "nn-NO"
	CrawlRunParamsExtractOptionsLocaleNrZa      CrawlRunParamsExtractOptionsLocale = "nr-ZA"
	CrawlRunParamsExtractOptionsLocaleNsoZa     CrawlRunParamsExtractOptionsLocale = "nso-ZA"
	CrawlRunParamsExtractOptionsLocaleNyn       CrawlRunParamsExtractOptionsLocale = "nyn"
	CrawlRunParamsExtractOptionsLocaleNynUg     CrawlRunParamsExtractOptionsLocale = "nyn-UG"
	CrawlRunParamsExtractOptionsLocaleOcFr      CrawlRunParamsExtractOptionsLocale = "oc-FR"
	CrawlRunParamsExtractOptionsLocaleOm        CrawlRunParamsExtractOptionsLocale = "om"
	CrawlRunParamsExtractOptionsLocaleOmEt      CrawlRunParamsExtractOptionsLocale = "om-ET"
	CrawlRunParamsExtractOptionsLocaleOmKe      CrawlRunParamsExtractOptionsLocale = "om-KE"
	CrawlRunParamsExtractOptionsLocaleOr        CrawlRunParamsExtractOptionsLocale = "or"
	CrawlRunParamsExtractOptionsLocaleOrIn      CrawlRunParamsExtractOptionsLocale = "or-IN"
	CrawlRunParamsExtractOptionsLocaleOsRu      CrawlRunParamsExtractOptionsLocale = "os-RU"
	CrawlRunParamsExtractOptionsLocalePa        CrawlRunParamsExtractOptionsLocale = "pa"
	CrawlRunParamsExtractOptionsLocalePaArab    CrawlRunParamsExtractOptionsLocale = "pa-Arab"
	CrawlRunParamsExtractOptionsLocalePaArabPk  CrawlRunParamsExtractOptionsLocale = "pa-Arab-PK"
	CrawlRunParamsExtractOptionsLocalePaGuru    CrawlRunParamsExtractOptionsLocale = "pa-Guru"
	CrawlRunParamsExtractOptionsLocalePaGuruIn  CrawlRunParamsExtractOptionsLocale = "pa-Guru-IN"
	CrawlRunParamsExtractOptionsLocalePaIn      CrawlRunParamsExtractOptionsLocale = "pa-IN"
	CrawlRunParamsExtractOptionsLocalePaPk      CrawlRunParamsExtractOptionsLocale = "pa-PK"
	CrawlRunParamsExtractOptionsLocalePapAn     CrawlRunParamsExtractOptionsLocale = "pap-AN"
	CrawlRunParamsExtractOptionsLocalePl        CrawlRunParamsExtractOptionsLocale = "pl"
	CrawlRunParamsExtractOptionsLocalePlPl      CrawlRunParamsExtractOptionsLocale = "pl-PL"
	CrawlRunParamsExtractOptionsLocalePs        CrawlRunParamsExtractOptionsLocale = "ps"
	CrawlRunParamsExtractOptionsLocalePsAf      CrawlRunParamsExtractOptionsLocale = "ps-AF"
	CrawlRunParamsExtractOptionsLocalePt        CrawlRunParamsExtractOptionsLocale = "pt"
	CrawlRunParamsExtractOptionsLocalePtBr      CrawlRunParamsExtractOptionsLocale = "pt-BR"
	CrawlRunParamsExtractOptionsLocalePtGw      CrawlRunParamsExtractOptionsLocale = "pt-GW"
	CrawlRunParamsExtractOptionsLocalePtMz      CrawlRunParamsExtractOptionsLocale = "pt-MZ"
	CrawlRunParamsExtractOptionsLocalePtPt      CrawlRunParamsExtractOptionsLocale = "pt-PT"
	CrawlRunParamsExtractOptionsLocaleRm        CrawlRunParamsExtractOptionsLocale = "rm"
	CrawlRunParamsExtractOptionsLocaleRmCh      CrawlRunParamsExtractOptionsLocale = "rm-CH"
	CrawlRunParamsExtractOptionsLocaleRo        CrawlRunParamsExtractOptionsLocale = "ro"
	CrawlRunParamsExtractOptionsLocaleRoMd      CrawlRunParamsExtractOptionsLocale = "ro-MD"
	CrawlRunParamsExtractOptionsLocaleRoRo      CrawlRunParamsExtractOptionsLocale = "ro-RO"
	CrawlRunParamsExtractOptionsLocaleRof       CrawlRunParamsExtractOptionsLocale = "rof"
	CrawlRunParamsExtractOptionsLocaleRofTz     CrawlRunParamsExtractOptionsLocale = "rof-TZ"
	CrawlRunParamsExtractOptionsLocaleRu        CrawlRunParamsExtractOptionsLocale = "ru"
	CrawlRunParamsExtractOptionsLocaleRuMd      CrawlRunParamsExtractOptionsLocale = "ru-MD"
	CrawlRunParamsExtractOptionsLocaleRuRu      CrawlRunParamsExtractOptionsLocale = "ru-RU"
	CrawlRunParamsExtractOptionsLocaleRuUa      CrawlRunParamsExtractOptionsLocale = "ru-UA"
	CrawlRunParamsExtractOptionsLocaleRw        CrawlRunParamsExtractOptionsLocale = "rw"
	CrawlRunParamsExtractOptionsLocaleRwRw      CrawlRunParamsExtractOptionsLocale = "rw-RW"
	CrawlRunParamsExtractOptionsLocaleRwk       CrawlRunParamsExtractOptionsLocale = "rwk"
	CrawlRunParamsExtractOptionsLocaleRwkTz     CrawlRunParamsExtractOptionsLocale = "rwk-TZ"
	CrawlRunParamsExtractOptionsLocaleSaIn      CrawlRunParamsExtractOptionsLocale = "sa-IN"
	CrawlRunParamsExtractOptionsLocaleSaq       CrawlRunParamsExtractOptionsLocale = "saq"
	CrawlRunParamsExtractOptionsLocaleSaqKe     CrawlRunParamsExtractOptionsLocale = "saq-KE"
	CrawlRunParamsExtractOptionsLocaleScIt      CrawlRunParamsExtractOptionsLocale = "sc-IT"
	CrawlRunParamsExtractOptionsLocaleSdIn      CrawlRunParamsExtractOptionsLocale = "sd-IN"
	CrawlRunParamsExtractOptionsLocaleSeNo      CrawlRunParamsExtractOptionsLocale = "se-NO"
	CrawlRunParamsExtractOptionsLocaleSeh       CrawlRunParamsExtractOptionsLocale = "seh"
	CrawlRunParamsExtractOptionsLocaleSehMz     CrawlRunParamsExtractOptionsLocale = "seh-MZ"
	CrawlRunParamsExtractOptionsLocaleSes       CrawlRunParamsExtractOptionsLocale = "ses"
	CrawlRunParamsExtractOptionsLocaleSesMl     CrawlRunParamsExtractOptionsLocale = "ses-ML"
	CrawlRunParamsExtractOptionsLocaleSg        CrawlRunParamsExtractOptionsLocale = "sg"
	CrawlRunParamsExtractOptionsLocaleSgCf      CrawlRunParamsExtractOptionsLocale = "sg-CF"
	CrawlRunParamsExtractOptionsLocaleShi       CrawlRunParamsExtractOptionsLocale = "shi"
	CrawlRunParamsExtractOptionsLocaleShiLatn   CrawlRunParamsExtractOptionsLocale = "shi-Latn"
	CrawlRunParamsExtractOptionsLocaleShiLatnMa CrawlRunParamsExtractOptionsLocale = "shi-Latn-MA"
	CrawlRunParamsExtractOptionsLocaleShiTfng   CrawlRunParamsExtractOptionsLocale = "shi-Tfng"
	CrawlRunParamsExtractOptionsLocaleShiTfngMa CrawlRunParamsExtractOptionsLocale = "shi-Tfng-MA"
	CrawlRunParamsExtractOptionsLocaleShsCa     CrawlRunParamsExtractOptionsLocale = "shs-CA"
	CrawlRunParamsExtractOptionsLocaleSi        CrawlRunParamsExtractOptionsLocale = "si"
	CrawlRunParamsExtractOptionsLocaleSiLk      CrawlRunParamsExtractOptionsLocale = "si-LK"
	CrawlRunParamsExtractOptionsLocaleSidEt     CrawlRunParamsExtractOptionsLocale = "sid-ET"
	CrawlRunParamsExtractOptionsLocaleSk        CrawlRunParamsExtractOptionsLocale = "sk"
	CrawlRunParamsExtractOptionsLocaleSkSk      CrawlRunParamsExtractOptionsLocale = "sk-SK"
	CrawlRunParamsExtractOptionsLocaleSl        CrawlRunParamsExtractOptionsLocale = "sl"
	CrawlRunParamsExtractOptionsLocaleSlSi      CrawlRunParamsExtractOptionsLocale = "sl-SI"
	CrawlRunParamsExtractOptionsLocaleSn        CrawlRunParamsExtractOptionsLocale = "sn"
	CrawlRunParamsExtractOptionsLocaleSnZw      CrawlRunParamsExtractOptionsLocale = "sn-ZW"
	CrawlRunParamsExtractOptionsLocaleSo        CrawlRunParamsExtractOptionsLocale = "so"
	CrawlRunParamsExtractOptionsLocaleSoDj      CrawlRunParamsExtractOptionsLocale = "so-DJ"
	CrawlRunParamsExtractOptionsLocaleSoEt      CrawlRunParamsExtractOptionsLocale = "so-ET"
	CrawlRunParamsExtractOptionsLocaleSoKe      CrawlRunParamsExtractOptionsLocale = "so-KE"
	CrawlRunParamsExtractOptionsLocaleSoSo      CrawlRunParamsExtractOptionsLocale = "so-SO"
	CrawlRunParamsExtractOptionsLocaleSq        CrawlRunParamsExtractOptionsLocale = "sq"
	CrawlRunParamsExtractOptionsLocaleSqAl      CrawlRunParamsExtractOptionsLocale = "sq-AL"
	CrawlRunParamsExtractOptionsLocaleSqMk      CrawlRunParamsExtractOptionsLocale = "sq-MK"
	CrawlRunParamsExtractOptionsLocaleSr        CrawlRunParamsExtractOptionsLocale = "sr"
	CrawlRunParamsExtractOptionsLocaleSrCyrl    CrawlRunParamsExtractOptionsLocale = "sr-Cyrl"
	CrawlRunParamsExtractOptionsLocaleSrCyrlBa  CrawlRunParamsExtractOptionsLocale = "sr-Cyrl-BA"
	CrawlRunParamsExtractOptionsLocaleSrCyrlMe  CrawlRunParamsExtractOptionsLocale = "sr-Cyrl-ME"
	CrawlRunParamsExtractOptionsLocaleSrCyrlRs  CrawlRunParamsExtractOptionsLocale = "sr-Cyrl-RS"
	CrawlRunParamsExtractOptionsLocaleSrLatn    CrawlRunParamsExtractOptionsLocale = "sr-Latn"
	CrawlRunParamsExtractOptionsLocaleSrLatnBa  CrawlRunParamsExtractOptionsLocale = "sr-Latn-BA"
	CrawlRunParamsExtractOptionsLocaleSrLatnMe  CrawlRunParamsExtractOptionsLocale = "sr-Latn-ME"
	CrawlRunParamsExtractOptionsLocaleSrLatnRs  CrawlRunParamsExtractOptionsLocale = "sr-Latn-RS"
	CrawlRunParamsExtractOptionsLocaleSrMe      CrawlRunParamsExtractOptionsLocale = "sr-ME"
	CrawlRunParamsExtractOptionsLocaleSrRs      CrawlRunParamsExtractOptionsLocale = "sr-RS"
	CrawlRunParamsExtractOptionsLocaleSSZa      CrawlRunParamsExtractOptionsLocale = "ss-ZA"
	CrawlRunParamsExtractOptionsLocaleStZa      CrawlRunParamsExtractOptionsLocale = "st-ZA"
	CrawlRunParamsExtractOptionsLocaleSv        CrawlRunParamsExtractOptionsLocale = "sv"
	CrawlRunParamsExtractOptionsLocaleSvFi      CrawlRunParamsExtractOptionsLocale = "sv-FI"
	CrawlRunParamsExtractOptionsLocaleSvSe      CrawlRunParamsExtractOptionsLocale = "sv-SE"
	CrawlRunParamsExtractOptionsLocaleSw        CrawlRunParamsExtractOptionsLocale = "sw"
	CrawlRunParamsExtractOptionsLocaleSwKe      CrawlRunParamsExtractOptionsLocale = "sw-KE"
	CrawlRunParamsExtractOptionsLocaleSwTz      CrawlRunParamsExtractOptionsLocale = "sw-TZ"
	CrawlRunParamsExtractOptionsLocaleTa        CrawlRunParamsExtractOptionsLocale = "ta"
	CrawlRunParamsExtractOptionsLocaleTaIn      CrawlRunParamsExtractOptionsLocale = "ta-IN"
	CrawlRunParamsExtractOptionsLocaleTaLk      CrawlRunParamsExtractOptionsLocale = "ta-LK"
	CrawlRunParamsExtractOptionsLocaleTe        CrawlRunParamsExtractOptionsLocale = "te"
	CrawlRunParamsExtractOptionsLocaleTeIn      CrawlRunParamsExtractOptionsLocale = "te-IN"
	CrawlRunParamsExtractOptionsLocaleTeo       CrawlRunParamsExtractOptionsLocale = "teo"
	CrawlRunParamsExtractOptionsLocaleTeoKe     CrawlRunParamsExtractOptionsLocale = "teo-KE"
	CrawlRunParamsExtractOptionsLocaleTeoUg     CrawlRunParamsExtractOptionsLocale = "teo-UG"
	CrawlRunParamsExtractOptionsLocaleTgTj      CrawlRunParamsExtractOptionsLocale = "tg-TJ"
	CrawlRunParamsExtractOptionsLocaleTh        CrawlRunParamsExtractOptionsLocale = "th"
	CrawlRunParamsExtractOptionsLocaleThTh      CrawlRunParamsExtractOptionsLocale = "th-TH"
	CrawlRunParamsExtractOptionsLocaleTi        CrawlRunParamsExtractOptionsLocale = "ti"
	CrawlRunParamsExtractOptionsLocaleTiEr      CrawlRunParamsExtractOptionsLocale = "ti-ER"
	CrawlRunParamsExtractOptionsLocaleTiEt      CrawlRunParamsExtractOptionsLocale = "ti-ET"
	CrawlRunParamsExtractOptionsLocaleTigEr     CrawlRunParamsExtractOptionsLocale = "tig-ER"
	CrawlRunParamsExtractOptionsLocaleTkTm      CrawlRunParamsExtractOptionsLocale = "tk-TM"
	CrawlRunParamsExtractOptionsLocaleTlPh      CrawlRunParamsExtractOptionsLocale = "tl-PH"
	CrawlRunParamsExtractOptionsLocaleTnZa      CrawlRunParamsExtractOptionsLocale = "tn-ZA"
	CrawlRunParamsExtractOptionsLocaleTo        CrawlRunParamsExtractOptionsLocale = "to"
	CrawlRunParamsExtractOptionsLocaleToTo      CrawlRunParamsExtractOptionsLocale = "to-TO"
	CrawlRunParamsExtractOptionsLocaleTr        CrawlRunParamsExtractOptionsLocale = "tr"
	CrawlRunParamsExtractOptionsLocaleTrCy      CrawlRunParamsExtractOptionsLocale = "tr-CY"
	CrawlRunParamsExtractOptionsLocaleTrTr      CrawlRunParamsExtractOptionsLocale = "tr-TR"
	CrawlRunParamsExtractOptionsLocaleTsZa      CrawlRunParamsExtractOptionsLocale = "ts-ZA"
	CrawlRunParamsExtractOptionsLocaleTtRu      CrawlRunParamsExtractOptionsLocale = "tt-RU"
	CrawlRunParamsExtractOptionsLocaleTzm       CrawlRunParamsExtractOptionsLocale = "tzm"
	CrawlRunParamsExtractOptionsLocaleTzmLatn   CrawlRunParamsExtractOptionsLocale = "tzm-Latn"
	CrawlRunParamsExtractOptionsLocaleTzmLatnMa CrawlRunParamsExtractOptionsLocale = "tzm-Latn-MA"
	CrawlRunParamsExtractOptionsLocaleUgCn      CrawlRunParamsExtractOptionsLocale = "ug-CN"
	CrawlRunParamsExtractOptionsLocaleUk        CrawlRunParamsExtractOptionsLocale = "uk"
	CrawlRunParamsExtractOptionsLocaleUkUa      CrawlRunParamsExtractOptionsLocale = "uk-UA"
	CrawlRunParamsExtractOptionsLocaleUnmUs     CrawlRunParamsExtractOptionsLocale = "unm-US"
	CrawlRunParamsExtractOptionsLocaleUr        CrawlRunParamsExtractOptionsLocale = "ur"
	CrawlRunParamsExtractOptionsLocaleUrIn      CrawlRunParamsExtractOptionsLocale = "ur-IN"
	CrawlRunParamsExtractOptionsLocaleUrPk      CrawlRunParamsExtractOptionsLocale = "ur-PK"
	CrawlRunParamsExtractOptionsLocaleUz        CrawlRunParamsExtractOptionsLocale = "uz"
	CrawlRunParamsExtractOptionsLocaleUzArab    CrawlRunParamsExtractOptionsLocale = "uz-Arab"
	CrawlRunParamsExtractOptionsLocaleUzArabAf  CrawlRunParamsExtractOptionsLocale = "uz-Arab-AF"
	CrawlRunParamsExtractOptionsLocaleUzCyrl    CrawlRunParamsExtractOptionsLocale = "uz-Cyrl"
	CrawlRunParamsExtractOptionsLocaleUzCyrlUz  CrawlRunParamsExtractOptionsLocale = "uz-Cyrl-UZ"
	CrawlRunParamsExtractOptionsLocaleUzLatn    CrawlRunParamsExtractOptionsLocale = "uz-Latn"
	CrawlRunParamsExtractOptionsLocaleUzLatnUz  CrawlRunParamsExtractOptionsLocale = "uz-Latn-UZ"
	CrawlRunParamsExtractOptionsLocaleUzUz      CrawlRunParamsExtractOptionsLocale = "uz-UZ"
	CrawlRunParamsExtractOptionsLocaleVeZa      CrawlRunParamsExtractOptionsLocale = "ve-ZA"
	CrawlRunParamsExtractOptionsLocaleVi        CrawlRunParamsExtractOptionsLocale = "vi"
	CrawlRunParamsExtractOptionsLocaleViVn      CrawlRunParamsExtractOptionsLocale = "vi-VN"
	CrawlRunParamsExtractOptionsLocaleVun       CrawlRunParamsExtractOptionsLocale = "vun"
	CrawlRunParamsExtractOptionsLocaleVunTz     CrawlRunParamsExtractOptionsLocale = "vun-TZ"
	CrawlRunParamsExtractOptionsLocaleWaBe      CrawlRunParamsExtractOptionsLocale = "wa-BE"
	CrawlRunParamsExtractOptionsLocaleWaeCh     CrawlRunParamsExtractOptionsLocale = "wae-CH"
	CrawlRunParamsExtractOptionsLocaleWalEt     CrawlRunParamsExtractOptionsLocale = "wal-ET"
	CrawlRunParamsExtractOptionsLocaleWoSn      CrawlRunParamsExtractOptionsLocale = "wo-SN"
	CrawlRunParamsExtractOptionsLocaleXhZa      CrawlRunParamsExtractOptionsLocale = "xh-ZA"
	CrawlRunParamsExtractOptionsLocaleXog       CrawlRunParamsExtractOptionsLocale = "xog"
	CrawlRunParamsExtractOptionsLocaleXogUg     CrawlRunParamsExtractOptionsLocale = "xog-UG"
	CrawlRunParamsExtractOptionsLocaleYiUs      CrawlRunParamsExtractOptionsLocale = "yi-US"
	CrawlRunParamsExtractOptionsLocaleYo        CrawlRunParamsExtractOptionsLocale = "yo"
	CrawlRunParamsExtractOptionsLocaleYoNg      CrawlRunParamsExtractOptionsLocale = "yo-NG"
	CrawlRunParamsExtractOptionsLocaleYueHk     CrawlRunParamsExtractOptionsLocale = "yue-HK"
	CrawlRunParamsExtractOptionsLocaleZh        CrawlRunParamsExtractOptionsLocale = "zh"
	CrawlRunParamsExtractOptionsLocaleZhCn      CrawlRunParamsExtractOptionsLocale = "zh-CN"
	CrawlRunParamsExtractOptionsLocaleZhHk      CrawlRunParamsExtractOptionsLocale = "zh-HK"
	CrawlRunParamsExtractOptionsLocaleZhHans    CrawlRunParamsExtractOptionsLocale = "zh-Hans"
	CrawlRunParamsExtractOptionsLocaleZhHansCn  CrawlRunParamsExtractOptionsLocale = "zh-Hans-CN"
	CrawlRunParamsExtractOptionsLocaleZhHansHk  CrawlRunParamsExtractOptionsLocale = "zh-Hans-HK"
	CrawlRunParamsExtractOptionsLocaleZhHansMo  CrawlRunParamsExtractOptionsLocale = "zh-Hans-MO"
	CrawlRunParamsExtractOptionsLocaleZhHansSg  CrawlRunParamsExtractOptionsLocale = "zh-Hans-SG"
	CrawlRunParamsExtractOptionsLocaleZhHant    CrawlRunParamsExtractOptionsLocale = "zh-Hant"
	CrawlRunParamsExtractOptionsLocaleZhHantHk  CrawlRunParamsExtractOptionsLocale = "zh-Hant-HK"
	CrawlRunParamsExtractOptionsLocaleZhHantMo  CrawlRunParamsExtractOptionsLocale = "zh-Hant-MO"
	CrawlRunParamsExtractOptionsLocaleZhHantTw  CrawlRunParamsExtractOptionsLocale = "zh-Hant-TW"
	CrawlRunParamsExtractOptionsLocaleZhSg      CrawlRunParamsExtractOptionsLocale = "zh-SG"
	CrawlRunParamsExtractOptionsLocaleZhTw      CrawlRunParamsExtractOptionsLocale = "zh-TW"
	CrawlRunParamsExtractOptionsLocaleZu        CrawlRunParamsExtractOptionsLocale = "zu"
	CrawlRunParamsExtractOptionsLocaleZuZa      CrawlRunParamsExtractOptionsLocale = "zu-ZA"
	CrawlRunParamsExtractOptionsLocaleAuto      CrawlRunParamsExtractOptionsLocale = "auto"
)

type CrawlRunParamsExtractOptionsNetworkCapture struct {
	Validation                  param.Opt[bool]    `json:"validation,omitzero"`
	WaitForRequestsCount        param.Opt[float64] `json:"wait_for_requests_count,omitzero"`
	WaitForRequestsCountTimeout param.Opt[float64] `json:"wait_for_requests_count_timeout,omitzero"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method,omitzero"`
	// Resource type for network capture filtering
	ResourceType CrawlRunParamsExtractOptionsNetworkCaptureResourceTypeUnion `json:"resource_type,omitzero"`
	StatusCode   CrawlRunParamsExtractOptionsNetworkCaptureStatusCodeUnion   `json:"status_code,omitzero"`
	URL          CrawlRunParamsExtractOptionsNetworkCaptureURL               `json:"url,omitzero"`
	paramObj
}

func (r CrawlRunParamsExtractOptionsNetworkCapture) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRunParamsExtractOptionsNetworkCapture
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRunParamsExtractOptionsNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlRunParamsExtractOptionsNetworkCapture](
		"method", "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE", "PATCH",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRunParamsExtractOptionsNetworkCaptureResourceTypeUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRunParamsExtractOptionsNetworkCaptureResourceTypeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlRunParamsExtractOptionsNetworkCaptureResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRunParamsExtractOptionsNetworkCaptureResourceTypeUnion) asAny() any {
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
type CrawlRunParamsExtractOptionsNetworkCaptureStatusCodeUnion struct {
	OfFloat      param.Opt[float64] `json:",omitzero,inline"`
	OfFloatArray []float64          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRunParamsExtractOptionsNetworkCaptureStatusCodeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfFloat, u.OfFloatArray)
}
func (u *CrawlRunParamsExtractOptionsNetworkCaptureStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRunParamsExtractOptionsNetworkCaptureStatusCodeUnion) asAny() any {
	if !param.IsOmitted(u.OfFloat) {
		return &u.OfFloat.Value
	} else if !param.IsOmitted(u.OfFloatArray) {
		return &u.OfFloatArray
	}
	return nil
}

// The property Value is required.
type CrawlRunParamsExtractOptionsNetworkCaptureURL struct {
	Value string `json:"value" api:"required"`
	// Any of "exact", "contains".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r CrawlRunParamsExtractOptionsNetworkCaptureURL) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRunParamsExtractOptionsNetworkCaptureURL
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRunParamsExtractOptionsNetworkCaptureURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CrawlRunParamsExtractOptionsNetworkCaptureURL](
		"type", "exact", "contains",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRunParamsExtractOptionsParserUnion struct {
	OfAnyMap map[string]any    `json:",omitzero,inline"`
	OfString param.Opt[string] `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRunParamsExtractOptionsParserUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap, u.OfString)
}
func (u *CrawlRunParamsExtractOptionsParserUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRunParamsExtractOptionsParserUnion) asAny() any {
	if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	}
	return nil
}

// Referrer policy for the request
type CrawlRunParamsExtractOptionsReferrerType string

const (
	CrawlRunParamsExtractOptionsReferrerTypeRandom     CrawlRunParamsExtractOptionsReferrerType = "random"
	CrawlRunParamsExtractOptionsReferrerTypeNoReferer  CrawlRunParamsExtractOptionsReferrerType = "no-referer"
	CrawlRunParamsExtractOptionsReferrerTypeSameOrigin CrawlRunParamsExtractOptionsReferrerType = "same-origin"
	CrawlRunParamsExtractOptionsReferrerTypeGoogle     CrawlRunParamsExtractOptionsReferrerType = "google"
	CrawlRunParamsExtractOptionsReferrerTypeBing       CrawlRunParamsExtractOptionsReferrerType = "bing"
	CrawlRunParamsExtractOptionsReferrerTypeFacebook   CrawlRunParamsExtractOptionsReferrerType = "facebook"
	CrawlRunParamsExtractOptionsReferrerTypeTwitter    CrawlRunParamsExtractOptionsReferrerType = "twitter"
	CrawlRunParamsExtractOptionsReferrerTypeInstagram  CrawlRunParamsExtractOptionsReferrerType = "instagram"
)

type CrawlRunParamsExtractOptionsSession struct {
	ID                  param.Opt[string]  `json:"id,omitzero"`
	PrefetchUserbrowser param.Opt[bool]    `json:"prefetch_userbrowser,omitzero"`
	Retry               param.Opt[bool]    `json:"retry,omitzero"`
	Timeout             param.Opt[float64] `json:"timeout,omitzero"`
	paramObj
}

func (r CrawlRunParamsExtractOptionsSession) MarshalJSON() (data []byte, err error) {
	type shadow CrawlRunParamsExtractOptionsSession
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CrawlRunParamsExtractOptionsSession) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CrawlRunParamsExtractOptionsSkillUnion struct {
	OfString      param.Opt[string] `json:",omitzero,inline"`
	OfStringArray []string          `json:",omitzero,inline"`
	paramUnion
}

func (u CrawlRunParamsExtractOptionsSkillUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfStringArray)
}
func (u *CrawlRunParamsExtractOptionsSkillUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CrawlRunParamsExtractOptionsSkillUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfStringArray) {
		return &u.OfStringArray
	}
	return nil
}

// Sitemap and other methods will be used together to find URLs.
type CrawlRunParamsSitemap string

const (
	CrawlRunParamsSitemapSkip    CrawlRunParamsSitemap = "skip"
	CrawlRunParamsSitemapInclude CrawlRunParamsSitemap = "include"
	CrawlRunParamsSitemapOnly    CrawlRunParamsSitemap = "only"
)
