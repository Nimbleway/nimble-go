// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomnimblewaynimblego

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/Nimbleway/nimble-go/internal/apijson"
	"github.com/Nimbleway/nimble-go/internal/requestconfig"
	"github.com/Nimbleway/nimble-go/option"
	"github.com/Nimbleway/nimble-go/packages/respjson"
)

// JobRunArtifactService contains methods and other services that help with
// interacting with the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobRunArtifactService] method instead.
type JobRunArtifactService struct {
	Options []option.RequestOption
}

// NewJobRunArtifactService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewJobRunArtifactService(opts ...option.RequestOption) (r JobRunArtifactService) {
	r = JobRunArtifactService{}
	r.Options = opts
	return
}

// List Run Artifacts
func (r *JobRunArtifactService) List(ctx context.Context, runID string, opts ...option.RequestOption) (res *JobRunArtifactListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if runID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/jobs/runs/%s/artifacts", runID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get Run Artifact Download URL
func (r *JobRunArtifactService) DownloadURL(ctx context.Context, artifactID string, query JobRunArtifactDownloadURLParams, opts ...option.RequestOption) (res *JobRunArtifactDownloadURLResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.RunID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	if artifactID == "" {
		err = errors.New("missing required artifact_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/jobs/runs/%s/artifacts/%s/download-url", query.RunID, artifactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get Run Artifact
func (r *JobRunArtifactService) Get(ctx context.Context, artifactID string, query JobRunArtifactGetParams, opts ...option.RequestOption) (res *JobRunArtifactGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.RunID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	if artifactID == "" {
		err = errors.New("missing required artifact_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/jobs/runs/%s/artifacts/%s", query.RunID, artifactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Preview Run Artifact
func (r *JobRunArtifactService) Preview(ctx context.Context, artifactID string, query JobRunArtifactPreviewParams, opts ...option.RequestOption) (res *JobRunArtifactPreviewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.RunID == "" {
		err = errors.New("missing required run_id parameter")
		return nil, err
	}
	if artifactID == "" {
		err = errors.New("missing required artifact_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/jobs/runs/%s/artifacts/%s/preview", query.RunID, artifactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Artifacts produced by a run.
type JobRunArtifactListResponse struct {
	// Artifacts produced by the run.
	Items []JobRunArtifactListResponseItem `json:"items" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Items       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobRunArtifactListResponse) RawJSON() string { return r.JSON.raw }
func (r *JobRunArtifactListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A file produced by a run.
//
// Intentional subset of the bakery Artifact: `data_format` and `s3_path` are
// hidden from SDK consumers — internal storage details, not part of the public
// contract. Use the download-url endpoint to fetch the file. Bakery emits `id` as
// an int (crawlit native); the SDK contract is a string.
type JobRunArtifactListResponseItem struct {
	// Artifact identifier.
	ID string `json:"id" api:"required"`
	// When the artifact was created.
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	Description string    `json:"description" api:"required"`
	// Artifact type.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobRunArtifactListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *JobRunArtifactListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A pre-signed URL for downloading an artifact.
type JobRunArtifactDownloadURLResponse struct {
	// When the download URL expires.
	ExpiresAt time.Time `json:"expires_at" api:"required" format:"date-time"`
	// Pre-signed URL to download the artifact.
	URL string `json:"url" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpiresAt   respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobRunArtifactDownloadURLResponse) RawJSON() string { return r.JSON.raw }
func (r *JobRunArtifactDownloadURLResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A file produced by a run.
//
// Intentional subset of the bakery Artifact: `data_format` and `s3_path` are
// hidden from SDK consumers — internal storage details, not part of the public
// contract. Use the download-url endpoint to fetch the file. Bakery emits `id` as
// an int (crawlit native); the SDK contract is a string.
type JobRunArtifactGetResponse struct {
	// Artifact identifier.
	ID string `json:"id" api:"required"`
	// When the artifact was created.
	CreatedAt   time.Time `json:"created_at" api:"required" format:"date-time"`
	Description string    `json:"description" api:"required"`
	// Artifact type.
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobRunArtifactGetResponse) RawJSON() string { return r.JSON.raw }
func (r *JobRunArtifactGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A tabular preview of an artifact's contents.
type JobRunArtifactPreviewResponse struct {
	// Column names in the preview.
	Columns []string `json:"columns" api:"required"`
	// Total number of rows in the artifact.
	RowCount int64 `json:"row_count" api:"required"`
	// Sample rows from the artifact.
	Rows []map[string]any `json:"rows" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Columns     respjson.Field
		RowCount    respjson.Field
		Rows        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobRunArtifactPreviewResponse) RawJSON() string { return r.JSON.raw }
func (r *JobRunArtifactPreviewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobRunArtifactDownloadURLParams struct {
	RunID string `path:"run_id" api:"required" json:"-"`
	paramObj
}

type JobRunArtifactGetParams struct {
	RunID string `path:"run_id" api:"required" json:"-"`
	paramObj
}

type JobRunArtifactPreviewParams struct {
	RunID string `path:"run_id" api:"required" json:"-"`
	paramObj
}
