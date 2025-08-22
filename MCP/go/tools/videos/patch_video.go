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

func Patch_videoHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
		url := fmt.Sprintf("%s/videos/%s", cfg.BaseURL, videoId)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
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

func CreatePatch_videoTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_videos_videoId",
		mcp.WithDescription("Update a video"),
		mcp.WithString("videoId", mcp.Required(), mcp.Description("The video ID for the video you want to delete.")),
		mcp.WithArray("tags", mcp.Description("Input parameter: A list of terms or words you want to tag the video with. Make sure the list includes all the tags you want as whatever you send in this list will overwrite the existing list for the video.")),
		mcp.WithString("title", mcp.Description("Input parameter: The title you want to use for your video.")),
		mcp.WithString("description", mcp.Description("Input parameter: A brief description of the video.")),
		mcp.WithArray("metadata", mcp.Description("Input parameter: A list (array) of dictionaries where each dictionary contains a key value pair that describes the video. As with tags, you must send the complete list of metadata you want as whatever you send here will overwrite the existing metadata for the video. [Dynamic Metadata](https://api.video/blog/endpoints/dynamic-metadata) allows you to define a key that allows any value pair.")),
		mcp.WithBoolean("mp4Support", mcp.Description("Input parameter: Whether the player supports the mp4 format.")),
		mcp.WithBoolean("panoramic", mcp.Description("Input parameter: Whether the video is a 360 degree or immersive video.")),
		mcp.WithString("playerId", mcp.Description("Input parameter: The unique ID for the player you want to associate with your video.")),
		mcp.WithBoolean("public", mcp.Description("Input parameter: Whether the video is publicly available or not. False means it is set to private. Default is true. Tutorials on [private videos](https://api.video/blog/endpoints/private-videos).")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Patch_videoHandler(cfg),
	}
}
