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

// JobService contains methods and other services that help with interacting with
// the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobService] method instead.
type JobService struct {
	Options []option.RequestOption
	Runs    JobRunService
}

// NewJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJobService(opts ...option.RequestOption) (r JobService) {
	r = JobService{}
	r.Options = opts
	r.Runs = NewJobRunService(opts...)
	return
}

// Create Job Public V2
func (r *JobService) New(ctx context.Context, body JobNewParams, opts ...option.RequestOption) (res *JobNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update Job Public V2
func (r *JobService) Update(ctx context.Context, jobID string, body JobUpdateParams, opts ...option.RequestOption) (res *JobUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Jobs Public V2
func (r *JobService) List(ctx context.Context, query JobListParams, opts ...option.RequestOption) (res *JobListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete Job Public V2
func (r *JobService) Delete(ctx context.Context, jobID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return err
	}
	path := fmt.Sprintf("v2/jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get Job Public V2
func (r *JobService) Get(ctx context.Context, jobID string, opts ...option.RequestOption) (res *JobGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A configured job: an agent plus its schedule, inputs, and destination.
type JobNewResponse struct {
	// Unique job identifier (job\_<n>).
	ID string `json:"id" api:"required"`
	// Job name.
	Name string `json:"name" api:"required"`
	// When the job was created.
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Free-text description of the job.
	Description string `json:"description" api:"nullable"`
	// Where a job writes its results.
	Destination JobNewResponseDestination `json:"destination" api:"nullable"`
	// Human-friendly job name shown in the UI.
	DisplayName string `json:"display_name" api:"nullable"`
	// Name of the extract template this job runs.
	ExtractTemplateName string `json:"extract_template_name" api:"nullable"`
	// Configuration for the input data a job processes.
	Inputs JobNewResponseInputs `json:"inputs" api:"nullable"`
	// Timestamp of the most recent run.
	LastRunAt time.Time `json:"last_run_at" api:"nullable" format:"date-time"`
	// Status of the most recent run.
	//
	// Any of "PENDING", "RUNNING", "SUCCESS", "FAILED", "CANCELLED", "TIMEOUT",
	// "WARNING".
	LastRunStatus JobNewResponseLastRunStatus `json:"last_run_status" api:"nullable"`
	// Cron-based schedule controlling when a job runs automatically.
	Schedule JobNewResponseSchedule `json:"schedule" api:"nullable"`
	// When the job was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Name                respjson.Field
		CreatedAt           respjson.Field
		Description         respjson.Field
		Destination         respjson.Field
		DisplayName         respjson.Field
		ExtractTemplateName respjson.Field
		Inputs              respjson.Field
		LastRunAt           respjson.Field
		LastRunStatus       respjson.Field
		Schedule            respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobNewResponse) RawJSON() string { return r.JSON.raw }
func (r *JobNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where a job writes its results.
type JobNewResponseDestination struct {
	// Destination path the output is written to.
	Path string `json:"path" api:"required"`
	// Destination kind: a local 'file' or an 's3' bucket.
	//
	// Any of "file", "s3".
	Type string `json:"type" api:"required"`
	// Output file format.
	//
	// Any of "jsonl", "csv", "parquet".
	Format string `json:"format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Type        respjson.Field
		Format      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobNewResponseDestination) RawJSON() string { return r.JSON.raw }
func (r *JobNewResponseDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the input data a job processes.
type JobNewResponseInputs struct {
	// How inputs are supplied: an 's3' bucket, 'inline' records, or an uploaded
	// 'file'.
	//
	// Any of "s3", "inline", "file".
	Type string `json:"type" api:"required"`
	// Inline list of input records. Used when type is 'inline'.
	Data []map[string]any `json:"data" api:"nullable"`
	// Path to the input file; must start with 's3' or 'file\_'. Used for 's3'/'file'
	// types.
	FilePath string `json:"file_path" api:"nullable"`
	// Inline input records keyed by source node id, e.g. {'source_a': [{...}]}. Used
	// when type is 'inline' on a dynamic-workflow job, which has one source node per
	// input file. Mutually exclusive with 'data'.
	NodeData map[string][]map[string]any `json:"node_data" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Data        respjson.Field
		FilePath    respjson.Field
		NodeData    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobNewResponseInputs) RawJSON() string { return r.JSON.raw }
func (r *JobNewResponseInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the most recent run.
type JobNewResponseLastRunStatus string

const (
	JobNewResponseLastRunStatusPending   JobNewResponseLastRunStatus = "PENDING"
	JobNewResponseLastRunStatusRunning   JobNewResponseLastRunStatus = "RUNNING"
	JobNewResponseLastRunStatusSuccess   JobNewResponseLastRunStatus = "SUCCESS"
	JobNewResponseLastRunStatusFailed    JobNewResponseLastRunStatus = "FAILED"
	JobNewResponseLastRunStatusCancelled JobNewResponseLastRunStatus = "CANCELLED"
	JobNewResponseLastRunStatusTimeout   JobNewResponseLastRunStatus = "TIMEOUT"
	JobNewResponseLastRunStatusWarning   JobNewResponseLastRunStatus = "WARNING"
)

// Cron-based schedule controlling when a job runs automatically.
type JobNewResponseSchedule struct {
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
func (r JobNewResponseSchedule) RawJSON() string { return r.JSON.raw }
func (r *JobNewResponseSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A configured job: an agent plus its schedule, inputs, and destination.
type JobUpdateResponse struct {
	// Unique job identifier (job\_<n>).
	ID string `json:"id" api:"required"`
	// Job name.
	Name string `json:"name" api:"required"`
	// When the job was created.
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Free-text description of the job.
	Description string `json:"description" api:"nullable"`
	// Where a job writes its results.
	Destination JobUpdateResponseDestination `json:"destination" api:"nullable"`
	// Human-friendly job name shown in the UI.
	DisplayName string `json:"display_name" api:"nullable"`
	// Name of the extract template this job runs.
	ExtractTemplateName string `json:"extract_template_name" api:"nullable"`
	// Configuration for the input data a job processes.
	Inputs JobUpdateResponseInputs `json:"inputs" api:"nullable"`
	// Timestamp of the most recent run.
	LastRunAt time.Time `json:"last_run_at" api:"nullable" format:"date-time"`
	// Status of the most recent run.
	//
	// Any of "PENDING", "RUNNING", "SUCCESS", "FAILED", "CANCELLED", "TIMEOUT",
	// "WARNING".
	LastRunStatus JobUpdateResponseLastRunStatus `json:"last_run_status" api:"nullable"`
	// Cron-based schedule controlling when a job runs automatically.
	Schedule JobUpdateResponseSchedule `json:"schedule" api:"nullable"`
	// When the job was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Name                respjson.Field
		CreatedAt           respjson.Field
		Description         respjson.Field
		Destination         respjson.Field
		DisplayName         respjson.Field
		ExtractTemplateName respjson.Field
		Inputs              respjson.Field
		LastRunAt           respjson.Field
		LastRunStatus       respjson.Field
		Schedule            respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *JobUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where a job writes its results.
type JobUpdateResponseDestination struct {
	// Destination path the output is written to.
	Path string `json:"path" api:"required"`
	// Destination kind: a local 'file' or an 's3' bucket.
	//
	// Any of "file", "s3".
	Type string `json:"type" api:"required"`
	// Output file format.
	//
	// Any of "jsonl", "csv", "parquet".
	Format string `json:"format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Type        respjson.Field
		Format      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobUpdateResponseDestination) RawJSON() string { return r.JSON.raw }
func (r *JobUpdateResponseDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the input data a job processes.
type JobUpdateResponseInputs struct {
	// How inputs are supplied: an 's3' bucket, 'inline' records, or an uploaded
	// 'file'.
	//
	// Any of "s3", "inline", "file".
	Type string `json:"type" api:"required"`
	// Inline list of input records. Used when type is 'inline'.
	Data []map[string]any `json:"data" api:"nullable"`
	// Path to the input file; must start with 's3' or 'file\_'. Used for 's3'/'file'
	// types.
	FilePath string `json:"file_path" api:"nullable"`
	// Inline input records keyed by source node id, e.g. {'source_a': [{...}]}. Used
	// when type is 'inline' on a dynamic-workflow job, which has one source node per
	// input file. Mutually exclusive with 'data'.
	NodeData map[string][]map[string]any `json:"node_data" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Data        respjson.Field
		FilePath    respjson.Field
		NodeData    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobUpdateResponseInputs) RawJSON() string { return r.JSON.raw }
func (r *JobUpdateResponseInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the most recent run.
type JobUpdateResponseLastRunStatus string

const (
	JobUpdateResponseLastRunStatusPending   JobUpdateResponseLastRunStatus = "PENDING"
	JobUpdateResponseLastRunStatusRunning   JobUpdateResponseLastRunStatus = "RUNNING"
	JobUpdateResponseLastRunStatusSuccess   JobUpdateResponseLastRunStatus = "SUCCESS"
	JobUpdateResponseLastRunStatusFailed    JobUpdateResponseLastRunStatus = "FAILED"
	JobUpdateResponseLastRunStatusCancelled JobUpdateResponseLastRunStatus = "CANCELLED"
	JobUpdateResponseLastRunStatusTimeout   JobUpdateResponseLastRunStatus = "TIMEOUT"
	JobUpdateResponseLastRunStatusWarning   JobUpdateResponseLastRunStatus = "WARNING"
)

// Cron-based schedule controlling when a job runs automatically.
type JobUpdateResponseSchedule struct {
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
func (r JobUpdateResponseSchedule) RawJSON() string { return r.JSON.raw }
func (r *JobUpdateResponseSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobListResponse struct {
	// Items returned in this page.
	Items []JobListResponseItem `json:"items" api:"required"`
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
func (r JobListResponse) RawJSON() string { return r.JSON.raw }
func (r *JobListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A configured job: an agent plus its schedule, inputs, and destination.
type JobListResponseItem struct {
	// Unique job identifier (job\_<n>).
	ID string `json:"id" api:"required"`
	// Job name.
	Name string `json:"name" api:"required"`
	// When the job was created.
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Free-text description of the job.
	Description string `json:"description" api:"nullable"`
	// Where a job writes its results.
	Destination JobListResponseItemDestination `json:"destination" api:"nullable"`
	// Human-friendly job name shown in the UI.
	DisplayName string `json:"display_name" api:"nullable"`
	// Name of the extract template this job runs.
	ExtractTemplateName string `json:"extract_template_name" api:"nullable"`
	// Configuration for the input data a job processes.
	Inputs JobListResponseItemInputs `json:"inputs" api:"nullable"`
	// Timestamp of the most recent run.
	LastRunAt time.Time `json:"last_run_at" api:"nullable" format:"date-time"`
	// Status of the most recent run.
	//
	// Any of "PENDING", "RUNNING", "SUCCESS", "FAILED", "CANCELLED", "TIMEOUT",
	// "WARNING".
	LastRunStatus string `json:"last_run_status" api:"nullable"`
	// Cron-based schedule controlling when a job runs automatically.
	Schedule JobListResponseItemSchedule `json:"schedule" api:"nullable"`
	// When the job was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Name                respjson.Field
		CreatedAt           respjson.Field
		Description         respjson.Field
		Destination         respjson.Field
		DisplayName         respjson.Field
		ExtractTemplateName respjson.Field
		Inputs              respjson.Field
		LastRunAt           respjson.Field
		LastRunStatus       respjson.Field
		Schedule            respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *JobListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where a job writes its results.
type JobListResponseItemDestination struct {
	// Destination path the output is written to.
	Path string `json:"path" api:"required"`
	// Destination kind: a local 'file' or an 's3' bucket.
	//
	// Any of "file", "s3".
	Type string `json:"type" api:"required"`
	// Output file format.
	//
	// Any of "jsonl", "csv", "parquet".
	Format string `json:"format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Type        respjson.Field
		Format      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobListResponseItemDestination) RawJSON() string { return r.JSON.raw }
func (r *JobListResponseItemDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the input data a job processes.
type JobListResponseItemInputs struct {
	// How inputs are supplied: an 's3' bucket, 'inline' records, or an uploaded
	// 'file'.
	//
	// Any of "s3", "inline", "file".
	Type string `json:"type" api:"required"`
	// Inline list of input records. Used when type is 'inline'.
	Data []map[string]any `json:"data" api:"nullable"`
	// Path to the input file; must start with 's3' or 'file\_'. Used for 's3'/'file'
	// types.
	FilePath string `json:"file_path" api:"nullable"`
	// Inline input records keyed by source node id, e.g. {'source_a': [{...}]}. Used
	// when type is 'inline' on a dynamic-workflow job, which has one source node per
	// input file. Mutually exclusive with 'data'.
	NodeData map[string][]map[string]any `json:"node_data" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Data        respjson.Field
		FilePath    respjson.Field
		NodeData    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobListResponseItemInputs) RawJSON() string { return r.JSON.raw }
func (r *JobListResponseItemInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Cron-based schedule controlling when a job runs automatically.
type JobListResponseItemSchedule struct {
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
func (r JobListResponseItemSchedule) RawJSON() string { return r.JSON.raw }
func (r *JobListResponseItemSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A configured job: an agent plus its schedule, inputs, and destination.
type JobGetResponse struct {
	// Unique job identifier (job\_<n>).
	ID string `json:"id" api:"required"`
	// Job name.
	Name string `json:"name" api:"required"`
	// When the job was created.
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Free-text description of the job.
	Description string `json:"description" api:"nullable"`
	// Where a job writes its results.
	Destination JobGetResponseDestination `json:"destination" api:"nullable"`
	// Human-friendly job name shown in the UI.
	DisplayName string `json:"display_name" api:"nullable"`
	// Name of the extract template this job runs.
	ExtractTemplateName string `json:"extract_template_name" api:"nullable"`
	// Configuration for the input data a job processes.
	Inputs JobGetResponseInputs `json:"inputs" api:"nullable"`
	// Timestamp of the most recent run.
	LastRunAt time.Time `json:"last_run_at" api:"nullable" format:"date-time"`
	// Status of the most recent run.
	//
	// Any of "PENDING", "RUNNING", "SUCCESS", "FAILED", "CANCELLED", "TIMEOUT",
	// "WARNING".
	LastRunStatus JobGetResponseLastRunStatus `json:"last_run_status" api:"nullable"`
	// Cron-based schedule controlling when a job runs automatically.
	Schedule JobGetResponseSchedule `json:"schedule" api:"nullable"`
	// When the job was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		Name                respjson.Field
		CreatedAt           respjson.Field
		Description         respjson.Field
		Destination         respjson.Field
		DisplayName         respjson.Field
		ExtractTemplateName respjson.Field
		Inputs              respjson.Field
		LastRunAt           respjson.Field
		LastRunStatus       respjson.Field
		Schedule            respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobGetResponse) RawJSON() string { return r.JSON.raw }
func (r *JobGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where a job writes its results.
type JobGetResponseDestination struct {
	// Destination path the output is written to.
	Path string `json:"path" api:"required"`
	// Destination kind: a local 'file' or an 's3' bucket.
	//
	// Any of "file", "s3".
	Type string `json:"type" api:"required"`
	// Output file format.
	//
	// Any of "jsonl", "csv", "parquet".
	Format string `json:"format"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Type        respjson.Field
		Format      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobGetResponseDestination) RawJSON() string { return r.JSON.raw }
func (r *JobGetResponseDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for the input data a job processes.
type JobGetResponseInputs struct {
	// How inputs are supplied: an 's3' bucket, 'inline' records, or an uploaded
	// 'file'.
	//
	// Any of "s3", "inline", "file".
	Type string `json:"type" api:"required"`
	// Inline list of input records. Used when type is 'inline'.
	Data []map[string]any `json:"data" api:"nullable"`
	// Path to the input file; must start with 's3' or 'file\_'. Used for 's3'/'file'
	// types.
	FilePath string `json:"file_path" api:"nullable"`
	// Inline input records keyed by source node id, e.g. {'source_a': [{...}]}. Used
	// when type is 'inline' on a dynamic-workflow job, which has one source node per
	// input file. Mutually exclusive with 'data'.
	NodeData map[string][]map[string]any `json:"node_data" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Data        respjson.Field
		FilePath    respjson.Field
		NodeData    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobGetResponseInputs) RawJSON() string { return r.JSON.raw }
func (r *JobGetResponseInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the most recent run.
type JobGetResponseLastRunStatus string

const (
	JobGetResponseLastRunStatusPending   JobGetResponseLastRunStatus = "PENDING"
	JobGetResponseLastRunStatusRunning   JobGetResponseLastRunStatus = "RUNNING"
	JobGetResponseLastRunStatusSuccess   JobGetResponseLastRunStatus = "SUCCESS"
	JobGetResponseLastRunStatusFailed    JobGetResponseLastRunStatus = "FAILED"
	JobGetResponseLastRunStatusCancelled JobGetResponseLastRunStatus = "CANCELLED"
	JobGetResponseLastRunStatusTimeout   JobGetResponseLastRunStatus = "TIMEOUT"
	JobGetResponseLastRunStatusWarning   JobGetResponseLastRunStatus = "WARNING"
)

// Cron-based schedule controlling when a job runs automatically.
type JobGetResponseSchedule struct {
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
func (r JobGetResponseSchedule) RawJSON() string { return r.JSON.raw }
func (r *JobGetResponseSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobNewParams struct {
	// Name of the extract template to run.
	ExtractTemplateName string `json:"extract_template_name" api:"required"`
	// Job name.
	Name string `json:"name" api:"required"`
	// Free-text description of the job.
	Description param.Opt[string] `json:"description,omitzero"`
	// Human-friendly job name shown in the UI.
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// Where a job writes its results.
	Destination JobNewParamsDestination `json:"destination,omitzero"`
	// Configuration for the input data a job processes.
	Inputs JobNewParamsInputs `json:"inputs,omitzero"`
	// Cron-based schedule controlling when a job runs automatically.
	Schedule JobNewParamsSchedule `json:"schedule,omitzero"`
	paramObj
}

func (r JobNewParams) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where a job writes its results.
//
// The properties Path, Type are required.
type JobNewParamsDestination struct {
	// Destination path the output is written to.
	Path string `json:"path" api:"required"`
	// Destination kind: a local 'file' or an 's3' bucket.
	//
	// Any of "file", "s3".
	Type string `json:"type,omitzero" api:"required"`
	// Output file format.
	//
	// Any of "jsonl", "csv", "parquet".
	Format string `json:"format,omitzero"`
	paramObj
}

func (r JobNewParamsDestination) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParamsDestination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewParamsDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JobNewParamsDestination](
		"type", "file", "s3",
	)
	apijson.RegisterFieldValidator[JobNewParamsDestination](
		"format", "jsonl", "csv", "parquet",
	)
}

// Configuration for the input data a job processes.
//
// The property Type is required.
type JobNewParamsInputs struct {
	// How inputs are supplied: an 's3' bucket, 'inline' records, or an uploaded
	// 'file'.
	//
	// Any of "s3", "inline", "file".
	Type string `json:"type,omitzero" api:"required"`
	// Path to the input file; must start with 's3' or 'file\_'. Used for 's3'/'file'
	// types.
	FilePath param.Opt[string] `json:"file_path,omitzero"`
	// Inline list of input records. Used when type is 'inline'.
	Data []map[string]any `json:"data,omitzero"`
	// Inline input records keyed by source node id, e.g. {'source_a': [{...}]}. Used
	// when type is 'inline' on a dynamic-workflow job, which has one source node per
	// input file. Mutually exclusive with 'data'.
	NodeData map[string][]map[string]any `json:"node_data,omitzero"`
	paramObj
}

func (r JobNewParamsInputs) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParamsInputs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewParamsInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JobNewParamsInputs](
		"type", "s3", "inline", "file",
	)
}

// Cron-based schedule controlling when a job runs automatically.
//
// The properties Cron, Enabled are required.
type JobNewParamsSchedule struct {
	// Cron expression defining when the job runs.
	Cron string `json:"cron" api:"required"`
	// Whether the schedule is currently active.
	Enabled bool `json:"enabled" api:"required"`
	paramObj
}

func (r JobNewParamsSchedule) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParamsSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewParamsSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobUpdateParams struct {
	// New description.
	Description param.Opt[string] `json:"description,omitzero"`
	// New display name.
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// Where a job writes its results.
	Destination JobUpdateParamsDestination `json:"destination,omitzero"`
	// Configuration for the input data a job processes.
	Inputs JobUpdateParamsInputs `json:"inputs,omitzero"`
	// Cron-based schedule controlling when a job runs automatically.
	Schedule JobUpdateParamsSchedule `json:"schedule,omitzero"`
	paramObj
}

func (r JobUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow JobUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Where a job writes its results.
//
// The properties Path, Type are required.
type JobUpdateParamsDestination struct {
	// Destination path the output is written to.
	Path string `json:"path" api:"required"`
	// Destination kind: a local 'file' or an 's3' bucket.
	//
	// Any of "file", "s3".
	Type string `json:"type,omitzero" api:"required"`
	// Output file format.
	//
	// Any of "jsonl", "csv", "parquet".
	Format string `json:"format,omitzero"`
	paramObj
}

func (r JobUpdateParamsDestination) MarshalJSON() (data []byte, err error) {
	type shadow JobUpdateParamsDestination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobUpdateParamsDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JobUpdateParamsDestination](
		"type", "file", "s3",
	)
	apijson.RegisterFieldValidator[JobUpdateParamsDestination](
		"format", "jsonl", "csv", "parquet",
	)
}

// Configuration for the input data a job processes.
//
// The property Type is required.
type JobUpdateParamsInputs struct {
	// How inputs are supplied: an 's3' bucket, 'inline' records, or an uploaded
	// 'file'.
	//
	// Any of "s3", "inline", "file".
	Type string `json:"type,omitzero" api:"required"`
	// Path to the input file; must start with 's3' or 'file\_'. Used for 's3'/'file'
	// types.
	FilePath param.Opt[string] `json:"file_path,omitzero"`
	// Inline list of input records. Used when type is 'inline'.
	Data []map[string]any `json:"data,omitzero"`
	// Inline input records keyed by source node id, e.g. {'source_a': [{...}]}. Used
	// when type is 'inline' on a dynamic-workflow job, which has one source node per
	// input file. Mutually exclusive with 'data'.
	NodeData map[string][]map[string]any `json:"node_data,omitzero"`
	paramObj
}

func (r JobUpdateParamsInputs) MarshalJSON() (data []byte, err error) {
	type shadow JobUpdateParamsInputs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobUpdateParamsInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JobUpdateParamsInputs](
		"type", "s3", "inline", "file",
	)
}

// Cron-based schedule controlling when a job runs automatically.
//
// The properties Cron, Enabled are required.
type JobUpdateParamsSchedule struct {
	// Cron expression defining when the job runs.
	Cron string `json:"cron" api:"required"`
	// Whether the schedule is currently active.
	Enabled bool `json:"enabled" api:"required"`
	paramObj
}

func (r JobUpdateParamsSchedule) MarshalJSON() (data []byte, err error) {
	type shadow JobUpdateParamsSchedule
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobUpdateParamsSchedule) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobListParams struct {
	Limit  param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JobListParams]'s query parameters as `url.Values`.
func (r JobListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
