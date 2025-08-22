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

func OrdersupdateHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		idVal, ok := args["id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: id"), nil
		}
		id, ok := idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: id"), nil
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
		var requestBody models.Order
		
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
		url := fmt.Sprintf("%s/pos/orders/%s%s", cfg.BaseURL, id, queryString)
		req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(bodyBytes))
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
		var result models.UpdateOrderResponse
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

func CreateOrdersupdateTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("patch_pos_orders_id",
		mcp.WithDescription("Update Order"),
		mcp.WithString("id", mcp.Required(), mcp.Description("ID of the record you are acting upon.")),
		mcp.WithString("x-apideck-consumer-id", mcp.Required(), mcp.Description("ID of the consumer which you want to get or push data from")),
		mcp.WithString("x-apideck-app-id", mcp.Required(), mcp.Description("The ID of your Unify application")),
		mcp.WithString("x-apideck-service-id", mcp.Description("Provide the service id you want to call (e.g., pipedrive). Only needed when a consumer has activated multiple integrations for a Unified API.")),
		mcp.WithBoolean("raw", mcp.Description("Include raw response. Mostly used for debugging purposes")),
		mcp.WithString("updated_by", mcp.Description("Input parameter: The user who last updated the object.")),
		mcp.WithArray("refunds", mcp.Description("")),
		mcp.WithString("updated_at", mcp.Description("Input parameter: The date and time when the object was last updated.")),
		mcp.WithString("id", mcp.Description("Input parameter: A unique identifier for an object.")),
		mcp.WithObject("custom_mappings", mcp.Description("Input parameter: When custom mappings are configured on the resource, the result is included here.")),
		mcp.WithArray("discounts", mcp.Description("")),
		mcp.WithNumber("total_discount", mcp.Description("")),
		mcp.WithArray("tenders", mcp.Description("")),
		mcp.WithString("merchant_id", mcp.Required(), mcp.Description("")),
		mcp.WithArray("service_charges", mcp.Description("Input parameter: Optional service charges or gratuity tip applied to the order.")),
		mcp.WithString("employee_id", mcp.Description("")),
		mcp.WithString("note", mcp.Description("Input parameter: A note with information about this order, may be printed on the order receipt and displayed in apps")),
		mcp.WithNumber("total_tax", mcp.Description("")),
		mcp.WithString("source", mcp.Description("Input parameter: Source of order. Indicates the way that the order was placed.")),
		mcp.WithArray("taxes", mcp.Description("")),
		mcp.WithString("created_at", mcp.Description("Input parameter: The date and time when the object was created.")),
		mcp.WithString("order_number", mcp.Description("")),
		mcp.WithArray("line_items", mcp.Description("")),
		mcp.WithString("reference_id", mcp.Description("Input parameter: An optional user-defined reference ID that associates this record with another entity in an external system. For example, a customer ID from an external customer management system.")),
		mcp.WithString("payment_status", mcp.Description("Input parameter: Is this order paid or not?")),
		mcp.WithString("idempotency_key", mcp.Description("Input parameter: A value you specify that uniquely identifies this request among requests you have sent.")),
		mcp.WithString("order_type_id", mcp.Description("")),
		mcp.WithNumber("total_amount", mcp.Description("")),
		mcp.WithString("table", mcp.Description("")),
		mcp.WithNumber("total_refund", mcp.Description("")),
		mcp.WithString("version", mcp.Description("")),
		mcp.WithBoolean("refunded", mcp.Description("")),
		mcp.WithString("closed_date", mcp.Description("")),
		mcp.WithString("created_by", mcp.Description("Input parameter: The user who created the object.")),
		mcp.WithArray("fulfillments", mcp.Description("")),
		mcp.WithString("location_id", mcp.Required(), mcp.Description("")),
		mcp.WithString("voided_at", mcp.Description("")),
		mcp.WithString("seat", mcp.Description("")),
		mcp.WithString("title", mcp.Description("")),
		mcp.WithString("order_date", mcp.Description("")),
		mcp.WithString("currency", mcp.Description("Input parameter: Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).")),
		mcp.WithArray("payments", mcp.Description("")),
		mcp.WithNumber("total_service_charge", mcp.Description("")),
		mcp.WithString("status", mcp.Description("Input parameter: Order status. Clover specific: If no value is set, the status defaults to hidden, which indicates a hidden order. A hidden order is not displayed in user interfaces and can only be retrieved by its id. When creating an order via the REST API the value must be manually set to 'open'. More info [https://docs.clover.com/reference/orderupdateorder]()")),
		mcp.WithBoolean("voided", mcp.Description("")),
		mcp.WithArray("customers", mcp.Description("")),
		mcp.WithString("customer_id", mcp.Description("")),
		mcp.WithNumber("total_tip", mcp.Description("")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    OrdersupdateHandler(cfg),
	}
}
