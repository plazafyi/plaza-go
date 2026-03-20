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
		Lat:             githubcomplazafyiplazago.F(0.000000),
		Lng:             githubcomplazafyiplazago.F(0.000000),
		Time:            githubcomplazafyiplazago.F(0.000000),
		Mode:            githubcomplazafyiplazago.F("mode"),
		OutputFields:    githubcomplazafyiplazago.F("output[fields]"),
		OutputGeometry:  githubcomplazafyiplazago.F(true),
		OutputInclude:   githubcomplazafyiplazago.F("output[include]"),
		OutputPrecision: githubcomplazafyiplazago.F(int64(0)),
		OutputSimplify:  githubcomplazafyiplazago.F(0.000000),
	})
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoutingIsochronePostWithOptionalParams(t *testing.T) {
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
	_, err := client.Routing.IsochronePost(context.TODO(), githubcomplazafyiplazago.RoutingIsochronePostParams{
		Lat:             githubcomplazafyiplazago.F(0.000000),
		Lng:             githubcomplazafyiplazago.F(0.000000),
		Time:            githubcomplazafyiplazago.F(0.000000),
		Mode:            githubcomplazafyiplazago.F("mode"),
		OutputFields:    githubcomplazafyiplazago.F("output[fields]"),
		OutputGeometry:  githubcomplazafyiplazago.F(true),
		OutputInclude:   githubcomplazafyiplazago.F("output[include]"),
		OutputPrecision: githubcomplazafyiplazago.F(int64(0)),
		OutputSimplify:  githubcomplazafyiplazago.F(0.000000),
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
			Destinations: githubcomplazafyiplazago.F([]githubcomplazafyiplazago.MatrixRequestDestinationParam{{
				Lat: githubcomplazafyiplazago.F(48.858400),
				Lng: githubcomplazafyiplazago.F(2.294500),
			}}),
			Origins: githubcomplazafyiplazago.F([]githubcomplazafyiplazago.MatrixRequestOriginParam{{
				Lat: githubcomplazafyiplazago.F(48.856600),
				Lng: githubcomplazafyiplazago.F(2.352200),
			}, {
				Lat: githubcomplazafyiplazago.F(48.860600),
				Lng: githubcomplazafyiplazago.F(2.337600),
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
		Lat:             githubcomplazafyiplazago.F(0.000000),
		Lng:             githubcomplazafyiplazago.F(0.000000),
		OutputFields:    githubcomplazafyiplazago.F("output[fields]"),
		OutputInclude:   githubcomplazafyiplazago.F("output[include]"),
		OutputPrecision: githubcomplazafyiplazago.F(int64(0)),
		Radius:          githubcomplazafyiplazago.F(int64(0)),
	})
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestRoutingNearestPostWithOptionalParams(t *testing.T) {
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
	_, err := client.Routing.NearestPost(context.TODO(), githubcomplazafyiplazago.RoutingNearestPostParams{
		Lat:             githubcomplazafyiplazago.F(0.000000),
		Lng:             githubcomplazafyiplazago.F(0.000000),
		OutputFields:    githubcomplazafyiplazago.F("output[fields]"),
		OutputInclude:   githubcomplazafyiplazago.F("output[include]"),
		OutputPrecision: githubcomplazafyiplazago.F(int64(0)),
		Radius:          githubcomplazafyiplazago.F(int64(0)),
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
			Destination: githubcomplazafyiplazago.F(githubcomplazafyiplazago.RouteRequestDestinationParam{
				Lat: githubcomplazafyiplazago.F(48.858400),
				Lng: githubcomplazafyiplazago.F(2.294500),
			}),
			Origin: githubcomplazafyiplazago.F(githubcomplazafyiplazago.RouteRequestOriginParam{
				Lat: githubcomplazafyiplazago.F(48.856600),
				Lng: githubcomplazafyiplazago.F(2.352200),
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
			Waypoints: githubcomplazafyiplazago.F([]githubcomplazafyiplazago.RouteRequestWaypointParam{{
				Lat: githubcomplazafyiplazago.F(48.856600),
				Lng: githubcomplazafyiplazago.F(2.352200),
			}}),
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
