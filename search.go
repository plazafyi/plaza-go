// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomplazafyiplazago

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/plazafyi/plaza-go/internal/apiquery"
	"github.com/plazafyi/plaza-go/internal/param"
	"github.com/plazafyi/plaza-go/internal/requestconfig"
	"github.com/plazafyi/plaza-go/option"
)

// SearchService contains methods and other services that help with interacting
// with the plaza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSearchService] method instead.
type SearchService struct {
	Options []option.RequestOption
}

// NewSearchService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSearchService(opts ...option.RequestOption) (r *SearchService) {
	r = &SearchService{}
	r.Options = opts
	return
}

// Search OSM features by name
func (r *SearchService) Query(ctx context.Context, query SearchQueryParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search OSM features by name
func (r *SearchService) QueryPost(ctx context.Context, body SearchQueryPostParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type SearchQueryParams struct {
	// Search query string
	Q param.Field[string] `query:"q" api:"required"`
	// Cursor for pagination
	Cursor param.Field[string] `query:"cursor"`
	// Maximum results (default 25, max 100)
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated property fields to include
	OutputFields param.Field[string] `query:"output[fields]"`
	// Extra computed fields: bbox, distance, center
	OutputInclude param.Field[string] `query:"output[include]"`
	// Coordinate decimal precision (1-15, default 7)
	OutputPrecision param.Field[int64] `query:"output[precision]"`
	// Sort by: distance, name, osm_id
	OutputSort param.Field[string] `query:"output[sort]"`
}

// URLQuery serializes [SearchQueryParams]'s query parameters as `url.Values`.
func (r SearchQueryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SearchQueryPostParams struct {
	// Search query string
	Q param.Field[string] `query:"q" api:"required"`
	// Cursor for pagination
	Cursor param.Field[string] `query:"cursor"`
	// Maximum results (default 25, max 100)
	Limit param.Field[int64] `query:"limit"`
	// Comma-separated property fields to include
	OutputFields param.Field[string] `query:"output[fields]"`
	// Extra computed fields: bbox, distance, center
	OutputInclude param.Field[string] `query:"output[include]"`
	// Coordinate decimal precision (1-15, default 7)
	OutputPrecision param.Field[int64] `query:"output[precision]"`
	// Sort by: distance, name, osm_id
	OutputSort param.Field[string] `query:"output[sort]"`
}

// URLQuery serializes [SearchQueryPostParams]'s query parameters as `url.Values`.
func (r SearchQueryPostParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
