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

// AgentTemplateService contains methods and other services that help with
// interacting with the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentTemplateService] method instead.
type AgentTemplateService struct {
	Options []option.RequestOption
}

// NewAgentTemplateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentTemplateService(opts ...option.RequestOption) (r AgentTemplateService) {
	r = AgentTemplateService{}
	r.Options = opts
	return
}

// List the pre-built agent templates available to your account. Use a template's
// `template_name` with `POST /v2/agents` to create an agent instance from it.
func (r *AgentTemplateService) List(ctx context.Context, query AgentTemplateListParams, opts ...option.RequestOption) (res *AgentTemplateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/agents/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve a single agent template by its stable `template_name`.
func (r *AgentTemplateService) Get(ctx context.Context, templateName string, opts ...option.RequestOption) (res *AgentTemplateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateName == "" {
		err = errors.New("missing required template_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/agents/templates/%s", templateName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AgentTemplateListResponse struct {
	// Items returned in this page.
	Items []AgentTemplateListResponseItem `json:"items" api:"required"`
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
func (r AgentTemplateListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentTemplateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentTemplateListResponseItem struct {
	// Unique template identifier (wsat\_<uuid>).
	ID string `json:"id" api:"required"`
	// When the template was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Template description shown to users.
	Description string `json:"description" api:"required"`
	// Human-friendly template name shown to users.
	DisplayName string `json:"display_name" api:"required"`
	// Default effort level for runs created from this template.
	//
	// Any of "low", "medium", "high", "x-high", "5x-high", "max".
	Effort string `json:"effort" api:"required"`
	// Ordered goals for the template.
	Goals []AgentTemplateListResponseItemGoal `json:"goals" api:"required"`
	// Icon identifier used when presenting the template.
	Icon string `json:"icon" api:"required"`
	// JSON schema describing the structured output the agent should produce.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Skill or operating context for the template.
	Skill string `json:"skill" api:"required"`
	// Ordered source groups for the template.
	Sources []AgentTemplateListResponseItemSource `json:"sources" api:"required"`
	// Suggested prompts for the template.
	SuggestedQuestions []AgentTemplateListResponseItemSuggestedQuestion `json:"suggested_questions" api:"required"`
	// Stable template name used to create agent instances.
	TemplateName string `json:"template_name" api:"required"`
	// When the template was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Primary use case supported by the template.
	//
	// Any of "research", "enrichment", "dataset_building".
	UseCase string `json:"use_case" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		DisplayName        respjson.Field
		Effort             respjson.Field
		Goals              respjson.Field
		Icon               respjson.Field
		OutputSchema       respjson.Field
		Skill              respjson.Field
		Sources            respjson.Field
		SuggestedQuestions respjson.Field
		TemplateName       respjson.Field
		UpdatedAt          respjson.Field
		UseCase            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTemplateListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *AgentTemplateListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentTemplateListResponseItemGoal struct {
	// Unique goal identifier (wsag\_<uuid>).
	ID string `json:"id" api:"required"`
	// Goal text.
	Goal string `json:"goal" api:"required"`
	// Zero-based goal position.
	Order int64 `json:"order" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Goal        respjson.Field
		Order       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTemplateListResponseItemGoal) RawJSON() string { return r.JSON.raw }
func (r *AgentTemplateListResponseItemGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentTemplateListResponseItemSource struct {
	// Unique source group identifier (wsas\_<uuid>).
	ID string `json:"id" api:"required"`
	// Domains included in this source group.
	Domains []string `json:"domains" api:"required"`
	// Zero-based source group position.
	Order int64 `json:"order" api:"required"`
	// Source group title.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Domains     respjson.Field
		Order       respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTemplateListResponseItemSource) RawJSON() string { return r.JSON.raw }
func (r *AgentTemplateListResponseItemSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentTemplateListResponseItemSuggestedQuestion struct {
	// Unique suggested question identifier (wsasq\_<uuid>).
	ID string `json:"id" api:"required"`
	// Zero-based suggested question position.
	Order int64 `json:"order" api:"required"`
	// Suggested prompt text.
	Question string `json:"question" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Order       respjson.Field
		Question    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTemplateListResponseItemSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *AgentTemplateListResponseItemSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentTemplateGetResponse struct {
	// Unique template identifier (wsat\_<uuid>).
	ID string `json:"id" api:"required"`
	// When the template was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Template description shown to users.
	Description string `json:"description" api:"required"`
	// Human-friendly template name shown to users.
	DisplayName string `json:"display_name" api:"required"`
	// Default effort level for runs created from this template.
	//
	// Any of "low", "medium", "high", "x-high", "5x-high", "max".
	Effort AgentTemplateGetResponseEffort `json:"effort" api:"required"`
	// Ordered goals for the template.
	Goals []AgentTemplateGetResponseGoal `json:"goals" api:"required"`
	// Icon identifier used when presenting the template.
	Icon string `json:"icon" api:"required"`
	// JSON schema describing the structured output the agent should produce.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Skill or operating context for the template.
	Skill string `json:"skill" api:"required"`
	// Ordered source groups for the template.
	Sources []AgentTemplateGetResponseSource `json:"sources" api:"required"`
	// Suggested prompts for the template.
	SuggestedQuestions []AgentTemplateGetResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	// Stable template name used to create agent instances.
	TemplateName string `json:"template_name" api:"required"`
	// When the template was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Primary use case supported by the template.
	//
	// Any of "research", "enrichment", "dataset_building".
	UseCase AgentTemplateGetResponseUseCase `json:"use_case" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		DisplayName        respjson.Field
		Effort             respjson.Field
		Goals              respjson.Field
		Icon               respjson.Field
		OutputSchema       respjson.Field
		Skill              respjson.Field
		Sources            respjson.Field
		SuggestedQuestions respjson.Field
		TemplateName       respjson.Field
		UpdatedAt          respjson.Field
		UseCase            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTemplateGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentTemplateGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default effort level for runs created from this template.
type AgentTemplateGetResponseEffort string

const (
	AgentTemplateGetResponseEffortLow    AgentTemplateGetResponseEffort = "low"
	AgentTemplateGetResponseEffortMedium AgentTemplateGetResponseEffort = "medium"
	AgentTemplateGetResponseEffortHigh   AgentTemplateGetResponseEffort = "high"
	AgentTemplateGetResponseEffortXHigh  AgentTemplateGetResponseEffort = "x-high"
	AgentTemplateGetResponseEffort5xHigh AgentTemplateGetResponseEffort = "5x-high"
	AgentTemplateGetResponseEffortMax    AgentTemplateGetResponseEffort = "max"
)

type AgentTemplateGetResponseGoal struct {
	// Unique goal identifier (wsag\_<uuid>).
	ID string `json:"id" api:"required"`
	// Goal text.
	Goal string `json:"goal" api:"required"`
	// Zero-based goal position.
	Order int64 `json:"order" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Goal        respjson.Field
		Order       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTemplateGetResponseGoal) RawJSON() string { return r.JSON.raw }
func (r *AgentTemplateGetResponseGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentTemplateGetResponseSource struct {
	// Unique source group identifier (wsas\_<uuid>).
	ID string `json:"id" api:"required"`
	// Domains included in this source group.
	Domains []string `json:"domains" api:"required"`
	// Zero-based source group position.
	Order int64 `json:"order" api:"required"`
	// Source group title.
	Title string `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Domains     respjson.Field
		Order       respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTemplateGetResponseSource) RawJSON() string { return r.JSON.raw }
func (r *AgentTemplateGetResponseSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentTemplateGetResponseSuggestedQuestion struct {
	// Unique suggested question identifier (wsasq\_<uuid>).
	ID string `json:"id" api:"required"`
	// Zero-based suggested question position.
	Order int64 `json:"order" api:"required"`
	// Suggested prompt text.
	Question string `json:"question" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Order       respjson.Field
		Question    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentTemplateGetResponseSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *AgentTemplateGetResponseSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary use case supported by the template.
type AgentTemplateGetResponseUseCase string

const (
	AgentTemplateGetResponseUseCaseResearch        AgentTemplateGetResponseUseCase = "research"
	AgentTemplateGetResponseUseCaseEnrichment      AgentTemplateGetResponseUseCase = "enrichment"
	AgentTemplateGetResponseUseCaseDatasetBuilding AgentTemplateGetResponseUseCase = "dataset_building"
)

type AgentTemplateListParams struct {
	Limit  param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentTemplateListParams]'s query parameters as
// `url.Values`.
func (r AgentTemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
