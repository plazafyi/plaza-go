// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomplazafyiplazago

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

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
func (r *RoutingService) Isochrone(ctx context.Context, params RoutingIsochroneParams, opts ...option.RequestOption) (res *RoutingIsochroneResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/isochrone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
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
func (r *RoutingService) Nearest(ctx context.Context, body RoutingNearestParams, opts ...option.RequestOption) (res *NearestResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/nearest"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Calculate a route between two points
func (r *RoutingService) Route(ctx context.Context, params RoutingRouteParams, opts ...option.RequestOption) (res *RouteResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/route"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Request body for isochrone calculation. Computes areas reachable from a point
// within the given travel time(s).
type IsochroneRequestParam struct {
	// GeoJSON Point geometry per RFC 7946. Coordinates use [longitude, latitude]
	// order. Optional third element is altitude in meters.
	Geometry param.Field[PointGeometryParam] `json:"geometry" api:"required"`
	// Travel time budgets in seconds. Each value produces one contour polygon.
	Time param.Field[[]int64] `json:"time" api:"required"`
	// Travel mode (default: `auto`)
	Mode param.Field[IsochroneRequestMode] `json:"mode"`
}

func (r IsochroneRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Travel mode (default: `auto`)
type IsochroneRequestMode string

const (
	IsochroneRequestModeAuto    IsochroneRequestMode = "auto"
	IsochroneRequestModeFoot    IsochroneRequestMode = "foot"
	IsochroneRequestModeBicycle IsochroneRequestMode = "bicycle"
)

func (r IsochroneRequestMode) IsKnown() bool {
	switch r {
	case IsochroneRequestModeAuto, IsochroneRequestModeFoot, IsochroneRequestModeBicycle:
		return true
	}
	return false
}

// Request body for distance matrix calculation. Computes travel durations (and
// optionally distances) between every origin-destination pair. Maximum 2,500 pairs
// (origins × destinations), each list capped at 50 coordinates.
type MatrixRequestParam struct {
	// Array of destination coordinates as GeoJSON Points (max 50)
	Destinations param.Field[[]PointGeometryParam] `json:"destinations" api:"required"`
	// Array of origin coordinates as GeoJSON Points (max 50)
	Origins param.Field[[]PointGeometryParam] `json:"origins" api:"required"`
	// Comma-separated list of annotations to include: `duration` (always included),
	// `distance`. Example: `duration,distance`.
	Annotations param.Field[string] `json:"annotations"`
	// Fallback speed in km/h for pairs where no route exists. When set, unreachable
	// pairs get estimated values instead of null.
	FallbackSpeed param.Field[float64] `json:"fallback_speed"`
	// Travel mode (default: `auto`)
	Mode param.Field[MatrixRequestMode] `json:"mode"`
}

func (r MatrixRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Travel mode (default: `auto`)
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

type MatrixResult map[string]interface{}

// Request body for nearest-road-segment lookup. Snaps a point to the road network.
type NearestRequestParam struct {
	// GeoJSON Point geometry per RFC 7946. Coordinates use [longitude, latitude]
	// order. Optional third element is altitude in meters.
	Geometry param.Field[PointGeometryParam] `json:"geometry" api:"required"`
	// Maximum search radius in meters (default: 100)
	Radius param.Field[float64] `json:"radius"`
}

func (r NearestRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// GeoJSON Point Feature representing the nearest point on the road network to the
// input coordinate. Used for snapping GPS coordinates to roads.
type NearestResult struct {
	// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
	// determines the coordinate structure.
	Geometry Geometry `json:"geometry" api:"required"`
	// Snap result metadata
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

// Snap result metadata
type NearestResultProperties struct {
	// Distance from the input coordinate to the snapped point in meters
	DistanceM float64 `json:"distance_m"`
	// ID of the road network edge that was snapped to
	EdgeID int64 `json:"edge_id"`
	// Length of the matched road edge in meters
	EdgeLengthM float64 `json:"edge_length_m"`
	// OSM highway tag value (e.g. `residential`, `primary`, `motorway`)
	Highway string `json:"highway" api:"nullable"`
	// OSM way ID of the matched road segment
	OsmWayID int64 `json:"osm_way_id"`
	// OSM surface tag value (e.g. `asphalt`, `gravel`, `paved`)
	Surface string                      `json:"surface" api:"nullable"`
	JSON    nearestResultPropertiesJSON `json:"-"`
}

// nearestResultPropertiesJSON contains the JSON metadata for the struct
// [NearestResultProperties]
type nearestResultPropertiesJSON struct {
	DistanceM   apijson.Field
	EdgeID      apijson.Field
	EdgeLengthM apijson.Field
	Highway     apijson.Field
	OsmWayID    apijson.Field
	Surface     apijson.Field
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

// Request body for route calculation. Origin and destination are GeoJSON Point
// geometries. Supports optional waypoints, alternative routes, turn-by-turn steps,
// and EV routing parameters.
type RouteRequestParam struct {
	// GeoJSON Point geometry per RFC 7946. Coordinates use [longitude, latitude]
	// order. Optional third element is altitude in meters.
	Destination param.Field[PointGeometryParam] `json:"destination" api:"required"`
	// GeoJSON Point geometry per RFC 7946. Coordinates use [longitude, latitude]
	// order. Optional third element is altitude in meters.
	Origin param.Field[PointGeometryParam] `json:"origin" api:"required"`
	// Number of alternative routes to return (0-3, default 0). When > 0, response is a
	// FeatureCollection of route Features.
	Alternatives param.Field[int64] `json:"alternatives"`
	// Include per-edge annotations (speed, duration) on the route (default: false)
	Annotations param.Field[bool] `json:"annotations"`
	// Departure time for traffic-aware routing (ISO 8601)
	DepartAt param.Field[time.Time] `json:"depart_at" format:"date-time"`
	// Electric vehicle parameters for EV-aware routing
	Ev param.Field[RouteRequestEvParam] `json:"ev"`
	// Comma-separated road types to exclude (e.g. `toll,motorway,ferry`)
	Exclude param.Field[string] `json:"exclude"`
	// Geometry encoding format. Default: `geojson`.
	Geometries param.Field[RouteRequestGeometries] `json:"geometries"`
	// Travel mode (default: `auto`)
	Mode param.Field[RouteRequestMode] `json:"mode"`
	// Level of geometry detail: `full` (all points), `simplified` (Douglas-Peucker),
	// `false` (no geometry). Default: `full`.
	Overview param.Field[RouteRequestOverview] `json:"overview"`
	// Include turn-by-turn navigation steps (default: false)
	Steps param.Field[bool] `json:"steps"`
	// Traffic prediction model (only used when `depart_at` is set)
	TrafficModel param.Field[RouteRequestTrafficModel] `json:"traffic_model"`
	// Intermediate waypoints to visit in order (maximum 25)
	Waypoints param.Field[[]PointGeometryParam] `json:"waypoints"`
}

func (r RouteRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Electric vehicle parameters for EV-aware routing
type RouteRequestEvParam struct {
	// Total battery capacity in watt-hours (required for EV routing)
	BatteryCapacityWh param.Field[float64] `json:"battery_capacity_wh" api:"required"`
	// Acceptable connector types (e.g. `["ccs", "chademo"]`)
	ConnectorTypes param.Field[[]string] `json:"connector_types"`
	// Starting charge as a fraction 0-1 (default: 0.8)
	InitialChargePct param.Field[float64] `json:"initial_charge_pct"`
	// Minimum acceptable charge at destination as a fraction 0-1 (default: 0.10)
	MinChargePct param.Field[float64] `json:"min_charge_pct"`
	// Minimum charger power in kilowatts
	MinPowerKw param.Field[float64] `json:"min_power_kw"`
}

func (r RouteRequestEvParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Geometry encoding format. Default: `geojson`.
type RouteRequestGeometries string

const (
	RouteRequestGeometriesGeojson   RouteRequestGeometries = "geojson"
	RouteRequestGeometriesPolyline  RouteRequestGeometries = "polyline"
	RouteRequestGeometriesPolyline6 RouteRequestGeometries = "polyline6"
)

func (r RouteRequestGeometries) IsKnown() bool {
	switch r {
	case RouteRequestGeometriesGeojson, RouteRequestGeometriesPolyline, RouteRequestGeometriesPolyline6:
		return true
	}
	return false
}

// Travel mode (default: `auto`)
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

// Level of geometry detail: `full` (all points), `simplified` (Douglas-Peucker),
// `false` (no geometry). Default: `full`.
type RouteRequestOverview string

const (
	RouteRequestOverviewFull       RouteRequestOverview = "full"
	RouteRequestOverviewSimplified RouteRequestOverview = "simplified"
	RouteRequestOverviewFalse      RouteRequestOverview = "false"
)

func (r RouteRequestOverview) IsKnown() bool {
	switch r {
	case RouteRequestOverviewFull, RouteRequestOverviewSimplified, RouteRequestOverviewFalse:
		return true
	}
	return false
}

// Traffic prediction model (only used when `depart_at` is set)
type RouteRequestTrafficModel string

const (
	RouteRequestTrafficModelBestGuess   RouteRequestTrafficModel = "best_guess"
	RouteRequestTrafficModelOptimistic  RouteRequestTrafficModel = "optimistic"
	RouteRequestTrafficModelPessimistic RouteRequestTrafficModel = "pessimistic"
)

func (r RouteRequestTrafficModel) IsKnown() bool {
	switch r {
	case RouteRequestTrafficModelBestGuess, RouteRequestTrafficModelOptimistic, RouteRequestTrafficModelPessimistic:
		return true
	}
	return false
}

// GeoJSON Feature representing a calculated route. The geometry is a LineString or
// MultiLineString of the route path. When `alternatives > 0`, the response is a
// FeatureCollection containing multiple route Features.
type RouteResult struct {
	// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
	// determines the coordinate structure.
	Geometry Geometry `json:"geometry" api:"required"`
	// Route metadata
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

// Route metadata
type RouteResultProperties struct {
	// Total route distance in meters
	DistanceM float64 `json:"distance_m" api:"required"`
	// Estimated travel duration in seconds
	DurationS float64 `json:"duration_s" api:"required"`
	// Per-edge annotations (present when `annotations: true` in request)
	Annotations map[string]interface{} `json:"annotations" api:"nullable"`
	// Battery charge level at route waypoints as [distance_fraction, charge_pct] pairs
	// (EV routes only)
	ChargeProfile [][]float64 `json:"charge_profile" api:"nullable"`
	// Recommended charging stops along the route (EV routes only)
	ChargingStops []map[string]interface{} `json:"charging_stops" api:"nullable"`
	// Edge-level route details (present when `annotations: true`)
	Edges []map[string]interface{} `json:"edges" api:"nullable"`
	// Total energy consumed in watt-hours (EV routes only)
	EnergyUsedWh float64                   `json:"energy_used_wh" api:"nullable"`
	JSON         routeResultPropertiesJSON `json:"-"`
}

// routeResultPropertiesJSON contains the JSON metadata for the struct
// [RouteResultProperties]
type routeResultPropertiesJSON struct {
	DistanceM     apijson.Field
	DurationS     apijson.Field
	Annotations   apijson.Field
	ChargeProfile apijson.Field
	ChargingStops apijson.Field
	Edges         apijson.Field
	EnergyUsedWh  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
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

// GeoJSON FeatureCollection of isochrone polygons — areas reachable within the
// specified travel time(s). Each Feature is a Polygon contour with travel time and
// area metadata in properties.
type RoutingIsochroneResponse struct {
	// Array of isochrone polygon Features, one per contour
	Features []GeoJsonFeature `json:"features" api:"required"`
	// Always `FeatureCollection`
	Type RoutingIsochroneResponseType `json:"type" api:"required"`
	JSON routingIsochroneResponseJSON `json:"-"`
}

// routingIsochroneResponseJSON contains the JSON metadata for the struct
// [RoutingIsochroneResponse]
type routingIsochroneResponseJSON struct {
	Features    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingIsochroneResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingIsochroneResponseJSON) RawJSON() string {
	return r.raw
}

// Always `FeatureCollection`
type RoutingIsochroneResponseType string

const (
	RoutingIsochroneResponseTypeFeatureCollection RoutingIsochroneResponseType = "FeatureCollection"
)

func (r RoutingIsochroneResponseType) IsKnown() bool {
	switch r {
	case RoutingIsochroneResponseTypeFeatureCollection:
		return true
	}
	return false
}

type RoutingIsochroneParams struct {
	// Request body for isochrone calculation. Computes areas reachable from a point
	// within the given travel time(s).
	IsochroneRequest IsochroneRequestParam `json:"isochrone_request" api:"required"`
	// Response format: json (default), geojson, csv, ndjson
	Format param.Field[string] `query:"format"`
}

func (r RoutingIsochroneParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.IsochroneRequest)
}

// URLQuery serializes [RoutingIsochroneParams]'s query parameters as `url.Values`.
func (r RoutingIsochroneParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoutingMatrixParams struct {
	// Request body for distance matrix calculation. Computes travel durations (and
	// optionally distances) between every origin-destination pair. Maximum 2,500 pairs
	// (origins × destinations), each list capped at 50 coordinates.
	MatrixRequest MatrixRequestParam `json:"matrix_request" api:"required"`
}

func (r RoutingMatrixParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MatrixRequest)
}

type RoutingNearestParams struct {
	// Request body for nearest-road-segment lookup. Snaps a point to the road network.
	NearestRequest NearestRequestParam `json:"nearest_request" api:"required"`
}

func (r RoutingNearestParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.NearestRequest)
}

type RoutingRouteParams struct {
	// Request body for route calculation. Origin and destination are GeoJSON Point
	// geometries. Supports optional waypoints, alternative routes, turn-by-turn steps,
	// and EV routing parameters.
	RouteRequest RouteRequestParam `json:"route_request" api:"required"`
	// Response format for alternatives: json (default), geojson, csv, ndjson
	Format param.Field[string] `query:"format"`
}

func (r RoutingRouteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.RouteRequest)
}

// URLQuery serializes [RoutingRouteParams]'s query parameters as `url.Values`.
func (r RoutingRouteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
