// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomplazafyiplazago

import (
	"context"
	"net/http"
	"slices"

	"github.com/plazafyi/plaza-go/internal/apijson"
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
func (r *QueryService) Overpass(ctx context.Context, body QueryOverpassParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/overpass"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Execute a SPARQL query
func (r *QueryService) Sparql(ctx context.Context, body QuerySparqlParams, opts ...option.RequestOption) (res *SparqlResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/sparql"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

// SPARQL query request. Queries OSM data using SPARQL syntax. Results are returned
// as a JSON object with a `results` array.
type SparqlQueryParam struct {
	// SPARQL query string
	Query param.Field[string] `json:"query" api:"required"`
}

func (r SparqlQueryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// SPARQL query result. Contains a `results` array of GeoJSON Feature objects.
// Unlike REST feature endpoints, SPARQL results may omit `@type`, `@id`, and
// compound `id` fields depending on the query shape.
type SparqlResult struct {
	// Array of GeoJSON Features matching the SPARQL query. Features include `@type`
	// and `@id` metadata when the source element type is known, but may contain only
	// tags as properties for untyped results.
	Results []SparqlResultResult `json:"results" api:"required"`
	JSON    sparqlResultJSON     `json:"-"`
}

// sparqlResultJSON contains the JSON metadata for the struct [SparqlResult]
type sparqlResultJSON struct {
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SparqlResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sparqlResultJSON) RawJSON() string {
	return r.raw
}

// GeoJSON Feature (may lack @type/@id metadata for untyped results)
type SparqlResultResult struct {
	// GeoJSON Geometry object per RFC 7946. Coordinates use [longitude, latitude]
	// order. 3D coordinates [lng, lat, elevation] are used for elevation endpoints.
	Geometry GeoJsonGeometry `json:"geometry" api:"required"`
	// OSM tags as key-value pairs, optionally with `@type` and `@id` metadata
	Properties map[string]interface{} `json:"properties" api:"required"`
	// Always `Feature`
	Type SparqlResultResultsType `json:"type" api:"required"`
	// Compound identifier in `type/osm_id` format (present when element type is known)
	ID   string                 `json:"id" api:"nullable"`
	JSON sparqlResultResultJSON `json:"-"`
}

// sparqlResultResultJSON contains the JSON metadata for the struct
// [SparqlResultResult]
type sparqlResultResultJSON struct {
	Geometry    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SparqlResultResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sparqlResultResultJSON) RawJSON() string {
	return r.raw
}

// Always `Feature`
type SparqlResultResultsType string

const (
	SparqlResultResultsTypeFeature SparqlResultResultsType = "Feature"
)

func (r SparqlResultResultsType) IsKnown() bool {
	switch r {
	case SparqlResultResultsTypeFeature:
		return true
	}
	return false
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
	// Step type: `overpass`, `sparql`, `filter`, or `transform`
	Type param.Field[QueryExecuteParamsStepsType] `json:"type" api:"required"`
	// Query string for this step (required for overpass/sparql steps)
	Query param.Field[string] `json:"query"`
}

func (r QueryExecuteParamsStep) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Step type: `overpass`, `sparql`, `filter`, or `transform`
type QueryExecuteParamsStepsType string

const (
	QueryExecuteParamsStepsTypeOverpass  QueryExecuteParamsStepsType = "overpass"
	QueryExecuteParamsStepsTypeSparql    QueryExecuteParamsStepsType = "sparql"
	QueryExecuteParamsStepsTypeFilter    QueryExecuteParamsStepsType = "filter"
	QueryExecuteParamsStepsTypeTransform QueryExecuteParamsStepsType = "transform"
)

func (r QueryExecuteParamsStepsType) IsKnown() bool {
	switch r {
	case QueryExecuteParamsStepsTypeOverpass, QueryExecuteParamsStepsTypeSparql, QueryExecuteParamsStepsTypeFilter, QueryExecuteParamsStepsTypeTransform:
		return true
	}
	return false
}

type QueryOverpassParams struct {
	// Overpass QL query request. The query is executed against Plaza's OSM database
	// and results are returned as GeoJSON.
	OverpassQuery OverpassQueryParam `json:"overpass_query" api:"required"`
}

func (r QueryOverpassParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.OverpassQuery)
}

type QuerySparqlParams struct {
	// SPARQL query request. Queries OSM data using SPARQL syntax. Results are returned
	// as a JSON object with a `results` array.
	SparqlQuery SparqlQueryParam `json:"sparql_query" api:"required"`
}

func (r QuerySparqlParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SparqlQuery)
}
