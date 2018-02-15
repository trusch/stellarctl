package transaction

import (
	"time"

	"github.com/stellar/go/clients/horizon"
)

const NativeAssetCode = "XLM"

// Transaction represents a stellar transaction
type Transaction struct {
	Network          string       `json:"network,omitempty" yaml:"network,omitempty" mapstructure:"network,emitempty"`
	SourceAccount    string       `json:"source_account,omitempty" yaml:"source_account,omitempty" mapstructure:"source_account,omitempty,emitempty"`
	SequenceID       uint64       `json:"sequence_id,omitempty" yaml:"sequence_id,omitempty" mapstructure:"sequence_id,omitempty,emitempty"`
	Operations       []Operation  `json:"operations,omitempty" yaml:"operations,omitempty" mapstructure:"operations,omitempty,emitempty"`
	Memo             string       `json:"memo,omitempty" yaml:"memo,omitempty" mapstructure:"memo,omitempty,emitempty"`
	NotBefore        time.Time    `json:"not_before,omitempty" yaml:"not_before,omitempty" mapstructure:"not_before,omitempty,emitempty"`
	NotAfter         time.Time    `json:"not_after,omitempty" yaml:"not_after,omitempty" mapstructure:"not_after,omitempty,emitempty"`
	Signatures       []*Signature `json:"signatures,omitempty" yaml:"signatures,omitempty" mapstructure:"signatures,omitempty,emitempty"`
	sequenceProvider *horizon.Client
}

func New(sequenceProvider *horizon.Client) *Transaction {
	return &Transaction{sequenceProvider: sequenceProvider}
}

func (tx *Transaction) SetSequenceProvider(cli *horizon.Client) {
	tx.sequenceProvider = cli
}

const (
	DefaultNetwork = "default"
	TestNetwork    = "test"
)

// Signature contains a signature of a tx
type Signature struct {
	Signature string `json:"signature,omitempty" yaml:"signature,omitempty" mapstructure:"signature,omitempty"`
	Hint      []byte `json:"hint,omitempty" yaml:"hint,omitempty" mapstructure:"hint,omitempty"`
}

// OperationType is the type which represents the operation type constants
type OperationType string

const (
	CreateAccount      OperationType = "create-account"
	Payment            OperationType = "payment"
	PathPayment        OperationType = "path-payment"
	ManageOffer        OperationType = "manage-offer"
	CreatePassiveOffer OperationType = "create-passive-offer"
	SetOptions         OperationType = "set-options"
	ChangeTrust        OperationType = "change-trust"
	AllowTrust         OperationType = "allow-trust"
	AccountMerge       OperationType = "account-merge"
	Inflation          OperationType = "inflation"
	ManageData         OperationType = "manage-data"
)

// Operation defines the common interface of all ops
type Operation interface {
	GetOperationType() OperationType
}

// Asset represents a asset in the network
type Asset struct {
	Code   string `json:"code,omitempty" yaml:"code,omitempty" mapstructure:"code,omitempty"`
	Issuer string `json:"issuer,omitempty" yaml:"issuer,omitempty" mapstructure:"issuer,omitempty"`
	Native bool   `json:"native,omitempty" yaml:"native,omitempty" mapstructure:"native,omitempty"`
}

// CreateAccountOperation creates a account on the ledger and funds it
type CreateAccountOperation struct {
	Type            OperationType `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	Destination     string        `json:"destination,omitempty" yaml:"destination,omitempty" mapstructure:"destination,omitempty"`
	StartingBalance string        `json:"starting_balance,omitempty" yaml:"starting_balance,omitempty" mapstructure:"starting_balance,omitempty"`
}

// GetOperationType implements the operation interface
func (op *CreateAccountOperation) GetOperationType() OperationType {
	return CreateAccount
}

// PaymentOperation sends assets to another account
type PaymentOperation struct {
	Type          OperationType `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	Destination   string        `json:"destination,omitempty" yaml:"destination,omitempty" mapstructure:"destination,omitempty"`
	SourceAccount string        `json:"source_account,omitempty" yaml:"source_account,omitempty" mapstructure:"source_account,omitempty"`
	Asset         Asset         `json:"asset,omitempty" yaml:"asset,omitempty" mapstructure:"asset,omitempty"`
	Amount        string        `json:"amount,omitempty" yaml:"amount,omitempty" mapstructure:"amount,omitempty"`
}

// GetOperationType implements the operation interface
func (op *PaymentOperation) GetOperationType() OperationType {
	return Payment
}

// PathPaymentOperation sends assets to another account
type PathPaymentOperation struct {
	Type              OperationType `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	SourceAccount     string        `json:"source_account,omitempty" yaml:"source_account,omitempty" mapstructure:"source_account,omitempty"`
	SendAsset         Asset         `json:"send_asset,omitempty" yaml:"send_asset,omitempty" mapstructure:"send_asset,omitempty"`
	SendMax           string        `json:"send_max,omitempty" yaml:"send_max,omitempty" mapstructure:"send_max,omitempty"`
	Destination       string        `json:"destination,omitempty" yaml:"destination,omitempty" mapstructure:"destination,omitempty"`
	DestinationAsset  Asset         `json:"destination_asset,omitempty" yaml:"destination_asset,omitempty" mapstructure:"destination_asset,omitempty"`
	DestinationAmount string        `json:"destination_amount,omitempty" yaml:"destination_amount,omitempty" mapstructure:"destination_amount,omitempty"`
	Path              []Asset       `json:"path,omitempty" yaml:"path,omitempty" mapstructure:"path,omitempty"`
}

// GetOperationType implements the operation interface
func (op *PathPaymentOperation) GetOperationType() OperationType {
	return PathPayment
}

// ManageOfferOperation sends assets to another account
type ManageOfferOperation struct {
	Type    OperationType `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	Selling Asset         `json:"selling,omitempty" yaml:"selling,omitempty" mapstructure:"selling,omitempty"`
	Buying  Asset         `json:"buying,omitempty" yaml:"buying,omitempty" mapstructure:"buying,omitempty"`
	Amount  string        `json:"amount,omitempty" yaml:"amount,omitempty" mapstructure:"amount,omitempty"`
	Price   string        `json:"price,omitempty" yaml:"price,omitempty" mapstructure:"price,omitempty"`
	OfferID uint64        `json:"offer_id,omitempty" yaml:"offer_id,omitempty" mapstructure:"offer_id,omitempty"`
}

// GetOperationType implements the operation interface
func (op *ManageOfferOperation) GetOperationType() OperationType {
	return ManageOffer
}

// CreatePassiveOfferOperation sends assets to another account
type CreatePassiveOfferOperation struct {
	Type    OperationType `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	Selling Asset         `json:"selling,omitempty" yaml:"selling,omitempty" mapstructure:"selling,omitempty"`
	Buying  Asset         `json:"buying,omitempty" yaml:"buying,omitempty" mapstructure:"buying,omitempty"`
	Amount  string        `json:"amount,omitempty" yaml:"amount,omitempty" mapstructure:"amount,omitempty"`
	Price   string        `json:"price,omitempty" yaml:"price,omitempty" mapstructure:"price,omitempty"`
}

// GetOperationType implements the operation interface
func (op *CreatePassiveOfferOperation) GetOperationType() OperationType {
	return CreatePassiveOffer
}

// SetOptionsOperation sends assets to another account
type SetOptionsOperation struct {
	Type                 OperationType `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	InflationDestination string        `json:"inflation_destination,omitempty" yaml:"inflation_destination,omitempty" mapstructure:"inflation_destination,omitempty"`
	SetFlags             uint32        `json:"set_flags,omitempty" yaml:"set_flags,omitempty" mapstructure:"set_flags,omitempty"`
	ClearFlags           uint32        `json:"clear_flags,omitempty" yaml:"clear_flags,omitempty" mapstructure:"clear_flags,omitempty"`
	MasterWeight         uint32        `json:"master_weight,omitempty" yaml:"master_weight,omitempty" mapstructure:"master_weight,omitempty"`
	LowThreshold         uint32        `json:"low_threshold,omitempty" yaml:"low_threshold,omitempty" mapstructure:"low_threshold,omitempty"`
	MediumThreshold      uint32        `json:"medium_threshold,omitempty" yaml:"medium_threshold,omitempty" mapstructure:"medium_threshold,omitempty"`
	HighThreshold        uint32        `json:"high_threshold,omitempty" yaml:"high_threshold,omitempty" mapstructure:"high_threshold,omitempty"`
	HomeDomain           string        `json:"home_domain,omitempty" yaml:"home_domain,omitempty" mapstructure:"home_domain,omitempty"`
	Signer               *struct {
		PublicKey string `json:"public_key,omitempty" yaml:"public_key,omitempty" mapstructure:"public_key,omitempty"`
		Weight    uint32 `json:"weight,omitempty" yaml:"weight,omitempty" mapstructure:"weight,omitempty"`
	} `json:"signer,omitempty" yaml:"signer,omitempty" mapstructure:"signer,omitempty"`
}

// GetOperationType implements the operation interface
func (op *SetOptionsOperation) GetOperationType() OperationType {
	return SetOptions
}

// ChangeTrustOperation sends assets to another account
type ChangeTrustOperation struct {
	Type  OperationType `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	Line  Asset         `json:"line,omitempty" yaml:"line,omitempty" mapstructure:"line,omitempty"`
	Limit string        `json:"limit,omitempty" yaml:"limit,omitempty" mapstructure:"limit,omitempty"`
}

// GetOperationType implements the operation interface
func (op *ChangeTrustOperation) GetOperationType() OperationType {
	return ChangeTrust
}

// AllowTrustOperation sends assets to another account
type AllowTrustOperation struct {
	Type      OperationType `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	Trustor   string        `json:"trustor,omitempty" yaml:"trustor,omitempty" mapstructure:"trustor,omitempty"`
	Asset     Asset         `json:"asset,omitempty" yaml:"asset,omitempty" mapstructure:"asset,omitempty"`
	Authorize bool          `json:"authorize,omitempty" yaml:"authorize,omitempty" mapstructure:"authorize,omitempty"`
}

// GetOperationType implements the operation interface
func (op *AllowTrustOperation) GetOperationType() OperationType {
	return AllowTrust
}

// AccountMergeOperation sends assets to another account
type AccountMergeOperation struct {
	Type        OperationType `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	Destination string        `json:"destination,omitempty" yaml:"destination,omitempty" mapstructure:"destination,omitempty"`
}

// GetOperationType implements the operation interface
func (op *AccountMergeOperation) GetOperationType() OperationType {
	return AccountMerge
}

// InflationOperation sends assets to another account
type InflationOperation struct {
	Type OperationType `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
}

// GetOperationType implements the operation interface
func (op *InflationOperation) GetOperationType() OperationType {
	return Inflation
}

// ManageDataOperation sends assets to another account
type ManageDataOperation struct {
	Type  OperationType `json:"type,omitempty" yaml:"type,omitempty" mapstructure:"type,omitempty"`
	Name  string        `json:"name,omitempty" yaml:"name,omitempty" mapstructure:"name,omitempty"`
	Value string        `json:"value,omitempty" yaml:"value,omitempty" mapstructure:"value,omitempty"`
}

// GetOperationType implements the operation interface
func (op *ManageDataOperation) GetOperationType() OperationType {
	return ManageData
}
