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

// GeocodeService contains methods and other services that help with interacting
// with the plaza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewGeocodeService] method instead.
type GeocodeService struct {
	Options []option.RequestOption
}

// NewGeocodeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewGeocodeService(opts ...option.RequestOption) (r *GeocodeService) {
	r = &GeocodeService{}
	r.Options = opts
	return
}

// Autocomplete a partial address
func (r *GeocodeService) Autocomplete(ctx context.Context, query GeocodeAutocompleteParams, opts ...option.RequestOption) (res *AutocompleteResult, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/geo+json")}, opts...)
	path := "api/v1/geocode/autocomplete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Batch geocode multiple addresses
func (r *GeocodeService) Batch(ctx context.Context, body GeocodeBatchParams, opts ...option.RequestOption) (res *GeocodeBatchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/geocode/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Forward geocode an address
func (r *GeocodeService) Forward(ctx context.Context, query GeocodeForwardParams, opts ...option.RequestOption) (res *GeocodeResult, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/geo+json")}, opts...)
	path := "api/v1/geocode"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Reverse geocode a coordinate
func (r *GeocodeService) Reverse(ctx context.Context, query GeocodeReverseParams, opts ...option.RequestOption) (res *ReverseGeocodeResult, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/geo+json")}, opts...)
	path := "api/v1/geocode/reverse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// GeoJSON FeatureCollection of autocomplete suggestions
type AutocompleteResult struct {
	Features []GeocodingFeature     `json:"features" api:"required"`
	Type     AutocompleteResultType `json:"type" api:"required"`
	JSON     autocompleteResultJSON `json:"-"`
}

// autocompleteResultJSON contains the JSON metadata for the struct
// [AutocompleteResult]
type autocompleteResultJSON struct {
	Features    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AutocompleteResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r autocompleteResultJSON) RawJSON() string {
	return r.raw
}

type AutocompleteResultType string

const (
	AutocompleteResultTypeFeatureCollection AutocompleteResultType = "FeatureCollection"
)

func (r AutocompleteResultType) IsKnown() bool {
	switch r {
	case AutocompleteResultTypeFeatureCollection:
		return true
	}
	return false
}

// GeoJSON FeatureCollection of geocoding results
type GeocodeResult struct {
	Features []GeocodingFeature `json:"features" api:"required"`
	Type     GeocodeResultType  `json:"type" api:"required"`
	JSON     geocodeResultJSON  `json:"-"`
}

// geocodeResultJSON contains the JSON metadata for the struct [GeocodeResult]
type geocodeResultJSON struct {
	Features    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeocodeResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geocodeResultJSON) RawJSON() string {
	return r.raw
}

type GeocodeResultType string

const (
	GeocodeResultTypeFeatureCollection GeocodeResultType = "FeatureCollection"
)

func (r GeocodeResultType) IsKnown() bool {
	switch r {
	case GeocodeResultTypeFeatureCollection:
		return true
	}
	return false
}

type GeocodingFeature struct {
	Geometry   GeoJsonGeometry            `json:"geometry" api:"required"`
	Properties GeocodingFeatureProperties `json:"properties" api:"required"`
	Type       GeocodingFeatureType       `json:"type" api:"required"`
	JSON       geocodingFeatureJSON       `json:"-"`
}

// geocodingFeatureJSON contains the JSON metadata for the struct
// [GeocodingFeature]
type geocodingFeatureJSON struct {
	Geometry    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeocodingFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geocodingFeatureJSON) RawJSON() string {
	return r.raw
}

type GeocodingFeatureProperties struct {
	// ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code" api:"nullable"`
	// Formatted address or place name
	DisplayName string `json:"display_name"`
	// Distance in meters
	DistanceM float64 `json:"distance_m" api:"nullable"`
	// OpenStreetMap ID
	OsmID int64 `json:"osm_id" api:"nullable"`
	// OSM element type
	OsmType string `json:"osm_type" api:"nullable"`
	// Match confidence score
	Score float64 `json:"score" api:"nullable"`
	// Result source (address, place, interpolation)
	Source string                         `json:"source" api:"nullable"`
	JSON   geocodingFeaturePropertiesJSON `json:"-"`
}

// geocodingFeaturePropertiesJSON contains the JSON metadata for the struct
// [GeocodingFeatureProperties]
type geocodingFeaturePropertiesJSON struct {
	CountryCode apijson.Field
	DisplayName apijson.Field
	DistanceM   apijson.Field
	OsmID       apijson.Field
	OsmType     apijson.Field
	Score       apijson.Field
	Source      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeocodingFeatureProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geocodingFeaturePropertiesJSON) RawJSON() string {
	return r.raw
}

type GeocodingFeatureType string

const (
	GeocodingFeatureTypeFeature GeocodingFeatureType = "Feature"
)

func (r GeocodingFeatureType) IsKnown() bool {
	switch r {
	case GeocodingFeatureTypeFeature:
		return true
	}
	return false
}

// GeoJSON FeatureCollection of reverse geocoding results
type ReverseGeocodeResult struct {
	Features []GeocodingFeature       `json:"features" api:"required"`
	Type     ReverseGeocodeResultType `json:"type" api:"required"`
	JSON     reverseGeocodeResultJSON `json:"-"`
}

// reverseGeocodeResultJSON contains the JSON metadata for the struct
// [ReverseGeocodeResult]
type reverseGeocodeResultJSON struct {
	Features    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ReverseGeocodeResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r reverseGeocodeResultJSON) RawJSON() string {
	return r.raw
}

type ReverseGeocodeResultType string

const (
	ReverseGeocodeResultTypeFeatureCollection ReverseGeocodeResultType = "FeatureCollection"
)

func (r ReverseGeocodeResultType) IsKnown() bool {
	switch r {
	case ReverseGeocodeResultTypeFeatureCollection:
		return true
	}
	return false
}

type GeocodeBatchResponse = interface{}

type GeocodeAutocompleteParams struct {
	// Partial address query
	Q param.Field[string] `query:"q" api:"required"`
	// ISO 3166-1 alpha-2 country code filter
	CountryCode param.Field[string] `query:"country_code"`
	// Language code for localized names (e.g. en, de, fr)
	Lang param.Field[string] `query:"lang"`
	// Focus latitude
	Lat param.Field[float64] `query:"lat"`
	// Filter by layer: address, poi, or admin
	Layer param.Field[string] `query:"layer"`
	// Maximum results (default 10, max 20)
	Limit param.Field[int64] `query:"limit"`
	// Focus longitude
	Lng param.Field[float64] `query:"lng"`
}

// URLQuery serializes [GeocodeAutocompleteParams]'s query parameters as
// `url.Values`.
func (r GeocodeAutocompleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GeocodeBatchParams struct {
	Addresses param.Field[[]string] `json:"addresses" api:"required"`
}

func (r GeocodeBatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type GeocodeForwardParams struct {
	// Address or place name
	Q param.Field[string] `query:"q" api:"required"`
	// Bounding box filter: south,west,north,east
	Bbox param.Field[string] `query:"bbox"`
	// ISO 3166-1 alpha-2 country code filter
	CountryCode param.Field[string] `query:"country_code"`
	// Language code for localized names (e.g. en, de, fr)
	Lang param.Field[string] `query:"lang"`
	// Focus latitude
	Lat param.Field[float64] `query:"lat"`
	// Filter by layer: address, poi, or admin
	Layer param.Field[string] `query:"layer"`
	// Maximum results (default 20, max 100)
	Limit param.Field[int64] `query:"limit"`
	// Focus longitude
	Lng param.Field[float64] `query:"lng"`
}

// URLQuery serializes [GeocodeForwardParams]'s query parameters as `url.Values`.
func (r GeocodeForwardParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GeocodeReverseParams struct {
	// Latitude
	Lat param.Field[float64] `query:"lat" api:"required"`
	// Longitude
	Lng param.Field[float64] `query:"lng" api:"required"`
	// Language code for localized names (e.g. en, de, fr)
	Lang param.Field[string] `query:"lang"`
	// Filter by layer: house or poi
	Layer param.Field[string] `query:"layer"`
	// Maximum results (default 1, max 20)
	Limit param.Field[int64] `query:"limit"`
	// Search radius in meters (default 200, max 5000)
	Radius param.Field[int64] `query:"radius"`
}

// URLQuery serializes [GeocodeReverseParams]'s query parameters as `url.Values`.
func (r GeocodeReverseParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
