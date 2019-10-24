package connection

// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/03-connection/keeper
// ALIASGEN: github.com/cosmos/cosmos-sdk/x/ibc/03-connection/types

import (
	"github.com/cosmos/cosmos-sdk/x/ibc/03-connection/keeper"
	"github.com/cosmos/cosmos-sdk/x/ibc/03-connection/types"
)

const (
	NONE                              = types.NONE
	INIT                              = types.INIT
	TRYOPEN                           = types.TRYOPEN
	OPEN                              = types.OPEN
	DefaultCodespace                  = types.DefaultCodespace
	CodeConnectionExists              = types.CodeConnectionExists
	CodeConnectionNotFound            = types.CodeConnectionNotFound
	CodeClientConnectionPathsNotFound = types.CodeClientConnectionPathsNotFound
	CodeConnectionPath                = types.CodeConnectionPath
	CodeInvalidCounterpartyConnection = types.CodeInvalidCounterpartyConnection
	AttributeKeyConnectionID          = types.AttributeKeyConnectionID
	AttributeKeyCounterpartyClientID  = types.AttributeKeyCounterpartyClientID
	SubModuleName                     = types.SubModuleName
	StoreKey                          = types.StoreKey
	RouterKey                         = types.RouterKey
	QuerierRoute                      = types.QuerierRoute
	QueryConnection                   = types.QueryConnection
	QueryClientConnections            = types.QueryClientConnections
)

var (
	// functions aliases
	NewKeeper                        = keeper.NewKeeper
	QuerierConnection                = keeper.QuerierConnection
	QuerierClientConnections         = keeper.QuerierClientConnections
	RegisterCodec                    = types.RegisterCodec
	SetMsgConnectionCodec            = types.SetMsgConnectionCodec
	NewConnectionEnd                 = types.NewConnectionEnd
	NewCounterparty                  = types.NewCounterparty
	ErrConnectionExists              = types.ErrConnectionExists
	ErrConnectionNotFound            = types.ErrConnectionNotFound
	ErrClientConnectionPathsNotFound = types.ErrClientConnectionPathsNotFound
	ErrConnectionPath                = types.ErrConnectionPath
	ErrInvalidCounterpartyConnection = types.ErrInvalidCounterpartyConnection
	ErrInvalidConnectionState        = types.ErrInvalidConnectionState
	ConnectionPath                   = types.ConnectionPath
	ClientConnectionsPath            = types.ClientConnectionsPath
	KeyConnection                    = types.KeyConnection
	KeyClientConnections             = types.KeyClientConnections
	NewMsgConnectionOpenInit         = types.NewMsgConnectionOpenInit
	NewMsgConnectionOpenTry          = types.NewMsgConnectionOpenTry
	NewMsgConnectionOpenAck          = types.NewMsgConnectionOpenAck
	NewMsgConnectionOpenConfirm      = types.NewMsgConnectionOpenConfirm
	NewQueryConnectionParams         = types.NewQueryConnectionParams
	NewQueryClientConnectionsParams  = types.NewQueryClientConnectionsParams

	// variable aliases
	SubModuleCdc                   = types.SubModuleCdc
	EventTypeConnectionOpenInit    = types.EventTypeConnectionOpenInit
	EventTypeConnectionOpenTry     = types.EventTypeConnectionOpenTry
	EventTypeConnectionOpenAck     = types.EventTypeConnectionOpenAck
	EventTypeConnectionOpenConfirm = types.EventTypeConnectionOpenConfirm
	AttributeValueCategory         = types.AttributeValueCategory
)

type (
	Keeper                       = keeper.Keeper
	State                        = types.State
	End                          = types.ConnectionEnd
	Counterparty                 = types.Counterparty
	ClientKeeper                 = types.ClientKeeper
	MsgConnectionOpenInit        = types.MsgConnectionOpenInit
	MsgConnectionOpenTry         = types.MsgConnectionOpenTry
	MsgConnectionOpenAck         = types.MsgConnectionOpenAck
	MsgConnectionOpenConfirm     = types.MsgConnectionOpenConfirm
	QueryConnectionParams        = types.QueryConnectionParams
	QueryClientConnectionsParams = types.QueryClientConnectionsParams
)
