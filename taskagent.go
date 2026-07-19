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
//
// Deprecated: deprecated
func (r *TaskAgentService) New(ctx context.Context, body TaskAgentNewParams, opts ...option.RequestOption) (res *TaskAgentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/task-agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update Agent
//
// Deprecated: deprecated
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
//
// Deprecated: deprecated
func (r *TaskAgentService) List(ctx context.Context, query TaskAgentListParams, opts ...option.RequestOption) (res *TaskAgentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/task-agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deactivate Agent
//
// Deprecated: deprecated
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
//
// Deprecated: deprecated
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
//
// Deprecated: deprecated
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
	// Unique web search agent identifier (wsa\_<uuid>).
	ID string `json:"id" api:"required"`
	// When the agent was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Agent description shown to users.
	Description string `json:"description" api:"required"`
	// Human-friendly agent name shown to users.
	DisplayName string `json:"display_name" api:"required"`
	// Domain expertise or operating context for the agent.
	DomainExpertise string `json:"domain_expertise" api:"required"`
	// Default effort level for this agent's runs.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort TaskAgentNewResponseEffort `json:"effort" api:"required"`
	// Ordered goals for the agent to follow.
	Goals []TaskAgentNewResponseGoal `json:"goals" api:"required"`
	// Icon identifier used when presenting the agent.
	Icon string `json:"icon" api:"required"`
	// Whether the agent can be used to start new runs.
	IsActive bool `json:"is_active" api:"required"`
	// JSON schema describing the structured output the agent should produce.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Source guidance for the agent.
	Sources TaskAgentNewResponseSources `json:"sources" api:"required"`
	// Suggested prompts users can run with this agent.
	SuggestedQuestions []TaskAgentNewResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	// When the agent was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Primary use case supported by the agent.
	//
	// Any of "research", "enrichment", "dataset_building".
	UseCase TaskAgentNewResponseUseCase `json:"use_case" api:"required"`
	// Stable agent name.
	AgentName string `json:"agent_name" api:"nullable"`
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
		AgentName          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default effort level for this agent's runs.
type TaskAgentNewResponseEffort string

const (
	TaskAgentNewResponseEffortLow    TaskAgentNewResponseEffort = "low"
	TaskAgentNewResponseEffortMedium TaskAgentNewResponseEffort = "medium"
	TaskAgentNewResponseEffortHigh   TaskAgentNewResponseEffort = "high"
	TaskAgentNewResponseEffortXHigh  TaskAgentNewResponseEffort = "x-high"
	TaskAgentNewResponseEffortMax    TaskAgentNewResponseEffort = "max"
)

type TaskAgentNewResponseGoal struct {
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
func (r TaskAgentNewResponseGoal) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentNewResponseGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source guidance for the agent.
type TaskAgentNewResponseSources struct {
	// Source groups the agent is allowed to use.
	Allow []TaskAgentNewResponseSourcesAllow `json:"allow"`
	// Free-text guidance describing sources or domains to avoid.
	Avoid string `json:"avoid" api:"nullable"`
	// Source groups the agent should not use.
	Block []TaskAgentNewResponseSourcesBlock `json:"block"`
	// Free-text guidance describing sources or domains to prioritize.
	Prioritize string `json:"prioritize" api:"nullable"`
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
func (r TaskAgentNewResponseSourcesAllow) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentNewResponseSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentNewResponseSourcesBlock struct {
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
func (r TaskAgentNewResponseSourcesBlock) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentNewResponseSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentNewResponseSuggestedQuestion struct {
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
func (r TaskAgentNewResponseSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentNewResponseSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary use case supported by the agent.
type TaskAgentNewResponseUseCase string

const (
	TaskAgentNewResponseUseCaseResearch        TaskAgentNewResponseUseCase = "research"
	TaskAgentNewResponseUseCaseEnrichment      TaskAgentNewResponseUseCase = "enrichment"
	TaskAgentNewResponseUseCaseDatasetBuilding TaskAgentNewResponseUseCase = "dataset_building"
)

type TaskAgentUpdateResponse struct {
	// Unique web search agent identifier (wsa\_<uuid>).
	ID string `json:"id" api:"required"`
	// When the agent was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Agent description shown to users.
	Description string `json:"description" api:"required"`
	// Human-friendly agent name shown to users.
	DisplayName string `json:"display_name" api:"required"`
	// Domain expertise or operating context for the agent.
	DomainExpertise string `json:"domain_expertise" api:"required"`
	// Default effort level for this agent's runs.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort TaskAgentUpdateResponseEffort `json:"effort" api:"required"`
	// Ordered goals for the agent to follow.
	Goals []TaskAgentUpdateResponseGoal `json:"goals" api:"required"`
	// Icon identifier used when presenting the agent.
	Icon string `json:"icon" api:"required"`
	// Whether the agent can be used to start new runs.
	IsActive bool `json:"is_active" api:"required"`
	// JSON schema describing the structured output the agent should produce.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Source guidance for the agent.
	Sources TaskAgentUpdateResponseSources `json:"sources" api:"required"`
	// Suggested prompts users can run with this agent.
	SuggestedQuestions []TaskAgentUpdateResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	// When the agent was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Primary use case supported by the agent.
	//
	// Any of "research", "enrichment", "dataset_building".
	UseCase TaskAgentUpdateResponseUseCase `json:"use_case" api:"required"`
	// Stable agent name.
	AgentName string `json:"agent_name" api:"nullable"`
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
		AgentName          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default effort level for this agent's runs.
type TaskAgentUpdateResponseEffort string

const (
	TaskAgentUpdateResponseEffortLow    TaskAgentUpdateResponseEffort = "low"
	TaskAgentUpdateResponseEffortMedium TaskAgentUpdateResponseEffort = "medium"
	TaskAgentUpdateResponseEffortHigh   TaskAgentUpdateResponseEffort = "high"
	TaskAgentUpdateResponseEffortXHigh  TaskAgentUpdateResponseEffort = "x-high"
	TaskAgentUpdateResponseEffortMax    TaskAgentUpdateResponseEffort = "max"
)

type TaskAgentUpdateResponseGoal struct {
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
func (r TaskAgentUpdateResponseGoal) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentUpdateResponseGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source guidance for the agent.
type TaskAgentUpdateResponseSources struct {
	// Source groups the agent is allowed to use.
	Allow []TaskAgentUpdateResponseSourcesAllow `json:"allow"`
	// Free-text guidance describing sources or domains to avoid.
	Avoid string `json:"avoid" api:"nullable"`
	// Source groups the agent should not use.
	Block []TaskAgentUpdateResponseSourcesBlock `json:"block"`
	// Free-text guidance describing sources or domains to prioritize.
	Prioritize string `json:"prioritize" api:"nullable"`
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
func (r TaskAgentUpdateResponseSourcesAllow) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentUpdateResponseSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentUpdateResponseSourcesBlock struct {
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
func (r TaskAgentUpdateResponseSourcesBlock) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentUpdateResponseSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentUpdateResponseSuggestedQuestion struct {
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
func (r TaskAgentUpdateResponseSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentUpdateResponseSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary use case supported by the agent.
type TaskAgentUpdateResponseUseCase string

const (
	TaskAgentUpdateResponseUseCaseResearch        TaskAgentUpdateResponseUseCase = "research"
	TaskAgentUpdateResponseUseCaseEnrichment      TaskAgentUpdateResponseUseCase = "enrichment"
	TaskAgentUpdateResponseUseCaseDatasetBuilding TaskAgentUpdateResponseUseCase = "dataset_building"
)

type TaskAgentListResponse struct {
	// Items returned in this page.
	Items []TaskAgentListResponseItem `json:"items" api:"required"`
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
func (r TaskAgentListResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentListResponseItem struct {
	// Unique web search agent identifier (wsa\_<uuid>).
	ID string `json:"id" api:"required"`
	// When the agent was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Agent description shown to users.
	Description string `json:"description" api:"required"`
	// Human-friendly agent name shown to users.
	DisplayName string `json:"display_name" api:"required"`
	// Domain expertise or operating context for the agent.
	DomainExpertise string `json:"domain_expertise" api:"required"`
	// Default effort level for this agent's runs.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort string `json:"effort" api:"required"`
	// Ordered goals for the agent to follow.
	Goals []TaskAgentListResponseItemGoal `json:"goals" api:"required"`
	// Icon identifier used when presenting the agent.
	Icon string `json:"icon" api:"required"`
	// Whether the agent can be used to start new runs.
	IsActive bool `json:"is_active" api:"required"`
	// JSON schema describing the structured output the agent should produce.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Source guidance for the agent.
	Sources TaskAgentListResponseItemSources `json:"sources" api:"required"`
	// Suggested prompts users can run with this agent.
	SuggestedQuestions []TaskAgentListResponseItemSuggestedQuestion `json:"suggested_questions" api:"required"`
	// When the agent was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Primary use case supported by the agent.
	//
	// Any of "research", "enrichment", "dataset_building".
	UseCase string `json:"use_case" api:"required"`
	// Stable agent name.
	AgentName string `json:"agent_name" api:"nullable"`
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
		AgentName          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentListResponseItemGoal struct {
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
func (r TaskAgentListResponseItemGoal) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentListResponseItemGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source guidance for the agent.
type TaskAgentListResponseItemSources struct {
	// Source groups the agent is allowed to use.
	Allow []TaskAgentListResponseItemSourcesAllow `json:"allow"`
	// Free-text guidance describing sources or domains to avoid.
	Avoid string `json:"avoid" api:"nullable"`
	// Source groups the agent should not use.
	Block []TaskAgentListResponseItemSourcesBlock `json:"block"`
	// Free-text guidance describing sources or domains to prioritize.
	Prioritize string `json:"prioritize" api:"nullable"`
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
func (r TaskAgentListResponseItemSources) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentListResponseItemSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentListResponseItemSourcesAllow struct {
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
func (r TaskAgentListResponseItemSourcesAllow) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentListResponseItemSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentListResponseItemSourcesBlock struct {
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
func (r TaskAgentListResponseItemSourcesBlock) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentListResponseItemSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentListResponseItemSuggestedQuestion struct {
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
func (r TaskAgentListResponseItemSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentListResponseItemSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentGetResponse struct {
	// Unique web search agent identifier (wsa\_<uuid>).
	ID string `json:"id" api:"required"`
	// When the agent was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Agent description shown to users.
	Description string `json:"description" api:"required"`
	// Human-friendly agent name shown to users.
	DisplayName string `json:"display_name" api:"required"`
	// Domain expertise or operating context for the agent.
	DomainExpertise string `json:"domain_expertise" api:"required"`
	// Default effort level for this agent's runs.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort TaskAgentGetResponseEffort `json:"effort" api:"required"`
	// Ordered goals for the agent to follow.
	Goals []TaskAgentGetResponseGoal `json:"goals" api:"required"`
	// Icon identifier used when presenting the agent.
	Icon string `json:"icon" api:"required"`
	// Whether the agent can be used to start new runs.
	IsActive bool `json:"is_active" api:"required"`
	// JSON schema describing the structured output the agent should produce.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Source guidance for the agent.
	Sources TaskAgentGetResponseSources `json:"sources" api:"required"`
	// Suggested prompts users can run with this agent.
	SuggestedQuestions []TaskAgentGetResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	// When the agent was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Primary use case supported by the agent.
	//
	// Any of "research", "enrichment", "dataset_building".
	UseCase TaskAgentGetResponseUseCase `json:"use_case" api:"required"`
	// Stable agent name.
	AgentName string `json:"agent_name" api:"nullable"`
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
		AgentName          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default effort level for this agent's runs.
type TaskAgentGetResponseEffort string

const (
	TaskAgentGetResponseEffortLow    TaskAgentGetResponseEffort = "low"
	TaskAgentGetResponseEffortMedium TaskAgentGetResponseEffort = "medium"
	TaskAgentGetResponseEffortHigh   TaskAgentGetResponseEffort = "high"
	TaskAgentGetResponseEffortXHigh  TaskAgentGetResponseEffort = "x-high"
	TaskAgentGetResponseEffortMax    TaskAgentGetResponseEffort = "max"
)

type TaskAgentGetResponseGoal struct {
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
func (r TaskAgentGetResponseGoal) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentGetResponseGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source guidance for the agent.
type TaskAgentGetResponseSources struct {
	// Source groups the agent is allowed to use.
	Allow []TaskAgentGetResponseSourcesAllow `json:"allow"`
	// Free-text guidance describing sources or domains to avoid.
	Avoid string `json:"avoid" api:"nullable"`
	// Source groups the agent should not use.
	Block []TaskAgentGetResponseSourcesBlock `json:"block"`
	// Free-text guidance describing sources or domains to prioritize.
	Prioritize string `json:"prioritize" api:"nullable"`
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
func (r TaskAgentGetResponseSourcesAllow) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentGetResponseSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentGetResponseSourcesBlock struct {
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
func (r TaskAgentGetResponseSourcesBlock) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentGetResponseSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentGetResponseSuggestedQuestion struct {
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
func (r TaskAgentGetResponseSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentGetResponseSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary use case supported by the agent.
type TaskAgentGetResponseUseCase string

const (
	TaskAgentGetResponseUseCaseResearch        TaskAgentGetResponseUseCase = "research"
	TaskAgentGetResponseUseCaseEnrichment      TaskAgentGetResponseUseCase = "enrichment"
	TaskAgentGetResponseUseCaseDatasetBuilding TaskAgentGetResponseUseCase = "dataset_building"
)

type TaskAgentRunResponse struct {
	// Run identifier, format "task*run*{uuid}".
	ID string `json:"id" api:"required"`
	// When the run was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Effort level used for the run.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort TaskAgentRunResponseEffort `json:"effort" api:"required"`
	// Interaction ID.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Current run status.
	//
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status TaskAgentRunResponseStatus `json:"status" api:"required"`
	// Web Search Agent instance this run belongs to.
	WebSearchAgentID string `json:"web_search_agent_id" api:"required"`
	// When the run completed.
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Error details when the run failed.
	Error TaskAgentRunResponseError `json:"error" api:"nullable"`
	// Prompt submitted for the run.
	Prompt string `json:"prompt" api:"nullable"`
	// When the run started executing.
	StartedAt time.Time `json:"started_at" api:"nullable" format:"date-time"`
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
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Effort level used for the run.
type TaskAgentRunResponseEffort string

const (
	TaskAgentRunResponseEffortLow    TaskAgentRunResponseEffort = "low"
	TaskAgentRunResponseEffortMedium TaskAgentRunResponseEffort = "medium"
	TaskAgentRunResponseEffortHigh   TaskAgentRunResponseEffort = "high"
	TaskAgentRunResponseEffortXHigh  TaskAgentRunResponseEffort = "x-high"
	TaskAgentRunResponseEffortMax    TaskAgentRunResponseEffort = "max"
)

// Current run status.
type TaskAgentRunResponseStatus string

const (
	TaskAgentRunResponseStatusQueued    TaskAgentRunResponseStatus = "queued"
	TaskAgentRunResponseStatusRunning   TaskAgentRunResponseStatus = "running"
	TaskAgentRunResponseStatusCompleted TaskAgentRunResponseStatus = "completed"
	TaskAgentRunResponseStatusFailed    TaskAgentRunResponseStatus = "failed"
	TaskAgentRunResponseStatusCancelled TaskAgentRunResponseStatus = "cancelled"
)

// Error details when the run failed.
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
	// Stable agent name.
	AgentName param.Opt[string] `json:"agent_name,omitzero"`
	// Agent description shown to users.
	Description param.Opt[string] `json:"description,omitzero"`
	// Human-friendly agent name shown to users.
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// Domain expertise or operating context for the agent.
	DomainExpertise param.Opt[string] `json:"domain_expertise,omitzero"`
	// Icon identifier used when presenting the agent.
	Icon param.Opt[string] `json:"icon,omitzero"`
	// Template name to materialize this instance from. When set, the scalar fields and
	// child rows are copied from the template.
	Template param.Opt[string] `json:"template,omitzero"`
	// Whether the agent can be used to start new runs.
	IsActive param.Opt[bool] `json:"is_active,omitzero"`
	// JSON schema describing the structured output the agent should produce.
	OutputSchema map[string]any `json:"output_schema,omitzero"`
	// Primary use case supported by the agent.
	//
	// Any of "research", "enrichment", "dataset_building".
	UseCase TaskAgentNewParamsUseCase `json:"use_case,omitzero"`
	// Default effort level for this agent's runs.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort TaskAgentNewParamsEffort `json:"effort,omitzero"`
	// Ordered goals for the agent to follow.
	Goals []string `json:"goals,omitzero"`
	// Source guidance for the agent.
	Sources TaskAgentNewParamsSources `json:"sources,omitzero"`
	// Suggested prompts users can run with this agent.
	SuggestedQuestions []string `json:"suggested_questions,omitzero"`
	paramObj
}

func (r TaskAgentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TaskAgentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskAgentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default effort level for this agent's runs.
type TaskAgentNewParamsEffort string

const (
	TaskAgentNewParamsEffortLow    TaskAgentNewParamsEffort = "low"
	TaskAgentNewParamsEffortMedium TaskAgentNewParamsEffort = "medium"
	TaskAgentNewParamsEffortHigh   TaskAgentNewParamsEffort = "high"
	TaskAgentNewParamsEffortXHigh  TaskAgentNewParamsEffort = "x-high"
	TaskAgentNewParamsEffortMax    TaskAgentNewParamsEffort = "max"
)

// Source guidance for the agent.
type TaskAgentNewParamsSources struct {
	// Free-text guidance describing sources or domains to avoid.
	Avoid param.Opt[string] `json:"avoid,omitzero"`
	// Free-text guidance describing sources or domains to prioritize.
	Prioritize param.Opt[string] `json:"prioritize,omitzero"`
	// Source groups the agent is allowed to use.
	Allow []TaskAgentNewParamsSourcesAllow `json:"allow,omitzero"`
	// Source groups the agent should not use.
	Block []TaskAgentNewParamsSourcesBlock `json:"block,omitzero"`
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
	// Domains included in this source group.
	Domains []string `json:"domains,omitzero" api:"required"`
	// Source group title.
	Title string `json:"title" api:"required"`
	// Zero-based source group position.
	Order param.Opt[int64] `json:"order,omitzero"`
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
	// Domains included in this source group.
	Domains []string `json:"domains,omitzero" api:"required"`
	// Source group title.
	Title string `json:"title" api:"required"`
	// Zero-based source group position.
	Order param.Opt[int64] `json:"order,omitzero"`
	paramObj
}

func (r TaskAgentNewParamsSourcesBlock) MarshalJSON() (data []byte, err error) {
	type shadow TaskAgentNewParamsSourcesBlock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskAgentNewParamsSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary use case supported by the agent.
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
	// User prompt or task instructions for the run.
	Input string `json:"input" api:"required"`
	// Previous interaction identifier used to continue a conversation.
	PreviousInteractionID param.Opt[string] `json:"previous_interaction_id,omitzero"`
	// Whether to stream run events when supported.
	EnableEvents param.Opt[bool] `json:"enable_events,omitzero"`
	// Canonical effort tier names for the research graph.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort TaskAgentRunParamsEffort `json:"effort,omitzero"`
	// Existing records to ENRICH: a list of partial rows, or a single object,
	// mirroring output_schema's shape.
	InputData TaskAgentRunParamsInputDataUnion `json:"input_data,omitzero"`
	// JSON schema overriding the agent's default structured output for this run.
	OutputSchema map[string]any `json:"output_schema,omitzero"`
	// Source guidance overriding the agent default.
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

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type TaskAgentRunParamsInputDataUnion struct {
	OfMapOfAnyMap []map[string]any `json:",omitzero,inline"`
	OfAnyMap      map[string]any   `json:",omitzero,inline"`
	paramUnion
}

func (u TaskAgentRunParamsInputDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMapOfAnyMap, u.OfAnyMap)
}
func (u *TaskAgentRunParamsInputDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *TaskAgentRunParamsInputDataUnion) asAny() any {
	if !param.IsOmitted(u.OfMapOfAnyMap) {
		return &u.OfMapOfAnyMap
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Source guidance overriding the agent default.
type TaskAgentRunParamsSources struct {
	// Free-text guidance describing sources or domains to avoid.
	Avoid param.Opt[string] `json:"avoid,omitzero"`
	// Free-text guidance describing sources or domains to prioritize.
	Prioritize param.Opt[string] `json:"prioritize,omitzero"`
	// Source groups the agent is allowed to use.
	Allow []TaskAgentRunParamsSourcesAllow `json:"allow,omitzero"`
	// Source groups the agent should not use.
	Block []TaskAgentRunParamsSourcesBlock `json:"block,omitzero"`
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
	// Domains included in this source group.
	Domains []string `json:"domains,omitzero" api:"required"`
	// Source group title.
	Title string `json:"title" api:"required"`
	// Zero-based source group position.
	Order param.Opt[int64] `json:"order,omitzero"`
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
	// Domains included in this source group.
	Domains []string `json:"domains,omitzero" api:"required"`
	// Source group title.
	Title string `json:"title" api:"required"`
	// Zero-based source group position.
	Order param.Opt[int64] `json:"order,omitzero"`
	paramObj
}

func (r TaskAgentRunParamsSourcesBlock) MarshalJSON() (data []byte, err error) {
	type shadow TaskAgentRunParamsSourcesBlock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskAgentRunParamsSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
