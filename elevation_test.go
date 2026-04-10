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
		ElevationLookupRequest: githubcomplazafyiplazago.ElevationLookupRequestParam{
			Geometry: githubcomplazafyiplazago.F[githubcomplazafyiplazago.ElevationLookupRequestGeometryUnionParam](githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}),
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
			Geometry: githubcomplazafyiplazago.F(githubcomplazafyiplazago.LineStringGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([][]float64{{2.352200, 48.856600}, {2.340000, 48.858000}, {2.294500, 48.858400}}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.LineStringGeometryTypeLineString),
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
