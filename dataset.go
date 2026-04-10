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

// Create a new dataset
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

// List datasets
func (r *DatasetService) List(ctx context.Context, query DatasetListParams, opts ...option.RequestOption) (res *DatasetList, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/datasets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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

// Metadata for a custom dataset. Datasets contain user-uploaded geospatial
// features separate from the OSM data.
type Dataset struct {
	// Dataset UUID
	ID string `json:"id" api:"required" format:"uuid"`
	// Creation timestamp (UTC)
	InsertedAt time.Time `json:"inserted_at" api:"required" format:"date-time"`
	// Human-readable dataset name
	Name string `json:"name" api:"required"`
	// Dataset scope: plaza (managed by Plaza) or user (user-owned)
	Scope DatasetScope `json:"scope" api:"required"`
	// URL-friendly identifier
	Slug string `json:"slug" api:"required"`
	// Current processing status
	Status DatasetStatus `json:"status" api:"required"`
	// Last update timestamp (UTC)
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Number of addresses in this dataset
	AddressCount int64 `json:"address_count"`
	// Required attribution text
	Attribution string `json:"attribution" api:"nullable"`
	// Dataset description
	Description string `json:"description" api:"nullable"`
	// Number of routing edges in this dataset
	EdgeCount int64 `json:"edge_count"`
	// Error message if status is 'error'
	ErrorMessage string `json:"error_message" api:"nullable"`
	// Number of features in this dataset
	FeatureCount int64 `json:"feature_count"`
	// License identifier (e.g. CC-BY-4.0)
	License string `json:"license" api:"nullable"`
	// Detected or user-defined property schema
	SchemaDefinition interface{} `json:"schema_definition" api:"nullable"`
	// Data format (geojson)
	SourceFormat string `json:"source_format" api:"nullable"`
	// URL of the original data source
	SourceURL string `json:"source_url" api:"nullable" format:"uri"`
	// Total storage consumed in bytes
	StorageBytes int64 `json:"storage_bytes"`
	// Whether strict schema validation is enabled
	StrictMode bool        `json:"strict_mode"`
	JSON       datasetJSON `json:"-"`
}

// datasetJSON contains the JSON metadata for the struct [Dataset]
type datasetJSON struct {
	ID               apijson.Field
	InsertedAt       apijson.Field
	Name             apijson.Field
	Scope            apijson.Field
	Slug             apijson.Field
	Status           apijson.Field
	UpdatedAt        apijson.Field
	AddressCount     apijson.Field
	Attribution      apijson.Field
	Description      apijson.Field
	EdgeCount        apijson.Field
	ErrorMessage     apijson.Field
	FeatureCount     apijson.Field
	License          apijson.Field
	SchemaDefinition apijson.Field
	SourceFormat     apijson.Field
	SourceURL        apijson.Field
	StorageBytes     apijson.Field
	StrictMode       apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *Dataset) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r datasetJSON) RawJSON() string {
	return r.raw
}

// Dataset scope: plaza (managed by Plaza) or user (user-owned)
type DatasetScope string

const (
	DatasetScopePlaza DatasetScope = "plaza"
	DatasetScopeUser  DatasetScope = "user"
)

func (r DatasetScope) IsKnown() bool {
	switch r {
	case DatasetScopePlaza, DatasetScopeUser:
		return true
	}
	return false
}

// Current processing status
type DatasetStatus string

const (
	DatasetStatusPending    DatasetStatus = "pending"
	DatasetStatusProcessing DatasetStatus = "processing"
	DatasetStatusReady      DatasetStatus = "ready"
	DatasetStatusError      DatasetStatus = "error"
)

func (r DatasetStatus) IsKnown() bool {
	switch r {
	case DatasetStatusPending, DatasetStatusProcessing, DatasetStatusReady, DatasetStatusError:
		return true
	}
	return false
}

// List of datasets visible to the authenticated user.
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
	// Enable strict schema validation (default true)
	StrictMode param.Field[bool] `json:"strict_mode"`
}

func (r DatasetNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DatasetListParams struct {
	// Filter by scope: plaza, user. Default shows user's own + plaza datasets.
	Scope param.Field[string] `query:"scope"`
}

// URLQuery serializes [DatasetListParams]'s query parameters as `url.Values`.
func (r DatasetListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
