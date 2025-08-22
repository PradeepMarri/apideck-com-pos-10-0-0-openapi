package main

import (
	"github.com/pos-api/mcp-server/config"
	"github.com/pos-api/mcp-server/models"
	tools_payments "github.com/pos-api/mcp-server/tools/payments"
	tools_modifiers "github.com/pos-api/mcp-server/tools/modifiers"
	tools_tenders "github.com/pos-api/mcp-server/tools/tenders"
	tools_locations "github.com/pos-api/mcp-server/tools/locations"
	tools_order_types "github.com/pos-api/mcp-server/tools/order_types"
	tools_orders "github.com/pos-api/mcp-server/tools/orders"
	tools_merchants "github.com/pos-api/mcp-server/tools/merchants"
	tools_modifier_groups "github.com/pos-api/mcp-server/tools/modifier_groups"
	tools_items "github.com/pos-api/mcp-server/tools/items"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_payments.CreatePaymentsdeleteTool(cfg),
		tools_payments.CreatePaymentsoneTool(cfg),
		tools_payments.CreatePaymentsupdateTool(cfg),
		tools_modifiers.CreateModifiersallTool(cfg),
		tools_modifiers.CreateModifiersaddTool(cfg),
		tools_payments.CreatePaymentsaddTool(cfg),
		tools_payments.CreatePaymentsallTool(cfg),
		tools_tenders.CreateTendersallTool(cfg),
		tools_tenders.CreateTendersaddTool(cfg),
		tools_locations.CreateLocationsaddTool(cfg),
		tools_locations.CreateLocationsallTool(cfg),
		tools_order_types.CreateOrdertypesdeleteTool(cfg),
		tools_order_types.CreateOrdertypesoneTool(cfg),
		tools_order_types.CreateOrdertypesupdateTool(cfg),
		tools_orders.CreateOrdersallTool(cfg),
		tools_orders.CreateOrdersaddTool(cfg),
		tools_merchants.CreateMerchantsallTool(cfg),
		tools_merchants.CreateMerchantsaddTool(cfg),
		tools_modifier_groups.CreateModifiergroupsallTool(cfg),
		tools_modifier_groups.CreateModifiergroupsaddTool(cfg),
		tools_orders.CreateOrderspayTool(cfg),
		tools_items.CreateItemsdeleteTool(cfg),
		tools_items.CreateItemsoneTool(cfg),
		tools_items.CreateItemsupdateTool(cfg),
		tools_merchants.CreateMerchantsoneTool(cfg),
		tools_merchants.CreateMerchantsupdateTool(cfg),
		tools_merchants.CreateMerchantsdeleteTool(cfg),
		tools_modifier_groups.CreateModifiergroupsdeleteTool(cfg),
		tools_modifier_groups.CreateModifiergroupsoneTool(cfg),
		tools_modifier_groups.CreateModifiergroupsupdateTool(cfg),
		tools_order_types.CreateOrdertypesallTool(cfg),
		tools_order_types.CreateOrdertypesaddTool(cfg),
		tools_locations.CreateLocationsdeleteTool(cfg),
		tools_locations.CreateLocationsoneTool(cfg),
		tools_locations.CreateLocationsupdateTool(cfg),
		tools_orders.CreateOrdersdeleteTool(cfg),
		tools_orders.CreateOrdersoneTool(cfg),
		tools_orders.CreateOrdersupdateTool(cfg),
		tools_items.CreateItemsallTool(cfg),
		tools_items.CreateItemsaddTool(cfg),
		tools_tenders.CreateTendersdeleteTool(cfg),
		tools_tenders.CreateTendersoneTool(cfg),
		tools_tenders.CreateTendersupdateTool(cfg),
		tools_modifiers.CreateModifiersdeleteTool(cfg),
		tools_modifiers.CreateModifiersoneTool(cfg),
		tools_modifiers.CreateModifiersupdateTool(cfg),
	}
}
