/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type XrpFee struct {
	// Number of transactions provisionally included in the in-progress ledger.
	CurrentLedgerSize string `json:"current_ledger_size,omitempty"`
	// Number of transactions currently queued for the next ledger.
	CurrentQueueSize string `json:"current_queue_size,omitempty"`
	Drops *XrpFeeDrops `json:"drops,omitempty"`
	// The approximate number of transactions expected to be included in the current ledger. This is based on the number of transactions in the previous ledger.
	ExpectedLedgerSize string `json:"expected_ledger_size,omitempty"`
	// The Ledger Index of the current open ledger these stats describe.
	LedgerCurrentIndex float64 `json:"ledger_current_index,omitempty"`
	Levels *XrpFeeLevels `json:"levels,omitempty"`
	// The maximum number of transactions that the transaction queue can currently hold.
	MaxQueueSize string `json:"max_queue_size,omitempty"`
}