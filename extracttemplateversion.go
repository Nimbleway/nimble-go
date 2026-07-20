// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomnimblewaynimblego

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/Nimbleway/nimble-go/internal/apijson"
	"github.com/Nimbleway/nimble-go/internal/apiquery"
	"github.com/Nimbleway/nimble-go/internal/requestconfig"
	"github.com/Nimbleway/nimble-go/option"
	"github.com/Nimbleway/nimble-go/packages/param"
	"github.com/Nimbleway/nimble-go/packages/respjson"
)

// ExtractTemplateVersionService contains methods and other services that help with
// interacting with the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtractTemplateVersionService] method instead.
type ExtractTemplateVersionService struct {
	Options []option.RequestOption
}

// NewExtractTemplateVersionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewExtractTemplateVersionService(opts ...option.RequestOption) (r ExtractTemplateVersionService) {
	r = ExtractTemplateVersionService{}
	r.Options = opts
	return
}

// List Extract Template Versions Public V2
func (r *ExtractTemplateVersionService) List(ctx context.Context, extractTemplateName string, query ExtractTemplateVersionListParams, opts ...option.RequestOption) (res *ExtractTemplateVersionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if extractTemplateName == "" {
		err = errors.New("missing required extract_template_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/extract/templates/%s/versions", extractTemplateName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get Extract Template Version Public V2
func (r *ExtractTemplateVersionService) Get(ctx context.Context, versionID string, query ExtractTemplateVersionGetParams, opts ...option.RequestOption) (res *ExtractTemplateVersionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ExtractTemplateName == "" {
		err = errors.New("missing required extract_template_name parameter")
		return nil, err
	}
	if versionID == "" {
		err = errors.New("missing required version_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/extract/templates/%s/versions/%s", query.ExtractTemplateName, versionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ExtractTemplateVersionListResponse struct {
	// Items returned in this page.
	Items []ExtractTemplateVersionListResponseItem `json:"items" api:"required"`
	// Maximum number of items returned.
	Limit int64 `json:"limit" api:"required"`
	// Number of items skipped before this page.
	Offset int64 `json:"offset" api:"required"`
	// Total number of items matching the query.
	Total int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Limit       respjson.Field
		Offset      respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateVersionListResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateVersionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateVersionListResponseItem struct {
	// Unique extract template version identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the version was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// JSON schema describing accepted input parameters.
	InputSchema map[string]any `json:"input_schema" api:"required"`
	// Metadata associated with this version.
	Metadata ExtractTemplateVersionListResponseItemMetadata `json:"metadata" api:"required"`
	// Extract template name this version belongs to.
	Name string `json:"name" api:"required"`
	// JSON schema describing extracted output.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Monotonic version number for the extract template.
	VersionNumber int64 `json:"version_number" api:"required"`
	// Sample input and output pairs for the version.
	Samples []ExtractTemplateVersionListResponseItemSample `json:"samples" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		InputSchema   respjson.Field
		Metadata      respjson.Field
		Name          respjson.Field
		OutputSchema  respjson.Field
		VersionNumber respjson.Field
		Samples       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateVersionListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateVersionListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata associated with this version.
type ExtractTemplateVersionListResponseItemMetadata struct {
	// Data source associated with the version.
	DataSource string `json:"data_source" api:"nullable"`
	// Version description shown to users.
	Description string `json:"description" api:"nullable"`
	// Human-friendly version display name.
	DisplayName string `json:"display_name" api:"nullable"`
	// Domain associated with the version.
	Domain string `json:"domain" api:"nullable"`
	// Entity type produced by the version.
	EntityType string `json:"entity_type" api:"nullable"`
	// Tags associated with the version.
	Tags []string `json:"tags"`
	// Business vertical associated with the version.
	Vertical string `json:"vertical" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataSource  respjson.Field
		Description respjson.Field
		DisplayName respjson.Field
		Domain      respjson.Field
		EntityType  respjson.Field
		Tags        respjson.Field
		Vertical    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateVersionListResponseItemMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateVersionListResponseItemMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateVersionListResponseItemSample struct {
	// Sample input parameters for the version.
	Input any `json:"input"`
	// Sample output produced by the version.
	Output any `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Input       respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateVersionListResponseItemSample) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateVersionListResponseItemSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateVersionGetResponse struct {
	// Unique extract template version identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the version was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// JSON schema describing accepted input parameters.
	InputSchema map[string]any `json:"input_schema" api:"required"`
	// Metadata associated with this version.
	Metadata ExtractTemplateVersionGetResponseMetadata `json:"metadata" api:"required"`
	// Extract template name this version belongs to.
	Name string `json:"name" api:"required"`
	// JSON schema describing extracted output.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Monotonic version number for the extract template.
	VersionNumber int64 `json:"version_number" api:"required"`
	// Sample input and output pairs for the version.
	Samples []ExtractTemplateVersionGetResponseSample `json:"samples" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		InputSchema   respjson.Field
		Metadata      respjson.Field
		Name          respjson.Field
		OutputSchema  respjson.Field
		VersionNumber respjson.Field
		Samples       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateVersionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateVersionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata associated with this version.
type ExtractTemplateVersionGetResponseMetadata struct {
	// Data source associated with the version.
	DataSource string `json:"data_source" api:"nullable"`
	// Version description shown to users.
	Description string `json:"description" api:"nullable"`
	// Human-friendly version display name.
	DisplayName string `json:"display_name" api:"nullable"`
	// Domain associated with the version.
	Domain string `json:"domain" api:"nullable"`
	// Entity type produced by the version.
	EntityType string `json:"entity_type" api:"nullable"`
	// Tags associated with the version.
	Tags []string `json:"tags"`
	// Business vertical associated with the version.
	Vertical string `json:"vertical" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataSource  respjson.Field
		Description respjson.Field
		DisplayName respjson.Field
		Domain      respjson.Field
		EntityType  respjson.Field
		Tags        respjson.Field
		Vertical    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateVersionGetResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateVersionGetResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateVersionGetResponseSample struct {
	// Sample input parameters for the version.
	Input any `json:"input"`
	// Sample output produced by the version.
	Output any `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Input       respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateVersionGetResponseSample) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateVersionGetResponseSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateVersionListParams struct {
	Limit  param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExtractTemplateVersionListParams]'s query parameters as
// `url.Values`.
func (r ExtractTemplateVersionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExtractTemplateVersionGetParams struct {
	ExtractTemplateName string `path:"extract_template_name" api:"required" json:"-"`
	paramObj
}
