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

func TestElementGet(t *testing.T) {
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
	_, err := client.Elements.Get(
		context.TODO(),
		"type",
		int64(0),
	)
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestElementBatch(t *testing.T) {
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
	_, err := client.Elements.Batch(context.TODO(), githubcomplazafyiplazago.ElementBatchParams{
		BatchRequest: githubcomplazafyiplazago.BatchRequestParam{
			Elements: githubcomplazafyiplazago.F([]githubcomplazafyiplazago.BatchRequestElementParam{{
				ID:   githubcomplazafyiplazago.F(int64(21154906)),
				Type: githubcomplazafyiplazago.F(githubcomplazafyiplazago.BatchRequestElementsTypeNode),
			}, {
				ID:   githubcomplazafyiplazago.F(int64(4589123)),
				Type: githubcomplazafyiplazago.F(githubcomplazafyiplazago.BatchRequestElementsTypeWay),
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

func TestElementLookup(t *testing.T) {
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
	_, err := client.Elements.Lookup(context.TODO())
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestElementNearbyWithOptionalParams(t *testing.T) {
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
	_, err := client.Elements.Nearby(context.TODO(), githubcomplazafyiplazago.ElementNearbyParams{
		Lat:             githubcomplazafyiplazago.F(0.000000),
		Limit:           githubcomplazafyiplazago.F(int64(0)),
		Lng:             githubcomplazafyiplazago.F(0.000000),
		Near:            githubcomplazafyiplazago.F("near"),
		OutputBuffer:    githubcomplazafyiplazago.F(0.000000),
		OutputCentroid:  githubcomplazafyiplazago.F(true),
		OutputFields:    githubcomplazafyiplazago.F("output[fields]"),
		OutputGeometry:  githubcomplazafyiplazago.F(true),
		OutputInclude:   githubcomplazafyiplazago.F("output[include]"),
		OutputPrecision: githubcomplazafyiplazago.F(int64(0)),
		OutputSimplify:  githubcomplazafyiplazago.F(0.000000),
		OutputSort:      githubcomplazafyiplazago.F("output[sort]"),
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

func TestElementNearbyPostWithOptionalParams(t *testing.T) {
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
	_, err := client.Elements.NearbyPost(context.TODO(), githubcomplazafyiplazago.ElementNearbyPostParams{
		Lat:             githubcomplazafyiplazago.F(0.000000),
		Limit:           githubcomplazafyiplazago.F(int64(0)),
		Lng:             githubcomplazafyiplazago.F(0.000000),
		Near:            githubcomplazafyiplazago.F("near"),
		OutputBuffer:    githubcomplazafyiplazago.F(0.000000),
		OutputCentroid:  githubcomplazafyiplazago.F(true),
		OutputFields:    githubcomplazafyiplazago.F("output[fields]"),
		OutputGeometry:  githubcomplazafyiplazago.F(true),
		OutputInclude:   githubcomplazafyiplazago.F("output[include]"),
		OutputPrecision: githubcomplazafyiplazago.F(int64(0)),
		OutputSimplify:  githubcomplazafyiplazago.F(0.000000),
		OutputSort:      githubcomplazafyiplazago.F("output[sort]"),
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

func TestElementQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Elements.Query(context.TODO(), githubcomplazafyiplazago.ElementQueryParams{
		Bbox:            githubcomplazafyiplazago.F("bbox"),
		Contains:        githubcomplazafyiplazago.F("contains"),
		Crosses:         githubcomplazafyiplazago.F("crosses"),
		Cursor:          githubcomplazafyiplazago.F("cursor"),
		Format:          githubcomplazafyiplazago.F("format"),
		H3:              githubcomplazafyiplazago.F("h3"),
		Intersects:      githubcomplazafyiplazago.F("intersects"),
		Limit:           githubcomplazafyiplazago.F(int64(0)),
		Near:            githubcomplazafyiplazago.F("near"),
		OutputBuffer:    githubcomplazafyiplazago.F(0.000000),
		OutputCentroid:  githubcomplazafyiplazago.F(true),
		OutputFields:    githubcomplazafyiplazago.F("output[fields]"),
		OutputGeometry:  githubcomplazafyiplazago.F(true),
		OutputInclude:   githubcomplazafyiplazago.F("output[include]"),
		OutputPrecision: githubcomplazafyiplazago.F(int64(0)),
		OutputSimplify:  githubcomplazafyiplazago.F(0.000000),
		OutputSort:      githubcomplazafyiplazago.F("output[sort]"),
		Radius:          githubcomplazafyiplazago.F(0.000000),
		Touches:         githubcomplazafyiplazago.F("touches"),
		Type:            githubcomplazafyiplazago.F("type"),
		Within:          githubcomplazafyiplazago.F("within"),
	})
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestElementQueryPostWithOptionalParams(t *testing.T) {
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
	_, err := client.Elements.QueryPost(context.TODO(), githubcomplazafyiplazago.ElementQueryPostParams{
		Bbox:            githubcomplazafyiplazago.F("bbox"),
		Contains:        githubcomplazafyiplazago.F("contains"),
		Crosses:         githubcomplazafyiplazago.F("crosses"),
		Cursor:          githubcomplazafyiplazago.F("cursor"),
		Format:          githubcomplazafyiplazago.F("format"),
		H3:              githubcomplazafyiplazago.F("h3"),
		Intersects:      githubcomplazafyiplazago.F("intersects"),
		Limit:           githubcomplazafyiplazago.F(int64(0)),
		Near:            githubcomplazafyiplazago.F("near"),
		OutputBuffer:    githubcomplazafyiplazago.F(0.000000),
		OutputCentroid:  githubcomplazafyiplazago.F(true),
		OutputFields:    githubcomplazafyiplazago.F("output[fields]"),
		OutputGeometry:  githubcomplazafyiplazago.F(true),
		OutputInclude:   githubcomplazafyiplazago.F("output[include]"),
		OutputPrecision: githubcomplazafyiplazago.F(int64(0)),
		OutputSimplify:  githubcomplazafyiplazago.F(0.000000),
		OutputSort:      githubcomplazafyiplazago.F("output[sort]"),
		Radius:          githubcomplazafyiplazago.F(0.000000),
		Touches:         githubcomplazafyiplazago.F("touches"),
		Type:            githubcomplazafyiplazago.F("type"),
		Within:          githubcomplazafyiplazago.F("within"),
	})
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
