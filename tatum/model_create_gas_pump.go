/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type CreateGasPump struct {
	// The blockchain to work with
	Chain string `json:"chain"`
	// The blockchain address that will own the precalculated gas pump addresses and will be used to pay gas fees for operations made on the gas pump addresses; can be referred to as \"master address\"
	Owner string `json:"owner"`
	// The start index of the range of gas pump addresses to precalculate
	From int32 `json:"from"`
	// The end index of the range of gas pump addresses to precalculate; must be greater than or equal to the value in the <code>from</code> parameter
	To int32 `json:"to"`
}