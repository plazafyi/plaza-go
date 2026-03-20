// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomplazafyiplazago_test

import (
	"context"
	"errors"
	"os"
	"testing"

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
		Lat:  githubcomplazafyiplazago.F(0.000000),
		Lng:  githubcomplazafyiplazago.F(0.000000),
		Time: githubcomplazafyiplazago.F(0.000000),
		Mode: githubcomplazafyiplazago.F("mode"),
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
			Destinations: githubcomplazafyiplazago.F(githubcomplazafyiplazago.GeoJsonGeometryParam{
				Coordinates: githubcomplazafyiplazago.F[githubcomplazafyiplazago.GeoJsonGeometryCoordinatesUnionParam](githubcomplazafyiplazago.GeoJsonGeometryCoordinatesArrayParam([]float64{0.000000})),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.GeoJsonGeometryTypePoint),
			}),
			Origins: githubcomplazafyiplazago.F(githubcomplazafyiplazago.GeoJsonGeometryParam{
				Coordinates: githubcomplazafyiplazago.F[githubcomplazafyiplazago.GeoJsonGeometryCoordinatesUnionParam](githubcomplazafyiplazago.GeoJsonGeometryCoordinatesArrayParam([]float64{0.000000})),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.GeoJsonGeometryTypePoint),
			}),
			Mode: githubcomplazafyiplazago.F(githubcomplazafyiplazago.MatrixRequestModeAuto),
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
		Lat:    githubcomplazafyiplazago.F(0.000000),
		Lng:    githubcomplazafyiplazago.F(0.000000),
		Radius: githubcomplazafyiplazago.F(int64(0)),
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
			Destination: githubcomplazafyiplazago.F(githubcomplazafyiplazago.GeoJsonGeometryParam{
				Coordinates: githubcomplazafyiplazago.F[githubcomplazafyiplazago.GeoJsonGeometryCoordinatesUnionParam](githubcomplazafyiplazago.GeoJsonGeometryCoordinatesArrayParam([]float64{0.000000})),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.GeoJsonGeometryTypePoint),
			}),
			Origin: githubcomplazafyiplazago.F(githubcomplazafyiplazago.GeoJsonGeometryParam{
				Coordinates: githubcomplazafyiplazago.F[githubcomplazafyiplazago.GeoJsonGeometryCoordinatesUnionParam](githubcomplazafyiplazago.GeoJsonGeometryCoordinatesArrayParam([]float64{0.000000})),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.GeoJsonGeometryTypePoint),
			}),
			Mode: githubcomplazafyiplazago.F(githubcomplazafyiplazago.RouteRequestModeAuto),
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
