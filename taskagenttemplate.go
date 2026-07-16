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

// TaskAgentTemplateService contains methods and other services that help with
// interacting with the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskAgentTemplateService] method instead.
type TaskAgentTemplateService struct {
	Options []option.RequestOption
}

// NewTaskAgentTemplateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTaskAgentTemplateService(opts ...option.RequestOption) (r TaskAgentTemplateService) {
	r = TaskAgentTemplateService{}
	r.Options = opts
	return
}

// List Templates
//
// Deprecated: deprecated
func (r *TaskAgentTemplateService) List(ctx context.Context, query TaskAgentTemplateListParams, opts ...option.RequestOption) (res *TaskAgentTemplateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/task-agents/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get Template
//
// Deprecated: deprecated
func (r *TaskAgentTemplateService) Get(ctx context.Context, templateName string, opts ...option.RequestOption) (res *TaskAgentTemplateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateName == "" {
		err = errors.New("missing required template_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/task-agents/templates/%s", templateName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type TaskAgentTemplateListResponse struct {
	// Items returned in this page.
	Items []TaskAgentTemplateListResponseItem `json:"items" api:"required"`
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
func (r TaskAgentTemplateListResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentTemplateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentTemplateListResponseItem struct {
	// Unique template identifier (wsat\_<uuid>).
	ID string `json:"id" api:"required"`
	// When the template was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Template description shown to users.
	Description string `json:"description" api:"required"`
	// Human-friendly template name shown to users.
	DisplayName string `json:"display_name" api:"required"`
	// Domain expertise or operating context for the template.
	DomainExpertise string `json:"domain_expertise" api:"required"`
	// Default effort level for runs created from this template.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort string `json:"effort" api:"required"`
	// Ordered goals for the template.
	Goals []TaskAgentTemplateListResponseItemGoal `json:"goals" api:"required"`
	// Icon identifier used when presenting the template.
	Icon string `json:"icon" api:"required"`
	// JSON schema describing the structured output the agent should produce.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Ordered source groups for the template.
	Sources []TaskAgentTemplateListResponseItemSource `json:"sources" api:"required"`
	// Suggested prompts for the template.
	SuggestedQuestions []TaskAgentTemplateListResponseItemSuggestedQuestion `json:"suggested_questions" api:"required"`
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
		DomainExpertise    respjson.Field
		Effort             respjson.Field
		Goals              respjson.Field
		Icon               respjson.Field
		OutputSchema       respjson.Field
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
func (r TaskAgentTemplateListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentTemplateListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentTemplateListResponseItemGoal struct {
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
func (r TaskAgentTemplateListResponseItemGoal) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentTemplateListResponseItemGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentTemplateListResponseItemSource struct {
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
func (r TaskAgentTemplateListResponseItemSource) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentTemplateListResponseItemSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentTemplateListResponseItemSuggestedQuestion struct {
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
func (r TaskAgentTemplateListResponseItemSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentTemplateListResponseItemSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentTemplateGetResponse struct {
	// Unique template identifier (wsat\_<uuid>).
	ID string `json:"id" api:"required"`
	// When the template was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Template description shown to users.
	Description string `json:"description" api:"required"`
	// Human-friendly template name shown to users.
	DisplayName string `json:"display_name" api:"required"`
	// Domain expertise or operating context for the template.
	DomainExpertise string `json:"domain_expertise" api:"required"`
	// Default effort level for runs created from this template.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort TaskAgentTemplateGetResponseEffort `json:"effort" api:"required"`
	// Ordered goals for the template.
	Goals []TaskAgentTemplateGetResponseGoal `json:"goals" api:"required"`
	// Icon identifier used when presenting the template.
	Icon string `json:"icon" api:"required"`
	// JSON schema describing the structured output the agent should produce.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Ordered source groups for the template.
	Sources []TaskAgentTemplateGetResponseSource `json:"sources" api:"required"`
	// Suggested prompts for the template.
	SuggestedQuestions []TaskAgentTemplateGetResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	// Stable template name used to create agent instances.
	TemplateName string `json:"template_name" api:"required"`
	// When the template was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Primary use case supported by the template.
	//
	// Any of "research", "enrichment", "dataset_building".
	UseCase TaskAgentTemplateGetResponseUseCase `json:"use_case" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		DisplayName        respjson.Field
		DomainExpertise    respjson.Field
		Effort             respjson.Field
		Goals              respjson.Field
		Icon               respjson.Field
		OutputSchema       respjson.Field
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
func (r TaskAgentTemplateGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentTemplateGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default effort level for runs created from this template.
type TaskAgentTemplateGetResponseEffort string

const (
	TaskAgentTemplateGetResponseEffortLow    TaskAgentTemplateGetResponseEffort = "low"
	TaskAgentTemplateGetResponseEffortMedium TaskAgentTemplateGetResponseEffort = "medium"
	TaskAgentTemplateGetResponseEffortHigh   TaskAgentTemplateGetResponseEffort = "high"
	TaskAgentTemplateGetResponseEffortXHigh  TaskAgentTemplateGetResponseEffort = "x-high"
	TaskAgentTemplateGetResponseEffortMax    TaskAgentTemplateGetResponseEffort = "max"
)

type TaskAgentTemplateGetResponseGoal struct {
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
func (r TaskAgentTemplateGetResponseGoal) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentTemplateGetResponseGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentTemplateGetResponseSource struct {
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
func (r TaskAgentTemplateGetResponseSource) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentTemplateGetResponseSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentTemplateGetResponseSuggestedQuestion struct {
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
func (r TaskAgentTemplateGetResponseSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentTemplateGetResponseSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary use case supported by the template.
type TaskAgentTemplateGetResponseUseCase string

const (
	TaskAgentTemplateGetResponseUseCaseResearch        TaskAgentTemplateGetResponseUseCase = "research"
	TaskAgentTemplateGetResponseUseCaseEnrichment      TaskAgentTemplateGetResponseUseCase = "enrichment"
	TaskAgentTemplateGetResponseUseCaseDatasetBuilding TaskAgentTemplateGetResponseUseCase = "dataset_building"
)

type TaskAgentTemplateListParams struct {
	Limit  param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaskAgentTemplateListParams]'s query parameters as
// `url.Values`.
func (r TaskAgentTemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
