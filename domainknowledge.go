// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomnimblewaynimblego

import (
	"context"
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

// DomainKnowledgeService contains methods and other services that help with
// interacting with the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDomainKnowledgeService] method instead.
type DomainKnowledgeService struct {
	Options []option.RequestOption
}

// NewDomainKnowledgeService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDomainKnowledgeService(opts ...option.RequestOption) (r DomainKnowledgeService) {
	r = DomainKnowledgeService{}
	r.Options = opts
	return
}

// Resolves the suggested driver for a given URL or agent name. Exactly one of
// `url` or `agent` must be provided.
func (r *DomainKnowledgeService) GetDriver(ctx context.Context, query DomainKnowledgeGetDriverParams, opts ...option.RequestOption) (res *DomainKnowledgeGetDriverResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v2/domain-knowledge/driver"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type DomainKnowledgeGetDriverResponse struct {
	// List of detected antibots for the domain
	Antibots []string `json:"antibots" api:"required"`
	// Description of the driver
	Description string `json:"description" api:"required"`
	// Resolved driver name
	Driver string `json:"driver" api:"required"`
	// The input agent name (present when agent query param was used)
	Agent string `json:"agent"`
	// Whether the page needs to be rendered to be properly resolved.
	NeedToRender bool `json:"need_to_render"`
	// The input URL (present when url query param was used)
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Antibots     respjson.Field
		Description  respjson.Field
		Driver       respjson.Field
		Agent        respjson.Field
		NeedToRender respjson.Field
		URL          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DomainKnowledgeGetDriverResponse) RawJSON() string { return r.JSON.raw }
func (r *DomainKnowledgeGetDriverResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DomainKnowledgeGetDriverParams struct {
	// Agent name to resolve driver for (e.g. nimble-ecommerce).
	Agent param.Opt[string] `query:"agent,omitzero" json:"-"`
	// Target domain to resolve driver for (e.g. amazon.com).
	URL param.Opt[string] `query:"url,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DomainKnowledgeGetDriverParams]'s query parameters as
// `url.Values`.
func (r DomainKnowledgeGetDriverParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
