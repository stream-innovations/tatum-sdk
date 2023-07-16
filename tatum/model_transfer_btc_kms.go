/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type TransferBtcKms struct {
	// Sender account ID
	SenderAccountId string `json:"senderAccountId"`
	// Blockchain address to send assets to. For BTC, LTC, DOGE and BCH, it is possible to enter list of multiple recipient blockchain addresses as a comma separated string.
	Address string `json:"address"`
	// Amount to be withdrawn to blockchain.
	Amount string `json:"amount"`
	// Compliance check, if withdrawal is not compliant, it will not be processed.
	Compliant bool `json:"compliant,omitempty"`
	// Fee to be submitted as a transaction fee to blockchain. If none is set, default value of 0.0005 BTC is used. Minimum fee is 0.00001 BTC.
	Fee string `json:"fee,omitempty"`
	// For BTC, LTC, DOGE and BCH, it is possible to enter list of multiple recipient blockchain amounts. List of recipient addresses must be present in the address field and total sum of amounts must be equal to the amount field.
	MultipleAmounts []string `json:"multipleAmounts,omitempty"`
	// Used to parametrize withdrawal as a change address for left coins from transaction. XPub or attr must be used.
	Attr string `json:"attr,omitempty"`
	// Signature hash of the mnemonic, which will be used to sign transactions locally. All signature Ids should be present, which might be used to sign transaction. Tatum KMS does not support keyPair type of off-chain transaction, only mnemonic based. 
	SignatureId string `json:"signatureId"`
	// Extended public key (xpub) of the wallet associated with the accounts. Should be present, when mnemonic is used.
	Xpub string `json:"xpub"`
	// Identifier of the payment, shown for created Transaction within Tatum sender account.
	PaymentId string `json:"paymentId,omitempty"`
	// Note visible to owner of withdrawing account
	SenderNote string `json:"senderNote,omitempty"`
}
