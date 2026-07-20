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
	"time"

	"github.com/Nimbleway/nimble-go/internal/apijson"
	"github.com/Nimbleway/nimble-go/internal/apiquery"
	"github.com/Nimbleway/nimble-go/internal/requestconfig"
	"github.com/Nimbleway/nimble-go/option"
	"github.com/Nimbleway/nimble-go/packages/param"
	"github.com/Nimbleway/nimble-go/packages/respjson"
)

// AgentRunService contains methods and other services that help with interacting
// with the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentRunService] method instead.
type AgentRunService struct {
	Options []option.RequestOption
}

// NewAgentRunService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentRunService(opts ...option.RequestOption) (r AgentRunService) {
	r = AgentRunService{}
	r.Options = opts
	return
}

// Start an agent run. The run executes asynchronously: the response returns
// immediately with status `queued`, then poll `GET .../runs/{run_id}` until
// `completed` and fetch the output from `GET .../runs/{run_id}/result` — or set
// `enable_events: true` and follow `GET .../runs/{run_id}/events` for live
// progress.
//
// To enrich existing records instead of researching from scratch, pass them in
// `input_data`; this requires an `output_schema` (on the request or the agent).
func (r *AgentRunService) New(ctx context.Context, agentID string, body AgentRunNewParams, opts ...option.RequestOption) (res *AgentRunNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/agents/%s/runs", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// List the runs of an agent, newest first, paginated with `offset`/`limit`.
func (r *AgentRunService) List(ctx context.Context, agentID string, query AgentRunListParams, opts ...option.RequestOption) (res *AgentRunListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/agents/%s/runs", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Retrieve a run's current state. Poll this endpoint after creating a run: the run
// is finished once `status` is `completed`, `failed`, or `cancelled`.
func (r *AgentRunService) Get(ctx context.Context, runID string, query AgentRunGetParams, opts ...option.RequestOption) (res *AgentRunGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return nil, err
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/agents/%s/runs/%s", query.AgentID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Fetch the output of a completed run. The `output` is `type: "text"` (a prose
// answer) or `type: "json"` (structured data matching the output schema), plus
// `trust` metadata with per-claim citations for the answer.
//
// While the run is still `queued` or `running` this endpoint returns `409`; if the
// run `failed` or was `cancelled` it returns `422` with the run and error details.
func (r *AgentRunService) Result(ctx context.Context, runID string, query AgentRunResultParams, opts ...option.RequestOption) (res *AgentRunResultResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return nil, err
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/agents/%s/runs/%s/result", query.AgentID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Stream a run's progress as
// [server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events)
// (`text/event-stream`). Create the run with `enable_events: true` to have events
// published. A keep-alive comment is sent every 15 seconds.
func (r *AgentRunService) StreamEvents(ctx context.Context, runID string, query AgentRunStreamEventsParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if query.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return err
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return err
	}
	path := fmt.Sprintf("v2/agents/%s/runs/%s/events", query.AgentID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

type AgentRunNewResponse struct {
	// Run identifier, format "task*run*{uuid}".
	ID string `json:"id" api:"required"`
	// When the run was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Effort level used for the run.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort AgentRunNewResponseEffort `json:"effort" api:"required"`
	// Interaction ID.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Current run status.
	//
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status AgentRunNewResponseStatus `json:"status" api:"required"`
	// Web Search Agent instance this run belongs to.
	WebSearchAgentID string `json:"web_search_agent_id" api:"required"`
	// When the run completed.
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Error details when the run failed.
	Error AgentRunNewResponseError `json:"error" api:"nullable"`
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
func (r AgentRunNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentRunNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Effort level used for the run.
type AgentRunNewResponseEffort string

const (
	AgentRunNewResponseEffortLow    AgentRunNewResponseEffort = "low"
	AgentRunNewResponseEffortMedium AgentRunNewResponseEffort = "medium"
	AgentRunNewResponseEffortHigh   AgentRunNewResponseEffort = "high"
	AgentRunNewResponseEffortXHigh  AgentRunNewResponseEffort = "x-high"
	AgentRunNewResponseEffortMax    AgentRunNewResponseEffort = "max"
)

// Current run status.
type AgentRunNewResponseStatus string

const (
	AgentRunNewResponseStatusQueued    AgentRunNewResponseStatus = "queued"
	AgentRunNewResponseStatusRunning   AgentRunNewResponseStatus = "running"
	AgentRunNewResponseStatusCompleted AgentRunNewResponseStatus = "completed"
	AgentRunNewResponseStatusFailed    AgentRunNewResponseStatus = "failed"
	AgentRunNewResponseStatusCancelled AgentRunNewResponseStatus = "cancelled"
)

// Error details when the run failed.
type AgentRunNewResponseError struct {
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
func (r AgentRunNewResponseError) RawJSON() string { return r.JSON.raw }
func (r *AgentRunNewResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunListResponse struct {
	// Items returned in this page.
	Items []AgentRunListResponseItem `json:"items" api:"required"`
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
func (r AgentRunListResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentRunListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunListResponseItem struct {
	// Run identifier, format "task*run*{uuid}".
	ID string `json:"id" api:"required"`
	// When the run was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Effort level used for the run.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort string `json:"effort" api:"required"`
	// Interaction ID.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Current run status.
	//
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status string `json:"status" api:"required"`
	// Web Search Agent instance this run belongs to.
	WebSearchAgentID string `json:"web_search_agent_id" api:"required"`
	// When the run completed.
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Error details when the run failed.
	Error AgentRunListResponseItemError `json:"error" api:"nullable"`
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
func (r AgentRunListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *AgentRunListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details when the run failed.
type AgentRunListResponseItemError struct {
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
func (r AgentRunListResponseItemError) RawJSON() string { return r.JSON.raw }
func (r *AgentRunListResponseItemError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunGetResponse struct {
	// Run identifier, format "task*run*{uuid}".
	ID string `json:"id" api:"required"`
	// When the run was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Effort level used for the run.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort AgentRunGetResponseEffort `json:"effort" api:"required"`
	// Interaction ID.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Current run status.
	//
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status AgentRunGetResponseStatus `json:"status" api:"required"`
	// Web Search Agent instance this run belongs to.
	WebSearchAgentID string `json:"web_search_agent_id" api:"required"`
	// When the run completed.
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Error details when the run failed.
	Error AgentRunGetResponseError `json:"error" api:"nullable"`
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
func (r AgentRunGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentRunGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Effort level used for the run.
type AgentRunGetResponseEffort string

const (
	AgentRunGetResponseEffortLow    AgentRunGetResponseEffort = "low"
	AgentRunGetResponseEffortMedium AgentRunGetResponseEffort = "medium"
	AgentRunGetResponseEffortHigh   AgentRunGetResponseEffort = "high"
	AgentRunGetResponseEffortXHigh  AgentRunGetResponseEffort = "x-high"
	AgentRunGetResponseEffortMax    AgentRunGetResponseEffort = "max"
)

// Current run status.
type AgentRunGetResponseStatus string

const (
	AgentRunGetResponseStatusQueued    AgentRunGetResponseStatus = "queued"
	AgentRunGetResponseStatusRunning   AgentRunGetResponseStatus = "running"
	AgentRunGetResponseStatusCompleted AgentRunGetResponseStatus = "completed"
	AgentRunGetResponseStatusFailed    AgentRunGetResponseStatus = "failed"
	AgentRunGetResponseStatusCancelled AgentRunGetResponseStatus = "cancelled"
)

// Error details when the run failed.
type AgentRunGetResponseError struct {
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
func (r AgentRunGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *AgentRunGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentRunResultResponseUnion contains all possible properties and values from
// [AgentRunResultResponseTaskRunResultPublicV2],
// [AgentRunResultResponseTaskRunFailedResultPublicV2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AgentRunResultResponseUnion struct {
	// This field is from variant [AgentRunResultResponseTaskRunResultPublicV2].
	Output AgentRunResultResponseTaskRunResultPublicV2OutputUnion `json:"output"`
	// This field is a union of [AgentRunResultResponseTaskRunResultPublicV2Run],
	// [AgentRunResultResponseTaskRunFailedResultPublicV2Run]
	Run AgentRunResultResponseUnionRun `json:"run"`
	// This field is from variant [AgentRunResultResponseTaskRunFailedResultPublicV2].
	Error AgentRunResultResponseTaskRunFailedResultPublicV2Error `json:"error"`
	JSON  struct {
		Output respjson.Field
		Run    respjson.Field
		Error  respjson.Field
		raw    string
	} `json:"-"`
}

func (u AgentRunResultResponseUnion) AsTaskRunResultPublicV2() (v AgentRunResultResponseTaskRunResultPublicV2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentRunResultResponseUnion) AsTaskRunFailedResultPublicV2() (v AgentRunResultResponseTaskRunFailedResultPublicV2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentRunResultResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentRunResultResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentRunResultResponseUnionRun is an implicit subunion of
// [AgentRunResultResponseUnion]. AgentRunResultResponseUnionRun provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AgentRunResultResponseUnion].
type AgentRunResultResponseUnionRun struct {
	ID               string    `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	Effort           string    `json:"effort"`
	InteractionID    string    `json:"interaction_id"`
	IsActive         bool      `json:"is_active"`
	Status           string    `json:"status"`
	WebSearchAgentID string    `json:"web_search_agent_id"`
	CompletedAt      time.Time `json:"completed_at"`
	// This field is a union of [AgentRunResultResponseTaskRunResultPublicV2RunError],
	// [AgentRunResultResponseTaskRunFailedResultPublicV2RunError]
	Error     AgentRunResultResponseUnionRunError `json:"error"`
	Prompt    string                              `json:"prompt"`
	StartedAt time.Time                           `json:"started_at"`
	JSON      struct {
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
		raw              string
	} `json:"-"`
}

func (r *AgentRunResultResponseUnionRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentRunResultResponseUnionRunError is an implicit subunion of
// [AgentRunResultResponseUnion]. AgentRunResultResponseUnionRunError provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AgentRunResultResponseUnion].
type AgentRunResultResponseUnionRunError struct {
	Message string `json:"message"`
	RefID   string `json:"ref_id"`
	JSON    struct {
		Message respjson.Field
		RefID   respjson.Field
		raw     string
	} `json:"-"`
}

func (r *AgentRunResultResponseUnionRunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResultResponseTaskRunResultPublicV2 struct {
	// Output from the completed task.
	Output AgentRunResultResponseTaskRunResultPublicV2OutputUnion `json:"output" api:"required"`
	// Task run object with status 'completed'.
	Run AgentRunResultResponseTaskRunResultPublicV2Run `json:"run" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		Run         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResultResponseTaskRunResultPublicV2) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResultResponseTaskRunResultPublicV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentRunResultResponseTaskRunResultPublicV2OutputUnion contains all possible
// properties and values from
// [AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2],
// [AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AgentRunResultResponseTaskRunResultPublicV2OutputUnion struct {
	// This field is a union of [string],
	// [AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2ContentUnion]
	Content AgentRunResultResponseTaskRunResultPublicV2OutputUnionContent `json:"content"`
	// This field is a union of
	// [AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2Trust],
	// [AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2Trust]
	Trust AgentRunResultResponseTaskRunResultPublicV2OutputUnionTrust `json:"trust"`
	Type  string                                                      `json:"type"`
	JSON  struct {
		Content respjson.Field
		Trust   respjson.Field
		Type    respjson.Field
		raw     string
	} `json:"-"`
}

func (u AgentRunResultResponseTaskRunResultPublicV2OutputUnion) AsTaskRunTextOutputPublicV2() (v AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentRunResultResponseTaskRunResultPublicV2OutputUnion) AsTaskRunJsonOutputPublicV2() (v AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentRunResultResponseTaskRunResultPublicV2OutputUnion) RawJSON() string { return u.JSON.raw }

func (r *AgentRunResultResponseTaskRunResultPublicV2OutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentRunResultResponseTaskRunResultPublicV2OutputUnionContent is an implicit
// subunion of [AgentRunResultResponseTaskRunResultPublicV2OutputUnion].
// AgentRunResultResponseTaskRunResultPublicV2OutputUnionContent provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AgentRunResultResponseTaskRunResultPublicV2OutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfAgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2ContentMapItem
// OfAnyArray]
type AgentRunResultResponseTaskRunResultPublicV2OutputUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2ContentMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                                   respjson.Field
		OfAgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2ContentMapItem respjson.Field
		OfAnyArray                                                                                 respjson.Field
		raw                                                                                        string
	} `json:"-"`
}

func (r *AgentRunResultResponseTaskRunResultPublicV2OutputUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentRunResultResponseTaskRunResultPublicV2OutputUnionTrust is an implicit
// subunion of [AgentRunResultResponseTaskRunResultPublicV2OutputUnion].
// AgentRunResultResponseTaskRunResultPublicV2OutputUnionTrust provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AgentRunResultResponseTaskRunResultPublicV2OutputUnion].
type AgentRunResultResponseTaskRunResultPublicV2OutputUnionTrust struct {
	// This field is a union of
	// [[]AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustClaim],
	// [[]AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2TrustClaim]
	Claims     AgentRunResultResponseTaskRunResultPublicV2OutputUnionTrustClaims `json:"claims"`
	Confidence string                                                            `json:"confidence"`
	Reasoning  string                                                            `json:"reasoning"`
	// This field is a union of
	// [[]AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustSource],
	// [[]AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2TrustSource]
	Sources AgentRunResultResponseTaskRunResultPublicV2OutputUnionTrustSources `json:"sources"`
	JSON    struct {
		Claims     respjson.Field
		Confidence respjson.Field
		Reasoning  respjson.Field
		Sources    respjson.Field
		raw        string
	} `json:"-"`
}

func (r *AgentRunResultResponseTaskRunResultPublicV2OutputUnionTrust) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentRunResultResponseTaskRunResultPublicV2OutputUnionTrustClaims is an implicit
// subunion of [AgentRunResultResponseTaskRunResultPublicV2OutputUnion].
// AgentRunResultResponseTaskRunResultPublicV2OutputUnionTrustClaims provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AgentRunResultResponseTaskRunResultPublicV2OutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfClaims]
type AgentRunResultResponseTaskRunResultPublicV2OutputUnionTrustClaims struct {
	// This field will be present if the value is a
	// [[]AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustClaim]
	// instead of an object.
	OfClaims []AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustClaim `json:",inline"`
	JSON     struct {
		OfClaims respjson.Field
		raw      string
	} `json:"-"`
}

func (r *AgentRunResultResponseTaskRunResultPublicV2OutputUnionTrustClaims) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentRunResultResponseTaskRunResultPublicV2OutputUnionTrustSources is an
// implicit subunion of [AgentRunResultResponseTaskRunResultPublicV2OutputUnion].
// AgentRunResultResponseTaskRunResultPublicV2OutputUnionTrustSources provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AgentRunResultResponseTaskRunResultPublicV2OutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfSources]
type AgentRunResultResponseTaskRunResultPublicV2OutputUnionTrustSources struct {
	// This field will be present if the value is a
	// [[]AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustSource]
	// instead of an object.
	OfSources []AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustSource `json:",inline"`
	JSON      struct {
		OfSources respjson.Field
		raw       string
	} `json:"-"`
}

func (r *AgentRunResultResponseTaskRunResultPublicV2OutputUnionTrustSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2 struct {
	// The final prose answer.
	Content string `json:"content" api:"required"`
	// Trust and citation metadata for the output.
	Trust AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2Trust `json:"trust" api:"required"`
	// Output content type.
	//
	// Any of "text".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Trust       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2) RawJSON() string {
	return r.JSON.raw
}
func (r *AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trust and citation metadata for the output.
type AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2Trust struct {
	// Per-claim trust, keyed by callout markers in the answer text.
	Claims []AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustClaim `json:"claims" api:"required"`
	// Overall confidence in the answer.
	//
	// Any of "high", "medium", "low", "pre_existing".
	Confidence string `json:"confidence" api:"required"`
	// Why this confidence level was assigned.
	Reasoning string `json:"reasoning" api:"required"`
	// Sources consulted while producing the answer.
	Sources []AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustSource `json:"sources" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Claims      respjson.Field
		Confidence  respjson.Field
		Reasoning   respjson.Field
		Sources     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2Trust) RawJSON() string {
	return r.JSON.raw
}
func (r *AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2Trust) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trust metadata for one claim in a prose answer, keyed by callout marker.
type AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustClaim struct {
	// Callout marker number referencing this claim in the answer text.
	Callout int64 `json:"callout" api:"required"`
	// Citations backing this claim.
	Citations []AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustClaimCitation `json:"citations" api:"required"`
	// Confidence in this claim.
	//
	// Any of "high", "medium", "low", "pre_existing".
	Confidence string `json:"confidence" api:"required"`
	// Why this confidence level was assigned.
	Reasoning string `json:"reasoning" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Callout     respjson.Field
		Citations   respjson.Field
		Confidence  respjson.Field
		Reasoning   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustClaim) RawJSON() string {
	return r.JSON.raw
}
func (r *AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustClaim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A citation backing a specific claim in the answer.
type AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustClaimCitation struct {
	// URL of the cited page.
	URL string `json:"url" api:"required"`
	// Verbatim excerpts supporting the claim.
	Excerpts []string `json:"excerpts" api:"nullable"`
	// Extract template used to read the source, when one was used.
	ExtractTemplateName string `json:"extract_template_name" api:"nullable"`
	// What _kind_ of source this is (classified by the compress LLM), independent of
	// TrustSourceType (how authoritative it is for a specific claim). Deliberately
	// uses "official" rather than "primary" so the two axes can never collide.
	//
	// Also doubles as the sub-question's `source_intent` (what kind of source a
	// question _needs_) — the two concepts overlap enough that a single enum lets
	// `classify_source_importance` compare "what we got" against "what we asked for"
	// directly.
	//
	// Any of "official", "news", "social", "academic", "aggregator", "other".
	SourceCategory string `json:"source_category" api:"nullable"`
	// What _kind_ of source this is (classified by the compress LLM), independent of
	// TrustSourceType (how authoritative it is for a specific claim). Deliberately
	// uses "official" rather than "primary" so the two axes can never collide.
	//
	// Also doubles as the sub-question's `source_intent` (what kind of source a
	// question _needs_) — the two concepts overlap enough that a single enum lets
	// `classify_source_importance` compare "what we got" against "what we asked for"
	// directly.
	//
	// Any of "official", "news", "social", "academic", "aggregator", "other".
	SourceIntent string `json:"source_intent" api:"nullable"`
	// How authoritative the source is: 'primary' or 'secondary'.
	//
	// Any of "primary", "secondary".
	SourceType string `json:"source_type" api:"nullable"`
	// Title of the cited page.
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL                 respjson.Field
		Excerpts            respjson.Field
		ExtractTemplateName respjson.Field
		SourceCategory      respjson.Field
		SourceIntent        respjson.Field
		SourceType          respjson.Field
		Title               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustClaimCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustClaimCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A source consulted while producing the answer.
type AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustSource struct {
	// How authoritative the source is: 'primary' or 'secondary'.
	//
	// Any of "primary", "secondary".
	Type string `json:"type" api:"required"`
	// URL of the source page.
	URL string `json:"url" api:"required"`
	// Extract template used to read the source, when one was used.
	ExtractTemplateName string `json:"extract_template_name" api:"nullable"`
	// What _kind_ of source this is (classified by the compress LLM), independent of
	// TrustSourceType (how authoritative it is for a specific claim). Deliberately
	// uses "official" rather than "primary" so the two axes can never collide.
	//
	// Also doubles as the sub-question's `source_intent` (what kind of source a
	// question _needs_) — the two concepts overlap enough that a single enum lets
	// `classify_source_importance` compare "what we got" against "what we asked for"
	// directly.
	//
	// Any of "official", "news", "social", "academic", "aggregator", "other".
	SourceCategory string `json:"source_category" api:"nullable"`
	// What _kind_ of source this is (classified by the compress LLM), independent of
	// TrustSourceType (how authoritative it is for a specific claim). Deliberately
	// uses "official" rather than "primary" so the two axes can never collide.
	//
	// Also doubles as the sub-question's `source_intent` (what kind of source a
	// question _needs_) — the two concepts overlap enough that a single enum lets
	// `classify_source_importance` compare "what we got" against "what we asked for"
	// directly.
	//
	// Any of "official", "news", "social", "academic", "aggregator", "other".
	SourceIntent string `json:"source_intent" api:"nullable"`
	// Title of the source page.
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                respjson.Field
		URL                 respjson.Field
		ExtractTemplateName respjson.Field
		SourceCategory      respjson.Field
		SourceIntent        respjson.Field
		Title               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustSource) RawJSON() string {
	return r.JSON.raw
}
func (r *AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunTextOutputPublicV2TrustSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2 struct {
	// The final structured output.
	Content AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2ContentUnion `json:"content" api:"required"`
	// Trust and citation metadata for the output.
	Trust AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2Trust `json:"trust" api:"required"`
	// Output content type.
	//
	// Any of "json".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Trust       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2) RawJSON() string {
	return r.JSON.raw
}
func (r *AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2ContentUnion
// contains all possible properties and values from [map[string]any], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfAgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2ContentMapItem
// OfAnyArray]
type AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2ContentUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfAgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2ContentMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfAgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2ContentMapItem respjson.Field
		OfAnyArray                                                                                 respjson.Field
		raw                                                                                        string
	} `json:"-"`
}

func (u AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2ContentUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2ContentUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2ContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2ContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trust and citation metadata for the output.
type AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2Trust struct {
	// Per-value trust, keyed by JSON path in the structured output.
	Claims []AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2TrustClaim `json:"claims" api:"required"`
	// Overall confidence in the answer.
	//
	// Any of "high", "medium", "low", "pre_existing".
	Confidence string `json:"confidence" api:"required"`
	// Why this confidence level was assigned.
	Reasoning string `json:"reasoning" api:"required"`
	// Sources consulted while producing the answer.
	Sources []AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2TrustSource `json:"sources" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Claims      respjson.Field
		Confidence  respjson.Field
		Reasoning   respjson.Field
		Sources     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2Trust) RawJSON() string {
	return r.JSON.raw
}
func (r *AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2Trust) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trust metadata for one value in a structured (JSON) answer, keyed by JSON path.
type AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2TrustClaim struct {
	// Citations backing this value.
	Citations []AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2TrustClaimCitation `json:"citations" api:"required"`
	// Confidence in this value.
	//
	// Any of "high", "medium", "low", "pre_existing".
	Confidence string `json:"confidence" api:"required"`
	// JSON path of the value in the structured output this claim refers to.
	Path string `json:"path" api:"required"`
	// Why this confidence level was assigned.
	Reasoning string `json:"reasoning" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Citations   respjson.Field
		Confidence  respjson.Field
		Path        respjson.Field
		Reasoning   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2TrustClaim) RawJSON() string {
	return r.JSON.raw
}
func (r *AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2TrustClaim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A citation backing a specific claim in the answer.
type AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2TrustClaimCitation struct {
	// URL of the cited page.
	URL string `json:"url" api:"required"`
	// Verbatim excerpts supporting the claim.
	Excerpts []string `json:"excerpts" api:"nullable"`
	// Extract template used to read the source, when one was used.
	ExtractTemplateName string `json:"extract_template_name" api:"nullable"`
	// What _kind_ of source this is (classified by the compress LLM), independent of
	// TrustSourceType (how authoritative it is for a specific claim). Deliberately
	// uses "official" rather than "primary" so the two axes can never collide.
	//
	// Also doubles as the sub-question's `source_intent` (what kind of source a
	// question _needs_) — the two concepts overlap enough that a single enum lets
	// `classify_source_importance` compare "what we got" against "what we asked for"
	// directly.
	//
	// Any of "official", "news", "social", "academic", "aggregator", "other".
	SourceCategory string `json:"source_category" api:"nullable"`
	// What _kind_ of source this is (classified by the compress LLM), independent of
	// TrustSourceType (how authoritative it is for a specific claim). Deliberately
	// uses "official" rather than "primary" so the two axes can never collide.
	//
	// Also doubles as the sub-question's `source_intent` (what kind of source a
	// question _needs_) — the two concepts overlap enough that a single enum lets
	// `classify_source_importance` compare "what we got" against "what we asked for"
	// directly.
	//
	// Any of "official", "news", "social", "academic", "aggregator", "other".
	SourceIntent string `json:"source_intent" api:"nullable"`
	// How authoritative the source is: 'primary' or 'secondary'.
	//
	// Any of "primary", "secondary".
	SourceType string `json:"source_type" api:"nullable"`
	// Title of the cited page.
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL                 respjson.Field
		Excerpts            respjson.Field
		ExtractTemplateName respjson.Field
		SourceCategory      respjson.Field
		SourceIntent        respjson.Field
		SourceType          respjson.Field
		Title               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2TrustClaimCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2TrustClaimCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A source consulted while producing the answer.
type AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2TrustSource struct {
	// How authoritative the source is: 'primary' or 'secondary'.
	//
	// Any of "primary", "secondary".
	Type string `json:"type" api:"required"`
	// URL of the source page.
	URL string `json:"url" api:"required"`
	// Extract template used to read the source, when one was used.
	ExtractTemplateName string `json:"extract_template_name" api:"nullable"`
	// What _kind_ of source this is (classified by the compress LLM), independent of
	// TrustSourceType (how authoritative it is for a specific claim). Deliberately
	// uses "official" rather than "primary" so the two axes can never collide.
	//
	// Also doubles as the sub-question's `source_intent` (what kind of source a
	// question _needs_) — the two concepts overlap enough that a single enum lets
	// `classify_source_importance` compare "what we got" against "what we asked for"
	// directly.
	//
	// Any of "official", "news", "social", "academic", "aggregator", "other".
	SourceCategory string `json:"source_category" api:"nullable"`
	// What _kind_ of source this is (classified by the compress LLM), independent of
	// TrustSourceType (how authoritative it is for a specific claim). Deliberately
	// uses "official" rather than "primary" so the two axes can never collide.
	//
	// Also doubles as the sub-question's `source_intent` (what kind of source a
	// question _needs_) — the two concepts overlap enough that a single enum lets
	// `classify_source_importance` compare "what we got" against "what we asked for"
	// directly.
	//
	// Any of "official", "news", "social", "academic", "aggregator", "other".
	SourceIntent string `json:"source_intent" api:"nullable"`
	// Title of the source page.
	Title string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                respjson.Field
		URL                 respjson.Field
		ExtractTemplateName respjson.Field
		SourceCategory      respjson.Field
		SourceIntent        respjson.Field
		Title               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2TrustSource) RawJSON() string {
	return r.JSON.raw
}
func (r *AgentRunResultResponseTaskRunResultPublicV2OutputTaskRunJsonOutputPublicV2TrustSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Task run object with status 'completed'.
type AgentRunResultResponseTaskRunResultPublicV2Run struct {
	// Run identifier, format "task*run*{uuid}".
	ID string `json:"id" api:"required"`
	// When the run was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Effort level used for the run.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort string `json:"effort" api:"required"`
	// Interaction ID.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Current run status.
	//
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status string `json:"status" api:"required"`
	// Web Search Agent instance this run belongs to.
	WebSearchAgentID string `json:"web_search_agent_id" api:"required"`
	// When the run completed.
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Error details when the run failed.
	Error AgentRunResultResponseTaskRunResultPublicV2RunError `json:"error" api:"nullable"`
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
func (r AgentRunResultResponseTaskRunResultPublicV2Run) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResultResponseTaskRunResultPublicV2Run) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details when the run failed.
type AgentRunResultResponseTaskRunResultPublicV2RunError struct {
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
func (r AgentRunResultResponseTaskRunResultPublicV2RunError) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResultResponseTaskRunResultPublicV2RunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResultResponseTaskRunFailedResultPublicV2 struct {
	// Structured error detail.
	Error AgentRunResultResponseTaskRunFailedResultPublicV2Error `json:"error" api:"required"`
	// Task run object with status 'failed'.
	Run AgentRunResultResponseTaskRunFailedResultPublicV2Run `json:"run" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Run         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResultResponseTaskRunFailedResultPublicV2) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResultResponseTaskRunFailedResultPublicV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured error detail.
type AgentRunResultResponseTaskRunFailedResultPublicV2Error struct {
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
func (r AgentRunResultResponseTaskRunFailedResultPublicV2Error) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResultResponseTaskRunFailedResultPublicV2Error) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Task run object with status 'failed'.
type AgentRunResultResponseTaskRunFailedResultPublicV2Run struct {
	// Run identifier, format "task*run*{uuid}".
	ID string `json:"id" api:"required"`
	// When the run was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Effort level used for the run.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort string `json:"effort" api:"required"`
	// Interaction ID.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Current run status.
	//
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status string `json:"status" api:"required"`
	// Web Search Agent instance this run belongs to.
	WebSearchAgentID string `json:"web_search_agent_id" api:"required"`
	// When the run completed.
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Error details when the run failed.
	Error AgentRunResultResponseTaskRunFailedResultPublicV2RunError `json:"error" api:"nullable"`
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
func (r AgentRunResultResponseTaskRunFailedResultPublicV2Run) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResultResponseTaskRunFailedResultPublicV2Run) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details when the run failed.
type AgentRunResultResponseTaskRunFailedResultPublicV2RunError struct {
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
func (r AgentRunResultResponseTaskRunFailedResultPublicV2RunError) RawJSON() string {
	return r.JSON.raw
}
func (r *AgentRunResultResponseTaskRunFailedResultPublicV2RunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunNewParams struct {
	// User prompt or task instructions for the run.
	Input string `json:"input" api:"required"`
	// Previous interaction identifier used to continue a conversation.
	PreviousInteractionID param.Opt[string] `json:"previous_interaction_id,omitzero"`
	// Whether to stream run events when supported.
	EnableEvents param.Opt[bool] `json:"enable_events,omitzero"`
	// Canonical effort tier names for the research graph.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort AgentRunNewParamsEffort `json:"effort,omitzero"`
	// Existing records to ENRICH: a list of partial rows, or a single object,
	// mirroring output_schema's shape.
	InputData AgentRunNewParamsInputDataUnion `json:"input_data,omitzero"`
	// JSON schema overriding the agent's default structured output for this run.
	OutputSchema map[string]any `json:"output_schema,omitzero"`
	// Source guidance overriding the agent default.
	Sources AgentRunNewParamsSources `json:"sources,omitzero"`
	paramObj
}

func (r AgentRunNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AgentRunNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRunNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Canonical effort tier names for the research graph.
type AgentRunNewParamsEffort string

const (
	AgentRunNewParamsEffortLow    AgentRunNewParamsEffort = "low"
	AgentRunNewParamsEffortMedium AgentRunNewParamsEffort = "medium"
	AgentRunNewParamsEffortHigh   AgentRunNewParamsEffort = "high"
	AgentRunNewParamsEffortXHigh  AgentRunNewParamsEffort = "x-high"
	AgentRunNewParamsEffortMax    AgentRunNewParamsEffort = "max"
)

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentRunNewParamsInputDataUnion struct {
	OfMapOfAnyMap []map[string]any `json:",omitzero,inline"`
	OfAnyMap      map[string]any   `json:",omitzero,inline"`
	paramUnion
}

func (u AgentRunNewParamsInputDataUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMapOfAnyMap, u.OfAnyMap)
}
func (u *AgentRunNewParamsInputDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentRunNewParamsInputDataUnion) asAny() any {
	if !param.IsOmitted(u.OfMapOfAnyMap) {
		return &u.OfMapOfAnyMap
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Source guidance overriding the agent default.
type AgentRunNewParamsSources struct {
	// Free-text guidance describing sources or domains to avoid.
	Avoid param.Opt[string] `json:"avoid,omitzero"`
	// Free-text guidance describing sources or domains to prioritize.
	Prioritize param.Opt[string] `json:"prioritize,omitzero"`
	// Source groups the agent is allowed to use.
	Allow []AgentRunNewParamsSourcesAllow `json:"allow,omitzero"`
	// Source groups the agent should not use.
	Block []AgentRunNewParamsSourcesBlock `json:"block,omitzero"`
	paramObj
}

func (r AgentRunNewParamsSources) MarshalJSON() (data []byte, err error) {
	type shadow AgentRunNewParamsSources
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRunNewParamsSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Domains, Title are required.
type AgentRunNewParamsSourcesAllow struct {
	// Domains included in this source group.
	Domains []string `json:"domains,omitzero" api:"required"`
	// Source group title.
	Title string `json:"title" api:"required"`
	// Zero-based source group position.
	Order param.Opt[int64] `json:"order,omitzero"`
	paramObj
}

func (r AgentRunNewParamsSourcesAllow) MarshalJSON() (data []byte, err error) {
	type shadow AgentRunNewParamsSourcesAllow
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRunNewParamsSourcesAllow) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Domains, Title are required.
type AgentRunNewParamsSourcesBlock struct {
	// Domains included in this source group.
	Domains []string `json:"domains,omitzero" api:"required"`
	// Source group title.
	Title string `json:"title" api:"required"`
	// Zero-based source group position.
	Order param.Opt[int64] `json:"order,omitzero"`
	paramObj
}

func (r AgentRunNewParamsSourcesBlock) MarshalJSON() (data []byte, err error) {
	type shadow AgentRunNewParamsSourcesBlock
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRunNewParamsSourcesBlock) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunListParams struct {
	Limit  param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentRunListParams]'s query parameters as `url.Values`.
func (r AgentRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentRunGetParams struct {
	AgentID string `path:"agent_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type AgentRunResultParams struct {
	AgentID string `path:"agent_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type AgentRunStreamEventsParams struct {
	AgentID string `path:"agent_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
