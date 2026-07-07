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

// JobRunService contains methods and other services that help with interacting
// with the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobRunService] method instead.
type JobRunService struct {
	Options   []option.RequestOption
	Artifacts JobRunArtifactService
}

// NewJobRunService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJobRunService(opts ...option.RequestOption) (r JobRunService) {
	r = JobRunService{}
	r.Options = opts
	r.Artifacts = NewJobRunArtifactService(opts...)
	return
}

// List Runs for Job
func (r *JobRunService) List(ctx context.Context, jobID string, query JobRunListParams, opts ...option.RequestOption) (res *JobRunListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/jobs/%s/runs", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Cancel Run
func (r *JobRunService) Cancel(ctx context.Context, runID string, opts ...option.RequestOption) (res *JobRunCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/jobs/runs/%s/cancel", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Get Run
func (r *JobRunService) Get(ctx context.Context, runID string, opts ...option.RequestOption) (res *JobRunGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/jobs/runs/%s", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A page of job runs.
type JobRunListResponse struct {
	// Runs on this page.
	Items []JobRunListResponseItem `json:"items" api:"required"`
	// Current page number.
	Page int64 `json:"page" api:"required"`
	// Number of items per page.
	PerPage int64 `json:"per_page" api:"required"`
	// Total number of runs matching the query.
	Total int64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		Page        respjson.Field
		PerPage     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobRunListResponse) RawJSON() string { return r.JSON.raw }
func (r *JobRunListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single execution of a job.
type JobRunListResponseItem struct {
	// Unique run identifier (run\_<n>).
	ID string `json:"id" api:"required"`
	// When the run was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Identifier of the job this run belongs to.
	JobID string `json:"job_id" api:"required"`
	// Current run status.
	//
	// Any of "PENDING", "RUNNING", "SUCCESS", "FAILED", "CANCELLED", "TIMEOUT",
	// "WARNING".
	Status string `json:"status" api:"required"`
	// What triggered the run: 'schedule' or 'manual'.
	//
	// Any of "schedule", "manual".
	TriggeredBy string `json:"triggered_by" api:"required"`
	// When the run finished.
	FinishedAt time.Time `json:"finished_at" api:"nullable" format:"date-time"`
	// Number of input records processed.
	InputCount int64 `json:"input_count" api:"nullable"`
	// Number of result records produced.
	ResultCount int64 `json:"result_count" api:"nullable"`
	// When the run started executing.
	StartedAt time.Time `json:"started_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		JobID       respjson.Field
		Status      respjson.Field
		TriggeredBy respjson.Field
		FinishedAt  respjson.Field
		InputCount  respjson.Field
		ResultCount respjson.Field
		StartedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobRunListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *JobRunListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Result of cancelling a run.
type JobRunCancelResponse struct {
	// Identifier of the cancelled run.
	ID string `json:"id" api:"required"`
	// Run status after cancellation.
	//
	// Any of "PENDING", "RUNNING", "SUCCESS", "FAILED", "CANCELLED", "TIMEOUT",
	// "WARNING".
	Status JobRunCancelResponseStatus `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobRunCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *JobRunCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Run status after cancellation.
type JobRunCancelResponseStatus string

const (
	JobRunCancelResponseStatusPending   JobRunCancelResponseStatus = "PENDING"
	JobRunCancelResponseStatusRunning   JobRunCancelResponseStatus = "RUNNING"
	JobRunCancelResponseStatusSuccess   JobRunCancelResponseStatus = "SUCCESS"
	JobRunCancelResponseStatusFailed    JobRunCancelResponseStatus = "FAILED"
	JobRunCancelResponseStatusCancelled JobRunCancelResponseStatus = "CANCELLED"
	JobRunCancelResponseStatusTimeout   JobRunCancelResponseStatus = "TIMEOUT"
	JobRunCancelResponseStatusWarning   JobRunCancelResponseStatus = "WARNING"
)

// Full detail for a single run.
type JobRunGetResponse struct {
	// Unique run identifier (run\_<n>).
	ID string `json:"id" api:"required"`
	// When the run was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Context of the job this run belongs to.
	Job JobRunGetResponseJob `json:"job" api:"required"`
	// Current run status.
	//
	// Any of "PENDING", "RUNNING", "SUCCESS", "FAILED", "CANCELLED", "TIMEOUT",
	// "WARNING".
	Status JobRunGetResponseStatus `json:"status" api:"required"`
	// What triggered the run: 'schedule' or 'manual'.
	//
	// Any of "schedule", "manual".
	TriggeredBy JobRunGetResponseTriggeredBy `json:"triggered_by" api:"required"`
	// Error details for a failed run.
	Error JobRunGetResponseError `json:"error" api:"nullable"`
	// When the run finished.
	FinishedAt time.Time `json:"finished_at" api:"nullable" format:"date-time"`
	// Sample of the run's input records.
	InputsSample []any `json:"inputs_sample" api:"nullable"`
	// When the run started executing.
	StartedAt time.Time `json:"started_at" api:"nullable" format:"date-time"`
	// Aggregate metrics for a run.
	Summary JobRunGetResponseSummary `json:"summary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Job          respjson.Field
		Status       respjson.Field
		TriggeredBy  respjson.Field
		Error        respjson.Field
		FinishedAt   respjson.Field
		InputsSample respjson.Field
		StartedAt    respjson.Field
		Summary      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobRunGetResponse) RawJSON() string { return r.JSON.raw }
func (r *JobRunGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Context of the job this run belongs to.
type JobRunGetResponseJob struct {
	// Unique job identifier (job\_<n>).
	ID string `json:"id" api:"required"`
	// Internal job name.
	Name string `json:"name" api:"required"`
	// Name of the agent this job runs.
	AgentName string `json:"agent_name" api:"nullable"`
	// Human-friendly job name shown in the UI.
	DisplayName string `json:"display_name" api:"nullable"`
	// Cron-based schedule controlling when a job runs automatically.
	Schedule JobRunGetResponseJobSchedule `json:"schedule" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		AgentName   respjson.Field
		DisplayName respjson.Field
		Schedule    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobRunGetResponseJob) RawJSON() string { return r.JSON.raw }
func (r *JobRunGetResponseJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cron-based schedule controlling when a job runs automatically.
type JobRunGetResponseJobSchedule struct {
	// Cron expression defining when the job runs.
	Cron string `json:"cron" api:"required"`
	// Whether the schedule is currently active.
	Enabled bool `json:"enabled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cron        respjson.Field
		Enabled     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobRunGetResponseJobSchedule) RawJSON() string { return r.JSON.raw }
func (r *JobRunGetResponseJobSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current run status.
type JobRunGetResponseStatus string

const (
	JobRunGetResponseStatusPending   JobRunGetResponseStatus = "PENDING"
	JobRunGetResponseStatusRunning   JobRunGetResponseStatus = "RUNNING"
	JobRunGetResponseStatusSuccess   JobRunGetResponseStatus = "SUCCESS"
	JobRunGetResponseStatusFailed    JobRunGetResponseStatus = "FAILED"
	JobRunGetResponseStatusCancelled JobRunGetResponseStatus = "CANCELLED"
	JobRunGetResponseStatusTimeout   JobRunGetResponseStatus = "TIMEOUT"
	JobRunGetResponseStatusWarning   JobRunGetResponseStatus = "WARNING"
)

// What triggered the run: 'schedule' or 'manual'.
type JobRunGetResponseTriggeredBy string

const (
	JobRunGetResponseTriggeredBySchedule JobRunGetResponseTriggeredBy = "schedule"
	JobRunGetResponseTriggeredByManual   JobRunGetResponseTriggeredBy = "manual"
)

// Error details for a failed run.
type JobRunGetResponseError struct {
	// Sample of individual error records from the run.
	ErrorsSample []map[string]any `json:"errors_sample" api:"nullable"`
	// Human-readable error message.
	Message string `json:"message" api:"nullable"`
	// Pipeline step where the error occurred.
	Step string `json:"step" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ErrorsSample respjson.Field
		Message      respjson.Field
		Step         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobRunGetResponseError) RawJSON() string { return r.JSON.raw }
func (r *JobRunGetResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Aggregate metrics for a run.
type JobRunGetResponseSummary struct {
	// Number of input records processed.
	InputCount int64 `json:"input_count" api:"nullable"`
	// Fraction of inputs that produced a result (result_count / input_count), from 0.0
	// to 1.0.
	MatchRate float64 `json:"match_rate" api:"nullable"`
	// Number of result records produced.
	ResultCount int64 `json:"result_count" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputCount  respjson.Field
		MatchRate   respjson.Field
		ResultCount respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobRunGetResponseSummary) RawJSON() string { return r.JSON.raw }
func (r *JobRunGetResponseSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobRunListParams struct {
	// Filter by status
	Status  param.Opt[string] `query:"status,omitzero" json:"-"`
	Page    param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PerPage param.Opt[int64]  `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JobRunListParams]'s query parameters as `url.Values`.
func (r JobRunListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
