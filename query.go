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

// Execute a multi-step query pipeline
func (r *QueryService) Execute(ctx context.Context, body QueryExecuteParams, opts ...option.RequestOption) (res *QueryExecuteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Execute an Overpass QL query
func (r *QueryService) Overpass(ctx context.Context, params QueryOverpassParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/overpass"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Overpass QL query request. The query is executed against Plaza's OSM database
// and results are returned as GeoJSON.
type OverpassQueryParam struct {
	// Overpass QL query string
	Data param.Field[string] `json:"data" api:"required"`
}

func (r OverpassQueryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Pipeline execution result containing the output of each step.
type QueryExecuteResponse struct {
	// Results from each pipeline step in execution order
	Steps []map[string]interface{} `json:"steps" api:"required"`
	JSON  queryExecuteResponseJSON `json:"-"`
}

// queryExecuteResponseJSON contains the JSON metadata for the struct
// [QueryExecuteResponse]
type queryExecuteResponseJSON struct {
	Steps       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *QueryExecuteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r queryExecuteResponseJSON) RawJSON() string {
	return r.raw
}

type QueryExecuteParams struct {
	// Ordered list of query steps to execute
	Steps param.Field[[]QueryExecuteParamsStep] `json:"steps" api:"required"`
}

func (r QueryExecuteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A single pipeline step
type QueryExecuteParamsStep struct {
	// Step type: `overpass`, `filter`, or `transform`
	Type param.Field[QueryExecuteParamsStepsType] `json:"type" api:"required"`
	// Query string for this step (required for overpass steps)
	Query param.Field[string] `json:"query"`
}

func (r QueryExecuteParamsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Step type: `overpass`, `filter`, or `transform`
type QueryExecuteParamsStepsType string

const (
	QueryExecuteParamsStepsTypeOverpass  QueryExecuteParamsStepsType = "overpass"
	QueryExecuteParamsStepsTypeFilter    QueryExecuteParamsStepsType = "filter"
	QueryExecuteParamsStepsTypeTransform QueryExecuteParamsStepsType = "transform"
)

func (r QueryExecuteParamsStepsType) IsKnown() bool {
	switch r {
	case QueryExecuteParamsStepsTypeOverpass, QueryExecuteParamsStepsTypeFilter, QueryExecuteParamsStepsTypeTransform:
		return true
	}
	return false
}

type QueryOverpassParams struct {
	// Overpass QL query request. The query is executed against Plaza's OSM database
	// and results are returned as GeoJSON.
	OverpassQuery OverpassQueryParam `json:"overpass_query" api:"required"`
	// Response format: json (default), geojson, csv, ndjson
	Format param.Field[string] `query:"format"`
}

func (r QueryOverpassParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.OverpassQuery)
}

// URLQuery serializes [QueryOverpassParams]'s query parameters as `url.Values`.
func (r QueryOverpassParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
