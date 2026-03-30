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

// AgentGenerationService contains methods and other services that help with
// interacting with the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentGenerationService] method instead.
type AgentGenerationService struct {
	Options []option.RequestOption
}

// NewAgentGenerationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentGenerationService(opts ...option.RequestOption) (r AgentGenerationService) {
	r = AgentGenerationService{}
	r.Options = opts
	return
}

// Create Agent Generation
func (r *AgentGenerationService) New(ctx context.Context, body AgentGenerationNewParams, opts ...option.RequestOption) (res *AgentGenerationNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/agents/generations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get Agent Generation
func (r *AgentGenerationService) Get(ctx context.Context, generationID string, opts ...option.RequestOption) (res *AgentGenerationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if generationID == "" {
		err = errors.New("missing required generation_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/agents/generations/%s", generationID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AgentGenerationNewResponse struct {
	ID                 string    `json:"id" api:"required" format:"uuid"`
	Status             string    `json:"status" api:"required"`
	AgentName          string    `json:"agent_name" api:"nullable"`
	CompletedAt        time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	CreatedAt          time.Time `json:"created_at" api:"nullable" format:"date-time"`
	Error              string    `json:"error" api:"nullable"`
	GeneratedVersion   any       `json:"generated_version" api:"nullable"`
	GeneratedVersionID string    `json:"generated_version_id" api:"nullable" format:"uuid"`
	SourceVersionID    string    `json:"source_version_id" api:"nullable" format:"uuid"`
	StartedAt          time.Time `json:"started_at" api:"nullable" format:"date-time"`
	Summary            string    `json:"summary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Status             respjson.Field
		AgentName          respjson.Field
		CompletedAt        respjson.Field
		CreatedAt          respjson.Field
		Error              respjson.Field
		GeneratedVersion   respjson.Field
		GeneratedVersionID respjson.Field
		SourceVersionID    respjson.Field
		StartedAt          respjson.Field
		Summary            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGenerationNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGenerationNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGenerationGetResponse struct {
	ID                 string    `json:"id" api:"required" format:"uuid"`
	Status             string    `json:"status" api:"required"`
	AgentName          string    `json:"agent_name" api:"nullable"`
	CompletedAt        time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	CreatedAt          time.Time `json:"created_at" api:"nullable" format:"date-time"`
	Error              string    `json:"error" api:"nullable"`
	GeneratedVersion   any       `json:"generated_version" api:"nullable"`
	GeneratedVersionID string    `json:"generated_version_id" api:"nullable" format:"uuid"`
	SourceVersionID    string    `json:"source_version_id" api:"nullable" format:"uuid"`
	StartedAt          time.Time `json:"started_at" api:"nullable" format:"date-time"`
	Summary            string    `json:"summary" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		Status             respjson.Field
		AgentName          respjson.Field
		CompletedAt        respjson.Field
		CreatedAt          respjson.Field
		Error              respjson.Field
		GeneratedVersion   respjson.Field
		GeneratedVersionID respjson.Field
		SourceVersionID    respjson.Field
		StartedAt          respjson.Field
		Summary            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentGenerationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentGenerationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentGenerationNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfCreateAgentGenerationRequest *AgentGenerationNewParamsBodyCreateAgentGenerationRequest `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfCreateAgentRefinementRequest *AgentGenerationNewParamsBodyCreateAgentRefinementRequest `json:",inline"`

	paramObj
}

func (u AgentGenerationNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfCreateAgentGenerationRequest, u.OfCreateAgentRefinementRequest)
}
func (r *AgentGenerationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties AgentName, Prompt, URL are required.
type AgentGenerationNewParamsBodyCreateAgentGenerationRequest struct {
	AgentName    string `json:"agent_name" api:"required"`
	Prompt       string `json:"prompt" api:"required"`
	URL          string `json:"url" api:"required"`
	Metadata     any    `json:"metadata,omitzero"`
	InputSchema  any    `json:"input_schema,omitzero"`
	OutputSchema any    `json:"output_schema,omitzero"`
	paramObj
}

func (r AgentGenerationNewParamsBodyCreateAgentGenerationRequest) MarshalJSON() (data []byte, err error) {
	type shadow AgentGenerationNewParamsBodyCreateAgentGenerationRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGenerationNewParamsBodyCreateAgentGenerationRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties FromAgent, Prompt are required.
type AgentGenerationNewParamsBodyCreateAgentRefinementRequest struct {
	FromAgent string `json:"from_agent" api:"required"`
	Prompt    string `json:"prompt" api:"required"`
	paramObj
}

func (r AgentGenerationNewParamsBodyCreateAgentRefinementRequest) MarshalJSON() (data []byte, err error) {
	type shadow AgentGenerationNewParamsBodyCreateAgentRefinementRequest
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentGenerationNewParamsBodyCreateAgentRefinementRequest) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
