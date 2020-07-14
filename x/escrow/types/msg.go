package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = MsgSendWithUnlock{}

type MsgSendWithUnlock struct {
	FromAddress sdk.AccAddress
	ToAddress   sdk.AccAddress
	Amount      sdk.Coins
	UnlockTime  time.Time
}

func NewMsgSendWithUnlock(fromAddr, toAddr sdk.AccAddress, amount sdk.Coins, unlockT time.Time) MsgSendWithUnlock {
	return MsgSendWithUnlock{FromAddress: fromAddr, ToAddress: toAddr, Amount: amount, UnlockTime: unlockT}
}

// Route Implements Msg.
func (msg MsgSendWithUnlock) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgSendWithUnlock) Type() string { return ModuleName }

// ValidateBasic Implements Msg.
func (msg MsgSendWithUnlock) ValidateBasic() error {
	if msg.FromAddress.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing sender address")
	}
	if msg.ToAddress.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing recipient address")
	}
	if !msg.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}
	if !msg.Amount.IsAllPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSendWithUnlock) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgSendWithUnlock) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.FromAddress}
}

// MsgPayout is a message for placing a bet
type MsgPayout struct {
	Receiver    sdk.AccAddress `json:"receiver"`
	DayId       int64          `json:"dayId"`
}

// NewMsgPayout creates a new MsgPayout instance.
func NewMsgPayout(
	receiver sdk.AccAddress,
	dayId int64,
) MsgPayout {
	return MsgPayout{
		Receiver: receiver,
		DayId:    dayId,
	}
}

// Route implements the sdk.Msg interface for MsgPayout.
func (msg MsgPayout) Route() string { return RouterKey }

// Type implements the sdk.Msg interface for MsgPayout.
func (msg MsgPayout) Type() string { return ModuleName }

// ValidateBasic implements the sdk.Msg interface for MsgPayout.
func (msg MsgPayout) ValidateBasic() error {
	if msg.Receiver.Empty() {
		return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgPayout: Bettor address must not be empty.")
	}
	// if msg.DayId != 0 {
	// 	return sdkerrors.Wrapf(ErrInvalidBasicMsg, "MsgPayout: Unknown DayId.")
	// }
	return nil
}

// GetSigners implements the sdk.Msg interface for MsgPayout.
func (msg MsgPayout) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Receiver}
}

// GetSignBytes implements the sdk.Msg interface for MsgPayout.
func (msg MsgPayout) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}
