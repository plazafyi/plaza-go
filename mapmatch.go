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
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/geo+json")}, opts...)
	path := "api/v1/map-match"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// GPS trace to match against the road network
type MapMatchRequestParam struct {
	// GPS trace (GeoJSON LineString geometry)
	Trace param.Field[GeoJsonGeometryParam] `json:"trace" api:"required"`
	// Search radius per coordinate in meters (optional, default 50)
	Radiuses param.Field[[]float64] `json:"radiuses"`
}

func (r MapMatchRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Map matching result with snapped geometry
type MapMatchResult struct {
	Geometry   GeoJsonGeometry          `json:"geometry" api:"required"`
	Properties MapMatchResultProperties `json:"properties" api:"required"`
	Type       MapMatchResultType       `json:"type" api:"required"`
	// Matched route legs between consecutive trace points
	Legs []map[string]interface{} `json:"legs"`
	JSON mapMatchResultJSON       `json:"-"`
}

// mapMatchResultJSON contains the JSON metadata for the struct [MapMatchResult]
type mapMatchResultJSON struct {
	Geometry    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	Legs        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MapMatchResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mapMatchResultJSON) RawJSON() string {
	return r.raw
}

type MapMatchResultProperties struct {
	// Match confidence score
	Confidence float64 `json:"confidence"`
	// Total matched distance in meters
	Distance float64 `json:"distance"`
	// Estimated duration in seconds
	Duration float64                      `json:"duration"`
	JSON     mapMatchResultPropertiesJSON `json:"-"`
}

// mapMatchResultPropertiesJSON contains the JSON metadata for the struct
// [MapMatchResultProperties]
type mapMatchResultPropertiesJSON struct {
	Confidence  apijson.Field
	Distance    apijson.Field
	Duration    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MapMatchResultProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r mapMatchResultPropertiesJSON) RawJSON() string {
	return r.raw
}

type MapMatchResultType string

const (
	MapMatchResultTypeFeature MapMatchResultType = "Feature"
)

func (r MapMatchResultType) IsKnown() bool {
	switch r {
	case MapMatchResultTypeFeature:
		return true
	}
	return false
}

type MapMatchMatchParams struct {
	// GPS trace to match against the road network
	MapMatchRequest MapMatchRequestParam `json:"map_match_request" api:"required"`
}

func (r MapMatchMatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.MapMatchRequest)
}
