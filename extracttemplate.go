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
	shimjson "github.com/Nimbleway/nimble-go/internal/encoding/json"
	"github.com/Nimbleway/nimble-go/internal/requestconfig"
	"github.com/Nimbleway/nimble-go/option"
	"github.com/Nimbleway/nimble-go/packages/param"
	"github.com/Nimbleway/nimble-go/packages/respjson"
	"github.com/Nimbleway/nimble-go/shared/constant"
)

// ExtractTemplateService contains methods and other services that help with
// interacting with the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtractTemplateService] method instead.
type ExtractTemplateService struct {
	Options     []option.RequestOption
	Generations ExtractTemplateGenerationService
	Versions    ExtractTemplateVersionService
}

// NewExtractTemplateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewExtractTemplateService(opts ...option.RequestOption) (r ExtractTemplateService) {
	r = ExtractTemplateService{}
	r.Options = opts
	r.Generations = NewExtractTemplateGenerationService(opts...)
	r.Versions = NewExtractTemplateVersionService(opts...)
	return
}

// Patch Extract Template Public V2
func (r *ExtractTemplateService) Update(ctx context.Context, extractTemplateName string, body ExtractTemplateUpdateParams, opts ...option.RequestOption) (res *ExtractTemplateUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if extractTemplateName == "" {
		err = errors.New("missing required extract_template_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/extract/templates/%s", extractTemplateName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List Extract Templates Public V2
func (r *ExtractTemplateService) List(ctx context.Context, query ExtractTemplateListParams, opts ...option.RequestOption) (res *ExtractTemplateListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/extract/templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete Extract Template Public V2
func (r *ExtractTemplateService) Delete(ctx context.Context, extractTemplateName string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if extractTemplateName == "" {
		err = errors.New("missing required extract_template_name parameter")
		return err
	}
	path := fmt.Sprintf("v2/extract/templates/%s", extractTemplateName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Execute Extraction Template Async Endpoint
func (r *ExtractTemplateService) Async(ctx context.Context, body ExtractTemplateAsyncParams, opts ...option.RequestOption) (res *ExtractTemplateAsyncResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/extract/templates/async"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Execute Extraction Template Batch Endpoint
func (r *ExtractTemplateService) Batch(ctx context.Context, body ExtractTemplateBatchParams, opts ...option.RequestOption) (res *ExtractTemplateBatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/extract/templates/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get Extract Template Public V2
func (r *ExtractTemplateService) Get(ctx context.Context, extractTemplateName string, opts ...option.RequestOption) (res *ExtractTemplateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if extractTemplateName == "" {
		err = errors.New("missing required extract_template_name parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/extract/templates/%s", extractTemplateName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Execute Extraction Template Realtime Endpoint
func (r *ExtractTemplateService) Run(ctx context.Context, body ExtractTemplateRunParams, opts ...option.RequestOption) (res *ExtractTemplateRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/extract/templates/run"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ExtractTemplateUpdateResponse struct {
	// Unique extract template identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the extract template was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Stable extract template name.
	Name string `json:"name" api:"required"`
	// When the extract template was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Published version details, when available.
	PublishedVersion ExtractTemplateUpdateResponsePublishedVersion `json:"published_version" api:"nullable"`
	// Identifier of the published version.
	PublishedVersionID string `json:"published_version_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Name               respjson.Field
		UpdatedAt          respjson.Field
		PublishedVersion   respjson.Field
		PublishedVersionID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Published version details, when available.
type ExtractTemplateUpdateResponsePublishedVersion struct {
	// Unique extract template version identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the version was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// JSON schema describing accepted input parameters.
	InputSchema map[string]any `json:"input_schema" api:"required"`
	// Metadata associated with this version.
	Metadata ExtractTemplateUpdateResponsePublishedVersionMetadata `json:"metadata" api:"required"`
	// Extract template name this version belongs to.
	Name string `json:"name" api:"required"`
	// JSON schema describing extracted output.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Monotonic version number for the extract template.
	VersionNumber int64 `json:"version_number" api:"required"`
	// Sample input and output pairs for the version.
	Samples []ExtractTemplateUpdateResponsePublishedVersionSample `json:"samples" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		InputSchema   respjson.Field
		Metadata      respjson.Field
		Name          respjson.Field
		OutputSchema  respjson.Field
		VersionNumber respjson.Field
		Samples       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateUpdateResponsePublishedVersion) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateUpdateResponsePublishedVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata associated with this version.
type ExtractTemplateUpdateResponsePublishedVersionMetadata struct {
	// Data source associated with the version.
	DataSource string `json:"data_source" api:"nullable"`
	// Version description shown to users.
	Description string `json:"description" api:"nullable"`
	// Human-friendly version display name.
	DisplayName string `json:"display_name" api:"nullable"`
	// Domain associated with the version.
	Domain string `json:"domain" api:"nullable"`
	// Entity type produced by the version.
	EntityType string `json:"entity_type" api:"nullable"`
	// Tags associated with the version.
	Tags []string `json:"tags"`
	// Business vertical associated with the version.
	Vertical string `json:"vertical" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataSource  respjson.Field
		Description respjson.Field
		DisplayName respjson.Field
		Domain      respjson.Field
		EntityType  respjson.Field
		Tags        respjson.Field
		Vertical    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateUpdateResponsePublishedVersionMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateUpdateResponsePublishedVersionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateUpdateResponsePublishedVersionSample struct {
	// Sample input parameters for the version.
	Input any `json:"input"`
	// Sample output produced by the version.
	Output any `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Input       respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateUpdateResponsePublishedVersionSample) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateUpdateResponsePublishedVersionSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateListResponse struct {
	// Items returned in this page.
	Items []ExtractTemplateListResponseItem `json:"items" api:"required"`
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
func (r ExtractTemplateListResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateListResponseItem struct {
	// Unique extract template identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the extract template was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Stable extract template name.
	Name string `json:"name" api:"required"`
	// When the extract template was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Published version details, when available.
	PublishedVersion ExtractTemplateListResponseItemPublishedVersion `json:"published_version" api:"nullable"`
	// Identifier of the published version.
	PublishedVersionID string `json:"published_version_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Name               respjson.Field
		UpdatedAt          respjson.Field
		PublishedVersion   respjson.Field
		PublishedVersionID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Published version details, when available.
type ExtractTemplateListResponseItemPublishedVersion struct {
	// Unique extract template version identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the version was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// JSON schema describing accepted input parameters.
	InputSchema map[string]any `json:"input_schema" api:"required"`
	// Metadata associated with this version.
	Metadata ExtractTemplateListResponseItemPublishedVersionMetadata `json:"metadata" api:"required"`
	// Extract template name this version belongs to.
	Name string `json:"name" api:"required"`
	// JSON schema describing extracted output.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Monotonic version number for the extract template.
	VersionNumber int64 `json:"version_number" api:"required"`
	// Sample input and output pairs for the version.
	Samples []ExtractTemplateListResponseItemPublishedVersionSample `json:"samples" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		InputSchema   respjson.Field
		Metadata      respjson.Field
		Name          respjson.Field
		OutputSchema  respjson.Field
		VersionNumber respjson.Field
		Samples       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateListResponseItemPublishedVersion) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateListResponseItemPublishedVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata associated with this version.
type ExtractTemplateListResponseItemPublishedVersionMetadata struct {
	// Data source associated with the version.
	DataSource string `json:"data_source" api:"nullable"`
	// Version description shown to users.
	Description string `json:"description" api:"nullable"`
	// Human-friendly version display name.
	DisplayName string `json:"display_name" api:"nullable"`
	// Domain associated with the version.
	Domain string `json:"domain" api:"nullable"`
	// Entity type produced by the version.
	EntityType string `json:"entity_type" api:"nullable"`
	// Tags associated with the version.
	Tags []string `json:"tags"`
	// Business vertical associated with the version.
	Vertical string `json:"vertical" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataSource  respjson.Field
		Description respjson.Field
		DisplayName respjson.Field
		Domain      respjson.Field
		EntityType  respjson.Field
		Tags        respjson.Field
		Vertical    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateListResponseItemPublishedVersionMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateListResponseItemPublishedVersionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateListResponseItemPublishedVersionSample struct {
	// Sample input parameters for the version.
	Input any `json:"input"`
	// Sample output produced by the version.
	Output any `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Input       respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateListResponseItemPublishedVersionSample) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateListResponseItemPublishedVersionSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateAsyncResponse struct {
	Status constant.Success `json:"status" default:"success"`
	Task   map[string]any   `json:"task" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Task        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateAsyncResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateAsyncResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response when a batch of extract tasks is created successfully.
type ExtractTemplateBatchResponse struct {
	// Unique identifier for the batch.
	BatchID string `json:"batch_id" api:"required"`
	// Number of tasks in the batch.
	BatchSize float64 `json:"batch_size" api:"required"`
	// List of created tasks.
	Tasks []ExtractTemplateBatchResponseTask `json:"tasks" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BatchID     respjson.Field
		BatchSize   respjson.Field
		Tasks       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateBatchResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateBatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateBatchResponseTask struct {
	// Unique task identifier.
	ID    string `json:"id" api:"required"`
	Query any    `json:"_query" api:"required"`
	// Timestamp when the task was created.
	CreatedAt string `json:"created_at" api:"required"`
	// Original input data for the task.
	Input any `json:"input" api:"required"`
	// Current state of the task.
	//
	// Any of "pending", "queued", "in_progress", "success", "error".
	State string `json:"state" api:"required"`
	// URL for checking the task status.
	StatusURL string `json:"status_url" api:"required" format:"uri"`
	// Account name that owns the task.
	AccountName string `json:"account_name"`
	// Any of "web", "serp", "ecommerce", "social", "media", "agent", "extract",
	// "fast-serp", "labs".
	APIType string `json:"api_type"`
	// Batch ID if this task is part of a batch.
	BatchID string `json:"batch_id" api:"nullable"`
	// URL for downloading the task results.
	DownloadURL string `json:"download_url" api:"nullable" format:"uri"`
	// Error message if the task failed.
	Error string `json:"error" api:"nullable"`
	// Classification of the error type.
	ErrorType string `json:"error_type" api:"nullable"`
	// Timestamp when the task was last modified.
	ModifiedAt string `json:"modified_at"`
	// Storage location of the output data.
	OutputURL string `json:"output_url" api:"nullable"`
	// Queue name the task was submitted to.
	Queue string `json:"queue"`
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
		Queue       respjson.Field
		StatusCode  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateBatchResponseTask) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateBatchResponseTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateGetResponse struct {
	// Unique extract template identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the extract template was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// Stable extract template name.
	Name string `json:"name" api:"required"`
	// When the extract template was last updated.
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Published version details, when available.
	PublishedVersion ExtractTemplateGetResponsePublishedVersion `json:"published_version" api:"nullable"`
	// Identifier of the published version.
	PublishedVersionID string `json:"published_version_id" api:"nullable" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Name               respjson.Field
		UpdatedAt          respjson.Field
		PublishedVersion   respjson.Field
		PublishedVersionID respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Published version details, when available.
type ExtractTemplateGetResponsePublishedVersion struct {
	// Unique extract template version identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the version was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// JSON schema describing accepted input parameters.
	InputSchema map[string]any `json:"input_schema" api:"required"`
	// Metadata associated with this version.
	Metadata ExtractTemplateGetResponsePublishedVersionMetadata `json:"metadata" api:"required"`
	// Extract template name this version belongs to.
	Name string `json:"name" api:"required"`
	// JSON schema describing extracted output.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Monotonic version number for the extract template.
	VersionNumber int64 `json:"version_number" api:"required"`
	// Sample input and output pairs for the version.
	Samples []ExtractTemplateGetResponsePublishedVersionSample `json:"samples" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		InputSchema   respjson.Field
		Metadata      respjson.Field
		Name          respjson.Field
		OutputSchema  respjson.Field
		VersionNumber respjson.Field
		Samples       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateGetResponsePublishedVersion) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateGetResponsePublishedVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata associated with this version.
type ExtractTemplateGetResponsePublishedVersionMetadata struct {
	// Data source associated with the version.
	DataSource string `json:"data_source" api:"nullable"`
	// Version description shown to users.
	Description string `json:"description" api:"nullable"`
	// Human-friendly version display name.
	DisplayName string `json:"display_name" api:"nullable"`
	// Domain associated with the version.
	Domain string `json:"domain" api:"nullable"`
	// Entity type produced by the version.
	EntityType string `json:"entity_type" api:"nullable"`
	// Tags associated with the version.
	Tags []string `json:"tags"`
	// Business vertical associated with the version.
	Vertical string `json:"vertical" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataSource  respjson.Field
		Description respjson.Field
		DisplayName respjson.Field
		Domain      respjson.Field
		EntityType  respjson.Field
		Tags        respjson.Field
		Vertical    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateGetResponsePublishedVersionMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateGetResponsePublishedVersionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateGetResponsePublishedVersionSample struct {
	// Sample input parameters for the version.
	Input any `json:"input"`
	// Sample output produced by the version.
	Output any `json:"output"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Input       respjson.Field
		Output      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateGetResponsePublishedVersionSample) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateGetResponsePublishedVersionSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateRunResponse struct {
	Data     ExtractTemplateRunResponseData     `json:"data" api:"required"`
	Metadata ExtractTemplateRunResponseMetadata `json:"metadata" api:"required"`
	// The status of the task.
	//
	// Any of "success", "skipped", "fatal", "error", "postponed", "ignored",
	// "rejected", "blocked".
	Status ExtractTemplateRunResponseStatus `json:"status" api:"required"`
	// Unique identifier for the task.
	TaskID string `json:"task_id" api:"required"`
	// The final URL.
	URL   string                          `json:"url" api:"required"`
	Debug ExtractTemplateRunResponseDebug `json:"debug"`
	// Pagination information if applicable.
	Pagination ExtractTemplateRunResponsePaginationUnion `json:"pagination"`
	// The HTTP status code of the task.
	StatusCode float64 `json:"status_code"`
	// List of warnings generated during the task.
	Warnings []string `json:"warnings"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Metadata    respjson.Field
		Status      respjson.Field
		TaskID      respjson.Field
		URL         respjson.Field
		Debug       respjson.Field
		Pagination  respjson.Field
		StatusCode  respjson.Field
		Warnings    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateRunResponseData struct {
	// Browser actions execution results. Present only when browser_actions were
	// specified in the request.
	BrowserActions ExtractTemplateRunResponseDataBrowserActions `json:"browser_actions"`
	// The cookies collected from browser actions during the task.
	Cookies []any `json:"cookies"`
	// The evaluation results from browser actions during the task.
	Eval []any `json:"eval"`
	// The http requests from browser actions made during the task.
	Fetch []any `json:"fetch"`
	// The headers received during the task.
	Headers map[string]string `json:"headers"`
	// The HTML content of the page.
	HTML string `json:"html"`
	// List of all unique URLs found on the page.
	Links []string `json:"links"`
	// The Markdown version of the HTML content.
	Markdown string `json:"markdown"`
	// The network capture data collected during the task.
	NetworkCapture []ExtractTemplateRunResponseDataNetworkCapture `json:"network_capture"`
	// Individual HTML content of each pagination page, before merging.
	PagesHTML []string `json:"pages_html"`
	// The parsing results extracted from the HTML & network content.
	Parsing ExtractTemplateRunResponseDataParsingUnion `json:"parsing"`
	// The list of redirects that occurred during the task.
	Redirects []ExtractTemplateRunResponseDataRedirect `json:"redirects"`
	// Screenshots taken during the task, from browser actions, or the screenshot
	// format.
	Screenshots []any `json:"screenshots"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BrowserActions respjson.Field
		Cookies        respjson.Field
		Eval           respjson.Field
		Fetch          respjson.Field
		Headers        respjson.Field
		HTML           respjson.Field
		Links          respjson.Field
		Markdown       respjson.Field
		NetworkCapture respjson.Field
		PagesHTML      respjson.Field
		Parsing        respjson.Field
		Redirects      respjson.Field
		Screenshots    respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponseData) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateRunResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Browser actions execution results. Present only when browser_actions were
// specified in the request.
type ExtractTemplateRunResponseDataBrowserActions struct {
	Results       []ExtractTemplateRunResponseDataBrowserActionsResult `json:"results" api:"required"`
	Success       bool                                                 `json:"success" api:"required"`
	TotalDuration float64                                              `json:"total_duration" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results       respjson.Field
		Success       respjson.Field
		TotalDuration respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponseDataBrowserActions) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateRunResponseDataBrowserActions) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateRunResponseDataBrowserActionsResult struct {
	Duration float64 `json:"duration" api:"required"`
	// Any of "goto", "wait", "wait_for_element", "wait_for_navigation", "click",
	// "fill", "press", "scroll", "auto_scroll", "screenshot", "get_cookies", "eval",
	// "fetch".
	Name string `json:"name" api:"required"`
	// Any of "no-run", "in-progress", "done", "error", "skipped".
	Status string `json:"status" api:"required"`
	Error  string `json:"error"`
	Result any    `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Duration    respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		Error       respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponseDataBrowserActionsResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateRunResponseDataBrowserActionsResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateRunResponseDataNetworkCapture struct {
	Filter       ExtractTemplateRunResponseDataNetworkCaptureFilter   `json:"filter" api:"required"`
	Results      []ExtractTemplateRunResponseDataNetworkCaptureResult `json:"results" api:"required"`
	ErrorMessage string                                               `json:"errorMessage"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Filter       respjson.Field
		Results      respjson.Field
		ErrorMessage respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponseDataNetworkCapture) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateRunResponseDataNetworkCapture) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateRunResponseDataNetworkCaptureFilter struct {
	Validation           bool    `json:"validation" api:"required"`
	WaitForRequestsCount float64 `json:"wait_for_requests_count" api:"required"`
	// Any of "GET", "HEAD", "POST", "PUT", "DELETE", "CONNECT", "OPTIONS", "TRACE",
	// "PATCH".
	Method string `json:"method"`
	// Resource type for network capture filtering
	ResourceType                ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeUnion `json:"resource_type"`
	StatusCode                  ExtractTemplateRunResponseDataNetworkCaptureFilterStatusCodeUnion   `json:"status_code"`
	URL                         ExtractTemplateRunResponseDataNetworkCaptureFilterURL               `json:"url"`
	WaitForRequestsCountTimeout float64                                                             `json:"wait_for_requests_count_timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Validation                  respjson.Field
		WaitForRequestsCount        respjson.Field
		Method                      respjson.Field
		ResourceType                respjson.Field
		StatusCode                  respjson.Field
		URL                         respjson.Field
		WaitForRequestsCountTimeout respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponseDataNetworkCaptureFilter) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateRunResponseDataNetworkCaptureFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeUnion contains all
// possible properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid:
// OfExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString
// OfExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray]
type ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray []string `json:",inline"`
	JSON                                                                           struct {
		OfExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString         respjson.Field
		OfExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray respjson.Field
		raw                                                                            string
	} `json:"-"`
}

func (u ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeUnion) AsExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeUnion) AsExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeArrayItemArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Resource type for network capture filtering
type ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString string

const (
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringDocument           ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "document"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringStylesheet         ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "stylesheet"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringImage              ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "image"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringMedia              ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "media"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringFont               ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "font"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringScript             ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "script"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringTexttrack          ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "texttrack"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringXhr                ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "xhr"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringFetch              ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "fetch"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringPrefetch           ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "prefetch"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringEventsource        ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "eventsource"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringWebsocket          ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "websocket"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringManifest           ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "manifest"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringSignedexchange     ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "signedexchange"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringPing               ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "ping"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringCspviolationreport ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "cspviolationreport"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringPreflight          ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "preflight"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringOther              ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "other"
	ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeStringFedcm              ExtractTemplateRunResponseDataNetworkCaptureFilterResourceTypeString = "fedcm"
)

// ExtractTemplateRunResponseDataNetworkCaptureFilterStatusCodeUnion contains all
// possible properties and values from [float64], [[]float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfFloatArray]
type ExtractTemplateRunResponseDataNetworkCaptureFilterStatusCodeUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [[]float64] instead of an object.
	OfFloatArray []float64 `json:",inline"`
	JSON         struct {
		OfFloat      respjson.Field
		OfFloatArray respjson.Field
		raw          string
	} `json:"-"`
}

func (u ExtractTemplateRunResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractTemplateRunResponseDataNetworkCaptureFilterStatusCodeUnion) AsFloatArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractTemplateRunResponseDataNetworkCaptureFilterStatusCodeUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *ExtractTemplateRunResponseDataNetworkCaptureFilterStatusCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateRunResponseDataNetworkCaptureFilterURL struct {
	// Any of "exact", "contains".
	Type  string `json:"type" api:"required"`
	Value string `json:"value" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponseDataNetworkCaptureFilterURL) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateRunResponseDataNetworkCaptureFilterURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateRunResponseDataNetworkCaptureResult struct {
	Request  ExtractTemplateRunResponseDataNetworkCaptureResultRequest  `json:"request" api:"required"`
	Response ExtractTemplateRunResponseDataNetworkCaptureResultResponse `json:"response" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Request     respjson.Field
		Response    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponseDataNetworkCaptureResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateRunResponseDataNetworkCaptureResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateRunResponseDataNetworkCaptureResultRequest struct {
	Headers map[string]string `json:"headers" api:"required"`
	Method  string            `json:"method" api:"required"`
	// Resource type for network capture filtering
	//
	// Any of "document", "stylesheet", "image", "media", "font", "script",
	// "texttrack", "xhr", "fetch", "prefetch", "eventsource", "websocket", "manifest",
	// "signedexchange", "ping", "cspviolationreport", "preflight", "other", "fedcm".
	ResourceType string `json:"resource_type" api:"required"`
	URL          string `json:"url" api:"required"`
	Body         string `json:"body"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headers      respjson.Field
		Method       respjson.Field
		ResourceType respjson.Field
		URL          respjson.Field
		Body         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponseDataNetworkCaptureResultRequest) RawJSON() string {
	return r.JSON.raw
}
func (r *ExtractTemplateRunResponseDataNetworkCaptureResultRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateRunResponseDataNetworkCaptureResultResponse struct {
	Body    string            `json:"body" api:"required"`
	Headers map[string]string `json:"headers" api:"required"`
	// Any of "none", "base64".
	Serialization string  `json:"serialization" api:"required"`
	Status        float64 `json:"status" api:"required"`
	StatusText    string  `json:"status_text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body          respjson.Field
		Headers       respjson.Field
		Serialization respjson.Field
		Status        respjson.Field
		StatusText    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponseDataNetworkCaptureResultResponse) RawJSON() string {
	return r.JSON.raw
}
func (r *ExtractTemplateRunResponseDataNetworkCaptureResultResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractTemplateRunResponseDataParsingUnion contains all possible properties and
// values from [ExtractTemplateRunResponseDataParsingParsingSuccessResult],
// [ExtractTemplateRunResponseDataParsingParsingErrorResult], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractTemplateRunResponseDataParsingMapItem]
type ExtractTemplateRunResponseDataParsingUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfExtractTemplateRunResponseDataParsingMapItem any `json:",inline"`
	// This field is from variant
	// [ExtractTemplateRunResponseDataParsingParsingSuccessResult].
	Entities map[string]any `json:"entities"`
	Status   string         `json:"status"`
	// This field is from variant
	// [ExtractTemplateRunResponseDataParsingParsingErrorResult].
	Error string `json:"error"`
	JSON  struct {
		OfExtractTemplateRunResponseDataParsingMapItem respjson.Field
		Entities                                       respjson.Field
		Status                                         respjson.Field
		Error                                          respjson.Field
		raw                                            string
	} `json:"-"`
}

func (u ExtractTemplateRunResponseDataParsingUnion) AsExtractTemplateRunResponseDataParsingParsingSuccessResult() (v ExtractTemplateRunResponseDataParsingParsingSuccessResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractTemplateRunResponseDataParsingUnion) AsExtractTemplateRunResponseDataParsingParsingErrorResult() (v ExtractTemplateRunResponseDataParsingParsingErrorResult) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractTemplateRunResponseDataParsingUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractTemplateRunResponseDataParsingUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractTemplateRunResponseDataParsingUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateRunResponseDataParsingParsingSuccessResult struct {
	Entities map[string]any   `json:"entities" api:"required"`
	Status   constant.Success `json:"status" default:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Entities    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponseDataParsingParsingSuccessResult) RawJSON() string {
	return r.JSON.raw
}
func (r *ExtractTemplateRunResponseDataParsingParsingSuccessResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateRunResponseDataParsingParsingErrorResult struct {
	Error  string         `json:"error" api:"required"`
	Status constant.Error `json:"status" default:"error"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponseDataParsingParsingErrorResult) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateRunResponseDataParsingParsingErrorResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateRunResponseDataRedirect struct {
	StatusCode float64 `json:"status_code" api:"required"`
	URL        string  `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		StatusCode  respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponseDataRedirect) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateRunResponseDataRedirect) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateRunResponseMetadata struct {
	// The name of the agent used for the query.
	Agent string `json:"agent"`
	// The driver used for the task.
	Driver string `json:"driver"`
	// The localization identifier for the query.
	LocalizationID string `json:"localization_id"`
	// The duration in milliseconds of the query processing.
	QueryDuration float64 `json:"query_duration"`
	// The time when the query was received.
	QueryTime string `json:"query_time"`
	// Additional response parameters.
	ResponseParameters any `json:"response_parameters"`
	// A tag associated with the query.
	Tag string `json:"tag"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Agent              respjson.Field
		Driver             respjson.Field
		LocalizationID     respjson.Field
		QueryDuration      respjson.Field
		QueryTime          respjson.Field
		ResponseParameters respjson.Field
		Tag                respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateRunResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the task.
type ExtractTemplateRunResponseStatus string

const (
	ExtractTemplateRunResponseStatusSuccess   ExtractTemplateRunResponseStatus = "success"
	ExtractTemplateRunResponseStatusSkipped   ExtractTemplateRunResponseStatus = "skipped"
	ExtractTemplateRunResponseStatusFatal     ExtractTemplateRunResponseStatus = "fatal"
	ExtractTemplateRunResponseStatusError     ExtractTemplateRunResponseStatus = "error"
	ExtractTemplateRunResponseStatusPostponed ExtractTemplateRunResponseStatus = "postponed"
	ExtractTemplateRunResponseStatusIgnored   ExtractTemplateRunResponseStatus = "ignored"
	ExtractTemplateRunResponseStatusRejected  ExtractTemplateRunResponseStatus = "rejected"
	ExtractTemplateRunResponseStatusBlocked   ExtractTemplateRunResponseStatus = "blocked"
)

type ExtractTemplateRunResponseDebug struct {
	// Performance metrics collected during the task.
	PerformanceMetrics map[string]float64 `json:"performance_metrics"`
	// Total bytes used by the proxy during the task.
	ProxyTotalBytesUsage float64 `json:"proxy_total_bytes_usage"`
	// The transformed output after applying any transformations.
	TransformedOutput any `json:"transformed_output"`
	// The userbrowser instance using during the task.
	Userbrowser any `json:"userbrowser"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PerformanceMetrics   respjson.Field
		ProxyTotalBytesUsage respjson.Field
		TransformedOutput    respjson.Field
		Userbrowser          respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponseDebug) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateRunResponseDebug) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ExtractTemplateRunResponsePaginationUnion contains all possible properties and
// values from [ExtractTemplateRunResponsePaginationNextPageParams],
// [[]ExtractTemplateRunResponsePaginationArrayItem].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfExtractTemplateRunResponsePaginationArray]
type ExtractTemplateRunResponsePaginationUnion struct {
	// This field will be present if the value is a
	// [[]ExtractTemplateRunResponsePaginationArrayItem] instead of an object.
	OfExtractTemplateRunResponsePaginationArray []ExtractTemplateRunResponsePaginationArrayItem `json:",inline"`
	// This field is from variant [ExtractTemplateRunResponsePaginationNextPageParams].
	NextPageParams map[string]any `json:"next_page_params"`
	JSON           struct {
		OfExtractTemplateRunResponsePaginationArray respjson.Field
		NextPageParams                              respjson.Field
		raw                                         string
	} `json:"-"`
}

func (u ExtractTemplateRunResponsePaginationUnion) AsExtractTemplateRunResponsePaginationNextPageParams() (v ExtractTemplateRunResponsePaginationNextPageParams) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ExtractTemplateRunResponsePaginationUnion) AsExtractTemplateRunResponsePaginationArray() (v []ExtractTemplateRunResponsePaginationArrayItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ExtractTemplateRunResponsePaginationUnion) RawJSON() string { return u.JSON.raw }

func (r *ExtractTemplateRunResponsePaginationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateRunResponsePaginationNextPageParams struct {
	NextPageParams map[string]any `json:"next_page_params" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponsePaginationNextPageParams) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateRunResponsePaginationNextPageParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateRunResponsePaginationArrayItem struct {
	NextPageParams map[string]any `json:"next_page_params" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextPageParams respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateRunResponsePaginationArrayItem) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateRunResponsePaginationArrayItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateUpdateParams struct {
	// A JSON Patch document per RFC 6902 — a JSON array of patch operations.
	Body []ExtractTemplateUpdateParamsBody
	paramObj
}

func (r ExtractTemplateUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *ExtractTemplateUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single JSON Patch operation per RFC 6902.
//
// The properties Op, Path are required.
type ExtractTemplateUpdateParamsBody struct {
	// Any of "add", "remove", "replace", "move", "copy", "test".
	Op    string            `json:"op,omitzero" api:"required"`
	Path  string            `json:"path" api:"required"`
	From  param.Opt[string] `json:"from,omitzero"`
	Value any               `json:"value,omitzero"`
	paramObj
}

func (r ExtractTemplateUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow ExtractTemplateUpdateParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractTemplateUpdateParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ExtractTemplateUpdateParamsBody](
		"op", "add", "remove", "replace", "move", "copy", "test",
	)
}

type ExtractTemplateListParams struct {
	Limit  param.Opt[int64] `query:"limit,omitzero" json:"-"`
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ExtractTemplateListParams]'s query parameters as
// `url.Values`.
func (r ExtractTemplateListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ExtractTemplateAsyncParams struct {
	Params   map[string]any `json:"params,omitzero" api:"required"`
	Template string         `json:"template" api:"required"`
	// URL to call back when async operation completes
	CallbackURL  param.Opt[string] `json:"callback_url,omitzero"`
	Localization param.Opt[bool]   `json:"localization,omitzero"`
	// Whether to compress stored data
	StorageCompress param.Opt[bool] `json:"storage_compress,omitzero"`
	// Custom name for the stored object
	StorageObjectName param.Opt[string] `json:"storage_object_name,omitzero"`
	// Type of storage to use for results
	StorageType param.Opt[string] `json:"storage_type,omitzero"`
	// URL for storage location
	StorageURL param.Opt[string] `json:"storage_url,omitzero"`
	// Response formats to include. All disabled by default.
	//
	// Any of "html", "markdown", "screenshot", "headers", "links".
	Formats []string `json:"formats,omitzero"`
	paramObj
}

func (r ExtractTemplateAsyncParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractTemplateAsyncParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractTemplateAsyncParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateBatchParams struct {
	Inputs       []ExtractTemplateBatchParamsInput      `json:"inputs,omitzero" api:"required"`
	SharedInputs ExtractTemplateBatchParamsSharedInputs `json:"shared_inputs,omitzero" api:"required"`
	paramObj
}

func (r ExtractTemplateBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractTemplateBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractTemplateBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateBatchParamsInput struct {
	Localization param.Opt[bool] `json:"localization,omitzero"`
	// Response formats to include. All disabled by default.
	//
	// Any of "html", "markdown", "screenshot", "headers", "links".
	Formats []string       `json:"formats,omitzero"`
	Params  map[string]any `json:"params,omitzero"`
	paramObj
}

func (r ExtractTemplateBatchParamsInput) MarshalJSON() (data []byte, err error) {
	type shadow ExtractTemplateBatchParamsInput
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractTemplateBatchParamsInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Template is required.
type ExtractTemplateBatchParamsSharedInputs struct {
	Template     string          `json:"template" api:"required"`
	Localization param.Opt[bool] `json:"localization,omitzero"`
	// Response formats to include. All disabled by default.
	//
	// Any of "html", "markdown", "screenshot", "headers", "links".
	Formats []string       `json:"formats,omitzero"`
	Params  map[string]any `json:"params,omitzero"`
	paramObj
}

func (r ExtractTemplateBatchParamsSharedInputs) MarshalJSON() (data []byte, err error) {
	type shadow ExtractTemplateBatchParamsSharedInputs
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractTemplateBatchParamsSharedInputs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateRunParams struct {
	Params       map[string]any  `json:"params,omitzero" api:"required"`
	Template     string          `json:"template" api:"required"`
	Localization param.Opt[bool] `json:"localization,omitzero"`
	// Response formats to include. All disabled by default.
	//
	// Any of "html", "markdown", "screenshot", "headers", "links".
	Formats []string `json:"formats,omitzero"`
	paramObj
}

func (r ExtractTemplateRunParams) MarshalJSON() (data []byte, err error) {
	type shadow ExtractTemplateRunParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractTemplateRunParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
