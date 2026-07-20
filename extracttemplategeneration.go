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
	"github.com/Nimbleway/nimble-go/packages/param"
	"github.com/Nimbleway/nimble-go/packages/respjson"
)

// ExtractTemplateGenerationService contains methods and other services that help
// with interacting with the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExtractTemplateGenerationService] method instead.
type ExtractTemplateGenerationService struct {
	Options []option.RequestOption
}

// NewExtractTemplateGenerationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExtractTemplateGenerationService(opts ...option.RequestOption) (r ExtractTemplateGenerationService) {
	r = ExtractTemplateGenerationService{}
	r.Options = opts
	return
}

// Create Extract Template Generation Public V2
func (r *ExtractTemplateGenerationService) New(ctx context.Context, body ExtractTemplateGenerationNewParams, opts ...option.RequestOption) (res *ExtractTemplateGenerationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/extract/templates/generations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get Extract Template Generation Public V2
func (r *ExtractTemplateGenerationService) Get(ctx context.Context, generationID string, opts ...option.RequestOption) (res *ExtractTemplateGenerationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if generationID == "" {
		err = errors.New("missing required generation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v2/extract/templates/generations/%s", generationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ExtractTemplateGenerationNewResponse struct {
	// Unique extract template generation identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// Current generation status.
	Status string `json:"status" api:"required"`
	// When the generation completed.
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// When the generation was created.
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error message when generation failed.
	Error string `json:"error" api:"nullable"`
	// Generated version details, when available.
	GeneratedVersion ExtractTemplateGenerationNewResponseGeneratedVersion `json:"generated_version" api:"nullable"`
	// Identifier of the generated version.
	GeneratedVersionID string `json:"generated_version_id" api:"nullable" format:"uuid"`
	// Extract template name associated with the generation.
	Name string `json:"name" api:"nullable"`
	// Identifier of the version being refined.
	SourceVersionID string `json:"source_version_id" api:"nullable" format:"uuid"`
	// When the generation started executing.
	StartedAt time.Time `json:"started_at" api:"nullable" format:"date-time"`
	// Summary of the generation result.
	Summary string `json:"summary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Status             respjson.Field
		CompletedAt        respjson.Field
		CreatedAt          respjson.Field
		Error              respjson.Field
		GeneratedVersion   respjson.Field
		GeneratedVersionID respjson.Field
		Name               respjson.Field
		SourceVersionID    respjson.Field
		StartedAt          respjson.Field
		Summary            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateGenerationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateGenerationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Generated version details, when available.
type ExtractTemplateGenerationNewResponseGeneratedVersion struct {
	// Unique extract template version identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the version was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// JSON schema describing accepted input parameters.
	InputSchema map[string]any `json:"input_schema" api:"required"`
	// Metadata associated with this version.
	Metadata ExtractTemplateGenerationNewResponseGeneratedVersionMetadata `json:"metadata" api:"required"`
	// Extract template name this version belongs to.
	Name string `json:"name" api:"required"`
	// JSON schema describing extracted output.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Monotonic version number for the extract template.
	VersionNumber int64 `json:"version_number" api:"required"`
	// Sample input and output pairs for the version.
	Samples []ExtractTemplateGenerationNewResponseGeneratedVersionSample `json:"samples" api:"nullable"`
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
func (r ExtractTemplateGenerationNewResponseGeneratedVersion) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateGenerationNewResponseGeneratedVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata associated with this version.
type ExtractTemplateGenerationNewResponseGeneratedVersionMetadata struct {
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
func (r ExtractTemplateGenerationNewResponseGeneratedVersionMetadata) RawJSON() string {
	return r.JSON.raw
}
func (r *ExtractTemplateGenerationNewResponseGeneratedVersionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateGenerationNewResponseGeneratedVersionSample struct {
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
func (r ExtractTemplateGenerationNewResponseGeneratedVersionSample) RawJSON() string {
	return r.JSON.raw
}
func (r *ExtractTemplateGenerationNewResponseGeneratedVersionSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateGenerationGetResponse struct {
	// Unique extract template generation identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// Current generation status.
	Status string `json:"status" api:"required"`
	// When the generation completed.
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// When the generation was created.
	CreatedAt time.Time `json:"created_at" api:"nullable" format:"date-time"`
	// Error message when generation failed.
	Error string `json:"error" api:"nullable"`
	// Generated version details, when available.
	GeneratedVersion ExtractTemplateGenerationGetResponseGeneratedVersion `json:"generated_version" api:"nullable"`
	// Identifier of the generated version.
	GeneratedVersionID string `json:"generated_version_id" api:"nullable" format:"uuid"`
	// Extract template name associated with the generation.
	Name string `json:"name" api:"nullable"`
	// Identifier of the version being refined.
	SourceVersionID string `json:"source_version_id" api:"nullable" format:"uuid"`
	// When the generation started executing.
	StartedAt time.Time `json:"started_at" api:"nullable" format:"date-time"`
	// Summary of the generation result.
	Summary string `json:"summary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Status             respjson.Field
		CompletedAt        respjson.Field
		CreatedAt          respjson.Field
		Error              respjson.Field
		GeneratedVersion   respjson.Field
		GeneratedVersionID respjson.Field
		Name               respjson.Field
		SourceVersionID    respjson.Field
		StartedAt          respjson.Field
		Summary            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExtractTemplateGenerationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateGenerationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Generated version details, when available.
type ExtractTemplateGenerationGetResponseGeneratedVersion struct {
	// Unique extract template version identifier.
	ID string `json:"id" api:"required" format:"uuid"`
	// When the version was created.
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// JSON schema describing accepted input parameters.
	InputSchema map[string]any `json:"input_schema" api:"required"`
	// Metadata associated with this version.
	Metadata ExtractTemplateGenerationGetResponseGeneratedVersionMetadata `json:"metadata" api:"required"`
	// Extract template name this version belongs to.
	Name string `json:"name" api:"required"`
	// JSON schema describing extracted output.
	OutputSchema map[string]any `json:"output_schema" api:"required"`
	// Monotonic version number for the extract template.
	VersionNumber int64 `json:"version_number" api:"required"`
	// Sample input and output pairs for the version.
	Samples []ExtractTemplateGenerationGetResponseGeneratedVersionSample `json:"samples" api:"nullable"`
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
func (r ExtractTemplateGenerationGetResponseGeneratedVersion) RawJSON() string { return r.JSON.raw }
func (r *ExtractTemplateGenerationGetResponseGeneratedVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata associated with this version.
type ExtractTemplateGenerationGetResponseGeneratedVersionMetadata struct {
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
func (r ExtractTemplateGenerationGetResponseGeneratedVersionMetadata) RawJSON() string {
	return r.JSON.raw
}
func (r *ExtractTemplateGenerationGetResponseGeneratedVersionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateGenerationGetResponseGeneratedVersionSample struct {
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
func (r ExtractTemplateGenerationGetResponseGeneratedVersionSample) RawJSON() string {
	return r.JSON.raw
}
func (r *ExtractTemplateGenerationGetResponseGeneratedVersionSample) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExtractTemplateGenerationNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfCreateExtractTemplateGenerationRequestPublicV2 *ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateGenerationRequestPublicV2 `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfCreateExtractTemplateRefinementRequestPublicV2 *ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateRefinementRequestPublicV2 `json:",inline"`

	paramObj
}

func (u ExtractTemplateGenerationNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCreateExtractTemplateGenerationRequestPublicV2, u.OfCreateExtractTemplateRefinementRequestPublicV2)
}
func (r *ExtractTemplateGenerationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Prompt, URL are required.
type ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateGenerationRequestPublicV2 struct {
	// Instructions for generating the extract template.
	Prompt string `json:"prompt" api:"required"`
	// Example URL used to generate the extract template.
	URL string `json:"url" api:"required"`
	// Optional stable name for the generated extract template.
	Name param.Opt[string] `json:"name,omitzero"`
	// Metadata to attach to the generated extract template.
	Metadata ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateGenerationRequestPublicV2Metadata `json:"metadata,omitzero"`
	// Optional JSON schema describing expected input parameters.
	InputSchema map[string]any `json:"input_schema,omitzero"`
	// Optional JSON schema describing desired extracted output.
	OutputSchema map[string]any `json:"output_schema,omitzero"`
	paramObj
}

func (r ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateGenerationRequestPublicV2) MarshalJSON() (data []byte, err error) {
	type shadow ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateGenerationRequestPublicV2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateGenerationRequestPublicV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata to attach to the generated extract template.
type ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateGenerationRequestPublicV2Metadata struct {
	// Description for the generated template.
	Description param.Opt[string] `json:"description,omitzero"`
	// Human-friendly display name for the generated template.
	DisplayName param.Opt[string] `json:"display_name,omitzero"`
	// Tags to associate with the generated template.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateGenerationRequestPublicV2Metadata) MarshalJSON() (data []byte, err error) {
	type shadow ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateGenerationRequestPublicV2Metadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateGenerationRequestPublicV2Metadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FromExtractTemplate, Prompt are required.
type ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateRefinementRequestPublicV2 struct {
	// Name of the source extract template to refine.
	FromExtractTemplate string `json:"from_extract_template" api:"required"`
	// Instructions for refining the extract template.
	Prompt string `json:"prompt" api:"required"`
	paramObj
}

func (r ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateRefinementRequestPublicV2) MarshalJSON() (data []byte, err error) {
	type shadow ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateRefinementRequestPublicV2
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExtractTemplateGenerationNewParamsBodyCreateExtractTemplateRefinementRequestPublicV2) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
