// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomnimblewaynimblego

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/Nimbleway/nimble-go/internal/apijson"
	"github.com/Nimbleway/nimble-go/internal/apiquery"
	"github.com/Nimbleway/nimble-go/internal/requestconfig"
	"github.com/Nimbleway/nimble-go/option"
	"github.com/Nimbleway/nimble-go/packages/param"
	"github.com/Nimbleway/nimble-go/packages/respjson"
)

// TaskService contains methods and other services that help with interacting with
// the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTaskService] method instead.
type TaskService struct {
	Options []option.RequestOption
}

// NewTaskService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTaskService(opts ...option.RequestOption) (r TaskService) {
	r = TaskService{}
	r.Options = opts
	return
}

// Retrieve a paginated list of tasks for the authenticated account.
func (r *TaskService) List(ctx context.Context, query TaskListParams, opts ...option.RequestOption) (res *TaskListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tasks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Retrieve the details of a specific task by its ID.
func (r *TaskService) Get(ctx context.Context, taskID string, opts ...option.RequestOption) (res *TaskGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return
	}
	path := fmt.Sprintf("v1/tasks/%s", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve the results of a completed task.
func (r *TaskService) Results(ctx context.Context, taskID string, opts ...option.RequestOption) (res *TaskResultsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return
	}
	path := fmt.Sprintf("v1/tasks/%s/results", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Paginated list of tasks response.
type TaskListResponse struct {
	// Array of task objects.
	Data       []TaskListResponseData     `json:"data" api:"required"`
	Pagination TaskListResponsePagination `json:"pagination" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskListResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskListResponseData struct {
	// Unique task identifier.
	ID    string `json:"id" api:"required"`
	Query any    `json:"_query" api:"required"`
	// Timestamp when the task was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Original input data for the task.
	Input any `json:"input" api:"required"`
	// Current state of the task.
	//
	// Any of "pending", "success", "error".
	State string `json:"state" api:"required"`
	// URL for checking the task status.
	StatusURL string `json:"status_url" api:"required" format:"uri"`
	// Account name that owns the task.
	AccountName string `json:"account_name"`
	// Any of "web", "serp", "ecommerce", "social", "media", "agent", "extract".
	APIType string `json:"api_type"`
	// Batch ID if this task is part of a batch.
	BatchID string `json:"batch_id"`
	// URL for downloading the task results.
	DownloadURL string `json:"download_url" format:"uri"`
	// Error message if the task failed.
	Error string `json:"error"`
	// Classification of the error type.
	ErrorType string `json:"error_type"`
	// Timestamp when the task was last modified.
	ModifiedAt string `json:"modified_at"`
	// Storage location of the output data.
	OutputURL string `json:"output_url"`
	// HTTP status code from the task execution.
	StatusCode float64 `json:"status_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Query       respjson.Field
		CreatedAt   respjson.Field
		Input       respjson.Field
		State       respjson.Field
		StatusURL   respjson.Field
		AccountName respjson.Field
		APIType     respjson.Field
		BatchID     respjson.Field
		DownloadURL respjson.Field
		Error       respjson.Field
		ErrorType   respjson.Field
		ModifiedAt  respjson.Field
		OutputURL   respjson.Field
		StatusCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskListResponseData) RawJSON() string { return r.JSON.raw }
func (r *TaskListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskListResponsePagination struct {
	// Whether there are more tasks available.
	HasNext bool `json:"has_next" api:"required"`
	// Cursor to use for fetching the next page.
	NextCursor string `json:"next_cursor" api:"required"`
	// Total number of tasks.
	Total float64 `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasNext     respjson.Field
		NextCursor  respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskListResponsePagination) RawJSON() string { return r.JSON.raw }
func (r *TaskListResponsePagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing task details.
type TaskGetResponse struct {
	Task TaskGetResponseTask `json:"task" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Task        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TaskGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskGetResponseTask struct {
	// Unique task identifier.
	ID    string `json:"id" api:"required"`
	Query any    `json:"_query" api:"required"`
	// Timestamp when the task was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Original input data for the task.
	Input any `json:"input" api:"required"`
	// Current state of the task.
	//
	// Any of "pending", "success", "error".
	State string `json:"state" api:"required"`
	// URL for checking the task status.
	StatusURL string `json:"status_url" api:"required" format:"uri"`
	// Account name that owns the task.
	AccountName string `json:"account_name"`
	// Any of "web", "serp", "ecommerce", "social", "media", "agent", "extract".
	APIType string `json:"api_type"`
	// Batch ID if this task is part of a batch.
	BatchID string `json:"batch_id"`
	// URL for downloading the task results.
	DownloadURL string `json:"download_url" format:"uri"`
	// Error message if the task failed.
	Error string `json:"error"`
	// Classification of the error type.
	ErrorType string `json:"error_type"`
	// Timestamp when the task was last modified.
	ModifiedAt string `json:"modified_at"`
	// Storage location of the output data.
	OutputURL string `json:"output_url"`
	// HTTP status code from the task execution.
	StatusCode float64 `json:"status_code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Query       respjson.Field
		CreatedAt   respjson.Field
		Input       respjson.Field
		State       respjson.Field
		StatusURL   respjson.Field
		AccountName respjson.Field
		APIType     respjson.Field
		BatchID     respjson.Field
		DownloadURL respjson.Field
		Error       respjson.Field
		ErrorType   respjson.Field
		ModifiedAt  respjson.Field
		OutputURL   respjson.Field
		StatusCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskGetResponseTask) RawJSON() string { return r.JSON.raw }
func (r *TaskGetResponseTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskResultsResponse map[string]any

type TaskListParams struct {
	// Cursor for pagination. Use the next_cursor from the previous response.
	Cursor param.Opt[string] `query:"cursor,omitzero" format:"uuid" json:"-"`
	// Number of tasks to return per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TaskListParams]'s query parameters as `url.Values`.
func (r TaskListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
