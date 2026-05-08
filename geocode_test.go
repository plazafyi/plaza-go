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

func TestGeocodeAutocompleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Geocode.Autocomplete(context.TODO(), githubcomplazafyiplazago.GeocodeAutocompleteParams{
		AutocompleteRequest: githubcomplazafyiplazago.AutocompleteRequestParam{
			Q:           githubcomplazafyiplazago.F("221B Bak"),
			CountryCode: githubcomplazafyiplazago.F("xx"),
			Focus: githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}),
			Lang:  githubcomplazafyiplazago.F("lang"),
			Layer: githubcomplazafyiplazago.F("layer"),
			Limit: githubcomplazafyiplazago.F(int64(1)),
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

func TestGeocodeBatch(t *testing.T) {
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
	_, err := client.Geocode.Batch(context.TODO(), githubcomplazafyiplazago.GeocodeBatchParams{
		Addresses: githubcomplazafyiplazago.F([]string{"string"}),
	})
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGeocodeForwardWithOptionalParams(t *testing.T) {
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
	_, err := client.Geocode.Forward(context.TODO(), githubcomplazafyiplazago.GeocodeForwardParams{
		GeocodeForwardRequest: githubcomplazafyiplazago.GeocodeForwardRequestParam{
			Q:           githubcomplazafyiplazago.F("221B Baker Street, London"),
			CountryCode: githubcomplazafyiplazago.F("xx"),
			Focus: githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}),
			Lang:  githubcomplazafyiplazago.F("lang"),
			Layer: githubcomplazafyiplazago.F("layer"),
			Limit: githubcomplazafyiplazago.F(int64(1)),
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

func TestGeocodeReverseWithOptionalParams(t *testing.T) {
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
	_, err := client.Geocode.Reverse(context.TODO(), githubcomplazafyiplazago.GeocodeReverseParams{
		GeocodeReverseRequest: githubcomplazafyiplazago.GeocodeReverseRequestParam{
			Geometry: githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}),
			Lang:   githubcomplazafyiplazago.F("lang"),
			Limit:  githubcomplazafyiplazago.F(int64(1)),
			Radius: githubcomplazafyiplazago.F(1.000000),
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
