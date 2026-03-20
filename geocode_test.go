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
		Q:           githubcomplazafyiplazago.F("q"),
		CountryCode: githubcomplazafyiplazago.F("country_code"),
		Lang:        githubcomplazafyiplazago.F("lang"),
		Lat:         githubcomplazafyiplazago.F(0.000000),
		Layer:       githubcomplazafyiplazago.F("layer"),
		Limit:       githubcomplazafyiplazago.F(int64(0)),
		Lng:         githubcomplazafyiplazago.F(0.000000),
	})
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGeocodeAutocompletePostWithOptionalParams(t *testing.T) {
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
	_, err := client.Geocode.AutocompletePost(context.TODO(), githubcomplazafyiplazago.GeocodeAutocompletePostParams{
		Q:           githubcomplazafyiplazago.F("q"),
		CountryCode: githubcomplazafyiplazago.F("country_code"),
		Lang:        githubcomplazafyiplazago.F("lang"),
		Lat:         githubcomplazafyiplazago.F(0.000000),
		Layer:       githubcomplazafyiplazago.F("layer"),
		Limit:       githubcomplazafyiplazago.F(int64(0)),
		Lng:         githubcomplazafyiplazago.F(0.000000),
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
		Q:           githubcomplazafyiplazago.F("q"),
		Bbox:        githubcomplazafyiplazago.F("bbox"),
		CountryCode: githubcomplazafyiplazago.F("country_code"),
		Lang:        githubcomplazafyiplazago.F("lang"),
		Lat:         githubcomplazafyiplazago.F(0.000000),
		Layer:       githubcomplazafyiplazago.F("layer"),
		Limit:       githubcomplazafyiplazago.F(int64(0)),
		Lng:         githubcomplazafyiplazago.F(0.000000),
	})
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestGeocodeForwardPostWithOptionalParams(t *testing.T) {
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
	_, err := client.Geocode.ForwardPost(context.TODO(), githubcomplazafyiplazago.GeocodeForwardPostParams{
		Q:           githubcomplazafyiplazago.F("q"),
		Bbox:        githubcomplazafyiplazago.F("bbox"),
		CountryCode: githubcomplazafyiplazago.F("country_code"),
		Lang:        githubcomplazafyiplazago.F("lang"),
		Lat:         githubcomplazafyiplazago.F(0.000000),
		Layer:       githubcomplazafyiplazago.F("layer"),
		Limit:       githubcomplazafyiplazago.F(int64(0)),
		Lng:         githubcomplazafyiplazago.F(0.000000),
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
		Lang:   githubcomplazafyiplazago.F("lang"),
		Lat:    githubcomplazafyiplazago.F(0.000000),
		Layer:  githubcomplazafyiplazago.F("layer"),
		Limit:  githubcomplazafyiplazago.F(int64(0)),
		Lng:    githubcomplazafyiplazago.F(0.000000),
		Near:   githubcomplazafyiplazago.F("near"),
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

func TestGeocodeReversePostWithOptionalParams(t *testing.T) {
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
	_, err := client.Geocode.ReversePost(context.TODO(), githubcomplazafyiplazago.GeocodeReversePostParams{
		Lang:   githubcomplazafyiplazago.F("lang"),
		Lat:    githubcomplazafyiplazago.F(0.000000),
		Layer:  githubcomplazafyiplazago.F("layer"),
		Limit:  githubcomplazafyiplazago.F(int64(0)),
		Lng:    githubcomplazafyiplazago.F(0.000000),
		Near:   githubcomplazafyiplazago.F("near"),
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
