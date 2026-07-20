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
//
// Deprecated: deprecated
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

// Fetch a run by id, scoped to the instance.
//
// A run resolves only when (run_id, agent_id) match — otherwise 404. This means a
// stale URL with a swapped agent_id won't leak runs across instances even if the
// run_id is real.
//
// Deprecated: deprecated
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
//
// Deprecated: deprecated
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
//
// Deprecated: deprecated
func (r *TaskAgentRunService) StreamEvents(ctx context.Context, runID string, query TaskAgentRunStreamEventsParams, opts ...option.RequestOption) (err error) {
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
	path := fmt.Sprintf("v1/task-agents/%s/runs/%s/events", query.AgentID, runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

type TaskAgentRunListResponse struct {
	// Items returned in this page.
	Items []TaskAgentRunListResponseItem `json:"items" api:"required"`
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
func (r TaskAgentRunListResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunListResponseItem struct {
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
	Error TaskAgentRunListResponseItemError `json:"error" api:"nullable"`
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
func (r TaskAgentRunListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details when the run failed.
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

type TaskAgentRunGetResponse struct {
	// Run identifier, format "task*run*{uuid}".
	ID string `json:"id" api:"required"`
	// When the run was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Effort level used for the run.
	//
	// Any of "low", "medium", "high", "x-high", "max".
	Effort TaskAgentRunGetResponseEffort `json:"effort" api:"required"`
	// Interaction ID.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Current run status.
	//
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status TaskAgentRunGetResponseStatus `json:"status" api:"required"`
	// Web Search Agent instance this run belongs to.
	WebSearchAgentID string `json:"web_search_agent_id" api:"required"`
	// When the run completed.
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Error details when the run failed.
	Error TaskAgentRunGetResponseError `json:"error" api:"nullable"`
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
func (r TaskAgentRunGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Effort level used for the run.
type TaskAgentRunGetResponseEffort string

const (
	TaskAgentRunGetResponseEffortLow    TaskAgentRunGetResponseEffort = "low"
	TaskAgentRunGetResponseEffortMedium TaskAgentRunGetResponseEffort = "medium"
	TaskAgentRunGetResponseEffortHigh   TaskAgentRunGetResponseEffort = "high"
	TaskAgentRunGetResponseEffortXHigh  TaskAgentRunGetResponseEffort = "x-high"
	TaskAgentRunGetResponseEffortMax    TaskAgentRunGetResponseEffort = "max"
)

// Current run status.
type TaskAgentRunGetResponseStatus string

const (
	TaskAgentRunGetResponseStatusQueued    TaskAgentRunGetResponseStatus = "queued"
	TaskAgentRunGetResponseStatusRunning   TaskAgentRunGetResponseStatus = "running"
	TaskAgentRunGetResponseStatusCompleted TaskAgentRunGetResponseStatus = "completed"
	TaskAgentRunGetResponseStatusFailed    TaskAgentRunGetResponseStatus = "failed"
	TaskAgentRunGetResponseStatusCancelled TaskAgentRunGetResponseStatus = "cancelled"
)

// Error details when the run failed.
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
// from [TaskAgentRunGetResultResponseTaskRunResultPublicV1],
// [TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TaskAgentRunGetResultResponseUnion struct {
	// This field is from variant [TaskAgentRunGetResultResponseTaskRunResultPublicV1].
	Output TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnion `json:"output"`
	// This field is a union of
	// [TaskAgentRunGetResultResponseTaskRunResultPublicV1Run],
	// [TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1Run]
	Run TaskAgentRunGetResultResponseUnionRun `json:"run"`
	// This field is from variant
	// [TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1].
	Error TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1Error `json:"error"`
	JSON  struct {
		Output respjson.Field
		Run    respjson.Field
		Error  respjson.Field
		raw    string
	} `json:"-"`
}

func (u TaskAgentRunGetResultResponseUnion) AsTaskRunResultPublicV1() (v TaskAgentRunGetResultResponseTaskRunResultPublicV1) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskAgentRunGetResultResponseUnion) AsTaskRunFailedResultPublicV1() (v TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1) {
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
	// This field is a union of
	// [TaskAgentRunGetResultResponseTaskRunResultPublicV1RunError],
	// [TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1RunError]
	Error     TaskAgentRunGetResultResponseUnionRunError `json:"error"`
	Prompt    string                                     `json:"prompt"`
	StartedAt time.Time                                  `json:"started_at"`
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

type TaskAgentRunGetResultResponseTaskRunResultPublicV1 struct {
	// Output from the completed task.
	Output TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnion `json:"output" api:"required"`
	// Task run object with status 'completed'.
	Run TaskAgentRunGetResultResponseTaskRunResultPublicV1Run `json:"run" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		Run         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunGetResultResponseTaskRunResultPublicV1) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnion contains all
// possible properties and values from
// [TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1],
// [TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnion struct {
	// This field is a union of [string],
	// [TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1ContentUnion]
	Content TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionContent `json:"content"`
	// This field is a union of
	// [TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1Trust],
	// [TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1Trust]
	Trust TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionTrust `json:"trust"`
	Type  string                                                             `json:"type"`
	JSON  struct {
		Content respjson.Field
		Trust   respjson.Field
		Type    respjson.Field
		raw     string
	} `json:"-"`
}

func (u TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnion) AsTaskRunTextOutputPublicV1() (v TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnion) AsTaskRunJsonOutputPublicV1() (v TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionContent is an
// implicit subunion of
// [TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnion].
// TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionContent provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfTaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1ContentMapItem
// OfAnyArray]
type TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1ContentMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                                          respjson.Field
		OfTaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1ContentMapItem respjson.Field
		OfAnyArray                                                                                        respjson.Field
		raw                                                                                               string
	} `json:"-"`
}

func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionTrust is an
// implicit subunion of
// [TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnion].
// TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionTrust provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnion].
type TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionTrust struct {
	// This field is a union of
	// [[]TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustClaim],
	// [[]TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1TrustClaim]
	Claims     TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionTrustClaims `json:"claims"`
	Confidence string                                                                   `json:"confidence"`
	Reasoning  string                                                                   `json:"reasoning"`
	// This field is a union of
	// [[]TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustSource],
	// [[]TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1TrustSource]
	Sources TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionTrustSources `json:"sources"`
	JSON    struct {
		Claims     respjson.Field
		Confidence respjson.Field
		Reasoning  respjson.Field
		Sources    respjson.Field
		raw        string
	} `json:"-"`
}

func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionTrust) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionTrustClaims is an
// implicit subunion of
// [TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnion].
// TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionTrustClaims
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfClaims]
type TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionTrustClaims struct {
	// This field will be present if the value is a
	// [[]TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustClaim]
	// instead of an object.
	OfClaims []TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustClaim `json:",inline"`
	JSON     struct {
		OfClaims respjson.Field
		raw      string
	} `json:"-"`
}

func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionTrustClaims) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionTrustSources is an
// implicit subunion of
// [TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnion].
// TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionTrustSources
// provides convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfSources]
type TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionTrustSources struct {
	// This field will be present if the value is a
	// [[]TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustSource]
	// instead of an object.
	OfSources []TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustSource `json:",inline"`
	JSON      struct {
		OfSources respjson.Field
		raw       string
	} `json:"-"`
}

func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputUnionTrustSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1 struct {
	// The final prose answer.
	Content string `json:"content" api:"required"`
	// Trust and citation metadata for the output.
	Trust TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1Trust `json:"trust" api:"required"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trust and citation metadata for the output.
type TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1Trust struct {
	// Per-claim trust, keyed by callout markers in the answer text.
	Claims []TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustClaim `json:"claims" api:"required"`
	// Overall confidence in the answer.
	//
	// Any of "high", "medium", "low", "pre_existing".
	Confidence string `json:"confidence" api:"required"`
	// Why this confidence level was assigned.
	Reasoning string `json:"reasoning" api:"required"`
	// Sources consulted while producing the answer.
	Sources []TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustSource `json:"sources" api:"required"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1Trust) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1Trust) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trust metadata for one claim in a prose answer, keyed by callout marker.
type TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustClaim struct {
	// Callout marker number referencing this claim in the answer text.
	Callout int64 `json:"callout" api:"required"`
	// Citations backing this claim.
	Citations []TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustClaimCitation `json:"citations" api:"required"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustClaim) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustClaim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A citation backing a specific claim in the answer.
type TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustClaimCitation struct {
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
func (r TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustClaimCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustClaimCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A source consulted while producing the answer.
type TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustSource struct {
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
func (r TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustSource) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunTextOutputPublicV1TrustSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1 struct {
	// The final structured output.
	Content TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1ContentUnion `json:"content" api:"required"`
	// Trust and citation metadata for the output.
	Trust TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1Trust `json:"trust" api:"required"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1ContentUnion
// contains all possible properties and values from [map[string]any], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfTaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1ContentMapItem
// OfAnyArray]
type TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1ContentUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfTaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1ContentMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfTaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1ContentMapItem respjson.Field
		OfAnyArray                                                                                        respjson.Field
		raw                                                                                               string
	} `json:"-"`
}

func (u TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1ContentUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1ContentUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1ContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1ContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trust and citation metadata for the output.
type TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1Trust struct {
	// Per-value trust, keyed by JSON path in the structured output.
	Claims []TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1TrustClaim `json:"claims" api:"required"`
	// Overall confidence in the answer.
	//
	// Any of "high", "medium", "low", "pre_existing".
	Confidence string `json:"confidence" api:"required"`
	// Why this confidence level was assigned.
	Reasoning string `json:"reasoning" api:"required"`
	// Sources consulted while producing the answer.
	Sources []TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1TrustSource `json:"sources" api:"required"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1Trust) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1Trust) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Trust metadata for one value in a structured (JSON) answer, keyed by JSON path.
type TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1TrustClaim struct {
	// Citations backing this value.
	Citations []TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1TrustClaimCitation `json:"citations" api:"required"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1TrustClaim) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1TrustClaim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A citation backing a specific claim in the answer.
type TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1TrustClaimCitation struct {
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
func (r TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1TrustClaimCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1TrustClaimCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A source consulted while producing the answer.
type TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1TrustSource struct {
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
func (r TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1TrustSource) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1OutputTaskRunJsonOutputPublicV1TrustSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Task run object with status 'completed'.
type TaskAgentRunGetResultResponseTaskRunResultPublicV1Run struct {
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
	Error TaskAgentRunGetResultResponseTaskRunResultPublicV1RunError `json:"error" api:"nullable"`
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
func (r TaskAgentRunGetResultResponseTaskRunResultPublicV1Run) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1Run) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details when the run failed.
type TaskAgentRunGetResultResponseTaskRunResultPublicV1RunError struct {
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
func (r TaskAgentRunGetResultResponseTaskRunResultPublicV1RunError) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunResultPublicV1RunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1 struct {
	// Structured error detail.
	Error TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1Error `json:"error" api:"required"`
	// Task run object with status 'failed'.
	Run TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1Run `json:"run" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Run         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Structured error detail.
type TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1Error struct {
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
func (r TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1Error) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1Error) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Task run object with status 'failed'.
type TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1Run struct {
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
	Error TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1RunError `json:"error" api:"nullable"`
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
func (r TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1Run) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1Run) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Error details when the run failed.
type TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1RunError struct {
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
func (r TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1RunError) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseTaskRunFailedResultPublicV1RunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunListParams struct {
	Limit  param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaskAgentRunListParams]'s query parameters as `url.Values`.
func (r TaskAgentRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
