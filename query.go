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

// Execute an Overpass QL query
func (r *QueryService) Overpass(ctx context.Context, body QueryOverpassParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/geo+json")}, opts...)
	path := "api/v1/overpass"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Execute a SPARQL query
func (r *QueryService) Sparql(ctx context.Context, body QuerySparqlParams, opts ...option.RequestOption) (res *SparqlResult, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/geo+json")}, opts...)
	path := "api/v1/sparql"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type OverpassQueryParam struct {
	// Overpass QL query string
	Data param.Field[string] `json:"data" api:"required"`
}

func (r OverpassQueryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SparqlQueryParam struct {
	// SPARQL query string
	Query param.Field[string] `json:"query" api:"required"`
}

func (r SparqlQueryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// GeoJSON FeatureCollection of SPARQL query results
type SparqlResult struct {
	// GeoJSON features from SPARQL query
	Features []GeoJsonFeature `json:"features" api:"required"`
	Type     SparqlResultType `json:"type" api:"required"`
	JSON     sparqlResultJSON `json:"-"`
}

// sparqlResultJSON contains the JSON metadata for the struct [SparqlResult]
type sparqlResultJSON struct {
	Features    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SparqlResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sparqlResultJSON) RawJSON() string {
	return r.raw
}

type SparqlResultType string

const (
	SparqlResultTypeFeatureCollection SparqlResultType = "FeatureCollection"
)

func (r SparqlResultType) IsKnown() bool {
	switch r {
	case SparqlResultTypeFeatureCollection:
		return true
	}
	return false
}

type QueryOverpassParams struct {
	OverpassQuery OverpassQueryParam `json:"overpass_query" api:"required"`
}

func (r QueryOverpassParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.OverpassQuery)
}

type QuerySparqlParams struct {
	SparqlQuery SparqlQueryParam `json:"sparql_query" api:"required"`
}

func (r QuerySparqlParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SparqlQuery)
}
