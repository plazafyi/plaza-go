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

func TestElevationBatchWithOptionalParams(t *testing.T) {
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
		Coordinates: githubcomplazafyiplazago.F([]githubcomplazafyiplazago.ElevationBatchParamsCoordinate{{
			Lat: githubcomplazafyiplazago.F(48.856600),
			Lng: githubcomplazafyiplazago.F(2.352200),
		}, {
			Lat: githubcomplazafyiplazago.F(45.764000),
			Lng: githubcomplazafyiplazago.F(4.835700),
		}}),
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
		Format:          githubcomplazafyiplazago.F("format"),
		Lat:             githubcomplazafyiplazago.F(0.000000),
		Lng:             githubcomplazafyiplazago.F(0.000000),
		Locations:       githubcomplazafyiplazago.F("locations"),
		OutputFields:    githubcomplazafyiplazago.F("output[fields]"),
		OutputInclude:   githubcomplazafyiplazago.F("output[include]"),
		OutputPrecision: githubcomplazafyiplazago.F(int64(0)),
	})
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestElevationLookupPostWithOptionalParams(t *testing.T) {
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
	_, err := client.Elevation.LookupPost(context.TODO(), githubcomplazafyiplazago.ElevationLookupPostParams{
		Format:          githubcomplazafyiplazago.F("format"),
		Lat:             githubcomplazafyiplazago.F(0.000000),
		Lng:             githubcomplazafyiplazago.F(0.000000),
		Locations:       githubcomplazafyiplazago.F("locations"),
		OutputFields:    githubcomplazafyiplazago.F("output[fields]"),
		OutputInclude:   githubcomplazafyiplazago.F("output[include]"),
		OutputPrecision: githubcomplazafyiplazago.F(int64(0)),
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
			Coordinates: githubcomplazafyiplazago.F([]githubcomplazafyiplazago.ElevationProfileRequestCoordinateParam{{
				Lat: githubcomplazafyiplazago.F(48.856600),
				Lng: githubcomplazafyiplazago.F(2.352200),
			}, {
				Lat: githubcomplazafyiplazago.F(48.858000),
				Lng: githubcomplazafyiplazago.F(2.340000),
			}, {
				Lat: githubcomplazafyiplazago.F(48.858400),
				Lng: githubcomplazafyiplazago.F(2.294500),
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
