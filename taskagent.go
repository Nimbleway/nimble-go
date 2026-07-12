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

// Create a Web Search Agent instance.
//
// `account_id` is JWT-derived and never read from the request body.
func (r *TaskAgentService) New(ctx context.Context, body TaskAgentNewParams, opts ...option.RequestOption) (res *TaskAgentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/task-agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update Agent
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

// List Web Search Agent instances.
//
// Callers are strictly scoped to their (account, workspace). If `workspace_id` is
// omitted, the user's default workspace is used.
func (r *TaskAgentService) List(ctx context.Context, query TaskAgentListParams, opts ...option.RequestOption) (res *[]TaskAgentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/task-agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deactivate Agent
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

// Get Agent
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

// Create a research run for a Web Search Agent instance.
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
	ID              string    `json:"id" api:"required"`
	CreatedAt       time.Time `json:"created_at" api:"required" format:"date-time"`
	Description     string    `json:"description" api:"required"`
	DisplayName     string    `json:"display_name" api:"required"`
	DomainExpertise string    `json:"domain_expertise" api:"required"`
	// Canonical effort tier names for the research graph.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort       TaskAgentNewResponseEffort `json:"effort" api:"required"`
	Goals        []TaskAgentNewResponseGoal `json:"goals" api:"required"`
	Icon         string                     `json:"icon" api:"required"`
	IsActive     bool                       `json:"is_active" api:"required"`
	OutputSchema map[string]any             `json:"output_schema" api:"required"`
	// Response variant of AgentSources — preserves per-row id on allow rows.
	Sources            TaskAgentNewResponseSources             `json:"sources" api:"required"`
	SuggestedQuestions []TaskAgentNewResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	UpdatedAt          time.Time                               `json:"updated_at" api:"required" format:"date-time"`
	// Any of "research", "enrichment", "dataset_building".
	UseCase     TaskAgentNewResponseUseCase `json:"use_case" api:"required"`
	AccountID   string                      `json:"account_id" api:"nullable"`
	AgentName   string                      `json:"agent_name" api:"nullable"`
	WorkspaceID string                      `json:"workspace_id" api:"nullable" format:"uuid"`
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
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Canonical effort tier names for the research graph.
type TaskAgentNewResponseEffort string

const (
	TaskAgentNewResponseEffortLow    TaskAgentNewResponseEffort = "low"
	TaskAgentNewResponseEffortMedium TaskAgentNewResponseEffort = "medium"
	TaskAgentNewResponseEffortHigh   TaskAgentNewResponseEffort = "high"
	TaskAgentNewResponseEffortXHigh  TaskAgentNewResponseEffort = "x-high"
	TaskAgentNewResponseEffortMax    TaskAgentNewResponseEffort = "max"
)

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

// Response variant of AgentSources — preserves per-row id on allow rows.
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

// Lenient response shape — domains are plain strings (no re-validation).
type TaskAgentNewResponseSourcesBlock struct {
	Domains []string `json:"domains" api:"required"`
	Order   int64    `json:"order" api:"required"`
	Title   string   `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domains     respjson.Field
		Order       respjson.Field
		Title       respjson.Field
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
	ID              string    `json:"id" api:"required"`
	CreatedAt       time.Time `json:"created_at" api:"required" format:"date-time"`
	Description     string    `json:"description" api:"required"`
	DisplayName     string    `json:"display_name" api:"required"`
	DomainExpertise string    `json:"domain_expertise" api:"required"`
	// Canonical effort tier names for the research graph.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort       TaskAgentUpdateResponseEffort `json:"effort" api:"required"`
	Goals        []TaskAgentUpdateResponseGoal `json:"goals" api:"required"`
	Icon         string                        `json:"icon" api:"required"`
	IsActive     bool                          `json:"is_active" api:"required"`
	OutputSchema map[string]any                `json:"output_schema" api:"required"`
	// Response variant of AgentSources — preserves per-row id on allow rows.
	Sources            TaskAgentUpdateResponseSources             `json:"sources" api:"required"`
	SuggestedQuestions []TaskAgentUpdateResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	UpdatedAt          time.Time                                  `json:"updated_at" api:"required" format:"date-time"`
	// Any of "research", "enrichment", "dataset_building".
	UseCase     TaskAgentUpdateResponseUseCase `json:"use_case" api:"required"`
	AccountID   string                         `json:"account_id" api:"nullable"`
	AgentName   string                         `json:"agent_name" api:"nullable"`
	WorkspaceID string                         `json:"workspace_id" api:"nullable" format:"uuid"`
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
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Canonical effort tier names for the research graph.
type TaskAgentUpdateResponseEffort string

const (
	TaskAgentUpdateResponseEffortLow    TaskAgentUpdateResponseEffort = "low"
	TaskAgentUpdateResponseEffortMedium TaskAgentUpdateResponseEffort = "medium"
	TaskAgentUpdateResponseEffortHigh   TaskAgentUpdateResponseEffort = "high"
	TaskAgentUpdateResponseEffortXHigh  TaskAgentUpdateResponseEffort = "x-high"
	TaskAgentUpdateResponseEffortMax    TaskAgentUpdateResponseEffort = "max"
)

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

// Response variant of AgentSources — preserves per-row id on allow rows.
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

// Lenient response shape — domains are plain strings (no re-validation).
type TaskAgentUpdateResponseSourcesBlock struct {
	Domains []string `json:"domains" api:"required"`
	Order   int64    `json:"order" api:"required"`
	Title   string   `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domains     respjson.Field
		Order       respjson.Field
		Title       respjson.Field
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
	ID              string    `json:"id" api:"required"`
	CreatedAt       time.Time `json:"created_at" api:"required" format:"date-time"`
	Description     string    `json:"description" api:"required"`
	DisplayName     string    `json:"display_name" api:"required"`
	DomainExpertise string    `json:"domain_expertise" api:"required"`
	// Canonical effort tier names for the research graph.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort       TaskAgentListResponseEffort `json:"effort" api:"required"`
	Goals        []TaskAgentListResponseGoal `json:"goals" api:"required"`
	Icon         string                      `json:"icon" api:"required"`
	IsActive     bool                        `json:"is_active" api:"required"`
	OutputSchema map[string]any              `json:"output_schema" api:"required"`
	// Response variant of AgentSources — preserves per-row id on allow rows.
	Sources            TaskAgentListResponseSources             `json:"sources" api:"required"`
	SuggestedQuestions []TaskAgentListResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	UpdatedAt          time.Time                                `json:"updated_at" api:"required" format:"date-time"`
	// Any of "research", "enrichment", "dataset_building".
	UseCase     TaskAgentListResponseUseCase `json:"use_case" api:"required"`
	AccountID   string                       `json:"account_id" api:"nullable"`
	AgentName   string                       `json:"agent_name" api:"nullable"`
	WorkspaceID string                       `json:"workspace_id" api:"nullable" format:"uuid"`
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
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentListResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Canonical effort tier names for the research graph.
type TaskAgentListResponseEffort string

const (
	TaskAgentListResponseEffortLow    TaskAgentListResponseEffort = "low"
	TaskAgentListResponseEffortMedium TaskAgentListResponseEffort = "medium"
	TaskAgentListResponseEffortHigh   TaskAgentListResponseEffort = "high"
	TaskAgentListResponseEffortXHigh  TaskAgentListResponseEffort = "x-high"
	TaskAgentListResponseEffortMax    TaskAgentListResponseEffort = "max"
)

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

// Response variant of AgentSources — preserves per-row id on allow rows.
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

// Lenient response shape — domains are plain strings (no re-validation).
type TaskAgentListResponseSourcesBlock struct {
	Domains []string `json:"domains" api:"required"`
	Order   int64    `json:"order" api:"required"`
	Title   string   `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domains     respjson.Field
		Order       respjson.Field
		Title       respjson.Field
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
	ID              string    `json:"id" api:"required"`
	CreatedAt       time.Time `json:"created_at" api:"required" format:"date-time"`
	Description     string    `json:"description" api:"required"`
	DisplayName     string    `json:"display_name" api:"required"`
	DomainExpertise string    `json:"domain_expertise" api:"required"`
	// Canonical effort tier names for the research graph.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort       TaskAgentGetResponseEffort `json:"effort" api:"required"`
	Goals        []TaskAgentGetResponseGoal `json:"goals" api:"required"`
	Icon         string                     `json:"icon" api:"required"`
	IsActive     bool                       `json:"is_active" api:"required"`
	OutputSchema map[string]any             `json:"output_schema" api:"required"`
	// Response variant of AgentSources — preserves per-row id on allow rows.
	Sources            TaskAgentGetResponseSources             `json:"sources" api:"required"`
	SuggestedQuestions []TaskAgentGetResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	UpdatedAt          time.Time                               `json:"updated_at" api:"required" format:"date-time"`
	// Any of "research", "enrichment", "dataset_building".
	UseCase     TaskAgentGetResponseUseCase `json:"use_case" api:"required"`
	AccountID   string                      `json:"account_id" api:"nullable"`
	AgentName   string                      `json:"agent_name" api:"nullable"`
	WorkspaceID string                      `json:"workspace_id" api:"nullable" format:"uuid"`
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
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Canonical effort tier names for the research graph.
type TaskAgentGetResponseEffort string

const (
	TaskAgentGetResponseEffortLow    TaskAgentGetResponseEffort = "low"
	TaskAgentGetResponseEffortMedium TaskAgentGetResponseEffort = "medium"
	TaskAgentGetResponseEffortHigh   TaskAgentGetResponseEffort = "high"
	TaskAgentGetResponseEffortXHigh  TaskAgentGetResponseEffort = "x-high"
	TaskAgentGetResponseEffortMax    TaskAgentGetResponseEffort = "max"
)

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

// Response variant of AgentSources — preserves per-row id on allow rows.
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

// Lenient response shape — domains are plain strings (no re-validation).
type TaskAgentGetResponseSourcesBlock struct {
	Domains []string `json:"domains" api:"required"`
	Order   int64    `json:"order" api:"required"`
	Title   string   `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domains     respjson.Field
		Order       respjson.Field
		Title       respjson.Field
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

// Task run status returned by list/create/get endpoints.
type TaskAgentRunResponse struct {
	// Run identifier, format "task*run*{uuid}".
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Canonical effort tier names for the research graph.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort TaskAgentRunResponseEffort `json:"effort" api:"required"`
	// Interaction ID — pass as previous_interaction_id to reuse context.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Lowercase status values used in API responses (distinct from the DB-level
	// TaskRunStatus enum).
	//
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status TaskAgentRunResponseStatus `json:"status" api:"required"`
	// Web Search Agent instance this run belongs to. Every task run is agent-bound
	// (see AGENTS-1666). Use this to build the nested URL
	// /api/v2/web-search-agents/{web_search_agent_id}/runs/{id}.
	WebSearchAgentID string    `json:"web_search_agent_id" api:"required"`
	CompletedAt      time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Error detail for a failed run.
	Error TaskAgentRunResponseError `json:"error" api:"nullable"`
	// Original user prompt before enrichment. Populated for Web Search Agent runs.
	Prompt      string    `json:"prompt" api:"nullable"`
	StartedAt   time.Time `json:"started_at" api:"nullable" format:"date-time"`
	WorkspaceID string    `json:"workspace_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CreatedAt        respjson.Field
		Effort           respjson.Field
		InteractionID    respjson.Field
		IsActive         respjson.Field
		Status           respjson.Field
		WebSearchAgentID respjson.Field
		CompletedAt      respjson.Field
		Error            respjson.Field
		Prompt           respjson.Field
		StartedAt        respjson.Field
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

// Canonical effort tier names for the research graph.
type TaskAgentRunResponseEffort string

const (
	TaskAgentRunResponseEffortLow    TaskAgentRunResponseEffort = "low"
	TaskAgentRunResponseEffortMedium TaskAgentRunResponseEffort = "medium"
	TaskAgentRunResponseEffortHigh   TaskAgentRunResponseEffort = "high"
	TaskAgentRunResponseEffortXHigh  TaskAgentRunResponseEffort = "x-high"
	TaskAgentRunResponseEffortMax    TaskAgentRunResponseEffort = "max"
)

// Lowercase status values used in API responses (distinct from the DB-level
// TaskRunStatus enum).
type TaskAgentRunResponseStatus string

const (
	TaskAgentRunResponseStatusQueued    TaskAgentRunResponseStatus = "queued"
	TaskAgentRunResponseStatusRunning   TaskAgentRunResponseStatus = "running"
	TaskAgentRunResponseStatusCompleted TaskAgentRunResponseStatus = "completed"
	TaskAgentRunResponseStatusFailed    TaskAgentRunResponseStatus = "failed"
	TaskAgentRunResponseStatusCancelled TaskAgentRunResponseStatus = "cancelled"
)

// Error detail for a failed run.
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
	// Template name to materialize this instance from. When set, the scalar fields and
	// child rows are copied from the template.
	Template     param.Opt[string] `json:"template,omitzero"`
	WorkspaceID  param.Opt[string] `json:"workspace_id,omitzero" format:"uuid"`
	IsActive     param.Opt[bool]   `json:"is_active,omitzero"`
	OutputSchema map[string]any    `json:"output_schema,omitzero"`
	// Any of "research", "enrichment", "dataset_building".
	UseCase TaskAgentNewParamsUseCase `json:"use_case,omitzero"`
	// Canonical effort tier names for the research graph.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort TaskAgentNewParamsEffort `json:"effort,omitzero"`
	Goals  []string                 `json:"goals,omitzero"`
	// Source preferences for a web search agent instance.
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

// Canonical effort tier names for the research graph.
type TaskAgentNewParamsEffort string

const (
	TaskAgentNewParamsEffortLow    TaskAgentNewParamsEffort = "low"
	TaskAgentNewParamsEffortMedium TaskAgentNewParamsEffort = "medium"
	TaskAgentNewParamsEffortHigh   TaskAgentNewParamsEffort = "high"
	TaskAgentNewParamsEffortXHigh  TaskAgentNewParamsEffort = "x-high"
	TaskAgentNewParamsEffortMax    TaskAgentNewParamsEffort = "max"
)

// Source preferences for a web search agent instance.
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
	// A JSON Patch document per RFC 6902 — a JSON array of patch operations.
	Body []TaskAgentUpdateParamsBody
	paramObj
}

func (r TaskAgentUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *TaskAgentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single JSON Patch operation per RFC 6902.
//
// The properties Op, Path are required.
type TaskAgentUpdateParamsBody struct {
	// Any of "add", "remove", "replace", "move", "copy", "test".
	Op    string            `json:"op,omitzero" api:"required"`
	Path  string            `json:"path" api:"required"`
	From  param.Opt[string] `json:"from,omitzero"`
	Value any               `json:"value,omitzero"`
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
		"op", "add", "remove", "replace", "move", "copy", "test",
	)
}

type TaskAgentListParams struct {
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" format:"uuid" json:"-"`
	Limit       param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Offset      param.Opt[int64]  `query:"offset,omitzero" json:"-"`
	// Canonical effort tier names for the research graph.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	FilterEffort TaskAgentListParamsFilterEffort `query:"filter_effort,omitzero" json:"-"`
	// Any of "research", "enrichment", "dataset_building".
	FilterUseCase TaskAgentListParamsFilterUseCase `query:"filter_use_case,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaskAgentListParams]'s query parameters as `url.Values`.
func (r TaskAgentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Canonical effort tier names for the research graph.
type TaskAgentListParamsFilterEffort string

const (
	TaskAgentListParamsFilterEffortLow    TaskAgentListParamsFilterEffort = "low"
	TaskAgentListParamsFilterEffortMedium TaskAgentListParamsFilterEffort = "medium"
	TaskAgentListParamsFilterEffortHigh   TaskAgentListParamsFilterEffort = "high"
	TaskAgentListParamsFilterEffortXHigh  TaskAgentListParamsFilterEffort = "x-high"
	TaskAgentListParamsFilterEffortMax    TaskAgentListParamsFilterEffort = "max"
)

type TaskAgentListParamsFilterUseCase string

const (
	TaskAgentListParamsFilterUseCaseResearch        TaskAgentListParamsFilterUseCase = "research"
	TaskAgentListParamsFilterUseCaseEnrichment      TaskAgentListParamsFilterUseCase = "enrichment"
	TaskAgentListParamsFilterUseCaseDatasetBuilding TaskAgentListParamsFilterUseCase = "dataset_building"
)

type TaskAgentRunParams struct {
	Input                 string            `json:"input" api:"required"`
	PreviousInteractionID param.Opt[string] `json:"previous_interaction_id,omitzero"`
	EnableEvents          param.Opt[bool]   `json:"enable_events,omitzero"`
	// Canonical effort tier names for the research graph.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort       TaskAgentRunParamsEffort `json:"effort,omitzero"`
	OutputSchema map[string]any           `json:"output_schema,omitzero"`
	// Source preferences for a web search agent instance.
	Sources TaskAgentRunParamsSources `json:"sources,omitzero"`
	paramObj
}

func (r TaskAgentRunParams) MarshalJSON() (data []byte, err error) {
	type shadow TaskAgentRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskAgentRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Canonical effort tier names for the research graph.
type TaskAgentRunParamsEffort string

const (
	TaskAgentRunParamsEffortLow    TaskAgentRunParamsEffort = "low"
	TaskAgentRunParamsEffortMedium TaskAgentRunParamsEffort = "medium"
	TaskAgentRunParamsEffortHigh   TaskAgentRunParamsEffort = "high"
	TaskAgentRunParamsEffortXHigh  TaskAgentRunParamsEffort = "x-high"
	TaskAgentRunParamsEffortMax    TaskAgentRunParamsEffort = "max"
)

// Source preferences for a web search agent instance.
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
