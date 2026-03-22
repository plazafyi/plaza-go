// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomplazafyiplazago

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/plazafyi/plaza-go/internal/apijson"
	"github.com/plazafyi/plaza-go/internal/apiquery"
	"github.com/plazafyi/plaza-go/internal/param"
	"github.com/plazafyi/plaza-go/internal/requestconfig"
	"github.com/plazafyi/plaza-go/option"
)

// QueryService contains methods and other services that help with interacting with
// the plaza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueryService] method instead.
type QueryService struct {
	Options []option.RequestOption
}

// NewQueryService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQueryService(opts ...option.RequestOption) (r *QueryService) {
	r = &QueryService{}
	r.Options = opts
	return
}

// Execute a PlazaQL query
func (r *QueryService) Execute(ctx context.Context, params QueryExecuteParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// PlazaQL query request. The query is executed against Plaza's OSM database and
// results are returned as GeoJSON.
type PlazaqlQueryParam struct {
	// PlazaQL query string
	Data param.Field[string] `json:"data" api:"required"`
}

func (r PlazaqlQueryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type QueryExecuteParams struct {
	// PlazaQL query request. The query is executed against Plaza's OSM database and
	// results are returned as GeoJSON.
	PlazaqlQuery PlazaqlQueryParam `json:"plazaql_query" api:"required"`
	// Response format: json (default), geojson, csv, ndjson
	Format param.Field[string] `query:"format"`
}

func (r QueryExecuteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.PlazaqlQuery)
}

// URLQuery serializes [QueryExecuteParams]'s query parameters as `url.Values`.
func (r QueryExecuteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
