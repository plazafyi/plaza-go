// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomplazafyiplazago_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/plazafyi/plaza-go"
	"github.com/plazafyi/plaza-go/internal/testutil"
	"github.com/plazafyi/plaza-go/option"
)

func TestRoutingIsochroneWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomplazafyiplazago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Routing.Isochrone(context.TODO(), githubcomplazafyiplazago.RoutingIsochroneParams{
		IsochroneRequest: githubcomplazafyiplazago.IsochroneRequestParam{
			Geometry: githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}),
			Time: githubcomplazafyiplazago.F([]int64{int64(1)}),
			Mode: githubcomplazafyiplazago.F(githubcomplazafyiplazago.IsochroneRequestModeAuto),
		},
		Format: githubcomplazafyiplazago.F("format"),
	})
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoutingMatrixWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomplazafyiplazago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Routing.Matrix(context.TODO(), githubcomplazafyiplazago.RoutingMatrixParams{
		MatrixRequest: githubcomplazafyiplazago.MatrixRequestParam{
			Destinations: githubcomplazafyiplazago.F([]githubcomplazafyiplazago.PointGeometryParam{{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.294500, 48.858400}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}}),
			Origins: githubcomplazafyiplazago.F([]githubcomplazafyiplazago.PointGeometryParam{{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}, {
				Coordinates: githubcomplazafyiplazago.F([]float64{2.337600, 48.860600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}}),
			Annotations:   githubcomplazafyiplazago.F("annotations"),
			FallbackSpeed: githubcomplazafyiplazago.F(1.000000),
			Mode:          githubcomplazafyiplazago.F(githubcomplazafyiplazago.MatrixRequestModeAuto),
		},
	})
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoutingNearestWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomplazafyiplazago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Routing.Nearest(context.TODO(), githubcomplazafyiplazago.RoutingNearestParams{
		NearestRequest: githubcomplazafyiplazago.NearestRequestParam{
			Geometry: githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}),
			Radius: githubcomplazafyiplazago.F(1.000000),
		},
	})
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoutingRouteWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomplazafyiplazago.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Routing.Route(context.TODO(), githubcomplazafyiplazago.RoutingRouteParams{
		RouteRequest: githubcomplazafyiplazago.RouteRequestParam{
			Destination: githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.294500, 48.858400}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}),
			Origin: githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}),
			Alternatives: githubcomplazafyiplazago.F(int64(0)),
			Annotations:  githubcomplazafyiplazago.F(true),
			DepartAt:     githubcomplazafyiplazago.F(time.Now()),
			Ev: githubcomplazafyiplazago.F(githubcomplazafyiplazago.RouteRequestEvParam{
				BatteryCapacityWh: githubcomplazafyiplazago.F(75000.000000),
				ConnectorTypes:    githubcomplazafyiplazago.F([]string{"string"}),
				InitialChargePct:  githubcomplazafyiplazago.F(0.000000),
				MinChargePct:      githubcomplazafyiplazago.F(0.000000),
				MinPowerKw:        githubcomplazafyiplazago.F(0.000000),
			}),
			Exclude:      githubcomplazafyiplazago.F("exclude"),
			Geometries:   githubcomplazafyiplazago.F(githubcomplazafyiplazago.RouteRequestGeometriesGeojson),
			Mode:         githubcomplazafyiplazago.F(githubcomplazafyiplazago.RouteRequestModeAuto),
			Overview:     githubcomplazafyiplazago.F(githubcomplazafyiplazago.RouteRequestOverviewFull),
			Steps:        githubcomplazafyiplazago.F(true),
			TrafficModel: githubcomplazafyiplazago.F(githubcomplazafyiplazago.RouteRequestTrafficModelBestGuess),
			Waypoints: githubcomplazafyiplazago.F([]githubcomplazafyiplazago.PointGeometryParam{{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}}),
		},
		Format: githubcomplazafyiplazago.F("format"),
	})
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
