package yourmodule

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type KVPair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Implement ValidateBasic() function to validate the key-value pair
func (kvp KVPair) ValidateBasic() error {
	if strings.TrimSpace(kvp.Key) == "" {
		return fmt.Errorf("key cannot be empty")
	}
	if strings.TrimSpace(kvp.Value) == "" {
		return fmt.Errorf("value cannot be empty")
	}
	return nil
}

// Implement String() function to display the key-value pair
func (kvp KVPair) String() string {
	return fmt.Sprintf("Key: %s, Value: %s", kvp.Key, kvp.Value)
}
