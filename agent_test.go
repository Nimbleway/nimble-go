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

func TestAgentListWithOptionalParams(t *testing.T) {
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
	_, err := client.Agent.List(context.TODO(), githubcomnimblewaynimblego.AgentListParams{
		Limit:     githubcomnimblewaynimblego.Int(1),
		ManagedBy: githubcomnimblewaynimblego.AgentListParamsManagedByNimble,
		Offset:    githubcomnimblewaynimblego.Int(0),
		Privacy:   githubcomnimblewaynimblego.AgentListParamsPrivacyPublic,
		Search:    githubcomnimblewaynimblego.String("search"),
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentGenerateWithOptionalParams(t *testing.T) {
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
	_, err := client.Agent.Generate(context.TODO(), githubcomnimblewaynimblego.AgentGenerateParams{
		OfCreateAgentGenerationRequest: &githubcomnimblewaynimblego.AgentGenerateParamsBodyCreateAgentGenerationRequest{
			Prompt:      "prompt",
			URL:         "url",
			AgentName:   githubcomnimblewaynimblego.String("agent_name"),
			InputSchema: map[string]any{},
			Metadata: githubcomnimblewaynimblego.AgentGenerateParamsBodyCreateAgentGenerationRequestMetadata{
				Description: githubcomnimblewaynimblego.String("description"),
				DisplayName: githubcomnimblewaynimblego.String("display_name"),
				Tags:        []string{"string"},
			},
			OutputSchema: map[string]any{},
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

func TestAgentGet(t *testing.T) {
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
	_, err := client.Agent.Get(context.TODO(), "template_name")
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentGetGeneration(t *testing.T) {
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
	_, err := client.Agent.GetGeneration(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentPublish(t *testing.T) {
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
	_, err := client.Agent.Publish(
		context.TODO(),
		"agent_name",
		githubcomnimblewaynimblego.AgentPublishParams{
			VersionID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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

func TestAgentRunWithOptionalParams(t *testing.T) {
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
	_, err := client.Agent.Run(context.TODO(), githubcomnimblewaynimblego.AgentRunParams{
		Agent: "agent",
		Params: map[string]any{
			"foo": "bar",
		},
		Formats:      []string{"html", "markdown"},
		Localization: githubcomnimblewaynimblego.Bool(true),
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentRunAsyncWithOptionalParams(t *testing.T) {
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
	_, err := client.Agent.RunAsync(context.TODO(), githubcomnimblewaynimblego.AgentRunAsyncParams{
		Agent: "agent",
		Params: map[string]any{
			"foo": "bar",
		},
		CallbackURL:       githubcomnimblewaynimblego.String("https://example.com/webhook/callback"),
		Formats:           []string{"html", "markdown"},
		Localization:      githubcomnimblewaynimblego.Bool(true),
		StorageCompress:   githubcomnimblewaynimblego.Bool(true),
		StorageObjectName: githubcomnimblewaynimblego.String("result-2024-01-15.json"),
		StorageType:       githubcomnimblewaynimblego.String("s3"),
		StorageURL:        githubcomnimblewaynimblego.String("s3://bucket-name/path/to/object"),
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentRunBatchWithOptionalParams(t *testing.T) {
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
	_, err := client.Agent.RunBatch(context.TODO(), githubcomnimblewaynimblego.AgentRunBatchParams{
		Inputs: []githubcomnimblewaynimblego.AgentRunBatchParamsInput{{
			Formats:      []string{"html", "markdown"},
			Localization: githubcomnimblewaynimblego.Bool(true),
			Params: map[string]any{
				"foo": "bar",
			},
		}},
		SharedInputs: githubcomnimblewaynimblego.AgentRunBatchParamsSharedInputs{
			Agent:        "agent",
			Formats:      []string{"html", "markdown"},
			Localization: githubcomnimblewaynimblego.Bool(true),
			Params: map[string]any{
				"foo": "bar",
			},
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
