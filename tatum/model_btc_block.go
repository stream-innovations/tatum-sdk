/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type BtcBlock struct {
	// Hash of block.
	Hash string `json:"hash,omitempty"`
	// The number of blocks preceding a particular block on a block chain.
	Height float64 `json:"height,omitempty"`
	// The number of blocks following a particular block on a block chain, including current one.
	Depth float64 `json:"depth,omitempty"`
	// Block version.
	Version float64 `json:"version,omitempty"`
	// Hash of the previous block.
	PrevBlock string `json:"prevBlock,omitempty"`
	// The root node of a merkle tree, a descendant of all the hashed pairs in the tree.
	MerkleRoot string `json:"merkleRoot,omitempty"`
	// Time of the block.
	Time float64 `json:"time,omitempty"`
	Bits float64 `json:"bits,omitempty"`
	// Arbitrary number that is used in Bitcoin's proof of work consensus algorithm.
	Nonce float64 `json:"nonce,omitempty"`
	Txs []BtcTx `json:"txs,omitempty"`
}
