/**
 * MCP Server function for Create a video
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

function getPost_VideosHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.description !== undefined) {
        queryParams.push(`description=${args.description}`);
    }
    if (args.publishedAt !== undefined) {
        queryParams.push(`publishedAt=${args.publishedAt}`);
    }
    if (args.source !== undefined) {
        queryParams.push(`source=${args.source}`);
    }
    if (args.title !== undefined) {
        queryParams.push(`title=${args.title}`);
    }
    if (args.playerId !== undefined) {
        queryParams.push(`playerId=${args.playerId}`);
    }
    if (args.panoramic !== undefined) {
        queryParams.push(`panoramic=${args.panoramic}`);
    }
    if (args.public !== undefined) {
        queryParams.push(`public=${args.public}`);
    }
    if (args.mp4Support !== undefined) {
        queryParams.push(`mp4Support=${args.mp4Support}`);
    }
    if (args.tags !== undefined) {
        queryParams.push(`tags=${args.tags}`);
    }
    if (args.metadata !== undefined) {
        queryParams.push(`metadata=${args.metadata}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/post_videos${queryString}`;
            
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

function createPost_VideosTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "post_videos",
            description: "Create a video",
            parameters: {
        description: { type: "string", required: false, description: "Input parameter: A brief description of your video." },
        publishedAt: { type: "string", required: false, description: "Input parameter: The API uses ISO-8601 format for time, and includes 3 places for milliseconds." },
        source: { type: "string", required: false, description: "Input parameter: If you add a video already on the web, this is where you enter the url for the video." },
        title: { type: "string", required: true, description: "Input parameter: The title of your new video." },
        playerId: { type: "string", required: false, description: "Input parameter: The unique identification number for your video player." },
        panoramic: { type: "string", required: false, description: "Input parameter: Indicates if your video is a 360/immersive video." },
        public: { type: "string", required: false, description: "Input parameter: Whether your video can be viewed by everyone, or requires authentication to see it. A setting of false will require a unique token for each view. Default is true. Tutorials on [private videos](https://api.video/blog/endpoints/private-videos)." },
        mp4Support: { type: "string", required: false, description: "Input parameter: Enables mp4 version in addition to streamed version." },
        tags: { type: "string", required: false, description: "Input parameter: A list of tags you want to use to describe your video." },
        metadata: { type: "string", required: false, description: "Input parameter: A list of key value pairs that you use to provide metadata for your video. These pairs can be made dynamic, allowing you to segment your audience. Read more on [dynamic metadata](https://api.video/blog/endpoints/dynamic-metadata)." },
            }
        },
        handler: getPost_VideosHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPost_VideosHandler,
    createPost_VideosTool
};