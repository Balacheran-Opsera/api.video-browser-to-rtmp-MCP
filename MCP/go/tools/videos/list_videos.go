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

func List_videosHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["title"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("title=%v", val))
		}
		if val, ok := args["tags"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("tags=%v", val))
		}
		if val, ok := args["metadata"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("metadata=%v", val))
		}
		if val, ok := args["description"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("description=%v", val))
		}
		if val, ok := args["liveStreamId"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("liveStreamId=%v", val))
		}
		if val, ok := args["sortBy"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sortBy=%v", val))
		}
		if val, ok := args["sortOrder"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("sortOrder=%v", val))
		}
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
		url := fmt.Sprintf("%s/videos%s", cfg.BaseURL, queryString)
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

func CreateList_videosTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("get_videos",
		mcp.WithDescription("List all videos"),
		mcp.WithString("title", mcp.Description("The title of a specific video you want to find. The search will match exactly to what term you provide and return any videos that contain the same term as part of their titles.")),
		mcp.WithArray("tags", mcp.Description("A tag is a category you create and apply to videos. You can search for videos with particular tags by listing one or more here. Only videos that have all the tags you list will be returned.")),
		mcp.WithArray("metadata", mcp.Description("Videos can be tagged with metadata tags in key:value pairs. You can search for videos with specific key value pairs using this parameter. [Dynamic Metadata](https://api.video/blog/endpoints/dynamic-metadata) allows you to define a key that allows any value pair.")),
		mcp.WithString("description", mcp.Description("If you described a video with a term or sentence, you can add it here to return videos containing this string.")),
		mcp.WithString("liveStreamId", mcp.Description("If you know the ID for a live stream, you can retrieve the stream by adding the ID for it here.")),
		mcp.WithString("sortBy", mcp.Description("Allowed: publishedAt, title. You can search by the time videos were published at, or by title.")),
		mcp.WithString("sortOrder", mcp.Description("Allowed: asc, desc. asc is ascending and sorts from A to Z. desc is descending and sorts from Z to A.")),
		mcp.WithNumber("currentPage", mcp.Description("Choose the number of search results to return per page. Minimum value: 1")),
		mcp.WithNumber("pageSize", mcp.Description("Results per page. Allowed values 1-100, default is 25.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    List_videosHandler(cfg),
	}
}
