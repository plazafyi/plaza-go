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

// ElevationService contains methods and other services that help with interacting
// with the plaza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewElevationService] method instead.
type ElevationService struct {
	Options []option.RequestOption
}

// NewElevationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewElevationService(opts ...option.RequestOption) (r *ElevationService) {
	r = &ElevationService{}
	r.Options = opts
	return
}

// Look up elevation for multiple coordinates
func (r *ElevationService) Batch(ctx context.Context, body ElevationBatchParams, opts ...option.RequestOption) (res *ElevationBatchResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/elevation/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Look up elevation at one or more points
func (r *ElevationService) Lookup(ctx context.Context, query ElevationLookupParams, opts ...option.RequestOption) (res *ElevationLookupResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/elevation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Look up elevation at one or more points
func (r *ElevationService) LookupPost(ctx context.Context, body ElevationLookupPostParams, opts ...option.RequestOption) (res *ElevationLookupResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/elevation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Elevation profile along coordinates
func (r *ElevationService) Profile(ctx context.Context, body ElevationProfileParams, opts ...option.RequestOption) (res *ElevationProfileResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/elevation/profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// GeoJSON FeatureCollection of elevation Point Features with 3D coordinates. Order
// matches the input coordinates array.
type ElevationBatchResult struct {
	// Elevation results in the same order as input coordinates
	Features []ElevationLookupResult  `json:"features" api:"required"`
	Type     ElevationBatchResultType `json:"type" api:"required"`
	JSON     elevationBatchResultJSON `json:"-"`
}

// elevationBatchResultJSON contains the JSON metadata for the struct
// [ElevationBatchResult]
type elevationBatchResultJSON struct {
	Features    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ElevationBatchResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r elevationBatchResultJSON) RawJSON() string {
	return r.raw
}

type ElevationBatchResultType string

const (
	ElevationBatchResultTypeFeatureCollection ElevationBatchResultType = "FeatureCollection"
)

func (r ElevationBatchResultType) IsKnown() bool {
	switch r {
	case ElevationBatchResultTypeFeatureCollection:
		return true
	}
	return false
}

// GeoJSON Point Feature with a 3D coordinate [lng, lat, elevation] per RFC 7946
// §3.1.1. The elevation is also available in `properties.elevation_m` for
// convenience.
type ElevationLookupResult struct {
	// GeoJSON Geometry object per RFC 7946. Coordinates use [longitude, latitude]
	// order. 3D coordinates [lng, lat, elevation] are used for elevation endpoints.
	Geometry   GeoJsonGeometry                 `json:"geometry" api:"required"`
	Properties ElevationLookupResultProperties `json:"properties" api:"required"`
	Type       ElevationLookupResultType       `json:"type" api:"required"`
	JSON       elevationLookupResultJSON       `json:"-"`
}

// elevationLookupResultJSON contains the JSON metadata for the struct
// [ElevationLookupResult]
type elevationLookupResultJSON struct {
	Geometry    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ElevationLookupResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r elevationLookupResultJSON) RawJSON() string {
	return r.raw
}

type ElevationLookupResultProperties struct {
	// Elevation in meters above mean sea level (WGS84 EGM96 geoid)
	ElevationM float64                             `json:"elevation_m" api:"required"`
	JSON       elevationLookupResultPropertiesJSON `json:"-"`
}

// elevationLookupResultPropertiesJSON contains the JSON metadata for the struct
// [ElevationLookupResultProperties]
type elevationLookupResultPropertiesJSON struct {
	ElevationM  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ElevationLookupResultProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r elevationLookupResultPropertiesJSON) RawJSON() string {
	return r.raw
}

type ElevationLookupResultType string

const (
	ElevationLookupResultTypeFeature ElevationLookupResultType = "Feature"
)

func (r ElevationLookupResultType) IsKnown() bool {
	switch r {
	case ElevationLookupResultTypeFeature:
		return true
	}
	return false
}

// Request body for elevation profile along a path. Provide at least 2 coordinates
// defining the path. Maximum 50 coordinates per request.
type ElevationProfileRequestParam struct {
	// Path coordinates in order of travel (min 2, max 50)
	Coordinates param.Field[[]ElevationProfileRequestCoordinateParam] `json:"coordinates" api:"required"`
}

func (r ElevationProfileRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Geographic coordinate as a JSON object with `lat` and `lng` fields.
type ElevationProfileRequestCoordinateParam struct {
	// Latitude in decimal degrees (-90 to 90)
	Lat param.Field[float64] `json:"lat" api:"required"`
	// Longitude in decimal degrees (-180 to 180)
	Lng param.Field[float64] `json:"lng" api:"required"`
}

func (r ElevationProfileRequestCoordinateParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// GeoJSON LineString Feature with 3D coordinates [lng, lat, elevation]
// representing the elevation profile along the input path. Summary statistics are
// in properties.
type ElevationProfileResult struct {
	// GeoJSON Geometry object per RFC 7946. Coordinates use [longitude, latitude]
	// order. 3D coordinates [lng, lat, elevation] are used for elevation endpoints.
	Geometry GeoJsonGeometry `json:"geometry" api:"required"`
	// Elevation profile summary statistics
	Properties ElevationProfileResultProperties `json:"properties" api:"required"`
	Type       ElevationProfileResultType       `json:"type" api:"required"`
	JSON       elevationProfileResultJSON       `json:"-"`
}

// elevationProfileResultJSON contains the JSON metadata for the struct
// [ElevationProfileResult]
type elevationProfileResultJSON struct {
	Geometry    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ElevationProfileResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r elevationProfileResultJSON) RawJSON() string {
	return r.raw
}

// Elevation profile summary statistics
type ElevationProfileResultProperties struct {
	// Average elevation along the profile in meters
	AvgElevationM float64 `json:"avg_elevation_m" api:"required"`
	// Maximum elevation along the profile in meters
	MaxElevationM float64 `json:"max_elevation_m" api:"required"`
	// Minimum elevation along the profile in meters
	MinElevationM float64 `json:"min_elevation_m" api:"required"`
	// Total cumulative elevation gain in meters
	TotalAscentM float64 `json:"total_ascent_m" api:"required"`
	// Total cumulative elevation loss in meters
	TotalDescentM float64                              `json:"total_descent_m" api:"required"`
	JSON          elevationProfileResultPropertiesJSON `json:"-"`
}

// elevationProfileResultPropertiesJSON contains the JSON metadata for the struct
// [ElevationProfileResultProperties]
type elevationProfileResultPropertiesJSON struct {
	AvgElevationM apijson.Field
	MaxElevationM apijson.Field
	MinElevationM apijson.Field
	TotalAscentM  apijson.Field
	TotalDescentM apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ElevationProfileResultProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r elevationProfileResultPropertiesJSON) RawJSON() string {
	return r.raw
}

type ElevationProfileResultType string

const (
	ElevationProfileResultTypeFeature ElevationProfileResultType = "Feature"
)

func (r ElevationProfileResultType) IsKnown() bool {
	switch r {
	case ElevationProfileResultTypeFeature:
		return true
	}
	return false
}

type ElevationBatchParams struct {
	// Coordinates to look up elevations for (max 50)
	Coordinates param.Field[[]ElevationBatchParamsCoordinate] `json:"coordinates" api:"required"`
}

func (r ElevationBatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Geographic coordinate as a JSON object with `lat` and `lng` fields.
type ElevationBatchParamsCoordinate struct {
	// Latitude in decimal degrees (-90 to 90)
	Lat param.Field[float64] `json:"lat" api:"required"`
	// Longitude in decimal degrees (-180 to 180)
	Lng param.Field[float64] `json:"lng" api:"required"`
}

func (r ElevationBatchParamsCoordinate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ElevationLookupParams struct {
	// Latitude (single point)
	Lat param.Field[float64] `query:"lat"`
	// Longitude (single point)
	Lng param.Field[float64] `query:"lng"`
	// Pipe-separated lng,lat pairs (batch)
	Locations param.Field[string] `query:"locations"`
	// Comma-separated property fields to include
	OutputFields param.Field[string] `query:"output[fields]"`
	// Extra computed fields: bbox, center
	OutputInclude param.Field[string] `query:"output[include]"`
	// Coordinate decimal precision (1-15, default 7)
	OutputPrecision param.Field[int64] `query:"output[precision]"`
}

// URLQuery serializes [ElevationLookupParams]'s query parameters as `url.Values`.
func (r ElevationLookupParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ElevationLookupPostParams struct {
	// Latitude (single point)
	Lat param.Field[float64] `query:"lat"`
	// Longitude (single point)
	Lng param.Field[float64] `query:"lng"`
	// Pipe-separated lng,lat pairs (batch)
	Locations param.Field[string] `query:"locations"`
	// Comma-separated property fields to include
	OutputFields param.Field[string] `query:"output[fields]"`
	// Extra computed fields: bbox, center
	OutputInclude param.Field[string] `query:"output[include]"`
	// Coordinate decimal precision (1-15, default 7)
	OutputPrecision param.Field[int64] `query:"output[precision]"`
}

// URLQuery serializes [ElevationLookupPostParams]'s query parameters as
// `url.Values`.
func (r ElevationLookupPostParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ElevationProfileParams struct {
	// Request body for elevation profile along a path. Provide at least 2 coordinates
	// defining the path. Maximum 50 coordinates per request.
	ElevationProfileRequest ElevationProfileRequestParam `json:"elevation_profile_request" api:"required"`
}

func (r ElevationProfileParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ElevationProfileRequest)
}
