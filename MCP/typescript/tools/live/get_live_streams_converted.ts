/**
 * MCP Server function for List all live streams
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

function getGet_Live_StreamsHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.streamKey !== undefined) {
        queryParams.push(`streamKey=${args.streamKey}`);
    }
    if (args.name !== undefined) {
        queryParams.push(`name=${args.name}`);
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
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/get_live_streams${queryString}`;
            
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

function createGet_Live_StreamsTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "get_live_streams",
            description: "List all live streams",
            parameters: {
        streamKey: { type: "string", required: false, description: "The unique stream key that allows you to stream videos." },
        name: { type: "string", required: false, description: "You can filter live streams by their name or a part of their name." },
        sortBy: { type: "string", required: false, description: "Allowed: createdAt, publishedAt, name. createdAt - the time a livestream was created using the specified streamKey. publishedAt - the time a livestream was published using the specified streamKey. name - the name of the livestream. If you choose one of the time based options, the time is presented in ISO-8601 format." },
        sortOrder: { type: "string", required: false, description: "Allowed: asc, desc. Ascending for date and time means that earlier values precede later ones. Descending means that later values preced earlier ones. For title, it is 0-9 and A-Z ascending and Z-A, 9-0 descending." },
        currentPage: { type: "string", required: false, description: "Choose the number of search results to return per page. Minimum value: 1" },
        pageSize: { type: "string", required: false, description: "Results per page. Allowed values 1-100, default is 25." },
            }
        },
        handler: getGet_Live_StreamsHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getGet_Live_StreamsHandler,
    createGet_Live_StreamsTool
};