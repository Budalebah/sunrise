package types

import (
	"testing"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/stretchr/testify/require"
	"github.com/sunriselayer/sunrise/testutil/sample"
)

func TestMsgPublishData_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  MsgPublishData
		err  error
	}{
		{
			name: "invalid address",
			msg: MsgPublishData{
				Sender: "invalid_address",
			},
			err: sdkerrors.ErrInvalidAddress,
		}, {
			name: "valid address",
			msg: MsgPublishData{
				Sender:            sample.AccAddress(),
				ParityShardCount:  0,
				ShardDoubleHashes: [][]byte{{0x01}},
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
