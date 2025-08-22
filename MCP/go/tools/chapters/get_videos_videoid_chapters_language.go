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

func Get_videos_videoid_chapters_languageHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		languageVal, ok := args["language"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: language"), nil
		}
		language, ok := languageVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: language"), nil
		}
		url := fmt.Sprintf("%s/videos/%s/chapters/%s", cfg.BaseURL, videoId, language)
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
		var result models.Chapter
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

func CreateGet_videos_videoid_chapters_languageTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_videos_videoId_chapters_language",
		mcp.WithDescription("Show a chapter"),
		mcp.WithString("videoId", mcp.Required(), mcp.Description("The unique identifier for the video you want to show a chapter for.")),
		mcp.WithString("language", mcp.Required(), mcp.Description("A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Get_videos_videoid_chapters_languageHandler(cfg),
	}
}
