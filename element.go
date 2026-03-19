// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomplazafyiplazago

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/plazafyi/plaza-go/internal/apijson"
	"github.com/plazafyi/plaza-go/internal/apiquery"
	"github.com/plazafyi/plaza-go/internal/param"
	"github.com/plazafyi/plaza-go/internal/requestconfig"
	"github.com/plazafyi/plaza-go/option"
)

// ElementService contains methods and other services that help with interacting
// with the plaza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewElementService] method instead.
type ElementService struct {
	Options []option.RequestOption
}

// NewElementService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewElementService(opts ...option.RequestOption) (r *ElementService) {
	r = &ElementService{}
	r.Options = opts
	return
}

// Get feature by type and ID
func (r *ElementService) Get(ctx context.Context, type_ string, id int64, opts ...option.RequestOption) (res *GeoJsonFeature, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/geo+json")}, opts...)
	if type_ == "" {
		err = errors.New("missing required type parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/features/%s/%v", type_, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Fetch multiple features by type and ID
func (r *ElementService) Batch(ctx context.Context, body ElementBatchParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/geo+json")}, opts...)
	path := "api/v1/features/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Find features near a geographic point
func (r *ElementService) Nearby(ctx context.Context, query ElementNearbyParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/geo+json")}, opts...)
	path := "api/v1/features/nearby"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Query features by bounding box or H3 cell
func (r *ElementService) Query(ctx context.Context, query ElementQueryParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/geo+json")}, opts...)
	path := "api/v1/features"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type BatchRequestParam struct {
	Elements param.Field[[]BatchRequestElementParam] `json:"elements" api:"required"`
}

func (r BatchRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BatchRequestElementParam struct {
	ID   param.Field[int64]                    `json:"id" api:"required"`
	Type param.Field[BatchRequestElementsType] `json:"type" api:"required"`
}

func (r BatchRequestElementParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type BatchRequestElementsType string

const (
	BatchRequestElementsTypeNode     BatchRequestElementsType = "node"
	BatchRequestElementsTypeWay      BatchRequestElementsType = "way"
	BatchRequestElementsTypeRelation BatchRequestElementsType = "relation"
)

func (r BatchRequestElementsType) IsKnown() bool {
	switch r {
	case BatchRequestElementsTypeNode, BatchRequestElementsTypeWay, BatchRequestElementsTypeRelation:
		return true
	}
	return false
}

type ElementBatchParams struct {
	BatchRequest BatchRequestParam `json:"batch_request" api:"required"`
}

func (r ElementBatchParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.BatchRequest)
}

type ElementNearbyParams struct {
	// Latitude (-90 to 90)
	Lat param.Field[float64] `query:"lat" api:"required"`
	// Longitude (-180 to 180)
	Lng param.Field[float64] `query:"lng" api:"required"`
	// Maximum results (default 20, max 100)
	Limit param.Field[int64] `query:"limit"`
	// Search radius in meters (default 500, max 10000)
	Radius param.Field[int64] `query:"radius"`
}

// URLQuery serializes [ElementNearbyParams]'s query parameters as `url.Values`.
func (r ElementNearbyParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ElementQueryParams struct {
	// Bounding box: south,west,north,east. At least one of bbox or h3 is required.
	Bbox param.Field[string] `query:"bbox"`
	// Cursor for pagination
	Cursor param.Field[string] `query:"cursor"`
	// H3 cell index. At least one of bbox or h3 is required.
	H3 param.Field[string] `query:"h3"`
	// Maximum results (default 100, max 10000)
	Limit param.Field[int64] `query:"limit"`
	// Element types (comma-separated: node,way,relation)
	Type param.Field[string] `query:"type"`
}

// URLQuery serializes [ElementQueryParams]'s query parameters as `url.Values`.
func (r ElementQueryParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
