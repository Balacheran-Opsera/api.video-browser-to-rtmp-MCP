/**
 * MCP Server function for Update a live stream
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

function getPatch_Live_Streams_Live_Stream_IdHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.liveStreamId !== undefined) {
        queryParams.push(`liveStreamId=${args.liveStreamId}`);
    }
    if (args.name !== undefined) {
        queryParams.push(`name=${args.name}`);
    }
    if (args.playerId !== undefined) {
        queryParams.push(`playerId=${args.playerId}`);
    }
    if (args.public !== undefined) {
        queryParams.push(`public=${args.public}`);
    }
    if (args.record !== undefined) {
        queryParams.push(`record=${args.record}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/patch_live_streams_live_stream_id${queryString}`;
            
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

function createPatch_Live_Streams_Live_Stream_IdTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "patch_live_streams_live_stream_id",
            description: "Update a live stream",
            parameters: {
        liveStreamId: { type: "string", required: true, description: "The unique ID for the live stream that you want to update information for such as player details, or whether you want the recording on or off." },
        name: { type: "string", required: false, description: "Input parameter: The name you want to use for your live stream." },
        playerId: { type: "string", required: false, description: "Input parameter: The unique ID for the player associated with a live stream that you want to update." },
        public: { type: "string", required: false, description: "Input parameter: BETA FEATURE Please limit all public = false ("private\" },
        record: { type: "string", required: false, description: "Input parameter: Use this to indicate whether you want the recording on or off. On is true, off is false." },
            }
        },
        handler: getPatch_Live_Streams_Live_Stream_IdHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPatch_Live_Streams_Live_Stream_IdHandler,
    createPatch_Live_Streams_Live_Stream_IdTool
};