package types

import (
	"testing"

	"github.com/b9lab/toll-road/testutil/sample"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateRoadOperator_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgCreateRoadOperator
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgCreateRoadOperator{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgCreateRoadOperator{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgUpdateRoadOperator_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgUpdateRoadOperator
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgUpdateRoadOperator{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgUpdateRoadOperator{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}

func TestMsgDeleteRoadOperator_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgDeleteRoadOperator
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgDeleteRoadOperator{
				Creator: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgDeleteRoadOperator{
				Creator: sample.AccAddress(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.msg.ValidateBasic()
			if tt.err != nil {
				require.ErrorIs(t, err, tt.err)
				return
			}
			require.NoError(t, err)
		})
	}
}
