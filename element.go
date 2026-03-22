// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomplazafyiplazago

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/plazafyi/plaza-go/internal/apijson"
	"github.com/plazafyi/plaza-go/internal/apiquery"
	"github.com/plazafyi/plaza-go/internal/param"
	"github.com/plazafyi/plaza-go/internal/requestconfig"
	"github.com/plazafyi/plaza-go/option"
)

// ElementService contains methods and other services that help with interacting
// with the plaza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewElementService] method instead.
type ElementService struct {
	Options []option.RequestOption
}

// NewElementService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewElementService(opts ...option.RequestOption) (r *ElementService) {
	r = &ElementService{}
	r.Options = opts
	return
}

// Get feature by type and ID
func (r *ElementService) Get(ctx context.Context, type_ string, id int64, opts ...option.RequestOption) (res *GeoJsonFeature, err error) {
	opts = slices.Concat(r.Options, opts)
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/features/%s/%v", type_, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Fetch multiple features by type and ID
func (r *ElementService) Batch(ctx context.Context, body ElementBatchParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/features/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get feature by type and ID
func (r *ElementService) Lookup(ctx context.Context, opts ...option.RequestOption) (res *GeoJsonFeature, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/features/lookup"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Find features near a geographic point
func (r *ElementService) Nearby(ctx context.Context, query ElementNearbyParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/features/nearby"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Find features near a geographic point
func (r *ElementService) NearbyPost(ctx context.Context, body ElementNearbyPostParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/features/nearby"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Query features by spatial predicate, bounding box, or H3 cell
func (r *ElementService) Query(ctx context.Context, query ElementQueryParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/features"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Query features by spatial predicate, bounding box, or H3 cell
func (r *ElementService) QueryPost(ctx context.Context, body ElementQueryPostParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/features"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Fetch multiple OSM elements by their type and ID in a single request. Maximum
// 100 elements per batch.
type BatchRequestParam struct {
	// Array of element references to fetch
	Elements param.Field[[]BatchRequestElementParam] `json:"elements" api:"required"`
}

func (r BatchRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Reference to a single OSM element
type BatchRequestElementParam struct {
	// OSM element ID
	ID param.Field[int64] `json:"id" api:"required"`
	// OSM element type
	Type param.Field[BatchRequestElementsType] `json:"type" api:"required"`
}

func (r BatchRequestElementParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// OSM element type
type BatchRequestElementsType string

const (
	BatchRequestElementsTypeNode     BatchRequestElementsType = "node"
	BatchRequestElementsTypeWay      BatchRequestElementsType = "way"
	BatchRequestElementsTypeRelation BatchRequestElementsType = "relation"
)

func (r BatchRequestElementsType) IsKnown() bool {
	switch r {
	case BatchRequestElementsTypeNode, BatchRequestElementsTypeWay, BatchRequestElementsTypeRelation:
		return true
	}
	return false
}

type ElementBatchParams struct {
	// Fetch multiple OSM elements by their type and ID in a single request. Maximum
	// 100 elements per batch.
	BatchRequest BatchRequestParam `json:"batch_request" api:"required"`
}

func (r ElementBatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BatchRequest)
}

type ElementNearbyParams struct {
	// Legacy shorthand. Latitude (-90 to 90). Use near param instead.
	Lat param.Field[float64] `query:"lat"`
	// Maximum results (default 20, max 100)
	Limit param.Field[int64] `query:"limit"`
	// Legacy shorthand. Longitude (-180 to 180). Use near param instead.
	Lng param.Field[float64] `query:"lng"`
	// Point geometry for proximity search (lat,lng or GeoJSON). Alternative to lat/lng
	// params.
	Near param.Field[string] `query:"near"`
	// Buffer geometry by meters
	OutputBuffer param.Field[float64] `query:"output[buffer]"`
	// Replace geometry with centroid
	OutputCentroid param.Field[bool] `query:"output[centroid]"`
	// Comma-separated property fields to include
	OutputFields param.Field[string] `query:"output[fields]"`
	// Include geometry (default true)
	OutputGeometry param.Field[bool] `query:"output[geometry]"`
	// Extra computed fields: bbox, distance, center
	OutputInclude param.Field[string] `query:"output[include]"`
	// Coordinate decimal precision (1-15, default 7)
	OutputPrecision param.Field[int64] `query:"output[precision]"`
	// Simplify geometry tolerance in meters
	OutputSimplify param.Field[float64] `query:"output[simplify]"`
	// Sort by: distance, name, osm_id
	OutputSort param.Field[string] `query:"output[sort]"`
	// Search radius in meters (default 500, max 10000)
	Radius param.Field[int64] `query:"radius"`
}

// URLQuery serializes [ElementNearbyParams]'s query parameters as `url.Values`.
func (r ElementNearbyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ElementNearbyPostParams struct {
	// Legacy shorthand. Latitude (-90 to 90). Use near param instead.
	Lat param.Field[float64] `query:"lat"`
	// Maximum results (default 20, max 100)
	Limit param.Field[int64] `query:"limit"`
	// Legacy shorthand. Longitude (-180 to 180). Use near param instead.
	Lng param.Field[float64] `query:"lng"`
	// Point geometry for proximity search (lat,lng or GeoJSON). Alternative to lat/lng
	// params.
	Near param.Field[string] `query:"near"`
	// Buffer geometry by meters
	OutputBuffer param.Field[float64] `query:"output[buffer]"`
	// Replace geometry with centroid
	OutputCentroid param.Field[bool] `query:"output[centroid]"`
	// Comma-separated property fields to include
	OutputFields param.Field[string] `query:"output[fields]"`
	// Include geometry (default true)
	OutputGeometry param.Field[bool] `query:"output[geometry]"`
	// Extra computed fields: bbox, distance, center
	OutputInclude param.Field[string] `query:"output[include]"`
	// Coordinate decimal precision (1-15, default 7)
	OutputPrecision param.Field[int64] `query:"output[precision]"`
	// Simplify geometry tolerance in meters
	OutputSimplify param.Field[float64] `query:"output[simplify]"`
	// Sort by: distance, name, osm_id
	OutputSort param.Field[string] `query:"output[sort]"`
	// Search radius in meters (default 500, max 10000)
	Radius param.Field[int64] `query:"radius"`
}

// URLQuery serializes [ElementNearbyPostParams]'s query parameters as
// `url.Values`.
func (r ElementNearbyPostParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ElementQueryParams struct {
	// Legacy shorthand. Bounding box: south,west,north,east. Use spatial predicates
	// (near, within, intersects) instead.
	Bbox param.Field[string] `query:"bbox"`
	// Geometry that features must contain
	Contains param.Field[string] `query:"contains"`
	// Geometry that features must cross
	Crosses param.Field[string] `query:"crosses"`
	// Cursor for pagination
	Cursor param.Field[string] `query:"cursor"`
	// Response format. json (default) returns paginated GeoJSON. geojson/csv/ndjson
	// stream via chunked transfer encoding.
	Format param.Field[string] `query:"format"`
	// Legacy shorthand. H3 cell index. Use spatial predicates instead.
	H3 param.Field[string] `query:"h3"`
	// Geometry that features must intersect
	Intersects param.Field[string] `query:"intersects"`
	// Maximum results (default 100, max 10000)
	Limit param.Field[int64] `query:"limit"`
	// Point geometry for proximity search (lat,lng). Requires radius.
	Near param.Field[string] `query:"near"`
	// Buffer geometry by meters
	OutputBuffer param.Field[float64] `query:"output[buffer]"`
	// Replace geometry with centroid
	OutputCentroid param.Field[bool] `query:"output[centroid]"`
	// Comma-separated property fields to include
	OutputFields param.Field[string] `query:"output[fields]"`
	// Include geometry (default true)
	OutputGeometry param.Field[bool] `query:"output[geometry]"`
	// Extra computed fields: bbox, distance, center
	OutputInclude param.Field[string] `query:"output[include]"`
	// Coordinate decimal precision (1-15, default 7)
	OutputPrecision param.Field[int64] `query:"output[precision]"`
	// Simplify geometry tolerance in meters
	OutputSimplify param.Field[float64] `query:"output[simplify]"`
	// Sort by: distance, name, osm_id
	OutputSort param.Field[string] `query:"output[sort]"`
	// Search radius in meters (for near) or buffer distance (for other predicates)
	Radius param.Field[float64] `query:"radius"`
	// Geometry that features must touch
	Touches param.Field[string] `query:"touches"`
	// Element types (comma-separated: node,way,relation)
	Type param.Field[string] `query:"type"`
	// Geometry that features must be within
	Within param.Field[string] `query:"within"`
}

// URLQuery serializes [ElementQueryParams]'s query parameters as `url.Values`.
func (r ElementQueryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ElementQueryPostParams struct {
	// Legacy shorthand. Bounding box: south,west,north,east. Use spatial predicates
	// (near, within, intersects) instead.
	Bbox param.Field[string] `query:"bbox"`
	// Geometry that features must contain
	Contains param.Field[string] `query:"contains"`
	// Geometry that features must cross
	Crosses param.Field[string] `query:"crosses"`
	// Cursor for pagination
	Cursor param.Field[string] `query:"cursor"`
	// Response format. json (default) returns paginated GeoJSON. geojson/csv/ndjson
	// stream via chunked transfer encoding.
	Format param.Field[string] `query:"format"`
	// Legacy shorthand. H3 cell index. Use spatial predicates instead.
	H3 param.Field[string] `query:"h3"`
	// Geometry that features must intersect
	Intersects param.Field[string] `query:"intersects"`
	// Maximum results (default 100, max 10000)
	Limit param.Field[int64] `query:"limit"`
	// Point geometry for proximity search (lat,lng). Requires radius.
	Near param.Field[string] `query:"near"`
	// Buffer geometry by meters
	OutputBuffer param.Field[float64] `query:"output[buffer]"`
	// Replace geometry with centroid
	OutputCentroid param.Field[bool] `query:"output[centroid]"`
	// Comma-separated property fields to include
	OutputFields param.Field[string] `query:"output[fields]"`
	// Include geometry (default true)
	OutputGeometry param.Field[bool] `query:"output[geometry]"`
	// Extra computed fields: bbox, distance, center
	OutputInclude param.Field[string] `query:"output[include]"`
	// Coordinate decimal precision (1-15, default 7)
	OutputPrecision param.Field[int64] `query:"output[precision]"`
	// Simplify geometry tolerance in meters
	OutputSimplify param.Field[float64] `query:"output[simplify]"`
	// Sort by: distance, name, osm_id
	OutputSort param.Field[string] `query:"output[sort]"`
	// Search radius in meters (for near) or buffer distance (for other predicates)
	Radius param.Field[float64] `query:"radius"`
	// Geometry that features must touch
	Touches param.Field[string] `query:"touches"`
	// Element types (comma-separated: node,way,relation)
	Type param.Field[string] `query:"type"`
	// Geometry that features must be within
	Within param.Field[string] `query:"within"`
}

// URLQuery serializes [ElementQueryPostParams]'s query parameters as `url.Values`.
func (r ElementQueryPostParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
