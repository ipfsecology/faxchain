package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgTokenIssue{}, "filechain/token/MsgIssue", nil)
	cdc.RegisterConcrete(MsgTokenBurn{}, "filechain/token/MsgBurn", nil)
	cdc.RegisterConcrete(MsgTokenMint{}, "filechain/token/MsgMint", nil)
	cdc.RegisterConcrete(MsgMultiSend{}, "filechain/token/MsgMultiTransfer", nil)
	cdc.RegisterConcrete(MsgSend{}, "filechain/token/MsgTransfer", nil)
	cdc.RegisterConcrete(MsgTransferOwnership{}, "filechain/token/MsgTransferOwnership", nil)
	cdc.RegisterConcrete(MsgConfirmOwnership{}, "filechain/token/MsgConfirmOwnership", nil)
	cdc.RegisterConcrete(MsgTokenModify{}, "filechain/token/MsgModify", nil)

	// for test
	//cdc.RegisterConcrete(MsgTokenDestroy{}, "filechain/token/MsgDestroy", nil)
}

// generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
