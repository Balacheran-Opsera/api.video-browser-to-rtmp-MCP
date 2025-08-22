/**
 * MCP Server function for List all videos
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

function getGet_VideosHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.title !== undefined) {
        queryParams.push(`title=${args.title}`);
    }
    if (args.description !== undefined) {
        queryParams.push(`description=${args.description}`);
    }
    if (args.liveStreamId !== undefined) {
        queryParams.push(`liveStreamId=${args.liveStreamId}`);
    }
    if (args.sortBy !== undefined) {
        queryParams.push(`sortBy=${args.sortBy}`);
    }
    if (args.sortOrder !== undefined) {
        queryParams.push(`sortOrder=${args.sortOrder}`);
    }
    if (args.currentPage !== undefined) {
        queryParams.push(`currentPage=${args.currentPage}`);
    }
    if (args.pageSize !== undefined) {
        queryParams.push(`pageSize=${args.pageSize}`);
    }
    if (args.tags !== undefined) {
        queryParams.push(`tags=${args.tags}`);
    }
    if (args.metadata !== undefined) {
        queryParams.push(`metadata=${args.metadata}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/get_videos${queryString}`;
            
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

function createGet_VideosTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "get_videos",
            description: "List all videos",
            parameters: {
        title: { type: "string", required: false, description: "The title of a specific video you want to find. The search will match exactly to what term you provide and return any videos that contain the same term as part of their titles." },
        description: { type: "string", required: false, description: "If you described a video with a term or sentence, you can add it here to return videos containing this string." },
        liveStreamId: { type: "string", required: false, description: "If you know the ID for a live stream, you can retrieve the stream by adding the ID for it here." },
        sortBy: { type: "string", required: false, description: "Allowed: publishedAt, title. You can search by the time videos were published at, or by title." },
        sortOrder: { type: "string", required: false, description: "Allowed: asc, desc. asc is ascending and sorts from A to Z. desc is descending and sorts from Z to A." },
        currentPage: { type: "string", required: false, description: "Choose the number of search results to return per page. Minimum value: 1" },
        pageSize: { type: "string", required: false, description: "Results per page. Allowed values 1-100, default is 25." },
        tags: { type: "string", required: false, description: "A tag is a category you create and apply to videos. You can search for videos with particular tags by listing one or more here. Only videos that have all the tags you list will be returned." },
        metadata: { type: "string", required: false, description: "Videos can be tagged with metadata tags in key:value pairs. You can search for videos with specific key value pairs using this parameter. [Dynamic Metadata](https://api.video/blog/endpoints/dynamic-metadata) allows you to define a key that allows any value pair." },
            }
        },
        handler: getGet_VideosHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getGet_VideosHandler,
    createGet_VideosTool
};