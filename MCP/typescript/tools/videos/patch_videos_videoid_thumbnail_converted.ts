/**
 * MCP Server function for Pick a thumbnail
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

function getPatch_Videos_Video_Id_ThumbnailHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
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
    if (args.timecode !== undefined) {
        queryParams.push(`timecode=${args.timecode}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/patch_videos_video_id_thumbnail${queryString}`;
            
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

function createPatch_Videos_Video_Id_ThumbnailTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "patch_videos_video_id_thumbnail",
            description: "Pick a thumbnail",
            parameters: {
        videoId: { type: "string", required: true, description: "Unique identifier of the video you want to add a thumbnail to, where you use a section of your video as the thumbnail." },
        timecode: { type: "string", required: true, description: "Input parameter: Frame in video to be used as a placeholder before the video plays. Example: '"00:01:00.000" for 1 minute into the video.' Valid Patterns: "hh:mm:ss.ms" "hh:mm:ss:frameNumber" "124" (integer value is reported as seconds) If selection is out of range, "00:00:00.00" will be chosen." },
            }
        },
        handler: getPatch_Videos_Video_Id_ThumbnailHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPatch_Videos_Video_Id_ThumbnailHandler,
    createPatch_Videos_Video_Id_ThumbnailTool
};