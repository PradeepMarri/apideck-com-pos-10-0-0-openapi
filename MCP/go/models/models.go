package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GetPosPaymentsResponse represents the GetPosPaymentsResponse schema from the OpenAPI specification
type GetPosPaymentsResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []PosPayment `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// DeleteModifierGroupResponse represents the DeleteModifierGroupResponse schema from the OpenAPI specification
type DeleteModifierGroupResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// UpdateOrderResponse represents the UpdateOrderResponse schema from the OpenAPI specification
type UpdateOrderResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// TooManyRequestsResponse represents the TooManyRequestsResponse schema from the OpenAPI specification
type TooManyRequestsResponse struct {
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail map[string]interface{} `json:"detail,omitempty"`
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 6585)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
}

// Item represents the Item schema from the OpenAPI specification
type Item struct {
	Idempotency_key string `json:"idempotency_key,omitempty"` // A value you specify that uniquely identifies this request among requests you have sent.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Variations []interface{} `json:"variations,omitempty"`
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Modifier_groups []interface{} `json:"modifier_groups,omitempty"`
	Present_at_all_locations bool `json:"present_at_all_locations,omitempty"`
	Name string `json:"name"`
	Price_currency string `json:"price_currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Id string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	Available_online bool `json:"available_online,omitempty"`
	Categories []interface{} `json:"categories,omitempty"`
	Deleted bool `json:"deleted,omitempty"` // Flag to indicate if the object is deleted.
	Price_amount float64 `json:"price_amount,omitempty"`
	Hidden bool `json:"hidden,omitempty"`
	Version string `json:"version,omitempty"` // The user who last updated the object.
	Code string `json:"code,omitempty"` // Product code, e.g. UPC or EAN
	Available bool `json:"available,omitempty"`
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Is_revenue bool `json:"is_revenue,omitempty"` // True if this item should be counted as revenue. For example, gift cards and donations would not be counted as revenue.
	Tax_ids []string `json:"tax_ids,omitempty"` // A list of Tax IDs for the product.
	Product_type string `json:"product_type,omitempty"`
	Use_default_tax_rates bool `json:"use_default_tax_rates,omitempty"`
	Pricing_type string `json:"pricing_type,omitempty"`
	Options []interface{} `json:"options,omitempty"` // List of options pertaining to this item's attribute variation
	Sku string `json:"sku,omitempty"` // SKU of the item
	Abbreviation string `json:"abbreviation,omitempty"`
	Cost float64 `json:"cost,omitempty"`
	Absent_at_location_ids []string `json:"absent_at_location_ids,omitempty"` // A list of locations where the object is not present, even if present_at_all_locations is true. This can include locations that are deactivated.
	Available_for_pickup bool `json:"available_for_pickup,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
}

// GetModifierResponse represents the GetModifierResponse schema from the OpenAPI specification
type GetModifierResponse struct {
	Data Modifier `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// CreateOrderTypeResponse represents the CreateOrderTypeResponse schema from the OpenAPI specification
type CreateOrderTypeResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// Tender represents the Tender schema from the OpenAPI specification
type Tender struct {
	Allows_tipping bool `json:"allows_tipping,omitempty"` // Allow tipping on payment from tender
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Editable bool `json:"editable,omitempty"`
	Label string `json:"label,omitempty"`
	Active bool `json:"active,omitempty"`
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Key string `json:"key,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Hidden bool `json:"hidden,omitempty"`
	Opens_cash_drawer bool `json:"opens_cash_drawer,omitempty"` // If this tender opens the cash drawer
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
}

// ModifierGroup represents the ModifierGroup schema from the OpenAPI specification
type ModifierGroup struct {
	Deleted bool `json:"deleted,omitempty"` // Flag to indicate if the object is deleted.
	Alternate_name string `json:"alternate_name,omitempty"`
	Modifiers []interface{} `json:"modifiers,omitempty"`
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Name string `json:"name,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Maximum_allowed int `json:"maximum_allowed,omitempty"`
	Minimum_required int `json:"minimum_required,omitempty"`
	Selection_type string `json:"selection_type,omitempty"`
	Present_at_all_locations bool `json:"present_at_all_locations,omitempty"`
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
}

// GetModifierGroupResponse represents the GetModifierGroupResponse schema from the OpenAPI specification
type GetModifierGroupResponse struct {
	Data ModifierGroup `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// UnexpectedErrorResponse represents the UnexpectedErrorResponse schema from the OpenAPI specification
type UnexpectedErrorResponse struct {
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
}

// PosBankAccount represents the PosBankAccount schema from the OpenAPI specification
type PosBankAccount struct {
	Ach_details map[string]interface{} `json:"ach_details,omitempty"` // ACH-specific details about `BANK_ACCOUNT` type payments with the `transfer_type` of `ACH`.
	Bank_name string `json:"bank_name,omitempty"` // The name of the bank associated with the bank account.
	Country string `json:"country,omitempty"` // Country code according to ISO 3166-1 alpha-2.
	Fingerprint string `json:"fingerprint,omitempty"` // Uniquely identifies the bank account for this seller and can be used to determine if payments are from the same bank account.
	Statement_description string `json:"statement_description,omitempty"` // The statement description as sent to the bank.
	Transfer_type string `json:"transfer_type,omitempty"` // The type of the bank transfer. The type can be `ACH` or `UNKNOWN`.
	Account_ownership_type string `json:"account_ownership_type,omitempty"` // The ownership type of the bank account performing the transfer. The type can be `INDIVIDUAL`, `COMPANY`, or `UNKNOWN`.
}

// GetItemsResponse represents the GetItemsResponse schema from the OpenAPI specification
type GetItemsResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Item `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// Merchant represents the Merchant schema from the OpenAPI specification
type Merchant struct {
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Address Address `json:"address,omitempty"`
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Name string `json:"name,omitempty"` // The name of the merchant
	Owner_id string `json:"owner_id,omitempty"`
	Status string `json:"status,omitempty"` // Status of this merchant.
	Language string `json:"language,omitempty"` // language code according to ISO 639-1. For the United States - EN
	Main_location_id string `json:"main_location_id,omitempty"` // The main location ID of the merchant
	Service_charges []ServiceCharge `json:"service_charges,omitempty"`
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
}

// BadRequestResponse represents the BadRequestResponse schema from the OpenAPI specification
type BadRequestResponse struct {
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
}

// GetLocationsResponse represents the GetLocationsResponse schema from the OpenAPI specification
type GetLocationsResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Location `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// UpdateModifierGroupResponse represents the UpdateModifierGroupResponse schema from the OpenAPI specification
type UpdateModifierGroupResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// GetModifierGroupsResponse represents the GetModifierGroupsResponse schema from the OpenAPI specification
type GetModifierGroupsResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []ModifierGroup `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// GetTendersResponse represents the GetTendersResponse schema from the OpenAPI specification
type GetTendersResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Tender `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
}

// DeleteItemResponse represents the DeleteItemResponse schema from the OpenAPI specification
type DeleteItemResponse struct {
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// CreateModifierResponse represents the CreateModifierResponse schema from the OpenAPI specification
type CreateModifierResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// OrderType represents the OrderType schema from the OpenAPI specification
type OrderType struct {
	Name string `json:"name,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	DefaultField bool `json:"default,omitempty"`
	Id string `json:"id,omitempty"` // A unique identifier for an object.
}

// UpdateModifierResponse represents the UpdateModifierResponse schema from the OpenAPI specification
type UpdateModifierResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// CreateItemResponse represents the CreateItemResponse schema from the OpenAPI specification
type CreateItemResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// CreateModifierGroupResponse represents the CreateModifierGroupResponse schema from the OpenAPI specification
type CreateModifierGroupResponse struct {
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// CreateOrderResponse represents the CreateOrderResponse schema from the OpenAPI specification
type CreateOrderResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// Order represents the Order schema from the OpenAPI specification
type Order struct {
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Fulfillments []map[string]interface{} `json:"fulfillments,omitempty"`
	Location_id string `json:"location_id"`
	Voided_at string `json:"voided_at,omitempty"`
	Seat string `json:"seat,omitempty"`
	Title string `json:"title,omitempty"`
	Order_date string `json:"order_date,omitempty"`
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Payments []map[string]interface{} `json:"payments,omitempty"`
	Total_service_charge int `json:"total_service_charge,omitempty"`
	Status string `json:"status,omitempty"` // Order status. Clover specific: If no value is set, the status defaults to hidden, which indicates a hidden order. A hidden order is not displayed in user interfaces and can only be retrieved by its id. When creating an order via the REST API the value must be manually set to 'open'. More info [https://docs.clover.com/reference/orderupdateorder]()
	Voided bool `json:"voided,omitempty"`
	Customers []map[string]interface{} `json:"customers,omitempty"`
	Customer_id string `json:"customer_id,omitempty"`
	Total_tip int `json:"total_tip,omitempty"`
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Refunds []map[string]interface{} `json:"refunds,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Discounts []map[string]interface{} `json:"discounts,omitempty"`
	Total_discount int `json:"total_discount,omitempty"`
	Tenders []map[string]interface{} `json:"tenders,omitempty"`
	Merchant_id string `json:"merchant_id"`
	Service_charges []ServiceCharge `json:"service_charges,omitempty"` // Optional service charges or gratuity tip applied to the order.
	Employee_id string `json:"employee_id,omitempty"`
	Note string `json:"note,omitempty"` // A note with information about this order, may be printed on the order receipt and displayed in apps
	Total_tax int `json:"total_tax,omitempty"`
	Source string `json:"source,omitempty"` // Source of order. Indicates the way that the order was placed.
	Taxes []interface{} `json:"taxes,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Order_number string `json:"order_number,omitempty"`
	Line_items []map[string]interface{} `json:"line_items,omitempty"`
	Reference_id string `json:"reference_id,omitempty"` // An optional user-defined reference ID that associates this record with another entity in an external system. For example, a customer ID from an external customer management system.
	Payment_status string `json:"payment_status,omitempty"` // Is this order paid or not?
	Idempotency_key string `json:"idempotency_key,omitempty"` // A value you specify that uniquely identifies this request among requests you have sent.
	Order_type_id string `json:"order_type_id,omitempty"`
	Total_amount int `json:"total_amount,omitempty"`
	Table string `json:"table,omitempty"`
	Total_refund int `json:"total_refund,omitempty"`
	Version string `json:"version,omitempty"`
	Refunded bool `json:"refunded,omitempty"`
	Closed_date string `json:"closed_date,omitempty"`
}

// UnauthorizedResponse represents the UnauthorizedResponse schema from the OpenAPI specification
type UnauthorizedResponse struct {
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
}

// UpdateMerchantResponse represents the UpdateMerchantResponse schema from the OpenAPI specification
type UpdateMerchantResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// NotImplementedResponse represents the NotImplementedResponse schema from the OpenAPI specification
type NotImplementedResponse struct {
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
}

// PosWebhookEvent represents the PosWebhookEvent schema from the OpenAPI specification
type PosWebhookEvent struct {
	Entity_id string `json:"entity_id,omitempty"` // The service provider's ID of the entity that triggered this event
	Event_id string `json:"event_id,omitempty"` // Unique reference to this request event
	Unified_api string `json:"unified_api,omitempty"` // Name of Apideck Unified API
	Entity_type string `json:"entity_type,omitempty"` // The type entity that triggered this event
	Occurred_at string `json:"occurred_at,omitempty"` // ISO Datetime for when the original event occurred
	Service_id string `json:"service_id,omitempty"` // Service provider identifier
	Execution_attempt float64 `json:"execution_attempt,omitempty"` // The current count this request event has been attempted
	Entity_url string `json:"entity_url,omitempty"` // The url to retrieve entity detail.
	Consumer_id string `json:"consumer_id,omitempty"` // Unique consumer identifier. You can freely choose a consumer ID yourself. Most of the time, this is an ID of your internal data model that represents a user or account in your system (for example account:12345). If the consumer doesn't exist yet, Vault will upsert a consumer based on your ID.
	Event_type string `json:"event_type,omitempty"`
}

// UpdateOrderTypeResponse represents the UpdateOrderTypeResponse schema from the OpenAPI specification
type UpdateOrderTypeResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// Website represents the Website schema from the OpenAPI specification
type Website struct {
	Id string `json:"id,omitempty"` // Unique identifier for the website
	TypeField string `json:"type,omitempty"` // The type of website
	Url string `json:"url"` // The website URL
}

// GetMerchantResponse represents the GetMerchantResponse schema from the OpenAPI specification
type GetMerchantResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Merchant `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// GetPosPaymentResponse represents the GetPosPaymentResponse schema from the OpenAPI specification
type GetPosPaymentResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data PosPayment `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// DeleteMerchantResponse represents the DeleteMerchantResponse schema from the OpenAPI specification
type DeleteMerchantResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// UnprocessableResponse represents the UnprocessableResponse schema from the OpenAPI specification
type UnprocessableResponse struct {
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// UpdateLocationResponse represents the UpdateLocationResponse schema from the OpenAPI specification
type UpdateLocationResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// NotFoundResponse represents the NotFoundResponse schema from the OpenAPI specification
type NotFoundResponse struct {
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
}

// CreateTenderResponse represents the CreateTenderResponse schema from the OpenAPI specification
type CreateTenderResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// DeleteOrderResponse represents the DeleteOrderResponse schema from the OpenAPI specification
type DeleteOrderResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// GetOrderResponse represents the GetOrderResponse schema from the OpenAPI specification
type GetOrderResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Order `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// GetOrderTypeResponse represents the GetOrderTypeResponse schema from the OpenAPI specification
type GetOrderTypeResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data OrderType `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// PhoneNumber represents the PhoneNumber schema from the OpenAPI specification
type PhoneNumber struct {
	Id string `json:"id,omitempty"` // Unique identifier of the phone number
	Number string `json:"number"` // The phone number
	TypeField string `json:"type,omitempty"` // The type of phone number
	Area_code string `json:"area_code,omitempty"` // The area code of the phone number, e.g. 323
	Country_code string `json:"country_code,omitempty"` // The country code of the phone number, e.g. +1
	Extension string `json:"extension,omitempty"` // The extension of the phone number
}

// Address represents the Address schema from the OpenAPI specification
type Address struct {
	Phone_number string `json:"phone_number,omitempty"` // Phone number of the address
	Country string `json:"country,omitempty"` // country code according to ISO 3166-1 alpha-2.
	Contact_name string `json:"contact_name,omitempty"` // Name of the contact person at the address
	TypeField string `json:"type,omitempty"` // The type of address.
	County string `json:"county,omitempty"` // Address field that holds a sublocality, such as a county
	Salutation string `json:"salutation,omitempty"` // Salutation of the contact person at the address
	State string `json:"state,omitempty"` // Name of state
	City string `json:"city,omitempty"` // Name of city.
	Line3 string `json:"line3,omitempty"` // Line 3 of the address
	Row_version string `json:"row_version,omitempty"` // A binary value used to detect updates to a object and prevent data conflicts. It is incremented each time an update is made to the object.
	Website string `json:"website,omitempty"` // Website of the address
	Fax string `json:"fax,omitempty"` // Fax number of the address
	Latitude string `json:"latitude,omitempty"` // Latitude of the address
	StringField string `json:"string,omitempty"` // The address string. Some APIs don't provide structured address data.
	Id string `json:"id,omitempty"` // Unique identifier for the address.
	Name string `json:"name,omitempty"` // The name of the address.
	Postal_code string `json:"postal_code,omitempty"` // Zip code or equivalent.
	Email string `json:"email,omitempty"` // Email address of the address
	Line2 string `json:"line2,omitempty"` // Line 2 of the address
	Street_number string `json:"street_number,omitempty"` // Street number
	Line1 string `json:"line1,omitempty"` // Line 1 of the address e.g. number, street, suite, apt #, etc.
	Line4 string `json:"line4,omitempty"` // Line 4 of the address
	Longitude string `json:"longitude,omitempty"` // Longitude of the address
	Notes string `json:"notes,omitempty"` // Additional notes
}

// UpdateItemResponse represents the UpdateItemResponse schema from the OpenAPI specification
type UpdateItemResponse struct {
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// DeleteModifierResponse represents the DeleteModifierResponse schema from the OpenAPI specification
type DeleteModifierResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// PaymentCard represents the PaymentCard schema from the OpenAPI specification
type PaymentCard struct {
	Bin string `json:"bin,omitempty"` // The first six digits of the card number, known as the Bank Identification Number (BIN).
	Card_brand string `json:"card_brand,omitempty"` // The first six digits of the card number, known as the Bank Identification Number (BIN).
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Customer_id string `json:"customer_id,omitempty"`
	Cardholder_name string `json:"cardholder_name,omitempty"`
	Merchant_id string `json:"merchant_id,omitempty"`
	Exp_year int `json:"exp_year,omitempty"` // The four-digit year of the card's expiration date.
	Billing_address Address `json:"billing_address,omitempty"`
	Exp_month int `json:"exp_month,omitempty"` // The expiration month of the associated card as an integer between 1 and 12.
	Fingerprint string `json:"fingerprint,omitempty"`
	Prepaid_type string `json:"prepaid_type,omitempty"`
	Enabled bool `json:"enabled,omitempty"` // Indicates whether or not a card can be used for payments.
	Last_4 string `json:"last_4,omitempty"`
	Card_type string `json:"card_type,omitempty"`
	Reference_id string `json:"reference_id,omitempty"` // An optional user-defined reference ID that associates this record with another entity in an external system. For example, a customer ID from an external customer management system.
	Version string `json:"version,omitempty"`
}

// SocialLink represents the SocialLink schema from the OpenAPI specification
type SocialLink struct {
	Id string `json:"id,omitempty"` // Unique identifier of the social link
	TypeField string `json:"type,omitempty"` // Type of the social link, e.g. twitter
	Url string `json:"url"` // URL of the social link, e.g. https://www.twitter.com/apideck
}

// PaymentRequiredResponse represents the PaymentRequiredResponse schema from the OpenAPI specification
type PaymentRequiredResponse struct {
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
}

// Meta represents the Meta schema from the OpenAPI specification
type Meta struct {
	Cursors map[string]interface{} `json:"cursors,omitempty"` // Cursors to navigate to previous or next pages through the API
	Items_on_page int `json:"items_on_page,omitempty"` // Number of items returned in the data property of the response
}

// CreatePosPaymentResponse represents the CreatePosPaymentResponse schema from the OpenAPI specification
type CreatePosPaymentResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// GetModifiersResponse represents the GetModifiersResponse schema from the OpenAPI specification
type GetModifiersResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Modifier `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// GetOrderTypesResponse represents the GetOrderTypesResponse schema from the OpenAPI specification
type GetOrderTypesResponse struct {
	Data []OrderType `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// UpdateTenderResponse represents the UpdateTenderResponse schema from the OpenAPI specification
type UpdateTenderResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// DeleteTenderResponse represents the DeleteTenderResponse schema from the OpenAPI specification
type DeleteTenderResponse struct {
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// CustomMappings represents the CustomMappings schema from the OpenAPI specification
type CustomMappings struct {
}

// GetMerchantsResponse represents the GetMerchantsResponse schema from the OpenAPI specification
type GetMerchantsResponse struct {
	Data []Merchant `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// CreateMerchantResponse represents the CreateMerchantResponse schema from the OpenAPI specification
type CreateMerchantResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// DeleteLocationResponse represents the DeleteLocationResponse schema from the OpenAPI specification
type DeleteLocationResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
}

// PosPayment represents the PosPayment schema from the OpenAPI specification
type PosPayment struct {
	Device_id string `json:"device_id,omitempty"`
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Idempotency_key string `json:"idempotency_key,omitempty"` // A value you specify that uniquely identifies this request among requests you have sent.
	Order_id string `json:"order_id"`
	External_payment_id string `json:"external_payment_id,omitempty"`
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Card_details map[string]interface{} `json:"card_details,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Tip float64 `json:"tip,omitempty"`
	Status string `json:"status,omitempty"` // Status of this payment.
	App_fee float64 `json:"app_fee,omitempty"` // The amount the developer is taking as a fee for facilitating the payment on behalf of the seller.
	Tender_id string `json:"tender_id"`
	External_details map[string]interface{} `json:"external_details,omitempty"` // Details about an external payment.
	Source string `json:"source,omitempty"` // Source of this payment.
	Wallet map[string]interface{} `json:"wallet,omitempty"` // Wallet details for this payment. This field is currently not available. Reach out to our team for more info.
	Source_id string `json:"source_id"` // The ID for the source of funds for this payment. Square-only: This can be a payment token (card nonce) generated by the payment form or a card on file made linked to the customer. if recording a payment that the seller received outside of Square, specify either `CASH` or `EXTERNAL`.
	Tax float64 `json:"tax,omitempty"`
	Currency string `json:"currency"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Merchant_id string `json:"merchant_id,omitempty"`
	Processing_fees []interface{} `json:"processing_fees,omitempty"`
	Approved float64 `json:"approved,omitempty"` // The initial amount of money approved for this payment.
	Customer_id string `json:"customer_id"`
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Location_id string `json:"location_id,omitempty"`
	Amount float64 `json:"amount"`
	Refunded float64 `json:"refunded,omitempty"` // The initial amount of money approved for this payment.
	Cash map[string]interface{} `json:"cash,omitempty"` // Cash details for this payment
	Bank_account PosBankAccount `json:"bank_account,omitempty"` // Card details for this payment. This field is currently not available. Reach out to our team for more info.
	Total float64 `json:"total,omitempty"`
	Employee_id string `json:"employee_id,omitempty"`
	Service_charges []ServiceCharge `json:"service_charges,omitempty"` // Optional service charges or gratuity tip applied to the order.
	Change_back_cash_amount float64 `json:"change_back_cash_amount,omitempty"`
}

// DeletePosPaymentResponse represents the DeletePosPaymentResponse schema from the OpenAPI specification
type DeletePosPaymentResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// Links represents the Links schema from the OpenAPI specification
type Links struct {
	Current string `json:"current,omitempty"` // Link to navigate to the current page through the API
	Next string `json:"next,omitempty"` // Link to navigate to the previous page through the API
	Previous string `json:"previous,omitempty"` // Link to navigate to the previous page through the API
}

// ModifierGroupFilter represents the ModifierGroupFilter schema from the OpenAPI specification
type ModifierGroupFilter struct {
	Modifier_group_id string `json:"modifier_group_id,omitempty"` // Id of the job to filter on
}

// ServiceCharge represents the ServiceCharge schema from the OpenAPI specification
type ServiceCharge struct {
	Name string `json:"name,omitempty"` // Service charge name
	Percentage float64 `json:"percentage,omitempty"` // Service charge percentage. Use this field to calculate the amount of the service charge. Pass a percentage and amount at the same time.
	TypeField string `json:"type,omitempty"` // The type of the service charge.
	Active bool `json:"active,omitempty"`
	Amount float64 `json:"amount,omitempty"`
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Id string `json:"id,omitempty"` // A unique identifier for an object.
}

// CustomField represents the CustomField schema from the OpenAPI specification
type CustomField struct {
	Value interface{} `json:"value,omitempty"`
	Description string `json:"description,omitempty"` // More information about the custom field
	Id string `json:"id"` // Unique identifier for the custom field.
	Name string `json:"name,omitempty"` // Name of the custom field.
}

// GetLocationResponse represents the GetLocationResponse schema from the OpenAPI specification
type GetLocationResponse struct {
	Data Location `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// Location represents the Location schema from the OpenAPI specification
type Location struct {
	Name string `json:"name,omitempty"` // The name of the location
	Status string `json:"status,omitempty"` // Status of this location.
	Address Address `json:"address,omitempty"`
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Merchant_id string `json:"merchant_id,omitempty"`
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Business_name string `json:"business_name,omitempty"` // The business name of the location
}

// GetOrdersResponse represents the GetOrdersResponse schema from the OpenAPI specification
type GetOrdersResponse struct {
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Order `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
}

// DeleteOrderTypeResponse represents the DeleteOrderTypeResponse schema from the OpenAPI specification
type DeleteOrderTypeResponse struct {
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
}

// GetItemResponse represents the GetItemResponse schema from the OpenAPI specification
type GetItemResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Item `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// UnifiedId represents the UnifiedId schema from the OpenAPI specification
type UnifiedId struct {
	Id string `json:"id"` // The unique identifier of the resource
}

// UpdatePosPaymentResponse represents the UpdatePosPaymentResponse schema from the OpenAPI specification
type UpdatePosPaymentResponse struct {
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
}

// GetTenderResponse represents the GetTenderResponse schema from the OpenAPI specification
type GetTenderResponse struct {
	Data Tender `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// CreateLocationResponse represents the CreateLocationResponse schema from the OpenAPI specification
type CreateLocationResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data UnifiedId `json:"data"`
	Operation string `json:"operation"` // Operation performed
	Resource string `json:"resource"` // Unified API resource name
	Service string `json:"service"` // Apideck ID of service provider
}

// Modifier represents the Modifier schema from the OpenAPI specification
type Modifier struct {
	Id string `json:"id,omitempty"` // A unique identifier for an object.
	Idempotency_key string `json:"idempotency_key,omitempty"` // A value you specify that uniquely identifies this request among requests you have sent.
	Modifier_group_id string `json:"modifier_group_id"`
	Available bool `json:"available,omitempty"`
	Currency string `json:"currency,omitempty"` // Indicates the associated currency for an amount of money. Values correspond to [ISO 4217](https://en.wikipedia.org/wiki/ISO_4217).
	Price_amount float64 `json:"price_amount,omitempty"`
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Custom_mappings CustomMappings `json:"custom_mappings,omitempty"` // When custom mappings are configured on the resource, the result is included here.
	Name string `json:"name"`
	Updated_by string `json:"updated_by,omitempty"` // The user who last updated the object.
	Alternate_name string `json:"alternate_name,omitempty"`
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Created_by string `json:"created_by,omitempty"` // The user who created the object.
}

// Email represents the Email schema from the OpenAPI specification
type Email struct {
	TypeField string `json:"type,omitempty"` // Email type
	Email string `json:"email"` // Email address
	Id string `json:"id,omitempty"` // Unique identifier for the email address
}
