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

func TestAgentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.New(context.TODO(), githubcomnimblewaynimblego.AgentNewParams{
		AgentName:   githubcomnimblewaynimblego.String("agent_name"),
		Description: githubcomnimblewaynimblego.String("description"),
		DisplayName: githubcomnimblewaynimblego.String("display_name"),
		Effort:      githubcomnimblewaynimblego.AgentNewParamsEffortLow,
		Goals:       []string{"string"},
		Icon:        githubcomnimblewaynimblego.String("icon"),
		IsActive:    githubcomnimblewaynimblego.Bool(true),
		OutputSchema: map[string]any{
			"foo": "bar",
		},
		Skill: githubcomnimblewaynimblego.String("skill"),
		Sources: githubcomnimblewaynimblego.AgentNewParamsSources{
			Allow: []githubcomnimblewaynimblego.AgentNewParamsSourcesAllow{{
				Domains: []string{"string"},
				Title:   "title",
				Order:   githubcomnimblewaynimblego.Int(0),
			}},
			Avoid: githubcomnimblewaynimblego.String("avoid"),
			Block: []githubcomnimblewaynimblego.AgentNewParamsSourcesBlock{{
				Domains: []string{"string"},
				Title:   "title",
				Order:   githubcomnimblewaynimblego.Int(0),
			}},
			Prioritize: githubcomnimblewaynimblego.String("prioritize"),
		},
		SuggestedQuestions: []string{"string"},
		Template:           githubcomnimblewaynimblego.String("template"),
		UseCase:            githubcomnimblewaynimblego.AgentNewParamsUseCaseResearch,
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentUpdate(t *testing.T) {
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
	_, err := client.Agents.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		githubcomnimblewaynimblego.AgentUpdateParams{
			Body: []githubcomnimblewaynimblego.AgentUpdateParamsBody{{
				Op:    "add",
				Path:  "path",
				From:  githubcomnimblewaynimblego.String("from"),
				Value: map[string]any{},
			}},
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
	_, err := client.Agents.List(context.TODO(), githubcomnimblewaynimblego.AgentListParams{
		Limit:       githubcomnimblewaynimblego.Int(1),
		Offset:      githubcomnimblewaynimblego.Int(0),
		WorkspaceID: githubcomnimblewaynimblego.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentDelete(t *testing.T) {
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
	err := client.Agents.Delete(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
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
	_, err := client.Agents.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
