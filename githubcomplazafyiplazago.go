// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomplazafyiplazago

import (
	"reflect"

	"github.com/plazafyi/plaza-go/internal/apijson"
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
	// GeoJSON Geometry object per RFC 7946. Coordinates use [longitude, latitude]
	// order. 3D coordinates [lng, lat, elevation] are used for elevation endpoints.
	Geometry GeoJsonGeometry `json:"geometry" api:"required"`
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

// GeoJSON Geometry object per RFC 7946. Coordinates use [longitude, latitude]
// order. 3D coordinates [lng, lat, elevation] are used for elevation endpoints.
type GeoJsonGeometry struct {
	// Coordinates array. Nesting depth varies by geometry type: Point = [lng, lat],
	// LineString = [[lng, lat], ...], Polygon = [[[lng, lat], ...], ...], etc.
	Coordinates GeoJsonGeometryCoordinatesUnion `json:"coordinates" api:"required"`
	// Geometry type
	Type GeoJsonGeometryType `json:"type" api:"required"`
	JSON geoJsonGeometryJSON `json:"-"`
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

// Coordinates array. Nesting depth varies by geometry type: Point = [lng, lat],
// LineString = [[lng, lat], ...], Polygon = [[[lng, lat], ...], ...], etc.
//
// Union satisfied by [GeoJsonGeometryCoordinatesPoint],
// [GeoJsonGeometryCoordinatesLineStringOrMultiPoint],
// [GeoJsonGeometryCoordinatesPolygonOrMultiLineString] or
// [GeoJsonGeometryCoordinatesMultiPolygon].
type GeoJsonGeometryCoordinatesUnion interface {
	implementsGeoJsonGeometryCoordinatesUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*GeoJsonGeometryCoordinatesUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GeoJsonGeometryCoordinatesPoint{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GeoJsonGeometryCoordinatesLineStringOrMultiPoint{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GeoJsonGeometryCoordinatesPolygonOrMultiLineString{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(GeoJsonGeometryCoordinatesMultiPolygon{}),
		},
	)
}

type GeoJsonGeometryCoordinatesPoint []float64

func (r GeoJsonGeometryCoordinatesPoint) implementsGeoJsonGeometryCoordinatesUnion() {}

type GeoJsonGeometryCoordinatesLineStringOrMultiPoint [][]float64

func (r GeoJsonGeometryCoordinatesLineStringOrMultiPoint) implementsGeoJsonGeometryCoordinatesUnion() {
}

type GeoJsonGeometryCoordinatesPolygonOrMultiLineString [][][]float64

func (r GeoJsonGeometryCoordinatesPolygonOrMultiLineString) implementsGeoJsonGeometryCoordinatesUnion() {
}

type GeoJsonGeometryCoordinatesMultiPolygon [][][][]float64

func (r GeoJsonGeometryCoordinatesMultiPolygon) implementsGeoJsonGeometryCoordinatesUnion() {}

// Geometry type
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
