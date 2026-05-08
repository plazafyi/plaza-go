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

func TestOptimizeNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Optimize.New(context.TODO(), githubcomplazafyiplazago.OptimizeNewParams{
		OptimizeRequest: githubcomplazafyiplazago.OptimizeRequestParam{
			Waypoints: githubcomplazafyiplazago.F(githubcomplazafyiplazago.MultiPointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([][]float64{{2.352200, 48.856600}, {2.337600, 48.860600}, {2.294500, 48.858400}}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.MultiPointGeometryTypeMultiPoint),
			}),
			Mode:      githubcomplazafyiplazago.F(githubcomplazafyiplazago.OptimizeRequestModeAuto),
			Roundtrip: githubcomplazafyiplazago.F(false),
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

func TestOptimizeGet(t *testing.T) {
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
	_, err := client.Optimize.Get(context.TODO(), "job_id")
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
