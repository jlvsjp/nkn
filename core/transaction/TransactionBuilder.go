package transaction

import (
	"nkn-core/common"
	"nkn-core/core/asset"
	"nkn-core/core/code"
	"nkn-core/core/contract/program"
	"nkn-core/core/transaction/payload"
	"nkn-core/crypto"
	"nkn-core/smartcontract/types"
)

//initial a new transaction with asset registration payload
func NewRegisterAssetTransaction(asset *asset.Asset, amount common.Fixed64, issuer *crypto.PubKey, conroller common.Uint160) (*Transaction, error) {

	//TODO: check arguments

	assetRegPayload := &payload.RegisterAsset{
		Asset:  asset,
		Amount: amount,
		//Precision: precision,
		Issuer:     issuer,
		Controller: conroller,
	}

	return &Transaction{
		//nonce uint64 //TODO: genenrate nonce
		UTXOInputs:    []*UTXOTxInput{},
		Attributes:    []*TxAttribute{},
		TxType:        RegisterAsset,
		Payload:       assetRegPayload,
		Programs:      []*program.Program{},
	}, nil
}

//initial a new transaction with asset registration payload
func NewBookKeeperTransaction(pubKey *crypto.PubKey, isAdd bool, cert []byte, issuer *crypto.PubKey) (*Transaction, error) {

	bookKeeperPayload := &payload.BookKeeper{
		PubKey: pubKey,
		Action: payload.BookKeeperAction_SUB,
		Cert:   cert,
		Issuer: issuer,
	}

	if isAdd {
		bookKeeperPayload.Action = payload.BookKeeperAction_ADD
	}

	return &Transaction{
		TxType:        BookKeeper,
		Payload:       bookKeeperPayload,
		UTXOInputs:    []*UTXOTxInput{},
		Attributes:    []*TxAttribute{},
		Programs:      []*program.Program{},
	}, nil
}

func NewIssueAssetTransaction(outputs []*TxOutput) (*Transaction, error) {

	assetRegPayload := &payload.IssueAsset{}

	return &Transaction{
		TxType:        IssueAsset,
		Payload:       assetRegPayload,
		Attributes:    []*TxAttribute{},
		Outputs:       outputs,
		Programs:      []*program.Program{},
	}, nil
}

func NewTransferAssetTransaction(inputs []*UTXOTxInput, outputs []*TxOutput) (*Transaction, error) {

	//TODO: check arguments

	assetRegPayload := &payload.TransferAsset{}

	return &Transaction{
		TxType:        TransferAsset,
		Payload:       assetRegPayload,
		Attributes:    []*TxAttribute{},
		UTXOInputs:    inputs,
		Outputs:       outputs,
		Programs:      []*program.Program{},
	}, nil
}

//initial a new transaction with publish payload
func NewDeployTransaction(fc *code.FunctionCode, programHash common.Uint160, name, codeversion, author, email, desp string, language types.LangType) (*Transaction, error) {
	//TODO: check arguments
	DeployCodePayload := &payload.DeployCode{
		Code:        fc,
		Name:        name,
		CodeVersion: codeversion,
		Author:      author,
		Email:       email,
		Description: desp,
		Language:    language,
		ProgramHash: programHash,
	}

	return &Transaction{
		TxType:        DeployCode,
		Payload:       DeployCodePayload,
		Attributes:    []*TxAttribute{},
		UTXOInputs:    []*UTXOTxInput{},
		Programs:      []*program.Program{},
	}, nil
}

//initial a new transaction with invoke payload
func NewInvokeTransaction(fc []byte, codeHash common.Uint160, programhash common.Uint160) (*Transaction, error) {
	//TODO: check arguments
	InvokeCodePayload := &payload.InvokeCode{
		Code:        fc,
		CodeHash:    codeHash,
		ProgramHash: programhash,
	}

	return &Transaction{
		TxType:        InvokeCode,
		Payload:       InvokeCodePayload,
		Attributes:    []*TxAttribute{},
		UTXOInputs:    []*UTXOTxInput{},
		Programs:      []*program.Program{},
	}, nil
}
