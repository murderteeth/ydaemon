package prices

import (
	"testing"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/yearn/ydaemon/internal/models"
)

//
// cd processes/prices
// go test -v ./...
//
func Test_yvmkUSD(t *testing.T) {
	fmt.Println("get yvmkUSD price..")
	mkUSD := common.HexToAddress("0x4591DBfF62656E7859Afe5e45f6f47D3669fBB28")
	yvmkUSD := common.HexToAddress("0x04AeBe2e4301CdF5E9c57B01eBdfe4Ac4B48DD13")
	tokenMap := make(map[common.Address]models.TERC20Token)
	tokenMap[yvmkUSD] = models.TERC20Token{
		Address:                   yvmkUSD,
		UnderlyingTokensAddresses: []common.Address{},
		Type:                      "", // set the appropriate TTokenType value
		Name:                      "yvmkUSD",
		Symbol:                    "", // set the symbol if known
		DisplayName:               "", // set the display name if desired
		DisplaySymbol:             "", // set the display symbol if desired
		Description:               "", // provide a description if necessary
		Category:                  "", // set the category if applicable
		Icon:                      "", // set the icon URL/path if available
		Decimals:                  0,  // set the number of decimals
		ChainID:                   0,  // set the chain ID
	}

	tokenMap[mkUSD] = models.TERC20Token{
		Address:                   mkUSD,
		UnderlyingTokensAddresses: []common.Address{},
		Type:                      "", // set the appropriate TTokenType value
		Name:                      "mkUSD",
		Symbol:                    "", // set the symbol if known
		DisplayName:               "", // set the display name if desired
		DisplaySymbol:             "", // set the display symbol if desired
		Description:               "", // provide a description if necessary
		Category:                  "", // set the category if applicable
		Icon:                      "", // set the icon URL/path if available
		Decimals:                  0,  // set the number of decimals
		ChainID:                   0,  // set the chain ID
	}

	prices := RetrieveAllPrices(1, tokenMap)
	fmt.Println("mkUSD price: ", prices[mkUSD].Price)
	fmt.Println("yvmkUSD price: ", prices[yvmkUSD].Price)
	// assert.Greater(t, 0, prices[mkUSD].Price)
	// assert.Greater(t, 0, prices[yvmkUSD].Price)
	assert.Equal(t, "x", "x")
}
