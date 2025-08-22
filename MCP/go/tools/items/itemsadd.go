package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/pos-api/mcp-server/config"
	"github.com/pos-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func ItemsaddHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["raw"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("raw=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.Item
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/pos/items%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Fallback to single auth parameter
		if cfg.APIKey != "" {
			req.Header.Set("Authorization", cfg.APIKey)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["x-apideck-consumer-id"]; ok {
			req.Header.Set("x-apideck-consumer-id", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-apideck-app-id"]; ok {
			req.Header.Set("x-apideck-app-id", fmt.Sprintf("%v", val))
		}
		if val, ok := args["x-apideck-service-id"]; ok {
			req.Header.Set("x-apideck-service-id", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.CreateItemResponse
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateItemsaddTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_pos_items",
		mcp.WithDescription("Create Item"),
		mcp.WithBoolean("raw", mcp.Description("Include raw response. Mostly used for debugging purposes")),
		mcp.WithString("x-apideck-consumer-id", mcp.Required(), mcp.Description("ID of the consumer which you want to get or push data from")),
		mcp.WithString("x-apideck-app-id", mcp.Required(), mcp.Description("The ID of your Unify application")),
		mcp.WithString("x-apideck-service-id", mcp.Description("Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.")),
		mcp.WithString("created_at", mcp.Description("Input parameter: The date and time when the object was created.")),
		mcp.WithString("idempotency_key", mcp.Description("Input parameter: A value you specify that uniquely identifies this request among requests you have sent.")),
		mcp.WithString("updated_by", mcp.Description("Input parameter: The user who last updated the object.")),
		mcp.WithArray("variations", mcp.Description("")),
		mcp.WithString("created_by", mcp.Description("Input parameter: The user who created the object.")),
		mcp.WithString("updated_at", mcp.Description("Input parameter: The date and time when the object was last updated.")),
		mcp.WithArray("modifier_groups", mcp.Description("")),
		mcp.WithBoolean("present_at_all_locations", mcp.Description("")),
		mcp.WithString("name", mcp.Required(), mcp.Description("")),
		mcp.WithString("price_currency", mcp.Description("Input parameter: Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).")),
		mcp.WithString("id", mcp.Description("")),
		mcp.WithString("description", mcp.Description("")),
		mcp.WithBoolean("available_online", mcp.Description("")),
		mcp.WithArray("categories", mcp.Description("")),
		mcp.WithBoolean("deleted", mcp.Description("Input parameter: Flag to indicate if the object is deleted.")),
		mcp.WithString("price_amount", mcp.Description("")),
		mcp.WithBoolean("hidden", mcp.Description("")),
		mcp.WithString("version", mcp.Description("Input parameter: The user who last updated the object.")),
		mcp.WithString("code", mcp.Description("Input parameter: Product code, e.g. UPC or EAN")),
		mcp.WithBoolean("available", mcp.Description("")),
		mcp.WithObject("custom_mappings", mcp.Description("Input parameter: When custom mappings are configured on the resource, the result is included here.")),
		mcp.WithBoolean("is_revenue", mcp.Description("Input parameter: True if this item should be counted as revenue. For example, gift cards and donations would not be counted as revenue.")),
		mcp.WithArray("tax_ids", mcp.Description("Input parameter: A list of Tax IDs for the product.")),
		mcp.WithString("product_type", mcp.Description("")),
		mcp.WithBoolean("use_default_tax_rates", mcp.Description("")),
		mcp.WithString("pricing_type", mcp.Description("")),
		mcp.WithArray("options", mcp.Description("Input parameter: List of options pertaining to this item's attribute variation")),
		mcp.WithString("sku", mcp.Description("Input parameter: SKU of the item")),
		mcp.WithString("abbreviation", mcp.Description("")),
		mcp.WithString("cost", mcp.Description("")),
		mcp.WithArray("absent_at_location_ids", mcp.Description("Input parameter: A list of locations where the object is not present, even if present_at_all_locations is true. This can include locations that are deactivated.")),
		mcp.WithBoolean("available_for_pickup", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    ItemsaddHandler(cfg),
	}
}
