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

// MapMatchService contains methods and other services that help with interacting
// with the plaza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMapMatchService] method instead.
type MapMatchService struct {
	Options []option.RequestOption
}

// NewMapMatchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMapMatchService(opts ...option.RequestOption) (r *MapMatchService) {
	r = &MapMatchService{}
	r.Options = opts
	return
}

// Match GPS coordinates to the road network
func (r *MapMatchService) Match(ctx context.Context, body MapMatchMatchParams, opts ...option.RequestOption) (res *MapMatchResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/map-match"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// GPS trace to snap to the road network. Provide an array of coordinate objects
// representing the GPS points. Maximum 50 points per request.
type MapMatchRequestParam struct {
	// GPS coordinates to match, in order of travel (max 50 points)
	Coordinates param.Field[[]MapMatchRequestCoordinateParam] `json:"coordinates" api:"required"`
	// Search radius per coordinate in meters. Must have the same length as
	// `coordinates` or be omitted entirely. Default: 50m per point.
	Radiuses param.Field[[]float64] `json:"radiuses"`
}

func (r MapMatchRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Geographic coordinate as a JSON object with `lat` and `lng` fields.
type MapMatchRequestCoordinateParam struct {
	// Latitude in decimal degrees (-90 to 90)
	Lat param.Field[float64] `json:"lat" api:"required"`
	// Longitude in decimal degrees (-180 to 180)
	Lng param.Field[float64] `json:"lng" api:"required"`
}

func (r MapMatchRequestCoordinateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Map matching result as a GeoJSON FeatureCollection. Each Feature is a snapped
// tracepoint. The top-level `matchings` array contains the matched sub-routes
// connecting consecutive tracepoints.
type MapMatchResult struct {
	// Snapped tracepoint Features in input order
	Features []MapMatchResultFeature `json:"features" api:"required"`
	// Matched sub-routes. Each matching connects a contiguous sequence of tracepoints
	// that could be matched to roads.
	Matchings []map[string]interface{} `json:"matchings" api:"required"`
	Type      MapMatchResultType       `json:"type" api:"required"`
	JSON      mapMatchResultJSON       `json:"-"`
}

// mapMatchResultJSON contains the JSON metadata for the struct [MapMatchResult]
type mapMatchResultJSON struct {
	Features    apijson.Field
	Matchings   apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MapMatchResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mapMatchResultJSON) RawJSON() string {
	return r.raw
}

// GeoJSON Point Feature representing a GPS point snapped to the road network.
type MapMatchResultFeature struct {
	// GeoJSON Geometry object per RFC 7946. Coordinates use [longitude, latitude]
	// order. 3D coordinates [lng, lat, elevation] are used for elevation endpoints.
	Geometry   GeoJsonGeometry                  `json:"geometry" api:"required"`
	Properties MapMatchResultFeaturesProperties `json:"properties" api:"required"`
	Type       MapMatchResultFeaturesType       `json:"type" api:"required"`
	JSON       mapMatchResultFeatureJSON        `json:"-"`
}

// mapMatchResultFeatureJSON contains the JSON metadata for the struct
// [MapMatchResultFeature]
type mapMatchResultFeatureJSON struct {
	Geometry    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MapMatchResultFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mapMatchResultFeatureJSON) RawJSON() string {
	return r.raw
}

type MapMatchResultFeaturesProperties struct {
	// Distance from the original GPS point to the snapped point in meters
	DistanceM float64 `json:"distance_m"`
	// Road edge ID the point was snapped to
	EdgeID int64 `json:"edge_id"`
	// Index into the `matchings` array indicating which matching sub-route this point
	// belongs to
	MatchingsIndex int64 `json:"matchings_index"`
	// Road name at the snapped point
	Name string `json:"name" api:"nullable"`
	// Original GPS coordinate as [lng, lat]
	Original []float64 `json:"original"`
	// Index of this tracepoint in the original `coordinates` array
	WaypointIndex int64                                `json:"waypoint_index"`
	JSON          mapMatchResultFeaturesPropertiesJSON `json:"-"`
}

// mapMatchResultFeaturesPropertiesJSON contains the JSON metadata for the struct
// [MapMatchResultFeaturesProperties]
type mapMatchResultFeaturesPropertiesJSON struct {
	DistanceM      apijson.Field
	EdgeID         apijson.Field
	MatchingsIndex apijson.Field
	Name           apijson.Field
	Original       apijson.Field
	WaypointIndex  apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *MapMatchResultFeaturesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mapMatchResultFeaturesPropertiesJSON) RawJSON() string {
	return r.raw
}

type MapMatchResultFeaturesType string

const (
	MapMatchResultFeaturesTypeFeature MapMatchResultFeaturesType = "Feature"
)

func (r MapMatchResultFeaturesType) IsKnown() bool {
	switch r {
	case MapMatchResultFeaturesTypeFeature:
		return true
	}
	return false
}

type MapMatchResultType string

const (
	MapMatchResultTypeFeatureCollection MapMatchResultType = "FeatureCollection"
)

func (r MapMatchResultType) IsKnown() bool {
	switch r {
	case MapMatchResultTypeFeatureCollection:
		return true
	}
	return false
}

type MapMatchMatchParams struct {
	// GPS trace to snap to the road network. Provide an array of coordinate objects
	// representing the GPS points. Maximum 50 points per request.
	MapMatchRequest MapMatchRequestParam `json:"map_match_request" api:"required"`
}

func (r MapMatchMatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MapMatchRequest)
}
