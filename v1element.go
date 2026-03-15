// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plaza

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/plaza-go/internal/apijson"
	"github.com/stainless-sdks/plaza-go/internal/apiquery"
	"github.com/stainless-sdks/plaza-go/internal/requestconfig"
	"github.com/stainless-sdks/plaza-go/option"
	"github.com/stainless-sdks/plaza-go/packages/param"
	"github.com/stainless-sdks/plaza-go/packages/respjson"
)

// V1ElementService contains methods and other services that help with interacting
// with the plaza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ElementService] method instead.
type V1ElementService struct {
	options []option.RequestOption
}

// NewV1ElementService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1ElementService(opts ...option.RequestOption) (r V1ElementService) {
	r = V1ElementService{}
	r.options = opts
	return
}

// Get element by type and ID
func (r *V1ElementService) Get(ctx context.Context, id int64, query V1ElementGetParams, opts ...option.RequestOption) (res *GeoJsonFeature, err error) {
	opts = slices.Concat(r.options, opts)
	if query.Type == "" {
		err = errors.New("missing required type parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/elements/%s/%v", url.PathEscape(query.Type), id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Fetch multiple elements by type and ID
func (r *V1ElementService) FetchBatch(ctx context.Context, body V1ElementFetchBatchParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/elements/batch"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Query elements by bounding box or H3 cell
func (r *V1ElementService) Query(ctx context.Context, query V1ElementQueryParams, opts ...option.RequestOption) (res *FeatureCollection, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/elements"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type GeoJsonFeature struct {
	Geometry   GeoJsonGeometry `json:"geometry" api:"required"`
	Properties map[string]any  `json:"properties" api:"required"`
	// Any of "Feature".
	Type GeoJsonFeatureType `json:"type" api:"required"`
	// Feature identifier (type/osm_id)
	ID string `json:"id"`
	// OpenStreetMap ID
	OsmID int64 `json:"osm_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Geometry    respjson.Field
		Properties  respjson.Field
		Type        respjson.Field
		ID          respjson.Field
		OsmID       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeoJsonFeature) RawJSON() string { return r.JSON.raw }
func (r *GeoJsonFeature) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeoJsonFeatureType string

const (
	GeoJsonFeatureTypeFeature GeoJsonFeatureType = "Feature"
)

type GeoJsonGeometry struct {
	// GeoJSON coordinates array (nesting depth varies by geometry type)
	Coordinates GeoJsonGeometryCoordinatesUnion `json:"coordinates" api:"required"`
	// Any of "Point", "LineString", "Polygon", "MultiPoint", "MultiLineString",
	// "MultiPolygon".
	Type GeoJsonGeometryType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Coordinates respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GeoJsonGeometry) RawJSON() string { return r.JSON.raw }
func (r *GeoJsonGeometry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// GeoJsonGeometryCoordinatesUnion contains all possible properties and values from
// [[]float64], [[][]float64], [[][][]float64], [[][][][]float64].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloatArray OfArrayOfFloatArray OfArrayOfArrayOfFloatArray
// OfArrayOfArrayOfArrayOfFloatArray]
type GeoJsonGeometryCoordinatesUnion struct {
	// This field will be present if the value is a [[]float64] instead of an object.
	OfFloatArray []float64 `json:",inline"`
	// This field will be present if the value is a [[][]float64] instead of an object.
	OfArrayOfFloatArray [][]float64 `json:",inline"`
	// This field will be present if the value is a [[][][]float64] instead of an
	// object.
	OfArrayOfArrayOfFloatArray [][][]float64 `json:",inline"`
	// This field will be present if the value is a [[][][][]float64] instead of an
	// object.
	OfArrayOfArrayOfArrayOfFloatArray [][][][]float64 `json:",inline"`
	JSON                              struct {
		OfFloatArray                      respjson.Field
		OfArrayOfFloatArray               respjson.Field
		OfArrayOfArrayOfFloatArray        respjson.Field
		OfArrayOfArrayOfArrayOfFloatArray respjson.Field
		raw                               string
	} `json:"-"`
}

func (u GeoJsonGeometryCoordinatesUnion) AsFloatArray() (v []float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GeoJsonGeometryCoordinatesUnion) AsArrayOfFloatArray() (v [][]float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GeoJsonGeometryCoordinatesUnion) AsArrayOfArrayOfFloatArray() (v [][][]float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u GeoJsonGeometryCoordinatesUnion) AsArrayOfArrayOfArrayOfFloatArray() (v [][][][]float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u GeoJsonGeometryCoordinatesUnion) RawJSON() string { return u.JSON.raw }

func (r *GeoJsonGeometryCoordinatesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GeoJsonGeometryType string

const (
	GeoJsonGeometryTypePoint           GeoJsonGeometryType = "Point"
	GeoJsonGeometryTypeLineString      GeoJsonGeometryType = "LineString"
	GeoJsonGeometryTypePolygon         GeoJsonGeometryType = "Polygon"
	GeoJsonGeometryTypeMultiPoint      GeoJsonGeometryType = "MultiPoint"
	GeoJsonGeometryTypeMultiLineString GeoJsonGeometryType = "MultiLineString"
	GeoJsonGeometryTypeMultiPolygon    GeoJsonGeometryType = "MultiPolygon"
)

type V1ElementGetParams struct {
	Type string `path:"type" api:"required" json:"-"`
	paramObj
}

type V1ElementFetchBatchParams struct {
	Elements []V1ElementFetchBatchParamsElement `json:"elements,omitzero" api:"required"`
	paramObj
}

func (r V1ElementFetchBatchParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ElementFetchBatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ElementFetchBatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties ID, Type are required.
type V1ElementFetchBatchParamsElement struct {
	ID int64 `json:"id" api:"required"`
	// Any of "node", "way", "relation".
	Type string `json:"type,omitzero" api:"required"`
	paramObj
}

func (r V1ElementFetchBatchParamsElement) MarshalJSON() (data []byte, err error) {
	type shadow V1ElementFetchBatchParamsElement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ElementFetchBatchParamsElement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[V1ElementFetchBatchParamsElement](
		"type", "node", "way", "relation",
	)
}

type V1ElementQueryParams struct {
	// Bounding box: south,west,north,east. At least one of bbox or h3 is required.
	Bbox param.Opt[string] `query:"bbox,omitzero" json:"-"`
	// Cursor for pagination
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// H3 cell index. At least one of bbox or h3 is required.
	H3 param.Opt[string] `query:"h3,omitzero" json:"-"`
	// Maximum results (default 100, max 10000)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Element types (comma-separated: node,way,relation)
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1ElementQueryParams]'s query parameters as `url.Values`.
func (r V1ElementQueryParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
