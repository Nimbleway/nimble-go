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

// Create Job
func (r *JobService) New(ctx context.Context, body JobNewParams, opts ...option.RequestOption) (res *JobNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Update Job
func (r *JobService) Update(ctx context.Context, jobID string, body JobUpdateParams, opts ...option.RequestOption) (res *JobUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Jobs
func (r *JobService) List(ctx context.Context, query JobListParams, opts ...option.RequestOption) (res *JobListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete Job
func (r *JobService) Delete(ctx context.Context, jobID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return err
	}
	path := fmt.Sprintf("v1/jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Get Job
func (r *JobService) Get(ctx context.Context, jobID string, opts ...option.RequestOption) (res *JobGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Trigger Run
func (r *JobService) Run(ctx context.Context, jobID string, opts ...option.RequestOption) (res *JobRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/jobs/%s/runs", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type JobNewResponse struct {
	ID          string                    `json:"id" api:"required"`
	Name        string                    `json:"name" api:"required"`
	AgentName   string                    `json:"agent_name" api:"nullable"`
	CreatedAt   time.Time                 `json:"created_at" api:"nullable" format:"date-time"`
	Description string                    `json:"description" api:"nullable"`
	Destination JobNewResponseDestination `json:"destination" api:"nullable"`
	DisplayName string                    `json:"display_name" api:"nullable"`
	Inputs      JobNewResponseInputs      `json:"inputs" api:"nullable"`
	LastRunAt   time.Time                 `json:"last_run_at" api:"nullable" format:"date-time"`
	// Any of "PENDING", "RUNNING", "SUCCESS", "FAILED", "CANCELLED", "TIMEOUT",
	// "WARNING".
	LastRunStatus JobNewResponseLastRunStatus `json:"last_run_status" api:"nullable"`
	Schedule      JobNewResponseSchedule      `json:"schedule" api:"nullable"`
	UpdatedAt     time.Time                   `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Name          respjson.Field
		AgentName     respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		Destination   respjson.Field
		DisplayName   respjson.Field
		Inputs        respjson.Field
		LastRunAt     respjson.Field
		LastRunStatus respjson.Field
		Schedule      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobNewResponse) RawJSON() string { return r.JSON.raw }
func (r *JobNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobNewResponseDestination struct {
	Path string `json:"path" api:"required"`
	// Any of "file", "s3".
	Type string `json:"type" api:"required"`
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

type JobNewResponseInputs struct {
	// Any of "s3", "inline", "file".
	Type     string           `json:"type" api:"required"`
	Data     []map[string]any `json:"data" api:"nullable"`
	FilePath string           `json:"file_path" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Data        respjson.Field
		FilePath    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobNewResponseInputs) RawJSON() string { return r.JSON.raw }
func (r *JobNewResponseInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type JobNewResponseSchedule struct {
	Cron    string `json:"cron" api:"required"`
	Enabled bool   `json:"enabled" api:"required"`
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

type JobUpdateResponse struct {
	ID          string                       `json:"id" api:"required"`
	Name        string                       `json:"name" api:"required"`
	AgentName   string                       `json:"agent_name" api:"nullable"`
	CreatedAt   time.Time                    `json:"created_at" api:"nullable" format:"date-time"`
	Description string                       `json:"description" api:"nullable"`
	Destination JobUpdateResponseDestination `json:"destination" api:"nullable"`
	DisplayName string                       `json:"display_name" api:"nullable"`
	Inputs      JobUpdateResponseInputs      `json:"inputs" api:"nullable"`
	LastRunAt   time.Time                    `json:"last_run_at" api:"nullable" format:"date-time"`
	// Any of "PENDING", "RUNNING", "SUCCESS", "FAILED", "CANCELLED", "TIMEOUT",
	// "WARNING".
	LastRunStatus JobUpdateResponseLastRunStatus `json:"last_run_status" api:"nullable"`
	Schedule      JobUpdateResponseSchedule      `json:"schedule" api:"nullable"`
	UpdatedAt     time.Time                      `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Name          respjson.Field
		AgentName     respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		Destination   respjson.Field
		DisplayName   respjson.Field
		Inputs        respjson.Field
		LastRunAt     respjson.Field
		LastRunStatus respjson.Field
		Schedule      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *JobUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobUpdateResponseDestination struct {
	Path string `json:"path" api:"required"`
	// Any of "file", "s3".
	Type string `json:"type" api:"required"`
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

type JobUpdateResponseInputs struct {
	// Any of "s3", "inline", "file".
	Type     string           `json:"type" api:"required"`
	Data     []map[string]any `json:"data" api:"nullable"`
	FilePath string           `json:"file_path" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Data        respjson.Field
		FilePath    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobUpdateResponseInputs) RawJSON() string { return r.JSON.raw }
func (r *JobUpdateResponseInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type JobUpdateResponseSchedule struct {
	Cron    string `json:"cron" api:"required"`
	Enabled bool   `json:"enabled" api:"required"`
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
	Items   []JobListResponseItem `json:"items" api:"required"`
	Page    int64                 `json:"page" api:"required"`
	PerPage int64                 `json:"per_page" api:"required"`
	Total   int64                 `json:"total" api:"required"`
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
func (r JobListResponse) RawJSON() string { return r.JSON.raw }
func (r *JobListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobListResponseItem struct {
	ID          string                         `json:"id" api:"required"`
	Name        string                         `json:"name" api:"required"`
	AgentName   string                         `json:"agent_name" api:"nullable"`
	CreatedAt   time.Time                      `json:"created_at" api:"nullable" format:"date-time"`
	Description string                         `json:"description" api:"nullable"`
	Destination JobListResponseItemDestination `json:"destination" api:"nullable"`
	DisplayName string                         `json:"display_name" api:"nullable"`
	Inputs      JobListResponseItemInputs      `json:"inputs" api:"nullable"`
	LastRunAt   time.Time                      `json:"last_run_at" api:"nullable" format:"date-time"`
	// Any of "PENDING", "RUNNING", "SUCCESS", "FAILED", "CANCELLED", "TIMEOUT",
	// "WARNING".
	LastRunStatus string                      `json:"last_run_status" api:"nullable"`
	Schedule      JobListResponseItemSchedule `json:"schedule" api:"nullable"`
	UpdatedAt     time.Time                   `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Name          respjson.Field
		AgentName     respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		Destination   respjson.Field
		DisplayName   respjson.Field
		Inputs        respjson.Field
		LastRunAt     respjson.Field
		LastRunStatus respjson.Field
		Schedule      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *JobListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobListResponseItemDestination struct {
	Path string `json:"path" api:"required"`
	// Any of "file", "s3".
	Type string `json:"type" api:"required"`
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

type JobListResponseItemInputs struct {
	// Any of "s3", "inline", "file".
	Type     string           `json:"type" api:"required"`
	Data     []map[string]any `json:"data" api:"nullable"`
	FilePath string           `json:"file_path" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Data        respjson.Field
		FilePath    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobListResponseItemInputs) RawJSON() string { return r.JSON.raw }
func (r *JobListResponseItemInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobListResponseItemSchedule struct {
	Cron    string `json:"cron" api:"required"`
	Enabled bool   `json:"enabled" api:"required"`
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

type JobGetResponse struct {
	ID          string                    `json:"id" api:"required"`
	Name        string                    `json:"name" api:"required"`
	AgentName   string                    `json:"agent_name" api:"nullable"`
	CreatedAt   time.Time                 `json:"created_at" api:"nullable" format:"date-time"`
	Description string                    `json:"description" api:"nullable"`
	Destination JobGetResponseDestination `json:"destination" api:"nullable"`
	DisplayName string                    `json:"display_name" api:"nullable"`
	Inputs      JobGetResponseInputs      `json:"inputs" api:"nullable"`
	LastRunAt   time.Time                 `json:"last_run_at" api:"nullable" format:"date-time"`
	// Any of "PENDING", "RUNNING", "SUCCESS", "FAILED", "CANCELLED", "TIMEOUT",
	// "WARNING".
	LastRunStatus JobGetResponseLastRunStatus `json:"last_run_status" api:"nullable"`
	Schedule      JobGetResponseSchedule      `json:"schedule" api:"nullable"`
	UpdatedAt     time.Time                   `json:"updated_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Name          respjson.Field
		AgentName     respjson.Field
		CreatedAt     respjson.Field
		Description   respjson.Field
		Destination   respjson.Field
		DisplayName   respjson.Field
		Inputs        respjson.Field
		LastRunAt     respjson.Field
		LastRunStatus respjson.Field
		Schedule      respjson.Field
		UpdatedAt     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobGetResponse) RawJSON() string { return r.JSON.raw }
func (r *JobGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobGetResponseDestination struct {
	Path string `json:"path" api:"required"`
	// Any of "file", "s3".
	Type string `json:"type" api:"required"`
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

type JobGetResponseInputs struct {
	// Any of "s3", "inline", "file".
	Type     string           `json:"type" api:"required"`
	Data     []map[string]any `json:"data" api:"nullable"`
	FilePath string           `json:"file_path" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Data        respjson.Field
		FilePath    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobGetResponseInputs) RawJSON() string { return r.JSON.raw }
func (r *JobGetResponseInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

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

type JobGetResponseSchedule struct {
	Cron    string `json:"cron" api:"required"`
	Enabled bool   `json:"enabled" api:"required"`
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

type JobRunResponse struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	JobID     string    `json:"job_id" api:"required"`
	// Any of "PENDING", "RUNNING", "SUCCESS", "FAILED", "CANCELLED", "TIMEOUT",
	// "WARNING".
	Status JobRunResponseStatus `json:"status" api:"required"`
	// Any of "schedule", "manual".
	TriggeredBy JobRunResponseTriggeredBy `json:"triggered_by" api:"required"`
	FinishedAt  time.Time                 `json:"finished_at" api:"nullable" format:"date-time"`
	InputCount  int64                     `json:"input_count" api:"nullable"`
	ResultCount int64                     `json:"result_count" api:"nullable"`
	StartedAt   time.Time                 `json:"started_at" api:"nullable" format:"date-time"`
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
func (r JobRunResponse) RawJSON() string { return r.JSON.raw }
func (r *JobRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobRunResponseStatus string

const (
	JobRunResponseStatusPending   JobRunResponseStatus = "PENDING"
	JobRunResponseStatusRunning   JobRunResponseStatus = "RUNNING"
	JobRunResponseStatusSuccess   JobRunResponseStatus = "SUCCESS"
	JobRunResponseStatusFailed    JobRunResponseStatus = "FAILED"
	JobRunResponseStatusCancelled JobRunResponseStatus = "CANCELLED"
	JobRunResponseStatusTimeout   JobRunResponseStatus = "TIMEOUT"
	JobRunResponseStatusWarning   JobRunResponseStatus = "WARNING"
)

type JobRunResponseTriggeredBy string

const (
	JobRunResponseTriggeredBySchedule JobRunResponseTriggeredBy = "schedule"
	JobRunResponseTriggeredByManual   JobRunResponseTriggeredBy = "manual"
)

type JobNewParams struct {
	AgentName   string                  `json:"agent_name" api:"required"`
	Name        string                  `json:"name" api:"required"`
	Description param.Opt[string]       `json:"description,omitzero"`
	DisplayName param.Opt[string]       `json:"display_name,omitzero"`
	Destination JobNewParamsDestination `json:"destination,omitzero"`
	Inputs      JobNewParamsInputs      `json:"inputs,omitzero"`
	Schedule    JobNewParamsSchedule    `json:"schedule,omitzero"`
	paramObj
}

func (r JobNewParams) MarshalJSON() (data []byte, err error) {
	type shadow JobNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Path, Type are required.
type JobNewParamsDestination struct {
	Path string `json:"path" api:"required"`
	// Any of "file", "s3".
	Type string `json:"type,omitzero" api:"required"`
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

// The property Type is required.
type JobNewParamsInputs struct {
	// Any of "s3", "inline", "file".
	Type     string            `json:"type,omitzero" api:"required"`
	FilePath param.Opt[string] `json:"file_path,omitzero"`
	Data     []map[string]any  `json:"data,omitzero"`
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

// The properties Cron, Enabled are required.
type JobNewParamsSchedule struct {
	Cron    string `json:"cron" api:"required"`
	Enabled bool   `json:"enabled" api:"required"`
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
	Description param.Opt[string]          `json:"description,omitzero"`
	DisplayName param.Opt[string]          `json:"display_name,omitzero"`
	Destination JobUpdateParamsDestination `json:"destination,omitzero"`
	Inputs      JobUpdateParamsInputs      `json:"inputs,omitzero"`
	Schedule    JobUpdateParamsSchedule    `json:"schedule,omitzero"`
	paramObj
}

func (r JobUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow JobUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Path, Type are required.
type JobUpdateParamsDestination struct {
	Path string `json:"path" api:"required"`
	// Any of "file", "s3".
	Type string `json:"type,omitzero" api:"required"`
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

// The property Type is required.
type JobUpdateParamsInputs struct {
	// Any of "s3", "inline", "file".
	Type     string            `json:"type,omitzero" api:"required"`
	FilePath param.Opt[string] `json:"file_path,omitzero"`
	Data     []map[string]any  `json:"data,omitzero"`
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

// The properties Cron, Enabled are required.
type JobUpdateParamsSchedule struct {
	Cron    string `json:"cron" api:"required"`
	Enabled bool   `json:"enabled" api:"required"`
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
	// Filter by agent name
	AgentName param.Opt[string] `query:"agent_name,omitzero" json:"-"`
	// Search by name or display name
	Q       param.Opt[string] `query:"q,omitzero" json:"-"`
	Page    param.Opt[int64]  `query:"page,omitzero" json:"-"`
	PerPage param.Opt[int64]  `query:"per_page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JobListParams]'s query parameters as `url.Values`.
func (r JobListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
