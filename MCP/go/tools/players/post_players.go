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

func Post_playersHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		// Create properly typed request body using the generated schema
		var requestBody models.PlayerCreationPayload
		
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
		url := fmt.Sprintf("%s/players", cfg.BaseURL)
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
		var result models.Player
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

func CreatePost_playersTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_players",
		mcp.WithDescription("Create a player"),
		mcp.WithBoolean("forceAutoplay", mcp.Description("Input parameter: enable/disable player autoplay. Default: false")),
		mcp.WithString("link", mcp.Description("Input parameter: RGBA color for all controls. Default: rgba(255, 255, 255, 1)")),
		mcp.WithString("text", mcp.Description("Input parameter: RGBA color for timer text. Default: rgba(255, 255, 255, 1)")),
		mcp.WithString("backgroundTop", mcp.Description("Input parameter: RGBA color: top 50% of background. Default: rgba(0, 0, 0, .7)")),
		mcp.WithBoolean("enableApi", mcp.Description("Input parameter: enable/disable player SDK access. Default: true")),
		mcp.WithBoolean("enableControls", mcp.Description("Input parameter: enable/disable player controls. Default: true")),
		mcp.WithBoolean("forceLoop", mcp.Description("Input parameter: enable/disable looping. Default: false")),
		mcp.WithBoolean("hideTitle", mcp.Description("Input parameter: enable/disable title. Default: false")),
		mcp.WithString("linkHover", mcp.Description("Input parameter: RGBA color for all controls when hovered. Default: rgba(255, 255, 255, 1)")),
		mcp.WithString("trackUnplayed", mcp.Description("Input parameter: RGBA color playback bar: downloaded but unplayed (buffered) content. Default: rgba(255, 255, 255, .35)")),
		mcp.WithString("trackPlayed", mcp.Description("Input parameter: RGBA color playback bar: played content. Default: rgba(88, 131, 255, .95)")),
		mcp.WithString("trackBackground", mcp.Description("Input parameter: RGBA color playback bar: background. Default: rgba(255, 255, 255, .2)")),
		mcp.WithString("backgroundBottom", mcp.Description("Input parameter: RGBA color: bottom 50% of background. Default: rgba(0, 0, 0, .7)")),
		mcp.WithString("backgroundText", mcp.Description("Input parameter: RGBA color for title text. Default: rgba(255, 255, 255, 1)")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Post_playersHandler(cfg),
	}
}
