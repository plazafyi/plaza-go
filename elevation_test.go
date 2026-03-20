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

func TestElevationBatch(t *testing.T) {
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
	_, err := client.Elevation.Batch(context.TODO(), githubcomplazafyiplazago.ElevationBatchParams{
		ElevationProfileRequest: githubcomplazafyiplazago.ElevationProfileRequestParam{
			Geometry: githubcomplazafyiplazago.F(githubcomplazafyiplazago.GeoJsonGeometryParam{
				Coordinates: githubcomplazafyiplazago.F[githubcomplazafyiplazago.GeoJsonGeometryCoordinatesUnionParam](githubcomplazafyiplazago.GeoJsonGeometryCoordinatesArrayParam([]float64{0.000000})),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.GeoJsonGeometryTypePoint),
			}),
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

func TestElevationLookupWithOptionalParams(t *testing.T) {
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
	_, err := client.Elevation.Lookup(context.TODO(), githubcomplazafyiplazago.ElevationLookupParams{
		Lat:       githubcomplazafyiplazago.F(0.000000),
		Lng:       githubcomplazafyiplazago.F(0.000000),
		Locations: githubcomplazafyiplazago.F("locations"),
	})
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestElevationProfile(t *testing.T) {
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
	_, err := client.Elevation.Profile(context.TODO(), githubcomplazafyiplazago.ElevationProfileParams{
		ElevationProfileRequest: githubcomplazafyiplazago.ElevationProfileRequestParam{
			Geometry: githubcomplazafyiplazago.F(githubcomplazafyiplazago.GeoJsonGeometryParam{
				Coordinates: githubcomplazafyiplazago.F[githubcomplazafyiplazago.GeoJsonGeometryCoordinatesUnionParam](githubcomplazafyiplazago.GeoJsonGeometryCoordinatesArrayParam([]float64{0.000000})),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.GeoJsonGeometryTypePoint),
			}),
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
