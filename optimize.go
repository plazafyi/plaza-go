// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomplazafyiplazago

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/plazafyi/plaza-go/internal/apijson"
	"github.com/plazafyi/plaza-go/internal/apiquery"
	"github.com/plazafyi/plaza-go/internal/param"
	"github.com/plazafyi/plaza-go/internal/requestconfig"
	"github.com/plazafyi/plaza-go/option"
	"github.com/tidwall/gjson"
)

// OptimizeService contains methods and other services that help with interacting
// with the plaza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOptimizeService] method instead.
type OptimizeService struct {
	Options []option.RequestOption
}

// NewOptimizeService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOptimizeService(opts ...option.RequestOption) (r *OptimizeService) {
	r = &OptimizeService{}
	r.Options = opts
	return
}

// Optimize route through waypoints
func (r *OptimizeService) New(ctx context.Context, params OptimizeNewParams, opts ...option.RequestOption) (res *OptimizeResult, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "api/v1/optimize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get async optimization result
func (r *OptimizeService) Get(ctx context.Context, jobID string, opts ...option.RequestOption) (res *OptimizeJobStatus, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/optimize/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Completed optimization result as a GeoJSON FeatureCollection. Each Feature is a
// waypoint in optimized visit order. Top-level fields provide summary statistics.
type OptimizeCompletedResult struct {
	// Waypoints in optimized visit order
	Features []OptimizeCompletedResultFeature `json:"features" api:"required"`
	// Optimization method used (e.g. `nearest_neighbor`, `2opt`)
	Optimization string `json:"optimization" api:"required"`
	// Whether the route returns to the starting waypoint
	Roundtrip bool `json:"roundtrip" api:"required"`
	// Total travel time for the optimized route in seconds
	TotalCostS float64                     `json:"total_cost_s" api:"required"`
	Type       OptimizeCompletedResultType `json:"type" api:"required"`
	JSON       optimizeCompletedResultJSON `json:"-"`
}

// optimizeCompletedResultJSON contains the JSON metadata for the struct
// [OptimizeCompletedResult]
type optimizeCompletedResultJSON struct {
	Features     apijson.Field
	Optimization apijson.Field
	Roundtrip    apijson.Field
	TotalCostS   apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *OptimizeCompletedResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r optimizeCompletedResultJSON) RawJSON() string {
	return r.raw
}

func (r OptimizeCompletedResult) implementsOptimizeResult() {}

// GeoJSON Point Feature representing an optimized waypoint with cost data.
type OptimizeCompletedResultFeature struct {
	// GeoJSON Geometry object per RFC 7946. Coordinates use [longitude, latitude]
	// order. 3D coordinates [lng, lat, elevation] are used for elevation endpoints.
	Geometry   GeoJsonGeometry                           `json:"geometry" api:"required"`
	Properties OptimizeCompletedResultFeaturesProperties `json:"properties" api:"required"`
	Type       OptimizeCompletedResultFeaturesType       `json:"type" api:"required"`
	JSON       optimizeCompletedResultFeatureJSON        `json:"-"`
}

// optimizeCompletedResultFeatureJSON contains the JSON metadata for the struct
// [OptimizeCompletedResultFeature]
type optimizeCompletedResultFeatureJSON struct {
	Geometry    apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OptimizeCompletedResultFeature) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r optimizeCompletedResultFeatureJSON) RawJSON() string {
	return r.raw
}

type OptimizeCompletedResultFeaturesProperties struct {
	// Travel time in seconds from the previous waypoint to this one (0 for the first
	// waypoint)
	CostS float64 `json:"cost_s" api:"required"`
	// Cumulative travel time in seconds from the start to this waypoint
	CumulativeCostS float64 `json:"cumulative_cost_s" api:"required"`
	// Position of this waypoint in the optimized visit order (0-based)
	WaypointIndex int64                                         `json:"waypoint_index" api:"required"`
	JSON          optimizeCompletedResultFeaturesPropertiesJSON `json:"-"`
}

// optimizeCompletedResultFeaturesPropertiesJSON contains the JSON metadata for the
// struct [OptimizeCompletedResultFeaturesProperties]
type optimizeCompletedResultFeaturesPropertiesJSON struct {
	CostS           apijson.Field
	CumulativeCostS apijson.Field
	WaypointIndex   apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *OptimizeCompletedResultFeaturesProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r optimizeCompletedResultFeaturesPropertiesJSON) RawJSON() string {
	return r.raw
}

type OptimizeCompletedResultFeaturesType string

const (
	OptimizeCompletedResultFeaturesTypeFeature OptimizeCompletedResultFeaturesType = "Feature"
)

func (r OptimizeCompletedResultFeaturesType) IsKnown() bool {
	switch r {
	case OptimizeCompletedResultFeaturesTypeFeature:
		return true
	}
	return false
}

type OptimizeCompletedResultType string

const (
	OptimizeCompletedResultTypeFeatureCollection OptimizeCompletedResultType = "FeatureCollection"
)

func (r OptimizeCompletedResultType) IsKnown() bool {
	switch r {
	case OptimizeCompletedResultTypeFeatureCollection:
		return true
	}
	return false
}

// Status of an async optimization job. When `completed`, the `result` field
// contains the full OptimizeCompletedResult. When `processing`, the job is still
// running — poll again. Failed jobs return a standard Error response (HTTP 422),
// not this schema.
type OptimizeJobStatus struct {
	// Current job state
	Status OptimizeJobStatusStatus `json:"status" api:"required"`
	// Completed optimization result as a GeoJSON FeatureCollection. Each Feature is a
	// waypoint in optimized visit order. Top-level fields provide summary statistics.
	Result OptimizeCompletedResult `json:"result" api:"nullable"`
	JSON   optimizeJobStatusJSON   `json:"-"`
}

// optimizeJobStatusJSON contains the JSON metadata for the struct
// [OptimizeJobStatus]
type optimizeJobStatusJSON struct {
	Status      apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OptimizeJobStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r optimizeJobStatusJSON) RawJSON() string {
	return r.raw
}

// Current job state
type OptimizeJobStatusStatus string

const (
	OptimizeJobStatusStatusCompleted  OptimizeJobStatusStatus = "completed"
	OptimizeJobStatusStatusProcessing OptimizeJobStatusStatus = "processing"
)

func (r OptimizeJobStatusStatus) IsKnown() bool {
	switch r {
	case OptimizeJobStatusStatusCompleted, OptimizeJobStatusStatusProcessing:
		return true
	}
	return false
}

// Async optimization in progress. Poll `GET /api/v1/optimize/{job_id}` until the
// status changes to `completed` or `failed`.
type OptimizeProcessingResult struct {
	// Job ID for polling the result
	JobID string `json:"job_id" api:"required"`
	// Always `processing`
	Status OptimizeProcessingResultStatus `json:"status" api:"required"`
	JSON   optimizeProcessingResultJSON   `json:"-"`
}

// optimizeProcessingResultJSON contains the JSON metadata for the struct
// [OptimizeProcessingResult]
type optimizeProcessingResultJSON struct {
	JobID       apijson.Field
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OptimizeProcessingResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r optimizeProcessingResultJSON) RawJSON() string {
	return r.raw
}

func (r OptimizeProcessingResult) implementsOptimizeResult() {}

// Always `processing`
type OptimizeProcessingResultStatus string

const (
	OptimizeProcessingResultStatusProcessing OptimizeProcessingResultStatus = "processing"
)

func (r OptimizeProcessingResultStatus) IsKnown() bool {
	switch r {
	case OptimizeProcessingResultStatusProcessing:
		return true
	}
	return false
}

// Route optimization (Travelling Salesman) request. Finds the most efficient order
// to visit a set of waypoints. Minimum 2 waypoints, maximum 50. For large inputs,
// the request may be processed asynchronously.
type OptimizeRequestParam struct {
	// Waypoints to visit in optimized order (2-50 points)
	Waypoints param.Field[[]OptimizeRequestWaypointParam] `json:"waypoints" api:"required"`
	// Travel mode (default: `auto`)
	Mode param.Field[OptimizeRequestMode] `json:"mode"`
	// Whether the route should return to the starting waypoint (default: true)
	Roundtrip param.Field[bool] `json:"roundtrip"`
}

func (r OptimizeRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Geographic coordinate as a JSON object with `lat` and `lng` fields.
type OptimizeRequestWaypointParam struct {
	// Latitude in decimal degrees (-90 to 90)
	Lat param.Field[float64] `json:"lat" api:"required"`
	// Longitude in decimal degrees (-180 to 180)
	Lng param.Field[float64] `json:"lng" api:"required"`
}

func (r OptimizeRequestWaypointParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Travel mode (default: `auto`)
type OptimizeRequestMode string

const (
	OptimizeRequestModeAuto    OptimizeRequestMode = "auto"
	OptimizeRequestModeFoot    OptimizeRequestMode = "foot"
	OptimizeRequestModeBicycle OptimizeRequestMode = "bicycle"
)

func (r OptimizeRequestMode) IsKnown() bool {
	switch r {
	case OptimizeRequestModeAuto, OptimizeRequestModeFoot, OptimizeRequestModeBicycle:
		return true
	}
	return false
}

// Optimization response — either a completed FeatureCollection with the optimized
// route, or an async job reference to poll.
type OptimizeResult struct {
	// This field can have the runtime type of [[]OptimizeCompletedResultFeature].
	Features interface{} `json:"features"`
	// Job ID for polling the result
	JobID string `json:"job_id"`
	// Optimization method used (e.g. `nearest_neighbor`, `2opt`)
	Optimization string `json:"optimization"`
	// Whether the route returns to the starting waypoint
	Roundtrip bool `json:"roundtrip"`
	// Always `processing`
	Status OptimizeResultStatus `json:"status"`
	// Total travel time for the optimized route in seconds
	TotalCostS float64            `json:"total_cost_s"`
	Type       OptimizeResultType `json:"type"`
	JSON       optimizeResultJSON `json:"-"`
	union      OptimizeResultUnion
}

// optimizeResultJSON contains the JSON metadata for the struct [OptimizeResult]
type optimizeResultJSON struct {
	Features     apijson.Field
	JobID        apijson.Field
	Optimization apijson.Field
	Roundtrip    apijson.Field
	Status       apijson.Field
	TotalCostS   apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r optimizeResultJSON) RawJSON() string {
	return r.raw
}

func (r *OptimizeResult) UnmarshalJSON(data []byte) (err error) {
	*r = OptimizeResult{}
	err = apijson.UnmarshalRoot(data, &r.union)
	if err != nil {
		return err
	}
	return apijson.Port(r.union, &r)
}

// AsUnion returns a [OptimizeResultUnion] interface which you can cast to the
// specific types for more type safety.
//
// Possible runtime types of the union are [OptimizeCompletedResult],
// [OptimizeProcessingResult].
func (r OptimizeResult) AsUnion() OptimizeResultUnion {
	return r.union
}

// Optimization response — either a completed FeatureCollection with the optimized
// route, or an async job reference to poll.
//
// Union satisfied by [OptimizeCompletedResult] or [OptimizeProcessingResult].
type OptimizeResultUnion interface {
	implementsOptimizeResult()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*OptimizeResultUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OptimizeCompletedResult{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(OptimizeProcessingResult{}),
		},
	)
}

// Always `processing`
type OptimizeResultStatus string

const (
	OptimizeResultStatusProcessing OptimizeResultStatus = "processing"
)

func (r OptimizeResultStatus) IsKnown() bool {
	switch r {
	case OptimizeResultStatusProcessing:
		return true
	}
	return false
}

type OptimizeResultType string

const (
	OptimizeResultTypeFeatureCollection OptimizeResultType = "FeatureCollection"
)

func (r OptimizeResultType) IsKnown() bool {
	switch r {
	case OptimizeResultTypeFeatureCollection:
		return true
	}
	return false
}

type OptimizeNewParams struct {
	// Route optimization (Travelling Salesman) request. Finds the most efficient order
	// to visit a set of waypoints. Minimum 2 waypoints, maximum 50. For large inputs,
	// the request may be processed asynchronously.
	OptimizeRequest OptimizeRequestParam `json:"optimize_request" api:"required"`
	// Response format: json (default), geojson, csv, ndjson
	Format param.Field[string] `query:"format"`
}

func (r OptimizeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.OptimizeRequest)
}

// URLQuery serializes [OptimizeNewParams]'s query parameters as `url.Values`.
func (r OptimizeNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
