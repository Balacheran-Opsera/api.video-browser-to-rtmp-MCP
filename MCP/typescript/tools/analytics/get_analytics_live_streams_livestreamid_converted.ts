/**
 * MCP Server function for List live stream player sessions
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

function getGet_Analytics_Live_Streams_Live_Stream_IdHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
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
    if (args.period !== undefined) {
        queryParams.push(`period=${args.period}`);
    }
    if (args.currentPage !== undefined) {
        queryParams.push(`currentPage=${args.currentPage}`);
    }
    if (args.pageSize !== undefined) {
        queryParams.push(`pageSize=${args.pageSize}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/get_analytics_live_streams_live_stream_id${queryString}`;
            
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

function createGet_Analytics_Live_Streams_Live_Stream_IdTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "get_analytics_live_streams_live_stream_id",
            description: "List live stream player sessions",
            parameters: {
        liveStreamId: { type: "string", required: true, description: "The unique identifier for the live stream you want to retrieve analytics for." },
        period: { type: "string", required: false, description: "Period must have one of the following formats: - For a day : "2018-01-01", - For a week: "2018-W01", - For a month: "2018-01" - For a year: "2018" For a range period: - Date range: "2018-01-01/2018-01-15"" },
        currentPage: { type: "string", required: false, description: "Choose the number of search results to return per page. Minimum value: 1" },
        pageSize: { type: "string", required: false, description: "Results per page. Allowed values 1-100, default is 25." },
            }
        },
        handler: getGet_Analytics_Live_Streams_Live_Stream_IdHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getGet_Analytics_Live_Streams_Live_Stream_IdHandler,
    createGet_Analytics_Live_Streams_Live_Stream_IdTool
};