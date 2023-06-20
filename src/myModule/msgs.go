package yourmodule

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Define the CreateKeyValueMsg struct for the transaction message
type CreateKeyValueMsg struct {
	Sender sdk.AccAddress `json:"sender"`
	Key    string         `json:"key"`
	Value  string         `json:"value"`
}

// Implement the Route() method to return the module name
func (msg CreateKeyValueMsg) Route() string {
	return "yourmodule"
}

// Implement the Type() method to return the action name
func (msg CreateKeyValueMsg) Type() string {
	return "create_key_value"
}

// Implement the ValidateBasic() method to validate the transaction message
func (msg CreateKeyValueMsg) ValidateBasic() error {
	if msg.Sender.Empty() {
		return fmt.Errorf("sender address cannot be empty")
	}
	if strings.TrimSpace(msg.Key) == "" {
		return fmt.Errorf("key cannot be empty")
	}
	if strings.TrimSpace(msg.Value) == "" {
		return fmt.Errorf("value cannot be empty")
	}
	return nil
}

// Implement the GetSignBytes() method to get the bytes for signing the message
func (msg CreateKeyValueMsg) GetSignBytes() []byte {
	bz := sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
	return sdk.MustSortJSON(bz)
}

// Implement the GetSigners() method to return the signer address
func (msg CreateKeyValueMsg) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
