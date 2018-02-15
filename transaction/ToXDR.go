package transaction

import (
	"encoding/base64"

	"github.com/stellar/go/build"
	"github.com/stellar/go/xdr"
)

// ToTransactionBuilder returns the transaction builder for this tx
func (tx *Transaction) ToTransactionBuilder() (*build.TransactionBuilder, error) {
	mutators := []build.TransactionMutator{}
	mutators = append(mutators, build.SourceAccount{AddressOrSeed: tx.SourceAccount})
	if tx.SequenceID > 0 {
		mutators = append(mutators, build.Sequence{Sequence: tx.SequenceID})
	} else {
		mutators = append(mutators, build.AutoSequence{SequenceProvider: tx.sequenceProvider})
	}
	if tx.Network == DefaultNetwork {
		mutators = append(mutators, build.DefaultNetwork)
	} else if tx.Network == TestNetwork {
		mutators = append(mutators, build.TestNetwork)
	} else {
		mutators = append(mutators, build.Network{Passphrase: tx.Network})
	}
	if tx.Memo != "" {
		mutators = append(mutators, build.MemoText{Value: tx.Memo})
	}
	for _, op := range tx.Operations {
		switch op.GetOperationType() {
		case CreateAccount:
			{
				createOp := op.(*CreateAccountOperation)
				createAccountMutators := []interface{}{
					build.Destination{AddressOrSeed: createOp.Destination},
					build.NativeAmount{Amount: createOp.StartingBalance},
				}
				mutators = append(mutators, build.CreateAccount(createAccountMutators...))
			}
		case Payment:
			{
				payOp := op.(*PaymentOperation)
				paymentMutators := make([]interface{}, 0)
				if payOp.Asset.Native {
					paymentMutators = append(paymentMutators,
						build.Destination{AddressOrSeed: payOp.Destination},
						build.NativeAmount{Amount: payOp.Amount},
					)
				} else {
					paymentMutators = append(paymentMutators,
						build.Destination{AddressOrSeed: payOp.Destination},
						build.CreditAmount{
							Code:   payOp.Asset.Code,
							Issuer: payOp.Asset.Issuer,
							Amount: payOp.Amount,
						},
					)
				}
				if payOp.SourceAccount != "" {
					paymentMutators = append(paymentMutators, build.SourceAccount{AddressOrSeed: payOp.SourceAccount})
				}
				mutators = append(mutators, build.Payment(paymentMutators...))
			}
		case PathPayment:
			{
				payOp := op.(*PathPaymentOperation)
				path := make([]build.Asset, len(payOp.Path))
				for idx, asset := range payOp.Path {
					path[idx] = build.Asset{
						Code:   asset.Code,
						Issuer: asset.Issuer,
						Native: asset.Native,
					}
				}
				var amount interface{}
				if payOp.SendAsset.Native {
					amount = build.NativeAmount{Amount: payOp.DestinationAmount}
				} else {
					amount = build.CreditAmount{
						Code:   payOp.DestinationAsset.Code,
						Issuer: payOp.DestinationAsset.Issuer,
						Amount: payOp.DestinationAmount,
					}
				}
				payMutators := []interface{}{
					build.Destination{AddressOrSeed: payOp.Destination},
					amount,
					build.PayWithPath{
						Asset: build.Asset{
							Code:   payOp.SendAsset.Code,
							Issuer: payOp.SendAsset.Issuer,
							Native: payOp.SendAsset.Native,
						},
						MaxAmount: payOp.SendMax,
						Path:      path,
					},
				}
				if payOp.SourceAccount != "" {
					payMutators = append(payMutators, build.SourceAccount{AddressOrSeed: payOp.SourceAccount})
				}
				mutators = append(mutators, build.Payment(payMutators...))
			}
		case ManageOffer:
			{
				manageOfferOp := op.(*ManageOfferOperation)
				rate := build.Rate{
					Selling: build.Asset{
						Code:   manageOfferOp.Selling.Code,
						Issuer: manageOfferOp.Selling.Issuer,
						Native: manageOfferOp.Selling.Native,
					},
					Buying: build.Asset{
						Code:   manageOfferOp.Buying.Code,
						Issuer: manageOfferOp.Buying.Issuer,
						Native: manageOfferOp.Buying.Native,
					},
					Price: build.Price(manageOfferOp.Price),
				}
				if manageOfferOp.Amount == "0" {
					mutators = append(mutators, build.DeleteOffer(rate, build.OfferID(manageOfferOp.OfferID)))
				} else if manageOfferOp.OfferID > 0 {
					mutators = append(mutators, build.UpdateOffer(rate, build.Amount(manageOfferOp.Amount), build.OfferID(manageOfferOp.OfferID)))
				} else {
					mutators = append(mutators, build.CreateOffer(rate, build.Amount(manageOfferOp.Amount)))
				}
			}
		case CreatePassiveOffer:
			{
				createOp := op.(*CreatePassiveOfferOperation)
				rate := build.Rate{
					Selling: build.Asset{
						Code:   createOp.Selling.Code,
						Issuer: createOp.Selling.Issuer,
						Native: createOp.Selling.Native,
					},
					Buying: build.Asset{
						Code:   createOp.Buying.Code,
						Issuer: createOp.Buying.Issuer,
						Native: createOp.Buying.Native,
					},
					Price: build.Price(createOp.Price),
				}
				mutators = append(mutators, build.CreatePassiveOffer(rate, build.Amount(createOp.Amount)))
			}
		case SetOptions:
			{
				setOptionsOp := op.(*SetOptionsOperation)
				setOptionMutators := make([]interface{}, 0)
				if setOptionsOp.InflationDestination != "" {
					setOptionMutators = append(setOptionMutators, build.InflationDest(setOptionsOp.InflationDestination))
				}
				if setOptionsOp.SetFlags > 0 {
					setOptionMutators = append(setOptionMutators, build.SetFlag(setOptionsOp.SetFlags))
				}
				if setOptionsOp.ClearFlags > 0 {
					setOptionMutators = append(setOptionMutators, build.ClearFlag(setOptionsOp.ClearFlags))
				}
				if setOptionsOp.MasterWeight > 0 {
					setOptionMutators = append(setOptionMutators, build.MasterWeight(setOptionsOp.MasterWeight))
				}
				if setOptionsOp.LowThreshold > 0 || setOptionsOp.MediumThreshold > 0 || setOptionsOp.HighThreshold > 0 {
					setOptionMutators = append(setOptionMutators, build.SetThresholds(setOptionsOp.LowThreshold, setOptionsOp.MediumThreshold, setOptionsOp.HighThreshold))
				}
				if setOptionsOp.HomeDomain != "" {
					setOptionMutators = append(setOptionMutators, build.HomeDomain(setOptionsOp.HomeDomain))
				}
				if setOptionsOp.Signer != nil {
					setOptionMutators = append(setOptionMutators, build.Signer{
						Address: setOptionsOp.Signer.PublicKey,
						Weight:  setOptionsOp.Signer.Weight,
					})
				}
				mutators = append(mutators, build.SetOptions(setOptionMutators...))
			}
		case ChangeTrust:
			{
				changeTrustOp := op.(*ChangeTrustOperation)
				args := []interface{}{}
				if changeTrustOp.Limit != "" {
					args = []interface{}{build.Limit(changeTrustOp.Limit)}
				}
				mutators = append(mutators, build.Trust(changeTrustOp.Line.Code, changeTrustOp.Line.Issuer, args...))
			}
		case AllowTrust:
			{
				allowTrustOp := op.(*AllowTrustOperation)
				mut := build.AllowTrust(
					build.Trustor{Address: allowTrustOp.Trustor},
					build.AllowTrustAsset{Code: allowTrustOp.Asset.Code},
					build.Authorize{Value: allowTrustOp.Authorize},
				)
				mutators = append(mutators, mut)
			}
		case AccountMerge:
			{
				accountMergeOp := op.(*AccountMergeOperation)
				mutators = append(mutators, build.AccountMerge(build.Destination{AddressOrSeed: accountMergeOp.Destination}))
			}
		case Inflation:
			{
				mutators = append(mutators, build.Inflation())
			}
		case ManageData:
			{
				manageDataOp := op.(*ManageDataOperation)
				if manageDataOp.Value == "" {
					mutators = append(mutators, build.ClearData(manageDataOp.Name))
				} else {
					mutators = append(mutators, build.SetData(manageDataOp.Name, []byte(manageDataOp.Value)))
				}
			}
		}
	}
	transaction, err := build.Transaction(mutators...)
	if err != nil {
		return nil, err
	}
	if !tx.NotBefore.IsZero() || !tx.NotAfter.IsZero() {
		transaction.TX.TimeBounds = &xdr.TimeBounds{
			MinTime: xdr.Uint64(tx.NotBefore.Unix()),
			MaxTime: xdr.Uint64(tx.NotAfter.Unix()),
		}
	}
	return transaction, nil
}

func (tx *Transaction) Sign(seed string) error {
	transaction, err := tx.ToTransactionBuilder()
	if err != nil {
		return err
	}
	envelop, err := transaction.Sign(seed)
	if err != nil {
		return err
	}
	for _, sig := range envelop.E.Signatures {
		txSig := &Signature{
			Hint:      sig.Hint[:4],
			Signature: base64.StdEncoding.EncodeToString(sig.Signature),
		}
		tx.Signatures = append(tx.Signatures, txSig)
	}
	return nil
}

func (tx *Transaction) ToXDR() (string, error) {
	txBuilder, err := tx.ToTransactionBuilder()
	if err != nil {
		return "", err
	}
	env, err := txBuilder.Sign()
	if err != nil {
		return "", err
	}
	for _, sigBase64 := range tx.Signatures {
		sig, err := base64.StdEncoding.DecodeString(sigBase64.Signature)
		if err != nil {
			return "", err
		}
		hint := sigBase64.Hint
		env.E.Signatures = append(env.E.Signatures, xdr.DecoratedSignature{
			Signature: sig,
			Hint:      [4]byte{hint[0], hint[1], hint[2], hint[3]},
		})
	}
	return env.Base64()
}
