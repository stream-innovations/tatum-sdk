/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type KcsTx struct {
	// Hash of the block where this transaction was in.
	BlockHash string `json:"blockHash,omitempty"`
	// TRUE if the transaction was successful, FALSE, if the EVM reverted the transaction.
	Status bool `json:"status,omitempty"`
	// Block number where this transaction was in.
	BlockNumber float64 `json:"blockNumber,omitempty"`
	// Address of the sender.
	From string `json:"from,omitempty"`
	// Gas provided by the sender.
	Gas float64 `json:"gas,omitempty"`
	// Gas price provided by the sender in wei.
	GasPrice string `json:"gasPrice,omitempty"`
	// Hash of the transaction.
	TransactionHash string `json:"transactionHash,omitempty"`
	// The data sent along with the transaction.
	Input string `json:"input,omitempty"`
	// The number of transactions made by the sender prior to this one.
	Nonce float64 `json:"nonce,omitempty"`
	// Address of the receiver. 'null' when its a contract creation transaction.
	To string `json:"to,omitempty"`
	// Integer of the transactions index position in the block.
	TransactionIndex float64 `json:"transactionIndex,omitempty"`
	// Value transferred in wei.
	Value string `json:"value,omitempty"`
	// The amount of gas used by this specific transaction alone.
	GasUsed float64 `json:"gasUsed,omitempty"`
	// The total amount of gas used when this transaction was executed in the block.
	CumulativeGasUsed float64 `json:"cumulativeGasUsed,omitempty"`
	// The contract address created, if the transaction was a contract creation, otherwise null.
	ContractAddress string `json:"contractAddress,omitempty"`
	// Log events, that happened in this transaction.
	Logs []BscTxLogs `json:"logs,omitempty"`
}
