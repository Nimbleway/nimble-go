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

// List all available Web Search Agent templates.
func (r *TaskAgentTemplateService) List(ctx context.Context, query TaskAgentTemplateListParams, opts ...option.RequestOption) (res *[]TaskAgentTemplateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/task-agents/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Fetch a single Web Search Agent template by name.
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
	ID                 string                                           `json:"id" api:"required"`
	CreatedAt          time.Time                                        `json:"created_at" api:"required" format:"date-time"`
	Description        string                                           `json:"description" api:"required"`
	DisplayName        string                                           `json:"display_name" api:"required"`
	DomainExpertise    string                                           `json:"domain_expertise" api:"required"`
	Effort             string                                           `json:"effort" api:"required"`
	Goals              []TaskAgentTemplateListResponseGoal              `json:"goals" api:"required"`
	Icon               string                                           `json:"icon" api:"required"`
	OutputSchema       map[string]any                                   `json:"output_schema" api:"required"`
	Sources            []TaskAgentTemplateListResponseSource            `json:"sources" api:"required"`
	SuggestedQuestions []TaskAgentTemplateListResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	TemplateName       string                                           `json:"template_name" api:"required"`
	UpdatedAt          time.Time                                        `json:"updated_at" api:"required" format:"date-time"`
	// Any of "research", "enrichment", "dataset_building".
	UseCase TaskAgentTemplateListResponseUseCase `json:"use_case" api:"required"`
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
func (r TaskAgentTemplateListResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentTemplateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentTemplateListResponseGoal struct {
	ID    string `json:"id" api:"required"`
	Goal  string `json:"goal" api:"required"`
	Order int64  `json:"order" api:"required"`
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
func (r TaskAgentTemplateListResponseGoal) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentTemplateListResponseGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentTemplateListResponseSource struct {
	ID      string   `json:"id" api:"required"`
	Domains []string `json:"domains" api:"required"`
	Order   int64    `json:"order" api:"required"`
	Title   string   `json:"title" api:"required"`
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
func (r TaskAgentTemplateListResponseSource) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentTemplateListResponseSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentTemplateListResponseSuggestedQuestion struct {
	ID       string `json:"id" api:"required"`
	Order    int64  `json:"order" api:"required"`
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
func (r TaskAgentTemplateListResponseSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentTemplateListResponseSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentTemplateListResponseUseCase string

const (
	TaskAgentTemplateListResponseUseCaseResearch        TaskAgentTemplateListResponseUseCase = "research"
	TaskAgentTemplateListResponseUseCaseEnrichment      TaskAgentTemplateListResponseUseCase = "enrichment"
	TaskAgentTemplateListResponseUseCaseDatasetBuilding TaskAgentTemplateListResponseUseCase = "dataset_building"
)

type TaskAgentTemplateGetResponse struct {
	ID                 string                                          `json:"id" api:"required"`
	CreatedAt          time.Time                                       `json:"created_at" api:"required" format:"date-time"`
	Description        string                                          `json:"description" api:"required"`
	DisplayName        string                                          `json:"display_name" api:"required"`
	DomainExpertise    string                                          `json:"domain_expertise" api:"required"`
	Effort             string                                          `json:"effort" api:"required"`
	Goals              []TaskAgentTemplateGetResponseGoal              `json:"goals" api:"required"`
	Icon               string                                          `json:"icon" api:"required"`
	OutputSchema       map[string]any                                  `json:"output_schema" api:"required"`
	Sources            []TaskAgentTemplateGetResponseSource            `json:"sources" api:"required"`
	SuggestedQuestions []TaskAgentTemplateGetResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	TemplateName       string                                          `json:"template_name" api:"required"`
	UpdatedAt          time.Time                                       `json:"updated_at" api:"required" format:"date-time"`
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

type TaskAgentTemplateGetResponseGoal struct {
	ID    string `json:"id" api:"required"`
	Goal  string `json:"goal" api:"required"`
	Order int64  `json:"order" api:"required"`
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
	ID      string   `json:"id" api:"required"`
	Domains []string `json:"domains" api:"required"`
	Order   int64    `json:"order" api:"required"`
	Title   string   `json:"title" api:"required"`
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
	ID       string `json:"id" api:"required"`
	Order    int64  `json:"order" api:"required"`
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

type TaskAgentTemplateGetResponseUseCase string

const (
	TaskAgentTemplateGetResponseUseCaseResearch        TaskAgentTemplateGetResponseUseCase = "research"
	TaskAgentTemplateGetResponseUseCaseEnrichment      TaskAgentTemplateGetResponseUseCase = "enrichment"
	TaskAgentTemplateGetResponseUseCaseDatasetBuilding TaskAgentTemplateGetResponseUseCase = "dataset_building"
)

type TaskAgentTemplateListParams struct {
	Effort  param.Opt[string] `query:"effort,omitzero" json:"-"`
	UseCase param.Opt[string] `query:"use_case,omitzero" json:"-"`
	Limit   param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Offset  param.Opt[int64]  `query:"offset,omitzero" json:"-"`
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
