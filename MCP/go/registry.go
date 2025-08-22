package main

import (
	"github.com/api-video/mcp-server/config"
	"github.com/api-video/mcp-server/models"
	tools_chapters "github.com/api-video/mcp-server/tools/chapters"
	tools_captions "github.com/api-video/mcp-server/tools/captions"
	tools_analytics "github.com/api-video/mcp-server/tools/analytics"
	tools_live "github.com/api-video/mcp-server/tools/live"
	tools_videos "github.com/api-video/mcp-server/tools/videos"
	tools_players "github.com/api-video/mcp-server/tools/players"
	tools_webhooks "github.com/api-video/mcp-server/tools/webhooks"
	tools_authentication "github.com/api-video/mcp-server/tools/authentication"
	tools_videos_delegated_upload "github.com/api-video/mcp-server/tools/videos_delegated_upload"
	tools_account "github.com/api-video/mcp-server/tools/account"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_chapters.CreateGet_videos_videoid_chaptersTool(cfg),
		tools_captions.CreateGet_videos_videoid_captionsTool(cfg),
		tools_analytics.CreateGet_analytics_videos_videoidTool(cfg),
		tools_live.CreateDelete_live_streams_livestreamid_thumbnailTool(cfg),
		tools_chapters.CreateDelete_videos_videoid_chapters_languageTool(cfg),
		tools_chapters.CreateGet_videos_videoid_chapters_languageTool(cfg),
		tools_videos.CreateGet_video_statusTool(cfg),
		tools_players.CreateDelete_players_playerid_logoTool(cfg),
		tools_analytics.CreateGet_analytics_live_streams_livestreamidTool(cfg),
		tools_analytics.CreateGet_analytics_sessions_sessionid_eventsTool(cfg),
		tools_webhooks.CreateList_webhooksTool(cfg),
		tools_webhooks.CreatePost_webhooksTool(cfg),
		tools_authentication.CreatePost_auth_api_keyTool(cfg),
		tools_videos.CreateDelete_videoTool(cfg),
		tools_videos.CreateGet_videoTool(cfg),
		tools_videos.CreatePatch_videoTool(cfg),
		tools_webhooks.CreateDelete_webhookTool(cfg),
		tools_webhooks.CreateGet_webhookTool(cfg),
		tools_captions.CreateDelete_videos_videoid_captions_languageTool(cfg),
		tools_captions.CreateGet_videos_videoid_captions_languageTool(cfg),
		tools_captions.CreatePatch_videos_videoid_captions_languageTool(cfg),
		tools_videos_delegated_upload.CreateGet_upload_tokensTool(cfg),
		tools_videos_delegated_upload.CreatePost_upload_tokensTool(cfg),
		tools_live.CreatePatch_live_streams_livestreamidTool(cfg),
		tools_live.CreateDelete_live_streams_livestreamidTool(cfg),
		tools_live.CreateGet_live_streams_livestreamidTool(cfg),
		tools_videos_delegated_upload.CreateDelete_upload_tokens_uploadtokenTool(cfg),
		tools_videos_delegated_upload.CreateGet_upload_tokens_uploadtokenTool(cfg),
		tools_videos.CreatePatch_videos_videoid_thumbnailTool(cfg),
		tools_players.CreateDelete_players_playeridTool(cfg),
		tools_players.CreateGet_players_playeridTool(cfg),
		tools_players.CreatePatch_players_playeridTool(cfg),
		tools_authentication.CreatePost_auth_refreshTool(cfg),
		tools_players.CreateGet_playersTool(cfg),
		tools_players.CreatePost_playersTool(cfg),
		tools_account.CreateGet_accountTool(cfg),
		tools_live.CreateGet_live_streamsTool(cfg),
		tools_live.CreatePost_live_streamsTool(cfg),
		tools_videos.CreateList_videosTool(cfg),
		tools_videos.CreatePost_videoTool(cfg),
	}
}
