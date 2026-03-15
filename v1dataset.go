// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plaza

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/plazafyi/plaza-go/internal/apijson"
	"github.com/plazafyi/plaza-go/internal/apiquery"
	"github.com/plazafyi/plaza-go/internal/requestconfig"
	"github.com/plazafyi/plaza-go/option"
	"github.com/plazafyi/plaza-go/packages/param"
	"github.com/plazafyi/plaza-go/packages/respjson"
)

// V1DatasetService contains methods and other services that help with interacting
// with the plaza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1DatasetService] method instead.
type V1DatasetService struct {
	options []option.RequestOption
}

// NewV1DatasetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1DatasetService(opts ...option.RequestOption) (r V1DatasetService) {
	r = V1DatasetService{}
	r.options = opts
	return
}

// Create a new dataset (admin only)
func (r *V1DatasetService) New(ctx context.Context, body V1DatasetNewParams, opts ...option.RequestOption) (res *DatasetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/datasets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get dataset by ID
func (r *V1DatasetService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *DatasetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/datasets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all datasets
func (r *V1DatasetService) List(ctx context.Context, opts ...option.RequestOption) (res *V1DatasetListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/datasets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a dataset
func (r *V1DatasetService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/datasets/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Query features in a dataset
func (r *V1DatasetService) QueryFeatures(ctx context.Context, id string, query V1DatasetQueryFeaturesParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/datasets/%s/features", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type DatasetResponse struct {
	// Dataset ID
	ID string `json:"id" api:"required" format:"uuid"`
	// Creation timestamp
	InsertedAt time.Time `json:"inserted_at" api:"required" format:"date-time"`
	// Dataset name
	Name string `json:"name" api:"required"`
	// URL-friendly slug
	Slug string `json:"slug" api:"required"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Attribution text
	Attribution string `json:"attribution" api:"nullable"`
	// Dataset description
	Description string `json:"description" api:"nullable"`
	// License identifier
	License string `json:"license" api:"nullable"`
	// Source data URL
	SourceURL string `json:"source_url" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		InsertedAt  respjson.Field
		Name        respjson.Field
		Slug        respjson.Field
		UpdatedAt   respjson.Field
		Attribution respjson.Field
		Description respjson.Field
		License     respjson.Field
		SourceURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DatasetResponse) RawJSON() string { return r.JSON.raw }
func (r *DatasetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureCollection struct {
	Features []GeoJsonFeature `json:"features" api:"required"`
	// Any of "FeatureCollection".
	Type       FeatureCollectionType       `json:"type" api:"required"`
	Pagination FeatureCollectionPagination `json:"pagination"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Features    respjson.Field
		Type        respjson.Field
		Pagination  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeatureCollection) RawJSON() string { return r.JSON.raw }
func (r *FeatureCollection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureCollectionType string

const (
	FeatureCollectionTypeFeatureCollection FeatureCollectionType = "FeatureCollection"
)

type FeatureCollectionPagination struct {
	// Whether more results exist
	HasMore bool `json:"has_more"`
	// Requested result limit
	Limit int64 `json:"limit"`
	// Cursor for next page
	NextCursor string `json:"next_cursor" api:"nullable"`
	// Offset for next page
	NextOffset int64 `json:"next_offset" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		HasMore     respjson.Field
		Limit       respjson.Field
		NextCursor  respjson.Field
		NextOffset  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeatureCollectionPagination) RawJSON() string { return r.JSON.raw }
func (r *FeatureCollectionPagination) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DatasetListResponse struct {
	Datasets []DatasetResponse `json:"datasets" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Datasets    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1DatasetListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1DatasetListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DatasetNewParams struct {
	// Dataset name
	Name string `json:"name" api:"required"`
	// URL-friendly slug
	Slug string `json:"slug" api:"required"`
	// Attribution text
	Attribution param.Opt[string] `json:"attribution,omitzero"`
	// Dataset description
	Description param.Opt[string] `json:"description,omitzero"`
	// License identifier
	License param.Opt[string] `json:"license,omitzero"`
	// Source data URL
	SourceURL param.Opt[string] `json:"source_url,omitzero"`
	paramObj
}

func (r V1DatasetNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1DatasetNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1DatasetNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1DatasetQueryFeaturesParams struct {
	// Cursor for pagination
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum results
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1DatasetQueryFeaturesParams]'s query parameters as
// `url.Values`.
func (r V1DatasetQueryFeaturesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
