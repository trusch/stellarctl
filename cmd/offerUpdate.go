// Copyright © 2018 Tino Rusch
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/trusch/stellarctl/transaction"
)

// offerUpdateCmd represents the offerUpdate command
var offerUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update a SDEX offer",
	Long:  `update a SDEX offer`,
	RunE: func(cmd *cobra.Command, args []string) error {
		f := cmd.Flags()
		amount, _ := f.GetString("amount")
		price, _ := f.GetString("price")
		sellingAssetCode, _ := f.GetString("selling-asset-code")
		sellingAssetIssuer, _ := f.GetString("selling-asset-issuer")
		buyingAssetCode, _ := f.GetString("buying-asset-code")
		buyingAssetIssuer, _ := f.GetString("buying-asset-issuer")
		id, _ := f.GetUint64("id")

		op := &transaction.ManageOfferOperation{
			Type:    transaction.ManageOffer,
			OfferID: id,
			Selling: transaction.Asset{
				Code:   sellingAssetCode,
				Issuer: sellingAssetIssuer,
				Native: strings.ToUpper(sellingAssetCode) == transaction.NativeAssetCode,
			},
			Buying: transaction.Asset{
				Code:   buyingAssetCode,
				Issuer: buyingAssetIssuer,
				Native: strings.ToUpper(buyingAssetCode) == transaction.NativeAssetCode,
			},
			Price:  price,
			Amount: amount,
		}

		cli := getClient()

		net := "default"
		if testNet, _ := f.GetBool("testnet"); testNet {
			net = "test"
		}
		seed, _ := f.GetString("seed")
		tx := transaction.Transaction{
			Network:       net,
			SourceAccount: seed,
			Operations:    []transaction.Operation{op},
		}
		tx.SetSequenceProvider(cli)
		if err := tx.Sign(seed); err != nil {
			return err
		}
		xdr, err := tx.ToXDR()
		if err != nil {
			return err
		}
		_, err = cli.SubmitTransaction(xdr)
		return err
	},
}

func init() {
	offerCmd.AddCommand(offerUpdateCmd)
	offerUpdateCmd.Flags().Uint64("id", 0, "offer id")
	offerUpdateCmd.Flags().String("amount", "", "amount of buying asset to buy")
	offerUpdateCmd.Flags().String("price", "", "how many units of selling-asset to spend for one buying-asset")
	offerUpdateCmd.Flags().String("selling-asset-code", "", "selling asset code")
	offerUpdateCmd.Flags().String("selling-asset-issuer", "", "selling asset issuer")
	offerUpdateCmd.Flags().String("buying-asset-code", "", "buying asset code")
	offerUpdateCmd.Flags().String("buying-asset-issuer", "", "buying asset issuer")
	offerUpdateCmd.Flags().String("seed", "", "account seed")
}
