/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type Customer struct {
	// Customer external ID.
	ExternalId string `json:"externalId"`
	// Customer internal ID within Tatum.
	Id string `json:"id"`
	// Indicates whether customer is enabled or not
	Enabled bool `json:"enabled"`
	// Indicates whether customer is active or not
	Active bool `json:"active"`
	// All transaction will be accounted in this currency for all accounts of the customer. Currency can be overridden per account level. ISO-4217
	AccountingCurrency string `json:"accountingCurrency,omitempty"`
	// Country customer has to be compliant with
	CustomerCountry string `json:"customerCountry,omitempty"`
	// Country service provider has to be compliant with
	ProviderCountry string `json:"providerCountry,omitempty"`
}