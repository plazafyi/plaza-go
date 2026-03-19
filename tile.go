// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomplazafyiplazago

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/plazafyi/plaza-go/internal/requestconfig"
	"github.com/plazafyi/plaza-go/option"
)

// TileService contains methods and other services that help with interacting with
// the plaza API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTileService] method instead.
type TileService struct {
	Options []option.RequestOption
}

// NewTileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTileService(opts ...option.RequestOption) (r *TileService) {
	r = &TileService{}
	r.Options = opts
	return
}

// Get a Mapbox Vector Tile
func (r *TileService) Get(ctx context.Context, z int64, x int64, y int64, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "application/vnd.mapbox-vector-tile")}, opts...)
	path := fmt.Sprintf("api/v1/tiles/%v/%v/%v", z, x, y)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}
