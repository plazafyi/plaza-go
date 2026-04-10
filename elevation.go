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

// Look up elevation at one or more points
func (r *ElevationService) Lookup(ctx context.Context, params ElevationLookupParams, opts ...option.RequestOption) (res *ElevationLookupResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/elevation"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Elevation profile along coordinates
func (r *ElevationService) Profile(ctx context.Context, body ElevationProfileParams, opts ...option.RequestOption) (res *ElevationProfileResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/elevation/profile"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Request body for elevation lookup. Accepts a single Point or a MultiPoint
// geometry.
type ElevationLookupRequestParam struct {
	// Point or MultiPoint geometry to look up elevations for
	Geometry param.Field[ElevationLookupRequestGeometryUnionParam] `json:"geometry" api:"required"`
}

func (r ElevationLookupRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Point or MultiPoint geometry to look up elevations for
type ElevationLookupRequestGeometryParam struct {
	Coordinates param.Field[interface{}]                        `json:"coordinates" api:"required"`
	Type        param.Field[ElevationLookupRequestGeometryType] `json:"type" api:"required"`
}

func (r ElevationLookupRequestGeometryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r ElevationLookupRequestGeometryParam) implementsElevationLookupRequestGeometryUnionParam() {}

// Point or MultiPoint geometry to look up elevations for
//
// Satisfied by [PointGeometryParam], [MultiPointGeometryParam],
// [ElevationLookupRequestGeometryParam].
type ElevationLookupRequestGeometryUnionParam interface {
	implementsElevationLookupRequestGeometryUnionParam()
}

type ElevationLookupRequestGeometryType string

const (
	ElevationLookupRequestGeometryTypePoint      ElevationLookupRequestGeometryType = "Point"
	ElevationLookupRequestGeometryTypeMultiPoint ElevationLookupRequestGeometryType = "MultiPoint"
)

func (r ElevationLookupRequestGeometryType) IsKnown() bool {
	switch r {
	case ElevationLookupRequestGeometryTypePoint, ElevationLookupRequestGeometryTypeMultiPoint:
		return true
	}
	return false
}

// GeoJSON Point Feature with a 3D coordinate [lng, lat, elevation] per RFC 7946
// §3.1.1. The elevation is also available in `properties.elevation_m` for
// convenience.
type ElevationLookupResult struct {
	// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
	// determines the coordinate structure.
	Geometry   Geometry                        `json:"geometry" api:"required"`
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

// Request body for elevation profile along a path. Provide a GeoJSON LineString
// geometry defining the path.
type ElevationProfileRequestParam struct {
	// GeoJSON LineString geometry per RFC 7946. An ordered sequence of two or more
	// positions.
	Geometry param.Field[LineStringGeometryParam] `json:"geometry" api:"required"`
}

func (r ElevationProfileRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// GeoJSON LineString Feature with 3D coordinates [lng, lat, elevation]
// representing the elevation profile along the input path. Summary statistics are
// in properties.
type ElevationProfileResult struct {
	// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
	// determines the coordinate structure.
	Geometry Geometry `json:"geometry" api:"required"`
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

type ElevationLookupParams struct {
	// Request body for elevation lookup. Accepts a single Point or a MultiPoint
	// geometry.
	ElevationLookupRequest ElevationLookupRequestParam `json:"elevation_lookup_request" api:"required"`
	// Response format: json (default), geojson, csv, ndjson
	Format param.Field[string] `query:"format"`
}

func (r ElevationLookupParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ElevationLookupRequest)
}

// URLQuery serializes [ElevationLookupParams]'s query parameters as `url.Values`.
func (r ElevationLookupParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ElevationProfileParams struct {
	// Request body for elevation profile along a path. Provide a GeoJSON LineString
	// geometry defining the path.
	ElevationProfileRequest ElevationProfileRequestParam `json:"elevation_profile_request" api:"required"`
}

func (r ElevationProfileParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.ElevationProfileRequest)
}
