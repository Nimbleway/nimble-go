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

// TaskAgentRunService contains methods and other services that help with
// interacting with the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskAgentRunService] method instead.
type TaskAgentRunService struct {
	Options []option.RequestOption
}

// NewTaskAgentRunService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTaskAgentRunService(opts ...option.RequestOption) (r TaskAgentRunService) {
	r = TaskAgentRunService{}
	r.Options = opts
	return
}

// List runs for this instance.
//
// `status` accepts a lowercase `TaskRunStatusValue` (e.g. "completed") or a
// comma-separated list of them (e.g. "queued,running").
func (r *TaskAgentRunService) List(ctx context.Context, agentID string, query TaskAgentRunListParams, opts ...option.RequestOption) (res *TaskAgentRunListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentID == "" {
		err = errors.New("missing required agent_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/task-agents/%s/runs", agentID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Cancel an in-progress or queued run.
//
// Verb is POST + `/cancel` action segment per the AGENTS-1666 spec (replaces the
// old `DELETE …/runs/{run_id}`).
func (r *TaskAgentRunService) Cancel(ctx context.Context, runID string, body TaskAgentRunCancelParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return err
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/task-agents/%s/runs/%s/cancel", body.AgentID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Fetch a run by id, scoped to the instance.
//
// A run resolves only when (run_id, agent_id) match — otherwise 404. This means a
// stale URL with a swapped agent_id won't leak runs across instances even if the
// run_id is real.
func (r *TaskAgentRunService) Get(ctx context.Context, runID string, query TaskAgentRunGetParams, opts ...option.RequestOption) (res *TaskAgentRunGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return nil, err
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/task-agents/%s/runs/%s", query.AgentID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Fetch the result for a terminal run on this instance.
//
// Mirrors the previous flat `GET /tasks/runs/:run_id/result` semantics:
//
// - 404 when the run doesn't belong to the agent.
// - 408 when the run is still active.
// - 422 (with TaskRunFailedResult body) when the run failed or was cancelled.
// - 200 (with TaskRunResult body) on success.
func (r *TaskAgentRunService) GetResult(ctx context.Context, runID string, query TaskAgentRunGetResultParams, opts ...option.RequestOption) (res *TaskAgentRunGetResultResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return nil, err
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/task-agents/%s/runs/%s/result", query.AgentID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// SSE stream of real-time progress events for a run on this instance.
func (r *TaskAgentRunService) StreamEvents(ctx context.Context, runID string, query TaskAgentRunStreamEventsParams, opts ...option.RequestOption) (res *TaskAgentRunStreamEventsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AgentID == "" {
		err = errors.New("missing required agent_id parameter")
		return nil, err
	}
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/task-agents/%s/runs/%s/events", query.AgentID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Paginated list of task runs for GET /tasks/runs.
type TaskAgentRunListResponse struct {
	Items  []TaskAgentRunListResponseItem `json:"items" api:"required"`
	Total  int64                          `json:"total" api:"required"`
	Limit  int64                          `json:"limit"`
	Offset int64                          `json:"offset"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Total       respjson.Field
		Limit       respjson.Field
		Offset      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunListResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Task run status returned by list/create/get endpoints.
type TaskAgentRunListResponseItem struct {
	// Run identifier, format "task*run*{uuid}".
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Canonical effort tier names for the research graph.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort string `json:"effort" api:"required"`
	// Interaction ID — pass as previous_interaction_id to reuse context.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Lowercase status values used in API responses (distinct from the DB-level
	// TaskRunStatus enum).
	//
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status string `json:"status" api:"required"`
	// Web Search Agent instance this run belongs to. Every task run is agent-bound
	// (see AGENTS-1666). Use this to build the nested URL
	// /api/v2/web-search-agents/{web_search_agent_id}/runs/{id}.
	WebSearchAgentID string    `json:"web_search_agent_id" api:"required"`
	CompletedAt      time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Error detail for a failed run.
	Error TaskAgentRunListResponseItemError `json:"error" api:"nullable"`
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
func (r TaskAgentRunListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error detail for a failed run.
type TaskAgentRunListResponseItemError struct {
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
func (r TaskAgentRunListResponseItemError) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunListResponseItemError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Task run status returned by list/create/get endpoints.
type TaskAgentRunGetResponse struct {
	// Run identifier, format "task*run*{uuid}".
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Canonical effort tier names for the research graph.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort TaskAgentRunGetResponseEffort `json:"effort" api:"required"`
	// Interaction ID — pass as previous_interaction_id to reuse context.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Lowercase status values used in API responses (distinct from the DB-level
	// TaskRunStatus enum).
	//
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status TaskAgentRunGetResponseStatus `json:"status" api:"required"`
	// Web Search Agent instance this run belongs to. Every task run is agent-bound
	// (see AGENTS-1666). Use this to build the nested URL
	// /api/v2/web-search-agents/{web_search_agent_id}/runs/{id}.
	WebSearchAgentID string    `json:"web_search_agent_id" api:"required"`
	CompletedAt      time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Error detail for a failed run.
	Error TaskAgentRunGetResponseError `json:"error" api:"nullable"`
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
func (r TaskAgentRunGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Canonical effort tier names for the research graph.
type TaskAgentRunGetResponseEffort string

const (
	TaskAgentRunGetResponseEffortLow    TaskAgentRunGetResponseEffort = "low"
	TaskAgentRunGetResponseEffortMedium TaskAgentRunGetResponseEffort = "medium"
	TaskAgentRunGetResponseEffortHigh   TaskAgentRunGetResponseEffort = "high"
	TaskAgentRunGetResponseEffortXHigh  TaskAgentRunGetResponseEffort = "x-high"
	TaskAgentRunGetResponseEffortMax    TaskAgentRunGetResponseEffort = "max"
)

// Lowercase status values used in API responses (distinct from the DB-level
// TaskRunStatus enum).
type TaskAgentRunGetResponseStatus string

const (
	TaskAgentRunGetResponseStatusQueued    TaskAgentRunGetResponseStatus = "queued"
	TaskAgentRunGetResponseStatusRunning   TaskAgentRunGetResponseStatus = "running"
	TaskAgentRunGetResponseStatusCompleted TaskAgentRunGetResponseStatus = "completed"
	TaskAgentRunGetResponseStatusFailed    TaskAgentRunGetResponseStatus = "failed"
	TaskAgentRunGetResponseStatusCancelled TaskAgentRunGetResponseStatus = "cancelled"
)

// Error detail for a failed run.
type TaskAgentRunGetResponseError struct {
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
func (r TaskAgentRunGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseUnion contains all possible properties and values
// from [TaskAgentRunGetResultResponseTaskRunResult],
// [TaskAgentRunGetResultResponseTaskRunFailedResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TaskAgentRunGetResultResponseUnion struct {
	// This field is from variant [TaskAgentRunGetResultResponseTaskRunResult].
	Output TaskAgentRunGetResultResponseTaskRunResultOutputUnion `json:"output"`
	// This field is a union of [TaskAgentRunGetResultResponseTaskRunResultRun],
	// [TaskAgentRunGetResultResponseTaskRunFailedResultRun]
	Run TaskAgentRunGetResultResponseUnionRun `json:"run"`
	// This field is from variant [TaskAgentRunGetResultResponseTaskRunFailedResult].
	Error TaskAgentRunGetResultResponseTaskRunFailedResultError `json:"error"`
	JSON  struct {
		Output respjson.Field
		Run    respjson.Field
		Error  respjson.Field
		raw    string
	} `json:"-"`
}

func (u TaskAgentRunGetResultResponseUnion) AsTaskRunResult() (v TaskAgentRunGetResultResponseTaskRunResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskAgentRunGetResultResponseUnion) AsTaskRunFailedResult() (v TaskAgentRunGetResultResponseTaskRunFailedResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TaskAgentRunGetResultResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *TaskAgentRunGetResultResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseUnionRun is an implicit subunion of
// [TaskAgentRunGetResultResponseUnion]. TaskAgentRunGetResultResponseUnionRun
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskAgentRunGetResultResponseUnion].
type TaskAgentRunGetResultResponseUnionRun struct {
	ID               string    `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	Effort           string    `json:"effort"`
	InteractionID    string    `json:"interaction_id"`
	IsActive         bool      `json:"is_active"`
	Status           string    `json:"status"`
	WebSearchAgentID string    `json:"web_search_agent_id"`
	CompletedAt      time.Time `json:"completed_at"`
	// This field is a union of [TaskAgentRunGetResultResponseTaskRunResultRunError],
	// [TaskAgentRunGetResultResponseTaskRunFailedResultRunError]
	Error       TaskAgentRunGetResultResponseUnionRunError `json:"error"`
	Prompt      string                                     `json:"prompt"`
	StartedAt   time.Time                                  `json:"started_at"`
	WorkspaceID string                                     `json:"workspace_id"`
	JSON        struct {
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
		raw              string
	} `json:"-"`
}

func (r *TaskAgentRunGetResultResponseUnionRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseUnionRunError is an implicit subunion of
// [TaskAgentRunGetResultResponseUnion]. TaskAgentRunGetResultResponseUnionRunError
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskAgentRunGetResultResponseUnion].
type TaskAgentRunGetResultResponseUnionRunError struct {
	Message string `json:"message"`
	RefID   string `json:"ref_id"`
	JSON    struct {
		Message respjson.Field
		RefID   respjson.Field
		raw     string
	} `json:"-"`
}

func (r *TaskAgentRunGetResultResponseUnionRunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for GET /tasks/runs/{run_id}/result — status 'completed'.
type TaskAgentRunGetResultResponseTaskRunResult struct {
	// Output from the completed task.
	Output TaskAgentRunGetResultResponseTaskRunResultOutputUnion `json:"output" api:"required"`
	// Task run object with status 'completed'.
	Run TaskAgentRunGetResultResponseTaskRunResultRun `json:"run" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		Run         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunGetResultResponseTaskRunResult) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResultResponseTaskRunResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseTaskRunResultOutputUnion contains all possible
// properties and values from
// [TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutput],
// [TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutput].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TaskAgentRunGetResultResponseTaskRunResultOutputUnion struct {
	// This field is a union of [string],
	// [TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputContentUnion]
	Content TaskAgentRunGetResultResponseTaskRunResultOutputUnionContent `json:"content"`
	// This field is a union of
	// [TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrust],
	// [TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrust]
	Trust TaskAgentRunGetResultResponseTaskRunResultOutputUnionTrust `json:"trust"`
	Type  string                                                     `json:"type"`
	JSON  struct {
		Content respjson.Field
		Trust   respjson.Field
		Type    respjson.Field
		raw     string
	} `json:"-"`
}

func (u TaskAgentRunGetResultResponseTaskRunResultOutputUnion) AsTaskRunTextOutput() (v TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskAgentRunGetResultResponseTaskRunResultOutputUnion) AsTaskRunJsonOutput() (v TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TaskAgentRunGetResultResponseTaskRunResultOutputUnion) RawJSON() string { return u.JSON.raw }

func (r *TaskAgentRunGetResultResponseTaskRunResultOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseTaskRunResultOutputUnionContent is an implicit
// subunion of [TaskAgentRunGetResultResponseTaskRunResultOutputUnion].
// TaskAgentRunGetResultResponseTaskRunResultOutputUnionContent provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskAgentRunGetResultResponseTaskRunResultOutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfTaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputContentMapItem
// OfAnyArray]
type TaskAgentRunGetResultResponseTaskRunResultOutputUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputContentMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                          respjson.Field
		OfTaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputContentMapItem respjson.Field
		OfAnyArray                                                                        respjson.Field
		raw                                                                               string
	} `json:"-"`
}

func (r *TaskAgentRunGetResultResponseTaskRunResultOutputUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseTaskRunResultOutputUnionTrust is an implicit
// subunion of [TaskAgentRunGetResultResponseTaskRunResultOutputUnion].
// TaskAgentRunGetResultResponseTaskRunResultOutputUnionTrust provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskAgentRunGetResultResponseTaskRunResultOutputUnion].
type TaskAgentRunGetResultResponseTaskRunResultOutputUnionTrust struct {
	// This field is a union of
	// [[]TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustClaim],
	// [[]TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrustClaim]
	Claims     TaskAgentRunGetResultResponseTaskRunResultOutputUnionTrustClaims `json:"claims"`
	Confidence string                                                           `json:"confidence"`
	Reasoning  string                                                           `json:"reasoning"`
	// This field is a union of
	// [[]TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustSource],
	// [[]TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrustSource]
	Sources TaskAgentRunGetResultResponseTaskRunResultOutputUnionTrustSources `json:"sources"`
	JSON    struct {
		Claims     respjson.Field
		Confidence respjson.Field
		Reasoning  respjson.Field
		Sources    respjson.Field
		raw        string
	} `json:"-"`
}

func (r *TaskAgentRunGetResultResponseTaskRunResultOutputUnionTrust) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseTaskRunResultOutputUnionTrustClaims is an implicit
// subunion of [TaskAgentRunGetResultResponseTaskRunResultOutputUnion].
// TaskAgentRunGetResultResponseTaskRunResultOutputUnionTrustClaims provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskAgentRunGetResultResponseTaskRunResultOutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfClaims]
type TaskAgentRunGetResultResponseTaskRunResultOutputUnionTrustClaims struct {
	// This field will be present if the value is a
	// [[]TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustClaim]
	// instead of an object.
	OfClaims []TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustClaim `json:",inline"`
	JSON     struct {
		OfClaims respjson.Field
		raw      string
	} `json:"-"`
}

func (r *TaskAgentRunGetResultResponseTaskRunResultOutputUnionTrustClaims) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseTaskRunResultOutputUnionTrustSources is an implicit
// subunion of [TaskAgentRunGetResultResponseTaskRunResultOutputUnion].
// TaskAgentRunGetResultResponseTaskRunResultOutputUnionTrustSources provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskAgentRunGetResultResponseTaskRunResultOutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfSources]
type TaskAgentRunGetResultResponseTaskRunResultOutputUnionTrustSources struct {
	// This field will be present if the value is a
	// [[]TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustSource]
	// instead of an object.
	OfSources []TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustSource `json:",inline"`
	JSON      struct {
		OfSources respjson.Field
		raw       string
	} `json:"-"`
}

func (r *TaskAgentRunGetResultResponseTaskRunResultOutputUnionTrustSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Text output from a completed task.
type TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutput struct {
	// The final prose answer.
	Content string                                                                 `json:"content" api:"required"`
	Trust   TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrust `json:"trust" api:"required"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutput) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrust struct {
	Claims []TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustClaim `json:"claims" api:"required"`
	// Any of "high", "medium", "low".
	Confidence string                                                                         `json:"confidence" api:"required"`
	Reasoning  string                                                                         `json:"reasoning" api:"required"`
	Sources    []TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustSource `json:"sources" api:"required"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrust) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrust) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustClaim struct {
	Callout   int64                                                                                 `json:"callout" api:"required"`
	Citations []TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustClaimCitation `json:"citations" api:"required"`
	// Any of "high", "medium", "low".
	Confidence string `json:"confidence" api:"required"`
	Reasoning  string `json:"reasoning" api:"required"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustClaim) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustClaim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustClaimCitation struct {
	URL                 string   `json:"url" api:"required"`
	Excerpts            []string `json:"excerpts" api:"nullable"`
	ExtractTemplateName string   `json:"extract_template_name" api:"nullable"`
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
	// Any of "primary", "secondary".
	SourceType string `json:"source_type" api:"nullable"`
	Title      string `json:"title" api:"nullable"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustClaimCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustClaimCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustSource struct {
	// Any of "primary", "secondary".
	Type                string `json:"type" api:"required"`
	URL                 string `json:"url" api:"required"`
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
	Title        string `json:"title" api:"nullable"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustSource) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunTextOutputTrustSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured JSON output from a completed task, produced when
// task_spec.output_schema.type is 'json'.
type TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutput struct {
	// Data conforming to the caller-supplied JSON schema. A dict for object schemas; a
	// list for array schemas.
	Content TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputContentUnion `json:"content" api:"required"`
	Trust   TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrust        `json:"trust" api:"required"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutput) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputContentUnion
// contains all possible properties and values from [map[string]any], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfTaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputContentMapItem
// OfAnyArray]
type TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputContentUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfTaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputContentMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfTaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputContentMapItem respjson.Field
		OfAnyArray                                                                        respjson.Field
		raw                                                                               string
	} `json:"-"`
}

func (u TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputContentUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputContentUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrust struct {
	Claims []TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrustClaim `json:"claims" api:"required"`
	// Any of "high", "medium", "low".
	Confidence string                                                                         `json:"confidence" api:"required"`
	Reasoning  string                                                                         `json:"reasoning" api:"required"`
	Sources    []TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrustSource `json:"sources" api:"required"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrust) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrust) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrustClaim struct {
	Citations []TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrustClaimCitation `json:"citations" api:"required"`
	// Any of "high", "medium", "low".
	Confidence string `json:"confidence" api:"required"`
	Path       string `json:"path" api:"required"`
	Reasoning  string `json:"reasoning" api:"required"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrustClaim) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrustClaim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrustClaimCitation struct {
	URL                 string   `json:"url" api:"required"`
	Excerpts            []string `json:"excerpts" api:"nullable"`
	ExtractTemplateName string   `json:"extract_template_name" api:"nullable"`
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
	// Any of "primary", "secondary".
	SourceType string `json:"source_type" api:"nullable"`
	Title      string `json:"title" api:"nullable"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrustClaimCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrustClaimCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrustSource struct {
	// Any of "primary", "secondary".
	Type                string `json:"type" api:"required"`
	URL                 string `json:"url" api:"required"`
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
	Title        string `json:"title" api:"nullable"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrustSource) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultOutputTaskRunJsonOutputTrustSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Task run object with status 'completed'.
type TaskAgentRunGetResultResponseTaskRunResultRun struct {
	// Run identifier, format "task*run*{uuid}".
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Canonical effort tier names for the research graph.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort string `json:"effort" api:"required"`
	// Interaction ID — pass as previous_interaction_id to reuse context.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Lowercase status values used in API responses (distinct from the DB-level
	// TaskRunStatus enum).
	//
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status string `json:"status" api:"required"`
	// Web Search Agent instance this run belongs to. Every task run is agent-bound
	// (see AGENTS-1666). Use this to build the nested URL
	// /api/v2/web-search-agents/{web_search_agent_id}/runs/{id}.
	WebSearchAgentID string    `json:"web_search_agent_id" api:"required"`
	CompletedAt      time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Error detail for a failed run.
	Error TaskAgentRunGetResultResponseTaskRunResultRunError `json:"error" api:"nullable"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultRun) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResultResponseTaskRunResultRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error detail for a failed run.
type TaskAgentRunGetResultResponseTaskRunResultRunError struct {
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
func (r TaskAgentRunGetResultResponseTaskRunResultRunError) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResultResponseTaskRunResultRunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response for GET /tasks/runs/{run_id}/result when the run failed.
//
// Returned with HTTP 422 so callers can distinguish a failed run from a missing
// one (404) or an active one (408).
type TaskAgentRunGetResultResponseTaskRunFailedResult struct {
	// Structured error detail.
	Error TaskAgentRunGetResultResponseTaskRunFailedResultError `json:"error" api:"required"`
	// Task run object with status 'failed'.
	Run TaskAgentRunGetResultResponseTaskRunFailedResultRun `json:"run" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Run         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunGetResultResponseTaskRunFailedResult) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResultResponseTaskRunFailedResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured error detail.
type TaskAgentRunGetResultResponseTaskRunFailedResultError struct {
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
func (r TaskAgentRunGetResultResponseTaskRunFailedResultError) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResultResponseTaskRunFailedResultError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Task run object with status 'failed'.
type TaskAgentRunGetResultResponseTaskRunFailedResultRun struct {
	// Run identifier, format "task*run*{uuid}".
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Canonical effort tier names for the research graph.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort string `json:"effort" api:"required"`
	// Interaction ID — pass as previous_interaction_id to reuse context.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Lowercase status values used in API responses (distinct from the DB-level
	// TaskRunStatus enum).
	//
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status string `json:"status" api:"required"`
	// Web Search Agent instance this run belongs to. Every task run is agent-bound
	// (see AGENTS-1666). Use this to build the nested URL
	// /api/v2/web-search-agents/{web_search_agent_id}/runs/{id}.
	WebSearchAgentID string    `json:"web_search_agent_id" api:"required"`
	CompletedAt      time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Error detail for a failed run.
	Error TaskAgentRunGetResultResponseTaskRunFailedResultRunError `json:"error" api:"nullable"`
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
func (r TaskAgentRunGetResultResponseTaskRunFailedResultRun) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResultResponseTaskRunFailedResultRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error detail for a failed run.
type TaskAgentRunGetResultResponseTaskRunFailedResultRunError struct {
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
func (r TaskAgentRunGetResultResponseTaskRunFailedResultRunError) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResultResponseTaskRunFailedResultRunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunStreamEventsResponse = any

type TaskAgentRunListParams struct {
	Q      param.Opt[string] `query:"q,omitzero" json:"-"`
	Status param.Opt[string] `query:"status,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	Offset param.Opt[int64]  `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaskAgentRunListParams]'s query parameters as `url.Values`.
func (r TaskAgentRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TaskAgentRunCancelParams struct {
	AgentID string `path:"agent_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type TaskAgentRunGetParams struct {
	AgentID string `path:"agent_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type TaskAgentRunGetResultParams struct {
	AgentID string `path:"agent_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type TaskAgentRunStreamEventsParams struct {
	AgentID string `path:"agent_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
