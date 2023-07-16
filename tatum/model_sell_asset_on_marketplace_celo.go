/*
 * Tatum API Reference
 *
 * # Welcome to the Tatum API Reference!  ## What is Tatum? ## What is Tatum?  Tatum offers a flexible framework to build, run, and scale blockchain apps fast. To learn more about the Tatum blockchain development framework, visit [our website](https://tatum.io/framework).  The Tatum API features powerful endpoints that simplify a complex blockchain into single API requests. Code for all supported blockchains using unified API calls.  ## Need help?  To chat with other developers, get help from the Support team, and engage with the thriving Tatum community, join  our [Discord server](https://discord.com/invite/tatum). For more information about how to work with Tatum, review the [online documentation](https://docs.tatum.io/).  ## About this API Reference  The Tatum API Reference is based on OpenAPI Specification v3.1.0 with a few [vendor extensions](https://github.com/Redocly/redoc/blob/master/docs/redoc-vendor-extensions.md) applied.  # Authentication  When using the Tatum API, you authenticate yourself with an **API key**. <SecurityDefinitions /> 
 *
 * API version: 3.16.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package tatum

type SellAssetOnMarketplaceCelo struct {
	// Blockchain to work with.
	Chain string `json:"chain"`
	// Currency to pay for transaction gas
	FeeCurrency string `json:"feeCurrency"`
	// Address of the marketplace smart contract.
	ContractAddress string `json:"contractAddress"`
	// Address of the NFT asset to sell smart contract.
	NftAddress string `json:"nftAddress"`
	// Address of the seller of the NFT asset.
	Seller string `json:"seller"`
	// Optional address of the ERC20 token, which will be used as a selling currency of the NFT.
	Erc20Address string `json:"erc20Address,omitempty"`
	// ID of the listing. It's up to the developer to generate unique ID
	ListingId string `json:"listingId"`
	// Amount of the assets to be sent. For ERC-721 tokens, enter amount only in case of native currency cashback.
	Amount string `json:"amount,omitempty"`
	// ID of token, if transaction is for ERC-721 or ERC-1155.
	TokenId string `json:"tokenId"`
	// Price of the asset to sell. Marketplace fee will be obtained on top of this price.
	Price string `json:"price"`
	// True if asset is NFT of type ERC721, false if ERC1155.
	IsErc721 bool `json:"isErc721"`
	// Private key of sender address. Private key, or signature Id must be present.
	FromPrivateKey string `json:"fromPrivateKey"`
	// Nonce to be set to Ethereum transaction. If not present, last known nonce will be used.
	Nonce float64 `json:"nonce,omitempty"`
	Fee *DeployErc20Fee `json:"fee,omitempty"`
}
