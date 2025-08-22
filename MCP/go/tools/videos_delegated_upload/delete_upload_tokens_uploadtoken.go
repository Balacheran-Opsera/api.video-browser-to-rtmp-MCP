package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/api-video/mcp-server/config"
	"github.com/api-video/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Delete_upload_tokens_uploadtokenHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		uploadTokenVal, ok := args["uploadToken"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: uploadToken"), nil
		}
		uploadToken, ok := uploadTokenVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: uploadToken"), nil
		}
		url := fmt.Sprintf("%s/upload-tokens/%s", cfg.BaseURL, uploadToken)
		req, err := http.NewRequest("DELETE", url, nil)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		if cfg.BearerToken != "" {
			req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", cfg.BearerToken))
		}
		req.Header.Set("Accept", "application/json")

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
		var result map[string]interface{}
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

func CreateDelete_upload_tokens_uploadtokenTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("delete_upload-tokens_uploadToken",
		mcp.WithDescription("Delete an upload token"),
		mcp.WithString("uploadToken", mcp.Required(), mcp.Description("The unique identifier for the upload token you want to delete. Deleting a token will make it so the token can no longer be used for authentication.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Delete_upload_tokens_uploadtokenHandler(cfg),
	}
}
