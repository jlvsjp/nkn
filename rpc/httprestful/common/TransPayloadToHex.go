package common

import (
	. "github.com/nknorg/nkn/common"
	"github.com/nknorg/nkn/core/asset"
	. "github.com/nknorg/nkn/core/transaction"
	"github.com/nknorg/nkn/core/transaction/payload"
)

type PayloadInfo interface{}

//implement PayloadInfo define CoinbaseInfo
type CoinbaseInfo struct {
	Nonce  uint64
	Issuer IssuerInfo
}

//implement PayloadInfo define IssueAssetInfo
type IssueAssetInfo struct {
}

type IssuerInfo struct {
	X, Y string
}

//implement PayloadInfo define RegisterAssetInfo
type RegisterAssetInfo struct {
	Asset      *asset.Asset
	Amount     Fixed64
	Issuer     IssuerInfo
	Controller string
}

//implement PayloadInfo define TransferAssetInfo
type TransferAssetInfo struct {
}

//implement PayloadInfo define CoinbaseInfo
type PrepaidInfo struct {
	Amount string
	Rates  string
}

type BookkeeperInfo struct {
	PubKey     string
	Action     string
	Issuer     IssuerInfo
	Controller string
}

func TransPayloadToHex(p Payload) PayloadInfo {
	switch object := p.(type) {
	case *payload.Coinbase:
	case *payload.BookKeeper:
		obj := new(BookkeeperInfo)
		encodedPubKey, _ := object.PubKey.EncodePoint(true)
		obj.PubKey = BytesToHexString(encodedPubKey)
		if object.Action == payload.BookKeeperAction_ADD {
			obj.Action = "add"
		} else if object.Action == payload.BookKeeperAction_SUB {
			obj.Action = "sub"
		} else {
			obj.Action = "nil"
		}
		obj.Issuer.X = object.Issuer.X.String()
		obj.Issuer.Y = object.Issuer.Y.String()

		return obj
	case *payload.IssueAsset:
	case *payload.TransferAsset:
	case *payload.Prepaid:
		obj := new(PrepaidInfo)
		obj.Amount = object.Amount.String()
		obj.Rates = object.Rates.String()
		return obj
	case *payload.RegisterAsset:
		obj := new(RegisterAssetInfo)
		obj.Asset = object.Asset
		obj.Amount = object.Amount
		obj.Issuer.X = object.Issuer.X.String()
		obj.Issuer.Y = object.Issuer.Y.String()
		obj.Controller = BytesToHexString(object.Controller.ToArray())
		return obj
	}
	return nil
}
