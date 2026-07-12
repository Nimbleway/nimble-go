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

func TestTaskAgentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.TaskAgent.New(context.TODO(), githubcomnimblewaynimblego.TaskAgentNewParams{
		AgentName:       githubcomnimblewaynimblego.String("agent_name"),
		Description:     githubcomnimblewaynimblego.String("description"),
		DisplayName:     githubcomnimblewaynimblego.String("display_name"),
		DomainExpertise: githubcomnimblewaynimblego.String("domain_expertise"),
		Effort:          githubcomnimblewaynimblego.TaskAgentNewParamsEffortLow,
		Goals:           []string{"string"},
		Icon:            githubcomnimblewaynimblego.String("icon"),
		IsActive:        githubcomnimblewaynimblego.Bool(true),
		OutputSchema: map[string]any{
			"foo": "bar",
		},
		Sources: githubcomnimblewaynimblego.TaskAgentNewParamsSources{
			Allow: []githubcomnimblewaynimblego.TaskAgentNewParamsSourcesAllow{{
				Domains: []string{"string"},
				Title:   "title",
				Order:   githubcomnimblewaynimblego.Int(0),
			}},
			Avoid: githubcomnimblewaynimblego.String("avoid"),
			Block: []githubcomnimblewaynimblego.TaskAgentNewParamsSourcesBlock{{
				Domains: []string{"string"},
				Title:   "title",
				Order:   githubcomnimblewaynimblego.Int(0),
			}},
			Prioritize: githubcomnimblewaynimblego.String("prioritize"),
		},
		SuggestedQuestions: []string{"string"},
		Template:           githubcomnimblewaynimblego.String("template"),
		UseCase:            githubcomnimblewaynimblego.TaskAgentNewParamsUseCaseResearch,
		WorkspaceID:        githubcomnimblewaynimblego.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTaskAgentUpdate(t *testing.T) {
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
	_, err := client.TaskAgent.Update(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		githubcomnimblewaynimblego.TaskAgentUpdateParams{
			Body: []githubcomnimblewaynimblego.TaskAgentUpdateParamsBody{{
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

func TestTaskAgentListWithOptionalParams(t *testing.T) {
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
	_, err := client.TaskAgent.List(context.TODO(), githubcomnimblewaynimblego.TaskAgentListParams{
		FilterEffort:  githubcomnimblewaynimblego.TaskAgentListParamsFilterEffortLow,
		FilterUseCase: githubcomnimblewaynimblego.TaskAgentListParamsFilterUseCaseResearch,
		Limit:         githubcomnimblewaynimblego.Int(0),
		Offset:        githubcomnimblewaynimblego.Int(0),
		WorkspaceID:   githubcomnimblewaynimblego.String("182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e"),
	})
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTaskAgentDeactivate(t *testing.T) {
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
	err := client.TaskAgent.Deactivate(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTaskAgentGet(t *testing.T) {
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
	_, err := client.TaskAgent.Get(context.TODO(), "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e")
	if err != nil {
		var apierr *githubcomnimblewaynimblego.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestTaskAgentRunWithOptionalParams(t *testing.T) {
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
	_, err := client.TaskAgent.Run(
		context.TODO(),
		"182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		githubcomnimblewaynimblego.TaskAgentRunParams{
			Input:        "input",
			Effort:       githubcomnimblewaynimblego.TaskAgentRunParamsEffortLow,
			EnableEvents: githubcomnimblewaynimblego.Bool(true),
			OutputSchema: map[string]any{
				"foo": "bar",
			},
			PreviousInteractionID: githubcomnimblewaynimblego.String("previous_interaction_id"),
			Sources: githubcomnimblewaynimblego.TaskAgentRunParamsSources{
				Allow: []githubcomnimblewaynimblego.TaskAgentRunParamsSourcesAllow{{
					Domains: []string{"string"},
					Title:   "title",
					Order:   githubcomnimblewaynimblego.Int(0),
				}},
				Avoid: githubcomnimblewaynimblego.String("avoid"),
				Block: []githubcomnimblewaynimblego.TaskAgentRunParamsSourcesBlock{{
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
