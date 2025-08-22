/**
 * MCP Server function for List all active upload tokens.
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

function getGet_Upload_TokensHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
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
            const url = `${config.baseUrl}/api/v2/get_upload_tokens${queryString}`;
            
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

function createGet_Upload_TokensTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "get_upload_tokens",
            description: "List all active upload tokens.",
            parameters: {
        sortBy: { type: "string", required: false, description: "Allowed: createdAt, ttl. You can use these to sort by when a token was created, or how much longer the token will be active (ttl - time to live). Date and time is presented in ISO-8601 format." },
        sortOrder: { type: "string", required: false, description: "Allowed: asc, desc. Ascending is 0-9 or A-Z. Descending is 9-0 or Z-A." },
        currentPage: { type: "string", required: false, description: "Choose the number of search results to return per page. Minimum value: 1" },
        pageSize: { type: "string", required: false, description: "Results per page. Allowed values 1-100, default is 25." },
            }
        },
        handler: getGet_Upload_TokensHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getGet_Upload_TokensHandler,
    createGet_Upload_TokensTool
};