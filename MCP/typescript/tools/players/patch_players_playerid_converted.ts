/**
 * MCP Server function for Update a player
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

function getPatch_Players_Player_IdHandler(config: APIConfig): (ctx: any, request: MCPRequest) => Promise<MCPToolResult> {
    return async function(ctx: any, request: MCPRequest): Promise<MCPToolResult> {
        try {
            const args = request?.params?.arguments || {};
            if (typeof args !== 'object') {
                return new MCPToolResultImpl("Invalid arguments object", true);
            }
            
            const queryParams: string[] = [];
    if (args.playerId !== undefined) {
        queryParams.push(`playerId=${args.playerId}`);
    }
    if (args.backgroundTop !== undefined) {
        queryParams.push(`backgroundTop=${args.backgroundTop}`);
    }
    if (args.linkHover !== undefined) {
        queryParams.push(`linkHover=${args.linkHover}`);
    }
    if (args.trackUnplayed !== undefined) {
        queryParams.push(`trackUnplayed=${args.trackUnplayed}`);
    }
    if (args.trackPlayed !== undefined) {
        queryParams.push(`trackPlayed=${args.trackPlayed}`);
    }
    if (args.trackBackground !== undefined) {
        queryParams.push(`trackBackground=${args.trackBackground}`);
    }
    if (args.backgroundBottom !== undefined) {
        queryParams.push(`backgroundBottom=${args.backgroundBottom}`);
    }
    if (args.backgroundText !== undefined) {
        queryParams.push(`backgroundText=${args.backgroundText}`);
    }
    if (args.link !== undefined) {
        queryParams.push(`link=${args.link}`);
    }
    if (args.text !== undefined) {
        queryParams.push(`text=${args.text}`);
    }
    if (args.enableApi !== undefined) {
        queryParams.push(`enableApi=${args.enableApi}`);
    }
    if (args.enableControls !== undefined) {
        queryParams.push(`enableControls=${args.enableControls}`);
    }
    if (args.forceLoop !== undefined) {
        queryParams.push(`forceLoop=${args.forceLoop}`);
    }
    if (args.hideTitle !== undefined) {
        queryParams.push(`hideTitle=${args.hideTitle}`);
    }
    if (args.forceAutoplay !== undefined) {
        queryParams.push(`forceAutoplay=${args.forceAutoplay}`);
    }
            
            const queryString = queryParams.length > 0 ? '?' + queryParams.join('&') : '';
            const url = `${config.baseUrl}/api/v2/patch_players_player_id${queryString}`;
            
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

function createPatch_Players_Player_IdTool(config: APIConfig): Tool {
    return {
        definition: {
            name: "patch_players_player_id",
            description: "Update a player",
            parameters: {
        playerId: { type: "string", required: true, description: "The unique identifier for the player." },
        backgroundTop: { type: "string", required: false, description: "Input parameter: RGBA color: top 50% of background. Default: rgba(0, 0, 0, .7)" },
        linkHover: { type: "string", required: false, description: "Input parameter: RGBA color for all controls when hovered. Default: rgba(255, 255, 255, 1)" },
        trackUnplayed: { type: "string", required: false, description: "Input parameter: RGBA color playback bar: downloaded but unplayed (buffered) content. Default: rgba(255, 255, 255, .35)" },
        trackPlayed: { type: "string", required: false, description: "Input parameter: RGBA color playback bar: played content. Default: rgba(88, 131, 255, .95)" },
        trackBackground: { type: "string", required: false, description: "Input parameter: RGBA color playback bar: background. Default: rgba(255, 255, 255, .2)" },
        backgroundBottom: { type: "string", required: false, description: "Input parameter: RGBA color: bottom 50% of background. Default: rgba(0, 0, 0, .7)" },
        backgroundText: { type: "string", required: false, description: "Input parameter: RGBA color for title text. Default: rgba(255, 255, 255, 1)" },
        link: { type: "string", required: false, description: "Input parameter: RGBA color for all controls. Default: rgba(255, 255, 255, 1)" },
        text: { type: "string", required: false, description: "Input parameter: RGBA color for timer text. Default: rgba(255, 255, 255, 1)" },
        enableApi: { type: "string", required: false, description: "Input parameter: enable/disable player SDK access. Default: true" },
        enableControls: { type: "string", required: false, description: "Input parameter: enable/disable player controls. Default: true" },
        forceLoop: { type: "string", required: false, description: "Input parameter: enable/disable looping. Default: false" },
        hideTitle: { type: "string", required: false, description: "Input parameter: enable/disable title. Default: false" },
        forceAutoplay: { type: "string", required: false, description: "Input parameter: enable/disable player autoplay. Default: false" },
            }
        },
        handler: getPatch_Players_Player_IdHandler(config)
    };
}

export {
    APIConfig,
    MCPRequest,
    MCPToolResult,
    MCPToolResultImpl,
    Tool,
    ToolDefinition,
    getPatch_Players_Player_IdHandler,
    createPatch_Players_Player_IdTool
};