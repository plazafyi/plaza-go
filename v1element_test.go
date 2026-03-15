// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plaza_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/plazafyi/plaza-go"
	"github.com/plazafyi/plaza-go/internal/testutil"
	"github.com/plazafyi/plaza-go/option"
)

func TestV1ElementGet(t *testing.T) {
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
	_, err := client.V1.Elements.Get(
		context.TODO(),
		0,
		plaza.V1ElementGetParams{
			Type: "type",
		},
	)
	if err != nil {
		var apierr *plaza.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ElementFetchBatch(t *testing.T) {
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
	_, err := client.V1.Elements.FetchBatch(context.TODO(), plaza.V1ElementFetchBatchParams{
		Elements: []plaza.V1ElementFetchBatchParamsElement{{
			ID:   0,
			Type: "node",
		}},
	})
	if err != nil {
		var apierr *plaza.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestV1ElementQueryWithOptionalParams(t *testing.T) {
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
	_, err := client.V1.Elements.Query(context.TODO(), plaza.V1ElementQueryParams{
		Bbox:   plaza.String("bbox"),
		Cursor: plaza.String("cursor"),
		H3:     plaza.String("h3"),
		Limit:  plaza.Int(0),
		Type:   plaza.String("type"),
	})
	if err != nil {
		var apierr *plaza.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
