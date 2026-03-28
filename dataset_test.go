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

func TestDatasetNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Datasets.New(context.TODO(), githubcomplazafyiplazago.DatasetNewParams{
		Name:        githubcomplazafyiplazago.F("NYC Bike Lanes"),
		Slug:        githubcomplazafyiplazago.F("nyc-bike-lanes"),
		Attribution: githubcomplazafyiplazago.F("attribution"),
		Description: githubcomplazafyiplazago.F("description"),
		License:     githubcomplazafyiplazago.F("license"),
		SourceURL:   githubcomplazafyiplazago.F("https://example.com"),
	})
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDatasetGet(t *testing.T) {
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
	_, err := client.Datasets.Get(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDatasetList(t *testing.T) {
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
	_, err := client.Datasets.List(context.TODO())
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDatasetDelete(t *testing.T) {
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
	err := client.Datasets.Delete(context.TODO(), "id")
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDatasetFeaturesWithOptionalParams(t *testing.T) {
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
	_, err := client.Datasets.Features(
		context.TODO(),
		"id",
		githubcomplazafyiplazago.DatasetFeaturesParams{
			Cursor:          githubcomplazafyiplazago.F("cursor"),
			Format:          githubcomplazafyiplazago.F("format"),
			Limit:           githubcomplazafyiplazago.F(int64(0)),
			OutputBuffer:    githubcomplazafyiplazago.F(0.000000),
			OutputCentroid:  githubcomplazafyiplazago.F(true),
			OutputFields:    githubcomplazafyiplazago.F("output[fields]"),
			OutputGeometry:  githubcomplazafyiplazago.F(true),
			OutputInclude:   githubcomplazafyiplazago.F("output[include]"),
			OutputPrecision: githubcomplazafyiplazago.F(int64(0)),
			OutputSimplify:  githubcomplazafyiplazago.F(0.000000),
			OutputSort:      githubcomplazafyiplazago.F("output[sort]"),
		},
	)
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
