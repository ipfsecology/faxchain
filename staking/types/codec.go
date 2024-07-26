package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// RegisterCodec registers concrete types for codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateValidator{}, "filechain/staking/MsgCreateValidator", nil)
	cdc.RegisterConcrete(MsgEditValidator{}, "filechain/staking/MsgEditValidator", nil)
	cdc.RegisterConcrete(MsgDestroyValidator{}, "filechain/staking/MsgDestroyValidator", nil)
	cdc.RegisterConcrete(MsgDeposit{}, "filechain/staking/MsgDeposit", nil)
	cdc.RegisterConcrete(MsgWithdraw{}, "filechain/staking/MsgWithdraw", nil)
	cdc.RegisterConcrete(MsgAddShares{}, "filechain/staking/MsgAddShares", nil)
	cdc.RegisterConcrete(MsgRegProxy{}, "filechain/staking/MsgRegProxy", nil)
	cdc.RegisterConcrete(MsgBindProxy{}, "filechain/staking/MsgBindProxy", nil)
	cdc.RegisterConcrete(MsgUnbindProxy{}, "filechain/staking/MsgUnbindProxy", nil)
}

// ModuleCdc is generic sealed codec to be used throughout this module
var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	RegisterCodec(ModuleCdc)
	codec.RegisterCrypto(ModuleCdc)
	ModuleCdc.Seal()
}
