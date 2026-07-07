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
	shimjson "github.com/Nimbleway/nimble-go/internal/encoding/json"
	"github.com/Nimbleway/nimble-go/internal/requestconfig"
	"github.com/Nimbleway/nimble-go/option"
	"github.com/Nimbleway/nimble-go/packages/param"
	"github.com/Nimbleway/nimble-go/packages/respjson"
)

// TaskAgentService contains methods and other services that help with interacting
// with the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskAgentService] method instead.
type TaskAgentService struct {
	Options   []option.RequestOption
	Templates TaskAgentTemplateService
	Runs      TaskAgentRunService
}

// NewTaskAgentService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTaskAgentService(opts ...option.RequestOption) (r TaskAgentService) {
	r = TaskAgentService{}
	r.Options = opts
	r.Templates = NewTaskAgentTemplateService(opts...)
	r.Runs = NewTaskAgentRunService(opts...)
	return
}

// Create a new workspace-scoped Web Search Agent. Pass `template` to clone from a
// named template.
func (r *TaskAgentService) New(ctx context.Context, body TaskAgentNewParams, opts ...option.RequestOption) (res *TaskAgentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/task-agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Apply a JSON Patch document (`application/json-patch+json`) to an agent you own.
// Each operation must be a `replace` with path `/field_name`.
func (r *TaskAgentService) Update(ctx context.Context, agentID string, body TaskAgentUpdateParams, opts ...option.RequestOption) (res *TaskAgentUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/task-agents/%s", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List active Web Search Agents visible to the caller. Includes agents scoped to
// the caller's workspace.
func (r *TaskAgentService) List(ctx context.Context, query TaskAgentListParams, opts ...option.RequestOption) (res *[]TaskAgentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/task-agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deactivate an agent you own. The agent is marked inactive but not deleted.
func (r *TaskAgentService) Deactivate(ctx context.Context, agentID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/task-agents/%s", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Fetch a single Web Search Agent by id.
func (r *TaskAgentService) Get(ctx context.Context, agentID string, opts ...option.RequestOption) (res *TaskAgentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/task-agents/%s", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Create and enqueue a research run for a Web Search Agent.
func (r *TaskAgentService) Run(ctx context.Context, agentID string, body TaskAgentRunParams, opts ...option.RequestOption) (res *TaskAgentRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/task-agents/%s/runs", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type TaskAgentNewResponse struct {
	ID                 string                                  `json:"id" api:"required"`
	CreatedAt          time.Time                               `json:"created_at" api:"required" format:"date-time"`
	Description        string                                  `json:"description" api:"required"`
	DisplayName        string                                  `json:"display_name" api:"required"`
	DomainExpertise    string                                  `json:"domain_expertise" api:"required"`
	Effort             string                                  `json:"effort" api:"required"`
	Goals              []TaskAgentNewResponseGoal              `json:"goals" api:"required"`
	Icon               string                                  `json:"icon" api:"required"`
	IsActive           bool                                    `json:"is_active" api:"required"`
	OutputSchema       map[string]any                          `json:"output_schema" api:"required"`
	Sources            TaskAgentNewResponseSources             `json:"sources" api:"required"`
	SuggestedQuestions []TaskAgentNewResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	UpdatedAt          time.Time                               `json:"updated_at" api:"required" format:"date-time"`
	// Any of "research", "enrichment", "dataset_building".
	UseCase       TaskAgentNewResponseUseCase `json:"use_case" api:"required"`
	AccountID     string                      `json:"account_id" api:"nullable"`
	AgentName     string                      `json:"agent_name" api:"nullable"`
	WorkspaceID   string                      `json:"workspace_id" api:"nullable" format:"uuid"`
	WorkspaceName string                      `json:"workspace_name" api:"nullable"`
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
		IsActive           respjson.Field
		OutputSchema       respjson.Field
		Sources            respjson.Field
		SuggestedQuestions respjson.Field
		UpdatedAt          respjson.Field
		UseCase            respjson.Field
		AccountID          respjson.Field
		AgentName          respjson.Field
		WorkspaceID        respjson.Field
		WorkspaceName      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentNewResponseGoal struct {
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
func (r TaskAgentNewResponseGoal) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentNewResponseGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentNewResponseSources struct {
	Allow      []TaskAgentNewResponseSourcesAllow `json:"allow"`
	Avoid      string                             `json:"avoid" api:"nullable"`
	Block      []TaskAgentNewResponseSourcesBlock `json:"block"`
	Prioritize string                             `json:"prioritize" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allow       respjson.Field
		Avoid       respjson.Field
		Block       respjson.Field
		Prioritize  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentNewResponseSources) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentNewResponseSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentNewResponseSourcesAllow struct {
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
func (r TaskAgentNewResponseSourcesAllow) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentNewResponseSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentNewResponseSourcesBlock struct {
	Domains []string `json:"domains" api:"required"`
	Title   string   `json:"title" api:"required"`
	Order   int64    `json:"order"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domains     respjson.Field
		Title       respjson.Field
		Order       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentNewResponseSourcesBlock) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentNewResponseSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentNewResponseSuggestedQuestion struct {
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
func (r TaskAgentNewResponseSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentNewResponseSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentNewResponseUseCase string

const (
	TaskAgentNewResponseUseCaseResearch        TaskAgentNewResponseUseCase = "research"
	TaskAgentNewResponseUseCaseEnrichment      TaskAgentNewResponseUseCase = "enrichment"
	TaskAgentNewResponseUseCaseDatasetBuilding TaskAgentNewResponseUseCase = "dataset_building"
)

type TaskAgentUpdateResponse struct {
	ID                 string                                     `json:"id" api:"required"`
	CreatedAt          time.Time                                  `json:"created_at" api:"required" format:"date-time"`
	Description        string                                     `json:"description" api:"required"`
	DisplayName        string                                     `json:"display_name" api:"required"`
	DomainExpertise    string                                     `json:"domain_expertise" api:"required"`
	Effort             string                                     `json:"effort" api:"required"`
	Goals              []TaskAgentUpdateResponseGoal              `json:"goals" api:"required"`
	Icon               string                                     `json:"icon" api:"required"`
	IsActive           bool                                       `json:"is_active" api:"required"`
	OutputSchema       map[string]any                             `json:"output_schema" api:"required"`
	Sources            TaskAgentUpdateResponseSources             `json:"sources" api:"required"`
	SuggestedQuestions []TaskAgentUpdateResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	UpdatedAt          time.Time                                  `json:"updated_at" api:"required" format:"date-time"`
	// Any of "research", "enrichment", "dataset_building".
	UseCase       TaskAgentUpdateResponseUseCase `json:"use_case" api:"required"`
	AccountID     string                         `json:"account_id" api:"nullable"`
	AgentName     string                         `json:"agent_name" api:"nullable"`
	WorkspaceID   string                         `json:"workspace_id" api:"nullable" format:"uuid"`
	WorkspaceName string                         `json:"workspace_name" api:"nullable"`
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
		IsActive           respjson.Field
		OutputSchema       respjson.Field
		Sources            respjson.Field
		SuggestedQuestions respjson.Field
		UpdatedAt          respjson.Field
		UseCase            respjson.Field
		AccountID          respjson.Field
		AgentName          respjson.Field
		WorkspaceID        respjson.Field
		WorkspaceName      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentUpdateResponseGoal struct {
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
func (r TaskAgentUpdateResponseGoal) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentUpdateResponseGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentUpdateResponseSources struct {
	Allow      []TaskAgentUpdateResponseSourcesAllow `json:"allow"`
	Avoid      string                                `json:"avoid" api:"nullable"`
	Block      []TaskAgentUpdateResponseSourcesBlock `json:"block"`
	Prioritize string                                `json:"prioritize" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allow       respjson.Field
		Avoid       respjson.Field
		Block       respjson.Field
		Prioritize  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentUpdateResponseSources) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentUpdateResponseSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentUpdateResponseSourcesAllow struct {
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
func (r TaskAgentUpdateResponseSourcesAllow) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentUpdateResponseSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentUpdateResponseSourcesBlock struct {
	Domains []string `json:"domains" api:"required"`
	Title   string   `json:"title" api:"required"`
	Order   int64    `json:"order"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domains     respjson.Field
		Title       respjson.Field
		Order       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentUpdateResponseSourcesBlock) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentUpdateResponseSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentUpdateResponseSuggestedQuestion struct {
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
func (r TaskAgentUpdateResponseSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentUpdateResponseSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentUpdateResponseUseCase string

const (
	TaskAgentUpdateResponseUseCaseResearch        TaskAgentUpdateResponseUseCase = "research"
	TaskAgentUpdateResponseUseCaseEnrichment      TaskAgentUpdateResponseUseCase = "enrichment"
	TaskAgentUpdateResponseUseCaseDatasetBuilding TaskAgentUpdateResponseUseCase = "dataset_building"
)

type TaskAgentListResponse struct {
	ID                 string                                   `json:"id" api:"required"`
	CreatedAt          time.Time                                `json:"created_at" api:"required" format:"date-time"`
	Description        string                                   `json:"description" api:"required"`
	DisplayName        string                                   `json:"display_name" api:"required"`
	DomainExpertise    string                                   `json:"domain_expertise" api:"required"`
	Effort             string                                   `json:"effort" api:"required"`
	Goals              []TaskAgentListResponseGoal              `json:"goals" api:"required"`
	Icon               string                                   `json:"icon" api:"required"`
	IsActive           bool                                     `json:"is_active" api:"required"`
	OutputSchema       map[string]any                           `json:"output_schema" api:"required"`
	Sources            TaskAgentListResponseSources             `json:"sources" api:"required"`
	SuggestedQuestions []TaskAgentListResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	UpdatedAt          time.Time                                `json:"updated_at" api:"required" format:"date-time"`
	// Any of "research", "enrichment", "dataset_building".
	UseCase       TaskAgentListResponseUseCase `json:"use_case" api:"required"`
	AccountID     string                       `json:"account_id" api:"nullable"`
	AgentName     string                       `json:"agent_name" api:"nullable"`
	WorkspaceID   string                       `json:"workspace_id" api:"nullable" format:"uuid"`
	WorkspaceName string                       `json:"workspace_name" api:"nullable"`
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
		IsActive           respjson.Field
		OutputSchema       respjson.Field
		Sources            respjson.Field
		SuggestedQuestions respjson.Field
		UpdatedAt          respjson.Field
		UseCase            respjson.Field
		AccountID          respjson.Field
		AgentName          respjson.Field
		WorkspaceID        respjson.Field
		WorkspaceName      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentListResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentListResponseGoal struct {
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
func (r TaskAgentListResponseGoal) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentListResponseGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentListResponseSources struct {
	Allow      []TaskAgentListResponseSourcesAllow `json:"allow"`
	Avoid      string                              `json:"avoid" api:"nullable"`
	Block      []TaskAgentListResponseSourcesBlock `json:"block"`
	Prioritize string                              `json:"prioritize" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allow       respjson.Field
		Avoid       respjson.Field
		Block       respjson.Field
		Prioritize  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentListResponseSources) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentListResponseSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentListResponseSourcesAllow struct {
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
func (r TaskAgentListResponseSourcesAllow) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentListResponseSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentListResponseSourcesBlock struct {
	Domains []string `json:"domains" api:"required"`
	Title   string   `json:"title" api:"required"`
	Order   int64    `json:"order"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domains     respjson.Field
		Title       respjson.Field
		Order       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentListResponseSourcesBlock) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentListResponseSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentListResponseSuggestedQuestion struct {
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
func (r TaskAgentListResponseSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentListResponseSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentListResponseUseCase string

const (
	TaskAgentListResponseUseCaseResearch        TaskAgentListResponseUseCase = "research"
	TaskAgentListResponseUseCaseEnrichment      TaskAgentListResponseUseCase = "enrichment"
	TaskAgentListResponseUseCaseDatasetBuilding TaskAgentListResponseUseCase = "dataset_building"
)

type TaskAgentGetResponse struct {
	ID                 string                                  `json:"id" api:"required"`
	CreatedAt          time.Time                               `json:"created_at" api:"required" format:"date-time"`
	Description        string                                  `json:"description" api:"required"`
	DisplayName        string                                  `json:"display_name" api:"required"`
	DomainExpertise    string                                  `json:"domain_expertise" api:"required"`
	Effort             string                                  `json:"effort" api:"required"`
	Goals              []TaskAgentGetResponseGoal              `json:"goals" api:"required"`
	Icon               string                                  `json:"icon" api:"required"`
	IsActive           bool                                    `json:"is_active" api:"required"`
	OutputSchema       map[string]any                          `json:"output_schema" api:"required"`
	Sources            TaskAgentGetResponseSources             `json:"sources" api:"required"`
	SuggestedQuestions []TaskAgentGetResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	UpdatedAt          time.Time                               `json:"updated_at" api:"required" format:"date-time"`
	// Any of "research", "enrichment", "dataset_building".
	UseCase       TaskAgentGetResponseUseCase `json:"use_case" api:"required"`
	AccountID     string                      `json:"account_id" api:"nullable"`
	AgentName     string                      `json:"agent_name" api:"nullable"`
	WorkspaceID   string                      `json:"workspace_id" api:"nullable" format:"uuid"`
	WorkspaceName string                      `json:"workspace_name" api:"nullable"`
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
		IsActive           respjson.Field
		OutputSchema       respjson.Field
		Sources            respjson.Field
		SuggestedQuestions respjson.Field
		UpdatedAt          respjson.Field
		UseCase            respjson.Field
		AccountID          respjson.Field
		AgentName          respjson.Field
		WorkspaceID        respjson.Field
		WorkspaceName      respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentGetResponseGoal struct {
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
func (r TaskAgentGetResponseGoal) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentGetResponseGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentGetResponseSources struct {
	Allow      []TaskAgentGetResponseSourcesAllow `json:"allow"`
	Avoid      string                             `json:"avoid" api:"nullable"`
	Block      []TaskAgentGetResponseSourcesBlock `json:"block"`
	Prioritize string                             `json:"prioritize" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allow       respjson.Field
		Avoid       respjson.Field
		Block       respjson.Field
		Prioritize  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentGetResponseSources) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentGetResponseSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentGetResponseSourcesAllow struct {
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
func (r TaskAgentGetResponseSourcesAllow) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentGetResponseSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentGetResponseSourcesBlock struct {
	Domains []string `json:"domains" api:"required"`
	Title   string   `json:"title" api:"required"`
	Order   int64    `json:"order"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domains     respjson.Field
		Title       respjson.Field
		Order       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentGetResponseSourcesBlock) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentGetResponseSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentGetResponseSuggestedQuestion struct {
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
func (r TaskAgentGetResponseSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentGetResponseSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentGetResponseUseCase string

const (
	TaskAgentGetResponseUseCaseResearch        TaskAgentGetResponseUseCase = "research"
	TaskAgentGetResponseUseCaseEnrichment      TaskAgentGetResponseUseCase = "enrichment"
	TaskAgentGetResponseUseCaseDatasetBuilding TaskAgentGetResponseUseCase = "dataset_building"
)

type TaskAgentRunResponse struct {
	// Run identifier.
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "quickest", "quick", "research", "pro", "max".
	Effort TaskAgentRunResponseEffort `json:"effort" api:"required"`
	// Interaction ID — pass as previous_interaction_id to reuse context.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status      TaskAgentRunResponseStatus `json:"status" api:"required"`
	CompletedAt time.Time                  `json:"completed_at" api:"nullable" format:"date-time"`
	Error       TaskAgentRunResponseError  `json:"error" api:"nullable"`
	Prompt      string                     `json:"prompt" api:"nullable"`
	StartedAt   time.Time                  `json:"started_at" api:"nullable" format:"date-time"`
	// Web Search Agent instance this run belongs to.
	WebSearchAgentID string `json:"web_search_agent_id" api:"nullable"`
	WorkspaceID      string `json:"workspace_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Effort           respjson.Field
		InteractionID    respjson.Field
		IsActive         respjson.Field
		Status           respjson.Field
		CompletedAt      respjson.Field
		Error            respjson.Field
		Prompt           respjson.Field
		StartedAt        respjson.Field
		WebSearchAgentID respjson.Field
		WorkspaceID      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunResponseEffort string

const (
	TaskAgentRunResponseEffortQuickest TaskAgentRunResponseEffort = "quickest"
	TaskAgentRunResponseEffortQuick    TaskAgentRunResponseEffort = "quick"
	TaskAgentRunResponseEffortResearch TaskAgentRunResponseEffort = "research"
	TaskAgentRunResponseEffortPro      TaskAgentRunResponseEffort = "pro"
	TaskAgentRunResponseEffortMax      TaskAgentRunResponseEffort = "max"
)

type TaskAgentRunResponseStatus string

const (
	TaskAgentRunResponseStatusQueued    TaskAgentRunResponseStatus = "queued"
	TaskAgentRunResponseStatusRunning   TaskAgentRunResponseStatus = "running"
	TaskAgentRunResponseStatusCompleted TaskAgentRunResponseStatus = "completed"
	TaskAgentRunResponseStatusFailed    TaskAgentRunResponseStatus = "failed"
	TaskAgentRunResponseStatusCancelled TaskAgentRunResponseStatus = "cancelled"
)

type TaskAgentRunResponseError struct {
	// Human-readable error description.
	Message string `json:"message" api:"required"`
	// Reference ID (equals the run id).
	RefID string `json:"ref_id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		RefID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunResponseError) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentNewParams struct {
	AgentName       param.Opt[string] `json:"agent_name,omitzero"`
	Description     param.Opt[string] `json:"description,omitzero"`
	DisplayName     param.Opt[string] `json:"display_name,omitzero"`
	DomainExpertise param.Opt[string] `json:"domain_expertise,omitzero"`
	Icon            param.Opt[string] `json:"icon,omitzero"`
	// Template name to materialise this instance from. When set, scalar fields and
	// child rows are copied from the template.
	Template     param.Opt[string] `json:"template,omitzero"`
	WorkspaceID  param.Opt[string] `json:"workspace_id,omitzero" format:"uuid"`
	Effort       param.Opt[string] `json:"effort,omitzero"`
	IsActive     param.Opt[bool]   `json:"is_active,omitzero"`
	OutputSchema map[string]any    `json:"output_schema,omitzero"`
	// Any of "research", "enrichment", "dataset_building".
	UseCase            TaskAgentNewParamsUseCase `json:"use_case,omitzero"`
	Goals              []string                  `json:"goals,omitzero"`
	Sources            TaskAgentNewParamsSources `json:"sources,omitzero"`
	SuggestedQuestions []string                  `json:"suggested_questions,omitzero"`
	paramObj
}

func (r TaskAgentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TaskAgentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskAgentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentNewParamsSources struct {
	Avoid      param.Opt[string]                `json:"avoid,omitzero"`
	Prioritize param.Opt[string]                `json:"prioritize,omitzero"`
	Allow      []TaskAgentNewParamsSourcesAllow `json:"allow,omitzero"`
	Block      []TaskAgentNewParamsSourcesBlock `json:"block,omitzero"`
	paramObj
}

func (r TaskAgentNewParamsSources) MarshalJSON() (data []byte, err error) {
	type shadow TaskAgentNewParamsSources
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskAgentNewParamsSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Domains, Title are required.
type TaskAgentNewParamsSourcesAllow struct {
	Domains []string         `json:"domains,omitzero" api:"required"`
	Title   string           `json:"title" api:"required"`
	Order   param.Opt[int64] `json:"order,omitzero"`
	paramObj
}

func (r TaskAgentNewParamsSourcesAllow) MarshalJSON() (data []byte, err error) {
	type shadow TaskAgentNewParamsSourcesAllow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskAgentNewParamsSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Domains, Title are required.
type TaskAgentNewParamsSourcesBlock struct {
	Domains []string         `json:"domains,omitzero" api:"required"`
	Title   string           `json:"title" api:"required"`
	Order   param.Opt[int64] `json:"order,omitzero"`
	paramObj
}

func (r TaskAgentNewParamsSourcesBlock) MarshalJSON() (data []byte, err error) {
	type shadow TaskAgentNewParamsSourcesBlock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskAgentNewParamsSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentNewParamsUseCase string

const (
	TaskAgentNewParamsUseCaseResearch        TaskAgentNewParamsUseCase = "research"
	TaskAgentNewParamsUseCaseEnrichment      TaskAgentNewParamsUseCase = "enrichment"
	TaskAgentNewParamsUseCaseDatasetBuilding TaskAgentNewParamsUseCase = "dataset_building"
)

type TaskAgentUpdateParams struct {
	Body []TaskAgentUpdateParamsBody
	paramObj
}

func (r TaskAgentUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *TaskAgentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Op, Path, Value are required.
type TaskAgentUpdateParamsBody struct {
	// Any of "replace".
	Op    string `json:"op,omitzero" api:"required"`
	Path  string `json:"path" api:"required"`
	Value any    `json:"value,omitzero" api:"required"`
	paramObj
}

func (r TaskAgentUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow TaskAgentUpdateParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskAgentUpdateParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TaskAgentUpdateParamsBody](
		"op", "replace",
	)
}

type TaskAgentListParams struct {
	Effort  param.Opt[string] `query:"effort,omitzero" json:"-"`
	UseCase param.Opt[string] `query:"use_case,omitzero" json:"-"`
	Limit   param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Offset  param.Opt[int64]  `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaskAgentListParams]'s query parameters as `url.Values`.
func (r TaskAgentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TaskAgentRunParams struct {
	Input        string                    `json:"input" api:"required"`
	EnableEvents param.Opt[bool]           `json:"enable_events,omitzero"`
	OutputSchema map[string]any            `json:"output_schema,omitzero"`
	Sources      TaskAgentRunParamsSources `json:"sources,omitzero"`
	paramObj
}

func (r TaskAgentRunParams) MarshalJSON() (data []byte, err error) {
	type shadow TaskAgentRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskAgentRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunParamsSources struct {
	Avoid      param.Opt[string]                `json:"avoid,omitzero"`
	Prioritize param.Opt[string]                `json:"prioritize,omitzero"`
	Allow      []TaskAgentRunParamsSourcesAllow `json:"allow,omitzero"`
	Block      []TaskAgentRunParamsSourcesBlock `json:"block,omitzero"`
	paramObj
}

func (r TaskAgentRunParamsSources) MarshalJSON() (data []byte, err error) {
	type shadow TaskAgentRunParamsSources
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskAgentRunParamsSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Domains, Title are required.
type TaskAgentRunParamsSourcesAllow struct {
	Domains []string         `json:"domains,omitzero" api:"required"`
	Title   string           `json:"title" api:"required"`
	Order   param.Opt[int64] `json:"order,omitzero"`
	paramObj
}

func (r TaskAgentRunParamsSourcesAllow) MarshalJSON() (data []byte, err error) {
	type shadow TaskAgentRunParamsSourcesAllow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskAgentRunParamsSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Domains, Title are required.
type TaskAgentRunParamsSourcesBlock struct {
	Domains []string         `json:"domains,omitzero" api:"required"`
	Title   string           `json:"title" api:"required"`
	Order   param.Opt[int64] `json:"order,omitzero"`
	paramObj
}

func (r TaskAgentRunParamsSourcesBlock) MarshalJSON() (data []byte, err error) {
	type shadow TaskAgentRunParamsSourcesBlock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskAgentRunParamsSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
