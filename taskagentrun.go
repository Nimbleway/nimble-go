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

// List task runs for the caller's workspace and the given agent, newest first.
func (r *TaskAgentRunService) List(ctx context.Context, agentID string, query TaskAgentRunListParams, opts ...option.RequestOption) (res *[]TaskAgentRunListResponse, err error) {
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

// Poll run status. Repeat until status is 'completed', 'failed', or 'cancelled'.
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

// Fetch the result for a terminal run. Returns 408 if still active, 422 with
// `AgentRunFailedResult` if failed.
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

// Server-Sent Events stream of real-time progress events for a run. The run must
// have been created with `enable_events=true`.
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

type TaskAgentRunListResponse struct {
	// Run identifier.
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "quickest", "quick", "research", "pro", "max".
	Effort TaskAgentRunListResponseEffort `json:"effort" api:"required"`
	// Interaction ID — pass as previous_interaction_id to reuse context.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status      TaskAgentRunListResponseStatus `json:"status" api:"required"`
	CompletedAt time.Time                      `json:"completed_at" api:"nullable" format:"date-time"`
	Error       TaskAgentRunListResponseError  `json:"error" api:"nullable"`
	Prompt      string                         `json:"prompt" api:"nullable"`
	StartedAt   time.Time                      `json:"started_at" api:"nullable" format:"date-time"`
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
func (r TaskAgentRunListResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunListResponseEffort string

const (
	TaskAgentRunListResponseEffortQuickest TaskAgentRunListResponseEffort = "quickest"
	TaskAgentRunListResponseEffortQuick    TaskAgentRunListResponseEffort = "quick"
	TaskAgentRunListResponseEffortResearch TaskAgentRunListResponseEffort = "research"
	TaskAgentRunListResponseEffortPro      TaskAgentRunListResponseEffort = "pro"
	TaskAgentRunListResponseEffortMax      TaskAgentRunListResponseEffort = "max"
)

type TaskAgentRunListResponseStatus string

const (
	TaskAgentRunListResponseStatusQueued    TaskAgentRunListResponseStatus = "queued"
	TaskAgentRunListResponseStatusRunning   TaskAgentRunListResponseStatus = "running"
	TaskAgentRunListResponseStatusCompleted TaskAgentRunListResponseStatus = "completed"
	TaskAgentRunListResponseStatusFailed    TaskAgentRunListResponseStatus = "failed"
	TaskAgentRunListResponseStatusCancelled TaskAgentRunListResponseStatus = "cancelled"
)

type TaskAgentRunListResponseError struct {
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
func (r TaskAgentRunListResponseError) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunListResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResponse struct {
	// Run identifier.
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "quickest", "quick", "research", "pro", "max".
	Effort TaskAgentRunGetResponseEffort `json:"effort" api:"required"`
	// Interaction ID — pass as previous_interaction_id to reuse context.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status      TaskAgentRunGetResponseStatus `json:"status" api:"required"`
	CompletedAt time.Time                     `json:"completed_at" api:"nullable" format:"date-time"`
	Error       TaskAgentRunGetResponseError  `json:"error" api:"nullable"`
	Prompt      string                        `json:"prompt" api:"nullable"`
	StartedAt   time.Time                     `json:"started_at" api:"nullable" format:"date-time"`
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
func (r TaskAgentRunGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResponseEffort string

const (
	TaskAgentRunGetResponseEffortQuickest TaskAgentRunGetResponseEffort = "quickest"
	TaskAgentRunGetResponseEffortQuick    TaskAgentRunGetResponseEffort = "quick"
	TaskAgentRunGetResponseEffortResearch TaskAgentRunGetResponseEffort = "research"
	TaskAgentRunGetResponseEffortPro      TaskAgentRunGetResponseEffort = "pro"
	TaskAgentRunGetResponseEffortMax      TaskAgentRunGetResponseEffort = "max"
)

type TaskAgentRunGetResponseStatus string

const (
	TaskAgentRunGetResponseStatusQueued    TaskAgentRunGetResponseStatus = "queued"
	TaskAgentRunGetResponseStatusRunning   TaskAgentRunGetResponseStatus = "running"
	TaskAgentRunGetResponseStatusCompleted TaskAgentRunGetResponseStatus = "completed"
	TaskAgentRunGetResponseStatusFailed    TaskAgentRunGetResponseStatus = "failed"
	TaskAgentRunGetResponseStatusCancelled TaskAgentRunGetResponseStatus = "cancelled"
)

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
// from [TaskAgentRunGetResultResponseAgentRunResult],
// [TaskAgentRunGetResultResponseAgentRunFailedResult].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TaskAgentRunGetResultResponseUnion struct {
	// This field is from variant [TaskAgentRunGetResultResponseAgentRunResult].
	Output TaskAgentRunGetResultResponseAgentRunResultOutputUnion `json:"output"`
	// This field is a union of [TaskAgentRunGetResultResponseAgentRunResultRun],
	// [TaskAgentRunGetResultResponseAgentRunFailedResultRun]
	Run TaskAgentRunGetResultResponseUnionRun `json:"run"`
	// This field is from variant [TaskAgentRunGetResultResponseAgentRunFailedResult].
	Error TaskAgentRunGetResultResponseAgentRunFailedResultError `json:"error"`
	JSON  struct {
		Output respjson.Field
		Run    respjson.Field
		Error  respjson.Field
		raw    string
	} `json:"-"`
}

func (u TaskAgentRunGetResultResponseUnion) AsAgentRunResult() (v TaskAgentRunGetResultResponseAgentRunResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskAgentRunGetResultResponseUnion) AsAgentRunFailedResult() (v TaskAgentRunGetResultResponseAgentRunFailedResult) {
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
	ID            string    `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	Effort        string    `json:"effort"`
	InteractionID string    `json:"interaction_id"`
	IsActive      bool      `json:"is_active"`
	Status        string    `json:"status"`
	CompletedAt   time.Time `json:"completed_at"`
	// This field is a union of [TaskAgentRunGetResultResponseAgentRunResultRunError],
	// [TaskAgentRunGetResultResponseAgentRunFailedResultRunError]
	Error            TaskAgentRunGetResultResponseUnionRunError `json:"error"`
	Prompt           string                                     `json:"prompt"`
	StartedAt        time.Time                                  `json:"started_at"`
	WebSearchAgentID string                                     `json:"web_search_agent_id"`
	WorkspaceID      string                                     `json:"workspace_id"`
	JSON             struct {
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

type TaskAgentRunGetResultResponseAgentRunResult struct {
	Output TaskAgentRunGetResultResponseAgentRunResultOutputUnion `json:"output" api:"required"`
	Run    TaskAgentRunGetResultResponseAgentRunResultRun         `json:"run" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Output      respjson.Field
		Run         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunGetResultResponseAgentRunResult) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResultResponseAgentRunResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseAgentRunResultOutputUnion contains all possible
// properties and values from
// [TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutput],
// [TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutput].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TaskAgentRunGetResultResponseAgentRunResultOutputUnion struct {
	// This field is a union of [string],
	// [TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputContentUnion]
	Content TaskAgentRunGetResultResponseAgentRunResultOutputUnionContent `json:"content"`
	// This field is a union of
	// [TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrust],
	// [TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrust]
	Trust TaskAgentRunGetResultResponseAgentRunResultOutputUnionTrust `json:"trust"`
	Type  string                                                      `json:"type"`
	JSON  struct {
		Content respjson.Field
		Trust   respjson.Field
		Type    respjson.Field
		raw     string
	} `json:"-"`
}

func (u TaskAgentRunGetResultResponseAgentRunResultOutputUnion) AsAgentRunTextOutput() (v TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskAgentRunGetResultResponseAgentRunResultOutputUnion) AsAgentRunJsonOutput() (v TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutput) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TaskAgentRunGetResultResponseAgentRunResultOutputUnion) RawJSON() string { return u.JSON.raw }

func (r *TaskAgentRunGetResultResponseAgentRunResultOutputUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseAgentRunResultOutputUnionContent is an implicit
// subunion of [TaskAgentRunGetResultResponseAgentRunResultOutputUnion].
// TaskAgentRunGetResultResponseAgentRunResultOutputUnionContent provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskAgentRunGetResultResponseAgentRunResultOutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString
// OfTaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputContentMapItem
// OfAnyArray]
type TaskAgentRunGetResultResponseAgentRunResultOutputUnionContent struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfTaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputContentMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                                                            respjson.Field
		OfTaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputContentMapItem respjson.Field
		OfAnyArray                                                                          respjson.Field
		raw                                                                                 string
	} `json:"-"`
}

func (r *TaskAgentRunGetResultResponseAgentRunResultOutputUnionContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseAgentRunResultOutputUnionTrust is an implicit
// subunion of [TaskAgentRunGetResultResponseAgentRunResultOutputUnion].
// TaskAgentRunGetResultResponseAgentRunResultOutputUnionTrust provides convenient
// access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskAgentRunGetResultResponseAgentRunResultOutputUnion].
type TaskAgentRunGetResultResponseAgentRunResultOutputUnionTrust struct {
	// This field is a union of
	// [[]TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustClaim],
	// [[]TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustClaim]
	Claims     TaskAgentRunGetResultResponseAgentRunResultOutputUnionTrustClaims `json:"claims"`
	Confidence string                                                            `json:"confidence"`
	Reasoning  string                                                            `json:"reasoning"`
	// This field is a union of
	// [[]TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustSource],
	// [[]TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustSource]
	Sources TaskAgentRunGetResultResponseAgentRunResultOutputUnionTrustSources `json:"sources"`
	JSON    struct {
		Claims     respjson.Field
		Confidence respjson.Field
		Reasoning  respjson.Field
		Sources    respjson.Field
		raw        string
	} `json:"-"`
}

func (r *TaskAgentRunGetResultResponseAgentRunResultOutputUnionTrust) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseAgentRunResultOutputUnionTrustClaims is an implicit
// subunion of [TaskAgentRunGetResultResponseAgentRunResultOutputUnion].
// TaskAgentRunGetResultResponseAgentRunResultOutputUnionTrustClaims provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskAgentRunGetResultResponseAgentRunResultOutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfClaims]
type TaskAgentRunGetResultResponseAgentRunResultOutputUnionTrustClaims struct {
	// This field will be present if the value is a
	// [[]TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustClaim]
	// instead of an object.
	OfClaims []TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustClaim `json:",inline"`
	JSON     struct {
		OfClaims respjson.Field
		raw      string
	} `json:"-"`
}

func (r *TaskAgentRunGetResultResponseAgentRunResultOutputUnionTrustClaims) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseAgentRunResultOutputUnionTrustSources is an
// implicit subunion of [TaskAgentRunGetResultResponseAgentRunResultOutputUnion].
// TaskAgentRunGetResultResponseAgentRunResultOutputUnionTrustSources provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [TaskAgentRunGetResultResponseAgentRunResultOutputUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfSources]
type TaskAgentRunGetResultResponseAgentRunResultOutputUnionTrustSources struct {
	// This field will be present if the value is a
	// [[]TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustSource]
	// instead of an object.
	OfSources []TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustSource `json:",inline"`
	JSON      struct {
		OfSources respjson.Field
		raw       string
	} `json:"-"`
}

func (r *TaskAgentRunGetResultResponseAgentRunResultOutputUnionTrustSources) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutput struct {
	// The final prose answer.
	Content string                                                                   `json:"content" api:"required"`
	Trust   TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrust `json:"trust" api:"required"`
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
func (r TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutput) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrust struct {
	Claims []TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustClaim `json:"claims" api:"required"`
	// Any of "high", "medium", "low".
	Confidence string                                                                           `json:"confidence" api:"required"`
	Reasoning  string                                                                           `json:"reasoning" api:"required"`
	Sources    []TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustSource `json:"sources" api:"required"`
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
func (r TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrust) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrust) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustClaim struct {
	Callout   int64                                                                                   `json:"callout" api:"required"`
	Citations []TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustClaimCitation `json:"citations" api:"required"`
	// Any of "high", "medium", "low".
	Confidence string                                                                              `json:"confidence" api:"required"`
	Reasoning  string                                                                              `json:"reasoning" api:"required"`
	Source     TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustClaimSource `json:"source" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Callout     respjson.Field
		Citations   respjson.Field
		Confidence  respjson.Field
		Reasoning   respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustClaim) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustClaim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustClaimCitation struct {
	URL                 string   `json:"url" api:"required"`
	Excerpts            []string `json:"excerpts" api:"nullable"`
	ExtractTemplateName string   `json:"extract_template_name" api:"nullable"`
	Title               string   `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL                 respjson.Field
		Excerpts            respjson.Field
		ExtractTemplateName respjson.Field
		Title               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustClaimCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustClaimCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustClaimSource struct {
	// Any of "primary", "secondary".
	Type                string `json:"type" api:"required"`
	URL                 string `json:"url" api:"required"`
	ExtractTemplateName string `json:"extract_template_name" api:"nullable"`
	Title               string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                respjson.Field
		URL                 respjson.Field
		ExtractTemplateName respjson.Field
		Title               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustClaimSource) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustClaimSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustSource struct {
	// Any of "primary", "secondary".
	Type                string `json:"type" api:"required"`
	URL                 string `json:"url" api:"required"`
	ExtractTemplateName string `json:"extract_template_name" api:"nullable"`
	Title               string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                respjson.Field
		URL                 respjson.Field
		ExtractTemplateName respjson.Field
		Title               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustSource) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunTextOutputTrustSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutput struct {
	Content TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputContentUnion `json:"content" api:"required"`
	Trust   TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrust        `json:"trust" api:"required"`
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
func (r TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutput) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputContentUnion
// contains all possible properties and values from [map[string]any], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfTaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputContentMapItem
// OfAnyArray]
type TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputContentUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfTaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputContentMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfTaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputContentMapItem respjson.Field
		OfAnyArray                                                                          respjson.Field
		raw                                                                                 string
	} `json:"-"`
}

func (u TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputContentUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputContentUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputContentUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputContentUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrust struct {
	Claims []TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustClaim `json:"claims" api:"required"`
	// Any of "high", "medium", "low".
	Confidence string                                                                           `json:"confidence" api:"required"`
	Reasoning  string                                                                           `json:"reasoning" api:"required"`
	Sources    []TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustSource `json:"sources" api:"required"`
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
func (r TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrust) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrust) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustClaim struct {
	Citations []TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustClaimCitation `json:"citations" api:"required"`
	// Any of "high", "medium", "low".
	Confidence string                                                                              `json:"confidence" api:"required"`
	Path       string                                                                              `json:"path" api:"required"`
	Reasoning  string                                                                              `json:"reasoning" api:"required"`
	Source     TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustClaimSource `json:"source" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Citations   respjson.Field
		Confidence  respjson.Field
		Path        respjson.Field
		Reasoning   respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustClaim) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustClaim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustClaimCitation struct {
	URL                 string   `json:"url" api:"required"`
	Excerpts            []string `json:"excerpts" api:"nullable"`
	ExtractTemplateName string   `json:"extract_template_name" api:"nullable"`
	Title               string   `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL                 respjson.Field
		Excerpts            respjson.Field
		ExtractTemplateName respjson.Field
		Title               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustClaimCitation) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustClaimCitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustClaimSource struct {
	// Any of "primary", "secondary".
	Type                string `json:"type" api:"required"`
	URL                 string `json:"url" api:"required"`
	ExtractTemplateName string `json:"extract_template_name" api:"nullable"`
	Title               string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                respjson.Field
		URL                 respjson.Field
		ExtractTemplateName respjson.Field
		Title               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustClaimSource) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustClaimSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustSource struct {
	// Any of "primary", "secondary".
	Type                string `json:"type" api:"required"`
	URL                 string `json:"url" api:"required"`
	ExtractTemplateName string `json:"extract_template_name" api:"nullable"`
	Title               string `json:"title" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type                respjson.Field
		URL                 respjson.Field
		ExtractTemplateName respjson.Field
		Title               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustSource) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseAgentRunResultOutputAgentRunJsonOutputTrustSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunResultRun struct {
	// Run identifier.
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "quickest", "quick", "research", "pro", "max".
	Effort string `json:"effort" api:"required"`
	// Interaction ID — pass as previous_interaction_id to reuse context.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status      string                                              `json:"status" api:"required"`
	CompletedAt time.Time                                           `json:"completed_at" api:"nullable" format:"date-time"`
	Error       TaskAgentRunGetResultResponseAgentRunResultRunError `json:"error" api:"nullable"`
	Prompt      string                                              `json:"prompt" api:"nullable"`
	StartedAt   time.Time                                           `json:"started_at" api:"nullable" format:"date-time"`
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
func (r TaskAgentRunGetResultResponseAgentRunResultRun) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResultResponseAgentRunResultRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunResultRunError struct {
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
func (r TaskAgentRunGetResultResponseAgentRunResultRunError) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResultResponseAgentRunResultRunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunFailedResult struct {
	Error TaskAgentRunGetResultResponseAgentRunFailedResultError `json:"error" api:"required"`
	Run   TaskAgentRunGetResultResponseAgentRunFailedResultRun   `json:"run" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Run         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskAgentRunGetResultResponseAgentRunFailedResult) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResultResponseAgentRunFailedResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunFailedResultError struct {
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
func (r TaskAgentRunGetResultResponseAgentRunFailedResultError) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResultResponseAgentRunFailedResultError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunFailedResultRun struct {
	// Run identifier.
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Any of "quickest", "quick", "research", "pro", "max".
	Effort string `json:"effort" api:"required"`
	// Interaction ID — pass as previous_interaction_id to reuse context.
	InteractionID string `json:"interaction_id" api:"required"`
	// True while status is 'queued' or 'running'.
	IsActive bool `json:"is_active" api:"required"`
	// Any of "queued", "running", "completed", "failed", "cancelled".
	Status      string                                                    `json:"status" api:"required"`
	CompletedAt time.Time                                                 `json:"completed_at" api:"nullable" format:"date-time"`
	Error       TaskAgentRunGetResultResponseAgentRunFailedResultRunError `json:"error" api:"nullable"`
	Prompt      string                                                    `json:"prompt" api:"nullable"`
	StartedAt   time.Time                                                 `json:"started_at" api:"nullable" format:"date-time"`
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
func (r TaskAgentRunGetResultResponseAgentRunFailedResultRun) RawJSON() string { return r.JSON.raw }
func (r *TaskAgentRunGetResultResponseAgentRunFailedResultRun) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunGetResultResponseAgentRunFailedResultRunError struct {
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
func (r TaskAgentRunGetResultResponseAgentRunFailedResultRunError) RawJSON() string {
	return r.JSON.raw
}
func (r *TaskAgentRunGetResultResponseAgentRunFailedResultRunError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskAgentRunStreamEventsResponse = any

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

type TaskAgentRunCancelParams struct {
	AgentID string `path:"agent_id" api:"required" json:"-"`
	paramObj
}

type TaskAgentRunGetParams struct {
	AgentID string `path:"agent_id" api:"required" json:"-"`
	paramObj
}

type TaskAgentRunGetResultParams struct {
	AgentID string `path:"agent_id" api:"required" json:"-"`
	paramObj
}

type TaskAgentRunStreamEventsParams struct {
	AgentID string `path:"agent_id" api:"required" json:"-"`
	paramObj
}
