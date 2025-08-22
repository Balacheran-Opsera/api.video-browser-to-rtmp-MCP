package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Name string `json:"name,omitempty"`
	Status int `json:"status,omitempty"`
	Title string `json:"title,omitempty"`
	TypeField string `json:"type,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Pagination Pagination `json:"pagination,omitempty"`
	Data []GeneratedType `json:"data,omitempty"`
}

// Videosourcelivestream represents the Videosourcelivestream schema from the OpenAPI specification
type Videosourcelivestream struct {
	Livestreamid string `json:"liveStreamId,omitempty"` // The unique identifier for the live stream.
	Links []Videosourcelivestreamlink `json:"links,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	File string `json:"file"` // The name of the file you want to use for your logo.
	Link string `json:"link"` // The path to the file you want to upload and use as a logo.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Data []GeneratedType `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Panoramic bool `json:"panoramic,omitempty"` // Whether the video is a 360 degree or immersive video.
	Playerid string `json:"playerId,omitempty"` // The unique ID for the player you want to associate with your video.
	Public bool `json:"public,omitempty"` // Whether the video is publicly available or not. False means it is set to private. Default is true. Tutorials on [private videos](https://api.video/blog/endpoints/private-videos).
	Tags []string `json:"tags,omitempty"` // A list of terms or words you want to tag the video with. Make sure the list includes all the tags you want as whatever you send in this list will overwrite the existing list for the video.
	Title string `json:"title,omitempty"` // The title you want to use for your video.
	Description string `json:"description,omitempty"` // A brief description of the video.
	Metadata []Metadata `json:"metadata,omitempty"` // A list (array) of dictionaries where each dictionary contains a key value pair that describes the video. As with tags, you must send the complete list of metadata you want as whatever you send here will overwrite the existing metadata for the video. [Dynamic Metadata](https://api.video/blog/endpoints/dynamic-metadata) allows you to define a key that allows any value pair.
	Mp4support bool `json:"mp4Support,omitempty"` // Whether the player supports the mp4 format.
}

// VideoAssets represents the VideoAssets schema from the OpenAPI specification
type VideoAssets struct {
	Player string `json:"player,omitempty"` // Raw url of the player.
	Thumbnail string `json:"thumbnail,omitempty"` // Poster of the video.
	Hls string `json:"hls,omitempty"` // This is the manifest URL. For HTTP Live Streaming (HLS), when a HLS video stream is initiated, the first file to download is the manifest. This file has the extension M3U8, and provides the video player with information about the various bitrates available for streaming.
	Iframe string `json:"iframe,omitempty"` // Code to use video from a third party website
	Mp4 string `json:"mp4,omitempty"` // Available only if mp4Support is enabled. Raw mp4 url.
}

// Livestreamsessionreferrer represents the Livestreamsessionreferrer schema from the OpenAPI specification
type Livestreamsessionreferrer struct {
	Medium string `json:"medium,omitempty"` // The type of search that brought the viewer to the live stream. Organic would be they found it on their own, paid would be they found it via an advertisement.
	Searchterm string `json:"searchTerm,omitempty"` // What term they searched for that led them to the live stream.
	Source string `json:"source,omitempty"` // Where the viewer came from to see the live stream (usually where they searched from).
	Url string `json:"url,omitempty"` // The website the viewer of the live stream was referred to in order to view the live stream.
}

// Bytesrange represents the Bytesrange schema from the OpenAPI specification
type Bytesrange struct {
	From int `json:"from,omitempty"` // The starting point for the range of bytes for a chunk of a video.
	To int `json:"to,omitempty"` // The ending point for the range of bytes for a chunk of a video.
	Total int `json:"total,omitempty"` // The total number of bytes in the provided range.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	File string `json:"file"` // The image to be added as a thumbnail.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Apikey string `json:"apiKey"` // Your account API key. You can use your sandbox API key, or you can use your production API key.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	File string `json:"file"` // The VTT file describing the chapters you want to upload.
}

// Videosessiondevice represents the Videosessiondevice schema from the OpenAPI specification
type Videosessiondevice struct {
	Vendor string `json:"vendor,omitempty"` // If known, what the brand of the device is, like Apple, Dell, etc.
	Model string `json:"model,omitempty"` // The specific model of the device, if known.
	TypeField string `json:"type,omitempty"` // What the type is like desktop, laptop, mobile.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Data []GeneratedType `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// Livestreamassets represents the Livestreamassets schema from the OpenAPI specification
type Livestreamassets struct {
	Hls string `json:"hls,omitempty"` // The http live streaming (HLS) link for your live video stream.
	Iframe string `json:"iframe,omitempty"` // The embed code for the iframe containing your live video stream.
	Player string `json:"player,omitempty"` // A link to the video player that is playing your live stream.
	Thumbnail string `json:"thumbnail,omitempty"` // A link to the thumbnail for your video.
}

// PlayerCreationPayload represents the PlayerCreationPayload schema from the OpenAPI specification
type PlayerCreationPayload struct {
	Trackbackground string `json:"trackBackground,omitempty"` // RGBA color playback bar: background. Default: rgba(255, 255, 255, .2)
	Backgroundbottom string `json:"backgroundBottom,omitempty"` // RGBA color: bottom 50% of background. Default: rgba(0, 0, 0, .7)
	Backgroundtext string `json:"backgroundText,omitempty"` // RGBA color for title text. Default: rgba(255, 255, 255, 1)
	Forceautoplay bool `json:"forceAutoplay,omitempty"` // enable/disable player autoplay. Default: false
	Link string `json:"link,omitempty"` // RGBA color for all controls. Default: rgba(255, 255, 255, 1)
	Text string `json:"text,omitempty"` // RGBA color for timer text. Default: rgba(255, 255, 255, 1)
	Backgroundtop string `json:"backgroundTop,omitempty"` // RGBA color: top 50% of background. Default: rgba(0, 0, 0, .7)
	Enableapi bool `json:"enableApi,omitempty"` // enable/disable player SDK access. Default: true
	Enablecontrols bool `json:"enableControls,omitempty"` // enable/disable player controls. Default: true
	Forceloop bool `json:"forceLoop,omitempty"` // enable/disable looping. Default: false
	Hidetitle bool `json:"hideTitle,omitempty"` // enable/disable title. Default: false
	Linkhover string `json:"linkHover,omitempty"` // RGBA color for all controls when hovered. Default: rgba(255, 255, 255, 1)
	Trackunplayed string `json:"trackUnplayed,omitempty"` // RGBA color playback bar: downloaded but unplayed (buffered) content. Default: rgba(255, 255, 255, .35)
	Trackplayed string `json:"trackPlayed,omitempty"` // RGBA color playback bar: played content. Default: rgba(88, 131, 255, .95)
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Access_token string `json:"access_token,omitempty"` // The access token containing security credentials allowing you to acccess the API. The token lasts for one hour.
	Expires_in int `json:"expires_in,omitempty"` // Lists the time in seconds when your access token expires. It lasts for one hour.
	Refresh_token string `json:"refresh_token,omitempty"` // A token you can use to get the next access token when your current access token expires.
	Token_type string `json:"token_type,omitempty"` // The type of token you have.
}

// Subtitle represents the Subtitle schema from the OpenAPI specification
type Subtitle struct {
	DefaultField bool `json:"default,omitempty"` // Whether you will have subtitles or not. True for yes you will have subtitles, false for no you will not have subtitles.
	Src string `json:"src,omitempty"`
	Srclang string `json:"srclang,omitempty"`
	Uri string `json:"uri,omitempty"`
}

// Paginationlink represents the Paginationlink schema from the OpenAPI specification
type Paginationlink struct {
	Uri string `json:"uri,omitempty"`
	Rel string `json:"rel,omitempty"`
}

// Player represents the Player schema from the OpenAPI specification
type Player struct {
	Trackunplayed string `json:"trackUnplayed,omitempty"` // RGBA color playback bar: downloaded but unplayed (buffered) content. Default: rgba(255, 255, 255, .35)
	Trackplayed string `json:"trackPlayed,omitempty"` // RGBA color playback bar: played content. Default: rgba(88, 131, 255, .95)
	Trackbackground string `json:"trackBackground,omitempty"` // RGBA color playback bar: background. Default: rgba(255, 255, 255, .2)
	Backgroundbottom string `json:"backgroundBottom,omitempty"` // RGBA color: bottom 50% of background. Default: rgba(0, 0, 0, .7)
	Backgroundtext string `json:"backgroundText,omitempty"` // RGBA color for title text. Default: rgba(255, 255, 255, 1)
	Forceautoplay bool `json:"forceAutoplay,omitempty"` // enable/disable player autoplay. Default: false
	Link string `json:"link,omitempty"` // RGBA color for all controls. Default: rgba(255, 255, 255, 1)
	Text string `json:"text,omitempty"` // RGBA color for timer text. Default: rgba(255, 255, 255, 1)
	Backgroundtop string `json:"backgroundTop,omitempty"` // RGBA color: top 50% of background. Default: rgba(0, 0, 0, .7)
	Enableapi bool `json:"enableApi,omitempty"` // enable/disable player SDK access. Default: true
	Enablecontrols bool `json:"enableControls,omitempty"` // enable/disable player controls. Default: true
	Forceloop bool `json:"forceLoop,omitempty"` // enable/disable looping. Default: false
	Hidetitle bool `json:"hideTitle,omitempty"` // enable/disable title. Default: false
	Linkhover string `json:"linkHover,omitempty"` // RGBA color for all controls when hovered. Default: rgba(255, 255, 255, 1)
	Shapemargin int `json:"shapeMargin,omitempty"` // Deprecated
	Updatedat string `json:"updatedAt,omitempty"` // When the player was last updated, presented in ISO-8601 format.
	Assets map[string]interface{} `json:"assets,omitempty"`
	Createdat string `json:"createdAt,omitempty"` // When the player was created, presented in ISO-8601 format.
	Linkactive string `json:"linkActive,omitempty"` // Deprecated
	Shapebackgroundbottom string `json:"shapeBackgroundBottom,omitempty"` // Deprecated
	Shapebackgroundtop string `json:"shapeBackgroundTop,omitempty"` // Deprecated
	Shaperadius int `json:"shapeRadius,omitempty"` // Deprecated
	Playerid string `json:"playerId,omitempty"`
	Shapeaspect string `json:"shapeAspect,omitempty"` // Deprecated
}

// Playerinput represents the Playerinput schema from the OpenAPI specification
type Playerinput struct {
	Backgroundbottom string `json:"backgroundBottom,omitempty"` // RGBA color: bottom 50% of background. Default: rgba(0, 0, 0, .7)
	Backgroundtext string `json:"backgroundText,omitempty"` // RGBA color for title text. Default: rgba(255, 255, 255, 1)
	Forceautoplay bool `json:"forceAutoplay,omitempty"` // enable/disable player autoplay. Default: false
	Link string `json:"link,omitempty"` // RGBA color for all controls. Default: rgba(255, 255, 255, 1)
	Text string `json:"text,omitempty"` // RGBA color for timer text. Default: rgba(255, 255, 255, 1)
	Backgroundtop string `json:"backgroundTop,omitempty"` // RGBA color: top 50% of background. Default: rgba(0, 0, 0, .7)
	Enableapi bool `json:"enableApi,omitempty"` // enable/disable player SDK access. Default: true
	Enablecontrols bool `json:"enableControls,omitempty"` // enable/disable player controls. Default: true
	Forceloop bool `json:"forceLoop,omitempty"` // enable/disable looping. Default: false
	Hidetitle bool `json:"hideTitle,omitempty"` // enable/disable title. Default: false
	Linkhover string `json:"linkHover,omitempty"` // RGBA color for all controls when hovered. Default: rgba(255, 255, 255, 1)
	Trackunplayed string `json:"trackUnplayed,omitempty"` // RGBA color playback bar: downloaded but unplayed (buffered) content. Default: rgba(255, 255, 255, .35)
	Trackplayed string `json:"trackPlayed,omitempty"` // RGBA color playback bar: played content. Default: rgba(88, 131, 255, .95)
	Trackbackground string `json:"trackBackground,omitempty"` // RGBA color playback bar: background. Default: rgba(255, 255, 255, .2)
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	File string `json:"file"` // The path to the video you want to upload.
	Videoid string `json:"videoId,omitempty"` // The video id returned by the first call to this endpoint in a large video upload scenario.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	DefaultField bool `json:"default,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Referrer Livestreamsessionreferrer `json:"referrer,omitempty"`
	Session Livestreamsessionsession `json:"session,omitempty"`
	Client Livestreamsessionclient `json:"client,omitempty"` // What kind of browser the viewer is using for the live stream session.
	Device Livestreamsessiondevice `json:"device,omitempty"` // What type of device the user is on when in the live stream session.
	Location Livestreamsessionlocation `json:"location,omitempty"` // The location of the viewer of the live stream.
	Os Videosessionos `json:"os,omitempty"` // The operating system the viewer is on.
}

// Videosessionos represents the Videosessionos schema from the OpenAPI specification
type Videosessionos struct {
	Version string `json:"version,omitempty"` // The version of the operating system.
	Name string `json:"name,omitempty"` // The name of the operating system.
	Shortname string `json:"shortname,omitempty"` // The nickname for the operating system, often representing the version.
}

// Videosessionreferrer represents the Videosessionreferrer schema from the OpenAPI specification
type Videosessionreferrer struct {
	Url string `json:"url,omitempty"` // The link the viewer used to reach the video session.
	Medium string `json:"medium,omitempty"` // How they arrived at the site, for example organic or paid. Organic meaning they found it themselves and paid meaning they followed a link from an advertisement.
	Searchterm string `json:"searchTerm,omitempty"` // The search term they typed to arrive at the video session.
	Source string `json:"source,omitempty"` // The source the referrer came from to the video session. For example if they searched through google to find the stream.
}

// Livestreamsessionlocation represents the Livestreamsessionlocation schema from the OpenAPI specification
type Livestreamsessionlocation struct {
	City string `json:"city,omitempty"` // The city of the viewer of the live stream.
	Country string `json:"country,omitempty"` // The country of the viewer of the live stream.
}

// Livestreamsessionsession represents the Livestreamsessionsession schema from the OpenAPI specification
type Livestreamsessionsession struct {
	Endedat string `json:"endedAt,omitempty"` // When the session ended, with the date and time presented in ISO-8601 format.
	Loadedat string `json:"loadedAt,omitempty"` // When the session started, with the date and time presented in ISO-8601 format.
	Sessionid string `json:"sessionId,omitempty"` // A unique identifier for your session. You can use this to track what happens during a specific session.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	File string `json:"file"` // The video text track (VTT) you want to upload.
}

// PlayerUpdatePayload represents the PlayerUpdatePayload schema from the OpenAPI specification
type PlayerUpdatePayload struct {
	Trackbackground string `json:"trackBackground,omitempty"` // RGBA color playback bar: background. Default: rgba(255, 255, 255, .2)
	Backgroundbottom string `json:"backgroundBottom,omitempty"` // RGBA color: bottom 50% of background. Default: rgba(0, 0, 0, .7)
	Backgroundtext string `json:"backgroundText,omitempty"` // RGBA color for title text. Default: rgba(255, 255, 255, 1)
	Forceautoplay bool `json:"forceAutoplay,omitempty"` // enable/disable player autoplay. Default: false
	Link string `json:"link,omitempty"` // RGBA color for all controls. Default: rgba(255, 255, 255, 1)
	Text string `json:"text,omitempty"` // RGBA color for timer text. Default: rgba(255, 255, 255, 1)
	Backgroundtop string `json:"backgroundTop,omitempty"` // RGBA color: top 50% of background. Default: rgba(0, 0, 0, .7)
	Enableapi bool `json:"enableApi,omitempty"` // enable/disable player SDK access. Default: true
	Enablecontrols bool `json:"enableControls,omitempty"` // enable/disable player controls. Default: true
	Forceloop bool `json:"forceLoop,omitempty"` // enable/disable looping. Default: false
	Hidetitle bool `json:"hideTitle,omitempty"` // enable/disable title. Default: false
	Linkhover string `json:"linkHover,omitempty"` // RGBA color for all controls when hovered. Default: rgba(255, 255, 255, 1)
	Trackunplayed string `json:"trackUnplayed,omitempty"` // RGBA color playback bar: downloaded but unplayed (buffered) content. Default: rgba(255, 255, 255, .35)
	Trackplayed string `json:"trackPlayed,omitempty"` // RGBA color playback bar: played content. Default: rgba(88, 131, 255, .95)
}

// Videosessionclient represents the Videosessionclient schema from the OpenAPI specification
type Videosessionclient struct {
	TypeField string `json:"type,omitempty"` // The type of client used to view the video session.
	Version string `json:"version,omitempty"` // The version of the browser used to view the video session.
	Name string `json:"name,omitempty"` // The name of the browser used to view the video session.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Refreshtoken string `json:"refreshToken"` // The refresh token is either the first refresh token you received when you authenticated with the auth/api-key endpoint, or it's the refresh token from the last time you used the auth/refresh endpoint. Place this in the body of your request to obtain a new access token (which is valid for an hour) and a new refresh token.
}

// Videostatusencoding represents the Videostatusencoding schema from the OpenAPI specification
type Videostatusencoding struct {
	Qualities []Quality `json:"qualities,omitempty"` // Available qualities the video can be viewed in.
	Metadata Videostatusencodingmetadata `json:"metadata,omitempty"`
	Playable bool `json:"playable,omitempty"` // Whether the video is playable or not.
}

// Videosessionlocation represents the Videosessionlocation schema from the OpenAPI specification
type Videosessionlocation struct {
	City string `json:"city,omitempty"` // The city of the viewer.
	Country string `json:"country,omitempty"` // The country of the viewer.
}

// Account represents the Account schema from the OpenAPI specification
type Account struct {
	Environment string `json:"environment,omitempty"` // Deprecated. Whether you are using your production or sandbox API key will impact what environment is displayed here, as well as stats and features information. If you use your sandbox key, the environment is "sandbox." If you use your production key, the environment is "production."
	Features []string `json:"features,omitempty"` // Deprecated. What features are enabled for your account. Choices include: app.dynamic_metadata - the ability to dynamically tag videos to better segment and understand your audiences, app.event_log - the ability to create and retrieve a log detailing how your videos were interacted with, player.white_label - the ability to customise your player, stats.player_events - the ability to see statistics about how your player is being used, transcode.mp4_support - the ability to reformat content into mp4 using the H264 codec.
	Quota map[string]interface{} `json:"quota,omitempty"` // Deprecated
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Createdat string `json:"createdAt,omitempty"` // When the token was created, displayed in ISO-8601 format.
	Expiresat string `json:"expiresAt,omitempty"` // When the token expires, displayed in ISO-8601 format.
	Token string `json:"token,omitempty"` // The unique identifier for the token you will use to authenticate an upload.
	Ttl int `json:"ttl,omitempty"` // Time-to-live - how long the upload token is valid for.
}

// Chapter represents the Chapter schema from the OpenAPI specification
type Chapter struct {
	Language string `json:"language,omitempty"`
	Src string `json:"src,omitempty"` // The link to your VTT file, which contains your chapters information for the video.
	Uri string `json:"uri,omitempty"`
}

// Video represents the Video schema from the OpenAPI specification
type Video struct {
	Source VideoSource `json:"source,omitempty"` // Source information about the video.
	Videoid string `json:"videoId,omitempty"` // The unique identifier of the video object.
	Assets VideoAssets `json:"assets,omitempty"` // Collection of details about the video object that you can use to work with the video object.
	Panoramic bool `json:"panoramic,omitempty"` // Defines if video is panoramic.
	Tags []interface{} `json:"tags,omitempty"` // One array of tags (each tag is a string) in order to categorize a video. Tags may include spaces.
	Description string `json:"description,omitempty"` // A description for the video content.
	Playerid string `json:"playerId,omitempty"` // The id of the player that will be applied on the video.
	Title string `json:"title,omitempty"` // The title of the video content.
	Updatedat string `json:"updatedAt,omitempty"` // The date and time the video was updated. Date and time are provided using ISO-8601 UTC format.
	Metadata []Metadata `json:"metadata,omitempty"` // Metadata you can use to categorise and filter videos. Metadata is a list of dictionaries, where each dictionary represents a key value pair for categorising a video. [Dynamic Metadata](https://api.video/blog/endpoints/dynamic-metadata) allows you to define a key that allows any value pair.
	Mp4support bool `json:"mp4Support,omitempty"` // This lets you know whether mp4 is supported. If enabled, an mp4 URL will be provided in the response for the video.
	Public bool `json:"public,omitempty"` // Defines if the content is publicly reachable or if a unique token is needed for each play session. Default is true. Tutorials on [private videos](https://api.video/blog/endpoints/private-videos).
	Publishedat string `json:"publishedAt,omitempty"` // The date and time the API created the video. Date and time are provided using ISO-8601 UTC format.
}

// Videosessionsession represents the Videosessionsession schema from the OpenAPI specification
type Videosessionsession struct {
	Endedat string `json:"endedAt,omitempty"` // When the video session ended, presented in ISO-8601 format.
	Loadedat string `json:"loadedAt,omitempty"` // When the video session started, presented in ISO-8601 format.
	Sessionid string `json:"sessionId,omitempty"` // The unique identifier for the session that you can use to track what happens during it.
}

// Videostatus represents the Videostatus schema from the OpenAPI specification
type Videostatus struct {
	Encoding Videostatusencoding `json:"encoding,omitempty"`
	Ingest Videostatusingest `json:"ingest,omitempty"` // Details about the capturing, transferring, and storing of your video for use immediately or in the future.
}

// Link represents the Link schema from the OpenAPI specification
type Link struct {
	Rel string `json:"rel,omitempty"`
	Uri string `json:"uri,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Data []Video `json:"data"`
	Pagination Pagination `json:"pagination"`
}

// Videostatusingest represents the Videostatusingest schema from the OpenAPI specification
type Videostatusingest struct {
	Filesize int `json:"filesize,omitempty"` // The size of your file in bytes.
	Receivedbytes []Bytesrange `json:"receivedBytes,omitempty"` // The total number of bytes received, listed for each chunk of the upload.
	Status string `json:"status,omitempty"` // There are three possible ingest statuses. missing - you are missing information required to ingest the video. uploading - the video is in the process of being uploaded. uploaded - the video is ready for use.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Timecode string `json:"timecode"` // Frame in video to be used as a placeholder before the video plays. Example: '"00:01:00.000" for 1 minute into the video.' Valid Patterns: "hh:mm:ss.ms" "hh:mm:ss:frameNumber" "124" (integer value is reported as seconds) If selection is out of range, "00:00:00.00" will be chosen.
}

// Videosourcelivestreamlink represents the Videosourcelivestreamlink schema from the OpenAPI specification
type Videosourcelivestreamlink struct {
	Rel string `json:"rel,omitempty"`
	Uri string `json:"uri,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Data []Webhook `json:"data,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	From int `json:"from,omitempty"`
	To int `json:"to,omitempty"`
	TypeField string `json:"type,omitempty"` // Possible values are: ready, play, pause, resume, seek.backward, seek.forward, end
	At int `json:"at,omitempty"`
	Emittedat string `json:"emittedAt,omitempty"` // When an event occurred, presented in ISO-8601 format.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Ttl int `json:"ttl,omitempty"` // Time in seconds that the token will be active. A value of 0 means that the token has no exipration date. The default is to have no expiration.
}

// Metadata represents the Metadata schema from the OpenAPI specification
type Metadata struct {
	Key string `json:"key,omitempty"` // The constant that defines the data set.
	Value string `json:"value,omitempty"` // A variable which belongs to the data set.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Streamkey string `json:"streamKey,omitempty"` // The unique, private stream key that you use to begin streaming.
	Assets Livestreamassets `json:"assets,omitempty"`
	Broadcasting bool `json:"broadcasting,omitempty"` // Whether or not you are broadcasting the live video you recorded for others to see. True means you are broadcasting to viewers, false means you are not.
	Livestreamid string `json:"liveStreamId,omitempty"` // The unique identifier for the live stream. Live stream IDs begin with "li."
	Name string `json:"name,omitempty"` // The name of your live stream.
	Playerid string `json:"playerId,omitempty"` // The unique identifier for the player.
	Public bool `json:"public,omitempty"` // BETA FEATURE Please limit all public = false ("private") livestreams to 3,000 users. Whether your video can be viewed by everyone, or requires authentication to see it. A setting of false will require a unique token for each view.
	Record bool `json:"record,omitempty"` // Whether you are recording or not.
}

// Livestreamsessionclient represents the Livestreamsessionclient schema from the OpenAPI specification
type Livestreamsessionclient struct {
	Name string `json:"name,omitempty"` // The name of the browser used to view the live stream session.
	TypeField string `json:"type,omitempty"` // The type of client used to view the live stream session.
	Version string `json:"version,omitempty"` // The version of the browser used to view the live stream session.
}

// Pagination represents the Pagination schema from the OpenAPI specification
type Pagination struct {
	Currentpageitems int `json:"currentPageItems,omitempty"` // The number of items on the current page.
	Itemstotal int `json:"itemsTotal,omitempty"` // Total number of items that exist.
	Links []Paginationlink `json:"links"`
	Pagesize int `json:"pageSize,omitempty"` // Maximum number of item per page.
	Pagestotal int `json:"pagesTotal,omitempty"` // Number of items listed in the current page.
	Currentpage int `json:"currentPage,omitempty"` // The current page index.
}

// Quality represents the Quality schema from the OpenAPI specification
type Quality struct {
	Quality string `json:"quality,omitempty"` // The quality of the video you have, in pixels. Choices include 360p, 480p, 720p, 1080p, and 2160p.
	Status string `json:"status,omitempty"` // The status of your video. Statuses include waiting - the video is waiting to be encoded. encoding - the video is in the process of being encoded. encoded - the video was successfully encoded. failed - the video failed to be encoded.
}

// Videostatusencodingmetadata represents the Videostatusencodingmetadata schema from the OpenAPI specification
type Videostatusencodingmetadata struct {
	Width int `json:"width,omitempty"` // The width of the video in pixels.
	Aspectratio string `json:"aspectRatio,omitempty"`
	Duration int `json:"duration,omitempty"` // The length of the video.
	Framerate int `json:"framerate,omitempty"` // The frequency with which consecutive images or frames appear on a display. Shown in this API as frames per second (fps).
	Height int `json:"height,omitempty"` // The height of the video in pixels.
	Samplerate int `json:"samplerate,omitempty"` // How many samples per second a digital audio system uses to record an audio signal. The higher the rate, the higher the frequencies that can be recorded. They are presented in this API using hertz.
	Videocodec string `json:"videoCodec,omitempty"` // The method used to compress and decompress digital video. API Video supports all codecs in the libavcodec library.
	Audiocodec string `json:"audioCodec,omitempty"` // The method used to compress and decompress digital audio for your video.
	Bitrate float64 `json:"bitrate,omitempty"` // The number of bits processed per second.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Problems []GeneratedType `json:"problems,omitempty"`
	Status int `json:"status,omitempty"`
	Title string `json:"title,omitempty"`
	TypeField string `json:"type,omitempty"`
	Name string `json:"name,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Data []Chapter `json:"data,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Source string `json:"source,omitempty"` // If you add a video already on the web, this is where you enter the url for the video.
	Mp4support bool `json:"mp4Support,omitempty"` // Enables mp4 version in addition to streamed version.
	Title string `json:"title"` // The title of your new video.
	Metadata []Metadata `json:"metadata,omitempty"` // A list of key value pairs that you use to provide metadata for your video. These pairs can be made dynamic, allowing you to segment your audience. Read more on [dynamic metadata](https://api.video/blog/endpoints/dynamic-metadata).
	Playerid string `json:"playerId,omitempty"` // The unique identification number for your video player.
	Tags []string `json:"tags,omitempty"` // A list of tags you want to use to describe your video.
	Description string `json:"description,omitempty"` // A brief description of your video.
	Panoramic bool `json:"panoramic,omitempty"` // Indicates if your video is a 360/immersive video.
	Public bool `json:"public,omitempty"` // Whether your video can be viewed by everyone, or requires authentication to see it. A setting of false will require a unique token for each view. Default is true. Tutorials on [private videos](https://api.video/blog/endpoints/private-videos).
	Publishedat string `json:"publishedAt,omitempty"` // The API uses ISO-8601 format for time, and includes 3 places for milliseconds.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Record bool `json:"record,omitempty"` // Whether you are recording or not. True for record, false for not record.
	Name string `json:"name"` // Add a name for your live stream here.
	Playerid string `json:"playerId,omitempty"` // The unique identifier for the player.
	Public bool `json:"public,omitempty"` // BETA FEATURE Please limit all public = false ("private") livestreams to 3,000 users. Whether your video can be viewed by everyone, or requires authentication to see it. A setting of false will require a unique token for each view.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Data []Player `json:"data,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	File string `json:"file"` // The image to be added as a thumbnail.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Record bool `json:"record,omitempty"` // Use this to indicate whether you want the recording on or off. On is true, off is false.
	Name string `json:"name,omitempty"` // The name you want to use for your live stream.
	Playerid string `json:"playerId,omitempty"` // The unique ID for the player associated with a live stream that you want to update.
	Public bool `json:"public,omitempty"` // BETA FEATURE Please limit all public = false ("private") livestreams to 3,000 users. Whether your video can be viewed by everyone, or requires authentication to see it. A setting of false will require a unique token for each view.
}

// Livestreamsessiondevice represents the Livestreamsessiondevice schema from the OpenAPI specification
type Livestreamsessiondevice struct {
	Model string `json:"model,omitempty"` // The specific model of the device, if known.
	TypeField string `json:"type,omitempty"` // What the type is like desktop, laptop, mobile.
	Vendor string `json:"vendor,omitempty"` // If known, what the brand of the device is, like Apple, Dell, etc.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Client Videosessionclient `json:"client,omitempty"` // What kind of browser the viewer is using for the video session.
	Device Videosessiondevice `json:"device,omitempty"` // What type of device the user is on when in the video session.
	Location Videosessionlocation `json:"location,omitempty"` // The location of the viewer.
	Os Videosessionos `json:"os,omitempty"` // The operating system the viewer is on.
	Referrer Videosessionreferrer `json:"referrer,omitempty"`
	Session Videosessionsession `json:"session,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Data []GeneratedType `json:"data,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Data []GeneratedType `json:"data,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	File string `json:"file"` // The path to the video you would like to upload. The path must be local. If you want to use a video from an online source, you must use the "/videos" endpoint and add the "source" parameter when you create a new video.
}

// Webhook represents the Webhook schema from the OpenAPI specification
type Webhook struct {
	Events []string `json:"events,omitempty"` // A list of events that will trigger the webhook.
	Url string `json:"url,omitempty"` // URL of the webhook
	Webhookid string `json:"webhookId,omitempty"` // Unique identifier of the webhook
	Createdat string `json:"createdAt,omitempty"` // When an webhook was created, presented in ISO-8601 format.
}

// VideoSource represents the VideoSource schema from the OpenAPI specification
type VideoSource struct {
	Livestream Videosourcelivestream `json:"liveStream,omitempty"` // This appears if the video is from a Live Record.
	TypeField string `json:"type,omitempty"`
	Uri string `json:"uri,omitempty"` // The URL where the video is stored.
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Data []Subtitle `json:"data,omitempty"`
	Pagination Pagination `json:"pagination,omitempty"`
}

// GeneratedType represents the GeneratedType schema from the OpenAPI specification
type GeneratedType struct {
	Events []string `json:"events"` // A list of the webhooks that you are subscribing to. There are Currently four webhook options: * ```video.encoding.quality.completed``` When a new video is uploaded into your account, it will be encoded into several different HLS sizes/bitrates. When each version is encoded, your webhook will get a notification. It will look like ```{ \"type\": \"video.encoding.quality.completed\", \"emittedAt\": \"2021-01-29T16:46:25.217+01:00\", \"videoId\": \"viXXXXXXXX\", \"encoding\": \"hls\", \"quality\": \"720p\"} ```. This request says that the 720p HLS encoding was completed. * ```live-stream.broadcast.started``` When a livestream begins broadcasting, the broadcasting parameter changes from false to true, and this webhook fires. * ```live-stream.broadcast.ended``` This event fores when the livestream has finished broadcasting, and the broadcasting parameter flips from false to true. * ```video.source.recorded``` This event is similar to ```video.encoding.quality.completed```, but tells you if a livestream has been recorded as a VOD.
	Url string `json:"url"` // The the url to which HTTP notifications are sent. It could be any http or https URL.
}
