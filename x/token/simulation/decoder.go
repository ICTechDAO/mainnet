package simulation

// DONTCOVER

import (
	"bytes"
	"fmt"

	gogotypes "github.com/gogo/protobuf/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/kv"

	"github.com/gauss/gauss/v6/x/token/types"
)

// NewDecodeStore unmarshals the KVPair's Value to the corresponding token type
func NewDecodeStore(cdc codec.BinaryCodec) func(kvA, kvB kv.Pair) string {
	return func(kvA, kvB kv.Pair) string {
		switch {
		case bytes.Equal(kvA.Key[:1], types.SymbolPrefix):
			var tokenA, tokenB types.Token
			cdc.MustUnmarshal(kvA.Value, &tokenA)
			cdc.MustUnmarshal(kvB.Value, &tokenB)
			return fmt.Sprintf("%v\n%v", tokenA, tokenB)
		case bytes.Equal(kvA.Key[:1], types.OwnerSymbolKey):
			var symbolA, symbolB gogotypes.Value
			cdc.MustUnmarshal(kvA.Value, &symbolA)
			cdc.MustUnmarshal(kvB.Value, &symbolB)
			return fmt.Sprintf("%v\n%v", symbolA, symbolB)
		default:
			panic(fmt.Sprintf("invalid %s key prefix %X", types.ModuleName, kvA.Key[:1]))
		}
	}
}
