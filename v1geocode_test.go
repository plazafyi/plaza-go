// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plaza_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/plaza-go"
	"github.com/stainless-sdks/plaza-go/internal/testutil"
	"github.com/stainless-sdks/plaza-go/option"
)

func TestV1GeocodeAutocompleteWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := plaza.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Geocode.Autocomplete(context.TODO(), plaza.V1GeocodeAutocompleteParams{
		Q:     "q",
		Lat:   plaza.Float(0),
		Limit: plaza.Int(0),
		Lng:   plaza.Float(0),
	})
	if err != nil {
		var apierr *plaza.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1GeocodeForwardWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := plaza.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Geocode.Forward(context.TODO(), plaza.V1GeocodeForwardParams{
		Q:     "q",
		Lat:   plaza.Float(0),
		Limit: plaza.Int(0),
		Lng:   plaza.Float(0),
	})
	if err != nil {
		var apierr *plaza.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1GeocodeReverseWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := plaza.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.V1.Geocode.Reverse(context.TODO(), plaza.V1GeocodeReverseParams{
		Lat:    0,
		Lng:    0,
		Radius: plaza.Int(0),
	})
	if err != nil {
		var apierr *plaza.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
