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

// RoutingService contains methods and other services that help with interacting
// with the plaza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRoutingService] method instead.
type RoutingService struct {
	Options []option.RequestOption
}

// NewRoutingService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewRoutingService(opts ...option.RequestOption) (r *RoutingService) {
	r = &RoutingService{}
	r.Options = opts
	return
}

// Calculate an isochrone from a point
func (r *RoutingService) Isochrone(ctx context.Context, query RoutingIsochroneParams, opts ...option.RequestOption) (res *GeoJsonFeature, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/geo+json")}, opts...)
	path := "api/v1/isochrone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Calculate a distance matrix between points
func (r *RoutingService) Matrix(ctx context.Context, body RoutingMatrixParams, opts ...option.RequestOption) (res *MatrixResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/matrix"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Snap a coordinate to the nearest road
func (r *RoutingService) Nearest(ctx context.Context, query RoutingNearestParams, opts ...option.RequestOption) (res *NearestResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/nearest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Calculate a route between two points
func (r *RoutingService) Route(ctx context.Context, body RoutingRouteParams, opts ...option.RequestOption) (res *RouteResult, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/geo+json")}, opts...)
	path := "api/v1/route"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type MatrixRequestParam struct {
	// Destination points (GeoJSON MultiPoint geometry)
	Destinations param.Field[GeoJsonGeometryParam] `json:"destinations" api:"required"`
	// Origin points (GeoJSON MultiPoint geometry)
	Origins param.Field[GeoJsonGeometryParam] `json:"origins" api:"required"`
	// Travel mode
	Mode param.Field[MatrixRequestMode] `json:"mode"`
}

func (r MatrixRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Travel mode
type MatrixRequestMode string

const (
	MatrixRequestModeAuto    MatrixRequestMode = "auto"
	MatrixRequestModeFoot    MatrixRequestMode = "foot"
	MatrixRequestModeBicycle MatrixRequestMode = "bicycle"
)

func (r MatrixRequestMode) IsKnown() bool {
	switch r {
	case MatrixRequestModeAuto, MatrixRequestModeFoot, MatrixRequestModeBicycle:
		return true
	}
	return false
}

type MatrixResult struct {
	// Distance matrix (meters), origins x destinations
	Distances [][]float64 `json:"distances" api:"required"`
	// Duration matrix (seconds), origins x destinations
	Durations [][]float64      `json:"durations" api:"required"`
	JSON      matrixResultJSON `json:"-"`
}

// matrixResultJSON contains the JSON metadata for the struct [MatrixResult]
type matrixResultJSON struct {
	Distances   apijson.Field
	Durations   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MatrixResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r matrixResultJSON) RawJSON() string {
	return r.raw
}

// GeoJSON Point Feature snapped to the nearest road segment
type NearestResult struct {
	Geometry   GeoJsonGeometry         `json:"geometry" api:"required"`
	Properties NearestResultProperties `json:"properties" api:"required"`
	Type       NearestResultType       `json:"type" api:"required"`
	JSON       nearestResultJSON       `json:"-"`
}

// nearestResultJSON contains the JSON metadata for the struct [NearestResult]
type nearestResultJSON struct {
	Geometry    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NearestResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nearestResultJSON) RawJSON() string {
	return r.raw
}

type NearestResultProperties struct {
	// Distance to nearest road in meters
	DistanceM float64 `json:"distance_m"`
	// Road edge ID
	EdgeID int64                       `json:"edge_id" api:"nullable"`
	JSON   nearestResultPropertiesJSON `json:"-"`
}

// nearestResultPropertiesJSON contains the JSON metadata for the struct
// [NearestResultProperties]
type nearestResultPropertiesJSON struct {
	DistanceM   apijson.Field
	EdgeID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *NearestResultProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r nearestResultPropertiesJSON) RawJSON() string {
	return r.raw
}

type NearestResultType string

const (
	NearestResultTypeFeature NearestResultType = "Feature"
)

func (r NearestResultType) IsKnown() bool {
	switch r {
	case NearestResultTypeFeature:
		return true
	}
	return false
}

type RouteRequestParam struct {
	// Destination point (GeoJSON Point geometry)
	Destination param.Field[GeoJsonGeometryParam] `json:"destination" api:"required"`
	// Origin point (GeoJSON Point geometry)
	Origin param.Field[GeoJsonGeometryParam] `json:"origin" api:"required"`
	Mode   param.Field[RouteRequestMode]     `json:"mode"`
}

func (r RouteRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RouteRequestMode string

const (
	RouteRequestModeAuto    RouteRequestMode = "auto"
	RouteRequestModeFoot    RouteRequestMode = "foot"
	RouteRequestModeBicycle RouteRequestMode = "bicycle"
)

func (r RouteRequestMode) IsKnown() bool {
	switch r {
	case RouteRequestModeAuto, RouteRequestModeFoot, RouteRequestModeBicycle:
		return true
	}
	return false
}

type RouteResult struct {
	Geometry   GeoJsonGeometry       `json:"geometry" api:"required"`
	Properties RouteResultProperties `json:"properties" api:"required"`
	Type       RouteResultType       `json:"type" api:"required"`
	JSON       routeResultJSON       `json:"-"`
}

// routeResultJSON contains the JSON metadata for the struct [RouteResult]
type routeResultJSON struct {
	Geometry    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeResultJSON) RawJSON() string {
	return r.raw
}

type RouteResultProperties struct {
	// Total distance in meters
	Distance float64 `json:"distance"`
	// Estimated duration in seconds
	Duration float64 `json:"duration"`
	// Travel mode used
	Mode string                    `json:"mode"`
	JSON routeResultPropertiesJSON `json:"-"`
}

// routeResultPropertiesJSON contains the JSON metadata for the struct
// [RouteResultProperties]
type routeResultPropertiesJSON struct {
	Distance    apijson.Field
	Duration    apijson.Field
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RouteResultProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routeResultPropertiesJSON) RawJSON() string {
	return r.raw
}

type RouteResultType string

const (
	RouteResultTypeFeature RouteResultType = "Feature"
)

func (r RouteResultType) IsKnown() bool {
	switch r {
	case RouteResultTypeFeature:
		return true
	}
	return false
}

type RoutingIsochroneParams struct {
	// Latitude
	Lat param.Field[float64] `query:"lat" api:"required"`
	// Longitude
	Lng param.Field[float64] `query:"lng" api:"required"`
	// Travel time in seconds (1-7200)
	Time param.Field[float64] `query:"time" api:"required"`
	// Travel mode (auto, foot, bicycle)
	Mode param.Field[string] `query:"mode"`
}

// URLQuery serializes [RoutingIsochroneParams]'s query parameters as `url.Values`.
func (r RoutingIsochroneParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoutingMatrixParams struct {
	MatrixRequest MatrixRequestParam `json:"matrix_request" api:"required"`
}

func (r RoutingMatrixParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MatrixRequest)
}

type RoutingNearestParams struct {
	// Latitude
	Lat param.Field[float64] `query:"lat" api:"required"`
	// Longitude
	Lng param.Field[float64] `query:"lng" api:"required"`
	// Search radius in meters (default 500, max 5000)
	Radius param.Field[int64] `query:"radius"`
}

// URLQuery serializes [RoutingNearestParams]'s query parameters as `url.Values`.
func (r RoutingNearestParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoutingRouteParams struct {
	RouteRequest RouteRequestParam `json:"route_request" api:"required"`
}

func (r RoutingRouteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RouteRequest)
}
