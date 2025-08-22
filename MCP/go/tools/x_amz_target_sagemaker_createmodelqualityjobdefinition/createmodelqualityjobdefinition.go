package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/amazon-sagemaker-service/mcp-server/config"
	"github.com/amazon-sagemaker-service/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func CreatemodelqualityjobdefinitionHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.CreateModelQualityJobDefinitionRequest
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/#X-Amz-Target=SageMaker.CreateModelQualityJobDefinition", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Amz-Target"]; ok {
			req.Header.Set("X-Amz-Target", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.CreateModelQualityJobDefinitionResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateCreatemodelqualityjobdefinitionTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_#X-Amz-Target=SageMaker_CreateModelQualityJobDefinition",
		mcp.WithDescription("Creates a definition for a job that monitors model quality and drift. For information about model monitor, see <a href="https://docs.aws.amazon.com/sagemaker/latest/dg/model-monitor.html">Amazon SageMaker Model Monitor</a>."),
		mcp.WithString("X-Amz-Target", mcp.Required(), mcp.Description("")),
		mcp.WithObject("StoppingCondition", mcp.Description("Input parameter: A time limit for how long the monitoring job is allowed to run before stopping.")),
		mcp.WithString("JobDefinitionName", mcp.Required(), mcp.Description("")),
		mcp.WithString("ModelQualityAppSpecification", mcp.Required(), mcp.Description("")),
		mcp.WithObject("ModelQualityJobOutputConfig", mcp.Required(), mcp.Description("Input parameter: The output configuration for monitoring jobs.")),
		mcp.WithString("RoleArn", mcp.Required(), mcp.Description("")),
		mcp.WithString("Tags", mcp.Description("")),
		mcp.WithObject("JobResources", mcp.Required(), mcp.Description("Input parameter: Identifies the resources to deploy for a monitoring job.")),
		mcp.WithString("ModelQualityBaselineConfig", mcp.Description("")),
		mcp.WithString("ModelQualityJobInput", mcp.Required(), mcp.Description("")),
		mcp.WithString("NetworkConfig", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreatemodelqualityjobdefinitionHandler(cfg),
	}
}
