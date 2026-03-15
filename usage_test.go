// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package plaza_test

import (
	"context"
	"os"
	"testing"

	"github.com/plazafyi/plaza-go"
	"github.com/plazafyi/plaza-go/internal/testutil"
	"github.com/plazafyi/plaza-go/option"
)

func TestUsage(t *testing.T) {
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
	datasets, err := client.V1.Datasets.List(context.TODO())
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", datasets.Datasets)
}
