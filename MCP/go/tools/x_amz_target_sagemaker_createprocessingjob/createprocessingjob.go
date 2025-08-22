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

func CreateprocessingjobHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.CreateProcessingJobRequest
		
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
		url := fmt.Sprintf("%s/#X-Amz-Target=SageMaker.CreateProcessingJob", cfg.BaseURL)
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
		var result models.CreateProcessingJobResponse
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

func CreateCreateprocessingjobTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_#X-Amz-Target=SageMaker_CreateProcessingJob",
		mcp.WithDescription("Creates a processing job."),
		mcp.WithString("X-Amz-Target", mcp.Required(), mcp.Description("")),
		mcp.WithString("RoleArn", mcp.Required(), mcp.Description("")),
		mcp.WithString("AppSpecification", mcp.Required(), mcp.Description("")),
		mcp.WithString("ProcessingResources", mcp.Required(), mcp.Description("")),
		mcp.WithString("Tags", mcp.Description("")),
		mcp.WithString("Environment", mcp.Description("")),
		mcp.WithString("ProcessingOutputConfig", mcp.Description("")),
		mcp.WithString("StoppingCondition", mcp.Description("")),
		mcp.WithObject("ExperimentConfig", mcp.Description("Input parameter: <p>Associates a SageMaker job as a trial component with an experiment and trial. Specified when you call the following APIs:</p> <ul> <li> <p> <a href=\"https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateProcessingJob.html\">CreateProcessingJob</a> </p> </li> <li> <p> <a href=\"https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTrainingJob.html\">CreateTrainingJob</a> </p> </li> <li> <p> <a href=\"https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_CreateTransformJob.html\">CreateTransformJob</a> </p> </li> </ul>")),
		mcp.WithString("NetworkConfig", mcp.Description("")),
		mcp.WithString("ProcessingInputs", mcp.Description("")),
		mcp.WithString("ProcessingJobName", mcp.Required(), mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    CreateprocessingjobHandler(cfg),
	}
}
