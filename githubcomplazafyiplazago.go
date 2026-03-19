// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomplazafyiplazago

import (
	"reflect"

	"github.com/plazafyi/plaza-go/internal/apijson"
	"github.com/plazafyi/plaza-go/internal/param"
	"github.com/tidwall/gjson"
)

// Bare GeoJSON FeatureCollection. Pagination metadata is returned in HTTP headers
// (X-Limit, X-Has-More, X-Next-Cursor, X-Next-Offset, Link).
type FeatureCollection struct {
	Features []GeoJsonFeature      `json:"features" api:"required"`
	Type     FeatureCollectionType `json:"type" api:"required"`
	JSON     featureCollectionJSON `json:"-"`
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

type GeoJsonFeature struct {
	Geometry   GeoJsonGeometry        `json:"geometry" api:"required"`
	Properties map[string]interface{} `json:"properties" api:"required"`
	Type       GeoJsonFeatureType     `json:"type" api:"required"`
	// Feature identifier (type/osm_id)
	ID string `json:"id"`
	// OpenStreetMap ID
	OsmID int64              `json:"osm_id"`
	JSON  geoJsonFeatureJSON `json:"-"`
}

// geoJsonFeatureJSON contains the JSON metadata for the struct [GeoJsonFeature]
type geoJsonFeatureJSON struct {
	Geometry    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	ID          apijson.Field
	OsmID       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeoJsonFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geoJsonFeatureJSON) RawJSON() string {
	return r.raw
}

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

type GeoJsonGeometry struct {
	// GeoJSON coordinates array (nesting depth varies by geometry type)
	Coordinates GeoJsonGeometryCoordinatesUnion `json:"coordinates" api:"required"`
	Type        GeoJsonGeometryType             `json:"type" api:"required"`
	JSON        geoJsonGeometryJSON             `json:"-"`
}

// geoJsonGeometryJSON contains the JSON metadata for the struct [GeoJsonGeometry]
type geoJsonGeometryJSON struct {
	Coordinates apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *GeoJsonGeometry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r geoJsonGeometryJSON) RawJSON() string {
	return r.raw
}

// GeoJSON coordinates array (nesting depth varies by geometry type)
//
// Union satisfied by [GeoJsonGeometryCoordinatesArray],
// [GeoJsonGeometryCoordinatesArray], [GeoJsonGeometryCoordinatesArray] or
// [GeoJsonGeometryCoordinatesArray].
type GeoJsonGeometryCoordinatesUnion interface {
	implementsGeoJsonGeometryCoordinatesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GeoJsonGeometryCoordinatesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GeoJsonGeometryCoordinatesArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GeoJsonGeometryCoordinatesArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GeoJsonGeometryCoordinatesArray{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GeoJsonGeometryCoordinatesArray{}),
		},
	)
}

type GeoJsonGeometryCoordinatesArray []float64

func (r GeoJsonGeometryCoordinatesArray) implementsGeoJsonGeometryCoordinatesUnion() {}

type GeoJsonGeometryType string

const (
	GeoJsonGeometryTypePoint           GeoJsonGeometryType = "Point"
	GeoJsonGeometryTypeLineString      GeoJsonGeometryType = "LineString"
	GeoJsonGeometryTypePolygon         GeoJsonGeometryType = "Polygon"
	GeoJsonGeometryTypeMultiPoint      GeoJsonGeometryType = "MultiPoint"
	GeoJsonGeometryTypeMultiLineString GeoJsonGeometryType = "MultiLineString"
	GeoJsonGeometryTypeMultiPolygon    GeoJsonGeometryType = "MultiPolygon"
)

func (r GeoJsonGeometryType) IsKnown() bool {
	switch r {
	case GeoJsonGeometryTypePoint, GeoJsonGeometryTypeLineString, GeoJsonGeometryTypePolygon, GeoJsonGeometryTypeMultiPoint, GeoJsonGeometryTypeMultiLineString, GeoJsonGeometryTypeMultiPolygon:
		return true
	}
	return false
}

type GeoJsonGeometryParam struct {
	// GeoJSON coordinates array (nesting depth varies by geometry type)
	Coordinates param.Field[GeoJsonGeometryCoordinatesUnionParam] `json:"coordinates" api:"required"`
	Type        param.Field[GeoJsonGeometryType]                  `json:"type" api:"required"`
}

func (r GeoJsonGeometryParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// GeoJSON coordinates array (nesting depth varies by geometry type)
//
// Satisfied by [GeoJsonGeometryCoordinatesArrayParam],
// [GeoJsonGeometryCoordinatesArrayParam], [GeoJsonGeometryCoordinatesArrayParam],
// [GeoJsonGeometryCoordinatesArrayParam].
type GeoJsonGeometryCoordinatesUnionParam interface {
	implementsGeoJsonGeometryCoordinatesUnionParam()
}

type GeoJsonGeometryCoordinatesArrayParam []float64

func (r GeoJsonGeometryCoordinatesArrayParam) implementsGeoJsonGeometryCoordinatesUnionParam() {}
