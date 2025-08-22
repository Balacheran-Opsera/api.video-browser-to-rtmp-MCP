package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"bytes"

	"github.com/api-video/mcp-server/config"
	"github.com/api-video/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Post_videoHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.GeneratedType
		
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
		url := fmt.Sprintf("%s/videos", cfg.BaseURL)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
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
		var result models.Video
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

func CreatePost_videoTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_videos",
		mcp.WithDescription("Create a video"),
		mcp.WithArray("tags", mcp.Description("Input parameter: A list of tags you want to use to describe your video.")),
		mcp.WithString("description", mcp.Description("Input parameter: A brief description of your video.")),
		mcp.WithBoolean("panoramic", mcp.Description("Input parameter: Indicates if your video is a 360/immersive video.")),
		mcp.WithBoolean("public", mcp.Description("Input parameter: Whether your video can be viewed by everyone, or requires authentication to see it. A setting of false will require a unique token for each view. Default is true. Tutorials on [private videos](https://api.video/blog/endpoints/private-videos).")),
		mcp.WithString("publishedAt", mcp.Description("Input parameter: The API uses ISO-8601 format for time, and includes 3 places for milliseconds.")),
		mcp.WithString("source", mcp.Description("Input parameter: If you add a video already on the web, this is where you enter the url for the video.")),
		mcp.WithBoolean("mp4Support", mcp.Description("Input parameter: Enables mp4 version in addition to streamed version.")),
		mcp.WithString("title", mcp.Required(), mcp.Description("Input parameter: The title of your new video.")),
		mcp.WithArray("metadata", mcp.Description("Input parameter: A list of key value pairs that you use to provide metadata for your video. These pairs can be made dynamic, allowing you to segment your audience. Read more on [dynamic metadata](https://api.video/blog/endpoints/dynamic-metadata).")),
		mcp.WithString("playerId", mcp.Description("Input parameter: The unique identification number for your video player.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_videoHandler(cfg),
	}
}
