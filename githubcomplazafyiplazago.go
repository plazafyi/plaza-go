// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomplazafyiplazago

import (
	"reflect"

	"github.com/plazafyi/plaza-go/internal/apijson"
	"github.com/plazafyi/plaza-go/internal/param"
	"github.com/tidwall/gjson"
)

// GeoJSON FeatureCollection (RFC 7946). For paginated endpoints, metadata is
// returned in HTTP response headers rather than the body:
//
// | Header          | Description                                      |
// | --------------- | ------------------------------------------------ |
// | `X-Limit`       | Requested result limit                           |
// | `X-Has-More`    | `true` if more results exist                     |
// | `X-Next-Cursor` | Opaque cursor for next page (cursor pagination)  |
// | `X-Next-Offset` | Numeric offset for next page (offset pagination) |
// | `Link`          | RFC 8288 `rel="next"` link to the next page      |
//
// Content-Type is `application/geo+json`.
type FeatureCollection struct {
	// Array of GeoJSON Feature objects
	Features []GeoJsonFeature `json:"features" api:"required"`
	// Always `FeatureCollection`
	Type FeatureCollectionType `json:"type" api:"required"`
	JSON featureCollectionJSON `json:"-"`
}

// featureCollectionJSON contains the JSON metadata for the struct
// [FeatureCollection]
type featureCollectionJSON struct {
	Features    apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FeatureCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r featureCollectionJSON) RawJSON() string {
	return r.raw
}

// Always `FeatureCollection`
type FeatureCollectionType string

const (
	FeatureCollectionTypeFeatureCollection FeatureCollectionType = "FeatureCollection"
)

func (r FeatureCollectionType) IsKnown() bool {
	switch r {
	case FeatureCollectionTypeFeatureCollection:
		return true
	}
	return false
}

// GeoJSON Feature representing an OSM element. Tags from the original OSM element
// are flattened directly into `properties` (not nested under a `tags` key).
// Metadata fields `@type` and `@id` identify the OSM element type and ID within
// properties.
type GeoJsonFeature struct {
	// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
	// determines the coordinate structure.
	Geometry Geometry `json:"geometry" api:"required"`
	// OSM tags flattened as key-value pairs, plus `@type` (node/way/relation) and
	// `@id` (OSM ID) metadata fields. May include `distance_m` for proximity queries.
	Properties map[string]interface{} `json:"properties" api:"required"`
	// Always `Feature`
	Type GeoJsonFeatureType `json:"type" api:"required"`
	// Compound identifier in `type/osm_id` format
	ID   string             `json:"id"`
	JSON geoJsonFeatureJSON `json:"-"`
}

// geoJsonFeatureJSON contains the JSON metadata for the struct [GeoJsonFeature]
type geoJsonFeatureJSON struct {
	Geometry    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeoJsonFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geoJsonFeatureJSON) RawJSON() string {
	return r.raw
}

// Always `Feature`
type GeoJsonFeatureType string

const (
	GeoJsonFeatureTypeFeature GeoJsonFeatureType = "Feature"
)

func (r GeoJsonFeatureType) IsKnown() bool {
	switch r {
	case GeoJsonFeatureTypeFeature:
		return true
	}
	return false
}

// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
// determines the coordinate structure.
type Geometry struct {
	// This field can have the runtime type of [[]float64], [[][]float64],
	// [[][][]float64], [[][][][]float64].
	Coordinates interface{}  `json:"coordinates" api:"required"`
	Type        GeometryType `json:"type" api:"required"`
	JSON        geometryJSON `json:"-"`
	union       GeometryUnion
}

// geometryJSON contains the JSON metadata for the struct [Geometry]
type geometryJSON struct {
	Coordinates apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r geometryJSON) RawJSON() string {
	return r.raw
}

func (r *Geometry) UnmarshalJSON(data []byte) (err error) {
	*r = Geometry{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [GeometryUnion] interface which you can cast to the specific
// types for more type safety.
//
// Possible runtime types of the union are [PointGeometry], [LineStringGeometry],
// [PolygonGeometry], [MultiPointGeometry], [MultiLineStringGeometry],
// [MultiPolygonGeometry].
func (r Geometry) AsUnion() GeometryUnion {
	return r.union
}

// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
// determines the coordinate structure.
//
// Union satisfied by [PointGeometry], [LineStringGeometry], [PolygonGeometry],
// [MultiPointGeometry], [MultiLineStringGeometry] or [MultiPolygonGeometry].
type GeometryUnion interface {
	implementsGeometry()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GeometryUnion)(nil)).Elem(),
		"type",
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PointGeometry{}),
			DiscriminatorValue: "Point",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(LineStringGeometry{}),
			DiscriminatorValue: "LineString",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(PolygonGeometry{}),
			DiscriminatorValue: "Polygon",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MultiPointGeometry{}),
			DiscriminatorValue: "MultiPoint",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MultiLineStringGeometry{}),
			DiscriminatorValue: "MultiLineString",
		},
		apijson.UnionVariant{
			TypeFilter:         gjson.JSON,
			Type:               reflect.TypeOf(MultiPolygonGeometry{}),
			DiscriminatorValue: "MultiPolygon",
		},
	)
}

type GeometryType string

const (
	GeometryTypePoint           GeometryType = "Point"
	GeometryTypeLineString      GeometryType = "LineString"
	GeometryTypePolygon         GeometryType = "Polygon"
	GeometryTypeMultiPoint      GeometryType = "MultiPoint"
	GeometryTypeMultiLineString GeometryType = "MultiLineString"
	GeometryTypeMultiPolygon    GeometryType = "MultiPolygon"
)

func (r GeometryType) IsKnown() bool {
	switch r {
	case GeometryTypePoint, GeometryTypeLineString, GeometryTypePolygon, GeometryTypeMultiPoint, GeometryTypeMultiLineString, GeometryTypeMultiPolygon:
		return true
	}
	return false
}

// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
// determines the coordinate structure.
type GeometryParam struct {
	Coordinates param.Field[interface{}]  `json:"coordinates" api:"required"`
	Type        param.Field[GeometryType] `json:"type" api:"required"`
}

func (r GeometryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r GeometryParam) implementsGeometryUnionParam() {}

// GeoJSON Geometry object per RFC 7946. Discriminated union — the `type` field
// determines the coordinate structure.
//
// Satisfied by [PointGeometryParam], [LineStringGeometryParam],
// [PolygonGeometryParam], [MultiPointGeometryParam],
// [MultiLineStringGeometryParam], [MultiPolygonGeometryParam], [GeometryParam].
type GeometryUnionParam interface {
	implementsGeometryUnionParam()
}

// GeoJSON LineString geometry per RFC 7946. An ordered sequence of two or more
// positions.
type LineStringGeometry struct {
	// Array of [lng, lat] or [lng, lat, alt] positions
	Coordinates [][]float64            `json:"coordinates" api:"required"`
	Type        LineStringGeometryType `json:"type" api:"required"`
	JSON        lineStringGeometryJSON `json:"-"`
}

// lineStringGeometryJSON contains the JSON metadata for the struct
// [LineStringGeometry]
type lineStringGeometryJSON struct {
	Coordinates apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LineStringGeometry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r lineStringGeometryJSON) RawJSON() string {
	return r.raw
}

func (r LineStringGeometry) implementsGeometry() {}

type LineStringGeometryType string

const (
	LineStringGeometryTypeLineString LineStringGeometryType = "LineString"
)

func (r LineStringGeometryType) IsKnown() bool {
	switch r {
	case LineStringGeometryTypeLineString:
		return true
	}
	return false
}

// GeoJSON LineString geometry per RFC 7946. An ordered sequence of two or more
// positions.
type LineStringGeometryParam struct {
	// Array of [lng, lat] or [lng, lat, alt] positions
	Coordinates param.Field[[][]float64]            `json:"coordinates" api:"required"`
	Type        param.Field[LineStringGeometryType] `json:"type" api:"required"`
}

func (r LineStringGeometryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r LineStringGeometryParam) implementsGeometryUnionParam() {}

// GeoJSON MultiLineString geometry per RFC 7946. An array of LineString coordinate
// arrays.
type MultiLineStringGeometry struct {
	// Array of LineString coordinate arrays
	Coordinates [][][]float64               `json:"coordinates" api:"required"`
	Type        MultiLineStringGeometryType `json:"type" api:"required"`
	JSON        multiLineStringGeometryJSON `json:"-"`
}

// multiLineStringGeometryJSON contains the JSON metadata for the struct
// [MultiLineStringGeometry]
type multiLineStringGeometryJSON struct {
	Coordinates apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MultiLineStringGeometry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r multiLineStringGeometryJSON) RawJSON() string {
	return r.raw
}

func (r MultiLineStringGeometry) implementsGeometry() {}

type MultiLineStringGeometryType string

const (
	MultiLineStringGeometryTypeMultiLineString MultiLineStringGeometryType = "MultiLineString"
)

func (r MultiLineStringGeometryType) IsKnown() bool {
	switch r {
	case MultiLineStringGeometryTypeMultiLineString:
		return true
	}
	return false
}

// GeoJSON MultiLineString geometry per RFC 7946. An array of LineString coordinate
// arrays.
type MultiLineStringGeometryParam struct {
	// Array of LineString coordinate arrays
	Coordinates param.Field[[][][]float64]               `json:"coordinates" api:"required"`
	Type        param.Field[MultiLineStringGeometryType] `json:"type" api:"required"`
}

func (r MultiLineStringGeometryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MultiLineStringGeometryParam) implementsGeometryUnionParam() {}

// GeoJSON MultiPoint geometry per RFC 7946. An array of positions.
type MultiPointGeometry struct {
	// Array of [lng, lat] or [lng, lat, alt] positions
	Coordinates [][]float64            `json:"coordinates" api:"required"`
	Type        MultiPointGeometryType `json:"type" api:"required"`
	JSON        multiPointGeometryJSON `json:"-"`
}

// multiPointGeometryJSON contains the JSON metadata for the struct
// [MultiPointGeometry]
type multiPointGeometryJSON struct {
	Coordinates apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MultiPointGeometry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r multiPointGeometryJSON) RawJSON() string {
	return r.raw
}

func (r MultiPointGeometry) implementsGeometry() {}

type MultiPointGeometryType string

const (
	MultiPointGeometryTypeMultiPoint MultiPointGeometryType = "MultiPoint"
)

func (r MultiPointGeometryType) IsKnown() bool {
	switch r {
	case MultiPointGeometryTypeMultiPoint:
		return true
	}
	return false
}

// GeoJSON MultiPoint geometry per RFC 7946. An array of positions.
type MultiPointGeometryParam struct {
	// Array of [lng, lat] or [lng, lat, alt] positions
	Coordinates param.Field[[][]float64]            `json:"coordinates" api:"required"`
	Type        param.Field[MultiPointGeometryType] `json:"type" api:"required"`
}

func (r MultiPointGeometryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MultiPointGeometryParam) implementsGeometryUnionParam() {}

func (r MultiPointGeometryParam) implementsElevationLookupRequestGeometryUnionParam() {}

// GeoJSON MultiPolygon geometry per RFC 7946. An array of Polygon coordinate
// arrays.
type MultiPolygonGeometry struct {
	// Array of Polygon coordinate arrays
	Coordinates [][][][]float64          `json:"coordinates" api:"required"`
	Type        MultiPolygonGeometryType `json:"type" api:"required"`
	JSON        multiPolygonGeometryJSON `json:"-"`
}

// multiPolygonGeometryJSON contains the JSON metadata for the struct
// [MultiPolygonGeometry]
type multiPolygonGeometryJSON struct {
	Coordinates apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MultiPolygonGeometry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r multiPolygonGeometryJSON) RawJSON() string {
	return r.raw
}

func (r MultiPolygonGeometry) implementsGeometry() {}

type MultiPolygonGeometryType string

const (
	MultiPolygonGeometryTypeMultiPolygon MultiPolygonGeometryType = "MultiPolygon"
)

func (r MultiPolygonGeometryType) IsKnown() bool {
	switch r {
	case MultiPolygonGeometryTypeMultiPolygon:
		return true
	}
	return false
}

// GeoJSON MultiPolygon geometry per RFC 7946. An array of Polygon coordinate
// arrays.
type MultiPolygonGeometryParam struct {
	// Array of Polygon coordinate arrays
	Coordinates param.Field[[][][][]float64]          `json:"coordinates" api:"required"`
	Type        param.Field[MultiPolygonGeometryType] `json:"type" api:"required"`
}

func (r MultiPolygonGeometryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r MultiPolygonGeometryParam) implementsGeometryUnionParam() {}

// GeoJSON Point geometry per RFC 7946. Coordinates use [longitude, latitude]
// order. Optional third element is altitude in meters.
type PointGeometry struct {
	// [longitude, latitude] or [longitude, latitude, altitude]
	Coordinates []float64         `json:"coordinates" api:"required"`
	Type        PointGeometryType `json:"type" api:"required"`
	JSON        pointGeometryJSON `json:"-"`
}

// pointGeometryJSON contains the JSON metadata for the struct [PointGeometry]
type pointGeometryJSON struct {
	Coordinates apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PointGeometry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pointGeometryJSON) RawJSON() string {
	return r.raw
}

func (r PointGeometry) implementsGeometry() {}

type PointGeometryType string

const (
	PointGeometryTypePoint PointGeometryType = "Point"
)

func (r PointGeometryType) IsKnown() bool {
	switch r {
	case PointGeometryTypePoint:
		return true
	}
	return false
}

// GeoJSON Point geometry per RFC 7946. Coordinates use [longitude, latitude]
// order. Optional third element is altitude in meters.
type PointGeometryParam struct {
	// [longitude, latitude] or [longitude, latitude, altitude]
	Coordinates param.Field[[]float64]         `json:"coordinates" api:"required"`
	Type        param.Field[PointGeometryType] `json:"type" api:"required"`
}

func (r PointGeometryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PointGeometryParam) implementsGeometryUnionParam() {}

func (r PointGeometryParam) implementsElevationLookupRequestGeometryUnionParam() {}

// GeoJSON Polygon geometry per RFC 7946. An array of linear rings where the first
// ring is the exterior boundary and subsequent rings are holes. Each ring must
// have at least 4 positions with the first and last being identical.
type PolygonGeometry struct {
	// Array of linear rings (first = exterior, rest = holes)
	Coordinates [][][]float64       `json:"coordinates" api:"required"`
	Type        PolygonGeometryType `json:"type" api:"required"`
	JSON        polygonGeometryJSON `json:"-"`
}

// polygonGeometryJSON contains the JSON metadata for the struct [PolygonGeometry]
type polygonGeometryJSON struct {
	Coordinates apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PolygonGeometry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r polygonGeometryJSON) RawJSON() string {
	return r.raw
}

func (r PolygonGeometry) implementsGeometry() {}

type PolygonGeometryType string

const (
	PolygonGeometryTypePolygon PolygonGeometryType = "Polygon"
)

func (r PolygonGeometryType) IsKnown() bool {
	switch r {
	case PolygonGeometryTypePolygon:
		return true
	}
	return false
}

// GeoJSON Polygon geometry per RFC 7946. An array of linear rings where the first
// ring is the exterior boundary and subsequent rings are holes. Each ring must
// have at least 4 positions with the first and last being identical.
type PolygonGeometryParam struct {
	// Array of linear rings (first = exterior, rest = holes)
	Coordinates param.Field[[][][]float64]       `json:"coordinates" api:"required"`
	Type        param.Field[PolygonGeometryType] `json:"type" api:"required"`
}

func (r PolygonGeometryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

func (r PolygonGeometryParam) implementsGeometryUnionParam() {}
