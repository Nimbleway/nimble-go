// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package githubcomnimblewaynimblego_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/Nimbleway/nimble-go"
	"github.com/Nimbleway/nimble-go/internal/testutil"
	"github.com/Nimbleway/nimble-go/option"
)

func TestJobNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomnimblewaynimblego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Jobs.New(context.TODO(), githubcomnimblewaynimblego.JobNewParams{
		ExtractTemplateName: "extract_template_name",
		Name:                "name",
		Description:         githubcomnimblewaynimblego.String("description"),
		Destination: githubcomnimblewaynimblego.JobNewParamsDestination{
			Path:   "path",
			Type:   "file",
			Format: "jsonl",
		},
		DisplayName: githubcomnimblewaynimblego.String("display_name"),
		Inputs: githubcomnimblewaynimblego.JobNewParamsInputs{
			Type: "s3",
			Data: []map[string]any{{
				"foo": "bar",
			}},
			FilePath: githubcomnimblewaynimblego.String("file_path"),
			NodeData: map[string][]map[string]any{
				"foo": {{
					"foo": "bar",
				}},
			},
		},
		Schedule: githubcomnimblewaynimblego.JobNewParamsSchedule{
			Cron:    "cron",
			Enabled: true,
		},
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomnimblewaynimblego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Jobs.Update(
		context.TODO(),
		"job_id",
		githubcomnimblewaynimblego.JobUpdateParams{
			Description: githubcomnimblewaynimblego.String("description"),
			Destination: githubcomnimblewaynimblego.JobUpdateParamsDestination{
				Path:   "path",
				Type:   "file",
				Format: "jsonl",
			},
			DisplayName: githubcomnimblewaynimblego.String("display_name"),
			Inputs: githubcomnimblewaynimblego.JobUpdateParamsInputs{
				Type: "s3",
				Data: []map[string]any{{
					"foo": "bar",
				}},
				FilePath: githubcomnimblewaynimblego.String("file_path"),
				NodeData: map[string][]map[string]any{
					"foo": {{
						"foo": "bar",
					}},
				},
			},
			Schedule: githubcomnimblewaynimblego.JobUpdateParamsSchedule{
				Cron:    "cron",
				Enabled: true,
			},
		},
	)
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomnimblewaynimblego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Jobs.List(context.TODO(), githubcomnimblewaynimblego.JobListParams{
		Limit:  githubcomnimblewaynimblego.Int(1),
		Offset: githubcomnimblewaynimblego.Int(0),
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomnimblewaynimblego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.Jobs.Delete(context.TODO(), "job_id")
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := githubcomnimblewaynimblego.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Jobs.Get(context.TODO(), "job_id")
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
