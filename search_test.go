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

func TestSearchQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.Search.Query(context.TODO(), githubcomplazafyiplazago.SearchQueryParams{
		Q:               githubcomplazafyiplazago.F("q"),
		Cursor:          githubcomplazafyiplazago.F("cursor"),
		Limit:           githubcomplazafyiplazago.F(int64(0)),
		OutputFields:    githubcomplazafyiplazago.F("output[fields]"),
		OutputInclude:   githubcomplazafyiplazago.F("output[include]"),
		OutputPrecision: githubcomplazafyiplazago.F(int64(0)),
		OutputSort:      githubcomplazafyiplazago.F("output[sort]"),
	})
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSearchQueryPostWithOptionalParams(t *testing.T) {
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
	_, err := client.Search.QueryPost(context.TODO(), githubcomplazafyiplazago.SearchQueryPostParams{
		Q:               githubcomplazafyiplazago.F("q"),
		Cursor:          githubcomplazafyiplazago.F("cursor"),
		Limit:           githubcomplazafyiplazago.F(int64(0)),
		OutputFields:    githubcomplazafyiplazago.F("output[fields]"),
		OutputInclude:   githubcomplazafyiplazago.F("output[include]"),
		OutputPrecision: githubcomplazafyiplazago.F(int64(0)),
		OutputSort:      githubcomplazafyiplazago.F("output[sort]"),
	})
	if err != nil {
		var apierr *githubcomplazafyiplazago.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
