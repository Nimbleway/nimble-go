// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomnimblewaynimblego

import (
	"context"
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

// Execute WSA Realtime Endpoint
func (r *AgentService) Async(ctx context.Context, body AgentAsyncParams, opts ...option.RequestOption) (res *AgentAsyncResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/agent/async"
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

type AgentListResponse struct {
	DisplayName string `json:"display_name,required"`
	IsPublic    bool   `json:"is_public,required"`
	Name        string `json:"name,required"`
	Description string `json:"description,nullable"`
	Domain      string `json:"domain,nullable"`
	EntityType  string `json:"entity_type,nullable"`
	Vertical    string `json:"vertical,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		IsPublic    respjson.Field
		Name        respjson.Field
		Description respjson.Field
		Domain      respjson.Field
		EntityType  respjson.Field
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
	Default     string   `json:"default,nullable"`
	Description string   `json:"description,nullable"`
	Examples    []string `json:"examples,nullable"`
	Name        string   `json:"name"`
	Required    bool     `json:"required"`
	Rules       []string `json:"rules,nullable"`
	Type        string   `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Default     respjson.Field
		Description respjson.Field
		Examples    respjson.Field
		Name        respjson.Field
		Required    respjson.Field
		Rules       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGetResponseInputProperty) RawJSON() string { return r.JSON.raw }
func (r *AgentGetResponseInputProperty) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentListParams struct {
	// Number of results per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Pagination offset
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
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
