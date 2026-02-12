// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pagination

import (
	"net/http"

	"github.com/Nimbleway/nimble-go/internal/apijson"
	"github.com/Nimbleway/nimble-go/internal/requestconfig"
	"github.com/Nimbleway/nimble-go/option"
	"github.com/Nimbleway/nimble-go/packages/param"
	"github.com/Nimbleway/nimble-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type CrawlPaginationPagination struct {
	NextCursor string `json:"next_cursor"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NextCursor  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CrawlPaginationPagination) RawJSON() string { return r.JSON.raw }
func (r *CrawlPaginationPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CrawlPagination[T any] struct {
	Data       []T                       `json:"data"`
	Pagination CrawlPaginationPagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	cfg *requestconfig.RequestConfig
	res *http.Response
}

// Returns the unmodified JSON received from the API
func (r CrawlPagination[T]) RawJSON() string { return r.JSON.raw }
func (r *CrawlPagination[T]) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GetNextPage returns the next page as defined by this pagination style. When
// there is no next page, this function will return a 'nil' for the page value, but
// will not return an error
func (r *CrawlPagination[T]) GetNextPage() (res *CrawlPagination[T], err error) {
	if len(r.Data) == 0 {
		return nil, nil
	}
	next := r.Pagination.NextCursor
	if len(next) == 0 {
		return nil, nil
	}
	cfg := r.cfg.Clone(r.cfg.Context)
	err = cfg.Apply(option.WithQuery("cursor", next))
	if err != nil {
		return nil, err
	}
	var raw *http.Response
	cfg.ResponseInto = &raw
	cfg.ResponseBodyInto = &res
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

func (r *CrawlPagination[T]) SetPageConfig(cfg *requestconfig.RequestConfig, res *http.Response) {
	if r == nil {
		r = &CrawlPagination[T]{}
	}
	r.cfg = cfg
	r.res = res
}

type CrawlPaginationAutoPager[T any] struct {
	page *CrawlPagination[T]
	cur  T
	idx  int
	run  int
	err  error
	paramObj
}

func NewCrawlPaginationAutoPager[T any](page *CrawlPagination[T], err error) *CrawlPaginationAutoPager[T] {
	return &CrawlPaginationAutoPager[T]{
		page: page,
		err:  err,
	}
}

func (r *CrawlPaginationAutoPager[T]) Next() bool {
	if r.page == nil || len(r.page.Data) == 0 {
		return false
	}
	if r.idx >= len(r.page.Data) {
		r.idx = 0
		r.page, r.err = r.page.GetNextPage()
		if r.err != nil || r.page == nil || len(r.page.Data) == 0 {
			return false
		}
	}
	r.cur = r.page.Data[r.idx]
	r.run += 1
	r.idx += 1
	return true
}

func (r *CrawlPaginationAutoPager[T]) Current() T {
	return r.cur
}

func (r *CrawlPaginationAutoPager[T]) Err() error {
	return r.err
}

func (r *CrawlPaginationAutoPager[T]) Index() int {
	return r.run
}
