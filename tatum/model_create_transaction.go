/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type CreateTransaction struct {
	// Internal sender account ID within Tatum platform
	SenderAccountId string `json:"senderAccountId"`
	// Internal recipient account ID within Tatum platform
	RecipientAccountId string `json:"recipientAccountId"`
	// Amount to be sent.
	Amount string `json:"amount"`
	// Anonymous transaction does not show sender account to recipient, default is false
	Anonymous bool `json:"anonymous,omitempty"`
	// Enable compliant checks. Transaction will not be processed, if compliant check fails.
	Compliant bool `json:"compliant,omitempty"`
	// For bookkeeping to distinct transaction purpose.
	TransactionCode string `json:"transactionCode,omitempty"`
	// Payment ID, External identifier of the payment, which can be used to pair transactions within Tatum accounts.
	PaymentId string `json:"paymentId,omitempty"`
	// Note visible to both, sender and recipient
	RecipientNote string `json:"recipientNote,omitempty"`
	// Exchange rate of the base pair. Only applicable for Tatum's Virtual currencies Ledger transactions. Override default exchange rate for the Virtual Currency.
	BaseRate float64 `json:"baseRate,omitempty"`
	// Note visible to sender
	SenderNote string `json:"senderNote,omitempty"`
}
