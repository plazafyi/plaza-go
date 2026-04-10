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

func TestFeatureGet(t *testing.T) {
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
	_, err := client.Features.Get(
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

func TestFeatureBatch(t *testing.T) {
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
	_, err := client.Features.Batch(context.TODO(), githubcomplazafyiplazago.FeatureBatchParams{
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

func TestFeatureQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Features.Query(context.TODO(), githubcomplazafyiplazago.FeatureQueryParams{
		Cursor: githubcomplazafyiplazago.F("cursor"),
		Format: githubcomplazafyiplazago.F("format"),
		H3:     githubcomplazafyiplazago.F("h3"),
		Limit:  githubcomplazafyiplazago.F(int64(0)),
		Type:   githubcomplazafyiplazago.F("type"),
		SpatialPredicate: githubcomplazafyiplazago.SpatialPredicateParam{
			Around: githubcomplazafyiplazago.F[githubcomplazafyiplazago.GeometryUnionParam](githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}),
			Contains: githubcomplazafyiplazago.F[githubcomplazafyiplazago.GeometryUnionParam](githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}),
			Crosses: githubcomplazafyiplazago.F[githubcomplazafyiplazago.GeometryUnionParam](githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}),
			Intersects: githubcomplazafyiplazago.F[githubcomplazafyiplazago.GeometryUnionParam](githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}),
			NotContains: githubcomplazafyiplazago.F[githubcomplazafyiplazago.GeometryUnionParam](githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}),
			NotIntersects: githubcomplazafyiplazago.F[githubcomplazafyiplazago.GeometryUnionParam](githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}),
			NotWithin: githubcomplazafyiplazago.F[githubcomplazafyiplazago.GeometryUnionParam](githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}),
			Radius: githubcomplazafyiplazago.F(500.000000),
			Touches: githubcomplazafyiplazago.F[githubcomplazafyiplazago.GeometryUnionParam](githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
			}),
			Within: githubcomplazafyiplazago.F[githubcomplazafyiplazago.GeometryUnionParam](githubcomplazafyiplazago.PointGeometryParam{
				Coordinates: githubcomplazafyiplazago.F([]float64{2.352200, 48.856600}),
				Type:        githubcomplazafyiplazago.F(githubcomplazafyiplazago.PointGeometryTypePoint),
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
