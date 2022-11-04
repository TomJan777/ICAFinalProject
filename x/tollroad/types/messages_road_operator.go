package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateRoadOperator = "create_road_operator"
	TypeMsgUpdateRoadOperator = "update_road_operator"
	TypeMsgDeleteRoadOperator = "delete_road_operator"
)

var _ sdk.Msg = &MsgCreateRoadOperator{}

func NewMsgCreateRoadOperator(
	creator string,
	name string,
	token string,
	active bool,

) *MsgCreateRoadOperator {
	return &MsgCreateRoadOperator{
		Creator: creator,
		Name:    name,
		Token:   token,
		Active:  active,
	}
}

func (msg *MsgCreateRoadOperator) Route() string {
	return RouterKey
}

func (msg *MsgCreateRoadOperator) Type() string {
	return TypeMsgCreateRoadOperator
}

func (msg *MsgCreateRoadOperator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateRoadOperator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateRoadOperator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateRoadOperator{}

func NewMsgUpdateRoadOperator(
	creator string,
	index string,
	name string,
	token string,
	active bool,

) *MsgUpdateRoadOperator {
	return &MsgUpdateRoadOperator{
		Creator: creator,
		Index:   index,
		Name:    name,
		Token:   token,
		Active:  active,
	}
}

func (msg *MsgUpdateRoadOperator) Route() string {
	return RouterKey
}

func (msg *MsgUpdateRoadOperator) Type() string {
	return TypeMsgUpdateRoadOperator
}

func (msg *MsgUpdateRoadOperator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateRoadOperator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateRoadOperator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteRoadOperator{}

func NewMsgDeleteRoadOperator(
	creator string,
	index string,

) *MsgDeleteRoadOperator {
	return &MsgDeleteRoadOperator{
		Creator: creator,
		Index:   index,
	}
}
func (msg *MsgDeleteRoadOperator) Route() string {
	return RouterKey
}

func (msg *MsgDeleteRoadOperator) Type() string {
	return TypeMsgDeleteRoadOperator
}

func (msg *MsgDeleteRoadOperator) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteRoadOperator) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteRoadOperator) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
