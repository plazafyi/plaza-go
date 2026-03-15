// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plaza

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/plaza-go/internal/apijson"
	"github.com/stainless-sdks/plaza-go/internal/apiquery"
	"github.com/stainless-sdks/plaza-go/internal/requestconfig"
	"github.com/stainless-sdks/plaza-go/option"
	"github.com/stainless-sdks/plaza-go/packages/param"
	"github.com/stainless-sdks/plaza-go/packages/respjson"
)

// V1Service contains methods and other services that help with interacting with
// the plaza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1Service] method instead.
type V1Service struct {
	options  []option.RequestOption
	Datasets V1DatasetService
	Elements V1ElementService
	Geocode  V1GeocodeService
}

// NewV1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1Service(opts ...option.RequestOption) (r V1Service) {
	r = V1Service{}
	r.options = opts
	r.Datasets = NewV1DatasetService(opts...)
	r.Elements = NewV1ElementService(opts...)
	r.Geocode = NewV1GeocodeService(opts...)
	return
}

// Calculate a distance matrix between points
func (r *V1Service) CalculateDistanceMatrix(ctx context.Context, body V1CalculateDistanceMatrixParams, opts ...option.RequestOption) (res *V1CalculateDistanceMatrixResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/matrix"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Calculate an isochrone from a point
func (r *V1Service) CalculateIsochrone(ctx context.Context, query V1CalculateIsochroneParams, opts ...option.RequestOption) (res *GeoJsonFeature, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/isochrone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Calculate a route between two points
func (r *V1Service) CalculateRoute(ctx context.Context, body V1CalculateRouteParams, opts ...option.RequestOption) (res *V1CalculateRouteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/route"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Execute an Overpass QL query
func (r *V1Service) ExecuteOverpass(ctx context.Context, body V1ExecuteOverpassParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/overpass"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Execute a query via REST parameters
func (r *V1Service) ExecuteQuery(ctx context.Context, query V1ExecuteQueryParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/query"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Execute a SPARQL query
func (r *V1Service) ExecuteSparql(ctx context.Context, body V1ExecuteSparqlParams, opts ...option.RequestOption) (res *V1ExecuteSparqlResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/sparql"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find elements near a point
func (r *V1Service) FindNearby(ctx context.Context, query V1FindNearbyParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/nearby"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Get a Mapbox Vector Tile
func (r *V1Service) GetTile(ctx context.Context, y int64, query V1GetTileParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.mapbox-vector-tile")}, opts...)
	path := fmt.Sprintf("api/v1/tiles/%v/%v/%v", query.Z, query.X, y)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Reverse geocode a coordinate to the nearest named feature
func (r *V1Service) ReverseGeocode(ctx context.Context, query V1ReverseGeocodeParams, opts ...option.RequestOption) (res *GeoJsonFeature, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/reverse-geocode"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Search OSM features by name
func (r *V1Service) SearchFeatures(ctx context.Context, query V1SearchFeaturesParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Snap a coordinate to the nearest road
func (r *V1Service) SnapToNearest(ctx context.Context, query V1SnapToNearestParams, opts ...option.RequestOption) (res *V1SnapToNearestResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/nearest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type V1CalculateDistanceMatrixResponse struct {
	// Distance matrix (meters), origins x destinations
	Distances [][]float64 `json:"distances" api:"required"`
	// Duration matrix (seconds), origins x destinations
	Durations [][]float64 `json:"durations" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distances   respjson.Field
		Durations   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CalculateDistanceMatrixResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CalculateDistanceMatrixResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CalculateRouteResponse struct {
	Geometry   GeoJsonGeometry                    `json:"geometry" api:"required"`
	Properties V1CalculateRouteResponseProperties `json:"properties" api:"required"`
	// Any of "Feature".
	Type V1CalculateRouteResponseType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Geometry    respjson.Field
		Properties  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CalculateRouteResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CalculateRouteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CalculateRouteResponseProperties struct {
	// Total distance in meters
	Distance float64 `json:"distance"`
	// Estimated duration in seconds
	Duration float64 `json:"duration"`
	// Travel mode used
	Mode string `json:"mode"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance    respjson.Field
		Duration    respjson.Field
		Mode        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CalculateRouteResponseProperties) RawJSON() string { return r.JSON.raw }
func (r *V1CalculateRouteResponseProperties) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CalculateRouteResponseType string

const (
	V1CalculateRouteResponseTypeFeature V1CalculateRouteResponseType = "Feature"
)

type V1ExecuteSparqlResponse struct {
	// Array of GeoJSON features from SPARQL query
	Results []GeoJsonFeature `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ExecuteSparqlResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ExecuteSparqlResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SnapToNearestResponse struct {
	// Distance to nearest road in meters
	Distance float64 `json:"distance" api:"required"`
	// Snapped latitude
	Lat float64 `json:"lat" api:"required"`
	// Snapped longitude
	Lng float64 `json:"lng" api:"required"`
	// Road edge ID
	EdgeID int64 `json:"edge_id" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Distance    respjson.Field
		Lat         respjson.Field
		Lng         respjson.Field
		EdgeID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SnapToNearestResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SnapToNearestResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CalculateDistanceMatrixParams struct {
	// List of destination coordinates
	Destinations []V1CalculateDistanceMatrixParamsDestination `json:"destinations,omitzero" api:"required"`
	// List of origin coordinates
	Origins []V1CalculateDistanceMatrixParamsOrigin `json:"origins,omitzero" api:"required"`
	// Travel mode
	//
	// Any of "auto", "foot", "bicycle".
	Mode V1CalculateDistanceMatrixParamsMode `json:"mode,omitzero"`
	paramObj
}

func (r V1CalculateDistanceMatrixParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CalculateDistanceMatrixParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CalculateDistanceMatrixParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Lat, Lng are required.
type V1CalculateDistanceMatrixParamsDestination struct {
	Lat float64 `json:"lat" api:"required"`
	Lng float64 `json:"lng" api:"required"`
	paramObj
}

func (r V1CalculateDistanceMatrixParamsDestination) MarshalJSON() (data []byte, err error) {
	type shadow V1CalculateDistanceMatrixParamsDestination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CalculateDistanceMatrixParamsDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Lat, Lng are required.
type V1CalculateDistanceMatrixParamsOrigin struct {
	Lat float64 `json:"lat" api:"required"`
	Lng float64 `json:"lng" api:"required"`
	paramObj
}

func (r V1CalculateDistanceMatrixParamsOrigin) MarshalJSON() (data []byte, err error) {
	type shadow V1CalculateDistanceMatrixParamsOrigin
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CalculateDistanceMatrixParamsOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Travel mode
type V1CalculateDistanceMatrixParamsMode string

const (
	V1CalculateDistanceMatrixParamsModeAuto    V1CalculateDistanceMatrixParamsMode = "auto"
	V1CalculateDistanceMatrixParamsModeFoot    V1CalculateDistanceMatrixParamsMode = "foot"
	V1CalculateDistanceMatrixParamsModeBicycle V1CalculateDistanceMatrixParamsMode = "bicycle"
)

type V1CalculateIsochroneParams struct {
	// Latitude
	Lat float64 `query:"lat" api:"required" json:"-"`
	// Longitude
	Lng float64 `query:"lng" api:"required" json:"-"`
	// Travel time in seconds (1-7200)
	Time float64 `query:"time" api:"required" json:"-"`
	// Travel mode (auto, foot, bicycle)
	Mode param.Opt[string] `query:"mode,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CalculateIsochroneParams]'s query parameters as
// `url.Values`.
func (r V1CalculateIsochroneParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1CalculateRouteParams struct {
	Destination V1CalculateRouteParamsDestination `json:"destination,omitzero" api:"required"`
	Origin      V1CalculateRouteParamsOrigin      `json:"origin,omitzero" api:"required"`
	// Any of "auto", "foot", "bicycle".
	Mode V1CalculateRouteParamsMode `json:"mode,omitzero"`
	paramObj
}

func (r V1CalculateRouteParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CalculateRouteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CalculateRouteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Lat, Lng are required.
type V1CalculateRouteParamsDestination struct {
	Lat float64 `json:"lat" api:"required"`
	Lng float64 `json:"lng" api:"required"`
	paramObj
}

func (r V1CalculateRouteParamsDestination) MarshalJSON() (data []byte, err error) {
	type shadow V1CalculateRouteParamsDestination
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CalculateRouteParamsDestination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Lat, Lng are required.
type V1CalculateRouteParamsOrigin struct {
	Lat float64 `json:"lat" api:"required"`
	Lng float64 `json:"lng" api:"required"`
	paramObj
}

func (r V1CalculateRouteParamsOrigin) MarshalJSON() (data []byte, err error) {
	type shadow V1CalculateRouteParamsOrigin
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CalculateRouteParamsOrigin) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CalculateRouteParamsMode string

const (
	V1CalculateRouteParamsModeAuto    V1CalculateRouteParamsMode = "auto"
	V1CalculateRouteParamsModeFoot    V1CalculateRouteParamsMode = "foot"
	V1CalculateRouteParamsModeBicycle V1CalculateRouteParamsMode = "bicycle"
)

type V1ExecuteOverpassParams struct {
	// Overpass QL query string
	Data string `json:"data" api:"required"`
	paramObj
}

func (r V1ExecuteOverpassParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ExecuteOverpassParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ExecuteOverpassParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ExecuteQueryParams struct {
	// Bounding box filter
	Bbox param.Opt[string] `query:"bbox,omitzero" json:"-"`
	// Element type filter
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1ExecuteQueryParams]'s query parameters as `url.Values`.
func (r V1ExecuteQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1ExecuteSparqlParams struct {
	// SPARQL query string
	Query string `json:"query" api:"required"`
	paramObj
}

func (r V1ExecuteSparqlParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ExecuteSparqlParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ExecuteSparqlParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1FindNearbyParams struct {
	// Latitude
	Lat float64 `query:"lat" api:"required" json:"-"`
	// Longitude
	Lng float64 `query:"lng" api:"required" json:"-"`
	// Maximum results (default 100, max 1000)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Search radius in meters (default 500)
	Radius param.Opt[float64] `query:"radius,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1FindNearbyParams]'s query parameters as `url.Values`.
func (r V1FindNearbyParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1GetTileParams struct {
	Z int64 `path:"z" api:"required" json:"-"`
	X int64 `path:"x" api:"required" json:"-"`
	paramObj
}

type V1ReverseGeocodeParams struct {
	// Latitude
	Lat float64 `query:"lat" api:"required" json:"-"`
	// Longitude
	Lng float64 `query:"lng" api:"required" json:"-"`
	paramObj
}

// URLQuery serializes [V1ReverseGeocodeParams]'s query parameters as `url.Values`.
func (r V1ReverseGeocodeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1SearchFeaturesParams struct {
	// Search query string
	Q string `query:"q" api:"required" json:"-"`
	// Cursor for pagination
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum results (default 25, max 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1SearchFeaturesParams]'s query parameters as `url.Values`.
func (r V1SearchFeaturesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1SnapToNearestParams struct {
	// Latitude
	Lat float64 `query:"lat" api:"required" json:"-"`
	// Longitude
	Lng float64 `query:"lng" api:"required" json:"-"`
	// Search radius in meters (default 500, max 5000)
	Radius param.Opt[int64] `query:"radius,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1SnapToNearestParams]'s query parameters as `url.Values`.
func (r V1SnapToNearestParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
