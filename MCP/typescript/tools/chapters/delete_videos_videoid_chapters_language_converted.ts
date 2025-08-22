/**
 * MCP Server function for Delete a chapter
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

function getDelete_Videos_Video_Id_Chapters_LanguageHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
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
    if (args.language !== undefined) {
        queryParams.push(`language=${args.language}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/delete_videos_video_id_chapters_language${queryString}`;
            
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

function createDelete_Videos_Video_Id_Chapters_LanguageTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "delete_videos_video_id_chapters_language",
            description: "Delete a chapter",
            parameters: {
        videoId: { type: "string", required: true, description: "The unique identifier for the video you want to delete a chapter from." },
        language: { type: "string", required: true, description: "A valid [BCP 47](https://github.com/libyal/libfwnt/wiki/Language-Code-identifiers) language representation." },
            }
        },
        handler: getDelete_Videos_Video_Id_Chapters_LanguageHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getDelete_Videos_Video_Id_Chapters_LanguageHandler,
    createDelete_Videos_Video_Id_Chapters_LanguageTool
};