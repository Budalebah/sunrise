package encoding

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/std"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
)

type ModuleRegister interface {
	RegisterInterfaces(codectypes.InterfaceRegistry)
}

// Config specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type Config struct {
	InterfaceRegistry codectypes.InterfaceRegistry
	Codec             codec.Codec
	TxConfig          client.TxConfig
}

// MakeConfig creates an encoding config for the app.
func MakeConfig(regs ...ModuleRegister) Config {
	// create the codec
	interfaceRegistry := codectypes.NewInterfaceRegistry()

	// register the standard types from the sdk
	std.RegisterInterfaces(interfaceRegistry)

	// register specific modules
	for _, reg := range regs {
		reg.RegisterInterfaces(interfaceRegistry)
	}

	// create the final configuration
	cdc := codec.NewProtoCodec(interfaceRegistry)
	dec := tx.DefaultTxDecoder(cdc)
	dec = indexWrapperDecoder(dec)

	txCfg, _ := tx.NewTxConfigWithOptions(cdc, tx.ConfigOptions{})

	return Config{
		InterfaceRegistry: interfaceRegistry,
		Codec:             cdc,
		TxConfig:          txCfg,
	}
}
