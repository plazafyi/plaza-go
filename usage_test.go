// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomplazafyiplazago_test

import (
	"context"
	"os"
	"testing"

	"github.com/plazafyi/plaza-go"
	"github.com/plazafyi/plaza-go/internal/testutil"
	"github.com/plazafyi/plaza-go/option"
)

func TestUsage(t *testing.T) {
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
	featureCollection, err := client.Elements.Nearby(context.TODO(), githubcomplazafyiplazago.ElementNearbyParams{
		Lat:    githubcomplazafyiplazago.F(48.858400),
		Lng:    githubcomplazafyiplazago.F(0.000000),
		Radius: githubcomplazafyiplazago.F(int64(500)),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v\n", featureCollection.Features)
}
