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

// FeatureService contains methods and other services that help with interacting
// with the plaza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFeatureService] method instead.
type FeatureService struct {
	Options []option.RequestOption
}

// NewFeatureService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFeatureService(opts ...option.RequestOption) (r *FeatureService) {
	r = &FeatureService{}
	r.Options = opts
	return
}

// Get feature by type and ID
func (r *FeatureService) Get(ctx context.Context, type_ string, id int64, opts ...option.RequestOption) (res *GeoJsonFeature, err error) {
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
func (r *FeatureService) Batch(ctx context.Context, body FeatureBatchParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/features/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Query features by spatial predicate, bounding box, or H3 cell
func (r *FeatureService) Query(ctx context.Context, params FeatureQueryParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/features"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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

// Spatial predicates for filtering features by geographic relationship. Predicates
// are mutually exclusive — use exactly one per request. The parameter name is the
// spatial operation, the value is a GeoJSON geometry to test against.
//
// | Predicate        | Meaning                                    |
// | ---------------- | ------------------------------------------ |
// | `around`         | Within radius meters (requires `radius`)   |
// | `intersects`     | Feature overlaps the input geometry        |
// | `within`         | Feature is fully inside the input geometry |
// | `contains`       | Feature fully contains the input geometry  |
// | `crosses`        | Feature crosses the input geometry         |
// | `touches`        | Feature shares boundary but not interior   |
// | `not_intersects` | Feature does not overlap                   |
// | `not_within`     | Feature is not fully inside                |
// | `not_contains`   | Feature does not fully contain             |
type SpatialPredicateParam struct {
	// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
	// determines the coordinate structure.
	Around param.Field[GeometryUnionParam] `json:"around"`
	// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
	// determines the coordinate structure.
	Contains param.Field[GeometryUnionParam] `json:"contains"`
	// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
	// determines the coordinate structure.
	Crosses param.Field[GeometryUnionParam] `json:"crosses"`
	// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
	// determines the coordinate structure.
	Intersects param.Field[GeometryUnionParam] `json:"intersects"`
	// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
	// determines the coordinate structure.
	NotContains param.Field[GeometryUnionParam] `json:"not_contains"`
	// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
	// determines the coordinate structure.
	NotIntersects param.Field[GeometryUnionParam] `json:"not_intersects"`
	// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
	// determines the coordinate structure.
	NotWithin param.Field[GeometryUnionParam] `json:"not_within"`
	// Search radius in meters. Required for `around`, optional buffer for other
	// predicates.
	Radius param.Field[float64] `json:"radius"`
	// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
	// determines the coordinate structure.
	Touches param.Field[GeometryUnionParam] `json:"touches"`
	// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
	// determines the coordinate structure.
	Within param.Field[GeometryUnionParam] `json:"within"`
}

func (r SpatialPredicateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FeatureBatchParams struct {
	// Fetch multiple OSM elements by their type and ID in a single request. Maximum
	// 100 elements per batch.
	BatchRequest BatchRequestParam `json:"batch_request" api:"required"`
}

func (r FeatureBatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BatchRequest)
}

type FeatureQueryParams struct {
	// Cursor for pagination
	Cursor param.Field[string] `query:"cursor"`
	// Response format. json (default) returns paginated GeoJSON. geojson/csv/ndjson
	// stream via chunked transfer encoding.
	Format param.Field[string] `query:"format"`
	// Legacy shorthand. H3 cell index. Use spatial predicates instead.
	H3 param.Field[string] `query:"h3"`
	// Maximum results (default 100, max 10000)
	Limit param.Field[int64] `query:"limit"`
	// Element types (comma-separated: node,way,relation)
	Type param.Field[string] `query:"type"`
	// Spatial predicates for filtering features by geographic relationship. Predicates
	// are mutually exclusive — use exactly one per request. The parameter name is the
	// spatial operation, the value is a GeoJSON geometry to test against.
	//
	// | Predicate        | Meaning                                    |
	// | ---------------- | ------------------------------------------ |
	// | `around`         | Within radius meters (requires `radius`)   |
	// | `intersects`     | Feature overlaps the input geometry        |
	// | `within`         | Feature is fully inside the input geometry |
	// | `contains`       | Feature fully contains the input geometry  |
	// | `crosses`        | Feature crosses the input geometry         |
	// | `touches`        | Feature shares boundary but not interior   |
	// | `not_intersects` | Feature does not overlap                   |
	// | `not_within`     | Feature is not fully inside                |
	// | `not_contains`   | Feature does not fully contain             |
	SpatialPredicate SpatialPredicateParam `json:"spatial_predicate"`
}

func (r FeatureQueryParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.SpatialPredicate)
}

// URLQuery serializes [FeatureQueryParams]'s query parameters as `url.Values`.
func (r FeatureQueryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
