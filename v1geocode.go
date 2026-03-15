// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plaza

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/plazafyi/plaza-go/internal/apijson"
	"github.com/plazafyi/plaza-go/internal/apiquery"
	"github.com/plazafyi/plaza-go/internal/requestconfig"
	"github.com/plazafyi/plaza-go/option"
	"github.com/plazafyi/plaza-go/packages/param"
	"github.com/plazafyi/plaza-go/packages/respjson"
)

// V1GeocodeService contains methods and other services that help with interacting
// with the plaza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1GeocodeService] method instead.
type V1GeocodeService struct {
	options []option.RequestOption
}

// NewV1GeocodeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1GeocodeService(opts ...option.RequestOption) (r V1GeocodeService) {
	r = V1GeocodeService{}
	r.options = opts
	return
}

// Autocomplete a partial address
func (r *V1GeocodeService) Autocomplete(ctx context.Context, query V1GeocodeAutocompleteParams, opts ...option.RequestOption) (res *V1GeocodeAutocompleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/geocode/autocomplete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Forward geocode an address
func (r *V1GeocodeService) Forward(ctx context.Context, query V1GeocodeForwardParams, opts ...option.RequestOption) (res *V1GeocodeForwardResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/geocode"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Reverse geocode a coordinate
func (r *V1GeocodeService) Reverse(ctx context.Context, query V1GeocodeReverseParams, opts ...option.RequestOption) (res *V1GeocodeReverseResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/geocode/reverse"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type V1GeocodeAutocompleteResponse struct {
	Results []V1GeocodeAutocompleteResponseResult `json:"results" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1GeocodeAutocompleteResponse) RawJSON() string { return r.JSON.raw }
func (r *V1GeocodeAutocompleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1GeocodeAutocompleteResponseResult struct {
	// Suggested address or place name
	DisplayName string `json:"display_name" api:"required"`
	// Latitude
	Lat float64 `json:"lat" api:"nullable"`
	// Longitude
	Lng float64 `json:"lng" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		Lat         respjson.Field
		Lng         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1GeocodeAutocompleteResponseResult) RawJSON() string { return r.JSON.raw }
func (r *V1GeocodeAutocompleteResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1GeocodeForwardResponse struct {
	Results []V1GeocodeForwardResponseResult `json:"results" api:"required"`
	// Number of results
	Count int64 `json:"count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Results     respjson.Field
		Count       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1GeocodeForwardResponse) RawJSON() string { return r.JSON.raw }
func (r *V1GeocodeForwardResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1GeocodeForwardResponseResult struct {
	// Formatted address or place name
	DisplayName string `json:"display_name" api:"required"`
	// Latitude
	Lat float64 `json:"lat" api:"required"`
	// Longitude
	Lng float64 `json:"lng" api:"required"`
	// OpenStreetMap ID
	OsmID int64 `json:"osm_id" api:"nullable"`
	// OSM element type
	OsmType string `json:"osm_type" api:"nullable"`
	// Match confidence score
	Score float64 `json:"score" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		Lat         respjson.Field
		Lng         respjson.Field
		OsmID       respjson.Field
		OsmType     respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1GeocodeForwardResponseResult) RawJSON() string { return r.JSON.raw }
func (r *V1GeocodeForwardResponseResult) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1GeocodeReverseResponse struct {
	// Formatted address
	Address string `json:"address" api:"required"`
	// Distance in meters from query point
	Distance float64 `json:"distance" api:"required"`
	// Latitude
	Lat float64 `json:"lat" api:"required"`
	// Longitude
	Lng float64 `json:"lng" api:"required"`
	// OpenStreetMap ID
	OsmID int64 `json:"osm_id" api:"nullable"`
	// OSM element type
	OsmType string `json:"osm_type" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address     respjson.Field
		Distance    respjson.Field
		Lat         respjson.Field
		Lng         respjson.Field
		OsmID       respjson.Field
		OsmType     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1GeocodeReverseResponse) RawJSON() string { return r.JSON.raw }
func (r *V1GeocodeReverseResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1GeocodeAutocompleteParams struct {
	// Partial address query
	Q string `query:"q" api:"required" json:"-"`
	// Focus latitude
	Lat param.Opt[float64] `query:"lat,omitzero" json:"-"`
	// Maximum results (default 10, max 20)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Focus longitude
	Lng param.Opt[float64] `query:"lng,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1GeocodeAutocompleteParams]'s query parameters as
// `url.Values`.
func (r V1GeocodeAutocompleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1GeocodeForwardParams struct {
	// Address or place name
	Q string `query:"q" api:"required" json:"-"`
	// Focus latitude
	Lat param.Opt[float64] `query:"lat,omitzero" json:"-"`
	// Maximum results (default 20, max 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Focus longitude
	Lng param.Opt[float64] `query:"lng,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1GeocodeForwardParams]'s query parameters as `url.Values`.
func (r V1GeocodeForwardParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1GeocodeReverseParams struct {
	// Latitude
	Lat float64 `query:"lat" api:"required" json:"-"`
	// Longitude
	Lng float64 `query:"lng" api:"required" json:"-"`
	// Search radius in meters (default 200, max 5000)
	Radius param.Opt[int64] `query:"radius,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1GeocodeReverseParams]'s query parameters as `url.Values`.
func (r V1GeocodeReverseParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
