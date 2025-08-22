/**
 * MCP Server function for Update a video
 */

import axios, { AxiosResponse } from 'axios';

interface APIConfig {
    baseUrl: string;
    apiKey: string;
}

interface MCPRequest {
    params?: {
        arguments?: Record<string, any>;
    };
}

interface MCPToolResult {
    content: string;
    isError: boolean;
}

interface ToolDefinition {
    name: string;
    description: string;
    parameters: Record<string, {
        type: string;
        required: boolean;
        description: string;
    }>;
}

interface Tool {
    definition: ToolDefinition;
    handler: (ctx: any, request: MCPRequest) => Promise<MCPToolResult>;
}

class MCPToolResultImpl implements MCPToolResult {
    constructor(
        public content: string,
        public isError: boolean = false
    ) {}
}

function getPatch_Videos_Video_IdHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.videoId !== undefined) {
        queryParams.push(`videoId=${args.videoId}`);
    }
    if (args.title !== undefined) {
        queryParams.push(`title=${args.title}`);
    }
    if (args.description !== undefined) {
        queryParams.push(`description=${args.description}`);
    }
    if (args.playerId !== undefined) {
        queryParams.push(`playerId=${args.playerId}`);
    }
    if (args.mp4Support !== undefined) {
        queryParams.push(`mp4Support=${args.mp4Support}`);
    }
    if (args.panoramic !== undefined) {
        queryParams.push(`panoramic=${args.panoramic}`);
    }
    if (args.public !== undefined) {
        queryParams.push(`public=${args.public}`);
    }
    if (args.tags !== undefined) {
        queryParams.push(`tags=${args.tags}`);
    }
    if (args.metadata !== undefined) {
        queryParams.push(`metadata=${args.metadata}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/patch_videos_video_id${queryString}`;
            
            const headers = {
                'Authorization': `Bearer ${config.apiKey}`,
                'Accept': 'application/json'
            };
            
            const response: AxiosResponse = await axios.get(url, { headers });
            
            if (response.status >= 400) {
                return new MCPToolResultImpl(`API error: ${response.data}`, true);
            }
            
            const prettyJSON = JSON.stringify(response.data, null, 2);
            return new MCPToolResultImpl(prettyJSON);
            
        } catch (error: any) {
            if (error.response) {
                return new MCPToolResultImpl(`Request failed: ${error.response.data}`, true);
            }
            return new MCPToolResultImpl(`Unexpected error: ${error.message}`, true);
        }
    };
}

function createPatch_Videos_Video_IdTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "patch_videos_video_id",
            description: "Update a video",
            parameters: {
        videoId: { type: "string", required: true, description: "The video ID for the video you want to delete." },
        title: { type: "string", required: false, description: "Input parameter: The title you want to use for your video." },
        description: { type: "string", required: false, description: "Input parameter: A brief description of the video." },
        playerId: { type: "string", required: false, description: "Input parameter: The unique ID for the player you want to associate with your video." },
        mp4Support: { type: "string", required: false, description: "Input parameter: Whether the player supports the mp4 format." },
        panoramic: { type: "string", required: false, description: "Input parameter: Whether the video is a 360 degree or immersive video." },
        public: { type: "string", required: false, description: "Input parameter: Whether the video is publicly available or not. False means it is set to private. Default is true. Tutorials on [private videos](https://api.video/blog/endpoints/private-videos)." },
        tags: { type: "string", required: false, description: "Input parameter: A list of terms or words you want to tag the video with. Make sure the list includes all the tags you want as whatever you send in this list will overwrite the existing list for the video." },
        metadata: { type: "string", required: false, description: "Input parameter: A list (array) of dictionaries where each dictionary contains a key value pair that describes the video. As with tags, you must send the complete list of metadata you want as whatever you send here will overwrite the existing metadata for the video. [Dynamic Metadata](https://api.video/blog/endpoints/dynamic-metadata) allows you to define a key that allows any value pair." },
            }
        },
        handler: getPatch_Videos_Video_IdHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPatch_Videos_Video_IdHandler,
    createPatch_Videos_Video_IdTool
};