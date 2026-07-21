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

// AgentService contains methods and other services that help with interacting with
// the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentService] method instead.
type AgentService struct {
	Options   []option.RequestOption
	Templates AgentTemplateService
	Runs      AgentRunService
}

// NewAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentService(opts ...option.RequestOption) (r AgentService) {
	r = AgentService{}
	r.Options = opts
	r.Templates = NewAgentTemplateService(opts...)
	r.Runs = NewAgentRunService(opts...)
	return
}

// Create a Web Search Agent. Either pass `template` to materialize a pre-built
// template (its fields, goals, sources, and suggested questions are copied), or
// define the agent from scratch with `display_name`, `goals`, `sources`, and an
// optional `output_schema` for structured results.
func (r *AgentService) New(ctx context.Context, body AgentNewParams, opts ...option.RequestOption) (res *AgentNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update an agent with a
// [JSON Patch](https://datatracker.ietf.org/doc/html/rfc6902) document — an array
// of `{op, path, value}` operations applied to the agent, e.g.
// `[{"op": "replace", "path": "/display_name", "value": "My agent"}]`. Returns the
// updated agent.
func (r *AgentService) Update(ctx context.Context, agentID string, body AgentUpdateParams, opts ...option.RequestOption) (res *AgentUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/agents/%s", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List the active Web Search Agents in your account. Results are scoped to the
// workspace resolved from your token (or the optional `workspace_id` query
// parameter) and paginated with `offset`/`limit`.
func (r *AgentService) List(ctx context.Context, query AgentListParams, opts ...option.RequestOption) (res *AgentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deactivate an agent. This is a soft delete: the agent can no longer start new
// runs, but its existing runs and their results remain retrievable.
func (r *AgentService) Delete(ctx context.Context, agentID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return err
	}
	path := fmt.Sprintf("v2/agents/%s", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Retrieve a single Web Search Agent by ID.
func (r *AgentService) Get(ctx context.Context, agentID string, opts ...option.RequestOption) (res *AgentGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/agents/%s", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Creates a minimal persistent Web Search Agent and starts a run for it. The
// response includes `web_search_agent_id` for later agent and run queries.
func (r *AgentService) Run(ctx context.Context, body AgentRunParams, opts ...option.RequestOption) (res *AgentRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/agents/runs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type AgentNewResponse struct {
	// Unique web search agent identifier (wsa\_<uuid>).
	ID string `json:"id" api:"required"`
	// When the agent was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Agent description shown to users.
	Description string `json:"description" api:"required"`
	// Human-friendly agent name shown to users.
	DisplayName string `json:"display_name" api:"required"`
	// Default effort level for this agent's runs.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort AgentNewResponseEffort `json:"effort" api:"required"`
	// Ordered goals for the agent to follow.
	Goals []AgentNewResponseGoal `json:"goals" api:"required"`
	// Icon identifier used when presenting the agent.
	Icon string `json:"icon" api:"required"`
	// Whether the agent can be used to start new runs.
	IsActive bool `json:"is_active" api:"required"`
	// JSON schema describing the structured output the agent should produce.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Skill or operating context for the agent.
	Skill string `json:"skill" api:"required"`
	// Source guidance for the agent.
	Sources AgentNewResponseSources `json:"sources" api:"required"`
	// Suggested prompts users can run with this agent.
	SuggestedQuestions []AgentNewResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	// When the agent was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Primary use case supported by the agent.
	//
	// Any of "research", "enrichment", "dataset_building".
	UseCase AgentNewResponseUseCase `json:"use_case" api:"required"`
	// Stable agent name.
	AgentName string `json:"agent_name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		DisplayName        respjson.Field
		Effort             respjson.Field
		Goals              respjson.Field
		Icon               respjson.Field
		IsActive           respjson.Field
		OutputSchema       respjson.Field
		Skill              respjson.Field
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
func (r AgentNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default effort level for this agent's runs.
type AgentNewResponseEffort string

const (
	AgentNewResponseEffortLow    AgentNewResponseEffort = "low"
	AgentNewResponseEffortMedium AgentNewResponseEffort = "medium"
	AgentNewResponseEffortHigh   AgentNewResponseEffort = "high"
	AgentNewResponseEffortXHigh  AgentNewResponseEffort = "x-high"
	AgentNewResponseEffortMax    AgentNewResponseEffort = "max"
)

type AgentNewResponseGoal struct {
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
func (r AgentNewResponseGoal) RawJSON() string { return r.JSON.raw }
func (r *AgentNewResponseGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source guidance for the agent.
type AgentNewResponseSources struct {
	// Source groups the agent is allowed to use.
	Allow []AgentNewResponseSourcesAllow `json:"allow"`
	// Free-text guidance describing sources or domains to avoid.
	Avoid string `json:"avoid" api:"nullable"`
	// Source groups the agent should not use.
	Block []AgentNewResponseSourcesBlock `json:"block"`
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
func (r AgentNewResponseSources) RawJSON() string { return r.JSON.raw }
func (r *AgentNewResponseSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentNewResponseSourcesAllow struct {
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
func (r AgentNewResponseSourcesAllow) RawJSON() string { return r.JSON.raw }
func (r *AgentNewResponseSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentNewResponseSourcesBlock struct {
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
func (r AgentNewResponseSourcesBlock) RawJSON() string { return r.JSON.raw }
func (r *AgentNewResponseSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentNewResponseSuggestedQuestion struct {
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
func (r AgentNewResponseSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *AgentNewResponseSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary use case supported by the agent.
type AgentNewResponseUseCase string

const (
	AgentNewResponseUseCaseResearch        AgentNewResponseUseCase = "research"
	AgentNewResponseUseCaseEnrichment      AgentNewResponseUseCase = "enrichment"
	AgentNewResponseUseCaseDatasetBuilding AgentNewResponseUseCase = "dataset_building"
)

type AgentUpdateResponse struct {
	// Unique web search agent identifier (wsa\_<uuid>).
	ID string `json:"id" api:"required"`
	// When the agent was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Agent description shown to users.
	Description string `json:"description" api:"required"`
	// Human-friendly agent name shown to users.
	DisplayName string `json:"display_name" api:"required"`
	// Default effort level for this agent's runs.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort AgentUpdateResponseEffort `json:"effort" api:"required"`
	// Ordered goals for the agent to follow.
	Goals []AgentUpdateResponseGoal `json:"goals" api:"required"`
	// Icon identifier used when presenting the agent.
	Icon string `json:"icon" api:"required"`
	// Whether the agent can be used to start new runs.
	IsActive bool `json:"is_active" api:"required"`
	// JSON schema describing the structured output the agent should produce.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Skill or operating context for the agent.
	Skill string `json:"skill" api:"required"`
	// Source guidance for the agent.
	Sources AgentUpdateResponseSources `json:"sources" api:"required"`
	// Suggested prompts users can run with this agent.
	SuggestedQuestions []AgentUpdateResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	// When the agent was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Primary use case supported by the agent.
	//
	// Any of "research", "enrichment", "dataset_building".
	UseCase AgentUpdateResponseUseCase `json:"use_case" api:"required"`
	// Stable agent name.
	AgentName string `json:"agent_name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		DisplayName        respjson.Field
		Effort             respjson.Field
		Goals              respjson.Field
		Icon               respjson.Field
		IsActive           respjson.Field
		OutputSchema       respjson.Field
		Skill              respjson.Field
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
func (r AgentUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default effort level for this agent's runs.
type AgentUpdateResponseEffort string

const (
	AgentUpdateResponseEffortLow    AgentUpdateResponseEffort = "low"
	AgentUpdateResponseEffortMedium AgentUpdateResponseEffort = "medium"
	AgentUpdateResponseEffortHigh   AgentUpdateResponseEffort = "high"
	AgentUpdateResponseEffortXHigh  AgentUpdateResponseEffort = "x-high"
	AgentUpdateResponseEffortMax    AgentUpdateResponseEffort = "max"
)

type AgentUpdateResponseGoal struct {
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
func (r AgentUpdateResponseGoal) RawJSON() string { return r.JSON.raw }
func (r *AgentUpdateResponseGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source guidance for the agent.
type AgentUpdateResponseSources struct {
	// Source groups the agent is allowed to use.
	Allow []AgentUpdateResponseSourcesAllow `json:"allow"`
	// Free-text guidance describing sources or domains to avoid.
	Avoid string `json:"avoid" api:"nullable"`
	// Source groups the agent should not use.
	Block []AgentUpdateResponseSourcesBlock `json:"block"`
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
func (r AgentUpdateResponseSources) RawJSON() string { return r.JSON.raw }
func (r *AgentUpdateResponseSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentUpdateResponseSourcesAllow struct {
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
func (r AgentUpdateResponseSourcesAllow) RawJSON() string { return r.JSON.raw }
func (r *AgentUpdateResponseSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentUpdateResponseSourcesBlock struct {
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
func (r AgentUpdateResponseSourcesBlock) RawJSON() string { return r.JSON.raw }
func (r *AgentUpdateResponseSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentUpdateResponseSuggestedQuestion struct {
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
func (r AgentUpdateResponseSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *AgentUpdateResponseSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary use case supported by the agent.
type AgentUpdateResponseUseCase string

const (
	AgentUpdateResponseUseCaseResearch        AgentUpdateResponseUseCase = "research"
	AgentUpdateResponseUseCaseEnrichment      AgentUpdateResponseUseCase = "enrichment"
	AgentUpdateResponseUseCaseDatasetBuilding AgentUpdateResponseUseCase = "dataset_building"
)

type AgentListResponse struct {
	// Items returned in this page.
	Items []AgentListResponseItem `json:"items" api:"required"`
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
func (r AgentListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentListResponseItem struct {
	// Unique web search agent identifier (wsa\_<uuid>).
	ID string `json:"id" api:"required"`
	// When the agent was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Agent description shown to users.
	Description string `json:"description" api:"required"`
	// Human-friendly agent name shown to users.
	DisplayName string `json:"display_name" api:"required"`
	// Default effort level for this agent's runs.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort string `json:"effort" api:"required"`
	// Ordered goals for the agent to follow.
	Goals []AgentListResponseItemGoal `json:"goals" api:"required"`
	// Icon identifier used when presenting the agent.
	Icon string `json:"icon" api:"required"`
	// Whether the agent can be used to start new runs.
	IsActive bool `json:"is_active" api:"required"`
	// JSON schema describing the structured output the agent should produce.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Skill or operating context for the agent.
	Skill string `json:"skill" api:"required"`
	// Source guidance for the agent.
	Sources AgentListResponseItemSources `json:"sources" api:"required"`
	// Suggested prompts users can run with this agent.
	SuggestedQuestions []AgentListResponseItemSuggestedQuestion `json:"suggested_questions" api:"required"`
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
		Effort             respjson.Field
		Goals              respjson.Field
		Icon               respjson.Field
		IsActive           respjson.Field
		OutputSchema       respjson.Field
		Skill              respjson.Field
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
func (r AgentListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentListResponseItemGoal struct {
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
func (r AgentListResponseItemGoal) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponseItemGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source guidance for the agent.
type AgentListResponseItemSources struct {
	// Source groups the agent is allowed to use.
	Allow []AgentListResponseItemSourcesAllow `json:"allow"`
	// Free-text guidance describing sources or domains to avoid.
	Avoid string `json:"avoid" api:"nullable"`
	// Source groups the agent should not use.
	Block []AgentListResponseItemSourcesBlock `json:"block"`
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
func (r AgentListResponseItemSources) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponseItemSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentListResponseItemSourcesAllow struct {
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
func (r AgentListResponseItemSourcesAllow) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponseItemSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentListResponseItemSourcesBlock struct {
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
func (r AgentListResponseItemSourcesBlock) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponseItemSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentListResponseItemSuggestedQuestion struct {
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
func (r AgentListResponseItemSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *AgentListResponseItemSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGetResponse struct {
	// Unique web search agent identifier (wsa\_<uuid>).
	ID string `json:"id" api:"required"`
	// When the agent was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Agent description shown to users.
	Description string `json:"description" api:"required"`
	// Human-friendly agent name shown to users.
	DisplayName string `json:"display_name" api:"required"`
	// Default effort level for this agent's runs.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort AgentGetResponseEffort `json:"effort" api:"required"`
	// Ordered goals for the agent to follow.
	Goals []AgentGetResponseGoal `json:"goals" api:"required"`
	// Icon identifier used when presenting the agent.
	Icon string `json:"icon" api:"required"`
	// Whether the agent can be used to start new runs.
	IsActive bool `json:"is_active" api:"required"`
	// JSON schema describing the structured output the agent should produce.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Skill or operating context for the agent.
	Skill string `json:"skill" api:"required"`
	// Source guidance for the agent.
	Sources AgentGetResponseSources `json:"sources" api:"required"`
	// Suggested prompts users can run with this agent.
	SuggestedQuestions []AgentGetResponseSuggestedQuestion `json:"suggested_questions" api:"required"`
	// When the agent was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Primary use case supported by the agent.
	//
	// Any of "research", "enrichment", "dataset_building".
	UseCase AgentGetResponseUseCase `json:"use_case" api:"required"`
	// Stable agent name.
	AgentName string `json:"agent_name" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Description        respjson.Field
		DisplayName        respjson.Field
		Effort             respjson.Field
		Goals              respjson.Field
		Icon               respjson.Field
		IsActive           respjson.Field
		OutputSchema       respjson.Field
		Skill              respjson.Field
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
func (r AgentGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default effort level for this agent's runs.
type AgentGetResponseEffort string

const (
	AgentGetResponseEffortLow    AgentGetResponseEffort = "low"
	AgentGetResponseEffortMedium AgentGetResponseEffort = "medium"
	AgentGetResponseEffortHigh   AgentGetResponseEffort = "high"
	AgentGetResponseEffortXHigh  AgentGetResponseEffort = "x-high"
	AgentGetResponseEffortMax    AgentGetResponseEffort = "max"
)

type AgentGetResponseGoal struct {
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
func (r AgentGetResponseGoal) RawJSON() string { return r.JSON.raw }
func (r *AgentGetResponseGoal) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Source guidance for the agent.
type AgentGetResponseSources struct {
	// Source groups the agent is allowed to use.
	Allow []AgentGetResponseSourcesAllow `json:"allow"`
	// Free-text guidance describing sources or domains to avoid.
	Avoid string `json:"avoid" api:"nullable"`
	// Source groups the agent should not use.
	Block []AgentGetResponseSourcesBlock `json:"block"`
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
func (r AgentGetResponseSources) RawJSON() string { return r.JSON.raw }
func (r *AgentGetResponseSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGetResponseSourcesAllow struct {
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
func (r AgentGetResponseSourcesAllow) RawJSON() string { return r.JSON.raw }
func (r *AgentGetResponseSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGetResponseSourcesBlock struct {
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
func (r AgentGetResponseSourcesBlock) RawJSON() string { return r.JSON.raw }
func (r *AgentGetResponseSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGetResponseSuggestedQuestion struct {
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
func (r AgentGetResponseSuggestedQuestion) RawJSON() string { return r.JSON.raw }
func (r *AgentGetResponseSuggestedQuestion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary use case supported by the agent.
type AgentGetResponseUseCase string

const (
	AgentGetResponseUseCaseResearch        AgentGetResponseUseCase = "research"
	AgentGetResponseUseCaseEnrichment      AgentGetResponseUseCase = "enrichment"
	AgentGetResponseUseCaseDatasetBuilding AgentGetResponseUseCase = "dataset_building"
)

type AgentRunResponse struct {
	// Run identifier, format "task*run*{uuid}".
	ID string `json:"id" api:"required"`
	// When the run was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Effort level used for the run.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort AgentRunResponseEffort `json:"effort" api:"required"`
	// Interaction ID.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Current run status.
	//
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status AgentRunResponseStatus `json:"status" api:"required"`
	// Web Search Agent instance this run belongs to.
	WebSearchAgentID string `json:"web_search_agent_id" api:"required"`
	// When the run completed.
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Error details when the run failed.
	Error AgentRunResponseError `json:"error" api:"nullable"`
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
func (r AgentRunResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Effort level used for the run.
type AgentRunResponseEffort string

const (
	AgentRunResponseEffortLow    AgentRunResponseEffort = "low"
	AgentRunResponseEffortMedium AgentRunResponseEffort = "medium"
	AgentRunResponseEffortHigh   AgentRunResponseEffort = "high"
	AgentRunResponseEffortXHigh  AgentRunResponseEffort = "x-high"
	AgentRunResponseEffortMax    AgentRunResponseEffort = "max"
)

// Current run status.
type AgentRunResponseStatus string

const (
	AgentRunResponseStatusQueued    AgentRunResponseStatus = "queued"
	AgentRunResponseStatusRunning   AgentRunResponseStatus = "running"
	AgentRunResponseStatusCompleted AgentRunResponseStatus = "completed"
	AgentRunResponseStatusFailed    AgentRunResponseStatus = "failed"
	AgentRunResponseStatusCancelled AgentRunResponseStatus = "cancelled"
)

// Error details when the run failed.
type AgentRunResponseError struct {
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
func (r AgentRunResponseError) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentNewParams struct {
	// Stable agent name.
	AgentName param.Opt[string] `json:"agent_name,omitzero"`
	// Agent description shown to users.
	Description param.Opt[string] `json:"description,omitzero"`
	// Human-friendly agent name shown to users.
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// Icon identifier used when presenting the agent.
	Icon param.Opt[string] `json:"icon,omitzero"`
	// Skill or operating context for the agent.
	Skill param.Opt[string] `json:"skill,omitzero"`
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
	UseCase AgentNewParamsUseCase `json:"use_case,omitzero"`
	// Default effort level for this agent's runs.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort AgentNewParamsEffort `json:"effort,omitzero"`
	// Ordered goals for the agent to follow.
	Goals []string `json:"goals,omitzero"`
	// Source guidance for the agent.
	Sources AgentNewParamsSources `json:"sources,omitzero"`
	// Suggested prompts users can run with this agent.
	SuggestedQuestions []string `json:"suggested_questions,omitzero"`
	paramObj
}

func (r AgentNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Default effort level for this agent's runs.
type AgentNewParamsEffort string

const (
	AgentNewParamsEffortLow    AgentNewParamsEffort = "low"
	AgentNewParamsEffortMedium AgentNewParamsEffort = "medium"
	AgentNewParamsEffortHigh   AgentNewParamsEffort = "high"
	AgentNewParamsEffortXHigh  AgentNewParamsEffort = "x-high"
	AgentNewParamsEffortMax    AgentNewParamsEffort = "max"
)

// Source guidance for the agent.
type AgentNewParamsSources struct {
	// Free-text guidance describing sources or domains to avoid.
	Avoid param.Opt[string] `json:"avoid,omitzero"`
	// Free-text guidance describing sources or domains to prioritize.
	Prioritize param.Opt[string] `json:"prioritize,omitzero"`
	// Source groups the agent is allowed to use.
	Allow []AgentNewParamsSourcesAllow `json:"allow,omitzero"`
	// Source groups the agent should not use.
	Block []AgentNewParamsSourcesBlock `json:"block,omitzero"`
	paramObj
}

func (r AgentNewParamsSources) MarshalJSON() (data []byte, err error) {
	type shadow AgentNewParamsSources
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentNewParamsSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Domains, Title are required.
type AgentNewParamsSourcesAllow struct {
	// Domains included in this source group.
	Domains []string `json:"domains,omitzero" api:"required"`
	// Source group title.
	Title string `json:"title" api:"required"`
	// Zero-based source group position.
	Order param.Opt[int64] `json:"order,omitzero"`
	paramObj
}

func (r AgentNewParamsSourcesAllow) MarshalJSON() (data []byte, err error) {
	type shadow AgentNewParamsSourcesAllow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentNewParamsSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Domains, Title are required.
type AgentNewParamsSourcesBlock struct {
	// Domains included in this source group.
	Domains []string `json:"domains,omitzero" api:"required"`
	// Source group title.
	Title string `json:"title" api:"required"`
	// Zero-based source group position.
	Order param.Opt[int64] `json:"order,omitzero"`
	paramObj
}

func (r AgentNewParamsSourcesBlock) MarshalJSON() (data []byte, err error) {
	type shadow AgentNewParamsSourcesBlock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentNewParamsSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Primary use case supported by the agent.
type AgentNewParamsUseCase string

const (
	AgentNewParamsUseCaseResearch        AgentNewParamsUseCase = "research"
	AgentNewParamsUseCaseEnrichment      AgentNewParamsUseCase = "enrichment"
	AgentNewParamsUseCaseDatasetBuilding AgentNewParamsUseCase = "dataset_building"
)

type AgentUpdateParams struct {
	// A JSON Patch document per RFC 6902 — a JSON array of patch operations.
	Body []AgentUpdateParamsBody
	paramObj
}

func (r AgentUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *AgentUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single JSON Patch operation per RFC 6902.
//
// The properties Op, Path are required.
type AgentUpdateParamsBody struct {
	// Any of "add", "remove", "replace", "move", "copy", "test".
	Op    string            `json:"op,omitzero" api:"required"`
	Path  string            `json:"path" api:"required"`
	From  param.Opt[string] `json:"from,omitzero"`
	Value any               `json:"value,omitzero"`
	paramObj
}

func (r AgentUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow AgentUpdateParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentUpdateParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AgentUpdateParamsBody](
		"op", "add", "remove", "replace", "move", "copy", "test",
	)
}

type AgentListParams struct {
	WorkspaceID param.Opt[string] `query:"workspace_id,omitzero" format:"uuid" json:"-"`
	Limit       param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Offset      param.Opt[int64]  `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentListParams]'s query parameters as `url.Values`.
func (r AgentListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentRunParams struct {
	// User prompt or task instructions for the run.
	Input string `json:"input" api:"required"`
	// Previous interaction identifier used to continue a conversation.
	PreviousInteractionID param.Opt[string] `json:"previous_interaction_id,omitzero"`
	// Whether to stream run events when supported.
	EnableEvents param.Opt[bool] `json:"enable_events,omitzero"`
	// Canonical effort tier names for the research graph.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort AgentRunParamsEffort `json:"effort,omitzero"`
	// Existing records to ENRICH: a list of partial rows, or a single object,
	// mirroring output_schema's shape.
	InputData AgentRunParamsInputDataUnion `json:"input_data,omitzero"`
	// JSON schema overriding the agent's default structured output for this run.
	OutputSchema map[string]any `json:"output_schema,omitzero"`
	// Source guidance overriding the agent default.
	Sources AgentRunParamsSources `json:"sources,omitzero"`
	paramObj
}

func (r AgentRunParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Canonical effort tier names for the research graph.
type AgentRunParamsEffort string

const (
	AgentRunParamsEffortLow    AgentRunParamsEffort = "low"
	AgentRunParamsEffortMedium AgentRunParamsEffort = "medium"
	AgentRunParamsEffortHigh   AgentRunParamsEffort = "high"
	AgentRunParamsEffortXHigh  AgentRunParamsEffort = "x-high"
	AgentRunParamsEffortMax    AgentRunParamsEffort = "max"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentRunParamsInputDataUnion struct {
	OfMapOfAnyMap []map[string]any `json:",omitzero,inline"`
	OfAnyMap      map[string]any   `json:",omitzero,inline"`
	paramUnion
}

func (u AgentRunParamsInputDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMapOfAnyMap, u.OfAnyMap)
}
func (u *AgentRunParamsInputDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentRunParamsInputDataUnion) asAny() any {
	if !param.IsOmitted(u.OfMapOfAnyMap) {
		return &u.OfMapOfAnyMap
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Source guidance overriding the agent default.
type AgentRunParamsSources struct {
	// Free-text guidance describing sources or domains to avoid.
	Avoid param.Opt[string] `json:"avoid,omitzero"`
	// Free-text guidance describing sources or domains to prioritize.
	Prioritize param.Opt[string] `json:"prioritize,omitzero"`
	// Source groups the agent is allowed to use.
	Allow []AgentRunParamsSourcesAllow `json:"allow,omitzero"`
	// Source groups the agent should not use.
	Block []AgentRunParamsSourcesBlock `json:"block,omitzero"`
	paramObj
}

func (r AgentRunParamsSources) MarshalJSON() (data []byte, err error) {
	type shadow AgentRunParamsSources
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRunParamsSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Domains, Title are required.
type AgentRunParamsSourcesAllow struct {
	// Domains included in this source group.
	Domains []string `json:"domains,omitzero" api:"required"`
	// Source group title.
	Title string `json:"title" api:"required"`
	// Zero-based source group position.
	Order param.Opt[int64] `json:"order,omitzero"`
	paramObj
}

func (r AgentRunParamsSourcesAllow) MarshalJSON() (data []byte, err error) {
	type shadow AgentRunParamsSourcesAllow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRunParamsSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Domains, Title are required.
type AgentRunParamsSourcesBlock struct {
	// Domains included in this source group.
	Domains []string `json:"domains,omitzero" api:"required"`
	// Source group title.
	Title string `json:"title" api:"required"`
	// Zero-based source group position.
	Order param.Opt[int64] `json:"order,omitzero"`
	paramObj
}

func (r AgentRunParamsSourcesBlock) MarshalJSON() (data []byte, err error) {
	type shadow AgentRunParamsSourcesBlock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRunParamsSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
