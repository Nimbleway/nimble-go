// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomnimblewaynimblego

import (
	"github.com/Nimbleway/nimble-go/option"
)

// FastSerpService contains methods and other services that help with interacting
// with the nimble API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFastSerpService] method instead.
type FastSerpService struct {
	Options []option.RequestOption
}

// NewFastSerpService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFastSerpService(opts ...option.RequestOption) (r FastSerpService) {
	r = FastSerpService{}
	r.Options = opts
	return
}
