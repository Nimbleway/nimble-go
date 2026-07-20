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

func TestAgentRunNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Runs.New(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		githubcomnimblewaynimblego.AgentRunNewParams{
			Input:        "input",
			Effort:       githubcomnimblewaynimblego.AgentRunNewParamsEffortLow,
			EnableEvents: githubcomnimblewaynimblego.Bool(true),
			InputData: githubcomnimblewaynimblego.AgentRunNewParamsInputDataUnion{
				OfMapOfAnyMap: []map[string]any{{
					"foo": "bar",
				}},
			},
			OutputSchema: map[string]any{
				"foo": "bar",
			},
			PreviousInteractionID: githubcomnimblewaynimblego.String("previous_interaction_id"),
			Sources: githubcomnimblewaynimblego.AgentRunNewParamsSources{
				Allow: []githubcomnimblewaynimblego.AgentRunNewParamsSourcesAllow{{
					Domains: []string{"string"},
					Title:   "title",
					Order:   githubcomnimblewaynimblego.Int(0),
				}},
				Avoid: githubcomnimblewaynimblego.String("avoid"),
				Block: []githubcomnimblewaynimblego.AgentRunNewParamsSourcesBlock{{
					Domains: []string{"string"},
					Title:   "title",
					Order:   githubcomnimblewaynimblego.Int(0),
				}},
				Prioritize: githubcomnimblewaynimblego.String("prioritize"),
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

func TestAgentRunListWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Runs.List(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		githubcomnimblewaynimblego.AgentRunListParams{
			Limit:  githubcomnimblewaynimblego.Int(1),
			Offset: githubcomnimblewaynimblego.Int(0),
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

func TestAgentRunGet(t *testing.T) {
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
	_, err := client.Agents.Runs.Get(
		context.TODO(),
		"run_id",
		githubcomnimblewaynimblego.AgentRunGetParams{
			AgentID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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

func TestAgentRunResult(t *testing.T) {
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
	_, err := client.Agents.Runs.Result(
		context.TODO(),
		"run_id",
		githubcomnimblewaynimblego.AgentRunResultParams{
			AgentID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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

func TestAgentRunStreamEvents(t *testing.T) {
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
	err := client.Agents.Runs.StreamEvents(
		context.TODO(),
		"run_id",
		githubcomnimblewaynimblego.AgentRunStreamEventsParams{
			AgentID: "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
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
