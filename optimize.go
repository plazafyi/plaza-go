// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomplazafyiplazago

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"slices"

	"github.com/plazafyi/plaza-go/internal/apijson"
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
func (r *OptimizeService) New(ctx context.Context, body OptimizeNewParams, opts ...option.RequestOption) (res *OptimizeResult, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/geo+json")}, opts...)
	path := "api/v1/optimize"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
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

// Completed optimization — GeoJSON Feature with optimized route
type OptimizeCompletedResult struct {
	Geometry   GeoJsonGeometry                   `json:"geometry" api:"required"`
	Properties OptimizeCompletedResultProperties `json:"properties" api:"required"`
	// Job status
	Status OptimizeCompletedResultStatus `json:"status" api:"required"`
	Type   OptimizeCompletedResultType   `json:"type" api:"required"`
	JSON   optimizeCompletedResultJSON   `json:"-"`
}

// optimizeCompletedResultJSON contains the JSON metadata for the struct
// [OptimizeCompletedResult]
type optimizeCompletedResultJSON struct {
	Geometry    apijson.Field
	Properties  apijson.Field
	Status      apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OptimizeCompletedResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r optimizeCompletedResultJSON) RawJSON() string {
	return r.raw
}

func (r OptimizeCompletedResult) implementsOptimizeResult() {}

type OptimizeCompletedResultProperties struct {
	// Total distance in meters
	Distance float64 `json:"distance"`
	// Estimated duration in seconds
	Duration float64 `json:"duration"`
	// Optimized waypoint ordering
	WaypointOrder []int64                               `json:"waypoint_order"`
	JSON          optimizeCompletedResultPropertiesJSON `json:"-"`
}

// optimizeCompletedResultPropertiesJSON contains the JSON metadata for the struct
// [OptimizeCompletedResultProperties]
type optimizeCompletedResultPropertiesJSON struct {
	Distance      apijson.Field
	Duration      apijson.Field
	WaypointOrder apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OptimizeCompletedResultProperties) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r optimizeCompletedResultPropertiesJSON) RawJSON() string {
	return r.raw
}

// Job status
type OptimizeCompletedResultStatus string

const (
	OptimizeCompletedResultStatusCompleted OptimizeCompletedResultStatus = "completed"
)

func (r OptimizeCompletedResultStatus) IsKnown() bool {
	switch r {
	case OptimizeCompletedResultStatusCompleted:
		return true
	}
	return false
}

type OptimizeCompletedResultType string

const (
	OptimizeCompletedResultTypeFeature OptimizeCompletedResultType = "Feature"
)

func (r OptimizeCompletedResultType) IsKnown() bool {
	switch r {
	case OptimizeCompletedResultTypeFeature:
		return true
	}
	return false
}

// Status of an async optimization job
type OptimizeJobStatus struct {
	// Job status
	Status OptimizeJobStatusStatus `json:"status" api:"required"`
	// Error message when failed
	Error string `json:"error" api:"nullable"`
	// Optimization result when completed
	Result interface{}           `json:"result" api:"nullable"`
	JSON   optimizeJobStatusJSON `json:"-"`
}

// optimizeJobStatusJSON contains the JSON metadata for the struct
// [OptimizeJobStatus]
type optimizeJobStatusJSON struct {
	Status      apijson.Field
	Error       apijson.Field
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

// Job status
type OptimizeJobStatusStatus string

const (
	OptimizeJobStatusStatusCompleted  OptimizeJobStatusStatus = "completed"
	OptimizeJobStatusStatusProcessing OptimizeJobStatusStatus = "processing"
	OptimizeJobStatusStatusFailed     OptimizeJobStatusStatus = "failed"
)

func (r OptimizeJobStatusStatus) IsKnown() bool {
	switch r {
	case OptimizeJobStatusStatusCompleted, OptimizeJobStatusStatusProcessing, OptimizeJobStatusStatusFailed:
		return true
	}
	return false
}

// Async optimization in progress — poll with the job_id
type OptimizeProcessingResult struct {
	// Job ID for polling
	JobID string `json:"job_id" api:"required"`
	// Job status
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

// Job status
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

// Route optimization request through waypoints
type OptimizeRequestParam struct {
	// Waypoints to visit (GeoJSON MultiPoint geometry, minimum 2 points)
	Waypoints param.Field[GeoJsonGeometryParam] `json:"waypoints" api:"required"`
	// Travel mode (default: auto)
	Mode param.Field[OptimizeRequestMode] `json:"mode"`
	// Whether route returns to start (default: true)
	Roundtrip param.Field[bool] `json:"roundtrip"`
}

func (r OptimizeRequestParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Travel mode (default: auto)
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

// Optimization response — either a completed GeoJSON Feature route or an async job
// reference
type OptimizeResult struct {
	// Job status
	Status   OptimizeResultStatus `json:"status" api:"required"`
	Geometry GeoJsonGeometry      `json:"geometry"`
	// Job ID for polling
	JobID string `json:"job_id"`
	// This field can have the runtime type of [OptimizeCompletedResultProperties].
	Properties interface{}        `json:"properties"`
	Type       OptimizeResultType `json:"type"`
	JSON       optimizeResultJSON `json:"-"`
	union      OptimizeResultUnion
}

// optimizeResultJSON contains the JSON metadata for the struct [OptimizeResult]
type optimizeResultJSON struct {
	Status      apijson.Field
	Geometry    apijson.Field
	JobID       apijson.Field
	Properties  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
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

// Optimization response — either a completed GeoJSON Feature route or an async job
// reference
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

// Job status
type OptimizeResultStatus string

const (
	OptimizeResultStatusCompleted  OptimizeResultStatus = "completed"
	OptimizeResultStatusProcessing OptimizeResultStatus = "processing"
)

func (r OptimizeResultStatus) IsKnown() bool {
	switch r {
	case OptimizeResultStatusCompleted, OptimizeResultStatusProcessing:
		return true
	}
	return false
}

type OptimizeResultType string

const (
	OptimizeResultTypeFeature OptimizeResultType = "Feature"
)

func (r OptimizeResultType) IsKnown() bool {
	switch r {
	case OptimizeResultTypeFeature:
		return true
	}
	return false
}

type OptimizeNewParams struct {
	// Route optimization request through waypoints
	OptimizeRequest OptimizeRequestParam `json:"optimize_request" api:"required"`
}

func (r OptimizeNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.OptimizeRequest)
}
