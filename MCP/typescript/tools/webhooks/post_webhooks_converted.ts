/**
 * MCP Server function for Create Webhook
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

function getPost_WebhooksHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.url !== undefined) {
        queryParams.push(`url=${args.url}`);
    }
    if (args.events !== undefined) {
        queryParams.push(`events=${args.events}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/post_webhooks${queryString}`;
            
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

function createPost_WebhooksTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "post_webhooks",
            description: "Create Webhook",
            parameters: {
        url: { type: "string", required: true, description: "Input parameter: The the url to which HTTP notifications are sent. It could be any http or https URL." },
        events: { type: "string", required: true, description: "Input parameter: A list of the webhooks that you are subscribing to. There are Currently four webhook options: * ```video.encoding.quality.completed``` When a new video is uploaded into your account, it will be encoded into several different HLS sizes/bitrates. When each version is encoded, your webhook will get a notification. It will look like ```{ \\"type\\": \\"video.encoding.quality.completed\\", \\"emittedAt\\": \\"2021-01-29T16:46:25.217+01:00\\", \\"videoId\\": \\"viXXXXXXXX\\", \\"encoding\\": \\"hls\\", \\"quality\\": \\"720p\\"} ```. This request says that the 720p HLS encoding was completed. * ```live-stream.broadcast.started``` When a livestream begins broadcasting, the broadcasting parameter changes from false to true, and this webhook fires. * ```live-stream.broadcast.ended``` This event fores when the livestream has finished broadcasting, and the broadcasting parameter flips from false to true. * ```video.source.recorded``` This event is similar to ```video.encoding.quality.completed```, but tells you if a livestream has been recorded as a VOD." },
            }
        },
        handler: getPost_WebhooksHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPost_WebhooksHandler,
    createPost_WebhooksTool
};