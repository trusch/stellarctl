package transaction

import (
	"fmt"
	"time"

	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// FromYAML parses a transaction from a YAML encoded form
func FromYAML(data []byte) (*Transaction, error) {
	doc := make(map[string]interface{})
	if err := yaml.Unmarshal(data, &doc); err != nil {
		return nil, err
	}
	tx := &Transaction{}

	if network, ok := doc["network"].(string); ok {
		tx.Network = network
	} else {
		tx.Network = DefaultNetwork
	}

	if source, ok := doc["source_account"].(string); ok {
		tx.SourceAccount = source
	} else {
		return nil, errors.New("need toplevel 'source_account' as string")
	}

	if seq, ok := doc["sequence_id"].(int); ok {
		tx.SequenceID = uint64(seq)
	} else {
		return nil, errors.New("need toplevel 'sequence_id' as string")
	}

	if memo, ok := doc["memo"].(string); ok {
		tx.Memo = memo
	}

	if notBeforeStr, ok := doc["not_before"].(string); ok {
		if t, err := time.Parse(time.RFC3339, notBeforeStr); err == nil {
			tx.NotBefore = t
		} else {
			return nil, errors.Wrap(err, "need 'not_before' to be valid RFC3339 timestamp")
		}
	}

	if notAfterStr, ok := doc["not_after"].(string); ok {
		if t, err := time.Parse(time.RFC3339, notAfterStr); err == nil {
			tx.NotAfter = t
		} else {
			return nil, errors.Wrap(err, "need 'not_after' to be valid RFC3339 timestamp")
		}
	}

	if signatureSlice, ok := doc["signatures"].([]interface{}); ok {
		for _, sig := range signatureSlice {
			if signature, ok := sig.(map[interface{}]interface{}); ok {
				s := &Signature{}
				if err := mapstructure.Decode(signature, s); err != nil {
					return nil, errors.Wrap(err, "can not parse signature")
				}
				tx.Signatures = append(tx.Signatures, s)
			} else {
				return nil, fmt.Errorf("supplied signature not valid (got type %T)", sig)
			}
		}
	}

	if operationSlice, ok := doc["operations"].([]interface{}); ok {
		for _, opInterface := range operationSlice {
			if operationDoc, ok := opInterface.(map[interface{}]interface{}); ok {
				opType, ok := operationDoc["type"].(string)
				if !ok {
					return nil, fmt.Errorf("invalid operation type: %v", operationDoc["type"])
				}
				switch OperationType(opType) {
				case CreateAccount:
					{
						operation := &CreateAccountOperation{}
						if err := mapstructure.Decode(operationDoc, operation); err == nil {
							tx.Operations = append(tx.Operations, operation)
						} else {
							return nil, errors.Wrap(err, "can not parse operation")
						}
					}
				case Payment:
					{
						operation := &PaymentOperation{}
						if err := mapstructure.Decode(operationDoc, operation); err == nil {
							tx.Operations = append(tx.Operations, operation)
						} else {
							return nil, errors.Wrap(err, "can not parse operation")
						}
					}
				case PathPayment:
					{
						operation := &PathPaymentOperation{}
						if err := mapstructure.Decode(operationDoc, operation); err == nil {
							tx.Operations = append(tx.Operations, operation)
						} else {
							return nil, errors.Wrap(err, "can not parse operation")
						}
					}
				case ManageOffer:
					{
						operation := &ManageOfferOperation{}
						if err := mapstructure.Decode(operationDoc, operation); err == nil {
							tx.Operations = append(tx.Operations, operation)
						} else {
							return nil, errors.Wrap(err, "can not parse operation")
						}
					}
				case CreatePassiveOffer:
					{
						operation := &CreatePassiveOfferOperation{}
						if err := mapstructure.Decode(operationDoc, operation); err == nil {
							tx.Operations = append(tx.Operations, operation)
						} else {
							return nil, errors.Wrap(err, "can not parse operation")
						}
					}
				case SetOptions:
					{
						operation := &SetOptionsOperation{}
						if err := mapstructure.Decode(operationDoc, operation); err == nil {
							tx.Operations = append(tx.Operations, operation)
						} else {
							return nil, errors.Wrap(err, "can not parse operation")
						}
					}
				case ChangeTrust:
					{
						operation := &ChangeTrustOperation{}
						if err := mapstructure.Decode(operationDoc, operation); err == nil {
							tx.Operations = append(tx.Operations, operation)
						} else {
							return nil, errors.Wrap(err, "can not parse operation")
						}
					}
				case AllowTrust:
					{
						operation := &AllowTrustOperation{}
						if err := mapstructure.Decode(operationDoc, operation); err == nil {
							tx.Operations = append(tx.Operations, operation)
						} else {
							return nil, errors.Wrap(err, "can not parse operation")
						}
					}
				case AccountMerge:
					{
						operation := &AccountMergeOperation{}
						if err := mapstructure.Decode(operationDoc, operation); err == nil {
							tx.Operations = append(tx.Operations, operation)
						} else {
							return nil, errors.Wrap(err, "can not parse operation")
						}
					}
				case Inflation:
					{
						operation := &InflationOperation{}
						if err := mapstructure.Decode(operationDoc, operation); err == nil {
							tx.Operations = append(tx.Operations, operation)
						} else {
							return nil, errors.Wrap(err, "can not parse operation")
						}
					}
				case ManageData:
					{
						operation := &ManageDataOperation{}
						if err := mapstructure.Decode(operationDoc, operation); err == nil {
							tx.Operations = append(tx.Operations, operation)
						} else {
							return nil, errors.Wrap(err, "can not parse operation")
						}
					}
				}
			} else {
				return nil, fmt.Errorf("supplied operation is not map[string]interface{} but a %T", opInterface)
			}
		}
	}
	return tx, nil
}
