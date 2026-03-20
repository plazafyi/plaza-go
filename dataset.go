// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomplazafyiplazago

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
	"github.com/plazafyi/plaza-go/internal/param"
	"github.com/plazafyi/plaza-go/internal/requestconfig"
	"github.com/plazafyi/plaza-go/option"
)

// DatasetService contains methods and other services that help with interacting
// with the plaza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDatasetService] method instead.
type DatasetService struct {
	Options []option.RequestOption
}

// NewDatasetService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDatasetService(opts ...option.RequestOption) (r *DatasetService) {
	r = &DatasetService{}
	r.Options = opts
	return
}

// Create a new dataset (admin only)
func (r *DatasetService) New(ctx context.Context, body DatasetNewParams, opts ...option.RequestOption) (res *Dataset, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/datasets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get dataset by ID
func (r *DatasetService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *Dataset, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/datasets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List all datasets
func (r *DatasetService) List(ctx context.Context, opts ...option.RequestOption) (res *DatasetList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/datasets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a dataset
func (r *DatasetService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("api/v1/datasets/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Query features in a dataset
func (r *DatasetService) Features(ctx context.Context, id string, query DatasetFeaturesParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/datasets/%s/features", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Metadata for a custom dataset. Datasets contain user-uploaded geospatial
// features separate from the OSM data.
type Dataset struct {
	// Dataset UUID
	ID string `json:"id" api:"required" format:"uuid"`
	// Creation timestamp (UTC)
	InsertedAt time.Time `json:"inserted_at" api:"required" format:"date-time"`
	// Human-readable dataset name
	Name string `json:"name" api:"required"`
	// URL-friendly identifier
	Slug string `json:"slug" api:"required"`
	// Last update timestamp (UTC)
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Required attribution text
	Attribution string `json:"attribution" api:"nullable"`
	// Dataset description
	Description string `json:"description" api:"nullable"`
	// License identifier (e.g. CC-BY-4.0)
	License string `json:"license" api:"nullable"`
	// URL of the original data source
	SourceURL string      `json:"source_url" api:"nullable" format:"uri"`
	JSON      datasetJSON `json:"-"`
}

// datasetJSON contains the JSON metadata for the struct [Dataset]
type datasetJSON struct {
	ID          apijson.Field
	InsertedAt  apijson.Field
	Name        apijson.Field
	Slug        apijson.Field
	UpdatedAt   apijson.Field
	Attribution apijson.Field
	Description apijson.Field
	License     apijson.Field
	SourceURL   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Dataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetJSON) RawJSON() string {
	return r.raw
}

// List of all available datasets.
type DatasetList struct {
	// Array of dataset metadata objects
	Datasets []Dataset       `json:"datasets" api:"required"`
	JSON     datasetListJSON `json:"-"`
}

// datasetListJSON contains the JSON metadata for the struct [DatasetList]
type datasetListJSON struct {
	Datasets    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DatasetList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetListJSON) RawJSON() string {
	return r.raw
}

type DatasetNewParams struct {
	// Human-readable dataset name
	Name param.Field[string] `json:"name" api:"required"`
	// URL-friendly identifier (lowercase, hyphens, no spaces)
	Slug param.Field[string] `json:"slug" api:"required"`
	// Required attribution text
	Attribution param.Field[string] `json:"attribution"`
	// Dataset description
	Description param.Field[string] `json:"description"`
	// License identifier (e.g. CC-BY-4.0)
	License param.Field[string] `json:"license"`
	// Source data URL
	SourceURL param.Field[string] `json:"source_url" format:"uri"`
}

func (r DatasetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatasetFeaturesParams struct {
	// Cursor for pagination
	Cursor param.Field[string] `query:"cursor"`
	// Maximum results
	Limit param.Field[int64] `query:"limit"`
	// Buffer geometry by meters
	OutputBuffer param.Field[float64] `query:"output[buffer]"`
	// Replace geometry with centroid
	OutputCentroid param.Field[bool] `query:"output[centroid]"`
	// Comma-separated property fields to include
	OutputFields param.Field[string] `query:"output[fields]"`
	// Include geometry (default true)
	OutputGeometry param.Field[bool] `query:"output[geometry]"`
	// Extra computed fields: bbox, distance, center
	OutputInclude param.Field[string] `query:"output[include]"`
	// Coordinate decimal precision (1-15, default 7)
	OutputPrecision param.Field[int64] `query:"output[precision]"`
	// Simplify geometry tolerance in meters
	OutputSimplify param.Field[float64] `query:"output[simplify]"`
	// Sort by: distance, name, osm_id
	OutputSort param.Field[string] `query:"output[sort]"`
}

// URLQuery serializes [DatasetFeaturesParams]'s query parameters as `url.Values`.
func (r DatasetFeaturesParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
