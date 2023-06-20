package yourmodule

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Define the QueryResult struct for query response
type QueryResult struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Implement the QueryKeyValue function to retrieve the value of a given key
func QueryKeyValue(ctx sdk.Context, req abci.RequestQuery, keeper Keeper) ([]byte, error) {
	var params struct {
		Key string `json:"key"`
	}

	err := keeper.cdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(params.Key) == "" {
		return nil, fmt.Errorf("key cannot be empty")
	}

	value, found := keeper.GetKeyValue(ctx, params.Key)
	if !found {
		return nil, fmt.Errorf("key does not exist")
	}

	res := QueryResult{
		Key:   params.Key,
		Value: string(value),
	}

	bz, err := codec.MarshalJSONIndent(keeper.cdc, res)
	if err != nil {
		return nil, err
	}

	return bz, nil
}
