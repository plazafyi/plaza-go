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
	path := "api/v1/geocode/autocomplete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Autocomplete a partial address
func (r *GeocodeService) AutocompletePost(ctx context.Context, body GeocodeAutocompletePostParams, opts ...option.RequestOption) (res *AutocompleteResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/geocode/autocomplete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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
	path := "api/v1/geocode"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Forward geocode an address
func (r *GeocodeService) ForwardPost(ctx context.Context, body GeocodeForwardPostParams, opts ...option.RequestOption) (res *GeocodeResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/geocode"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Reverse geocode a coordinate
func (r *GeocodeService) Reverse(ctx context.Context, query GeocodeReverseParams, opts ...option.RequestOption) (res *ReverseGeocodeResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/geocode/reverse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Reverse geocode a coordinate
func (r *GeocodeService) ReversePost(ctx context.Context, body GeocodeReversePostParams, opts ...option.RequestOption) (res *ReverseGeocodeResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/geocode/reverse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// GeoJSON FeatureCollection of autocomplete suggestions for partial address input.
// Optimized for low-latency type-ahead UIs. Content-Type: `application/geo+json`.
type AutocompleteResult struct {
	// Autocomplete suggestions ordered by relevance
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

// GeoJSON FeatureCollection of forward geocoding results, ordered by relevance.
// Content-Type: `application/geo+json`.
type GeocodeResult struct {
	// Geocoding results ordered by relevance score
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

// GeoJSON Feature representing a geocoding result. The geometry is always a Point.
// Properties include the formatted display name, OSM metadata, confidence score,
// and source type.
type GeocodingFeature struct {
	// GeoJSON Geometry object per RFC 7946. Coordinates use [longitude, latitude]
	// order. 3D coordinates [lng, lat, elevation] are used for elevation endpoints.
	Geometry GeoJsonGeometry `json:"geometry" api:"required"`
	// Geocoding result properties
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

// Geocoding result properties
type GeocodingFeatureProperties struct {
	// Formatted address or place name
	DisplayName string `json:"display_name" api:"required"`
	// POI category (e.g. restaurant, cafe, park). Present for place results.
	Category string `json:"category" api:"nullable"`
	// City or town name. Present for address results.
	City string `json:"city" api:"nullable"`
	// Interpolation confidence (0-1). Present only for interpolated results.
	Confidence float64 `json:"confidence" api:"nullable"`
	// Country name. Present for reverse geocode address results.
	Country string `json:"country" api:"nullable"`
	// ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code" api:"nullable"`
	// Distance from the query point in meters (reverse geocode / nearby only)
	DistanceM float64 `json:"distance_m" api:"nullable"`
	// Complete formatted address from the database. Present for reverse geocode
	// address results.
	FullAddress string `json:"full_address" api:"nullable"`
	// House or building number. Present for address and interpolated results.
	HouseNumber string `json:"house_number" api:"nullable"`
	// Whether this result was estimated by address interpolation rather than an exact
	// database match.
	Interpolated bool `json:"interpolated" api:"nullable"`
	// Place name (raw). Present for reverse geocode place results.
	Name string `json:"name" api:"nullable"`
	// OpenStreetMap element ID (null for interpolated results)
	OsmID int64 `json:"osm_id" api:"nullable"`
	// OSM element type (node, way, relation)
	OsmType GeocodingFeaturePropertiesOsmType `json:"osm_type" api:"nullable"`
	// Postal code. Present for reverse geocode address results.
	Postcode string `json:"postcode" api:"nullable"`
	// Relevance score (higher is better). Incorporates text match quality, spatial
	// proximity boost, and popularity signals. Not bounded to 0-1.
	Score float64 `json:"score" api:"nullable"`
	// Result source indicating how the result was found: structured (exact field
	// match), bm25 (full-text search), fuzzy (trigram similarity), address (reverse
	// geocode address), place (reverse geocode POI), interpolation (estimated from
	// neighboring addresses)
	Source GeocodingFeaturePropertiesSource `json:"source" api:"nullable"`
	// State or province name. Present for reverse geocode address results.
	State string `json:"state" api:"nullable"`
	// Street name. Present for address and interpolated results.
	Street string `json:"street" api:"nullable"`
	// POI subcategory. Present for place results.
	Subcategory string `json:"subcategory" api:"nullable"`
	// Raw OSM tags. Present for place results.
	Tags map[string]string `json:"tags" api:"nullable"`
	// Wikipedia article reference (e.g. en:Eiffel Tower). Present for notable places.
	Wikipedia string                         `json:"wikipedia" api:"nullable"`
	JSON      geocodingFeaturePropertiesJSON `json:"-"`
}

// geocodingFeaturePropertiesJSON contains the JSON metadata for the struct
// [GeocodingFeatureProperties]
type geocodingFeaturePropertiesJSON struct {
	DisplayName  apijson.Field
	Category     apijson.Field
	City         apijson.Field
	Confidence   apijson.Field
	Country      apijson.Field
	CountryCode  apijson.Field
	DistanceM    apijson.Field
	FullAddress  apijson.Field
	HouseNumber  apijson.Field
	Interpolated apijson.Field
	Name         apijson.Field
	OsmID        apijson.Field
	OsmType      apijson.Field
	Postcode     apijson.Field
	Score        apijson.Field
	Source       apijson.Field
	State        apijson.Field
	Street       apijson.Field
	Subcategory  apijson.Field
	Tags         apijson.Field
	Wikipedia    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *GeocodingFeatureProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geocodingFeaturePropertiesJSON) RawJSON() string {
	return r.raw
}

// OSM element type (node, way, relation)
type GeocodingFeaturePropertiesOsmType string

const (
	GeocodingFeaturePropertiesOsmTypeNode     GeocodingFeaturePropertiesOsmType = "node"
	GeocodingFeaturePropertiesOsmTypeWay      GeocodingFeaturePropertiesOsmType = "way"
	GeocodingFeaturePropertiesOsmTypeRelation GeocodingFeaturePropertiesOsmType = "relation"
)

func (r GeocodingFeaturePropertiesOsmType) IsKnown() bool {
	switch r {
	case GeocodingFeaturePropertiesOsmTypeNode, GeocodingFeaturePropertiesOsmTypeWay, GeocodingFeaturePropertiesOsmTypeRelation:
		return true
	}
	return false
}

// Result source indicating how the result was found: structured (exact field
// match), bm25 (full-text search), fuzzy (trigram similarity), address (reverse
// geocode address), place (reverse geocode POI), interpolation (estimated from
// neighboring addresses)
type GeocodingFeaturePropertiesSource string

const (
	GeocodingFeaturePropertiesSourceStructured    GeocodingFeaturePropertiesSource = "structured"
	GeocodingFeaturePropertiesSourceBm25          GeocodingFeaturePropertiesSource = "bm25"
	GeocodingFeaturePropertiesSourceFuzzy         GeocodingFeaturePropertiesSource = "fuzzy"
	GeocodingFeaturePropertiesSourceAddress       GeocodingFeaturePropertiesSource = "address"
	GeocodingFeaturePropertiesSourcePlace         GeocodingFeaturePropertiesSource = "place"
	GeocodingFeaturePropertiesSourceInterpolation GeocodingFeaturePropertiesSource = "interpolation"
)

func (r GeocodingFeaturePropertiesSource) IsKnown() bool {
	switch r {
	case GeocodingFeaturePropertiesSourceStructured, GeocodingFeaturePropertiesSourceBm25, GeocodingFeaturePropertiesSourceFuzzy, GeocodingFeaturePropertiesSourceAddress, GeocodingFeaturePropertiesSourcePlace, GeocodingFeaturePropertiesSourceInterpolation:
		return true
	}
	return false
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

// GeoJSON FeatureCollection of reverse geocoding results, ordered by distance from
// the query point. Content-Type: `application/geo+json`.
type ReverseGeocodeResult struct {
	// Reverse geocoding results ordered by distance
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

// Batch geocoding result. Each entry in `results` is a FeatureCollection
// corresponding to the input address at the same index. Order is preserved.
type GeocodeBatchResponse struct {
	// Number of addresses processed (always equals length of results)
	Count int64 `json:"count" api:"required"`
	// Array of FeatureCollections, one per input address. Empty FeatureCollections
	// indicate no match.
	Results []GeocodeResult          `json:"results" api:"required"`
	JSON    geocodeBatchResponseJSON `json:"-"`
}

// geocodeBatchResponseJSON contains the JSON metadata for the struct
// [GeocodeBatchResponse]
type geocodeBatchResponseJSON struct {
	Count       apijson.Field
	Results     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeocodeBatchResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geocodeBatchResponseJSON) RawJSON() string {
	return r.raw
}

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

type GeocodeAutocompletePostParams struct {
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

// URLQuery serializes [GeocodeAutocompletePostParams]'s query parameters as
// `url.Values`.
func (r GeocodeAutocompletePostParams) URLQuery() (v url.Values) {
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

type GeocodeForwardPostParams struct {
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

// URLQuery serializes [GeocodeForwardPostParams]'s query parameters as
// `url.Values`.
func (r GeocodeForwardPostParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type GeocodeReverseParams struct {
	// Language code for localized names (e.g. en, de, fr)
	Lang param.Field[string] `query:"lang"`
	// Legacy shorthand. Latitude. Use near param instead.
	Lat param.Field[float64] `query:"lat"`
	// Filter by layer: house or poi
	Layer param.Field[string] `query:"layer"`
	// Maximum results (default 1, max 20)
	Limit param.Field[int64] `query:"limit"`
	// Legacy shorthand. Longitude. Use near param instead.
	Lng param.Field[float64] `query:"lng"`
	// Point geometry for reverse geocode (lat,lng or GeoJSON). Alternative to lat/lng
	// params.
	Near param.Field[string] `query:"near"`
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

type GeocodeReversePostParams struct {
	// Language code for localized names (e.g. en, de, fr)
	Lang param.Field[string] `query:"lang"`
	// Legacy shorthand. Latitude. Use near param instead.
	Lat param.Field[float64] `query:"lat"`
	// Filter by layer: house or poi
	Layer param.Field[string] `query:"layer"`
	// Maximum results (default 1, max 20)
	Limit param.Field[int64] `query:"limit"`
	// Legacy shorthand. Longitude. Use near param instead.
	Lng param.Field[float64] `query:"lng"`
	// Point geometry for reverse geocode (lat,lng or GeoJSON). Alternative to lat/lng
	// params.
	Near param.Field[string] `query:"near"`
	// Search radius in meters (default 200, max 5000)
	Radius param.Field[int64] `query:"radius"`
}

// URLQuery serializes [GeocodeReversePostParams]'s query parameters as
// `url.Values`.
func (r GeocodeReversePostParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
