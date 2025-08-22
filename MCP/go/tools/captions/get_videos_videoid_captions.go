package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/api-video/mcp-server/config"
	"github.com/api-video/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Get_videos_videoid_captionsHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		videoIdVal, ok := args["videoId"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: videoId"), nil
		}
		videoId, ok := videoIdVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: videoId"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["currentPage"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("currentPage=%v", val))
		}
		if val, ok := args["pageSize"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("pageSize=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		url := fmt.Sprintf("%s/videos/%s/captions%s", cfg.BaseURL, videoId, queryString)
		req, err := http.NewRequest("GET", url, nil)
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
		var result models.GeneratedType
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

func CreateGet_videos_videoid_captionsTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_videos_videoId_captions",
		mcp.WithDescription("List video captions"),
		mcp.WithString("videoId", mcp.Required(), mcp.Description("The unique identifier for the video you want to retrieve a list of captions for.")),
		mcp.WithNumber("currentPage", mcp.Description("Choose the number of search results to return per page. Minimum value: 1")),
		mcp.WithNumber("pageSize", mcp.Description("Results per page. Allowed values 1-100, default is 25.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_videos_videoid_captionsHandler(cfg),
	}
}
