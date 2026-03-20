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
func (r *RoutingService) Isochrone(ctx context.Context, query RoutingIsochroneParams, opts ...option.RequestOption) (res *RoutingIsochroneResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/isochrone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Calculate an isochrone from a point
func (r *RoutingService) IsochronePost(ctx context.Context, body RoutingIsochronePostParams, opts ...option.RequestOption) (res *RoutingIsochronePostResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/isochrone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

// Snap a coordinate to the nearest road
func (r *RoutingService) NearestPost(ctx context.Context, body RoutingNearestPostParams, opts ...option.RequestOption) (res *NearestResult, err error) {
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

// Request body for distance matrix calculation. Computes travel durations (and
// optionally distances) between every origin-destination pair. Maximum 2,500 pairs
// (origins × destinations), each list capped at 50 coordinates.
type MatrixRequestParam struct {
	// Array of destination coordinates (max 50)
	Destinations param.Field[[]MatrixRequestDestinationParam] `json:"destinations" api:"required"`
	// Array of origin coordinates (max 50)
	Origins param.Field[[]MatrixRequestOriginParam] `json:"origins" api:"required"`
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

// Geographic coordinate as a JSON object with `lat` and `lng` fields.
type MatrixRequestDestinationParam struct {
	// Latitude in decimal degrees (-90 to 90)
	Lat param.Field[float64] `json:"lat" api:"required"`
	// Longitude in decimal degrees (-180 to 180)
	Lng param.Field[float64] `json:"lng" api:"required"`
}

func (r MatrixRequestDestinationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Geographic coordinate as a JSON object with `lat` and `lng` fields.
type MatrixRequestOriginParam struct {
	// Latitude in decimal degrees (-90 to 90)
	Lat param.Field[float64] `json:"lat" api:"required"`
	// Longitude in decimal degrees (-180 to 180)
	Lng param.Field[float64] `json:"lng" api:"required"`
}

func (r MatrixRequestOriginParam) MarshalJSON() (data []byte, err error) {
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

// GeoJSON Point Feature representing the nearest point on the road network to the
// input coordinate. Used for snapping GPS coordinates to roads.
type NearestResult struct {
	// GeoJSON Geometry object per RFC 7946. Coordinates use [longitude, latitude]
	// order. 3D coordinates [lng, lat, elevation] are used for elevation endpoints.
	Geometry GeoJsonGeometry `json:"geometry" api:"required"`
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

// Request body for route calculation. Origin and destination are lat/lng
// coordinate objects. Supports optional waypoints, alternative routes,
// turn-by-turn steps, and EV routing parameters.
type RouteRequestParam struct {
	// Geographic coordinate as a JSON object with `lat` and `lng` fields.
	Destination param.Field[RouteRequestDestinationParam] `json:"destination" api:"required"`
	// Geographic coordinate as a JSON object with `lat` and `lng` fields.
	Origin param.Field[RouteRequestOriginParam] `json:"origin" api:"required"`
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
	Waypoints param.Field[[]RouteRequestWaypointParam] `json:"waypoints"`
}

func (r RouteRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Geographic coordinate as a JSON object with `lat` and `lng` fields.
type RouteRequestDestinationParam struct {
	// Latitude in decimal degrees (-90 to 90)
	Lat param.Field[float64] `json:"lat" api:"required"`
	// Longitude in decimal degrees (-180 to 180)
	Lng param.Field[float64] `json:"lng" api:"required"`
}

func (r RouteRequestDestinationParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Geographic coordinate as a JSON object with `lat` and `lng` fields.
type RouteRequestOriginParam struct {
	// Latitude in decimal degrees (-90 to 90)
	Lat param.Field[float64] `json:"lat" api:"required"`
	// Longitude in decimal degrees (-180 to 180)
	Lng param.Field[float64] `json:"lng" api:"required"`
}

func (r RouteRequestOriginParam) MarshalJSON() (data []byte, err error) {
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

// Geographic coordinate as a JSON object with `lat` and `lng` fields.
type RouteRequestWaypointParam struct {
	// Latitude in decimal degrees (-90 to 90)
	Lat param.Field[float64] `json:"lat" api:"required"`
	// Longitude in decimal degrees (-180 to 180)
	Lng param.Field[float64] `json:"lng" api:"required"`
}

func (r RouteRequestWaypointParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// GeoJSON Feature representing a calculated route. The geometry is a LineString or
// MultiLineString of the route path. When `alternatives > 0`, the response is a
// FeatureCollection containing multiple route Features.
type RouteResult struct {
	// GeoJSON Geometry object per RFC 7946. Coordinates use [longitude, latitude]
	// order. 3D coordinates [lng, lat, elevation] are used for elevation endpoints.
	Geometry GeoJsonGeometry `json:"geometry" api:"required"`
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

// GeoJSON Feature or FeatureCollection representing isochrone polygons — areas
// reachable within the specified travel time(s). Single time value returns a
// Feature; comma-separated times return a FeatureCollection with one polygon per
// contour.
type RoutingIsochroneResponse struct {
	// Array of isochrone polygon Features (multi-contour only)
	Features []GeoJsonFeature `json:"features" api:"nullable"`
	// GeoJSON Geometry object per RFC 7946. Coordinates use [longitude, latitude]
	// order. 3D coordinates [lng, lat, elevation] are used for elevation endpoints.
	Geometry GeoJsonGeometry `json:"geometry" api:"nullable"`
	// Isochrone metadata
	Properties RoutingIsochroneResponseProperties `json:"properties" api:"nullable"`
	// `Feature` for single contour, `FeatureCollection` for multiple contours
	Type RoutingIsochroneResponseType `json:"type"`
	JSON routingIsochroneResponseJSON `json:"-"`
}

// routingIsochroneResponseJSON contains the JSON metadata for the struct
// [RoutingIsochroneResponse]
type routingIsochroneResponseJSON struct {
	Features    apijson.Field
	Geometry    apijson.Field
	Properties  apijson.Field
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

// Isochrone metadata
type RoutingIsochroneResponseProperties struct {
	// Area of the isochrone polygon in square meters (multi-contour features only)
	AreaM2 float64 `json:"area_m2" api:"nullable"`
	// Maximum actual travel cost in seconds to the isochrone boundary (single contour
	// only)
	MaxCostS float64 `json:"max_cost_s" api:"nullable"`
	// Travel mode used for the isochrone calculation
	Mode RoutingIsochroneResponsePropertiesMode `json:"mode"`
	// Travel time budget in seconds
	TimeSeconds float64 `json:"time_seconds"`
	// Number of road network vertices within the isochrone
	VerticesReached int64                                  `json:"vertices_reached"`
	JSON            routingIsochroneResponsePropertiesJSON `json:"-"`
}

// routingIsochroneResponsePropertiesJSON contains the JSON metadata for the struct
// [RoutingIsochroneResponseProperties]
type routingIsochroneResponsePropertiesJSON struct {
	AreaM2          apijson.Field
	MaxCostS        apijson.Field
	Mode            apijson.Field
	TimeSeconds     apijson.Field
	VerticesReached apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RoutingIsochroneResponseProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingIsochroneResponsePropertiesJSON) RawJSON() string {
	return r.raw
}

// Travel mode used for the isochrone calculation
type RoutingIsochroneResponsePropertiesMode string

const (
	RoutingIsochroneResponsePropertiesModeAuto    RoutingIsochroneResponsePropertiesMode = "auto"
	RoutingIsochroneResponsePropertiesModeFoot    RoutingIsochroneResponsePropertiesMode = "foot"
	RoutingIsochroneResponsePropertiesModeBicycle RoutingIsochroneResponsePropertiesMode = "bicycle"
)

func (r RoutingIsochroneResponsePropertiesMode) IsKnown() bool {
	switch r {
	case RoutingIsochroneResponsePropertiesModeAuto, RoutingIsochroneResponsePropertiesModeFoot, RoutingIsochroneResponsePropertiesModeBicycle:
		return true
	}
	return false
}

// `Feature` for single contour, `FeatureCollection` for multiple contours
type RoutingIsochroneResponseType string

const (
	RoutingIsochroneResponseTypeFeature           RoutingIsochroneResponseType = "Feature"
	RoutingIsochroneResponseTypeFeatureCollection RoutingIsochroneResponseType = "FeatureCollection"
)

func (r RoutingIsochroneResponseType) IsKnown() bool {
	switch r {
	case RoutingIsochroneResponseTypeFeature, RoutingIsochroneResponseTypeFeatureCollection:
		return true
	}
	return false
}

// GeoJSON Feature or FeatureCollection representing isochrone polygons — areas
// reachable within the specified travel time(s). Single time value returns a
// Feature; comma-separated times return a FeatureCollection with one polygon per
// contour.
type RoutingIsochronePostResponse struct {
	// Array of isochrone polygon Features (multi-contour only)
	Features []GeoJsonFeature `json:"features" api:"nullable"`
	// GeoJSON Geometry object per RFC 7946. Coordinates use [longitude, latitude]
	// order. 3D coordinates [lng, lat, elevation] are used for elevation endpoints.
	Geometry GeoJsonGeometry `json:"geometry" api:"nullable"`
	// Isochrone metadata
	Properties RoutingIsochronePostResponseProperties `json:"properties" api:"nullable"`
	// `Feature` for single contour, `FeatureCollection` for multiple contours
	Type RoutingIsochronePostResponseType `json:"type"`
	JSON routingIsochronePostResponseJSON `json:"-"`
}

// routingIsochronePostResponseJSON contains the JSON metadata for the struct
// [RoutingIsochronePostResponse]
type routingIsochronePostResponseJSON struct {
	Features    apijson.Field
	Geometry    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingIsochronePostResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingIsochronePostResponseJSON) RawJSON() string {
	return r.raw
}

// Isochrone metadata
type RoutingIsochronePostResponseProperties struct {
	// Area of the isochrone polygon in square meters (multi-contour features only)
	AreaM2 float64 `json:"area_m2" api:"nullable"`
	// Maximum actual travel cost in seconds to the isochrone boundary (single contour
	// only)
	MaxCostS float64 `json:"max_cost_s" api:"nullable"`
	// Travel mode used for the isochrone calculation
	Mode RoutingIsochronePostResponsePropertiesMode `json:"mode"`
	// Travel time budget in seconds
	TimeSeconds float64 `json:"time_seconds"`
	// Number of road network vertices within the isochrone
	VerticesReached int64                                      `json:"vertices_reached"`
	JSON            routingIsochronePostResponsePropertiesJSON `json:"-"`
}

// routingIsochronePostResponsePropertiesJSON contains the JSON metadata for the
// struct [RoutingIsochronePostResponseProperties]
type routingIsochronePostResponsePropertiesJSON struct {
	AreaM2          apijson.Field
	MaxCostS        apijson.Field
	Mode            apijson.Field
	TimeSeconds     apijson.Field
	VerticesReached apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *RoutingIsochronePostResponseProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingIsochronePostResponsePropertiesJSON) RawJSON() string {
	return r.raw
}

// Travel mode used for the isochrone calculation
type RoutingIsochronePostResponsePropertiesMode string

const (
	RoutingIsochronePostResponsePropertiesModeAuto    RoutingIsochronePostResponsePropertiesMode = "auto"
	RoutingIsochronePostResponsePropertiesModeFoot    RoutingIsochronePostResponsePropertiesMode = "foot"
	RoutingIsochronePostResponsePropertiesModeBicycle RoutingIsochronePostResponsePropertiesMode = "bicycle"
)

func (r RoutingIsochronePostResponsePropertiesMode) IsKnown() bool {
	switch r {
	case RoutingIsochronePostResponsePropertiesModeAuto, RoutingIsochronePostResponsePropertiesModeFoot, RoutingIsochronePostResponsePropertiesModeBicycle:
		return true
	}
	return false
}

// `Feature` for single contour, `FeatureCollection` for multiple contours
type RoutingIsochronePostResponseType string

const (
	RoutingIsochronePostResponseTypeFeature           RoutingIsochronePostResponseType = "Feature"
	RoutingIsochronePostResponseTypeFeatureCollection RoutingIsochronePostResponseType = "FeatureCollection"
)

func (r RoutingIsochronePostResponseType) IsKnown() bool {
	switch r {
	case RoutingIsochronePostResponseTypeFeature, RoutingIsochronePostResponseTypeFeatureCollection:
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
	// Response format: json (default), geojson, csv, ndjson
	Format param.Field[string] `query:"format"`
	// Travel mode (auto, foot, bicycle)
	Mode param.Field[string] `query:"mode"`
	// Comma-separated property fields to include
	OutputFields param.Field[string] `query:"output[fields]"`
	// Include geometry (default true)
	OutputGeometry param.Field[bool] `query:"output[geometry]"`
	// Extra computed fields: bbox, center
	OutputInclude param.Field[string] `query:"output[include]"`
	// Coordinate decimal precision (1-15, default 7)
	OutputPrecision param.Field[int64] `query:"output[precision]"`
	// Simplify geometry tolerance in meters
	OutputSimplify param.Field[float64] `query:"output[simplify]"`
}

// URLQuery serializes [RoutingIsochroneParams]'s query parameters as `url.Values`.
func (r RoutingIsochroneParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoutingIsochronePostParams struct {
	// Latitude
	Lat param.Field[float64] `query:"lat" api:"required"`
	// Longitude
	Lng param.Field[float64] `query:"lng" api:"required"`
	// Travel time in seconds (1-7200)
	Time param.Field[float64] `query:"time" api:"required"`
	// Response format: json (default), geojson, csv, ndjson
	Format param.Field[string] `query:"format"`
	// Travel mode (auto, foot, bicycle)
	Mode param.Field[string] `query:"mode"`
	// Comma-separated property fields to include
	OutputFields param.Field[string] `query:"output[fields]"`
	// Include geometry (default true)
	OutputGeometry param.Field[bool] `query:"output[geometry]"`
	// Extra computed fields: bbox, center
	OutputInclude param.Field[string] `query:"output[include]"`
	// Coordinate decimal precision (1-15, default 7)
	OutputPrecision param.Field[int64] `query:"output[precision]"`
	// Simplify geometry tolerance in meters
	OutputSimplify param.Field[float64] `query:"output[simplify]"`
}

// URLQuery serializes [RoutingIsochronePostParams]'s query parameters as
// `url.Values`.
func (r RoutingIsochronePostParams) URLQuery() (v url.Values) {
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
	// Latitude
	Lat param.Field[float64] `query:"lat" api:"required"`
	// Longitude
	Lng param.Field[float64] `query:"lng" api:"required"`
	// Comma-separated property fields to include
	OutputFields param.Field[string] `query:"output[fields]"`
	// Extra computed fields: bbox, distance, center
	OutputInclude param.Field[string] `query:"output[include]"`
	// Coordinate decimal precision (1-15, default 7)
	OutputPrecision param.Field[int64] `query:"output[precision]"`
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

type RoutingNearestPostParams struct {
	// Latitude
	Lat param.Field[float64] `query:"lat" api:"required"`
	// Longitude
	Lng param.Field[float64] `query:"lng" api:"required"`
	// Comma-separated property fields to include
	OutputFields param.Field[string] `query:"output[fields]"`
	// Extra computed fields: bbox, distance, center
	OutputInclude param.Field[string] `query:"output[include]"`
	// Coordinate decimal precision (1-15, default 7)
	OutputPrecision param.Field[int64] `query:"output[precision]"`
	// Search radius in meters (default 500, max 5000)
	Radius param.Field[int64] `query:"radius"`
}

// URLQuery serializes [RoutingNearestPostParams]'s query parameters as
// `url.Values`.
func (r RoutingNearestPostParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type RoutingRouteParams struct {
	// Request body for route calculation. Origin and destination are lat/lng
	// coordinate objects. Supports optional waypoints, alternative routes,
	// turn-by-turn steps, and EV routing parameters.
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
